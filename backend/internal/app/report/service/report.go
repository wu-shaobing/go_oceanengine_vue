package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"oceanengine-backend/config"
	advRepo "oceanengine-backend/internal/app/advertiser/repository"
	"oceanengine-backend/internal/app/report/dto"
	"oceanengine-backend/internal/app/report/model"
	"oceanengine-backend/internal/app/report/repository"
	"oceanengine-backend/pkg/errcode"
	"oceanengine-backend/pkg/oceanengine"
)

// ReportService 报告服务
type ReportService struct {
	repo     repository.ReportRepository
	advRepo  advRepo.AdvertiserRepository
	oceanCfg *config.OceanConfig
}

// NewReportService 创建报告服务
func NewReportService(db *gorm.DB, oceanCfg *config.OceanConfig) *ReportService {
	return &ReportService{
		repo:     repository.NewReportRepository(db),
		advRepo:  advRepo.NewAdvertiserRepository(db),
		oceanCfg: oceanCfg,
	}
}

// GetAdvertiserReport 获取广告主报告
func (s *ReportService) GetAdvertiserReport(ctx context.Context, req *dto.ReportQueryReq) ([]*dto.ReportDetailResp, error) {
	list, err := s.repo.GetAdvertiserReport(ctx, req)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.ReportDetailResp, len(list))
	for i, r := range list {
		result[i] = &dto.ReportDetailResp{
			StatDate:    r.StatDate,
			Cost:        r.Cost,
			Show:        r.Show,
			Click:       r.Click,
			Convert:     r.Convert,
			CTR:         r.CTR,
			CVR:         r.CVR,
			CPM:         r.CPM,
			CPC:         r.CPC,
			ConvertCost: r.ConvertCost,
		}
	}

	return result, nil
}

// GetAdvertiserSummary 获取广告主汇总报告
func (s *ReportService) GetAdvertiserSummary(ctx context.Context, req *dto.ReportQueryReq) (*dto.ReportSummaryResp, error) {
	result, err := s.repo.GetAdvertiserSummary(ctx, req)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return result, nil
}

// GetCampaignReport 获取广告系列报告
func (s *ReportService) GetCampaignReport(ctx context.Context, req *dto.ReportQueryReq) ([]*dto.CampaignReportResp, error) {
	list, err := s.repo.GetCampaignReport(ctx, req)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.CampaignReportResp, len(list))
	for i, r := range list {
		result[i] = &dto.CampaignReportResp{
			CampaignID:  r.CampaignID,
			StatDate:    r.StatDate,
			Cost:        r.Cost,
			Show:        r.Show,
			Click:       r.Click,
			Convert:     r.Convert,
			CTR:         r.CTR,
			CVR:         r.CVR,
			CPM:         r.CPM,
			CPC:         r.CPC,
			ConvertCost: r.ConvertCost,
		}
	}

	return result, nil
}

// GetAdReport 获取广告组报告
func (s *ReportService) GetAdReport(ctx context.Context, req *dto.ReportQueryReq) ([]*dto.AdReportResp, error) {
	list, err := s.repo.GetAdReport(ctx, req)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.AdReportResp, len(list))
	for i, r := range list {
		result[i] = &dto.AdReportResp{
			AdID:        r.AdID,
			StatDate:    r.StatDate,
			Cost:        r.Cost,
			Show:        r.Show,
			Click:       r.Click,
			Convert:     r.Convert,
			CTR:         r.CTR,
			CVR:         r.CVR,
			CPM:         r.CPM,
			CPC:         r.CPC,
			ConvertCost: r.ConvertCost,
		}
	}

	return result, nil
}

// SyncReport 同步报告数据
func (s *ReportService) SyncReport(ctx context.Context, req *dto.ReportSyncReq) (*dto.ReportSyncResp, error) {
	// 获取广告主信息
	adv, err := s.advRepo.GetByID(ctx, req.AdvertiserID)
	if err != nil {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}

	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}

	// 创建SDK客户端
	client := oceanengine.NewClient(s.oceanCfg.AppID, s.oceanCfg.Secret)
	client.SetAccessToken(adv.AccessToken)

	reportService := oceanengine.NewReportService(client)
	syncCount := 0

	// 同步广告主维度报告
	reportReq := &oceanengine.ReportRequest{
		AdvertiserID: int64(adv.AdvertiserID),
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}
	reports, _, err := reportService.GetAdvertiserReport(ctx, reportReq)
	if err == nil {
		for _, r := range reports {
			report := &model.AdvertiserReport{
				AdvertiserID: req.AdvertiserID,
				StatDate:     r.StatDatetime,
				Cost:         r.Cost,
				Show:         r.ShowCnt,
				Click:        r.ClickCnt,
				Convert:      r.ConvertCnt,
				CTR:          r.CTR,
				CVR:          r.ConvertRate,
				CPM:          r.CPM,
				CPC:          r.CPC,
				ConvertCost:  r.ConvertCost,
			}
			if err := s.repo.SaveAdvertiserReport(ctx, report); err == nil {
				syncCount++
			}
		}
	}

	// 同步广告系列维度报告
	if req.Dimension == "" || req.Dimension == "CAMPAIGN" {
		campaignReports, _, err := reportService.GetCampaignReport(ctx, reportReq)
		if err == nil {
			for _, r := range campaignReports {
				report := &model.CampaignReport{
					AdvertiserID: req.AdvertiserID,
					CampaignID:   uint64(r.CampaignID),
					StatDate:     r.StatDatetime,
					Cost:         r.Cost,
					Show:         r.ShowCnt,
					Click:        r.ClickCnt,
					Convert:      r.ConvertCnt,
					CTR:          r.CTR,
					CVR:          r.ConvertRate,
					CPM:          r.CPM,
					CPC:          r.CPC,
					ConvertCost:  r.ConvertCost,
				}
				if err := s.repo.SaveCampaignReport(ctx, report); err == nil {
					syncCount++
				}
			}
		}
	}

	return &dto.ReportSyncResp{
		SyncCount: syncCount,
		SyncAt:    time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// GetExportTaskList 获取导出任务列表
func (s *ReportService) GetExportTaskList(ctx context.Context, req *dto.ExportTaskListReq) ([]*dto.ExportTaskResp, int64, error) {
	list, total, err := s.repo.GetExportTaskList(ctx, req)
	if err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.ExportTaskResp, len(list))
	for i, t := range list {
		result[i] = &dto.ExportTaskResp{
			ID:           t.ID,
			AdvertiserID: t.AdvertiserID,
			TaskType:     t.TaskType,
			Status:       t.Status,
			FileName:     t.FileName,
			FileSize:     t.FileSize,
			StartDate:    t.StartDate,
			EndDate:      t.EndDate,
			ErrorMsg:     t.ErrorMsg,
			CreatedAt:    t.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// CreateExportTask 创建导出任务
func (s *ReportService) CreateExportTask(ctx context.Context, req *dto.ExportCreateReq) (*dto.ExportTaskResp, error) {
	task := &model.ExportTask{
		AdvertiserID: req.AdvertiserID,
		TaskType:     req.TaskType,
		Status:       model.ExportStatusPending,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}

	if err := s.repo.CreateExportTask(ctx, task); err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// TODO: 启动异步导出任务

	return &dto.ExportTaskResp{
		ID:           task.ID,
		AdvertiserID: task.AdvertiserID,
		TaskType:     task.TaskType,
		Status:       task.Status,
		StartDate:    task.StartDate,
		EndDate:      task.EndDate,
		CreatedAt:    task.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetExportTask 获取导出任务详情
func (s *ReportService) GetExportTask(ctx context.Context, id uint64) (*dto.ExportTaskResp, error) {
	task, err := s.repo.GetExportTaskByID(ctx, id)
	if err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if task == nil {
		return nil, errcode.New(errcode.ErrNotFound)
	}

	return &dto.ExportTaskResp{
		ID:           task.ID,
		AdvertiserID: task.AdvertiserID,
		TaskType:     task.TaskType,
		Status:       task.Status,
		FileName:     task.FileName,
		FileSize:     task.FileSize,
		StartDate:    task.StartDate,
		EndDate:      task.EndDate,
		ErrorMsg:     task.ErrorMsg,
		CreatedAt:    task.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetCreativeReport 获取创意报告
func (s *ReportService) GetCreativeReport(ctx context.Context, req *dto.ReportQueryReq) ([]*dto.CreativeReportResp, error) {
	// 这里简化处理，实际应从创意报表中查询
	// TODO: 创建创意报表并实现查询逻辑
	return []*dto.CreativeReportResp{}, nil
}

// GetRealtimeReport 获取实时报告
func (s *ReportService) GetRealtimeReport(ctx context.Context, req *dto.RealtimeReportReq) ([]*dto.RealtimeReportResp, error) {
	// 获取广告主信息
	adv, err := s.advRepo.GetByID(ctx, req.AdvertiserID)
	if err != nil {
		return nil, errcode.New(errcode.ErrAdvertiserNotFound)
	}

	if adv.AccessToken == "" {
		return nil, errcode.New(errcode.ErrOETokenInvalid)
	}

	// 实时数据通常从今天的报表中获取
	today := time.Now().Format("2006-01-02")

	var result []*dto.RealtimeReportResp

	switch req.Level {
	case "advertiser":
		reportReq := &dto.ReportQueryReq{
			AdvertiserID: req.AdvertiserID,
			StartDate:    today,
			EndDate:      today,
		}
		reports, err := s.repo.GetAdvertiserReport(ctx, reportReq)
		if err == nil && len(reports) > 0 {
			r := reports[0]
			result = append(result, &dto.RealtimeReportResp{
				TodayCost:  r.Cost,
				TodayShow:  r.Show,
				TodayClick: r.Click,
				TodayConv:  r.Convert,
				TodayCTR:   r.CTR,
				TodayCVR:   r.CVR,
				UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			})
		}
	case "campaign", "ad":
		// 简化处理，返回空结果
		// 实际应根据 level 查询对应报表
	}

	if result == nil {
		result = []*dto.RealtimeReportResp{}
	}

	return result, nil
}
