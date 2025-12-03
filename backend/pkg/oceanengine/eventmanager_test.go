package oceanengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventManagerService(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewEventManagerService(client)

	assert.NotNil(t, service)
	assert.NotNil(t, service.client)
}

func TestEventManagerService_AssetsGetRequest(t *testing.T) {
	req := &AssetsGetRequest{
		AdvertiserID: 67890,
		AssetType:    "SITE",
		LandingType:  "EXTERNAL",
		Page:         1,
		PageSize:     20,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, "SITE", req.AssetType)
	assert.Equal(t, "EXTERNAL", req.LandingType)
}

func TestEventManagerService_AssetsCreateRequest(t *testing.T) {
	req := &AssetsCreateRequest{
		AdvertiserID: 67890,
		AssetType:    "SITE",
		AssetName:    "测试资产",
		LandingType:  "EXTERNAL",
		DownloadURL:  "https://example.com",
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, "测试资产", req.AssetName)
	assert.Equal(t, "https://example.com", req.DownloadURL)
}

func TestEventManagerService_EventsCreateRequest(t *testing.T) {
	req := &EventsCreateRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
		EventTypes:   []string{"form_submit", "button_click"},
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
	assert.Equal(t, 2, len(req.EventTypes))
}

func TestEventManagerService_TrackURLCreateRequest(t *testing.T) {
	req := &TrackURLCreateRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
		GroupName:    "测试监测链接组",
		TrackURLs: []TrackURLSetting{
			{
				ActionType: "click",
				TrackURL:   "https://track.example.com/click",
			},
			{
				ActionType: "impression",
				TrackURL:   "https://track.example.com/impression",
			},
		},
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, "测试监测链接组", req.GroupName)
	assert.Equal(t, 2, len(req.TrackURLs))
}

func TestEventManagerService_ShareRequest(t *testing.T) {
	req := &ShareRequest{
		AdvertiserID:        67890,
		AssetID:             12345,
		TargetAdvertiserIDs: []int64{11111, 22222},
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
	assert.Equal(t, 2, len(req.TargetAdvertiserIDs))
}

func TestEventManagerService_ConversionRequest(t *testing.T) {
	req := &ConversionRequest{
		EventType: "form_submit",
		Context: map[string]interface{}{
			"ad": map[string]interface{}{
				"callback": "test-callback-data",
			},
		},
		Timestamp: 1704067200,
	}

	assert.Equal(t, "form_submit", req.EventType)
	assert.NotNil(t, req.Context)
}

func TestEventManagerService_AddPublicKeyRequest(t *testing.T) {
	req := &AddPublicKeyRequest{
		AdvertiserID: 67890,
		PublicKey:    "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkq...\n-----END PUBLIC KEY-----",
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Contains(t, req.PublicKey, "BEGIN PUBLIC KEY")
}

func TestEventManagerTypes(t *testing.T) {
	// 测试 Asset 类型
	asset := Asset{
		AssetID:     12345,
		AssetName:   "测试资产",
		AssetType:   "SITE",
		LandingType: "EXTERNAL",
		Status:      1,
	}
	assert.Equal(t, int64(12345), asset.AssetID)
	assert.Equal(t, "SITE", asset.AssetType)

	// 测试 EventConfig 类型
	eventConfig := EventConfig{
		EventID:      11111,
		EventType:    "form_submit",
		EventName:    "表单提交",
		OptimizeGoal: "CONVERT",
		Status:       1,
	}
	assert.Equal(t, int64(11111), eventConfig.EventID)
	assert.Equal(t, "form_submit", eventConfig.EventType)

	// 测试 TrackURL 类型
	trackURL := TrackURL{
		TrackURLID:   22222,
		TrackURLName: "监测链接",
	}
	assert.Equal(t, int64(22222), trackURL.TrackURLID)

	// 测试 OptimizedGoal 类型
	goal := OptimizedGoal{
		OptimizeGoal:     "CONVERT",
		OptimizeGoalName: "表单提交",
		EventType:        "form_submit",
		DeepBidType:      "DEEP_BID_DEFAULT",
	}
	assert.Equal(t, "CONVERT", goal.OptimizeGoal)

	// 测试 PublicKey 类型
	publicKey := PublicKey{
		KeyID:     "33333",
		PublicKey: "test-key",
		Status:    1,
	}
	assert.Equal(t, "33333", publicKey.KeyID)
}

func TestEventManagerRequestParams(t *testing.T) {
	// 测试资产获取参数构建
	assetsReq := &AssetsGetRequest{
		AdvertiserID: 67890,
		AssetType:    "SITE",
		LandingType:  "EXTERNAL",
		Page:         1,
		PageSize:     20,
	}

	params := map[string]interface{}{
		"advertiser_id": assetsReq.AdvertiserID,
	}
	if assetsReq.AssetType != "" {
		params["asset_type"] = assetsReq.AssetType
	}
	if assetsReq.LandingType != "" {
		params["landing_type"] = assetsReq.LandingType
	}
	if assetsReq.Page > 0 {
		params["page"] = assetsReq.Page
	}
	if assetsReq.PageSize > 0 {
		params["page_size"] = assetsReq.PageSize
	}

	assert.Equal(t, int64(67890), params["advertiser_id"])
	assert.Equal(t, "SITE", params["asset_type"])
	assert.Equal(t, "EXTERNAL", params["landing_type"])
	assert.Equal(t, 1, params["page"])
	assert.Equal(t, 20, params["page_size"])
}

func TestEventManagerService_AllAssetsListRequest(t *testing.T) {
	req := &AllAssetsListRequest{
		AdvertiserID: 67890,
		AssetType:    "APP",
		LandingType:  "APP_ANDROID",
		Page:         1,
		PageSize:     10,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, "APP", req.AssetType)
	assert.Equal(t, "APP_ANDROID", req.LandingType)
}

func TestEventManagerService_AvailableEventsGetRequest(t *testing.T) {
	req := &AvailableEventsGetRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
}

func TestEventManagerService_EventConfigsGetRequest(t *testing.T) {
	req := &EventConfigsGetRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
}

func TestEventManagerService_TrackURLGetRequest(t *testing.T) {
	req := &TrackURLGetRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
}

func TestEventManagerService_TrackURLUpdateRequest(t *testing.T) {
	req := &TrackURLUpdateRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
		GroupID:      33333,
		GroupName:    "更新后的监测链接组",
		TrackURLs: []TrackURLSetting{
			{
				ActionType: "click",
				TrackURL:   "https://track.example.com/new-click",
			},
		},
	}

	assert.Equal(t, int64(33333), req.GroupID)
	assert.Equal(t, "更新后的监测链接组", req.GroupName)
}

func TestEventManagerService_ShareGetRequest(t *testing.T) {
	req := &ShareGetRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.AssetID)
}

func TestEventManagerService_EventConvertOptimizedGoalGetRequest(t *testing.T) {
	req := &EventConvertOptimizedGoalGetRequest{
		AdvertiserID: 67890,
		AssetID:      12345,
		LandingType:  "EXTERNAL",
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, "EXTERNAL", req.LandingType)
}

func TestEventManagerService_GetAllPublicKeysRequest(t *testing.T) {
	req := &GetAllPublicKeysRequest{
		AdvertiserID: 67890,
	}

	assert.Equal(t, int64(67890), req.AdvertiserID)
}
