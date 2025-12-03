package api

import (
	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// CaptchaAPI 验证码 API
type CaptchaAPI struct {
	captchaService *service.CaptchaService
}

// NewCaptchaAPI 创建验证码 API
func NewCaptchaAPI(captchaService *service.CaptchaService) *CaptchaAPI {
	return &CaptchaAPI{captchaService: captchaService}
}

// Get 获取验证码
// @Summary 获取验证码
// @Tags 认证
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=dto.CaptchaResp}
// @Router /api/v1/auth/captcha [get]
func (a *CaptchaAPI) Get(c *gin.Context) {
	id, b64s, err := a.captchaService.Generate()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, &dto.CaptchaResp{
		CaptchaID: id,
		Captcha:   b64s,
	})
}
