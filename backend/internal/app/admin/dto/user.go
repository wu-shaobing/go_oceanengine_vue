package dto

import "oceanengine-backend/pkg/utils"

// LoginReq 登录请求
type LoginReq struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaID   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}

// LoginResp 登录响应
type LoginResp struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	User         *UserInfo `json:"user"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID       uint64   `json:"id"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
}

// RefreshTokenReq 刷新 Token 请求
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserListReq 用户列表请求
type UserListReq struct {
	utils.Pagination
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Phone    string `form:"phone"`
	Status   *int8  `form:"status"`
	RoleID   uint64 `form:"role_id"`
}

// UserListResp 用户列表响应项
type UserListResp struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Status      int8   `json:"status"`
	RoleID      uint64 `json:"role_id"`
	RoleName    string `json:"role_name"`
	LastLoginAt string `json:"last_login_at"`
	CreatedAt   string `json:"created_at"`
}

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Nickname string `json:"nickname" binding:"max=128"`
	Phone    string `json:"phone" binding:"max=20"`
	Email    string `json:"email" binding:"omitempty,email,max=128"`
	Avatar   string `json:"avatar" binding:"max=255"`
	Status   int8   `json:"status"`
	RoleID   uint64 `json:"role_id" binding:"required"`
	Remark   string `json:"remark" binding:"max=500"`
}

// UserUpdateReq 更新用户请求
type UserUpdateReq struct {
	ID       uint64 `json:"id" binding:"required"`
	Nickname string `json:"nickname" binding:"max=128"`
	Phone    string `json:"phone" binding:"max=20"`
	Email    string `json:"email" binding:"omitempty,email,max=128"`
	Avatar   string `json:"avatar" binding:"max=255"`
	Status   int8   `json:"status"`
	RoleID   uint64 `json:"role_id"`
	Remark   string `json:"remark" binding:"max=500"`
}

// UserChangePasswordReq 修改密码请求
type UserChangePasswordReq struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=32"`
}

// UserResetPasswordReq 重置密码请求
type UserResetPasswordReq struct {
	ID       uint64 `json:"id" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// RoleListReq 角色列表请求
type RoleListReq struct {
	utils.Pagination
	Name   string `form:"name"`
	Code   string `form:"code"`
	Status *int8  `form:"status"`
}

// RoleCreateReq 创建角色请求
type RoleCreateReq struct {
	Name   string `json:"name" binding:"required,max=64"`
	Code   string `json:"code" binding:"required,max=64"`
	Sort   int    `json:"sort"`
	Status int8   `json:"status"`
	Remark string `json:"remark" binding:"max=500"`
}

// RoleUpdateReq 更新角色请求
type RoleUpdateReq struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" binding:"max=64"`
	Code   string `json:"code" binding:"max=64"`
	Sort   int    `json:"sort"`
	Status int8   `json:"status"`
	Remark string `json:"remark" binding:"max=500"`
}

// RoleMenuUpdateReq 更新角色菜单请求
type RoleMenuUpdateReq struct {
	MenuIDs []uint64 `json:"menu_ids"`
}

// MenuCreateReq 创建菜单请求
type MenuCreateReq struct {
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name" binding:"required,max=64"`
	Path       string `json:"path" binding:"max=255"`
	Component  string `json:"component" binding:"max=255"`
	Icon       string `json:"icon" binding:"max=64"`
	Sort       int    `json:"sort"`
	Type       int8   `json:"type" binding:"required"`
	Permission string `json:"permission" binding:"max=128"`
	Status     int8   `json:"status"`
	Hidden     int8   `json:"hidden"`
	Remark     string `json:"remark" binding:"max=500"`
}

// MenuUpdateReq 更新菜单请求
type MenuUpdateReq struct {
	ID         uint64 `json:"id"`
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name" binding:"max=64"`
	Path       string `json:"path" binding:"max=255"`
	Component  string `json:"component" binding:"max=255"`
	Icon       string `json:"icon" binding:"max=64"`
	Sort       int    `json:"sort"`
	Type       int8   `json:"type"`
	Permission string `json:"permission" binding:"max=128"`
	Status     int8   `json:"status"`
	Hidden     int8   `json:"hidden"`
	Remark     string `json:"remark" binding:"max=500"`
}

// MenuTree 菜单树
type MenuTree struct {
	ID         uint64      `json:"id"`
	ParentID   uint64      `json:"parent_id"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Component  string      `json:"component"`
	Icon       string      `json:"icon"`
	Sort       int         `json:"sort"`
	Type       int8        `json:"type"`
	Permission string      `json:"permission"`
	Status     int8        `json:"status"`
	Hidden     int8        `json:"hidden"`
	Children   []*MenuTree `json:"children"`
}

// OperationLogListReq 操作日志列表请求
type OperationLogListReq struct {
	utils.Pagination
	UserID    uint64 `form:"user_id"`
	Username  string `form:"username"`
	Module    string `form:"module"`
	Action    string `form:"action"`
	Status    *int   `form:"status"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

// OperationLogResp 操作日志响应
type OperationLogResp struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Module    string `json:"module"`
	Action    string `json:"action"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Request   string `json:"request"`
	Response  string `json:"response"`
	Status    int    `json:"status"`
	Duration  int64  `json:"duration"`
	CreatedAt string `json:"created_at"`
}

// OperationLogDeleteReq 删除操作日志请求
type OperationLogDeleteReq struct {
	BeforeTime string `json:"before_time" binding:"required"`
}

// CaptchaResp 验证码响应
type CaptchaResp struct {
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}
