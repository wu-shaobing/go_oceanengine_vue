package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// MenuAPI 菜单管理API
type MenuAPI struct {
	menuService *service.MenuService
}

// NewMenuAPI 创建菜单管理API
func NewMenuAPI(menuService *service.MenuService) *MenuAPI {
	return &MenuAPI{menuService: menuService}
}

// GetTree godoc
// @Summary 获取菜单树
// @Tags 系统管理-菜单
// @Produce json
// @Success 200 {object} response.Response{data=[]dto.MenuTree}
// @Router /api/v1/system/menus/tree [get]
func (a *MenuAPI) GetTree(c *gin.Context) {
	tree, err := a.menuService.GetTree(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, tree)
}

// GetUserMenuTree godoc
// @Summary 获取用户菜单树
// @Tags 系统管理-菜单
// @Produce json
// @Success 200 {object} response.Response{data=[]dto.MenuTree}
// @Router /api/v1/system/menus/user [get]
func (a *MenuAPI) GetUserMenuTree(c *gin.Context) {
	roleID := c.GetUint64("roleID")
	tree, err := a.menuService.GetUserMenuTree(c.Request.Context(), roleID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, tree)
}

// GetList godoc
// @Summary 获取菜单列表
// @Tags 系统管理-菜单
// @Produce json
// @Success 200 {object} response.Response{data=[]model.Menu}
// @Router /api/v1/system/menus [get]
func (a *MenuAPI) GetList(c *gin.Context) {
	list, err := a.menuService.GetList(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, list)
}

// GetByID godoc
// @Summary 获取菜单详情
// @Tags 系统管理-菜单
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} response.Response{data=model.Menu}
// @Router /api/v1/system/menus/{id} [get]
func (a *MenuAPI) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	menu, err := a.menuService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, menu)
}

// Create godoc
// @Summary 创建菜单
// @Tags 系统管理-菜单
// @Accept json
// @Produce json
// @Param data body dto.MenuCreateReq true "菜单信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menus [post]
func (a *MenuAPI) Create(c *gin.Context) {
	var req dto.MenuCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	operatorID := c.GetUint64("userID")
	if err := a.menuService.Create(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Update godoc
// @Summary 更新菜单
// @Tags 系统管理-菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param data body dto.MenuUpdateReq true "菜单信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menus/{id} [put]
func (a *MenuAPI) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.MenuUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	operatorID := c.GetUint64("userID")
	if err := a.menuService.Update(c.Request.Context(), &req, operatorID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Delete godoc
// @Summary 删除菜单
// @Tags 系统管理-菜单
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menus/{id} [delete]
func (a *MenuAPI) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.menuService.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
