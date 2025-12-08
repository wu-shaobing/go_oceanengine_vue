package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserList_Success 测试获取用户列表
func TestUserList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/users?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestUserList_Unauthorized 测试未授权访问用户列表
func TestUserList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()

	w := ts.MakeRequest("GET", "/api/v1/system/users", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestUserCreate_Success 测试创建用户
func TestUserCreate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"username": "testuser",
		"password": "test123456",
		"nickname": "测试用户",
		"email":    "test@test.com",
		"phone":    "13900139000",
		"role_id":  1,
		"status":   1,
	}

	w := ts.MakeRequest("POST", "/api/v1/system/users", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestUserCreate_DuplicateUsername 测试创建重复用户名
func TestUserCreate_DuplicateUsername(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"username": "admin", // 已存在的用户名
		"password": "test123456",
		"nickname": "测试用户",
		"email":    "test2@test.com",
		"phone":    "13900139001",
		"role_id":  1,
		"status":   1,
	}

	w := ts.MakeRequest("POST", "/api/v1/system/users", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	// 预期返回错误
	assert.NotEqual(t, 0, resp.Code)
}

// TestUserGetByID_Success 测试获取用户详情
func TestUserGetByID_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/users/1", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestUserUpdate_Success 测试更新用户
func TestUserUpdate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	updateReq := map[string]interface{}{
		"nickname": "更新后的昵称",
		"email":    "updated@test.com",
		"status":   1,
	}

	w := ts.MakeRequest("PUT", "/api/v1/system/users/1", updateReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestRoleList_Success 测试获取角色列表
func TestRoleList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/roles", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestRoleCreate_Success 测试创建角色
func TestRoleCreate_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	createReq := map[string]interface{}{
		"name":       "测试角色",
		"key":        "test_role",
		"sort":       10,
		"status":     1,
		"data_scope": "1",
		"remark":     "测试角色备注",
	}

	w := ts.MakeRequest("POST", "/api/v1/system/roles", createReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestMenuList_Success 测试获取菜单列表
func TestMenuList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/menus", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestMenuTree_Success 测试获取菜单树
func TestMenuTree_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/menus/tree", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestUserMenuTree_Success 测试获取用户菜单树
func TestUserMenuTree_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/menus/user", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestOperationLogList_Success 测试获取操作日志
func TestOperationLogList_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/logs/operation?page=1&page_size=10", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestChangePassword_Success 测试修改密码
func TestChangePassword_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	changeReq := map[string]string{
		"old_password": "admin123",
		"new_password": "newpassword123",
	}

	w := ts.MakeRequest("POST", "/api/v1/system/users/change-password", changeReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestChangePassword_WrongOldPassword 测试修改密码-旧密码错误
func TestChangePassword_WrongOldPassword(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	changeReq := map[string]string{
		"old_password": "wrongpassword",
		"new_password": "newpassword123",
	}

	w := ts.MakeRequest("POST", "/api/v1/system/users/change-password", changeReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.NotEqual(t, 0, resp.Code)
}

// TestResetPassword_Success 测试重置密码
func TestResetPassword_Success(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	resetReq := map[string]string{
		"new_password": "resetpassword123",
	}

	w := ts.MakeRequest("POST", "/api/v1/system/users/1/reset-password", resetReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestRoleMenus_Get 测试获取角色菜单
func TestRoleMenus_Get(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	w := ts.MakeRequest("GET", "/api/v1/system/roles/1/menus", nil, token)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

// TestRoleMenus_Update 测试更新角色菜单
func TestRoleMenus_Update(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	updateReq := map[string]interface{}{
		"menu_ids": []int{1, 2},
	}

	w := ts.MakeRequest("PUT", "/api/v1/system/roles/1/menus", updateReq, token)

	var resp Response
	err = ParseResponse(w, &resp)
	require.NoError(t, err)
}

// TestPagination 测试分页参数
func TestPagination(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Cleanup()
	ts.SeedTestData(t)

	token, err := ts.GenerateTestToken(1, "admin")
	require.NoError(t, err)

	// 测试不同的分页参数
	testCases := []struct {
		page     int
		pageSize int
	}{
		{1, 5},
		{1, 10},
		{2, 10},
		{1, 20},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("page=%d,pageSize=%d", tc.page, tc.pageSize), func(t *testing.T) {
			url := fmt.Sprintf("/api/v1/system/users?page=%d&page_size=%d", tc.page, tc.pageSize)
			w := ts.MakeRequest("GET", url, nil, token)
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}
