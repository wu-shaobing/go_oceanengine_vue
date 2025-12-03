package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// ServeMarketClient 服务市场API客户端
type ServeMarketClient struct {
	client *Client
}

// NewServeMarketClient 创建服务市场客户端
func (c *Client) ServeMarket() *ServeMarketClient {
	return &ServeMarketClient{client: c}
}

// ==================== 订单管理 ====================

// AppOrder 应用订单
type AppOrder struct {
	OrderID      string `json:"order_id"`
	AppID        string `json:"app_id"`
	AppName      string `json:"app_name"`
	OrderStatus  int    `json:"order_status"`
	OrderAmount  int64  `json:"order_amount"`
	CreateTime   string `json:"create_time"`
	ExpireTime   string `json:"expire_time"`
	AdvertiserID uint64 `json:"advertiser_id"`
}

// GetAppOrderList 获取应用订单列表
func (s *ServeMarketClient) GetAppOrderList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]AppOrder, int, error) {
	path := "/2/servemarket/app/order/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []AppOrder `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 功能点管理 ====================

// FuncPoint 功能点
type FuncPoint struct {
	FuncID      string `json:"func_id"`
	FuncName    string `json:"func_name"`
	FuncType    int    `json:"func_type"`
	Status      int    `json:"status"`
	TotalCount  int    `json:"total_count"`
	UsedCount   int    `json:"used_count"`
	RemainCount int    `json:"remain_count"`
	ExpireTime  string `json:"expire_time"`
}

// GetFuncPointList 获取已购功能点列表
func (s *ServeMarketClient) GetFuncPointList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]FuncPoint, int, error) {
	path := "/2/servemarket/func/purchased/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []FuncPoint `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ConsumeFuncPoint 消耗功能点
func (s *ServeMarketClient) ConsumeFuncPoint(ctx context.Context, accessToken string, advertiserID uint64, funcID string, count int) error {
	path := "/2/servemarket/func/consume/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"func_id":       funcID,
		"count":         count,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 质量报告 ====================

// QualityReport 质量报告
type QualityReport struct {
	ReportID   string             `json:"report_id"`
	ReportType int                `json:"report_type"`
	TargetID   string             `json:"target_id"`
	TargetType int                `json:"target_type"`
	Score      float64            `json:"score"`
	CreateTime string             `json:"create_time"`
	Dimensions []QualityDimension `json:"dimensions"`
}

type QualityDimension struct {
	DimensionName string  `json:"dimension_name"`
	Score         float64 `json:"score"`
	Suggestion    string  `json:"suggestion"`
}

// GetQualityReport 获取质量报告
func (s *ServeMarketClient) GetQualityReport(ctx context.Context, accessToken string, advertiserID uint64, targetID string, targetType int) (*QualityReport, error) {
	path := "/2/servemarket/quality/report/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"target_id":     targetID,
		"target_type":   targetType,
	}

	var result struct {
		Data QualityReport `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// CreateQualityReport 创建质量报告
func (s *ServeMarketClient) CreateQualityReport(ctx context.Context, accessToken string, advertiserID uint64, targetID string, targetType int) (string, error) {
	path := "/2/servemarket/quality/report/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"target_id":     targetID,
		"target_type":   targetType,
	}

	var result struct {
		Data struct {
			ReportID string `json:"report_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.ReportID, nil
}

// ==================== 订阅管理 ====================

// RdsSubscription RDS订阅
type RdsSubscription struct {
	SubscriptionID   string `json:"subscription_id"`
	SubscriptionType int    `json:"subscription_type"`
	Status           int    `json:"status"`
	CreateTime       string `json:"create_time"`
	ExpireTime       string `json:"expire_time"`
	CallbackURL      string `json:"callback_url"`
}

// GetRdsSubscriptionList 获取RDS订阅列表
func (s *ServeMarketClient) GetRdsSubscriptionList(ctx context.Context, accessToken string, advertiserID uint64) ([]RdsSubscription, error) {
	path := "/2/servemarket/rds/subscription/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data struct {
			List []RdsSubscription `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// CreateRdsSubscription 创建RDS订阅
func (s *ServeMarketClient) CreateRdsSubscription(ctx context.Context, accessToken string, advertiserID uint64, subscriptionType int, callbackURL string) (string, error) {
	path := "/2/servemarket/rds/subscription/create/"
	data := map[string]interface{}{
		"advertiser_id":     advertiserID,
		"subscription_type": subscriptionType,
		"callback_url":      callbackURL,
	}

	var result struct {
		Data struct {
			SubscriptionID string `json:"subscription_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.SubscriptionID, nil
}

// UpdateRdsSubscription 更新RDS订阅
func (s *ServeMarketClient) UpdateRdsSubscription(ctx context.Context, accessToken string, advertiserID uint64, subscriptionID string, callbackURL string) error {
	path := "/2/servemarket/rds/subscription/update/"
	data := map[string]interface{}{
		"advertiser_id":   advertiserID,
		"subscription_id": subscriptionID,
		"callback_url":    callbackURL,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteRdsSubscription 删除RDS订阅
func (s *ServeMarketClient) DeleteRdsSubscription(ctx context.Context, accessToken string, advertiserID uint64, subscriptionID string) error {
	path := "/2/servemarket/rds/subscription/delete/"
	data := map[string]interface{}{
		"advertiser_id":   advertiserID,
		"subscription_id": subscriptionID,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 连山投前分析 ====================

// QualitySubmitRequest 投前分析提交请求
type QualitySubmitRequest struct {
	AdvertiserID uint64   `json:"advertiser_id"`
	MaterialIDs  []string `json:"material_ids"`
	AnalyzeType  int      `json:"analyze_type"` // 1:视频 2:图片
}

// SubmitQualityAnalysis 提交投前分析
func (s *ServeMarketClient) SubmitQualityAnalysis(ctx context.Context, accessToken string, req *QualitySubmitRequest) (uint64, error) {
	path := "/v3/file/quality/submit/"
	data := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"material_ids":  req.MaterialIDs,
		"analyze_type":  req.AnalyzeType,
	}

	var result struct {
		Data struct {
			TaskID uint64 `json:"task_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.TaskID, nil
}

// MaterialQuality 素材质量分析结果
type MaterialQuality struct {
	MaterialID  string              `json:"material_id"`
	Score       float64             `json:"score"`
	Status      int                 `json:"status"` // 0:处理中 1:已完成 2:失败
	Suggestions []QualitySuggestion `json:"suggestions"`
}

type QualitySuggestion struct {
	Dimension  string  `json:"dimension"`
	Score      float64 `json:"score"`
	Suggestion string  `json:"suggestion"`
}

// GetQualityAnalysisResult 获取投前分析结果
func (s *ServeMarketClient) GetQualityAnalysisResult(ctx context.Context, accessToken string, advertiserID uint64, taskID uint64) ([]MaterialQuality, error) {
	path := "/v3/file/quality/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
	}

	var result struct {
		Data struct {
			List []MaterialQuality `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== RDS账户订阅 ====================

// RdsAccount RDS订阅账户
type RdsAccount struct {
	AdvertiserID uint64 `json:"advertiser_id"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// AddRdsAccounts 新增RDS账户订阅
func (s *ServeMarketClient) AddRdsAccounts(ctx context.Context, accessToken string, subscriptionID string, advertiserIDs []uint64) error {
	path := "/2/servemarket/rds/accounts/add/"
	data := map[string]interface{}{
		"subscription_id": subscriptionID,
		"advertiser_ids":  advertiserIDs,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// RemoveRdsAccounts 删除RDS账户订阅
func (s *ServeMarketClient) RemoveRdsAccounts(ctx context.Context, accessToken string, subscriptionID string, advertiserIDs []uint64) error {
	path := "/2/servemarket/rds/accounts/remove/"
	data := map[string]interface{}{
		"subscription_id": subscriptionID,
		"advertiser_ids":  advertiserIDs,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetRdsAccountsList 获取RDS订阅账户列表
func (s *ServeMarketClient) GetRdsAccountsList(ctx context.Context, accessToken string, subscriptionID string, page, pageSize int) ([]RdsAccount, int, error) {
	path := "/2/servemarket/rds/accounts/list/"
	params := map[string]interface{}{
		"subscription_id": subscriptionID,
		"page":            page,
		"page_size":       pageSize,
	}

	var result struct {
		Data struct {
			List     []RdsAccount `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== Token校验 ====================

// CidVerifyResult Token校验结果
type CidVerifyResult struct {
	AppID        string `json:"app_id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	ExpireTime   int64  `json:"expire_time"`
	Scope        string `json:"scope"`
}

// VerifyAppAccessToken 校验App Access Token
func (s *ServeMarketClient) VerifyAppAccessToken(ctx context.Context, token string) (*CidVerifyResult, error) {
	path := "/2/servemarket/cid/verify_token/"
	params := map[string]interface{}{
		"token": token,
	}

	resp, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data CidVerifyResult `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 应用信息 ====================

// AppInfo 应用信息
type AppInfo struct {
	AppID       string `json:"app_id"`
	AppName     string `json:"app_name"`
	AppType     int    `json:"app_type"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
}

// GetAppInfo 获取应用信息
func (s *ServeMarketClient) GetAppInfo(ctx context.Context, accessToken string, appID string) (*AppInfo, error) {
	path := "/2/servemarket/app/info/"
	params := map[string]interface{}{
		"app_id": appID,
	}

	var result struct {
		Data AppInfo `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 功能点消耗记录 ====================

// FuncConsumeRecord 功能点消耗记录
type FuncConsumeRecord struct {
	RecordID     string `json:"record_id"`
	FuncID       string `json:"func_id"`
	FuncName     string `json:"func_name"`
	ConsumeCount int    `json:"consume_count"`
	ConsumeTime  string `json:"consume_time"`
	Remark       string `json:"remark"`
}

// GetFuncConsumeRecords 获取功能点消耗记录
func (s *ServeMarketClient) GetFuncConsumeRecords(ctx context.Context, accessToken string, advertiserID uint64, funcID string, startDate, endDate string, page, pageSize int) ([]FuncConsumeRecord, int, error) {
	path := "/2/servemarket/func/consume/records/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"func_id":       funcID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []FuncConsumeRecord `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}
