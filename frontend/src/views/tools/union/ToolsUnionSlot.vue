<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 85 })

const slots = ref([
  { id: 'US001', name: '开屏广告位', type: 'splash', appName: '主APP', impressions: 1250000, revenue: 25680, ecpm: 20.5, status: 'active' },
  { id: 'US002', name: '信息流广告位', type: 'feed', appName: '主APP', impressions: 3560000, revenue: 45230, ecpm: 12.7, status: 'active' },
  { id: 'US003', name: '激励视频广告位', type: 'reward', appName: '游戏APP', impressions: 890000, revenue: 35680, ecpm: 40.1, status: 'active' },
  { id: 'US004', name: '插屏广告位', type: 'interstitial', appName: '工具APP', impressions: 560000, revenue: 12340, ecpm: 22.0, status: 'paused' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '穿山甲广告位' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">穿山甲广告位管理</h1>
          <p class="mt-2 text-gray-600">管理流量变现广告位</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建广告位
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">广告位总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">626万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日收入</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥118,930</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均eCPM</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥19.0</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告位名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">所属应用</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日收入</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">eCPM</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="slot in slots" :key="slot.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ slot.name }}</div>
              <div class="text-xs text-gray-500">{{ slot.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     slot.type === 'splash' ? 'bg-red-100 text-red-700' :
                     slot.type === 'feed' ? 'bg-blue-100 text-blue-700' :
                     slot.type === 'reward' ? 'bg-green-100 text-green-700' :
                     'bg-purple-100 text-purple-700']">
                {{ slot.type === 'splash' ? '开屏' : slot.type === 'feed' ? '信息流' : slot.type === 'reward' ? '激励视频' : '插屏' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ slot.appName }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ (slot.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">¥{{ slot.revenue.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ slot.ecpm }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     slot.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ slot.status === 'active' ? '运行中' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-gray-600 hover:text-gray-800">数据</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
