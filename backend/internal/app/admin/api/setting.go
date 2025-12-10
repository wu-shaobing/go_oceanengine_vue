package api

import (
	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// SettingAPI 用户设置 API
type SettingAPI struct {
	settingService *service.SettingService
}

// NewSettingAPI 创建用户设置 API
func NewSettingAPI(settingService *service.SettingService) *SettingAPI {
	return &SettingAPI{settingService: settingService}
}

// Get godoc
// @Summary 获取用户设置
// @Tags 系统管理-设置
// @Produce json
// @Success 200 {object} response.Response{data=dto.UserSettingResp}
// @Router /api/v1/system/settings [get]
func (a *SettingAPI) Get(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	setting, err := a.settingService.GetUserSetting(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, setting)
}

// Update godoc
// @Summary 更新用户设置
// @Tags 系统管理-设置
// @Accept json
// @Produce json
// @Param data body dto.UserSettingUpdateReq true "设置信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/settings [put]
func (a *SettingAPI) Update(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.UserSettingUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.settingService.UpdateUserSetting(c.Request.Context(), userID, &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
