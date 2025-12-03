package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// AdvertiserService 广告主服务
type AdvertiserService struct {
	client *Client
}

// NewAdvertiserService 创建广告主服务
func NewAdvertiserService(client *Client) *AdvertiserService {
	return &AdvertiserService{client: client}
}

// AdvertiserInfo 广告主信息
type AdvertiserInfo struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	Company    string `json:"company"`
	Address    string `json:"address"`
	LicenseURL string `json:"license_url"`
	LicenseNo  string `json:"license_no"`
}

// AdvertiserInfoResponse 广告主信息响应
type AdvertiserInfoResponse struct {
	List []AdvertiserInfo `json:"list"`
}

// GetInfo 获取广告主信息
func (s *AdvertiserService) GetInfo(ctx context.Context, advertiserIDs []int64) ([]AdvertiserInfo, error) {
	idsStr := make([]string, len(advertiserIDs))
	for i, id := range advertiserIDs {
		idsStr[i] = strconv.FormatInt(id, 10)
	}

	body := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	resp, err := s.client.Post(ctx, "/2/advertiser/info/", body)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AdvertiserInfoResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// FundInfo 资金信息
type FundInfo struct {
	Balance          float64 `json:"balance"`
	ValidBalance     float64 `json:"valid_balance"`
	CashBalance      float64 `json:"cash_balance"`
	GrantBalance     float64 `json:"grant_balance"`
	ReturnGoodsCoast float64 `json:"return_goods_coast"`
}

// FundInfoResponse 资金信息响应
type FundInfoResponse struct {
	List []struct {
		AdvertiserID int64    `json:"advertiser_id"`
		FundInfo     FundInfo `json:"fund_info"`
	} `json:"list"`
}

// GetFund 获取资金信息
func (s *AdvertiserService) GetFund(ctx context.Context, advertiserIDs []int64) (map[int64]FundInfo, error) {
	body := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	resp, err := s.client.Post(ctx, "/2/advertiser/fund/get/", body)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result FundInfoResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	funds := make(map[int64]FundInfo)
	for _, item := range result.List {
		funds[item.AdvertiserID] = item.FundInfo
	}

	return funds, nil
}

// FundTransaction 资金流水
type FundTransaction struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	TransactionTime string  `json:"transaction_time"`
	Remark          string  `json:"remark"`
	BalanceAfter    float64 `json:"balance_after"`
}

// FundTransactionRequest 资金流水请求
type FundTransactionRequest struct {
	AdvertiserID    int64  `json:"advertiser_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	TransactionType string `json:"transaction_type,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        int    `json:"page_size,omitempty"`
}

// FundTransactionResponse 资金流水响应
type FundTransactionResponse struct {
	List     []FundTransaction `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetFundTransaction 获取资金流水
func (s *AdvertiserService) GetFundTransaction(ctx context.Context, req *FundTransactionRequest) (*FundTransactionResponse, error) {
	resp, err := s.client.Post(ctx, "/2/advertiser/fund/transaction/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result FundTransactionResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// Advertiser 创建广告主服务(链式调用)
func (c *Client) Advertiser() *AdvertiserService {
	return NewAdvertiserService(c)
}

// ==================== 资金管理扩展 ====================

// DailyFundStat 日资金统计
type DailyFundStat struct {
	Date           string  `json:"date"`
	Balance        float64 `json:"balance"`
	Cost           float64 `json:"cost"`
	CashCost       float64 `json:"cash_cost"`
	Reward         float64 `json:"reward"`
	TransferIn     float64 `json:"transfer_in"`
	TransferOut    float64 `json:"transfer_out"`
	FreezeAmount   float64 `json:"freeze_amount"`
	UnfreezeAmount float64 `json:"unfreeze_amount"`
}

// DailyFundStatRequest 日资金统计请求
type DailyFundStatRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// GetDailyFundStat 获取日资金统计
func (s *AdvertiserService) GetDailyFundStat(ctx context.Context, accessToken string, req *DailyFundStatRequest) ([]DailyFundStat, int, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_date":    req.StartDate,
		"end_date":      req.EndDate,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	var result struct {
		Data struct {
			List     []DailyFundStat `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/fund/daily_stat/", params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetFundWithToken 获取资金信息(带Token)
func (s *AdvertiserService) GetFundWithToken(ctx context.Context, accessToken string, advertiserIDs []int64) (map[int64]FundInfo, error) {
	params := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	var result struct {
		Data FundInfoResponse `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/fund/get/", params, &result)
	if err != nil {
		return nil, err
	}

	funds := make(map[int64]FundInfo)
	for _, item := range result.Data.List {
		funds[item.AdvertiserID] = item.FundInfo
	}
	return funds, nil
}

// GetInfoWithToken 获取广告主信息(带Token)
func (s *AdvertiserService) GetInfoWithToken(ctx context.Context, accessToken string, advertiserIDs []int64) ([]AdvertiserInfo, error) {
	params := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	var result struct {
		Data AdvertiserInfoResponse `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/info/", params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetTransactionDetail 获取交易详情
func (s *AdvertiserService) GetTransactionDetail(ctx context.Context, accessToken string, advertiserID int64, transactionSeq string) (*FundTransaction, error) {
	params := map[string]interface{}{
		"advertiser_id":   advertiserID,
		"transaction_seq": transactionSeq,
	}

	var result struct {
		Data FundTransaction `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/fund/transaction/detail/", params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 账户管理 ====================

// BudgetInfo 预算信息
type BudgetInfo struct {
	Budget     float64 `json:"budget"`
	BudgetMode string  `json:"budget_mode"`
}

// GetBudget 获取账户日预算
func (s *AdvertiserService) GetBudget(ctx context.Context, accessToken string, advertiserID int64) (*BudgetInfo, error) {
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data BudgetInfo `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/budget/get/", params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateBudget 更新账户日预算
func (s *AdvertiserService) UpdateBudget(ctx context.Context, accessToken string, advertiserID int64, budget float64, budgetMode string) error {
	body := map[string]interface{}{
		"advertiser_id": advertiserID,
		"budget":        budget,
	}
	if budgetMode != "" {
		body["budget_mode"] = budgetMode
	}
	return s.client.PostWithToken(ctx, accessToken, "/2/advertiser/budget/update/", body, nil)
}

// AdvertiserPublicInfo 广告主公开信息
type AdvertiserPublicInfo struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Company        string  `json:"company"`
	FirstIndustry  string  `json:"first_industry_name"`
	SecondIndustry string  `json:"second_industry_name"`
	Role           string  `json:"role"`
	Status         string  `json:"status"`
	Balance        float64 `json:"balance"`
	ValidBalance   float64 `json:"valid_balance"`
	CreateTime     string  `json:"create_time"`
}

// GetPublicInfo 获取广告主公开信息
func (s *AdvertiserService) GetPublicInfo(ctx context.Context, accessToken string, advertiserIDs []int64) ([]AdvertiserPublicInfo, error) {
	params := map[string]interface{}{
		"advertiser_ids": advertiserIDs,
	}

	var result struct {
		Data struct {
			List []AdvertiserPublicInfo `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, "/2/advertiser/public_info/", params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}
