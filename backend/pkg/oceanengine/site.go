package oceanengine

import (
	"context"
)

// SiteClient 建站管理客户端
type SiteClient struct {
	client *Client
}

// NewSiteClient 创建建站客户端
func NewSiteClient(client *Client) *SiteClient {
	return &SiteClient{client: client}
}

// Site 返回建站管理客户端
func (c *Client) Site() *SiteClient {
	return &SiteClient{client: c}
}

// ==================== 橙子建站 ====================

// OrangeSite 橙子建站落地页
type OrangeSite struct {
	SiteID     uint64 `json:"site_id"`
	SiteName   string `json:"site_name"`
	SiteURL    string `json:"site_url"`
	PreviewURL string `json:"preview_url"`
	Status     int    `json:"status"`
	SiteType   string `json:"site_type"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// GetOrangeSiteList 获取橙子建站落地页列表
func (s *SiteClient) GetOrangeSiteList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]OrangeSite, int, error) {
	path := "/v3.0/tools/orange_site/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []OrangeSite `json:"list"`
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

// GetOrangeSiteDetail 获取橙子建站落地页详情
func (s *SiteClient) GetOrangeSiteDetail(ctx context.Context, accessToken string, advertiserID, siteID uint64) (*OrangeSite, error) {
	path := "/v3.0/tools/orange_site/detail/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	var result struct {
		Data OrangeSite `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// OrangeSiteCreateRequest 创建橙子建站落地页请求
type OrangeSiteCreateRequest struct {
	AdvertiserID uint64                   `json:"advertiser_id"`
	SiteName     string                   `json:"site_name"`
	SiteType     string                   `json:"site_type"`
	Components   []map[string]interface{} `json:"components"`
}

// CreateOrangeSite 创建橙子建站落地页
func (s *SiteClient) CreateOrangeSite(ctx context.Context, accessToken string, req *OrangeSiteCreateRequest) (uint64, error) {
	path := "/v3.0/tools/orange_site/create/"
	var result struct {
		Data struct {
			SiteID uint64 `json:"site_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SiteID, nil
}

// UpdateOrangeSite 更新橙子建站落地页
func (s *SiteClient) UpdateOrangeSite(ctx context.Context, accessToken string, advertiserID, siteID uint64, siteName string, components []map[string]interface{}) error {
	path := "/v3.0/tools/orange_site/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	if siteName != "" {
		data["site_name"] = siteName
	}
	if len(components) > 0 {
		data["components"] = components
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// PublishOrangeSite 发布橙子建站落地页
func (s *SiteClient) PublishOrangeSite(ctx context.Context, accessToken string, advertiserID, siteID uint64) (string, error) {
	path := "/v3.0/tools/orange_site/publish/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	var result struct {
		Data struct {
			SiteURL string `json:"site_url"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.SiteURL, nil
}

// DeleteOrangeSite 删除橙子建站落地页
func (s *SiteClient) DeleteOrangeSite(ctx context.Context, accessToken string, advertiserID, siteID uint64) error {
	path := "/v3.0/tools/orange_site/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// CopyOrangeSite 复制橙子建站落地页
func (s *SiteClient) CopyOrangeSite(ctx context.Context, accessToken string, advertiserID, siteID uint64, siteName string) (uint64, error) {
	path := "/v3.0/tools/orange_site/copy/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
		"site_name":     siteName,
	}
	var result struct {
		Data struct {
			SiteID uint64 `json:"site_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SiteID, nil
}

// ==================== 第三方落地页管理 ====================

// ThirdPartySite 第三方落地页
type ThirdPartySite struct {
	SiteID      uint64 `json:"site_id"`
	SiteName    string `json:"site_name"`
	SiteURL     string `json:"site_url"`
	Status      int    `json:"status"`
	AuditStatus int    `json:"audit_status"`
	AuditReason string `json:"audit_reason"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// GetThirdPartySiteList 获取第三方落地页列表
func (s *SiteClient) GetThirdPartySiteList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]ThirdPartySite, int, error) {
	path := "/v3.0/tools/third_party_site/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ThirdPartySite `json:"list"`
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

// CreateThirdPartySite 创建第三方落地页
func (s *SiteClient) CreateThirdPartySite(ctx context.Context, accessToken string, advertiserID uint64, siteName, siteURL string) (uint64, error) {
	path := "/v3.0/tools/third_party_site/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_name":     siteName,
		"site_url":      siteURL,
	}
	var result struct {
		Data struct {
			SiteID uint64 `json:"site_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SiteID, nil
}

// UpdateThirdPartySite 更新第三方落地页
func (s *SiteClient) UpdateThirdPartySite(ctx context.Context, accessToken string, advertiserID, siteID uint64, siteName, siteURL string) error {
	path := "/v3.0/tools/third_party_site/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	if siteName != "" {
		data["site_name"] = siteName
	}
	if siteURL != "" {
		data["site_url"] = siteURL
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteThirdPartySite 删除第三方落地页
func (s *SiteClient) DeleteThirdPartySite(ctx context.Context, accessToken string, advertiserID, siteID uint64) error {
	path := "/v3.0/tools/third_party_site/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 落地页模板 ====================

// SiteTemplate 落地页模板
type SiteTemplate struct {
	TemplateID   uint64 `json:"template_id"`
	TemplateName string `json:"template_name"`
	TemplateType string `json:"template_type"`
	PreviewURL   string `json:"preview_url"`
	Industry     string `json:"industry"`
}

// GetSiteTemplateList 获取落地页模板列表
func (s *SiteClient) GetSiteTemplateList(ctx context.Context, accessToken string, advertiserID uint64, templateType string, page, pageSize int) ([]SiteTemplate, int, error) {
	path := "/v3.0/tools/site_template/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	if templateType != "" {
		params["template_type"] = templateType
	}
	var result struct {
		Data struct {
			List     []SiteTemplate `json:"list"`
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

// CreateSiteFromTemplate 从模板创建落地页
func (s *SiteClient) CreateSiteFromTemplate(ctx context.Context, accessToken string, advertiserID, templateID uint64, siteName string) (uint64, error) {
	path := "/v3.0/tools/site/create_from_template/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"template_id":   templateID,
		"site_name":     siteName,
	}
	var result struct {
		Data struct {
			SiteID uint64 `json:"site_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SiteID, nil
}

// ==================== 落地页组件 ====================

// SiteComponent 落地页组件
type SiteComponent struct {
	ComponentID   string                 `json:"component_id"`
	ComponentType string                 `json:"component_type"`
	ComponentName string                 `json:"component_name"`
	Properties    map[string]interface{} `json:"properties"`
}

// GetSiteComponentList 获取落地页组件列表
func (s *SiteClient) GetSiteComponentList(ctx context.Context, accessToken string, advertiserID uint64, componentType string) ([]SiteComponent, error) {
	path := "/v3.0/tools/site_component/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}
	if componentType != "" {
		params["component_type"] = componentType
	}
	var result struct {
		Data struct {
			List []SiteComponent `json:"list"`
		} `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 落地页分析 ====================

// SiteAnalysis 落地页分析数据
type SiteAnalysis struct {
	SiteID         uint64  `json:"site_id"`
	PV             int64   `json:"pv"`
	UV             int64   `json:"uv"`
	BounceRate     float64 `json:"bounce_rate"`
	AvgStayTime    float64 `json:"avg_stay_time"`
	ConversionRate float64 `json:"conversion_rate"`
	FormSubmit     int64   `json:"form_submit"`
	ButtonClick    int64   `json:"button_click"`
	PhoneClick     int64   `json:"phone_click"`
}

// GetSiteAnalysis 获取落地页分析数据
func (s *SiteClient) GetSiteAnalysis(ctx context.Context, accessToken string, advertiserID, siteID uint64, startDate, endDate string) (*SiteAnalysis, error) {
	path := "/v3.0/tools/site_analysis/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data SiteAnalysis `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// SiteHeatmap 落地页热力图数据
type SiteHeatmap struct {
	SiteID     uint64                   `json:"site_id"`
	HeatmapURL string                   `json:"heatmap_url"`
	ClickData  []map[string]interface{} `json:"click_data"`
}

// GetSiteHeatmap 获取落地页热力图
func (s *SiteClient) GetSiteHeatmap(ctx context.Context, accessToken string, advertiserID, siteID uint64, startDate, endDate string) (*SiteHeatmap, error) {
	path := "/v3.0/tools/site_heatmap/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data SiteHeatmap `json:"data"`
	}
	err := s.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 微页面（小程序落地页） ====================

// MiniPage 微页面/小程序落地页
type MiniPage struct {
	PageID     uint64 `json:"page_id"`
	PageName   string `json:"page_name"`
	PagePath   string `json:"page_path"`
	AppID      string `json:"app_id"`
	Status     int    `json:"status"`
	PreviewURL string `json:"preview_url"`
	CreateTime string `json:"create_time"`
}

// GetMiniPageList 获取微页面列表
func (s *SiteClient) GetMiniPageList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]MiniPage, int, error) {
	path := "/v3.0/tools/mini_page/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []MiniPage `json:"list"`
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

// CreateMiniPage 创建微页面
func (s *SiteClient) CreateMiniPage(ctx context.Context, accessToken string, advertiserID uint64, pageName, pagePath, appID string) (uint64, error) {
	path := "/v3.0/tools/mini_page/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page_name":     pageName,
		"page_path":     pagePath,
		"app_id":        appID,
	}
	var result struct {
		Data struct {
			PageID uint64 `json:"page_id"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.PageID, nil
}

// DeleteMiniPage 删除微页面
func (s *SiteClient) DeleteMiniPage(ctx context.Context, accessToken string, advertiserID, pageID uint64) error {
	path := "/v3.0/tools/mini_page/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page_id":       pageID,
	}
	return s.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 表单管理 ====================

// SiteForm 落地页表单
type SiteForm struct {
	FormID     uint64   `json:"form_id"`
	FormName   string   `json:"form_name"`
	SiteID     uint64   `json:"site_id"`
	Fields     []string `json:"fields"`
	SubmitNum  int      `json:"submit_num"`
	CreateTime string   `json:"create_time"`
}

// GetSiteFormList 获取落地页表单列表
func (s *SiteClient) GetSiteFormList(ctx context.Context, accessToken string, advertiserID, siteID uint64, page, pageSize int) ([]SiteForm, int, error) {
	path := "/v3.0/tools/site_form/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"site_id":       siteID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []SiteForm `json:"list"`
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

// FormSubmission 表单提交记录
type FormSubmission struct {
	SubmissionID uint64                 `json:"submission_id"`
	FormID       uint64                 `json:"form_id"`
	FormData     map[string]interface{} `json:"form_data"`
	SubmitTime   string                 `json:"submit_time"`
	SourceAdID   uint64                 `json:"source_ad_id"`
	SourceIP     string                 `json:"source_ip"`
}

// GetFormSubmissionList 获取表单提交记录
func (s *SiteClient) GetFormSubmissionList(ctx context.Context, accessToken string, advertiserID, formID uint64, startDate, endDate string, page, pageSize int) ([]FormSubmission, int, error) {
	path := "/v3.0/tools/form_submission/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"form_id":       formID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []FormSubmission `json:"list"`
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

// ExportFormSubmissions 导出表单提交数据
func (s *SiteClient) ExportFormSubmissions(ctx context.Context, accessToken string, advertiserID, formID uint64, startDate, endDate string) (string, error) {
	path := "/v3.0/tools/form_submission/export/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"form_id":       formID,
		"start_date":    startDate,
		"end_date":      endDate,
	}
	var result struct {
		Data struct {
			DownloadURL string `json:"download_url"`
		} `json:"data"`
	}
	err := s.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.DownloadURL, nil
}
