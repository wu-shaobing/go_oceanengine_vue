// Package main 演示巨量千川电商广告投放流程
//
// 本示例展示如何使用 SDK 完成以下操作:
// 1. 商品管理 - 获取商品列表
// 2. 广告组管理 - 创建/查询广告组
// 3. 广告计划管理 - 创建/查询广告计划
// 4. 数据报表 - 查询广告效果数据
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/api/qianchuan/ad"
	"github.com/bububa/oceanengine/marketing-api/api/qianchuan/campaign"
	"github.com/bububa/oceanengine/marketing-api/api/qianchuan/product"
	"github.com/bububa/oceanengine/marketing-api/api/qianchuan/report"
	"github.com/bububa/oceanengine/marketing-api/api/qianchuan/shop"
	"github.com/bububa/oceanengine/marketing-api/core"
	shopModel "github.com/bububa/oceanengine/marketing-api/model/qianchuan/shop"
	adModel "github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
	campaignModel "github.com/bububa/oceanengine/marketing-api/model/qianchuan/campaign"
	productModel "github.com/bububa/oceanengine/marketing-api/model/qianchuan/product"
	reportModel "github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
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

	// 广告主ID (千川账户ID)
	advertiserID := uint64(1234567890) // 替换为实际的广告主ID

	// 2. 获取店铺授权列表
	fmt.Println("=== 获取店铺授权列表 ===")
	shopReq := &shopModel.AuthorizedGetRequest{
		AdvertiserID: advertiserID,
	}
	shopResp, err := shop.AuthorizedGet(ctx, client, accessToken, shopReq)
	if err != nil {
		log.Printf("获取店铺列表失败: %v\n", err)
	} else if shopResp != nil {
		for _, s := range shopResp.ShopList {
			fmt.Printf("店铺: %s (ID: %d)\n", s.ShopName, s.ShopID)
		}
	}

	// 3. 获取商品列表
	fmt.Println("\n=== 获取商品列表 ===")
	productReq := &productModel.AvailableGetRequest{
		AdvertiserID: advertiserID,
		Page:         1,
		PageSize:     20,
	}

	productResp, err := product.AvailableGet(ctx, client, accessToken, productReq)
	if err != nil {
		log.Printf("获取商品列表失败: %v\n", err)
	} else {
		fmt.Printf("商品总数: %d\n", productResp.PageInfo.TotalNumber)
		for _, p := range productResp.ProductList {
			fmt.Printf("  - %s (ID: %d)\n", p.Name, p.ID)
		}
	}

	// 4. 创建广告组
	fmt.Println("\n=== 创建广告组 ===")
	campaignReq := &campaignModel.CreateRequest{
		AdvertiserID: advertiserID,
		CampaignName: "SDK测试-电商推广",
		// MarketingGoal: "LIVE_PROM_GOODS", // 直播带货
	}

	campaignID, err := campaign.Create(ctx, client, accessToken, campaignReq)
	if err != nil {
		log.Printf("创建广告组失败: %v (取决于账户权限)\n", err)
	} else {
		fmt.Printf("创建广告组成功, ID: %d\n", campaignID)
	}

	// 5. 查询广告组列表
	fmt.Println("\n=== 查询广告组列表 ===")
	campaignListReq := &campaignModel.ListGetRequest{
		AdvertiserID: advertiserID,
		Page:         1,
		PageSize:     10,
	}

	campaignResp, err := campaign.ListGet(ctx, client, accessToken, campaignListReq)
	if err != nil {
		log.Printf("查询广告组失败: %v\n", err)
	} else {
		fmt.Printf("广告组数量: %d\n", len(campaignResp.List))
		for _, c := range campaignResp.List {
			fmt.Printf("  - %s (ID: %d)\n", c.Name, c.ID)
		}
	}

	// 6. 查询广告计划列表
	fmt.Println("\n=== 查询广告计划列表 ===")
	adListReq := &adModel.GetRequest{
		AdvertiserID: advertiserID,
		Page:         1,
		PageSize:     10,
	}

	adResp, err := ad.Get(ctx, client, accessToken, adListReq)
	if err != nil {
		log.Printf("查询广告计划失败: %v\n", err)
	} else {
		fmt.Printf("广告计划数量: %d\n", len(adResp.List))
	}

	// 7. 获取广告数据报表
	fmt.Println("\n=== 获取广告数据报表 ===")
	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-01-31")
	reportReq := &reportModel.GetRequest{
		AdvertiserID: advertiserID,
		StartDate:    startDate,
		EndDate:      endDate,
		Fields:       []string{"stat_cost", "show_cnt", "click_cnt", "pay_order_count"},
	}

	reportData, err := report.AdGet(ctx, client, accessToken, reportReq)
	if err != nil {
		log.Printf("获取报表失败: %v\n", err)
	} else {
		fmt.Printf("报表数据条数: %d\n", len(reportData.List))
	}

	fmt.Println("\n千川示例运行完成!")
}

// ExampleCreateLiveAd 创建直播间推广广告示例
func ExampleCreateLiveAd(ctx context.Context, client *core.SDKClient, accessToken string, advertiserID uint64) {
	// 直播间推广需要配置:
	// 1. 选择推广目标 (直播间成交/直播间观看等)
	// 2. 设置定向人群
	// 3. 设置出价和预算
	// 4. 关联直播间

	// adCreateReq := &adModel.CreateRequest{
	//     AdvertiserID: advertiserID,
	//     MarketingGoal: "LIVE_PROM_GOODS",
	//     // ... 其他配置
	// }
	//
	// adID, err := ad.Create(ctx, client, accessToken, adCreateReq)

	fmt.Println("直播间推广示例 - 请参考API文档配置完整参数")
}

// ExampleProductAd 商品推广广告示例
func ExampleProductAd(ctx context.Context, client *core.SDKClient, accessToken string, advertiserID uint64) {
	// 商品推广需要:
	// 1. 获取可投放商品列表
	// 2. 选择商品创建计划
	// 3. 设置投放策略

	fmt.Println("商品推广示例 - 请参考API文档配置完整参数")
}
