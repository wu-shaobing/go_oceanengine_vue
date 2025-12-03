package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// CampaignService 广告系列服务
type CampaignService struct {
	client *Client
}

// NewCampaignService 创建广告系列服务
func NewCampaignService(client *Client) *CampaignService {
	return &CampaignService{client: client}
}

// Campaign 广告系列
type Campaign struct {
	ID           int64   `json:"campaign_id"`
	Name         string  `json:"campaign_name"`
	Budget       float64 `json:"budget"`
	BudgetMode   string  `json:"budget_mode"`
	LandingType  string  `json:"landing_type"`
	Status       string  `json:"status"`
	OptStatus    string  `json:"opt_status"`
	CreateTime   string  `json:"campaign_create_time"`
	ModifyTime   string  `json:"campaign_modify_time"`
	AdvertiserID int64   `json:"advertiser_id"`
}

// CampaignListRequest 广告系列列表请求
type CampaignListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		IDs            []int64 `json:"ids,omitempty"`
		CampaignName   string  `json:"campaign_name,omitempty"`
		LandingType    string  `json:"landing_type,omitempty"`
		Status         string  `json:"status,omitempty"`
		CampaignStatus string  `json:"campaign_status,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// CampaignListResponse 广告系列列表响应
type CampaignListResponse struct {
	List     []Campaign `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetList 获取广告系列列表
func (s *CampaignService) GetList(ctx context.Context, req *CampaignListRequest) (*CampaignListResponse, error) {
	resp, err := s.client.Post(ctx, "/2/campaign/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CampaignListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CampaignCreateRequest 创建广告系列请求
type CampaignCreateRequest struct {
	AdvertiserID  int64   `json:"advertiser_id"`
	CampaignName  string  `json:"campaign_name"`
	Budget        float64 `json:"budget,omitempty"`
	BudgetMode    string  `json:"budget_mode"`
	LandingType   string  `json:"landing_type"`
	UniqueFK      string  `json:"unique_fk,omitempty"`
	OperationType string  `json:"operation,omitempty"`
}

// CampaignCreateResponse 创建广告系列响应
type CampaignCreateResponse struct {
	CampaignID int64 `json:"campaign_id"`
}

// Create 创建广告系列
func (s *CampaignService) Create(ctx context.Context, req *CampaignCreateRequest) (*CampaignCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/campaign/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CampaignCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CampaignUpdateRequest 更新广告系列请求
type CampaignUpdateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CampaignID   int64   `json:"campaign_id"`
	CampaignName string  `json:"campaign_name,omitempty"`
	Budget       float64 `json:"budget,omitempty"`
	BudgetMode   string  `json:"budget_mode,omitempty"`
}

// Update 更新广告系列
func (s *CampaignService) Update(ctx context.Context, req *CampaignUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/campaign/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// CampaignStatusUpdateRequest 更新状态请求
type CampaignStatusUpdateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CampaignIDs  []int64 `json:"campaign_ids"`
	OptStatus    string  `json:"opt_status"` // enable, disable
}

// UpdateStatus 更新广告系列状态
func (s *CampaignService) UpdateStatus(ctx context.Context, req *CampaignStatusUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/campaign/update/status/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// Ad 广告组
type Ad struct {
	ID            int64   `json:"ad_id"`
	Name          string  `json:"name"`
	CampaignID    int64   `json:"campaign_id"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	Bid           float64 `json:"bid"`
	Pricing       string  `json:"pricing"`
	Status        string  `json:"status"`
	OptStatus     string  `json:"opt_status"`
	DeliveryRange string  `json:"delivery_range"`
	CreateTime    string  `json:"ad_create_time"`
	ModifyTime    string  `json:"ad_modify_time"`
	AdvertiserID  int64   `json:"advertiser_id"`
}

// AdListRequest 广告列表请求
type AdListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		IDs         []int64 `json:"ids,omitempty"`
		AdName      string  `json:"ad_name,omitempty"`
		CampaignIDs []int64 `json:"campaign_ids,omitempty"`
		Status      string  `json:"status,omitempty"`
		AdStatus    string  `json:"ad_status,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// AdListResponse 广告列表响应
type AdListResponse struct {
	List     []Ad `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// AdService 广告服务
type AdService struct {
	client *Client
}

// NewAdService 创建广告服务
func NewAdService(client *Client) *AdService {
	return &AdService{client: client}
}

// GetList 获取广告列表
func (s *AdService) GetList(ctx context.Context, req *AdListRequest) (*AdListResponse, error) {
	resp, err := s.client.Post(ctx, "/2/ad/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AdListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AdStatusUpdateRequest 更新广告状态请求
type AdStatusUpdateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
	OptStatus    string  `json:"opt_status"` // enable, disable
}

// UpdateStatus 更新广告状态
func (s *AdService) UpdateStatus(ctx context.Context, req *AdStatusUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/ad/update/status/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}
