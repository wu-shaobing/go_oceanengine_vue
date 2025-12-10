package service

import (
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// SettingService 用户设置服务
type SettingService struct {
	db *gorm.DB
}

// NewSettingService 创建用户设置服务
func NewSettingService(db *gorm.DB) *SettingService {
	return &SettingService{db: db}
}

// GetUserSetting 获取用户设置
func (s *SettingService) GetUserSetting(ctx context.Context, userID uint64) (*dto.UserSettingResp, error) {
	var setting model.UserSetting
	err := s.db.WithContext(ctx).Where("user_id = ?", userID).First(&setting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 返回默认设置
			defaultSetting := model.DefaultUserSettingData()
			return &dto.UserSettingResp{
				Language:             defaultSetting.Language,
				Timezone:             defaultSetting.Timezone,
				Theme:                defaultSetting.Theme,
				NotificationsEnabled: defaultSetting.NotificationsEnabled,
				EmailAlertsEnabled:   defaultSetting.EmailAlertsEnabled,
				SmsAlertsEnabled:     defaultSetting.SmsAlertsEnabled,
				AutoRefreshEnabled:   defaultSetting.AutoRefreshEnabled,
				RefreshInterval:      defaultSetting.RefreshInterval,
			}, nil
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 解析 JSON 设置
	var settingData model.UserSettingData
	if err := json.Unmarshal(setting.Settings, &settingData); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.UserSettingResp{
		Language:             settingData.Language,
		Timezone:             settingData.Timezone,
		Theme:                settingData.Theme,
		NotificationsEnabled: settingData.NotificationsEnabled,
		EmailAlertsEnabled:   settingData.EmailAlertsEnabled,
		SmsAlertsEnabled:     settingData.SmsAlertsEnabled,
		AutoRefreshEnabled:   settingData.AutoRefreshEnabled,
		RefreshInterval:      settingData.RefreshInterval,
	}, nil
}

// UpdateUserSetting 更新用户设置
func (s *SettingService) UpdateUserSetting(ctx context.Context, userID uint64, req *dto.UserSettingUpdateReq) error {
	// 获取当前设置或默认设置
	var setting model.UserSetting
	var settingData *model.UserSettingData

	err := s.db.WithContext(ctx).Where("user_id = ?", userID).First(&setting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 使用默认设置
			settingData = model.DefaultUserSettingData()
		} else {
			return errcode.Wrap(errcode.ErrInternalServer, err)
		}
	} else {
		// 解析现有设置
		settingData = &model.UserSettingData{}
		if err := json.Unmarshal(setting.Settings, settingData); err != nil {
			settingData = model.DefaultUserSettingData()
		}
	}

	// 更新设置（只更新非 nil 的字段）
	if req.Language != nil {
		settingData.Language = *req.Language
	}
	if req.Timezone != nil {
		settingData.Timezone = *req.Timezone
	}
	if req.Theme != nil {
		settingData.Theme = *req.Theme
	}
	if req.NotificationsEnabled != nil {
		settingData.NotificationsEnabled = *req.NotificationsEnabled
	}
	if req.EmailAlertsEnabled != nil {
		settingData.EmailAlertsEnabled = *req.EmailAlertsEnabled
	}
	if req.SmsAlertsEnabled != nil {
		settingData.SmsAlertsEnabled = *req.SmsAlertsEnabled
	}
	if req.AutoRefreshEnabled != nil {
		settingData.AutoRefreshEnabled = *req.AutoRefreshEnabled
	}
	if req.RefreshInterval != nil {
		settingData.RefreshInterval = *req.RefreshInterval
	}

	// 序列化设置
	settingsJSON, err := json.Marshal(settingData)
	if err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 使用 upsert 操作
	if setting.ID == 0 {
		// 创建新记录
		newSetting := &model.UserSetting{
			UserID:   userID,
			Settings: settingsJSON,
		}
		if err := s.db.WithContext(ctx).Create(newSetting).Error; err != nil {
			return errcode.Wrap(errcode.ErrInternalServer, err)
		}
	} else {
		// 更新现有记录
		if err := s.db.WithContext(ctx).Model(&setting).Update("settings", settingsJSON).Error; err != nil {
			return errcode.Wrap(errcode.ErrInternalServer, err)
		}
	}

	return nil
}
