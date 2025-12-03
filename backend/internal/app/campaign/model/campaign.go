package model

import (
	"time"

	"gorm.io/gorm"
)

// Campaign 广告系列模型
type Campaign struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID  uint64         `gorm:"index;not null" json:"advertiser_id"`
	CampaignID    uint64         `gorm:"uniqueIndex;not null" json:"campaign_id"`
	Name          string         `gorm:"size:255;not null" json:"name"`
	Budget        float64        `gorm:"type:decimal(14,2);default:0" json:"budget"`
	BudgetMode    string         `gorm:"size:32;default:'BUDGET_MODE_INFINITE'" json:"budget_mode"`
	MarketingGoal string         `gorm:"size:64" json:"marketing_goal"`
	DeliveryMode  string         `gorm:"size:32" json:"delivery_mode"`
	Status        string         `gorm:"size:32;default:'ENABLE'" json:"status"`
	LandingType   string         `gorm:"size:32" json:"landing_type"`
	StartTime     *time.Time     `gorm:"type:datetime" json:"start_time"`
	EndTime       *time.Time     `gorm:"type:datetime" json:"end_time"`
	Remark        string         `gorm:"size:500" json:"remark"`
	LastSyncAt    *time.Time     `gorm:"type:datetime" json:"last_sync_at"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Campaign) TableName() string {
	return "ad_campaign"
}

// 预算模式常量
const (
	BudgetModeInfinite = "BUDGET_MODE_INFINITE" // 不限预算
	BudgetModeDay      = "BUDGET_MODE_DAY"      // 日预算
	BudgetModeTotal    = "BUDGET_MODE_TOTAL"    // 总预算
)

// 投放状态常量
const (
	CampaignStatusEnable  = "ENABLE"  // 启用
	CampaignStatusDisable = "DISABLE" // 暂停
	CampaignStatusDelete  = "DELETE"  // 删除
	CampaignStatusAudit   = "AUDIT"   // 审核中
	CampaignStatusReject  = "REJECT"  // 审核拒绝
)

// 营销目标常量
const (
	MarketingGoalLive        = "LIVE"        // 直播
	MarketingGoalVideo       = "VIDEO"       // 视频
	MarketingGoalAPP         = "APP"         // 应用推广
	MarketingGoalDPA         = "DPA"         // 商品推广
	MarketingGoalWebsite     = "WEBSITE"     // 网站转化
	MarketingGoalProducts    = "PRODUCTS"    // 商品
	MarketingGoalConversions = "CONVERSIONS" // 线索收集
)
