package dto

import "oceanengine-backend/pkg/utils"

// ImageListReq 图片列表请求
type ImageListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	Keyword      string `form:"keyword"`
}

// ImageUploadReq 图片上传请求
type ImageUploadReq struct {
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	UploadType   string `form:"upload_type" binding:"required"` // file, url
	URL          string `form:"url"`                            // 图片URL（URL上传时）
}

// ImageListResp 图片列表响应项
type ImageListResp struct {
	ID           uint64 `json:"id"`
	ImageID      string `json:"image_id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	Filename     string `json:"filename"`
	Size         int64  `json:"size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Format       string `json:"format"`
	URL          string `json:"url"`
	CreatedAt    string `json:"created_at"`
}

// VideoListReq 视频列表请求
type VideoListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	Keyword      string `form:"keyword"`
}

// VideoUploadReq 视频上传请求
type VideoUploadReq struct {
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
}

// VideoListResp 视频列表响应项
type VideoListResp struct {
	ID           uint64  `json:"id"`
	VideoID      string  `json:"video_id"`
	AdvertiserID uint64  `json:"advertiser_id"`
	Filename     string  `json:"filename"`
	Size         int64   `json:"size"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	Duration     float64 `json:"duration"`
	Format       string  `json:"format"`
	URL          string  `json:"url"`
	PosterURL    string  `json:"poster_url"`
	CreatedAt    string  `json:"created_at"`
}

// MaterialDeleteReq 删除素材请求
type MaterialDeleteReq struct {
	AdvertiserID uint64 `json:"advertiser_id" binding:"required"`
	MaterialType string `json:"material_type" binding:"required"` // image, video
	MaterialID   string `json:"material_id" binding:"required"`
}

// ImageUploadResp 图片上传响应
type ImageUploadResp struct {
	ImageID string `json:"image_id"`
	URL     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Size    int64  `json:"size"`
	Format  string `json:"format"`
}

// VideoUploadResp 视频上传响应
type VideoUploadResp struct {
	VideoID   string  `json:"video_id"`
	URL       string  `json:"url"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Duration  float64 `json:"duration"`
	Size      int64   `json:"size"`
	Format    string  `json:"format"`
	PosterURL string  `json:"poster_url"`
}
