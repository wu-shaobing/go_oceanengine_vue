package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// AsyncReportService 异步报表服务
type AsyncReportService struct {
	client *Client
}

// NewAsyncReportService 创建异步报表服务
func NewAsyncReportService(client *Client) *AsyncReportService {
	return &AsyncReportService{client: client}
}

// AsyncReportTaskCreateRequest 创建异步报表任务请求
type AsyncReportTaskCreateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	TaskName     string `json:"task_name,omitempty"`
	TaskType     string `json:"task_type"` // REPORT 固定值
	TaskParams   struct {
		DataLevel       string   `json:"data_level"`                 // ACCOUNT, CAMPAIGN, AD, CREATIVE
		DataType        string   `json:"data_type,omitempty"`        // EFFECT, BASIC
		TimeGranularity string   `json:"time_granularity,omitempty"` // HOURLY, DAILY, TOTAL
		StartDate       string   `json:"start_date"`
		EndDate         string   `json:"end_date"`
		Fields          []string `json:"fields,omitempty"`
		Filtering       struct {
			CampaignIDs []int64 `json:"campaign_ids,omitempty"`
			AdIDs       []int64 `json:"ad_ids,omitempty"`
			CreativeIDs []int64 `json:"creative_ids,omitempty"`
		} `json:"filtering,omitempty"`
		OrderField string `json:"order_field,omitempty"`
		OrderType  string `json:"order_type,omitempty"` // ASC, DESC
	} `json:"task_params"`
}

// AsyncReportTaskCreateResponse 创建异步报表任务响应
type AsyncReportTaskCreateResponse struct {
	TaskID string `json:"task_id"`
}

// CreateTask 创建异步报表任务
func (s *AsyncReportService) CreateTask(ctx context.Context, req *AsyncReportTaskCreateRequest) (*AsyncReportTaskCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/async_task/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AsyncReportTaskCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AsyncReportTaskGetRequest 获取异步报表任务请求
type AsyncReportTaskGetRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	TaskIDs      []string `json:"task_ids,omitempty"`
	Filtering    struct {
		TaskName   string `json:"task_name,omitempty"`
		TaskType   string `json:"task_type,omitempty"`
		TaskStatus string `json:"task_status,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// AsyncReportTask 异步报表任务
type AsyncReportTask struct {
	TaskID       string `json:"task_id"`
	TaskName     string `json:"task_name"`
	TaskType     string `json:"task_type"`
	TaskStatus   string `json:"task_status"` // PROCESSING, SUCCESS, FAILED
	CreateTime   string `json:"create_time"`
	DownloadURL  string `json:"download_url,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// AsyncReportTaskGetResponse 获取异步报表任务响应
type AsyncReportTaskGetResponse struct {
	List     []AsyncReportTask `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetTasks 获取异步报表任务列表
func (s *AsyncReportService) GetTasks(ctx context.Context, req *AsyncReportTaskGetRequest) (*AsyncReportTaskGetResponse, error) {
	resp, err := s.client.Post(ctx, "/2/async_task/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AsyncReportTaskGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AsyncReportTaskDownloadRequest 下载异步报表任务请求
type AsyncReportTaskDownloadRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	TaskID       string `json:"task_id"`
}

// AsyncReportTaskDownloadResponse 下载异步报表任务响应
type AsyncReportTaskDownloadResponse struct {
	List     []map[string]interface{} `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// Download 下载异步报表任务结果
func (s *AsyncReportService) Download(ctx context.Context, req *AsyncReportTaskDownloadRequest) (*AsyncReportTaskDownloadResponse, error) {
	resp, err := s.client.Post(ctx, "/2/async_task/download/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AsyncReportTaskDownloadResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// IntegratedReportRequest 统一报表请求
type IntegratedReportRequest struct {
	AdvertiserID    int64    `json:"advertiser_id"`
	StartDate       string   `json:"start_date"`
	EndDate         string   `json:"end_date"`
	TimeGranularity string   `json:"time_granularity,omitempty"` // HOURLY, DAILY, TOTAL
	GroupBy         []string `json:"group_by,omitempty"`
	Fields          []string `json:"fields,omitempty"`
	Filtering       struct {
		CampaignIDs []int64 `json:"campaign_ids,omitempty"`
		AdIDs       []int64 `json:"ad_ids,omitempty"`
		CreativeIDs []int64 `json:"creative_ids,omitempty"`
	} `json:"filtering,omitempty"`
	OrderField string `json:"order_field,omitempty"`
	OrderType  string `json:"order_type,omitempty"` // ASC, DESC
	Page       int    `json:"page,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

// IntegratedReportResponse 统一报表响应
type IntegratedReportResponse struct {
	List     []map[string]interface{} `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetIntegratedReport 获取统一报表 (同步接口)
func (s *AsyncReportService) GetIntegratedReport(ctx context.Context, req *IntegratedReportRequest) (*IntegratedReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/integrated/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result IntegratedReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// VideoAnalysisReportRequest 视频分析报表请求
type VideoAnalysisReportRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	MaterialIDs  []string `json:"material_ids,omitempty"`
	Fields       []string `json:"fields,omitempty"`
	Page         int      `json:"page,omitempty"`
	PageSize     int      `json:"page_size,omitempty"`
}

// VideoAnalysisReportResponse 视频分析报表响应
type VideoAnalysisReportResponse struct {
	List     []map[string]interface{} `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetVideoAnalysisReport 获取视频分析报表
func (s *AsyncReportService) GetVideoAnalysisReport(ctx context.Context, req *VideoAnalysisReportRequest) (*VideoAnalysisReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/video/analysis/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result VideoAnalysisReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AudienceAnalysisReportRequest 受众分析报表请求
type AudienceAnalysisReportRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Dimensions   []string `json:"dimensions,omitempty"` // age, gender, city, etc.
}

// AudienceAnalysisReportResponse 受众分析报表响应
type AudienceAnalysisReportResponse struct {
	List []struct {
		Dimension string `json:"dimension"`
		Data      []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Ratio float64 `json:"ratio"`
		} `json:"data"`
	} `json:"list"`
}

// GetAudienceAnalysisReport 获取受众分析报表
func (s *AsyncReportService) GetAudienceAnalysisReport(ctx context.Context, req *AudienceAnalysisReportRequest) (*AudienceAnalysisReportResponse, error) {
	resp, err := s.client.Post(ctx, "/2/report/audience/analysis/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AudienceAnalysisReportResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}
