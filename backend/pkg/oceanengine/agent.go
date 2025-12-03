package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// AgentService 代理商服务
type AgentService struct {
	client *Client
}

// NewAgentService 创建代理商服务
func NewAgentService(client *Client) *AgentService {
	return &AgentService{client: client}
}

// Agent 创建代理商服务(链式调用)
func (c *Client) Agent() *AgentService {
	return NewAgentService(c)
}

// AgentInfo 代理商信息
type AgentInfo struct {
	AgentID       int64  `json:"agent_id"`
	AgentName     string `json:"agent_name"`
	CompanyName   string `json:"company_name"`
	Role          string `json:"role"`
	Status        string `json:"status"`
	CreateTime    string `json:"create_time"`
	CustomerCount int    `json:"customer_count"`
}

// GetAgentInfo 获取代理商信息
func (s *AgentService) GetAgentInfo(ctx context.Context, accessToken string, agentID int64) (*AgentInfo, error) {
	params := map[string]interface{}{
		"agent_id": agentID,
	}

	var result struct {
		Data AgentInfo `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/agent/info/", params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// AgentAdvertiser 代理商下的广告主
type AgentAdvertiser struct {
	AdvertiserID   int64   `json:"advertiser_id"`
	AdvertiserName string  `json:"advertiser_name"`
	Company        string  `json:"company"`
	Role           string  `json:"role"`
	Status         string  `json:"status"`
	Balance        float64 `json:"balance"`
	CreateTime     string  `json:"create_time"`
}

// AgentAdvertiserListRequest 代理商广告主列表请求
type AgentAdvertiserListRequest struct {
	AgentID      int64  `json:"agent_id"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
	Status       string `json:"status,omitempty"`
	AdvertiserID int64  `json:"advertiser_id,omitempty"`
}

// GetAdvertiserList 获取代理商下的广告主列表
func (s *AgentService) GetAdvertiserList(ctx context.Context, accessToken string, req *AgentAdvertiserListRequest) ([]AgentAdvertiser, int, error) {
	params := map[string]interface{}{
		"agent_id": req.AgentID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.Status != "" {
		params["status"] = req.Status
	}
	if req.AdvertiserID > 0 {
		params["advertiser_id"] = req.AdvertiserID
	}

	var result struct {
		Data struct {
			List     []AgentAdvertiser `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
				TotalPage   int `json:"total_page"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/agent/advertiser/select/", params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// TransferRequest 转账请求
type TransferRequest struct {
	AgentID              int64   `json:"agent_id"`
	AdvertiserID         int64   `json:"advertiser_id"`
	TransferType         string  `json:"transfer_type"` // agent_to_advertiser, advertiser_to_agent
	Amount               float64 `json:"amount"`
	TransferSerialNumber string  `json:"transfer_serial_no,omitempty"`
}

// TransferResponse 转账响应
type TransferResponse struct {
	TransactionSeq string  `json:"transaction_seq"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
}

// Transfer 代理商与广告主之间转账
func (s *AgentService) Transfer(ctx context.Context, accessToken string, req *TransferRequest) (*TransferResponse, error) {
	body := map[string]interface{}{
		"agent_id":      req.AgentID,
		"advertiser_id": req.AdvertiserID,
		"transfer_type": req.TransferType,
		"amount":        req.Amount,
	}
	if req.TransferSerialNumber != "" {
		body["transfer_serial_no"] = req.TransferSerialNumber
	}

	var result struct {
		Data TransferResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, "/2/agent/advertiser/recharge/", body, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// AgentFundInfo 代理商资金信息
type AgentFundInfo struct {
	AgentID       int64   `json:"agent_id"`
	Balance       float64 `json:"balance"`
	ValidBalance  float64 `json:"valid_balance"`
	FreezeBalance float64 `json:"freeze_balance"`
	CashBalance   float64 `json:"cash_balance"`
	GrantBalance  float64 `json:"grant_balance"`
}

// GetAgentFund 获取代理商资金信息
func (s *AgentService) GetAgentFund(ctx context.Context, accessToken string, agentID int64) (*AgentFundInfo, error) {
	params := map[string]interface{}{
		"agent_id": agentID,
	}

	var result struct {
		Data AgentFundInfo `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/agent/fund/get/", params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// AgentFundFlow 代理商资金流水
type AgentFundFlow struct {
	TransactionSeq  string  `json:"transaction_seq"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	TransactionTime string  `json:"transaction_time"`
	BalanceAfter    float64 `json:"balance_after"`
	AdvertiserID    int64   `json:"advertiser_id,omitempty"`
	AdvertiserName  string  `json:"advertiser_name,omitempty"`
	Remark          string  `json:"remark,omitempty"`
}

// AgentFundFlowRequest 代理商资金流水请求
type AgentFundFlowRequest struct {
	AgentID         int64  `json:"agent_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	TransactionType string `json:"transaction_type,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        int    `json:"page_size,omitempty"`
}

// GetAgentFundFlow 获取代理商资金流水
func (s *AgentService) GetAgentFundFlow(ctx context.Context, accessToken string, req *AgentFundFlowRequest) ([]AgentFundFlow, int, error) {
	body := map[string]interface{}{
		"agent_id":   req.AgentID,
		"start_date": req.StartDate,
		"end_date":   req.EndDate,
	}
	if req.TransactionType != "" {
		body["transaction_type"] = req.TransactionType
	}
	if req.Page > 0 {
		body["page"] = req.Page
	}
	if req.PageSize > 0 {
		body["page_size"] = req.PageSize
	}

	resp, err := s.client.Post(ctx, "/2/agent/fund/transaction/get/", body)
	if err != nil {
		return nil, 0, err
	}

	if !resp.IsSuccess() {
		return nil, 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List     []AgentFundFlow `json:"list"`
		PageInfo struct {
			TotalNumber int `json:"total_number"`
		} `json:"page_info"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("unmarshal response failed: %w", err)
	}

	return result.List, result.PageInfo.TotalNumber, nil
}

// CreateAdvertiserRequest 创建广告主请求
type CreateAdvertiserRequest struct {
	AgentID          int64   `json:"agent_id"`
	AdvertiserName   string  `json:"advertiser_name"`
	Company          string  `json:"company"`
	BrandName        string  `json:"brand_name,omitempty"`
	Industry         []int64 `json:"industry,omitempty"`
	ContactName      string  `json:"contact_name"`
	ContactPhone     string  `json:"contact_phone"`
	LicenseURL       string  `json:"license_url,omitempty"`
	LicenseNo        string  `json:"license_no,omitempty"`
	PromotionArea    string  `json:"promotion_area,omitempty"`
	FirstIndustryID  int64   `json:"first_industry_id,omitempty"`
	SecondIndustryID int64   `json:"second_industry_id,omitempty"`
}

// CreateAdvertiser 创建广告主
func (s *AgentService) CreateAdvertiser(ctx context.Context, accessToken string, req *CreateAdvertiserRequest) (int64, error) {
	var result struct {
		Data struct {
			AdvertiserID int64 `json:"advertiser_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, "/2/agent/advertiser/create/", req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.AdvertiserID, nil
}

// UpdateAdvertiserRequest 更新广告主请求
type UpdateAdvertiserRequest struct {
	AgentID        int64  `json:"agent_id"`
	AdvertiserID   int64  `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name,omitempty"`
	Company        string `json:"company,omitempty"`
	ContactName    string `json:"contact_name,omitempty"`
	ContactPhone   string `json:"contact_phone,omitempty"`
}

// UpdateAdvertiser 更新广告主信息
func (s *AgentService) UpdateAdvertiser(ctx context.Context, accessToken string, req *UpdateAdvertiserRequest) error {
	return s.client.PostWithToken(ctx, accessToken, "/2/agent/advertiser/update/", req, nil)
}

// GetChildAgentList 获取子代理商列表
func (s *AgentService) GetChildAgentList(ctx context.Context, accessToken string, agentID int64, page, pageSize int) ([]AgentInfo, int, error) {
	params := map[string]interface{}{
		"agent_id":  agentID,
		"page":      page,
		"page_size": pageSize,
	}

	var result struct {
		Data struct {
			List     []AgentInfo `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/agent/child_agent/select/", params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// RefundRequest 退款请求
type RefundRequest struct {
	AgentID              int64   `json:"agent_id"`
	AdvertiserID         int64   `json:"advertiser_id"`
	Amount               float64 `json:"amount"`
	TransferSerialNumber string  `json:"transfer_serial_no,omitempty"`
}

// Refund 广告主退款到代理商
func (s *AgentService) Refund(ctx context.Context, accessToken string, req *RefundRequest) (*TransferResponse, error) {
	body := map[string]interface{}{
		"agent_id":      req.AgentID,
		"advertiser_id": req.AdvertiserID,
		"amount":        req.Amount,
	}
	if req.TransferSerialNumber != "" {
		body["transfer_serial_no"] = req.TransferSerialNumber
	}

	var result struct {
		Data TransferResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, "/2/agent/advertiser/refund/", body, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetDailyFundStat 获取日流水统计
func (s *AgentService) GetDailyFundStat(ctx context.Context, accessToken string, agentID int64, startDate, endDate string, page, pageSize int) ([]map[string]interface{}, int, error) {
	body := map[string]interface{}{
		"agent_id":   agentID,
		"start_date": startDate,
		"end_date":   endDate,
		"page":       page,
		"page_size":  pageSize,
	}

	resp, err := s.client.Post(ctx, "/2/agent/fund/daily_stat/", body)
	if err != nil {
		return nil, 0, err
	}

	if !resp.IsSuccess() {
		return nil, 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List     []map[string]interface{} `json:"list"`
		PageInfo struct {
			TotalNumber int `json:"total_number"`
		} `json:"page_info"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("unmarshal response failed: %w", err)
	}

	return result.List, result.PageInfo.TotalNumber, nil
}
