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
      // 个人设置
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('@/views/settings/Settings.vue'),
        meta: { title: '个人设置', hidden: true }
      },
      // 账户信息
      {
        path: 'account',
        name: 'Account',
        component: () => import('@/views/settings/Account.vue'),
        meta: { title: '账户信息', hidden: true }
      },
      // 广告主管理
      {
        path: 'advertisers',
        name: 'Advertisers',
        component: () => import('@/views/advertiser/AdvertiserList.vue'),
        meta: { title: '广告主管理', icon: 'users' }
      },
      {
        path: 'advertisers/create',
        name: 'AdvertiserCreate',
        component: () => import('@/views/advertiser/AdvertiserCreate.vue'),
        meta: { title: '创建广告主', hidden: true }
      },
      {
        path: 'advertisers/:id',
        name: 'AdvertiserDetail',
        component: () => import('@/views/advertiser/AdvertiserDetail.vue'),
        meta: { title: '广告主详情', hidden: true }
      },
      {
        path: 'advertiser/audit',
        name: 'AdvertiserAudit',
        component: () => import('@/views/advertiser/AdvertiserAudit.vue'),
        meta: { title: '广告主审核', hidden: true }
      },
      {
        path: 'advertiser/contract',
        name: 'AdvertiserContract',
        component: () => import('@/views/advertiser/AdvertiserContract.vue'),
        meta: { title: '广告主合同', hidden: true }
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
      {
        path: 'ads/create',
        name: 'AdCreate',
        component: () => import('@/views/ad/AdCreate.vue'),
        meta: { title: '创建广告', hidden: true }
      },
      {
        path: 'ads/:id',
        name: 'AdDetail',
        component: () => import('@/views/ad/AdDetail.vue'),
        meta: { title: '广告详情', hidden: true }
      },
      {
        path: 'ads/:id/edit',
        name: 'AdEdit',
        component: () => import('@/views/ad/AdEdit.vue'),
        meta: { title: '编辑广告', hidden: true }
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
        path: 'agent/advertiser/create',
        name: 'AgentAdvertiserCreate',
        component: () => import('@/views/agent/AgentAdvertisers.vue'),
        meta: { title: '创建广告主', hidden: true }
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
      {
        path: 'wallet/invoice',
        name: 'WalletInvoice',
        component: () => import('@/views/wallet/WalletInvoice.vue'),
        meta: { title: '发票管理', hidden: true }
      },
      {
        path: 'wallet/contract',
        name: 'WalletContract',
        component: () => import('@/views/wallet/WalletContract.vue'),
        meta: { title: '合同管理', hidden: true }
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
      {
        path: 'workspace/task',
        name: 'WorkspaceTask',
        component: () => import('@/views/workspace/WorkspaceTask.vue'),
        meta: { title: '任务列表', hidden: true }
      },
      {
        path: 'workspace/alert',
        name: 'WorkspaceAlert',
        component: () => import('@/views/workspace/WorkspaceAlert.vue'),
        meta: { title: '告警管理', hidden: true }
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
      {
        path: 'agent/auth',
        name: 'AgentAuth',
        component: () => import('@/views/agent/AgentAuth.vue'),
        meta: { title: '代理商授权', hidden: true }
      },
      {
        path: 'agent/account',
        name: 'AgentAccount',
        component: () => import('@/views/agent/AgentAccount.vue'),
        meta: { title: '代理商账户', hidden: true }
      },
      // 关键词模块
      {
        path: 'keyword/list',
        name: 'KeywordList',
        component: () => import('@/views/keyword/KeywordList.vue'),
        meta: { title: '关键词列表', hidden: true }
      },
      {
        path: 'keyword/create',
        name: 'KeywordCreate',
        component: () => import('@/views/keyword/KeywordCreate.vue'),
        meta: { title: '创建关键词', hidden: true }
      },
      {
        path: 'keyword/edit/:id',
        name: 'KeywordEdit',
        component: () => import('@/views/keyword/KeywordEdit.vue'),
        meta: { title: '编辑关键词', hidden: true }
      },
      {
        path: 'keyword/suggest',
        name: 'KeywordSuggest',
        component: () => import('@/views/keyword/KeywordSuggest.vue'),
        meta: { title: '推荐关键词', hidden: true }
      },
      {
        path: 'keyword/bidding',
        name: 'KeywordBidding',
        component: () => import('@/views/keyword/KeywordBidding.vue'),
        meta: { title: '关键词出价', hidden: true }
      },
      {
        path: 'keyword/blueflow',
        name: 'BlueflowKeywords',
        component: () => import('@/views/keyword/BlueflowKeywords.vue'),
        meta: { title: '蓝流量关键词', hidden: true }
      },
      {
        path: 'keyword/exclude',
        name: 'KeywordExclude',
        component: () => import('@/views/keyword/KeywordExclude.vue'),
        meta: { title: '否定关键词', hidden: true }
      },
      {
        path: 'keyword/privative-word/project',
        name: 'PrivativeWordProject',
        component: () => import('@/views/keyword/PrivativeWordProject.vue'),
        meta: { title: '项目否定词', hidden: true }
      },
      {
        path: 'keyword/privative-word/campaign',
        name: 'PrivativeWordCampaign',
        component: () => import('@/views/keyword/PrivativeWordCampaign.vue'),
        meta: { title: '计划否定词', hidden: true }
      },
      {
        path: 'keyword/blueflow-packages',
        name: 'BlueflowPackages',
        component: () => import('@/views/keyword/BlueflowPackages.vue'),
        meta: { title: '蓝流量包', hidden: true }
      },
      {
        path: 'keyword/report',
        name: 'KeywordReport',
        component: () => import('@/views/keyword/KeywordReport.vue'),
        meta: { title: '关键词报表', hidden: true }
      },
      // DPA商品模块
      {
        path: 'dpa/products',
        name: 'DpaProductList',
        component: () => import('@/views/dpa/DpaProductList.vue'),
        meta: { title: 'DPA商品库', hidden: true }
      },
      {
        path: 'dpa/product/create',
        name: 'DpaProductCreate',
        component: () => import('@/views/dpa/DpaProductCreate.vue'),
        meta: { title: '创建商品', hidden: true }
      },
      // 站点模块
      {
        path: 'site',
        name: 'SiteList',
        component: () => import('@/views/site/SiteList.vue'),
        meta: { title: '橙子建站', hidden: true }
      },
      {
        path: 'site/list',
        name: 'SiteListAlias',
        component: () => import('@/views/site/SiteList.vue'),
        meta: { title: '站点列表', hidden: true }
      },
      // 线索模块
      {
        path: 'clue/forms',
        name: 'ClueFormList',
        component: () => import('@/views/clue/ClueFormList.vue'),
        meta: { title: '青鸟表单', hidden: true }
      },
      {
        path: 'clue/form/list',
        name: 'ClueFormListAlias',
        component: () => import('@/views/clue/ClueFormList.vue'),
        meta: { title: '青鸟表单', hidden: true }
      },
      {
        path: 'clue/coupon/create',
        name: 'ClueCouponCreate',
        component: () => import('@/views/clue/ClueCouponCreate.vue'),
        meta: { title: '创建卡券', hidden: true }
      },
      {
        path: 'clue/coupon/edit/:id',
        name: 'ClueCouponEdit',
        component: () => import('@/views/clue/ClueCouponCreate.vue'),
        meta: { title: '编辑卡券', hidden: true }
      },
      {
        path: 'clue/form/create',
        name: 'ClueFormCreate',
        component: () => import('@/views/clue/ClueFormCreate.vue'),
        meta: { title: '创建表单', hidden: true }
      },
      {
        path: 'clue/form/edit/:id',
        name: 'ClueFormEdit',
        component: () => import('@/views/clue/ClueFormCreate.vue'),
        meta: { title: '编辑表单', hidden: true }
      },
      {
        path: 'clue/form/detail/:id',
        name: 'ClueFormDetailView',
        component: () => import('@/views/clue/ClueFormDetail.vue'),
        meta: { title: '表单详情', hidden: true }
      },
      {
        path: 'clue/smartphone/list',
        name: 'ClueSmartphoneList',
        component: () => import('@/views/clue/ClueSmartphoneList.vue'),
        meta: { title: '智能电话', hidden: true }
      },
      {
        path: 'clue/smartphone/create',
        name: 'ClueSmartphoneCreate',
        component: () => import('@/views/clue/ClueSmartphoneCreate.vue'),
        meta: { title: '创建智能电话', hidden: true }
      },
      {
        path: 'clue/coupon/list',
        name: 'ClueCouponList',
        component: () => import('@/views/clue/ClueCouponList.vue'),
        meta: { title: '卡券列表', hidden: true }
      },
      {
        path: 'clue/coupon/code-upload',
        name: 'ClueCouponCodeUpload',
        component: () => import('@/views/clue/ClueCouponCodeUpload.vue'),
        meta: { title: '上传卡码', hidden: true }
      },
      {
        path: 'clue/callback-config',
        name: 'ClueCallbackConfig',
        component: () => import('@/views/clue/ClueCallbackConfig.vue'),
        meta: { title: '回调配置', hidden: true }
      },
      {
        path: 'clue/feyu-callback',
        name: 'ClueFeyuCallback',
        component: () => import('@/views/clue/ClueFeyuCallback.vue'),
        meta: { title: '飞鱼回调', hidden: true }
      },
      {
        path: 'clue/assign',
        name: 'ClueAssign',
        component: () => import('@/views/clue/ClueAssign.vue'),
        meta: { title: '线索分配', hidden: true }
      },
      {
        path: 'clue/key-action',
        name: 'ClueKeyAction',
        component: () => import('@/views/clue/ClueKeyAction.vue'),
        meta: { title: '关键行为', hidden: true }
      },
      // 落地页模块
      {
        path: 'site/create',
        name: 'SiteCreate',
        component: () => import('@/views/site/SiteCreate.vue'),
        meta: { title: '创建落地页', hidden: true }
      },
      {
        path: 'site/edit/:id',
        name: 'SiteEdit',
        component: () => import('@/views/site/SiteEdit.vue'),
        meta: { title: '编辑落地页', hidden: true }
      },
      {
        path: 'site/landing-page',
        name: 'SiteLandingPage',
        component: () => import('@/views/site/SiteLandingPage.vue'),
        meta: { title: '落地页管理', hidden: true }
      },
      {
        path: 'site/form/create',
        name: 'SiteFormCreate',
        component: () => import('@/views/site/SiteFormCreate.vue'),
        meta: { title: '创建表单', hidden: true }
      },
      {
        path: 'site/thirdsite',
        name: 'ThirdsiteList',
        component: () => import('@/views/site/ThirdsiteList.vue'),
        meta: { title: '第三方站点', hidden: true }
      },
      {
        path: 'site/thirdsite/create',
        name: 'ThirdsiteCreate',
        component: () => import('@/views/site/ThirdsiteCreate.vue'),
        meta: { title: '创建第三方站点', hidden: true }
      },
      {
        path: 'site/form/list',
        name: 'SiteFormList',
        component: () => import('@/views/site/SiteFormList.vue'),
        meta: { title: '表单列表', hidden: true }
      },
      {
        path: 'site/stats',
        name: 'SiteStats',
        component: () => import('@/views/site/SiteStats.vue'),
        meta: { title: '站点统计', hidden: true }
      },
      {
        path: 'site/preview/:id',
        name: 'SitePreview',
        component: () => import('@/views/site/SitePreview.vue'),
        meta: { title: '站点预览', hidden: true }
      },
      {
        path: 'site/template',
        name: 'SiteTemplate',
        component: () => import('@/views/site/SiteTemplate.vue'),
        meta: { title: '站点模板', hidden: true }
      },
      {
        path: 'site/template/list',
        name: 'SiteTemplateList',
        component: () => import('@/views/site/SiteTemplateList.vue'),
        meta: { title: '模板列表', hidden: true }
      },
      // DPA商品详情
      {
        path: 'dpa/product/:id',
        name: 'DpaProductDetail',
        component: () => import('@/views/dpa/DpaProductDetail.vue'),
        meta: { title: '商品详情', hidden: true }
      },
      {
        path: 'dpa/product/edit/:id',
        name: 'DpaProductEdit',
        component: () => import('@/views/dpa/DpaProductUpdate.vue'),
        meta: { title: '编辑商品', hidden: true }
      },
      {
        path: 'dpa/product/upload',
        name: 'DpaProductUpload',
        component: () => import('@/views/dpa/DpaProductUpload.vue'),
        meta: { title: '上传商品', hidden: true }
      },
      {
        path: 'dpa/dict',
        name: 'DpaDict',
        component: () => import('@/views/dpa/DpaDict.vue'),
        meta: { title: '商品字典', hidden: true }
      },
      {
        path: 'dpa/assets',
        name: 'DpaAssetsList',
        component: () => import('@/views/dpa/DpaAssetsList.vue'),
        meta: { title: '资产列表', hidden: true }
      },
      {
        path: 'dpa/report',
        name: 'DpaReport',
        component: () => import('@/views/dpa/DpaReport.vue'),
        meta: { title: 'DPA报表', hidden: true }
      },
      {
        path: 'dpa/template/create',
        name: 'DpaTemplateCreate',
        component: () => import('@/views/dpa/DpaTemplateCreate.vue'),
        meta: { title: '创建DPA模板', hidden: true }
      },
      {
        path: 'dpa/behavior',
        name: 'DpaBehavior',
        component: () => import('@/views/dpa/DpaBehavior.vue'),
        meta: { title: '行为管理', hidden: true }
      },
      {
        path: 'dpa/video-template',
        name: 'DpaVideoTemplate',
        component: () => import('@/views/dpa/DpaVideoTemplate.vue'),
        meta: { title: '视频模板', hidden: true }
      },
      {
        path: 'dpa/product/sync',
        name: 'DpaProductSync',
        component: () => import('@/views/dpa/DpaProductSync.vue'),
        meta: { title: '商品同步', hidden: true }
      },
      // 诊断工具
      {
        path: 'tools/diagnosis/suggestion',
        name: 'ToolsDiagnosisSuggestion',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisSuggestion.vue'),
        meta: { title: '计划诊断', hidden: true }
      },
      {
        path: 'tools/diagnosis/creative',
        name: 'ToolsDiagnosisCreative',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisCreative.vue'),
        meta: { title: '创意诊断', hidden: true }
      },
      {
        path: 'tools/diagnosis/account',
        name: 'ToolsDiagnosisAccount',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisAccount.vue'),
        meta: { title: '账户诊断', hidden: true }
      },
      {
        path: 'tools/diagnosis/ad',
        name: 'ToolsDiagnosisAd',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisAd.vue'),
        meta: { title: '广告诊断', hidden: true }
      },
      {
        path: 'tools/diagnosis/accept',
        name: 'ToolsDiagnosisAccept',
        component: () => import('@/views/tools/diagnosis/ToolsDiagnosisAccept.vue'),
        meta: { title: '起量诊断', hidden: true }
      },
      // 兴趣工具
      {
        path: 'tools/interest/category',
        name: 'ToolsInterestCategory',
        component: () => import('@/views/tools/interest/ToolsInterestCategory.vue'),
        meta: { title: '兴趣类目', hidden: true }
      },
      {
        path: 'tools/interest/keyword',
        name: 'ToolsInterestKeyword',
        component: () => import('@/views/tools/interest/ToolsInterestKeyword.vue'),
        meta: { title: '兴趣关键词', hidden: true }
      },
      {
        path: 'tools/interest/behavior',
        name: 'ToolsInterestBehavior',
        component: () => import('@/views/tools/interest/ToolsInterestBehavior.vue'),
        meta: { title: '行为兴趣', hidden: true }
      },
      {
        path: 'tools/interest/action',
        name: 'ToolsInterestAction',
        component: () => import('@/views/tools/interest/ToolsInterestAction.vue'),
        meta: { title: '行为定向', hidden: true }
      },
      // 转化工具
      {
        path: 'tools/adconvert',
        name: 'ToolsAdconvertList',
        component: () => import('@/views/tools/adconvert/ToolsAdconvertList.vue'),
        meta: { title: '转化目标', hidden: true }
      },
      {
        path: 'tools/adconvert/create',
        name: 'ToolsAdconvertCreate',
        component: () => import('@/views/tools/adconvert/ToolsAdconvertCreate.vue'),
        meta: { title: '创建转化', hidden: true }
      },
      {
        path: 'tools/adconvert/detail/:id',
        name: 'ToolsAdconvertDetail',
        component: () => import('@/views/tools/adconvert/ToolsAdconvertDetail.vue'),
        meta: { title: '转化详情', hidden: true }
      },
      // 人群工具
      {
        path: 'tools/audience',
        name: 'ToolsAudiencePackageList',
        component: () => import('@/views/tools/audience/ToolsAudiencePackageList.vue'),
        meta: { title: '人群包', hidden: true }
      },
      {
        path: 'tools/audience/create',
        name: 'ToolsAudiencePackageCreate',
        component: () => import('@/views/tools/audience/ToolsAudiencePackageCreate.vue'),
        meta: { title: '创建人群包', hidden: true }
      },
      {
        path: 'tools/audience/lookalike',
        name: 'ToolsAudienceLookalike',
        component: () => import('@/views/tools/audience/ToolsAudienceLookalike.vue'),
        meta: { title: '相似人群', hidden: true }
      },
      {
        path: 'tools/audience/list',
        name: 'ToolsAudienceList',
        component: () => import('@/views/tools/audience/ToolsAudienceList.vue'),
        meta: { title: '人群列表', hidden: true }
      },
      {
        path: 'tools/audience/upload',
        name: 'ToolsAudienceUpload',
        component: () => import('@/views/tools/audience/ToolsAudienceUpload.vue'),
        meta: { title: '上传人群', hidden: true }
      },
      // RTA工具
      {
        path: 'tools/rta',
        name: 'ToolsRtaList',
        component: () => import('@/views/tools/rta/ToolsRtaList.vue'),
        meta: { title: 'RTA策略', hidden: true }
      },
      {
        path: 'tools/rta/create',
        name: 'ToolsRtaCreate',
        component: () => import('@/views/tools/rta/ToolsRtaCreate.vue'),
        meta: { title: '创建RTA策略', hidden: true }
      },
      {
        path: 'tools/rta/logs',
        name: 'ToolsRtaLogs',
        component: () => import('@/views/tools/rta/ToolsRtaLogs.vue'),
        meta: { title: 'RTA日志', hidden: true }
      },
      // 评论工具
      {
        path: 'tools/comment',
        name: 'ToolsCommentList',
        component: () => import('@/views/tools/comment/ToolsCommentList.vue'),
        meta: { title: '评论管理', hidden: true }
      },
      {
        path: 'tools/comment/reply',
        name: 'ToolsCommentReply',
        component: () => import('@/views/tools/comment/ToolsCommentReply.vue'),
        meta: { title: '评论回复', hidden: true }
      },
      {
        path: 'tools/comment/filter',
        name: 'ToolsCommentFilter',
        component: () => import('@/views/tools/comment/ToolsCommentFilter.vue'),
        meta: { title: '评论过滤', hidden: true }
      },
      // 抖音工具
      {
        path: 'tools/aweme/info',
        name: 'ToolsAwemeInfo',
        component: () => import('@/views/tools/aweme/ToolsAwemeInfo.vue'),
        meta: { title: '账号信息', hidden: true }
      },
      {
        path: 'tools/aweme/account',
        name: 'ToolsAwemeAccount',
        component: () => import('@/views/tools/aweme/ToolsAwemeAccount.vue'),
        meta: { title: '抖音账号', hidden: true }
      },
      {
        path: 'tools/aweme/similar',
        name: 'ToolsAwemeSimilar',
        component: () => import('@/views/tools/aweme/ToolsAwemeSimilar.vue'),
        meta: { title: '相似达人', hidden: true }
      },
      {
        path: 'tools/aweme/category',
        name: 'ToolsAwemeCategory',
        component: () => import('@/views/tools/aweme/ToolsAwemeCategory.vue'),
        meta: { title: '抖音分类', hidden: true }
      },
      {
        path: 'tools/aweme/video',
        name: 'ToolsAwemeVideo',
        component: () => import('@/views/tools/aweme/ToolsAwemeVideo.vue'),
        meta: { title: '视频搜索', hidden: true }
      },
      {
        path: 'tools/aweme/author-info',
        name: 'ToolsAwemeAuthorInfo',
        component: () => import('@/views/tools/aweme/ToolsAwemeAuthorInfo.vue'),
        meta: { title: '达人信息', hidden: true }
      },
      {
        path: 'tools/aweme/top-author',
        name: 'ToolsAwemeTopAuthor',
        component: () => import('@/views/tools/aweme/ToolsAwemeTopAuthor.vue'),
        meta: { title: '热门达人', hidden: true }
      },
      // 创意工具
      {
        path: 'tools/creative/word',
        name: 'ToolsCreativeWord',
        component: () => import('@/views/tools/creative/ToolsCreativeWord.vue'),
        meta: { title: '文案生成', hidden: true }
      },
      {
        path: 'tools/creative/test',
        name: 'ToolsCreativeTest',
        component: () => import('@/views/tools/creative/ToolsCreativeTest.vue'),
        meta: { title: '创意测试', hidden: true }
      },
      {
        path: 'tools/creative/word-expand',
        name: 'ToolsCreativeWordExpand',
        component: () => import('@/views/tools/creative/ToolsCreativeWordExpand.vue'),
        meta: { title: '文案扩展', hidden: true }
      },
      {
        path: 'tools/creative/review',
        name: 'ToolsCreativeReview',
        component: () => import('@/views/tools/creative/ToolsCreativeReview.vue'),
        meta: { title: '创意审核', hidden: true }
      },
      // 穿山甲联盟
      {
        path: 'tools/union',
        name: 'ToolsUnionList',
        component: () => import('@/views/tools/union/ToolsUnionList.vue'),
        meta: { title: '穿山甲', hidden: true }
      },
      {
        path: 'tools/union/create',
        name: 'ToolsUnionCreate',
        component: () => import('@/views/tools/union/ToolsUnionApp.vue'),
        meta: { title: '创建应用', hidden: true }
      },
      {
        path: 'tools/union/report',
        name: 'ToolsUnionReport',
        component: () => import('@/views/tools/union/ToolsUnionReport.vue'),
        meta: { title: '联盟报表', hidden: true }
      },
      {
        path: 'tools/union/slot',
        name: 'ToolsUnionSlot',
        component: () => import('@/views/tools/union/ToolsUnionSlot.vue'),
        meta: { title: '广告位', hidden: true }
      },
      // 自定义报表
      {
        path: 'report/custom',
        name: 'ReportCustomList',
        component: () => import('@/views/report/ReportCustomList.vue'),
        meta: { title: '自定义报表', hidden: true }
      },
      {
        path: 'report/custom/create',
        name: 'ReportCustomCreate',
        component: () => import('@/views/report/ReportCustom.vue'),
        meta: { title: '创建自定义报表', hidden: true }
      },
      {
        path: 'report/comparison',
        name: 'ReportComparison',
        component: () => import('@/views/report/ReportComparison.vue'),
        meta: { title: '对比报表', hidden: true }
      },
      {
        path: 'report/ad-account',
        name: 'ReportAdAccount',
        component: () => import('@/views/report/ReportAdAccount.vue'),
        meta: { title: '账户报表', hidden: true }
      },
      {
        path: 'report/schedule',
        name: 'ReportSchedule',
        component: () => import('@/views/report/ReportSchedule.vue'),
        meta: { title: '定时报表', hidden: true }
      },
      {
        path: 'report/realtime',
        name: 'ReportRealtime',
        component: () => import('@/views/report/ReportRealtime.vue'),
        meta: { title: '实时报表', hidden: true }
      },
      {
        path: 'report/ad-creative',
        name: 'ReportAdCreative',
        component: () => import('@/views/report/ReportAdCreative.vue'),
        meta: { title: '创意报表', hidden: true }
      },
      {
        path: 'report/geo',
        name: 'ReportGeo',
        component: () => import('@/views/report/ReportGeo.vue'),
        meta: { title: '地域报表', hidden: true }
      },
      {
        path: 'report/audience',
        name: 'ReportAudience',
        component: () => import('@/views/report/ReportAudience.vue'),
        meta: { title: '受众报表', hidden: true }
      },
      {
        path: 'report/ad-campaign',
        name: 'ReportAdCampaign',
        component: () => import('@/views/report/ReportAdCampaign.vue'),
        meta: { title: '计划报表', hidden: true }
      },
      {
        path: 'report/export',
        name: 'ReportExport',
        component: () => import('@/views/report/ReportExport.vue'),
        meta: { title: '报表导出', hidden: true }
      },
      // 素材管理
      {
        path: 'material/video',
        name: 'MaterialVideoList',
        component: () => import('@/views/material/MaterialVideoList.vue'),
        meta: { title: '视频素材', hidden: true }
      },
      {
        path: 'material/video/upload',
        name: 'MaterialVideoUpload',
        component: () => import('@/views/material/MaterialVideoCreate.vue'),
        meta: { title: '上传视频', hidden: true }
      },
      {
        path: 'material/image',
        name: 'MaterialImageList',
        component: () => import('@/views/material/MaterialImageList.vue'),
        meta: { title: '图片素材', hidden: true }
      },
      {
        path: 'material/image/upload',
        name: 'MaterialImageUpload',
        component: () => import('@/views/material/MaterialImageCreate.vue'),
        meta: { title: '上传图片', hidden: true }
      },
      {
        path: 'material/template',
        name: 'MaterialTemplate',
        component: () => import('@/views/material/MaterialTemplate.vue'),
        meta: { title: '素材模板', hidden: true }
      },
      {
        path: 'material/image/edit/:id',
        name: 'MaterialImageEdit',
        component: () => import('@/views/material/MaterialImageEdit.vue'),
        meta: { title: '编辑图片', hidden: true }
      },
      {
        path: 'material/recycle',
        name: 'MaterialRecycle',
        component: () => import('@/views/material/MaterialRecycle.vue'),
        meta: { title: '回收站', hidden: true }
      },
      {
        path: 'material/ai-generate',
        name: 'MaterialAIGenerate',
        component: () => import('@/views/material/MaterialAIGenerate.vue'),
        meta: { title: 'AI生成', hidden: true }
      },
      {
        path: 'material/analysis',
        name: 'MaterialAnalysis',
        component: () => import('@/views/material/MaterialAnalysis.vue'),
        meta: { title: '素材分析', hidden: true }
      },
      {
        path: 'material/video/edit/:id',
        name: 'MaterialVideoEdit',
        component: () => import('@/views/material/MaterialVideoEdit.vue'),
        meta: { title: '编辑视频', hidden: true }
      },
      {
        path: 'material/folder',
        name: 'MaterialFolder',
        component: () => import('@/views/material/MaterialFolder.vue'),
        meta: { title: '素材文件夹', hidden: true }
      },
      // 提价工具
      {
        path: 'tools/adraise',
        name: 'ToolsAdraise',
        component: () => import('@/views/tools/adraise/ToolsAdraise.vue'),
        meta: { title: '智能提价', hidden: true }
      },
      {
        path: 'tools/adraise/history',
        name: 'ToolsAdraiseHistory',
        component: () => import('@/views/tools/adraise/ToolsAdraiseHistory.vue'),
        meta: { title: '提价历史', hidden: true }
      },
      {
        path: 'tools/adraise/strategy',
        name: 'ToolsAdraiseStrategy',
        component: () => import('@/views/tools/adraise/ToolsAdraiseStrategy.vue'),
        meta: { title: '提价策略', hidden: true }
      },
      {
        path: 'tools/adraise/create',
        name: 'ToolsAdraiseCreate',
        component: () => import('@/views/tools/adraise/ToolsAdraiseCreate.vue'),
        meta: { title: '创建提价', hidden: true }
      },
      // 高级工具
      {
        path: 'tools/advanced/security-score',
        name: 'ToolsSecurityScore',
        component: () => import('@/views/tools/advanced/ToolsSecurityScore.vue'),
        meta: { title: '安全评分', hidden: true }
      },
      {
        path: 'tools/advanced/promotion-card',
        name: 'ToolsPromotionCard',
        component: () => import('@/views/tools/advanced/ToolsPromotionCard.vue'),
        meta: { title: '推广卡片', hidden: true }
      },
      {
        path: 'tools/advanced/gray-release',
        name: 'ToolsGrayRelease',
        component: () => import('@/views/tools/advanced/ToolsGrayRelease.vue'),
        meta: { title: '灰度发布', hidden: true }
      },
      {
        path: 'tools/advanced/ies-account',
        name: 'ToolsIesAccount',
        component: () => import('@/views/tools/advanced/ToolsIesAccount.vue'),
        meta: { title: 'IES账号', hidden: true }
      },
      {
        path: 'tools/advanced/automation',
        name: 'ToolsAdvancedAutomation',
        component: () => import('@/views/tools/advanced/ToolsAdvancedAutomation.vue'),
        meta: { title: '自动化工具', hidden: true }
      },
      {
        path: 'tools/advanced/action-text',
        name: 'ToolsActionText',
        component: () => import('@/views/tools/advanced/ToolsActionText.vue'),
        meta: { title: '行动文案', hidden: true }
      },
      {
        path: 'tools/advanced/native-anchor',
        name: 'ToolsNativeAnchor',
        component: () => import('@/views/tools/advanced/ToolsNativeAnchor.vue'),
        meta: { title: '原生键词', hidden: true }
      },
      {
        path: 'tools/advanced/aweme-auth',
        name: 'ToolsAwemeAuth',
        component: () => import('@/views/tools/advanced/ToolsAwemeAuth.vue'),
        meta: { title: '抖音授权', hidden: true }
      },
      {
        path: 'tools/advanced/copy',
        name: 'ToolsAdvancedCopy',
        component: () => import('@/views/tools/advanced/ToolsAdvancedCopy.vue'),
        meta: { title: '复制工具', hidden: true }
      },
      {
        path: 'tools/advanced/batch',
        name: 'ToolsAdvancedBatch',
        component: () => import('@/views/tools/advanced/ToolsAdvancedBatch.vue'),
        meta: { title: '批量操作', hidden: true }
      },
      {
        path: 'tools/advanced/anchor-check',
        name: 'ToolsAnchorCheck',
        component: () => import('@/views/tools/advanced/ToolsAnchorCheck.vue'),
        meta: { title: '键词检查', hidden: true }
      },
      // 应用工具
      {
        path: 'tools/app/package-parse',
        name: 'ToolsAppPackageParse',
        component: () => import('@/views/tools/app/ToolsAppPackageParse.vue'),
        meta: { title: '包解析', hidden: true }
      },
      {
        path: 'tools/app/deeplink',
        name: 'ToolsAppDeeplink',
        component: () => import('@/views/tools/app/ToolsAppDeeplink.vue'),
        meta: { title: '深度链接', hidden: true }
      },
      {
        path: 'tools/app/promotion',
        name: 'ToolsAppPromotion',
        component: () => import('@/views/tools/app/ToolsAppPromotion.vue'),
        meta: { title: '应用推广', hidden: true }
      },
      {
        path: 'tools/app/share',
        name: 'ToolsAppShare',
        component: () => import('@/views/tools/app/ToolsAppShare.vue'),
        meta: { title: '分享管理', hidden: true }
      },
      {
        path: 'tools/app/extend-package',
        name: 'ToolsAppExtendPackage',
        component: () => import('@/views/tools/app/ToolsAppExtendPackage.vue'),
        meta: { title: '扩展包', hidden: true }
      },
      {
        path: 'tools/app/booking',
        name: 'ToolsAppBooking',
        component: () => import('@/views/tools/app/ToolsAppBooking.vue'),
        meta: { title: '预约管理', hidden: true }
      },
      // 基础工具扩展
      {
        path: 'tools/basic/schedule',
        name: 'ToolsBasicSchedule',
        component: () => import('@/views/tools/basic/ToolsBasicSchedule.vue'),
        meta: { title: '时段工具', hidden: true }
      },
      {
        path: 'tools/basic/bid',
        name: 'ToolsBasicBid',
        component: () => import('@/views/tools/basic/ToolsBasicBid.vue'),
        meta: { title: '出价工具', hidden: true }
      },
      {
        path: 'tools/basic/budget',
        name: 'ToolsBasicBudget',
        component: () => import('@/views/tools/basic/ToolsBasicBudget.vue'),
        meta: { title: '预算工具', hidden: true }
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
        path: 'other/douplus/order',
        name: 'OtherDouplusOrder',
        component: () => import('@/views/other/OtherDouplusOrder.vue'),
        meta: { title: 'Dou+订单', hidden: true }
      },
      {
        path: 'other/douplus/data',
        name: 'OtherDouplusData',
        component: () => import('@/views/other/OtherDouplusData.vue'),
        meta: { title: 'Dou+数据', hidden: true }
      },
      {
        path: 'other/track',
        name: 'OtherTrack',
        component: () => import('@/views/other/OtherTrack.vue'),
        meta: { title: '转化追踪', hidden: true }
      },
      {
        path: 'other/track/list',
        name: 'OtherTrackList',
        component: () => import('@/views/other/OtherTrackList.vue'),
        meta: { title: '追踪列表', hidden: true }
      },
      {
        path: 'other/track/create',
        name: 'OtherTrackCreate',
        component: () => import('@/views/other/OtherTrackCreate.vue'),
        meta: { title: '创建追踪', hidden: true }
      },
      {
        path: 'other/track/callback',
        name: 'OtherTrackCallback',
        component: () => import('@/views/other/OtherTrackCallback.vue'),
        meta: { title: '追踪回调', hidden: true }
      },
      {
        path: 'other/pixel-manage',
        name: 'OtherPixelManage',
        component: () => import('@/views/other/OtherPixelManage.vue'),
        meta: { title: '像素管理', hidden: true }
      },
      {
        path: 'other/event-track',
        name: 'OtherEventTrack',
        component: () => import('@/views/other/OtherEventTrack.vue'),
        meta: { title: '事件追踪', hidden: true }
      },
      {
        path: 'other/servemarket',
        name: 'OtherServeMarket',
        component: () => import('@/views/other/OtherServeMarket.vue'),
        meta: { title: '服务市场', hidden: true }
      },
      {
        path: 'other/servemarket/list',
        name: 'OtherServeMarketList',
        component: () => import('@/views/other/OtherServeMarketList.vue'),
        meta: { title: '服务列表', hidden: true }
      },
      {
        path: 'other/servemarket/detail/:id',
        name: 'OtherServeMarketDetail',
        component: () => import('@/views/other/OtherServeMarketDetail.vue'),
        meta: { title: '服务详情', hidden: true }
      },
      {
        path: 'other/api-log',
        name: 'OtherApiLog',
        component: () => import('@/views/other/OtherApiLog.vue'),
        meta: { title: 'API日志', hidden: true }
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
        path: 'qianchuan/material/video',
        name: 'QianchuanMaterialVideo',
        component: () => import('@/views/qianchuan/MaterialVideo.vue'),
        meta: { title: '视频素材', hidden: true }
      },
      {
        path: 'qianchuan/material/image',
        name: 'QianchuanMaterialImage',
        component: () => import('@/views/qianchuan/MaterialImage.vue'),
        meta: { title: '图片素材', hidden: true }
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
        path: 'qianchuan/uni/:id',
        name: 'QianchuanUniDetail',
        component: () => import('@/views/qianchuan/UniDetail.vue'),
        meta: { title: '全域详情', hidden: true }
      },
      {
        path: 'qianchuan/aweme-order',
        name: 'QianchuanAwemeOrder',
        component: () => import('@/views/qianchuan/AwemeOrderList.vue'),
        meta: { title: '随心推', hidden: true }
      },
      {
        path: 'qianchuan/aweme-order/create',
        name: 'QianchuanAwemeOrderCreate',
        component: () => import('@/views/qianchuan/AwemeOrderCreate.vue'),
        meta: { title: '创建随心推', hidden: true }
      },
      {
        path: 'qianchuan/aweme-order/:id',
        name: 'QianchuanAwemeOrderDetail',
        component: () => import('@/views/qianchuan/AwemeOrderDetail.vue'),
        meta: { title: '随心推详情', hidden: true }
      },
      {
        path: 'qianchuan/report',
        name: 'QianchuanReport',
        component: () => import('@/views/qianchuan/ReportAdvertiser.vue'),
        meta: { title: '数据报表', hidden: true }
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
      {
        path: 'qianchuan/keyword/recommend',
        name: 'QianchuanKeywordRecommend',
        component: () => import('@/views/qianchuan/KeywordRecommend.vue'),
        meta: { title: '关键词推荐', hidden: true }
      },
      {
        path: 'qianchuan/finance/detail',
        name: 'QianchuanFinanceDetail',
        component: () => import('@/views/qianchuan/FinanceDetail.vue'),
        meta: { title: '财务明细', hidden: true }
      },
      // 千川扩展报表
      {
        path: 'qianchuan/report/creative',
        name: 'QianchuanReportCreative',
        component: () => import('@/views/qianchuan/ReportCreative.vue'),
        meta: { title: '创意报表', hidden: true }
      },
      {
        path: 'qianchuan/report/room',
        name: 'QianchuanReportRoom',
        component: () => import('@/views/qianchuan/ReportRoom.vue'),
        meta: { title: '直播间报表', hidden: true }
      },
      {
        path: 'qianchuan/live-room/detail/:id',
        name: 'QianchuanLiveRoomDetail',
        component: () => import('@/views/qianchuan/LiveRoomDetail.vue'),
        meta: { title: '直播间详情', hidden: true }
      },
      {
        path: 'qianchuan/report/ad',
        name: 'QianchuanReportAd',
        component: () => import('@/views/qianchuan/ReportAd.vue'),
        meta: { title: '广告报表', hidden: true }
      },
      {
        path: 'qianchuan/tools/dmp',
        name: 'QianchuanToolsDmp',
        component: () => import('@/views/qianchuan/ToolsDmp.vue'),
        meta: { title: 'DMP工具', hidden: true }
      },
      {
        path: 'qianchuan/aweme-auth',
        name: 'QianchuanAwemeAuth',
        component: () => import('@/views/qianchuan/AwemeAuth.vue'),
        meta: { title: '抖音授权', hidden: true }
      },
      {
        path: 'qianchuan/report/material',
        name: 'QianchuanReportMaterial',
        component: () => import('@/views/qianchuan/ReportMaterial.vue'),
        meta: { title: '素材报表', hidden: true }
      },
      {
        path: 'qianchuan/report/custom',
        name: 'QianchuanReportCustom',
        component: () => import('@/views/qianchuan/ReportCustom.vue'),
        meta: { title: '自定义报表', hidden: true }
      },
      {
        path: 'qianchuan/report/keyword',
        name: 'QianchuanReportKeyword',
        component: () => import('@/views/qianchuan/ReportKeyword.vue'),
        meta: { title: '关键词报表', hidden: true }
      },
      {
        path: 'qianchuan/report/uni',
        name: 'QianchuanReportUni',
        component: () => import('@/views/qianchuan/ReportUni.vue'),
        meta: { title: '全域报表', hidden: true }
      },
      {
        path: 'qianchuan/tools/industry',
        name: 'QianchuanToolsIndustry',
        component: () => import('@/views/qianchuan/ToolsIndustry.vue'),
        meta: { title: '行业工具', hidden: true }
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
        path: 'local/promotion/edit/:id',
        name: 'LocalPromotionEdit',
        component: () => import('@/views/local/PromotionCreate.vue'),
        meta: { title: '编辑广告', hidden: true }
      },
      {
        path: 'local/clue',
        name: 'LocalClue',
        component: () => import('@/views/local/ClueList.vue'),
        meta: { title: '线索管理', hidden: true }
      },
      {
        path: 'local/report',
        name: 'LocalReport',
        component: () => import('@/views/local/ReportProject.vue'),
        meta: { title: '数据报表', hidden: true }
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
      {
        path: 'enterprise/bind',
        name: 'EnterpriseBind',
        component: () => import('@/views/enterprise/BindList.vue'),
        meta: { title: '绑定管理', hidden: true }
      },
      {
        path: 'enterprise/log',
        name: 'EnterpriseLog',
        component: () => import('@/views/enterprise/OperationLog.vue'),
        meta: { title: '操作日志', hidden: true }
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
        path: 'star/task/create',
        name: 'StarTaskCreate',
        component: () => import('@/views/star/TaskCreate.vue'),
        meta: { title: '创建任务', hidden: true }
      },
      {
        path: 'star/demand',
        name: 'StarDemand',
        component: () => import('@/views/star/DemandList.vue'),
        meta: { title: '需求管理', hidden: true }
      },
      {
        path: 'star/demand/create',
        name: 'StarDemandCreate',
        component: () => import('@/views/star/DemandCreate.vue'),
        meta: { title: '发布需求', hidden: true }
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
      },
      {
        path: 'system/logs',
        name: 'SystemLogs',
        component: () => import('@/views/system/SystemLogs.vue'),
        meta: { title: '操作日志', hidden: true }
      },
      {
        path: 'system/settings',
        name: 'SystemSettings',
        component: () => import('@/views/system/SystemSettings.vue'),
        meta: { title: '系统设置', hidden: true }
      },
      {
        path: 'system/notifications',
        name: 'SystemNotifications',
        component: () => import('@/views/system/SystemNotifications.vue'),
        meta: { title: '消息通知', hidden: true }
      },
      {
        path: 'system/dict',
        name: 'DictManage',
        component: () => import('@/views/system/DictManage.vue'),
        meta: { title: '字典管理', hidden: true }
      }
    ]
  },
  // 404 兑底路由
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

/**
 * 所有路由
 */
export const routes: RouteRecordRaw[] = [...publicRoutes, ...protectedRoutes]
