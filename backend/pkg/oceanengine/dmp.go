package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// DMPService DMP人群包服务
type DMPService struct {
	client *Client
}

// NewDMPService 创建DMP服务
func NewDMPService(client *Client) *DMPService {
	return &DMPService{client: client}
}

// CustomAudience 自定义人群
type CustomAudience struct {
	CustomAudienceID int64  `json:"custom_audience_id"`
	Name             string `json:"name"`
	Source           string `json:"source"`
	Status           int    `json:"status"`
	CoverNum         int64  `json:"cover_num"`
	Description      string `json:"description"`
	CreateTime       string `json:"create_time"`
	ModifyTime       string `json:"modify_time"`
}

// CustomAudienceListRequest 获取人群包列表请求
type CustomAudienceListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		CustomAudienceIDs []int64 `json:"custom_audience_ids,omitempty"`
		Name              string  `json:"name,omitempty"`
		LandingType       string  `json:"landing_type,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// CustomAudienceListResponse 获取人群包列表响应
type CustomAudienceListResponse struct {
	List     []CustomAudience `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetCustomAudienceList 获取人群包列表
func (s *DMPService) GetCustomAudienceList(ctx context.Context, req *CustomAudienceListRequest) (*CustomAudienceListResponse, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/select/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CustomAudienceListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CustomAudienceCreateRequest 创建人群包请求
type CustomAudienceCreateRequest struct {
	AdvertiserID           int64   `json:"advertiser_id"`
	Name                   string  `json:"name"`
	Description            string  `json:"description,omitempty"`
	DataSourceType         string  `json:"data_source_type"` // FILE, PIXEL, APP, SHOP, etc.
	DataSourceID           int64   `json:"data_source_id,omitempty"`
	DataSourceFilePath     string  `json:"data_source_file_path,omitempty"`
	DataSourceFileType     string  `json:"data_source_file_type,omitempty"` // IMEI, IDFA, MAC, PHONE, etc.
	CalculationType        string  `json:"calculation_type,omitempty"`
	RetargetingTagsInclude []int64 `json:"retargeting_tags_include,omitempty"`
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
}

// CustomAudienceCreateResponse 创建人群包响应
type CustomAudienceCreateResponse struct {
	CustomAudienceID int64 `json:"custom_audience_id"`
}

// CreateCustomAudience 创建人群包
func (s *DMPService) CreateCustomAudience(ctx context.Context, req *CustomAudienceCreateRequest) (*CustomAudienceCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CustomAudienceCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CustomAudienceUpdateRequest 更新人群包请求
type CustomAudienceUpdateRequest struct {
	AdvertiserID     int64  `json:"advertiser_id"`
	CustomAudienceID int64  `json:"custom_audience_id"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
}

// UpdateCustomAudience 更新人群包
func (s *DMPService) UpdateCustomAudience(ctx context.Context, req *CustomAudienceUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// CustomAudienceDeleteRequest 删除人群包请求
type CustomAudienceDeleteRequest struct {
	AdvertiserID     int64 `json:"advertiser_id"`
	CustomAudienceID int64 `json:"custom_audience_id"`
}

// DeleteCustomAudience 删除人群包
func (s *DMPService) DeleteCustomAudience(ctx context.Context, req *CustomAudienceDeleteRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/delete/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// CustomAudiencePublishRequest 发布人群包请求
type CustomAudiencePublishRequest struct {
	AdvertiserID        int64   `json:"advertiser_id"`
	CustomAudienceID    int64   `json:"custom_audience_id"`
	TargetAdvertiserIDs []int64 `json:"target_advertiser_ids,omitempty"`
}

// PublishCustomAudience 发布人群包
func (s *DMPService) PublishCustomAudience(ctx context.Context, req *CustomAudiencePublishRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/publish/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// CustomAudiencePushRequest 推送人群包请求
type CustomAudiencePushRequest struct {
	AdvertiserID        int64   `json:"advertiser_id"`
	CustomAudienceID    int64   `json:"custom_audience_id"`
	TargetAdvertiserIDs []int64 `json:"target_advertiser_ids"`
}

// PushCustomAudience 推送人群包到其他广告主
func (s *DMPService) PushCustomAudience(ctx context.Context, req *CustomAudiencePushRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/push_v2/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// DataSourceFileUploadRequest 上传数据源文件请求
type DataSourceFileUploadRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	File         []byte `json:"-"` // file content
	Filename     string `json:"-"`
}

// DataSourceFileUploadResponse 上传数据源文件响应
type DataSourceFileUploadResponse struct {
	Path string `json:"path"`
}

// UploadDataSourceFile 上传数据源文件
func (s *DMPService) UploadDataSourceFile(ctx context.Context, advertiserID int64, filename string, data []byte) (*DataSourceFileUploadResponse, error) {
	// 使用 FileService 的上传逻辑
	// 此处简化处理，实际需要根据API文档实现multipart上传
	reqBody := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/dmp/data_source/file/upload/", reqBody)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result DataSourceFileUploadResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// LookalikeAudienceCreateRequest 创建相似人群请求
type LookalikeAudienceCreateRequest struct {
	AdvertiserID     int64   `json:"advertiser_id"`
	Name             string  `json:"name"`
	Description      string  `json:"description,omitempty"`
	CustomAudienceID int64   `json:"custom_audience_id"`
	LookalikeLevel   int     `json:"lookalike_level"`         // 1-10
	LocationType     string  `json:"location_type,omitempty"` // ALL, CUSTOM
	LocationIDs      []int64 `json:"location_ids,omitempty"`
}

// LookalikeAudienceCreateResponse 创建相似人群响应
type LookalikeAudienceCreateResponse struct {
	CustomAudienceID int64 `json:"custom_audience_id"`
}

// CreateLookalikeAudience 创建相似人群
func (s *DMPService) CreateLookalikeAudience(ctx context.Context, req *LookalikeAudienceCreateRequest) (*LookalikeAudienceCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/lookalike/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result LookalikeAudienceCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 人群包详情 ====================

// CustomAudienceReadRequest 获取人群包详情请求
type CustomAudienceReadRequest struct {
	AdvertiserID      int64   `json:"advertiser_id"`
	CustomAudienceIDs []int64 `json:"custom_audience_ids"`
}

// GetCustomAudienceDetail 获取人群包详情
func (s *DMPService) GetCustomAudienceDetail(ctx context.Context, req *CustomAudienceReadRequest) ([]CustomAudience, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/read/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []CustomAudience `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 数据源管理 ====================

// DataSource 数据源
type DataSource struct {
	DataSourceID   int64  `json:"data_source_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	DataSourceType string `json:"data_source_type"`
	Status         int    `json:"status"`
	CoverNum       int64  `json:"cover_num"`
	CreateTime     string `json:"create_time"`
	ModifyTime     string `json:"modify_time"`
}

// DataSourceCreateRequest 创建数据源请求
type DataSourceCreateRequest struct {
	AdvertiserID   int64  `json:"advertiser_id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	DataSourceType string `json:"data_source_type"` // FILE, PIXEL, APP
	FilePath       string `json:"file_path,omitempty"`
	FileType       string `json:"file_type,omitempty"`
}

// DataSourceCreateResponse 创建数据源响应
type DataSourceCreateResponse struct {
	DataSourceID string `json:"data_source_id"`
}

// CreateDataSource 创建数据源
func (s *DMPService) CreateDataSource(ctx context.Context, req *DataSourceCreateRequest) (*DataSourceCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/data_source/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result DataSourceCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// DataSourceUpdateRequest 更新数据源请求
type DataSourceUpdateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	DataSourceID string `json:"data_source_id"`
	FilePath     string `json:"file_path,omitempty"`
	Operation    string `json:"operation,omitempty"` // ADD, DELETE
}

// UpdateDataSource 更新数据源
func (s *DMPService) UpdateDataSource(ctx context.Context, req *DataSourceUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/data_source/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// DataSourceReadRequest 获取数据源详情请求
type DataSourceReadRequest struct {
	AdvertiserID  int64    `json:"advertiser_id"`
	DataSourceIDs []string `json:"data_source_ids"`
}

// GetDataSourceDetail 获取数据源详情
func (s *DMPService) GetDataSourceDetail(ctx context.Context, req *DataSourceReadRequest) ([]DataSource, error) {
	resp, err := s.client.Post(ctx, "/2/dmp/data_source/read/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []DataSource `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 云图相关接口 ====================

// BrandInfo 云图品牌信息
type BrandInfo struct {
	BrandID   int64  `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

// GetBrandList 获取广告账户关联云图账户信息
func (s *DMPService) GetBrandList(ctx context.Context, advertiserID int64) ([]BrandInfo, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/dmp/brand/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []BrandInfo `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// CustomAudienceCopyRequest 推送人群包到云图请求
type CustomAudienceCopyRequest struct {
	AdvertiserID     int64 `json:"advertiser_id"`
	CustomAudienceID int64 `json:"custom_audience_id"`
	BrandID          int64 `json:"brand_id"`
}

// CopyCustomAudienceToBrand 推送dmp人群包到云图账户
func (s *DMPService) CopyCustomAudienceToBrand(ctx context.Context, req *CustomAudienceCopyRequest) error {
	resp, err := s.client.Post(ctx, "/2/dmp/custom_audience/copy/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 行为兴趣定向 ====================

// InterestCategory 兴趣分类
type InterestCategory struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
	Level    int    `json:"level"`
}

// GetInterestCategories 获取兴趣分类列表
func (s *DMPService) GetInterestCategories(ctx context.Context, advertiserID int64) ([]InterestCategory, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/tools/interest_category/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []InterestCategory `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ActionCategory 行为分类
type ActionCategory struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
	Level    int    `json:"level"`
}

// GetActionCategories 获取行为分类列表
func (s *DMPService) GetActionCategories(ctx context.Context, advertiserID int64, actionScene string, actionDays int) ([]ActionCategory, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
		"action_scene":  actionScene,
		"action_days":   actionDays,
	}

	resp, err := s.client.Post(ctx, "/2/tools/action_category/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []ActionCategory `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// InterestKeyword 兴趣关键词
type InterestKeyword struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SearchInterestKeywords 搜索兴趣关键词
func (s *DMPService) SearchInterestKeywords(ctx context.Context, advertiserID int64, query string) ([]InterestKeyword, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    query,
	}

	resp, err := s.client.Post(ctx, "/2/tools/interest_keyword/select/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []InterestKeyword `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 抖音达人定向 ====================

// AwemeAuthor 抖音达人
type AwemeAuthor struct {
	AuthorID     int64  `json:"author_id"`
	AuthorName   string `json:"author_name"`
	Avatar       string `json:"avatar"`
	FansCount    int64  `json:"fans_count"`
	CategoryName string `json:"category_name"`
}

// SearchAwemeAuthors 搜索抖音达人
func (s *DMPService) SearchAwemeAuthors(ctx context.Context, advertiserID int64, query string) ([]AwemeAuthor, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
		"query_word":    query,
	}

	resp, err := s.client.Post(ctx, "/2/tools/aweme_author/select/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AwemeAuthor `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// AwemeAuthorCategory 抖音达人分类
type AwemeAuthorCategory struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

// GetAwemeAuthorCategories 获取抖音达人分类
func (s *DMPService) GetAwemeAuthorCategories(ctx context.Context, advertiserID int64) ([]AwemeAuthorCategory, error) {
	req := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	resp, err := s.client.Post(ctx, "/2/tools/aweme_author_category/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []AwemeAuthorCategory `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== 人群估算 ====================

// AudienceEstimateRequest 人群估算请求
type AudienceEstimateRequest struct {
	AdvertiserID           int64    `json:"advertiser_id"`
	Gender                 string   `json:"gender,omitempty"`
	Age                    []string `json:"age,omitempty"`
	City                   []int64  `json:"city,omitempty"`
	InterestCategories     []int64  `json:"interest_categories,omitempty"`
	ActionCategories       []int64  `json:"action_categories,omitempty"`
	RetargetingTags        []int64  `json:"retargeting_tags,omitempty"`
	RetargetingTagsExclude []int64  `json:"retargeting_tags_exclude,omitempty"`
}

// AudienceEstimateResponse 人群估算响应
type AudienceEstimateResponse struct {
	UserCount int64 `json:"user_count"`
}

// EstimateAudience 估算人群覆盖
func (s *DMPService) EstimateAudience(ctx context.Context, req *AudienceEstimateRequest) (*AudienceEstimateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/tools/estimate_audience/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AudienceEstimateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}
