<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'

const loading = ref(true)
const searchKeyword = ref('')
const typeFilter = ref('all')

const stats = reactive({
  total: 45,
  custom: 28,
  lookalike: 12,
  retargeting: 5
})

const audiences = ref<any[]>([])
const pagination = reactive({
  current: 1,
  total: 45,
  pageSize: 10
})

const columns = [
  { key: 'id', title: '人群ID' },
  { key: 'name', title: '人群名称' },
  { key: 'type', title: '类型' },
  { key: 'size', title: '人群规模', align: 'right' as const },
  { key: 'usedIn', title: '使用广告数', align: 'right' as const },
  { key: 'status', title: '状态' },
  { key: 'updateTime', title: '更新时间' },
  { key: 'actions', title: '操作' }
]

const audienceTypes = ['custom', 'lookalike', 'retargeting']
const typeLabels: Record<string, string> = {
  custom: '自定义人群',
  lookalike: '相似人群',
  retargeting: '再营销人群'
}

const fetchAudiences = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  audiences.value = Array.from({ length: 10 }, (_, i) => ({
    id: `AUD${500001 + i + (pagination.current - 1) * 10}`,
    name: `人群包 ${i + 1 + (pagination.current - 1) * 10}`,
    type: audienceTypes[Math.floor(Math.random() * 3)],
    size: Math.floor(Math.random() * 5000000) + 100000,
    usedIn: Math.floor(Math.random() * 20),
    status: Math.random() > 0.2 ? 'ready' : 'processing',
    updateTime: '2024-01-15 14:30'
  }))
  
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchAudiences()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchAudiences()
}

const formatNumber = (value: number) => {
  if (value >= 1000000) {
    return `${(value / 1000000).toFixed(1)}M`
  } else if (value >= 1000) {
    return `${(value / 1000).toFixed(0)}K`
  }
  return value.toString()
}

onMounted(() => {
  fetchAudiences()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '受众管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">受众管理</h1>
          <p class="mt-2 text-gray-600">管理您的目标受众人群</p>
        </div>
        <button
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建人群
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
      <StatsCard title="人群总数" :value="stats.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="自定义人群" :value="stats.custom" color="green">
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="相似人群" :value="stats.lookalike" color="purple">
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="再营销人群" :value="stats.retargeting" color="orange">
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">人群列表</h3>
          <div class="flex items-center gap-3">
            <div class="relative">
              <input
                v-model="searchKeyword"
                type="text"
                placeholder="搜索人群..."
                @keyup.enter="handleSearch"
                class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
              <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
            </div>
            <select
              v-model="typeFilter"
              @change="handleSearch"
              class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="all">全部类型</option>
              <option value="custom">自定义人群</option>
              <option value="lookalike">相似人群</option>
              <option value="retargeting">再营销人群</option>
            </select>
          </div>
        </div>
      </div>

      <DataTable :columns="columns" :data="audiences" :loading="loading" selectable>
        <template #id="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #name="{ value }">
          <span class="text-blue-600 hover:text-blue-800 cursor-pointer">{{ value }}</span>
        </template>

        <template #type="{ value }">
          <span class="px-2 py-1 text-xs rounded-full" :class="{
            'bg-blue-100 text-blue-700': value === 'custom',
            'bg-purple-100 text-purple-700': value === 'lookalike',
            'bg-orange-100 text-orange-700': value === 'retargeting'
          }">
            {{ typeLabels[value] }}
          </span>
        </template>

        <template #size="{ value }">
          <span class="text-gray-900">{{ formatNumber(value) }}</span>
        </template>

        <template #usedIn="{ value }">
          <span class="text-gray-600">{{ value }}</span>
        </template>

        <template #status="{ value }">
          <StatusBadge
            :status="value === 'ready' ? 'success' : 'info'"
            :text="value === 'ready' ? '已就绪' : '处理中'"
            show-icon
          />
        </template>

        <template #actions>
          <div class="flex items-center gap-2">
            <button class="text-blue-600 hover:text-blue-800">详情</button>
            <button class="text-gray-600 hover:text-gray-800">编辑</button>
            <button class="text-red-600 hover:text-red-800">删除</button>
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
