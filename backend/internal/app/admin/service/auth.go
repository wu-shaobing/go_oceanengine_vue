package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/auth"
	"oceanengine-backend/pkg/errcode"
)

// AuthService 认证服务
type AuthService struct {
	db         *gorm.DB
	jwtManager *auth.JWTManager
}

// NewAuthService 创建认证服务
func NewAuthService(db *gorm.DB, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{
		db:         db,
		jwtManager: jwtManager,
	}
}

// Login 用户登录
func (s *AuthService) Login(ctx context.Context, req *dto.LoginReq, clientIP string) (*dto.LoginResp, error) {
	// 1. 获取用户
	var user model.User
	err := s.db.WithContext(ctx).
		Preload("Role").
		Where("username = ?", req.Username).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcode.New(errcode.ErrPasswordWrong)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 2. 检查用户状态
	if user.Status == model.UserStatusDisabled {
		return nil, errcode.New(errcode.ErrAccountDisabled)
	}
	if user.Status == model.UserStatusLocked {
		return nil, errcode.New(errcode.ErrAccountLocked)
	}

	// 3. 验证密码
	if !auth.VerifyPassword(req.Password, user.Password) {
		return nil, errcode.New(errcode.ErrPasswordWrong)
	}

	// 4. 获取角色信息
	var roleKey string
	var dataScope string
	if user.Role != nil {
		roleKey = user.Role.Key
		dataScope = string(rune('0' + user.Role.DataScope))
	}

	// 5. 生成 Token
	claims := &auth.Claims{
		UserID:    int64(user.ID),
		Username:  user.Username,
		RoleKey:   roleKey,
		RoleID:    int64(user.RoleID),
		DataScope: dataScope,
	}

	accessToken, err := s.jwtManager.GenerateToken(claims)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(int64(user.ID))
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 6. 更新登录信息
	now := time.Now()
	s.db.WithContext(ctx).Model(&user).Updates(map[string]interface{}{
		"last_login_at": now,
		"last_login_ip": clientIP,
	})

	return &dto.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.jwtManager.GetAccessExpire().Seconds()),
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Roles:    []string{roleKey},
		},
	}, nil
}

// RefreshToken 刷新 Token
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*dto.LoginResp, error) {
	// 1. 解析刷新 Token
	claims, err := s.jwtManager.ParseToken(refreshToken)
	if err != nil {
		return nil, errcode.New(errcode.ErrRefreshTokenInvalid)
	}

	// 2. 获取最新用户信息
	var user model.User
	err = s.db.WithContext(ctx).
		Preload("Role").
		First(&user, claims.UserID).Error
	if err != nil {
		return nil, errcode.New(errcode.ErrUserNotFound)
	}

	// 3. 检查用户状态
	if user.Status != model.UserStatusEnabled {
		return nil, errcode.New(errcode.ErrAccountDisabled)
	}

	// 4. 获取角色信息
	var roleKey string
	var dataScope string
	if user.Role != nil {
		roleKey = user.Role.Key
		dataScope = string(rune('0' + user.Role.DataScope))
	}

	// 5. 生成新 Token
	newClaims := &auth.Claims{
		UserID:    int64(user.ID),
		Username:  user.Username,
		RoleKey:   roleKey,
		RoleID:    int64(user.RoleID),
		DataScope: dataScope,
	}

	accessToken, err := s.jwtManager.GenerateToken(newClaims)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	newRefreshToken, err := s.jwtManager.GenerateRefreshToken(int64(user.ID))
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(s.jwtManager.GetAccessExpire().Seconds()),
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Roles:    []string{roleKey},
		},
	}, nil
}

// GetUserInfo 获取当前用户信息
func (s *AuthService) GetUserInfo(ctx context.Context, userID int64) (*dto.UserInfo, error) {
	var user model.User
	err := s.db.WithContext(ctx).
		Preload("Role").
		First(&user, userID).Error
	if err != nil {
		return nil, errcode.New(errcode.ErrUserNotFound)
	}

	var roleKey string
	if user.Role != nil {
		roleKey = user.Role.Key
	}

	return &dto.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Roles:    []string{roleKey},
	}, nil
}
