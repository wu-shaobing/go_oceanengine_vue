package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// QianchuanClient 千川API客户端
type QianchuanClient struct {
	client *Client
}

// NewQianchuanClient 创建千川客户端
func (c *Client) Qianchuan() *QianchuanClient {
	return &QianchuanClient{client: c}
}

// ==================== 账户管理 ====================

// AccountInfo 千川账户信息
type QianchuanAccountInfo struct {
	AdvertiserID   uint64 `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	AccountType    int    `json:"account_type"` // 1-自运营 2-代运营
	Balance        int64  `json:"balance"`      // 账户余额(分)
	ValidBalance   int64  `json:"valid_balance"`
}

// GetAccountInfo 获取千川账户信息
func (q *QianchuanClient) GetAccountInfo(ctx context.Context, accessToken string, advertiserID uint64) (*QianchuanAccountInfo, error) {
	path := "/v1.0/qianchuan/advertiser/info/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data QianchuanAccountInfo `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetBalance 获取账户余额
func (q *QianchuanClient) GetBalance(ctx context.Context, accessToken string, advertiserID uint64) (int64, error) {
	path := "/v1.0/qianchuan/finance/wallet/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data struct {
			Balance      int64 `json:"balance"`
			ValidBalance int64 `json:"valid_balance"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.Balance, nil
}

// FinanceDetail 财务明细
type FinanceDetail struct {
	TransactionID   string  `json:"transaction_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Balance         float64 `json:"balance"`
	CreateTime      string  `json:"create_time"`
	Remark          string  `json:"remark"`
}

// GetFinanceDetail 获取财务明细
func (q *QianchuanClient) GetFinanceDetail(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]FinanceDetail, int, error) {
	path := "/v1.0/qianchuan/finance/flow/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if startDate != "" {
		params["start_date"] = startDate
	}
	if endDate != "" {
		params["end_date"] = endDate
	}

	var result struct {
		Data struct {
			List     []FinanceDetail `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 店铺管理 ====================

// Shop 店铺信息
type Shop struct {
	ShopID   uint64 `json:"shop_id"`
	ShopName string `json:"shop_name"`
	Platform int    `json:"platform"` // 1-抖店 2-小店
	Status   int    `json:"status"`
}

// GetShopList 获取店铺列表
func (q *QianchuanClient) GetShopList(ctx context.Context, accessToken string, advertiserID uint64) ([]Shop, error) {
	path := "/v1.0/qianchuan/shop/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data struct {
			List []Shop `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 广告投放 ====================

// QianchuanAd 千川广告计划
type QianchuanAd struct {
	AdID          uint64  `json:"ad_id"`
	AdName        string  `json:"ad_name"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	CampaignID    uint64  `json:"campaign_id"`
	MarketingGoal string  `json:"marketing_goal"`
	Status        string  `json:"status"`
	OptStatus     string  `json:"opt_status"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	DeliveryRange string  `json:"delivery_range"`
	CreateTime    string  `json:"create_time"`
	ModifyTime    string  `json:"modify_time"`
}

// AdListRequest 广告列表请求
type QianchuanAdListRequest struct {
	AdvertiserID uint64    `json:"advertiser_id"`
	Filtering    *AdFilter `json:"filtering,omitempty"`
	Page         int       `json:"page,omitempty"`
	PageSize     int       `json:"page_size,omitempty"`
}

type AdFilter struct {
	AdIDs         []uint64 `json:"ad_ids,omitempty"`
	AdName        string   `json:"ad_name,omitempty"`
	Status        string   `json:"status,omitempty"`
	MarketingGoal string   `json:"marketing_goal,omitempty"`
}

// GetAdList 获取广告计划列表
func (q *QianchuanClient) GetAdList(ctx context.Context, accessToken string, req *QianchuanAdListRequest) ([]QianchuanAd, int, error) {
	path := "/v1.0/qianchuan/ad/get/"
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.Filtering != nil {
		params["filtering"] = req.Filtering
	}

	var result struct {
		Data struct {
			List     []QianchuanAd `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
				TotalPage   int `json:"total_page"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreateAd 创建广告计划
func (q *QianchuanClient) CreateAd(ctx context.Context, accessToken string, advertiserID uint64, adData map[string]interface{}) (uint64, error) {
	path := "/v1.0/qianchuan/ad/create/"
	adData["advertiser_id"] = advertiserID

	var result struct {
		Data struct {
			AdID uint64 `json:"ad_id"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, adData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.AdID, nil
}

// UpdateAdStatus 更新广告状态
func (q *QianchuanClient) UpdateAdStatus(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, optStatus string) error {
	path := "/v1.0/qianchuan/ad/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"opt_status":    optStatus, // enable, disable
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 数据报表 ====================

// QianchuanReport 千川报表数据
type QianchuanReport struct {
	StatDatetime string  `json:"stat_datetime"`
	Cost         float64 `json:"cost"`
	ShowCnt      int64   `json:"show_cnt"`
	ClickCnt     int64   `json:"click_cnt"`
	ConvertCnt   int64   `json:"convert_cnt"`
	PayOrderCnt  int64   `json:"pay_order_cnt"`
	PayOrderAmt  float64 `json:"pay_order_amount"`
	Ctr          float64 `json:"ctr"`
	Cpm          float64 `json:"cpm"`
	CpaConvert   float64 `json:"cpa_convert"`
	Roi          float64 `json:"roi"`
}

// GetAdReport 获取广告报表
func (q *QianchuanClient) GetAdReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, adIDs []uint64) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/ad/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(adIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"ad_ids": adIDs,
		}
	}

	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetLiveReport 获取直播间报表
func (q *QianchuanClient) GetLiveReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/live_room/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}

	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 商品管理 ====================

// Product 商品信息
type Product struct {
	ProductID     uint64  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	MarketPrice   float64 `json:"market_price"`
	DiscountPrice float64 `json:"discount_price"`
	Img           string  `json:"img"`
	Status        int     `json:"status"`
}

// GetProductList 获取商品列表
func (q *QianchuanClient) GetProductList(ctx context.Context, accessToken string, advertiserID uint64, awemeID uint64, page, pageSize int) ([]Product, int, error) {
	path := "/v1.0/qianchuan/product/available/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []Product `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 素材管理 ====================

// VideoUploadResponse 视频上传响应
type VideoUploadResponse struct {
	VideoID    string `json:"video_id"`
	MaterialID uint64 `json:"material_id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Duration   int    `json:"duration"`
	Size       int64  `json:"size"`
}

// ImageUploadResponse 图片上传响应
type ImageUploadResponse struct {
	ImageID    string `json:"image_id"`
	MaterialID uint64 `json:"material_id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Size       int64  `json:"size"`
	URL        string `json:"url"`
}

// UploadVideo 上传视频素材
func (q *QianchuanClient) UploadVideo(ctx context.Context, accessToken string, advertiserID uint64, filePath string) (*VideoUploadResponse, error) {
	path := "/v1.0/qianchuan/file/video/ad/"
	extraFields := map[string]string{
		"advertiser_id": fmt.Sprintf("%d", advertiserID),
	}

	resp, err := q.client.UploadFile(ctx, accessToken, path, "video_file", filePath, extraFields)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data VideoUploadResponse `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}
	return &result.Data, nil
}

// UploadVideoFromReader 从Reader上传视频素材
func (q *QianchuanClient) UploadVideoFromReader(ctx context.Context, accessToken string, advertiserID uint64, fileName string, reader io.Reader) (*VideoUploadResponse, error) {
	path := "/v1.0/qianchuan/file/video/ad/"
	extraFields := map[string]string{
		"advertiser_id": fmt.Sprintf("%d", advertiserID),
	}

	resp, err := q.client.UploadFileFromReader(ctx, accessToken, path, "video_file", fileName, reader, extraFields)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data VideoUploadResponse `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}
	return &result.Data, nil
}

// UploadImage 上传图片素材
func (q *QianchuanClient) UploadImage(ctx context.Context, accessToken string, advertiserID uint64, filePath string) (*ImageUploadResponse, error) {
	path := "/v1.0/qianchuan/file/image/ad/"
	extraFields := map[string]string{
		"advertiser_id": fmt.Sprintf("%d", advertiserID),
	}

	resp, err := q.client.UploadFile(ctx, accessToken, path, "image_file", filePath, extraFields)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data ImageUploadResponse `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}
	return &result.Data, nil
}

// UploadImageFromReader 从Reader上传图片素材
func (q *QianchuanClient) UploadImageFromReader(ctx context.Context, accessToken string, advertiserID uint64, fileName string, reader io.Reader) (*ImageUploadResponse, error) {
	path := "/v1.0/qianchuan/file/image/ad/"
	extraFields := map[string]string{
		"advertiser_id": fmt.Sprintf("%d", advertiserID),
	}

	resp, err := q.client.UploadFileFromReader(ctx, accessToken, path, "image_file", fileName, reader, extraFields)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data ImageUploadResponse `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}
	return &result.Data, nil
}

// ==================== 广告组管理 ====================

// QianchuanCampaign 千川广告组
type QianchuanCampaign struct {
	CampaignID    uint64  `json:"campaign_id"`
	CampaignName  string  `json:"campaign_name"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	Budget        float64 `json:"budget"`
	BudgetMode    string  `json:"budget_mode"`
	MarketingGoal string  `json:"marketing_goal"`
	Status        string  `json:"status"`
	CreateTime    string  `json:"create_time"`
	ModifyTime    string  `json:"modify_time"`
}

// CreateCampaign 创建广告组
func (q *QianchuanClient) CreateCampaign(ctx context.Context, accessToken string, advertiserID uint64, name string, budget float64, budgetMode, marketingGoal string) (uint64, error) {
	path := "/v1.0/qianchuan/campaign/create/"
	data := map[string]interface{}{
		"advertiser_id":  advertiserID,
		"campaign_name":  name,
		"budget":         budget,
		"budget_mode":    budgetMode,
		"marketing_goal": marketingGoal,
	}
	var result struct {
		Data struct {
			CampaignID uint64 `json:"campaign_id"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.CampaignID, nil
}

// UpdateCampaign 更新广告组
func (q *QianchuanClient) UpdateCampaign(ctx context.Context, accessToken string, advertiserID, campaignID uint64, updateData map[string]interface{}) error {
	path := "/v1.0/qianchuan/campaign/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["campaign_id"] = campaignID
	return q.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// UpdateCampaignStatus 更新广告组状态
func (q *QianchuanClient) UpdateCampaignStatus(ctx context.Context, accessToken string, advertiserID uint64, campaignIDs []uint64, optStatus string) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/campaign/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"campaign_ids":  campaignIDs,
		"opt_status":    optStatus,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_campaign_ids"`
			FailIDs    []uint64 `json:"fail_campaign_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetCampaignList 获取广告组列表
func (q *QianchuanClient) GetCampaignList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int, filtering map[string]interface{}) ([]QianchuanCampaign, int, error) {
	path := "/v1.0/qianchuan/campaign/list/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if filtering != nil {
		params["filtering"] = filtering
	}
	var result struct {
		Data struct {
			List     []QianchuanCampaign `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 广告计划扩展 ====================

// UpdateAd 更新广告计划
func (q *QianchuanClient) UpdateAd(ctx context.Context, accessToken string, advertiserID, adID uint64, updateData map[string]interface{}) error {
	path := "/v1.0/qianchuan/ad/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["ad_id"] = adID
	return q.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// GetAdDetail 获取广告计划详情
func (q *QianchuanClient) GetAdDetail(ctx context.Context, accessToken string, advertiserID, adID uint64) (*QianchuanAd, error) {
	path := "/v1.0/qianchuan/ad/detail/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data QianchuanAd `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateAdBid 更新广告出价
func (q *QianchuanClient) UpdateAdBid(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, bid float64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/bid/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"bid":           bid,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateAdBudget 更新广告预算
func (q *QianchuanClient) UpdateAdBudget(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, budget float64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/budget/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"budget":        budget,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateAdRoiGoal 更新ROI目标
func (q *QianchuanClient) UpdateAdRoiGoal(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, roiGoal float64) error {
	path := "/v1.0/qianchuan/ad/roi_goal/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"roi_goal":      roiGoal,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetAdRejectReason 获取计划审核建议
func (q *QianchuanClient) GetAdRejectReason(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64) ([]map[string]interface{}, error) {
	path := "/v1.0/qianchuan/ad/reject_reason/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 广告创意管理 ====================

// QianchuanCreative 千川创意
type QianchuanCreative struct {
	CreativeID   uint64 `json:"creative_id"`
	AdID         uint64 `json:"ad_id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	CreativeName string `json:"creative_name"`
	Status       string `json:"status"`
	OptStatus    string `json:"opt_status"`
	CreativeType string `json:"creative_type"`
	ImageMode    string `json:"image_mode"`
	CreateTime   string `json:"create_time"`
	ModifyTime   string `json:"modify_time"`
}

// GetCreativeList 获取创意列表
func (q *QianchuanClient) GetCreativeList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int, filtering map[string]interface{}) ([]QianchuanCreative, int, error) {
	path := "/v1.0/qianchuan/creative/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if filtering != nil {
		params["filtering"] = filtering
	}
	var result struct {
		Data struct {
			List     []QianchuanCreative `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// UpdateCreativeStatus 更新创意状态
func (q *QianchuanClient) UpdateCreativeStatus(ctx context.Context, accessToken string, advertiserID uint64, creativeIDs []uint64, optStatus string) error {
	path := "/v1.0/qianchuan/creative/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"creative_ids":  creativeIDs,
		"opt_status":    optStatus,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetCreativeRejectReason 获取创意审核建议
func (q *QianchuanClient) GetCreativeRejectReason(ctx context.Context, accessToken string, advertiserID uint64, creativeIDs []uint64) ([]map[string]interface{}, error) {
	path := "/v1.0/qianchuan/creative/reject_reason/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"creative_ids":  creativeIDs,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 随心推投放 ====================

// AwemeOrder 随心推订单
type AwemeOrder struct {
	OrderID       uint64  `json:"order_id"`
	AwemeID       uint64  `json:"aweme_id"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	Budget        float64 `json:"budget"`
	DeliveryRange string  `json:"delivery_range"`
	MarketingGoal string  `json:"marketing_goal"`
	Status        string  `json:"status"`
	CreateTime    string  `json:"create_time"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
}

// CreateAwemeOrder 创建随心推订单
func (q *QianchuanClient) CreateAwemeOrder(ctx context.Context, accessToken string, advertiserID uint64, orderData map[string]interface{}) (*AwemeOrder, error) {
	path := "/v1.0/qianchuan/aweme/order/create/"
	orderData["advertiser_id"] = advertiserID
	var result struct {
		Data AwemeOrder `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, orderData, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// TerminateAwemeOrder 终止随心推订单
func (q *QianchuanClient) TerminateAwemeOrder(ctx context.Context, accessToken string, advertiserID uint64, orderIDs []uint64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/aweme/order/terminate/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_ids":     orderIDs,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_order_ids"`
			FailIDs    []uint64 `json:"fail_order_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetAwemeOrderList 获取随心推订单列表
func (q *QianchuanClient) GetAwemeOrderList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int, filtering map[string]interface{}) ([]AwemeOrder, int, error) {
	path := "/v1.0/qianchuan/aweme/order/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if filtering != nil {
		params["filtering"] = filtering
	}
	var result struct {
		Data struct {
			List     []AwemeOrder `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetAwemeOrderDetail 获取随心推订单详情
func (q *QianchuanClient) GetAwemeOrderDetail(ctx context.Context, accessToken string, advertiserID, orderID uint64) (*AwemeOrder, error) {
	path := "/v1.0/qianchuan/aweme/order/detail/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_id":      orderID,
	}
	var result struct {
		Data AwemeOrder `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetAwemeVideoList 获取随心推可投视频列表
func (q *QianchuanClient) GetAwemeVideoList(ctx context.Context, accessToken string, advertiserID, awemeID uint64, page, pageSize int) ([]map[string]interface{}, int, error) {
	path := "/v1.0/qianchuan/aweme/video/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []map[string]interface{} `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 全域推广 ====================

// UniPromotion 全域推广计划
type UniPromotion struct {
	AdID          uint64  `json:"ad_id"`
	AdName        string  `json:"ad_name"`
	AdvertiserID  uint64  `json:"advertiser_id"`
	AwemeID       uint64  `json:"aweme_id"`
	MarketingGoal string  `json:"marketing_goal"`
	Status        string  `json:"status"`
	OptStatus     string  `json:"opt_status"`
	Budget        float64 `json:"budget"`
	RoiGoal       float64 `json:"roi_goal"`
	CreateTime    string  `json:"create_time"`
	ModifyTime    string  `json:"modify_time"`
}

// CreateUniPromotion 创建全域推广计划
func (q *QianchuanClient) CreateUniPromotion(ctx context.Context, accessToken string, advertiserID uint64, promotionData map[string]interface{}) (uint64, error) {
	path := "/v1.0/qianchuan/uni_promotion/create/"
	promotionData["advertiser_id"] = advertiserID
	var result struct {
		Data struct {
			AdID uint64 `json:"ad_id"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, promotionData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.AdID, nil
}

// UpdateUniPromotion 更新全域推广计划
func (q *QianchuanClient) UpdateUniPromotion(ctx context.Context, accessToken string, advertiserID, adID uint64, updateData map[string]interface{}) error {
	path := "/v1.0/qianchuan/uni_promotion/update/"
	updateData["advertiser_id"] = advertiserID
	updateData["ad_id"] = adID
	return q.client.PostWithToken(ctx, accessToken, path, updateData, nil)
}

// UpdateUniPromotionStatus 更新全域推广状态
func (q *QianchuanClient) UpdateUniPromotionStatus(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, optStatus string) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/uni_promotion/status/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"opt_status":    optStatus,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetUniPromotionList 获取全域推广列表
func (q *QianchuanClient) GetUniPromotionList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int, filtering map[string]interface{}) ([]UniPromotion, int, error) {
	path := "/v1.0/qianchuan/uni_promotion/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if filtering != nil {
		params["filtering"] = filtering
	}
	var result struct {
		Data struct {
			List     []UniPromotion `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetUniPromotionDetail 获取全域推广详情
func (q *QianchuanClient) GetUniPromotionDetail(ctx context.Context, accessToken string, advertiserID, adID uint64) (*UniPromotion, error) {
	path := "/v1.0/qianchuan/uni_promotion/detail/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data UniPromotion `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 关键词管理 ====================

// Keyword 搜索关键词
type Keyword struct {
	Word      string  `json:"word"`
	MatchType string  `json:"match_type"`
	Bid       float64 `json:"bid"`
	Status    string  `json:"status"`
}

// GetKeywords 获取计划关键词
func (q *QianchuanClient) GetKeywords(ctx context.Context, accessToken string, advertiserID, adID uint64) ([]Keyword, error) {
	path := "/v1.0/qianchuan/ad/keywords/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data struct {
			Keywords []Keyword `json:"keywords"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.Keywords, nil
}

// UpdateKeywords 更新关键词
func (q *QianchuanClient) UpdateKeywords(ctx context.Context, accessToken string, advertiserID, adID uint64, keywords []Keyword) error {
	path := "/v1.0/qianchuan/ad/keywords/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"keywords":      keywords,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetRecommendKeywords 获取推荐关键词
func (q *QianchuanClient) GetRecommendKeywords(ctx context.Context, accessToken string, advertiserID, adID uint64) ([]Keyword, error) {
	path := "/v1.0/qianchuan/ad/recommend_keywords/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data struct {
			Keywords []Keyword `json:"keywords"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.Keywords, nil
}

// GetAdvertiserReport 获取广告主维度报表
func (q *QianchuanClient) GetAdvertiserReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/advertiser/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== DMP人群管理 ====================

// Audience 人群包
type QianchuanAudience struct {
	AudienceID   uint64 `json:"audience_id"`
	Name         string `json:"name"`
	AudienceType int    `json:"audience_type"`
	CoverNum     int64  `json:"cover_num"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// GetAudienceList 获取人群包列表
func (q *QianchuanClient) GetAudienceList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]QianchuanAudience, int, error) {
	path := "/v1.0/qianchuan/dmp/audiences/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []QianchuanAudience `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// PushAudience 推送人群包
func (q *QianchuanClient) PushAudience(ctx context.Context, accessToken string, advertiserID uint64, audienceID uint64, targetAdvertiserIDs []uint64) error {
	path := "/v1.0/qianchuan/dmp/audience/push/"
	data := map[string]interface{}{
		"advertiser_id":         advertiserID,
		"audience_id":           audienceID,
		"target_advertiser_ids": targetAdvertiserIDs,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteAudience 删除人群包
func (q *QianchuanClient) DeleteAudience(ctx context.Context, accessToken string, advertiserID uint64, audienceID uint64) error {
	path := "/v1.0/qianchuan/dmp/audience/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"audience_id":   audienceID,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 扩展报表 ====================

// GetCreativeReport 获取创意报表
func (q *QianchuanClient) GetCreativeReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, creativeIDs []uint64) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/creative/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(creativeIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"creative_ids": creativeIDs,
		}
	}
	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetMaterialReport 获取素材报表
func (q *QianchuanClient) GetMaterialReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/material/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetUniPromotionReport 获取全域推广报表
func (q *QianchuanClient) GetUniPromotionReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, adIDs []uint64) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/uni_promotion/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(adIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"ad_ids": adIDs,
		}
	}
	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetAwemeOrderReport 获取随心推订单报表
func (q *QianchuanClient) GetAwemeOrderReport(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, orderIDs []uint64) ([]QianchuanReport, error) {
	path := "/v1.0/qianchuan/report/order/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	if len(orderIDs) > 0 {
		params["filtering"] = map[string]interface{}{
			"order_ids": orderIDs,
		}
	}
	var result struct {
		Data struct {
			List []QianchuanReport `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 账户预算管理 ====================

// GetAccountBudget 获取账户日预算
func (q *QianchuanClient) GetAccountBudget(ctx context.Context, accessToken string, advertiserID uint64) (float64, error) {
	path := "/v1.0/qianchuan/advertiser/budget/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			Budget float64 `json:"budget"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.Budget, nil
}

// UpdateAccountBudget 更新账户日预算
func (q *QianchuanClient) UpdateAccountBudget(ctx context.Context, accessToken string, advertiserID uint64, budget float64) error {
	path := "/v1.0/qianchuan/advertiser/budget/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"budget":        budget,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 工具接口 ====================

// GetIndustryList 获取行业列表
func (q *QianchuanClient) GetIndustryList(ctx context.Context, accessToken string) ([]map[string]interface{}, error) {
	path := "/v1.0/qianchuan/tools/industry/get/"
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// EstimateAudience 获取定向受众预估
func (q *QianchuanClient) EstimateAudience(ctx context.Context, accessToken string, advertiserID uint64, targeting map[string]interface{}) (int64, error) {
	path := "/v1.0/qianchuan/tools/estimate_audience/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"targeting":     targeting,
	}
	var result struct {
		Data struct {
			UserCount int64 `json:"user_count"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.UserCount, nil
}

// GetInterestCategories 获取兴趣类目
func (q *QianchuanClient) GetInterestCategories(ctx context.Context, accessToken string, advertiserID uint64) ([]map[string]interface{}, error) {
	path := "/v1.0/qianchuan/tools/interest_action/interest/category/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetActionCategories 获取行为类目
func (q *QianchuanClient) GetActionCategories(ctx context.Context, accessToken string, advertiserID uint64) ([]map[string]interface{}, error) {
	path := "/v1.0/qianchuan/tools/interest_action/action/category/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 抵音号管理 ====================

// QianchuanAweme 千川抵音号
type QianchuanAweme struct {
	AwemeID       uint64 `json:"aweme_id"`
	AwemeName     string `json:"aweme_name"`
	AwemeAvatar   string `json:"aweme_avatar"`
	FollowerCount int64  `json:"follower_count"`
	AuthStatus    int    `json:"auth_status"`
	AuthType      string `json:"auth_type"`
	BindStatus    int    `json:"bind_status"`
}

// GetAuthorizedAwemeList 获取已授权抵音号列表
func (q *QianchuanClient) GetAuthorizedAwemeList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]QianchuanAweme, int, error) {
	path := "/v1.0/qianchuan/aweme/authorized/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []QianchuanAweme `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetAwemeInfo 获取抵音号详情
func (q *QianchuanClient) GetAwemeInfo(ctx context.Context, accessToken string, advertiserID, awemeID uint64) (*QianchuanAweme, error) {
	path := "/v1.0/qianchuan/aweme/info/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
	}
	var result struct {
		Data QianchuanAweme `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 第三方平台产品管理 ====================

// ExternalProduct 外部平台商品
type ExternalProduct struct {
	ExternalProductID string  `json:"external_product_id"`
	ProductName       string  `json:"product_name"`
	ProductImg        string  `json:"product_img"`
	Price             float64 `json:"price"`
	Platform          string  `json:"platform"`
	Status            int     `json:"status"`
}

// GetExternalProductList 获取第三方平台商品列表
func (q *QianchuanClient) GetExternalProductList(ctx context.Context, accessToken string, advertiserID, awemeID uint64, platform string, page, pageSize int) ([]ExternalProduct, int, error) {
	path := "/v1.0/qianchuan/external_product/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"platform":      platform,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ExternalProduct `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 直播管理 ====================

// LiveRoom 直播间信息
type LiveRoom struct {
	RoomID      uint64 `json:"room_id"`
	RoomTitle   string `json:"room_title"`
	AwemeID     uint64 `json:"aweme_id"`
	AwemeName   string `json:"aweme_name"`
	Status      int    `json:"status"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	WatchCount  int64  `json:"watch_count"`
	OnlineCount int64  `json:"online_count"`
}

// GetLiveRoomList 获取直播间列表
func (q *QianchuanClient) GetLiveRoomList(ctx context.Context, accessToken string, advertiserID uint64, awemeID uint64, page, pageSize int) ([]LiveRoom, int, error) {
	path := "/v1.0/qianchuan/aweme/live_room/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if awemeID > 0 {
		params["aweme_id"] = awemeID
	}
	var result struct {
		Data struct {
			List     []LiveRoom `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 计划高级更新API ====================

// UpdateScheduleDate 更新计划投放时间
func (q *QianchuanClient) UpdateScheduleDate(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, startTime, endTime string) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/schedule_date/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"start_time":    startTime,
		"end_time":      endTime,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateScheduleTime 更新计划投放时段
func (q *QianchuanClient) UpdateScheduleTime(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, scheduleTime string) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/schedule_time/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"schedule_time": scheduleTime,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateScheduleFixedRange 更新计划投放时长
func (q *QianchuanClient) UpdateScheduleFixedRange(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, scheduleFixedRange int) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/schedule_fixed_range/update/"
	data := map[string]interface{}{
		"advertiser_id":        advertiserID,
		"ad_ids":               adIDs,
		"schedule_fixed_range": scheduleFixedRange,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateRegion 更新计划地域定向
func (q *QianchuanClient) UpdateRegion(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, district string, city []uint64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/ad/region/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"district":      district,
	}
	if len(city) > 0 {
		data["city"] = city
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// GetLqAdList 获取低效计划列表
func (q *QianchuanClient) GetLqAdList(ctx context.Context, accessToken string, advertiserID uint64) ([]uint64, error) {
	path := "/v1.0/qianchuan/ad/lq_ad/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			AdIDs []uint64 `json:"ad_ids"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.AdIDs, nil
}

// SuggestRoiGoalResult ROI目标建议结果
type SuggestRoiGoalResult struct {
	SuggestRoiGoal float64 `json:"suggest_roi_goal"`
	MinRoiGoal     float64 `json:"min_roi_goal"`
	MaxRoiGoal     float64 `json:"max_roi_goal"`
}

// GetSuggestRoiGoal 获取支付ROI目标建议
func (q *QianchuanClient) GetSuggestRoiGoal(ctx context.Context, accessToken string, advertiserID, adID uint64) (*SuggestRoiGoalResult, error) {
	path := "/v1.0/qianchuan/ad/suggest_roi_goal/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data SuggestRoiGoalResult `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// SuggestBidResult 出价建议结果
type SuggestBidResult struct {
	SuggestBid float64 `json:"suggest_bid"`
	MinBid     float64 `json:"min_bid"`
	MaxBid     float64 `json:"max_bid"`
}

// GetSuggestBid 获取非ROI目标建议出价
func (q *QianchuanClient) GetSuggestBid(ctx context.Context, accessToken string, advertiserID, adID uint64) (*SuggestBidResult, error) {
	path := "/v1.0/qianchuan/ad/suggest_bid/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data SuggestBidResult `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// SuggestBudgetResult 预算建议结果
type SuggestBudgetResult struct {
	SuggestBudget float64 `json:"suggest_budget"`
	MinBudget     float64 `json:"min_budget"`
	MaxBudget     float64 `json:"max_budget"`
}

// GetSuggestBudget 获取建议预算
func (q *QianchuanClient) GetSuggestBudget(ctx context.Context, accessToken string, advertiserID, adID uint64) (*SuggestBudgetResult, error) {
	path := "/v1.0/qianchuan/ad/suggest_budget/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data SuggestBudgetResult `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// CompensateStatus 计划成本保障状态
type CompensateStatus struct {
	AdID             uint64 `json:"ad_id"`
	CompensateStatus int    `json:"compensate_status"`
}

// GetCompensateStatus 获取计划成本保障状态
func (q *QianchuanClient) GetCompensateStatus(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64) ([]CompensateStatus, error) {
	path := "/v1.0/qianchuan/ad/compensate_status/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
	}
	var result struct {
		Data struct {
			List []CompensateStatus `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// LearningStatus 计划学习期状态
type LearningStatus struct {
	AdID           uint64 `json:"ad_id"`
	LearningPhase  string `json:"learning_phase"`
	LearningStatus string `json:"learning_status"`
}

// GetLearningStatus 获取计划学习期状态
func (q *QianchuanClient) GetLearningStatus(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64) ([]LearningStatus, error) {
	path := "/v1.0/qianchuan/ad/learning_status/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
	}
	var result struct {
		Data struct {
			List []LearningStatus `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 直播间报表扩展 ====================

// LiveStat 今日直播数据
type LiveStat struct {
	LiveCount        int64   `json:"live_count"`
	WatchCount       int64   `json:"watch_count"`
	FansAddCount     int64   `json:"fans_add_count"`
	Gmv              float64 `json:"gmv"`
	PayOrderCount    int64   `json:"pay_order_count"`
	AvgWatchDuration float64 `json:"avg_watch_duration"`
}

// GetTodayLiveStat 获取今日直播数据
func (q *QianchuanClient) GetTodayLiveStat(ctx context.Context, accessToken string, advertiserID uint64) (*LiveStat, error) {
	path := "/v1.0/qianchuan/report/live/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data LiveStat `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// LiveRoomDetail 直播间详情
type LiveRoomDetail struct {
	RoomID           uint64  `json:"room_id"`
	RoomTitle        string  `json:"room_title"`
	AwemeID          uint64  `json:"aweme_id"`
	AwemeName        string  `json:"aweme_name"`
	StartTime        string  `json:"start_time"`
	EndTime          string  `json:"end_time"`
	Status           int     `json:"status"`
	WatchCount       int64   `json:"watch_count"`
	OnlineCount      int64   `json:"online_count"`
	FansAddCount     int64   `json:"fans_add_count"`
	Gmv              float64 `json:"gmv"`
	PayOrderCount    int64   `json:"pay_order_count"`
	AvgWatchDuration float64 `json:"avg_watch_duration"`
}

// GetLiveRoomDetail 获取直播间详情
func (q *QianchuanClient) GetLiveRoomDetail(ctx context.Context, accessToken string, advertiserID uint64, roomID uint64) (*LiveRoomDetail, error) {
	path := "/v1.0/qianchuan/live/room/detail/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"room_id":       roomID,
	}
	var result struct {
		Data LiveRoomDetail `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// LiveFlowPerformance 直播间流量表现
type LiveFlowPerformance struct {
	RoomID           uint64  `json:"room_id"`
	PaidFlowRatio    float64 `json:"paid_flow_ratio"`
	NaturalFlowRatio float64 `json:"natural_flow_ratio"`
	TotalFlow        int64   `json:"total_flow"`
}

// GetLiveFlowPerformance 获取直播间流量表现
func (q *QianchuanClient) GetLiveFlowPerformance(ctx context.Context, accessToken string, advertiserID uint64, roomID uint64) (*LiveFlowPerformance, error) {
	path := "/v1.0/qianchuan/live/room/flow_performance/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"room_id":       roomID,
	}
	var result struct {
		Data LiveFlowPerformance `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// LiveUserInsight 直播间用户洞察
type LiveUserInsight struct {
	GenderDistribution []map[string]interface{} `json:"gender_distribution"`
	AgeDistribution    []map[string]interface{} `json:"age_distribution"`
	CityDistribution   []map[string]interface{} `json:"city_distribution"`
}

// GetLiveUserInsight 获取直播间用户洞察
func (q *QianchuanClient) GetLiveUserInsight(ctx context.Context, accessToken string, advertiserID uint64, roomID uint64) (*LiveUserInsight, error) {
	path := "/v1.0/qianchuan/live/room/user/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"room_id":       roomID,
	}
	var result struct {
		Data LiveUserInsight `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// LiveProduct 直播间商品
type LiveProduct struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ProductImg  string  `json:"product_img"`
	Price       float64 `json:"price"`
	SaleCount   int64   `json:"sale_count"`
	Gmv         float64 `json:"gmv"`
	ClickCount  int64   `json:"click_count"`
}

// GetLiveProductList 获取直播间商品列表
func (q *QianchuanClient) GetLiveProductList(ctx context.Context, accessToken string, advertiserID uint64, roomID uint64, page, pageSize int) ([]LiveProduct, int, error) {
	path := "/v1.0/qianchuan/live/room/product/list/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"room_id":       roomID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LiveProduct `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 关键词管理扩展 ====================

// WordPackage 词包推荐关键词
type WordPackage struct {
	PackageID   uint64 `json:"package_id"`
	PackageName string `json:"package_name"`
	WordCount   int    `json:"word_count"`
}

// GetKeywordPackages 获取词包推荐关键词
func (q *QianchuanClient) GetKeywordPackages(ctx context.Context, accessToken string, advertiserID uint64) ([]WordPackage, error) {
	path := "/v1.0/qianchuan/ad/keyword_package/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []WordPackage `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// KeywordCheckResult 关键词校验结果
type KeywordCheckResult struct {
	Word    string `json:"word"`
	IsValid bool   `json:"is_valid"`
	Reason  string `json:"reason"`
}

// CheckKeywords 关键词合规校验
func (q *QianchuanClient) CheckKeywords(ctx context.Context, accessToken string, advertiserID uint64, words []string) ([]KeywordCheckResult, error) {
	path := "/v1.0/qianchuan/ad/keyword/check/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"words":         words,
	}
	var result struct {
		Data struct {
			List []KeywordCheckResult `json:"list"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// PrivateWord 否定词
type PrivateWord struct {
	Word      string `json:"word"`
	MatchType string `json:"match_type"`
}

// GetPrivateWords 获取否定词列表
func (q *QianchuanClient) GetPrivateWords(ctx context.Context, accessToken string, advertiserID, adID uint64) ([]PrivateWord, error) {
	path := "/v1.0/qianchuan/ad/privatewords/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	var result struct {
		Data struct {
			PrivateWords []PrivateWord `json:"private_words"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.PrivateWords, nil
}

// UpdatePrivateWords 全量更新否定词
func (q *QianchuanClient) UpdatePrivateWords(ctx context.Context, accessToken string, advertiserID, adID uint64, phraseWords, preciseWords []string) error {
	path := "/v1.0/qianchuan/ad/privatewords/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"phrase_words":  phraseWords,
		"precise_words": preciseWords,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 全域推广扩展API ====================

// UniPromotionMaterial 全域推广素材
type UniPromotionMaterial struct {
	MaterialID   uint64 `json:"material_id"`
	MaterialType string `json:"material_type"`
	VideoID      string `json:"video_id"`
	ImageID      string `json:"image_id"`
	Status       int    `json:"status"`
}

// GetUniPromotionMaterial 获取全域推广计划下素材
func (q *QianchuanClient) GetUniPromotionMaterial(ctx context.Context, accessToken string, advertiserID, adID uint64, page, pageSize int) ([]UniPromotionMaterial, int, error) {
	path := "/v1.0/qianchuan/uni_promotion/material/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []UniPromotionMaterial `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// DeleteUniPromotionMaterial 删除全域推广计划下素材
func (q *QianchuanClient) DeleteUniPromotionMaterial(ctx context.Context, accessToken string, advertiserID, adID uint64, materialIDs []uint64) error {
	path := "/v1.0/qianchuan/uni_promotion/material/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"material_ids":  materialIDs,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// UpdateUniPromotionBudget 更新全域推广计划预算
func (q *QianchuanClient) UpdateUniPromotionBudget(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, budget float64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/uni_promotion/ad_budget/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"budget":        budget,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateUniPromotionRoiGoal 更新全域推广控成本计划支付ROI目标
func (q *QianchuanClient) UpdateUniPromotionRoiGoal(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, roi2Goal float64) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/uni_promotion/ad_roi2_goal/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"roi2_goal":     roi2Goal,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateUniPromotionScheduleDate 更新全域推广计划投放时间
func (q *QianchuanClient) UpdateUniPromotionScheduleDate(ctx context.Context, accessToken string, advertiserID uint64, adIDs []uint64, startTime, endTime string) ([]uint64, []uint64, error) {
	path := "/v1.0/qianchuan/uni_promotion/ad_schedule_date/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_ids":        adIDs,
		"start_time":    startTime,
		"end_time":      endTime,
	}
	var result struct {
		Data struct {
			SuccessIDs []uint64 `json:"success_ad_ids"`
			FailIDs    []uint64 `json:"fail_ad_ids"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.SuccessIDs, result.Data.FailIDs, nil
}

// UpdateUniPromotionName 更新商品全域推广计划名称
func (q *QianchuanClient) UpdateUniPromotionName(ctx context.Context, accessToken string, advertiserID, adID uint64, adName string) error {
	path := "/v1.0/qianchuan/uni_promotion/ad_name/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"ad_name":       adName,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// UniPromotionAuthorizedAweme 全域推广可投抵音号
type UniPromotionAuthorizedAweme struct {
	AwemeID       uint64 `json:"aweme_id"`
	AwemeName     string `json:"aweme_name"`
	AwemeAvatar   string `json:"aweme_avatar"`
	FollowerCount int64  `json:"follower_count"`
}

// GetUniPromotionAuthorizedAweme 获取可投全域推广抵音号列表
func (q *QianchuanClient) GetUniPromotionAuthorizedAweme(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]UniPromotionAuthorizedAweme, int, error) {
	path := "/v1.0/qianchuan/uni_promotion/authorized/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []UniPromotionAuthorizedAweme `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 工具API扩展 ====================

// AdQuota 计划配额信息
type AdQuota struct {
	TotalQuota  int `json:"total_quota"`
	UsedQuota   int `json:"used_quota"`
	RemainQuota int `json:"remain_quota"`
}

// GetAdQuota 获取在投计划配额信息
func (q *QianchuanClient) GetAdQuota(ctx context.Context, accessToken string, advertiserID uint64) (*AdQuota, error) {
	path := "/v1.0/qianchuan/ad/quota/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data AdQuota `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GrayAbility 白名单能力
type GrayAbility struct {
	AbilityKey  string `json:"ability_key"`
	AbilityName string `json:"ability_name"`
	IsGray      bool   `json:"is_gray"`
}

// GetGrayAbilities 获取白名单能力
func (q *QianchuanClient) GetGrayAbilities(ctx context.Context, accessToken string, advertiserID uint64, abilityKeys []string) ([]GrayAbility, error) {
	path := "/v1.0/qianchuan/tools/gray/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ability_keys":  abilityKeys,
	}
	var result struct {
		Data struct {
			List []GrayAbility `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// LogSearchResult 日志查询结果
type LogSearchResult struct {
	OperateTime  string `json:"operate_time"`
	OperateType  string `json:"operate_type"`
	OperatorName string `json:"operator_name"`
	ObjectType   string `json:"object_type"`
	ObjectID     uint64 `json:"object_id"`
	ObjectName   string `json:"object_name"`
	BeforeValue  string `json:"before_value"`
	AfterValue   string `json:"after_value"`
}

// SearchLogs 日志查询
func (q *QianchuanClient) SearchLogs(ctx context.Context, accessToken string, advertiserID uint64, startTime, endTime string, page, pageSize int) ([]LogSearchResult, int, error) {
	path := "/v1.0/qianchuan/tools/log/search/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_time":    startTime,
		"end_time":      endTime,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LogSearchResult `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 商品竞争分析API ====================

// ProductAnalysis 商品竞争分析项
type ProductAnalysis struct {
	ProductID    uint64  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductImg   string  `json:"product_img"`
	Price        float64 `json:"price"`
	CompeteCount int     `json:"compete_count"`
	CompeteLevel string  `json:"compete_level"`
	AvgBid       float64 `json:"avg_bid"`
	AvgRoi       float64 `json:"avg_roi"`
}

// GetProductAnalysisList 获取商品竞争分析列表
func (q *QianchuanClient) GetProductAnalysisList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]ProductAnalysis, int, error) {
	path := "/v1.0/qianchuan/product/analyse/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ProductAnalysis `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CompareStatsData 效果对比数据
type CompareStatsData struct {
	MyStats       map[string]interface{} `json:"my_stats"`
	CompeteStats  map[string]interface{} `json:"compete_stats"`
	IndustryStats map[string]interface{} `json:"industry_stats"`
}

// GetProductCompareStats 商品竞争分析详情-效果对比
func (q *QianchuanClient) GetProductCompareStats(ctx context.Context, accessToken string, advertiserID, productID uint64, startDate, endDate string) (*CompareStatsData, error) {
	path := "/v1.0/qianchuan/product/analyse/compare_stats/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"product_id":    productID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data CompareStatsData `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// CompareCreative 创意比对数据
type CompareCreative struct {
	MyCreatives      []map[string]interface{} `json:"my_creatives"`
	CompeteCreatives []map[string]interface{} `json:"compete_creatives"`
}

// GetProductCompareCreative 商品竞争分析详情-创意比对
func (q *QianchuanClient) GetProductCompareCreative(ctx context.Context, accessToken string, advertiserID, productID uint64) (*CompareCreative, error) {
	path := "/v1.0/qianchuan/product/analyse/compare_creative/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"product_id":    productID,
	}
	var result struct {
		Data CompareCreative `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 随心推扩展API ====================

// EstimateProfit 投放效果预估
type EstimateProfit struct {
	EstimatedShow    int64   `json:"estimated_show"`
	EstimatedClick   int64   `json:"estimated_click"`
	EstimatedConvert int64   `json:"estimated_convert"`
	EstimatedCost    float64 `json:"estimated_cost"`
	EstimatedCPM     float64 `json:"estimated_cpm"`
}

// GetAwemeEstimateProfit 获取随心推投放效果预估
func (q *QianchuanClient) GetAwemeEstimateProfit(ctx context.Context, accessToken string, advertiserID, awemeID, itemID uint64, budget float64, deliveryGoal string) (*EstimateProfit, error) {
	path := "/v1.0/qianchuan/aweme/order/estimate_profit/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"item_id":       itemID,
		"budget":        budget,
		"delivery_goal": deliveryGoal,
	}
	var result struct {
		Data EstimateProfit `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetAwemeSuggestBid 获取随心推短视频建议出价
func (q *QianchuanClient) GetAwemeSuggestBid(ctx context.Context, accessToken string, advertiserID, awemeID uint64, deliveryGoal string) (float64, error) {
	path := "/v1.0/qianchuan/aweme/order/suggest_bid/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"delivery_goal": deliveryGoal,
	}
	var result struct {
		Data struct {
			SuggestBid float64 `json:"suggest_bid"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SuggestBid, nil
}

// GetAwemeSuggestRoiGoal 获取随心推ROI建议出价
func (q *QianchuanClient) GetAwemeSuggestRoiGoal(ctx context.Context, accessToken string, advertiserID, awemeID uint64) (float64, error) {
	path := "/v1.0/qianchuan/aweme/order/suggest_roi_goal/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
	}
	var result struct {
		Data struct {
			SuggestRoiGoal float64 `json:"suggest_roi_goal"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SuggestRoiGoal, nil
}

// AwemeOrderQuota 随心推订单配额
type AwemeOrderQuota struct {
	TotalQuota  int `json:"total_quota"`
	UsedQuota   int `json:"used_quota"`
	RemainQuota int `json:"remain_quota"`
}

// GetAwemeOrderQuota 查询随心推使用中订单配额信息
func (q *QianchuanClient) GetAwemeOrderQuota(ctx context.Context, accessToken string, advertiserID uint64) (*AwemeOrderQuota, error) {
	path := "/v1.0/qianchuan/aweme/order/quota/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data AwemeOrderQuota `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// AddAwemeOrderBudget 追加随心推订单预算
func (q *QianchuanClient) AddAwemeOrderBudget(ctx context.Context, accessToken string, advertiserID, orderID uint64, addBudget float64) error {
	path := "/v1.0/qianchuan/aweme/order/budget/add/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_id":      orderID,
		"add_budget":    addBudget,
	}
	return q.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetAwemeSuggestDeliveryTime 获取建议延长时长
func (q *QianchuanClient) GetAwemeSuggestDeliveryTime(ctx context.Context, accessToken string, advertiserID, orderID uint64) (float64, error) {
	path := "/v1.0/qianchuan/aweme/order/suggest_delivery_time/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"order_id":      orderID,
	}
	var result struct {
		Data struct {
			SuggestDeliveryTime float64 `json:"suggest_delivery_time"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SuggestDeliveryTime, nil
}

// AwemeInterestKeyword 随心推兴趣标签
type AwemeInterestKeyword struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// GetAwemeInterestKeywords 获取随心推兴趣标签
func (q *QianchuanClient) GetAwemeInterestKeywords(ctx context.Context, accessToken string, advertiserID uint64) ([]AwemeInterestKeyword, error) {
	path := "/v1.0/qianchuan/aweme/interest_action/interest_keyword/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []AwemeInterestKeyword `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 图文素材管理 ====================

// Carousel 图文素材
type Carousel struct {
	CarouselID   string   `json:"carousel_id"`
	CarouselName string   `json:"carousel_name"`
	ImageIDs     []string `json:"image_ids"`
	Status       int      `json:"status"`
	CreateTime   string   `json:"create_time"`
}

// GetCarouselList 获取千川素材库图文
func (q *QianchuanClient) GetCarouselList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]Carousel, int, error) {
	path := "/v1.0/qianchuan/carousel/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []Carousel `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// AwemeCarousel 抖音号下图文
type AwemeCarousel struct {
	ItemID     uint64   `json:"item_id"`
	Title      string   `json:"title"`
	ImageURLs  []string `json:"image_urls"`
	CreateTime string   `json:"create_time"`
}

// GetAwemeCarouselList 获取抖音号下图文
func (q *QianchuanClient) GetAwemeCarouselList(ctx context.Context, accessToken string, advertiserID, awemeID uint64, page, pageSize int) ([]AwemeCarousel, int, error) {
	path := "/v1.0/qianchuan/carousel/aweme/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"aweme_id":      awemeID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []AwemeCarousel `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetVideoOriginalList 获取首发素材
func (q *QianchuanClient) GetVideoOriginalList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]string, int, error) {
	path := "/v1.0/qianchuan/file/video/original/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			VideoIDs []string `json:"video_ids"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.VideoIDs, result.Data.PageInfo.TotalNumber, nil
}

// GetVideoEfficiencyList 获取低效素材
func (q *QianchuanClient) GetVideoEfficiencyList(ctx context.Context, accessToken string, advertiserID uint64) ([]string, error) {
	path := "/v1.0/qianchuan/file/video/efficiency/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			VideoIDs []string `json:"video_ids"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.VideoIDs, nil
}

// ==================== 行为兴趣词管理 ====================

// InterestActionWord 行为兴趣词对象
type InterestActionWord struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentID uint64 `json:"parent_id"`
	Level    int    `json:"level"`
}

// GetActionCategoriesV2 行为类目查询(V2带场景参数)
func (q *QianchuanClient) GetActionCategoriesV2(ctx context.Context, accessToken string, advertiserID uint64, actionScene string, actionDays int) ([]InterestActionWord, error) {
	path := "/v1.0/qianchuan/tools/interest_action/action/category/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"action_scene":  actionScene,
		"action_days":   actionDays,
	}
	var result struct {
		Data struct {
			List []InterestActionWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetActionKeywords 行为关键词查询
func (q *QianchuanClient) GetActionKeywords(ctx context.Context, accessToken string, advertiserID uint64, queryWord, actionScene string, actionDays int) ([]InterestActionWord, error) {
	path := "/v1.0/qianchuan/tools/interest_action/action/keyword/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    queryWord,
		"action_scene":  actionScene,
		"action_days":   actionDays,
	}
	var result struct {
		Data struct {
			List []InterestActionWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetInterestCategoriesV2 兴趣类目查询(V2返回InterestActionWord)
func (q *QianchuanClient) GetInterestCategoriesV2(ctx context.Context, accessToken string, advertiserID uint64) ([]InterestActionWord, error) {
	path := "/v1.0/qianchuan/tools/interest_action/interest/category/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []InterestActionWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// GetInterestKeywords 兴趣关键词查询
func (q *QianchuanClient) GetInterestKeywords(ctx context.Context, accessToken string, advertiserID uint64, queryWord string) ([]InterestActionWord, error) {
	path := "/v1.0/qianchuan/tools/interest_action/interest/keyword/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    queryWord,
	}
	var result struct {
		Data struct {
			List []InterestActionWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// Id2WordResult ID转词结果
type Id2WordResult struct {
	InterestCategories []InterestActionWord `json:"interest_categories"`
	InterestKeywords   []InterestActionWord `json:"interest_keywords"`
	ActionCategories   []InterestActionWord `json:"action_categories"`
	ActionKeywords     []InterestActionWord `json:"action_keywords"`
}

// Id2Word 兴趣行为类目关键词ID转词
func (q *QianchuanClient) Id2Word(ctx context.Context, accessToken string, advertiserID uint64, interestCategoryIDs, interestKeywordIDs, actionCategoryIDs, actionKeywordIDs []uint64) (*Id2WordResult, error) {
	path := "/v1.0/qianchuan/tools/interest_action/id2word/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	if len(interestCategoryIDs) > 0 {
		data["interest_category_ids"] = interestCategoryIDs
	}
	if len(interestKeywordIDs) > 0 {
		data["interest_keyword_ids"] = interestKeywordIDs
	}
	if len(actionCategoryIDs) > 0 {
		data["action_category_ids"] = actionCategoryIDs
	}
	if len(actionKeywordIDs) > 0 {
		data["action_keyword_ids"] = actionKeywordIDs
	}
	var result struct {
		Data Id2WordResult `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetKeywordSuggest 获取行为兴趣推荐关键词
func (q *QianchuanClient) GetKeywordSuggest(ctx context.Context, accessToken string, advertiserID uint64, keywords []string) ([]InterestActionWord, error) {
	path := "/v1.0/qianchuan/tools/interest_action/keyword/suggest/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"keywords":      keywords,
	}
	var result struct {
		Data struct {
			List []InterestActionWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 动态创意词包 ====================

// CreativeWord 动态创意词包
type CreativeWord struct {
	CreativeWordID uint64   `json:"creative_word_id"`
	Name           string   `json:"name"`
	Words          []string `json:"words"`
	CreateTime     string   `json:"create_time"`
}

// GetCreativeWordList 查询动态创意词包
func (q *QianchuanClient) GetCreativeWordList(ctx context.Context, accessToken string, advertiserID uint64) ([]CreativeWord, error) {
	path := "/v1.0/qianchuan/tools/creative_word/select/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []CreativeWord `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 抖音达人工具 ====================

// AwemeAuthorInfo 抖音达人信息
type AwemeAuthorInfo struct {
	AuthorID     uint64 `json:"author_id"`
	AuthorName   string `json:"author_name"`
	Avatar       string `json:"avatar"`
	FansCount    int64  `json:"fans_count"`
	CategoryName string `json:"category_name"`
}

// GetAwemeCategoryTopAuthors 查询抖音类目下的推荐达人
func (q *QianchuanClient) GetAwemeCategoryTopAuthors(ctx context.Context, accessToken string, advertiserID uint64, categoryID uint64) ([]AwemeAuthorInfo, error) {
	path := "/v1.0/qianchuan/tools/aweme_category_top_author/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"category_id":   categoryID,
	}
	var result struct {
		Data struct {
			List []AwemeAuthorInfo `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// AwemeCategory 抖音类目
type AwemeCategory struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentID uint64 `json:"parent_id"`
	Level    int    `json:"level"`
}

// GetAwemeMultiLevelCategories 查询抖音类目列表
func (q *QianchuanClient) GetAwemeMultiLevelCategories(ctx context.Context, accessToken string, advertiserID uint64) ([]AwemeCategory, error) {
	path := "/v1.0/qianchuan/tools/aweme_multi_level_category/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	var result struct {
		Data struct {
			List []AwemeCategory `json:"list"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// SearchSimilarAuthors 查询抖音类似帐号
func (q *QianchuanClient) SearchSimilarAuthors(ctx context.Context, accessToken string, advertiserID uint64, authorIDs []uint64) ([]AwemeAuthorInfo, error) {
	path := "/v1.0/qianchuan/tools/aweme_similar_author/search/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"author_ids":    authorIDs,
	}
	var result struct {
		Data struct {
			List []AwemeAuthorInfo `json:"list"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// SearchAwemeInfo 查询抖音帐号和类目信息
func (q *QianchuanClient) SearchAwemeInfo(ctx context.Context, accessToken string, advertiserID uint64, queryWord string) ([]AwemeAuthorInfo, []AwemeCategory, error) {
	path := "/v1.0/qianchuan/tools/aweme_info/search/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    queryWord,
	}
	var result struct {
		Data struct {
			Authors    []AwemeAuthorInfo `json:"authors"`
			Categories []AwemeCategory   `json:"categories"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, nil, err
	}
	return result.Data.Authors, result.Data.Categories, nil
}

// GetAwemeAuthorInfo 查询抖音号id对应的达人信息
func (q *QianchuanClient) GetAwemeAuthorInfo(ctx context.Context, accessToken string, advertiserID uint64, authorIDs []uint64) ([]AwemeAuthorInfo, error) {
	path := "/v1.0/qianchuan/tools/aweme_author_info/get/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"author_ids":    authorIDs,
	}
	var result struct {
		Data struct {
			List []AwemeAuthorInfo `json:"list"`
		} `json:"data"`
	}
	err := q.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// LiveAuthorizeInfo 授权直播达人信息
type LiveAuthorizeInfo struct {
	AuthorID        uint64 `json:"author_id"`
	AuthorName      string `json:"author_name"`
	Avatar          string `json:"avatar"`
	AuthorizeStatus int    `json:"authorize_status"`
}

// GetLiveAuthorizeList 查询授权直播抖音达人列表
func (q *QianchuanClient) GetLiveAuthorizeList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]LiveAuthorizeInfo, int, error) {
	path := "/v1.0/qianchuan/tools/live_authorize/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []LiveAuthorizeInfo `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := q.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}
