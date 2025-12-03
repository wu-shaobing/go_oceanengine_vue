package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// AdvToolsService 高级工具服务
type AdvToolsService struct {
	client *Client
}

// NewAdvToolsService 创建高级工具服务
func NewAdvToolsService(client *Client) *AdvToolsService {
	return &AdvToolsService{client: client}
}

// ==================== RTA策略管理 ====================

// RtaInfo RTA策略信息
type RtaInfo struct {
	RtaID      int64  `json:"rta_id"`
	RtaName    string `json:"rta_name"`
	Status     int    `json:"status"`
	ScopeType  string `json:"scope_type"`
	CreateTime string `json:"create_time"`
	ModifyTime string `json:"modify_time"`
}

// RtaGetInfoRequest 获取RTA策略数据请求
type RtaGetInfoRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	RtaIDs       []int64 `json:"rta_ids,omitempty"`
	Page         int     `json:"page,omitempty"`
	PageSize     int     `json:"page_size,omitempty"`
}

// RtaGetInfoResponse 获取RTA策略数据响应
type RtaGetInfoResponse struct {
	List     []RtaInfo `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetRtaInfo 获取RTA策略数据
func (s *AdvToolsService) GetRtaInfo(ctx context.Context, req *RtaGetInfoRequest) (*RtaGetInfoResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if len(req.RtaIDs) > 0 {
		params["rta_ids"] = req.RtaIDs
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/tools/rta/get_info/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result RtaGetInfoResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// RtaGetRequest 获取可用RTA策略请求
type RtaGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
}

// GetAvailableRta 获取可用的RTA策略
func (s *AdvToolsService) GetAvailableRta(ctx context.Context, req *RtaGetRequest) ([]RtaInfo, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/rta/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []RtaInfo `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// RtaStatusUpdateRequest 批量启停RTA策略请求
type RtaStatusUpdateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	RtaIDs       []int64 `json:"rta_ids"`
	Status       int     `json:"status"` // 1-启用 0-停用
}

// UpdateRtaStatus 批量启停账户下RTA策略
func (s *AdvToolsService) UpdateRtaStatus(ctx context.Context, req *RtaStatusUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/tools/rta/status/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// RtaSetScopeRequest 设置RTA策略生效范围请求
type RtaSetScopeRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	RtaID        int64   `json:"rta_id"`
	ScopeType    string  `json:"scope_type"` // ALL-全部 SPECIFIC-指定
	AdIDs        []int64 `json:"ad_ids,omitempty"`
}

// SetRtaScope 设置账户下RTA策略生效范围
func (s *AdvToolsService) SetRtaScope(ctx context.Context, req *RtaSetScopeRequest) error {
	resp, err := s.client.Post(ctx, "/2/tools/rta/set_scope/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// RtaScope RTA策略绑定信息
type RtaScope struct {
	RtaID     int64   `json:"rta_id"`
	ScopeType string  `json:"scope_type"`
	AdIDs     []int64 `json:"ad_ids"`
}

// RtaScopeGetRequest 获取RTA策略绑定信息请求
type RtaScopeGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	RtaIDs       []int64 `json:"rta_ids,omitempty"`
}

// GetRtaScope 获取RTA策略绑定信息列表
func (s *AdvToolsService) GetRtaScope(ctx context.Context, req *RtaScopeGetRequest) ([]RtaScope, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if len(req.RtaIDs) > 0 {
		params["rta_ids"] = req.RtaIDs
	}

	resp, err := s.client.Get(ctx, "/2/tools/rta/scope/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []RtaScope `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 一键起量管理 ====================

// AdRaiseStatus 起量状态
type AdRaiseStatus struct {
	AdID         int64   `json:"ad_id"`
	RaiseStatus  string  `json:"raise_status"`
	RaiseBudget  float64 `json:"raise_budget"`
	RaiseVersion int     `json:"raise_version"`
}

// AdRaiseSetRequest 启动一键起量请求
type AdRaiseSetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdID         int64   `json:"ad_id"`
	Budget       float64 `json:"budget"`
}

// SetAdRaise 启动一键起量
func (s *AdvToolsService) SetAdRaise(ctx context.Context, req *AdRaiseSetRequest) (string, error) {
	resp, err := s.client.Post(ctx, "/2/tools/ad_raise/set/", req)
	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		RaiseStatus string `json:"raise_status"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return "", fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.RaiseStatus, nil
}

// AdRaiseEstimateRequest 获取起量预估值请求
type AdRaiseEstimateRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdID         int64   `json:"ad_id"`
	Budget       float64 `json:"budget"`
}

// GetAdRaiseEstimate 获取当前起量预估值
func (s *AdvToolsService) GetAdRaiseEstimate(ctx context.Context, req *AdRaiseEstimateRequest) (int64, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_id":         req.AdID,
		"budget":        req.Budget,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_raise/estimate/", params)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		EstimateCount int64 `json:"estimate_count"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.EstimateCount, nil
}

// AdRaiseStatusRequest 获取起量状态请求
type AdRaiseStatusRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// GetAdRaiseStatus 获取当前起量状态
func (s *AdvToolsService) GetAdRaiseStatus(ctx context.Context, req *AdRaiseStatusRequest) (map[int64]string, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_ids":        req.AdIDs,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_raise/status/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result map[int64]string
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result, nil
}

// AdRaiseResult 起量后验数据
type AdRaiseResult struct {
	AdID           int64   `json:"ad_id"`
	Cost           float64 `json:"cost"`
	ShowCount      int64   `json:"show_count"`
	ClickCount     int64   `json:"click_count"`
	ConvertCount   int64   `json:"convert_count"`
	RaiseBudget    float64 `json:"raise_budget"`
	RaiseStartTime string  `json:"raise_start_time"`
	RaiseEndTime   string  `json:"raise_end_time"`
}

// AdRaiseResultRequest 获取起量后验数据请求
type AdRaiseResultRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// GetAdRaiseResult 获取起量的后验数据
func (s *AdvToolsService) GetAdRaiseResult(ctx context.Context, req *AdRaiseResultRequest) ([]AdRaiseResult, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_ids":        req.AdIDs,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_raise/result/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdRaiseResult `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// SuggestBudget 建议起量预算
type SuggestBudget struct {
	AdID          int64   `json:"ad_id"`
	SuggestBudget float64 `json:"suggest_budget"`
}

// SuggestBudgetGetRequest 获取广告建议起量预算请求
type SuggestBudgetGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// GetSuggestBudget 获取广告建议起量预算
func (s *AdvToolsService) GetSuggestBudget(ctx context.Context, req *SuggestBudgetGetRequest) ([]SuggestBudget, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_ids":        req.AdIDs,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_raise/suggest_budget/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []SuggestBudget `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 定向包管理 ====================

// AudiencePackage 定向包
type AudiencePackage struct {
	AudiencePackageID   int64                  `json:"audience_package_id"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	LandingType         string                 `json:"landing_type"`
	DeliveryRange       string                 `json:"delivery_range"`
	RetargetingTagsIncl []int64                `json:"retargeting_tags_include"`
	RetargetingTagsExcl []int64                `json:"retargeting_tags_exclude"`
	Audience            map[string]interface{} `json:"audience"`
	Status              int                    `json:"status"`
	CreateTime          string                 `json:"create_time"`
	ModifyTime          string                 `json:"modify_time"`
}

// AudiencePackageGetRequest 获取定向包请求
type AudiencePackageGetRequest struct {
	AdvertiserID       int64   `json:"advertiser_id"`
	AudiencePackageIDs []int64 `json:"audience_package_ids,omitempty"`
	LandingType        string  `json:"landing_type,omitempty"`
	Page               int     `json:"page,omitempty"`
	PageSize           int     `json:"page_size,omitempty"`
}

// AudiencePackageGetResponse 获取定向包响应
type AudiencePackageGetResponse struct {
	List     []AudiencePackage `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetAudiencePackage 获取定向包
func (s *AdvToolsService) GetAudiencePackage(ctx context.Context, req *AudiencePackageGetRequest) (*AudiencePackageGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if len(req.AudiencePackageIDs) > 0 {
		params["audience_package_ids"] = req.AudiencePackageIDs
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

	resp, err := s.client.Get(ctx, "/2/audience_package/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AudiencePackageGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AudiencePackageCreateRequest 创建定向包请求
type AudiencePackageCreateRequest struct {
	AdvertiserID        int64                  `json:"advertiser_id"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description,omitempty"`
	LandingType         string                 `json:"landing_type"`
	DeliveryRange       string                 `json:"delivery_range,omitempty"`
	RetargetingTagsIncl []int64                `json:"retargeting_tags_include,omitempty"`
	RetargetingTagsExcl []int64                `json:"retargeting_tags_exclude,omitempty"`
	Audience            map[string]interface{} `json:"audience,omitempty"`
}

// CreateAudiencePackage 创建定向包
func (s *AdvToolsService) CreateAudiencePackage(ctx context.Context, req *AudiencePackageCreateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/audience_package/create/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AudiencePackageID int64 `json:"audience_package_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AudiencePackageID, nil
}

// AudiencePackageUpdateRequest 更新定向包请求
type AudiencePackageUpdateRequest struct {
	AdvertiserID        int64                  `json:"advertiser_id"`
	AudiencePackageID   int64                  `json:"audience_package_id"`
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	RetargetingTagsIncl []int64                `json:"retargeting_tags_include,omitempty"`
	RetargetingTagsExcl []int64                `json:"retargeting_tags_exclude,omitempty"`
	Audience            map[string]interface{} `json:"audience,omitempty"`
}

// UpdateAudiencePackage 更新定向包
func (s *AdvToolsService) UpdateAudiencePackage(ctx context.Context, req *AudiencePackageUpdateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/audience_package/update/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AudiencePackageID int64 `json:"audience_package_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AudiencePackageID, nil
}

// AudiencePackageDeleteRequest 删除定向包请求
type AudiencePackageDeleteRequest struct {
	AdvertiserID      int64 `json:"advertiser_id"`
	AudiencePackageID int64 `json:"audience_package_id"`
}

// DeleteAudiencePackage 删除定向包
func (s *AdvToolsService) DeleteAudiencePackage(ctx context.Context, req *AudiencePackageDeleteRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/audience_package/delete/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AudiencePackageID int64 `json:"audience_package_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AudiencePackageID, nil
}

// AdBindRequest 计划绑定/解绑定向包请求
type AdBindRequest struct {
	AdvertiserID      int64   `json:"advertiser_id"`
	AudiencePackageID int64   `json:"audience_package_id"`
	AdIDs             []int64 `json:"ad_ids"`
}

// BindAudiencePackage 计划绑定定向包
func (s *AdvToolsService) BindAudiencePackage(ctx context.Context, req *AdBindRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/audience_package/ad/bind/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AudiencePackageID int64 `json:"audience_package_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AudiencePackageID, nil
}

// UnbindAudiencePackage 定向包解绑
func (s *AdvToolsService) UnbindAudiencePackage(ctx context.Context, req *AdBindRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/audience_package/ad/unbind/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		AudiencePackageID int64 `json:"audience_package_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.AudiencePackageID, nil
}

// ==================== 原生锚点管理 ====================

// NativeAnchor 原生锚点
type NativeAnchor struct {
	AnchorID   int64  `json:"anchor_id"`
	AnchorName string `json:"anchor_name"`
	AnchorType string `json:"anchor_type"`
	Status     int    `json:"status"`
	PreviewURL string `json:"preview_url"`
	CreateTime string `json:"create_time"`
	ModifyTime string `json:"modify_time"`
}

// NativeAnchorGetRequest 获取原生锚点请求
type NativeAnchorGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AnchorIDs    []int64 `json:"anchor_ids,omitempty"`
	AnchorType   string  `json:"anchor_type,omitempty"`
	Page         int     `json:"page,omitempty"`
	PageSize     int     `json:"page_size,omitempty"`
}

// NativeAnchorGetResponse 获取原生锚点响应
type NativeAnchorGetResponse struct {
	List     []NativeAnchor `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetNativeAnchor 获取账户下原生锚点
func (s *AdvToolsService) GetNativeAnchor(ctx context.Context, req *NativeAnchorGetRequest) (*NativeAnchorGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if len(req.AnchorIDs) > 0 {
		params["anchor_ids"] = req.AnchorIDs
	}
	if req.AnchorType != "" {
		params["anchor_type"] = req.AnchorType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/tools/native_anchor/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result NativeAnchorGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// NativeAnchorDetailRequest 获取原生锚点详情请求
type NativeAnchorDetailRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AnchorID     int64 `json:"anchor_id"`
}

// GetNativeAnchorDetail 获取原生锚点详情
func (s *AdvToolsService) GetNativeAnchorDetail(ctx context.Context, req *NativeAnchorDetailRequest) (*NativeAnchor, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"anchor_id":     req.AnchorID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/native_anchor/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result NativeAnchor
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// NativeAnchorCreateRequest 创建原生锚点请求
type NativeAnchorCreateRequest struct {
	AdvertiserID int64                  `json:"advertiser_id"`
	AnchorName   string                 `json:"anchor_name"`
	AnchorType   string                 `json:"anchor_type"`
	AnchorConfig map[string]interface{} `json:"anchor_config"`
}

// NativeAnchorCreateResponse 创建原生锚点响应
type NativeAnchorCreateResponse struct {
	AnchorID int64 `json:"anchor_id"`
}

// CreateNativeAnchor 原生锚点创建
func (s *AdvToolsService) CreateNativeAnchor(ctx context.Context, req *NativeAnchorCreateRequest) (*NativeAnchorCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/tools/native_anchor/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result NativeAnchorCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// NativeAnchorUpdateRequest 更新原生锚点请求
type NativeAnchorUpdateRequest struct {
	AdvertiserID int64                  `json:"advertiser_id"`
	AnchorID     int64                  `json:"anchor_id"`
	AnchorName   string                 `json:"anchor_name,omitempty"`
	AnchorConfig map[string]interface{} `json:"anchor_config,omitempty"`
}

// NativeAnchorUpdateResponse 更新原生锚点响应
type NativeAnchorUpdateResponse struct {
	AnchorID int64 `json:"anchor_id"`
}

// UpdateNativeAnchor 更新原生锚点
func (s *AdvToolsService) UpdateNativeAnchor(ctx context.Context, req *NativeAnchorUpdateRequest) (*NativeAnchorUpdateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/tools/native_anchor/update/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result NativeAnchorUpdateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// NativeAnchorDeleteRequest 删除原生锚点请求
type NativeAnchorDeleteRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AnchorID     int64 `json:"anchor_id"`
}

// DeleteNativeAnchor 删除原生锚点
func (s *AdvToolsService) DeleteNativeAnchor(ctx context.Context, req *NativeAnchorDeleteRequest) error {
	resp, err := s.client.Post(ctx, "/2/tools/native_anchor/delete/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 计划诊断工具 ====================

// DiagnosisSuggestion 诊断建议
type DiagnosisSuggestion struct {
	SuggestionID   int64  `json:"suggestion_id"`
	SuggestionType string `json:"suggestion_type"`
	SuggestionDesc string `json:"suggestion_desc"`
	Dimension      string `json:"dimension"`
	Priority       int    `json:"priority"`
	EstimateEffect string `json:"estimate_effect"`
}

// DiagnosisSuggestionGetRequest 获取计划诊断建议请求
type DiagnosisSuggestionGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	AdID         int64 `json:"ad_id"`
}

// DiagnosisSuggestionGetResponse 获取计划诊断建议响应
type DiagnosisSuggestionGetResponse struct {
	List []DiagnosisSuggestion `json:"list"`
}

// GetDiagnosisSuggestion 获取计划诊断建议
func (s *AdvToolsService) GetDiagnosisSuggestion(ctx context.Context, req *DiagnosisSuggestionGetRequest) (*DiagnosisSuggestionGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_id":         req.AdID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/diagnosis/suggestion/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result DiagnosisSuggestionGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// DiagnosisSuggestionAcceptRequest 采纳计划诊断建议请求
type DiagnosisSuggestionAcceptRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdID         int64   `json:"ad_id"`
	SuggestionID []int64 `json:"suggestion_ids"`
}

// DiagnosisSuggestionAcceptResponse 采纳计划诊断建议响应
type DiagnosisSuggestionAcceptResponse struct {
	SuccessIDs []int64 `json:"success_ids"`
	FailIDs    []int64 `json:"fail_ids"`
}

// AcceptDiagnosisSuggestion 采纳计划诊断建议
func (s *AdvToolsService) AcceptDiagnosisSuggestion(ctx context.Context, req *DiagnosisSuggestionAcceptRequest) (*DiagnosisSuggestionAcceptResponse, error) {
	resp, err := s.client.Post(ctx, "/2/tools/diagnosis/suggestion/accept/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result DiagnosisSuggestionAcceptResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 其他工具 ====================

// QuotaInfo 配额信息
type QuotaInfo struct {
	CurrentCount int64 `json:"current_count"`
	MaxCount     int64 `json:"max_count"`
}

// QuotaGetRequest 查询在投计划配额请求
type QuotaGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
}

// GetQuota 查询在投计划配额
func (s *AdvToolsService) GetQuota(ctx context.Context, req *QuotaGetRequest) (*QuotaInfo, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	resp, err := s.client.Get(ctx, "/2/tools/quota/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result QuotaInfo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AdQuality 广告质量度
type AdQuality struct {
	AdID         int64  `json:"ad_id"`
	QualityScore int    `json:"quality_score"`
	QualityLevel string `json:"quality_level"`
}

// AdQualityGetRequest 查询广告质量度请求
type AdQualityGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// GetAdQuality 查询广告质量度
func (s *AdvToolsService) GetAdQuality(ctx context.Context, req *AdQualityGetRequest) ([]AdQuality, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_ids":        req.AdIDs,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_quality/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdQuality `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// AdStatExtraInfo 广告计划学习期状态
type AdStatExtraInfo struct {
	AdID              int64  `json:"ad_id"`
	LearningPhase     string `json:"learning_phase"`
	LearningPhaseDesc string `json:"learning_phase_desc"`
}

// AdStatExtraInfoGetRequest 查询广告计划学习期状态请求
type AdStatExtraInfoGetRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// GetAdStatExtraInfo 查询广告计划学习期状态
func (s *AdvToolsService) GetAdStatExtraInfo(ctx context.Context, req *AdStatExtraInfoGetRequest) ([]AdStatExtraInfo, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"ad_ids":        req.AdIDs,
	}

	resp, err := s.client.Get(ctx, "/2/tools/ad_stat_extra_info/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdStatExtraInfo `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}
