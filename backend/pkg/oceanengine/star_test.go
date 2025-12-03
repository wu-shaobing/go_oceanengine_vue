package oceanengine

import (
	"testing"
)

// TestStarClient_NewClient 测试星图客户端创建
func TestStarClient_NewClient(t *testing.T) {
	client := NewClient("test_app", "test_secret")

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	starClient := client.Star()
	if starClient == nil {
		t.Fatal("Star() returned nil")
	}
}

// TestStarAccountInfoStruct 测试星图账户信息结构体
func TestStarAccountInfoStruct(t *testing.T) {
	info := StarAccountInfo{
		AdvertiserID:   123456,
		AdvertiserName: "测试账户",
		Balance:        10000,
		AccountType:    1,
	}

	if info.AdvertiserID != 123456 {
		t.Errorf("StarAccountInfo.AdvertiserID = %d, want 123456", info.AdvertiserID)
	}
	if info.AdvertiserName != "测试账户" {
		t.Errorf("StarAccountInfo.AdvertiserName = %s, want 测试账户", info.AdvertiserName)
	}
}

// TestStarTaskStruct 测试星图任务结构体
func TestStarTaskStruct(t *testing.T) {
	task := StarTask{
		TaskID:     111,
		TaskName:   "测试任务",
		TaskType:   1,
		TaskStatus: 2,
		Budget:     10000,
	}

	if task.TaskID != 111 {
		t.Errorf("StarTask.TaskID = %d, want 111", task.TaskID)
	}
	if task.TaskName != "测试任务" {
		t.Errorf("StarTask.TaskName = %s, want 测试任务", task.TaskName)
	}
}

// TestFundBalanceStruct 测试资金余额结构体
func TestFundBalanceStruct(t *testing.T) {
	balance := FundBalance{
		AdvertiserID: 123456,
		Balance:      5000,
		FreezeAmount: 1000,
		ValidAmount:  4000,
	}

	if balance.AdvertiserID != 123456 {
		t.Errorf("FundBalance.AdvertiserID = %d, want 123456", balance.AdvertiserID)
	}
	if balance.ValidAmount != 4000 {
		t.Errorf("FundBalance.ValidAmount = %d, want 4000", balance.ValidAmount)
	}
}

// TestDemandStruct 测试需求结构体
func TestDemandStruct(t *testing.T) {
	demand := Demand{
		DemandID:   222,
		DemandName: "测试需求",
		Status:     1,
		Budget:     5000,
	}

	if demand.DemandID != 222 {
		t.Errorf("Demand.DemandID = %d, want 222", demand.DemandID)
	}
	if demand.DemandName != "测试需求" {
		t.Errorf("Demand.DemandName = %s, want 测试需求", demand.DemandName)
	}
}

// TestStarClueStruct 测试星图线索结构体
func TestStarClueStruct(t *testing.T) {
	clue := StarClue{
		ClueID:   333,
		TaskID:   111,
		Phone:    "138****1234",
		Name:     "测试用户",
		ClueType: 1,
	}

	if clue.ClueID != 333 {
		t.Errorf("StarClue.ClueID = %d, want 333", clue.ClueID)
	}
	if clue.Name != "测试用户" {
		t.Errorf("StarClue.Name = %s, want 测试用户", clue.Name)
	}
}

// TestStarReportStruct 测试星图报表结构体
func TestStarReportStruct(t *testing.T) {
	report := StarReport{
		Date:         "2024-01-01",
		PlayCount:    10000,
		LikeCount:    500,
		CommentCount: 100,
		ShareCount:   50,
	}

	if report.PlayCount != 10000 {
		t.Errorf("StarReport.PlayCount = %d, want 10000", report.PlayCount)
	}
}

// TestStarTaskItemStruct 测试星图任务视频结构体
func TestStarTaskItemStruct(t *testing.T) {
	item := StarTaskItem{
		ItemID:     444,
		TaskID:     111,
		TalentName: "测试达人",
		PlayCount:  5000,
		LikeCount:  200,
	}

	if item.ItemID != 444 {
		t.Errorf("StarTaskItem.ItemID = %d, want 444", item.ItemID)
	}
	if item.TalentName != "测试达人" {
		t.Errorf("StarTaskItem.TalentName = %s, want 测试达人", item.TalentName)
	}
}
