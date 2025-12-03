package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/internal/app/creative/dto"
	"oceanengine-backend/internal/app/creative/service"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/response"
)

// CreativeHandler 创意处理器
type CreativeHandler struct {
	service *service.CreativeService
}

// NewCreativeHandler 创建创意处理器
func NewCreativeHandler(db *gorm.DB) *CreativeHandler {
	return &CreativeHandler{
		service: service.NewCreativeService(db),
	}
}

// List 获取创意列表
// @Summary 获取创意列表
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param advertiser_id query int false "广告主ID"
// @Param ad_id query int false "广告组ID"
// @Param title query string false "创意标题"
// @Param status query string false "状态"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.CreativeListResp}}
// @Router /api/v1/creatives [get]
func (h *CreativeHandler) List(c *gin.Context) {
	var req dto.CreativeListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	list, total, err := h.service.GetList(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithList(c, list, total, req.Page, req.PageSize)
}

// Get 获取创意详情
// @Summary 获取创意详情
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param id path int true "创意ID"
// @Success 200 {object} response.Response{data=dto.CreativeDetailResp}
// @Router /api/v1/creatives/{id} [get]
func (h *CreativeHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// Update 更新创意
// @Summary 更新创意
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param id path int true "创意ID"
// @Param body body dto.CreativeUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/creatives/{id} [put]
func (h *CreativeHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	var req dto.CreativeUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if err := h.service.Update(c.Request.Context(), id, &req); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c)
}

// UpdateStatus 批量更新状态
// @Summary 批量更新状态
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param body body dto.CreativeStatusUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/creatives/status [put]
func (h *CreativeHandler) UpdateStatus(c *gin.Context) {
	var req dto.CreativeStatusUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if err := h.service.UpdateStatus(c.Request.Context(), &req); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c)
}

// Create 创建创意
// @Summary 创建创意
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param body body dto.CreativeCreateReq true "创意信息"
// @Success 200 {object} response.Response{data=map[string]interface{}}
// @Router /api/v1/creatives [post]
func (h *CreativeHandler) Create(c *gin.Context) {
	var req dto.CreativeCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// Delete 删除创意
// @Summary 删除创意
// @Tags 创意管理
// @Accept json
// @Produce json
// @Param id path int true "创意ID"
// @Success 200 {object} response.Response
// @Router /api/v1/creatives/{id} [delete]
func (h *CreativeHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c)
}
