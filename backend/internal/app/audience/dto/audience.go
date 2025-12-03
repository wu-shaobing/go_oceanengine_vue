package dto

import "oceanengine-backend/pkg/utils"

// AudiencePackageListReq 定向包列表请求
type AudiencePackageListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	Keyword      string `form:"keyword"`
	LandingType  string `form:"landing_type"`
}

// AudiencePackageListResp 定向包列表响应项
type AudiencePackageListResp struct {
	ID            uint64 `json:"id"`
	PackageID     uint64 `json:"package_id"`
	AdvertiserID  uint64 `json:"advertiser_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	LandingType   string `json:"landing_type"`
	DeliveryRange string `json:"delivery_range"`
	CreatedAt     string `json:"created_at"`
}

// AudiencePackageCreateReq 创建定向包请求
type AudiencePackageCreateReq struct {
	AdvertiserID  uint64                 `json:"advertiser_id" binding:"required"`
	Name          string                 `json:"name" binding:"required,max=255"`
	Description   string                 `json:"description" binding:"max=500"`
	LandingType   string                 `json:"landing_type" binding:"required"`
	DeliveryRange string                 `json:"delivery_range"`
	Audience      map[string]interface{} `json:"audience"` // 定向设置
}

// AudiencePackageUpdateReq 更新定向包请求
type AudiencePackageUpdateReq struct {
	ID            uint64                 `json:"id"`
	Name          string                 `json:"name" binding:"max=255"`
	Description   string                 `json:"description" binding:"max=500"`
	DeliveryRange string                 `json:"delivery_range"`
	Audience      map[string]interface{} `json:"audience"`
}

// CustomAudienceListReq 自定义人群列表请求
type CustomAudienceListReq struct {
	utils.Pagination
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	Keyword      string `form:"keyword"`
	Source       string `form:"source"`
	Status       *int   `form:"status"`
}

// CustomAudienceListResp 自定义人群列表响应项
type CustomAudienceListResp struct {
	ID               uint64 `json:"id"`
	CustomAudienceID uint64 `json:"custom_audience_id"`
	AdvertiserID     uint64 `json:"advertiser_id"`
	Name             string `json:"name"`
	Source           string `json:"source"`
	Status           int    `json:"status"`
	CoverNum         int64  `json:"cover_num"`
	Tag              string `json:"tag"`
	PushStatus       int    `json:"push_status"`
	CreatedAt        string `json:"created_at"`
}

// CustomAudienceCreateReq 创建自定义人群请求
type CustomAudienceCreateReq struct {
	AdvertiserID   uint64   `json:"advertiser_id" binding:"required"`
	Name           string   `json:"name" binding:"required,max=255"`
	Source         string   `json:"source" binding:"required"`
	Tag            string   `json:"tag" binding:"max=100"`
	DataFileIDs    []string `json:"data_file_ids"`   // 数据文件ID列表
	RetargetingTag string   `json:"retargeting_tag"` // 重定向标签
}

// CustomAudiencePushReq 推送人群请求
type CustomAudiencePushReq struct {
	CustomAudienceID uint64   `json:"custom_audience_id" binding:"required"`
	TargetAdIDs      []uint64 `json:"target_ad_ids"`
}
