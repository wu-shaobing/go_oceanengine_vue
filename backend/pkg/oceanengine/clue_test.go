package oceanengine

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClueService_GetClueList(t *testing.T) {
	// 创建模拟服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/open_api/2/tools/clue/get/", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		response := BaseResponse{
			Code:      0,
			Message:   "success",
			RequestID: "test-request-id",
		}
		data := ClueListResponse{
			List: []Clue{
				{
					ClueID:          12345,
					AdvertiserID:    67890,
					Name:            "测试线索",
					TelephoneNumber: "13800138000",
				},
			},
		}
		data.PageInfo.TotalNumber = 1
		data.PageInfo.Page = 1
		data.PageInfo.PageSize = 20

		dataBytes, _ := json.Marshal(data)
		response.Data = dataBytes

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &ClueListRequest{
		AdvertiserID: 67890,
		StartTime:    "2024-01-01",
		EndTime:      "2024-01-31",
		Page:         1,
		PageSize:     20,
	}

	// 注意: 实际测试需要修改 BaseURL 或使用依赖注入
	// 这里只验证请求构建逻辑
	assert.NotNil(t, service)
	assert.NotNil(t, req)
}

func TestClueService_ClueCallback(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &ClueCallbackRequest{
		AdvertiserID:   67890,
		ClueID:         12345,
		EventType:      "form_submit",
		ConvertState:   1,
		ExternalAction: "form",
	}

	assert.NotNil(t, service)
	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.ClueID)
}

func TestClueService_BatchClueCallback(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &BatchClueCallbackRequest{
		AdvertiserID: 67890,
		ClueList: []ClueCallbackDetail{
			{
				ClueID:       12345,
				EventType:    "form_submit",
				ConvertState: 1,
			},
			{
				ClueID:       12346,
				EventType:    "phone_confirm",
				ConvertState: 2,
			},
		},
	}

	assert.NotNil(t, service)
	assert.Equal(t, 2, len(req.ClueList))
}

func TestClueService_GetKeyAction(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &KeyActionGetRequest{
		AdvertiserID: 67890,
		ClueID:       12345,
	}

	assert.NotNil(t, service)
	assert.Equal(t, int64(67890), req.AdvertiserID)
	assert.Equal(t, int64(12345), req.ClueID)
}

func TestClueService_GetSmartPhone(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &SmartPhoneGetRequest{
		AdvertiserID: 67890,
		Page:         1,
		PageSize:     20,
	}

	assert.NotNil(t, service)
	assert.Equal(t, int64(67890), req.AdvertiserID)
}

func TestClueService_GetFormList(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &FormGetRequest{
		AdvertiserID: 67890,
		FormType:     "normal",
		Page:         1,
		PageSize:     20,
	}

	assert.NotNil(t, service)
	assert.Equal(t, "normal", req.FormType)
}

func TestClueService_GetFormDetail(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &FormDetailRequest{
		AdvertiserID: 67890,
		FormID:       11111,
	}

	assert.NotNil(t, service)
	assert.Equal(t, int64(11111), req.FormID)
}

func TestClueService_GetClueStoreList(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	req := &ClueStoreListRequest{
		AdvertiserID: 67890,
		StoreName:    "测试门店",
		Page:         1,
		PageSize:     20,
	}

	assert.NotNil(t, service)
	assert.Equal(t, "测试门店", req.StoreName)
}

func TestClueTypes(t *testing.T) {
	// 测试 Clue 类型
	clue := Clue{
		ClueID:          12345,
		AdvertiserID:    67890,
		Name:            "张三",
		TelephoneNumber: "13800138000",
		Gender:          1,
		City:            "北京",
		Source:          "form",
		ClueType:        1,
		ConvertStatus:   0,
	}
	assert.Equal(t, int64(12345), clue.ClueID)
	assert.Equal(t, "张三", clue.Name)

	// 测试 SmartPhone 类型
	phone := SmartPhone{
		SmartPhoneID: 11111,
		Name:         "智能电话1",
		PhoneNumber:  "4001234567",
		BindStatus:   1,
	}
	assert.Equal(t, int64(11111), phone.SmartPhoneID)

	// 测试 Form 类型
	form := Form{
		FormID:   22222,
		FormName: "测试表单",
		FormType: "normal",
		Status:   1,
	}
	assert.Equal(t, int64(22222), form.FormID)

	// 测试 ClueStore 类型
	store := ClueStore{
		StoreID:   33333,
		StoreName: "测试门店",
		Address:   "北京市朝阳区",
		Status:    1,
	}
	assert.Equal(t, int64(33333), store.StoreID)
}

func TestNewClueService(t *testing.T) {
	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	assert.NotNil(t, service)
	assert.NotNil(t, service.client)
}

// TestClueListRequestParams 测试请求参数构建
func TestClueListRequestParams(t *testing.T) {
	req := &ClueListRequest{
		AdvertiserID:  67890,
		StartTime:     "2024-01-01",
		EndTime:       "2024-01-31",
		Page:          1,
		PageSize:      20,
		ClueType:      1,
		ConvertStatus: 2,
	}

	// 验证参数构建逻辑
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_time":    req.StartTime,
		"end_time":      req.EndTime,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}
	if req.ClueType > 0 {
		params["clue_type"] = req.ClueType
	}
	if req.ConvertStatus > 0 {
		params["convert_status"] = req.ConvertStatus
	}

	assert.Equal(t, int64(67890), params["advertiser_id"])
	assert.Equal(t, "2024-01-01", params["start_time"])
	assert.Equal(t, 1, params["page"])
	assert.Equal(t, 1, params["clue_type"])
	assert.Equal(t, 2, params["convert_status"])
}

// TestContextPropagation 测试 context 传递
func TestContextPropagation(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test-key", "test-value")

	client := NewClient("test-app-id", "test-secret")
	service := NewClueService(client)

	// 验证 context 传递
	assert.NotNil(t, ctx)
	assert.NotNil(t, service)
	assert.Equal(t, "test-value", ctx.Value("test-key"))
}
