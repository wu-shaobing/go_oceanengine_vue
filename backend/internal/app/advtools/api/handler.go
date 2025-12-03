package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// AdvToolsHandler 高级工具处理器
type AdvToolsHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewAdvToolsHandler 创建高级工具处理器
func NewAdvToolsHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *AdvToolsHandler {
	return &AdvToolsHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *AdvToolsHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *AdvToolsHandler) getAdvertiserID(c *gin.Context) int64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// ==================== RTA策略管理 ====================

// GetRtaInfo 获取RTA策略数据
func (h *AdvToolsHandler) GetRtaInfo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.RtaGetInfoRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := advToolsSvc.GetRtaInfo(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetAvailableRta 获取可用RTA策略
func (h *AdvToolsHandler) GetAvailableRta(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.RtaGetRequest{
		AdvertiserID: advertiserID,
	}

	list, err := advToolsSvc.GetAvailableRta(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// UpdateRtaStatus 更新RTA策略状态
func (h *AdvToolsHandler) UpdateRtaStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.RtaStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	err := advToolsSvc.UpdateRtaStatus(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// SetRtaScope 设置RTA策略生效范围
func (h *AdvToolsHandler) SetRtaScope(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.RtaSetScopeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	err := advToolsSvc.SetRtaScope(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// GetRtaScope 获取RTA策略绑定信息
func (h *AdvToolsHandler) GetRtaScope(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.RtaScopeGetRequest{
		AdvertiserID: advertiserID,
	}

	list, err := advToolsSvc.GetRtaScope(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 一键起量管理 ====================

// SetAdRaise 启动一键起量
func (h *AdvToolsHandler) SetAdRaise(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AdRaiseSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	status, err := advToolsSvc.SetAdRaise(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"raise_status": status})
}

// GetAdRaiseEstimate 获取起量预估值
func (h *AdvToolsHandler) GetAdRaiseEstimate(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseInt(c.Query("ad_id"), 10, 64)
	budget, _ := strconv.ParseFloat(c.Query("budget"), 64)

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.AdRaiseEstimateRequest{
		AdvertiserID: advertiserID,
		AdID:         adID,
		Budget:       budget,
	}

	estimate, err := advToolsSvc.GetAdRaiseEstimate(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"estimate_count": estimate})
}

// GetAdRaiseStatus 获取起量状态
func (h *AdvToolsHandler) GetAdRaiseStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req struct {
		AdIDs []int64 `json:"ad_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	statusReq := &oceanengine.AdRaiseStatusRequest{
		AdvertiserID: advertiserID,
		AdIDs:        req.AdIDs,
	}

	result, err := advToolsSvc.GetAdRaiseStatus(c.Request.Context(), statusReq)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetAdRaiseResult 获取起量后验数据
func (h *AdvToolsHandler) GetAdRaiseResult(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req struct {
		AdIDs []int64 `json:"ad_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	resultReq := &oceanengine.AdRaiseResultRequest{
		AdvertiserID: advertiserID,
		AdIDs:        req.AdIDs,
	}

	list, err := advToolsSvc.GetAdRaiseResult(c.Request.Context(), resultReq)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetSuggestBudget 获取建议起量预算
func (h *AdvToolsHandler) GetSuggestBudget(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req struct {
		AdIDs []int64 `json:"ad_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	budgetReq := &oceanengine.SuggestBudgetGetRequest{
		AdvertiserID: advertiserID,
		AdIDs:        req.AdIDs,
	}

	list, err := advToolsSvc.GetSuggestBudget(c.Request.Context(), budgetReq)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 定向包管理 ====================

// GetAudiencePackage 获取定向包列表
func (h *AdvToolsHandler) GetAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	landingType := c.Query("landing_type")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.AudiencePackageGetRequest{
		AdvertiserID: advertiserID,
		LandingType:  landingType,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := advToolsSvc.GetAudiencePackage(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// CreateAudiencePackage 创建定向包
func (h *AdvToolsHandler) CreateAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AudiencePackageCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	id, err := advToolsSvc.CreateAudiencePackage(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"audience_package_id": id})
}

// UpdateAudiencePackage 更新定向包
func (h *AdvToolsHandler) UpdateAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AudiencePackageUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	id, err := advToolsSvc.UpdateAudiencePackage(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"audience_package_id": id})
}

// DeleteAudiencePackage 删除定向包
func (h *AdvToolsHandler) DeleteAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AudiencePackageDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	id, err := advToolsSvc.DeleteAudiencePackage(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"audience_package_id": id})
}

// BindAudiencePackage 计划绑定定向包
func (h *AdvToolsHandler) BindAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AdBindRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	id, err := advToolsSvc.BindAudiencePackage(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"audience_package_id": id})
}

// UnbindAudiencePackage 解绑定向包
func (h *AdvToolsHandler) UnbindAudiencePackage(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AdBindRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	id, err := advToolsSvc.UnbindAudiencePackage(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"audience_package_id": id})
}

// ==================== 原生锚点管理 ====================

// GetNativeAnchor 获取原生锚点列表
func (h *AdvToolsHandler) GetNativeAnchor(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	anchorType := c.Query("anchor_type")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.NativeAnchorGetRequest{
		AdvertiserID: advertiserID,
		AnchorType:   anchorType,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := advToolsSvc.GetNativeAnchor(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetNativeAnchorDetail 获取原生锚点详情
func (h *AdvToolsHandler) GetNativeAnchorDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	anchorID, _ := strconv.ParseInt(c.Param("anchor_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || anchorID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.NativeAnchorDetailRequest{
		AdvertiserID: advertiserID,
		AnchorID:     anchorID,
	}

	result, err := advToolsSvc.GetNativeAnchorDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// CreateNativeAnchor 创建原生锚点
func (h *AdvToolsHandler) CreateNativeAnchor(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.NativeAnchorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	result, err := advToolsSvc.CreateNativeAnchor(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdateNativeAnchor 更新原生锚点
func (h *AdvToolsHandler) UpdateNativeAnchor(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.NativeAnchorUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	result, err := advToolsSvc.UpdateNativeAnchor(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// DeleteNativeAnchor 删除原生锚点
func (h *AdvToolsHandler) DeleteNativeAnchor(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.NativeAnchorDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	err := advToolsSvc.DeleteNativeAnchor(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 诊断工具 ====================

// GetDiagnosisSuggestion 获取计划诊断建议
func (h *AdvToolsHandler) GetDiagnosisSuggestion(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseInt(c.Query("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.DiagnosisSuggestionGetRequest{
		AdvertiserID: advertiserID,
		AdID:         adID,
	}

	result, err := advToolsSvc.GetDiagnosisSuggestion(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": result.List})
}

// AcceptDiagnosisSuggestion 采纳诊断建议
func (h *AdvToolsHandler) AcceptDiagnosisSuggestion(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.DiagnosisSuggestionAcceptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	result, err := advToolsSvc.AcceptDiagnosisSuggestion(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// ==================== 其他工具 ====================

// GetQuota 获取在投计划配额
func (h *AdvToolsHandler) GetQuota(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	req := &oceanengine.QuotaGetRequest{
		AdvertiserID: advertiserID,
	}

	result, err := advToolsSvc.GetQuota(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetAdQuality 获取广告质量度
func (h *AdvToolsHandler) GetAdQuality(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req struct {
		AdIDs []int64 `json:"ad_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	qualityReq := &oceanengine.AdQualityGetRequest{
		AdvertiserID: advertiserID,
		AdIDs:        req.AdIDs,
	}

	list, err := advToolsSvc.GetAdQuality(c.Request.Context(), qualityReq)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetAdStatExtraInfo 获取广告学习期状态
func (h *AdvToolsHandler) GetAdStatExtraInfo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	var req struct {
		AdIDs []int64 `json:"ad_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	advToolsSvc := oceanengine.NewAdvToolsService(h.client)
	extraReq := &oceanengine.AdStatExtraInfoGetRequest{
		AdvertiserID: advertiserID,
		AdIDs:        req.AdIDs,
	}

	list, err := advToolsSvc.GetAdStatExtraInfo(c.Request.Context(), extraReq)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}
