package oceanengine

import (
	"context"
)

// DPAClient DPA商品广告客户端
type DPAClient struct {
	client *Client
}

// NewDPAClient 创建DPA客户端
func NewDPAClient(client *Client) *DPAClient {
	return &DPAClient{client: client}
}

// ==================== 商品库管理 ====================

// ProductLibrary 商品库
type ProductLibrary struct {
	LibraryID   uint64 `json:"library_id"`
	LibraryName string `json:"library_name"`
	ProductNum  int    `json:"product_num"`
	Status      int    `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// CreateProductLibrary 创建商品库
func (d *DPAClient) CreateProductLibrary(ctx context.Context, accessToken string, advertiserID uint64, libraryName string) (uint64, error) {
	path := "/v3.0/dpa/product_library/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_name":  libraryName,
	}
	var result struct {
		Data struct {
			LibraryID uint64 `json:"library_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.LibraryID, nil
}

// GetProductLibraryList 获取商品库列表
func (d *DPAClient) GetProductLibraryList(ctx context.Context, accessToken string, advertiserID uint64, page, pageSize int) ([]ProductLibrary, int, error) {
	path := "/v3.0/dpa/product_library/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ProductLibrary `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// UpdateProductLibrary 更新商品库
func (d *DPAClient) UpdateProductLibrary(ctx context.Context, accessToken string, advertiserID, libraryID uint64, libraryName string) error {
	path := "/v3.0/dpa/product_library/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"library_name":  libraryName,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteProductLibrary 删除商品库
func (d *DPAClient) DeleteProductLibrary(ctx context.Context, accessToken string, advertiserID, libraryID uint64) error {
	path := "/v3.0/dpa/product_library/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 商品管理 ====================

// DPAProduct DPA商品
type DPAProduct struct {
	ProductID       uint64            `json:"product_id"`
	ProductOuterID  string            `json:"product_outer_id"`
	ProductName     string            `json:"product_name"`
	ProductPrice    float64           `json:"product_price"`
	ProductURL      string            `json:"product_url"`
	ProductImageURL string            `json:"product_image_url"`
	CategoryID      uint64            `json:"category_id"`
	CategoryName    string            `json:"category_name"`
	Status          int               `json:"status"`
	Attributes      map[string]string `json:"attributes"`
	CreateTime      string            `json:"create_time"`
	UpdateTime      string            `json:"update_time"`
}

// DPAProductCreateRequest 创建商品请求
type DPAProductCreateRequest struct {
	AdvertiserID    uint64            `json:"advertiser_id"`
	LibraryID       uint64            `json:"library_id"`
	ProductOuterID  string            `json:"product_outer_id"`
	ProductName     string            `json:"product_name"`
	ProductPrice    float64           `json:"product_price"`
	ProductURL      string            `json:"product_url"`
	ProductImageURL string            `json:"product_image_url"`
	CategoryID      uint64            `json:"category_id,omitempty"`
	Attributes      map[string]string `json:"attributes,omitempty"`
}

// CreateProduct 创建商品
func (d *DPAClient) CreateProduct(ctx context.Context, accessToken string, req *DPAProductCreateRequest) (uint64, error) {
	path := "/v3.0/dpa/product/create/"
	var result struct {
		Data struct {
			ProductID uint64 `json:"product_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.ProductID, nil
}

// GetProductList 获取商品列表
func (d *DPAClient) GetProductList(ctx context.Context, accessToken string, advertiserID, libraryID uint64, page, pageSize int) ([]DPAProduct, int, error) {
	path := "/v3.0/dpa/product/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []DPAProduct `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// DPAProductUpdateRequest 更新商品请求
type DPAProductUpdateRequest struct {
	AdvertiserID    uint64            `json:"advertiser_id"`
	LibraryID       uint64            `json:"library_id"`
	ProductID       uint64            `json:"product_id"`
	ProductName     string            `json:"product_name,omitempty"`
	ProductPrice    float64           `json:"product_price,omitempty"`
	ProductURL      string            `json:"product_url,omitempty"`
	ProductImageURL string            `json:"product_image_url,omitempty"`
	CategoryID      uint64            `json:"category_id,omitempty"`
	Attributes      map[string]string `json:"attributes,omitempty"`
}

// UpdateProduct 更新商品
func (d *DPAClient) UpdateProduct(ctx context.Context, accessToken string, req *DPAProductUpdateRequest) error {
	path := "/v3.0/dpa/product/update/"
	return d.client.PostWithToken(ctx, accessToken, path, req, nil)
}

// DeleteProduct 删除商品
func (d *DPAClient) DeleteProduct(ctx context.Context, accessToken string, advertiserID, libraryID, productID uint64) error {
	path := "/v3.0/dpa/product/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"product_id":    productID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// BatchDeleteProducts 批量删除商品
func (d *DPAClient) BatchDeleteProducts(ctx context.Context, accessToken string, advertiserID, libraryID uint64, productIDs []uint64) error {
	path := "/v3.0/dpa/product/batch_delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"product_ids":   productIDs,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 商品分类管理 ====================

// ProductCategory 商品分类
type ProductCategory struct {
	CategoryID   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     uint64 `json:"parent_id"`
	Level        int    `json:"level"`
	ProductNum   int    `json:"product_num"`
}

// GetProductCategoryList 获取商品分类列表
func (d *DPAClient) GetProductCategoryList(ctx context.Context, accessToken string, advertiserID, libraryID uint64) ([]ProductCategory, error) {
	path := "/v3.0/dpa/product_category/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
	}
	var result struct {
		Data struct {
			List []ProductCategory `json:"list"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// CreateProductCategory 创建商品分类
func (d *DPAClient) CreateProductCategory(ctx context.Context, accessToken string, advertiserID, libraryID uint64, categoryName string, parentID uint64) (uint64, error) {
	path := "/v3.0/dpa/product_category/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"category_name": categoryName,
	}
	if parentID > 0 {
		data["parent_id"] = parentID
	}
	var result struct {
		Data struct {
			CategoryID uint64 `json:"category_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.CategoryID, nil
}

// UpdateProductCategory 更新商品分类
func (d *DPAClient) UpdateProductCategory(ctx context.Context, accessToken string, advertiserID, libraryID, categoryID uint64, categoryName string) error {
	path := "/v3.0/dpa/product_category/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"category_id":   categoryID,
		"category_name": categoryName,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteProductCategory 删除商品分类
func (d *DPAClient) DeleteProductCategory(ctx context.Context, accessToken string, advertiserID, libraryID, categoryID uint64) error {
	path := "/v3.0/dpa/product_category/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"category_id":   categoryID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 商品圈选/筛选 ====================

// ProductFilter 商品筛选条件
type ProductFilter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// ProductSet 商品集
type ProductSet struct {
	SetID        uint64          `json:"set_id"`
	SetName      string          `json:"set_name"`
	LibraryID    uint64          `json:"library_id"`
	Filters      []ProductFilter `json:"filters"`
	ProductCount int             `json:"product_count"`
	CreateTime   string          `json:"create_time"`
}

// CreateProductSet 创建商品集
func (d *DPAClient) CreateProductSet(ctx context.Context, accessToken string, advertiserID, libraryID uint64, setName string, filters []ProductFilter) (uint64, error) {
	path := "/v3.0/dpa/product_set/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"set_name":      setName,
		"filters":       filters,
	}
	var result struct {
		Data struct {
			SetID uint64 `json:"set_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.SetID, nil
}

// GetProductSetList 获取商品集列表
func (d *DPAClient) GetProductSetList(ctx context.Context, accessToken string, advertiserID, libraryID uint64, page, pageSize int) ([]ProductSet, int, error) {
	path := "/v3.0/dpa/product_set/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ProductSet `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// UpdateProductSet 更新商品集
func (d *DPAClient) UpdateProductSet(ctx context.Context, accessToken string, advertiserID, libraryID, setID uint64, setName string, filters []ProductFilter) error {
	path := "/v3.0/dpa/product_set/update/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"set_id":        setID,
	}
	if setName != "" {
		data["set_name"] = setName
	}
	if len(filters) > 0 {
		data["filters"] = filters
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteProductSet 删除商品集
func (d *DPAClient) DeleteProductSet(ctx context.Context, accessToken string, advertiserID, libraryID, setID uint64) error {
	path := "/v3.0/dpa/product_set/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"set_id":        setID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== DPA模板管理 ====================

// DPATemplate DPA模板
type DPATemplate struct {
	TemplateID   uint64 `json:"template_id"`
	TemplateName string `json:"template_name"`
	TemplateType string `json:"template_type"`
	Status       int    `json:"status"`
	PreviewURL   string `json:"preview_url"`
	CreateTime   string `json:"create_time"`
}

// GetTemplateList 获取DPA模板列表
func (d *DPAClient) GetTemplateList(ctx context.Context, accessToken string, advertiserID uint64, templateType string, page, pageSize int) ([]DPATemplate, int, error) {
	path := "/v3.0/dpa/template/get/"
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
			List     []DPATemplate `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== DPA创意管理 ====================

// DPACreativeCreateRequest 创建DPA创意请求
type DPACreativeCreateRequest struct {
	AdvertiserID uint64   `json:"advertiser_id"`
	AdID         uint64   `json:"ad_id"`
	LibraryID    uint64   `json:"library_id"`
	SetID        uint64   `json:"set_id,omitempty"`
	TemplateID   uint64   `json:"template_id"`
	ProductIDs   []uint64 `json:"product_ids,omitempty"`
	Title        string   `json:"title,omitempty"`
}

// DPACreative DPA创意
type DPACreative struct {
	CreativeID uint64 `json:"creative_id"`
	AdID       uint64 `json:"ad_id"`
	LibraryID  uint64 `json:"library_id"`
	SetID      uint64 `json:"set_id"`
	TemplateID uint64 `json:"template_id"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// CreateDPACreative 创建DPA创意
func (d *DPAClient) CreateDPACreative(ctx context.Context, accessToken string, req *DPACreativeCreateRequest) (uint64, error) {
	path := "/v3.0/dpa/creative/create/"
	var result struct {
		Data struct {
			CreativeID uint64 `json:"creative_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.CreativeID, nil
}

// GetDPACreativeList 获取DPA创意列表
func (d *DPAClient) GetDPACreativeList(ctx context.Context, accessToken string, advertiserID uint64, adID uint64, page, pageSize int) ([]DPACreative, int, error) {
	path := "/v3.0/dpa/creative/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []DPACreative `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 数据同步 ====================

// SyncTask 同步任务
type SyncTask struct {
	TaskID     uint64 `json:"task_id"`
	LibraryID  uint64 `json:"library_id"`
	Status     int    `json:"status"`
	SourceType string `json:"source_type"`
	SourceURL  string `json:"source_url"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// CreateSyncTask 创建数据同步任务
func (d *DPAClient) CreateSyncTask(ctx context.Context, accessToken string, advertiserID, libraryID uint64, sourceType, sourceURL string) (uint64, error) {
	path := "/v3.0/dpa/sync_task/create/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"source_type":   sourceType,
		"source_url":    sourceURL,
	}
	var result struct {
		Data struct {
			TaskID uint64 `json:"task_id"`
		} `json:"data"`
	}
	err := d.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.TaskID, nil
}

// GetSyncTaskList 获取同步任务列表
func (d *DPAClient) GetSyncTaskList(ctx context.Context, accessToken string, advertiserID, libraryID uint64, page, pageSize int) ([]SyncTask, int, error) {
	path := "/v3.0/dpa/sync_task/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []SyncTask `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// TriggerSyncTask 触发同步任务
func (d *DPAClient) TriggerSyncTask(ctx context.Context, accessToken string, advertiserID, libraryID, taskID uint64) error {
	path := "/v3.0/dpa/sync_task/trigger/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"task_id":       taskID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteSyncTask 删除同步任务
func (d *DPAClient) DeleteSyncTask(ctx context.Context, accessToken string, advertiserID, libraryID, taskID uint64) error {
	path := "/v3.0/dpa/sync_task/delete/"
	data := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"task_id":       taskID,
	}
	return d.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 商品报表 ====================

// ProductReport 商品报表
type ProductReport struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Cost        float64 `json:"cost"`
	Show        int64   `json:"show"`
	Click       int64   `json:"click"`
	Convert     int64   `json:"convert"`
	CTR         float64 `json:"ctr"`
	CVR         float64 `json:"cvr"`
	CPC         float64 `json:"cpc"`
	CPM         float64 `json:"cpm"`
	CPA         float64 `json:"cpa"`
}

// GetProductReport 获取商品报表
func (d *DPAClient) GetProductReport(ctx context.Context, accessToken string, advertiserID, libraryID uint64, startDate, endDate string, page, pageSize int) ([]ProductReport, int, error) {
	path := "/v3.0/dpa/report/product/get/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"library_id":    libraryID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}
	var result struct {
		Data struct {
			List     []ProductReport `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := d.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}
