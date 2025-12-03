import request, { PageResponse } from './request'

// ==================== 数据源类型 ====================

export interface DataSource {
  data_source_id: string
  name: string
  description: string
  source_type: string
  status: number
  cover_num: number
  create_time: string
  update_time: string
}

export interface DataSourceCreateParams {
  advertiser_id: number
  name: string
  description?: string
  data_format: string
  file_paths: string[]
}

// ==================== 人群包类型 ====================

export interface CustomAudience {
  audience_id: number
  advertiser_id: number
  name: string
  description: string
  audience_type: string
  source: string
  status: number
  cover_num: number
  is_valid: boolean
  expired_time?: string
  create_time: string
  update_time: string
}

export interface CustomAudienceDetail extends CustomAudience {
  data_source_ids: string[]
  rules?: AudienceRule[]
}

export interface AudienceRule {
  rule_type: string
  operator: string
  value: string[]
}

export interface CustomAudienceCreateParams {
  advertiser_id: number
  name: string
  description?: string
  data_source_id?: string
  custom_audience_type: string
}

export interface AudiencePushParams {
  advertiser_id: number
  audience_id: number
  target_advertiser_ids: number[]
}

// ==================== 云图相关类型 ====================

export interface BrandInfo {
  brand_id: string
  brand_name: string
  status: number
}

export interface AudienceCopyParams {
  advertiser_id: number
  audience_ids: number[]
  brand_id: string
}

// ==================== API 方法 ====================

export const dmpApi = {
  // ========== 数据源管理 ==========
  
  // 上传数据源文件
  uploadDataSourceFile(data: FormData) {
    return request.post<{ file_path: string }>('/dmp/datasource/file/upload', data)
  },

  // 创建数据源
  createDataSource(data: DataSourceCreateParams) {
    return request.post<{ data_source_id: string }>('/dmp/datasource/create', data)
  },

  // 更新数据源
  updateDataSource(data: { advertiser_id: number; data_source_id: string; name?: string; description?: string }) {
    return request.put<void>('/dmp/datasource/update', data)
  },

  // 获取数据源详情
  getDataSourceDetail(params: { advertiser_id: number; data_source_ids: string[] }) {
    return request.get<DataSource[]>('/dmp/datasource/read', params)
  },

  // ========== 人群包管理 ==========

  // 获取人群包列表
  getCustomAudienceList(params: { advertiser_id: number; page: number; page_size: number; status?: number }) {
    return request.get<PageResponse<CustomAudience>>('/dmp/customaudience/select', params)
  },

  // 获取人群包详情
  getCustomAudienceDetail(params: { advertiser_id: number; audience_ids: number[] }) {
    return request.get<CustomAudienceDetail[]>('/dmp/customaudience/read', params)
  },

  // 发布人群包
  publishCustomAudience(data: { advertiser_id: number; audience_id: number }) {
    return request.post<void>('/dmp/customaudience/publish', data)
  },

  // 推送人群包
  pushCustomAudience(data: AudiencePushParams) {
    return request.post<void>('/dmp/customaudience/push', data)
  },

  // 删除人群包
  deleteCustomAudience(data: { advertiser_id: number; audience_id: number }) {
    return request.delete<void>('/dmp/customaudience/delete', data)
  },

  // ========== 云图相关 ==========

  // 获取广告账户关联云图账户信息
  getBrandInfo(advertiser_id: number) {
    return request.get<BrandInfo[]>('/dmp/brand/get', { advertiser_id })
  },

  // 推送DMP人群包到云图账户
  copyAudienceToBrand(data: AudienceCopyParams) {
    return request.post<void>('/dmp/customaudience/copy', data)
  },

  // ========== 定向包管理 ==========

  // 获取定向包列表
  getOrientationPackageList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<OrientationPackage>>('/dmp/orientation/package/get', params)
  },

  // 创建定向包
  createOrientationPackage(data: OrientationPackageCreateParams) {
    return request.post<{ package_id: number }>('/dmp/orientation/package/create', data)
  },

  // 更新定向包
  updateOrientationPackage(package_id: number, data: Partial<OrientationPackageCreateParams>) {
    return request.put<void>(`/dmp/orientation/package/${package_id}`, data)
  },

  // 删除定向包
  deleteOrientationPackage(package_id: number, advertiser_id: number) {
    return request.delete<void>(`/dmp/orientation/package/${package_id}`, { advertiser_id })
  }
}

// ==================== 定向包类型 ====================

export interface OrientationPackage {
  package_id: number
  advertiser_id: number
  name: string
  description: string
  status: number
  landing_type: string[]
  audience_ids_include: number[]
  audience_ids_exclude: number[]
  gender: string
  age: string[]
  city: number[]
  location_type: string
  create_time: string
  update_time: string
}

export interface OrientationPackageCreateParams {
  advertiser_id: number
  name: string
  description?: string
  landing_type?: string[]
  audience_ids_include?: number[]
  audience_ids_exclude?: number[]
  gender?: string
  age?: string[]
  city?: number[]
  location_type?: string
}
