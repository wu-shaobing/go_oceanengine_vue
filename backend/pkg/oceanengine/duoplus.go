package oceanengine

import (
	"context"
)

// DuoplusClient Dou+投放客户端
type DuoplusClient struct {
	client *Client
}

// NewDuoplusClient 创建Dou+客户端
func NewDuoplusClient(client *Client) *DuoplusClient {
	return &DuoplusClient{client: client}
}

// ==================== Dou+订单管理 ====================

// DuoplusOrder Dou+订单
type DuoplusOrder struct {
	OrderID      uint64   `json:"order_id"`
	AwemeItemID  uint64   `json:"aweme_item_id"`
	AwemeID      uint64   `json:"aweme_id"`
	OrderStatus  int      `json:"order_status"`
	Budget       float64  `json:"budget"`
	Cost         float64  `json:"cost"`
	DeliveryGoal string   `json:"delivery_goal"`
	DeliveryTime int      `json:"delivery_time"`
	Audience     string   `json:"audience"`
	AgeRange     string   `json:"age_range"`
	Gender       string   `json:"gender"`
	City         []string `json:"city"`
	InterestTags []string `json:"interest_tags"`
	CreateTime   string   `json:"create_time"`
	StartTime    string   `json:"start_time"`
	EndTime      string   `json:"end_time"`
}

// DuoplusOrderCreateRequest 创建Dou+订单请求
type DuoplusOrderCreateRequest struct {
	AwemeID      uint64   `json:"aweme_id"`
	AwemeItemID  uint64   `json:"aweme_item_id"`
	Budget       float64  `json:"budget"`
	DeliveryGoal string   `json:"delivery_goal"`
	DeliveryTime int      `json:"delivery_time"`
	Audience     string   `json:"audience,omitempty"`
	AgeRange     string   `json:"age_range,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	City         []string `json:"city,omitempty"`
	InterestTags []string `json:"interest_tags,omitempty"`
}

// CreateOrder 创建Dou+订单
func (d *DuoplusClient) CreateOrder(ctx context.Context, accessToken string, req *DuoplusOrderCreateRequest) (uint64, error) {
	path := "/v1.0/duoplus/order/create/"
	var result struct {
		Data struct {
			OrderID uint64 `json:"order_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.OrderID, nil
}

// GetOrderList 获取Dou+订单列表
func (d *DuoplusClient) GetOrderList(ctx context.Context, accessToken string, awemeID uint64, page, pageSize int) ([]DuoplusOrder, int, error) {
	path := "/v1.0/duoplus/order/get/"
	params := map[string]interface{}{
		"aweme_id":  awemeID,
		"page":      page,
		"page_size": pageSize,
	}
	var result struct {
		Data struct {
			List     []DuoplusOrder `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetOrderDetail 获取Dou+订单详情
func (d *DuoplusClient) GetOrderDetail(ctx context.Context, accessToken string, orderID uint64) (*DuoplusOrder, error) {
	path := "/v1.0/duoplus/order/detail/get/"
	params := map[string]interface{}{
		"order_id": orderID,
	}
	var result struct {
		Data DuoplusOrder `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// PauseOrder 暂停Dou+订单
func (d *DuoplusClient) PauseOrder(ctx context.Context, accessToken string, orderID uint64) error {
	path := "/v1.0/duoplus/order/pause/"
	data := map[string]interface{}{
		"order_id": orderID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ResumeOrder 恢复Dou+订单
func (d *DuoplusClient) ResumeOrder(ctx context.Context, accessToken string, orderID uint64) error {
	path := "/v1.0/duoplus/order/resume/"
	data := map[string]interface{}{
		"order_id": orderID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// TerminateOrder 终止Dou+订单
func (d *DuoplusClient) TerminateOrder(ctx context.Context, accessToken string, orderID uint64) error {
	path := "/v1.0/duoplus/order/terminate/"
	data := map[string]interface{}{
		"order_id": orderID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// AddOrderBudget 追加Dou+订单预算
func (d *DuoplusClient) AddOrderBudget(ctx context.Context, accessToken string, orderID uint64, addBudget float64) error {
	path := "/v1.0/duoplus/order/budget/add/"
	data := map[string]interface{}{
		"order_id":   orderID,
		"add_budget": addBudget,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== Dou+数据报表 ====================

// DuoplusOrderStats Dou+订单统计数据
type DuoplusOrderStats struct {
	OrderID      uint64  `json:"order_id"`
	Cost         float64 `json:"cost"`
	ShowCount    int64   `json:"show_count"`
	ClickCount   int64   `json:"click_count"`
	PlayCount    int64   `json:"play_count"`
	LikeCount    int64   `json:"like_count"`
	CommentCount int64   `json:"comment_count"`
	ShareCount   int64   `json:"share_count"`
	FollowCount  int64   `json:"follow_count"`
	ProfileVisit int64   `json:"profile_visit"`
	CTR          float64 `json:"ctr"`
	CPM          float64 `json:"cpm"`
}

// GetOrderStats 获取Dou+订单统计数据
func (d *DuoplusClient) GetOrderStats(ctx context.Context, accessToken string, orderID uint64) (*DuoplusOrderStats, error) {
	path := "/v1.0/duoplus/order/stats/get/"
	params := map[string]interface{}{
		"order_id": orderID,
	}
	var result struct {
		Data DuoplusOrderStats `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DuoplusItemStats Dou+视频统计数据
type DuoplusItemStats struct {
	AwemeItemID    uint64  `json:"aweme_item_id"`
	TotalCost      float64 `json:"total_cost"`
	TotalShow      int64   `json:"total_show"`
	TotalClick     int64   `json:"total_click"`
	TotalPlay      int64   `json:"total_play"`
	TotalLike      int64   `json:"total_like"`
	TotalComment   int64   `json:"total_comment"`
	TotalShare     int64   `json:"total_share"`
	TotalFollow    int64   `json:"total_follow"`
	OrganicShow    int64   `json:"organic_show"`
	OrganicLike    int64   `json:"organic_like"`
	OrganicComment int64   `json:"organic_comment"`
	OrganicShare   int64   `json:"organic_share"`
}

// GetItemStats 获取Dou+视频统计数据
func (d *DuoplusClient) GetItemStats(ctx context.Context, accessToken string, awemeItemID uint64) (*DuoplusItemStats, error) {
	path := "/v1.0/duoplus/item/stats/get/"
	params := map[string]interface{}{
		"aweme_item_id": awemeItemID,
	}
	var result struct {
		Data DuoplusItemStats `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== Dou+可投放视频 ====================

// DuoplusVideo Dou+可投放视频
type DuoplusVideo struct {
	AwemeItemID  uint64 `json:"aweme_item_id"`
	Title        string `json:"title"`
	CoverURL     string `json:"cover_url"`
	Duration     int    `json:"duration"`
	PlayCount    int64  `json:"play_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	ShareCount   int64  `json:"share_count"`
	CreateTime   string `json:"create_time"`
	Status       int    `json:"status"`
}

// GetAvailableVideos 获取可投放视频列表
func (d *DuoplusClient) GetAvailableVideos(ctx context.Context, accessToken string, awemeID uint64, page, pageSize int) ([]DuoplusVideo, int, error) {
	path := "/v1.0/duoplus/video/available/get/"
	params := map[string]interface{}{
		"aweme_id":  awemeID,
		"page":      page,
		"page_size": pageSize,
	}
	var result struct {
		Data struct {
			List     []DuoplusVideo `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== Dou+预估与建议 ====================

// DuoplusEstimate Dou+投放效果预估
type DuoplusEstimate struct {
	EstimatedShow    int64 `json:"estimated_show"`
	EstimatedLike    int64 `json:"estimated_like"`
	EstimatedComment int64 `json:"estimated_comment"`
	EstimatedShare   int64 `json:"estimated_share"`
	EstimatedFollow  int64 `json:"estimated_follow"`
}

// GetEstimate 获取Dou+投放效果预估
func (d *DuoplusClient) GetEstimate(ctx context.Context, accessToken string, awemeItemID uint64, budget float64, deliveryGoal string, deliveryTime int) (*DuoplusEstimate, error) {
	path := "/v1.0/duoplus/estimate/get/"
	params := map[string]interface{}{
		"aweme_item_id": awemeItemID,
		"budget":        budget,
		"delivery_goal": deliveryGoal,
		"delivery_time": deliveryTime,
	}
	var result struct {
		Data DuoplusEstimate `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DuoplusSuggest Dou+投放建议
type DuoplusSuggest struct {
	SuggestedBudget    float64 `json:"suggested_budget"`
	SuggestedGoal      string  `json:"suggested_goal"`
	SuggestedAudience  string  `json:"suggested_audience"`
	SuggestedTimeRange string  `json:"suggested_time_range"`
}

// GetSuggest 获取Dou+投放建议
func (d *DuoplusClient) GetSuggest(ctx context.Context, accessToken string, awemeItemID uint64) (*DuoplusSuggest, error) {
	path := "/v1.0/duoplus/suggest/get/"
	params := map[string]interface{}{
		"aweme_item_id": awemeItemID,
	}
	var result struct {
		Data DuoplusSuggest `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== Dou+定向工具 ====================

// DuoplusInterestTag Dou+兴趣标签
type DuoplusInterestTag struct {
	TagID   uint64 `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// GetInterestTags 获取Dou+兴趣标签列表
func (d *DuoplusClient) GetInterestTags(ctx context.Context, accessToken string) ([]DuoplusInterestTag, error) {
	path := "/v1.0/duoplus/tools/interest_tag/get/"
	var result struct {
		Data struct {
			List []DuoplusInterestTag `json:"list"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// DuoplusCity Dou+城市
type DuoplusCity struct {
	CityID   uint64 `json:"city_id"`
	CityName string `json:"city_name"`
	Level    int    `json:"level"`
	ParentID uint64 `json:"parent_id"`
}

// GetCities 获取Dou+城市列表
func (d *DuoplusClient) GetCities(ctx context.Context, accessToken string) ([]DuoplusCity, error) {
	path := "/v1.0/duoplus/tools/city/get/"
	var result struct {
		Data struct {
			List []DuoplusCity `json:"list"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== Dou+账户信息 ====================

// DuoplusAccount Dou+账户信息
type DuoplusAccount struct {
	AwemeID       uint64  `json:"aweme_id"`
	AwemeName     string  `json:"aweme_name"`
	Avatar        string  `json:"avatar"`
	FansCount     int64   `json:"fans_count"`
	Balance       float64 `json:"balance"`
	TotalRecharge float64 `json:"total_recharge"`
	TotalCost     float64 `json:"total_cost"`
}

// GetAccountInfo 获取Dou+账户信息
func (d *DuoplusClient) GetAccountInfo(ctx context.Context, accessToken string, awemeID uint64) (*DuoplusAccount, error) {
	path := "/v1.0/duoplus/account/info/get/"
	params := map[string]interface{}{
		"aweme_id": awemeID,
	}
	var result struct {
		Data DuoplusAccount `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DuoplusAccountStats Dou+账户统计
type DuoplusAccountStats struct {
	AwemeID        uint64  `json:"aweme_id"`
	TotalOrders    int     `json:"total_orders"`
	RunningOrders  int     `json:"running_orders"`
	FinishedOrders int     `json:"finished_orders"`
	TotalCost      float64 `json:"total_cost"`
	TotalShow      int64   `json:"total_show"`
	TotalFollow    int64   `json:"total_follow"`
}

// GetAccountStats 获取Dou+账户统计
func (d *DuoplusClient) GetAccountStats(ctx context.Context, accessToken string, awemeID uint64, startDate, endDate string) (*DuoplusAccountStats, error) {
	path := "/v1.0/duoplus/account/stats/get/"
	params := map[string]interface{}{
		"aweme_id":   awemeID,
		"start_date": startDate,
		"end_date":   endDate,
	}
	var result struct {
		Data DuoplusAccountStats `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== Dou+钱包 ====================

// DuoplusTransaction Dou+交易记录
type DuoplusTransaction struct {
	TransactionID   string  `json:"transaction_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Balance         float64 `json:"balance"`
	Remark          string  `json:"remark"`
	CreateTime      string  `json:"create_time"`
}

// GetTransactionList 获取Dou+交易记录
func (d *DuoplusClient) GetTransactionList(ctx context.Context, accessToken string, awemeID uint64, startDate, endDate string, page, pageSize int) ([]DuoplusTransaction, int, error) {
	path := "/v1.0/duoplus/wallet/transaction/get/"
	params := map[string]interface{}{
		"aweme_id":   awemeID,
		"start_date": startDate,
		"end_date":   endDate,
		"page":       page,
		"page_size":  pageSize,
	}
	var result struct {
		Data struct {
			List     []DuoplusTransaction `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetWalletBalance 获取Dou+钱包余额
func (d *DuoplusClient) GetWalletBalance(ctx context.Context, accessToken string, awemeID uint64) (float64, error) {
	path := "/v1.0/duoplus/wallet/balance/get/"
	params := map[string]interface{}{
		"aweme_id": awemeID,
	}
	var result struct {
		Data struct {
			Balance float64 `json:"balance"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.Balance, nil
}
