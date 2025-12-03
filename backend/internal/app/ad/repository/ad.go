package repository

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/ad/dto"
	"oceanengine-backend/internal/app/ad/model"
)

// AdRepository 广告组仓储接口
type AdRepository interface {
	GetList(ctx context.Context, req *dto.AdListReq) ([]*model.Ad, int64, error)
	GetByID(ctx context.Context, id uint64) (*model.Ad, error)
	GetByAdID(ctx context.Context, adID uint64) (*model.Ad, error)
	GetByCampaignID(ctx context.Context, campaignID uint64) ([]*model.Ad, error)
	Create(ctx context.Context, ad *model.Ad) error
	Update(ctx context.Context, ad *model.Ad) error
	UpdateStatus(ctx context.Context, ids []uint64, status string) error
	Delete(ctx context.Context, id uint64) error
	ExistsByAdID(ctx context.Context, adID uint64) (bool, error)
}

type adRepository struct {
	db *gorm.DB
}

// NewAdRepository 创建广告组仓储
func NewAdRepository(db *gorm.DB) AdRepository {
	return &adRepository{db: db}
}

// GetList 获取广告组列表
func (r *adRepository) GetList(ctx context.Context, req *dto.AdListReq) ([]*model.Ad, int64, error) {
	var list []*model.Ad
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Ad{})

	if req.AdvertiserID > 0 {
		query = query.Where("advertiser_id = ?", req.AdvertiserID)
	}
	if req.CampaignID > 0 {
		query = query.Where("campaign_id = ?", req.CampaignID)
	}
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != "" {
		query = query.Where("opt_status = ?", req.Status)
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

// GetByID 根据ID获取广告组
func (r *adRepository) GetByID(ctx context.Context, id uint64) (*model.Ad, error) {
	var ad model.Ad
	if err := r.db.WithContext(ctx).First(&ad, id).Error; err != nil {
		return nil, err
	}
	return &ad, nil
}

// GetByAdID 根据广告组ID获取
func (r *adRepository) GetByAdID(ctx context.Context, adID uint64) (*model.Ad, error) {
	var ad model.Ad
	if err := r.db.WithContext(ctx).Where("ad_id = ?", adID).First(&ad).Error; err != nil {
		return nil, err
	}
	return &ad, nil
}

// GetByCampaignID 根据广告系列ID获取广告组列表
func (r *adRepository) GetByCampaignID(ctx context.Context, campaignID uint64) ([]*model.Ad, error) {
	var list []*model.Ad
	if err := r.db.WithContext(ctx).Where("campaign_id = ?", campaignID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// Create 创建广告组
func (r *adRepository) Create(ctx context.Context, ad *model.Ad) error {
	return r.db.WithContext(ctx).Create(ad).Error
}

// Update 更新广告组
func (r *adRepository) Update(ctx context.Context, ad *model.Ad) error {
	return r.db.WithContext(ctx).Save(ad).Error
}

// UpdateStatus 批量更新状态
func (r *adRepository) UpdateStatus(ctx context.Context, ids []uint64, status string) error {
	return r.db.WithContext(ctx).Model(&model.Ad{}).
		Where("id IN ?", ids).
		Update("opt_status", status).Error
}

// Delete 删除广告组
func (r *adRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Ad{}, id).Error
}

// ExistsByAdID 检查广告组ID是否存在
func (r *adRepository) ExistsByAdID(ctx context.Context, adID uint64) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Ad{}).Where("ad_id = ?", adID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
