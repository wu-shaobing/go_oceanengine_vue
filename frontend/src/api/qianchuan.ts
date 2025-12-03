import request, { PageResponse } from './request'

// ==================== 账户相关 ====================

export interface QianchuanAccount {
  advertiser_id: number
  advertiser_name: string
  account_type: string
  company: string
  status: string
  balance: number
}

export interface QianchuanShop {
  shop_id: number
  shop_name: string
  shop_logo: string
  status: string
}

export interface AwemeAuth {
  aweme_id: string
  aweme_name: string
  aweme_avatar: string
  auth_status: string
  auth_time: string
}

// ==================== 广告相关 ====================

export interface QianchuanCampaign {
  campaign_id: number
  campaign_name: string
  budget: number
  budget_mode: string
  status: string
  marketing_goal: string
  created_at: string
}

export interface QianchuanAd {
  ad_id: number
  ad_name: string
  campaign_id: number
  status: string
  opt_status: string
  budget: number
  bid: number
  delivery_range: string
  created_at: string
}

export interface QianchuanCreative {
  creative_id: number
  ad_id: number
  creative_material_mode: string
  image_mode: string
  video_id: string
  image_ids: string[]
  title: string
  status: string
}

// ==================== 随心推 ====================

export interface AwemeOrder {
  order_id: string
  aweme_id: string
  aweme_name: string
  product_type: string
  budget: number
  status: string
  start_time: string
  end_time: string
  created_at: string
}

export interface AwemeOrderCreateParams {
  aweme_id: string
  product_type: string
  budget: number
  duration: number
  target_type: string
  target_config?: object
}

// ==================== 报表相关 ====================

export interface QianchuanReportData {
  stat_datetime: string
  cost: number
  show_cnt: number
  click_cnt: number
  ctr: number
  convert_cnt: number
  convert_cost: number
  convert_rate: number
  pay_order_count?: number
  pay_order_amount?: number
  roi?: number
}

// ==================== 素材相关 ====================

export interface QianchuanMaterial {
  material_id: string
  material_type: string
  width: number
  height: number
  url: string
  create_time: string
  signature: string
}

// ==================== 工具相关 ====================

export interface Industry {
  industry_id: number
  industry_name: string
  level: number
  parent_id: number
}

export interface DmpAudience {
  custom_audience_id: number
  name: string
  cover_num: number
  source: string
  status: string
  create_time: string
}

export interface Product {
  product_id: number
  product_name: string
  product_img: string
  market_price: number
  discount_price: number
  status: string
}

// ==================== 关键词相关 ====================

export interface Keyword {
  word: string
  match_type: string
  bid: number
  status: string
}

export interface InterestActionWord {
  id: number
  name: string
  parent_id?: number
  level: number
}

// ==================== API 方法 ====================

export const qianchuanApi = {
  // 账户
  getAccountInfo(advertiser_id: number) {
    return request.get<QianchuanAccount>('/qianchuan/account', { advertiser_id })
  },

  getShopList(advertiser_id: number) {
    return request.get<QianchuanShop[]>('/qianchuan/shops', { advertiser_id })
  },

  getAwemeAuthList(advertiser_id: number) {
    return request.get<AwemeAuth[]>('/qianchuan/aweme/auth', { advertiser_id })
  },

  getBalance(advertiser_id: number) {
    return request.get<{ balance: number; cash: number; grant: number }>('/qianchuan/balance', { advertiser_id })
  },

  // 广告系列
  getCampaignList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<QianchuanCampaign>>('/qianchuan/campaigns', params)
  },

  createCampaign(data: { advertiser_id: number; campaign_name: string; budget_mode: string; budget?: number; marketing_goal: string }) {
    return request.post<{ campaign_id: number }>('/qianchuan/campaigns', data)
  },

  updateCampaign(campaign_id: number, data: Partial<QianchuanCampaign>) {
    return request.put<void>(`/qianchuan/campaigns/${campaign_id}`, data)
  },

  // 广告计划
  getAdList(params: { advertiser_id: number; page: number; page_size: number; campaign_id?: number; status?: string }) {
    return request.get<PageResponse<QianchuanAd>>('/qianchuan/ads', params)
  },

  getAdDetail(ad_id: number) {
    return request.get<QianchuanAd>(`/qianchuan/ads/${ad_id}`)
  },

  createAd(data: object) {
    return request.post<{ ad_id: number }>('/qianchuan/ads', data)
  },

  updateAd(ad_id: number, data: object) {
    return request.put<void>(`/qianchuan/ads/${ad_id}`, data)
  },

  updateAdStatus(ad_ids: number[], opt_status: string) {
    return request.post<void>('/qianchuan/ads/status', { ad_ids, opt_status })
  },

  // 创意
  getCreativeList(params: { advertiser_id: number; ad_id?: number; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanCreative>>('/qianchuan/creatives', params)
  },

  // 全域推广
  getUniList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanAd>>('/qianchuan/uni/list', params)
  },

  createUni(data: object) {
    return request.post<{ ad_id: number }>('/qianchuan/uni', data)
  },

  getUniDetail(ad_id: number) {
    return request.get<QianchuanAd>(`/qianchuan/uni/${ad_id}`)
  },

  // 随心推
  getAwemeOrderList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<AwemeOrder>>('/qianchuan/aweme/orders', params)
  },

  getAwemeOrderDetail(order_id: string) {
    return request.get<AwemeOrder>(`/qianchuan/aweme/orders/${order_id}`)
  },

  createAwemeOrder(data: AwemeOrderCreateParams) {
    return request.post<{ order_id: string }>('/qianchuan/aweme/orders', data)
  },

  // 报表
  getAdvertiserReport(params: { advertiser_id: number; start_date: string; end_date: string }) {
    return request.get<QianchuanReportData[]>('/qianchuan/reports/advertiser', params)
  },

  getAdReport(params: { advertiser_id: number; ad_ids?: number[]; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanReportData & { ad_id: number; ad_name: string }>>('/qianchuan/reports/ad', params)
  },

  getCreativeReport(params: { advertiser_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanReportData & { creative_id: number }>>('/qianchuan/reports/creative', params)
  },

  getMaterialReport(params: { advertiser_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanReportData & { material_id: string }>>('/qianchuan/reports/material', params)
  },

  getKeywordReport(params: { advertiser_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanReportData & { keyword: string }>>('/qianchuan/reports/keyword', params)
  },

  getLiveReport(params: { advertiser_id: number; start_date: string; end_date: string }) {
    return request.get<QianchuanReportData[]>('/qianchuan/reports/live', params)
  },

  getRoomReport(params: { advertiser_id: number; room_id: string; start_date: string; end_date: string }) {
    return request.get<QianchuanReportData>('/qianchuan/reports/room', params)
  },

  getUniReport(params: { advertiser_id: number; start_date: string; end_date: string }) {
    return request.get<QianchuanReportData[]>('/qianchuan/reports/uni', params)
  },

  // 素材
  getMaterialList(params: { advertiser_id: number; material_type: string; page: number; page_size: number }) {
    return request.get<PageResponse<QianchuanMaterial>>('/qianchuan/materials', params)
  },

  uploadImage(advertiser_id: number, file: File) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    return request.post<{ image_id: string; url: string }>('/qianchuan/materials/image', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  uploadVideo(advertiser_id: number, file: File) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    return request.post<{ video_id: string; url: string }>('/qianchuan/materials/video', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  // 工具
  getIndustryList() {
    return request.get<Industry[]>('/qianchuan/tools/industries')
  },

  getDmpAudienceList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<DmpAudience>>('/qianchuan/tools/dmp', params)
  },

  getKeywordRecommend(params: { advertiser_id: number; seed_words: string[] }) {
    return request.get<{ keyword: string; heat: number }[]>('/qianchuan/tools/keyword/recommend', params)
  },

  getProductList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<Product>>('/qianchuan/products', params)
  },

  // 财务
  getBudget(advertiser_id: number) {
    return request.get<{ budget: number; budget_daily_used: number }>('/qianchuan/budget', { advertiser_id })
  },

  updateBudget(advertiser_id: number, budget: number) {
    return request.post<void>('/qianchuan/budget', { advertiser_id, budget })
  },

  getFinanceDetail(params: { advertiser_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<{
      transaction_seq: string
      transaction_type: string
      amount: number
      transaction_time: string
    }>>('/qianchuan/finance/detail', params)
  },

  // 关键词管理
  getKeywordList(params: { advertiser_id: number; ad_id: number }) {
    return request.get<{ list: Keyword[] }>('/qianchuan/keywords', params)
  },

  updateKeywords(data: { advertiser_id: number; ad_id: number; keywords: Keyword[] }) {
    return request.put<void>('/qianchuan/keywords', data)
  },

  getActionKeywords(params: { advertiser_id: number; query_word: string; action_scene?: string; action_days?: number }) {
    return request.get<{ list: InterestActionWord[] }>('/qianchuan/keywords/action', params)
  },

  getInterestKeywords(params: { advertiser_id: number; query_word: string }) {
    return request.get<{ list: InterestActionWord[] }>('/qianchuan/keywords/interest', params)
  },

  getKeywordSuggest(data: { advertiser_id: number; keywords: string[] }) {
    return request.post<{ list: InterestActionWord[] }>('/qianchuan/keywords/suggest', data)
  }
}
