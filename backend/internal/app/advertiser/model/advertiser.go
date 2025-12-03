package model

import (
	"time"

	"gorm.io/gorm"
)

// Advertiser 广告主表
type Advertiser struct {
	ID            uint64         `gorm:"primaryKey" json:"id"`
	AdvertiserID  uint64         `gorm:"uniqueIndex;not null" json:"advertiser_id"` // Ocean Engine 广告主ID
	Name          string         `gorm:"size:255;not null;index" json:"name"`
	Company       string         `gorm:"size:255" json:"company"`
	Status        string         `gorm:"size:50;index" json:"status"`
	Role          string         `gorm:"size:50" json:"role"`
	Balance       float64        `gorm:"type:decimal(15,2);default:0" json:"balance"`
	ValidBalance  float64        `gorm:"type:decimal(15,2);default:0" json:"valid_balance"`
	CashBalance   float64        `gorm:"type:decimal(15,2);default:0" json:"cash_balance"`
	Industry      string         `gorm:"size:100" json:"industry"`
	LicenseURL    string         `gorm:"size:500" json:"license_url"`
	LicenseNo     string         `gorm:"size:100" json:"license_no"`
	ContactName   string         `gorm:"size:64" json:"contact_name"`
	ContactPhone  string         `gorm:"size:20" json:"contact_phone"`
	ContactEmail  string         `gorm:"size:128" json:"contact_email"`
	Address       string         `gorm:"size:500" json:"address"`
	AccessToken   string         `gorm:"size:500" json:"-"`
	RefreshToken  string         `gorm:"size:500" json:"-"`
	TokenExpireAt *time.Time     `json:"-"`
	LastSyncAt    *time.Time     `json:"last_sync_at"`
	Remark        string         `gorm:"size:500" json:"remark"`
	CreatedAt     time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy     uint64         `gorm:"default:0" json:"created_by"`
	UpdatedBy     uint64         `gorm:"default:0" json:"updated_by"`
}

// TableName 表名
func (Advertiser) TableName() string {
	return "ad_advertiser"
}

// AdvertiserFund 资金流水表
type AdvertiserFund struct {
	ID              uint64     `gorm:"primaryKey" json:"id"`
	AdvertiserID    uint64     `gorm:"index;not null" json:"advertiser_id"`
	TransactionType string     `gorm:"size:50;index;not null" json:"transaction_type"`
	Amount          float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	BalanceBefore   float64    `gorm:"type:decimal(15,2);default:0" json:"balance_before"`
	BalanceAfter    float64    `gorm:"type:decimal(15,2);default:0" json:"balance_after"`
	TransactionSeq  string     `gorm:"size:100" json:"transaction_seq"`
	TransactionTime *time.Time `gorm:"index" json:"transaction_time"`
	Remark          string     `gorm:"size:500" json:"remark"`
	CreatedAt       time.Time  `json:"created_at"`
}

// TableName 表名
func (AdvertiserFund) TableName() string {
	return "ad_advertiser_fund"
}

// 广告主状态常量
const (
	AdvertiserStatusEnable  = "STATUS_ENABLE"
	AdvertiserStatusDisable = "STATUS_DISABLE"
)

// 交易类型常量
const (
	TransactionTypeRecharge = "recharge" // 充值
	TransactionTypeConsume  = "consume"  // 消费
	TransactionTypeRefund   = "refund"   // 退款
)
