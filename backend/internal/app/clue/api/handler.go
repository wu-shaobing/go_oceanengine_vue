package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// ClueHandler 线索管理处理器
type ClueHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewClueHandler 创建线索处理器
func NewClueHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *ClueHandler {
	return &ClueHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *ClueHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *ClueHandler) getAdvertiserID(c *gin.Context) int64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// ==================== 飞鱼线索 ====================

// GetClueList 获取线索列表
func (h *ClueHandler) GetClueList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || startTime == "" || endTime == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.ClueListRequest{
		AdvertiserID: advertiserID,
		StartTime:    startTime,
		EndTime:      endTime,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := clueService.GetClueList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// ClueCallback 回传有效线索
func (h *ClueHandler) ClueCallback(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.ClueCallbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.ClueID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	err := clueService.ClueCallback(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "回传成功"})
}

// BatchClueCallback 批量回传线索
func (h *ClueHandler) BatchClueCallback(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.BatchClueCallbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.ClueList) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	result, err := clueService.BatchClueCallback(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetKeyAction 获取活动记录
func (h *ClueHandler) GetKeyAction(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	clueIDStr := c.Query("clue_id")
	clueID, _ := strconv.ParseInt(clueIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || clueID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.KeyActionGetRequest{
		AdvertiserID: advertiserID,
		ClueID:       clueID,
	}

	result, err := clueService.GetKeyAction(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetSmartPhone 查询智能电话列表
func (h *ClueHandler) GetSmartPhone(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.SmartPhoneGetRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := clueService.GetSmartPhone(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetFormList 查询表单列表
func (h *ClueHandler) GetFormList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	formType := c.Query("form_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.FormGetRequest{
		AdvertiserID: advertiserID,
		FormType:     formType,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := clueService.GetFormList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetFormDetail 查询表单详情
func (h *ClueHandler) GetFormDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	formIDStr := c.Param("form_id")
	formID, _ := strconv.ParseInt(formIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || formID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.FormDetailRequest{
		AdvertiserID: advertiserID,
		FormID:       formID,
	}

	result, err := clueService.GetFormDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetClueStoreList 获取线索店铺列表
func (h *ClueHandler) GetClueStoreList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	storeName := c.Query("store_name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	clueService := oceanengine.NewClueService(h.client)
	req := &oceanengine.ClueStoreListRequest{
		AdvertiserID: advertiserID,
		StoreName:    storeName,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := clueService.GetClueStoreList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// ==================== 青鸟线索通 - 表单管理 ====================

// GetQingniaoFormList 获取青鸟表单列表
func (h *ClueHandler) GetQingniaoFormList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	formName := c.Query("form_name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.FormListRequest{
		AdvertiserID: advertiserID,
		FormName:     formName,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetFormList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// CreateQingniaoForm 创建青鸟表单
func (h *ClueHandler) CreateQingniaoForm(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.FormCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	formID, err := qingniaoService.CreateForm(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"form_id": formID})
}

// UpdateQingniaoForm 更新青鸟表单
func (h *ClueHandler) UpdateQingniaoForm(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.FormUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.FormID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	formID, err := qingniaoService.UpdateForm(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"form_id": formID})
}

// DeleteQingniaoForm 删除青鸟表单
func (h *ClueHandler) DeleteQingniaoForm(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	formIDStr := c.Param("form_id")
	formID, _ := strconv.ParseInt(formIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || formID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.FormDeleteRequest{
		AdvertiserID: advertiserID,
		FormID:       formID,
	}

	_, err := qingniaoService.DeleteForm(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "删除成功"})
}

// ==================== 青鸟线索通 - 卡券管理 ====================

// GetCouponList 获取卡券列表
func (h *ClueHandler) GetCouponList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.CouponListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetCouponList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// CreateCoupon 创建卡券
func (h *ClueHandler) CreateCoupon(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CouponCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	couponID, err := qingniaoService.CreateCoupon(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"coupon_id": couponID})
}

// GetCouponDetail 获取卡券详情
func (h *ClueHandler) GetCouponDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	couponIDStr := c.Param("coupon_id")
	couponID, _ := strconv.ParseInt(couponIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || couponID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.CouponDetailRequest{
		AdvertiserID: advertiserID,
		CouponID:     couponID,
	}

	result, err := qingniaoService.GetCouponDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdateCoupon 更新卡券
func (h *ClueHandler) UpdateCoupon(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CouponUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CouponID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	err := qingniaoService.UpdateCoupon(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "更新成功"})
}

// UploadCouponCode 上传券码
func (h *ClueHandler) UploadCouponCode(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CouponCodeUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CouponID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	result, err := qingniaoService.UploadCouponCode(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// ConsumeCouponCode 核销券码
func (h *ClueHandler) ConsumeCouponCode(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.CouponCodeConsumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.CouponID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	err := qingniaoService.ConsumeCouponCode(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "核销成功"})
}

// ==================== 青鸟线索通 - 智能电话 ====================

// GetQingniaoSmartPhoneList 获取青鸟智能电话列表
func (h *ClueHandler) GetQingniaoSmartPhoneList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.QingniaoSmartPhoneListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetSmartPhoneList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// CreateSmartPhone 创建智能电话
func (h *ClueHandler) CreateSmartPhone(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.SmartPhoneCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	result, err := qingniaoService.CreateSmartPhone(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// DeleteSmartPhone 删除智能电话
func (h *ClueHandler) DeleteSmartPhone(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	smartPhoneIDStr := c.Param("smart_phone_id")
	smartPhoneID, _ := strconv.ParseInt(smartPhoneIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || smartPhoneID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.SmartPhoneDeleteRequest{
		AdvertiserID: advertiserID,
		SmartPhoneID: smartPhoneID,
	}

	err := qingniaoService.DeleteSmartPhone(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"message": "删除成功"})
}

// GetSmartPhoneRecords 获取智能电话拨打记录
func (h *ClueHandler) GetSmartPhoneRecords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	smartPhoneIDStr := c.Query("smart_phone_id")
	smartPhoneID, _ := strconv.ParseInt(smartPhoneIDStr, 10, 64)
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || startTime == "" || endTime == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.SmartPhoneRecordRequest{
		AdvertiserID: advertiserID,
		SmartPhoneID: smartPhoneID,
		StartTime:    startTime,
		EndTime:      endTime,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetSmartPhoneRecords(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// ==================== 青鸟线索通 - 微信加粉 ====================

// GetWechatPoolList 获取微信库微信号列表
func (h *ClueHandler) GetWechatPoolList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.WechatPoolListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetWechatPoolList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetWechatInstanceList 获取微信号码包列表
func (h *ClueHandler) GetWechatInstanceList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.WechatInstanceListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}

	result, err := qingniaoService.GetWechatInstanceList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, result.List, int64(result.PageInfo.TotalNumber), page, pageSize)
}

// GetWechatInstanceDetail 获取微信号码包详情
func (h *ClueHandler) GetWechatInstanceDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	instanceIDStr := c.Param("instance_id")
	instanceID, _ := strconv.ParseInt(instanceIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || instanceID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	req := &oceanengine.WechatInstanceDetailRequest{
		AdvertiserID: advertiserID,
		InstanceID:   instanceID,
	}

	result, err := qingniaoService.GetWechatInstanceDetail(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdateWechatInstance 更新微信号码包
func (h *ClueHandler) UpdateWechatInstance(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.WechatInstanceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.InstanceID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	qingniaoService := oceanengine.NewQingniaoService(h.client)
	result, err := qingniaoService.UpdateWechatInstance(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}
