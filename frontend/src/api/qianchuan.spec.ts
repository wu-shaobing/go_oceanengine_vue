import { describe, it, expect, vi, beforeEach } from 'vitest'
import { qianchuanApi } from './qianchuan'
import request from './request'

// Mock request module
vi.mock('./request', () => ({
  default: {
    get: vi.fn(),
    post: vi.fn(),
    put: vi.fn(),
    delete: vi.fn()
  }
}))

describe('qianchuanApi', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  describe('账户相关', () => {
    it('getAccountInfo 应该调用正确的 API', async () => {
      const mockData = {
        advertiser_id: 12345,
        advertiser_name: '测试账户',
        account_type: 'QIANCHUAN'
      }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getAccountInfo(12345)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/account', { advertiser_id: 12345 })
      expect(result).toEqual(mockData)
    })

    it('getShopList 应该调用正确的 API', async () => {
      const mockData = [{ shop_id: 1, shop_name: '测试店铺' }]
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getShopList(12345)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/shops', { advertiser_id: 12345 })
      expect(result).toEqual(mockData)
    })

    it('getAwemeAuthList 应该调用正确的 API', async () => {
      const mockData = [{ aweme_id: '123', aweme_name: '测试抖音号' }]
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getAwemeAuthList(12345)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/aweme/auth', { advertiser_id: 12345 })
      expect(result).toEqual(mockData)
    })

    it('getBalance 应该调用正确的 API', async () => {
      const mockData = { balance: 10000, cash: 8000, grant: 2000 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getBalance(12345)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/balance', { advertiser_id: 12345 })
      expect(result).toEqual(mockData)
    })
  })

  describe('广告系列相关', () => {
    it('getCampaignList 应该调用正确的 API 和参数', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, page: 1, page_size: 20 }
      const result = await qianchuanApi.getCampaignList(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/campaigns', params)
      expect(result).toEqual(mockData)
    })

    it('createCampaign 应该调用正确的 API', async () => {
      const mockData = { campaign_id: 1 }
      vi.mocked(request.post).mockResolvedValue(mockData)

      const data = {
        advertiser_id: 12345,
        campaign_name: '测试系列',
        budget_mode: 'BUDGET_MODE_DAY',
        marketing_goal: 'VIDEO_PROM_GOODS'
      }
      const result = await qianchuanApi.createCampaign(data)

      expect(request.post).toHaveBeenCalledWith('/qianchuan/campaigns', data)
      expect(result).toEqual(mockData)
    })
  })

  describe('广告计划相关', () => {
    it('getAdList 应该调用正确的 API', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, page: 1, page_size: 20 }
      const result = await qianchuanApi.getAdList(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/ads', params)
      expect(result).toEqual(mockData)
    })

    it('getAdDetail 应该调用正确的 API', async () => {
      const mockData = { ad_id: 1, ad_name: '测试计划' }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getAdDetail(1)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/ads/1')
      expect(result).toEqual(mockData)
    })

    it('createAd 应该调用正确的 API', async () => {
      const mockData = { ad_id: 1 }
      vi.mocked(request.post).mockResolvedValue(mockData)

      const data = { campaign_id: 1, ad_name: '测试计划' }
      const result = await qianchuanApi.createAd(data)

      expect(request.post).toHaveBeenCalledWith('/qianchuan/ads', data)
      expect(result).toEqual(mockData)
    })

    it('updateAdStatus 应该调用正确的 API', async () => {
      vi.mocked(request.post).mockResolvedValue(undefined)

      await qianchuanApi.updateAdStatus([1, 2, 3], 'enable')

      expect(request.post).toHaveBeenCalledWith('/qianchuan/ads/status', {
        ad_ids: [1, 2, 3],
        opt_status: 'enable'
      })
    })
  })

  describe('随心推相关', () => {
    it('getAwemeOrderList 应该调用正确的 API', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, page: 1, page_size: 20 }
      const result = await qianchuanApi.getAwemeOrderList(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/aweme/orders', params)
      expect(result).toEqual(mockData)
    })

    it('getAwemeOrderDetail 应该调用正确的 API', async () => {
      const mockData = { order_id: 'order123', status: 'success' }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getAwemeOrderDetail('order123')

      expect(request.get).toHaveBeenCalledWith('/qianchuan/aweme/orders/order123')
      expect(result).toEqual(mockData)
    })
  })

  describe('报表相关', () => {
    it('getAdvertiserReport 应该调用正确的 API', async () => {
      const mockData = [{ stat_datetime: '2024-01-01', cost: 100 }]
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, start_date: '2024-01-01', end_date: '2024-01-31' }
      const result = await qianchuanApi.getAdvertiserReport(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/reports/advertiser', params)
      expect(result).toEqual(mockData)
    })

    it('getAdReport 应该调用正确的 API', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = {
        advertiser_id: 12345,
        start_date: '2024-01-01',
        end_date: '2024-01-31',
        page: 1,
        page_size: 20
      }
      const result = await qianchuanApi.getAdReport(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/reports/ad', params)
      expect(result).toEqual(mockData)
    })
  })

  describe('工具相关', () => {
    it('getIndustryList 应该调用正确的 API', async () => {
      const mockData = [{ industry_id: 1, industry_name: '电商' }]
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getIndustryList()

      expect(request.get).toHaveBeenCalledWith('/qianchuan/tools/industries')
      expect(result).toEqual(mockData)
    })

    it('getDmpAudienceList 应该调用正确的 API', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, page: 1, page_size: 20 }
      const result = await qianchuanApi.getDmpAudienceList(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/tools/dmp', params)
      expect(result).toEqual(mockData)
    })

    it('getProductList 应该调用正确的 API', async () => {
      const mockData = { list: [], total: 0, page: 1, page_size: 20 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const params = { advertiser_id: 12345, page: 1, page_size: 20 }
      const result = await qianchuanApi.getProductList(params)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/products', params)
      expect(result).toEqual(mockData)
    })
  })

  describe('财务相关', () => {
    it('getBudget 应该调用正确的 API', async () => {
      const mockData = { budget: 10000, budget_daily_used: 5000 }
      vi.mocked(request.get).mockResolvedValue(mockData)

      const result = await qianchuanApi.getBudget(12345)

      expect(request.get).toHaveBeenCalledWith('/qianchuan/budget', { advertiser_id: 12345 })
      expect(result).toEqual(mockData)
    })

    it('updateBudget 应该调用正确的 API', async () => {
      vi.mocked(request.post).mockResolvedValue(undefined)

      await qianchuanApi.updateBudget(12345, 20000)

      expect(request.post).toHaveBeenCalledWith('/qianchuan/budget', {
        advertiser_id: 12345,
        budget: 20000
      })
    })
  })
})
