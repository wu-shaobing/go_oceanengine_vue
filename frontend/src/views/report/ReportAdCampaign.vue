<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 128 })
const dateRange = ref('7d')

const campaignData = ref([
  { id: 'CAM001', name: '双11促销活动', budget: 50000, cost: 45680, progress: 91.4, impressions: 3560000, clicks: 125600, ctr: 3.5, conversions: 1890 },
  { id: 'CAM002', name: '品牌曝光活动', budget: 30000, cost: 28560, progress: 95.2, impressions: 2560000, clicks: 68900, ctr: 2.7, conversions: 456 },
  { id: 'CAM003', name: '新品上市推广', budget: 20000, cost: 15680, progress: 78.4, impressions: 1280000, clicks: 45600, ctr: 3.6, conversions: 890 },
  { id: 'CAM004', name: '日常投放计划', budget: 10000, cost: 8560, progress: 85.6, impressions: 680000, clicks: 19800, ctr: 2.9, conversions: 325 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleExport = () => {
  alert('导出广告系列报表')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表' }, { name: '广告系列报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告系列报表</h1>
          <p class="mt-2 text-gray-600">按广告系列维度查看投放数据</p>
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

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃系列</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总预算</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥110,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥98,480</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">预算消耗率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">89.5%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">系列名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">预算</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="campaign in campaignData" :key="campaign.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ campaign.name }}</div>
              <div class="text-xs text-gray-500">{{ campaign.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">¥{{ campaign.budget.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ campaign.cost.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-16 h-2 bg-gray-200 rounded-full">
                  <div class="h-full bg-blue-500 rounded-full" :style="{ width: `${campaign.progress}%` }"></div>
                </div>
                <span class="text-xs text-gray-500">{{ campaign.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (campaign.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (campaign.clicks / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm" :class="campaign.ctr >= 3 ? 'text-green-600' : 'text-yellow-600'">{{ campaign.ctr }}%</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ campaign.conversions.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
