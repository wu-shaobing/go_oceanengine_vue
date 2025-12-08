package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCampaignList_Success 测试获取广告系列列表
func TestCampaignList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/campaigns?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestCampaignList_Unauthorized 测试未授权访问广告系列列表
func TestCampaignList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/campaigns", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestCampaignList_WithFilters 测试带过滤条件的广告系列列表
func TestCampaignList_WithFilters(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 测试按广告主ID过滤
	w := ts.MakeRequest("GET", "/api/v1/campaigns?advertiser_id=1&page=1&page_size=10", nil, token)
	assert.Equal(t, http.StatusOK, w.Code)

	// 测试按状态过滤
	w = ts.MakeRequest("GET", "/api/v1/campaigns?status=enable&page=1&page_size=10", nil, token)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestCampaignCreate_Success 测试创建广告系列
func TestCampaignCreate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"advertiser_id":  1,
		"campaign_name":  "测试广告系列",
		"budget_mode":    "BUDGET_MODE_DAY",
		"budget":         1000.00,
		"landing_type":   "LINK",
		"marketing_goal": "VIDEO_PROM_GOODS",
	}

	w := ts.MakeRequest("POST", "/api/v1/campaigns", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestCampaignCreate_MissingFields 测试创建广告系列缺少必填字段
func TestCampaignCreate_MissingFields(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 缺少campaign_name
	createReq := map[string]interface{}{
		"advertiser_id": 1,
		"budget_mode":   "BUDGET_MODE_DAY",
	}

	w := ts.MakeRequest("POST", "/api/v1/campaigns", createReq, token)

	// 预期返回错误
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestCampaignGet_NotFound 测试获取不存在的广告系列
func TestCampaignGet_NotFound(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/campaigns/999999", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestCampaignUpdate_Success 测试更新广告系列
func TestCampaignUpdate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	updateReq := map[string]interface{}{
		"campaign_name": "更新后的广告系列名称",
		"budget":        2000.00,
	}

	w := ts.MakeRequest("PUT", "/api/v1/campaigns/1", updateReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestCampaignUpdateStatus_Success 测试更新广告系列状态
func TestCampaignUpdateStatus_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	statusReq := map[string]interface{}{
		"campaign_ids": []int{1},
		"opt_status":   "disable",
	}

	w := ts.MakeRequest("PUT", "/api/v1/campaigns/status", statusReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestCampaignDelete_Success 测试删除广告系列
func TestCampaignDelete_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("DELETE", "/api/v1/campaigns/1", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestCampaignSync_Success 测试同步广告系列
func TestCampaignSync_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("POST", "/api/v1/campaigns/sync/1", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// --- 广告组测试 ---

// TestAdList_Success 测试获取广告组列表
func TestAdList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/ads?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestAdList_Unauthorized 测试未授权访问广告组列表
func TestAdList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/ads", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestAdCreate_Success 测试创建广告组
func TestAdCreate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"advertiser_id": 1,
		"campaign_id":   1,
		"name":          "测试广告组",
		"budget_mode":   "BUDGET_MODE_DAY",
		"budget":        500.00,
	}

	w := ts.MakeRequest("POST", "/api/v1/ads", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestAdUpdateStatus_Success 测试更新广告组状态
func TestAdUpdateStatus_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	statusReq := map[string]interface{}{
		"ad_ids":     []int{1},
		"opt_status": "disable",
	}

	w := ts.MakeRequest("PUT", "/api/v1/ads/status", statusReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// --- 创意测试 ---

// TestCreativeList_Success 测试获取创意列表
func TestCreativeList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/creatives?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestCreativeList_Unauthorized 测试未授权访问创意列表
func TestCreativeList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/creatives", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestCreativeUpdateStatus_Success 测试更新创意状态
func TestCreativeUpdateStatus_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	statusReq := map[string]interface{}{
		"creative_ids": []int{1},
		"opt_status":   "disable",
	}

	w := ts.MakeRequest("PUT", "/api/v1/creatives/status", statusReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}
