package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/database"
	"oceanengine-backend/pkg/oauth"
	"oceanengine-backend/pkg/oceanengine"
	"oceanengine-backend/pkg/response"
)

// QianchuanOAuthHandler 千川OAuth处理器
type QianchuanOAuthHandler struct {
	cfg          *config.QianchuanConfig
	stateManager *oauth.StateManager
	client       *oceanengine.Client
}

// NewQianchuanOAuthHandler 创建千川OAuth处理器
func NewQianchuanOAuthHandler(cfg *config.QianchuanConfig, redisClient *redis.Client) *QianchuanOAuthHandler {
	return &QianchuanOAuthHandler{
		cfg:          cfg,
		stateManager: oauth.NewStateManager(redisClient),
		client:       oceanengine.NewClient(cfg.AppID, cfg.Secret),
	}
}

// NewQianchuanOAuthHandlerDefault 使用默认Redis创建
func NewQianchuanOAuthHandlerDefault(cfg *config.QianchuanConfig) *QianchuanOAuthHandler {
	return NewQianchuanOAuthHandler(cfg, database.GetRedis())
}

// GetAuthURLResponse 授权URL响应
type GetAuthURLResponse struct {
	AuthURL string `json:"auth_url"`
	State   string `json:"state"`
}

// GetAuthURL 获取千川授权URL
// @Summary 获取千川授权URL
// @Tags 千川OAuth
// @Accept json
// @Produce json
// @Param redirect_url query string false "授权成功后的跳转地址"
// @Success 200 {object} response.Response{data=GetAuthURLResponse}
// @Router /api/v1/qianchuan/oauth/url [get]
func (h *QianchuanOAuthHandler) GetAuthURL(c *gin.Context) {
	// 获取自定义跳转地址
	redirectURL := c.Query("redirect_url")

	// 生成并保存 state
	stateData := map[string]string{
		"redirect_url": redirectURL,
	}

	state, err := h.stateManager.GenerateAndSave(c.Request.Context(), stateData)
	if err != nil {
		response.InternalError(c, "生成授权状态失败: "+err.Error())
		return
	}

	// 构建授权URL
	authURL := h.buildAuthURL(state)

	response.OKWithData(c, GetAuthURLResponse{
		AuthURL: authURL,
		State:   state,
	})
}

// buildAuthURL 构建千川授权URL
func (h *QianchuanOAuthHandler) buildAuthURL(state string) string {
	// 官方千川授权URL格式:
	// https://qianchuan.jinritemai.com/openapi/qc/audit/oauth.html?app_id=xxx&state=xxx&material_auth=1&rid=xxx
	baseURL := h.cfg.AuthURL
	if baseURL == "" {
		baseURL = "https://qianchuan.jinritemai.com/openapi/qc/audit/oauth.html"
	}

	params := url.Values{}
	params.Set("app_id", h.cfg.AppID)
	params.Set("state", state)

	// 启用素材授权
	if h.cfg.MaterialAuth {
		params.Set("material_auth", "1")
	}

	// 设置回调地址
	if h.cfg.RedirectURI != "" {
		params.Set("redirect_uri", h.cfg.RedirectURI)
	}

	return baseURL + "?" + params.Encode()
}

// OAuthCallback 处理OAuth回调
// @Summary 处理千川OAuth回调
// @Tags 千川OAuth
// @Accept json
// @Produce json
// @Param auth_code query string true "授权码"
// @Param state query string true "状态参数"
// @Success 200 {object} response.Response
// @Router /api/v1/qianchuan/oauth/callback [get]
func (h *QianchuanOAuthHandler) OAuthCallback(c *gin.Context) {
	authCode := c.Query("auth_code")
	state := c.Query("state")

	// 检查必要参数
	if authCode == "" {
		h.redirectWithError(c, "missing_auth_code", "授权码缺失")
		return
	}

	if state == "" {
		h.redirectWithError(c, "missing_state", "状态参数缺失")
		return
	}

	// 验证 state
	stateData, err := h.stateManager.ValidateState(c.Request.Context(), state)
	if err != nil {
		h.redirectWithError(c, "invalid_state", "无效或过期的授权请求")
		return
	}

	// 使用授权码换取 access_token
	oauthService := oceanengine.NewOAuthService(h.client)
	tokenResp, err := oauthService.GetAccessToken(c.Request.Context(), authCode)
	if err != nil {
		h.redirectWithError(c, "token_error", "获取访问令牌失败: "+err.Error())
		return
	}

	// 保存授权信息到数据库
	if err := h.saveAuthorization(c, tokenResp); err != nil {
		h.redirectWithError(c, "save_error", "保存授权信息失败: "+err.Error())
		return
	}

	// 如果有自定义跳转地址，则跳转
	if redirectURL, ok := stateData["redirect_url"]; ok && redirectURL != "" {
		// 添加成功参数
		if parsedURL, err := url.Parse(redirectURL); err == nil {
			query := parsedURL.Query()
			query.Set("auth_result", "success")
			query.Set("advertiser_count", fmt.Sprintf("%d", len(tokenResp.AdvertiserIDs)))
			parsedURL.RawQuery = query.Encode()
			c.Redirect(http.StatusFound, parsedURL.String())
			return
		}
	}

	// 返回JSON响应
	response.OKWithData(c, gin.H{
		"message":          "授权成功",
		"advertiser_ids":   tokenResp.AdvertiserIDs,
		"advertiser_count": len(tokenResp.AdvertiserIDs),
	})
}

// redirectWithError 重定向到错误页面
func (h *QianchuanOAuthHandler) redirectWithError(c *gin.Context, errorCode, errorMsg string) {
	// 可以配置一个前端错误页面URL
	// 这里暂时返回JSON错误
	response.BadRequest(c, fmt.Sprintf("[%s] %s", errorCode, errorMsg))
}

// saveAuthorization 保存授权信息
func (h *QianchuanOAuthHandler) saveAuthorization(c *gin.Context, tokenResp *oceanengine.AccessTokenResponse) error {
	// TODO: 将授权信息保存到数据库
	// 可以复用 advertiser service 的逻辑
	// 这里只是示例，实际应该注入 service

	// 设置 Token 用于后续API调用
	h.client.SetAccessToken(tokenResp.AccessToken)

	// 获取各广告主的详细信息
	advService := oceanengine.NewAdvertiserService(h.client)
	for _, advID := range tokenResp.AdvertiserIDs {
		infos, err := advService.GetInfo(c.Request.Context(), []int64{advID})
		if err != nil {
			continue // 跳过获取失败的
		}
		if len(infos) == 0 {
			continue
		}

		// TODO: 保存到数据库
		_ = infos[0]
	}

	return nil
}

// RefreshTokenRequest 刷新Token请求
type RefreshTokenRequest struct {
	AdvertiserID int64  `json:"advertiser_id" binding:"required"`
	RefreshToken string `json:"refresh_token"` // 可选，如果不提供则从数据库获取
}

// RefreshToken 刷新Token
// @Summary 刷新千川访问令牌
// @Tags 千川OAuth
// @Accept json
// @Produce json
// @Param body body RefreshTokenRequest true "刷新Token请求"
// @Success 200 {object} response.Response
// @Router /api/v1/qianchuan/oauth/refresh [post]
func (h *QianchuanOAuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	refreshToken := req.RefreshToken

	// 如果没有提供refresh_token，从数据库获取
	if refreshToken == "" {
		// 从数据库查询
		var adv struct {
			RefreshToken string
		}
		if err := database.GetDB().Table("ad_advertiser").
			Select("refresh_token").
			Where("advertiser_id = ? AND deleted_at IS NULL", req.AdvertiserID).
			First(&adv).Error; err != nil {
			response.NotFound(c, "未找到该广告主")
			return
		}

		if adv.RefreshToken == "" {
			response.BadRequest(c, "该广告主没有可用的刷新令牌，请重新授权")
			return
		}
		refreshToken = adv.RefreshToken
	}

	// 调用API刷新Token
	oauthService := oceanengine.NewOAuthService(h.client)
	tokenResp, err := oauthService.RefreshAccessToken(c.Request.Context(), refreshToken)
	if err != nil {
		response.InternalError(c, "刷新令牌失败: "+err.Error())
		return
	}

	// 更新数据库中的Token
	newExpireAt := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	if err := database.GetDB().Table("ad_advertiser").
		Where("advertiser_id = ?", req.AdvertiserID).
		Updates(map[string]interface{}{
			"access_token":     tokenResp.AccessToken,
			"refresh_token":    tokenResp.RefreshToken,
			"token_expires_at": newExpireAt,
			"updated_at":       time.Now(),
		}).Error; err != nil {
		response.InternalError(c, "更新令牌失败: "+err.Error())
		return
	}

	response.OKWithData(c, gin.H{
		"message":    "刷新成功",
		"expires_in": tokenResp.ExpiresIn,
		"expire_at":  newExpireAt.Format(time.RFC3339),
	})
}
