import type { RouteRecordRaw } from 'vue-router'
import AppLayout from '@/components/layout/AppLayout.vue'

/**
 * 公开路由（无需认证）
 */
export const publicRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/LoginView.vue'),
    meta: { title: '登录', requiresAuth: false }
  },
  {
    path: '/403',
    name: 'Forbidden',
    component: () => import('@/views/error/403.vue'),
    meta: { title: '无权限', requiresAuth: false }
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
    meta: { title: '页面不存在', requiresAuth: false }
  }
]

/**
 * 需要认证的路由
 */
export const protectedRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    component: AppLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/DashboardView.vue'),
        meta: { title: '工作台', icon: 'home' }
      },
      // 广告主管理
      {
        path: 'advertisers',
        name: 'Advertisers',
        component: () => import('@/views/advertiser/AdvertiserList.vue'),
        meta: { title: '广告主管理', icon: 'users' }
      },
      {
        path: 'advertisers/:id',
        name: 'AdvertiserDetail',
        component: () => import('@/views/advertiser/AdvertiserDetail.vue'),
        meta: { title: '广告主详情', hidden: true }
      },
      // 广告系列管理
      {
        path: 'campaigns',
        name: 'Campaigns',
        component: () => import('@/views/campaign/CampaignList.vue'),
        meta: { title: '广告计划', icon: 'chart' }
      },
      {
        path: 'campaigns/create',
        name: 'CampaignCreate',
        component: () => import('@/views/campaign/CampaignCreate.vue'),
        meta: { title: '创建广告计划', hidden: true }
      },
      {
        path: 'campaigns/:id',
        name: 'CampaignDetail',
        component: () => import('@/views/campaign/CampaignDetail.vue'),
        meta: { title: '广告计划详情', hidden: true }
      },
      {
        path: 'campaigns/:id/edit',
        name: 'CampaignEdit',
        component: () => import('@/views/campaign/CampaignEdit.vue'),
        meta: { title: '编辑广告计划', hidden: true }
      },
      // 广告管理
      {
        path: 'ads',
        name: 'Ads',
        component: () => import('@/views/ad/AdList.vue'),
        meta: { title: '广告', icon: 'document' }
      },
      // 创意管理
      {
        path: 'creatives',
        name: 'Creatives',
        component: () => import('@/views/creative/CreativeList.vue'),
        meta: { title: '创意管理', icon: 'image' }
      },
      // 媒体库
      {
        path: 'media',
        name: 'Media',
        component: () => import('@/views/media/MediaLibrary.vue'),
        meta: { title: '媒体库', icon: 'folder' }
      },
      // 工具
      {
        path: 'tools/targeting',
        name: 'ToolsTargeting',
        component: () => import('@/views/tools/TargetingTools.vue'),
        meta: { title: '定向工具', icon: 'target' }
      },
      // 人群管理
      {
        path: 'audiences',
        name: 'Audiences',
        component: () => import('@/views/audience/AudienceList.vue'),
        meta: { title: '人群包', icon: 'group' }
      },
      {
        path: 'audiences/create',
        name: 'AudienceCreate',
        component: () => import('@/views/audience/AudienceEdit.vue'),
        meta: { title: '新建人群', hidden: true }
      },
      {
        path: 'audiences/:id/edit',
        name: 'AudienceEdit',
        component: () => import('@/views/audience/AudienceEdit.vue'),
        meta: { title: '编辑人群', hidden: true }
      },
      // 数据报表
      {
        path: 'reports',
        name: 'Reports',
        component: () => import('@/views/report/ReportDashboard.vue'),
        meta: { title: '数据报表', icon: 'chart-bar' }
      },
      {
        path: 'reports/:id',
        name: 'ReportDetail',
        component: () => import('@/views/report/ReportDetail.vue'),
        meta: { title: '报表详情', hidden: true }
      },
      // 代理商管理
      {
        path: 'agent',
        name: 'AgentDashboard',
        component: () => import('@/views/agent/AgentDashboard.vue'),
        meta: { title: '代理商工作台', icon: 'office' }
      },
      {
        path: 'agent/advertisers',
        name: 'AgentAdvertisers',
        component: () => import('@/views/agent/AgentAdvertisers.vue'),
        meta: { title: '广告主管理', hidden: true }
      },
      {
        path: 'agent/recharge',
        name: 'AgentRecharge',
        component: () => import('@/views/agent/AgentRecharge.vue'),
        meta: { title: '代理商充值', hidden: true }
      },
      {
        path: 'agent/transfer',
        name: 'AgentTransfer',
        component: () => import('@/views/agent/AgentTransfer.vue'),
        meta: { title: '转账记录', hidden: true }
      },
      // 共享钱包
      {
        path: 'wallet',
        name: 'WalletDailyStat',
        component: () => import('@/views/wallet/WalletDailyStat.vue'),
        meta: { title: '钱包日流水', icon: 'wallet' }
      },
      {
        path: 'wallet/transactions',
        name: 'WalletTransactions',
        component: () => import('@/views/wallet/WalletTransactions.vue'),
        meta: { title: '流水明细', hidden: true }
      },
      {
        path: 'wallet/transfer',
        name: 'WalletTransfer',
        component: () => import('@/views/wallet/WalletTransfer.vue'),
        meta: { title: '钱包转账', hidden: true }
      },
      // 基础工具
      {
        path: 'tools/region',
        name: 'ToolsRegion',
        component: () => import('@/views/tools/basic/ToolsRegion.vue'),
        meta: { title: '地域列表', hidden: true }
      },
      {
        path: 'tools/industry',
        name: 'ToolsIndustry',
        component: () => import('@/views/tools/basic/ToolsIndustry.vue'),
        meta: { title: '行业列表', hidden: true }
      },
      {
        path: 'tools/quota',
        name: 'ToolsQuota',
        component: () => import('@/views/tools/basic/ToolsQuota.vue'),
        meta: { title: '配额查询', hidden: true }
      },
      {
        path: 'tools/bid-suggest',
        name: 'ToolsBidSuggest',
        component: () => import('@/views/tools/basic/ToolsBidSuggest.vue'),
        meta: { title: '建议日预算', hidden: true }
      },
      {
        path: 'tools/ad-quality',
        name: 'ToolsAdQuality',
        component: () => import('@/views/tools/basic/ToolsAdQuality.vue'),
        meta: { title: '广告质量度', hidden: true }
      },
      {
        path: 'tools/ad-learning',
        name: 'ToolsAdLearning',
        component: () => import('@/views/tools/basic/ToolsAdLearning.vue'),
        meta: { title: '学习期状态', hidden: true }
      },
      {
        path: 'tools/preview-qrcode',
        name: 'ToolsPreviewQrcode',
        component: () => import('@/views/tools/basic/ToolsPreviewQrcode.vue'),
        meta: { title: '广告预览', hidden: true }
      },
      {
        path: 'tools/country-info',
        name: 'ToolsCountryInfo',
        component: () => import('@/views/tools/basic/ToolsCountryInfo.vue'),
        meta: { title: '国家信息', hidden: true }
      },
      // 工作台模块
      {
        path: 'workspace/transfer-targets',
        name: 'WorkspaceTransferTargets',
        component: () => import('@/views/workspace/WorkspaceTransferTargets.vue'),
        meta: { title: '可转账列表', hidden: true }
      },
      {
        path: 'workspace/transfer/create',
        name: 'WorkspaceTransferCreate',
        component: () => import('@/views/workspace/WorkspaceTransferCreate.vue'),
        meta: { title: '发起转账', hidden: true }
      },
      // 代理商扩展
      {
        path: 'agent/child-agents',
        name: 'AgentChildAgents',
        component: () => import('@/views/agent/AgentChildAgents.vue'),
        meta: { title: '二级代理商', hidden: true }
      },
      {
        path: 'agent/info',
        name: 'AgentInfo',
        component: () => import('@/views/agent/AgentInfo.vue'),
        meta: { title: '代理商信息', hidden: true }
      },
      // 关键词模块
      {
        path: 'keyword/create',
        name: 'KeywordCreate',
        component: () => import('@/views/keyword/KeywordCreate.vue'),
        meta: { title: '创建关键词', hidden: true }
      },
      {
        path: 'keyword/suggest',
        name: 'KeywordSuggest',
        component: () => import('@/views/keyword/KeywordSuggest.vue'),
        meta: { title: '推荐关键词', hidden: true }
      },
      // DPA商品模块
      {
        path: 'dpa/products',
        name: 'DpaProductList',
        component: () => import('@/views/dpa/DpaProductList.vue'),
        meta: { title: 'DPA商品库', hidden: true }
      },
      // 站点模块
      {
        path: 'site',
        name: 'SiteList',
        component: () => import('@/views/site/SiteList.vue'),
        meta: { title: '橙子建站', hidden: true }
      },
      // 线索模块
      {
        path: 'clue/forms',
        name: 'ClueFormList',
        component: () => import('@/views/clue/ClueFormList.vue'),
        meta: { title: '青鸟表单', hidden: true }
      },
      {
        path: 'clue/form/create',
        name: 'ClueFormCreate',
        component: () => import('@/views/clue/ClueFormCreate.vue'),
        meta: { title: '创建表单', hidden: true }
      },
      // 落地页模块
      {
        path: 'site/create',
        name: 'SiteCreate',
        component: () => import('@/views/site/SiteCreate.vue'),
        meta: { title: '创建落地页', hidden: true }
      },
      // DPA商品详情
      {
        path: 'dpa/product/:id',
        name: 'DpaProductDetail',
        component: () => import('@/views/dpa/DpaProductDetail.vue'),
        meta: { title: '商品详情', hidden: true }
      },
      // 诊断工具
      {
        path: 'tools/diagnosis/suggestion',
        name: 'ToolsDiagnosisSuggestion',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisSuggestion.vue'),
        meta: { title: '计划诊断', hidden: true }
      },
      // 兴趣工具
      {
        path: 'tools/interest/category',
        name: 'ToolsInterestCategory',
        component: () => import('@/views/tools/interest/ToolsInterestCategory.vue'),
        meta: { title: '兴趣类目', hidden: true }
      },
      // 转化工具
      {
        path: 'tools/adconvert',
        name: 'ToolsAdconvertList',
        component: () => import('@/views/tools/adconvert/ToolsAdconvertList.vue'),
        meta: { title: '转化目标', hidden: true }
      },
      // 人群工具
      {
        path: 'tools/audience',
        name: 'ToolsAudiencePackageList',
        component: () => import('@/views/tools/audience/ToolsAudiencePackageList.vue'),
        meta: { title: '人群包', hidden: true }
      },
      // RTA工具
      {
        path: 'tools/rta',
        name: 'ToolsRtaList',
        component: () => import('@/views/tools/rta/ToolsRtaList.vue'),
        meta: { title: 'RTA策略', hidden: true }
      },
      // 评论工具
      {
        path: 'tools/comment',
        name: 'ToolsCommentList',
        component: () => import('@/views/tools/comment/ToolsCommentList.vue'),
        meta: { title: '评论管理', hidden: true }
      },
      // 抖音工具
      {
        path: 'tools/aweme/info',
        name: 'ToolsAwemeInfo',
        component: () => import('@/views/tools/aweme/ToolsAwemeInfo.vue'),
        meta: { title: '账号信息', hidden: true }
      },
      // 创意工具
      {
        path: 'tools/creative/word',
        name: 'ToolsCreativeWord',
        component: () => import('@/views/tools/creative/ToolsCreativeWord.vue'),
        meta: { title: '文案生成', hidden: true }
      },
      // 穿山甲联盟
      {
        path: 'tools/union',
        name: 'ToolsUnionList',
        component: () => import('@/views/tools/union/ToolsUnionList.vue'),
        meta: { title: '穿山甲', hidden: true }
      },
      // 自定义报表
      {
        path: 'report/custom',
        name: 'ReportCustomList',
        component: () => import('@/views/report/ReportCustomList.vue'),
        meta: { title: '自定义报表', hidden: true }
      },
      // 素材管理
      {
        path: 'material/video',
        name: 'MaterialVideoList',
        component: () => import('@/views/material/MaterialVideoList.vue'),
        meta: { title: '视频素材', hidden: true }
      },
      {
        path: 'material/image',
        name: 'MaterialImageList',
        component: () => import('@/views/material/MaterialImageList.vue'),
        meta: { title: '图片素材', hidden: true }
      },
      {
        path: 'material/template',
        name: 'MaterialTemplate',
        component: () => import('@/views/material/MaterialTemplate.vue'),
        meta: { title: '素材模板', hidden: true }
      },
      // 提价工具
      {
        path: 'tools/adraise',
        name: 'ToolsAdraise',
        component: () => import('@/views/tools/adraise/ToolsAdraise.vue'),
        meta: { title: '智能提价', hidden: true }
      },
      // DPA分类
      {
        path: 'dpa/category',
        name: 'DpaCategoryList',
        component: () => import('@/views/dpa/DpaCategoryList.vue'),
        meta: { title: '商品分类', hidden: true }
      },
      // 站点分组
      {
        path: 'site/group',
        name: 'LandingGroupList',
        component: () => import('@/views/site/LandingGroupList.vue'),
        meta: { title: '站点分组', hidden: true }
      },
      // 数据任务
      {
        path: 'report/task',
        name: 'ReportDataTask',
        component: () => import('@/views/report/ReportDataTask.vue'),
        meta: { title: '数据任务', hidden: true }
      },
      // 其他功能
      {
        path: 'other/douplus',
        name: 'OtherDouplus',
        component: () => import('@/views/other/OtherDouplus.vue'),
        meta: { title: 'Dou+管理', hidden: true }
      },
      {
        path: 'other/track',
        name: 'OtherTrack',
        component: () => import('@/views/other/OtherTrack.vue'),
        meta: { title: '转化追踪', hidden: true }
      },
      {
        path: 'other/servemarket',
        name: 'OtherServeMarket',
        component: () => import('@/views/other/OtherServeMarket.vue'),
        meta: { title: '服务市场', hidden: true }
      },
      // ==================== 千川模块 ====================
      {
        path: 'qianchuan',
        name: 'QianchuanDashboard',
        component: () => import('@/views/qianchuan/QianchuanDashboard.vue'),
        meta: { title: '千川工作台', icon: 'qianchuan' }
      },
      {
        path: 'qianchuan/account',
        name: 'QianchuanAccount',
        component: () => import('@/views/qianchuan/AccountInfo.vue'),
        meta: { title: '账户信息', hidden: true }
      },
      {
        path: 'qianchuan/budget',
        name: 'QianchuanBudget',
        component: () => import('@/views/qianchuan/AccountBudget.vue'),
        meta: { title: '账户预算', hidden: true }
      },
      {
        path: 'qianchuan/shop',
        name: 'QianchuanShop',
        component: () => import('@/views/qianchuan/ShopList.vue'),
        meta: { title: '店铺列表', hidden: true }
      },
      {
        path: 'qianchuan/campaign',
        name: 'QianchuanCampaign',
        component: () => import('@/views/qianchuan/CampaignList.vue'),
        meta: { title: '千川计划', hidden: true }
      },
      {
        path: 'qianchuan/campaign/create',
        name: 'QianchuanCampaignCreate',
        component: () => import('@/views/qianchuan/CampaignCreate.vue'),
        meta: { title: '创建计划', hidden: true }
      },
      {
        path: 'qianchuan/ad',
        name: 'QianchuanAd',
        component: () => import('@/views/qianchuan/AdList.vue'),
        meta: { title: '千川广告', hidden: true }
      },
      {
        path: 'qianchuan/ad/create',
        name: 'QianchuanAdCreate',
        component: () => import('@/views/qianchuan/AdCreate.vue'),
        meta: { title: '创建广告', hidden: true }
      },
      {
        path: 'qianchuan/ad/:id',
        name: 'QianchuanAdDetail',
        component: () => import('@/views/qianchuan/AdDetail.vue'),
        meta: { title: '广告详情', hidden: true }
      },
      {
        path: 'qianchuan/creative',
        name: 'QianchuanCreative',
        component: () => import('@/views/qianchuan/CreativeList.vue'),
        meta: { title: '千川创意', hidden: true }
      },
      {
        path: 'qianchuan/material',
        name: 'QianchuanMaterial',
        component: () => import('@/views/qianchuan/MaterialList.vue'),
        meta: { title: '素材管理', hidden: true }
      },
      {
        path: 'qianchuan/product',
        name: 'QianchuanProduct',
        component: () => import('@/views/qianchuan/ProductList.vue'),
        meta: { title: '商品管理', hidden: true }
      },
      {
        path: 'qianchuan/uni',
        name: 'QianchuanUni',
        component: () => import('@/views/qianchuan/UniList.vue'),
        meta: { title: '全域推广', hidden: true }
      },
      {
        path: 'qianchuan/uni/create',
        name: 'QianchuanUniCreate',
        component: () => import('@/views/qianchuan/UniCreate.vue'),
        meta: { title: '创建全域', hidden: true }
      },
      {
        path: 'qianchuan/aweme-order',
        name: 'QianchuanAwemeOrder',
        component: () => import('@/views/qianchuan/AwemeOrderList.vue'),
        meta: { title: '随心推', hidden: true }
      },
      {
        path: 'qianchuan/report/advertiser',
        name: 'QianchuanReportAdvertiser',
        component: () => import('@/views/qianchuan/ReportAdvertiser.vue'),
        meta: { title: '账户报表', hidden: true }
      },
      {
        path: 'qianchuan/report/live',
        name: 'QianchuanReportLive',
        component: () => import('@/views/qianchuan/ReportLive.vue'),
        meta: { title: '直播报表', hidden: true }
      },
      {
        path: 'qianchuan/finance',
        name: 'QianchuanFinance',
        component: () => import('@/views/qianchuan/FinanceWallet.vue'),
        meta: { title: '财务管理', hidden: true }
      },
      {
        path: 'qianchuan/keyword',
        name: 'QianchuanKeyword',
        component: () => import('@/views/qianchuan/KeywordList.vue'),
        meta: { title: '关键词', hidden: true }
      },
      // ==================== 本地推模块 ====================
      {
        path: 'local',
        name: 'LocalDashboard',
        component: () => import('@/views/local/LocalDashboard.vue'),
        meta: { title: '本地推工作台', icon: 'local' }
      },
      {
        path: 'local/project',
        name: 'LocalProject',
        component: () => import('@/views/local/ProjectList.vue'),
        meta: { title: '项目管理', hidden: true }
      },
      {
        path: 'local/project/create',
        name: 'LocalProjectCreate',
        component: () => import('@/views/local/ProjectCreate.vue'),
        meta: { title: '创建项目', hidden: true }
      },
      {
        path: 'local/project/:id',
        name: 'LocalProjectDetail',
        component: () => import('@/views/local/ProjectDetail.vue'),
        meta: { title: '项目详情', hidden: true }
      },
      {
        path: 'local/promotion',
        name: 'LocalPromotion',
        component: () => import('@/views/local/PromotionList.vue'),
        meta: { title: '推广管理', hidden: true }
      },
      {
        path: 'local/promotion/create',
        name: 'LocalPromotionCreate',
        component: () => import('@/views/local/PromotionCreate.vue'),
        meta: { title: '创建广告', hidden: true }
      },
      {
        path: 'local/promotion/:id',
        name: 'LocalPromotionDetail',
        component: () => import('@/views/local/PromotionDetail.vue'),
        meta: { title: '广告详情', hidden: true }
      },
      {
        path: 'local/clue',
        name: 'LocalClue',
        component: () => import('@/views/local/ClueList.vue'),
        meta: { title: '线索管理', hidden: true }
      },
      {
        path: 'local/report/project',
        name: 'LocalReportProject',
        component: () => import('@/views/local/ReportProject.vue'),
        meta: { title: '项目报表', hidden: true }
      },
      {
        path: 'local/report/promotion',
        name: 'LocalReportPromotion',
        component: () => import('@/views/local/ReportPromotion.vue'),
        meta: { title: '广告报表', hidden: true }
      },
      {
        path: 'local/report/material',
        name: 'LocalReportMaterial',
        component: () => import('@/views/local/ReportMaterial.vue'),
        meta: { title: '素材报表', hidden: true }
      },
      {
        path: 'local/material/video',
        name: 'LocalVideoList',
        component: () => import('@/views/local/VideoList.vue'),
        meta: { title: '视频管理', hidden: true }
      },
      // ==================== 企业号模块 ====================
      {
        path: 'enterprise',
        name: 'EnterpriseDashboard',
        component: () => import('@/views/enterprise/EnterpriseDashboard.vue'),
        meta: { title: '企业号工作台', icon: 'enterprise' }
      },
      {
        path: 'enterprise/info',
        name: 'EnterpriseInfo',
        component: () => import('@/views/enterprise/EnterpriseInfo.vue'),
        meta: { title: '账号信息', hidden: true }
      },
      {
        path: 'enterprise/comment',
        name: 'EnterpriseComment',
        component: () => import('@/views/enterprise/CommentList.vue'),
        meta: { title: '评论管理', hidden: true }
      },
      {
        path: 'enterprise/item',
        name: 'EnterpriseItem',
        component: () => import('@/views/enterprise/ItemList.vue'),
        meta: { title: '视频列表', hidden: true }
      },
      {
        path: 'enterprise/overview',
        name: 'EnterpriseOverview',
        component: () => import('@/views/enterprise/OverviewData.vue'),
        meta: { title: '数据概览', hidden: true }
      },
      // ==================== 星图模块 ====================
      {
        path: 'star',
        name: 'StarDashboard',
        component: () => import('@/views/star/StarDashboard.vue'),
        meta: { title: '星图工作台', icon: 'star' }
      },
      {
        path: 'star/task',
        name: 'StarTask',
        component: () => import('@/views/star/TaskList.vue'),
        meta: { title: '任务管理', hidden: true }
      },
      {
        path: 'star/demand',
        name: 'StarDemand',
        component: () => import('@/views/star/DemandList.vue'),
        meta: { title: '需求管理', hidden: true }
      },
      {
        path: 'star/fund',
        name: 'StarFund',
        component: () => import('@/views/star/FundBalance.vue'),
        meta: { title: '资金管理', hidden: true }
      },
      {
        path: 'star/report',
        name: 'StarReport',
        component: () => import('@/views/star/ReportOverview.vue'),
        meta: { title: '数据报表', hidden: true }
      },
      {
        path: 'star/account',
        name: 'StarAccountInfo',
        component: () => import('@/views/star/AccountInfo.vue'),
        meta: { title: '账户信息', hidden: true }
      },
      {
        path: 'star/clue',
        name: 'StarClueList',
        component: () => import('@/views/star/ClueList.vue'),
        meta: { title: '线索管理', hidden: true }
      },
      {
        path: 'star/task/:id',
        name: 'StarTaskDetail',
        component: () => import('@/views/star/TaskDetail.vue'),
        meta: { title: '任务详情', hidden: true }
      },
      {
        path: 'star/tasks',
        name: 'StarTaskItem',
        component: () => import('@/views/star/TaskItem.vue'),
        meta: { title: '任务列表', hidden: true }
      },
      {
        path: 'star/fund/daily',
        name: 'StarFundDaily',
        component: () => import('@/views/star/FundDaily.vue'),
        meta: { title: '资金日流水', hidden: true }
      },
      {
        path: 'star/fund/transaction',
        name: 'StarFundTransaction',
        component: () => import('@/views/star/FundTransaction.vue'),
        meta: { title: '交易明细', hidden: true }
      },
      {
        path: 'star/demand/order',
        name: 'StarDemandOrder',
        component: () => import('@/views/star/DemandOrder.vue'),
        meta: { title: '需求单管理', hidden: true }
      },
      {
        path: 'star/report/audience',
        name: 'StarReportAudience',
        component: () => import('@/views/star/ReportAudience.vue'),
        meta: { title: '受众分析', hidden: true }
      },
      {
        path: 'star/report/daily',
        name: 'StarReportDaily',
        component: () => import('@/views/star/ReportDaily.vue'),
        meta: { title: '每日报告', hidden: true }
      },
      {
        path: 'star/agent/advertisers',
        name: 'StarAgentAdvertisers',
        component: () => import('@/views/star/AgentAdvertisers.vue'),
        meta: { title: '代理商广告主', hidden: true }
      },
      // ==================== 服务市场模块 ====================
      {
        path: 'servemarket',
        name: 'ServeMarketDashboard',
        component: () => import('@/views/servemarket/ServeMarketDashboard.vue'),
        meta: { title: '服务市场', icon: 'market' }
      },
      {
        path: 'servemarket/order',
        name: 'ServeMarketOrder',
        component: () => import('@/views/servemarket/ServeMarketOrder.vue'),
        meta: { title: '订单管理', hidden: true }
      },
      {
        path: 'servemarket/func',
        name: 'ServeMarketFunc',
        component: () => import('@/views/servemarket/ServeMarketFunc.vue'),
        meta: { title: '功能点管理', hidden: true }
      },
      {
        path: 'servemarket/quality',
        name: 'ServeMarketQuality',
        component: () => import('@/views/servemarket/ServeMarketQuality.vue'),
        meta: { title: '质量报告', hidden: true }
      },
      {
        path: 'servemarket/subscribe',
        name: 'ServeMarketSubscribe',
        component: () => import('@/views/servemarket/ServeMarketSubscribe.vue'),
        meta: { title: '订阅管理', hidden: true }
      },
      // ==================== 系统管理 ====================
      {
        path: 'system/users',
        name: 'UserManage',
        component: () => import('@/views/system/UserManage.vue'),
        meta: { title: '用户管理', icon: 'user' }
      },
      {
        path: 'system/roles',
        name: 'RoleManage',
        component: () => import('@/views/system/RoleManage.vue'),
        meta: { title: '角色管理', icon: 'shield' }
      },
      {
        path: 'system/menus',
        name: 'MenuManage',
        component: () => import('@/views/system/MenuManage.vue'),
        meta: { title: '菜单管理', icon: 'menu' }
      }
    ]
  },
  // 404 兜底路由
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

/**
 * 所有路由
 */
export const routes: RouteRecordRaw[] = [...publicRoutes, ...protectedRoutes]
