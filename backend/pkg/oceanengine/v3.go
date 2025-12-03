package oceanengine

import (
	"context"
)

// V3Client v3体验版客户端
type V3Client struct {
	client *Client
}

// NewV3Client 创建v3体验版客户端
func NewV3Client(client *Client) *V3Client {
	return &V3Client{client: client}
}

// ==================== 项目管理模块 ====================

// ProjectCreateRequest 创建项目请求
type ProjectCreateRequest struct {
	AdvertiserID    uint64                 `json:"advertiser_id"`
	Name            string                 `json:"name"`
	Operation       string                 `json:"operation"`                   // enable/disable
	LandingType     string                 `json:"landing_type"`                // 推广类型
	DeliveryRange   string                 `json:"delivery_range"`              // 投放范围
	Audience        map[string]interface{} `json:"audience,omitempty"`          // 定向设置
	DeliverySetting map[string]interface{} `json:"delivery_setting"`            // 投放设置
	TrackURLSetting map[string]interface{} `json:"track_url_setting,omitempty"` // 监测链接
}

// ProjectUpdateRequest 修改项目请求
type ProjectUpdateRequest struct {
	AdvertiserID    uint64                 `json:"advertiser_id"`
	ProjectID       uint64                 `json:"project_id"`
	Name            string                 `json:"name,omitempty"`
	Operation       string                 `json:"operation,omitempty"`
	Audience        map[string]interface{} `json:"audience,omitempty"`
	DeliverySetting map[string]interface{} `json:"delivery_setting,omitempty"`
	TrackURLSetting map[string]interface{} `json:"track_url_setting,omitempty"`
}

// ProjectUpdateResult 项目更新结果
type ProjectUpdateResult struct {
	ProjectID  uint64   `json:"project_id"`
	SuccessIDs []uint64 `json:"success_list"`
	FailList   []struct {
		ProjectID uint64 `json:"project_id"`
		Code      int    `json:"code"`
		Message   string `json:"message"`
	} `json:"fail_list"`
}

// Project 项目信息
type Project struct {
	ProjectID       uint64                 `json:"project_id"`
	Name            string                 `json:"name"`
	AdvertiserID    uint64                 `json:"advertiser_id"`
	Status          string                 `json:"status"`
	Operation       string                 `json:"operation"`
	LandingType     string                 `json:"landing_type"`
	DeliveryRange   string                 `json:"delivery_range"`
	Audience        map[string]interface{} `json:"audience"`
	DeliverySetting map[string]interface{} `json:"delivery_setting"`
	CreateTime      string                 `json:"create_time"`
	ModifyTime      string                 `json:"modify_time"`
}

// V3ProjectListRequest 获取项目列表请求
type V3ProjectListRequest struct {
	AdvertiserID    uint64   `json:"advertiser_id"`
	ProjectIDs      []uint64 `json:"project_ids,omitempty"`
	ProjectName     string   `json:"project_name,omitempty"`
	LandingType     string   `json:"landing_type,omitempty"`
	Status          string   `json:"status,omitempty"`
	CreateTimeStart string   `json:"create_time_start,omitempty"`
	CreateTimeEnd   string   `json:"create_time_end,omitempty"`
	Page            int      `json:"page,omitempty"`
	PageSize        int      `json:"page_size,omitempty"`
}

// CreateProject 创建项目
func (v *V3Client) CreateProject(ctx context.Context, accessToken string, req *ProjectCreateRequest) (uint64, error) {
	path := "/v3.0/project/create/"
	var result struct {
		Data struct {
			ProjectID uint64 `json:"project_id"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.ProjectID, nil
}

// UpdateProject 修改项目
func (v *V3Client) UpdateProject(ctx context.Context, accessToken string, req *ProjectUpdateRequest) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/update/"
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetProjectList 获取项目列表
func (v *V3Client) GetProjectList(ctx context.Context, accessToken string, req *V3ProjectListRequest) ([]Project, int, error) {
	path := "/v3.0/project/list/"
	var result struct {
		Data struct {
			List     []Project `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
				TotalPage   int `json:"total_page"`
				Page        int `json:"page"`
				PageSize    int `json:"page_size"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// UpdateProjectStatus 批量更新项目状态
func (v *V3Client) UpdateProjectStatus(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64, optStatus string) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
		"opt_status":    optStatus,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DeleteProjects 批量删除项目
func (v *V3Client) DeleteProjects(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ProjectCostProtectStatus 项目成本保障状态
type ProjectCostProtectStatus struct {
	ProjectID         uint64  `json:"project_id"`
	CostProtectStatus int     `json:"cost_protect_status"`
	CompensateRatio   float64 `json:"compensate_ratio"`
	CompensateAmount  float64 `json:"compensate_amount"`
}

// GetProjectCostProtectStatus 批量获取项目成本保障状态
func (v *V3Client) GetProjectCostProtectStatus(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64) ([]ProjectCostProtectStatus, error) {
	path := "/v3.0/project/cost_protect_status/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
	}
	var result struct {
		Data struct {
			List []ProjectCostProtectStatus `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// UpdateProjectBudget 批量更新项目预算
func (v *V3Client) UpdateProjectBudget(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/budget/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateProjectScheduleTime 批量更新项目投放时间
func (v *V3Client) UpdateProjectScheduleTime(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64, startTime, endTime string) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/schedule_time/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
		"start_time":    startTime,
		"end_time":      endTime,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateProjectWeekSchedule 批量更新项目投放时段
func (v *V3Client) UpdateProjectWeekSchedule(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64, weekSchedule string) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/week_schedule/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
		"week_schedule": weekSchedule,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateProjectRoiGoal 批量修改项目ROI系数
func (v *V3Client) UpdateProjectRoiGoal(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*ProjectUpdateResult, error) {
	path := "/v3.0/project/roi_goal/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data ProjectUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// BudgetGroup 预算组
type BudgetGroup struct {
	BudgetGroupID   uint64   `json:"budget_group_id"`
	BudgetGroupName string   `json:"budget_group_name"`
	Budget          float64  `json:"budget"`
	BudgetMode      string   `json:"budget_mode"`
	ProjectIDs      []uint64 `json:"project_ids"`
}

// CreateBudgetGroup 创建预算组
func (v *V3Client) CreateBudgetGroup(ctx context.Context, accessToken string, advertiserID uint64, name string, budget float64, budgetMode string, projectIDs []uint64) (uint64, error) {
	path := "/v3.0/project/budget_group/create/"
	data := map[string]interface{}{
		"advertiser_id":     advertiserID,
		"budget_group_name": name,
		"budget":            budget,
		"budget_mode":       budgetMode,
		"project_ids":       projectIDs,
	}
	var result struct {
		Data struct {
			BudgetGroupID uint64 `json:"budget_group_id"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.BudgetGroupID, nil
}

// UpdateBudgetGroup 更新预算组
func (v *V3Client) UpdateBudgetGroup(ctx context.Context, accessToken string, advertiserID uint64, budgetGroupID uint64, name string, budget float64, projectIDs []uint64) (uint64, error) {
	path := "/v3.0/project/budget_group/update/"
	data := map[string]interface{}{
		"advertiser_id":     advertiserID,
		"budget_group_id":   budgetGroupID,
		"budget_group_name": name,
		"budget":            budget,
		"project_ids":       projectIDs,
	}
	var result struct {
		Data struct {
			BudgetGroupID uint64 `json:"budget_group_id"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.BudgetGroupID, nil
}

// DeleteBudgetGroups 批量删除预算组
func (v *V3Client) DeleteBudgetGroups(ctx context.Context, accessToken string, advertiserID uint64, budgetGroupIDs []uint64) ([]uint64, []uint64, error) {
	path := "/v3.0/project/budget_group/delete/"
	data := map[string]interface{}{
		"advertiser_id":    advertiserID,
		"budget_group_ids": budgetGroupIDs,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_budget_group_ids"`
			FailIDs    []uint64 `json:"fail_budget_group_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetBudgetGroupList 获取预算组列表
func (v *V3Client) GetBudgetGroupList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]BudgetGroup, int, error) {
	path := "/v3.0/project/budget_group/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []BudgetGroup `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 广告管理模块 ====================

// PromotionCreateRequest 创建广告请求
type PromotionCreateRequest struct {
	AdvertiserID  uint64                 `json:"advertiser_id"`
	ProjectID     uint64                 `json:"project_id"`
	Name          string                 `json:"name,omitempty"`
	Operation     string                 `json:"operation"`        // enable/disable
	Source        string                 `json:"source,omitempty"` // 来源
	VideoMaterial map[string]interface{} `json:"video_material,omitempty"`
	ImageMaterial map[string]interface{} `json:"image_material,omitempty"`
	TitleMaterial map[string]interface{} `json:"title_material,omitempty"`
}

// PromotionUpdateRequest 修改广告请求
type PromotionUpdateRequest struct {
	AdvertiserID  uint64                 `json:"advertiser_id"`
	PromotionID   uint64                 `json:"promotion_id"`
	Name          string                 `json:"name,omitempty"`
	Operation     string                 `json:"operation,omitempty"`
	VideoMaterial map[string]interface{} `json:"video_material,omitempty"`
	ImageMaterial map[string]interface{} `json:"image_material,omitempty"`
	TitleMaterial map[string]interface{} `json:"title_material,omitempty"`
}

// PromotionUpdateResult 广告更新结果
type PromotionUpdateResult struct {
	PromotionID uint64   `json:"promotion_id"`
	SuccessIDs  []uint64 `json:"success_list"`
	FailList    []struct {
		PromotionID uint64 `json:"promotion_id"`
		Code        int    `json:"code"`
		Message     string `json:"message"`
	} `json:"fail_list"`
}

// Promotion 广告信息
type Promotion struct {
	PromotionID   uint64                 `json:"promotion_id"`
	ProjectID     uint64                 `json:"project_id"`
	Name          string                 `json:"name"`
	AdvertiserID  uint64                 `json:"advertiser_id"`
	Status        string                 `json:"status"`
	Operation     string                 `json:"operation"`
	VideoMaterial map[string]interface{} `json:"video_material"`
	ImageMaterial map[string]interface{} `json:"image_material"`
	TitleMaterial map[string]interface{} `json:"title_material"`
	CreateTime    string                 `json:"create_time"`
	ModifyTime    string                 `json:"modify_time"`
}

// V3PromotionListRequest 获取广告列表请求
type V3PromotionListRequest struct {
	AdvertiserID    uint64   `json:"advertiser_id"`
	PromotionIDs    []uint64 `json:"promotion_ids,omitempty"`
	ProjectID       uint64   `json:"project_id,omitempty"`
	PromotionName   string   `json:"promotion_name,omitempty"`
	Status          string   `json:"status,omitempty"`
	CreateTimeStart string   `json:"create_time_start,omitempty"`
	CreateTimeEnd   string   `json:"create_time_end,omitempty"`
	Page            int      `json:"page,omitempty"`
	PageSize        int      `json:"page_size,omitempty"`
}

// CreatePromotion 创建广告
func (v *V3Client) CreatePromotion(ctx context.Context, accessToken string, req *PromotionCreateRequest) (uint64, error) {
	path := "/v3.0/promotion/create/"
	var result struct {
		Data struct {
			PromotionID uint64 `json:"promotion_id"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.PromotionID, nil
}

// UpdatePromotion 修改广告
func (v *V3Client) UpdatePromotion(ctx context.Context, accessToken string, req *PromotionUpdateRequest) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/update/"
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetPromotionList 获取广告列表
func (v *V3Client) GetPromotionList(ctx context.Context, accessToken string, req *V3PromotionListRequest) ([]Promotion, int, error) {
	path := "/v3.0/promotion/list/"
	var result struct {
		Data struct {
			List     []Promotion `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
				TotalPage   int `json:"total_page"`
				Page        int `json:"page"`
				PageSize    int `json:"page_size"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// UpdatePromotionBudget 批量更新广告预算
func (v *V3Client) UpdatePromotionBudget(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/budget/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdatePromotionBid 批量更新广告出价
func (v *V3Client) UpdatePromotionBid(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/bid/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdatePromotionStatus 批量更新广告启用状态
func (v *V3Client) UpdatePromotionStatus(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64, optStatus string) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
		"opt_status":    optStatus,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DeletePromotions 批量删除广告
func (v *V3Client) DeletePromotions(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// RejectReason 审核建议
type RejectReason struct {
	PromotionID  uint64 `json:"promotion_id"`
	RejectReason string `json:"reject_reason"`
	RejectItem   string `json:"reject_item"`
}

// GetPromotionRejectReason 获取计划审核建议
func (v *V3Client) GetPromotionRejectReason(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64) ([]RejectReason, error) {
	path := "/v3.0/promotion/reject_reason/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
	}
	var result struct {
		Data struct {
			List []RejectReason `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// UpdatePromotionMaterialStatus 批量更新广告素材启用状态
func (v *V3Client) UpdatePromotionMaterialStatus(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/material/status/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdatePromotionDeepBid 批量修改深度出价
func (v *V3Client) UpdatePromotionDeepBid(ctx context.Context, accessToken string, advertiserID uint64, data []map[string]interface{}) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/deep_bid/update/"
	reqData := map[string]interface{}{
		"advertiser_id": advertiserID,
		"data":          data,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, reqData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// PromotionCostProtectStatus 广告成本保障状态
type PromotionCostProtectStatus struct {
	PromotionID       uint64  `json:"promotion_id"`
	CostProtectStatus int     `json:"cost_protect_status"`
	CompensateRatio   float64 `json:"compensate_ratio"`
	CompensateAmount  float64 `json:"compensate_amount"`
}

// GetPromotionCostProtectStatus 批量获取计划成本保障状态
func (v *V3Client) GetPromotionCostProtectStatus(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64) ([]PromotionCostProtectStatus, error) {
	path := "/v3.0/promotion/cost_protect_status/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
	}
	var result struct {
		Data struct {
			List []PromotionCostProtectStatus `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// AutoGenerateConfig 白盒配置
type AutoGenerateConfig struct {
	ConfigID       uint64                 `json:"config_id"`
	ConfigName     string                 `json:"config_name"`
	Status         int                    `json:"status"`
	GenerateConfig map[string]interface{} `json:"generate_config"`
}

// CreateAutoGenerateConfig 新建/修改白盒配置
func (v *V3Client) CreateAutoGenerateConfig(ctx context.Context, accessToken string, advertiserID, promotionID uint64, config map[string]interface{}) (uint64, error) {
	path := "/v3.0/promotion/auto_generate_config/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"config":        config,
	}
	var result struct {
		Data struct {
			ConfigID uint64 `json:"config_id"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.ConfigID, nil
}

// GetAutoGenerateConfig 查询配置详情
func (v *V3Client) GetAutoGenerateConfig(ctx context.Context, accessToken string, advertiserID, promotionID uint64) (*AutoGenerateConfig, error) {
	path := "/v3.0/promotion/auto_generate_config/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
	}
	var result struct {
		Data AutoGenerateConfig `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdatePromotionScheduleTime 批量更新广告投放时段
func (v *V3Client) UpdatePromotionScheduleTime(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64, scheduleTime string) (*PromotionUpdateResult, error) {
	path := "/v3.0/promotion/schedule_time/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
		"schedule_time": scheduleTime,
	}
	var result struct {
		Data PromotionUpdateResult `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== v3报表模块 ====================

// V3ReportRequest 报表请求
type V3ReportRequest struct {
	AdvertiserID uint64   `json:"advertiser_id"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Fields       []string `json:"fields,omitempty"`
	GroupBy      []string `json:"group_by,omitempty"`
	OrderField   string   `json:"order_field,omitempty"`
	OrderType    string   `json:"order_type,omitempty"`
	Page         int      `json:"page,omitempty"`
	PageSize     int      `json:"page_size,omitempty"`
}

// V3ReportData 报表数据
type V3ReportData struct {
	List     []map[string]interface{} `json:"list"`
	PageInfo struct {
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
	} `json:"page_info"`
}

// GetProjectReport 项目数据报表
func (v *V3Client) GetProjectReport(ctx context.Context, accessToken string, req *V3ReportRequest) (*V3ReportData, error) {
	path := "/v3.0/report/project/get/"
	var result struct {
		Data V3ReportData `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_date":    req.StartDate,
		"end_date":      req.EndDate,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetPromotionReport 广告数据报表
func (v *V3Client) GetPromotionReport(ctx context.Context, accessToken string, req *V3ReportRequest) (*V3ReportData, error) {
	path := "/v3.0/report/promotion/get/"
	var result struct {
		Data V3ReportData `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_date":    req.StartDate,
		"end_date":      req.EndDate,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetMaterialReport 素材数据报表
func (v *V3Client) GetMaterialReport(ctx context.Context, accessToken string, req *V3ReportRequest) (*V3ReportData, error) {
	path := "/v3.0/report/material/get/"
	var result struct {
		Data V3ReportData `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_date":    req.StartDate,
		"end_date":      req.EndDate,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetCustomReport 自定义报表
func (v *V3Client) GetCustomReport(ctx context.Context, accessToken string, req *V3ReportRequest) (*V3ReportData, error) {
	path := "/v3.0/report/custom/get/"
	var result struct {
		Data V3ReportData `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// CustomConfigField 自定义报表字段
type CustomConfigField struct {
	FieldName   string `json:"field_name"`
	FieldDesc   string `json:"field_desc"`
	FieldType   string `json:"field_type"`
	IsDimension bool   `json:"is_dimension"`
	IsMetric    bool   `json:"is_metric"`
}

// GetCustomConfigFields 获取自定义报表可用指标和维度
func (v *V3Client) GetCustomConfigFields(ctx context.Context, accessToken string, advertiserID uint64) ([]CustomConfigField, error) {
	path := "/v3.0/report/custom/config/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []CustomConfigField `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 搜索广告工具 ====================
// 注: BlueFlowPackage, BlueFlowKeyword, SuggestKeyword 类型定义在 searchad.go 中

// GetBlueFlowPackages 获取蓝海流量包
func (v *V3Client) GetBlueFlowPackages(ctx context.Context, accessToken string, advertiserID uint64) ([]BlueFlowPackage, error) {
	path := "/v3.0/blueflow/package/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []BlueFlowPackage `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetBlueFlowKeywords 获取广告下可用蓝海关键词
func (v *V3Client) GetBlueFlowKeywords(ctx context.Context, accessToken string, advertiserID, promotionID uint64) ([]BlueFlowKeyword, error) {
	path := "/v3.0/blueflow/keyword/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
	}
	var result struct {
		Data struct {
			List []BlueFlowKeyword `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetSuggestKeywords 获取推荐关键词
func (v *V3Client) GetSuggestKeywords(ctx context.Context, accessToken string, advertiserID uint64, queryWord string) ([]SuggestKeyword, error) {
	path := "/v3.0/keyword/suggest/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    queryWord,
	}
	var result struct {
		Data struct {
			List []SuggestKeyword `json:"list"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// V3Keyword 关键词
type V3Keyword struct {
	KeywordID uint64  `json:"keyword_id"`
	Word      string  `json:"word"`
	MatchType string  `json:"match_type"`
	Bid       float64 `json:"bid"`
	Status    string  `json:"status"`
}

// CreateV3Keywords 体验版创建关键词
func (v *V3Client) CreateV3Keywords(ctx context.Context, accessToken string, advertiserID, promotionID uint64, keywords []map[string]interface{}) ([]uint64, []uint64, error) {
	path := "/v3.0/keyword/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"keywords":      keywords,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_keyword_ids"`
			FailIDs    []uint64 `json:"fail_keyword_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateV3Keywords 体验版更新关键词属性
func (v *V3Client) UpdateV3Keywords(ctx context.Context, accessToken string, advertiserID uint64, keywords []map[string]interface{}) ([]uint64, []uint64, error) {
	path := "/v3.0/keyword/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"keywords":      keywords,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_keyword_ids"`
			FailIDs    []uint64 `json:"fail_keyword_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// DeleteV3Keywords 体验版删除关键词
func (v *V3Client) DeleteV3Keywords(ctx context.Context, accessToken string, advertiserID uint64, keywordIDs []uint64) ([]uint64, []uint64, error) {
	path := "/v3.0/keyword/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"keyword_ids":   keywordIDs,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_keyword_ids"`
			FailIDs    []uint64 `json:"fail_keyword_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetV3Keywords 体验版获取关键词列表
func (v *V3Client) GetV3Keywords(ctx context.Context, accessToken string, advertiserID, promotionID uint64, page, pageSize int) ([]V3Keyword, int, error) {
	path := "/v3.0/keyword/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []V3Keyword `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 否定词管理v3 ====================

// V3PrivativeWord 否定词
type V3PrivativeWord struct {
	WordID    uint64 `json:"word_id"`
	Word      string `json:"word"`
	MatchType string `json:"match_type"`
}

// AddV3PrivativeWords 2.0项目批量新增否定词
func (v *V3Client) AddV3PrivativeWords(ctx context.Context, accessToken string, advertiserID, projectID uint64, words []map[string]interface{}) ([]uint64, []uint64, error) {
	path := "/v3.0/privativeword/add/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"words":         words,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_word_ids"`
			FailIDs    []uint64 `json:"fail_word_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateV3PrivativeWords 2.0项目批量更新否定词
func (v *V3Client) UpdateV3PrivativeWords(ctx context.Context, accessToken string, advertiserID, projectID uint64, words []map[string]interface{}) ([]uint64, []uint64, error) {
	path := "/v3.0/privativeword/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"words":         words,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_word_ids"`
			FailIDs    []uint64 `json:"fail_word_ids"`
		} `json:"data"`
	}
	err := v.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetV3PrivativeWords 2.0项目批量获取否定词
func (v *V3Client) GetV3PrivativeWords(ctx context.Context, accessToken string, advertiserID, projectID uint64, page, pageSize int) ([]V3PrivativeWord, int, error) {
	path := "/v3.0/privativeword/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []V3PrivativeWord `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := v.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}
