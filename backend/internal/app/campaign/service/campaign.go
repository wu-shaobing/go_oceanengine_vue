package service

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/config"
	advModel "oceanengine-backend/internal/app/advertiser/model"
	advRepo "oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/internal/app/campaign/dto"
	"oceanengine-backend/internal/app/campaign/model"
	"oceanengine-backend/internal/app/campaign/repository"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// CampaignService 广告系列服务
type CampaignService struct {
	repo     repository.CampaignRepository
	advRepo  advRepo.AdvertiserRepository
	oceanCfg *config.OceanConfig
}

// NewCampaignService 创建广告系列服务
func NewCampaignService(db *gorm.DB, oceanCfg *config.OceanConfig) *CampaignService {
	return &CampaignService{
		repo:     repository.NewCampaignRepository(db),
		advRepo:  advRepo.NewAdvertiserRepository(db),
		oceanCfg: oceanCfg,
	}
}

// GetList 获取广告系列列表
func (s *CampaignService) GetList(ctx context.Context, req *dto.CampaignListReq) ([]*dto.CampaignListResp, int64, error) {
	list, total, err := s.repo.GetList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.CampaignListResp, len(list))
	for i, c := range list {
		result[i] = &dto.CampaignListResp{
			ID:            c.ID,
			AdvertiserID:  c.AdvertiserID,
			CampaignID:    c.CampaignID,
			Name:          c.Name,
			Budget:        c.Budget,
			BudgetMode:    c.BudgetMode,
			MarketingGoal: c.MarketingGoal,
			Status:        c.Status,
			CreatedAt:     c.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetByID 获取广告系列详情
func (s *CampaignService) GetByID(ctx context.Context, id uint64) (*dto.CampaignDetailResp, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrCampaignNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.toDetailResp(c), nil
}

// Create 创建广告系列
func (s *CampaignService) Create(ctx context.Context, req *dto.CampaignCreateReq) (*dto.CampaignDetailResp, error) {
	// 验证广告主
	adv, err := s.advRepo.GetByID(ctx, req.AdvertiserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 调用巨量引擎API创建广告系列
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	campaignService := oceanengine.NewCampaignService(client)
	createReq := &oceanengine.CampaignCreateRequest{
		AdvertiserID: int64(adv.AdvertiserID),
		CampaignName: req.Name,
		Budget:       req.Budget,
		BudgetMode:   req.BudgetMode,
		LandingType:  req.LandingType,
	}

	createResp, err := campaignService.Create(ctx, createReq)
	if err != nil {
		return nil, errcode.WrapWithMessage(errcode.ErrOEAPIFailed, "创建广告系列失败", err)
	}

	// 保存到本地数据库
	campaign := &model.Campaign{
		AdvertiserID:  req.AdvertiserID,
		CampaignID:    uint64(createResp.CampaignID),
		Name:          req.Name,
		Budget:        req.Budget,
		BudgetMode:    req.BudgetMode,
		MarketingGoal: req.MarketingGoal,
		LandingType:   req.LandingType,
		Status:        model.CampaignStatusEnable,
	}

	if err := s.repo.Create(ctx, campaign); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.toDetailResp(campaign), nil
}

// Update 更新广告系列
func (s *CampaignService) Update(ctx context.Context, id uint64, req *dto.CampaignUpdateReq) error {
	campaign, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrCampaignNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 获取广告主信息
	adv, err := s.advRepo.GetByID(ctx, campaign.AdvertiserID)
	if err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 调用巨量引擎API更新
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	campaignService := oceanengine.NewCampaignService(client)
	updateReq := &oceanengine.CampaignUpdateRequest{
		AdvertiserID: int64(adv.AdvertiserID),
		CampaignID:   int64(campaign.CampaignID),
	}

	if req.Name != "" {
		updateReq.CampaignName = req.Name
		campaign.Name = req.Name
	}
	if req.Budget > 0 {
		updateReq.Budget = req.Budget
		campaign.Budget = req.Budget
	}
	if req.BudgetMode != "" {
		updateReq.BudgetMode = req.BudgetMode
		campaign.BudgetMode = req.BudgetMode
	}

	if err := campaignService.Update(ctx, updateReq); err != nil {
		return errcode.WrapWithMessage(errcode.ErrOEAPIFailed, "更新广告系列失败", err)
	}

	// 更新本地数据
	if req.Remark != "" {
		campaign.Remark = req.Remark
	}

	if err := s.repo.Update(ctx, campaign); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdateStatus 批量更新状态
func (s *CampaignService) UpdateStatus(ctx context.Context, req *dto.CampaignStatusUpdateReq) error {
	// 这里可以批量调用API更新状态
	if err := s.repo.UpdateStatus(ctx, req.IDs, req.Status); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// Delete 删除广告系列
func (s *CampaignService) Delete(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrCampaignNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Sync 同步广告系列数据
func (s *CampaignService) Sync(ctx context.Context, advertiserID uint64) (*dto.CampaignSyncResp, error) {
	// 获取广告主信息
	var adv *advModel.Advertiser
	var err error

	adv, err = s.advRepo.GetByID(ctx, advertiserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdvertiserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}

	// 创建SDK客户端
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	// 获取广告系列列表
	campaignService := oceanengine.NewCampaignService(client)
	listReq := &oceanengine.CampaignListRequest{
		AdvertiserID: int64(adv.AdvertiserID),
		PageSize:     100,
	}
	campaignsResp, err := campaignService.GetList(ctx, listReq)
	if err != nil {
		return nil, errcode.WrapWithMessage(errcode.ErrOEAPIFailed, "获取广告系列列表失败", err)
	}

	syncCount := 0
	for _, c := range campaignsResp.List {
		exists, _ := s.repo.ExistsByCampaignID(ctx, uint64(c.ID))
		if exists {
			// 更新
			campaign, _ := s.repo.GetByCampaignID(ctx, uint64(c.ID))
			if campaign != nil {
				campaign.Name = c.Name
				campaign.Budget = c.Budget
				campaign.BudgetMode = c.BudgetMode
				campaign.Status = c.Status
				now := time.Now()
				campaign.LastSyncAt = &now
				_ = s.repo.Update(ctx, campaign)
				syncCount++
			}
		} else {
			// 创建
			campaign := &model.Campaign{
				AdvertiserID: advertiserID,
				CampaignID:   uint64(c.ID),
				Name:         c.Name,
				Budget:       c.Budget,
				BudgetMode:   c.BudgetMode,
				Status:       c.Status,
			}
			now := time.Now()
			campaign.LastSyncAt = &now
			_ = s.repo.Create(ctx, campaign)
			syncCount++
		}
	}

	return &dto.CampaignSyncResp{
		SyncCount: syncCount,
		SyncAt:    time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// toDetailResp 转换为详情响应
func (s *CampaignService) toDetailResp(c *model.Campaign) *dto.CampaignDetailResp {
	var startTime, endTime, lastSyncAt string
	if c.StartTime != nil {
		startTime = c.StartTime.Format("2006-01-02 15:04:05")
	}
	if c.EndTime != nil {
		endTime = c.EndTime.Format("2006-01-02 15:04:05")
	}
	if c.LastSyncAt != nil {
		lastSyncAt = c.LastSyncAt.Format("2006-01-02 15:04:05")
	}

	return &dto.CampaignDetailResp{
		ID:            c.ID,
		AdvertiserID:  c.AdvertiserID,
		CampaignID:    c.CampaignID,
		Name:          c.Name,
		Budget:        c.Budget,
		BudgetMode:    c.BudgetMode,
		MarketingGoal: c.MarketingGoal,
		DeliveryMode:  c.DeliveryMode,
		Status:        c.Status,
		LandingType:   c.LandingType,
		StartTime:     startTime,
		EndTime:       endTime,
		Remark:        c.Remark,
		LastSyncAt:    lastSyncAt,
		CreatedAt:     c.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     c.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
