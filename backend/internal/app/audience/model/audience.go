package model

import (
	"time"

	"gorm.io/gorm"
)

// AudiencePackage 定向包表
type AudiencePackage struct {
	ID            uint64         `gorm:"primaryKey" json:"id"`
	PackageID     uint64         `gorm:"uniqueIndex" json:"package_id"` // Ocean Engine 定向包ID
	AdvertiserID  uint64         `gorm:"index" json:"advertiser_id"`    // 广告主ID
	Name          string         `gorm:"size:255" json:"name"`          // 定向包名称
	Description   string         `gorm:"size:500" json:"description"`   // 描述
	Status        string         `gorm:"size:50" json:"status"`         // 状态
	LandingType   string         `gorm:"size:50" json:"landing_type"`   // 推广类型
	Audience      string         `gorm:"type:text" json:"audience"`     // 定向设置（JSON）
	DeliveryRange string         `gorm:"size:50" json:"delivery_range"` // 投放范围
	ModifyTime    *time.Time     `json:"modify_time"`
	CreateTime    *time.Time     `json:"create_time"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (AudiencePackage) TableName() string {
	return "ad_audience_package"
}

// CustomAudience 自定义人群表
type CustomAudience struct {
	ID               uint64         `gorm:"primaryKey" json:"id"`
	CustomAudienceID uint64         `gorm:"uniqueIndex" json:"custom_audience_id"` // Ocean Engine 人群包ID
	AdvertiserID     uint64         `gorm:"index" json:"advertiser_id"`            // 广告主ID
	Name             string         `gorm:"size:255" json:"name"`                  // 人群包名称
	Source           string         `gorm:"size:50" json:"source"`                 // 来源类型
	Status           int            `gorm:"default:0" json:"status"`               // 状态
	CoverNum         int64          `gorm:"default:0" json:"cover_num"`            // 覆盖人数
	Tag              string         `gorm:"size:100" json:"tag"`                   // 标签
	PushStatus       int            `gorm:"default:0" json:"push_status"`          // 推送状态
	ModifyTime       *time.Time     `json:"modify_time"`
	CreateTime       *time.Time     `json:"create_time"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (CustomAudience) TableName() string {
	return "ad_custom_audience"
}

// 定向包状态
const (
	AudiencePackageStatusEnable  = "AUDIENCE_PACKAGE_STATUS_ENABLE"
	AudiencePackageStatusDisable = "AUDIENCE_PACKAGE_STATUS_DISABLE"
)

// 自定义人群状态
const (
	CustomAudienceStatusProcessing = 0 // 处理中
	CustomAudienceStatusReady      = 1 // 已完成
	CustomAudienceStatusFailed     = 2 // 失败
)

// 人群来源类型
const (
	AudienceSourceDMP       = "DMP"       // DMP人群
	AudienceSourcePixel     = "PIXEL"     // 像素人群
	AudienceSourceLookalike = "LOOKALIKE" // 相似人群
)
