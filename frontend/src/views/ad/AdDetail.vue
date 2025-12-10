<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import StatsCard from '@/components/business/StatsCard.vue'

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const adId = route.params.id as string

const ad = reactive({
  id: '',
  name: '',
  campaign: '',
  campaignId: '',
  status: 'active',
  budget: 0,
  bidType: 'cpc',
  bidAmount: 0,
  impressions: 0,
  clicks: 0,
  ctr: 0,
  spend: 0,
  conversions: 0,
  cpa: 0,
  startTime: '',
  endTime: '',
  createdAt: '',
  updatedAt: ''
})

const stats = reactive({
  impressions: 0,
  clicks: 0,
  ctr: 0,
  spend: 0,
  conversions: 0,
  cpa: 0
})

const fetchAdDetail = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // 模拟数据
  Object.assign(ad, {
    id: adId,
    name: `广告创意 ${adId}`,
    campaign: '品牌推广系列',
    campaignId: '1',
    status: 'active',
    budget: 5000,
    bidType: 'cpc',
    bidAmount: 2.5,
    impressions: 156789,
    clicks: 4532,
    ctr: 2.89,
    spend: 11330,
    conversions: 234,
    cpa: 48.42,
    startTime: '2025-01-01 00:00:00',
    endTime: '2025-12-31 23:59:59',
    createdAt: '2025-01-01 10:30:00',
    updatedAt: '2025-12-09 08:00:00'
  })
  
  Object.assign(stats, {
    impressions: ad.impressions,
    clicks: ad.clicks,
    ctr: ad.ctr,
    spend: ad.spend,
    conversions: ad.conversions,
    cpa: ad.cpa
  })
  
  loading.value = false
}

const getStatusConfig = (status: string) => {
  const map: Record<string, { status: 'success' | 'warning' | 'info', text: string }> = {
    active: { status: 'success', text: '投放中' },
    paused: { status: 'warning', text: '已暂停' },
    pending: { status: 'info', text: '审核中' }
  }
  return map[status] || { status: 'info', text: status }
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

const handlePause = () => {
  ad.status = ad.status === 'active' ? 'paused' : 'active'
  alert(`广告已${ad.status === 'active' ? '启动' : '暂停'}`)
}

const handleDelete = () => {
  if (confirm('确定要删除此广告吗？')) {
    alert('广告已删除')
    router.push('/ads')
  }
}

onMounted(() => {
  fetchAdDetail()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告管理', path: '/ads' }, { name: '广告详情' }]" />
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <h1 class="text-3xl font-bold text-gray-900">{{ ad.name }}</h1>
          <StatusBadge
            v-if="!loading"
            :status="getStatusConfig(ad.status).status"
            :text="getStatusConfig(ad.status).text"
            show-icon
          />
        </div>
        <div class="flex gap-3">
          <button
            @click="handlePause"
            :class="[
              'px-4 py-2 rounded-lg transition-colors',
              ad.status === 'active' 
                ? 'bg-orange-100 text-orange-700 hover:bg-orange-200' 
                : 'bg-green-100 text-green-700 hover:bg-green-200'
            ]"
          >
            {{ ad.status === 'active' ? '暂停投放' : '启动投放' }}
          </button>
          <router-link
            :to="`/ads/${adId}/edit`"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            编辑
          </router-link>
          <button
            @click="handleDelete"
            class="px-4 py-2 bg-red-100 text-red-700 rounded-lg hover:bg-red-200 transition-colors"
          >
            删除
          </button>
        </div>
      </div>
    </div>

    <!-- 数据概览 -->
    <div class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-6 gap-4">
      <StatsCard title="展示量" :value="formatNumber(stats.impressions)" color="blue">
        <template #icon>
          <svg class="h-6 w-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
          </svg>
        </template>
      </StatsCard>
      
      <StatsCard title="点击量" :value="formatNumber(stats.clicks)" color="green">
        <template #icon>
          <svg class="h-6 w-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122"/>
          </svg>
        </template>
      </StatsCard>
      
      <StatsCard title="点击率" :value="`${stats.ctr}%`" color="purple">
        <template #icon>
          <svg class="h-6 w-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
          </svg>
        </template>
      </StatsCard>
      
      <StatsCard title="消耗" :value="formatMoney(stats.spend)" color="orange">
        <template #icon>
          <svg class="h-6 w-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
      
      <StatsCard title="转化数" :value="formatNumber(stats.conversions)" color="green">
        <template #icon>
          <svg class="h-6 w-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
      
      <StatsCard title="转化成本" :value="formatMoney(stats.cpa)" color="red">
        <template #icon>
          <svg class="h-6 w-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h2>
        <div class="space-y-4">
          <div class="flex justify-between">
            <span class="text-gray-500">广告ID</span>
            <span class="text-gray-900 font-medium">{{ ad.id }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">广告名称</span>
            <span class="text-gray-900">{{ ad.name }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">所属系列</span>
            <router-link :to="`/campaigns/${ad.campaignId}`" class="text-blue-600 hover:text-blue-800">
              {{ ad.campaign }}
            </router-link>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">创建时间</span>
            <span class="text-gray-900">{{ ad.createdAt }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">更新时间</span>
            <span class="text-gray-900">{{ ad.updatedAt }}</span>
          </div>
        </div>
      </div>

      <!-- 投放设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">投放设置</h2>
        <div class="space-y-4">
          <div class="flex justify-between">
            <span class="text-gray-500">日预算</span>
            <span class="text-gray-900">{{ ad.budget ? formatMoney(ad.budget) : '不限' }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">出价方式</span>
            <span class="text-gray-900">{{ ad.bidType.toUpperCase() }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">出价金额</span>
            <span class="text-gray-900">{{ formatMoney(ad.bidAmount) }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">投放开始</span>
            <span class="text-gray-900">{{ ad.startTime }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">投放结束</span>
            <span class="text-gray-900">{{ ad.endTime || '长期投放' }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 创意预览 -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">创意预览</h2>
      <div class="flex gap-6">
        <div class="w-64 h-64 bg-gray-100 rounded-lg flex items-center justify-center">
          <svg class="h-16 w-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
          </svg>
        </div>
        <div class="flex-1 space-y-4">
          <div>
            <h3 class="text-sm font-medium text-gray-500">创意标题</h3>
            <p class="text-gray-900">限时优惠，立即抢购！</p>
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-500">创意描述</h3>
            <p class="text-gray-900">精选好物，品质保障，全场满减活动进行中...</p>
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-500">落地页</h3>
            <a href="#" class="text-blue-600 hover:text-blue-800">https://example.com/landing</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
