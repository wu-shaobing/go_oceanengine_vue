package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	advModel "oceanengine-backend/internal/app/advertiser/model"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// ServeMarketHandler 服务市场处理器
type ServeMarketHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewServeMarketHandler 创建服务市场处理器
func NewServeMarketHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *ServeMarketHandler {
	return &ServeMarketHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAdvertiser 获取广告主及Token
func (h *ServeMarketHandler) getAdvertiser(c *gin.Context) (*advModel.Advertiser, error) {
	advIDStr := c.Query("advertiser_id")
	if advIDStr == "" {
		return nil, errcode.New(errcode.ErrInvalidParam)
	}
	advID, err := strconv.ParseUint(advIDStr, 10, 64)
	if err != nil {
		return nil, errcode.New(errcode.ErrInvalidParam)
	}
	var adv advModel.Advertiser
	if err := h.db.Where("advertiser_id = ?", advID).First(&adv).Error; err != nil {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}
	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}
	return &adv, nil
}

// GetOrderList 获取订单列表
func (h *ServeMarketHandler) GetOrderList(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.ServeMarket().GetAppOrderList(c.Request.Context(), adv.AccessToken, adv.AdvertiserID, page, pageSize)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetOrderDetail 获取订单详情
func (h *ServeMarketHandler) GetOrderDetail(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	orderID := c.Param("order_id")
	if orderID == "" {
		response.BadRequest(c, "订单ID不能为空")
		return
	}

	// 获取订单列表中的特定订单
	list, _, err := h.client.ServeMarket().GetAppOrderList(c.Request.Context(), adv.AccessToken, adv.AdvertiserID, 1, 100)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	for _, order := range list {
		if order.OrderID == orderID {
			response.OKWithData(c, order)
			return
		}
	}

	response.NotFound(c, "订单不存在")
}

// GetFuncList 获取已购功能列表
func (h *ServeMarketHandler) GetFuncList(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.client.ServeMarket().GetFuncPointList(c.Request.Context(), adv.AccessToken, adv.AdvertiserID, page, pageSize)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetFuncDetail 获取功能详情
func (h *ServeMarketHandler) GetFuncDetail(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	funcID := c.Param("func_id")
	if funcID == "" {
		response.BadRequest(c, "功能ID不能为空")
		return
	}

	// 获取功能列表中的特定功能
	list, _, err := h.client.ServeMarket().GetFuncPointList(c.Request.Context(), adv.AccessToken, adv.AdvertiserID, 1, 100)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	for _, funcPoint := range list {
		if funcPoint.FuncID == funcID {
			response.OKWithData(c, funcPoint)
			return
		}
	}

	response.NotFound(c, "功能点不存在")
}

// GetQualityAnalysis 获取投前分析
func (h *ServeMarketHandler) GetQualityAnalysis(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	targetID := c.Query("target_id")
	targetTypeStr := c.DefaultQuery("target_type", "1")
	targetType, _ := strconv.Atoi(targetTypeStr)

	if targetID == "" {
		response.BadRequest(c, "目标ID不能为空")
		return
	}

	report, err := h.client.ServeMarket().GetQualityReport(c.Request.Context(), adv.AccessToken, adv.AdvertiserID, targetID, targetType)
	if err != nil {
		// 如果没有报告，返回空数据
		response.OKWithData(c, gin.H{
			"quality_score":  0,
			"creative_score": 0,
			"target_score":   0,
			"budget_score":   0,
			"suggestions":    []string{},
		})
		return
	}

	// 转换维度数据
	scores := gin.H{
		"quality_score": report.Score,
		"report_id":     report.ReportID,
		"create_time":   report.CreateTime,
	}
	var suggestions []string
	for _, dim := range report.Dimensions {
		scores[dim.DimensionName+"_score"] = dim.Score
		if dim.Suggestion != "" {
			suggestions = append(suggestions, dim.Suggestion)
		}
	}
	scores["suggestions"] = suggestions

	response.OKWithData(c, scores)
}

// GetSubscriptionList 获取RDS订阅列表
func (h *ServeMarketHandler) GetSubscriptionList(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	list, err := h.client.ServeMarket().GetRdsSubscriptionList(c.Request.Context(), adv.AccessToken, adv.AdvertiserID)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OKWithList(c, list, int64(len(list)), 1, len(list))
}

// CreateSubscriptionRequest 创建订阅请求
type CreateSubscriptionRequest struct {
	AdvertiserID     uint64 `json:"advertiser_id" binding:"required"`
	SubscriptionType int    `json:"subscription_type" binding:"required"`
	CallbackURL      string `json:"callback_url" binding:"required,url"`
}

// CreateSubscription 创建RDS订阅
func (h *ServeMarketHandler) CreateSubscription(c *gin.Context) {
	var req CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	var adv advModel.Advertiser
	if err := h.db.Where("advertiser_id = ?", req.AdvertiserID).First(&adv).Error; err != nil {
		response.Error(c, errcode.New(errcode.ErrAdvertiserNotFound))
		return
	}

	subscriptionID, err := h.client.ServeMarket().CreateRdsSubscription(
		c.Request.Context(),
		adv.AccessToken,
		adv.AdvertiserID,
		req.SubscriptionType,
		req.CallbackURL,
	)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OKWithData(c, gin.H{"subscription_id": subscriptionID})
}

// UpdateSubscriptionRequest 更新订阅请求
type UpdateSubscriptionRequest struct {
	AdvertiserID uint64 `json:"advertiser_id" binding:"required"`
	CallbackURL  string `json:"callback_url" binding:"required,url"`
}

// UpdateSubscription 更新RDS订阅
func (h *ServeMarketHandler) UpdateSubscription(c *gin.Context) {
	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		response.BadRequest(c, "订阅ID不能为空")
		return
	}

	var req UpdateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	var adv advModel.Advertiser
	if err := h.db.Where("advertiser_id = ?", req.AdvertiserID).First(&adv).Error; err != nil {
		response.Error(c, errcode.New(errcode.ErrAdvertiserNotFound))
		return
	}

	err := h.client.ServeMarket().UpdateRdsSubscription(
		c.Request.Context(),
		adv.AccessToken,
		adv.AdvertiserID,
		subscriptionID,
		req.CallbackURL,
	)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OK(c)
}

// DeleteSubscription 删除RDS订阅
func (h *ServeMarketHandler) DeleteSubscription(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		response.BadRequest(c, "订阅ID不能为空")
		return
	}

	err = h.client.ServeMarket().DeleteRdsSubscription(
		c.Request.Context(),
		adv.AccessToken,
		adv.AdvertiserID,
		subscriptionID,
	)
	if err != nil {
		response.Error(c, errcode.Wrap(errcode.ErrOEAPIFailed, err))
		return
	}

	response.OK(c)
}

// GetDashboard 获取仪表盘数据
func (h *ServeMarketHandler) GetDashboard(c *gin.Context) {
	adv, err := h.getAdvertiser(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	ctx := c.Request.Context()

	// 获取订单数据
	orders, totalOrders, _ := h.client.ServeMarket().GetAppOrderList(ctx, adv.AccessToken, adv.AdvertiserID, 1, 100)

	// 统计活跃服务和即将过期
	var activeServices, expireSoon int
	var totalSpend int64

	for _, order := range orders {
		if order.OrderStatus == 1 { // 活跃状态
			activeServices++
		}
		// 简单计算消费
		totalSpend += order.OrderAmount
	}

	// 获取功能点数据
	funcs, _, _ := h.client.ServeMarket().GetFuncPointList(ctx, adv.AccessToken, adv.AdvertiserID, 1, 100)
	for _, f := range funcs {
		if f.Status == 1 {
			activeServices++
		}
		// 检查即将过期 (这里简化处理)
		if f.ExpireTime != "" {
			expireSoon++
		}
	}

	response.OKWithData(c, gin.H{
		"total_orders":    totalOrders,
		"active_services": activeServices,
		"expire_soon":     expireSoon,
		"total_spend":     float64(totalSpend) / 100, // 分转元
	})
}
