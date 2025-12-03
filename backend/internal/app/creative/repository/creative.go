package repository

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/creative/dto"
	"oceanengine-backend/internal/app/creative/model"
)

// CreativeRepository 创意仓储接口
type CreativeRepository interface {
	GetList(ctx context.Context, req *dto.CreativeListReq) ([]*model.Creative, int64, error)
	GetByID(ctx context.Context, id uint64) (*model.Creative, error)
	GetByCreativeID(ctx context.Context, creativeID uint64) (*model.Creative, error)
	GetByAdID(ctx context.Context, adID uint64) ([]*model.Creative, error)
	Create(ctx context.Context, creative *model.Creative) error
	Update(ctx context.Context, creative *model.Creative) error
	UpdateStatus(ctx context.Context, ids []uint64, status string) error
	Delete(ctx context.Context, id uint64) error
	ExistsByCreativeID(ctx context.Context, creativeID uint64) (bool, error)
}

type creativeRepository struct {
	db *gorm.DB
}

// NewCreativeRepository 创建创意仓储
func NewCreativeRepository(db *gorm.DB) CreativeRepository {
	return &creativeRepository{db: db}
}

// GetList 获取创意列表
func (r *creativeRepository) GetList(ctx context.Context, req *dto.CreativeListReq) ([]*model.Creative, int64, error) {
	var list []*model.Creative
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Creative{})

	if req.AdvertiserID > 0 {
		query = query.Where("advertiser_id = ?", req.AdvertiserID)
	}
	if req.AdID > 0 {
		query = query.Where("ad_id = ?", req.AdID)
	}
	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
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

// GetByID 根据ID获取创意
func (r *creativeRepository) GetByID(ctx context.Context, id uint64) (*model.Creative, error) {
	var creative model.Creative
	if err := r.db.WithContext(ctx).First(&creative, id).Error; err != nil {
		return nil, err
	}
	return &creative, nil
}

// GetByCreativeID 根据创意ID获取
func (r *creativeRepository) GetByCreativeID(ctx context.Context, creativeID uint64) (*model.Creative, error) {
	var creative model.Creative
	if err := r.db.WithContext(ctx).Where("creative_id = ?", creativeID).First(&creative).Error; err != nil {
		return nil, err
	}
	return &creative, nil
}

// GetByAdID 根据广告组ID获取创意列表
func (r *creativeRepository) GetByAdID(ctx context.Context, adID uint64) ([]*model.Creative, error) {
	var list []*model.Creative
	if err := r.db.WithContext(ctx).Where("ad_id = ?", adID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// Create 创建创意
func (r *creativeRepository) Create(ctx context.Context, creative *model.Creative) error {
	return r.db.WithContext(ctx).Create(creative).Error
}

// Update 更新创意
func (r *creativeRepository) Update(ctx context.Context, creative *model.Creative) error {
	return r.db.WithContext(ctx).Save(creative).Error
}

// UpdateStatus 批量更新状态
func (r *creativeRepository) UpdateStatus(ctx context.Context, ids []uint64, status string) error {
	return r.db.WithContext(ctx).Model(&model.Creative{}).
		Where("id IN ?", ids).
		Update("opt_status", status).Error
}

// Delete 删除创意
func (r *creativeRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Creative{}, id).Error
}

// ExistsByCreativeID 检查创意ID是否存在
func (r *creativeRepository) ExistsByCreativeID(ctx context.Context, creativeID uint64) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Creative{}).Where("creative_id = ?", creativeID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
