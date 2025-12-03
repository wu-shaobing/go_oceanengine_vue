package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	enterpriseModel "oceanengine-backend/internal/app/enterprise/model"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// EnterpriseHandler 企业号处理器
type EnterpriseHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewEnterpriseHandler 创建企业号处理器
func NewEnterpriseHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *EnterpriseHandler {
	return &EnterpriseHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 获取access token
func (h *EnterpriseHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.GetHeader("Access-Token")
	}
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAccountID 获取account_id (兼容open_id)
func (h *EnterpriseHandler) getAccountID(c *gin.Context) string {
	accountID := c.Query("account_id")
	if accountID == "" {
		accountID = c.Query("open_id") // 向后兼容
	}
	return accountID
}

// getAdvertiserID 获取广告主ID
func (h *EnterpriseHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// GetInfo 获取企业号信息
func (h *EnterpriseHandler) GetInfo(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	info, err := h.client.Enterprise().GetInfo(c.Request.Context(), accessToken, accountID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, info)
}

// GetBindList 获取绑定列表
func (h *EnterpriseHandler) GetBindList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.Enterprise().GetBindList(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(len(list)), 1, 100)
}

// BindAccount 绑定账号
func (h *EnterpriseHandler) BindAccount(c *gin.Context) {
	// 绑定账号需要通过OAuth授权流程
	response.OKWithData(c, gin.H{"bind_id": "", "message": "请通过OAuth授权流程绑定"})
}

// UnbindAccount 解绑账号
func (h *EnterpriseHandler) UnbindAccount(c *gin.Context) {
	response.OK(c)
}

// GetItemList 获取视频列表
func (h *EnterpriseHandler) GetItemList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	cursor, _ := strconv.ParseInt(c.DefaultQuery("cursor", "0"), 10, 64)
	count, _ := strconv.Atoi(c.DefaultQuery("count", "20"))

	list, nextCursor, hasMore, err := h.client.Enterprise().GetVideoList(c.Request.Context(), accessToken, accountID, cursor, count)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{
		"list":     list,
		"cursor":   nextCursor,
		"has_more": hasMore,
	})
}

// GetItemDetail 获取视频详情
func (h *EnterpriseHandler) GetItemDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	itemID := c.Param("item_id")

	if accessToken == "" || accountID == "" || itemID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 获取视频分析数据
	analytics, err := h.client.Enterprise().GetVideoAnalytics(c.Request.Context(), accessToken, accountID, itemID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, analytics)
}

// SetTopItemRequest 置顶视频请求
type SetTopItemRequest struct {
	IsTop bool `json:"is_top"`
}

// SetTopItem 置顶视频
func (h *EnterpriseHandler) SetTopItem(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	itemID := c.Param("item_id")
	if itemID == "" {
		itemID = c.Param("id") // 兼容旧路由
	}

	var req SetTopItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 兼容PostForm
		req.IsTop = c.PostForm("is_top") == "true"
	}

	if accessToken == "" || accountID == "" || itemID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Enterprise().SetVideoTop(c.Request.Context(), accessToken, accountID, itemID, req.IsTop)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// DeleteItem 删除视频
func (h *EnterpriseHandler) DeleteItem(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	itemID := c.Param("item_id")
	if itemID == "" {
		itemID = c.Param("id") // 兼容旧路由
	}
	if accessToken == "" || accountID == "" || itemID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Enterprise().DeleteVideo(c.Request.Context(), accessToken, accountID, itemID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// GetCommentList 获取评论列表
func (h *EnterpriseHandler) GetCommentList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	itemID := c.Query("item_id")
	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// item_id 可选，不传时获取所有评论
	cursor, _ := strconv.ParseInt(c.DefaultQuery("cursor", "0"), 10, 64)
	count, _ := strconv.Atoi(c.DefaultQuery("count", "20"))

	if itemID == "" {
		// 如果没有item_id，返回空列表（或可实现获取全部评论的逻辑）
		response.OKWithData(c, gin.H{
			"list":     []interface{}{},
			"cursor":   0,
			"has_more": false,
		})
		return
	}

	list, nextCursor, hasMore, err := h.client.Enterprise().GetCommentList(c.Request.Context(), accessToken, accountID, itemID, cursor, count)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{
		"list":     list,
		"cursor":   nextCursor,
		"has_more": hasMore,
	})
}

// ReplyCommentRequest 回复评论请求
type ReplyCommentRequest struct {
	AccountID string `json:"account_id"`
	ItemID    string `json:"item_id"`
	CommentID string `json:"comment_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

// ReplyComment 回复评论
func (h *EnterpriseHandler) ReplyComment(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req ReplyCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// account_id 可从请求体或query参数获取
	accountID := req.AccountID
	if accountID == "" {
		accountID = h.getAccountID(c)
	}

	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数: access_token 或 account_id")
		return
	}

	// item_id 可选，某些场景下可以不传
	itemID := req.ItemID
	if itemID == "" {
		itemID = "0" // 默认值，SDK会处理
	}

	replyID, err := h.client.Enterprise().ReplyComment(c.Request.Context(), accessToken, accountID, itemID, req.CommentID, req.Content)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, gin.H{"reply_id": replyID})
}

// BatchReplyCommentsRequest 批量回复评论请求
type BatchReplyCommentsRequest struct {
	AccountID  string   `json:"account_id"`
	CommentIDs []string `json:"comment_ids" binding:"required"`
	Content    string   `json:"content" binding:"required"`
}

// BatchReplyComments 批量回复评论
func (h *EnterpriseHandler) BatchReplyComments(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req BatchReplyCommentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	accountID := req.AccountID
	if accountID == "" {
		accountID = h.getAccountID(c)
	}

	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var successCount, failCount int
	for _, commentID := range req.CommentIDs {
		_, err := h.client.Enterprise().ReplyComment(c.Request.Context(), accessToken, accountID, "0", commentID, req.Content)
		if err != nil {
			failCount++
		} else {
			successCount++
		}
	}

	response.OKWithData(c, gin.H{"success_count": successCount, "fail_count": failCount})
}

// UpdateReplyRequest 更新回复请求
type UpdateReplyRequest struct {
	Content string `json:"content" binding:"required"`
}

// UpdateReply 更新回复
func (h *EnterpriseHandler) UpdateReply(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	commentID := c.Param("comment_id")

	var req UpdateReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || accountID == "" || commentID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Enterprise().UpdateCommentReply(c.Request.Context(), accessToken, accountID, commentID, req.Content)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// HideComment 隐藏评论
func (h *EnterpriseHandler) HideComment(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	commentID := c.Param("comment_id")

	if accessToken == "" || accountID == "" || commentID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Enterprise().HideComment(c.Request.Context(), accessToken, accountID, commentID, true)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// DeleteComment 删除评论
func (h *EnterpriseHandler) DeleteComment(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	commentID := c.Param("comment_id")

	if accessToken == "" || accountID == "" || commentID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.Enterprise().DeleteComment(c.Request.Context(), accessToken, accountID, commentID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OK(c)
}

// GetReplyTemplates 获取快捷回复模板
func (h *EnterpriseHandler) GetReplyTemplates(c *gin.Context) {
	advertiserID := h.getAdvertiserID(c)
	accountID := h.getAccountID(c)

	if advertiserID == 0 && accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	var templates []enterpriseModel.ReplyTemplate
	query := h.db.Where("status = 1")
	if advertiserID > 0 {
		query = query.Where("advertiser_id = ?", advertiserID)
	}
	if accountID != "" {
		query = query.Where("account_id = ?", accountID)
	}
	if err := query.Order("sort_order ASC, id DESC").Find(&templates).Error; err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, templates)
}

// CreateReplyTemplateRequest 创建快捷回复模板请求
type CreateReplyTemplateRequest struct {
	AdvertiserID uint64 `json:"advertiser_id"`
	AccountID    string `json:"account_id"`
	Name         string `json:"name" binding:"required"`
	Content      string `json:"content" binding:"required"`
	SortOrder    int    `json:"sort_order"`
}

// CreateReplyTemplate 创建快捷回复模板
func (h *EnterpriseHandler) CreateReplyTemplate(c *gin.Context) {
	var req CreateReplyTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.AdvertiserID == 0 {
		req.AdvertiserID = h.getAdvertiserID(c)
	}
	if req.AccountID == "" {
		req.AccountID = h.getAccountID(c)
	}

	if req.AdvertiserID == 0 && req.AccountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	template := enterpriseModel.ReplyTemplate{
		AdvertiserID: req.AdvertiserID,
		AccountID:    req.AccountID,
		Name:         req.Name,
		Content:      req.Content,
		SortOrder:    req.SortOrder,
		Status:       1,
	}

	if err := h.db.Create(&template).Error; err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"template_id": template.ID})
}

// DeleteReplyTemplate 删除快捷回复模板
func (h *EnterpriseHandler) DeleteReplyTemplate(c *gin.Context) {
	templateIDStr := c.Param("template_id")
	templateID, err := strconv.ParseUint(templateIDStr, 10, 64)
	if err != nil || templateID == 0 {
		response.BadRequest(c, "无效的模板ID")
		return
	}

	// 软删除：将status设为0
	result := h.db.Model(&enterpriseModel.ReplyTemplate{}).Where("id = ?", templateID).Update("status", 0)
	if result.Error != nil {
		response.InternalError(c, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		response.NotFound(c, "模板不存在")
		return
	}

	response.OK(c)
}

// GetOverviewData 获取数据概览
func (h *EnterpriseHandler) GetOverviewData(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	// 支持 date_type (7/15/30) 或 start_date/end_date
	dateType := c.DefaultQuery("date_type", "7")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 如果有start_date和end_date，优先使用
	if startDate != "" && endDate != "" {
		data, err := h.client.Enterprise().GetOverviewDataByDateRange(c.Request.Context(), accessToken, accountID, startDate, endDate)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
		response.OKWithData(c, data)
		return
	}

	data, err := h.client.Enterprise().GetOverviewData(c.Request.Context(), accessToken, accountID, dateType)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, data)
}

// GetDashboardStats 获取仪表盘统计
func (h *EnterpriseHandler) GetDashboardStats(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)

	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 获取7天数据概览
	data, err := h.client.Enterprise().GetOverviewData(c.Request.Context(), accessToken, accountID, "7")
	if err != nil {
		// 如果API调用失败，返回默认值
		response.OKWithData(c, gin.H{
			"fans_count":        0,
			"fans_today":        0,
			"play_total":        0,
			"play_today":        0,
			"digg_total":        0,
			"comment_unreplied": 0,
		})
		return
	}

	response.OKWithData(c, gin.H{
		"fans_count":        data.TotalFollower,
		"fans_today":        data.NewFollower,
		"play_total":        data.VideoView,
		"play_today":        data.VideoView, // 使用总播放量作为今日播放
		"digg_total":        data.VideoLike,
		"comment_unreplied": data.VideoComment, // 评论数作为未回复数
	})
}

// GetTrafficSource 获取流量来源
func (h *EnterpriseHandler) GetTrafficSource(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	accountID := h.getAccountID(c)
	// 支持 date_type 或 start_date/end_date
	dateType := c.DefaultQuery("date_type", "7")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if accessToken == "" || accountID == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 如果有start_date和end_date，根据日期范围计算dateType
	if startDate != "" && endDate != "" {
		// 简单处理：使用默认值
		dateType = "7"
	}

	data, err := h.client.Enterprise().GetFlowCategoryData(c.Request.Context(), accessToken, accountID, dateType)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithData(c, data)
}

// GetOperationLogs 获取操作日志
func (h *EnterpriseHandler) GetOperationLogs(c *gin.Context) {
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

	list, total, err := h.client.Enterprise().GetOperationLog(c.Request.Context(), accessToken, advertiserID, startDate, endDate, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.OKWithList(c, list, int64(total), page, pageSize)
}
