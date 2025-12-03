package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/report/dto"
	"oceanengine-backend/internal/app/report/service"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/response"
)

// ReportHandler 报告处理器
type ReportHandler struct {
	service *service.ReportService
}

// NewReportHandler 创建报告处理器
func NewReportHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *ReportHandler {
	return &ReportHandler{
		service: service.NewReportService(db, oceanCfg),
	}
}

// GetAdvertiserReport 获取广告主报告
// @Summary 获取广告主报告
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param start_date query string true "开始日期"
// @Param end_date query string true "结束日期"
// @Success 200 {object} response.Response{data=[]dto.ReportDetailResp}
// @Router /api/v1/reports/advertiser [get]
func (h *ReportHandler) GetAdvertiserReport(c *gin.Context) {
	var req dto.ReportQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetAdvertiserReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetAdvertiserSummary 获取广告主汇总报告
// @Summary 获取广告主汇总报告
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param start_date query string true "开始日期"
// @Param end_date query string true "结束日期"
// @Success 200 {object} response.Response{data=dto.ReportSummaryResp}
// @Router /api/v1/reports/advertiser/summary [get]
func (h *ReportHandler) GetAdvertiserSummary(c *gin.Context) {
	var req dto.ReportQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetAdvertiserSummary(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetCampaignReport 获取广告系列报告
// @Summary 获取广告系列报告
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param start_date query string true "开始日期"
// @Param end_date query string true "结束日期"
// @Param campaign_id query int false "广告系列ID"
// @Success 200 {object} response.Response{data=[]dto.CampaignReportResp}
// @Router /api/v1/reports/campaign [get]
func (h *ReportHandler) GetCampaignReport(c *gin.Context) {
	var req dto.ReportQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetCampaignReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetAdReport 获取广告组报告
// @Summary 获取广告组报告
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param start_date query string true "开始日期"
// @Param end_date query string true "结束日期"
// @Param campaign_id query int false "广告系列ID"
// @Param ad_id query int false "广告组ID"
// @Success 200 {object} response.Response{data=[]dto.AdReportResp}
// @Router /api/v1/reports/ad [get]
func (h *ReportHandler) GetAdReport(c *gin.Context) {
	var req dto.ReportQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetAdReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// SyncReport 同步报告数据
// @Summary 同步报告数据
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param body body dto.ReportSyncReq true "同步请求"
// @Success 200 {object} response.Response{data=dto.ReportSyncResp}
// @Router /api/v1/reports/sync [post]
func (h *ReportHandler) SyncReport(c *gin.Context) {
	var req dto.ReportSyncReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.SyncReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetExportTaskList 获取导出任务列表
// @Summary 获取导出任务列表
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int false "广告主ID"
// @Param status query string false "状态"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.ExportTaskResp}}
// @Router /api/v1/reports/exports [get]
func (h *ReportHandler) GetExportTaskList(c *gin.Context) {
	var req dto.ExportTaskListReq
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

	list, total, err := h.service.GetExportTaskList(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithList(c, list, total, req.Page, req.PageSize)
}

// CreateExportTask 创建导出任务
// @Summary 创建导出任务
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param body body dto.ExportCreateReq true "创建请求"
// @Success 200 {object} response.Response{data=dto.ExportTaskResp}
// @Router /api/v1/reports/exports [post]
func (h *ReportHandler) CreateExportTask(c *gin.Context) {
	var req dto.ExportCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.CreateExportTask(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetExportTask 获取导出任务详情
// @Summary 获取导出任务详情
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param id path int true "任务ID"
// @Success 200 {object} response.Response{data=dto.ExportTaskResp}
// @Router /api/v1/reports/exports/{id} [get]
func (h *ReportHandler) GetExportTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetExportTask(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetCreativeReport 获取创意报告
// @Summary 获取创意报告
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param start_date query string true "开始日期"
// @Param end_date query string true "结束日期"
// @Success 200 {object} response.Response{data=[]dto.CreativeReportResp}
// @Router /api/v1/reports/creative [get]
func (h *ReportHandler) GetCreativeReport(c *gin.Context) {
	var req dto.ReportQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetCreativeReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetRealtimeReport 获取实时数据
// @Summary 获取实时数据
// @Tags 数据报告
// @Accept json
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param level query string true "层级 advertiser/campaign/ad"
// @Success 200 {object} response.Response{data=[]dto.RealtimeReportResp}
// @Router /api/v1/reports/realtime [get]
func (h *ReportHandler) GetRealtimeReport(c *gin.Context) {
	var req dto.RealtimeReportReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetRealtimeReport(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}
