package service

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/ad/dto"
	"oceanengine-backend/internal/app/ad/model"
	"oceanengine-backend/internal/app/ad/repository"
	advRepo "oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// AdService 广告组服务
type AdService struct {
	repo     repository.AdRepository
	advRepo  advRepo.AdvertiserRepository
	oceanCfg *config.OceanConfig
}

// NewAdService 创建广告组服务
func NewAdService(db *gorm.DB, oceanCfg *config.OceanConfig) *AdService {
	return &AdService{
		repo:     repository.NewAdRepository(db),
		advRepo:  advRepo.NewAdvertiserRepository(db),
		oceanCfg: oceanCfg,
	}
}

// GetList 获取广告组列表
func (s *AdService) GetList(ctx context.Context, req *dto.AdListReq) ([]*dto.AdListResp, int64, error) {
	list, total, err := s.repo.GetList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.AdListResp, len(list))
	for i, ad := range list {
		result[i] = &dto.AdListResp{
			ID:           ad.ID,
			AdvertiserID: ad.AdvertiserID,
			CampaignID:   ad.CampaignID,
			AdID:         ad.AdID,
			Name:         ad.Name,
			OptStatus:    ad.OptStatus,
			Status:       ad.Status,
			Budget:       ad.Budget,
			Bid:          ad.Bid,
			CreatedAt:    ad.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetByID 获取广告组详情
func (s *AdService) GetByID(ctx context.Context, id uint64) (*dto.AdDetailResp, error) {
	ad, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.toDetailResp(ad), nil
}

// Update 更新广告组
func (s *AdService) Update(ctx context.Context, id uint64, req *dto.AdUpdateReq) error {
	ad, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrAdNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if req.Name != "" {
		ad.Name = req.Name
	}
	if req.Budget > 0 {
		ad.Budget = req.Budget
	}
	if req.Bid > 0 {
		ad.Bid = req.Bid
	}
	if req.OptStatus != "" {
		ad.OptStatus = req.OptStatus
	}
	if req.Remark != "" {
		ad.Remark = req.Remark
	}

	if err := s.repo.Update(ctx, ad); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdateStatus 批量更新状态
func (s *AdService) UpdateStatus(ctx context.Context, req *dto.AdStatusUpdateReq) error {
	if err := s.repo.UpdateStatus(ctx, req.IDs, req.OptStatus); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// Create 创建广告组
func (s *AdService) Create(ctx context.Context, req *dto.AdCreateReq) (map[string]interface{}, error) {
	// 验证广告主是否存在并获取 access_token
	adv, err := s.advRepo.GetByID(ctx, req.AdvertiserID)
	if err != nil {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}

	var adID uint64

	// 如果有 access_token，调用 OE API 创建广告
	if adv.AccessToken != "" {
		client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
		client.SetAccessToken(adv.AccessToken)
		adSvc := oceanengine.NewAdService(client)

		oeReq := &oceanengine.AdCreateRequest{
			AdvertiserID:  int64(adv.AdvertiserID),
			CampaignID:    int64(req.CampaignID),
			Name:          req.Name,
			Budget:        req.Budget,
			BudgetMode:    req.BudgetMode,
			DeliveryRange: req.DeliveryRange,
			Pricing:       req.Pricing,
			Bid:           req.Bid,
		}

		oeResp, err := adSvc.Create(ctx, oeReq)
		if err != nil {
			// OE API 调用失败，回退到本地创建
			adID = uint64(time.Now().UnixNano())
		} else {
			adID = uint64(oeResp.AdID)
		}
	} else {
		// 无 token，本地生成唯一 ID
		adID = uint64(time.Now().UnixNano())
	}

	// 创建广告组记录
	ad := &model.Ad{
		AdvertiserID:  req.AdvertiserID,
		CampaignID:    req.CampaignID,
		AdID:          adID,
		Name:          req.Name,
		Budget:        req.Budget,
		BudgetMode:    req.BudgetMode,
		DeliveryRange: req.DeliveryRange,
		Pricing:       req.Pricing,
		Bid:           req.Bid,
		OptStatus:     "ENABLE",
		Status:        "AD_STATUS_DELIVERY_OK",
	}

	if err := s.repo.Create(ctx, ad); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return map[string]interface{}{
		"id":    ad.ID,
		"ad_id": ad.AdID,
	}, nil
}

// Delete 删除广告组
func (s *AdService) Delete(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrAdNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// toDetailResp 转换为详情响应
func (s *AdService) toDetailResp(ad *model.Ad) *dto.AdDetailResp {
	var startTime, endTime, lastSyncAt string
	if ad.StartTime != nil {
		startTime = ad.StartTime.Format("2006-01-02 15:04:05")
	}
	if ad.EndTime != nil {
		endTime = ad.EndTime.Format("2006-01-02 15:04:05")
	}
	if ad.LastSyncAt != nil {
		lastSyncAt = ad.LastSyncAt.Format("2006-01-02 15:04:05")
	}

	return &dto.AdDetailResp{
		ID:             ad.ID,
		AdvertiserID:   ad.AdvertiserID,
		CampaignID:     ad.CampaignID,
		AdID:           ad.AdID,
		Name:           ad.Name,
		OptStatus:      ad.OptStatus,
		Status:         ad.Status,
		Budget:         ad.Budget,
		BudgetMode:     ad.BudgetMode,
		DeliveryRange:  ad.DeliveryRange,
		Pricing:        ad.Pricing,
		Bid:            ad.Bid,
		ConvertID:      ad.ConvertID,
		StartTime:      startTime,
		EndTime:        endTime,
		ScheduleType:   ad.ScheduleType,
		AudienceType:   ad.AudienceType,
		TargetSettings: ad.TargetSettings,
		Remark:         ad.Remark,
		LastSyncAt:     lastSyncAt,
		CreatedAt:      ad.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      ad.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
