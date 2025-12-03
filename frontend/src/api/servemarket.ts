import request, { PageResponse } from './request'

// ==================== 订单相关 ====================

export interface ServeMarketOrder {
  order_id: string
  advertiser_id: number
  service_id: string
  service_name: string
  order_amount: number
  order_status: string
  start_time: string
  end_time: string
  create_time: string
}

// ==================== 功能相关 ====================

export interface ServeMarketFunc {
  func_id: string
  func_name: string
  func_type: string
  status: string
  expire_time: string
  purchase_time: string
}

// ==================== 投前分析 ====================

export interface ServeMarketQuality {
  quality_score: number
  creative_score: number
  target_score: number
  budget_score: number
  suggestions: string[]
}

// ==================== RDS 订阅 ====================

export interface RDSSubscription {
  subscription_id: string
  advertiser_id: number
  data_type: string
  callback_url: string
  status: string
  create_time: string
}

// ==================== 仪表盘数据 ====================

export interface ServeMarketDashboard {
  total_orders: number
  active_services: number
  expire_soon: number
  total_spend: number
}

// ==================== API 方法 ====================

export const serveMarketApi = {
  // 订单管理
  getOrderList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<ServeMarketOrder>>('/servemarket/orders', params)
  },

  getOrderDetail(order_id: string) {
    return request.get<ServeMarketOrder>(`/servemarket/orders/${order_id}`)
  },

  // 已购功能
  getFuncList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<ServeMarketFunc>>('/servemarket/funcs', params)
  },

  getFuncDetail(func_id: string) {
    return request.get<ServeMarketFunc>(`/servemarket/funcs/${func_id}`)
  },

  // 投前分析
  getQualityAnalysis(params: { advertiser_id: number; ad_id?: number }) {
    return request.get<ServeMarketQuality>('/servemarket/quality', params)
  },

  // RDS 订阅管理
  getSubscriptionList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<RDSSubscription>>('/servemarket/subscriptions', params)
  },

  createSubscription(data: { advertiser_id: number; data_type: string; callback_url: string }) {
    return request.post<{ subscription_id: string }>('/servemarket/subscriptions', data)
  },

  updateSubscription(subscription_id: string, data: { callback_url?: string; status?: string }) {
    return request.put<void>(`/servemarket/subscriptions/${subscription_id}`, data)
  },

  deleteSubscription(subscription_id: string) {
    return request.delete<void>(`/servemarket/subscriptions/${subscription_id}`)
  },

  // 仪表盘
  getDashboard(advertiser_id: number) {
    return request.get<ServeMarketDashboard>('/servemarket/dashboard', { advertiser_id })
  }
}
