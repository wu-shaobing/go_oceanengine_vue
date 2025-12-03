package service

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// OperationLogService 操作日志服务
type OperationLogService struct {
	db *gorm.DB
}

// NewOperationLogService 创建操作日志服务
func NewOperationLogService(db *gorm.DB) *OperationLogService {
	return &OperationLogService{db: db}
}

// GetList 获取操作日志列表
func (s *OperationLogService) GetList(ctx context.Context, req *dto.OperationLogListReq) ([]*dto.OperationLogResp, int64, error) {
	var logs []*model.OperationLog
	var total int64

	query := s.db.WithContext(ctx).Model(&model.OperationLog{})

	if req.UserID > 0 {
		query = query.Where("user_id = ?", req.UserID)
	}
	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Module != "" {
		query = query.Where("module = ?", req.Module)
	}
	if req.Action != "" {
		query = query.Where("action = ?", req.Action)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.StartTime != "" {
		query = query.Where("created_at >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		query = query.Where("created_at <= ?", req.EndTime)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&logs).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.OperationLogResp, len(logs))
	for i, log := range logs {
		result[i] = &dto.OperationLogResp{
			ID:        log.ID,
			UserID:    log.UserID,
			Username:  log.Username,
			Module:    log.Module,
			Action:    log.Action,
			Method:    log.Method,
			Path:      log.Path,
			IP:        log.IP,
			UserAgent: log.UserAgent,
			Request:   log.Body,
			Response:  log.Response,
			Status:    log.Status,
			Duration:  log.Latency,
			CreatedAt: log.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// Create 创建操作日志
func (s *OperationLogService) Create(ctx context.Context, log *model.OperationLog) error {
	if err := s.db.WithContext(ctx).Create(log).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// Delete 删除操作日志（按时间范围批量删除）
func (s *OperationLogService) Delete(ctx context.Context, beforeTime string) error {
	if beforeTime == "" {
		return errcode.New(errcode.ErrInvalidParams)
	}

	if err := s.db.WithContext(ctx).Where("created_at < ?", beforeTime).Delete(&model.OperationLog{}).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return nil
}

// GetModules 获取所有模块列表
func (s *OperationLogService) GetModules(ctx context.Context) ([]string, error) {
	var modules []string
	if err := s.db.WithContext(ctx).Model(&model.OperationLog{}).Distinct("module").Pluck("module", &modules).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return modules, nil
}
