package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/audience/dto"
	"oceanengine-backend/internal/app/audience/service"
	"oceanengine-backend/pkg/response"
)

// AudienceAPI 人群定向API
type AudienceAPI struct {
	audienceService *service.AudienceService
}

// NewAudienceAPI 创建人群定向API
func NewAudienceAPI(audienceService *service.AudienceService) *AudienceAPI {
	return &AudienceAPI{audienceService: audienceService}
}

// GetPackageList godoc
// @Summary 获取定向包列表
// @Tags 人群定向
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param keyword query string false "关键词"
// @Param landing_type query string false "推广类型"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.AudiencePackageListResp}
// @Router /api/v1/audiences/packages [get]
func (a *AudienceAPI) GetPackageList(c *gin.Context) {
	var req dto.AudiencePackageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.audienceService.GetPackageList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetPackageByID godoc
// @Summary 获取定向包详情
// @Tags 人群定向
// @Produce json
// @Param id path int true "定向包ID"
// @Success 200 {object} response.Response{data=model.AudiencePackage}
// @Router /api/v1/audiences/packages/{id} [get]
func (a *AudienceAPI) GetPackageByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	pkg, err := a.audienceService.GetPackageByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, pkg)
}

// CreatePackage godoc
// @Summary 创建定向包
// @Tags 人群定向
// @Accept json
// @Produce json
// @Param data body dto.AudiencePackageCreateReq true "定向包信息"
// @Success 200 {object} response.Response
// @Router /api/v1/audiences/packages [post]
func (a *AudienceAPI) CreatePackage(c *gin.Context) {
	var req dto.AudiencePackageCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.audienceService.CreatePackage(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// UpdatePackage godoc
// @Summary 更新定向包
// @Tags 人群定向
// @Accept json
// @Produce json
// @Param id path int true "定向包ID"
// @Param data body dto.AudiencePackageUpdateReq true "定向包信息"
// @Success 200 {object} response.Response
// @Router /api/v1/audiences/packages/{id} [put]
func (a *AudienceAPI) UpdatePackage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.AudiencePackageUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	if err := a.audienceService.UpdatePackage(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// DeletePackage godoc
// @Summary 删除定向包
// @Tags 人群定向
// @Produce json
// @Param id path int true "定向包ID"
// @Success 200 {object} response.Response
// @Router /api/v1/audiences/packages/{id} [delete]
func (a *AudienceAPI) DeletePackage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.audienceService.DeletePackage(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// GetCustomAudienceList godoc
// @Summary 获取自定义人群列表
// @Tags 人群定向
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param keyword query string false "关键词"
// @Param source query string false "来源类型"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.CustomAudienceListResp}
// @Router /api/v1/audiences/custom [get]
func (a *AudienceAPI) GetCustomAudienceList(c *gin.Context) {
	var req dto.CustomAudienceListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.audienceService.GetCustomAudienceList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetCustomAudienceByID godoc
// @Summary 获取自定义人群详情
// @Tags 人群定向
// @Produce json
// @Param id path int true "人群ID"
// @Success 200 {object} response.Response{data=model.CustomAudience}
// @Router /api/v1/audiences/custom/{id} [get]
func (a *AudienceAPI) GetCustomAudienceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	audience, err := a.audienceService.GetCustomAudienceByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, audience)
}

// DeleteCustomAudience godoc
// @Summary 删除自定义人群
// @Tags 人群定向
// @Produce json
// @Param id path int true "人群ID"
// @Success 200 {object} response.Response
// @Router /api/v1/audiences/custom/{id} [delete]
func (a *AudienceAPI) DeleteCustomAudience(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.audienceService.DeleteCustomAudience(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
