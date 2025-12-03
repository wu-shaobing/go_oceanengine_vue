package dto

import "oceanengine-backend/pkg/utils"

// AdvertiserListReq 广告主列表请求
type AdvertiserListReq struct {
	utils.Pagination
	Keyword string `form:"keyword"`
	Status  string `form:"status"`
	Sync    bool   `form:"sync"` // 是否同步远程数据
}

// AdvertiserListResp 广告主列表响应
type AdvertiserListResp struct {
	ID           uint64  `json:"id"`
	AdvertiserID uint64  `json:"advertiser_id"`
	Name         string  `json:"name"`
	Company      string  `json:"company"`
	Status       string  `json:"status"`
	Balance      float64 `json:"balance"`
	Industry     string  `json:"industry"`
	CreatedAt    string  `json:"created_at"`
}

// AdvertiserDetailResp 广告主详情响应
type AdvertiserDetailResp struct {
	ID           uint64  `json:"id"`
	AdvertiserID uint64  `json:"advertiser_id"`
	Name         string  `json:"name"`
	Company      string  `json:"company"`
	Status       string  `json:"status"`
	Role         string  `json:"role"`
	Balance      float64 `json:"balance"`
	ValidBalance float64 `json:"valid_balance"`
	CashBalance  float64 `json:"cash_balance"`
	Industry     string  `json:"industry"`
	ContactName  string  `json:"contact_name"`
	ContactPhone string  `json:"contact_phone"`
	ContactEmail string  `json:"contact_email"`
	Address      string  `json:"address"`
	LastSyncAt   string  `json:"last_sync_at"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// AdvertiserCreateReq 创建广告主请求 (OAuth 授权后)
type AdvertiserCreateReq struct {
	AdvertiserID uint64 `json:"advertiser_id" binding:"required"`
	Name         string `json:"name" binding:"required,max=255"`
	Company      string `json:"company" binding:"max=255"`
	Industry     string `json:"industry" binding:"max=100"`
	Remark       string `json:"remark" binding:"max=500"`
}

// AdvertiserUpdateReq 更新广告主请求
type AdvertiserUpdateReq struct {
	ID           uint64 `json:"id" binding:"required"`
	ContactName  string `json:"contact_name" binding:"max=64"`
	ContactPhone string `json:"contact_phone" binding:"max=20"`
	ContactEmail string `json:"contact_email" binding:"max=128"`
	Address      string `json:"address" binding:"max=500"`
	Remark       string `json:"remark" binding:"max=500"`
}

// AdvertiserBalanceResp 广告主余额响应
type AdvertiserBalanceResp struct {
	Balance      float64 `json:"balance"`
	ValidBalance float64 `json:"valid_balance"`
	CashBalance  float64 `json:"cash_balance"`
	GrantBalance float64 `json:"grant_balance"`
	SyncAt       string  `json:"sync_at"`
}

// OAuthAuthorizeResp OAuth 授权响应
type OAuthAuthorizeResp struct {
	AuthorizeURL string `json:"authorize_url"`
}

// OAuthURLResp OAuth URL 响应
type OAuthURLResp struct {
	AuthURL string `json:"auth_url"`
	State   string `json:"state"`
}

// OAuthCallbackReq OAuth 回调请求
type OAuthCallbackReq struct {
	AuthCode string `form:"auth_code" binding:"required"`
	State    string `form:"state"`
}

// AdvertiserSyncResp 广告主同步响应
type AdvertiserSyncResp struct {
	AdvertiserID uint64   `json:"advertiser_id"`
	SyncFields   []string `json:"sync_fields"`
	SyncAt       string   `json:"sync_at"`
}

// FundListReq 资金流水列表请求
type FundListReq struct {
	utils.Pagination
	AdvertiserID    uint64 `form:"advertiser_id" binding:"required"`
	StartDate       string `form:"start_date"`
	EndDate         string `form:"end_date"`
	TransactionType string `form:"transaction_type"`
}

// FundListResp 资金流水响应
type FundListResp struct {
	ID              uint64  `json:"id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	BalanceBefore   float64 `json:"balance_before"`
	BalanceAfter    float64 `json:"balance_after"`
	TransactionSeq  string  `json:"transaction_seq"`
	TransactionTime string  `json:"transaction_time"`
	Remark          string  `json:"remark"`
}
