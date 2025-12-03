package service

import (
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/audience/dto"
	"oceanengine-backend/internal/app/audience/model"
	"oceanengine-backend/pkg/errcode"
)

// AudienceService 人群定向服务
type AudienceService struct {
	db *gorm.DB
}

// NewAudienceService 创建人群定向服务
func NewAudienceService(db *gorm.DB) *AudienceService {
	return &AudienceService{db: db}
}

// GetPackageList 获取定向包列表
func (s *AudienceService) GetPackageList(ctx context.Context, req *dto.AudiencePackageListReq) ([]*dto.AudiencePackageListResp, int64, error) {
	var packages []*model.AudiencePackage
	var total int64

	query := s.db.WithContext(ctx).Model(&model.AudiencePackage{}).Where("advertiser_id = ?", req.AdvertiserID)

	if req.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.LandingType != "" {
		query = query.Where("landing_type = ?", req.LandingType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&packages).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.AudiencePackageListResp, len(packages))
	for i, pkg := range packages {
		result[i] = &dto.AudiencePackageListResp{
			ID:            pkg.ID,
			PackageID:     pkg.PackageID,
			AdvertiserID:  pkg.AdvertiserID,
			Name:          pkg.Name,
			Description:   pkg.Description,
			Status:        pkg.Status,
			LandingType:   pkg.LandingType,
			DeliveryRange: pkg.DeliveryRange,
			CreatedAt:     pkg.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetPackageByID 获取定向包详情
func (s *AudienceService) GetPackageByID(ctx context.Context, id uint64) (*model.AudiencePackage, error) {
	var pkg model.AudiencePackage
	if err := s.db.WithContext(ctx).First(&pkg, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &pkg, nil
}

// CreatePackage 创建定向包
func (s *AudienceService) CreatePackage(ctx context.Context, req *dto.AudiencePackageCreateReq) error {
	audienceJSON, _ := json.Marshal(req.Audience)

	pkg := &model.AudiencePackage{
		AdvertiserID:  req.AdvertiserID,
		Name:          req.Name,
		Description:   req.Description,
		LandingType:   req.LandingType,
		DeliveryRange: req.DeliveryRange,
		Audience:      string(audienceJSON),
		Status:        model.AudiencePackageStatusEnable,
	}

	if err := s.db.WithContext(ctx).Create(pkg).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdatePackage 更新定向包
func (s *AudienceService) UpdatePackage(ctx context.Context, req *dto.AudiencePackageUpdateReq) error {
	var pkg model.AudiencePackage
	if err := s.db.WithContext(ctx).First(&pkg, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.DeliveryRange != "" {
		updates["delivery_range"] = req.DeliveryRange
	}
	if req.Audience != nil {
		audienceJSON, _ := json.Marshal(req.Audience)
		updates["audience"] = string(audienceJSON)
	}

	if err := s.db.WithContext(ctx).Model(&pkg).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// DeletePackage 删除定向包
func (s *AudienceService) DeletePackage(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.AudiencePackage{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrNotFound)
	}
	return nil
}

// GetCustomAudienceList 获取自定义人群列表
func (s *AudienceService) GetCustomAudienceList(ctx context.Context, req *dto.CustomAudienceListReq) ([]*dto.CustomAudienceListResp, int64, error) {
	var audiences []*model.CustomAudience
	var total int64

	query := s.db.WithContext(ctx).Model(&model.CustomAudience{}).Where("advertiser_id = ?", req.AdvertiserID)

	if req.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.Source != "" {
		query = query.Where("source = ?", req.Source)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&audiences).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.CustomAudienceListResp, len(audiences))
	for i, aud := range audiences {
		result[i] = &dto.CustomAudienceListResp{
			ID:               aud.ID,
			CustomAudienceID: aud.CustomAudienceID,
			AdvertiserID:     aud.AdvertiserID,
			Name:             aud.Name,
			Source:           aud.Source,
			Status:           aud.Status,
			CoverNum:         aud.CoverNum,
			Tag:              aud.Tag,
			PushStatus:       aud.PushStatus,
			CreatedAt:        aud.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetCustomAudienceByID 获取自定义人群详情
func (s *AudienceService) GetCustomAudienceByID(ctx context.Context, id uint64) (*model.CustomAudience, error) {
	var audience model.CustomAudience
	if err := s.db.WithContext(ctx).First(&audience, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &audience, nil
}

// DeleteCustomAudience 删除自定义人群
func (s *AudienceService) DeleteCustomAudience(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.CustomAudience{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrNotFound)
	}
	return nil
}
