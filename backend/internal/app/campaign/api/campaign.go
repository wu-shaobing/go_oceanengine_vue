package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/campaign/dto"
	"oceanengine-backend/internal/app/campaign/service"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/response"
)

// CampaignHandler 广告系列处理器
type CampaignHandler struct {
	service *service.CampaignService
}

// NewCampaignHandler 创建广告系列处理器
func NewCampaignHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *CampaignHandler {
	return &CampaignHandler{
		service: service.NewCampaignService(db, oceanCfg),
	}
}

// List 获取广告系列列表
// @Summary 获取广告系列列表
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param advertiser_id query int false "广告主ID"
// @Param name query string false "系列名称"
// @Param status query string false "状态"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.CampaignListResp}}
// @Router /api/v1/campaigns [get]
func (h *CampaignHandler) List(c *gin.Context) {
	var req dto.CampaignListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	// 设置默认分页
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

// Get 获取广告系列详情
// @Summary 获取广告系列详情
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param id path int true "系列ID"
// @Success 200 {object} response.Response{data=dto.CampaignDetailResp}
// @Router /api/v1/campaigns/{id} [get]
func (h *CampaignHandler) Get(c *gin.Context) {
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

// Create 创建广告系列
// @Summary 创建广告系列
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param body body dto.CampaignCreateReq true "创建内容"
// @Success 200 {object} response.Response{data=dto.CampaignDetailResp}
// @Router /api/v1/campaigns [post]
func (h *CampaignHandler) Create(c *gin.Context) {
	var req dto.CampaignCreateReq
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

// Update 更新广告系列
// @Summary 更新广告系列
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param id path int true "系列ID"
// @Param body body dto.CampaignUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/campaigns/{id} [put]
func (h *CampaignHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	var req dto.CampaignUpdateReq
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
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param body body dto.CampaignStatusUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/campaigns/status [put]
func (h *CampaignHandler) UpdateStatus(c *gin.Context) {
	var req dto.CampaignStatusUpdateReq
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

// Delete 删除广告系列
// @Summary 删除广告系列
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param id path int true "系列ID"
// @Success 200 {object} response.Response
// @Router /api/v1/campaigns/{id} [delete]
func (h *CampaignHandler) Delete(c *gin.Context) {
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

// Sync 同步广告系列
// @Summary 同步广告系列
// @Tags 广告系列管理
// @Accept json
// @Produce json
// @Param advertiser_id path int true "广告主ID"
// @Success 200 {object} response.Response{data=dto.CampaignSyncResp}
// @Router /api/v1/campaigns/sync/{advertiser_id} [post]
func (h *CampaignHandler) Sync(c *gin.Context) {
	advertiserID, err := strconv.ParseUint(c.Param("advertiser_id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.Sync(c.Request.Context(), advertiserID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}
