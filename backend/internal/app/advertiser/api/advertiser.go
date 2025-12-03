package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/internal/app/advertiser/dto"
	"oceanengine-backend/internal/app/advertiser/service"
	"oceanengine-backend/pkg/database"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oauth"
	"oceanengine-backend/pkg/response"
	"oceanengine-backend/pkg/utils"
)

// AdvertiserHandler 广告主处理器
type AdvertiserHandler struct {
	service      *service.AdvertiserService
	stateManager *oauth.StateManager
}

// NewAdvertiserHandler 创建广告主处理器
func NewAdvertiserHandler(db *gorm.DB, oceanCfg *config.OceanConfig) *AdvertiserHandler {
	return &AdvertiserHandler{
		service:      service.NewAdvertiserService(db, oceanCfg),
		stateManager: oauth.NewStateManager(database.GetRedis()),
	}
}

// List 获取广告主列表
// @Summary 获取广告主列表
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param name query string false "广告主名称"
// @Param company query string false "公司名称"
// @Param status query string false "状态"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.AdvertiserListResp}}
// @Router /api/v1/advertisers [get]
func (h *AdvertiserHandler) List(c *gin.Context) {
	var req dto.AdvertiserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	// 设置默认分页
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	list, total, err := h.service.GetList(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithList(c, list, total, req.Page, req.PageSize)
}

// Get 获取广告主详情
// @Summary 获取广告主详情
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Success 200 {object} response.Response{data=dto.AdvertiserDetailResp}
// @Router /api/v1/advertisers/{id} [get]
func (h *AdvertiserHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// Update 更新广告主
// @Summary 更新广告主
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Param body body dto.AdvertiserUpdateReq true "更新内容"
// @Success 200 {object} response.Response
// @Router /api/v1/advertisers/{id} [put]
func (h *AdvertiserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	var req dto.AdvertiserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if err := h.service.Update(c.Request.Context(), id, &req); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c)
}

// Delete 删除广告主
// @Summary 删除广告主
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Success 200 {object} response.Response
// @Router /api/v1/advertisers/{id} [delete]
func (h *AdvertiserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c)
}

// Sync 同步广告主数据
// @Summary 同步广告主数据
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Success 200 {object} response.Response{data=dto.AdvertiserSyncResp}
// @Router /api/v1/advertisers/{id}/sync [post]
func (h *AdvertiserHandler) Sync(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.Sync(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetBalance 获取广告主余额
// @Summary 获取广告主余额
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Success 200 {object} response.Response{data=dto.AdvertiserBalanceResp}
// @Router /api/v1/advertisers/{id}/balance [get]
func (h *AdvertiserHandler) GetBalance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	data, err := h.service.GetBalance(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, data)
}

// GetFundList 获取资金流水列表
// @Summary 获取资金流水列表
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param id path int true "广告主ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.ListData{list=[]dto.FundListResp}}
// @Router /api/v1/advertisers/{id}/funds [get]
func (h *AdvertiserHandler) GetFundList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	var req dto.FundListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, errcode.New(errcode.ErrInvalidParams))
		return
	}

	req.AdvertiserID = id
	pagination := utils.NewPagination(req.Page, req.PageSize)
	req.Page = pagination.Page
	req.PageSize = pagination.PageSize

	list, total, err := h.service.GetFundList(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithList(c, list, total, req.Page, req.PageSize)
}

// GetOAuthURL 获取 OAuth 授权 URL
// @Summary 获取 OAuth 授权 URL
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param redirect_url query string false "授权成功后的跳转地址"
// @Success 200 {object} response.Response{data=dto.OAuthURLResp}
// @Router /api/v1/advertisers/oauth/url [get]
func (h *AdvertiserHandler) GetOAuthURL(c *gin.Context) {
	// 获取自定义跳转地址
	redirectURL := c.Query("redirect_url")

	// 生成并保存 state（防CSRF攻击）
	stateData := map[string]string{
		"redirect_url": redirectURL,
		"source":       "ocean", // 标识来源为巨量广告
	}

	state, err := h.stateManager.GenerateAndSave(c.Request.Context(), stateData)
	if err != nil {
		response.InternalError(c, "生成授权状态失败: "+err.Error())
		return
	}

	url := h.service.GetOAuthAuthorizeURL(state)

	response.OKWithData(c, dto.OAuthURLResp{
		AuthURL: url,
		State:   state,
	})
}

// OAuthCallback 处理 OAuth 回调
// @Summary 处理 OAuth 回调
// @Tags 广告主管理
// @Accept json
// @Produce json
// @Param auth_code query string true "授权码"
// @Param state query string true "状态码"
// @Success 200 {object} response.Response
// @Router /api/v1/advertisers/oauth/callback [get]
func (h *AdvertiserHandler) OAuthCallback(c *gin.Context) {
	authCode := c.Query("auth_code")
	state := c.Query("state")

	if authCode == "" {
		response.BadRequest(c, "缺少授权码")
		return
	}

	if state == "" {
		response.BadRequest(c, "缺少状态参数")
		return
	}

	// 验证 state 参数（防CSRF攻击）
	_, err := h.stateManager.ValidateState(c.Request.Context(), state)
	if err != nil {
		response.BadRequest(c, "无效或过期的授权请求")
		return
	}

	if err := h.service.HandleOAuthCallback(c.Request.Context(), authCode); err != nil {
		response.Fail(c, err)
		return
	}

	response.OKWithData(c, map[string]string{
		"message": "授权成功",
	})
}
