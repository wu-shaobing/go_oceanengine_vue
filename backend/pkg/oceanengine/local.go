package oceanengine

import (
	"context"
)

// LocalClient 本地推API客户端
type LocalClient struct {
	client *Client
}

// NewLocalClient 创建本地推客户端
func (c *Client) Local() *LocalClient {
	return &LocalClient{client: c}
}

// ==================== 项目管理 ====================

// LocalProject 本地推项目
type LocalProject struct {
	ProjectID     uint64  `json:"project_id"`
	ProjectName   string  `json:"project_name"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	Status        string  `json:"status"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	DeliveryRange string  `json:"delivery_range"`
	CreateTime    string  `json:"create_time"`
	ModifyTime    string  `json:"modify_time"`
}

// ProjectListRequest 项目列表请求
type ProjectListRequest struct {
	AdvertiserID uint64         `json:"advertiser_id"`
	Filtering    *ProjectFilter `json:"filtering,omitempty"`
	Page         int            `json:"page,omitempty"`
	PageSize     int            `json:"page_size,omitempty"`
}

type ProjectFilter struct {
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	Status     string   `json:"status,omitempty"`
}

// GetProjectList 获取项目列表
func (l *LocalClient) GetProjectList(ctx context.Context, accessToken string, req *ProjectListRequest) ([]LocalProject, int, error) {
	path := "/v1.0/local/project/list/"
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.Filtering != nil {
		params["filtering"] = req.Filtering
	}

	var result struct {
		Data struct {
			List     []LocalProject `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreateProject 创建项目
func (l *LocalClient) CreateProject(ctx context.Context, accessToken string, advertiserID uint64, projectData map[string]interface{}) (uint64, error) {
	path := "/v1.0/local/project/create/"
	projectData["advertiser_id"] = advertiserID

	var result struct {
		Data struct {
			ProjectID uint64 `json:"project_id"`
		} `json:"data"`
	}
	err := l.client.PostWithToken(ctx, accessToken, path, projectData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.ProjectID, nil
}

// GetProjectDetail 获取项目详情
func (l *LocalClient) GetProjectDetail(ctx context.Context, accessToken string, advertiserID, projectID uint64) (*LocalProject, error) {
	path := "/v1.0/local/project/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
	}

	var result struct {
		Data LocalProject `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 广告管理 ====================

// LocalPromotion 本地推广告
type LocalPromotion struct {
	PromotionID   uint64  `json:"promotion_id"`
	PromotionName string  `json:"promotion_name"`
	ProjectID     uint64  `json:"project_id"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	Status        string  `json:"status"`
	OptStatus     string  `json:"opt_status"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	CreateTime    string  `json:"create_time"`
	ModifyTime    string  `json:"modify_time"`
}

// PromotionListRequest 广告列表请求
type PromotionListRequest struct {
	AdvertiserID uint64           `json:"advertiser_id"`
	Filtering    *PromotionFilter `json:"filtering,omitempty"`
	Page         int              `json:"page,omitempty"`
	PageSize     int              `json:"page_size,omitempty"`
}

type PromotionFilter struct {
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	ProjectIDs   []uint64 `json:"project_ids,omitempty"`
	Status       string   `json:"status,omitempty"`
}

// GetPromotionList 获取广告列表
func (l *LocalClient) GetPromotionList(ctx context.Context, accessToken string, req *PromotionListRequest) ([]LocalPromotion, int, error) {
	path := "/v1.0/local/promotion/list/"
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.Filtering != nil {
		params["filtering"] = req.Filtering
	}

	var result struct {
		Data struct {
			List     []LocalPromotion `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreatePromotion 创建广告
func (l *LocalClient) CreatePromotion(ctx context.Context, accessToken string, advertiserID uint64, promotionData map[string]interface{}) (uint64, error) {
	path := "/v1.0/local/promotion/create/"
	promotionData["advertiser_id"] = advertiserID

	var result struct {
		Data struct {
			PromotionID uint64 `json:"promotion_id"`
		} `json:"data"`
	}
	err := l.client.PostWithToken(ctx, accessToken, path, promotionData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.PromotionID, nil
}

// GetPromotionDetail 获取广告详情
func (l *LocalClient) GetPromotionDetail(ctx context.Context, accessToken string, advertiserID, promotionID uint64) (*LocalPromotion, error) {
	path := "/v1.0/local/promotion/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
	}

	var result struct {
		Data LocalPromotion `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdatePromotionStatus 更新广告状态
func (l *LocalClient) UpdatePromotionStatus(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64, optStatus string) error {
	path := "/v1.0/local/promotion/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
		"opt_status":    optStatus,
	}
	return l.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 线索管理 ====================

// LocalClue 本地推线索信息
type LocalClue struct {
	ClueID       uint64 `json:"clue_id"`
	PromotionID  uint64 `json:"promotion_id"`
	ProjectID    uint64 `json:"project_id"`
	ClueType     int    `json:"clue_type"`
	ClueSource   int    `json:"clue_source"`
	CreateTime   string `json:"create_time"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Remark       string `json:"remark"`
	FollowStatus int    `json:"follow_status"`
}

// GetClueList 获取线索列表
func (l *LocalClient) GetClueList(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]LocalClue, int, error) {
	path := "/v1.0/local/clue/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []LocalClue `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 数据报表 ====================

// LocalReport 本地推报表
type LocalReport struct {
	StatDatetime string  `json:"stat_datetime"`
	Cost         float64 `json:"cost"`
	ShowCnt      int64   `json:"show_cnt"`
	ClickCnt     int64   `json:"click_cnt"`
	ConvertCnt   int64   `json:"convert_cnt"`
	Ctr          float64 `json:"ctr"`
	Cpm          float64 `json:"cpm"`
	Cpc          float64 `json:"cpc"`
	CpaConvert   float64 `json:"cpa_convert"`
}

// GetProjectReport 获取项目报表
func (l *LocalClient) GetProjectReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, projectIDs []uint64) ([]LocalReport, error) {
	path := "/v1.0/local/report/project/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(projectIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"project_ids": projectIDs,
		}
	}

	var result struct {
		Data struct {
			List []LocalReport `json:"list"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetPromotionReport 获取广告报表
func (l *LocalClient) GetPromotionReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, promotionIDs []uint64) ([]LocalReport, error) {
	path := "/v1.0/local/report/promotion/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(promotionIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"promotion_ids": promotionIDs,
		}
	}

	var result struct {
		Data struct {
			List []LocalReport `json:"list"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetMaterialReport 获取素材报表
func (l *LocalClient) GetMaterialReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string) ([]LocalReport, error) {
	path := "/v1.0/local/report/material/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}

	var result struct {
		Data struct {
			List []LocalReport `json:"list"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 项目管理扩展 ====================

// UpdateProject 更新项目
func (l *LocalClient) UpdateProject(ctx context.Context, accessToken string, advertiserID, projectID uint64, updateData map[string]interface{}) error {
	path := "/v1.0/local/project/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["project_id"] = projectID
	return l.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// UpdateProjectStatus 更新项目状态
func (l *LocalClient) UpdateProjectStatus(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64, optStatus string) ([]uint64, []uint64, error) {
	path := "/v1.0/local/project/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
		"opt_status":    optStatus,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_project_ids"`
			FailIDs    []uint64 `json:"fail_project_ids"`
		} `json:"data"`
	}
	err := l.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// DeleteProject 删除项目
func (l *LocalClient) DeleteProject(ctx context.Context, accessToken string, advertiserID uint64, projectIDs []uint64) error {
	path := "/v1.0/local/project/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_ids":   projectIDs,
	}
	return l.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 广告管理扩展 ====================

// UpdatePromotion 更新广告
func (l *LocalClient) UpdatePromotion(ctx context.Context, accessToken string, advertiserID, promotionID uint64, updateData map[string]interface{}) error {
	path := "/v1.0/local/promotion/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["promotion_id"] = promotionID
	return l.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// DeletePromotion 删除广告
func (l *LocalClient) DeletePromotion(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64) error {
	path := "/v1.0/local/promotion/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
	}
	return l.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// UpdatePromotionBudget 更新广告预算
func (l *LocalClient) UpdatePromotionBudget(ctx context.Context, accessToken string, advertiserID uint64, promotionIDs []uint64, budget float64) error {
	path := "/v1.0/local/promotion/budget/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_ids": promotionIDs,
		"budget":        budget,
	}
	return l.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 商品与门店管理 ====================

// LocalProduct 本地推商品
type LocalProduct struct {
	ProductID     uint64  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"discount_price"`
	Status        int     `json:"status"`
	Category      string  `json:"category"`
	ImageURL      string  `json:"image_url"`
}

// GetProductList 获取可投商品列表
func (l *LocalClient) GetProductList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LocalProduct, int, error) {
	path := "/v1.0/local/product/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalProduct `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// LocalStore 门店信息
type LocalStore struct {
	PoiID         string  `json:"poi_id"`
	PoiName       string  `json:"poi_name"`
	Address       string  `json:"address"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Phone         string  `json:"phone"`
	BusinessHours string  `json:"business_hours"`
	Status        int     `json:"status"`
}

// GetStoreList 获取门店列表
func (l *LocalClient) GetStoreList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LocalStore, int, error) {
	path := "/v1.0/local/poi/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalStore `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetProductsByPoiIDs 根据门店ID查询商品
func (l *LocalClient) GetProductsByPoiIDs(ctx context.Context, accessToken string, advertiserID uint64, poiIDs []string) ([]uint64, error) {
	path := "/v1.0/local/product/poi/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"poi_ids":       poiIDs,
	}
	var result struct {
		Data struct {
			ProductIDs []uint64 `json:"product_ids"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.ProductIDs, nil
}

// ==================== 抖音号管理 ====================

// LocalAweme 本地推抖音号
type LocalAweme struct {
	AwemeID       uint64 `json:"aweme_id"`
	AwemeName     string `json:"aweme_name"`
	AwemeAvatar   string `json:"aweme_avatar"`
	FollowerCount int64  `json:"follower_count"`
	AuthStatus    int    `json:"auth_status"`
}

// GetAuthorizedAwemeList 获取已授权抖音号列表
func (l *LocalClient) GetAuthorizedAwemeList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LocalAweme, int, error) {
	path := "/v1.0/local/aweme/authorized/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalAweme `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 人群包管理 ====================

// LocalAudience 本地推人群包
type LocalAudience struct {
	AudienceID uint64 `json:"audience_id"`
	Name       string `json:"name"`
	CoverNum   int64  `json:"cover_num"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// GetAudienceList 获取可用人群包列表
func (l *LocalClient) GetAudienceList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LocalAudience, int, error) {
	path := "/v1.0/local/customaudience/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalAudience `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 素材管理 ====================

// LocalVideo 本地推视频素材
type LocalVideo struct {
	VideoID    string  `json:"video_id"`
	VideoURL   string  `json:"video_url"`
	CoverURL   string  `json:"cover_url"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	Duration   float64 `json:"duration"`
	Status     int     `json:"status"`
	CreateTime string  `json:"create_time"`
}

// GetVideoList 获取素材库视频
func (l *LocalClient) GetVideoList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LocalVideo, int, error) {
	path := "/v1.0/local/file/video/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalVideo `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetAwemeVideoList 获取抖音主页视频
func (l *LocalClient) GetAwemeVideoList(ctx context.Context, accessToken string, advertiserID, awemeID uint64, page, pageSize int) ([]LocalVideo, int, error) {
	path := "/v1.0/local/file/video/aweme/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LocalVideo `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreateVideoUploadTask 创建异步视频上传任务
func (l *LocalClient) CreateVideoUploadTask(ctx context.Context, accessToken string, advertiserID uint64, videoURL string) (uint64, error) {
	path := "/v1.0/local/file/video/upload/task/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"video_url":     videoURL,
	}
	var result struct {
		Data struct {
			TaskID uint64 `json:"task_id"`
		} `json:"data"`
	}
	err := l.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.TaskID, nil
}

// GetVideoUploadTaskList 查询视频上传任务结果
func (l *LocalClient) GetVideoUploadTaskList(ctx context.Context, accessToken string, advertiserID uint64, taskIDs []uint64) ([]map[string]interface{}, error) {
	path := "/v1.0/local/file/video/upload/task/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_ids":      taskIDs,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 线索管理扩展 ====================

// GetClueDetail 获取线索详情
func (l *LocalClient) GetClueDetail(ctx context.Context, accessToken string, advertiserID uint64, clueID uint64) (*Clue, error) {
	path := "/v1.0/local/clue/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"clue_id":       clueID,
	}
	var result struct {
		Data Clue `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateClueFollowStatus 更新线索跟进状态
func (l *LocalClient) UpdateClueFollowStatus(ctx context.Context, accessToken string, advertiserID uint64, clueID uint64, followStatus int, remark string) error {
	path := "/v1.0/local/clue/follow/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"clue_id":       clueID,
		"follow_status": followStatus,
	}
	if remark != "" {
		data["remark"] = remark
	}
	return l.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ClueCallback 线索回传
func (l *LocalClient) ClueCallback(ctx context.Context, accessToken string, advertiserID uint64, clueData map[string]interface{}) error {
	path := "/v1.0/local/clue/life/callback/"
	clueData["advertiser_id"] = advertiserID
	return l.client.PostWithToken(ctx, accessToken, path, clueData, nil)
}

// ==================== 扩展报表 ====================

// GetClueReport 获取线索报表
func (l *LocalClient) GetClueReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]map[string]interface{}, int, error) {
	path := "/v1.0/local/report/clue/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []map[string]interface{} `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetStoreReport 获取门店报表
func (l *LocalClient) GetStoreReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, poiIDs []string) ([]map[string]interface{}, error) {
	path := "/v1.0/local/report/poi/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(poiIDs) > 0 {
		params["poi_ids"] = poiIDs
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := l.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}
