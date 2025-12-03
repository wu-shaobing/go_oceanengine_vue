import request, { PageResponse } from './request'

export interface Campaign {
  id: number
  campaign_id: number
  advertiser_id: number
  name: string
  budget_mode: string
  budget: number
  landing_type: string
  status: string
  opt_status: string
  created_at: string
}

export interface CampaignCreateRequest {
  advertiser_id: number
  name: string
  budget_mode: 'BUDGET_MODE_INFINITE' | 'BUDGET_MODE_DAY'
  budget?: number
  landing_type: string
  marketing_goal?: string
}

export interface CampaignListParams {
  advertiser_id: number
  page: number
  page_size: number
  status?: string
  name?: string
}

export const campaignApi = {
  getList(params: CampaignListParams) {
    return request.get<PageResponse<Campaign>>('/campaign', params)
  },

  getDetail(id: number) {
    return request.get<Campaign>(`/campaign/${id}`)
  },

  create(data: CampaignCreateRequest) {
    return request.post<{ id: number }>('/campaign', data)
  },

  update(id: number, data: Partial<CampaignCreateRequest>) {
    return request.put<void>(`/campaign/${id}`, data)
  },

  updateStatus(id: number, status: 'enable' | 'disable' | 'delete') {
    return request.post<void>(`/campaign/${id}/status`, { status })
  },

  batchUpdateStatus(ids: number[], status: 'enable' | 'disable' | 'delete') {
    return request.post<void>('/campaign/batch/status', { ids, status })
  }
}
