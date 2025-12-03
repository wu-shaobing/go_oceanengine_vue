package main

import (
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	adModel "oceanengine-backend/internal/app/ad/model"
	adminModel "oceanengine-backend/internal/app/admin/model"
	advertiserModel "oceanengine-backend/internal/app/advertiser/model"
	audienceModel "oceanengine-backend/internal/app/audience/model"
	campaignModel "oceanengine-backend/internal/app/campaign/model"
	creativeModel "oceanengine-backend/internal/app/creative/model"
	mediaModel "oceanengine-backend/internal/app/media/model"
	reportModel "oceanengine-backend/internal/app/report/model"
	"oceanengine-backend/pkg/database"
	"oceanengine-backend/pkg/logger"
)

func main() {
	// 命令行参数
	configPath := flag.String("config", "config/settings.yml", "配置文件路径")
	action := flag.String("action", "migrate", "执行操作: migrate, fresh, seed")
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

	// 初始化数据库
	db, err := database.Init(&cfg.Database, log)
	if err != nil {
		log.Fatal(fmt.Sprintf("初始化数据库失败: %v", err))
	}

	switch *action {
	case "migrate":
		runMigrate(log, db)
	case "fresh":
		runFresh(log, db)
	case "seed":
		runSeed(log, db)
	default:
		fmt.Printf("未知操作: %s\n", *action)
		os.Exit(1)
	}
}

// runMigrate 执行迁移
func runMigrate(log *zap.Logger, db *gorm.DB) {
	log.Info("开始执行数据库迁移...")

	// 迁移所有模型
	models := []interface{}{
		// 系统管理模块
		&adminModel.User{},
		&adminModel.Role{},
		&adminModel.Menu{},
		&adminModel.RoleMenu{},
		&adminModel.OperationLog{},
		// 广告主模块
		&advertiserModel.Advertiser{},
		&advertiserModel.AdvertiserFund{},
		// 广告系列模块
		&campaignModel.Campaign{},
		// 广告组模块
		&adModel.Ad{},
		// 创意模块
		&creativeModel.Creative{},
		// 报表模块
		&reportModel.AdvertiserReport{},
		&reportModel.CampaignReport{},
		&reportModel.AdReport{},
		&reportModel.ExportTask{},
		// 素材模块
		&mediaModel.MaterialImage{},
		&mediaModel.MaterialVideo{},
		// 人群定向模块
		&audienceModel.AudiencePackage{},
		&audienceModel.CustomAudience{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Error(fmt.Sprintf("迁移失败: %v", err))
			os.Exit(1)
		}
	}

	log.Info("数据库迁移完成")
}

// runFresh 清空并重建表
func runFresh(log *zap.Logger, db *gorm.DB) {
	log.Info("开始清空数据库...")

	// 获取所有表
	tables := []string{
		"sys_user", "sys_role", "sys_menu", "sys_role_menu", "sys_operation_log",
		"ad_advertiser", "ad_advertiser_fund",
		"ad_campaign", "ad_ad", "ad_creative",
		"rpt_advertiser_daily", "rpt_campaign_daily", "rpt_ad_daily",
		"ad_material_image", "ad_material_video",
		"ad_audience_package", "ad_custom_audience",
	}

	// 禁用外键检查
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Warn(fmt.Sprintf("删除表 %s 失败: %v", table, err))
		} else {
			log.Info(fmt.Sprintf("删除表 %s 成功", table))
		}
	}

	// 恢复外键检查
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// 重新迁移
	runMigrate(log, db)
}

// runSeed 填充初始数据
func runSeed(log *zap.Logger, db *gorm.DB) {
	log.Info("开始填充初始数据...")

	// 创建默认管理员角色
	adminRole := &adminModel.Role{
		Name:      "超级管理员",
		Key:       "admin",
		Sort:      1,
		Status:    1,
		DataScope: 1,
		Remark:    "系统超级管理员",
	}
	if err := db.FirstOrCreate(adminRole, adminModel.Role{Key: "admin"}).Error; err != nil {
		log.Error(fmt.Sprintf("创建管理员角色失败: %v", err))
	} else {
		log.Info("管理员角色创建成功")
	}

	// 创建默认管理员用户（密码: admin123）
	// 密码使用 bcrypt 加密: $2a$10$...
	adminUser := &adminModel.User{
		Username: "admin",
		Password: "$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6vKq6", // admin123
		Nickname: "管理员",
		Status:   1,
		RoleID:   adminRole.ID,
		Remark:   "系统管理员账号",
	}
	if err := db.FirstOrCreate(adminUser, adminModel.User{Username: "admin"}).Error; err != nil {
		log.Error(fmt.Sprintf("创建管理员用户失败: %v", err))
	} else {
		log.Info("管理员用户创建成功")
	}

	// 创建基础菜单
	menus := []adminModel.Menu{
		{ParentID: 0, Name: "系统管理", Path: "/system", Component: "Layout", Icon: "setting", Sort: 1, Type: 1, Visible: 1, Status: 1},
		{ParentID: 0, Name: "广告管理", Path: "/ads", Component: "Layout", Icon: "promotion", Sort: 2, Type: 1, Visible: 1, Status: 1},
		{ParentID: 0, Name: "数据报表", Path: "/report", Component: "Layout", Icon: "data-analysis", Sort: 3, Type: 1, Visible: 1, Status: 1},
	}

	for _, menu := range menus {
		if err := db.FirstOrCreate(&menu, adminModel.Menu{Name: menu.Name, ParentID: menu.ParentID}).Error; err != nil {
			log.Error(fmt.Sprintf("创建菜单 %s 失败: %v", menu.Name, err))
		}
	}
	log.Info("基础菜单创建成功")

	log.Info("初始数据填充完成")
}
