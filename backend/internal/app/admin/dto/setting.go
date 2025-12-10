package dto

import "oceanengine-backend/pkg/utils"

// ==================== 用户设置 ====================

// UserSettingResp 用户设置响应
type UserSettingResp struct {
	Language             string `json:"language"`
	Timezone             string `json:"timezone"`
	Theme                string `json:"theme"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	EmailAlertsEnabled   bool   `json:"email_alerts_enabled"`
	SmsAlertsEnabled     bool   `json:"sms_alerts_enabled"`
	AutoRefreshEnabled   bool   `json:"auto_refresh_enabled"`
	RefreshInterval      int    `json:"refresh_interval"`
}

// UserSettingUpdateReq 更新用户设置请求
type UserSettingUpdateReq struct {
	Language             *string `json:"language" binding:"omitempty,oneof=zh-CN zh-TW en-US"`
	Timezone             *string `json:"timezone" binding:"omitempty,oneof=Asia/Shanghai Asia/Hong_Kong UTC"`
	Theme                *string `json:"theme" binding:"omitempty,oneof=light dark auto"`
	NotificationsEnabled *bool   `json:"notifications_enabled"`
	EmailAlertsEnabled   *bool   `json:"email_alerts_enabled"`
	SmsAlertsEnabled     *bool   `json:"sms_alerts_enabled"`
	AutoRefreshEnabled   *bool   `json:"auto_refresh_enabled"`
	RefreshInterval      *int    `json:"refresh_interval" binding:"omitempty,min=10,max=300"`
}

// ==================== 消息通知 ====================

// NotificationListReq 通知列表请求
type NotificationListReq struct {
	utils.Pagination
	Type    string `form:"type"`     // success, warning, error, info
	IsRead  *bool  `form:"is_read"`  // 是否已读
	Keyword string `form:"keyword"`  // 标题/内容关键词
}

// NotificationResp 通知响应
type NotificationResp struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	IsRead    bool   `json:"is_read"`
	Link      string `json:"link"`
	CreatedAt string `json:"created_at"`
}

// NotificationStatsResp 通知统计响应
type NotificationStatsResp struct {
	Total     int64 `json:"total"`
	Unread    int64 `json:"unread"`
	TodayNew  int64 `json:"today_new"`
	Important int64 `json:"important"` // warning + error 类型数量
}

// NotificationMarkReadReq 标记已读请求
type NotificationMarkReadReq struct {
	IDs []uint64 `json:"ids" binding:"required,min=1"`
}

// NotificationCreateReq 创建通知请求（内部使用）
type NotificationCreateReq struct {
	UserID  uint64 `json:"user_id"` // 0 表示全局通知
	Title   string `json:"title" binding:"required,max=255"`
	Content string `json:"content"`
	Type    string `json:"type" binding:"required,oneof=success warning error info"`
	Link    string `json:"link" binding:"omitempty,max=512"`
}
