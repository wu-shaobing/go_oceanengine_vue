package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetAccessToken(t *testing.T) {
	tests := []struct {
		name           string
		xAccessToken   string
		accessToken    string
		queryToken     string
		expectedResult string
	}{
		{
			name:           "X-Access-Token header takes precedence",
			xAccessToken:   "x-token-value",
			accessToken:    "access-token-value",
			queryToken:     "query-token-value",
			expectedResult: "x-token-value",
		},
		{
			name:           "Access-Token header when X-Access-Token is empty",
			xAccessToken:   "",
			accessToken:    "access-token-value",
			queryToken:     "query-token-value",
			expectedResult: "access-token-value",
		},
		{
			name:           "Query parameter when headers are empty",
			xAccessToken:   "",
			accessToken:    "",
			queryToken:     "query-token-value",
			expectedResult: "query-token-value",
		},
		{
			name:           "Empty when all sources are empty",
			xAccessToken:   "",
			accessToken:    "",
			queryToken:     "",
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/?access_token="+tt.queryToken, nil)
			if tt.xAccessToken != "" {
				req.Header.Set("X-Access-Token", tt.xAccessToken)
			}
			if tt.accessToken != "" {
				req.Header.Set("Access-Token", tt.accessToken)
			}
			c.Request = req

			result := GetAccessToken(c)
			if result != tt.expectedResult {
				t.Errorf("GetAccessToken() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestGetAdvertiserID(t *testing.T) {
	tests := []struct {
		name           string
		queryValue     string
		paramValue     string
		expectedResult uint64
	}{
		{
			name:           "Query parameter",
			queryValue:     "12345",
			paramValue:     "",
			expectedResult: 12345,
		},
		{
			name:           "Param when query is empty",
			queryValue:     "",
			paramValue:     "67890",
			expectedResult: 67890,
		},
		{
			name:           "Zero when both are empty",
			queryValue:     "",
			paramValue:     "",
			expectedResult: 0,
		},
		{
			name:           "Invalid value returns zero",
			queryValue:     "invalid",
			paramValue:     "",
			expectedResult: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/?advertiser_id="+tt.queryValue, nil)
			c.Request = req
			if tt.paramValue != "" {
				c.Params = gin.Params{gin.Param{Key: "advertiser_id", Value: tt.paramValue}}
			}

			result := GetAdvertiserID(c)
			if result != tt.expectedResult {
				t.Errorf("GetAdvertiserID() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestGetAccountID(t *testing.T) {
	tests := []struct {
		name           string
		accountID      string
		openID         string
		expectedResult string
	}{
		{
			name:           "account_id takes precedence",
			accountID:      "account123",
			openID:         "open456",
			expectedResult: "account123",
		},
		{
			name:           "open_id when account_id is empty",
			accountID:      "",
			openID:         "open456",
			expectedResult: "open456",
		},
		{
			name:           "Empty when both are empty",
			accountID:      "",
			openID:         "",
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := "/?"
			if tt.accountID != "" {
				url += "account_id=" + tt.accountID + "&"
			}
			if tt.openID != "" {
				url += "open_id=" + tt.openID
			}
			req, _ := http.NewRequest("GET", url, nil)
			c.Request = req

			result := GetAccountID(c)
			if result != tt.expectedResult {
				t.Errorf("GetAccountID() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
