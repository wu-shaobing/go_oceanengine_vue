package handler

import (
	"net/http"

	"github.com/bububa/oceanengine/server/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler API处理器
type Handler struct {
	oceanEngine *service.OceanEngineService
}

// NewHandler 创建Handler
func NewHandler(oceanEngine *service.OceanEngineService) *Handler {
	return &Handler{
		oceanEngine: oceanEngine,
	}
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

// --- Health Check ---

// Health 健康检查
func (h *Handler) Health(c *gin.Context) {
	success(c, gin.H{
		"status": "healthy",
	})
}

// --- OAuth 相关 ---

// GetAuthURLRequest 获取授权链接请求
type GetAuthURLRequest struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
	State       string `json:"state"`
}

// GetAuthURL 获取OAuth授权链接
func (h *Handler) GetAuthURL(c *gin.Context) {
	var req GetAuthURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "参数错误: "+err.Error())
		return
	}

	url := h.oceanEngine.GetAuthURL(req.RedirectURL, req.State)
	success(c, gin.H{
		"auth_url": url,
	})
}

// GetAccessTokenRequest 获取AccessToken请求
type GetAccessTokenRequest struct {
	AuthCode string `json:"auth_code" binding:"required"`
}

// GetAccessToken 使用授权码获取AccessToken
func (h *Handler) GetAccessToken(c *gin.Context) {
	var req GetAccessTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "参数错误: "+err.Error())
		return
	}

	token, err := h.oceanEngine.GetAccessToken(c.Request.Context(), req.AuthCode)
	if err != nil {
		fail(c, 500, "获取Token失败: "+err.Error())
		return
	}

	success(c, token)
}

// RefreshTokenRequest 刷新Token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshToken 刷新Token
func (h *Handler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "参数错误: "+err.Error())
		return
	}

	token, err := h.oceanEngine.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		fail(c, 500, "刷新Token失败: "+err.Error())
		return
	}

	success(c, token)
}

// --- 广告主相关 ---

// GetAdvertisers 获取已授权广告主列表
func (h *Handler) GetAdvertisers(c *gin.Context) {
	accessToken := c.GetHeader("Access-Token")
	if accessToken == "" {
		fail(c, 401, "缺少Access-Token")
		return
	}

	advertisers, err := h.oceanEngine.GetAdvertisers(c.Request.Context(), accessToken)
	if err != nil {
		fail(c, 500, "获取广告主列表失败: "+err.Error())
		return
	}

	success(c, advertisers)
}

// GetAdvertiserInfoRequest 获取广告主信息请求
type GetAdvertiserInfoRequest struct {
	AdvertiserIDs []uint64 `json:"advertiser_ids" binding:"required"`
}

// GetAdvertiserInfo 获取广告主详细信息
func (h *Handler) GetAdvertiserInfo(c *gin.Context) {
	accessToken := c.GetHeader("Access-Token")
	if accessToken == "" {
		fail(c, 401, "缺少Access-Token")
		return
	}

	var req GetAdvertiserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "参数错误: "+err.Error())
		return
	}

	infos, err := h.oceanEngine.GetAdvertiserInfo(c.Request.Context(), accessToken, req.AdvertiserIDs)
	if err != nil {
		fail(c, 500, "获取广告主信息失败: "+err.Error())
		return
	}

	success(c, infos)
}
