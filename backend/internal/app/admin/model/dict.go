package model

import (
	"time"

	"gorm.io/gorm"
)

// DictType 字典类型
type DictType struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:100;not null;comment:字典名称" json:"name"`
	Type      string         `gorm:"size:100;uniqueIndex;not null;comment:字典类型" json:"type"`
	Status    int8           `gorm:"default:1;comment:状态(0禁用 1启用)" json:"status"`
	Remark    string         `gorm:"size:500;comment:备注" json:"remark"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (DictType) TableName() string {
	return "sys_dict_type"
}

// DictData 字典数据
type DictData struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	DictType  string         `gorm:"size:100;index;not null;comment:字典类型" json:"dict_type"`
	Label     string         `gorm:"size:100;not null;comment:字典标签" json:"label"`
	Value     string         `gorm:"size:100;not null;comment:字典键值" json:"value"`
	Sort      int            `gorm:"default:0;comment:排序" json:"sort"`
	Status    int8           `gorm:"default:1;comment:状态(0禁用 1启用)" json:"status"`
	IsDefault int8           `gorm:"default:0;comment:是否默认(0否 1是)" json:"is_default"`
	Remark    string         `gorm:"size:500;comment:备注" json:"remark"`
	CssClass  string         `gorm:"size:100;comment:样式属性" json:"css_class"`
	ListClass string         `gorm:"size:100;comment:列表样式" json:"list_class"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (DictData) TableName() string {
	return "sys_dict_data"
}

// 字典状态常量
const (
	DictStatusDisabled = 0 // 禁用
	DictStatusEnabled  = 1 // 启用
)
