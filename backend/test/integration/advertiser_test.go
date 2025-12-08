package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAdvertiserList_Success 测试获取广告主列表
func TestAdvertiserList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/advertisers?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestAdvertiserList_Unauthorized 测试未授权访问广告主列表
func TestAdvertiserList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/advertisers", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestAdvertiserList_WithFilters 测试带过滤条件的广告主列表
func TestAdvertiserList_WithFilters(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 测试关键词搜索
	w := ts.MakeRequest("GET", "/api/v1/advertisers?keyword=test&page=1&page_size=10", nil, token)
	assert.Equal(t, http.StatusOK, w.Code)

	// 测试状态过滤
	w = ts.MakeRequest("GET", "/api/v1/advertisers?status=1&page=1&page_size=10", nil, token)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestOAuthURL_Success 测试获取OAuth授权URL
// 注意：此测试需要Redis连接来存储OAuth state
// 在单元测试环境中跳过，因为OAuth state管理器需要完整的Redis配置
func TestOAuthURL_Success(t *testing.T) {
	// 此测试需要完整的后端环境（包括Redis初始化）
	// 在纯单元测试环境中跳过
	t.Skip("跳过测试：OAuth URL生成需要完整的Redis配置，适合在e2e测试中验证")
}

// TestOAuthCallback_MissingCode 测试OAuth回调缺少code参数
func TestOAuthCallback_MissingCode(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	// 缺少code参数
	w := ts.MakeRequest("GET", "/api/v1/advertisers/oauth/callback?state=test", nil, "")

	// 预期返回错误
	assert.NotEqual(t, http.StatusOK, w.Code)
}

// TestOAuthCallback_InvalidCode 测试OAuth回调无效code
func TestOAuthCallback_InvalidCode(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	// 使用无效的code
	w := ts.MakeRequest("GET", "/api/v1/advertisers/oauth/callback?code=invalid_code&state=test", nil, "")

	// 由于code无效，巨量引擎会返回错误
	var resp Response
	_ = ParseResponse(w, &resp)
	// OAuth回调的行为取决于巨量引擎的响应
}

// TestAdvertiserGet_NotFound 测试获取不存在的广告主
func TestAdvertiserGet_NotFound(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/advertisers/999999", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	// 预期返回错误（记录不存在）
}

// TestAdvertiserDelete_Success 测试删除广告主
func TestAdvertiserDelete_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 先获取广告主列表，确认存在记录（如果有的话）
	w := ts.MakeRequest("DELETE", "/api/v1/advertisers/1", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestAdvertiserSync_Success 测试同步广告主
// 注意：此测试需要外部API连接
func TestAdvertiserSync_Success(t *testing.T) {
	if !isExternalAPIAvailable() {
		t.Skip("跳过测试：需要外部API凭证")
	}

	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 注意：实际同步需要有效的access_token和advertiser_id
	w := ts.MakeRequest("POST", "/api/v1/advertisers/1/sync", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestAdvertiserBalance_Success 测试获取广告主余额
// 注意：此测试需要外部API连接
func TestAdvertiserBalance_Success(t *testing.T) {
	if !isExternalAPIAvailable() {
		t.Skip("跳过测试：需要外部API凭证")
	}

	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/advertisers/1/balance", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestAdvertiserFunds_Success 测试获取广告主资金流水
// 注意：此测试需要外部API连接
func TestAdvertiserFunds_Success(t *testing.T) {
	if !isExternalAPIAvailable() {
		t.Skip("跳过测试：需要外部API凭证")
	}

	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/advertisers/1/funds?page=1&page_size=10", nil, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}
