package api

import (
	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// NotificationAPI 消息通知 API
type NotificationAPI struct {
	notificationService *service.NotificationService
}

// NewNotificationAPI 创建消息通知 API
func NewNotificationAPI(notificationService *service.NotificationService) *NotificationAPI {
	return &NotificationAPI{notificationService: notificationService}
}

// GetList godoc
// @Summary 获取通知列表
// @Tags 系统管理-通知
// @Produce json
// @Param type query string false "通知类型"
// @Param is_read query bool false "是否已读"
// @Param keyword query string false "关键词"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.NotificationResp}
// @Router /api/v1/system/notifications [get]
func (a *NotificationAPI) GetList(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.NotificationListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.notificationService.GetList(c.Request.Context(), userID, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetStats godoc
// @Summary 获取通知统计
// @Tags 系统管理-通知
// @Produce json
// @Success 200 {object} response.Response{data=dto.NotificationStatsResp}
// @Router /api/v1/system/notifications/stats [get]
func (a *NotificationAPI) GetStats(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	stats, err := a.notificationService.GetStats(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, stats)
}

// MarkAsRead godoc
// @Summary 标记为已读
// @Tags 系统管理-通知
// @Accept json
// @Produce json
// @Param data body dto.NotificationMarkReadReq true "通知ID列表"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notifications/read [post]
func (a *NotificationAPI) MarkAsRead(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.NotificationMarkReadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.notificationService.MarkAsRead(c.Request.Context(), userID, req.IDs); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// MarkAllAsRead godoc
// @Summary 标记全部已读
// @Tags 系统管理-通知
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/notifications/read-all [post]
func (a *NotificationAPI) MarkAllAsRead(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	if err := a.notificationService.MarkAllAsRead(c.Request.Context(), userID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// Delete godoc
// @Summary 删除通知
// @Tags 系统管理-通知
// @Accept json
// @Produce json
// @Param data body dto.NotificationMarkReadReq true "通知ID列表"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notifications [delete]
func (a *NotificationAPI) Delete(c *gin.Context) {
	userID := c.GetUint64("userID")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.NotificationMarkReadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.notificationService.Delete(c.Request.Context(), userID, req.IDs); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
