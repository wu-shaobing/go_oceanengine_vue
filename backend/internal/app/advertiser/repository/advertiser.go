package repository

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/advertiser/dto"
	"oceanengine-backend/internal/app/advertiser/model"
)

// AdvertiserRepository 广告主仓库接口
type AdvertiserRepository interface {
	GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*model.Advertiser, int64, error)
	GetByID(ctx context.Context, id uint64) (*model.Advertiser, error)
	GetByAdvertiserID(ctx context.Context, advertiserID uint64) (*model.Advertiser, error)
	Create(ctx context.Context, advertiser *model.Advertiser) error
	Update(ctx context.Context, advertiser *model.Advertiser) error
	Delete(ctx context.Context, id uint64) error
	ExistsByAdvertiserID(ctx context.Context, advertiserID uint64) (bool, error)
}

// advertiserRepository 广告主仓库实现
type advertiserRepository struct {
	db *gorm.DB
}

// NewAdvertiserRepository 创建广告主仓库
func NewAdvertiserRepository(db *gorm.DB) AdvertiserRepository {
	return &advertiserRepository{db: db}
}

// GetList 获取广告主列表
func (r *advertiserRepository) GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*model.Advertiser, int64, error) {
	var list []*model.Advertiser
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Advertiser{})

	// 关键词搜索
	if req.Keyword != "" {
		query = query.Where("name LIKE ? OR company LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 状态筛选
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := req.GetOffset()
	limit := req.GetLimit()
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetByID 根据 ID 获取广告主
func (r *advertiserRepository) GetByID(ctx context.Context, id uint64) (*model.Advertiser, error) {
	var advertiser model.Advertiser
	if err := r.db.WithContext(ctx).First(&advertiser, id).Error; err != nil {
		return nil, err
	}
	return &advertiser, nil
}

// GetByAdvertiserID 根据 Ocean Engine 广告主 ID 获取
func (r *advertiserRepository) GetByAdvertiserID(ctx context.Context, advertiserID uint64) (*model.Advertiser, error) {
	var advertiser model.Advertiser
	if err := r.db.WithContext(ctx).Where("advertiser_id = ?", advertiserID).First(&advertiser).Error; err != nil {
		return nil, err
	}
	return &advertiser, nil
}

// Create 创建广告主
func (r *advertiserRepository) Create(ctx context.Context, advertiser *model.Advertiser) error {
	return r.db.WithContext(ctx).Create(advertiser).Error
}

// Update 更新广告主
func (r *advertiserRepository) Update(ctx context.Context, advertiser *model.Advertiser) error {
	return r.db.WithContext(ctx).Save(advertiser).Error
}

// Delete 删除广告主
func (r *advertiserRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Advertiser{}, id).Error
}

// ExistsByAdvertiserID 检查广告主是否存在
func (r *advertiserRepository) ExistsByAdvertiserID(ctx context.Context, advertiserID uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Advertiser{}).Where("advertiser_id = ?", advertiserID).Count(&count).Error
	return count > 0, err
}

// FundRepository 资金流水仓库接口
type FundRepository interface {
	GetList(ctx context.Context, req *dto.FundListReq) ([]*model.AdvertiserFund, int64, error)
	Create(ctx context.Context, fund *model.AdvertiserFund) error
	BatchCreate(ctx context.Context, funds []*model.AdvertiserFund) error
}

// fundRepository 资金流水仓库实现
type fundRepository struct {
	db *gorm.DB
}

// NewFundRepository 创建资金流水仓库
func NewFundRepository(db *gorm.DB) FundRepository {
	return &fundRepository{db: db}
}

// GetList 获取资金流水列表
func (r *fundRepository) GetList(ctx context.Context, req *dto.FundListReq) ([]*model.AdvertiserFund, int64, error) {
	var list []*model.AdvertiserFund
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AdvertiserFund{}).Where("advertiser_id = ?", req.AdvertiserID)

	// 日期筛选
	if req.StartDate != "" {
		query = query.Where("transaction_time >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		query = query.Where("transaction_time <= ?", req.EndDate+" 23:59:59")
	}

	// 交易类型筛选
	if req.TransactionType != "" {
		query = query.Where("transaction_type = ?", req.TransactionType)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := req.GetOffset()
	limit := req.GetLimit()
	if err := query.Order("transaction_time DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// Create 创建资金流水
func (r *fundRepository) Create(ctx context.Context, fund *model.AdvertiserFund) error {
	return r.db.WithContext(ctx).Create(fund).Error
}

// BatchCreate 批量创建资金流水
func (r *fundRepository) BatchCreate(ctx context.Context, funds []*model.AdvertiserFund) error {
	return r.db.WithContext(ctx).CreateInBatches(funds, 100).Error
}
