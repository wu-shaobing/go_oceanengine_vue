package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetAccessToken 从请求中统一获取 access_token
// 按优先级从以下位置获取：
// 1. X-Access-Token header
// 2. Access-Token header
// 3. access_token query parameter
func GetAccessToken(c *gin.Context) string {
	token := c.GetHeader("X-Access-Token")
	if token == "" {
		token = c.GetHeader("Access-Token")
	}
	if token == "" {
		token = c.Query("access_token")
	}
	return token
}

// GetAdvertiserID 从请求中获取 advertiser_id
// 从 query 或 path parameter 中获取
func GetAdvertiserID(c *gin.Context) uint64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

// GetAdvertiserIDInt64 从请求中获取 advertiser_id (int64类型)
func GetAdvertiserIDInt64(c *gin.Context) int64 {
	idStr := c.Query("advertiser_id")
	if idStr == "" {
		idStr = c.Param("advertiser_id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// GetAccountID 从请求中获取 account_id (通常用于企业号)
// 支持 account_id 和 open_id 两种参数名
func GetAccountID(c *gin.Context) string {
	accountID := c.Query("account_id")
	if accountID == "" {
		accountID = c.Query("open_id")
	}
	return accountID
}
