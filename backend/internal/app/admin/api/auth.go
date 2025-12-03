package api

import (
	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/internal/middleware"
	"oceanengine-backend/pkg/response"
)

// AuthAPI 认证 API
type AuthAPI struct {
	authService *service.AuthService
}

// NewAuthAPI 创建认证 API
func NewAuthAPI(authService *service.AuthService) *AuthAPI {
	return &AuthAPI{authService: authService}
}

// Login 登录
// @Summary 用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body dto.LoginReq true "登录信息"
// @Success 200 {object} response.Response{data=dto.LoginResp}
// @Router /api/v1/auth/login [post]
func (a *AuthAPI) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := a.authService.Login(c.Request.Context(), &req, c.ClientIP())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, resp)
}

// RefreshToken 刷新 Token
// @Summary 刷新 Token
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body dto.RefreshTokenReq true "刷新 Token"
// @Success 200 {object} response.Response{data=dto.LoginResp}
// @Router /api/v1/auth/refresh [post]
func (a *AuthAPI) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := a.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, resp)
}

// GetUserInfo 获取当前用户信息
// @Summary 获取当前用户信息
// @Tags 认证
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{data=dto.UserInfo}
// @Router /api/v1/auth/userinfo [get]
func (a *AuthAPI) GetUserInfo(c *gin.Context) {
	userID := middleware.GetUserID(c)

	info, err := a.authService.GetUserInfo(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, info)
}

// Logout 退出登录
// @Summary 退出登录
// @Tags 认证
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response
// @Router /api/v1/auth/logout [post]
func (a *AuthAPI) Logout(c *gin.Context) {
	// 可以在这里实现 token 黑名单等逻辑
	response.Success(c, nil)
}
