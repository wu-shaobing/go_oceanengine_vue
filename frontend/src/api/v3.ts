import request, { PageResponse } from './request'

// ==================== 项目类型 ====================

export interface V3Project {
  project_id: number
  advertiser_id: number
  project_name: string
  landing_type: string
  marketing_goal: string
  status: string
  opt_status: string
  budget_mode: string
  budget: number
  delivery_range: string
  create_time: string
  modify_time: string
}

export interface V3ProjectDetail extends V3Project {
  audience: V3Audience
  delivery_setting: V3DeliverySetting
  tracking_url?: string
}

export interface V3Audience {
  gender: string
  age: string[]
  city: number[]
  location_type: string
  interest_categories: number[]
  action_categories: number[]
  retargeting_tags_include: number[]
  retargeting_tags_exclude: number[]
}

export interface V3DeliverySetting {
  inventory_type: string[]
  schedule_type: string
  schedule_time?: string
  start_time?: string
  end_time?: string
}

export interface V3ProjectCreateParams {
  advertiser_id: number
  project_name: string
  landing_type: string
  marketing_goal: string
  budget_mode: string
  budget?: number
  delivery_range: string
  audience?: Partial<V3Audience>
  delivery_setting?: Partial<V3DeliverySetting>
}

// ==================== 广告(Promotion)类型 ====================

export interface V3Promotion {
  promotion_id: number
  advertiser_id: number
  project_id: number
  promotion_name: string
  status: string
  opt_status: string
  budget_mode: string
  budget: number
  bid_type: string
  bid: number
  cpa_bid?: number
  deep_bid_type?: string
  deep_cpabid?: number
  roi_goal?: number
  create_time: string
  modify_time: string
}

export interface V3PromotionDetail extends V3Promotion {
  materials: V3Material[]
  native_setting?: V3NativeSetting
}

export interface V3Material {
  material_id: string
  material_type: string
  video_id?: string
  image_ids?: string[]
  title?: string
  description?: string
}

export interface V3NativeSetting {
  anchor_id?: number
  call_to_action?: string
}

export interface V3PromotionCreateParams {
  advertiser_id: number
  project_id: number
  promotion_name: string
  budget_mode: string
  budget?: number
  bid_type: string
  bid: number
  cpa_bid?: number
  deep_bid_type?: string
  deep_cpabid?: number
  roi_goal?: number
  materials?: Partial<V3Material>[]
  native_setting?: Partial<V3NativeSetting>
}

// ==================== 预算组类型 ====================

export interface V3BudgetGroup {
  budget_group_id: number
  advertiser_id: number
  budget_group_name: string
  budget: number
  budget_mode: string
  status: string
  project_ids: number[]
  create_time: string
  modify_time: string
}

export interface V3BudgetGroupCreateParams {
  advertiser_id: number
  budget_group_name: string
  budget: number
  budget_mode: string
  project_ids?: number[]
}

// ==================== 报表类型 ====================

export interface V3ReportData {
  stat_datetime?: string
  project_id?: number
  promotion_id?: number
  cost: number
  show_cnt: number
  click_cnt: number
  convert_cnt: number
  ctr: number
  cpm: number
  cpc: number
  cpa: number
  attribution_convert_cnt?: number
  deep_convert_cnt?: number
}

// ==================== API 方法 ====================

export const v3Api = {
  // ========== 项目管理 ==========
  
  // 获取项目列表
  getProjectList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<V3Project>>('/v3/project/list', params)
  },

  // 获取项目详情
  getProjectDetail(project_id: number, advertiser_id: number) {
    return request.get<V3ProjectDetail>(`/v3/project/${project_id}`, { advertiser_id })
  },

  // 创建项目
  createProject(data: V3ProjectCreateParams) {
    return request.post<{ project_id: number }>('/v3/project/create', data)
  },

  // 更新项目
  updateProject(project_id: number, data: Partial<V3ProjectCreateParams>) {
    return request.put<void>(`/v3/project/${project_id}`, data)
  },

  // 更新项目状态
  updateProjectStatus(data: { advertiser_id: number; project_ids: number[]; opt_status: string }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/project/status/update', data)
  },

  // 删除项目
  deleteProject(data: { advertiser_id: number; project_ids: number[] }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/project/delete', data)
  },

  // 更新项目预算
  updateProjectBudget(data: { advertiser_id: number; project_ids: number[]; budget: number }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/project/budget/update', data)
  },

  // ========== 广告管理 ==========

  // 获取广告列表
  getPromotionList(params: { advertiser_id: number; project_id?: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<V3Promotion>>('/v3/promotion/list', params)
  },

  // 获取广告详情
  getPromotionDetail(promotion_id: number, advertiser_id: number) {
    return request.get<V3PromotionDetail>(`/v3/promotion/${promotion_id}`, { advertiser_id })
  },

  // 创建广告
  createPromotion(data: V3PromotionCreateParams) {
    return request.post<{ promotion_id: number }>('/v3/promotion/create', data)
  },

  // 更新广告
  updatePromotion(promotion_id: number, data: Partial<V3PromotionCreateParams>) {
    return request.put<void>(`/v3/promotion/${promotion_id}`, data)
  },

  // 更新广告状态
  updatePromotionStatus(data: { advertiser_id: number; promotion_ids: number[]; opt_status: string }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/status/update', data)
  },

  // 删除广告
  deletePromotion(data: { advertiser_id: number; promotion_ids: number[] }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/delete', data)
  },

  // 更新广告预算
  updatePromotionBudget(data: { advertiser_id: number; promotion_ids: number[]; budget: number }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/budget/update', data)
  },

  // 更新广告出价
  updatePromotionBid(data: { advertiser_id: number; promotion_ids: number[]; bid: number }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/bid/update', data)
  },

  // 更新广告深度出价
  updatePromotionDeepBid(data: { advertiser_id: number; promotion_ids: number[]; deep_cpabid?: number; roi_goal?: number }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/deep_bid/update', data)
  },

  // 更新广告投放时段
  updatePromotionScheduleTime(data: { advertiser_id: number; promotion_ids: number[]; schedule_time: string }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/promotion/schedule_time/update', data)
  },

  // 更新广告素材状态
  updatePromotionMaterialStatus(data: { advertiser_id: number; promotion_id: number; material_ids: string[]; status: string }) {
    return request.post<{ success_ids: string[]; fail_ids: string[] }>('/v3/promotion/material/status/update', data)
  },

  // 获取广告审核建议
  getPromotionRejectReason(params: { advertiser_id: number; promotion_ids: number[] }) {
    return request.get<PromotionRejectReason[]>('/v3/promotion/reject_reason', params)
  },

  // 获取广告成本保障状态
  getPromotionCostProtectStatus(params: { advertiser_id: number; promotion_ids: number[] }) {
    return request.get<CostProtectStatus[]>('/v3/promotion/cost_protect_status', params)
  },

  // ========== 预算组管理 ==========

  // 获取预算组列表
  getBudgetGroupList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<V3BudgetGroup>>('/v3/budget_group/list', params)
  },

  // 创建预算组
  createBudgetGroup(data: V3BudgetGroupCreateParams) {
    return request.post<{ budget_group_id: number }>('/v3/budget_group/create', data)
  },

  // 更新预算组
  updateBudgetGroup(budget_group_id: number, data: Partial<V3BudgetGroupCreateParams>) {
    return request.put<void>(`/v3/budget_group/${budget_group_id}`, data)
  },

  // 删除预算组
  deleteBudgetGroup(data: { advertiser_id: number; budget_group_ids: number[] }) {
    return request.post<{ success_ids: number[]; fail_ids: number[] }>('/v3/budget_group/delete', data)
  },

  // ========== 数据报表 ==========

  // 获取项目报表
  getProjectReport(params: { advertiser_id: number; project_ids?: number[]; start_date: string; end_date: string; group_by?: string[] }) {
    return request.get<V3ReportData[]>('/v3/report/project', params)
  },

  // 获取广告报表
  getPromotionReport(params: { advertiser_id: number; promotion_ids?: number[]; start_date: string; end_date: string; group_by?: string[] }) {
    return request.get<V3ReportData[]>('/v3/report/promotion', params)
  },

  // 获取素材报表
  getMaterialReport(params: { advertiser_id: number; promotion_id?: number; start_date: string; end_date: string }) {
    return request.get<V3ReportData[]>('/v3/report/material', params)
  },

  // 获取自定义报表
  getCustomReport(params: { advertiser_id: number; dimensions: string[]; metrics: string[]; start_date: string; end_date: string; filtering?: object }) {
    return request.get<V3ReportData[]>('/v3/report/custom', params)
  },

  // 获取自定义报表可用指标和维度
  getCustomReportConfig(advertiser_id: number) {
    return request.get<CustomReportConfig[]>('/v3/report/custom/config', { advertiser_id })
  },

  // ========== 异步报表任务 ==========

  // 创建异步报表任务
  createAsyncReportTask(data: AsyncReportTaskCreateParams) {
    return request.post<{ task_id: string }>('/v3/report/async/create', data)
  },

  // 获取异步报表任务列表
  getAsyncReportTaskList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<AsyncReportTask>>('/v3/report/async/list', params)
  },

  // 下载异步报表结果
  downloadAsyncReportResult(task_id: string, advertiser_id: number) {
    return request.get<{ download_url: string }>(`/v3/report/async/download/${task_id}`, { advertiser_id })
  },

  // ========== 白盒策略配置 ==========

  // 创建/修改白盒配置
  saveAutoGenerateConfig(data: AutoGenerateConfigParams) {
    return request.post<{ config_id: number }>('/v3/promotion/auto_generate_config', data)
  },

  // 获取白盒配置详情
  getAutoGenerateConfig(params: { advertiser_id: number; promotion_id: number }) {
    return request.get<AutoGenerateConfig>('/v3/promotion/auto_generate_config', params)
  },

  // 获取白盒策略模板列表
  getStrategyList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<StrategyTemplate>>('/v3/creative/strategy/list', params)
  },

  // ========== 关联云图 ==========

  // 获取关联云图的广告主账户信息
  getCdpBrandInfo(advertiser_id: number) {
    return request.get<CdpBrandInfo[]>('/v3/cdp/brand', { advertiser_id })
  }
}

// ==================== 其他类型定义 ====================

export interface PromotionRejectReason {
  promotion_id: number
  reject_reason: string
  reject_item: string
  suggest: string
}

export interface CostProtectStatus {
  promotion_id: number
  cost_protect_status: number
  cost_protect_type: string
}

export interface CustomReportConfig {
  name: string
  type: string
  description: string
}

export interface AsyncReportTaskCreateParams {
  advertiser_id: number
  task_name: string
  dimensions: string[]
  metrics: string[]
  start_date: string
  end_date: string
  filtering?: object
}

export interface AsyncReportTask {
  task_id: string
  task_name: string
  status: string
  create_time: string
  complete_time?: string
}

export interface AutoGenerateConfigParams {
  advertiser_id: number
  promotion_id: number
  auto_generate_type: string
  config: object
}

export interface AutoGenerateConfig {
  config_id: number
  promotion_id: number
  auto_generate_type: string
  config: object
  status: string
}

export interface StrategyTemplate {
  strategy_id: number
  strategy_name: string
  strategy_type: string
  description: string
  status: string
}

export interface CdpBrandInfo {
  brand_id: string
  brand_name: string
  status: string
}
