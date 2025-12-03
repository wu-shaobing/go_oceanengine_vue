package service

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	adRepo "oceanengine-backend/internal/app/ad/repository"
	advRepo "oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/internal/app/creative/dto"
	"oceanengine-backend/internal/app/creative/model"
	"oceanengine-backend/internal/app/creative/repository"
	mediaModel "oceanengine-backend/internal/app/media/model"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// CreativeService 创意服务
type CreativeService struct {
	db      *gorm.DB
	repo    repository.CreativeRepository
	adRepo  adRepo.AdRepository
	advRepo advRepo.AdvertiserRepository
}

// NewCreativeService 创建创意服务
func NewCreativeService(db *gorm.DB) *CreativeService {
	return &CreativeService{
		db:      db,
		repo:    repository.NewCreativeRepository(db),
		adRepo:  adRepo.NewAdRepository(db),
		advRepo: advRepo.NewAdvertiserRepository(db),
	}
}

// GetList 获取创意列表
func (s *CreativeService) GetList(ctx context.Context, req *dto.CreativeListReq) ([]*dto.CreativeListResp, int64, error) {
	list, total, err := s.repo.GetList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.CreativeListResp, len(list))
	for i, c := range list {
		result[i] = &dto.CreativeListResp{
			ID:           c.ID,
			AdvertiserID: c.AdvertiserID,
			AdID:         c.AdID,
			CreativeID:   c.CreativeID,
			Title:        c.Title,
			ImageMode:    c.ImageMode,
			OptStatus:    c.OptStatus,
			Status:       c.Status,
			ImageURL:     c.ImageURL,
			CreatedAt:    c.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// Create 创建创意
func (s *CreativeService) Create(ctx context.Context, req *dto.CreativeCreateReq) (map[string]interface{}, error) {
	// 校验广告组是否存在
	ad, err := s.adRepo.GetByAdID(ctx, req.AdID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrAdNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	if ad.AdvertiserID != req.AdvertiserID {
		return nil, errcode.New(errcode.ErrInvalidParams)
	}

	// 获取广告主信息，用于调用 OE API
	adv, err := s.advRepo.GetByID(ctx, req.AdvertiserID)
	if err != nil {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}

	var creativeID uint64

	// 如果有 access_token，调用 OE API 创建创意
	if adv.AccessToken != "" {
		client := oceanengine.NewClient("", "")
		client.SetAccessToken(adv.AccessToken)
		creativeSvc := oceanengine.NewCreativeService(client)

		oeReq := &oceanengine.CreativeCreateRequest{
			AdvertiserID: int64(adv.AdvertiserID),
			AdID:         int64(req.AdID),
			Title:        req.Title,
			Source:       req.Source,
			ImageMode:    req.ImageMode,
			ImageIDs:     req.ImageIDs,
			VideoID:      req.VideoID,
			CallToAction: req.ActionText,
			WebURL:       req.ActionURL,
		}

		oeResp, err := creativeSvc.Create(ctx, oeReq)
		if err != nil {
			// OE API 调用失败，回退到本地创建
			creativeID = uint64(time.Now().UnixNano())
		} else {
			creativeID = uint64(oeResp.CreativeID)
		}
	} else {
		// 无 token，本地生成唯一 ID
		creativeID = uint64(time.Now().UnixNano())
	}

	// 组装创意数据
	creative := &model.Creative{
		AdvertiserID: req.AdvertiserID,
		AdID:         req.AdID,
		CreativeID:   creativeID,
		Title:        req.Title,
		Source:       req.Source,
		ImageMode:    req.ImageMode,
		ActionText:   req.ActionText,
		LandingURL:   req.ActionURL,
		OptStatus:    model.CreativeOptStatusEnable,
		Status:       "CREATIVE_STATUS_OK",
	}

	// 如果有图片/视频素材，取首个设置URL（简化处理）
	if len(req.ImageIDs) > 0 {
		var img mediaModel.MaterialImage
		if err := s.db.WithContext(ctx).
			Where("advertiser_id = ? AND image_id = ?", req.AdvertiserID, req.ImageIDs[0]).
			First(&img).Error; err == nil {
			creative.ImageURL = img.URL
		}
	}
	if req.VideoID != "" {
		var vid mediaModel.MaterialVideo
		if err := s.db.WithContext(ctx).
			Where("advertiser_id = ? AND video_id = ?", req.AdvertiserID, req.VideoID).
			First(&vid).Error; err == nil {
			creative.VideoURL = vid.URL
			creative.ThumbURL = vid.PosterURL
		}
	}

	if err := s.repo.Create(ctx, creative); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return map[string]interface{}{
		"id":          creative.ID,
		"creative_id": creative.CreativeID,
	}, nil
}

// GetByID 获取创意详情
func (s *CreativeService) GetByID(ctx context.Context, id uint64) (*dto.CreativeDetailResp, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrCreativeNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.toDetailResp(c), nil
}

// Update 更新创意
func (s *CreativeService) Update(ctx context.Context, id uint64, req *dto.CreativeUpdateReq) error {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrCreativeNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if req.Title != "" {
		c.Title = req.Title
	}
	if req.OptStatus != "" {
		c.OptStatus = req.OptStatus
	}
	if req.Remark != "" {
		c.Remark = req.Remark
	}

	if err := s.repo.Update(ctx, c); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdateStatus 批量更新状态
func (s *CreativeService) UpdateStatus(ctx context.Context, req *dto.CreativeStatusUpdateReq) error {
	if err := s.repo.UpdateStatus(ctx, req.IDs, req.OptStatus); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// Delete 删除创意
func (s *CreativeService) Delete(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrCreativeNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// toDetailResp 转换为详情响应
func (s *CreativeService) toDetailResp(c *model.Creative) *dto.CreativeDetailResp {
	var lastSyncAt string
	if c.LastSyncAt != nil {
		lastSyncAt = c.LastSyncAt.Format("2006-01-02 15:04:05")
	}

	return &dto.CreativeDetailResp{
		ID:           c.ID,
		AdvertiserID: c.AdvertiserID,
		AdID:         c.AdID,
		CreativeID:   c.CreativeID,
		Title:        c.Title,
		Source:       c.Source,
		ImageMode:    c.ImageMode,
		OptStatus:    c.OptStatus,
		Status:       c.Status,
		ImageURL:     c.ImageURL,
		VideoURL:     c.VideoURL,
		ThumbURL:     c.ThumbURL,
		ActionType:   c.ActionType,
		ActionText:   c.ActionText,
		LandingURL:   c.LandingURL,
		Remark:       c.Remark,
		LastSyncAt:   lastSyncAt,
		CreatedAt:    c.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    c.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
