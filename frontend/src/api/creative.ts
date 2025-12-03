import request, { PageResponse } from './request'

// ==================== 创意相关 ====================

export interface Creative {
  id: number
  creative_id: number
  ad_id: number
  advertiser_id: number
  creative_material_mode: string
  image_mode: string
  video_id?: string
  image_ids?: string[]
  title: string
  title_list?: string[]
  source: string
  status: string
  opt_status: string
  created_at: string
}

export interface CreativeCreateParams {
  ad_id: number
  creative_material_mode: string
  image_mode: string
  title: string
  title_list?: string[]
  video_id?: string
  image_ids?: string[]
  source?: string
  third_industry_id?: number
  action_bar_text?: string
  call_to_action?: string
}

export interface CreativeListParams {
  advertiser_id: number
  ad_id?: number
  page: number
  page_size: number
  status?: string
}

// ==================== 创意组件 ====================

export interface CreativeComponent {
  component_id: string
  component_type: string
  component_name: string
  content: object
  status: string
  create_time: string
}

// ==================== 创意预览 ====================

export interface CreativePreview {
  preview_url: string
  qr_code_url: string
  expire_time: string
}

// ==================== API 方法 ====================

export const creativeApi = {
  // 创意列表
  getList(params: CreativeListParams) {
    return request.get<PageResponse<Creative>>('/creatives', params)
  },

  // 创意详情
  getDetail(id: number) {
    return request.get<Creative>(`/creatives/${id}`)
  },

  // 创建创意
  create(data: CreativeCreateParams) {
    return request.post<{ id: number; creative_id: number }>('/creatives', data)
  },

  // 更新创意
  update(id: number, data: Partial<CreativeCreateParams>) {
    return request.put<void>(`/creatives/${id}`, data)
  },

  // 删除创意
  delete(id: number) {
    return request.delete<void>(`/creatives/${id}`)
  },

  // 批量更新状态
  updateStatus(ids: number[], opt_status: 'enable' | 'disable' | 'delete') {
    return request.put<void>('/creatives/status', { ids, opt_status })
  },

  // 创意预览
  getPreview(creative_id: number) {
    return request.get<CreativePreview>(`/creatives/${creative_id}/preview`)
  },

  // 创意组件
  getComponentList(params: { advertiser_id: number; component_type?: string; page: number; page_size: number }) {
    return request.get<PageResponse<CreativeComponent>>('/creatives/components', params)
  },

  createComponent(data: { advertiser_id: number; component_type: string; component_name: string; content: object }) {
    return request.post<{ component_id: string }>('/creatives/components', data)
  },

  updateComponent(component_id: string, data: { component_name?: string; content?: object }) {
    return request.put<void>(`/creatives/components/${component_id}`, data)
  },

  deleteComponent(component_id: string) {
    return request.delete<void>(`/creatives/components/${component_id}`)
  },

  // 批量创建创意
  batchCreate(data: { ad_id: number; creatives: CreativeCreateParams[] }) {
    return request.post<{ creative_ids: number[] }>('/creatives/batch', data)
  },

  // 获取创意模板
  getTemplates(params: { advertiser_id: number; industry_id?: number }) {
    return request.get<{ template_id: string; template_name: string; preview_url: string }[]>('/creatives/templates', params)
  }
}
