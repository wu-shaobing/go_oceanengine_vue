// Package main 演示本地推广告投放流程
//
// 本示例展示如何使用 SDK 完成以下操作:
// 1. 线索管理 - 获取本地推线索
// 2. POI 门店管理
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/api/local/clue"
	"github.com/bububa/oceanengine/marketing-api/api/local/poi"
	"github.com/bububa/oceanengine/marketing-api/core"
	clueModel "github.com/bububa/oceanengine/marketing-api/model/local/clue"
	poiModel "github.com/bububa/oceanengine/marketing-api/model/local/poi"
)

func main() {
	// 1. 初始化 SDK 客户端
	appIDStr := os.Getenv("OCEANENGINE_APP_ID")
	appSecret := os.Getenv("OCEANENGINE_APP_SECRET")
	accessToken := os.Getenv("OCEANENGINE_ACCESS_TOKEN")

	if appIDStr == "" || appSecret == "" {
		log.Fatal("请设置环境变量 OCEANENGINE_APP_ID 和 OCEANENGINE_APP_SECRET")
	}

	appID, err := strconv.ParseUint(appIDStr, 10, 64)
	if err != nil {
		log.Fatalf("OCEANENGINE_APP_ID 必须是数字: %v", err)
	}

	client := core.NewSDKClient(appID, appSecret)
	ctx := context.Background()

	if accessToken == "" {
		log.Println("提示: 设置 OCEANENGINE_ACCESS_TOKEN 环境变量后可运行完整示例")
		return
	}

	// 广告主ID
	advertiserID := uint64(1234567890) // 替换为实际的广告主ID

	// 2. 获取本地推线索列表
	fmt.Println("=== 获取本地推线索 ===")
	clueReq := &clueModel.LifeGetRequest{
		LocalAccountIDs: []uint64{advertiserID},
		StartTime:       "2024-01-01 00:00:00",
		EndTime:         "2024-01-31 23:59:59",
		Page:            1,
		PageSize:        20,
	}

	clueResp, err := clue.LifeGet(ctx, client, accessToken, clueReq)
	if err != nil {
		log.Printf("获取线索列表失败: %v\n", err)
	} else if clueResp.PageInfo != nil {
		fmt.Printf("线索总数: %d\n", clueResp.PageInfo.TotalNumber)
	}

	// 3. 根据多门店ID获取门店ID
	fmt.Println("\n=== 获取门店ID ===")
	poiReq := &poiModel.MultiPoiIDsGetRequest{
		LocalAccountID: advertiserID,
		MultiPoiIDs:    []uint64{123456}, // 替换为实际的多门店ID
	}

	poiResp, err := poi.MultiPoiIDsGet(ctx, client, accessToken, poiReq)
	if err != nil {
		log.Printf("获取门店ID失败: %v\n", err)
	} else {
		fmt.Printf("门店信息: %+v\n", poiResp.MultiPoiInfo)
	}

	fmt.Println("\n本地推示例运行完成!")
}

// ExampleLocalPromotion 本地推广广告推广示例
func ExampleLocalPromotion(ctx context.Context, client *core.SDKClient, accessToken string, advertiserID uint64) {
	// 本地推广广告特点:
	// 1. 基于地理位置定向
	// 2. 关联 POI 门店信息
	// 3. 支持团购/优惠券推广
	// 4. 到店转化追踪

	fmt.Println("本地推广广告示例 - 请参考API文档配置完整参数")
}
