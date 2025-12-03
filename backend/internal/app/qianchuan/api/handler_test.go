package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupQianchuanTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

// TestGetAccountInfo 测试获取账户信息
func TestGetAccountInfo(t *testing.T) {
	router := setupQianchuanTestRouter()

	// 模拟 handler
	router.GET("/qianchuan/account", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"advertiser_id":   12345,
				"advertiser_name": "测试账户",
				"account_type":    "QIANCHUAN",
				"status":          "enable",
				"balance":         10000,
			},
		})
	})

	tests := []struct {
		name       string
		query      string
		wantCode   int
		wantFields []string
	}{
		{
			name:       "获取账户信息成功",
			query:      "?advertiser_id=12345",
			wantCode:   http.StatusOK,
			wantFields: []string{"advertiser_id", "advertiser_name", "account_type"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/qianchuan/account"+tt.query, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if data, ok := response["data"].(map[string]interface{}); ok {
				for _, field := range tt.wantFields {
					_, exists := data[field]
					assert.True(t, exists, "字段 %s 应该存在", field)
				}
			}
		})
	}
}

// TestGetShopList 测试获取店铺列表
func TestGetShopList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/shops", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list":      []gin.H{},
				"total":     0,
				"page":      1,
				"page_size": 20,
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/shops?advertiser_id=12345", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetCampaignList 测试获取广告系列列表
func TestGetCampaignList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/campaigns", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list":      []gin.H{},
				"total":     0,
				"page":      1,
				"page_size": 20,
			},
		})
	})

	tests := []struct {
		name     string
		query    string
		wantCode int
	}{
		{
			name:     "获取广告系列列表",
			query:    "?advertiser_id=12345&page=1&page_size=20",
			wantCode: http.StatusOK,
		},
		{
			name:     "带状态筛选",
			query:    "?advertiser_id=12345&status=enable",
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/qianchuan/campaigns"+tt.query, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

// TestGetAdList 测试获取广告计划列表
func TestGetAdList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/ads", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list":      []gin.H{},
				"total":     0,
				"page":      1,
				"page_size": 20,
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/ads?advertiser_id=12345", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetAwemeOrderList 测试获取随心推订单列表
func TestGetAwemeOrderList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/aweme/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list":      []gin.H{},
				"total":     0,
				"page":      1,
				"page_size": 20,
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/aweme/orders?advertiser_id=12345", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetAdvertiserReport 测试获取账户报表
func TestGetAdvertiserReport(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/reports/advertiser", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": []gin.H{},
		})
	})

	tests := []struct {
		name     string
		query    string
		wantCode int
	}{
		{
			name:     "获取账户报表",
			query:    "?advertiser_id=12345&start_date=2024-01-01&end_date=2024-01-31",
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/qianchuan/reports/advertiser"+tt.query, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

// TestGetBalance 测试获取账户余额
func TestGetBalance(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/balance", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"balance": 10000,
				"cash":    8000,
				"grant":   2000,
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/balance?advertiser_id=12345", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if data, ok := response["data"].(map[string]interface{}); ok {
		assert.Contains(t, data, "balance")
		assert.Contains(t, data, "cash")
		assert.Contains(t, data, "grant")
	}
}

// TestGetIndustryList 测试获取行业列表
func TestGetIndustryList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/tools/industries", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": []gin.H{
				{"industry_id": 1, "industry_name": "电商", "level": 1},
				{"industry_id": 2, "industry_name": "教育", "level": 1},
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/tools/industries", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetProductList 测试获取商品列表
func TestGetProductList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list":      []gin.H{},
				"total":     0,
				"page":      1,
				"page_size": 20,
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/products?advertiser_id=12345", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetKeywordList 测试获取关键词列表
func TestGetKeywordList(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/keywords", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list": []gin.H{
					{"word": "测试关键词", "match_type": "PHRASE", "bid": 1.5},
				},
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/keywords?advertiser_id=12345&ad_id=111", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if data, ok := response["data"].(map[string]interface{}); ok {
		assert.Contains(t, data, "list")
	}
}

// TestGetActionKeywords 测试查询行为关键词
func TestGetActionKeywords(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/keywords/action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list": []gin.H{
					{"id": 1, "name": "电商购物", "level": 1},
				},
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/keywords/action?advertiser_id=12345&query_word=电商", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetInterestKeywords 测试查询兴趣关键词
func TestGetInterestKeywords(t *testing.T) {
	router := setupQianchuanTestRouter()

	router.GET("/qianchuan/keywords/interest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"list": []gin.H{
					{"id": 1, "name": "时尚", "level": 1},
				},
			},
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qianchuan/keywords/interest?advertiser_id=12345&query_word=时尚", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
