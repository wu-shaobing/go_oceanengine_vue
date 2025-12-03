// Package main 演示星图达人营销功能
//
// 本示例展示如何使用 SDK 完成以下操作:
// 1. 获取星图账户信息
// 2. 获取星图订单线索
// 3. 获取需求列表
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/api/star"
	"github.com/bububa/oceanengine/marketing-api/core"
	starModel "github.com/bububa/oceanengine/marketing-api/model/star"
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
	_ = appID // 未使用，示例中仅用于验证

	client := core.NewSDKClient(appID, appSecret)
	ctx := context.Background()

	if accessToken == "" {
		log.Println("提示: 设置 OCEANENGINE_ACCESS_TOKEN 环境变量后可运行完整示例")
		return
	}

	// 星图账户ID
	starID := uint64(1234567890) // 替换为实际的账户ID
	orderID := uint64(9876543210) // 替换为实际的订单ID

	// 2. 获取星图账户信息
	fmt.Println("=== 获取星图账户信息 ===")
	infoReq := &starModel.InfoRequest{
		StarIDs: []uint64{starID},
	}

	info, err := star.Info(ctx, client, accessToken, infoReq)
	if err != nil {
		log.Printf("获取星图信息失败: %v\n", err)
	} else {
		fmt.Printf("星图账户: %+v\n", info)
	}

	// 3. 获取订单线索列表
	fmt.Println("\n=== 获取订单线索 ===")
	clueReq := &starModel.ClueGetRequest{
		StarID:  starID,
		OrderID: orderID,
	}

	clueResp, err := star.ClueList(ctx, client, accessToken, clueReq)
	if err != nil {
		log.Printf("获取线索列表失败: %v\n", err)
	} else {
		fmt.Printf("线索数量: %d\n", len(clueResp.List))
	}

	// 4. 获取需求列表
	fmt.Println("\n=== 获取需求列表 ===")
	demandReq := &starModel.DemandListRequest{
		StarID:   starID,
		Page:     1,
		PageSize: 10,
	}

	demandResp, err := star.DemandList(ctx, client, accessToken, demandReq)
	if err != nil {
		log.Printf("获取需求列表失败: %v\n", err)
	} else {
		fmt.Printf("需求数量: %d\n", len(demandResp.List))
	}

	fmt.Println("\n星图示例运行完成!")
}

// ExampleDemandOrderList 获取需求订单列表示例
func ExampleDemandOrderList(ctx context.Context, client *core.SDKClient, accessToken string, starID uint64) {
	// 需求订单查询:
	// 使用 star.DemandOrderList 获取需求订单

	fmt.Println("需求订单示例 - 请参考API文档配置完整参数")
}

// ExampleReportData 数据报表示例
func ExampleReportData(ctx context.Context, client *core.SDKClient, accessToken string, starID uint64) {
	// 数据报表查询:
	// 1. star.ReportOrderOverviewGet - 订单概览
	// 2. star.ReportOrderUserDistributionGet - 用户分布
	// 3. star.ReportDataTopicConfig - 话题配置

	fmt.Println("数据报表示例 - 请参考API文档配置完整参数")
}
