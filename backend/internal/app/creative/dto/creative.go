package dto

import "oceanengine-backend/pkg/utils"

// CreativeListReq 创意列表请求
type CreativeListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id"`
	AdID         uint64 `form:"ad_id"`
	Title        string `form:"title"`
	Status       string `form:"status"`
}

// CreativeListResp 创意列表响应
type CreativeListResp struct {
	ID           uint64 `json:"id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	AdID         uint64 `json:"ad_id"`
	CreativeID   uint64 `json:"creative_id"`
	Title        string `json:"title"`
	ImageMode    string `json:"image_mode"`
	OptStatus    string `json:"opt_status"`
	Status       string `json:"status"`
	ImageURL     string `json:"image_url"`
	CreatedAt    string `json:"created_at"`
}

// CreativeDetailResp 创意详情响应
type CreativeDetailResp struct {
	ID           uint64 `json:"id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	AdID         uint64 `json:"ad_id"`
	CreativeID   uint64 `json:"creative_id"`
	Title        string `json:"title"`
	Source       string `json:"source"`
	ImageMode    string `json:"image_mode"`
	OptStatus    string `json:"opt_status"`
	Status       string `json:"status"`
	ImageURL     string `json:"image_url"`
	VideoURL     string `json:"video_url"`
	ThumbURL     string `json:"thumb_url"`
	ActionType   string `json:"action_type"`
	ActionText   string `json:"action_text"`
	LandingURL   string `json:"landing_url"`
	Remark       string `json:"remark"`
	LastSyncAt   string `json:"last_sync_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// CreativeCreateReq 创建创意请求
type CreativeCreateReq struct {
	AdvertiserID         uint64   `json:"advertiser_id" binding:"required"`
	AdID                 uint64   `json:"ad_id" binding:"required"`
	CreativeMaterialMode string   `json:"creative_material_mode"`
	Title                string   `json:"title" binding:"required,max=255"`
	Source               string   `json:"source" binding:"max=255"`
	ImageMode            string   `json:"image_mode" binding:"required"`
	ImageIDs             []string `json:"image_ids"`
	VideoID              string   `json:"video_id"`
	ActionText           string   `json:"action_text" binding:"max=64"`
	ActionURL            string   `json:"action_url" binding:"max=1024"`
}

// CreativeUpdateReq 更新创意请求
type CreativeUpdateReq struct {
	Title     string `json:"title" binding:"max=255"`
	OptStatus string `json:"opt_status" binding:"omitempty,oneof=ENABLE DISABLE"`
	Remark    string `json:"remark" binding:"max=500"`
}

// CreativeStatusUpdateReq 批量更新创意状态请求
type CreativeStatusUpdateReq struct {
	IDs       []uint64 `json:"ids" binding:"required,min=1"`
	OptStatus string   `json:"opt_status" binding:"required,oneof=ENABLE DISABLE"`
}
