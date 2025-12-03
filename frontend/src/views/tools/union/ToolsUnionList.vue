<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 24 })

const apps = ref([
  { id: 'APP001', name: '趣味小说', platform: 'android', adTypes: ['开屏', '信息流', '激励视频'], dailyPV: 1250000, revenue: 35680, status: 'active' },
  { id: 'APP002', name: '健康计步', platform: 'ios', adTypes: ['Banner', '信息流'], dailyPV: 680000, revenue: 18920, status: 'active' },
  { id: 'APP003', name: '天气预报', platform: 'both', adTypes: ['开屏', '插屏'], dailyPV: 2100000, revenue: 52300, status: 'active' },
  { id: 'APP004', name: '工具箱', platform: 'android', adTypes: ['激励视频'], dailyPV: 450000, revenue: 12560, status: 'reviewing' }
])

const getPlatformLabel = (platform: string) => {
  switch (platform) {
    case 'android': return 'Android'
    case 'ios': return 'iOS'
    default: return '双平台'
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '穿山甲' }, { name: '应用管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">穿山甲联盟</h1>
          <p class="mt-2 text-gray-600">管理穿山甲广告联盟应用</p>
        </div>
        <router-link to="/tools/union/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          接入应用
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">接入应用</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日PV</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">448万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日收益</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥11.9万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">本月收益</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥285万</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平台</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告位类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">日均PV</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">昨日收益</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="app in apps" :key="app.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ app.name }}</div>
                <div class="text-xs text-gray-500">{{ app.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ getPlatformLabel(app.platform) }}</span>
              </td>
              <td class="px-6 py-4">
                <div class="flex flex-wrap gap-1">
                  <span v-for="type in app.adTypes" :key="type" 
                        class="px-2 py-0.5 bg-blue-50 text-blue-700 rounded text-xs">{{ type }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ (app.dailyPV / 10000).toFixed(0) }}万</td>
              <td class="px-6 py-4 text-sm font-medium text-green-600">¥{{ app.revenue.toLocaleString() }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  app.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                  {{ app.status === 'active' ? '已上线' : '审核中' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
                <button class="text-gray-600 hover:text-gray-800">数据</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
