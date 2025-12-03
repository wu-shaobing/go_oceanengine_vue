package oceanengine

import (
	"testing"
)

// TestQianchuanClient_NewClient 测试千川客户端创建
func TestQianchuanClient_NewClient(t *testing.T) {
	client := NewClient("test_app", "test_secret")

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	qcClient := client.Qianchuan()
	if qcClient == nil {
		t.Fatal("Qianchuan() returned nil")
	}
}

// TestKeywordStruct 测试关键词结构体
func TestKeywordStruct(t *testing.T) {
	keyword := Keyword{
		Word:      "测试关键词",
		MatchType: "PHRASE",
		Bid:       1.5,
		Status:    "ENABLE",
	}

	if keyword.Word != "测试关键词" {
		t.Errorf("Keyword.Word = %s, want 测试关键词", keyword.Word)
	}
	if keyword.MatchType != "PHRASE" {
		t.Errorf("Keyword.MatchType = %s, want PHRASE", keyword.MatchType)
	}
}

// TestQianchuanReportStruct 测试报表结构体
func TestQianchuanReportStruct(t *testing.T) {
	report := QianchuanReport{
		StatDatetime: "2024-01-01",
		Cost:         100.5,
		ShowCnt:      1000,
		ClickCnt:     50,
	}

	if report.Cost != 100.5 {
		t.Errorf("QianchuanReport.Cost = %f, want 100.5", report.Cost)
	}
}

// TestQianchuanCampaignStruct 测试广告组结构体
func TestQianchuanCampaignStruct(t *testing.T) {
	campaign := QianchuanCampaign{
		CampaignID:   111,
		CampaignName: "Test Campaign",
		AdvertiserID: 123456,
		Status:       "ENABLE",
	}

	if campaign.CampaignID != 111 {
		t.Errorf("QianchuanCampaign.CampaignID = %d, want 111", campaign.CampaignID)
	}
}

// TestQianchuanAdStruct 测试广告计划结构体
func TestQianchuanAdStruct(t *testing.T) {
	ad := QianchuanAd{
		AdID:         222,
		AdName:       "Test Ad",
		CampaignID:   111,
		AdvertiserID: 123456,
		Status:       "AD_STATUS_ENABLE",
	}

	if ad.AdID != 222 {
		t.Errorf("QianchuanAd.AdID = %d, want 222", ad.AdID)
	}
}

// TestQianchuanAudienceStruct 测试人群包结构体
func TestQianchuanAudienceStruct(t *testing.T) {
	audience := QianchuanAudience{
		AudienceID:   333,
		Name:         "Test Audience",
		AudienceType: 1,
		CoverNum:     10000,
		Status:       1,
	}

	if audience.AudienceID != 333 {
		t.Errorf("QianchuanAudience.AudienceID = %d, want 333", audience.AudienceID)
	}
}
