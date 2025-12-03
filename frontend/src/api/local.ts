import request, { PageResponse } from './request'

// ==================== 项目相关 ====================

export interface LocalProject {
  project_id: number
  project_name: string
  advertiser_id: number
  landing_type: string
  budget: number
  budget_mode: string
  status: string
  opt_status: string
  created_at: string
}

export interface LocalProjectCreateParams {
  advertiser_id: number
  project_name: string
  landing_type: string
  budget_mode: string
  budget?: number
  delivery_range?: string
  audience?: object
}

// ==================== 广告相关 ====================

export interface LocalPromotion {
  promotion_id: number
  project_id: number
  promotion_name: string
  status: string
  opt_status: string
  budget: number
  bid: number
  schedule_type: string
  start_time?: string
  end_time?: string
  created_at: string
}

export interface LocalPromotionCreateParams {
  project_id: number
  promotion_name: string
  budget: number
  bid: number
  schedule_type: string
  start_time?: string
  end_time?: string
  audience?: object
  creative?: object
}

// ==================== 线索相关 ====================

export interface LocalClue {
  clue_id: string
  promotion_id: number
  clue_type: string
  name: string
  telephone: string
  address?: string
  create_time: string
  follow_status: string
  remark?: string
}

// ==================== 报表相关 ====================

export interface LocalReportData {
  stat_datetime: string
  cost: number
  show_cnt: number
  click_cnt: number
  ctr: number
  clue_count: number
  clue_cost: number
  call_count: number
  form_count: number
}

// ==================== 素材相关 ====================

export interface LocalMaterial {
  material_id: string
  material_type: string
  url: string
  width: number
  height: number
  duration?: number
  create_time: string
}

// ==================== 门店相关 ====================

export interface LocalStore {
  store_id: string
  store_name: string
  address: string
  latitude: number
  longitude: number
  phone: string
  business_hours: string
  status: string
}

// ==================== API 方法 ====================

export const localApi = {
  // 项目管理
  getProjectList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<LocalProject>>('/local/projects', params)
  },

  getProjectDetail(project_id: number) {
    return request.get<LocalProject>(`/local/projects/${project_id}`)
  },

  createProject(data: LocalProjectCreateParams) {
    return request.post<{ project_id: number }>('/local/projects', data)
  },

  updateProject(project_id: number, data: Partial<LocalProjectCreateParams>) {
    return request.put<void>(`/local/projects/${project_id}`, data)
  },

  updateProjectStatus(project_ids: number[], opt_status: string) {
    return request.post<void>('/local/projects/status', { project_ids, opt_status })
  },

  deleteProject(project_id: number) {
    return request.delete<void>(`/local/projects/${project_id}`)
  },

  // 广告管理
  getPromotionList(params: { advertiser_id: number; project_id?: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<LocalPromotion>>('/local/promotions', params)
  },

  getPromotionDetail(promotion_id: number) {
    return request.get<LocalPromotion>(`/local/promotions/${promotion_id}`)
  },

  createPromotion(data: LocalPromotionCreateParams) {
    return request.post<{ promotion_id: number }>('/local/promotions', data)
  },

  updatePromotion(promotion_id: number, data: Partial<LocalPromotionCreateParams>) {
    return request.put<void>(`/local/promotions/${promotion_id}`, data)
  },

  updatePromotionStatus(promotion_ids: number[], opt_status: string) {
    return request.post<void>('/local/promotions/status', { promotion_ids, opt_status })
  },

  deletePromotion(promotion_id: number) {
    return request.delete<void>(`/local/promotions/${promotion_id}`)
  },

  // 线索管理
  getClueList(params: { advertiser_id: number; promotion_id?: number; page: number; page_size: number; follow_status?: string }) {
    return request.get<PageResponse<LocalClue>>('/local/clues', params)
  },

  getClueDetail(clue_id: string) {
    return request.get<LocalClue>(`/local/clues/${clue_id}`)
  },

  updateClueStatus(clue_id: string, follow_status: string, remark?: string) {
    return request.put<void>(`/local/clues/${clue_id}`, { follow_status, remark })
  },

  exportClues(params: { advertiser_id: number; start_date: string; end_date: string }) {
    return request.post<{ task_id: string }>('/local/clues/export', params)
  },

  // 报表
  getProjectReport(params: { advertiser_id: number; project_ids?: number[]; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<LocalReportData & { project_id: number; project_name: string }>>('/local/reports/project', params)
  },

  getPromotionReport(params: { advertiser_id: number; promotion_ids?: number[]; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<LocalReportData & { promotion_id: number; promotion_name: string }>>('/local/reports/promotion', params)
  },

  getMaterialReport(params: { advertiser_id: number; start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<LocalReportData & { material_id: string }>>('/local/reports/material', params)
  },

  // 素材管理
  getMaterialList(params: { advertiser_id: number; material_type: string; page: number; page_size: number }) {
    return request.get<PageResponse<LocalMaterial>>('/local/materials', params)
  },

  uploadVideo(advertiser_id: number, file: File) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    return request.post<{ video_id: string; url: string }>('/local/materials/video', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  uploadImage(advertiser_id: number, file: File) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    return request.post<{ image_id: string; url: string }>('/local/materials/image', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  deleteMaterial(material_id: string) {
    return request.delete<void>(`/local/materials/${material_id}`)
  },

  // 门店管理
  getStoreList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<LocalStore>>('/local/stores', params)
  },

  getStoreDetail(store_id: string) {
    return request.get<LocalStore>(`/local/stores/${store_id}`)
  },

  createStore(data: Omit<LocalStore, 'store_id' | 'status'> & { advertiser_id: number }) {
    return request.post<{ store_id: string }>('/local/stores', data)
  },

  updateStore(store_id: string, data: Partial<LocalStore>) {
    return request.put<void>(`/local/stores/${store_id}`, data)
  },

  deleteStore(store_id: string) {
    return request.delete<void>(`/local/stores/${store_id}`)
  }
}
