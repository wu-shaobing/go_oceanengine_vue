package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// EventManagerHandler 事件管理处理器
type EventManagerHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewEventManagerHandler 创建事件管理处理器
func NewEventManagerHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *EventManagerHandler {
	return &EventManagerHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *EventManagerHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *EventManagerHandler) getAdvertiserID(c *gin.Context) int64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// ==================== 资产管理 ====================

// GetAssets 获取已创建资产列表
func (h *EventManagerHandler) GetAssets(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetType := c.Query("asset_type")
	landingType := c.Query("landing_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.AssetsGetRequest{
		AdvertiserID: advertiserID,
		AssetType:    assetType,
		LandingType:  landingType,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := eventService.GetAssets(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetAllAssetsList 获取账户下资产列表
func (h *EventManagerHandler) GetAllAssetsList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetType := c.Query("asset_type")
	landingType := c.Query("landing_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.AllAssetsListRequest{
		AdvertiserID: advertiserID,
		AssetType:    assetType,
		LandingType:  landingType,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := eventService.GetAllAssetsList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// CreateAsset 创建事件资产
func (h *EventManagerHandler) CreateAsset(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AssetsCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	assetID, err := eventService.CreateAsset(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"asset_id": assetID})
}

// ==================== 事件管理 ====================

// GetAvailableEvents 获取可创建事件列表
func (h *EventManagerHandler) GetAvailableEvents(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetIDStr := c.Query("asset_id")
	assetID, _ := strconv.ParseInt(assetIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || assetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.AvailableEventsGetRequest{
		AdvertiserID: advertiserID,
		AssetID:      assetID,
	}

	result, err := eventService.GetAvailableEvents(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": result})
}

// GetEventConfigs 获取已创建事件列表
func (h *EventManagerHandler) GetEventConfigs(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetIDStr := c.Query("asset_id")
	assetID, _ := strconv.ParseInt(assetIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || assetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.EventConfigsGetRequest{
		AdvertiserID: advertiserID,
		AssetID:      assetID,
	}

	result, err := eventService.GetEventConfigs(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": result})
}

// CreateEvents 资产下创建事件
func (h *EventManagerHandler) CreateEvents(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.EventsCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AssetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	err := eventService.CreateEvents(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "创建成功"})
}

// ==================== 监测链接管理 ====================

// GetTrackURL 获取监测链接组
func (h *EventManagerHandler) GetTrackURL(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetIDStr := c.Query("asset_id")
	assetID, _ := strconv.ParseInt(assetIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || assetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.TrackURLGetRequest{
		AdvertiserID: advertiserID,
		AssetID:      assetID,
	}

	result, err := eventService.GetTrackURL(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// CreateTrackURL 创建监测链接组
func (h *EventManagerHandler) CreateTrackURL(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.TrackURLCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AssetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	err := eventService.CreateTrackURL(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "创建成功"})
}

// UpdateTrackURL 更新监测链接组
func (h *EventManagerHandler) UpdateTrackURL(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.TrackURLUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AssetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	err := eventService.UpdateTrackURL(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "更新成功"})
}

// ==================== 资产共享 ====================

// GetShare 查看共享范围
func (h *EventManagerHandler) GetShare(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetIDStr := c.Query("asset_id")
	assetID, _ := strconv.ParseInt(assetIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || assetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.ShareGetRequest{
		AdvertiserID: advertiserID,
		AssetID:      assetID,
	}

	result, err := eventService.GetShare(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// Share 资产共享
func (h *EventManagerHandler) Share(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.ShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AssetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	failList, err := eventService.Share(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"fail_list": failList})
}

// ShareCancel 取消资产共享
func (h *EventManagerHandler) ShareCancel(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.ShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AssetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	failList, err := eventService.ShareCancel(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"fail_list": failList})
}

// ==================== 优化目标 ====================

// GetOptimizedGoal 获取可用优化目标
func (h *EventManagerHandler) GetOptimizedGoal(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	assetIDStr := c.Query("asset_id")
	assetID, _ := strconv.ParseInt(assetIDStr, 10, 64)
	landingType := c.Query("landing_type")

	if accessToken == "" || advertiserID == 0 || assetID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.EventConvertOptimizedGoalGetRequest{
		AdvertiserID: advertiserID,
		AssetID:      assetID,
		LandingType:  landingType,
	}

	result, err := eventService.GetEventConvertOptimizedGoal(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// ==================== 转化回传 ====================

// Conversion 转化回传
func (h *EventManagerHandler) Conversion(c *gin.Context) {
	var req oceanengine.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.EventType == "" || req.Context == nil {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	_, err := eventService.Conversion(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "回传成功"})
}

// ==================== 转化回传鉴权 ====================

// AddPublicKey 新增公钥
func (h *EventManagerHandler) AddPublicKey(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.AddPublicKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	result, err := eventService.AddPublicKey(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetAllPublicKeys 查询全部公钥
func (h *EventManagerHandler) GetAllPublicKeys(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	req := &oceanengine.GetAllPublicKeysRequest{
		AdvertiserID: advertiserID,
	}

	result, err := eventService.GetAllPublicKeys(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": result})
}

// EnableAuth 开启鉴权
func (h *EventManagerHandler) EnableAuth(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	err := eventService.EnableAuth(c.Request.Context(), advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "开启成功"})
}

// DisableAuth 关闭鉴权
func (h *EventManagerHandler) DisableAuth(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	eventService := oceanengine.NewEventManagerService(h.client)
	err := eventService.DisableAuth(c.Request.Context(), advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "关闭成功"})
}
