package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/service"
	"oceanengine-backend/pkg/response"
)

// DictAPI 字典管理 API
type DictAPI struct {
	dictService *service.DictService
}

// NewDictAPI 创建字典管理 API
func NewDictAPI(dictService *service.DictService) *DictAPI {
	return &DictAPI{dictService: dictService}
}

// ========== 字典类型 ==========

// GetTypeList 获取字典类型列表
// @Summary 获取字典类型列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param name query string false "字典名称"
// @Param type query string false "字典类型"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.DictTypeResp}
// @Router /api/v1/system/dict/types [get]
func (a *DictAPI) GetTypeList(c *gin.Context) {
	var req dto.DictTypeListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.dictService.GetTypeList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetTypeByID 获取字典类型详情
// @Summary 获取字典类型详情
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典类型ID"
// @Success 200 {object} response.Response{data=dto.DictTypeResp}
// @Router /api/v1/system/dict/types/{id} [get]
func (a *DictAPI) GetTypeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	data, err := a.dictService.GetTypeByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, data)
}

// CreateType 创建字典类型
// @Summary 创建字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param body body dto.DictTypeCreateReq true "字典类型信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/types [post]
func (a *DictAPI) CreateType(c *gin.Context) {
	var req dto.DictTypeCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.dictService.CreateType(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// UpdateType 更新字典类型
// @Summary 更新字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典类型ID"
// @Param body body dto.DictTypeUpdateReq true "字典类型信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/types/{id} [put]
func (a *DictAPI) UpdateType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.DictTypeUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	if err := a.dictService.UpdateType(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// DeleteType 删除字典类型
// @Summary 删除字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典类型ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/types/{id} [delete]
func (a *DictAPI) DeleteType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.dictService.DeleteType(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// ========== 字典数据 ==========

// GetDataList 获取字典数据列表
// @Summary 获取字典数据列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param dict_type query string true "字典类型"
// @Param label query string false "标签"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]dto.DictDataResp}
// @Router /api/v1/system/dict/data [get]
func (a *DictAPI) GetDataList(c *gin.Context) {
	var req dto.DictDataListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := a.dictService.GetDataList(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithPage(c, list, total, req.GetPage(), req.GetPageSize())
}

// GetDataByType 根据类型获取字典数据
// @Summary 根据类型获取字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param type path string true "字典类型"
// @Success 200 {object} response.Response{data=[]dto.DictDataResp}
// @Router /api/v1/system/dict/data/{type} [get]
func (a *DictAPI) GetDataByType(c *gin.Context) {
	dictType := c.Param("type")
	if dictType == "" {
		response.BadRequest(c, "dict_type is required")
		return
	}

	data, err := a.dictService.GetDataByType(c.Request.Context(), dictType)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, data)
}

// CreateData 创建字典数据
// @Summary 创建字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param body body dto.DictDataCreateReq true "字典数据信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/data [post]
func (a *DictAPI) CreateData(c *gin.Context) {
	var req dto.DictDataCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := a.dictService.CreateData(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// UpdateData 更新字典数据
// @Summary 更新字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据ID"
// @Param body body dto.DictDataUpdateReq true "字典数据信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/data/{id} [put]
func (a *DictAPI) UpdateData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req dto.DictDataUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ID = id

	if err := a.dictService.UpdateData(c.Request.Context(), &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}

// DeleteData 删除字典数据
// @Summary 删除字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dict/data/{id} [delete]
func (a *DictAPI) DeleteData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := a.dictService.DeleteData(c.Request.Context(), id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c)
}
