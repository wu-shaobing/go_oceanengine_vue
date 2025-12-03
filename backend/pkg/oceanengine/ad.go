package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// AdCreateRequest 创建广告计划请求
type AdCreateRequest struct {
	AdvertiserID         int64       `json:"advertiser_id"`
	CampaignID           int64       `json:"campaign_id"`
	Name                 string      `json:"name"`
	DeliveryRange        string      `json:"delivery_range,omitempty"`         // DEFAULT, UNION
	Budget               float64     `json:"budget,omitempty"`                 // 预算
	BudgetMode           string      `json:"budget_mode,omitempty"`            // BUDGET_MODE_DAY, BUDGET_MODE_TOTAL, BUDGET_MODE_INFINITE
	Bid                  float64     `json:"bid,omitempty"`                    // 出价
	Pricing              string      `json:"pricing"`                          // PRICING_CPC, PRICING_CPM, PRICING_OCPM, PRICING_CPA
	StartTime            string      `json:"start_time,omitempty"`             // 投放开始时间
	EndTime              string      `json:"end_time,omitempty"`               // 投放结束时间
	ScheduleType         string      `json:"schedule_type,omitempty"`          // SCHEDULE_FROM_NOW, SCHEDULE_START_END
	FlowControlMode      string      `json:"flow_control_mode,omitempty"`      // 投放速度 FLOW_CONTROL_MODE_FAST, FLOW_CONTROL_MODE_SMOOTH
	ConvertID            int64       `json:"convert_id,omitempty"`             // 转化ID
	ExternalAction       string      `json:"external_action,omitempty"`        // 转化目标
	DeepExternalAction   string      `json:"deep_external_action,omitempty"`   // 深度转化目标
	DeepBidType          string      `json:"deep_bid_type,omitempty"`          // 深度出价类型
	SmartBidType         string      `json:"smart_bid_type,omitempty"`         // 投放场景 SMART_BID_CUSTOM, SMART_BID_CONSERVATIVE
	AdjustCpa            float64     `json:"adjust_cpa,omitempty"`             // CPA出价
	CpaBid               float64     `json:"cpa_bid,omitempty"`                // 目标转化出价
	RoiGoal              float64     `json:"roi_goal,omitempty"`               // ROI目标
	LubanRoiGoal         float64     `json:"luban_roi_goal,omitempty"`         // 鲁班ROI
	Audience             *AdAudience `json:"audience,omitempty"`               // 定向设置
	AwemeID              int64       `json:"aweme_id,omitempty"`               // 抖音号ID
	CreativeMaterialMode string      `json:"creative_material_mode,omitempty"` // 创意类型
}

// AdAudience 广告定向设置
type AdAudience struct {
	District               string   `json:"district,omitempty"`                 // 地域定向类型 CITY, COUNTY, NONE
	City                   []int64  `json:"city,omitempty"`                     // 城市ID列表
	LocationType           string   `json:"location_type,omitempty"`            // 位置类型
	Gender                 string   `json:"gender,omitempty"`                   // 性别 GENDER_MALE, GENDER_FEMALE, NONE
	Age                    []string `json:"age,omitempty"`                      // 年龄段
	RetargetingTags        []int64  `json:"retargeting_tags,omitempty"`         // 定向人群包
	RetargetingTagsExclude []int64  `json:"retargeting_tags_exclude,omitempty"` // 排除人群包
	InterestCategories     []int64  `json:"interest_categories,omitempty"`      // 兴趣分类
	InterestActionMode     string   `json:"interest_action_mode,omitempty"`     // 行为兴趣
	ActionScene            []string `json:"action_scene,omitempty"`             // 行为场景
	ActionDays             int      `json:"action_days,omitempty"`              // 行为天数
	ActionCategories       []int64  `json:"action_categories,omitempty"`        // 行为类目
	ActionWords            []int64  `json:"action_words,omitempty"`             // 行为关键词
	Platform               []string `json:"platform,omitempty"`                 // 平台 ANDROID, IOS
	AndroidOsv             string   `json:"android_osv,omitempty"`              // 最低安卓版本
	IosOsv                 string   `json:"ios_osv,omitempty"`                  // 最低iOS版本
	Ac                     []string `json:"ac,omitempty"`                       // 网络类型 WIFI, 2G, 3G, 4G
	Carrier                []string `json:"carrier,omitempty"`                  // 运营商
	ActivateType           []string `json:"activate_type,omitempty"`            // 新用户定向
	ArticleCategory        []string `json:"article_category,omitempty"`         // 文章分类
	DeviceBrand            []string `json:"device_brand,omitempty"`             // 设备品牌
	LaunchPrice            []int    `json:"launch_price,omitempty"`             // 设备价格
	AutoExtendEnabled      int      `json:"auto_extend_enabled,omitempty"`      // 智能放量
	AutoExtendTargets      []string `json:"auto_extend_targets,omitempty"`      // 可放开定向
}

// AdCreateResponse 创建广告计划响应
type AdCreateResponse struct {
	AdID int64 `json:"ad_id"`
}

// Create 创建广告计划
func (s *AdService) Create(ctx context.Context, req *AdCreateRequest) (*AdCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/ad/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AdCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AdUpdateRequest 更新广告计划请求
type AdUpdateRequest struct {
	AdvertiserID int64       `json:"advertiser_id"`
	AdID         int64       `json:"ad_id"`
	Name         string      `json:"name,omitempty"`
	Budget       float64     `json:"budget,omitempty"`
	BudgetMode   string      `json:"budget_mode,omitempty"`
	Bid          float64     `json:"bid,omitempty"`
	CpaBid       float64     `json:"cpa_bid,omitempty"`
	StartTime    string      `json:"start_time,omitempty"`
	EndTime      string      `json:"end_time,omitempty"`
	ScheduleTime string      `json:"schedule_time,omitempty"`
	Audience     *AdAudience `json:"audience,omitempty"`
}

// AdUpdateResponse 更新广告计划响应
type AdUpdateResponse struct {
	AdID int64 `json:"ad_id"`
}

// Update 更新广告计划
func (s *AdService) Update(ctx context.Context, req *AdUpdateRequest) (*AdUpdateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/ad/update/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AdUpdateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// AdUpdateBudgetRequest 更新广告预算请求
type AdUpdateBudgetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Data         []struct {
		AdID   int64   `json:"ad_id"`
		Budget float64 `json:"budget"`
	} `json:"data"`
}

// AdUpdateBudgetResponseItem 更新预算响应项
type AdUpdateBudgetResponseItem struct {
	AdID    int64  `json:"ad_id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UpdateBudget 更新广告预算
func (s *AdService) UpdateBudget(ctx context.Context, req *AdUpdateBudgetRequest) ([]AdUpdateBudgetResponseItem, error) {
	resp, err := s.client.Post(ctx, "/2/ad/update/budget/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Errors []AdUpdateBudgetResponseItem `json:"errors"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.Errors, nil
}

// AdUpdateBidRequest 更新广告出价请求
type AdUpdateBidRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Data         []struct {
		AdID   int64   `json:"ad_id"`
		Bid    float64 `json:"bid,omitempty"`
		CpaBid float64 `json:"cpa_bid,omitempty"`
	} `json:"data"`
}

// UpdateBid 更新广告出价
func (s *AdService) UpdateBid(ctx context.Context, req *AdUpdateBidRequest) ([]AdUpdateBudgetResponseItem, error) {
	resp, err := s.client.Post(ctx, "/2/ad/update/bid/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Errors []AdUpdateBudgetResponseItem `json:"errors"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.Errors, nil
}

// AdDetailRequest 广告详情请求
type AdDetailRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// AdDetail 广告详情
type AdDetail struct {
	AdID               int64       `json:"ad_id"`
	CampaignID         int64       `json:"campaign_id"`
	Name               string      `json:"name"`
	DeliveryRange      string      `json:"delivery_range"`
	Budget             float64     `json:"budget"`
	BudgetMode         string      `json:"budget_mode"`
	Bid                float64     `json:"bid"`
	CpaBid             float64     `json:"cpa_bid"`
	Pricing            string      `json:"pricing"`
	Status             string      `json:"status"`
	OptStatus          string      `json:"opt_status"`
	StartTime          string      `json:"start_time"`
	EndTime            string      `json:"end_time"`
	ScheduleType       string      `json:"schedule_type"`
	FlowControlMode    string      `json:"flow_control_mode"`
	ConvertID          int64       `json:"convert_id"`
	ExternalAction     string      `json:"external_action"`
	DeepExternalAction string      `json:"deep_external_action"`
	DeepBidType        string      `json:"deep_bid_type"`
	SmartBidType       string      `json:"smart_bid_type"`
	RoiGoal            float64     `json:"roi_goal"`
	Audience           *AdAudience `json:"audience"`
	AwemeID            int64       `json:"aweme_id"`
	CreateTime         string      `json:"ad_create_time"`
	ModifyTime         string      `json:"ad_modify_time"`
	AdvertiserID       int64       `json:"advertiser_id"`
}

// GetDetail 获取广告详情
func (s *AdService) GetDetail(ctx context.Context, req *AdDetailRequest) ([]AdDetail, error) {
	resp, err := s.client.Post(ctx, "/2/ad/detail/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdDetail `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// AdCostProtectStatusRequest 成本保障状态请求
type AdCostProtectStatusRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// AdCostProtectStatus 成本保障状态
type AdCostProtectStatus struct {
	AdID             int64  `json:"ad_id"`
	ProtectStatus    int    `json:"protect_status"` // 0-不在保障中 1-保障中
	ProtectStartTime string `json:"protect_start_time"`
	ProtectEndTime   string `json:"protect_end_time"`
}

// GetCostProtectStatus 获取成本保障状态
func (s *AdService) GetCostProtectStatus(ctx context.Context, req *AdCostProtectStatusRequest) ([]AdCostProtectStatus, error) {
	resp, err := s.client.Post(ctx, "/2/ad/cost/protect/status/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdCostProtectStatus `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// AdRejectReasonRequest 审核建议请求
type AdRejectReasonRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`
	AdIDs        []int64 `json:"ad_ids"`
}

// AdRejectReason 审核建议
type AdRejectReason struct {
	AdID         int64  `json:"ad_id"`
	RejectReason string `json:"reject_reason"`
}

// GetRejectReason 获取审核建议
func (s *AdService) GetRejectReason(ctx context.Context, req *AdRejectReasonRequest) ([]AdRejectReason, error) {
	resp, err := s.client.Post(ctx, "/2/ad/reject_reason/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AdRejectReason `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}
