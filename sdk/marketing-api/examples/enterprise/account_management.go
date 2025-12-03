// Package main 演示企业号账号管理功能
//
// 本示例展示如何使用 SDK 完成以下操作:
// 1. 获取企业号账号信息
// 2. 获取运营数据概览
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/api/enterprise"
	"github.com/bububa/oceanengine/marketing-api/core"
	enterpriseModel "github.com/bububa/oceanengine/marketing-api/model/enterprise"
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

	// 企业号 DouYin ID
	douyinIDs := []string{"your_douyin_id"} // 替换为实际的抖音号

	// 2. 获取企业号信息
	fmt.Println("=== 获取企业号信息 ===")
	infoReq := &enterpriseModel.InfoRequest{
		EDouyinIDs: douyinIDs,
	}

	infos, err := enterprise.Info(ctx, client, accessToken, infoReq)
	if err != nil {
		log.Printf("获取企业号信息失败: %v\n", err)
	} else {
		for _, info := range infos {
			fmt.Printf("企业号: %s\n", info.EDouyinID)
		}
	}

	// 3. 获取运营数据概览
	fmt.Println("\n=== 获取运营数据概览 ===")
	// 广告主ID (替换为实际的广告主ID)
	advertiserID := uint64(1234567890)
	reportReq := &enterpriseModel.ReportRequest{
		AdvertiserID: advertiserID,
		StartTime:    "2024-01-01",
		EndTime:      "2024-01-31",
	}

	reportData, err := enterprise.OverviewDataGet(ctx, client, accessToken, reportReq)
	if err != nil {
		log.Printf("获取运营数据失败: %v\n", err)
	} else {
		fmt.Printf("运营数据: %+v\n", reportData)
	}

	fmt.Println("\n企业号示例运行完成!")
}

// ExampleItemList 获取视频列表示例
func ExampleItemList(ctx context.Context, client *core.SDKClient, accessToken, douyinID string) {
	// 获取企业号视频列表:
	// 使用 enterprise.ItemList 获取视频列表

	fmt.Println("视频列表示例 - 请参考API文档配置完整参数")
}

// ExampleBindList 获取绑定列表示例
func ExampleBindList(ctx context.Context, client *core.SDKClient, accessToken string) {
	// 获取企业号绑定列表:
	// 使用 enterprise.BindListGet 获取绑定关系

	fmt.Println("绑定列表示例 - 请参考API文档配置完整参数")
}
