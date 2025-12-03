package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// UserAPI 用户管理API
type UserAPI struct {
	userService *service.UserService
}

// NewUserAPI 创建用户管理API
func NewUserAPI(userService *service.UserService) *UserAPI {
	return &UserAPI{userService: userService}
}

// GetList godoc
// @Summary 获取用户列表
// @Tags 系统管理-用户
// @Produce json
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param phone query string false "手机号"
// @Param status query int false "状态"
// @Param role_id query int false "角色ID"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=dto.UserListResp}
// @Router /api/v1/system/users [get]
func (a *UserAPI) GetList(c *gin.Context) {
	var req dto.UserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.userService.GetList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetByID godoc
// @Summary 获取用户详情
// @Tags 系统管理-用户
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=model.User}
// @Router /api/v1/system/users/{id} [get]
func (a *UserAPI) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	user, err := a.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, user)
}

// Create godoc
// @Summary 创建用户
// @Tags 系统管理-用户
// @Accept json
// @Produce json
// @Param data body dto.UserCreateReq true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/users [post]
func (a *UserAPI) Create(c *gin.Context) {
	var req dto.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	operatorID := c.GetUint64("userID")
	if err := a.userService.Create(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Update godoc
// @Summary 更新用户
// @Tags 系统管理-用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param data body dto.UserUpdateReq true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/users/{id} [put]
func (a *UserAPI) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	operatorID := c.GetUint64("userID")
	if err := a.userService.Update(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Delete godoc
// @Summary 删除用户
// @Tags 系统管理-用户
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/users/{id} [delete]
func (a *UserAPI) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.userService.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// ResetPassword godoc
// @Summary 重置密码
// @Tags 系统管理-用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param data body dto.UserResetPasswordReq true "密码信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/users/{id}/reset-password [post]
func (a *UserAPI) ResetPassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.UserResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	if err := a.userService.ResetPassword(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// ChangePassword godoc
// @Summary 修改密码
// @Tags 系统管理-用户
// @Accept json
// @Produce json
// @Param data body dto.UserChangePasswordReq true "密码信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/users/change-password [post]
func (a *UserAPI) ChangePassword(c *gin.Context) {
	var req dto.UserChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userID := c.GetUint64("userID")
	if err := a.userService.ChangePassword(c.Request.Context(), userID, &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
