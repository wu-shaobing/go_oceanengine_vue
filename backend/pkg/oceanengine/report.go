package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// ReportService 报表服务
type ReportService struct {
	client *Client
}

// NewReportService 创建报表服务
func NewReportService(client *Client) *ReportService {
	return &ReportService{client: client}
}

// ReportRequest 报表请求
type ReportRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	TimeGranular string   `json:"time_granularity,omitempty"` // STAT_TIME_GRANULARITY_DAILY, STAT_TIME_GRANULARITY_HOURLY
	GroupBy      []string `json:"group_by,omitempty"`
	Filtering    struct {
		CampaignIDs []int64 `json:"campaign_ids,omitempty"`
		AdIDs       []int64 `json:"ad_ids,omitempty"`
		CreativeIDs []int64 `json:"creative_ids,omitempty"`
	} `json:"filtering,omitempty"`
	OrderField string `json:"order_field,omitempty"`
	OrderType  string `json:"order_type,omitempty"`
	Page       int    `json:"page,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

// ReportMetrics 报表指标
type ReportMetrics struct {
	StatDatetime string  `json:"stat_datetime"`
	Cost         float64 `json:"cost"`
	ShowCnt      int64   `json:"show"`
	ClickCnt     int64   `json:"click"`
	ConvertCnt   int64   `json:"convert"`
	CTR          float64 `json:"ctr"`
	CPC          float64 `json:"cpc"`
	CPM          float64 `json:"cpm"`
	ConvertRate  float64 `json:"convert_rate"`
	ConvertCost  float64 `json:"convert_cost"`
}

// AdvertiserReportRow 广告主报表行
type AdvertiserReportRow struct {
	AdvertiserID int64 `json:"advertiser_id"`
	ReportMetrics
}

// CampaignReportRow 广告系列报表行
type CampaignReportRow struct {
	CampaignID   int64 `json:"campaign_id"`
	AdvertiserID int64 `json:"advertiser_id"`
	ReportMetrics
}

// AdReportRow 广告报表行
type AdReportRow struct {
	AdID         int64 `json:"ad_id"`
	CampaignID   int64 `json:"campaign_id"`
	AdvertiserID int64 `json:"advertiser_id"`
	ReportMetrics
}

// CreativeReportRow 创意报表行
type CreativeReportRow struct {
	CreativeID   int64 `json:"creative_id"`
	AdID         int64 `json:"ad_id"`
	CampaignID   int64 `json:"campaign_id"`
	AdvertiserID int64 `json:"advertiser_id"`
	ReportMetrics
}

// ReportResponse 报表响应
type ReportResponse struct {
	List     json.RawMessage `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetAdvertiserReport 获取广告主报表
func (s *ReportService) GetAdvertiserReport(ctx context.Context, req *ReportRequest) ([]AdvertiserReportRow, *ReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/advertiser/get/", req)
	if err != nil {
		return nil, nil, err
	}

	if !resp.IsSuccess() {
		return nil, nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	var rows []AdvertiserReportRow
	if err := json.Unmarshal(result.List, &rows); err != nil {
		return nil, nil, fmt.Errorf("unmarshal list failed: %w", err)
	}

	return rows, &result, nil
}

// GetCampaignReport 获取广告系列报表
func (s *ReportService) GetCampaignReport(ctx context.Context, req *ReportRequest) ([]CampaignReportRow, *ReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/campaign/get/", req)
	if err != nil {
		return nil, nil, err
	}

	if !resp.IsSuccess() {
		return nil, nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	var rows []CampaignReportRow
	if err := json.Unmarshal(result.List, &rows); err != nil {
		return nil, nil, fmt.Errorf("unmarshal list failed: %w", err)
	}

	return rows, &result, nil
}

// GetAdReport 获取广告报表
func (s *ReportService) GetAdReport(ctx context.Context, req *ReportRequest) ([]AdReportRow, *ReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/ad/get/", req)
	if err != nil {
		return nil, nil, err
	}

	if !resp.IsSuccess() {
		return nil, nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	var rows []AdReportRow
	if err := json.Unmarshal(result.List, &rows); err != nil {
		return nil, nil, fmt.Errorf("unmarshal list failed: %w", err)
	}

	return rows, &result, nil
}

// GetCreativeReport 获取创意报表
func (s *ReportService) GetCreativeReport(ctx context.Context, req *ReportRequest) ([]CreativeReportRow, *ReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/creative/get/", req)
	if err != nil {
		return nil, nil, err
	}

	if !resp.IsSuccess() {
		return nil, nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	var rows []CreativeReportRow
	if err := json.Unmarshal(result.List, &rows); err != nil {
		return nil, nil, fmt.Errorf("unmarshal list failed: %w", err)
	}

	return rows, &result, nil
}
