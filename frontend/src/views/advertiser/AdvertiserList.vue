<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAdvertiserStore } from '@/stores/advertiser'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import DataTable from '@/components/common/DataTable.vue'

const store = useAdvertiserStore()
const searchKeyword = ref('')
const statusFilter = ref('all')

const columns = [
  { key: 'id', title: '广告主 ID' },
  { key: 'name', title: '名称' },
  { key: 'company', title: '公司' },
  { key: 'balance', title: '余额', align: 'right' as const },
  { key: 'status', title: '状态' },
  { key: 'createTime', title: '创建时间' },
  { key: 'actions', title: '操作' }
]

const pagination = computed(() => ({
  current: store.currentPage,
  total: store.total,
  pageSize: store.pageSize
}))

const handleSearch = () => {
  store.fetchAdvertisers({
    keyword: searchKeyword.value,
    status: statusFilter.value
  })
}

const handlePageChange = (page: number) => {
  store.fetchAdvertisers({
    page,
    keyword: searchKeyword.value,
    status: statusFilter.value
  })
}

const formatBalance = (value: number) => {
  return `¥${value.toLocaleString()}`
}

onMounted(() => {
  store.fetchAdvertisers()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告主管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告主管理</h1>
          <p class="mt-2 text-gray-600">管理您的广告主账户</p>
        </div>
        <router-link
          to="/advertisers/create"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          添加广告主
        </router-link>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
      <StatsCard title="总账户数" :value="store.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="启用账户" :value="store.enabledCount" color="green">
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="总余额" :value="formatBalance(store.totalBalance)" color="orange">
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-lg border border-gray-200">
      <!-- Table Header -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">广告主列表</h3>
          <div class="flex items-center gap-3">
            <div class="relative">
              <input
                v-model="searchKeyword"
                type="text"
                placeholder="搜索广告主..."
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
              <option value="enabled">启用</option>
              <option value="disabled">禁用</option>
            </select>
          </div>
        </div>
      </div>

      <DataTable
        :columns="columns"
        :data="store.advertisers"
        :loading="store.loading"
        :pagination="pagination"
        selectable
        @page-change="handlePageChange"
      >
        <template #id="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #name="{ value }">
          <span class="text-gray-900">{{ value }}</span>
        </template>

        <template #company="{ value }">
          <span class="text-gray-600">{{ value }}</span>
        </template>

        <template #balance="{ value }">
          <span class="font-medium text-gray-900">{{ formatBalance(value) }}</span>
        </template>

        <template #status="{ value }">
          <StatusBadge
            :status="value === 'enabled' ? 'success' : 'error'"
            :text="value === 'enabled' ? '启用' : '禁用'"
            show-icon
          />
        </template>

        <template #actions="{ row }">
          <div class="flex items-center gap-2">
            <router-link :to="`/advertisers/${row.id}?mode=edit`" class="text-blue-600 hover:text-blue-800">
              编辑
            </router-link>
            <router-link :to="`/advertisers/${row.id}`" class="text-gray-600 hover:text-gray-800">
              查看
            </router-link>
          </div>
        </template>
      </DataTable>
    </div>
  </div>
</template>
