package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// DPAHandler DPA商品广告处理器
type DPAHandler struct {
	db       *gorm.DB
	oceanCfg *config.OceanConfig
	client   *oceanengine.Client
}

// NewDPAHandler 创建DPA处理器
func NewDPAHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *DPAHandler {
	return &DPAHandler{
		db:       db,
		oceanCfg: oceanCfg,
		client:   oceanengine.NewClient(oceanCfg.AppID, oceanCfg.Secret),
	}
}

// getAccessToken 从请求中获取 access_token
func (h *DPAHandler) getAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.GetHeader("Access-Token")
	}
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// getAdvertiserID 从请求中获取 advertiser_id
func (h *DPAHandler) getAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// ==================== 商品库管理 ====================

// GetProductLibraryList 获取商品库列表
func (h *DPAHandler) GetProductLibraryList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.DPA().GetProductLibraryList(c.Request.Context(), accessToken, advertiserID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateProductLibrary 创建商品库
func (h *DPAHandler) CreateProductLibrary(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		LibraryName  string `json:"library_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryName == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	libraryID, err := h.client.DPA().CreateProductLibrary(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryName)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"library_id": libraryID})
}

// UpdateProductLibrary 更新商品库
func (h *DPAHandler) UpdateProductLibrary(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	libraryIDStr := c.Param("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		LibraryName  string `json:"library_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || libraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().UpdateProductLibrary(c.Request.Context(), accessToken, req.AdvertiserID, libraryID, req.LibraryName)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// DeleteProductLibrary 删除商品库
func (h *DPAHandler) DeleteProductLibrary(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	libraryIDStr := c.Param("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)
	advertiserID := h.getAdvertiserID(c)

	if accessToken == "" || advertiserID == 0 || libraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().DeleteProductLibrary(c.Request.Context(), accessToken, advertiserID, libraryID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 商品管理 ====================

// GetProductList 获取商品列表
func (h *DPAHandler) GetProductList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || libraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.DPA().GetProductList(c.Request.Context(), accessToken, advertiserID, libraryID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateProduct 创建商品
func (h *DPAHandler) CreateProduct(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.DPAProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	productID, err := h.client.DPA().CreateProduct(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"product_id": productID})
}

// UpdateProduct 更新商品
func (h *DPAHandler) UpdateProduct(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	productIDStr := c.Param("product_id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 64)

	var req oceanengine.DPAProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ProductID = productID

	if accessToken == "" || req.AdvertiserID == 0 || productID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().UpdateProduct(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// DeleteProduct 删除商品
func (h *DPAHandler) DeleteProduct(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	productIDStr := c.Param("product_id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 64)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || libraryID == 0 || productID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().DeleteProduct(c.Request.Context(), accessToken, advertiserID, libraryID, productID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// BatchDeleteProducts 批量删除商品
func (h *DPAHandler) BatchDeleteProducts(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64   `json:"advertiser_id"`
		LibraryID    uint64   `json:"library_id"`
		ProductIDs   []uint64 `json:"product_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 || len(req.ProductIDs) == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().BatchDeleteProducts(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryID, req.ProductIDs)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 商品分类管理 ====================

// GetProductCategoryList 获取商品分类列表
func (h *DPAHandler) GetProductCategoryList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || libraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, err := h.client.DPA().GetProductCategoryList(c.Request.Context(), accessToken, advertiserID, libraryID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"list": list})
}

// CreateProductCategory 创建商品分类
func (h *DPAHandler) CreateProductCategory(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		LibraryID    uint64 `json:"library_id"`
		CategoryName string `json:"category_name"`
		ParentID     uint64 `json:"parent_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 || req.CategoryName == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	categoryID, err := h.client.DPA().CreateProductCategory(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryID, req.CategoryName, req.ParentID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"category_id": categoryID})
}

// UpdateProductCategory 更新商品分类
func (h *DPAHandler) UpdateProductCategory(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	categoryIDStr := c.Param("category_id")
	categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 64)

	var req struct {
		AdvertiserID uint64 `json:"advertiser_id"`
		LibraryID    uint64 `json:"library_id"`
		CategoryName string `json:"category_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 || categoryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().UpdateProductCategory(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryID, categoryID, req.CategoryName)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// DeleteProductCategory 删除商品分类
func (h *DPAHandler) DeleteProductCategory(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	categoryIDStr := c.Param("category_id")
	categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 64)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || libraryID == 0 || categoryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().DeleteProductCategory(c.Request.Context(), accessToken, advertiserID, libraryID, categoryID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== 商品集管理 ====================

// GetProductSetList 获取商品集列表
func (h *DPAHandler) GetProductSetList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || libraryID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.DPA().GetProductSetList(c.Request.Context(), accessToken, advertiserID, libraryID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateProductSet 创建商品集
func (h *DPAHandler) CreateProductSet(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req struct {
		AdvertiserID uint64                      `json:"advertiser_id"`
		LibraryID    uint64                      `json:"library_id"`
		SetName      string                      `json:"set_name"`
		Filters      []oceanengine.ProductFilter `json:"filters"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 || req.SetName == "" {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	setID, err := h.client.DPA().CreateProductSet(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryID, req.SetName, req.Filters)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"set_id": setID})
}

// UpdateProductSet 更新商品集
func (h *DPAHandler) UpdateProductSet(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	setIDStr := c.Param("set_id")
	setID, _ := strconv.ParseUint(setIDStr, 10, 64)

	var req struct {
		AdvertiserID uint64                      `json:"advertiser_id"`
		LibraryID    uint64                      `json:"library_id"`
		SetName      string                      `json:"set_name"`
		Filters      []oceanengine.ProductFilter `json:"filters"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.LibraryID == 0 || setID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().UpdateProductSet(c.Request.Context(), accessToken, req.AdvertiserID, req.LibraryID, setID, req.SetName, req.Filters)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// DeleteProductSet 删除商品集
func (h *DPAHandler) DeleteProductSet(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	setIDStr := c.Param("set_id")
	setID, _ := strconv.ParseUint(setIDStr, 10, 64)
	advertiserID := h.getAdvertiserID(c)
	libraryIDStr := c.Query("library_id")
	libraryID, _ := strconv.ParseUint(libraryIDStr, 10, 64)

	if accessToken == "" || advertiserID == 0 || libraryID == 0 || setID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	err := h.client.DPA().DeleteProductSet(c.Request.Context(), accessToken, advertiserID, libraryID, setID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c)
}

// ==================== DPA模板管理 ====================

// GetTemplateList 获取DPA模板列表
func (h *DPAHandler) GetTemplateList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	templateType := c.Query("template_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.DPA().GetTemplateList(c.Request.Context(), accessToken, advertiserID, templateType, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// ==================== DPA创意管理 ====================

// GetDPACreativeList 获取DPA创意列表
func (h *DPAHandler) GetDPACreativeList(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	advertiserID := h.getAdvertiserID(c)
	adIDStr := c.Query("ad_id")
	adID, _ := strconv.ParseUint(adIDStr, 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if accessToken == "" || advertiserID == 0 || adID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	list, total, err := h.client.DPA().GetDPACreativeList(c.Request.Context(), accessToken, advertiserID, adID, page, pageSize)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithList(c, list, int64(total), page, pageSize)
}

// CreateDPACreative 创建DPA创意
func (h *DPAHandler) CreateDPACreative(c *gin.Context) {
	accessToken := h.getAccessToken(c)

	var req oceanengine.DPACreativeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if accessToken == "" || req.AdvertiserID == 0 || req.AdID == 0 {
		response.BadRequest(c, "缺少必要参数")
		return
	}

	creativeID, err := h.client.DPA().CreateDPACreative(c.Request.Context(), accessToken, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKWithData(c, gin.H{"creative_id": creativeID})
}
