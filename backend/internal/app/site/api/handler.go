package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// SiteHandler 建站管理处理器
type SiteHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewSiteHandler 创建建站管理处理器
func NewSiteHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *SiteHandler {
	return &SiteHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 获取access token
func (h *SiteHandler) getAccessToken(c *gin.Context) string {
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
func (h *SiteHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.PostForm("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// ==================== 橙子建站 ====================

// GetOrangeSiteList 获取橙子建站落地页列表
func (h *SiteHandler) GetOrangeSiteList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetOrangeSiteList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetOrangeSiteDetail 获取橙子建站落地页详情
func (h *SiteHandler) GetOrangeSiteDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Site().GetOrangeSiteDetail(c.Request.Context(), accessToken, advertiserID, siteID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, detail)
}

// CreateOrangeSite 创建橙子建站落地页
func (h *SiteHandler) CreateOrangeSite(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.OrangeSiteCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	siteID, err := h.client.Site().CreateOrangeSite(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"site_id": siteID})
}

// UpdateOrangeSite 更新橙子建站落地页
func (h *SiteHandler) UpdateOrangeSite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		SiteName     string                   `json:"site_name"`
		Components   []map[string]interface{} `json:"components"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Site().UpdateOrangeSite(c.Request.Context(), accessToken, req.AdvertiserID, siteID, req.SiteName, req.Components)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// PublishOrangeSite 发布橙子建站落地页
func (h *SiteHandler) PublishOrangeSite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	siteURL, err := h.client.Site().PublishOrangeSite(c.Request.Context(), accessToken, advertiserID, siteID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"site_url": siteURL})
}

// DeleteOrangeSite 删除橙子建站落地页
func (h *SiteHandler) DeleteOrangeSite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Site().DeleteOrangeSite(c.Request.Context(), accessToken, advertiserID, siteID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// CopyOrangeSite 复制橙子建站落地页
func (h *SiteHandler) CopyOrangeSite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		SiteName     string `json:"site_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	newSiteID, err := h.client.Site().CopyOrangeSite(c.Request.Context(), accessToken, req.AdvertiserID, siteID, req.SiteName)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"site_id": newSiteID})
}

// ==================== 第三方落地页 ====================

// GetThirdPartySiteList 获取第三方落地页列表
func (h *SiteHandler) GetThirdPartySiteList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetThirdPartySiteList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateThirdPartySite 创建第三方落地页
func (h *SiteHandler) CreateThirdPartySite(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		SiteName     string `json:"site_name"`
		SiteURL      string `json:"site_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	siteID, err := h.client.Site().CreateThirdPartySite(c.Request.Context(), accessToken, req.AdvertiserID, req.SiteName, req.SiteURL)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"site_id": siteID})
}

// UpdateThirdPartySite 更新第三方落地页
func (h *SiteHandler) UpdateThirdPartySite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		SiteName     string `json:"site_name"`
		SiteURL      string `json:"site_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Site().UpdateThirdPartySite(c.Request.Context(), accessToken, req.AdvertiserID, siteID, req.SiteName, req.SiteURL)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// DeleteThirdPartySite 删除第三方落地页
func (h *SiteHandler) DeleteThirdPartySite(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Param("site_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Site().DeleteThirdPartySite(c.Request.Context(), accessToken, advertiserID, siteID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// ==================== 落地页模板 ====================

// GetSiteTemplateList 获取落地页模板列表
func (h *SiteHandler) GetSiteTemplateList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	templateType := c.Query("template_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetSiteTemplateList(c.Request.Context(), accessToken, advertiserID, templateType, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateSiteFromTemplate 从模板创建落地页
func (h *SiteHandler) CreateSiteFromTemplate(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		TemplateID   uint64 `json:"template_id"`
		SiteName     string `json:"site_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.TemplateID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	siteID, err := h.client.Site().CreateSiteFromTemplate(c.Request.Context(), accessToken, req.AdvertiserID, req.TemplateID, req.SiteName)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"site_id": siteID})
}

// ==================== 落地页组件 ====================

// GetSiteComponentList 获取落地页组件列表
func (h *SiteHandler) GetSiteComponentList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	componentType := c.Query("component_type")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Site().GetSiteComponentList(c.Request.Context(), accessToken, advertiserID, componentType)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, list)
}

// ==================== 落地页分析 ====================

// GetSiteAnalysis 获取落地页分析数据
func (h *SiteHandler) GetSiteAnalysis(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Query("site_id"), 10, 64)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	analysis, err := h.client.Site().GetSiteAnalysis(c.Request.Context(), accessToken, advertiserID, siteID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, analysis)
}

// GetSiteHeatmap 获取落地页热力图
func (h *SiteHandler) GetSiteHeatmap(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Query("site_id"), 10, 64)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	heatmap, err := h.client.Site().GetSiteHeatmap(c.Request.Context(), accessToken, advertiserID, siteID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, heatmap)
}

// ==================== 微页面 ====================

// GetMiniPageList 获取微页面列表
func (h *SiteHandler) GetMiniPageList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetMiniPageList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateMiniPage 创建微页面
func (h *SiteHandler) CreateMiniPage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		PageName     string `json:"page_name"`
		PagePath     string `json:"page_path"`
		AppID        string `json:"app_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	pageID, err := h.client.Site().CreateMiniPage(c.Request.Context(), accessToken, req.AdvertiserID, req.PageName, req.PagePath, req.AppID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"page_id": pageID})
}

// DeleteMiniPage 删除微页面
func (h *SiteHandler) DeleteMiniPage(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	pageID, _ := strconv.ParseUint(c.Param("page_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || pageID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Site().DeleteMiniPage(c.Request.Context(), accessToken, advertiserID, pageID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// ==================== 表单管理 ====================

// GetSiteFormList 获取落地页表单列表
func (h *SiteHandler) GetSiteFormList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	siteID, _ := strconv.ParseUint(c.Query("site_id"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || siteID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetSiteFormList(c.Request.Context(), accessToken, advertiserID, siteID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetFormSubmissionList 获取表单提交记录
func (h *SiteHandler) GetFormSubmissionList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	formID, _ := strconv.ParseUint(c.Query("form_id"), 10, 64)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || formID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Site().GetFormSubmissionList(c.Request.Context(), accessToken, advertiserID, formID, startDate, endDate, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}

// ExportFormSubmissions 导出表单提交数据
func (h *SiteHandler) ExportFormSubmissions(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		FormID       uint64 `json:"form_id"`
		StartDate    string `json:"start_date"`
		EndDate      string `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.FormID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	downloadURL, err := h.client.Site().ExportFormSubmissions(c.Request.Context(), accessToken, req.AdvertiserID, req.FormID, req.StartDate, req.EndDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"download_url": downloadURL})
}
