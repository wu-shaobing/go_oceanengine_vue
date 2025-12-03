import request, { PageResponse } from './request'

export interface Advertiser {
  id: number
  advertiser_id: number
  name: string
  company: string
  status: string
  balance: number
  valid_balance: number
  created_at: string
  last_sync_at: string
}

export interface AdvertiserListParams {
  page: number
  page_size: number
  keyword?: string
  status?: string
}

export const advertiserApi = {
  getList(params: AdvertiserListParams) {
    return request.get<PageResponse<Advertiser>>('/advertiser', params)
  },

  getDetail(id: number) {
    return request.get<Advertiser>(`/advertiser/${id}`)
  },

  sync(id: number) {
    return request.post<void>(`/advertiser/${id}/sync`)
  },

  getBalance(id: number) {
    return request.get<{
      balance: number
      valid_balance: number
      cash_balance: number
    }>(`/advertiser/${id}/balance`)
  },

  getFundTransactions(id: number, params: { 
    start_date: string
    end_date: string
    page: number
    page_size: number 
  }) {
    return request.get<PageResponse<{
      transaction_seq: string
      transaction_type: string
      amount: number
      transaction_time: string
    }>>(`/advertiser/${id}/funds`, params)
  }
}
