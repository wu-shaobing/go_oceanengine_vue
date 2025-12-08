<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })
const dateRange = ref('7d')

const accountData = ref([
  { id: 'ACC001', name: '品牌广告账户', cost: 125680, impressions: 8560000, clicks: 256800, ctr: 3.0, conversions: 3560, cpa: 35.3 },
  { id: 'ACC002', name: '效果广告账户', cost: 89500, impressions: 5620000, clicks: 185600, ctr: 3.3, conversions: 2890, cpa: 31.0 },
  { id: 'ACC003', name: '测试账户', cost: 12560, impressions: 890000, clicks: 25680, ctr: 2.9, conversions: 456, cpa: 27.5 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleExport = () => {
  alert('导出账户报表')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表' }, { name: '账户报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">账户级报表</h1>
          <p class="mt-2 text-gray-600">按账户维度查看广告数据</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
<button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleExport">导出</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥227,740</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">1,507万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">46.8万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">3.1%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总转化</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">6,906</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账户名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CPA</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="account in accountData" :key="account.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ account.name }}</div>
              <div class="text-xs text-gray-500">{{ account.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ account.cost.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (account.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (account.clicks / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm" :class="account.ctr >= 3 ? 'text-green-600' : 'text-yellow-600'">{{ account.ctr }}%</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ account.conversions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ account.cpa }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
