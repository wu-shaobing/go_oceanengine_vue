import request, { PageResponse } from './request'

// ==================== 资产管理类型 ====================

export interface Asset {
  asset_id: number
  asset_name: string
  asset_type: string
  landing_type: string
  status: string
  create_time: string
  update_time: string
}

export interface AssetDetail extends Asset {
  advertiser_id: number
  url?: string
  app_id?: string
  app_name?: string
  package_name?: string
  mini_program_id?: string
}

export interface AssetsGetParams {
  advertiser_id: number
  asset_type?: string
  landing_type?: string
  page: number
  page_size: number
}

export interface AssetCreateParams {
  advertiser_id: number
  asset_name: string
  asset_type: string
  landing_type: string
  url?: string
  app_id?: string
  package_name?: string
  mini_program_id?: string
}

// ==================== 事件管理类型 ====================

export interface AvailableEvent {
  event_type: string
  event_name: string
  description: string
  category: string
}

export interface EventConfig {
  event_id: number
  event_type: string
  event_name: string
  status: string
  priority: number
  create_time: string
}

export interface EventCreateParams {
  advertiser_id: number
  asset_id: number
  events: {
    event_type: string
    event_name?: string
    priority?: number
  }[]
}

// ==================== 监测链接类型 ====================

export interface TrackURL {
  track_url_group_id: number
  track_url_group_name: string
  track_urls: TrackURLItem[]
  create_time: string
}

export interface TrackURLItem {
  action_type: string
  url: string
  url_type: string
}

export interface TrackURLCreateParams {
  advertiser_id: number
  asset_id: number
  track_url_group_name: string
  track_urls: TrackURLItem[]
}

export interface TrackURLUpdateParams extends TrackURLCreateParams {
  track_url_group_id: number
}

// ==================== 资产共享类型 ====================

export interface ShareInfo {
  share_type: string
  shared_advertiser_ids: number[]
  shared_agent_ids: number[]
}

export interface ShareParams {
  advertiser_id: number
  asset_id: number
  share_type: string
  target_advertiser_ids?: number[]
  target_agent_ids?: number[]
}

// ==================== 优化目标类型 ====================

export interface OptimizedGoal {
  external_action: string
  external_action_name: string
  deep_external_action?: string
  deep_external_action_name?: string
  deep_bid_type: string
  optimization_goal: string
}

export interface DeepBidType {
  deep_bid_type: string
  deep_bid_type_name: string
  description: string
}

// ==================== 转化回传类型 ====================

export interface ConversionParams {
  event_type: string
  context: {
    ad: {
      callback: string
    }
  }
  timestamp?: number
  properties?: Record<string, unknown>
}

// ==================== 鉴权管理类型 ====================

export interface PublicKey {
  key_id: number
  public_key: string
  algorithm: string
  status: string
  create_time: string
  expire_time: string
}

export interface AddPublicKeyParams {
  advertiser_id: number
  public_key: string
  algorithm?: string
}

// ==================== API 方法 ====================

export const eventmanagerApi = {
  // ========== 资产管理 ==========

  // 获取已创建资产列表
  getAssets(params: AssetsGetParams) {
    return request.get<PageResponse<Asset>>('/eventmanager/assets', params)
  },

  // 获取账户下资产列表
  getAllAssetsList(params: AssetsGetParams) {
    return request.get<PageResponse<Asset>>('/eventmanager/assets/all', params)
  },

  // 创建事件资产
  createAsset(data: AssetCreateParams) {
    return request.post<{ asset_id: number }>('/eventmanager/assets', data)
  },

  // ========== 事件管理 ==========

  // 获取可创建事件列表
  getAvailableEvents(params: { advertiser_id: number; asset_id: number }) {
    return request.get<AvailableEvent[]>('/eventmanager/events/available', params)
  },

  // 获取已创建事件列表
  getEventConfigs(params: { advertiser_id: number; asset_id: number }) {
    return request.get<EventConfig[]>('/eventmanager/events/configs', params)
  },

  // 资产下创建事件
  createEvents(data: EventCreateParams) {
    return request.post<void>('/eventmanager/events', data)
  },

  // ========== 监测链接管理 ==========

  // 获取监测链接组
  getTrackURL(params: { advertiser_id: number; asset_id: number }) {
    return request.get<TrackURL>('/eventmanager/track-url', params)
  },

  // 创建监测链接组
  createTrackURL(data: TrackURLCreateParams) {
    return request.post<void>('/eventmanager/track-url', data)
  },

  // 更新监测链接组
  updateTrackURL(data: TrackURLUpdateParams) {
    return request.put<void>('/eventmanager/track-url', data)
  },

  // ========== 资产共享 ==========

  // 查看共享范围
  getShare(params: { advertiser_id: number; asset_id: number }) {
    return request.get<ShareInfo>('/eventmanager/share', params)
  },

  // 资产共享
  share(data: ShareParams) {
    return request.post<{ fail_list: number[] }>('/eventmanager/share', data)
  },

  // 取消资产共享
  shareCancel(data: ShareParams) {
    return request.post<{ fail_list: number[] }>('/eventmanager/share/cancel', data)
  },

  // ========== 优化目标 ==========

  // 获取可用优化目标
  getOptimizedGoal(params: { advertiser_id: number; asset_id: number; landing_type?: string }) {
    return request.get<{ optimized_goals: OptimizedGoal[]; deep_bid_types: DeepBidType[] }>('/eventmanager/optimized-goal', params)
  },

  // ========== 转化回传 ==========

  // 转化回传
  conversion(data: ConversionParams) {
    return request.post<void>('/eventmanager/conversion', data)
  },

  // ========== 转化回传鉴权 ==========

  // 新增公钥
  addPublicKey(data: AddPublicKeyParams) {
    return request.post<PublicKey>('/eventmanager/auth/public-key', data)
  },

  // 查询全部公钥
  getAllPublicKeys(advertiser_id: number) {
    return request.get<PublicKey[]>('/eventmanager/auth/public-keys', { advertiser_id })
  },

  // 开启鉴权
  enableAuth(advertiser_id: number) {
    return request.post<void>('/eventmanager/auth/enable', { advertiser_id })
  },

  // 关闭鉴权
  disableAuth(advertiser_id: number) {
    return request.post<void>('/eventmanager/auth/disable', { advertiser_id })
  }
}
