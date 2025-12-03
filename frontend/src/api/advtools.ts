import request, { PageResponse } from './request'

// ==================== RTA策略类型 ====================

export interface RtaInfo {
  rta_id: number
  advertiser_id: number
  rta_name: string
  status: string
  rta_type: string
  strategy_type: string
  create_time: string
  update_time: string
}

export interface AvailableRta {
  rta_id: number
  rta_name: string
  rta_type: string
  description: string
}

export interface RtaScope {
  rta_id: number
  scope_type: string
  scope_value: string[]
}

// ==================== 一键起量类型 ====================

export interface AdRaiseParams {
  advertiser_id: number
  ad_id: number
  budget: number
  duration: number
  raise_type?: string
}

export interface AdRaiseEstimate {
  ad_id: number
  estimate_cost: number
  estimate_show: number
  estimate_click: number
  estimate_convert: number
  suggest_budget: number
}

export interface AdRaiseStatus {
  ad_id: number
  raise_status: string
  budget: number
  spend: number
  start_time: string
  end_time: string
  current_progress: number
}

export interface AdRaiseResult {
  ad_id: number
  total_cost: number
  total_show: number
  total_click: number
  total_convert: number
  raise_status: string
  complete_time: string
}

export interface SuggestBudget {
  ad_id: number
  suggest_budget: number
  min_budget: number
  max_budget: number
  suggest_duration: number
}

// ==================== 定向包类型 ====================

export interface AudiencePackage {
  audience_package_id: number
  advertiser_id: number
  name: string
  description: string
  audience_type: string
  status: string
  landing_type: string[]
  retargeting_tags_include: number[]
  retargeting_tags_exclude: number[]
  gender: string
  age: string[]
  city: number[]
  location_type: string
  interest_categories: number[]
  action_categories: number[]
  create_time: string
  update_time: string
}

export interface AudiencePackageCreateParams {
  advertiser_id: number
  name: string
  description?: string
  landing_type?: string[]
  retargeting_tags_include?: number[]
  retargeting_tags_exclude?: number[]
  gender?: string
  age?: string[]
  city?: number[]
  location_type?: string
  interest_categories?: number[]
  action_categories?: number[]
}

// ==================== 原生锚点类型 ====================

export interface NativeAnchor {
  anchor_id: number
  advertiser_id: number
  anchor_name: string
  anchor_type: string
  anchor_status: string
  bind_count: number
  create_time: string
  update_time: string
}

export interface NativeAnchorDetail extends NativeAnchor {
  anchor_content: AnchorContent
  style_config: StyleConfig
}

export interface AnchorContent {
  title: string
  description?: string
  icon_url?: string
  action_text: string
  link_url?: string
  phone?: string
  app_info?: AppInfo
}

export interface AppInfo {
  app_name: string
  package_name: string
  download_url: string
}

export interface StyleConfig {
  background_color?: string
  text_color?: string
  button_color?: string
  position?: string
}

export interface NativeAnchorCreateParams {
  advertiser_id: number
  anchor_name: string
  anchor_type: string
  anchor_content: AnchorContent
  style_config?: StyleConfig
}

// ==================== 诊断工具类型 ====================

export interface DiagnosisSuggestion {
  suggestion_id: string
  ad_id: number
  suggestion_type: string
  title: string
  description: string
  impact_level: string
  current_value: string
  suggest_value: string
  expect_effect: string
  create_time: string
}

// ==================== 其他工具类型 ====================

export interface Quota {
  advertiser_id: number
  quota_type: string
  total_quota: number
  used_quota: number
  remain_quota: number
  reset_time?: string
}

export interface AdQuality {
  ad_id: number
  quality_score: number
  creative_score: number
  landing_page_score: number
  audience_score: number
  suggestions: string[]
  update_time: string
}

export interface AdStatExtraInfo {
  ad_id: number
  first_convert_time?: string
  last_convert_time?: string
  audience_analysis?: AudienceAnalysis
  convert_analysis?: ConvertAnalysis
}

export interface AudienceAnalysis {
  gender_distribution: Record<string, number>
  age_distribution: Record<string, number>
  city_distribution: Record<string, number>
}

export interface ConvertAnalysis {
  convert_path: ConvertPath[]
  avg_convert_time: number
}

export interface ConvertPath {
  step: number
  action: string
  count: number
  rate: number
}

// ==================== API 方法 ====================

export const advtoolsApi = {
  // ========== RTA策略管理 ==========
  
  // 获取RTA策略信息
  getRtaInfo(params: { advertiser_id: number; rta_id?: number }) {
    return request.get<RtaInfo[]>('/advtools/rta/info', params)
  },

  // 获取可用RTA策略
  getAvailableRta(advertiser_id: number) {
    return request.get<AvailableRta[]>('/advtools/rta/available', { advertiser_id })
  },

  // 更新RTA策略状态
  updateRtaStatus(data: { advertiser_id: number; rta_id: number; status: string }) {
    return request.put<void>('/advtools/rta/status', data)
  },

  // 设置RTA策略范围
  setRtaScope(data: { advertiser_id: number; rta_id: number; scope_type: string; scope_value: string[] }) {
    return request.put<void>('/advtools/rta/scope', data)
  },

  // 获取RTA策略范围
  getRtaScope(params: { advertiser_id: number; rta_id: number }) {
    return request.get<RtaScope>('/advtools/rta/scope', params)
  },

  // ========== 一键起量管理 ==========

  // 设置一键起量
  setAdRaise(data: AdRaiseParams) {
    return request.post<{ raise_id: number }>('/advtools/raise', data)
  },

  // 获取起量效果预估
  getAdRaiseEstimate(params: { advertiser_id: number; ad_id: number; budget: number }) {
    return request.get<AdRaiseEstimate>('/advtools/raise/estimate', params)
  },

  // 获取起量状态
  getAdRaiseStatus(params: { advertiser_id: number; ad_id: number }) {
    return request.get<AdRaiseStatus>('/advtools/raise/status', params)
  },

  // 获取起量结果
  getAdRaiseResult(params: { advertiser_id: number; ad_id: number }) {
    return request.get<AdRaiseResult>('/advtools/raise/result', params)
  },

  // 获取建议预算
  getSuggestBudget(params: { advertiser_id: number; ad_id: number }) {
    return request.get<SuggestBudget>('/advtools/raise/suggest-budget', params)
  },

  // ========== 定向包管理 ==========

  // 获取定向包列表
  getAudiencePackage(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<AudiencePackage>>('/advtools/audience', params)
  },

  // 创建定向包
  createAudiencePackage(data: AudiencePackageCreateParams) {
    return request.post<{ audience_package_id: number }>('/advtools/audience', data)
  },

  // 更新定向包
  updateAudiencePackage(package_id: number, data: Partial<AudiencePackageCreateParams>) {
    return request.put<void>(`/advtools/audience/${package_id}`, data)
  },

  // 删除定向包
  deleteAudiencePackage(package_id: number, advertiser_id: number) {
    return request.delete<void>(`/advtools/audience/${package_id}`, { advertiser_id })
  },

  // 绑定定向包到广告
  bindAudiencePackage(data: { advertiser_id: number; audience_package_id: number; ad_ids: number[] }) {
    return request.post<{ success_count: number; fail_list: number[] }>('/advtools/audience/bind', data)
  },

  // 解绑定向包
  unbindAudiencePackage(data: { advertiser_id: number; audience_package_id: number; ad_ids: number[] }) {
    return request.post<{ success_count: number; fail_list: number[] }>('/advtools/audience/unbind', data)
  },

  // ========== 原生锚点管理 ==========

  // 获取原生锚点列表
  getNativeAnchor(params: { advertiser_id: number; page: number; page_size: number; anchor_type?: string }) {
    return request.get<PageResponse<NativeAnchor>>('/advtools/anchor', params)
  },

  // 获取原生锚点详情
  getNativeAnchorDetail(anchor_id: number, advertiser_id: number) {
    return request.get<NativeAnchorDetail>(`/advtools/anchor/${anchor_id}`, { advertiser_id })
  },

  // 创建原生锚点
  createNativeAnchor(data: NativeAnchorCreateParams) {
    return request.post<{ anchor_id: number }>('/advtools/anchor', data)
  },

  // 更新原生锚点
  updateNativeAnchor(anchor_id: number, data: Partial<NativeAnchorCreateParams>) {
    return request.put<void>(`/advtools/anchor/${anchor_id}`, data)
  },

  // 删除原生锚点
  deleteNativeAnchor(anchor_id: number, advertiser_id: number) {
    return request.delete<void>(`/advtools/anchor/${anchor_id}`, { advertiser_id })
  },

  // ========== 诊断工具 ==========

  // 获取诊断建议
  getDiagnosisSuggestion(params: { advertiser_id: number; ad_id?: number; suggestion_type?: string }) {
    return request.get<DiagnosisSuggestion[]>('/advtools/diagnosis/suggestion', params)
  },

  // 接受诊断建议
  acceptDiagnosisSuggestion(data: { advertiser_id: number; suggestion_id: string }) {
    return request.post<void>('/advtools/diagnosis/accept', data)
  },

  // ========== 其他工具 ==========

  // 获取配额信息
  getQuota(params: { advertiser_id: number; quota_type?: string }) {
    return request.get<Quota[]>('/advtools/quota', params)
  },

  // 获取广告质量分
  getAdQuality(params: { advertiser_id: number; ad_id: number }) {
    return request.get<AdQuality>('/advtools/ad-quality', params)
  },

  // 获取广告扩展统计信息
  getAdStatExtraInfo(params: { advertiser_id: number; ad_id: number }) {
    return request.get<AdStatExtraInfo>('/advtools/ad-stat-extra', params)
  }
}
