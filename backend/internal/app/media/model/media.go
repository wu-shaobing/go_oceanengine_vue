package model

import (
	"time"

	"gorm.io/gorm"
)

// MaterialImage 图片素材表
type MaterialImage struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	ImageID      string         `gorm:"size:100;uniqueIndex" json:"image_id"` // Ocean Engine 图片ID
	AdvertiserID uint64         `gorm:"index" json:"advertiser_id"`           // 广告主ID
	Filename     string         `gorm:"size:255" json:"filename"`             // 文件名
	Size         int64          `gorm:"default:0" json:"size"`                // 文件大小（字节）
	Width        int            `gorm:"default:0" json:"width"`               // 宽度
	Height       int            `gorm:"default:0" json:"height"`              // 高度
	Format       string         `gorm:"size:20" json:"format"`                // 格式
	URL          string         `gorm:"size:500" json:"url"`                  // 图片URL
	MaterialID   string         `gorm:"size:100" json:"material_id"`          // 素材ID
	Signature    string         `gorm:"size:100" json:"signature"`            // MD5签名
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (MaterialImage) TableName() string {
	return "ad_material_image"
}

// MaterialVideo 视频素材表
type MaterialVideo struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	VideoID      string         `gorm:"size:100;uniqueIndex" json:"video_id"` // Ocean Engine 视频ID
	AdvertiserID uint64         `gorm:"index" json:"advertiser_id"`           // 广告主ID
	Filename     string         `gorm:"size:255" json:"filename"`             // 文件名
	Size         int64          `gorm:"default:0" json:"size"`                // 文件大小（字节）
	Width        int            `gorm:"default:0" json:"width"`               // 宽度
	Height       int            `gorm:"default:0" json:"height"`              // 高度
	Duration     float64        `gorm:"default:0" json:"duration"`            // 时长（秒）
	Format       string         `gorm:"size:20" json:"format"`                // 格式
	URL          string         `gorm:"size:500" json:"url"`                  // 视频URL
	PosterURL    string         `gorm:"size:500" json:"poster_url"`           // 封面URL
	MaterialID   string         `gorm:"size:100" json:"material_id"`          // 素材ID
	Signature    string         `gorm:"size:100" json:"signature"`            // MD5签名
	BitRate      int            `gorm:"default:0" json:"bit_rate"`            // 比特率
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (MaterialVideo) TableName() string {
	return "ad_material_video"
}

// 素材类型常量
const (
	MaterialTypeImage = "image"
	MaterialTypeVideo = "video"
)

// 图片格式
const (
	ImageFormatJPG  = "jpg"
	ImageFormatJPEG = "jpeg"
	ImageFormatPNG  = "png"
	ImageFormatGIF  = "gif"
)

// 视频格式
const (
	VideoFormatMP4 = "mp4"
	VideoFormatAVI = "avi"
	VideoFormatMOV = "mov"
	VideoFormatWMV = "wmv"
)
