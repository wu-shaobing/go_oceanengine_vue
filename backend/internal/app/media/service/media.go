package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	advRepo "oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/internal/app/media/dto"
	"oceanengine-backend/internal/app/media/model"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// MediaService 素材服务
type MediaService struct {
	db      *gorm.DB
	advRepo advRepo.AdvertiserRepository
}

// NewMediaService 创建素材服务
func NewMediaService(db *gorm.DB) *MediaService {
	return &MediaService{
		db:      db,
		advRepo: advRepo.NewAdvertiserRepository(db),
	}
}

// GetImageList 获取图片列表
func (s *MediaService) GetImageList(ctx context.Context, req *dto.ImageListReq) ([]*dto.ImageListResp, int64, error) {
	var images []*model.MaterialImage
	var total int64

	query := s.db.WithContext(ctx).Model(&model.MaterialImage{}).Where("advertiser_id = ?", req.AdvertiserID)

	if req.Keyword != "" {
		query = query.Where("filename LIKE ?", "%"+req.Keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&images).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.ImageListResp, len(images))
	for i, img := range images {
		result[i] = &dto.ImageListResp{
			ID:           img.ID,
			ImageID:      img.ImageID,
			AdvertiserID: img.AdvertiserID,
			Filename:     img.Filename,
			Size:         img.Size,
			Width:        img.Width,
			Height:       img.Height,
			Format:       img.Format,
			URL:          img.URL,
			CreatedAt:    img.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetImageByID 获取图片详情
func (s *MediaService) GetImageByID(ctx context.Context, id uint64) (*model.MaterialImage, error) {
	var image model.MaterialImage
	if err := s.db.WithContext(ctx).First(&image, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrMaterialNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &image, nil
}

// CreateImage 创建图片记录
func (s *MediaService) CreateImage(ctx context.Context, image *model.MaterialImage) error {
	if err := s.db.WithContext(ctx).Create(image).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// DeleteImage 删除图片
func (s *MediaService) DeleteImage(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.MaterialImage{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrMaterialNotFound)
	}
	return nil
}

// GetVideoList 获取视频列表
func (s *MediaService) GetVideoList(ctx context.Context, req *dto.VideoListReq) ([]*dto.VideoListResp, int64, error) {
	var videos []*model.MaterialVideo
	var total int64

	query := s.db.WithContext(ctx).Model(&model.MaterialVideo{}).Where("advertiser_id = ?", req.AdvertiserID)

	if req.Keyword != "" {
		query = query.Where("filename LIKE ?", "%"+req.Keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&videos).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.VideoListResp, len(videos))
	for i, vid := range videos {
		result[i] = &dto.VideoListResp{
			ID:           vid.ID,
			VideoID:      vid.VideoID,
			AdvertiserID: vid.AdvertiserID,
			Filename:     vid.Filename,
			Size:         vid.Size,
			Width:        vid.Width,
			Height:       vid.Height,
			Duration:     vid.Duration,
			Format:       vid.Format,
			URL:          vid.URL,
			PosterURL:    vid.PosterURL,
			CreatedAt:    vid.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetVideoByID 获取视频详情
func (s *MediaService) GetVideoByID(ctx context.Context, id uint64) (*model.MaterialVideo, error) {
	var video model.MaterialVideo
	if err := s.db.WithContext(ctx).First(&video, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrMaterialNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &video, nil
}

// CreateVideo 创建视频记录
func (s *MediaService) CreateVideo(ctx context.Context, video *model.MaterialVideo) error {
	if err := s.db.WithContext(ctx).Create(video).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// DeleteVideo 删除视频
func (s *MediaService) DeleteVideo(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.MaterialVideo{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrMaterialNotFound)
	}
	return nil
}

// UploadImage 上传图片
// 如果广告主有 access_token，则调用 OE API 上传
func (s *MediaService) UploadImage(ctx context.Context, advertiserID uint64, filename string, file io.Reader, size int64) (*dto.ImageUploadResp, error) {
	// 读取文件内容计算签名
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 计算MD5签名
	hash := md5.Sum(data)
	signature := fmt.Sprintf("%x", hash)

	// 检查是否已存在
	var existImage model.MaterialImage
	if err := s.db.WithContext(ctx).Where("advertiser_id = ? AND signature = ?", advertiserID, signature).First(&existImage).Error; err == nil {
		return &dto.ImageUploadResp{
			ImageID: existImage.ImageID,
			URL:     existImage.URL,
			Width:   existImage.Width,
			Height:  existImage.Height,
			Size:    existImage.Size,
			Format:  existImage.Format,
		}, nil
	}

	// 获取文件格式
	ext := strings.ToLower(filepath.Ext(filename))
	format := strings.TrimPrefix(ext, ".")
	if format == "jpeg" {
		format = "jpg"
	}

	var imageID, imageURL string
	var width, height int

	// 获取广告主信息，尝试调用 OE API
	adv, _ := s.advRepo.GetByID(ctx, advertiserID)
	if adv != nil && adv.AccessToken != "" {
		client := oceanengine.NewClient("", "")
		client.SetAccessToken(adv.AccessToken)
		fileSvc := oceanengine.NewFileService(client)

		oeResp, err := fileSvc.UploadImageByBytes(ctx, int64(adv.AdvertiserID), filename, data)
		if err == nil {
			imageID = oeResp.ImageID
			imageURL = oeResp.URL
			width = oeResp.Width
			height = oeResp.Height
		}
	}

	// 如果 OE API 调用失败或无 token，回退到本地生成
	if imageID == "" {
		imageID = fmt.Sprintf("img_%s", uuid.New().String()[:8])
		imageURL = fmt.Sprintf("https://cdn.oceanengine.com/material/%d/%s%s", advertiserID, imageID, ext)
	}

	// 保存数据库记录
	image := &model.MaterialImage{
		ImageID:      imageID,
		AdvertiserID: advertiserID,
		Filename:     filename,
		Size:         size,
		Width:        width,
		Height:       height,
		Format:       format,
		URL:          imageURL,
		Signature:    signature,
	}

	if err := s.db.WithContext(ctx).Create(image).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.ImageUploadResp{
		ImageID: imageID,
		URL:     imageURL,
		Width:   width,
		Height:  height,
		Size:    size,
		Format:  format,
	}, nil
}

// UploadVideo 上传视频
// 如果广告主有 access_token，则调用 OE API 上传
func (s *MediaService) UploadVideo(ctx context.Context, advertiserID uint64, filename string, file io.Reader, size int64) (*dto.VideoUploadResp, error) {
	// 读取文件内容计算签名
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 计算MD5签名
	hash := md5.Sum(data)
	signature := fmt.Sprintf("%x", hash)

	// 检查是否已存在
	var existVideo model.MaterialVideo
	if err := s.db.WithContext(ctx).Where("advertiser_id = ? AND signature = ?", advertiserID, signature).First(&existVideo).Error; err == nil {
		return &dto.VideoUploadResp{
			VideoID:   existVideo.VideoID,
			URL:       existVideo.URL,
			Width:     existVideo.Width,
			Height:    existVideo.Height,
			Duration:  existVideo.Duration,
			Size:      existVideo.Size,
			Format:    existVideo.Format,
			PosterURL: existVideo.PosterURL,
		}, nil
	}

	// 获取文件格式
	ext := strings.ToLower(filepath.Ext(filename))
	format := strings.TrimPrefix(ext, ".")

	var videoID, videoURL, posterURL string
	var width, height int
	var duration float64

	// 获取广告主信息，尝试调用 OE API
	adv, _ := s.advRepo.GetByID(ctx, advertiserID)
	if adv != nil && adv.AccessToken != "" {
		client := oceanengine.NewClient("", "")
		client.SetAccessToken(adv.AccessToken)
		fileSvc := oceanengine.NewFileService(client)

		oeResp, err := fileSvc.UploadVideoByBytes(ctx, int64(adv.AdvertiserID), filename, data)
		if err == nil {
			videoID = oeResp.VideoID
			videoURL = oeResp.URL
			posterURL = oeResp.PosterURL
			width = oeResp.Width
			height = oeResp.Height
			duration = oeResp.Duration
		}
	}

	// 如果 OE API 调用失败或无 token，回退到本地生成
	if videoID == "" {
		videoID = fmt.Sprintf("vid_%s", uuid.New().String()[:8])
		videoURL = fmt.Sprintf("https://cdn.oceanengine.com/material/%d/%s%s", advertiserID, videoID, ext)
		posterURL = fmt.Sprintf("https://cdn.oceanengine.com/material/%d/%s_poster.jpg", advertiserID, videoID)
	}

	// 保存数据库记录
	video := &model.MaterialVideo{
		VideoID:      videoID,
		AdvertiserID: advertiserID,
		Filename:     filename,
		Size:         size,
		Width:        width,
		Height:       height,
		Duration:     duration,
		Format:       format,
		URL:          videoURL,
		PosterURL:    posterURL,
		Signature:    signature,
		CreatedAt:    time.Now(),
	}

	if err := s.db.WithContext(ctx).Create(video).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.VideoUploadResp{
		VideoID:   videoID,
		URL:       videoURL,
		Width:     width,
		Height:    height,
		Duration:  duration,
		Size:      size,
		Format:    format,
		PosterURL: posterURL,
	}, nil
}
