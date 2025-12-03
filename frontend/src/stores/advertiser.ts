import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Advertiser, AdvertiserListParams } from '@/types'

// Mock data for demonstration
const mockAdvertisers: Advertiser[] = [
  { id: '1234567890', name: '美妆旗舰店', company: '上海美丽科技有限公司', balance: 125000, status: 'enabled', createTime: '2024-01-15' },
  { id: '1234567891', name: '数码专营店', company: '北京智能设备公司', balance: 89500, status: 'enabled', createTime: '2024-01-14' },
  { id: '1234567892', name: '服饰品牌店', company: '杭州时尚服饰集团', balance: 156800, status: 'disabled', createTime: '2024-01-13' },
  { id: '1234567893', name: '食品旗舰店', company: '广州美食有限公司', balance: 98200, status: 'enabled', createTime: '2024-01-12' },
  { id: '1234567894', name: '母婴用品店', company: '深圳亲子科技公司', balance: 112500, status: 'enabled', createTime: '2024-01-11' },
]

export const useAdvertiserStore = defineStore('advertiser', () => {
  const advertisers = ref<Advertiser[]>([])
  const loading = ref(false)
  const total = ref(0)
  const currentPage = ref(1)
  const pageSize = ref(10)
  const currentAdvertiserId = ref<number>(0)

  const totalBalance = computed(() => {
    return advertisers.value.reduce((sum, a) => sum + a.balance, 0)
  })

  const enabledCount = computed(() => {
    return advertisers.value.filter(a => a.status === 'enabled').length
  })

  const setCurrentAdvertiser = (id: number) => {
    currentAdvertiserId.value = id
  }

  const fetchAdvertisers = async (params?: AdvertiserListParams) => {
    loading.value = true
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 500))
      
      let filtered = [...mockAdvertisers]
      if (params?.status && params.status !== 'all') {
        filtered = filtered.filter(a => a.status === params.status)
      }
      if (params?.keyword) {
        const kw = params.keyword.toLowerCase()
        filtered = filtered.filter(a => 
          a.name.toLowerCase().includes(kw) || 
          a.company.toLowerCase().includes(kw)
        )
      }
      
      advertisers.value = filtered
      total.value = filtered.length
      currentPage.value = params?.page || 1
    } finally {
      loading.value = false
    }
  }

  const getAdvertiserById = (id: string) => {
    return advertisers.value.find(a => a.id === id) || mockAdvertisers.find(a => a.id === id)
  }

  return {
    advertisers,
    loading,
    total,
    currentPage,
    pageSize,
    currentAdvertiserId,
    totalBalance,
    enabledCount,
    fetchAdvertisers,
    getAdvertiserById,
    setCurrentAdvertiser
  }
})
