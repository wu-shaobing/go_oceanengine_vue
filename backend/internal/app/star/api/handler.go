package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// StarHandler 星图处理器
type StarHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewStarHandler 创建星图处理器
func NewStarHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *StarHandler {
	return &StarHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 获取access token
func (h *StarHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.GetHeader("Access-Token")
	}
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 获取广告主ID
func (h *StarHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.PostForm("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// GetAccountInfo 获取星图账户信息
func (h *StarHandler) GetAccountInfo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	info, err := h.client.Star().GetAccountInfo(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, info)
}

// GetAgentAdvertisers 获取代理商广告主列表
func (h *StarHandler) GetAgentAdvertisers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	// 代理商列表通过其他接口获取
	response.OKWithList(c, []interface{}{}, 0, page, pageSize)
}

// GetBatchBalance 批量获取余额
func (h *StarHandler) GetBatchBalance(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	if accessToken == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req struct {
		AdvertiserIDs []uint64 `json:"advertiser_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, err := h.client.Star().GetFundBalance(c.Request.Context(), accessToken, req.AdvertiserIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, list)
}

// GetFundDaily 获取日流水
func (h *StarHandler) GetFundDaily(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Star().GetFundDaily(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, list)
}

// GetFundTransactions 获取流水明细
func (h *StarHandler) GetFundTransactions(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetFundTransaction(c.Request.Context(), accessToken, advertiserID, startDate, endDate, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetTaskList 获取任务列表
func (h *StarHandler) GetTaskList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetTaskList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetTaskDetail 获取任务详情
func (h *StarHandler) GetTaskDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Star().GetTaskDetail(c.Request.Context(), accessToken, advertiserID, taskID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// GetTaskItems 获取任务视频列表
func (h *StarHandler) GetTaskItems(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetTaskItemList(c.Request.Context(), accessToken, advertiserID, taskID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// UpdateTaskStatusRequest 更新任务状态请求
type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// UpdateTaskStatus 更新任务状态
func (h *StarHandler) UpdateTaskStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Param("task_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.client.Star().UpdateTaskStatus(c.Request.Context(), accessToken, advertiserID, taskID, req.Status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// GetDemandList 获取需求列表
func (h *StarHandler) GetDemandList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetDemandList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetDemandDetail 获取需求详情
func (h *StarHandler) GetDemandDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	demandID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || demandID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Star().GetDemandDetail(c.Request.Context(), accessToken, advertiserID, demandID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// GetDemandOrders 获取需求订单列表
func (h *StarHandler) GetDemandOrders(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	demandID, _ := strconv.ParseUint(c.Param("demand_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || demandID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetDemandOrders(c.Request.Context(), accessToken, advertiserID, demandID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetReportOverview 获取投后报表概览
func (h *StarHandler) GetReportOverview(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Query("task_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	report, err := h.client.Star().GetReportOverview(c.Request.Context(), accessToken, advertiserID, taskID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, report)
}

// GetReportAudience 获取受众报表
func (h *StarHandler) GetReportAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Query("task_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	audience, err := h.client.Star().GetReportAudience(c.Request.Context(), accessToken, advertiserID, taskID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, audience)
}

// GetReportDaily 获取每日趋势报表
func (h *StarHandler) GetReportDaily(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Query("task_id"), 10, 64)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Star().GetReportDaily(c.Request.Context(), accessToken, advertiserID, taskID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, list)
}

// GetClueList 获取线索列表
func (h *StarHandler) GetClueList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	taskID, _ := strconv.ParseUint(c.Query("task_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || taskID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Star().GetClueList(c.Request.Context(), accessToken, advertiserID, taskID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// UpdateClueStatusRequest 更新线索状态请求
type UpdateClueStatusRequest struct {
	Status int    `json:"status" binding:"required"`
	Remark string `json:"remark"`
}

// UpdateClueStatus 更新线索状态
func (h *StarHandler) UpdateClueStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	clueID, _ := strconv.ParseUint(c.Param("clue_id"), 10, 64)

	if accessToken == "" || clueID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req UpdateClueStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.client.Star().UpdateClueStatus(c.Request.Context(), accessToken, clueID, req.Status, req.Remark)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// ExportCluesRequest 导出线索请求
type ExportCluesRequest struct {
	TaskID    uint64 `json:"task_id" binding:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// ExportClues 导出线索
func (h *StarHandler) ExportClues(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req ExportCluesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 获取所有线索数据
	allClues := make([]interface{}, 0)
	page := 1
	pageSize := 100

	for {
		list, total, err := h.client.Star().GetClueList(c.Request.Context(), accessToken, advertiserID, req.TaskID, page, pageSize)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}

		for _, clue := range list {
			allClues = append(allClues, clue)
		}

		if len(allClues) >= total || len(list) < pageSize {
			break
		}
		page++

		// 防止无限循环
		if page > 100 {
			break
		}
	}

	response.OKWithData(c, gin.H{
		"data":    allClues,
		"total":   len(allClues),
		"message": "线索数据导出完成，请在前端进行文件下载处理",
	})
}
