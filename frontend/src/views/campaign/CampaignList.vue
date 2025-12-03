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
  total: 156,
  active: 89,
  paused: 45,
  totalBudget: 2580000
})

const campaigns = ref<any[]>([])
const pagination = reactive({
  current: 1,
  total: 156,
  pageSize: 10
})

const columns = [
  { key: 'id', title: '广告系列 ID' },
  { key: 'name', title: '名称' },
  { key: 'budget', title: '预算', align: 'right' as const },
  { key: 'spend', title: '消耗', align: 'right' as const },
  { key: 'impressions', title: '展示', align: 'right' as const },
  { key: 'clicks', title: '点击', align: 'right' as const },
  { key: 'ctr', title: 'CTR', align: 'right' as const },
  { key: 'status', title: '状态' },
  { key: 'actions', title: '操作' }
]

const fetchCampaigns = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  campaigns.value = Array.from({ length: 10 }, (_, i) => ({
    id: `C${100001 + i + (pagination.current - 1) * 10}`,
    name: `广告系列 ${i + 1 + (pagination.current - 1) * 10}`,
    budget: Math.floor(Math.random() * 50000) + 10000,
    spend: Math.floor(Math.random() * 30000) + 5000,
    impressions: Math.floor(Math.random() * 500000) + 100000,
    clicks: Math.floor(Math.random() * 10000) + 2000,
    ctr: (Math.random() * 3 + 1).toFixed(2),
    status: Math.random() > 0.3 ? 'active' : 'paused',
    createTime: '2024-01-15'
  }))
  
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchCampaigns()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchCampaigns()
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

onMounted(() => {
  fetchCampaigns()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告系列' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告系列</h1>
          <p class="mt-2 text-gray-600">管理您的广告投放系列</p>
        </div>
        <router-link
          to="/campaigns/create"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建广告系列
        </router-link>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
      <StatsCard title="广告系列总数" :value="stats.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"/>
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

      <StatsCard title="总预算" :value="formatMoney(stats.totalBudget)" color="purple">
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">广告系列列表</h3>
          <div class="flex items-center gap-3">
            <div class="relative">
              <input
                v-model="searchKeyword"
                type="text"
                placeholder="搜索广告系列..."
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
            </select>
          </div>
        </div>
      </div>

      <DataTable
        :columns="columns"
        :data="campaigns"
        :loading="loading"
        selectable
      >
        <template #id="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #name="{ value, row }">
          <router-link :to="`/campaigns/${row.id}`" class="text-blue-600 hover:text-blue-800">
            {{ value }}
          </router-link>
        </template>

        <template #budget="{ value }">
          <span class="text-gray-900">{{ formatMoney(value) }}</span>
        </template>

        <template #spend="{ value }">
          <span class="text-gray-900">{{ formatMoney(value) }}</span>
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

        <template #status="{ value }">
          <StatusBadge
            :status="value === 'active' ? 'success' : 'warning'"
            :text="value === 'active' ? '投放中' : '已暂停'"
            show-icon
          />
        </template>

        <template #actions="{ row }">
          <div class="flex items-center gap-2">
            <router-link :to="`/campaigns/${row.id}`" class="text-blue-600 hover:text-blue-800">
              查看
            </router-link>
            <button class="text-gray-600 hover:text-gray-800">
              {{ row.status === 'active' ? '暂停' : '启用' }}
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
