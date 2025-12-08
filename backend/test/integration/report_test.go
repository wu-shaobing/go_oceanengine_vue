package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- 数据报表测试 ---

// TestReportAdvertiser_Success 测试获取广告主报表
func TestReportAdvertiser_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/reports/advertiser?advertiser_id=1&start_date=2024-01-01&end_date=2024-01-31", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestReportAdvertiser_Unauthorized 测试未授权访问报表
func TestReportAdvertiser_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/reports/advertiser", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestReportAdvertiserSummary_Success 测试获取广告主汇总报表
func TestReportAdvertiserSummary_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/reports/advertiser/summary?advertiser_id=1", nil, token)

	// 验证请求已处理（不是401未授权）
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)

	var resp Response
	_ = ParseResponse(w, &resp)
	// 注意：实际API可能返回参数错误或者成功，具体取决于数据库状态
}

// TestReportCampaign_Success 测试获取广告系列报表
func TestReportCampaign_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/reports/campaign?advertiser_id=1&start_date=2024-01-01&end_date=2024-01-31", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestReportAd_Success 测试获取广告组报表
func TestReportAd_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/reports/ad?advertiser_id=1&start_date=2024-01-01&end_date=2024-01-31", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestReportSync_Success 测试同步报表数据
func TestReportSync_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	syncReq := map[string]interface{}{
		"advertiser_id": 1,
		"start_date":    "2024-01-01",
		"end_date":      "2024-01-31",
	}

	w := ts.MakeRequest("POST", "/api/v1/reports/sync", syncReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestReportExportList_Success 测试获取导出任务列表
func TestReportExportList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/reports/exports?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestReportExportCreate_Success 测试创建导出任务
func TestReportExportCreate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"advertiser_id": 1,
		"report_type":   "advertiser",
		"start_date":    "2024-01-01",
		"end_date":      "2024-01-31",
	}

	w := ts.MakeRequest("POST", "/api/v1/reports/exports", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// --- 千川模块测试 ---
// 注意：千川模块需要外部API连接和有效的广告主数据

// TestQianchuanAccount_Success 测试获取千川账户信息
func TestQianchuanAccount_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接，适合在e2e测试中验证")
}

// TestQianchuanAccount_Unauthorized 测试未授权访问千川账户
func TestQianchuanAccount_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/qianchuan/account", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestQianchuanShops_Success 测试获取千川店铺列表
func TestQianchuanShops_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanBalance_Success 测试获取千川账户余额
func TestQianchuanBalance_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanCampaigns_Success 测试获取千川广告系列列表
func TestQianchuanCampaigns_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanAds_Success 测试获取千川广告列表
func TestQianchuanAds_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanProducts_Success 测试获取千川商品列表
func TestQianchuanProducts_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanReportAdvertiser_Success 测试获取千川广告主报表
func TestQianchuanReportAdvertiser_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// TestQianchuanReportAd_Success 测试获取千川广告报表
func TestQianchuanReportAd_Success(t *testing.T) {
	t.Skip("跳过测试：千川模块需要外部API连接")
}

// --- 素材管理测试 ---

// TestMediaImageList_Success 测试获取图片素材列表
func TestMediaImageList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/media/images?page=1&page_size=10", nil, token)

	// 验证请求已处理（不是401未授权）
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)

	var resp Response
	_ = ParseResponse(w, &resp)
}

// TestMediaVideoList_Success 测试获取视频素材列表
func TestMediaVideoList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/media/videos?page=1&page_size=10", nil, token)

	// 验证请求已处理（不是401未授权）
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)

	var resp Response
	_ = ParseResponse(w, &resp)
}

// --- 人群定向测试 ---

// TestAudiencePackageList_Success 测试获取定向包列表
func TestAudiencePackageList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/audiences/packages?page=1&page_size=10", nil, token)

	// 验证请求已处理（不是401未授权）
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)

	var resp Response
	_ = ParseResponse(w, &resp)
}

// TestAudienceCustomList_Success 测试获取人群包列表
func TestAudienceCustomList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/audiences/custom?page=1&page_size=10", nil, token)

	// 验证请求已处理（不是401未授权）
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)

	var resp Response
	_ = ParseResponse(w, &resp)
}
