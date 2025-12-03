package model

import (
	"time"

	"gorm.io/gorm"
)

// Ad 广告组模型
type Ad struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID   uint64         `gorm:"index;not null" json:"advertiser_id"`
	CampaignID     uint64         `gorm:"index;not null" json:"campaign_id"`
	AdID           uint64         `gorm:"uniqueIndex;not null" json:"ad_id"`
	Name           string         `gorm:"size:255;not null" json:"name"`
	OptStatus      string         `gorm:"size:32;default:'ENABLE'" json:"opt_status"`
	Status         string         `gorm:"size:32" json:"status"`
	Budget         float64        `gorm:"type:decimal(14,2);default:0" json:"budget"`
	BudgetMode     string         `gorm:"size:32" json:"budget_mode"`
	DeliveryRange  string         `gorm:"size:32" json:"delivery_range"`
	Pricing        string         `gorm:"size:32" json:"pricing"`
	Bid            float64        `gorm:"type:decimal(14,2)" json:"bid"`
	ConvertID      uint64         `gorm:"default:0" json:"convert_id"`
	StartTime      *time.Time     `gorm:"type:datetime" json:"start_time"`
	EndTime        *time.Time     `gorm:"type:datetime" json:"end_time"`
	ScheduleType   string         `gorm:"size:32" json:"schedule_type"`
	AudienceType   string         `gorm:"size:32" json:"audience_type"`
	TargetSettings string         `gorm:"type:text" json:"target_settings"` // JSON格式的定向设置
	Remark         string         `gorm:"size:500" json:"remark"`
	LastSyncAt     *time.Time     `gorm:"type:datetime" json:"last_sync_at"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Ad) TableName() string {
	return "ad_ad"
}

// 广告组状态常量
const (
	AdOptStatusEnable  = "ENABLE"  // 启用
	AdOptStatusDisable = "DISABLE" // 暂停
)

// 投放范围常量
const (
	DeliveryRangeDefault = "DEFAULT" // 默认
	DeliveryRangeUnion   = "UNION"   // 穿山甲
)

// 出价方式常量
const (
	PricingCPM  = "PRICING_CPM"  // 按展示付费
	PricingCPC  = "PRICING_CPC"  // 按点击付费
	PricingOCPM = "PRICING_OCPM" // 按转化付费
	PricingCPA  = "PRICING_CPA"  // 按行动付费
)
