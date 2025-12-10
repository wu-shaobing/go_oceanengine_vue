package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONData JSON 数据类型
type JSONData []byte

// Value implements driver.Valuer interface
func (j JSONData) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

// Scan implements sql.Scanner interface
func (j *JSONData) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*j = append((*j)[0:0], v...)
		return nil
	case string:
		*j = []byte(v)
		return nil
	default:
		return errors.New("invalid type for JSONData")
	}
}

// MarshalJSON implements json.Marshaler interface
func (j JSONData) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (j *JSONData) UnmarshalJSON(data []byte) error {
	*j = append((*j)[0:0], data...)
	return nil
}

// Unmarshal 解析 JSON 数据到结构体
func (j JSONData) Unmarshal(v interface{}) error {
	if len(j) == 0 {
		return nil
	}
	return json.Unmarshal(j, v)
}

// UserSetting 用户设置表
type UserSetting struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"uniqueIndex;not null" json:"user_id"`
	Settings  JSONData  `gorm:"type:json" json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 表名
func (UserSetting) TableName() string {
	return "sys_user_setting"
}

// UserSettingData 设置数据结构
type UserSettingData struct {
	Language             string `json:"language"`               // zh-CN, zh-TW, en-US
	Timezone             string `json:"timezone"`               // Asia/Shanghai, Asia/Hong_Kong, UTC
	Theme                string `json:"theme"`                  // light, dark, auto
	NotificationsEnabled bool   `json:"notifications_enabled"`  // 系统通知开关
	EmailAlertsEnabled   bool   `json:"email_alerts_enabled"`   // 邮件通知开关
	SmsAlertsEnabled     bool   `json:"sms_alerts_enabled"`     // 短信通知开关
	AutoRefreshEnabled   bool   `json:"auto_refresh_enabled"`   // 自动刷新开关
	RefreshInterval      int    `json:"refresh_interval"`       // 刷新间隔(秒)
}

// DefaultUserSettingData 默认设置
func DefaultUserSettingData() *UserSettingData {
	return &UserSettingData{
		Language:             "zh-CN",
		Timezone:             "Asia/Shanghai",
		Theme:                "light",
		NotificationsEnabled: true,
		EmailAlertsEnabled:   true,
		SmsAlertsEnabled:     false,
		AutoRefreshEnabled:   true,
		RefreshInterval:      30,
	}
}

// Notification 系统通知表
type Notification struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"index;not null" json:"user_id"` // 0 表示全局通知
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Type      string    `gorm:"size:32;index" json:"type"` // success, warning, error, info
	IsRead    bool      `gorm:"default:false;index" json:"is_read"`
	Link      string    `gorm:"size:512" json:"link"` // 可选的跳转链接
	CreatedAt time.Time `gorm:"index" json:"created_at"`
}

// TableName 表名
func (Notification) TableName() string {
	return "sys_notification"
}

// 通知类型常量
const (
	NotificationTypeSuccess = "success"
	NotificationTypeWarning = "warning"
	NotificationTypeError   = "error"
	NotificationTypeInfo    = "info"
)
