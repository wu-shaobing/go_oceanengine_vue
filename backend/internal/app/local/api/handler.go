package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// LocalHandler 本地推处理器
type LocalHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewLocalHandler 创建本地推处理器
func NewLocalHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *LocalHandler {
	return &LocalHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 获取access token
func (h *LocalHandler) getAccessToken(c *gin.Context) string {
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
func (h *LocalHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.PostForm("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// GetProjectList 获取项目列表
func (h *LocalHandler) GetProjectList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	req := &oceanengine.ProjectListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.client.Local().GetProjectList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetProjectDetail 获取项目详情
func (h *LocalHandler) GetProjectDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	projectID, _ := strconv.ParseUint(c.Param("project_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Local().GetProjectDetail(c.Request.Context(), accessToken, advertiserID, projectID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// CreateProject 创建项目
func (h *LocalHandler) CreateProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var projectData map[string]interface{}
	if err := c.ShouldBindJSON(&projectData); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	projectID, err := h.client.Local().CreateProject(c.Request.Context(), accessToken, advertiserID, projectData)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"project_id": projectID})
}

// UpdateProject 更新项目
func (h *LocalHandler) UpdateProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	projectID, _ := strconv.ParseUint(c.Param("project_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	updateData["project_id"] = projectID

	err := h.client.Local().UpdateProject(c.Request.Context(), accessToken, advertiserID, projectID, updateData)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// UpdateProjectStatusRequest 更新项目状态请求
type UpdateProjectStatusRequest struct {
	ProjectIDs []uint64 `json:"project_ids" binding:"required"`
	OptStatus  string   `json:"opt_status" binding:"required"`
}

// UpdateProjectStatus 更新项目状态 (支持批量)
func (h *LocalHandler) UpdateProjectStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req UpdateProjectStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	succIDs, failIDs, err := h.client.Local().UpdateProjectStatus(c.Request.Context(), accessToken, advertiserID, req.ProjectIDs, req.OptStatus)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"success_ids": succIDs, "fail_ids": failIDs})
}

// DeleteProject 删除项目
func (h *LocalHandler) DeleteProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	projectID, _ := strconv.ParseUint(c.Param("project_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Local().DeleteProject(c.Request.Context(), accessToken, advertiserID, []uint64{projectID})
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// GetPromotionList 获取广告列表
func (h *LocalHandler) GetPromotionList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	req := &oceanengine.PromotionListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.client.Local().GetPromotionList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetPromotionDetail 获取广告详情
func (h *LocalHandler) GetPromotionDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionID, _ := strconv.ParseUint(c.Param("promotion_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Local().GetPromotionDetail(c.Request.Context(), accessToken, advertiserID, promotionID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// CreatePromotion 创建广告
func (h *LocalHandler) CreatePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var promotionData map[string]interface{}
	if err := c.ShouldBindJSON(&promotionData); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	promotionID, err := h.client.Local().CreatePromotion(c.Request.Context(), accessToken, advertiserID, promotionData)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"promotion_id": promotionID})
}

// UpdatePromotion 更新广告
func (h *LocalHandler) UpdatePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionID, _ := strconv.ParseUint(c.Param("promotion_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	updateData["promotion_id"] = promotionID

	err := h.client.Local().UpdatePromotion(c.Request.Context(), accessToken, advertiserID, promotionID, updateData)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// UpdatePromotionStatusRequest 更新广告状态请求
type UpdatePromotionStatusRequest struct {
	PromotionIDs []uint64 `json:"promotion_ids" binding:"required"`
	OptStatus    string   `json:"opt_status" binding:"required"`
}

// UpdatePromotionStatus 更新广告状态 (支持批量)
func (h *LocalHandler) UpdatePromotionStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req UpdatePromotionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Local().UpdatePromotionStatus(c.Request.Context(), accessToken, advertiserID, req.PromotionIDs, req.OptStatus)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// DeletePromotion 删除广告
func (h *LocalHandler) DeletePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionID, _ := strconv.ParseUint(c.Param("promotion_id"), 10, 64)
	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Local().DeletePromotion(c.Request.Context(), accessToken, advertiserID, []uint64{promotionID})
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// GetClueList 获取线索列表
func (h *LocalHandler) GetClueList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	list, total, err := h.client.Local().GetClueList(c.Request.Context(), accessToken, advertiserID, startDate, endDate, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetClueDetail 获取线索详情
func (h *LocalHandler) GetClueDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	clueID, _ := strconv.ParseUint(c.Param("clue_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || clueID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Local().GetClueDetail(c.Request.Context(), accessToken, advertiserID, clueID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// UpdateClueStatus 更新线索状态
func (h *LocalHandler) UpdateClueStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	clueID, _ := strconv.ParseUint(c.Param("clue_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || clueID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req struct {
		FollowStatus int    `json:"follow_status"`
		Remark       string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.client.Local().UpdateClueFollowStatus(c.Request.Context(), accessToken, advertiserID, clueID, req.FollowStatus, req.Remark)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// ExportCluesRequest 导出线索请求
type ExportCluesRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

// ExportClues 导出线索数据
func (h *LocalHandler) ExportClues(c *gin.Context) {
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

	// 获取线索报表数据作为导出数据源
	list, total, err := h.client.Local().GetClueReport(c.Request.Context(), accessToken, advertiserID, req.StartDate, req.EndDate, 1, 1000)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"data":    list,
		"total":   total,
		"message": "线索数据导出完成，请在前端进行文件下载处理",
	})
}

// GetProjectReport 获取项目报表
func (h *LocalHandler) GetProjectReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	list, err := h.client.Local().GetProjectReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(len(list)), page, pageSize)
}

// GetPromotionReport 获取广告报表
func (h *LocalHandler) GetPromotionReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	list, err := h.client.Local().GetPromotionReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(len(list)), page, pageSize)
}

// GetMaterialReport 获取素材报表
func (h *LocalHandler) GetMaterialReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	list, err := h.client.Local().GetMaterialReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(len(list)), page, pageSize)
}

// GetMaterialList 获取素材列表
func (h *LocalHandler) GetMaterialList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Local().GetVideoList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// UploadVideoRequest 上传视频请求
type UploadVideoRequest struct {
	VideoURL string `json:"video_url" binding:"required"`
}

// UploadVideo 上传视频 (异步任务)
func (h *LocalHandler) UploadVideo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req UploadVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	taskID, err := h.client.Local().CreateVideoUploadTask(c.Request.Context(), accessToken, advertiserID, req.VideoURL)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"task_id": taskID, "message": "视频上传任务已创建，请查询任务状态获取结果"})
}

// UploadImage 上传图片 (通过文件上传接口)
func (h *LocalHandler) UploadImage(c *gin.Context) {
	// 本地推图片上传需通过通用文件上传接口
	response.OKWithData(c, gin.H{
		"image_id": "",
		"url":      "",
		"message":  "请使用通用文件上传接口 /api/v1/file/image/upload",
	})
}

// DeleteMaterial 删除素材
func (h *LocalHandler) DeleteMaterial(c *gin.Context) {
	// 本地推平台暂不支持素材删除，素材需要手动在投放平台管理
	response.OKWithData(c, gin.H{"message": "素材删除请在投放平台操作"})
}

// GetStoreList 获取门店列表
func (h *LocalHandler) GetStoreList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.Local().GetStoreList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetStoreDetail 获取门店详情
func (h *LocalHandler) GetStoreDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}
	// 门店详情通过列表接口筛选获取，暂返回空数据
	response.OKWithData(c, gin.H{"message": "门店详情请通过列表接口查询"})
}

// CreateStore 创建门店
func (h *LocalHandler) CreateStore(c *gin.Context) {
	// 本地推门店通过POI同步创建，不支持直接创建
	response.OKWithData(c, gin.H{
		"store_id": "",
		"message":  "门店需通过抹茶POI同步创建，无法直接创建",
	})
}

// UpdateStore 更新门店
func (h *LocalHandler) UpdateStore(c *gin.Context) {
	// 门店信息通过POI同步，不支持直接编辑
	response.OKWithData(c, gin.H{"message": "门店信息请在抹茶平台更新"})
}

// DeleteStore 删除门店
func (h *LocalHandler) DeleteStore(c *gin.Context) {
	// 门店删除需在抹茶平台操作
	response.OKWithData(c, gin.H{"message": "门店删除请在抹茶平台操作"})
}
