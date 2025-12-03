package oceanengine

import (
	"context"
)

// StarClient 星图API客户端
type StarClient struct {
	client *Client
}

// NewStarClient 创建星图客户端
func (c *Client) Star() *StarClient {
	return &StarClient{client: c}
}

// ==================== 账户信息 ====================

// StarAccountInfo 星图账户信息
type StarAccountInfo struct {
	AdvertiserID   uint64 `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	Balance        int64  `json:"balance"`
	AccountType    int    `json:"account_type"`
}

// GetAccountInfo 获取星图账户信息
func (s *StarClient) GetAccountInfo(ctx context.Context, accessToken string, advertiserID uint64) (*StarAccountInfo, error) {
	path := "/star/advertiser/info/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data StarAccountInfo `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 资金管理 ====================

// FundBalance 资金余额
type FundBalance struct {
	AdvertiserID uint64 `json:"advertiser_id"`
	Balance      int64  `json:"balance"`
	FreezeAmount int64  `json:"freeze_amount"`
	ValidAmount  int64  `json:"valid_amount"`
}

// GetFundBalance 获取资金余额
func (s *StarClient) GetFundBalance(ctx context.Context, accessToken string, advertiserIDs []uint64) ([]FundBalance, error) {
	path := "/star/fund/balance/get/"
	params := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	var result struct {
		Data struct {
			List []FundBalance `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// FundDaily 资金日流水
type FundDaily struct {
	Date          string `json:"date"`
	Income        int64  `json:"income"`
	Expense       int64  `json:"expense"`
	Balance       int64  `json:"balance"`
	FreezeIncome  int64  `json:"freeze_income"`
	FreezeExpense int64  `json:"freeze_expense"`
}

// GetFundDaily 获取资金日流水
func (s *StarClient) GetFundDaily(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string) ([]FundDaily, error) {
	path := "/star/fund/daily/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}

	var result struct {
		Data struct {
			List []FundDaily `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// StarFundTransaction 星图资金流水明细
type StarFundTransaction struct {
	TransactionID   string `json:"transaction_id"`
	TransactionType int    `json:"transaction_type"`
	Amount          int64  `json:"amount"`
	Balance         int64  `json:"balance"`
	CreateTime      string `json:"create_time"`
	Remark          string `json:"remark"`
}

// GetFundTransaction 获取资金流水明细
func (s *StarClient) GetFundTransaction(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]StarFundTransaction, int, error) {
	path := "/star/fund/transaction/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []StarFundTransaction `json:"list"`
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

// ==================== 任务管理 ====================

// StarTask 星图任务
type StarTask struct {
	TaskID      uint64 `json:"task_id"`
	TaskName    string `json:"task_name"`
	TaskType    int    `json:"task_type"`
	TaskStatus  int    `json:"task_status"`
	Budget      int64  `json:"budget"`
	CreateTime  string `json:"create_time"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	TalentCount int    `json:"talent_count"`
	FinishCount int    `json:"finish_count"`
}

// GetTaskList 获取任务列表
func (s *StarClient) GetTaskList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]StarTask, int, error) {
	path := "/star/task/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []StarTask `json:"list"`
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

// GetTaskDetail 获取任务详情
func (s *StarClient) GetTaskDetail(ctx context.Context, accessToken string, advertiserID uint64, taskID uint64) (*StarTask, error) {
	path := "/star/task/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
	}

	var result struct {
		Data StarTask `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 需求管理 ====================

// Demand 需求信息
type Demand struct {
	DemandID    uint64 `json:"demand_id"`
	DemandName  string `json:"demand_name"`
	DemandType  int    `json:"demand_type"`
	Status      int    `json:"status"`
	Budget      int64  `json:"budget"`
	CreateTime  string `json:"create_time"`
	TalentCount int    `json:"talent_count"`
}

// GetDemandList 获取需求列表
func (s *StarClient) GetDemandList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]Demand, int, error) {
	path := "/star/demand/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []Demand `json:"list"`
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

// DemandOrder 需求订单
type DemandOrder struct {
	OrderID     uint64 `json:"order_id"`
	DemandID    uint64 `json:"demand_id"`
	TalentID    uint64 `json:"talent_id"`
	TalentName  string `json:"talent_name"`
	OrderStatus int    `json:"order_status"`
	OrderAmount int64  `json:"order_amount"`
	CreateTime  string `json:"create_time"`
}

// GetDemandOrderList 获取需求订单列表
func (s *StarClient) GetDemandOrderList(ctx context.Context, accessToken string, advertiserID, demandID uint64, page, pageSize int) ([]DemandOrder, int, error) {
	path := "/star/demand/order/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"demand_id":     demandID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []DemandOrder `json:"list"`
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

// ==================== 投后报表 ====================

// StarReport 星图报表
type StarReport struct {
	Date         string  `json:"date"`
	PlayCount    int64   `json:"play_count"`
	LikeCount    int64   `json:"like_count"`
	CommentCount int64   `json:"comment_count"`
	ShareCount   int64   `json:"share_count"`
	Cost         float64 `json:"cost"`
	Cpm          float64 `json:"cpm"`
	Cpe          float64 `json:"cpe"`
}

// GetReportOverview 获取投后报表概览
func (s *StarClient) GetReportOverview(ctx context.Context, accessToken string, advertiserID uint64, taskID uint64) (*StarReport, error) {
	path := "/star/report/overview/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
	}

	var result struct {
		Data StarReport `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetReportDaily 获取每日报表
func (s *StarClient) GetReportDaily(ctx context.Context, accessToken string, advertiserID, taskID uint64, startDate, endDate string) ([]StarReport, error) {
	path := "/star/report/daily/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"start_date":    startDate,
		"end_date":      endDate,
	}

	var result struct {
		Data struct {
			List []StarReport `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// StarReportAudience 受众报表
type StarReportAudience struct {
	Gender   map[string]int64   `json:"gender"`
	Age      []map[string]int64 `json:"age"`
	Province []map[string]int64 `json:"province"`
}

// GetReportAudience 获取受众报表
func (s *StarClient) GetReportAudience(ctx context.Context, accessToken string, advertiserID uint64, taskID uint64) (*StarReportAudience, error) {
	path := "/star/report/audience/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
	}

	var result struct {
		Data StarReportAudience `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetDemandOrders 获取需求订单列表 (alias)
func (s *StarClient) GetDemandOrders(ctx context.Context, accessToken string, advertiserID, demandID uint64, page, pageSize int) ([]DemandOrder, int, error) {
	return s.GetDemandOrderList(ctx, accessToken, advertiserID, demandID, page, pageSize)
}

// ==================== 线索管理 ====================

// StarClue 星图线索
type StarClue struct {
	ClueID     uint64 `json:"clue_id"`
	TaskID     uint64 `json:"task_id"`
	TalentID   uint64 `json:"talent_id"`
	ClueType   int    `json:"clue_type"`
	CreateTime string `json:"create_time"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
}

// GetClueList 获取线索列表
func (s *StarClient) GetClueList(ctx context.Context, accessToken string, advertiserID uint64, taskID uint64, page, pageSize int) ([]StarClue, int, error) {
	path := "/star/clue/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []StarClue `json:"list"`
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

// ==================== 代理商管理 ====================

// StarAgent 星图代理商信息
type StarAgent struct {
	AgentID   uint64 `json:"agent_id"`
	AgentName string `json:"agent_name"`
	Status    int    `json:"status"`
}

// GetAgentInfo 获取代理商信息
func (s *StarClient) GetAgentInfo(ctx context.Context, accessToken string, agentID uint64) (*StarAgent, error) {
	path := "/star/agent/info/"
	params := map[string]interface{}{
		"agent_id": agentID,
	}
	var result struct {
		Data StarAgent `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetAdvertiserList 获取广告主列表
func (s *StarClient) GetAdvertiserList(ctx context.Context, accessToken string, agentID uint64, page, pageSize int) ([]StarAccountInfo, int, error) {
	path := "/star/agent/advertiser/select/"
	params := map[string]interface{}{
		"agent_id":  agentID,
		"page":      page,
		"page_size": pageSize,
	}
	var result struct {
		Data struct {
			List     []StarAccountInfo `json:"list"`
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

// GetChildAgentList 获取二级代理商列表
func (s *StarClient) GetChildAgentList(ctx context.Context, accessToken string, agentID uint64) ([]uint64, error) {
	path := "/star/agent/child_agent/select/"
	params := map[string]interface{}{
		"agent_id": agentID,
	}
	var result struct {
		Data struct {
			AgentIDs []uint64 `json:"agent_ids"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.AgentIDs, nil
}

// UpdateAdvertiser 修改广告主信息
func (s *StarClient) UpdateAdvertiser(ctx context.Context, accessToken string, advertiserID uint64, updateData map[string]interface{}) error {
	path := "/star/agent/advertiser/update/"
	updateData["advertiser_id"] = advertiserID
	return s.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// ==================== 资金扩展 ====================

// GetTransferRecord 查询代理商转账记录
func (s *StarClient) GetTransferRecord(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]map[string]interface{}, int, error) {
	path := "/star/agent/transfer/transaction/record/"
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
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 任务管理扩展 ====================

// StarTaskItem 星图任务视频维度数据
type StarTaskItem struct {
	ItemID       uint64  `json:"item_id"`
	TaskID       uint64  `json:"task_id"`
	TalentID     uint64  `json:"talent_id"`
	TalentName   string  `json:"talent_name"`
	VideoTitle   string  `json:"video_title"`
	PlayCount    int64   `json:"play_count"`
	LikeCount    int64   `json:"like_count"`
	CommentCount int64   `json:"comment_count"`
	ShareCount   int64   `json:"share_count"`
	PublishTime  string  `json:"publish_time"`
	Cost         float64 `json:"cost"`
}

// GetTaskItemList 获取任务视频维度数据
func (s *StarClient) GetTaskItemList(ctx context.Context, accessToken string, advertiserID, taskID uint64, page, pageSize int) ([]StarTaskItem, int, error) {
	path := "/star/task/item/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []StarTaskItem `json:"list"`
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

// CreateTask 创建任务
func (s *StarClient) CreateTask(ctx context.Context, accessToken string, advertiserID uint64, taskData map[string]interface{}) (uint64, error) {
	path := "/star/task/create/"
	taskData["advertiser_id"] = advertiserID
	var result struct {
		Data struct {
			TaskID uint64 `json:"task_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, taskData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.TaskID, nil
}

// UpdateTask 更新任务
func (s *StarClient) UpdateTask(ctx context.Context, accessToken string, advertiserID, taskID uint64, updateData map[string]interface{}) error {
	path := "/star/task/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["task_id"] = taskID
	return s.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// UpdateTaskStatus 更新任务状态
func (s *StarClient) UpdateTaskStatus(ctx context.Context, accessToken string, advertiserID, taskID uint64, status string) error {
	path := "/star/task/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"status":        status, // 取消/完成等
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// UpdateClueStatus 更新线索状态
func (s *StarClient) UpdateClueStatus(ctx context.Context, accessToken string, clueID uint64, status int, remark string) error {
	path := "/star/clue/status/update/"
	data := map[string]interface{}{
		"clue_id": clueID,
		"status":  status, // 跟进状态
	}
	if remark != "" {
		data["remark"] = remark
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 投后报表扩展 ====================

// StarAudienceReport 受众报表
type StarAudienceReport struct {
	GenderDistribution   []map[string]interface{} `json:"gender_distribution"`
	AgeDistribution      []map[string]interface{} `json:"age_distribution"`
	CityDistribution     []map[string]interface{} `json:"city_distribution"`
	InterestDistribution []map[string]interface{} `json:"interest_distribution"`
}

// GetAudienceReport 获取受众报表
func (s *StarClient) GetAudienceReport(ctx context.Context, accessToken string, advertiserID, taskID uint64, orderID uint64) (*StarAudienceReport, error) {
	path := "/star/report/order/user_distribution/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"order_id":      orderID,
	}
	var result struct {
		Data StarAudienceReport `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetDataTopicConfig 获取任务下可查询的数据指标
func (s *StarClient) GetDataTopicConfig(ctx context.Context, accessToken string, advertiserID, taskID uint64) ([]map[string]interface{}, error) {
	path := "/star/report/data_topic/config/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetCustomDataTopicConfig 获取投后数据主题累计数据
func (s *StarClient) GetCustomDataTopicConfig(ctx context.Context, accessToken string, advertiserID, taskID uint64, topics []string) (map[string]interface{}, error) {
	path := "/star/report/custom_data_topic/config/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"topics":        topics,
	}
	var result struct {
		Data map[string]interface{} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetCustomDataTopicDailyReport 获取投后每日趋势数据
func (s *StarClient) GetCustomDataTopicDailyReport(ctx context.Context, accessToken string, advertiserID, taskID uint64, startDate, endDate string, topics []string) ([]map[string]interface{}, error) {
	path := "/star/report/custom_data_topic/daily/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"task_id":       taskID,
		"start_date":    startDate,
		"end_date":      endDate,
		"topics":        topics,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 需求管理扩展 ====================

// CreateDemand 创建需求
func (s *StarClient) CreateDemand(ctx context.Context, accessToken string, advertiserID uint64, demandData map[string]interface{}) (uint64, error) {
	path := "/star/demand/create/"
	demandData["advertiser_id"] = advertiserID
	var result struct {
		Data struct {
			DemandID uint64 `json:"demand_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, demandData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.DemandID, nil
}

// GetDemandDetail 获取需求详情
func (s *StarClient) GetDemandDetail(ctx context.Context, accessToken string, advertiserID, demandID uint64) (*Demand, error) {
	path := "/star/demand/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"demand_id":     demandID,
	}
	var result struct {
		Data Demand `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// AcceptDemandOrder 接受需求订单
func (s *StarClient) AcceptDemandOrder(ctx context.Context, accessToken string, advertiserID, orderID uint64) error {
	path := "/star/demand/order/accept/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_id":      orderID,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// RejectDemandOrder 拒绝需求订单
func (s *StarClient) RejectDemandOrder(ctx context.Context, accessToken string, advertiserID, orderID uint64, reason string) error {
	path := "/star/demand/order/reject/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_id":      orderID,
		"reason":        reason,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 达人管理 ====================

// StarTalent 星图达人
type StarTalent struct {
	TalentID      uint64  `json:"talent_id"`
	TalentName    string  `json:"talent_name"`
	Avatar        string  `json:"avatar"`
	FollowerCount int64   `json:"follower_count"`
	Category      string  `json:"category"`
	Price         int64   `json:"price"`
	Score         float64 `json:"score"`
}

// SearchTalent 搜索达人
func (s *StarClient) SearchTalent(ctx context.Context, accessToken string, advertiserID uint64, keyword string, page, pageSize int) ([]StarTalent, int, error) {
	path := "/star/talent/search/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"keyword":       keyword,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []StarTalent `json:"list"`
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

// GetTalentDetail 获取达人详情
func (s *StarClient) GetTalentDetail(ctx context.Context, accessToken string, advertiserID, talentID uint64) (*StarTalent, error) {
	path := "/star/talent/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"talent_id":     talentID,
	}
	var result struct {
		Data StarTalent `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetRecommendTalentList 获取推荐达人列表
func (s *StarClient) GetRecommendTalentList(ctx context.Context, accessToken string, advertiserID uint64, category string, page, pageSize int) ([]StarTalent, int, error) {
	path := "/star/talent/recommend/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if category != "" {
		params["category"] = category
	}
	var result struct {
		Data struct {
			List     []StarTalent `json:"list"`
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
