package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/media/dto"
	"oceanengine-backend/internal/app/media/service"
	"oceanengine-backend/pkg/response"
)

// MediaAPI 素材管理API
type MediaAPI struct {
	mediaService *service.MediaService
}

// NewMediaAPI 创建素材管理API
func NewMediaAPI(mediaService *service.MediaService) *MediaAPI {
	return &MediaAPI{mediaService: mediaService}
}

// GetImageList godoc
// @Summary 获取图片列表
// @Tags 素材管理
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param keyword query string false "关键词"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.ImageListResp}
// @Router /api/v1/media/images [get]
func (a *MediaAPI) GetImageList(c *gin.Context) {
	var req dto.ImageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.mediaService.GetImageList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetImageByID godoc
// @Summary 获取图片详情
// @Tags 素材管理
// @Produce json
// @Param id path int true "图片ID"
// @Success 200 {object} response.Response{data=model.MaterialImage}
// @Router /api/v1/media/images/{id} [get]
func (a *MediaAPI) GetImageByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	image, err := a.mediaService.GetImageByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, image)
}

// DeleteImage godoc
// @Summary 删除图片
// @Tags 素材管理
// @Produce json
// @Param id path int true "图片ID"
// @Success 200 {object} response.Response
// @Router /api/v1/media/images/{id} [delete]
func (a *MediaAPI) DeleteImage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.mediaService.DeleteImage(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// GetVideoList godoc
// @Summary 获取视频列表
// @Tags 素材管理
// @Produce json
// @Param advertiser_id query int true "广告主ID"
// @Param keyword query string false "关键词"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.VideoListResp}
// @Router /api/v1/media/videos [get]
func (a *MediaAPI) GetVideoList(c *gin.Context) {
	var req dto.VideoListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.mediaService.GetVideoList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetVideoByID godoc
// @Summary 获取视频详情
// @Tags 素材管理
// @Produce json
// @Param id path int true "视频ID"
// @Success 200 {object} response.Response{data=model.MaterialVideo}
// @Router /api/v1/media/videos/{id} [get]
func (a *MediaAPI) GetVideoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	video, err := a.mediaService.GetVideoByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, video)
}

// DeleteVideo godoc
// @Summary 删除视频
// @Tags 素材管理
// @Produce json
// @Param id path int true "视频ID"
// @Success 200 {object} response.Response
// @Router /api/v1/media/videos/{id} [delete]
func (a *MediaAPI) DeleteVideo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.mediaService.DeleteVideo(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// UploadImage godoc
// @Summary 上传图片素材
// @Tags 素材管理
// @Accept multipart/form-data
// @Produce json
// @Param advertiser_id formData int true "广告主ID"
// @Param file formData file true "图片文件"
// @Success 200 {object} response.Response{data=dto.ImageUploadResp}
// @Router /api/v1/media/images/upload [post]
func (a *MediaAPI) UploadImage(c *gin.Context) {
	advIDStr := c.PostForm("advertiser_id")
	advID, err := strconv.ParseUint(advIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "advertiser_id is required")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "file is required")
		return
	}
	defer file.Close()

	result, err := a.mediaService.UploadImage(c.Request.Context(), advID, header.Filename, file, header.Size)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result)
}

// UploadVideo godoc
// @Summary 上传视频素材
// @Tags 素材管理
// @Accept multipart/form-data
// @Produce json
// @Param advertiser_id formData int true "广告主ID"
// @Param file formData file true "视频文件"
// @Success 200 {object} response.Response{data=dto.VideoUploadResp}
// @Router /api/v1/media/videos/upload [post]
func (a *MediaAPI) UploadVideo(c *gin.Context) {
	advIDStr := c.PostForm("advertiser_id")
	advID, err := strconv.ParseUint(advIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "advertiser_id is required")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "file is required")
		return
	}
	defer file.Close()

	result, err := a.mediaService.UploadVideo(c.Request.Context(), advID, header.Filename, file, header.Size)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result)
}
