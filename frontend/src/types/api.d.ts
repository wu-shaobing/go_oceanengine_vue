/**
 * API 响应基础类型
 */
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
}

/**
 * 分页响应
 */
export interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  page_size: number
}

/**
 * 分页请求参数
 */
export interface PageParams {
  page?: number
  page_size?: number
}

/**
 * 广告主相关类型
 */
export interface Advertiser {
  id: number
  advertiser_id: string
  name: string
  company: string
  status: 'active' | 'disabled' | 'pending'
  balance: number
  daily_budget: number
  industry: string
  contact: string
  email: string
  phone: string
  created_at: string
  updated_at: string
}

export interface AdvertiserListParams extends PageParams {
  keyword?: string
  status?: string
  industry?: string
}

/**
 * 广告计划相关类型
 */
export interface Campaign {
  id: number
  campaign_id: string
  advertiser_id: string
  name: string
  status: 'active' | 'paused' | 'deleted'
  budget: number
  budget_mode: 'daily' | 'total'
  landing_type: 'app' | 'link' | 'quick_app'
  start_time: string
  end_time: string
  created_at: string
  updated_at: string
}

export interface CampaignListParams extends PageParams {
  advertiser_id?: string
  keyword?: string
  status?: string
}

/**
 * 广告创意相关类型
 */
export interface Creative {
  id: number
  creative_id: string
  campaign_id: string
  name: string
  status: 'active' | 'paused' | 'under_review' | 'rejected'
  creative_type: 'image' | 'video' | 'carousel'
  title: string
  description: string
  image_url?: string
  video_url?: string
  click_url: string
  created_at: string
  updated_at: string
}

/**
 * 报表相关类型
 */
export interface ReportMetrics {
  impressions: number
  clicks: number
  conversions: number
  cost: number
  ctr: number
  cvr: number
  cpc: number
  cpa: number
}

export interface ReportData extends ReportMetrics {
  date: string
  advertiser_id?: string
  campaign_id?: string
  creative_id?: string
}

export interface ReportParams {
  start_date: string
  end_date: string
  advertiser_id?: string
  campaign_id?: string
  group_by?: 'day' | 'week' | 'month'
  metrics?: string[]
}

/**
 * 人群包相关类型
 */
export interface Audience {
  id: number
  audience_id: string
  name: string
  description: string
  type: 'custom' | 'lookalike' | 'retargeting' | 'dmp'
  status: 'active' | 'processing' | 'disabled'
  size: number
  tags: string[]
  created_at: string
  updated_at: string
}

export interface AudienceListParams extends PageParams {
  keyword?: string
  type?: string
  status?: string
}

/**
 * 用户相关类型
 */
export interface User {
  id: number
  username: string
  email: string
  phone: string
  avatar?: string
  role: string
  status: 'active' | 'disabled'
  last_login_at?: string
  created_at: string
  updated_at: string
}

export interface LoginParams {
  username: string
  password: string
  remember?: boolean
}

export interface LoginResult {
  token: string
  user: User
  expires_in: number
}

/**
 * 角色相关类型
 */
export interface Role {
  id: number
  name: string
  key: string
  description: string
  permissions: string[]
  status: 'active' | 'disabled'
  created_at: string
  updated_at: string
}

/**
 * 菜单相关类型
 */
export interface Menu {
  id: number
  name: string
  path: string
  icon?: string
  component?: string
  parent_id?: number
  sort: number
  status: 'active' | 'disabled'
  children?: Menu[]
}
