package model

import (
	"time"

	"gorm.io/gorm"
)

// User 系统用户表
type User struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"size:64;uniqueIndex" json:"username"`
	Password    string         `gorm:"size:128" json:"-"`
	Nickname    string         `gorm:"size:128" json:"nickname"`
	Phone       string         `gorm:"size:20;index" json:"phone"`
	Email       string         `gorm:"size:128;index" json:"email"`
	Avatar      string         `gorm:"size:255" json:"avatar"`
	Status      int8           `gorm:"default:1;index" json:"status"` // 0-禁用，1-启用
	RoleID      uint64         `gorm:"index" json:"role_id"`
	DeptID      uint64         `gorm:"default:0" json:"dept_id"`
	LastLoginAt *time.Time     `json:"last_login_at"`
	LastLoginIP string         `gorm:"size:50" json:"last_login_ip"`
	Remark      string         `gorm:"size:500" json:"remark"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy   uint64         `gorm:"default:0" json:"created_by"`
	UpdatedBy   uint64         `gorm:"default:0" json:"updated_by"`

	// 关联
	Role *Role `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}

// TableName 表名
func (User) TableName() string {
	return "sys_user"
}

// Role 角色表
type Role struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:64;not null" json:"name"`
	Key       string         `gorm:"size:64;uniqueIndex" json:"key"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int8           `gorm:"default:1;index" json:"status"`
	DataScope int8           `gorm:"default:1" json:"data_scope"` // 1-全部,2-自定义,3-本部门,4-本部门及以下,5-仅本人
	Remark    string         `gorm:"size:500" json:"remark"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy uint64         `gorm:"default:0" json:"created_by"`
	UpdatedBy uint64         `gorm:"default:0" json:"updated_by"`

	// 关联
	Menus []*Menu `gorm:"many2many:sys_role_menu" json:"menus,omitempty"`
}

// TableName 表名
func (Role) TableName() string {
	return "sys_role"
}

// Menu 菜单表
type Menu struct {
	ID         uint64         `gorm:"primaryKey" json:"id"`
	ParentID   uint64         `gorm:"default:0;index" json:"parent_id"`
	Name       string         `gorm:"size:64;not null" json:"name"`
	Path       string         `gorm:"size:255" json:"path"`
	Component  string         `gorm:"size:255" json:"component"`
	Icon       string         `gorm:"size:64" json:"icon"`
	Sort       int            `gorm:"default:0" json:"sort"`
	Permission string         `gorm:"size:128" json:"permission"`
	Type       int8           `gorm:"default:1" json:"type"` // 1-目录,2-菜单,3-按钮
	Visible    int8           `gorm:"default:1" json:"visible"`
	Status     int8           `gorm:"default:1" json:"status"`
	IsFrame    int8           `gorm:"default:0" json:"is_frame"`
	IsCache    int8           `gorm:"default:0" json:"is_cache"`
	Remark     string         `gorm:"size:500" json:"remark"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 子菜单
	Children []*Menu `gorm:"-" json:"children,omitempty"`
}

// TableName 表名
func (Menu) TableName() string {
	return "sys_menu"
}

// RoleMenu 角色菜单关联表
type RoleMenu struct {
	ID     uint64 `gorm:"primaryKey"`
	RoleID uint64 `gorm:"index;not null"`
	MenuID uint64 `gorm:"index;not null"`
}

// TableName 表名
func (RoleMenu) TableName() string {
	return "sys_role_menu"
}

// OperationLog 操作日志表
type OperationLog struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"index;default:0" json:"user_id"`
	Username  string    `gorm:"size:64" json:"username"`
	Module    string    `gorm:"size:64;index" json:"module"`
	Action    string    `gorm:"size:64" json:"action"`
	Method    string    `gorm:"size:16" json:"method"`
	Path      string    `gorm:"size:255" json:"path"`
	Query     string    `gorm:"type:text" json:"query"`
	Body      string    `gorm:"type:text" json:"body"`
	Response  string    `gorm:"type:text" json:"response"`
	IP        string    `gorm:"size:50" json:"ip"`
	UserAgent string    `gorm:"size:500" json:"user_agent"`
	Status    int       `gorm:"default:0" json:"status"`
	Latency   int64     `gorm:"default:0" json:"latency"` // 毫秒
	ErrorMsg  string    `gorm:"type:text" json:"error_msg"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
}

// TableName 表名
func (OperationLog) TableName() string {
	return "sys_operation_log"
}

// 用户状态常量
const (
	UserStatusDisabled = 0
	UserStatusEnabled  = 1
	UserStatusLocked   = 2
)

// 菜单类型常量
const (
	MenuTypeDir    = 1 // 目录
	MenuTypeMenu   = 2 // 菜单
	MenuTypeButton = 3 // 按钮
)
