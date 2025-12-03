package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// EventManagerService 事件管理服务
type EventManagerService struct {
	client *Client
}

// NewEventManagerService 创建事件管理服务
func NewEventManagerService(client *Client) *EventManagerService {
	return &EventManagerService{client: client}
}

// ==================== 资产管理 ====================

// Asset 资产信息
type Asset struct {
	AssetID      int64  `json:"asset_id"`
	AssetName    string `json:"asset_name"`
	AssetType    string `json:"asset_type"`
	Status       int    `json:"status"`
	LandingType  string `json:"landing_type"`
	CreateTime   string `json:"create_time"`
	ModifyTime   string `json:"modify_time"`
	AdvertiserID int64  `json:"advertiser_id"`
	AppType      string `json:"app_type,omitempty"`
	PackageName  string `json:"package_name,omitempty"`
	DownloadURL  string `json:"download_url,omitempty"`
}

// AssetDetail 资产详情
type AssetDetail struct {
	Asset
	EventConfigs []EventConfig `json:"event_configs"`
	TrackURLs    []TrackURL    `json:"track_urls"`
}

// AssetsGetRequest 获取资产列表请求
type AssetsGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AssetIDs     []int64 `json:"asset_ids,omitempty"`
	AssetType    string  `json:"asset_type,omitempty"`
	LandingType  string  `json:"landing_type,omitempty"`
	Page         int     `json:"page,omitempty"`
	PageSize     int     `json:"page_size,omitempty"`
}

// AssetsGetResponse 获取资产列表响应
type AssetsGetResponse struct {
	List     []Asset `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetAssets 获取已创建资产列表
func (s *EventManagerService) GetAssets(ctx context.Context, req *AssetsGetRequest) (*AssetsGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if len(req.AssetIDs) > 0 {
		params["asset_ids"] = req.AssetIDs
	}
	if req.AssetType != "" {
		params["asset_type"] = req.AssetType
	}
	if req.LandingType != "" {
		params["landing_type"] = req.LandingType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/assets/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AssetsGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AllAssetsListRequest 获取账户下资产列表请求
type AllAssetsListRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	AssetType    string `json:"asset_type,omitempty"`
	LandingType  string `json:"landing_type,omitempty"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// AllAssetsListResponse 获取账户下资产列表响应
type AllAssetsListResponse struct {
	List     []Asset `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetAllAssetsList 获取账户下资产列表（新）
func (s *EventManagerService) GetAllAssetsList(ctx context.Context, req *AllAssetsListRequest) (*AllAssetsListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.AssetType != "" {
		params["asset_type"] = req.AssetType
	}
	if req.LandingType != "" {
		params["landing_type"] = req.LandingType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/all_assets/list/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AllAssetsListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AllAssetsDetailRequest 获取资产详情请求
type AllAssetsDetailRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AssetIDs     []int64 `json:"asset_ids"`
}

// GetAllAssetsDetail 获取已创建资产详情（新）
func (s *EventManagerService) GetAllAssetsDetail(ctx context.Context, req *AllAssetsDetailRequest) ([]AssetDetail, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_ids":     req.AssetIDs,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/all_assets/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AssetDetail `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// AssetsCreateRequest 创建事件资产请求
type AssetsCreateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	AssetName    string `json:"asset_name"`
	AssetType    string `json:"asset_type"`
	LandingType  string `json:"landing_type"`
	AppType      string `json:"app_type,omitempty"`
	PackageName  string `json:"package_name,omitempty"`
	DownloadURL  string `json:"download_url,omitempty"`
}

// CreateAsset 创建事件资产
func (s *EventManagerService) CreateAsset(ctx context.Context, req *AssetsCreateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/event_manager/assets/create/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AssetID int64 `json:"asset_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AssetID, nil
}

// ==================== 事件管理 ====================

// EventConfig 事件配置
type EventConfig struct {
	EventID      int64  `json:"event_id"`
	EventName    string `json:"event_name"`
	EventType    string `json:"event_type"`
	OptimizeGoal string `json:"optimize_goal"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// AvailableEventsGetRequest 获取可创建事件列表请求
type AvailableEventsGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AssetID      int64 `json:"asset_id"`
}

// GetAvailableEvents 获取可创建事件列表
func (s *EventManagerService) GetAvailableEvents(ctx context.Context, req *AvailableEventsGetRequest) ([]EventConfig, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_id":      req.AssetID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/available_events/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []EventConfig `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// EventConfigsGetRequest 获取已创建事件列表请求
type EventConfigsGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AssetID      int64 `json:"asset_id"`
}

// GetEventConfigs 获取已创建事件列表
func (s *EventManagerService) GetEventConfigs(ctx context.Context, req *EventConfigsGetRequest) ([]EventConfig, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_id":      req.AssetID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/event_configs/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []EventConfig `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// EventsCreateRequest 资产下创建事件请求
type EventsCreateRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	AssetID      int64    `json:"asset_id"`
	EventTypes   []string `json:"event_types"`
}

// CreateEvents 资产下创建事件
func (s *EventManagerService) CreateEvents(ctx context.Context, req *EventsCreateRequest) error {
	resp, err := s.client.Post(ctx, "/2/event_manager/events/create/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 监测链接管理 ====================

// TrackURL 监测链接
type TrackURL struct {
	TrackURLID   int64  `json:"track_url_id"`
	TrackURLName string `json:"track_url_name"`
	TrackURL     string `json:"track_url"`
	ActionType   string `json:"action_type"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// TrackURLGroup 监测链接组
type TrackURLGroup struct {
	GroupID    int64      `json:"group_id"`
	GroupName  string     `json:"group_name"`
	TrackURLs  []TrackURL `json:"track_urls"`
	Status     int        `json:"status"`
	CreateTime string     `json:"create_time"`
}

// TrackURLCreateRequest 创建监测链接组请求
type TrackURLCreateRequest struct {
	AdvertiserID int64             `json:"advertiser_id"`
	AssetID      int64             `json:"asset_id"`
	GroupName    string            `json:"group_name"`
	TrackURLs    []TrackURLSetting `json:"track_urls"`
}

// TrackURLSetting 监测链接设置
type TrackURLSetting struct {
	ActionType string `json:"action_type"`
	TrackURL   string `json:"track_url"`
}

// CreateTrackURL 事件资产下创建监测链接组
func (s *EventManagerService) CreateTrackURL(ctx context.Context, req *TrackURLCreateRequest) error {
	resp, err := s.client.Post(ctx, "/2/event_manager/track_url/create/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// TrackURLUpdateRequest 更新监测链接组请求
type TrackURLUpdateRequest struct {
	AdvertiserID int64             `json:"advertiser_id"`
	AssetID      int64             `json:"asset_id"`
	GroupID      int64             `json:"group_id"`
	GroupName    string            `json:"group_name,omitempty"`
	TrackURLs    []TrackURLSetting `json:"track_urls,omitempty"`
}

// UpdateTrackURL 事件资产下更新监测链接组
func (s *EventManagerService) UpdateTrackURL(ctx context.Context, req *TrackURLUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/event_manager/track_url/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// TrackURLGetRequest 获取监测链接组请求
type TrackURLGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AssetID      int64 `json:"asset_id"`
}

// TrackURLGetResponse 获取监测链接组响应
type TrackURLGetResponse struct {
	List []TrackURLGroup `json:"list"`
}

// GetTrackURL 获取事件资产下的监测链接组
func (s *EventManagerService) GetTrackURL(ctx context.Context, req *TrackURLGetRequest) (*TrackURLGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_id":      req.AssetID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/track_url/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result TrackURLGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 资产共享 ====================

// ShareGetRequest 查看共享范围请求
type ShareGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AssetID      int64 `json:"asset_id"`
}

// ShareGetResponse 查看共享范围响应
type ShareGetResponse struct {
	SharedAdvertiserIDs []int64 `json:"shared_advertiser_ids"`
}

// GetShare 事件管理资产查看共享范围
func (s *EventManagerService) GetShare(ctx context.Context, req *ShareGetRequest) (*ShareGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_id":      req.AssetID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/share/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ShareGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ShareRequest 资产共享/取消共享请求
type ShareRequest struct {
	AdvertiserID        int64   `json:"advertiser_id"`
	AssetID             int64   `json:"asset_id"`
	TargetAdvertiserIDs []int64 `json:"target_advertiser_ids"`
}

// ShareError 共享错误
type ShareError struct {
	AdvertiserID int64  `json:"advertiser_id"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

// Share 事件管理资产共享
func (s *EventManagerService) Share(ctx context.Context, req *ShareRequest) ([]ShareError, error) {
	resp, err := s.client.Post(ctx, "/2/event_manager/share/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		FailList []ShareError `json:"fail_list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.FailList, nil
}

// ShareCancel 事件管理资产取消共享
func (s *EventManagerService) ShareCancel(ctx context.Context, req *ShareRequest) ([]ShareError, error) {
	resp, err := s.client.Post(ctx, "/2/event_manager/share/cancel/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		FailList []ShareError `json:"fail_list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.FailList, nil
}

// ==================== 优化目标 ====================

// OptimizedGoal 可用优化目标
type OptimizedGoal struct {
	OptimizeGoal     string `json:"optimize_goal"`
	OptimizeGoalName string `json:"optimize_goal_name"`
	EventType        string `json:"event_type"`
	DeepBidType      string `json:"deep_bid_type,omitempty"`
}

// EventConvertOptimizedGoalGetRequest 获取可用优化目标请求
type EventConvertOptimizedGoalGetRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	AssetID      int64  `json:"asset_id"`
	LandingType  string `json:"landing_type,omitempty"`
}

// EventConvertOptimizedGoalGetResponse 获取可用优化目标响应
type EventConvertOptimizedGoalGetResponse struct {
	List []OptimizedGoal `json:"list"`
}

// GetEventConvertOptimizedGoal 获取可用优化目标
func (s *EventManagerService) GetEventConvertOptimizedGoal(ctx context.Context, req *EventConvertOptimizedGoalGetRequest) (*EventConvertOptimizedGoalGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"asset_id":      req.AssetID,
	}
	if req.LandingType != "" {
		params["landing_type"] = req.LandingType
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/event_convert/optimized_goal/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result EventConvertOptimizedGoalGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 体验版优化目标 ====================

// V3OptimizedGoalGetRequest 获取可用优化目标请求（体验版）
type V3OptimizedGoalGetRequest struct {
	AdvertiserID  int64  `json:"advertiser_id"`
	MarketingGoal string `json:"marketing_goal"`
	LandingType   string `json:"landing_type"`
	AssetID       int64  `json:"asset_id,omitempty"`
	OptimizeGoal  string `json:"optimize_goal,omitempty"`
}

// V3OptimizedGoalGetResponse 获取可用优化目标响应（体验版）
type V3OptimizedGoalGetResponse struct {
	List []OptimizedGoal `json:"list"`
}

// GetV3OptimizedGoal 获取可用优化目标（体验版）
func (s *EventManagerService) GetV3OptimizedGoal(ctx context.Context, req *V3OptimizedGoalGetRequest) (*V3OptimizedGoalGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id":  req.AdvertiserID,
		"marketing_goal": req.MarketingGoal,
		"landing_type":   req.LandingType,
	}
	if req.AssetID > 0 {
		params["asset_id"] = req.AssetID
	}
	if req.OptimizeGoal != "" {
		params["optimize_goal"] = req.OptimizeGoal
	}

	resp, err := s.client.Get(ctx, "/v3.0/event_manager/optimized_goal/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result V3OptimizedGoalGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// DeepBidType 深度优化方式
type DeepBidType struct {
	DeepBidType     string `json:"deep_bid_type"`
	DeepBidTypeName string `json:"deep_bid_type_name"`
}

// V3DeepBidTypeGetRequest 获取可用深度优化方式请求（体验版）
type V3DeepBidTypeGetRequest struct {
	AdvertiserID  int64  `json:"advertiser_id"`
	MarketingGoal string `json:"marketing_goal"`
	LandingType   string `json:"landing_type"`
	OptimizeGoal  string `json:"optimize_goal"`
	AssetID       int64  `json:"asset_id,omitempty"`
}

// GetV3DeepBidType 获取可用深度优化方式（体验版）
func (s *EventManagerService) GetV3DeepBidType(ctx context.Context, req *V3DeepBidTypeGetRequest) ([]DeepBidType, error) {
	params := map[string]interface{}{
		"advertiser_id":  req.AdvertiserID,
		"marketing_goal": req.MarketingGoal,
		"landing_type":   req.LandingType,
		"optimize_goal":  req.OptimizeGoal,
	}
	if req.AssetID > 0 {
		params["asset_id"] = req.AssetID
	}

	resp, err := s.client.Get(ctx, "/v3.0/event_manager/deep_bid_type/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []DeepBidType `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 转化回传 ====================

// ConversionRequest 转化回传请求
type ConversionRequest struct {
	EventType  string                 `json:"event_type"`
	Context    map[string]interface{} `json:"context"`
	Timestamp  int64                  `json:"timestamp,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Conversion 转化回传
func (s *EventManagerService) Conversion(ctx context.Context, req *ConversionRequest) (int, error) {
	resp, err := s.client.Post(ctx, "/2/conversion/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return resp.Code, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return 0, nil
}

// ==================== 转化回传鉴权 ====================

// PublicKey 公钥信息
type PublicKey struct {
	KeyID      string `json:"key_id"`
	PublicKey  string `json:"public_key"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// AddPublicKeyRequest 新增公钥请求
type AddPublicKeyRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	PublicKey    string `json:"public_key"`
}

// AddPublicKey 新增公钥
func (s *EventManagerService) AddPublicKey(ctx context.Context, req *AddPublicKeyRequest) (*PublicKey, error) {
	resp, err := s.client.Post(ctx, "/2/event_manager/auth/public_key/add/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result PublicKey
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// DelPublicKeyRequest 删除公钥请求
type DelPublicKeyRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	KeyID        string `json:"key_id"`
}

// DelPublicKey 删除公钥
func (s *EventManagerService) DelPublicKey(ctx context.Context, req *DelPublicKeyRequest) error {
	resp, err := s.client.Post(ctx, "/2/event_manager/auth/public_key/del/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// GetPublicKeyRequest 查询公钥请求
type GetPublicKeyRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	KeyID        string `json:"key_id"`
}

// GetPublicKey 查询公钥
func (s *EventManagerService) GetPublicKey(ctx context.Context, req *GetPublicKeyRequest) (*PublicKey, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"key_id":        req.KeyID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/auth/public_key/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result PublicKey
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// GetAllPublicKeysRequest 查询全部公钥请求
type GetAllPublicKeysRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
}

// GetAllPublicKeys 查询全部公钥
func (s *EventManagerService) GetAllPublicKeys(ctx context.Context, req *GetAllPublicKeysRequest) ([]PublicKey, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	resp, err := s.client.Get(ctx, "/2/event_manager/auth/public_key/get_all/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []PublicKey `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// EnableAuth 开启鉴权
func (s *EventManagerService) EnableAuth(ctx context.Context, advertiserID int64) error {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/event_manager/auth/enable/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// DisableAuth 关闭鉴权
func (s *EventManagerService) DisableAuth(ctx context.Context, advertiserID int64) error {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/event_manager/auth/disable/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}
