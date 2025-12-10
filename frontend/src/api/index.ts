// 基础工具
export { default as request, type PageResponse, type ApiResponse } from './request'

// 基础模块
export { authApi, type LoginRequest, type LoginResponse, type UserInfo } from './auth'

// 系统管理
export {
  systemApi,
  type User,
  type UserListParams,
  type UserCreateParams,
  type UserUpdateParams,
  type Role,
  type RoleListParams,
  type RoleCreateParams,
  type RoleUpdateParams,
  type Menu,
  type MenuTree,
  type MenuCreateParams,
  type MenuUpdateParams,
  type OperationLog,
  type OperationLogListParams
} from './system'
export { advertiserApi, type Advertiser, type AdvertiserListParams } from './advertiser'
export { campaignApi, type Campaign, type CampaignCreateRequest, type CampaignListParams } from './campaign'
export { reportApi, type ReportData, type ReportParams } from './report'

// 广告管理
export { adApi, type Ad, type AdCreateParams, type AdAudience, type AdListParams } from './ad'
export { creativeApi, type Creative, type CreativeCreateParams, type CreativeListParams } from './creative'
export { materialApi, type ImageMaterial, type VideoMaterial, type MaterialGroup } from './material'

// 千川模块
export {
  qianchuanApi,
  type QianchuanAccount,
  type QianchuanShop,
  type QianchuanCampaign,
  type QianchuanAd,
  type QianchuanCreative,
  type AwemeOrder,
  type QianchuanReportData,
  type QianchuanMaterial,
  type Product
} from './qianchuan'

// 企业号模块
export {
  enterpriseApi,
  type EnterpriseInfo,
  type EnterpriseBind,
  type EnterpriseItem,
  type EnterpriseComment,
  type EnterpriseOverviewData
} from './enterprise'

// 本地推模块
export {
  localApi,
  type LocalProject,
  type LocalPromotion,
  type LocalClue,
  type LocalReportData,
  type LocalStore
} from './local'

// 星图模块
export {
  starApi,
  type StarAccount,
  type StarTask,
  type StarTaskDetail,
  type StarDemand,
  type StarReportOverview,
  type StarClue
} from './star'

// 服务市场模块
export {
  serveMarketApi,
  type ServeMarketOrder,
  type ServeMarketFunc,
  type ServeMarketQuality,
  type RDSSubscription
} from './servemarket'
