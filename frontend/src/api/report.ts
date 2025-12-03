import request, { PageResponse } from './request'

export interface ReportData {
  stat_datetime: string
  cost: number
  show: number
  click: number
  ctr: number
  convert: number
  convert_cost: number
  convert_rate: number
}

export interface ReportParams {
  advertiser_id: number
  start_date: string
  end_date: string
  group_by?: string[]
  page?: number
  page_size?: number
}

export const reportApi = {
  getAdvertiserReport(params: ReportParams) {
    return request.get<ReportData[]>('/report/advertiser', params)
  },

  getCampaignReport(params: ReportParams & { campaign_ids?: number[] }) {
    return request.get<PageResponse<ReportData & { 
      campaign_id: number
      campaign_name: string 
    }>>('/report/campaign', params)
  },

  getRealtime(advertiser_id: number) {
    return request.get<{
      cost: number
      show: number
      click: number
      convert: number
      update_time: string
    }>('/report/realtime', { advertiser_id })
  },

  exportReport(params: ReportParams & { type: 'advertiser' | 'campaign' | 'ad' }) {
    return request.post<{ task_id: string }>('/report/export', params)
  },

  getExportResult(taskId: string) {
    return request.get<{ status: string; url?: string }>(`/report/export/${taskId}`)
  }
}
