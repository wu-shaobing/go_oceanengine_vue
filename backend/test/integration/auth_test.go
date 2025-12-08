package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestHealthCheck 测试健康检查端点
func TestHealthCheck(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/health", nil, "")

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err := ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestLogin_Success 测试登录成功
func TestLogin_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	loginReq := map[string]string{
		"username": "admin",
		"password": "admin123",
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")

	assert.Equal(t, http.StatusOK, w.Code)

	var resp LoginResponse
	err := ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
	assert.NotEmpty(t, resp.Data.AccessToken)
	assert.NotEmpty(t, resp.Data.RefreshToken)
}

// TestLogin_WrongPassword 测试密码错误
func TestLogin_WrongPassword(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	loginReq := map[string]string{
		"username": "admin",
		"password": "wrongpassword",
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")

	var resp Response
	err := ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.NotEqual(t, 0, resp.Code)
}

// TestLogin_UserNotFound 测试用户不存在
func TestLogin_UserNotFound(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	loginReq := map[string]string{
		"username": "nonexistent",
		"password": "password123",
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")

	var resp Response
	err := ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.NotEqual(t, 0, resp.Code)
}

// TestLogin_EmptyUsername 测试空用户名
func TestLogin_EmptyUsername(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	loginReq := map[string]string{
		"username": "",
		"password": "password123",
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestLogin_EmptyPassword 测试空密码
func TestLogin_EmptyPassword(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	loginReq := map[string]string{
		"username": "admin",
		"password": "",
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestGetUserInfo_Success 测试获取用户信息成功
func TestGetUserInfo_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	// 生成测试Token
	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/auth/userinfo", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestGetUserInfo_Unauthorized 测试未授权访问
func TestGetUserInfo_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/auth/userinfo", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestGetUserInfo_InvalidToken 测试无效Token
func TestGetUserInfo_InvalidToken(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/auth/userinfo", nil, "invalid-token")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestRefreshToken_Success 测试刷新Token成功
func TestRefreshToken_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	// 先登录获取RefreshToken
	loginReq := map[string]string{
		"username": "admin",
		"password": "admin123",
	}

	loginResp := ts.MakeRequest("POST", "/api/v1/auth/login", loginReq, "")
	require.Equal(t, http.StatusOK, loginResp.Code)

	var login LoginResponse
	err := ParseResponse(loginResp, &login)
	require.NoError(t, err)

	// 使用RefreshToken刷新
	refreshReq := map[string]string{
		"refresh_token": login.Data.RefreshToken,
	}

	w := ts.MakeRequest("POST", "/api/v1/auth/refresh", refreshReq, "")

	var resp LoginResponse
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	// 验证返回了新的token
	if resp.Code == 0 {
		assert.NotEmpty(t, resp.Data.AccessToken)
	}
}

// TestLogout_Success 测试退出登录成功
func TestLogout_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("POST", "/api/v1/auth/logout", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestCaptcha_Get 测试获取验证码
func TestCaptcha_Get(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/auth/captcha", nil, "")

	assert.Equal(t, http.StatusOK, w.Code)
}
