package repository

import (
	"context"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/report/dto"
	"oceanengine-backend/internal/app/report/model"
)

// ReportRepository 报告仓储接口
type ReportRepository interface {
	// 广告主报告
	GetAdvertiserReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.AdvertiserReport, error)
	GetAdvertiserSummary(ctx context.Context, req *dto.ReportQueryReq) (*dto.ReportSummaryResp, error)
	SaveAdvertiserReport(ctx context.Context, report *model.AdvertiserReport) error

	// 广告系列报告
	GetCampaignReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.CampaignReport, error)
	SaveCampaignReport(ctx context.Context, report *model.CampaignReport) error

	// 广告组报告
	GetAdReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.AdReport, error)
	SaveAdReport(ctx context.Context, report *model.AdReport) error

	// 导出任务
	GetExportTaskList(ctx context.Context, req *dto.ExportTaskListReq) ([]*model.ExportTask, int64, error)
	GetExportTaskByID(ctx context.Context, id uint64) (*model.ExportTask, error)
	CreateExportTask(ctx context.Context, task *model.ExportTask) error
	UpdateExportTask(ctx context.Context, task *model.ExportTask) error
}

type reportRepository struct {
	db *gorm.DB
}

// NewReportRepository 创建报告仓储
func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

// GetAdvertiserReport 获取广告主报告
func (r *reportRepository) GetAdvertiserReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.AdvertiserReport, error) {
	var list []*model.AdvertiserReport

	query := r.db.WithContext(ctx).Model(&model.AdvertiserReport{}).
		Where("advertiser_id = ?", req.AdvertiserID).
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if err := query.Order("stat_date ASC").Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// GetAdvertiserSummary 获取广告主汇总报告
func (r *reportRepository) GetAdvertiserSummary(ctx context.Context, req *dto.ReportQueryReq) (*dto.ReportSummaryResp, error) {
	var result dto.ReportSummaryResp

	err := r.db.WithContext(ctx).Model(&model.AdvertiserReport{}).
		Select("COALESCE(SUM(cost), 0) as cost, COALESCE(SUM(show), 0) as show, COALESCE(SUM(click), 0) as click, COALESCE(SUM(convert), 0) as convert").
		Where("advertiser_id = ?", req.AdvertiserID).
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	// 计算派生指标
	if result.Show > 0 {
		result.CTR = float64(result.Click) / float64(result.Show) * 100
		result.CPM = result.Cost / float64(result.Show) * 1000
	}
	if result.Click > 0 {
		result.CVR = float64(result.Convert) / float64(result.Click) * 100
		result.CPC = result.Cost / float64(result.Click)
	}
	if result.Convert > 0 {
		result.ConvertCost = result.Cost / float64(result.Convert)
	}

	return &result, nil
}

// SaveAdvertiserReport 保存广告主报告
func (r *reportRepository) SaveAdvertiserReport(ctx context.Context, report *model.AdvertiserReport) error {
	return r.db.WithContext(ctx).
		Where("advertiser_id = ? AND stat_date = ?", report.AdvertiserID, report.StatDate).
		Assign(report).
		FirstOrCreate(report).Error
}

// GetCampaignReport 获取广告系列报告
func (r *reportRepository) GetCampaignReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.CampaignReport, error) {
	var list []*model.CampaignReport

	query := r.db.WithContext(ctx).Model(&model.CampaignReport{}).
		Where("advertiser_id = ?", req.AdvertiserID).
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if req.CampaignID > 0 {
		query = query.Where("campaign_id = ?", req.CampaignID)
	}

	if err := query.Order("stat_date ASC, campaign_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// SaveCampaignReport 保存广告系列报告
func (r *reportRepository) SaveCampaignReport(ctx context.Context, report *model.CampaignReport) error {
	return r.db.WithContext(ctx).
		Where("campaign_id = ? AND stat_date = ?", report.CampaignID, report.StatDate).
		Assign(report).
		FirstOrCreate(report).Error
}

// GetAdReport 获取广告组报告
func (r *reportRepository) GetAdReport(ctx context.Context, req *dto.ReportQueryReq) ([]*model.AdReport, error) {
	var list []*model.AdReport

	query := r.db.WithContext(ctx).Model(&model.AdReport{}).
		Where("advertiser_id = ?", req.AdvertiserID).
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if req.CampaignID > 0 {
		query = query.Where("campaign_id = ?", req.CampaignID)
	}
	if req.AdID > 0 {
		query = query.Where("ad_id = ?", req.AdID)
	}

	if err := query.Order("stat_date ASC, ad_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// SaveAdReport 保存广告组报告
func (r *reportRepository) SaveAdReport(ctx context.Context, report *model.AdReport) error {
	return r.db.WithContext(ctx).
		Where("ad_id = ? AND stat_date = ?", report.AdID, report.StatDate).
		Assign(report).
		FirstOrCreate(report).Error
}

// GetExportTaskList 获取导出任务列表
func (r *reportRepository) GetExportTaskList(ctx context.Context, req *dto.ExportTaskListReq) ([]*model.ExportTask, int64, error) {
	var list []*model.ExportTask
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ExportTask{})

	if req.AdvertiserID > 0 {
		query = query.Where("advertiser_id = ?", req.AdvertiserID)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetExportTaskByID 根据ID获取导出任务
func (r *reportRepository) GetExportTaskByID(ctx context.Context, id uint64) (*model.ExportTask, error) {
	var task model.ExportTask
	if err := r.db.WithContext(ctx).First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateExportTask 创建导出任务
func (r *reportRepository) CreateExportTask(ctx context.Context, task *model.ExportTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

// UpdateExportTask 更新导出任务
func (r *reportRepository) UpdateExportTask(ctx context.Context, task *model.ExportTask) error {
	return r.db.WithContext(ctx).Save(task).Error
}
