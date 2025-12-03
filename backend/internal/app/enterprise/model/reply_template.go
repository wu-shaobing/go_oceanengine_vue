package model

import (
	"time"
)

// ReplyTemplate 快捷回复模板
type ReplyTemplate struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	AdvertiserID uint64    `json:"advertiser_id" gorm:"index;not null"`
	AccountID    string    `json:"account_id" gorm:"index;type:varchar(64)"`
	Name         string    `json:"name" gorm:"type:varchar(100);not null"`
	Content      string    `json:"content" gorm:"type:text;not null"`
	SortOrder    int       `json:"sort_order" gorm:"default:0"`
	Status       int       `json:"status" gorm:"default:1"` // 1: active, 0: inactive
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName 表名
func (ReplyTemplate) TableName() string {
	return "enterprise_reply_templates"
}
