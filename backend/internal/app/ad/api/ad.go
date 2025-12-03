package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/ad/dto"
	"oceanengine-backend/internal/app/ad/service"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/response"
)

// AdHandler 广告组处理器
type AdHandler struct {
	service *service.AdService
}

// NewAdHandler 创建广告组处理器
func NewAdHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *AdHandler {
	return &AdHandler{
		service: service.NewAdService(db, oceanCfg),
	}
}

// List 获取广告组列表
// @Summary 获取广告组列表
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param advertiser_id query int false "广告主ID"
// @Param campaign_id query int false "系列ID"
// @Param name query string false "广告组名称"
// @Param status query string false "状态"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.AdListResp}}
// @Router /api/v1/ads [get]
func (h *AdHandler) List(c *gin.Context) {
	var req dto.AdListReq
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

// Get 获取广告组详情
// @Summary 获取广告组详情
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param id path int true "广告组ID"
// @Success 200 {object} response.Response{data=dto.AdDetailResp}
// @Router /api/v1/ads/{id} [get]
func (h *AdHandler) Get(c *gin.Context) {
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

// Update 更新广告组
// @Summary 更新广告组
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param id path int true "广告组ID"
// @Param body body dto.AdUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/ads/{id} [put]
func (h *AdHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	var req dto.AdUpdateReq
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
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param body body dto.AdStatusUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/ads/status [put]
func (h *AdHandler) UpdateStatus(c *gin.Context) {
	var req dto.AdStatusUpdateReq
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

// Create 创建广告组
// @Summary 创建广告组
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param body body dto.AdCreateReq true "广告组信息"
// @Success 200 {object} response.Response{data=map[string]interface{}}
// @Router /api/v1/ads [post]
func (h *AdHandler) Create(c *gin.Context) {
	var req dto.AdCreateReq
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

// Delete 删除广告组
// @Summary 删除广告组
// @Tags 广告组管理
// @Accept json
// @Produce json
// @Param id path int true "广告组ID"
// @Success 200 {object} response.Response
// @Router /api/v1/ads/{id} [delete]
func (h *AdHandler) Delete(c *gin.Context) {
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
