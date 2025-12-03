import request, { PageResponse } from './request'

// ==================== 账号相关 ====================

export interface StarAccount {
  star_id: number
  star_name: string
  avatar_url: string
  account_type: string
  status: string
  balance: number
}

export interface AgentAdvertiser {
  advertiser_id: number
  advertiser_name: string
  company: string
  status: string
  balance: number
}

// ==================== 资金相关 ====================

export interface FundBalance {
  star_id: number
  balance: number
  cash_balance: number
  grant_balance: number
}

export interface FundDaily {
  date: string
  income: number
  expense: number
  balance: number
}

export interface FundTransaction {
  transaction_seq: string
  transaction_type: string
  amount: number
  balance_after: number
  transaction_time: string
  remark: string
}

// ==================== 任务相关 ====================

export interface StarTask {
  task_id: string
  task_name: string
  advertiser_id: number
  advertiser_name: string
  status: string
  budget: number
  start_time: string
  end_time: string
  create_time: string
}

export interface StarTaskDetail extends StarTask {
  requirement: string
  target_audience: string
  content_type: string
  video_count: number
  completed_count: number
}

export interface StarTaskItem {
  item_id: string
  task_id: string
  author_id: string
  author_name: string
  author_avatar: string
  title: string
  cover_url: string
  video_url: string
  status: string
  publish_time: string
  play_count: number
  digg_count: number
  comment_count: number
  share_count: number
}

// ==================== 需求相关 ====================

export interface StarDemand {
  demand_id: string
  demand_name: string
  advertiser_id: number
  content_type: string
  budget: number
  status: string
  create_time: string
}

export interface StarDemandOrder {
  order_id: string
  demand_id: string
  author_id: string
  author_name: string
  order_amount: number
  status: string
  create_time: string
}

// ==================== 报表相关 ====================

export interface StarReportOverview {
  date: string
  cost: number
  play_count: number
  digg_count: number
  comment_count: number
  share_count: number
  convert_count: number
}

export interface StarReportAudience {
  gender: { male: number; female: number }
  age: { name: string; value: number }[]
  province: { name: string; value: number }[]
}

// ==================== 线索相关 ====================

export interface StarClue {
  clue_id: string
  task_id: string
  item_id: string
  clue_type: string
  name: string
  phone: string
  create_time: string
  follow_status: string
}

// ==================== API 方法 ====================

export const starApi = {
  // 账号信息
  getAccountInfo(star_id: number) {
    return request.get<StarAccount>('/star/account', { star_id })
  },

  getAgentAdvertisers(params: { star_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<AgentAdvertiser>>('/star/agent/advertisers', params)
  },

  // 资金管理
  getBatchBalance(star_ids: number[]) {
    return request.post<FundBalance[]>('/star/fund/balance', { star_ids })
  },

  getFundDaily(params: { star_id: number; start_date: string; end_date: string }) {
    return request.get<FundDaily[]>('/star/fund/daily', params)
  },

  getFundTransactions(params: { star_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<FundTransaction>>('/star/fund/transactions', params)
  },

  // 任务管理
  getTaskList(params: { star_id: number; status?: string; page: number; page_size: number }) {
    return request.get<PageResponse<StarTask>>('/star/tasks', params)
  },

  getTaskDetail(task_id: string) {
    return request.get<StarTaskDetail>(`/star/tasks/${task_id}`)
  },

  getTaskItems(params: { task_id: string; page: number; page_size: number }) {
    return request.get<PageResponse<StarTaskItem>>(`/star/tasks/${params.task_id}/items`, params)
  },

  updateTaskStatus(task_id: string, status: string) {
    return request.put<void>(`/star/tasks/${task_id}/status`, { status })
  },

  // 需求管理
  getDemandList(params: { star_id: number; status?: string; page: number; page_size: number }) {
    return request.get<PageResponse<StarDemand>>('/star/demands', params)
  },

  getDemandDetail(demand_id: string) {
    return request.get<StarDemand>(`/star/demands/${demand_id}`)
  },

  getDemandOrders(params: { demand_id: string; page: number; page_size: number }) {
    return request.get<PageResponse<StarDemandOrder>>(`/star/demands/${params.demand_id}/orders`, params)
  },

  // 报表
  getReportOverview(params: { star_id: number; start_date: string; end_date: string }) {
    return request.get<StarReportOverview[]>('/star/reports/overview', params)
  },

  getReportAudience(params: { star_id: number; task_id?: string }) {
    return request.get<StarReportAudience>('/star/reports/audience', params)
  },

  getReportDaily(params: { star_id: number; task_id?: string; start_date: string; end_date: string }) {
    return request.get<StarReportOverview[]>('/star/reports/daily', params)
  },

  // 线索管理
  getClueList(params: { star_id: number; task_id?: string; page: number; page_size: number; follow_status?: string }) {
    return request.get<PageResponse<StarClue>>('/star/clues', params)
  },

  updateClueStatus(clue_id: string, follow_status: string) {
    return request.put<void>(`/star/clues/${clue_id}`, { follow_status })
  },

  exportClues(params: { advertiser_id: number; task_id: number; start_date?: string; end_date?: string }) {
    return request.post<{ data: StarClue[]; total: number; message: string }>('/star/clues/export', params)
  }
}
