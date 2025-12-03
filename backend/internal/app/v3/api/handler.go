package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// V3Handler V3体验版处理器
type V3Handler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewV3Handler 创建V3处理器
func NewV3Handler(db *gorm.DB, oceanCfg *config.OceanConfig) *V3Handler {
	return &V3Handler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *V3Handler) getAccessToken(c *gin.Context) string {
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
func (h *V3Handler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// ==================== 项目管理 ====================

// GetProjectList 获取项目列表
func (h *V3Handler) GetProjectList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.Query("status")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3ProjectListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		PageSize:     pageSize,
		Status:       status,
	}

	list, total, err := h.client.V3().GetProjectList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetProjectDetail 获取项目详情
func (h *V3Handler) GetProjectDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	projectIDStr := c.Param("project_id")
	projectID, _ := strconv.ParseUint(projectIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3ProjectListRequest{
		AdvertiserID: advertiserID,
		ProjectIDs:   []uint64{projectID},
		Page:         1,
		PageSize:     1,
	}

	list, _, err := h.client.V3().GetProjectList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	if len(list) == 0 {
		response.NotFound(c, "项目不存在")
		return
	}

	response.OKWithData(c, list[0])
}

// CreateProject 创建项目
func (h *V3Handler) CreateProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.ProjectCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	projectID, err := h.client.V3().CreateProject(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"project_id": projectID})
}

// UpdateProject 更新项目
func (h *V3Handler) UpdateProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	projectIDStr := c.Param("project_id")
	projectID, _ := strconv.ParseUint(projectIDStr, 10, 64)

	var req oceanengine.ProjectUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ProjectID = projectID

	if accessToken == "" || req.AdvertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().UpdateProject(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdateProjectStatus 更新项目状态
func (h *V3Handler) UpdateProjectStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		ProjectIDs   []uint64 `json:"project_ids"`
		OptStatus    string   `json:"opt_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.ProjectIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().UpdateProjectStatus(c.Request.Context(), accessToken, req.AdvertiserID, req.ProjectIDs, req.OptStatus)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// DeleteProject 删除项目
func (h *V3Handler) DeleteProject(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		ProjectIDs   []uint64 `json:"project_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.ProjectIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().DeleteProjects(c.Request.Context(), accessToken, req.AdvertiserID, req.ProjectIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdateProjectBudget 更新项目预算
func (h *V3Handler) UpdateProjectBudget(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		ProjectIDs   []uint64                 `json:"project_ids"`
		Budget       float64                  `json:"budget"`
		Data         []map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	// 如果传入了简化参数,构建data
	data := req.Data
	if len(data) == 0 && len(req.ProjectIDs) > 0 {
		data = make([]map[string]interface{}, len(req.ProjectIDs))
		for i, pid := range req.ProjectIDs {
			data[i] = map[string]interface{}{
				"project_id": pid,
				"budget":     req.Budget,
			}
		}
	}

	result, err := h.client.V3().UpdateProjectBudget(c.Request.Context(), accessToken, req.AdvertiserID, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// ==================== 广告管理 ====================

// GetPromotionList 获取广告列表
func (h *V3Handler) GetPromotionList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.Query("status")
	projectIDStr := c.Query("project_id")
	projectID, _ := strconv.ParseUint(projectIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3PromotionListRequest{
		AdvertiserID: advertiserID,
		ProjectID:    projectID,
		Status:       status,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.client.V3().GetPromotionList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// GetPromotionDetail 获取广告详情
func (h *V3Handler) GetPromotionDetail(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDStr := c.Param("promotion_id")
	promotionID, _ := strconv.ParseUint(promotionIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3PromotionListRequest{
		AdvertiserID: advertiserID,
		PromotionIDs: []uint64{promotionID},
		Page:         1,
		PageSize:     1,
	}

	list, _, err := h.client.V3().GetPromotionList(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	if len(list) == 0 {
		response.NotFound(c, "广告不存在")
		return
	}

	response.OKWithData(c, list[0])
}

// CreatePromotion 创建广告
func (h *V3Handler) CreatePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.PromotionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	promotionID, err := h.client.V3().CreatePromotion(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"promotion_id": promotionID})
}

// UpdatePromotion 更新广告
func (h *V3Handler) UpdatePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	promotionIDStr := c.Param("promotion_id")
	promotionID, _ := strconv.ParseUint(promotionIDStr, 10, 64)

	var req oceanengine.PromotionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.PromotionID = promotionID

	if accessToken == "" || req.AdvertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().UpdatePromotion(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionStatus 更新广告状态
func (h *V3Handler) UpdatePromotionStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		PromotionIDs []uint64 `json:"promotion_ids"`
		OptStatus    string   `json:"opt_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.PromotionIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().UpdatePromotionStatus(c.Request.Context(), accessToken, req.AdvertiserID, req.PromotionIDs, req.OptStatus)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// DeletePromotion 删除广告
func (h *V3Handler) DeletePromotion(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		PromotionIDs []uint64 `json:"promotion_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.PromotionIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().DeletePromotions(c.Request.Context(), accessToken, req.AdvertiserID, req.PromotionIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionBudget 更新广告预算
func (h *V3Handler) UpdatePromotionBudget(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		PromotionIDs []uint64                 `json:"promotion_ids"`
		Budget       float64                  `json:"budget"`
		Data         []map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	data := req.Data
	if len(data) == 0 && len(req.PromotionIDs) > 0 {
		data = make([]map[string]interface{}, len(req.PromotionIDs))
		for i, pid := range req.PromotionIDs {
			data[i] = map[string]interface{}{
				"promotion_id": pid,
				"budget":       req.Budget,
			}
		}
	}

	result, err := h.client.V3().UpdatePromotionBudget(c.Request.Context(), accessToken, req.AdvertiserID, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionBid 更新广告出价
func (h *V3Handler) UpdatePromotionBid(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		PromotionIDs []uint64                 `json:"promotion_ids"`
		Bid          float64                  `json:"bid"`
		Data         []map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	data := req.Data
	if len(data) == 0 && len(req.PromotionIDs) > 0 {
		data = make([]map[string]interface{}, len(req.PromotionIDs))
		for i, pid := range req.PromotionIDs {
			data[i] = map[string]interface{}{
				"promotion_id": pid,
				"bid":          req.Bid,
			}
		}
	}

	result, err := h.client.V3().UpdatePromotionBid(c.Request.Context(), accessToken, req.AdvertiserID, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionDeepBid 更新广告深度出价
func (h *V3Handler) UpdatePromotionDeepBid(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		PromotionIDs []uint64                 `json:"promotion_ids"`
		DeepCpaBid   float64                  `json:"deep_cpabid"`
		RoiGoal      float64                  `json:"roi_goal"`
		Data         []map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	data := req.Data
	if len(data) == 0 && len(req.PromotionIDs) > 0 {
		data = make([]map[string]interface{}, len(req.PromotionIDs))
		for i, pid := range req.PromotionIDs {
			item := map[string]interface{}{
				"promotion_id": pid,
			}
			if req.DeepCpaBid > 0 {
				item["deep_cpabid"] = req.DeepCpaBid
			}
			if req.RoiGoal > 0 {
				item["roi_goal"] = req.RoiGoal
			}
			data[i] = item
		}
	}

	result, err := h.client.V3().UpdatePromotionDeepBid(c.Request.Context(), accessToken, req.AdvertiserID, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionScheduleTime 更新广告投放时段
func (h *V3Handler) UpdatePromotionScheduleTime(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		PromotionIDs []uint64 `json:"promotion_ids"`
		ScheduleTime string   `json:"schedule_time"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.PromotionIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	result, err := h.client.V3().UpdatePromotionScheduleTime(c.Request.Context(), accessToken, req.AdvertiserID, req.PromotionIDs, req.ScheduleTime)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// UpdatePromotionMaterialStatus 更新广告素材状态
func (h *V3Handler) UpdatePromotionMaterialStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		PromotionID  uint64                   `json:"promotion_id"`
		MaterialIDs  []string                 `json:"material_ids"`
		Status       string                   `json:"status"`
		Data         []map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	data := req.Data
	if len(data) == 0 && req.PromotionID > 0 && len(req.MaterialIDs) > 0 {
		data = []map[string]interface{}{
			{
				"promotion_id": req.PromotionID,
				"material_ids": req.MaterialIDs,
				"status":       req.Status,
			},
		}
	}

	result, err := h.client.V3().UpdatePromotionMaterialStatus(c.Request.Context(), accessToken, req.AdvertiserID, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, result)
}

// GetPromotionRejectReason 获取广告审核建议
func (h *V3Handler) GetPromotionRejectReason(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDsStr := c.QueryArray("promotion_ids")

	if accessToken == "" || advertiserID == 0 || len(promotionIDsStr) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	promotionIDs := make([]uint64, len(promotionIDsStr))
	for i, s := range promotionIDsStr {
		promotionIDs[i], _ = strconv.ParseUint(s, 10, 64)
	}

	list, err := h.client.V3().GetPromotionRejectReason(c.Request.Context(), accessToken, advertiserID, promotionIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetPromotionCostProtectStatus 获取广告成本保障状态
func (h *V3Handler) GetPromotionCostProtectStatus(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDsStr := c.QueryArray("promotion_ids")

	if accessToken == "" || advertiserID == 0 || len(promotionIDsStr) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	promotionIDs := make([]uint64, len(promotionIDsStr))
	for i, s := range promotionIDsStr {
		promotionIDs[i], _ = strconv.ParseUint(s, 10, 64)
	}

	list, err := h.client.V3().GetPromotionCostProtectStatus(c.Request.Context(), accessToken, advertiserID, promotionIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 预算组管理 ====================

// GetBudgetGroupList 获取预算组列表
func (h *V3Handler) GetBudgetGroupList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.V3().GetBudgetGroupList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateBudgetGroup 创建预算组
func (h *V3Handler) CreateBudgetGroup(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID    uint64   `json:"advertiser_id"`
		BudgetGroupName string   `json:"budget_group_name"`
		Budget          float64  `json:"budget"`
		BudgetMode      string   `json:"budget_mode"`
		ProjectIDs      []uint64 `json:"project_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	budgetGroupID, err := h.client.V3().CreateBudgetGroup(c.Request.Context(), accessToken, req.AdvertiserID, req.BudgetGroupName, req.Budget, req.BudgetMode, req.ProjectIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"budget_group_id": budgetGroupID})
}

// UpdateBudgetGroup 更新预算组
func (h *V3Handler) UpdateBudgetGroup(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	budgetGroupIDStr := c.Param("budget_group_id")
	budgetGroupID, _ := strconv.ParseUint(budgetGroupIDStr, 10, 64)

	var req struct {
		AdvertiserID    uint64   `json:"advertiser_id"`
		BudgetGroupName string   `json:"budget_group_name"`
		Budget          float64  `json:"budget"`
		ProjectIDs      []uint64 `json:"project_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || budgetGroupID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	resultID, err := h.client.V3().UpdateBudgetGroup(c.Request.Context(), accessToken, req.AdvertiserID, budgetGroupID, req.BudgetGroupName, req.Budget, req.ProjectIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"budget_group_id": resultID})
}

// DeleteBudgetGroup 删除预算组
func (h *V3Handler) DeleteBudgetGroup(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID   uint64   `json:"advertiser_id"`
		BudgetGroupIDs []uint64 `json:"budget_group_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.BudgetGroupIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().DeleteBudgetGroups(c.Request.Context(), accessToken, req.AdvertiserID, req.BudgetGroupIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_ids": successIDs,
		"fail_ids":    failIDs,
	})
}

// ==================== 数据报表 ====================

// GetProjectReport 获取项目报表
func (h *V3Handler) GetProjectReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || startDate == "" || endDate == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3ReportRequest{
		AdvertiserID: advertiserID,
		StartDate:    startDate,
		EndDate:      endDate,
		Page:         page,
		PageSize:     pageSize,
	}

	data, err := h.client.V3().GetProjectReport(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, data)
}

// GetPromotionReport 获取广告报表
func (h *V3Handler) GetPromotionReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || startDate == "" || endDate == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3ReportRequest{
		AdvertiserID: advertiserID,
		StartDate:    startDate,
		EndDate:      endDate,
		Page:         page,
		PageSize:     pageSize,
	}

	data, err := h.client.V3().GetPromotionReport(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, data)
}

// GetMaterialReport 获取素材报表
func (h *V3Handler) GetMaterialReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || startDate == "" || endDate == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	req := &oceanengine.V3ReportRequest{
		AdvertiserID: advertiserID,
		StartDate:    startDate,
		EndDate:      endDate,
		Page:         page,
		PageSize:     pageSize,
	}

	data, err := h.client.V3().GetMaterialReport(c.Request.Context(), accessToken, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, data)
}

// GetCustomReport 获取自定义报表
func (h *V3Handler) GetCustomReport(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.V3ReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	data, err := h.client.V3().GetCustomReport(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, data)
}

// GetCustomReportConfig 获取自定义报表可用指标和维度
func (h *V3Handler) GetCustomReportConfig(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.V3().GetCustomConfigFields(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 白盒配置 ====================

// SaveAutoGenerateConfig 保存白盒配置
func (h *V3Handler) SaveAutoGenerateConfig(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                 `json:"advertiser_id"`
		PromotionID  uint64                 `json:"promotion_id"`
		Config       map[string]interface{} `json:"config"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.PromotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	configID, err := h.client.V3().CreateAutoGenerateConfig(c.Request.Context(), accessToken, req.AdvertiserID, req.PromotionID, req.Config)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"config_id": configID})
}

// GetAutoGenerateConfig 获取白盒配置详情
func (h *V3Handler) GetAutoGenerateConfig(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDStr := c.Query("promotion_id")
	promotionID, _ := strconv.ParseUint(promotionIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	config, err := h.client.V3().GetAutoGenerateConfig(c.Request.Context(), accessToken, advertiserID, promotionID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, config)
}

// ==================== 搜索广告工具 ====================

// GetBlueFlowPackages 获取蓝海流量包
func (h *V3Handler) GetBlueFlowPackages(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.V3().GetBlueFlowPackages(c.Request.Context(), accessToken, advertiserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetBlueFlowKeywords 获取广告下可用蓝海关键词
func (h *V3Handler) GetBlueFlowKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDStr := c.Query("promotion_id")
	promotionID, _ := strconv.ParseUint(promotionIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.V3().GetBlueFlowKeywords(c.Request.Context(), accessToken, advertiserID, promotionID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// GetSuggestKeywords 获取推荐关键词
func (h *V3Handler) GetSuggestKeywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	queryWord := c.Query("query_word")

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.V3().GetSuggestKeywords(c.Request.Context(), accessToken, advertiserID, queryWord)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// ==================== 关键词管理 ====================

// GetV3Keywords 获取关键词列表
func (h *V3Handler) GetV3Keywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	promotionIDStr := c.Query("promotion_id")
	promotionID, _ := strconv.ParseUint(promotionIDStr, 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || promotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.V3().GetV3Keywords(c.Request.Context(), accessToken, advertiserID, promotionID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateV3Keywords 创建关键词
func (h *V3Handler) CreateV3Keywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		PromotionID  uint64                   `json:"promotion_id"`
		Keywords     []map[string]interface{} `json:"keywords"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.PromotionID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().CreateV3Keywords(c.Request.Context(), accessToken, req.AdvertiserID, req.PromotionID, req.Keywords)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_keyword_ids": successIDs,
		"fail_keyword_ids":    failIDs,
	})
}

// UpdateV3Keywords 更新关键词
func (h *V3Handler) UpdateV3Keywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		Keywords     []map[string]interface{} `json:"keywords"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().UpdateV3Keywords(c.Request.Context(), accessToken, req.AdvertiserID, req.Keywords)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_keyword_ids": successIDs,
		"fail_keyword_ids":    failIDs,
	})
}

// DeleteV3Keywords 删除关键词
func (h *V3Handler) DeleteV3Keywords(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		KeywordIDs   []uint64 `json:"keyword_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || len(req.KeywordIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().DeleteV3Keywords(c.Request.Context(), accessToken, req.AdvertiserID, req.KeywordIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_keyword_ids": successIDs,
		"fail_keyword_ids":    failIDs,
	})
}

// ==================== 否定词管理 ====================

// GetV3PrivativeWords 获取否定词列表
func (h *V3Handler) GetV3PrivativeWords(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	projectIDStr := c.Query("project_id")
	projectID, _ := strconv.ParseUint(projectIDStr, 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || projectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.V3().GetV3PrivativeWords(c.Request.Context(), accessToken, advertiserID, projectID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// AddV3PrivativeWords 添加否定词
func (h *V3Handler) AddV3PrivativeWords(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		ProjectID    uint64                   `json:"project_id"`
		Words        []map[string]interface{} `json:"words"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.ProjectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().AddV3PrivativeWords(c.Request.Context(), accessToken, req.AdvertiserID, req.ProjectID, req.Words)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_word_ids": successIDs,
		"fail_word_ids":    failIDs,
	})
}

// UpdateV3PrivativeWords 更新否定词
func (h *V3Handler) UpdateV3PrivativeWords(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                   `json:"advertiser_id"`
		ProjectID    uint64                   `json:"project_id"`
		Words        []map[string]interface{} `json:"words"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.ProjectID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	successIDs, failIDs, err := h.client.V3().UpdateV3PrivativeWords(c.Request.Context(), accessToken, req.AdvertiserID, req.ProjectID, req.Words)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"success_word_ids": successIDs,
		"fail_word_ids":    failIDs,
	})
}
