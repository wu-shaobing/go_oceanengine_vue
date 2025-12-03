package api

import (
	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// OperationLogAPI 操作日志API
type OperationLogAPI struct {
	logService *service.OperationLogService
}

// NewOperationLogAPI 创建操作日志API
func NewOperationLogAPI(logService *service.OperationLogService) *OperationLogAPI {
	return &OperationLogAPI{logService: logService}
}

// GetList godoc
// @Summary 获取操作日志列表
// @Tags 系统管理-操作日志
// @Produce json
// @Param user_id query int false "用户ID"
// @Param username query string false "用户名"
// @Param module query string false "模块"
// @Param action query string false "操作"
// @Param status query int false "状态"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.OperationLogResp}
// @Router /api/v1/system/logs/operation [get]
func (a *OperationLogAPI) GetList(c *gin.Context) {
	var req dto.OperationLogListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.logService.GetList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetModules godoc
// @Summary 获取模块列表
// @Tags 系统管理-操作日志
// @Produce json
// @Success 200 {object} response.Response{data=[]string}
// @Router /api/v1/system/logs/modules [get]
func (a *OperationLogAPI) GetModules(c *gin.Context) {
	modules, err := a.logService.GetModules(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, modules)
}

// Delete godoc
// @Summary 删除操作日志
// @Tags 系统管理-操作日志
// @Accept json
// @Produce json
// @Param data body dto.OperationLogDeleteReq true "删除条件"
// @Success 200 {object} response.Response
// @Router /api/v1/system/logs/operation [delete]
func (a *OperationLogAPI) Delete(c *gin.Context) {
	var req dto.OperationLogDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.logService.Delete(c.Request.Context(), req.BeforeTime); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
