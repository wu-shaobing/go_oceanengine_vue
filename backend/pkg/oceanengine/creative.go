package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// CreativeService 创意服务
type CreativeService struct {
	client *Client
}

// NewCreativeService 创建创意服务
func NewCreativeService(client *Client) *CreativeService {
	return &CreativeService{client: client}
}

// Creative 创意
type Creative struct {
	CreativeID      int64    `json:"creative_id"`
	AdID            int64    `json:"ad_id"`
	CampaignID      int64    `json:"campaign_id"`
	AdvertiserID    int64    `json:"advertiser_id"`
	Title           string   `json:"title"`
	Source          string   `json:"source"`
	ImageMode       string   `json:"image_mode"`
	Status          string   `json:"status"`
	OptStatus       string   `json:"opt_status"`
	ImageIDs        []string `json:"image_ids"`
	VideoID         string   `json:"video_id"`
	ThirdIndustryID int64    `json:"third_industry_id"`
	ActionType      string   `json:"ad_download_type"`
	AdDownloadURL   string   `json:"ad_download_url"`
	AppName         string   `json:"app_name"`
	WebURL          string   `json:"web_url"`
	CreateTime      string   `json:"creative_create_time"`
	ModifyTime      string   `json:"creative_modify_time"`
}

// CreativeListRequest 创意列表请求
type CreativeListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		CreativeIDs []int64 `json:"creative_ids,omitempty"`
		AdIDs       []int64 `json:"ad_ids,omitempty"`
		CampaignIDs []int64 `json:"campaign_ids,omitempty"`
		Title       string  `json:"title,omitempty"`
		Status      string  `json:"status,omitempty"`
		ImageMode   string  `json:"image_mode,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// CreativeListResponse 创意列表响应
type CreativeListResponse struct {
	List     []Creative `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetList 获取创意列表
func (s *CreativeService) GetList(ctx context.Context, req *CreativeListRequest) (*CreativeListResponse, error) {
	resp, err := s.client.Post(ctx, "/2/creative/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CreativeListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CreativeCreateRequest 创建创意请求
type CreativeCreateRequest struct {
	AdvertiserID         int64    `json:"advertiser_id"`
	AdID                 int64    `json:"ad_id"`
	Title                string   `json:"title,omitempty"`
	Source               string   `json:"source,omitempty"`
	ImageMode            string   `json:"image_mode"`
	ImageIDs             []string `json:"image_ids,omitempty"`
	VideoID              string   `json:"video_id,omitempty"`
	ThirdIndustryID      int64    `json:"third_industry_id,omitempty"`
	ActionType           string   `json:"ad_download_type,omitempty"`
	AdDownloadURL        string   `json:"ad_download_url,omitempty"`
	AppName              string   `json:"app_name,omitempty"`
	WebURL               string   `json:"web_url,omitempty"`
	CallToAction         string   `json:"call_to_action,omitempty"`
	CreativeDisplayMode  string   `json:"creative_display_mode,omitempty"`
	AdvancedCreativeType string   `json:"advanced_creative_type,omitempty"`
	AppThumbnailIDs      []string `json:"app_thumbnail_ids,omitempty"`
	TitleList            []string `json:"title_list,omitempty"`
	DescriptionList      []string `json:"description_list,omitempty"`
	SubTitle             string   `json:"sub_title,omitempty"`
	VideoPlayEffectType  string   `json:"video_play_effect_type,omitempty"`
}

// CreativeCreateResponse 创建创意响应
type CreativeCreateResponse struct {
	CreativeID int64 `json:"creative_id"`
}

// Create 创建创意
func (s *CreativeService) Create(ctx context.Context, req *CreativeCreateRequest) (*CreativeCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/creative/create_v2/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CreativeCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CustomCreativeCreateRequest 创建自定义创意请求
type CustomCreativeCreateRequest struct {
	AdvertiserID     int64    `json:"advertiser_id"`
	AdID             int64    `json:"ad_id"`
	ImageMode        string   `json:"image_mode"`
	Title            string   `json:"title,omitempty"`
	Source           string   `json:"source,omitempty"`
	ImageIDs         []string `json:"image_ids,omitempty"`
	VideoID          string   `json:"video_id,omitempty"`
	ThirdIndustryID  int64    `json:"third_industry_id,omitempty"`
	CarouselSettings []struct {
		ImageID string `json:"image_id"`
		Title   string `json:"title,omitempty"`
		WebURL  string `json:"web_url,omitempty"`
	} `json:"carousel_settings,omitempty"`
	ActionText string `json:"action_text,omitempty"`
	ActionURL  string `json:"action_url,omitempty"`
}

// CustomCreativeCreate 创建自定义创意
func (s *CreativeService) CustomCreativeCreate(ctx context.Context, req *CustomCreativeCreateRequest) (*CreativeCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/creative/custom/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CreativeCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CreativeUpdateRequest 更新创意请求
type CreativeUpdateRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	CreativeID   int64    `json:"creative_id"`
	Title        string   `json:"title,omitempty"`
	Source       string   `json:"source,omitempty"`
	ImageIDs     []string `json:"image_ids,omitempty"`
	VideoID      string   `json:"video_id,omitempty"`
	ActionText   string   `json:"action_text,omitempty"`
	ActionURL    string   `json:"action_url,omitempty"`
}

// CreativeUpdateResponse 更新创意响应
type CreativeUpdateResponse struct {
	CreativeID int64 `json:"creative_id"`
}

// Update 更新创意
func (s *CreativeService) Update(ctx context.Context, req *CreativeUpdateRequest) (*CreativeUpdateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/creative/update/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CreativeUpdateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CreativeStatusUpdateRequest 更新创意状态请求
type CreativeStatusUpdateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CreativeIDs  []int64 `json:"creative_ids"`
	OptStatus    string  `json:"opt_status"` // enable, disable
}

// CreativeStatusUpdateResponse 更新状态响应
type CreativeStatusUpdateResponse struct {
	CreativeIDs []int64 `json:"creative_ids"`
	Errors      []struct {
		CreativeID int64  `json:"creative_id"`
		Code       int    `json:"code"`
		Message    string `json:"message"`
	} `json:"errors"`
}

// UpdateStatus 更新创意状态
func (s *CreativeService) UpdateStatus(ctx context.Context, req *CreativeStatusUpdateRequest) (*CreativeStatusUpdateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/creative/update/status/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CreativeStatusUpdateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CreativeDetailRequest 创意详情请求
type CreativeDetailRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CreativeIDs  []int64 `json:"creative_ids"`
}

// CreativeDetail 创意详情
type CreativeDetail struct {
	CreativeID      int64    `json:"creative_id"`
	AdID            int64    `json:"ad_id"`
	CampaignID      int64    `json:"campaign_id"`
	AdvertiserID    int64    `json:"advertiser_id"`
	Title           string   `json:"title"`
	Source          string   `json:"source"`
	ImageMode       string   `json:"image_mode"`
	Status          string   `json:"status"`
	OptStatus       string   `json:"opt_status"`
	ImageIDs        []string `json:"image_ids"`
	VideoID         string   `json:"video_id"`
	ThirdIndustryID int64    `json:"third_industry_id"`
	ActionType      string   `json:"ad_download_type"`
	AdDownloadURL   string   `json:"ad_download_url"`
	AppName         string   `json:"app_name"`
	WebURL          string   `json:"web_url"`
	ActionText      string   `json:"action_text"`
	ActionURL       string   `json:"action_url"`
	TitleList       []string `json:"title_list"`
	ImageList       []string `json:"image_list"`
	DescriptionList []string `json:"description_list"`
	CreateTime      string   `json:"creative_create_time"`
	ModifyTime      string   `json:"creative_modify_time"`
}

// GetDetail 获取创意详情
func (s *CreativeService) GetDetail(ctx context.Context, req *CreativeDetailRequest) ([]CreativeDetail, error) {
	resp, err := s.client.Post(ctx, "/2/creative/read/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []CreativeDetail `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// CreativeRejectReasonRequest 创意审核建议请求
type CreativeRejectReasonRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CreativeIDs  []int64 `json:"creative_ids"`
}

// CreativeRejectReason 创意审核建议
type CreativeRejectReason struct {
	CreativeID   int64  `json:"creative_id"`
	RejectReason string `json:"reject_reason"`
}

// GetRejectReason 获取创意审核建议
func (s *CreativeService) GetRejectReason(ctx context.Context, req *CreativeRejectReasonRequest) ([]CreativeRejectReason, error) {
	resp, err := s.client.Post(ctx, "/2/creative/reject_reason/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []CreativeRejectReason `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// CreativeMaterialReadRequest 创意素材信息请求
type CreativeMaterialReadRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	CreativeIDs  []int64 `json:"creative_ids"`
}

// CreativeMaterial 创意素材信息
type CreativeMaterial struct {
	CreativeID   int64    `json:"creative_id"`
	ImageIDs     []string `json:"image_ids"`
	ImageURLs    []string `json:"image_urls"`
	VideoID      string   `json:"video_id"`
	VideoURL     string   `json:"video_url"`
	ThumbImageID string   `json:"thumb_image_id"`
	ThumbURL     string   `json:"thumb_url"`
}

// GetMaterial 获取创意素材信息
func (s *CreativeService) GetMaterial(ctx context.Context, req *CreativeMaterialReadRequest) ([]CreativeMaterial, error) {
	resp, err := s.client.Post(ctx, "/2/creative/material/read/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []CreativeMaterial `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ProceduralCreativeCreateRequest 创建程序化创意请求
type ProceduralCreativeCreateRequest struct {
	AdvertiserID    int64    `json:"advertiser_id"`
	AdID            int64    `json:"ad_id"`
	TitleList       []string `json:"title_list,omitempty"`
	ImageList       []string `json:"image_list,omitempty"`
	DescriptionList []string `json:"description_list,omitempty"`
	VideoList       []string `json:"video_list,omitempty"`
	Source          string   `json:"source,omitempty"`
}

// ProceduralCreativeCreate 创建程序化创意
func (s *CreativeService) ProceduralCreativeCreate(ctx context.Context, req *ProceduralCreativeCreateRequest) error {
	resp, err := s.client.Post(ctx, "/2/creative/procedural/create/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ProceduralCreativeUpdateRequest 更新程序化创意请求
type ProceduralCreativeUpdateRequest struct {
	AdvertiserID    int64    `json:"advertiser_id"`
	AdID            int64    `json:"ad_id"`
	TitleList       []string `json:"title_list,omitempty"`
	ImageList       []string `json:"image_list,omitempty"`
	DescriptionList []string `json:"description_list,omitempty"`
	VideoList       []string `json:"video_list,omitempty"`
	Source          string   `json:"source,omitempty"`
}

// ProceduralCreativeUpdate 更新程序化创意
func (s *CreativeService) ProceduralCreativeUpdate(ctx context.Context, req *ProceduralCreativeUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/creative/procedural/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}
