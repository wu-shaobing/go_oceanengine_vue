export interface Advertiser {
  id: string
  name: string
  company: string
  balance: number
  status: 'enabled' | 'disabled'
  email?: string
  phone?: string
  address?: string
  createTime: string
  updateTime?: string
}

export interface AdvertiserListParams {
  page?: number
  pageSize?: number
  status?: string
  keyword?: string
}

export interface AdvertiserListResponse {
  list: Advertiser[]
  total: number
  page: number
  pageSize: number
}
