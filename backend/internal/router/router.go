package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	adApi "oceanengine-backend/internal/app/ad/api"
	adminApi "oceanengine-backend/internal/app/admin/api"
	"oceanengine-backend/internal/app/admin/service"
	advApi "oceanengine-backend/internal/app/advertiser/api"
	advtoolsApi "oceanengine-backend/internal/app/advtools/api"
	audienceApi "oceanengine-backend/internal/app/audience/api"
	audienceService "oceanengine-backend/internal/app/audience/service"
	campaignApi "oceanengine-backend/internal/app/campaign/api"
	clueApi "oceanengine-backend/internal/app/clue/api"
	creativeApi "oceanengine-backend/internal/app/creative/api"
	dmpApi "oceanengine-backend/internal/app/dmp/api"
	dpaApi "oceanengine-backend/internal/app/dpa/api"
	enterpriseApi "oceanengine-backend/internal/app/enterprise/api"
	eventmanagerApi "oceanengine-backend/internal/app/eventmanager/api"
	localApi "oceanengine-backend/internal/app/local/api"
	mediaApi "oceanengine-backend/internal/app/media/api"
	mediaService "oceanengine-backend/internal/app/media/service"
	qianchuanApi "oceanengine-backend/internal/app/qianchuan/api"
	reportApi "oceanengine-backend/internal/app/report/api"
	serveMarketApi "oceanengine-backend/internal/app/servemarket/api"
	siteApi "oceanengine-backend/internal/app/site/api"
	starApi "oceanengine-backend/internal/app/star/api"
	v3Api "oceanengine-backend/internal/app/v3/api"
	"oceanengine-backend/internal/middleware"
	"oceanengine-backend/pkg/auth"
)

// Router 路由管理器
type Router struct {
	engine       *gin.Engine
	db           *gorm.DB
	logger       *zap.Logger
	jwtManager   *auth.JWTManager
	oceanCfg     *config.OceanConfig
	qianchuanCfg *config.QianchuanConfig
}

// NewRouter 创建路由
func NewRouter(db *gorm.DB, logger *zap.Logger, jwtManager *auth.JWTManager, oceanCfg *config.OceanConfig) *Router {
	return &Router{
		db:         db,
		logger:     logger,
		jwtManager: jwtManager,
		oceanCfg:   oceanCfg,
	}
}

// NewRouterWithQianchuan 创建带千川配置的路由
func NewRouterWithQianchuan(db *gorm.DB, logger *zap.Logger, jwtManager *auth.JWTManager, oceanCfg *config.OceanConfig, qianchuanCfg *config.QianchuanConfig) *Router {
	return &Router{
		db:           db,
		logger:       logger,
		jwtManager:   jwtManager,
		oceanCfg:     oceanCfg,
		qianchuanCfg: qianchuanCfg,
	}
}

// Setup 设置路由
func (r *Router) Setup(mode string) *gin.Engine {
	// 设置 Gin 模式
	gin.SetMode(mode)

	r.engine = gin.New()

	// 全局中间件
	r.engine.Use(middleware.Recovery(r.logger))
	r.engine.Use(middleware.RequestID())
	r.engine.Use(middleware.Logger(r.logger))
	// 安全响应头
	r.engine.Use(middleware.SecurityHeaders())
	// CORS（支持通过 CORS_ALLOWED_ORIGINS 配置）
	r.engine.Use(middleware.CORS(middleware.DefaultCORSConfig()))

	// 限流中间件
	rateLimiter := middleware.NewIPRateLimiter(middleware.DefaultRateLimitConfig())
	r.engine.Use(middleware.RateLimit(rateLimiter))

	// 健康检查
	r.engine.GET("/health", r.healthCheck)

	// API 路由组
	apiV1 := r.engine.Group("/api/v1")
	{
		// 公开路由
		r.registerPublicRoutes(apiV1)

		// 需要认证的路由
		protected := apiV1.Group("")
		protected.Use(middleware.JWTAuth(r.jwtManager))
		r.registerProtectedRoutes(protected)
	}

	return r.engine
}

// healthCheck 健康检查
func (r *Router) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// registerPublicRoutes 注册公开路由
func (r *Router) registerPublicRoutes(rg *gin.RouterGroup) {
	// 认证服务
	authService := service.NewAuthService(r.db, r.jwtManager)
	authAPI := adminApi.NewAuthAPI(authService)

	// 验证码服务
	captchaService := service.NewCaptchaService()
	captchaAPI := adminApi.NewCaptchaAPI(captchaService)

	// 认证相关
	authGroup := rg.Group("/auth")
	{
		authGroup.GET("/captcha", captchaAPI.Get)
		authGroup.POST("/login", authAPI.Login)
		authGroup.POST("/refresh", authAPI.RefreshToken)
	}

	// 广告主 OAuth 回调（公开）
	advHandler := advApi.NewAdvertiserHandler(r.db, r.oceanCfg)
	oauthGroup := rg.Group("/advertisers/oauth")
	{
		oauthGroup.GET("/callback", advHandler.OAuthCallback)
	}

	// 千川 OAuth 路由（公开）
	if r.qianchuanCfg != nil {
		qcOAuthHandler := qianchuanApi.NewQianchuanOAuthHandlerDefault(r.qianchuanCfg)
		qcOAuthGroup := rg.Group("/qianchuan/oauth")
		{
			qcOAuthGroup.GET("/url", qcOAuthHandler.GetAuthURL)
			qcOAuthGroup.GET("/callback", qcOAuthHandler.OAuthCallback)
			qcOAuthGroup.POST("/refresh", qcOAuthHandler.RefreshToken)
		}
	}
}

// registerProtectedRoutes 注册需要认证的路由
func (r *Router) registerProtectedRoutes(rg *gin.RouterGroup) {
	// 认证服务
	authService := service.NewAuthService(r.db, r.jwtManager)
	authAPI := adminApi.NewAuthAPI(authService)

	// 认证相关
	authGroup := rg.Group("/auth")
	{
		authGroup.GET("/userinfo", authAPI.GetUserInfo)
		authGroup.POST("/logout", authAPI.Logout)
	}

	// 系统管理
	r.registerSystemRoutes(rg)

	// 广告主管理
	r.registerAdvertiserRoutes(rg)

	// 广告系列管理
	r.registerCampaignRoutes(rg)

	// 广告组管理
	r.registerAdRoutes(rg)

	// 创意管理
	r.registerCreativeRoutes(rg)

	// 数据报告
	r.registerReportRoutes(rg)

	// 素材管理
	r.registerMediaRoutes(rg)

	// 人群定向
	r.registerAudienceRoutes(rg)

	// 千川模块
	r.registerQianchuanRoutes(rg)

	// 企业号模块
	r.registerEnterpriseRoutes(rg)

	// 本地推模块
	r.registerLocalRoutes(rg)

	// 星图模块
	r.registerStarRoutes(rg)

	// 服务市场模块
	r.registerServeMarketRoutes(rg)

	// 线索管理模块
	r.registerClueRoutes(rg)

	// 事件管理模块
	r.registerEventManagerRoutes(rg)

	// 高级工具模块
	r.registerAdvToolsRoutes(rg)

	// 建站管理模块
	r.registerSiteRoutes(rg)

	// V3体验版模块
	r.registerV3Routes(rg)

	// DMP人群包模块
	r.registerDMPRoutes(rg)

	// DPA商品广告模块
	r.registerDPARoutes(rg)
}

// registerSystemRoutes 注册系统管理路由
func (r *Router) registerSystemRoutes(rg *gin.RouterGroup) {
	// 初始化服务
	userService := service.NewUserService(r.db)
	roleService := service.NewRoleService(r.db)
	menuService := service.NewMenuService(r.db)
	logService := service.NewOperationLogService(r.db)

	// 初始化 API
	userAPI := adminApi.NewUserAPI(userService)
	roleAPI := adminApi.NewRoleAPI(roleService)
	menuAPI := adminApi.NewMenuAPI(menuService)
	logAPI := adminApi.NewOperationLogAPI(logService)

	system := rg.Group("/system")
	{
		// 用户管理
		users := system.Group("/users")
		{
			users.GET("", userAPI.GetList)
			users.POST("", userAPI.Create)
			users.GET("/:id", userAPI.GetByID)
			users.PUT("/:id", userAPI.Update)
			users.DELETE("/:id", userAPI.Delete)
			users.POST("/:id/reset-password", userAPI.ResetPassword)
			users.POST("/change-password", userAPI.ChangePassword)
		}

		// 角色管理
		roles := system.Group("/roles")
		{
			roles.GET("", roleAPI.GetList)
			roles.GET("/all", roleAPI.GetAll)
			roles.POST("", roleAPI.Create)
			roles.GET("/:id", roleAPI.GetByID)
			roles.PUT("/:id", roleAPI.Update)
			roles.DELETE("/:id", roleAPI.Delete)
			roles.GET("/:id/menus", roleAPI.GetRoleMenus)
			roles.PUT("/:id/menus", roleAPI.UpdateRoleMenus)
		}

		// 菜单管理
		menus := system.Group("/menus")
		{
			menus.GET("", menuAPI.GetList)
			menus.GET("/tree", menuAPI.GetTree)
			menus.GET("/user", menuAPI.GetUserMenuTree)
			menus.POST("", menuAPI.Create)
			menus.GET("/:id", menuAPI.GetByID)
			menus.PUT("/:id", menuAPI.Update)
			menus.DELETE("/:id", menuAPI.Delete)
		}

		// 操作日志
		logs := system.Group("/logs")
		{
			logs.GET("/operation", logAPI.GetList)
			logs.GET("/modules", logAPI.GetModules)
			logs.DELETE("/operation", logAPI.Delete)
		}
	}
}

// registerAdvertiserRoutes 注册广告主路由
func (r *Router) registerAdvertiserRoutes(rg *gin.RouterGroup) {
	advHandler := advApi.NewAdvertiserHandler(r.db, r.oceanCfg)

	advertisers := rg.Group("/advertisers")
	{
		advertisers.GET("", advHandler.List)
		advertisers.GET("/:id", advHandler.Get)
		advertisers.PUT("/:id", advHandler.Update)
		advertisers.DELETE("/:id", advHandler.Delete)
		advertisers.POST("/:id/sync", advHandler.Sync)
		advertisers.GET("/:id/balance", advHandler.GetBalance)
		advertisers.GET("/:id/funds", advHandler.GetFundList)

		// OAuth 相关
		oauth := advertisers.Group("/oauth")
		{
			oauth.GET("/url", advHandler.GetOAuthURL)
		}
	}
}

// registerCampaignRoutes 注册广告系列路由
func (r *Router) registerCampaignRoutes(rg *gin.RouterGroup) {
	campaignHandler := campaignApi.NewCampaignHandler(r.db, r.oceanCfg)

	campaigns := rg.Group("/campaigns")
	{
		campaigns.GET("", campaignHandler.List)
		campaigns.POST("", campaignHandler.Create)
		campaigns.GET("/:id", campaignHandler.Get)
		campaigns.PUT("/:id", campaignHandler.Update)
		campaigns.DELETE("/:id", campaignHandler.Delete)
		campaigns.PUT("/status", campaignHandler.UpdateStatus)
		campaigns.POST("/sync/:advertiser_id", campaignHandler.Sync)
	}
}

// registerAdRoutes 注册广告组路由
func (r *Router) registerAdRoutes(rg *gin.RouterGroup) {
	adHandler := adApi.NewAdHandler(r.db, r.oceanCfg)

	ads := rg.Group("/ads")
	{
		ads.GET("", adHandler.List)
		ads.POST("", adHandler.Create)
		ads.GET("/:id", adHandler.Get)
		ads.PUT("/:id", adHandler.Update)
		ads.DELETE("/:id", adHandler.Delete)
		ads.PUT("/status", adHandler.UpdateStatus)
	}
}

// registerCreativeRoutes 注册创意路由
func (r *Router) registerCreativeRoutes(rg *gin.RouterGroup) {
	creativeHandler := creativeApi.NewCreativeHandler(r.db)

	creatives := rg.Group("/creatives")
	{
		creatives.GET("", creativeHandler.List)
		creatives.POST("", creativeHandler.Create)
		creatives.GET("/:id", creativeHandler.Get)
		creatives.PUT("/:id", creativeHandler.Update)
		creatives.DELETE("/:id", creativeHandler.Delete)
		creatives.PUT("/status", creativeHandler.UpdateStatus)
	}
}

// registerReportRoutes 注册报表路由
func (r *Router) registerReportRoutes(rg *gin.RouterGroup) {
	reportHandler := reportApi.NewReportHandler(r.db, r.oceanCfg)

	reports := rg.Group("/reports")
	{
		reports.GET("/advertiser", reportHandler.GetAdvertiserReport)
		reports.GET("/advertiser/summary", reportHandler.GetAdvertiserSummary)
		reports.GET("/campaign", reportHandler.GetCampaignReport)
		reports.GET("/ad", reportHandler.GetAdReport)
		reports.POST("/sync", reportHandler.SyncReport)
		reports.GET("/exports", reportHandler.GetExportTaskList)
		reports.POST("/exports", reportHandler.CreateExportTask)
	}
}

// registerMediaRoutes 注册素材管理路由
func (r *Router) registerMediaRoutes(rg *gin.RouterGroup) {
	mediaSvc := mediaService.NewMediaService(r.db)
	mediaHandler := mediaApi.NewMediaAPI(mediaSvc)

	media := rg.Group("/media")
	{
		// 图片素材
		images := media.Group("/images")
		{
			images.GET("", mediaHandler.GetImageList)
			images.GET("/:id", mediaHandler.GetImageByID)
			images.DELETE("/:id", mediaHandler.DeleteImage)
		}

		// 视频素材
		videos := media.Group("/videos")
		{
			videos.GET("", mediaHandler.GetVideoList)
			videos.GET("/:id", mediaHandler.GetVideoByID)
			videos.DELETE("/:id", mediaHandler.DeleteVideo)
		}
	}
}

// registerAudienceRoutes 注册人群定向路由
func (r *Router) registerAudienceRoutes(rg *gin.RouterGroup) {
	audService := audienceService.NewAudienceService(r.db)
	audienceHandler := audienceApi.NewAudienceAPI(audService)

	audiences := rg.Group("/audiences")
	{
		// 定向包
		packages := audiences.Group("/packages")
		{
			packages.GET("", audienceHandler.GetPackageList)
			packages.POST("", audienceHandler.CreatePackage)
			packages.GET("/:id", audienceHandler.GetPackageByID)
			packages.PUT("/:id", audienceHandler.UpdatePackage)
			packages.DELETE("/:id", audienceHandler.DeletePackage)
		}

		// 人群包
		custom := audiences.Group("/custom")
		{
			custom.GET("", audienceHandler.GetCustomAudienceList)
			custom.GET("/:id", audienceHandler.GetCustomAudienceByID)
			custom.DELETE("/:id", audienceHandler.DeleteCustomAudience)
		}
	}
}

// registerQianchuanRoutes 注册千川路由
func (r *Router) registerQianchuanRoutes(rg *gin.RouterGroup) {
	handler := qianchuanApi.NewQianchuanHandler(r.db, r.oceanCfg)

	qianchuan := rg.Group("/qianchuan")
	{
		qianchuan.GET("/account", handler.GetAccountInfo)
		qianchuan.GET("/shops", handler.GetShopList)
		qianchuan.GET("/aweme/auth", handler.GetAwemeAuthList)
		qianchuan.GET("/balance", handler.GetBalance)
		qianchuan.GET("/campaigns", handler.GetCampaignList)
		qianchuan.POST("/campaigns", handler.CreateCampaign)
		qianchuan.GET("/ads", handler.GetAdList)
		qianchuan.GET("/ads/:ad_id", handler.GetAdDetail)
		qianchuan.POST("/ads", handler.CreateAd)
		qianchuan.POST("/ads/status", handler.UpdateAdStatus)
		qianchuan.GET("/creatives", handler.GetCreativeList)
		qianchuan.GET("/aweme/orders", handler.GetAwemeOrderList)
		qianchuan.GET("/aweme/orders/:order_id", handler.GetAwemeOrderDetail)
		qianchuan.POST("/aweme/orders", handler.CreateAwemeOrder)
		qianchuan.GET("/reports/advertiser", handler.GetAdvertiserReport)
		qianchuan.GET("/reports/ad", handler.GetAdReport)
		qianchuan.GET("/reports/creative", handler.GetCreativeReport)
		qianchuan.GET("/reports/material", handler.GetMaterialReport)
		qianchuan.GET("/reports/keyword", handler.GetKeywordReport)
		qianchuan.GET("/reports/live", handler.GetLiveReport)
		qianchuan.GET("/reports/room", handler.GetRoomReport)
		qianchuan.GET("/reports/uni", handler.GetUniReport)
		// 全域推广
		qianchuan.GET("/uni/list", handler.GetUniList)
		qianchuan.POST("/uni", handler.CreateUni)
		qianchuan.GET("/uni/:ad_id", handler.GetUniDetail)
		// 素材
		qianchuan.GET("/materials", handler.GetMaterialList)
		qianchuan.POST("/materials/image", handler.UploadImage)
		qianchuan.POST("/materials/video", handler.UploadVideo)
		// 工具
		qianchuan.GET("/tools/industries", handler.GetIndustryList)
		qianchuan.GET("/tools/dmp", handler.GetDmpList)
		qianchuan.GET("/tools/keyword/recommend", handler.GetKeywordRecommend)
		qianchuan.GET("/products", handler.GetProductList)
		qianchuan.GET("/budget", handler.GetBudget)
		qianchuan.POST("/budget", handler.UpdateBudget)
		qianchuan.GET("/finance/detail", handler.GetFinanceDetail)
		// 关键词管理
		qianchuan.GET("/keywords", handler.GetKeywordList)
		qianchuan.PUT("/keywords", handler.UpdateKeywords)
		qianchuan.GET("/keywords/action", handler.GetActionKeywords)
		qianchuan.GET("/keywords/interest", handler.GetInterestKeywords)
		qianchuan.POST("/keywords/suggest", handler.GetKeywordSuggest)
	}
}

// registerEnterpriseRoutes 注册企业号路由
func (r *Router) registerEnterpriseRoutes(rg *gin.RouterGroup) {
	handler := enterpriseApi.NewEnterpriseHandler(r.db, r.oceanCfg)

	enterprise := rg.Group("/enterprise")
	{
		enterprise.GET("/info", handler.GetInfo)
		enterprise.GET("/binds", handler.GetBindList)
		enterprise.POST("/binds", handler.BindAccount)
		enterprise.DELETE("/binds/:bind_id", handler.UnbindAccount)
		enterprise.GET("/items", handler.GetItemList)
		enterprise.GET("/items/:item_id", handler.GetItemDetail)
		enterprise.POST("/items/:item_id/top", handler.SetTopItem)
		enterprise.DELETE("/items/:item_id", handler.DeleteItem)
		enterprise.GET("/comments", handler.GetCommentList)
		enterprise.POST("/comments/reply", handler.ReplyComment)
		enterprise.POST("/comments/batch-reply", handler.BatchReplyComments)
		enterprise.PUT("/comments/:comment_id/reply", handler.UpdateReply)
		enterprise.POST("/comments/:comment_id/hide", handler.HideComment)
		enterprise.DELETE("/comments/:comment_id", handler.DeleteComment)
		enterprise.GET("/reply-templates", handler.GetReplyTemplates)
		enterprise.POST("/reply-templates", handler.CreateReplyTemplate)
		enterprise.DELETE("/reply-templates/:template_id", handler.DeleteReplyTemplate)
		enterprise.GET("/overview", handler.GetOverviewData)
		enterprise.GET("/dashboard", handler.GetDashboardStats)
		enterprise.GET("/traffic-source", handler.GetTrafficSource)
		enterprise.GET("/logs", handler.GetOperationLogs)
	}
}

// registerLocalRoutes 注册本地推路由
func (r *Router) registerLocalRoutes(rg *gin.RouterGroup) {
	handler := localApi.NewLocalHandler(r.db, r.oceanCfg)

	local := rg.Group("/local")
	{
		local.GET("/projects", handler.GetProjectList)
		local.GET("/projects/:project_id", handler.GetProjectDetail)
		local.POST("/projects", handler.CreateProject)
		local.PUT("/projects/:project_id", handler.UpdateProject)
		local.POST("/projects/status", handler.UpdateProjectStatus)
		local.DELETE("/projects/:project_id", handler.DeleteProject)
		local.GET("/promotions", handler.GetPromotionList)
		local.GET("/promotions/:promotion_id", handler.GetPromotionDetail)
		local.POST("/promotions", handler.CreatePromotion)
		local.PUT("/promotions/:promotion_id", handler.UpdatePromotion)
		local.POST("/promotions/status", handler.UpdatePromotionStatus)
		local.DELETE("/promotions/:promotion_id", handler.DeletePromotion)
		local.GET("/clues", handler.GetClueList)
		local.GET("/clues/:clue_id", handler.GetClueDetail)
		local.PUT("/clues/:clue_id", handler.UpdateClueStatus)
		local.POST("/clues/export", handler.ExportClues)
		local.GET("/reports/project", handler.GetProjectReport)
		local.GET("/reports/promotion", handler.GetPromotionReport)
		local.GET("/reports/material", handler.GetMaterialReport)
		local.GET("/materials", handler.GetMaterialList)
		local.POST("/materials/video", handler.UploadVideo)
		local.POST("/materials/image", handler.UploadImage)
		local.DELETE("/materials/:material_id", handler.DeleteMaterial)
		local.GET("/stores", handler.GetStoreList)
		local.GET("/stores/:store_id", handler.GetStoreDetail)
		local.POST("/stores", handler.CreateStore)
		local.PUT("/stores/:store_id", handler.UpdateStore)
		local.DELETE("/stores/:store_id", handler.DeleteStore)
	}
}

// registerStarRoutes 注册星图路由
func (r *Router) registerStarRoutes(rg *gin.RouterGroup) {
	handler := starApi.NewStarHandler(r.db, r.oceanCfg)

	star := rg.Group("/star")
	{
		star.GET("/account", handler.GetAccountInfo)
		star.GET("/agent/advertisers", handler.GetAgentAdvertisers)
		star.POST("/fund/balance", handler.GetBatchBalance)
		star.GET("/fund/daily", handler.GetFundDaily)
		star.GET("/fund/transactions", handler.GetFundTransactions)
		star.GET("/tasks", handler.GetTaskList)
		star.GET("/tasks/:task_id", handler.GetTaskDetail)
		star.GET("/tasks/:task_id/items", handler.GetTaskItems)
		star.PUT("/tasks/:task_id/status", handler.UpdateTaskStatus)
		star.GET("/demands", handler.GetDemandList)
		star.GET("/demands/:demand_id", handler.GetDemandDetail)
		star.GET("/demands/:demand_id/orders", handler.GetDemandOrders)
		star.GET("/reports/overview", handler.GetReportOverview)
		star.GET("/reports/audience", handler.GetReportAudience)
		star.GET("/reports/daily", handler.GetReportDaily)
		star.GET("/clues", handler.GetClueList)
		star.PUT("/clues/:clue_id", handler.UpdateClueStatus)
		star.POST("/clues/export", handler.ExportClues)
	}
}

// registerServeMarketRoutes 注册服务市场路由
func (r *Router) registerServeMarketRoutes(rg *gin.RouterGroup) {
	handler := serveMarketApi.NewServeMarketHandler(r.db, r.oceanCfg)

	servemarket := rg.Group("/servemarket")
	{
		servemarket.GET("/orders", handler.GetOrderList)
		servemarket.GET("/orders/:order_id", handler.GetOrderDetail)
		servemarket.GET("/funcs", handler.GetFuncList)
		servemarket.GET("/funcs/:func_id", handler.GetFuncDetail)
		servemarket.GET("/quality", handler.GetQualityAnalysis)
		servemarket.GET("/subscriptions", handler.GetSubscriptionList)
		servemarket.POST("/subscriptions", handler.CreateSubscription)
		servemarket.PUT("/subscriptions/:subscription_id", handler.UpdateSubscription)
		servemarket.DELETE("/subscriptions/:subscription_id", handler.DeleteSubscription)
		servemarket.GET("/dashboard", handler.GetDashboard)
	}
}

// registerClueRoutes 注册线索管理路由
func (r *Router) registerClueRoutes(rg *gin.RouterGroup) {
	handler := clueApi.NewClueHandler(r.db, r.oceanCfg)

	clue := rg.Group("/clue")
	{
		// 飞鱼线索
		clue.GET("/list", handler.GetClueList)
		clue.POST("/callback", handler.ClueCallback)
		clue.POST("/callback/batch", handler.BatchClueCallback)
		clue.GET("/key-action", handler.GetKeyAction)
		clue.GET("/smart-phone", handler.GetSmartPhone)
		clue.GET("/form/list", handler.GetFormList)
		clue.GET("/form/detail", handler.GetFormDetail)
		clue.GET("/store/list", handler.GetClueStoreList)

		// 青鸟线索通 - 表单管理
		qingniao := clue.Group("/qingniao")
		{
			qingniao.GET("/forms", handler.GetQingniaoFormList)
			qingniao.POST("/forms", handler.CreateQingniaoForm)
			qingniao.PUT("/forms/:form_id", handler.UpdateQingniaoForm)
			qingniao.DELETE("/forms/:form_id", handler.DeleteQingniaoForm)

			// 优惠券管理
			qingniao.GET("/coupons", handler.GetCouponList)
			qingniao.GET("/coupons/:coupon_id", handler.GetCouponDetail)
			qingniao.POST("/coupons", handler.CreateCoupon)
			qingniao.PUT("/coupons/:coupon_id", handler.UpdateCoupon)

			// 智能电话
			qingniao.GET("/smart-phone", handler.GetQingniaoSmartPhoneList)
			qingniao.POST("/smart-phone", handler.CreateSmartPhone)
			qingniao.DELETE("/smart-phone/:smart_phone_id", handler.DeleteSmartPhone)
			qingniao.GET("/smart-phone/records", handler.GetSmartPhoneRecords)

			// 微信相关
			qingniao.GET("/wechat/pool", handler.GetWechatPoolList)
			qingniao.GET("/wechat/instances", handler.GetWechatInstanceList)
			qingniao.GET("/wechat/instances/:instance_id", handler.GetWechatInstanceDetail)
			qingniao.PUT("/wechat/instances/:instance_id", handler.UpdateWechatInstance)
		}
	}
}

// registerEventManagerRoutes 注册事件管理路由
func (r *Router) registerEventManagerRoutes(rg *gin.RouterGroup) {
	handler := eventmanagerApi.NewEventManagerHandler(r.db, r.oceanCfg)

	eventmanager := rg.Group("/eventmanager")
	{
		// 资产管理
		eventmanager.GET("/assets", handler.GetAssets)
		eventmanager.GET("/assets/all", handler.GetAllAssetsList)
		eventmanager.POST("/assets", handler.CreateAsset)

		// 事件管理
		eventmanager.GET("/events/available", handler.GetAvailableEvents)
		eventmanager.GET("/events/configs", handler.GetEventConfigs)
		eventmanager.POST("/events", handler.CreateEvents)

		// 监测链接
		eventmanager.GET("/track-url", handler.GetTrackURL)
		eventmanager.POST("/track-url", handler.CreateTrackURL)
		eventmanager.PUT("/track-url", handler.UpdateTrackURL)

		// 资产共享
		eventmanager.GET("/share", handler.GetShare)
		eventmanager.POST("/share", handler.Share)
		eventmanager.POST("/share/cancel", handler.ShareCancel)

		// 优化目标
		eventmanager.GET("/optimized-goal", handler.GetOptimizedGoal)

		// 转化回传
		eventmanager.POST("/conversion", handler.Conversion)

		// 转化回传鉴权
		eventmanager.POST("/auth/public-key", handler.AddPublicKey)
		eventmanager.GET("/auth/public-keys", handler.GetAllPublicKeys)
		eventmanager.POST("/auth/enable", handler.EnableAuth)
		eventmanager.POST("/auth/disable", handler.DisableAuth)
	}
}

// registerAdvToolsRoutes 注册高级工具路由
func (r *Router) registerAdvToolsRoutes(rg *gin.RouterGroup) {
	handler := advtoolsApi.NewAdvToolsHandler(r.db, r.oceanCfg)

	advtools := rg.Group("/advtools")
	{
		// RTA策略管理
		rta := advtools.Group("/rta")
		{
			rta.GET("/info", handler.GetRtaInfo)
			rta.GET("/available", handler.GetAvailableRta)
			rta.PUT("/status", handler.UpdateRtaStatus)
			rta.PUT("/scope", handler.SetRtaScope)
			rta.GET("/scope", handler.GetRtaScope)
		}

		// 一键起量管理
		raise := advtools.Group("/raise")
		{
			raise.POST("", handler.SetAdRaise)
			raise.GET("/estimate", handler.GetAdRaiseEstimate)
			raise.GET("/status", handler.GetAdRaiseStatus)
			raise.GET("/result", handler.GetAdRaiseResult)
			raise.GET("/suggest-budget", handler.GetSuggestBudget)
		}

		// 定向包管理
		audience := advtools.Group("/audience")
		{
			audience.GET("", handler.GetAudiencePackage)
			audience.POST("", handler.CreateAudiencePackage)
			audience.PUT("/:package_id", handler.UpdateAudiencePackage)
			audience.DELETE("/:package_id", handler.DeleteAudiencePackage)
			audience.POST("/bind", handler.BindAudiencePackage)
			audience.POST("/unbind", handler.UnbindAudiencePackage)
		}

		// 原生锚点管理
		anchor := advtools.Group("/anchor")
		{
			anchor.GET("", handler.GetNativeAnchor)
			anchor.GET("/:anchor_id", handler.GetNativeAnchorDetail)
			anchor.POST("", handler.CreateNativeAnchor)
			anchor.PUT("/:anchor_id", handler.UpdateNativeAnchor)
			anchor.DELETE("/:anchor_id", handler.DeleteNativeAnchor)
		}

		// 诊断工具
		diagnosis := advtools.Group("/diagnosis")
		{
			diagnosis.GET("/suggestion", handler.GetDiagnosisSuggestion)
			diagnosis.POST("/accept", handler.AcceptDiagnosisSuggestion)
		}

		// 其他工具
		advtools.GET("/quota", handler.GetQuota)
		advtools.GET("/ad-quality", handler.GetAdQuality)
		advtools.GET("/ad-stat-extra", handler.GetAdStatExtraInfo)
	}
}

// registerSiteRoutes 注册建站管理路由
func (r *Router) registerSiteRoutes(rg *gin.RouterGroup) {
	handler := siteApi.NewSiteHandler(r.db, r.oceanCfg)

	site := rg.Group("/site")
	{
		// 橙子建站
		orange := site.Group("/orange")
		{
			orange.GET("", handler.GetOrangeSiteList)
			orange.GET("/:site_id", handler.GetOrangeSiteDetail)
			orange.POST("", handler.CreateOrangeSite)
			orange.PUT("/:site_id", handler.UpdateOrangeSite)
			orange.POST("/:site_id/publish", handler.PublishOrangeSite)
			orange.DELETE("/:site_id", handler.DeleteOrangeSite)
			orange.POST("/:site_id/copy", handler.CopyOrangeSite)
		}

		// 第三方落地页
		thirdparty := site.Group("/thirdparty")
		{
			thirdparty.GET("", handler.GetThirdPartySiteList)
			thirdparty.POST("", handler.CreateThirdPartySite)
			thirdparty.PUT("/:site_id", handler.UpdateThirdPartySite)
			thirdparty.DELETE("/:site_id", handler.DeleteThirdPartySite)
		}

		// 模板
		templates := site.Group("/templates")
		{
			templates.GET("", handler.GetSiteTemplateList)
			templates.POST("/create-site", handler.CreateSiteFromTemplate)
		}

		// 组件
		site.GET("/components", handler.GetSiteComponentList)

		// 分析
		analysis := site.Group("/analysis")
		{
			analysis.GET("", handler.GetSiteAnalysis)
			analysis.GET("/heatmap", handler.GetSiteHeatmap)
		}

		// 微页面
		mini := site.Group("/mini")
		{
			mini.GET("", handler.GetMiniPageList)
			mini.POST("", handler.CreateMiniPage)
			mini.DELETE("/:page_id", handler.DeleteMiniPage)
		}

		// 表单管理
		forms := site.Group("/forms")
		{
			forms.GET("", handler.GetSiteFormList)
			forms.GET("/submissions", handler.GetFormSubmissionList)
			forms.POST("/export", handler.ExportFormSubmissions)
		}
	}
}

// registerV3Routes 注册V3体验版路由
func (r *Router) registerV3Routes(rg *gin.RouterGroup) {
	handler := v3Api.NewV3Handler(r.db, r.oceanCfg)

	v3 := rg.Group("/v3")
	{
		// 项目管理
		projects := v3.Group("/projects")
		{
			projects.GET("", handler.GetProjectList)
			projects.GET("/:project_id", handler.GetProjectDetail)
			projects.POST("", handler.CreateProject)
			projects.PUT("/:project_id", handler.UpdateProject)
			projects.POST("/status", handler.UpdateProjectStatus)
			projects.DELETE("", handler.DeleteProject)
		}

		// 广告管理
		promotions := v3.Group("/promotions")
		{
			promotions.GET("", handler.GetPromotionList)
			promotions.GET("/:promotion_id", handler.GetPromotionDetail)
			promotions.POST("", handler.CreatePromotion)
			promotions.PUT("/:promotion_id", handler.UpdatePromotion)
			promotions.POST("/status", handler.UpdatePromotionStatus)
			promotions.POST("/budget", handler.UpdatePromotionBudget)
			promotions.POST("/bid", handler.UpdatePromotionBid)
			promotions.DELETE("", handler.DeletePromotion)
		}

	}
}

// registerDMPRoutes 注册DMP人群包路由
func (r *Router) registerDMPRoutes(rg *gin.RouterGroup) {
	handler := dmpApi.NewDMPHandler(r.db, r.oceanCfg)

	dmp := rg.Group("/dmp")
	{
		// 数据源管理
		datasource := dmp.Group("/datasource")
		{
			datasource.POST("/upload", handler.UploadDataSourceFile)
			datasource.POST("", handler.CreateDataSource)
			datasource.PUT("", handler.UpdateDataSource)
			datasource.GET("", handler.GetDataSourceDetail)
		}

		// 人群包管理
		audience := dmp.Group("/audience")
		{
			audience.GET("", handler.GetCustomAudienceList)
			audience.GET("/detail", handler.GetCustomAudienceDetail)
			audience.POST("", handler.CreateCustomAudience)
			audience.DELETE("", handler.DeleteCustomAudience)
			audience.POST("/publish", handler.PublishCustomAudience)
			audience.POST("/push", handler.PushCustomAudience)
		}
	}
}

// registerDPARoutes 注册DPA商品广告路由
func (r *Router) registerDPARoutes(rg *gin.RouterGroup) {
	handler := dpaApi.NewDPAHandler(r.db, r.oceanCfg)

	dpa := rg.Group("/dpa")
	{
		// 商品库管理
		libraries := dpa.Group("/libraries")
		{
			libraries.GET("", handler.GetProductLibraryList)
			libraries.POST("", handler.CreateProductLibrary)
			libraries.PUT("/:library_id", handler.UpdateProductLibrary)
			libraries.DELETE("/:library_id", handler.DeleteProductLibrary)
		}

		// 商品管理
		products := dpa.Group("/products")
		{
			products.GET("", handler.GetProductList)
			products.POST("", handler.CreateProduct)
			products.PUT("/:product_id", handler.UpdateProduct)
			products.DELETE("/:product_id", handler.DeleteProduct)
		}

		// 商品集管理
		sets := dpa.Group("/sets")
		{
			sets.GET("", handler.GetProductSetList)
			sets.POST("", handler.CreateProductSet)
			sets.PUT("/:set_id", handler.UpdateProductSet)
			sets.DELETE("/:set_id", handler.DeleteProductSet)
		}
	}
}
