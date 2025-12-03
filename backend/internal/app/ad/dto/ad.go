package dto

import "oceanengine-backend/pkg/utils"

// AdListReq 广告组列表请求
type AdListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id"`
	CampaignID   uint64 `form:"campaign_id"`
	Name         string `form:"name"`
	Status       string `form:"status"`
}

// AdListResp 广告组列表响应
type AdListResp struct {
	ID           uint64  `json:"id"`
	AdvertiserID uint64  `json:"advertiser_id"`
	CampaignID   uint64  `json:"campaign_id"`
	AdID         uint64  `json:"ad_id"`
	Name         string  `json:"name"`
	OptStatus    string  `json:"opt_status"`
	Status       string  `json:"status"`
	Budget       float64 `json:"budget"`
	Bid          float64 `json:"bid"`
	CreatedAt    string  `json:"created_at"`
}

// AdDetailResp 广告组详情响应
type AdDetailResp struct {
	ID             uint64  `json:"id"`
	AdvertiserID   uint64  `json:"advertiser_id"`
	CampaignID     uint64  `json:"campaign_id"`
	AdID           uint64  `json:"ad_id"`
	Name           string  `json:"name"`
	OptStatus      string  `json:"opt_status"`
	Status         string  `json:"status"`
	Budget         float64 `json:"budget"`
	BudgetMode     string  `json:"budget_mode"`
	DeliveryRange  string  `json:"delivery_range"`
	Pricing        string  `json:"pricing"`
	Bid            float64 `json:"bid"`
	ConvertID      uint64  `json:"convert_id"`
	StartTime      string  `json:"start_time"`
	EndTime        string  `json:"end_time"`
	ScheduleType   string  `json:"schedule_type"`
	AudienceType   string  `json:"audience_type"`
	TargetSettings string  `json:"target_settings"`
	Remark         string  `json:"remark"`
	LastSyncAt     string  `json:"last_sync_at"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// AdCreateReq 创建广告组请求
type AdCreateReq struct {
	AdvertiserID  uint64  `json:"advertiser_id" binding:"required"`
	CampaignID    uint64  `json:"campaign_id" binding:"required"`
	Name          string  `json:"name" binding:"required,max=255"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	DeliveryRange string  `json:"delivery_range"`
	Pricing       string  `json:"pricing" binding:"required"`
	Bid           float64 `json:"bid"`
}

// AdUpdateReq 更新广告组请求
type AdUpdateReq struct {
	Name      string  `json:"name" binding:"max=255"`
	Budget    float64 `json:"budget"`
	Bid       float64 `json:"bid"`
	OptStatus string  `json:"opt_status" binding:"omitempty,oneof=ENABLE DISABLE"`
	Remark    string  `json:"remark" binding:"max=500"`
}

// AdStatusUpdateReq 批量更新广告组状态请求
type AdStatusUpdateReq struct {
	IDs       []uint64 `json:"ids" binding:"required,min=1"`
	OptStatus string   `json:"opt_status" binding:"required,oneof=ENABLE DISABLE"`
}

// AdSyncResp 广告组同步响应
type AdSyncResp struct {
	SyncCount int    `json:"sync_count"`
	SyncAt    string `json:"sync_at"`
}
