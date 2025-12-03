<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'

const loading = ref(true)
const searchKeyword = ref('')
const statusFilter = ref('all')

const stats = reactive({
  total: 458,
  active: 312,
  paused: 98,
  pending: 48
})

const ads = ref<any[]>([])
const pagination = reactive({
  current: 1,
  total: 458,
  pageSize: 10
})

const columns = [
  { key: 'id', title: '广告 ID' },
  { key: 'name', title: '广告名称' },
  { key: 'campaign', title: '所属系列' },
  { key: 'impressions', title: '展示', align: 'right' as const },
  { key: 'clicks', title: '点击', align: 'right' as const },
  { key: 'ctr', title: 'CTR', align: 'right' as const },
  { key: 'spend', title: '消耗', align: 'right' as const },
  { key: 'status', title: '状态' },
  { key: 'actions', title: '操作' }
]

const fetchAds = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  ads.value = Array.from({ length: 10 }, (_, i) => ({
    id: `AD${200001 + i + (pagination.current - 1) * 10}`,
    name: `广告创意 ${i + 1 + (pagination.current - 1) * 10}`,
    campaign: `广告系列 ${Math.floor(Math.random() * 10) + 1}`,
    impressions: Math.floor(Math.random() * 200000) + 30000,
    clicks: Math.floor(Math.random() * 5000) + 800,
    ctr: (Math.random() * 3 + 1).toFixed(2),
    spend: Math.floor(Math.random() * 5000) + 1000,
    status: ['active', 'paused', 'pending'][Math.floor(Math.random() * 3)]
  }))
  
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchAds()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchAds()
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

const getStatusConfig = (status: string) => {
  const map: Record<string, { status: 'success' | 'warning' | 'info', text: string }> = {
    active: { status: 'success', text: '投放中' },
    paused: { status: 'warning', text: '已暂停' },
    pending: { status: 'info', text: '审核中' }
  }
  return map[status] || { status: 'info', text: status }
}

onMounted(() => {
  fetchAds()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告管理</h1>
          <p class="mt-2 text-gray-600">管理您的广告创意</p>
        </div>
        <router-link
          to="/ads/create"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建广告
        </router-link>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
      <StatsCard title="广告总数" :value="stats.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="投放中" :value="stats.active" color="green">
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="已暂停" :value="stats.paused" color="orange">
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="审核中" :value="stats.pending" color="purple">
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">广告列表</h3>
          <div class="flex items-center gap-3">
            <div class="relative">
              <input
                v-model="searchKeyword"
                type="text"
                placeholder="搜索广告..."
                @keyup.enter="handleSearch"
                class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
              <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
            </div>
            <select
              v-model="statusFilter"
              @change="handleSearch"
              class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="all">全部状态</option>
              <option value="active">投放中</option>
              <option value="paused">已暂停</option>
              <option value="pending">审核中</option>
            </select>
          </div>
        </div>
      </div>

      <DataTable :columns="columns" :data="ads" :loading="loading" selectable>
        <template #id="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #name="{ value, row }">
          <router-link :to="`/ads/${row.id}`" class="text-blue-600 hover:text-blue-800">
            {{ value }}
          </router-link>
        </template>

        <template #campaign="{ value }">
          <span class="text-gray-600">{{ value }}</span>
        </template>

        <template #impressions="{ value }">
          <span class="text-gray-600">{{ formatNumber(value) }}</span>
        </template>

        <template #clicks="{ value }">
          <span class="text-gray-600">{{ formatNumber(value) }}</span>
        </template>

        <template #ctr="{ value }">
          <span class="text-gray-600">{{ value }}%</span>
        </template>

        <template #spend="{ value }">
          <span class="text-gray-900">{{ formatMoney(value) }}</span>
        </template>

        <template #status="{ value }">
          <StatusBadge
            :status="getStatusConfig(value).status"
            :text="getStatusConfig(value).text"
            show-icon
          />
        </template>

        <template #actions="{ row }">
          <div class="flex items-center gap-2">
            <router-link :to="`/ads/${row.id}`" class="text-blue-600 hover:text-blue-800">
              查看
            </router-link>
            <button class="text-gray-600 hover:text-gray-800">
              复制
            </button>
          </div>
        </template>
      </DataTable>

      <div class="p-4 border-t border-gray-200">
        <Pagination
          :current="pagination.current"
          :total="pagination.total"
          :page-size="pagination.pageSize"
          @change="handlePageChange"
        />
      </div>
    </div>
  </div>
</template>
