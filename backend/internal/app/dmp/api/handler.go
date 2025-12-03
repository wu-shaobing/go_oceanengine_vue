package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// DMPHandler DMP人群包管理处理器
type DMPHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewDMPHandler 创建DMP处理器
func NewDMPHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *DMPHandler {
	return &DMPHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *DMPHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.GetHeader("Access-Token")
	}
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *DMPHandler) getAdvertiserID(c *gin.Context) int64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// ==================== 数据源管理 ====================

// UploadDataSourceFile 上传数据源文件
func (h *DMPHandler) UploadDataSourceFile(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "文件上传失败: "+err.Error())
		return
	}
	defer file.Close()

	// 读取文件内容
	fileBytes := make([]byte, header.Size)
	_, err = file.Read(fileBytes)
	if err != nil {
		response.InternalError(c, "读取文件失败: "+err.Error())
		return
	}

	result, err := h.client.DMP().UploadDataSourceFile(c.Request.Context(), advertiserID, header.Filename, fileBytes)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"file_path": result.Path})
}

// CreateDataSource 创建数据源
func (h *DMPHandler) CreateDataSource(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.DataSourceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	result, err := h.client.DMP().CreateDataSource(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"data_source_id": result.DataSourceID})
}

// UpdateDataSource 更新数据源
func (h *DMPHandler) UpdateDataSource(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.DataSourceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.DataSourceID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	err := h.client.DMP().UpdateDataSource(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// GetDataSourceDetail 获取数据源详情
func (h *DMPHandler) GetDataSourceDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	dataSourceIDs := c.QueryArray("data_source_ids")

	if accessToken == "" || advertiserID == 0 || len(dataSourceIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	req := &oceanengine.DataSourceReadRequest{
		AdvertiserID:  advertiserID,
		DataSourceIDs: dataSourceIDs,
	}

	list, err := h.client.DMP().GetDataSourceDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 人群包管理 ====================

// GetCustomAudienceList 获取人群包列表
func (h *DMPHandler) GetCustomAudienceList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	req := &oceanengine.CustomAudienceListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := h.client.DMP().GetCustomAudienceList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetCustomAudienceDetail 获取人群包详情
func (h *DMPHandler) GetCustomAudienceDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	audienceIDsStr := c.QueryArray("audience_ids")

	if accessToken == "" || advertiserID == 0 || len(audienceIDsStr) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	audienceIDs := make([]int64, len(audienceIDsStr))
	for i, s := range audienceIDsStr {
		audienceIDs[i], _ = strconv.ParseInt(s, 10, 64)
	}

	h.client.SetAccessToken(accessToken)
	req := &oceanengine.CustomAudienceReadRequest{
		AdvertiserID:      advertiserID,
		CustomAudienceIDs: audienceIDs,
	}

	list, err := h.client.DMP().GetCustomAudienceDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// CreateCustomAudience 创建人群包
func (h *DMPHandler) CreateCustomAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CustomAudienceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	result, err := h.client.DMP().CreateCustomAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"custom_audience_id": result.CustomAudienceID})
}

// PublishCustomAudience 发布人群包
func (h *DMPHandler) PublishCustomAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CustomAudiencePublishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CustomAudienceID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	err := h.client.DMP().PublishCustomAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// PushCustomAudience 推送人群包
func (h *DMPHandler) PushCustomAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CustomAudiencePushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CustomAudienceID == 0 || len(req.TargetAdvertiserIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	err := h.client.DMP().PushCustomAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// DeleteCustomAudience 删除人群包
func (h *DMPHandler) DeleteCustomAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CustomAudienceDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CustomAudienceID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	err := h.client.DMP().DeleteCustomAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 云图相关 ====================

// GetBrandList 获取广告账户关联云图账户信息
func (h *DMPHandler) GetBrandList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().GetBrandList(c.Request.Context(), advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// CopyAudienceToBrand 推送DMP人群包到云图账户
func (h *DMPHandler) CopyAudienceToBrand(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CustomAudienceCopyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CustomAudienceID == 0 || req.BrandID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	err := h.client.DMP().CopyCustomAudienceToBrand(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== Lookalike相似人群 ====================

// CreateLookalikeAudience 创建相似人群
func (h *DMPHandler) CreateLookalikeAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.LookalikeAudienceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	result, err := h.client.DMP().CreateLookalikeAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"custom_audience_id": result.CustomAudienceID})
}

// ==================== 行为兴趣定向 ====================

// GetInterestCategories 获取兴趣分类列表
func (h *DMPHandler) GetInterestCategories(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().GetInterestCategories(c.Request.Context(), advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetActionCategories 获取行为分类列表
func (h *DMPHandler) GetActionCategories(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	actionScene := c.DefaultQuery("action_scene", "E-COMMERCE")
	actionDays, _ := strconv.Atoi(c.DefaultQuery("action_days", "7"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().GetActionCategories(c.Request.Context(), advertiserID, actionScene, actionDays)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// SearchInterestKeywords 搜索兴趣关键词
func (h *DMPHandler) SearchInterestKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	query := c.Query("query")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().SearchInterestKeywords(c.Request.Context(), advertiserID, query)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 抖音达人定向 ====================

// SearchAwemeAuthors 搜索抖音达人
func (h *DMPHandler) SearchAwemeAuthors(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	query := c.Query("query")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().SearchAwemeAuthors(c.Request.Context(), advertiserID, query)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetAwemeAuthorCategories 获取抖音达人分类
func (h *DMPHandler) GetAwemeAuthorCategories(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	list, err := h.client.DMP().GetAwemeAuthorCategories(c.Request.Context(), advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 人群估算 ====================

// EstimateAudience 估算人群覆盖
func (h *DMPHandler) EstimateAudience(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AudienceEstimateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	h.client.SetAccessToken(accessToken)
	result, err := h.client.DMP().EstimateAudience(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"user_count": result.UserCount})
}
