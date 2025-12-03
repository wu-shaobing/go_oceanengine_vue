package model

import (
	"time"

	"gorm.io/gorm"
)

// AdvertiserReport 广告主报告
type AdvertiserReport struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID uint64    `gorm:"index;not null" json:"advertiser_id"`
	StatDate     string    `gorm:"size:10;index;not null" json:"stat_date"` // YYYY-MM-DD
	Cost         float64   `gorm:"type:decimal(14,2);default:0" json:"cost"`
	Show         int64     `gorm:"default:0" json:"show"`
	Click        int64     `gorm:"default:0" json:"click"`
	Convert      int64     `gorm:"default:0" json:"convert"`
	CTR          float64   `gorm:"type:decimal(10,4);default:0" json:"ctr"`
	CVR          float64   `gorm:"type:decimal(10,4);default:0" json:"cvr"`
	CPM          float64   `gorm:"type:decimal(10,2);default:0" json:"cpm"`
	CPC          float64   `gorm:"type:decimal(10,2);default:0" json:"cpc"`
	ConvertCost  float64   `gorm:"type:decimal(10,2);default:0" json:"convert_cost"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 表名
func (AdvertiserReport) TableName() string {
	return "rpt_advertiser"
}

// CampaignReport 广告系列报告
type CampaignReport struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID uint64    `gorm:"index;not null" json:"advertiser_id"`
	CampaignID   uint64    `gorm:"index;not null" json:"campaign_id"`
	StatDate     string    `gorm:"size:10;index;not null" json:"stat_date"`
	Cost         float64   `gorm:"type:decimal(14,2);default:0" json:"cost"`
	Show         int64     `gorm:"default:0" json:"show"`
	Click        int64     `gorm:"default:0" json:"click"`
	Convert      int64     `gorm:"default:0" json:"convert"`
	CTR          float64   `gorm:"type:decimal(10,4);default:0" json:"ctr"`
	CVR          float64   `gorm:"type:decimal(10,4);default:0" json:"cvr"`
	CPM          float64   `gorm:"type:decimal(10,2);default:0" json:"cpm"`
	CPC          float64   `gorm:"type:decimal(10,2);default:0" json:"cpc"`
	ConvertCost  float64   `gorm:"type:decimal(10,2);default:0" json:"convert_cost"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 表名
func (CampaignReport) TableName() string {
	return "rpt_campaign"
}

// AdReport 广告组报告
type AdReport struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID uint64    `gorm:"index;not null" json:"advertiser_id"`
	CampaignID   uint64    `gorm:"index;not null" json:"campaign_id"`
	AdID         uint64    `gorm:"index;not null" json:"ad_id"`
	StatDate     string    `gorm:"size:10;index;not null" json:"stat_date"`
	Cost         float64   `gorm:"type:decimal(14,2);default:0" json:"cost"`
	Show         int64     `gorm:"default:0" json:"show"`
	Click        int64     `gorm:"default:0" json:"click"`
	Convert      int64     `gorm:"default:0" json:"convert"`
	CTR          float64   `gorm:"type:decimal(10,4);default:0" json:"ctr"`
	CVR          float64   `gorm:"type:decimal(10,4);default:0" json:"cvr"`
	CPM          float64   `gorm:"type:decimal(10,2);default:0" json:"cpm"`
	CPC          float64   `gorm:"type:decimal(10,2);default:0" json:"cpc"`
	ConvertCost  float64   `gorm:"type:decimal(10,2);default:0" json:"convert_cost"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 表名
func (AdReport) TableName() string {
	return "rpt_ad"
}

// ExportTask 导出任务
type ExportTask struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	AdvertiserID uint64         `gorm:"index;not null" json:"advertiser_id"`
	TaskType     string         `gorm:"size:32;not null" json:"task_type"`
	Status       string         `gorm:"size:32;default:'PENDING'" json:"status"`
	FileName     string         `gorm:"size:255" json:"file_name"`
	FilePath     string         `gorm:"size:512" json:"file_path"`
	FileSize     int64          `gorm:"default:0" json:"file_size"`
	StartDate    string         `gorm:"size:10" json:"start_date"`
	EndDate      string         `gorm:"size:10" json:"end_date"`
	ErrorMsg     string         `gorm:"size:500" json:"error_msg"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (ExportTask) TableName() string {
	return "rpt_export_task"
}

// 导出任务状态
const (
	ExportStatusPending    = "PENDING"
	ExportStatusProcessing = "PROCESSING"
	ExportStatusCompleted  = "COMPLETED"
	ExportStatusFailed     = "FAILED"
)

// 导出任务类型
const (
	ExportTypeAdvertiser = "ADVERTISER"
	ExportTypeCampaign   = "CAMPAIGN"
	ExportTypeAd         = "AD"
)
