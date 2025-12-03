package dto

// ReportQueryReq 报告查询请求
type ReportQueryReq struct {
	AdvertiserID uint64 `form:"advertiser_id" binding:"required"`
	StartDate    string `form:"start_date" binding:"required"`
	EndDate      string `form:"end_date" binding:"required"`
	GroupBy      string `form:"group_by"` // STAT_DATE 按日期汇总
	CampaignID   uint64 `form:"campaign_id"`
	AdID         uint64 `form:"ad_id"`
}

// ReportSummaryResp 报告汇总响应
type ReportSummaryResp struct {
	Cost        float64 `json:"cost"`
	Show        int64   `json:"show"`
	Click       int64   `json:"click"`
	Convert     int64   `json:"convert"`
	CTR         float64 `json:"ctr"`
	CVR         float64 `json:"cvr"`
	CPM         float64 `json:"cpm"`
	CPC         float64 `json:"cpc"`
	ConvertCost float64 `json:"convert_cost"`
}

// ReportDetailResp 报告明细响应
type ReportDetailResp struct {
	StatDate    string  `json:"stat_date"`
	Cost        float64 `json:"cost"`
	Show        int64   `json:"show"`
	Click       int64   `json:"click"`
	Convert     int64   `json:"convert"`
	CTR         float64 `json:"ctr"`
	CVR         float64 `json:"cvr"`
	CPM         float64 `json:"cpm"`
	CPC         float64 `json:"cpc"`
	ConvertCost float64 `json:"convert_cost"`
}

// CampaignReportResp 广告系列报告响应
type CampaignReportResp struct {
	CampaignID   uint64  `json:"campaign_id"`
	CampaignName string  `json:"campaign_name"`
	StatDate     string  `json:"stat_date"`
	Cost         float64 `json:"cost"`
	Show         int64   `json:"show"`
	Click        int64   `json:"click"`
	Convert      int64   `json:"convert"`
	CTR          float64 `json:"ctr"`
	CVR          float64 `json:"cvr"`
	CPM          float64 `json:"cpm"`
	CPC          float64 `json:"cpc"`
	ConvertCost  float64 `json:"convert_cost"`
}

// AdReportResp 广告组报告响应
type AdReportResp struct {
	AdID        uint64  `json:"ad_id"`
	AdName      string  `json:"ad_name"`
	StatDate    string  `json:"stat_date"`
	Cost        float64 `json:"cost"`
	Show        int64   `json:"show"`
	Click       int64   `json:"click"`
	Convert     int64   `json:"convert"`
	CTR         float64 `json:"ctr"`
	CVR         float64 `json:"cvr"`
	CPM         float64 `json:"cpm"`
	CPC         float64 `json:"cpc"`
	ConvertCost float64 `json:"convert_cost"`
}

// ReportSyncReq 报告同步请求
type ReportSyncReq struct {
	AdvertiserID uint64 `json:"advertiser_id" binding:"required"`
	StartDate    string `json:"start_date" binding:"required"`
	EndDate      string `json:"end_date" binding:"required"`
	Dimension    string `json:"dimension"` // ADVERTISER, CAMPAIGN, AD
}

// ReportSyncResp 报告同步响应
type ReportSyncResp struct {
	SyncCount int    `json:"sync_count"`
	SyncAt    string `json:"sync_at"`
}

// ExportTaskListReq 导出任务列表请求
type ExportTaskListReq struct {
	AdvertiserID uint64 `form:"advertiser_id"`
	Status       string `form:"status"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// ExportTaskResp 导出任务响应
type ExportTaskResp struct {
	ID           uint64 `json:"id"`
	AdvertiserID uint64 `json:"advertiser_id"`
	TaskType     string `json:"task_type"`
	Status       string `json:"status"`
	FileName     string `json:"file_name"`
	FileSize     int64  `json:"file_size"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	ErrorMsg     string `json:"error_msg"`
	CreatedAt    string `json:"created_at"`
}

// ExportCreateReq 创建导出任务请求
type ExportCreateReq struct {
	AdvertiserID uint64 `json:"advertiser_id" binding:"required"`
	TaskType     string `json:"task_type" binding:"required,oneof=ADVERTISER CAMPAIGN AD"`
	StartDate    string `json:"start_date" binding:"required"`
	EndDate      string `json:"end_date" binding:"required"`
}

// CreativeReportResp 创意报告响应
type CreativeReportResp struct {
	CreativeID  uint64  `json:"creative_id"`
	StatDate    string  `json:"stat_date"`
	Cost        float64 `json:"cost"`
	Show        int64   `json:"show"`
	Click       int64   `json:"click"`
	Convert     int64   `json:"convert"`
	CTR         float64 `json:"ctr"`
	CVR         float64 `json:"cvr"`
	CPM         float64 `json:"cpm"`
	CPC         float64 `json:"cpc"`
	ConvertCost float64 `json:"convert_cost"`
}

// RealtimeReportReq 实时数据请求
type RealtimeReportReq struct {
	AdvertiserID uint64   `form:"advertiser_id" binding:"required"`
	Level        string   `form:"level" binding:"required,oneof=advertiser campaign ad"`
	IDs          []uint64 `form:"ids"`
}

// RealtimeReportResp 实时数据响应
type RealtimeReportResp struct {
	ID         uint64  `json:"id,omitempty"`
	TodayCost  float64 `json:"today_cost"`
	TodayShow  int64   `json:"today_show"`
	TodayClick int64   `json:"today_click"`
	TodayConv  int64   `json:"today_convert"`
	TodayCTR   float64 `json:"today_ctr"`
	TodayCVR   float64 `json:"today_cvr"`
	UpdatedAt  string  `json:"updated_at"`
}
