package model

import (
	"time"

	"gorm.io/gorm"
)

// Creative 创意模型
type Creative struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID uint64         `gorm:"index;not null" json:"advertiser_id"`
	AdID         uint64         `gorm:"index;not null" json:"ad_id"`
	CreativeID   uint64         `gorm:"uniqueIndex;not null" json:"creative_id"`
	Title        string         `gorm:"size:255" json:"title"`
	Source       string         `gorm:"size:255" json:"source"`
	ImageMode    string         `gorm:"size:64" json:"image_mode"`
	OptStatus    string         `gorm:"size:32;default:'ENABLE'" json:"opt_status"`
	Status       string         `gorm:"size:32" json:"status"`
	ImageURL     string         `gorm:"size:512" json:"image_url"`
	VideoURL     string         `gorm:"size:512" json:"video_url"`
	ThumbURL     string         `gorm:"size:512" json:"thumb_url"`
	ActionType   string         `gorm:"size:32" json:"action_type"`
	ActionText   string         `gorm:"size:64" json:"action_text"`
	LandingURL   string         `gorm:"size:1024" json:"landing_url"`
	Remark       string         `gorm:"size:500" json:"remark"`
	LastSyncAt   *time.Time     `gorm:"type:datetime" json:"last_sync_at"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Creative) TableName() string {
	return "ad_creative"
}

// 创意状态常量
const (
	CreativeOptStatusEnable  = "ENABLE"  // 启用
	CreativeOptStatusDisable = "DISABLE" // 暂停
)

// 素材类型常量
const (
	ImageModeLargeImage    = "CREATIVE_IMAGE_MODE_LARGE"          // 大图
	ImageModeSmallImage    = "CREATIVE_IMAGE_MODE_SMALL"          // 小图
	ImageModeGroupImage    = "CREATIVE_IMAGE_MODE_GROUP"          // 组图
	ImageModeVideo         = "CREATIVE_IMAGE_MODE_VIDEO"          // 横版视频
	ImageModeVerticalVideo = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL" // 竖版视频
)
