// Package main 演示巨量引擎广告投放基础流程
//
// 本示例展示如何使用 SDK 完成以下操作:
// 1. OAuth2 授权获取 AccessToken
// 2. 获取广告主信息
// 3. 创建广告组
// 4. 创建广告计划
// 5. 查询广告数据报表
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/api/ad"
	"github.com/bububa/oceanengine/marketing-api/api/advertiser"
	"github.com/bububa/oceanengine/marketing-api/api/campaign"
	"github.com/bububa/oceanengine/marketing-api/api/oauth"
	"github.com/bububa/oceanengine/marketing-api/api/report"
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	advertiserModel "github.com/bububa/oceanengine/marketing-api/model/advertiser"
	campaignModel "github.com/bububa/oceanengine/marketing-api/model/campaign"
	reportModel "github.com/bububa/oceanengine/marketing-api/model/report"
)

func main() {
	// 1. 初始化 SDK 客户端
	// 从环境变量获取配置，生产环境请妥善管理凭证
	appIDStr := os.Getenv("OCEANENGINE_APP_ID")
	appSecret := os.Getenv("OCEANENGINE_APP_SECRET")

	if appIDStr == "" || appSecret == "" {
		log.Fatal("请设置环境变量 OCEANENGINE_APP_ID 和 OCEANENGINE_APP_SECRET")
	}

	appID, err := strconv.ParseUint(appIDStr, 10, 64)
	if err != nil {
		log.Fatalf("OCEANENGINE_APP_ID 必须是数字: %v", err)
	}

	client := core.NewSDKClient(appID, appSecret)
	ctx := context.Background()

	// 2. 生成授权链接 (用户需要通过此链接授权)
	redirectURL := "https://your-callback-url.com/oauth/callback"
	state := "your_random_state_string"
	authURL := oauth.Url(client, redirectURL, state, false)
	fmt.Printf("请访问以下链接进行授权:\n%s\n\n", authURL)

	// 3. 使用授权码获取 AccessToken (实际场景中从回调获取 authCode)
	// authCode := "从回调URL获取的授权码"
	// tokenResp, err := oauth.AccessToken(ctx, client, authCode)
	// if err != nil {
	//     log.Fatalf("获取AccessToken失败: %v", err)
	// }
	// accessToken := tokenResp.AccessToken

	// 以下为演示代码，实际使用时需要真实的 accessToken
	accessToken := os.Getenv("OCEANENGINE_ACCESS_TOKEN")
	if accessToken == "" {
		log.Println("提示: 设置 OCEANENGINE_ACCESS_TOKEN 环境变量后可运行完整示例")
		return
	}

	// 4. 获取已授权的广告主列表
	advertisers, err := oauth.AdvertiserGet(ctx, client, accessToken)
	if err != nil {
		log.Fatalf("获取广告主列表失败: %v", err)
	}

	if len(advertisers) == 0 {
		log.Fatal("没有已授权的广告主")
	}

	advertiserID := advertisers[0].AdvertiserID
	fmt.Printf("使用广告主ID: %d\n", advertiserID)

	// 5. 获取广告主详细信息
	infoReq := &advertiserModel.InfoRequest{
		AdvertiserIDs: []uint64{advertiserID},
	}
	infos, err := advertiser.Info(ctx, client, accessToken, infoReq)
	if err != nil {
		log.Fatalf("获取广告主信息失败: %v", err)
	}

	for _, info := range infos {
		fmt.Printf("广告主: %s (ID: %d)\n", info.Name, info.ID)
	}

	// 6. 创建广告组示例
	campaignReq := &campaignModel.CreateRequest{
		AdvertiserID: advertiserID,
		CampaignName: "SDK测试广告组",
		CampaignType: enum.CampaignType_FEED, // 信息流广告
		LandingType:  enum.APP,               // 应用推广
	}

	campaignID, err := campaign.Create(ctx, client, accessToken, campaignReq)
	if err != nil {
		log.Printf("创建广告组失败: %v (这可能是正常的，取决于账户权限)\n", err)
	} else {
		fmt.Printf("创建广告组成功, ID: %d\n", campaignID)
	}

	// 7. 查询广告数据报表
	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-01-31")
	reportReq := &reportModel.GetRequest{
		AdvertiserID:    advertiserID,
		StartDate:       startDate,
		EndDate:         endDate,
		TimeGranularity: enum.STAT_TIME_GRANULARITY_DAILY,
		GroupBy:         []enum.StatGroupBy{enum.STAT_GROUP_BY_FIELD_STAT_TIME},
	}

	reportData, err := report.AdvertiserGet(ctx, client, accessToken, reportReq)
	if err != nil {
		log.Printf("获取报表数据失败: %v\n", err)
	} else {
		fmt.Printf("报表数据条数: %d\n", len(reportData.List))
	}

	// 8. 刷新 Token (AccessToken 即将过期时使用)
	// refreshToken := tokenResp.RefreshToken
	// newTokenResp, err := oauth.RefreshToken(ctx, client, refreshToken)

	fmt.Println("\n示例运行完成!")
}

// ExampleCreateAd 创建广告计划示例
func ExampleCreateAd(ctx context.Context, client *core.SDKClient, accessToken string, advertiserID, campaignID uint64) {
	// 注意: 这是一个简化的示例，实际创建广告需要更多参数
	// 请参考 API 文档了解完整的参数要求

	_ = ad.Create // 避免未使用警告

	// adReq := &adModel.CreateRequest{
	//     AdvertiserID: advertiserID,
	//     CampaignID:   campaignID,
	//     Name:         "SDK测试计划",
	//     // ... 其他必要参数
	// }
	//
	// adID, err := ad.Create(ctx, client, accessToken, adReq)
	// if err != nil {
	//     log.Printf("创建广告计划失败: %v", err)
	//     return
	// }
	// fmt.Printf("创建广告计划成功, ID: %d\n", adID)
}
