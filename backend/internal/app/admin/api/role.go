package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// RoleAPI 角色管理API
type RoleAPI struct {
	roleService *service.RoleService
}

// NewRoleAPI 创建角色管理API
func NewRoleAPI(roleService *service.RoleService) *RoleAPI {
	return &RoleAPI{roleService: roleService}
}

// GetList godoc
// @Summary 获取角色列表
// @Tags 系统管理-角色
// @Produce json
// @Param name query string false "角色名称"
// @Param code query string false "角色编码"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]model.Role}
// @Router /api/v1/system/roles [get]
func (a *RoleAPI) GetList(c *gin.Context) {
	var req dto.RoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.roleService.GetList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetAll godoc
// @Summary 获取所有角色
// @Tags 系统管理-角色
// @Produce json
// @Success 200 {object} response.Response{data=[]model.Role}
// @Router /api/v1/system/roles/all [get]
func (a *RoleAPI) GetAll(c *gin.Context) {
	list, err := a.roleService.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, list)
}

// GetByID godoc
// @Summary 获取角色详情
// @Tags 系统管理-角色
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} response.Response{data=model.Role}
// @Router /api/v1/system/roles/{id} [get]
func (a *RoleAPI) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	role, err := a.roleService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, role)
}

// Create godoc
// @Summary 创建角色
// @Tags 系统管理-角色
// @Accept json
// @Produce json
// @Param data body dto.RoleCreateReq true "角色信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/roles [post]
func (a *RoleAPI) Create(c *gin.Context) {
	var req dto.RoleCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	operatorID := c.GetUint64("userID")
	if err := a.roleService.Create(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Update godoc
// @Summary 更新角色
// @Tags 系统管理-角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleUpdateReq true "角色信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/roles/{id} [put]
func (a *RoleAPI) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.RoleUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	operatorID := c.GetUint64("userID")
	if err := a.roleService.Update(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Delete godoc
// @Summary 删除角色
// @Tags 系统管理-角色
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/roles/{id} [delete]
func (a *RoleAPI) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.roleService.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// GetRoleMenus godoc
// @Summary 获取角色菜单
// @Tags 系统管理-角色
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} response.Response{data=[]uint64}
// @Router /api/v1/system/roles/{id}/menus [get]
func (a *RoleAPI) GetRoleMenus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	menuIDs, err := a.roleService.GetRoleMenus(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, menuIDs)
}

// UpdateRoleMenus godoc
// @Summary 更新角色菜单
// @Tags 系统管理-角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleMenuUpdateReq true "菜单ID列表"
// @Success 200 {object} response.Response
// @Router /api/v1/system/roles/{id}/menus [put]
func (a *RoleAPI) UpdateRoleMenus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.RoleMenuUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.roleService.UpdateRoleMenus(c.Request.Context(), id, req.MenuIDs); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
