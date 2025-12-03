package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	"oceanengine-backend/pkg/database"
	"oceanengine-backend/pkg/logger"
	"oceanengine-backend/pkg/oceanengine"
)

// TaskRunner 任务运行器
type TaskRunner struct {
	cfg    *config.Config
	log    *zap.Logger
	db     *gorm.DB
	client *oceanengine.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func main() {
	// 命令行参数
	configPath := flag.String("config", "config/settings.yml", "配置文件路径")
	flag.Parse()

	// 加载配置
	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	log, err := logger.Init(&cfg.Logger)
	if err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	log.Info("启动定时任务服务")

	// 初始化数据库
	db, err := database.Init(&cfg.Database, log)
	if err != nil {
		log.Fatal(fmt.Sprintf("初始化数据库失败: %v", err))
	}

	// 初始化 Ocean Engine 客户端
	client := oceanengine.NewClient(cfg.Ocean.AppID, cfg.Ocean.Secret)

	// 创建任务运行器
	ctx, cancel := context.WithCancel(context.Background())
	runner := &TaskRunner{
		cfg:    cfg,
		log:    log,
		db:     db,
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}

	// 启动定时任务
	go runner.startScheduler()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("正在关闭定时任务服务...")
	cancel()
	time.Sleep(2 * time.Second)
	log.Info("定时任务服务已关闭")
}

// startScheduler 启动调度器
func (r *TaskRunner) startScheduler() {
	// 每小时同步广告主余额
	go r.runPeriodically("广告主余额同步", 1*time.Hour, r.syncAdvertiserBalance)

	// 每天凌晨2点同步昨日报表
	go r.runDailyAt("日报表同步", 2, 0, r.syncDailyReport)

	// 每5分钟检查并刷新即将过期的 Token
	go r.runPeriodically("Token刷新检查", 5*time.Minute, r.refreshExpiredTokens)

	// 每天清理30天前的操作日志
	go r.runDailyAt("操作日志清理", 3, 0, r.cleanOperationLogs)
}

// runPeriodically 周期性运行任务
func (r *TaskRunner) runPeriodically(name string, interval time.Duration, task func() error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	r.log.Info(fmt.Sprintf("[%s] 任务已启动，间隔: %v", name, interval))

	// 立即执行一次
	if err := task(); err != nil {
		r.log.Error(fmt.Sprintf("[%s] 执行失败: %v", name, err))
	}

	for {
		select {
		case <-r.ctx.Done():
			r.log.Info(fmt.Sprintf("[%s] 任务已停止", name))
			return
		case <-ticker.C:
			if err := task(); err != nil {
				r.log.Error(fmt.Sprintf("[%s] 执行失败: %v", name, err))
			} else {
				r.log.Info(fmt.Sprintf("[%s] 执行成功", name))
			}
		}
	}
}

// runDailyAt 每天指定时间运行任务
func (r *TaskRunner) runDailyAt(name string, hour, minute int, task func() error) {
	r.log.Info(fmt.Sprintf("[%s] 任务已启动，每天 %02d:%02d 执行", name, hour, minute))

	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if next.Before(now) {
			next = next.Add(24 * time.Hour)
		}

		timer := time.NewTimer(next.Sub(now))
		select {
		case <-r.ctx.Done():
			timer.Stop()
			r.log.Info(fmt.Sprintf("[%s] 任务已停止", name))
			return
		case <-timer.C:
			if err := task(); err != nil {
				r.log.Error(fmt.Sprintf("[%s] 执行失败: %v", name, err))
			} else {
				r.log.Info(fmt.Sprintf("[%s] 执行成功", name))
			}
		}
	}
}

// syncAdvertiserBalance 同步广告主余额
func (r *TaskRunner) syncAdvertiserBalance() error {
	r.log.Info("开始同步广告主余额...")

	// 1. 查询所有有效Token的广告主
	var advertisers []struct {
		ID           uint64
		AdvertiserID uint64
		AccessToken  string
	}
	if err := r.db.Table("ad_advertiser").
		Select("id, advertiser_id, access_token").
		Where("access_token != '' AND deleted_at IS NULL").
		Find(&advertisers).Error; err != nil {
		return fmt.Errorf("查询广告主失败: %w", err)
	}

	if len(advertisers) == 0 {
		r.log.Info("没有需要同步的广告主")
		return nil
	}

	successCount := 0
	failCount := 0

	// 2. 逐个同步余额
	for _, adv := range advertisers {
		balance, err := r.client.Qianchuan().GetBalance(r.ctx, adv.AccessToken, adv.AdvertiserID)
		if err != nil {
			r.log.Warn(fmt.Sprintf("获取广告主 %d 余额失败: %v", adv.AdvertiserID, err))
			failCount++
			continue
		}

		// 3. 更新数据库
		if err := r.db.Table("ad_advertiser").
			Where("id = ?", adv.ID).
			Updates(map[string]interface{}{
				"balance":      float64(balance) / 100, // 分转元
				"last_sync_at": time.Now(),
			}).Error; err != nil {
			r.log.Warn(fmt.Sprintf("更新广告主 %d 余额失败: %v", adv.AdvertiserID, err))
			failCount++
			continue
		}

		successCount++
	}

	r.log.Info(fmt.Sprintf("广告主余额同步完成，成功: %d, 失败: %d", successCount, failCount))
	return nil
}

// syncDailyReport 同步日报表
func (r *TaskRunner) syncDailyReport() error {
	r.log.Info("开始同步日报表...")

	// 1. 获取昨天的日期
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// 2. 查询所有有效Token的广告主
	var advertisers []struct {
		ID           uint64
		AdvertiserID uint64
		AccessToken  string
	}
	if err := r.db.Table("ad_advertiser").
		Select("id, advertiser_id, access_token").
		Where("access_token != '' AND deleted_at IS NULL").
		Find(&advertisers).Error; err != nil {
		return fmt.Errorf("查询广告主失败: %w", err)
	}

	if len(advertisers) == 0 {
		r.log.Info("没有需要同步的广告主")
		return nil
	}

	successCount := 0
	failCount := 0

	// 3. 逐个同步日报表
	for _, adv := range advertisers {
		report, err := r.client.Qianchuan().GetAdvertiserReport(r.ctx, adv.AccessToken, adv.AdvertiserID, yesterday, yesterday)
		if err != nil {
			r.log.Warn(fmt.Sprintf("获取广告主 %d 日报表失败: %v", adv.AdvertiserID, err))
			failCount++
			continue
		}

		// 4. 保存报表数据
		for _, data := range report {
			reportRecord := map[string]interface{}{
				"advertiser_id": adv.AdvertiserID,
				"stat_date":     yesterday,
				"cost":          data.Cost,
				"show_cnt":      data.ShowCnt,
				"click_cnt":     data.ClickCnt,
				"convert_cnt":   data.ConvertCnt,
				"pay_order_cnt": data.PayOrderCnt,
				"pay_order_amt": data.PayOrderAmt,
				"created_at":    time.Now(),
			}

			// 使用upsert逻辑
			if err := r.db.Table("ad_report_daily").Where(
				"advertiser_id = ? AND stat_date = ?",
				adv.AdvertiserID, yesterday,
			).Assign(reportRecord).FirstOrCreate(&map[string]interface{}{}).Error; err != nil {
				r.log.Warn(fmt.Sprintf("保存广告主 %d 报表失败: %v", adv.AdvertiserID, err))
			}
		}

		successCount++
	}

	r.log.Info(fmt.Sprintf("日报表同步完成(%s)，成功: %d, 失败: %d", yesterday, successCount, failCount))
	return nil
}

// refreshExpiredTokens 刷新即将过期的 Token
func (r *TaskRunner) refreshExpiredTokens() error {
	r.log.Debug("检查即将过期的 Token...")

	// 1. 查询即将过期的Token (未杨1小时内过期)
	expireTime := time.Now().Add(1 * time.Hour)

	var advertisers []struct {
		ID            uint64
		AdvertiserID  uint64
		RefreshToken  string
		TokenExpireAt *time.Time
	}
	if err := r.db.Table("ad_advertiser").
		Select("id, advertiser_id, refresh_token, token_expire_at").
		Where("refresh_token != '' AND token_expire_at IS NOT NULL AND token_expire_at < ? AND deleted_at IS NULL", expireTime).
		Find(&advertisers).Error; err != nil {
		return fmt.Errorf("查询广告主失败: %w", err)
	}

	if len(advertisers) == 0 {
		return nil
	}

	r.log.Info(fmt.Sprintf("发现 %d 个即将过期的Token", len(advertisers)))

	successCount := 0
	failCount := 0

	// 2. 逐个刷新Token
	for _, adv := range advertisers {
		tokenData, err := r.client.OAuth().RefreshAccessToken(r.ctx, adv.RefreshToken)
		if err != nil {
			r.log.Warn(fmt.Sprintf("刷新广告主 %d Token失败: %v", adv.AdvertiserID, err))
			failCount++
			continue
		}

		// 3. 更新数据库
		newExpireAt := time.Now().Add(time.Duration(tokenData.ExpiresIn) * time.Second)
		if err := r.db.Table("ad_advertiser").
			Where("id = ?", adv.ID).
			Updates(map[string]interface{}{
				"access_token":    tokenData.AccessToken,
				"refresh_token":   tokenData.RefreshToken,
				"token_expire_at": newExpireAt,
				"updated_at":      time.Now(),
			}).Error; err != nil {
			r.log.Warn(fmt.Sprintf("更新广告主 %d Token失败: %v", adv.AdvertiserID, err))
			failCount++
			continue
		}

		successCount++
		r.log.Info(fmt.Sprintf("刷新广告主 %d Token成功", adv.AdvertiserID))
	}

	if successCount > 0 || failCount > 0 {
		r.log.Info(fmt.Sprintf("Token刷新完成，成功: %d, 失败: %d", successCount, failCount))
	}
	return nil
}

// cleanOperationLogs 清理操作日志
func (r *TaskRunner) cleanOperationLogs() error {
	r.log.Info("开始清理操作日志...")

	// 删除30天前的操作日志
	cutoffDate := time.Now().AddDate(0, 0, -30)
	result := r.db.Table("sys_operation_log").Where("created_at < ?", cutoffDate).Delete(&struct{}{})
	if result.Error != nil {
		return result.Error
	}

	r.log.Info(fmt.Sprintf("清理操作日志完成，删除 %d 条记录", result.RowsAffected))
	return nil
}
