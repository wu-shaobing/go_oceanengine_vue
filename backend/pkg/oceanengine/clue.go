package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// ClueService 飞鱼线索管理服务
type ClueService struct {
	client *Client
}

// NewClueService 创建飞鱼线索服务
func NewClueService(client *Client) *ClueService {
	return &ClueService{client: client}
}

// ==================== 线索管理 ====================

// Clue 线索信息
type Clue struct {
	ClueID          int64             `json:"clue_id"`
	AdvertiserID    int64             `json:"advertiser_id"`
	SiteID          int64             `json:"site_id"`
	SiteName        string            `json:"site_name"`
	TelephoneNumber string            `json:"telephone_number"`
	FormName        string            `json:"form_name"`
	Name            string            `json:"name"`
	Gender          int               `json:"gender"`
	Age             string            `json:"age"`
	City            string            `json:"city"`
	Address         string            `json:"address"`
	Remark          string            `json:"remark"`
	Email           string            `json:"email"`
	WechatID        string            `json:"wechat_id"`
	QQID            string            `json:"qq_id"`
	ExtraInfo       map[string]string `json:"extra_info"`
	Source          string            `json:"source"`
	CreateTime      string            `json:"create_time"`
	FlowType        string            `json:"flow_type"`
	AdID            int64             `json:"ad_id"`
	CreativeID      int64             `json:"creative_id"`
	CampaignID      int64             `json:"campaign_id"`
	CampaignName    string            `json:"campaign_name"`
	RequestID       string            `json:"request_id"`
	ClueType        int               `json:"clue_type"`
	ConvertStatus   int               `json:"convert_status"`
}

// ClueListRequest 获取线索列表请求
type ClueListRequest struct {
	AdvertiserID  int64  `json:"advertiser_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Page          int    `json:"page,omitempty"`
	PageSize      int    `json:"page_size,omitempty"`
	ClueType      int    `json:"clue_type,omitempty"`
	ConvertStatus int    `json:"convert_status,omitempty"`
}

// ClueListResponse 获取线索列表响应
type ClueListResponse struct {
	List     []Clue `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetClueList 获取线索列表
func (s *ClueService) GetClueList(ctx context.Context, req *ClueListRequest) (*ClueListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_time":    req.StartTime,
		"end_time":      req.EndTime,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.ClueType > 0 {
		params["clue_type"] = req.ClueType
	}
	if req.ConvertStatus > 0 {
		params["convert_status"] = req.ConvertStatus
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ClueListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ClueCallbackRequest 回传有效线索请求
type ClueCallbackRequest struct {
	AdvertiserID   int64  `json:"advertiser_id"`
	ClueID         int64  `json:"clue_id"`
	EventType      string `json:"event_type"`
	ConvertState   int    `json:"convert_state,omitempty"`
	PayAmount      int64  `json:"pay_amount,omitempty"`
	ExternalAction string `json:"external_action,omitempty"`
	OccurTime      string `json:"occur_time,omitempty"`
	Remark         string `json:"remark,omitempty"`
}

// ClueCallback 回传有效线索
func (s *ClueService) ClueCallback(ctx context.Context, req *ClueCallbackRequest) error {
	resp, err := s.client.Post(ctx, "/2/tools/clue/callback/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 活动记录 ====================

// KeyAction 活动记录
type KeyAction struct {
	ClueID       int64  `json:"clue_id"`
	ActionType   string `json:"action_type"`
	ActionName   string `json:"action_name"`
	ActionTime   string `json:"action_time"`
	ActionDetail string `json:"action_detail"`
}

// KeyActionGetRequest 获取活动记录请求
type KeyActionGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	ClueID       int64 `json:"clue_id"`
}

// KeyActionGetResponse 获取活动记录响应
type KeyActionGetResponse struct {
	List []KeyAction `json:"list"`
}

// GetKeyAction 获取活动记录
func (s *ClueService) GetKeyAction(ctx context.Context, req *KeyActionGetRequest) (*KeyActionGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"clue_id":       req.ClueID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/key_action/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result KeyActionGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 智能电话 ====================

// SmartPhone 智能电话
type SmartPhone struct {
	SmartPhoneID int64  `json:"smart_phone_id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	BindStatus   int    `json:"bind_status"`
	CreateTime   string `json:"create_time"`
	ModifyTime   string `json:"modify_time"`
}

// SmartPhoneGetRequest 查询智能电话请求
type SmartPhoneGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// SmartPhoneGetResponse 查询智能电话响应
type SmartPhoneGetResponse struct {
	List     []SmartPhone `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetSmartPhone 查询已有智能电话
func (s *ClueService) GetSmartPhone(ctx context.Context, req *SmartPhoneGetRequest) (*SmartPhoneGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/smart_phone/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result SmartPhoneGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 表单管理 ====================

// Form 表单信息
type Form struct {
	FormID        int64             `json:"form_id"`
	FormName      string            `json:"form_name"`
	FormType      string            `json:"form_type"`
	Status        int               `json:"status"`
	CreateTime    string            `json:"create_time"`
	ModifyTime    string            `json:"modify_time"`
	FormElements  []FormElement     `json:"form_elements"`
	ButtonText    string            `json:"button_text"`
	SubmitMessage string            `json:"submit_message"`
	ExtraSettings map[string]string `json:"extra_settings"`
}

// FormElement 表单元素
type FormElement struct {
	ElementID   int64    `json:"element_id"`
	ElementType string   `json:"element_type"`
	ElementName string   `json:"element_name"`
	Required    bool     `json:"required"`
	Options     []string `json:"options,omitempty"`
	Placeholder string   `json:"placeholder,omitempty"`
}

// FormGetRequest 查询表单列表请求
type FormGetRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	FormType     string `json:"form_type,omitempty"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// FormGetResponse 查询表单列表响应
type FormGetResponse struct {
	List     []Form `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetFormList 查询已有表单列表
func (s *ClueService) GetFormList(ctx context.Context, req *FormGetRequest) (*FormGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.FormType != "" {
		params["form_type"] = req.FormType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/form/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result FormGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// FormDetailRequest 查询表单详情请求
type FormDetailRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	FormID       int64 `json:"form_id"`
}

// GetFormDetail 查询表单详情
func (s *ClueService) GetFormDetail(ctx context.Context, req *FormDetailRequest) (*Form, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"form_id":       req.FormID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/form/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result Form
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 线索店铺 ====================

// ClueStore 线索店铺信息
type ClueStore struct {
	StoreID     int64  `json:"store_id"`
	StoreName   string `json:"store_name"`
	Address     string `json:"address"`
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
	CreateTime  string `json:"create_time"`
}

// ClueStoreListRequest 获取店铺列表请求
type ClueStoreListRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	StoreName    string `json:"store_name,omitempty"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// ClueStoreListResponse 获取店铺列表响应
type ClueStoreListResponse struct {
	List     []ClueStore `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetClueStoreList 获取线索店铺列表
func (s *ClueService) GetClueStoreList(ctx context.Context, req *ClueStoreListRequest) (*ClueStoreListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.StoreName != "" {
		params["store_name"] = req.StoreName
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/tools/clue/store/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ClueStoreListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 批量线索回传 ====================

// BatchClueCallbackRequest 批量回传线索请求
type BatchClueCallbackRequest struct {
	AdvertiserID int64                `json:"advertiser_id"`
	ClueList     []ClueCallbackDetail `json:"clue_list"`
}

// ClueCallbackDetail 线索回传详情
type ClueCallbackDetail struct {
	ClueID         int64  `json:"clue_id"`
	EventType      string `json:"event_type"`
	ConvertState   int    `json:"convert_state,omitempty"`
	PayAmount      int64  `json:"pay_amount,omitempty"`
	ExternalAction string `json:"external_action,omitempty"`
	OccurTime      string `json:"occur_time,omitempty"`
}

// BatchClueCallbackResponse 批量回传线索响应
type BatchClueCallbackResponse struct {
	SuccessCount int `json:"success_count"`
	FailCount    int `json:"fail_count"`
	FailList     []struct {
		ClueID  int64  `json:"clue_id"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"fail_list"`
}

// BatchClueCallback 批量回传有效线索
func (s *ClueService) BatchClueCallback(ctx context.Context, req *BatchClueCallbackRequest) (*BatchClueCallbackResponse, error) {
	resp, err := s.client.Post(ctx, "/2/tools/clue/batch_callback/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result BatchClueCallbackResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}
