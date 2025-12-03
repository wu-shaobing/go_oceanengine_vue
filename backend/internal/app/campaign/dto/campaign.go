package dto

import "oceanengine-backend/pkg/utils"

// CampaignListReq 广告系列列表请求
type CampaignListReq struct {
	utils.Pagination
	AdvertiserID  uint64 `form:"advertiser_id"`
	Name          string `form:"name"`
	Status        string `form:"status"`
	MarketingGoal string `form:"marketing_goal"`
}

// CampaignListResp 广告系列列表响应
type CampaignListResp struct {
	ID            uint64  `json:"id"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	CampaignID    uint64  `json:"campaign_id"`
	Name          string  `json:"name"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	MarketingGoal string  `json:"marketing_goal"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
}

// CampaignDetailResp 广告系列详情响应
type CampaignDetailResp struct {
	ID            uint64  `json:"id"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	CampaignID    uint64  `json:"campaign_id"`
	Name          string  `json:"name"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	MarketingGoal string  `json:"marketing_goal"`
	DeliveryMode  string  `json:"delivery_mode"`
	Status        string  `json:"status"`
	LandingType   string  `json:"landing_type"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	Remark        string  `json:"remark"`
	LastSyncAt    string  `json:"last_sync_at"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// CampaignCreateReq 创建广告系列请求
type CampaignCreateReq struct {
	AdvertiserID  uint64  `json:"advertiser_id" binding:"required"`
	Name          string  `json:"name" binding:"required,max=255"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode" binding:"required,oneof=BUDGET_MODE_INFINITE BUDGET_MODE_DAY BUDGET_MODE_TOTAL"`
	MarketingGoal string  `json:"marketing_goal" binding:"required"`
	LandingType   string  `json:"landing_type"`
}

// CampaignUpdateReq 更新广告系列请求
type CampaignUpdateReq struct {
	Name       string  `json:"name" binding:"max=255"`
	Budget     float64 `json:"budget"`
	BudgetMode string  `json:"budget_mode" binding:"omitempty,oneof=BUDGET_MODE_INFINITE BUDGET_MODE_DAY BUDGET_MODE_TOTAL"`
	Status     string  `json:"status" binding:"omitempty,oneof=ENABLE DISABLE"`
	Remark     string  `json:"remark" binding:"max=500"`
}

// CampaignStatusUpdateReq 更新广告系列状态请求
type CampaignStatusUpdateReq struct {
	IDs    []uint64 `json:"ids" binding:"required,min=1"`
	Status string   `json:"status" binding:"required,oneof=ENABLE DISABLE"`
}

// CampaignSyncResp 广告系列同步响应
type CampaignSyncResp struct {
	SyncCount int    `json:"sync_count"`
	SyncAt    string `json:"sync_at"`
}
