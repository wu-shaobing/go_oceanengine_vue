package service

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/advertiser/dto"
	"oceanengine-backend/internal/app/advertiser/model"
	"oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// AdvertiserService 广告主服务
type AdvertiserService struct {
	repo     repository.AdvertiserRepository
	fundRepo repository.FundRepository
	oceanCfg *config.OceanConfig
}

// NewAdvertiserService 创建广告主服务
func NewAdvertiserService(db *gorm.DB, oceanCfg *config.OceanConfig) *AdvertiserService {
	return &AdvertiserService{
		repo:     repository.NewAdvertiserRepository(db),
		fundRepo: repository.NewFundRepository(db),
		oceanCfg: oceanCfg,
	}
}

// GetList 获取广告主列表
func (s *AdvertiserService) GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*dto.AdvertiserListResp, int64, error) {
	list, total, err := s.repo.GetList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.AdvertiserListResp, len(list))
	for i, adv := range list {
		result[i] = &dto.AdvertiserListResp{
			ID:           adv.ID,
			AdvertiserID: adv.AdvertiserID,
			Name:         adv.Name,
			Company:      adv.Company,
			Status:       adv.Status,
			Balance:      adv.Balance,
			Industry:     adv.Industry,
			CreatedAt:    adv.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetByID 获取广告主详情
func (s *AdvertiserService) GetByID(ctx context.Context, id uint64) (*dto.AdvertiserDetailResp, error) {
	adv, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.toDetailResp(adv), nil
}

// Update 更新广告主
func (s *AdvertiserService) Update(ctx context.Context, id uint64, req *dto.AdvertiserUpdateReq) error {
	adv, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 更新本地字段
	if req.ContactName != "" {
		adv.ContactName = req.ContactName
	}
	if req.ContactPhone != "" {
		adv.ContactPhone = req.ContactPhone
	}
	if req.ContactEmail != "" {
		adv.ContactEmail = req.ContactEmail
	}
	if req.Address != "" {
		adv.Address = req.Address
	}
	if req.Remark != "" {
		adv.Remark = req.Remark
	}

	if err := s.repo.Update(ctx, adv); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Delete 删除广告主
func (s *AdvertiserService) Delete(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Sync 同步广告主数据
func (s *AdvertiserService) Sync(ctx context.Context, id uint64) (*dto.AdvertiserSyncResp, error) {
	adv, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 检查 Token 是否有效
	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}

	// 创建 SDK 客户端
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	// 获取广告主信息
	advService := oceanengine.NewAdvertiserService(client)
	infos, err := advService.GetInfo(ctx, []int64{int64(adv.AdvertiserID)})
	if err != nil {
		return nil, errcode.WrapWithMessage(errcode.ErrAdvertiserSyncFailed, "获取广告主信息失败", err)
	}

	syncFields := []string{}
	if len(infos) > 0 {
		info := infos[0]
		if adv.Name != info.Name {
			adv.Name = info.Name
			syncFields = append(syncFields, "name")
		}
		if adv.Company != info.Company {
			adv.Company = info.Company
			syncFields = append(syncFields, "company")
		}
		if adv.Status != info.Status {
			adv.Status = info.Status
			syncFields = append(syncFields, "status")
		}
		if adv.Role != info.Role {
			adv.Role = info.Role
			syncFields = append(syncFields, "role")
		}
		if adv.Address != info.Address {
			adv.Address = info.Address
			syncFields = append(syncFields, "address")
		}
	}

	// 获取资金信息
	funds, err := advService.GetFund(ctx, []int64{int64(adv.AdvertiserID)})
	if err == nil {
		if fund, ok := funds[int64(adv.AdvertiserID)]; ok {
			if adv.Balance != fund.Balance {
				adv.Balance = fund.Balance
				syncFields = append(syncFields, "balance")
			}
			if adv.ValidBalance != fund.ValidBalance {
				adv.ValidBalance = fund.ValidBalance
				syncFields = append(syncFields, "valid_balance")
			}
			if adv.CashBalance != fund.CashBalance {
				adv.CashBalance = fund.CashBalance
				syncFields = append(syncFields, "cash_balance")
			}
		}
	}

	// 更新同步时间
	now := time.Now()
	adv.LastSyncAt = &now

	if err := s.repo.Update(ctx, adv); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.AdvertiserSyncResp{
		AdvertiserID: adv.AdvertiserID,
		SyncFields:   syncFields,
		SyncAt:       now.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetBalance 获取广告主余额
func (s *AdvertiserService) GetBalance(ctx context.Context, id uint64) (*dto.AdvertiserBalanceResp, error) {
	adv, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 检查 Token 是否有效
	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}

	// 创建 SDK 客户端
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	// 获取资金信息
	advService := oceanengine.NewAdvertiserService(client)
	funds, err := advService.GetFund(ctx, []int64{int64(adv.AdvertiserID)})
	if err != nil {
		return nil, errcode.WrapWithMessage(errcode.ErrOEAPIFailed, "获取余额失败", err)
	}

	fund, ok := funds[int64(adv.AdvertiserID)]
	if !ok {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}

	// 更新本地余额
	adv.Balance = fund.Balance
	adv.ValidBalance = fund.ValidBalance
	adv.CashBalance = fund.CashBalance
	_ = s.repo.Update(ctx, adv)

	return &dto.AdvertiserBalanceResp{
		Balance:      fund.Balance,
		ValidBalance: fund.ValidBalance,
		CashBalance:  fund.CashBalance,
		GrantBalance: fund.GrantBalance,
		SyncAt:       time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// GetFundList 获取资金流水列表
func (s *AdvertiserService) GetFundList(ctx context.Context, req *dto.FundListReq) ([]*dto.FundListResp, int64, error) {
	list, total, err := s.fundRepo.GetList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.FundListResp, len(list))
	for i, fund := range list {
		var transactionTime string
		if fund.TransactionTime != nil {
			transactionTime = fund.TransactionTime.Format("2006-01-02 15:04:05")
		}
		result[i] = &dto.FundListResp{
			ID:              fund.ID,
			TransactionType: fund.TransactionType,
			Amount:          fund.Amount,
			BalanceBefore:   fund.BalanceBefore,
			BalanceAfter:    fund.BalanceAfter,
			TransactionSeq:  fund.TransactionSeq,
			TransactionTime: transactionTime,
			Remark:          fund.Remark,
		}
	}

	return result, total, nil
}

// GetOAuthAuthorizeURL 获取 OAuth 授权 URL
func (s *AdvertiserService) GetOAuthAuthorizeURL(state string) string {
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	oauthService := oceanengine.NewOAuthService(client)
	// 使用带scope和material_auth的URL
	return oauthService.GetAuthURLWithScope(state, s.oceanCfg.RedirectURI, nil, s.oceanCfg.MaterialAuth)
}

// HandleOAuthCallback 处理 OAuth 回调
func (s *AdvertiserService) HandleOAuthCallback(ctx context.Context, authCode string) error {
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	oauthService := oceanengine.NewOAuthService(client)

	// 获取 Access Token
	tokenResp, err := oauthService.GetAccessToken(ctx, authCode)
	if err != nil {
		return errcode.WrapWithMessage(errcode.ErrAdvertiserAuthFailed, "获取Token失败", err)
	}

	// 设置 Token
	client.SetAccessToken(tokenResp.AccessToken)

	// 获取广告主信息
	advService := oceanengine.NewAdvertiserService(client)
	for _, advertiserID := range tokenResp.AdvertiserIDs {
		infos, err := advService.GetInfo(ctx, []int64{advertiserID})
		if err != nil {
			continue
		}
		if len(infos) == 0 {
			continue
		}

		info := infos[0]

		// 检查是否已存在
		exists, _ := s.repo.ExistsByAdvertiserID(ctx, uint64(advertiserID))
		if exists {
			// 更新现有广告主
			adv, _ := s.repo.GetByAdvertiserID(ctx, uint64(advertiserID))
			if adv != nil {
				adv.AccessToken = tokenResp.AccessToken
				adv.RefreshToken = tokenResp.RefreshToken
				expireAt := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
				adv.TokenExpireAt = &expireAt
				adv.Name = info.Name
				adv.Company = info.Company
				adv.Status = info.Status
				adv.Role = info.Role
				_ = s.repo.Update(ctx, adv)
			}
		} else {
			// 创建新广告主
			expireAt := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
			adv := &model.Advertiser{
				AdvertiserID:  uint64(advertiserID),
				Name:          info.Name,
				Company:       info.Company,
				Status:        info.Status,
				Role:          info.Role,
				AccessToken:   tokenResp.AccessToken,
				RefreshToken:  tokenResp.RefreshToken,
				TokenExpireAt: &expireAt,
			}
			_ = s.repo.Create(ctx, adv)
		}
	}

	return nil
}

// toDetailResp 转换为详情响应
func (s *AdvertiserService) toDetailResp(adv *model.Advertiser) *dto.AdvertiserDetailResp {
	var lastSyncAt string
	if adv.LastSyncAt != nil {
		lastSyncAt = adv.LastSyncAt.Format("2006-01-02 15:04:05")
	}

	return &dto.AdvertiserDetailResp{
		ID:           adv.ID,
		AdvertiserID: adv.AdvertiserID,
		Name:         adv.Name,
		Company:      adv.Company,
		Status:       adv.Status,
		Role:         adv.Role,
		Balance:      adv.Balance,
		ValidBalance: adv.ValidBalance,
		CashBalance:  adv.CashBalance,
		Industry:     adv.Industry,
		ContactName:  adv.ContactName,
		ContactPhone: adv.ContactPhone,
		ContactEmail: adv.ContactEmail,
		Address:      adv.Address,
		LastSyncAt:   lastSyncAt,
		CreatedAt:    adv.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    adv.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
