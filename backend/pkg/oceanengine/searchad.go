package oceanengine

import (
	"context"
)

// SearchAdClient 搜索广告客户端
type SearchAdClient struct {
	client *Client
}

// NewSearchAdClient 创建搜索广告客户端
func (c *Client) SearchAd() *SearchAdClient {
	return &SearchAdClient{client: c}
}

// ==================== 关键词管理 ====================

// SearchKeyword 搜索关键词
type SearchKeyword struct {
	KeywordID uint64  `json:"keyword_id"`
	AdID      uint64  `json:"ad_id"`
	Word      string  `json:"word"`
	MatchType string  `json:"match_type"` // PHRASE-短语匹配, PRECISE-精准匹配, EXTENSIVE-广泛匹配
	Bid       float64 `json:"bid"`
	Status    string  `json:"status"`
	Quality   int     `json:"quality"` // 质量分
	CPCBid    float64 `json:"cpc_bid"`
}

// KeywordCreateRequest 创建关键词请求
type KeywordCreateRequest struct {
	AdvertiserID uint64             `json:"advertiser_id"`
	AdID         uint64             `json:"ad_id"`
	Keywords     []KeywordForCreate `json:"keywords"`
}

// KeywordForCreate 创建关键词条目
type KeywordForCreate struct {
	Word      string  `json:"word"`
	MatchType string  `json:"match_type"`
	Bid       float64 `json:"bid,omitempty"`
}

// KeywordUpdateRequest 更新关键词请求
type KeywordUpdateRequest struct {
	AdvertiserID uint64             `json:"advertiser_id"`
	AdID         uint64             `json:"ad_id"`
	Keywords     []KeywordForUpdate `json:"keywords"`
}

// KeywordForUpdate 更新关键词条目
type KeywordForUpdate struct {
	KeywordID uint64  `json:"keyword_id"`
	Bid       float64 `json:"bid,omitempty"`
	MatchType string  `json:"match_type,omitempty"`
}

// KeywordResponse 关键词操作响应
type KeywordResponse struct {
	SuccessKeywordIDs []uint64 `json:"success_keyword_ids"`
	FailKeywordIDs    []uint64 `json:"fail_keyword_ids"`
}

// GetKeywordList 获取关键词列表
func (s *SearchAdClient) GetKeywordList(ctx context.Context, accessToken string, advertiserID, adID uint64, page, pageSize int) ([]SearchKeyword, int, error) {
	path := "/v2.0/keyword/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []SearchKeyword `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreateKeywords 创建关键词
func (s *SearchAdClient) CreateKeywords(ctx context.Context, accessToken string, req *KeywordCreateRequest) (*KeywordResponse, error) {
	path := "/v2.0/keyword/create/"
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateKeywords 更新关键词
func (s *SearchAdClient) UpdateKeywords(ctx context.Context, accessToken string, req *KeywordUpdateRequest) (*KeywordResponse, error) {
	path := "/v2.0/keyword/update/"
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DeleteKeywords 删除关键词
func (s *SearchAdClient) DeleteKeywords(ctx context.Context, accessToken string, advertiserID, adID uint64, keywordIDs []uint64) (*KeywordResponse, error) {
	path := "/v2.0/keyword/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"keyword_ids":   keywordIDs,
	}
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// SuggestKeyword 推荐关键词
type SuggestKeyword struct {
	Word         string  `json:"word"`
	MatchType    string  `json:"match_type"`
	SuggestBid   float64 `json:"suggest_bid"`
	Competition  string  `json:"competition"`
	SearchVolume int64   `json:"search_volume"`
	PredictShow  int64   `json:"predict_show"`
	PredictClick int64   `json:"predict_click"`
}

// GetSuggestKeywords 获取推荐关键词
func (s *SearchAdClient) GetSuggestKeywords(ctx context.Context, accessToken string, advertiserID uint64, seedWords []string, count int) ([]SuggestKeyword, error) {
	path := "/v2.0/keyword/suggest/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"seed_words":    seedWords,
		"count":         count,
	}
	var result struct {
		Data struct {
			List []SuggestKeyword `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== V3体验版关键词管理 ====================

// GetKeywordListV3 获取关键词列表(V3体验版)
func (s *SearchAdClient) GetKeywordListV3(ctx context.Context, accessToken string, advertiserID, promotionID uint64, page, pageSize int) ([]SearchKeyword, int, error) {
	path := "/v3.0/keyword/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []SearchKeyword `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// CreateKeywordsV3 创建关键词(V3体验版)
func (s *SearchAdClient) CreateKeywordsV3(ctx context.Context, accessToken string, advertiserID, promotionID uint64, keywords []KeywordForCreate) (*KeywordResponse, error) {
	path := "/v3.0/keyword/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"keywords":      keywords,
	}
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateKeywordsV3 更新关键词(V3体验版)
func (s *SearchAdClient) UpdateKeywordsV3(ctx context.Context, accessToken string, advertiserID, promotionID uint64, keywords []KeywordForUpdate) (*KeywordResponse, error) {
	path := "/v3.0/keyword/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"keywords":      keywords,
	}
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// DeleteKeywordsV3 删除关键词(V3体验版)
func (s *SearchAdClient) DeleteKeywordsV3(ctx context.Context, accessToken string, advertiserID, promotionID uint64, keywordIDs []uint64) (*KeywordResponse, error) {
	path := "/v3.0/keyword/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"keyword_ids":   keywordIDs,
	}
	var result struct {
		Data KeywordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetSuggestKeywordsV3 获取推荐关键词(V3体验版)
func (s *SearchAdClient) GetSuggestKeywordsV3(ctx context.Context, accessToken string, advertiserID uint64, seedWords []string, count int) ([]SuggestKeyword, error) {
	path := "/v3.0/keyword/suggest/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"seed_words":    seedWords,
		"count":         count,
	}
	var result struct {
		Data struct {
			List []SuggestKeyword `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 否定词管理 ====================

// PrivativeWord 否定词
type PrivativeWord struct {
	WordID    uint64 `json:"word_id"`
	Word      string `json:"word"`
	MatchType string `json:"match_type"` // PHRASE-短语否定, PRECISE-精确否定
}

// PrivativeWordResponse 否定词操作响应
type PrivativeWordResponse struct {
	SuccessWordIDs []uint64 `json:"success_word_ids"`
	FailWordIDs    []uint64 `json:"fail_word_ids"`
}

// GetPrivativeWords 获取否定词列表
func (s *SearchAdClient) GetPrivativeWords(ctx context.Context, accessToken string, advertiserID uint64, objectType string, objectID uint64, page, pageSize int) ([]PrivativeWord, int, error) {
	path := "/v2.0/privativeword/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"object_type":   objectType, // AD-计划级, CAMPAIGN-广告组级
		"object_id":     objectID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []PrivativeWord `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// AddAdPrivativeWords 批量新增计划否定词
func (s *SearchAdClient) AddAdPrivativeWords(ctx context.Context, accessToken string, advertiserID, adID uint64, phraseWords, preciseWords []string) (*PrivativeWordResponse, error) {
	path := "/v2.0/privativeword/ad/add/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}
	if len(phraseWords) > 0 {
		data["phrase_words"] = phraseWords
	}
	if len(preciseWords) > 0 {
		data["precise_words"] = preciseWords
	}
	var result struct {
		Data PrivativeWordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateAdPrivativeWords 设置计划否定词(全量更新)
func (s *SearchAdClient) UpdateAdPrivativeWords(ctx context.Context, accessToken string, advertiserID, adID uint64, phraseWords, preciseWords []string) error {
	path := "/v2.0/privativeword/ad/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"phrase_words":  phraseWords,
		"precise_words": preciseWords,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// AddCampaignPrivativeWords 批量新增广告组否定词
func (s *SearchAdClient) AddCampaignPrivativeWords(ctx context.Context, accessToken string, advertiserID, campaignID uint64, phraseWords, preciseWords []string) (*PrivativeWordResponse, error) {
	path := "/v2.0/privativeword/campaign/add/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"campaign_id":   campaignID,
	}
	if len(phraseWords) > 0 {
		data["phrase_words"] = phraseWords
	}
	if len(preciseWords) > 0 {
		data["precise_words"] = preciseWords
	}
	var result struct {
		Data PrivativeWordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateCampaignPrivativeWords 设置广告组否定词(全量更新)
func (s *SearchAdClient) UpdateCampaignPrivativeWords(ctx context.Context, accessToken string, advertiserID, campaignID uint64, phraseWords, preciseWords []string) error {
	path := "/v2.0/privativeword/campaign/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"campaign_id":   campaignID,
		"phrase_words":  phraseWords,
		"precise_words": preciseWords,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== V3体验版否定词管理 ====================

// AddProjectPrivativeWordsV3 批量新增项目否定词(V3)
func (s *SearchAdClient) AddProjectPrivativeWordsV3(ctx context.Context, accessToken string, advertiserID, projectID uint64, phraseWords, preciseWords []string) (*PrivativeWordResponse, error) {
	path := "/v3.0/privativeword/add/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
	}
	if len(phraseWords) > 0 {
		data["phrase_words"] = phraseWords
	}
	if len(preciseWords) > 0 {
		data["precise_words"] = preciseWords
	}
	var result struct {
		Data PrivativeWordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// UpdateProjectPrivativeWordsV3 设置项目否定词(V3全量更新)
func (s *SearchAdClient) UpdateProjectPrivativeWordsV3(ctx context.Context, accessToken string, advertiserID, projectID uint64, phraseWords, preciseWords []string) (*PrivativeWordResponse, error) {
	path := "/v3.0/privativeword/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"phrase_words":  phraseWords,
		"precise_words": preciseWords,
	}
	var result struct {
		Data PrivativeWordResponse `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetProjectPrivativeWordsV3 获取项目否定词列表(V3)
func (s *SearchAdClient) GetProjectPrivativeWordsV3(ctx context.Context, accessToken string, advertiserID, projectID uint64) ([]PrivativeWord, error) {
	path := "/v3.0/privativeword/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
	}
	var result struct {
		Data struct {
			List []PrivativeWord `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 蓝海流量包管理 ====================

// BlueFlowPackage 蓝海流量包
type BlueFlowPackage struct {
	PackageID   uint64 `json:"package_id"`
	PackageName string `json:"package_name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	KeywordNum  int    `json:"keyword_num"`
}

// GetBlueFlowPackageList 获取蓝海流量包列表
func (s *SearchAdClient) GetBlueFlowPackageList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]BlueFlowPackage, int, error) {
	path := "/v3.0/blueflow/package/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []BlueFlowPackage `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// BlueFlowKeyword 蓝海关键词
type BlueFlowKeyword struct {
	KeywordID    uint64  `json:"keyword_id"`
	Word         string  `json:"word"`
	MatchType    string  `json:"match_type"`
	SuggestBid   float64 `json:"suggest_bid"`
	Competition  string  `json:"competition"`
	SearchVolume int64   `json:"search_volume"`
}

// GetBlueFlowKeywordList 获取广告下可用蓝海关键词
func (s *SearchAdClient) GetBlueFlowKeywordList(ctx context.Context, accessToken string, advertiserID, promotionID uint64, page, pageSize int) ([]BlueFlowKeyword, int, error) {
	path := "/v3.0/blueflow/keyword/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"promotion_id":  promotionID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []BlueFlowKeyword `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 关键词出价系数管理 ====================

// KeywordBidRatio 关键词出价系数
type KeywordBidRatio struct {
	RatioID      uint64  `json:"ratio_id"`
	AdvertiserID uint64  `json:"advertiser_id"`
	ProjectID    uint64  `json:"project_id"`
	KeywordID    uint64  `json:"keyword_id"`
	Ratio        float64 `json:"ratio"`
	Status       int     `json:"status"`
	CreateTime   string  `json:"create_time"`
	UpdateTime   string  `json:"update_time"`
}

// CreateKeywordBidRatio 创建关键词出价系数
func (s *SearchAdClient) CreateKeywordBidRatio(ctx context.Context, accessToken string, advertiserID, projectID, keywordID uint64, ratio float64) (uint64, error) {
	path := "/v2.0/tools/keywords_bid_ratio/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"keyword_id":    keywordID,
		"ratio":         ratio,
	}
	var result struct {
		Data struct {
			RatioID uint64 `json:"ratio_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.RatioID, nil
}

// UpdateKeywordBidRatio 更新关键词出价系数
func (s *SearchAdClient) UpdateKeywordBidRatio(ctx context.Context, accessToken string, advertiserID, ratioID uint64, ratio float64) error {
	path := "/v2.0/tools/keywords_bid_ratio/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ratio_id":      ratioID,
		"ratio":         ratio,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteKeywordBidRatio 删除关键词出价系数
func (s *SearchAdClient) DeleteKeywordBidRatio(ctx context.Context, accessToken string, advertiserID uint64, ratioIDs []uint64) error {
	path := "/v2.0/tools/keywords_bid_ratio/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ratio_ids":     ratioIDs,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetKeywordBidRatioList 获取关键词出价系数列表
func (s *SearchAdClient) GetKeywordBidRatioList(ctx context.Context, accessToken string, advertiserID, projectID uint64, page, pageSize int) ([]KeywordBidRatio, int, error) {
	path := "/v2.0/tools/keywords_bid_ratio/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []KeywordBidRatio `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// GetKeywordBidRatioProjectInfo 获取项目下关键词出价系数信息
func (s *SearchAdClient) GetKeywordBidRatioProjectInfo(ctx context.Context, accessToken string, advertiserID, projectID uint64) (map[string]interface{}, error) {
	path := "/v2.0/tools/keywords_bid_ratio/project_info/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"project_id":    projectID,
	}
	var result struct {
		Data map[string]interface{} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}
