export * from './advertiser'

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

export interface PaginationParams {
  page: number
  pageSize: number
}

export interface PaginatedResponse<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

export interface Campaign {
  id: string
  name: string
  advertiserId: string
  status: 'enabled' | 'disabled' | 'pending'
  budget: number
  dailyBudget: number
  startTime: string
  endTime?: string
  createTime: string
}

export interface Ad {
  id: string
  name: string
  campaignId: string
  status: 'enabled' | 'disabled' | 'pending' | 'rejected'
  bid: number
  createTime: string
}

export interface Report {
  date: string
  impressions: number
  clicks: number
  cost: number
  conversions: number
  ctr: number
  cpc: number
  cpa: number
}

export interface DashboardStats {
  todayCost: number
  todayImpressions: number
  todayClicks: number
  todayConversions: number
  costTrend: number
  impressionsTrend: number
}
