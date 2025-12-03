package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/auth"
	"oceanengine-backend/pkg/errcode"
)

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetList 获取用户列表
func (s *UserService) GetList(ctx context.Context, req *dto.UserListReq) ([]*dto.UserListResp, int64, error) {
	var users []*model.User
	var total int64

	query := s.db.WithContext(ctx).Model(&model.User{}).Preload("Role")

	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Nickname != "" {
		query = query.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.RoleID > 0 {
		query = query.Where("role_id = ?", req.RoleID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&users).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.UserListResp, len(users))
	for i, u := range users {
		var lastLoginAt string
		if u.LastLoginAt != nil {
			lastLoginAt = u.LastLoginAt.Format("2006-01-02 15:04:05")
		}
		var roleName string
		if u.Role != nil {
			roleName = u.Role.Name
		}
		result[i] = &dto.UserListResp{
			ID:          u.ID,
			Username:    u.Username,
			Nickname:    u.Nickname,
			Phone:       u.Phone,
			Email:       u.Email,
			Avatar:      u.Avatar,
			Status:      u.Status,
			RoleID:      u.RoleID,
			RoleName:    roleName,
			LastLoginAt: lastLoginAt,
			CreatedAt:   u.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetByID 获取用户详情
func (s *UserService) GetByID(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	if err := s.db.WithContext(ctx).Preload("Role").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrUserNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &user, nil
}

// Create 创建用户
func (s *UserService) Create(ctx context.Context, req *dto.UserCreateReq, operatorID uint64) error {
	// 检查用户名是否已存在
	var count int64
	if err := s.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	if count > 0 {
		return errcode.New(errcode.ErrUserExists)
	}

	// 加密密码
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	user := &model.User{
		Username:  req.Username,
		Password:  hashedPassword,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Email:     req.Email,
		Avatar:    req.Avatar,
		Status:    req.Status,
		RoleID:    req.RoleID,
		Remark:    req.Remark,
		CreatedBy: operatorID,
		UpdatedBy: operatorID,
	}

	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Update 更新用户
func (s *UserService) Update(ctx context.Context, req *dto.UserUpdateReq, operatorID uint64) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrUserNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	updates := map[string]interface{}{
		"updated_by": operatorID,
	}
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.RoleID > 0 {
		updates["role_id"] = req.RoleID
	}
	updates["status"] = req.Status
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := s.db.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Delete 删除用户
func (s *UserService) Delete(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.User{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrUserNotFound)
	}
	return nil
}

// ResetPassword 重置密码
func (s *UserService) ResetPassword(ctx context.Context, req *dto.UserResetPasswordReq) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrUserNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.db.WithContext(ctx).Model(&user).Update("password", hashedPassword).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(ctx context.Context, userID uint64, req *dto.UserChangePasswordReq) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrUserNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 验证旧密码
	if !auth.VerifyPassword(req.OldPassword, user.Password) {
		return errcode.New(errcode.ErrPasswordWrong)
	}

	// 加密新密码
	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if err := s.db.WithContext(ctx).Model(&user).Update("password", hashedPassword).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}
