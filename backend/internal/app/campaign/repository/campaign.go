package repository

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/campaign/dto"
	"oceanengine-backend/internal/app/campaign/model"
)

// CampaignRepository 广告系列仓储接口
type CampaignRepository interface {
	GetList(ctx context.Context, req *dto.CampaignListReq) ([]*model.Campaign, int64, error)
	GetByID(ctx context.Context, id uint64) (*model.Campaign, error)
	GetByCampaignID(ctx context.Context, campaignID uint64) (*model.Campaign, error)
	GetByAdvertiserID(ctx context.Context, advertiserID uint64) ([]*model.Campaign, error)
	Create(ctx context.Context, campaign *model.Campaign) error
	Update(ctx context.Context, campaign *model.Campaign) error
	UpdateStatus(ctx context.Context, ids []uint64, status string) error
	Delete(ctx context.Context, id uint64) error
	BatchDelete(ctx context.Context, ids []uint64) error
	ExistsByCampaignID(ctx context.Context, campaignID uint64) (bool, error)
}

type campaignRepository struct {
	db *gorm.DB
}

// NewCampaignRepository 创建广告系列仓储
func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{db: db}
}

// GetList 获取广告系列列表
func (r *campaignRepository) GetList(ctx context.Context, req *dto.CampaignListReq) ([]*model.Campaign, int64, error) {
	var list []*model.Campaign
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Campaign{})

	if req.AdvertiserID > 0 {
		query = query.Where("advertiser_id = ?", req.AdvertiserID)
	}
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}
	if req.MarketingGoal != "" {
		query = query.Where("marketing_goal = ?", req.MarketingGoal)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetByID 根据ID获取广告系列
func (r *campaignRepository) GetByID(ctx context.Context, id uint64) (*model.Campaign, error) {
	var campaign model.Campaign
	if err := r.db.WithContext(ctx).First(&campaign, id).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

// GetByCampaignID 根据广告系列ID获取
func (r *campaignRepository) GetByCampaignID(ctx context.Context, campaignID uint64) (*model.Campaign, error) {
	var campaign model.Campaign
	if err := r.db.WithContext(ctx).Where("campaign_id = ?", campaignID).First(&campaign).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

// GetByAdvertiserID 根据广告主ID获取广告系列列表
func (r *campaignRepository) GetByAdvertiserID(ctx context.Context, advertiserID uint64) ([]*model.Campaign, error) {
	var list []*model.Campaign
	if err := r.db.WithContext(ctx).Where("advertiser_id = ?", advertiserID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// Create 创建广告系列
func (r *campaignRepository) Create(ctx context.Context, campaign *model.Campaign) error {
	return r.db.WithContext(ctx).Create(campaign).Error
}

// Update 更新广告系列
func (r *campaignRepository) Update(ctx context.Context, campaign *model.Campaign) error {
	return r.db.WithContext(ctx).Save(campaign).Error
}

// UpdateStatus 批量更新状态
func (r *campaignRepository) UpdateStatus(ctx context.Context, ids []uint64, status string) error {
	return r.db.WithContext(ctx).Model(&model.Campaign{}).
		Where("id IN ?", ids).
		Update("status", status).Error
}

// Delete 删除广告系列
func (r *campaignRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Campaign{}, id).Error
}

// BatchDelete 批量删除广告系列
func (r *campaignRepository) BatchDelete(ctx context.Context, ids []uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Campaign{}, ids).Error
}

// ExistsByCampaignID 检查广告系列ID是否存在
func (r *campaignRepository) ExistsByCampaignID(ctx context.Context, campaignID uint64) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Campaign{}).Where("campaign_id = ?", campaignID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
