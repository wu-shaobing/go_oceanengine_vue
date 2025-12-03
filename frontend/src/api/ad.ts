import request, { PageResponse } from './request'

// ==================== 广告计划相关 ====================

export interface Ad {
  id: number
  ad_id: number
  campaign_id: number
  advertiser_id: number
  name: string
  status: string
  opt_status: string
  pricing: string
  bid: number
  budget: number
  budget_mode: string
  delivery_range: string
  schedule_type: string
  start_time?: string
  end_time?: string
  created_at: string
  updated_at: string
}

export interface AdCreateParams {
  campaign_id: number
  name: string
  budget_mode: string
  budget?: number
  bid: number
  pricing: string
  delivery_range: string
  schedule_type: string
  start_time?: string
  end_time?: string
  audience?: AdAudience
  smart_bid_type?: string
  flow_control_mode?: string
}

export interface AdAudience {
  gender?: string
  age?: string[]
  city?: number[]
  location_type?: string
  interest_categories?: number[]
  interest_action_mode?: string
  action_categories?: number[]
  ac?: string[]
  platform?: string[]
  device_brand?: string[]
  device_price?: string[]
  launch_price?: number[]
  retargeting_tags_include?: number[]
  retargeting_tags_exclude?: number[]
  hide_if_exists?: number
  hide_if_converted?: string
}

export interface AdListParams {
  advertiser_id: number
  campaign_id?: number
  page: number
  page_size: number
  status?: string
  name?: string
}

// ==================== API 方法 ====================

export const adApi = {
  // 广告计划列表
  getList(params: AdListParams) {
    return request.get<PageResponse<Ad>>('/ads', params)
  },

  // 广告计划详情
  getDetail(id: number) {
    return request.get<Ad>(`/ads/${id}`)
  },

  // 创建广告计划
  create(data: AdCreateParams) {
    return request.post<{ id: number; ad_id: number }>('/ads', data)
  },

  // 更新广告计划
  update(id: number, data: Partial<AdCreateParams>) {
    return request.put<void>(`/ads/${id}`, data)
  },

  // 删除广告计划
  delete(id: number) {
    return request.delete<void>(`/ads/${id}`)
  },

  // 批量更新状态
  updateStatus(ids: number[], opt_status: 'enable' | 'disable' | 'delete') {
    return request.put<void>('/ads/status', { ids, opt_status })
  },

  // 批量更新预算
  updateBudget(data: { ad_id: number; budget: number }[]) {
    return request.put<void>('/ads/budget', { ads: data })
  },

  // 批量更新出价
  updateBid(data: { ad_id: number; bid: number }[]) {
    return request.put<void>('/ads/bid', { ads: data })
  },

  // 复制广告计划
  copy(ad_id: number, params?: { campaign_id?: number; copy_count?: number }) {
    return request.post<{ ad_ids: number[] }>(`/ads/${ad_id}/copy`, params)
  },

  // 获取定向模板
  getAudienceTemplates(advertiser_id: number) {
    return request.get<{ template_id: number; template_name: string; audience: AdAudience }[]>('/ads/audience/templates', { advertiser_id })
  },

  // 保存定向模板
  saveAudienceTemplate(data: { advertiser_id: number; template_name: string; audience: AdAudience }) {
    return request.post<{ template_id: number }>('/ads/audience/templates', data)
  },

  // 删除定向模板
  deleteAudienceTemplate(template_id: number) {
    return request.delete<void>(`/ads/audience/templates/${template_id}`)
  },

  // 获取预算建议
  getBudgetSuggestion(params: { advertiser_id: number; campaign_id?: number }) {
    return request.get<{ min_budget: number; suggested_budget: number; max_budget: number }>('/ads/budget/suggestion', params)
  },

  // 获取出价建议
  getBidSuggestion(params: { advertiser_id: number; pricing: string; audience?: AdAudience }) {
    return request.get<{ min_bid: number; suggested_bid: number; max_bid: number }>('/ads/bid/suggestion', params)
  }
}
