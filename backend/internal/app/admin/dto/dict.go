package dto

import "oceanengine-backend/pkg/utils"

// DictTypeListReq 字典类型列表请求
type DictTypeListReq struct {
	utils.Pagination
	Name   string `form:"name"`
	Type   string `form:"type"`
	Status *int8  `form:"status"`
}

// DictTypeCreateReq 创建字典类型请求
type DictTypeCreateReq struct {
	Name   string `json:"name" binding:"required,max=100"`
	Type   string `json:"type" binding:"required,max=100"`
	Status int8   `json:"status"`
	Remark string `json:"remark" binding:"max=500"`
}

// DictTypeUpdateReq 更新字典类型请求
type DictTypeUpdateReq struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" binding:"max=100"`
	Type   string `json:"type" binding:"max=100"`
	Status int8   `json:"status"`
	Remark string `json:"remark" binding:"max=500"`
}

// DictTypeResp 字典类型响应
type DictTypeResp struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Status    int8   `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
}

// DictDataListReq 字典数据列表请求
type DictDataListReq struct {
	utils.Pagination
	DictType string `form:"dict_type" binding:"required"`
	Label    string `form:"label"`
	Status   *int8  `form:"status"`
}

// DictDataCreateReq 创建字典数据请求
type DictDataCreateReq struct {
	DictType  string `json:"dict_type" binding:"required,max=100"`
	Label     string `json:"label" binding:"required,max=100"`
	Value     string `json:"value" binding:"required,max=100"`
	Sort      int    `json:"sort"`
	Status    int8   `json:"status"`
	IsDefault int8   `json:"is_default"`
	Remark    string `json:"remark" binding:"max=500"`
	CssClass  string `json:"css_class" binding:"max=100"`
	ListClass string `json:"list_class" binding:"max=100"`
}

// DictDataUpdateReq 更新字典数据请求
type DictDataUpdateReq struct {
	ID        uint64 `json:"id"`
	DictType  string `json:"dict_type" binding:"max=100"`
	Label     string `json:"label" binding:"max=100"`
	Value     string `json:"value" binding:"max=100"`
	Sort      int    `json:"sort"`
	Status    int8   `json:"status"`
	IsDefault int8   `json:"is_default"`
	Remark    string `json:"remark" binding:"max=500"`
	CssClass  string `json:"css_class" binding:"max=100"`
	ListClass string `json:"list_class" binding:"max=100"`
}

// DictDataResp 字典数据响应
type DictDataResp struct {
	ID        uint64 `json:"id"`
	DictType  string `json:"dict_type"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Sort      int    `json:"sort"`
	Status    int8   `json:"status"`
	IsDefault int8   `json:"is_default"`
	Remark    string `json:"remark"`
	CssClass  string `json:"css_class"`
	ListClass string `json:"list_class"`
	CreatedAt string `json:"created_at"`
}
