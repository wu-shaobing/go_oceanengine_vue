package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// NotificationService 消息通知服务
type NotificationService struct {
	db *gorm.DB
}

// NewNotificationService 创建消息通知服务
func NewNotificationService(db *gorm.DB) *NotificationService {
	return &NotificationService{db: db}
}

// GetList 获取通知列表
func (s *NotificationService) GetList(ctx context.Context, userID uint64, req *dto.NotificationListReq) ([]*dto.NotificationResp, int64, error) {
	var notifications []*model.Notification
	var total int64

	query := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? OR user_id = 0", userID) // 用户自己的通知 + 全局通知

	// 筛选条件
	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}
	if req.IsRead != nil {
		query = query.Where("is_read = ?", *req.IsRead)
	}
	if req.Keyword != "" {
		query = query.Where("(title LIKE ? OR content LIKE ?)", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 分页查询
	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Order("created_at DESC").Offset(offset).Limit(req.GetPageSize()).Find(&notifications).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 转换响应
	result := make([]*dto.NotificationResp, len(notifications))
	for i, n := range notifications {
		result[i] = &dto.NotificationResp{
			ID:        n.ID,
			Title:     n.Title,
			Content:   n.Content,
			Type:      n.Type,
			IsRead:    n.IsRead,
			Link:      n.Link,
			CreatedAt: n.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetStats 获取通知统计
func (s *NotificationService) GetStats(ctx context.Context, userID uint64) (*dto.NotificationStatsResp, error) {
	var stats dto.NotificationStatsResp

	baseQuery := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? OR user_id = 0", userID)

	// 总数
	if err := baseQuery.Count(&stats.Total).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 未读数
	if err := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("(user_id = ? OR user_id = 0) AND is_read = ?", userID, false).
		Count(&stats.Unread).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 今日新增
	today := time.Now().Truncate(24 * time.Hour)
	if err := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("(user_id = ? OR user_id = 0) AND created_at >= ?", userID, today).
		Count(&stats.TodayNew).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 重要通知数 (warning + error)
	if err := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("(user_id = ? OR user_id = 0) AND type IN (?, ?) AND is_read = ?",
			userID, model.NotificationTypeWarning, model.NotificationTypeError, false).
		Count(&stats.Important).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &stats, nil
}

// MarkAsRead 标记为已读
func (s *NotificationService) MarkAsRead(ctx context.Context, userID uint64, ids []uint64) error {
	result := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("id IN ? AND (user_id = ? OR user_id = 0)", ids, userID).
		Update("is_read", true)

	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}

	return nil
}

// MarkAllAsRead 标记全部已读
func (s *NotificationService) MarkAllAsRead(ctx context.Context, userID uint64) error {
	result := s.db.WithContext(ctx).Model(&model.Notification{}).
		Where("(user_id = ? OR user_id = 0) AND is_read = ?", userID, false).
		Update("is_read", true)

	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}

	return nil
}

// Delete 删除通知
func (s *NotificationService) Delete(ctx context.Context, userID uint64, ids []uint64) error {
	// 只能删除自己的通知（user_id = userID），不能删除全局通知
	result := s.db.WithContext(ctx).
		Where("id IN ? AND user_id = ?", ids, userID).
		Delete(&model.Notification{})

	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}

	return nil
}

// Create 创建通知（内部使用）
func (s *NotificationService) Create(ctx context.Context, req *dto.NotificationCreateReq) error {
	notification := &model.Notification{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Link:    req.Link,
		IsRead:  false,
	}

	if err := s.db.WithContext(ctx).Create(notification).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// CreateBatch 批量创建通知（内部使用）
func (s *NotificationService) CreateBatch(ctx context.Context, notifications []*model.Notification) error {
	if len(notifications) == 0 {
		return nil
	}

	if err := s.db.WithContext(ctx).Create(&notifications).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}
