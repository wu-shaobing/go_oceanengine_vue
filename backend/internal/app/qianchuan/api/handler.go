package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// QianchuanHandler 千川处理器
type QianchuanHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewQianchuanHandler 创建千川处理器
func NewQianchuanHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *QianchuanHandler {
	return &QianchuanHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *QianchuanHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *QianchuanHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// GetAccountInfo 获取千川账户信息
func (h *QianchuanHandler) GetAccountInfo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	info, err := h.client.Qianchuan().GetAccountInfo(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, info)
}

// GetShopList 获取店铺列表
func (h *QianchuanHandler) GetShopList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetShopList(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetAwemeAuthList 获取已授权抹音号列表
func (h *QianchuanHandler) GetAwemeAuthList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetAuthorizedAwemeList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetBalance 获取账户余额
func (h *QianchuanHandler) GetBalance(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	balance, err := h.client.Qianchuan().GetBalance(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"balance": balance})
}

// GetCampaignList 获取广告组列表
func (h *QianchuanHandler) GetCampaignList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetCampaignList(c.Request.Context(), accessToken, advertiserID, page, pageSize, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateCampaign 创建广告组
func (h *QianchuanHandler) CreateCampaign(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID  uint64  `json:"advertiser_id"`
		CampaignName  string  `json:"campaign_name"`
		Budget        float64 `json:"budget"`
		BudgetMode    string  `json:"budget_mode"`
		MarketingGoal string  `json:"marketing_goal"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	campaignID, err := h.client.Qianchuan().CreateCampaign(c.Request.Context(), accessToken, req.AdvertiserID, req.CampaignName, req.Budget, req.BudgetMode, req.MarketingGoal)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"campaign_id": campaignID})
}

// GetAdList 获取广告计划列表
func (h *QianchuanHandler) GetAdList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.QianchuanAdListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
	}
	list, total, err := h.client.Qianchuan().GetAdList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetAdDetail 获取广告计划详情
func (h *QianchuanHandler) GetAdDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseUint(c.Param("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Qianchuan().GetAdDetail(c.Request.Context(), accessToken, advertiserID, adID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, detail)
}

// CreateAd 创建广告计划
func (h *QianchuanHandler) CreateAd(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	advertiserID := uint64(req["advertiser_id"].(float64))
	adID, err := h.client.Qianchuan().CreateAd(c.Request.Context(), accessToken, advertiserID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"ad_id": adID})
}

// UpdateAdStatus 更新广告状态
func (h *QianchuanHandler) UpdateAdStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		AdIDs        []uint64 `json:"ad_ids"`
		OptStatus    string   `json:"opt_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	err := h.client.Qianchuan().UpdateAdStatus(c.Request.Context(), accessToken, req.AdvertiserID, req.AdIDs, req.OptStatus)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// GetCreativeList 获取创意列表
func (h *QianchuanHandler) GetCreativeList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetCreativeList(c.Request.Context(), accessToken, advertiserID, page, pageSize, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetAwemeOrderList 获取随心推订单列表
func (h *QianchuanHandler) GetAwemeOrderList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetAwemeOrderList(c.Request.Context(), accessToken, advertiserID, page, pageSize, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetAwemeOrderDetail 获取随心推订单详情
func (h *QianchuanHandler) GetAwemeOrderDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	orderID, _ := strconv.ParseUint(c.Param("order_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || orderID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Qianchuan().GetAwemeOrderDetail(c.Request.Context(), accessToken, advertiserID, orderID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, detail)
}

// CreateAwemeOrder 创建随心推订单
func (h *QianchuanHandler) CreateAwemeOrder(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	advertiserID := uint64(req["advertiser_id"].(float64))
	order, err := h.client.Qianchuan().CreateAwemeOrder(c.Request.Context(), accessToken, advertiserID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, order)
}

// GetAdvertiserReport 获取账户报表
func (h *QianchuanHandler) GetAdvertiserReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetAdReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, list)
}

// GetAdReport 获取广告报表
func (h *QianchuanHandler) GetAdReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetAdReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetMaterialList 获取素材列表
func (h *QianchuanHandler) GetMaterialList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetMaterialReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// UploadImage 上传图片素材
func (h *QianchuanHandler) UploadImage(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 获取上传的文件
	file, header, err := c.Request.FormFile("image_file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的图片文件")
		return
	}
	defer file.Close()

	result, err := h.client.Qianchuan().UploadImageFromReader(c.Request.Context(), accessToken, advertiserID, header.Filename, file)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UploadVideo 上传视频素材
func (h *QianchuanHandler) UploadVideo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 获取上传的文件
	file, header, err := c.Request.FormFile("video_file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的视频文件")
		return
	}
	defer file.Close()

	result, err := h.client.Qianchuan().UploadVideoFromReader(c.Request.Context(), accessToken, advertiserID, header.Filename, file)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetIndustryList 获取行业列表
func (h *QianchuanHandler) GetIndustryList(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	if accessToken == "" {
		response.BadRequest(c, "缺少 access_token")
		return
	}

	list, err := h.client.Qianchuan().GetIndustryList(c.Request.Context(), accessToken)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, list)
}

// GetDmpList 获取人群包列表
func (h *QianchuanHandler) GetDmpList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetAudienceList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetProductList 获取商品列表
func (h *QianchuanHandler) GetProductList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	awemeID, _ := strconv.ParseUint(c.Query("aweme_id"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetProductList(c.Request.Context(), accessToken, advertiserID, awemeID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetBudget 获取账户预算
func (h *QianchuanHandler) GetBudget(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	budget, err := h.client.Qianchuan().GetAccountBudget(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"budget": budget})
}

// UpdateBudget 更新账户预算
func (h *QianchuanHandler) UpdateBudget(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64  `json:"advertiser_id"`
		Budget       float64 `json:"budget"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	err := h.client.Qianchuan().UpdateAccountBudget(c.Request.Context(), accessToken, req.AdvertiserID, req.Budget)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// GetFinanceDetail 获取财务明细
func (h *QianchuanHandler) GetFinanceDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetFinanceDetail(c.Request.Context(), accessToken, advertiserID, startDate, endDate, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// ==================== 全域推广 ====================

// GetUniList 获取全域推广列表
func (h *QianchuanHandler) GetUniList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.Qianchuan().GetUniPromotionList(c.Request.Context(), accessToken, advertiserID, page, pageSize, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateUni 创建全域推广
func (h *QianchuanHandler) CreateUni(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	advertiserID := uint64(req["advertiser_id"].(float64))
	adID, err := h.client.Qianchuan().CreateUniPromotion(c.Request.Context(), accessToken, advertiserID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"ad_id": adID})
}

// GetUniDetail 获取全域推广详情
func (h *QianchuanHandler) GetUniDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseUint(c.Param("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	detail, err := h.client.Qianchuan().GetUniPromotionDetail(c.Request.Context(), accessToken, advertiserID, adID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, detail)
}

// ==================== 扩展报表 ====================

// GetCreativeReport 获取创意报表
func (h *QianchuanHandler) GetCreativeReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetCreativeReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetMaterialReport 获取素材报表
func (h *QianchuanHandler) GetMaterialReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetMaterialReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetKeywordReport 获取关键词报表
func (h *QianchuanHandler) GetKeywordReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseUint(c.Query("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	if adID == 0 {
		response.OKWithList(c, []interface{}{}, 0, 1, 20)
		return
	}

	keywords, err := h.client.Qianchuan().GetKeywords(c.Request.Context(), accessToken, advertiserID, adID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": keywords})
}

// GetLiveReport 获取直播报表
func (h *QianchuanHandler) GetLiveReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetLiveReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, list)
}

// GetRoomReport 获取直播间报表
func (h *QianchuanHandler) GetRoomReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 使用通用的直播间报表接口
	list, err := h.client.Qianchuan().GetLiveReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	if len(list) > 0 {
		response.OKWithData(c, list[0])
	} else {
		response.OKWithData(c, gin.H{})
	}
}

// GetUniReport 获取全域推广报表
func (h *QianchuanHandler) GetUniReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Qianchuan().GetUniPromotionReport(c.Request.Context(), accessToken, advertiserID, startDate, endDate, nil)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, list)
}

// ==================== 工具扩展 ====================

// GetKeywordRecommend 获取关键词推荐
func (h *QianchuanHandler) GetKeywordRecommend(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseUint(c.Query("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	if adID == 0 {
		// 如果没有ad_id，返回词包列表
		packages, err := h.client.Qianchuan().GetKeywordPackages(c.Request.Context(), accessToken, advertiserID)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
		response.OKWithData(c, packages)
		return
	}

	keywords, err := h.client.Qianchuan().GetRecommendKeywords(c.Request.Context(), accessToken, advertiserID, adID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, keywords)
}

// ==================== 关键词管理 ====================

// GetKeywordList 获取计划关键词列表
func (h *QianchuanHandler) GetKeywordList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adID, _ := strconv.ParseUint(c.Query("ad_id"), 10, 64)

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	keywords, err := h.client.Qianchuan().GetKeywords(c.Request.Context(), accessToken, advertiserID, adID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": keywords})
}

// UpdateKeywordsRequest 更新关键词请求
type UpdateKeywordsRequest struct {
	AdID     uint64                `json:"ad_id" binding:"required"`
	Keywords []oceanengine.Keyword `json:"keywords" binding:"required"`
}

// UpdateKeywords 更新计划关键词
func (h *QianchuanHandler) UpdateKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req UpdateKeywordsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.client.Qianchuan().UpdateKeywords(c.Request.Context(), accessToken, advertiserID, req.AdID, req.Keywords)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 行为兴趣关键词 ====================

// GetActionKeywords 查询行为关键词
func (h *QianchuanHandler) GetActionKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	queryWord := c.Query("query_word")
	actionScene := c.DefaultQuery("action_scene", "E-COMMERCE")
	actionDays, _ := strconv.Atoi(c.DefaultQuery("action_days", "30"))

	if accessToken == "" || advertiserID == 0 || queryWord == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	keywords, err := h.client.Qianchuan().GetActionKeywords(c.Request.Context(), accessToken, advertiserID, queryWord, actionScene, actionDays)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": keywords})
}

// GetInterestKeywords 查询兴趣关键词
func (h *QianchuanHandler) GetInterestKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	queryWord := c.Query("query_word")

	if accessToken == "" || advertiserID == 0 || queryWord == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	keywords, err := h.client.Qianchuan().GetInterestKeywords(c.Request.Context(), accessToken, advertiserID, queryWord)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": keywords})
}

// GetKeywordSuggest 获取行为兴趣推荐关键词
func (h *QianchuanHandler) GetKeywordSuggest(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var req struct {
		Keywords []string `json:"keywords" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	keywords, err := h.client.Qianchuan().GetKeywordSuggest(c.Request.Context(), accessToken, advertiserID, req.Keywords)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": keywords})
}
