<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const reports = ref([
  { id: 'RPT001', name: '每日投放概览', type: 'daily', dimensions: ['日期', '账户'], metrics: ['消耗', '展示', '点击'], schedule: '每日08:00', status: 'active', lastRun: '2025-11-28 08:00' },
  { id: 'RPT002', name: '周度素材分析', type: 'weekly', dimensions: ['素材', '创意类型'], metrics: ['消耗', 'CTR', 'CVR'], schedule: '每周一09:00', status: 'active', lastRun: '2025-11-25 09:00' },
  { id: 'RPT003', name: '月度ROI报表', type: 'monthly', dimensions: ['广告组', '转化类型'], metrics: ['消耗', '转化', 'ROI'], schedule: '每月1日10:00', status: 'active', lastRun: '2025-11-01 10:00' },
  { id: 'RPT004', name: '人群效果对比', type: 'custom', dimensions: ['人群包', '地域'], metrics: ['消耗', 'CPA'], schedule: '手动', status: 'draft', lastRun: '-' }
])

const getTypeLabel = (type: string) => {
  const types: Record<string, string> = { daily: '日报', weekly: '周报', monthly: '月报', custom: '自定义' }
  return types[type] || type
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleRun = (item: typeof reports.value[0]) => {
  alert(`运行报表: ${item.name}`)
}

const handleEdit = (item: typeof reports.value[0]) => {
  alert(`编辑报表: ${item.name}`)
}

const handleDownload = (item: typeof reports.value[0]) => {
  alert(`下载报表: ${item.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表中心' }, { name: '自定义报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">自定义报表</h1>
          <p class="mt-2 text-gray-600">创建和管理自定义数据报表</p>
        </div>
        <router-link to="/report/custom/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建报表
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总报表</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">定时任务</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日生成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">订阅人数</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">45</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">报表名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">维度/指标</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">定时</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">最近生成</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in reports" :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                <div class="text-xs text-gray-500">{{ item.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getTypeLabel(item.type) }}</span>
              </td>
              <td class="px-6 py-4">
                <div class="text-xs text-gray-600">
                  <span class="text-gray-400">维度:</span> {{ item.dimensions.join(', ') }}
                </div>
                <div class="text-xs text-gray-600">
                  <span class="text-gray-400">指标:</span> {{ item.metrics.join(', ') }}
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ item.schedule }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ item.lastRun }}</td>
              <td class="px-6 py-4 text-sm">
<button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleRun(item)">运行</button>
                <button class="text-gray-600 hover:text-gray-800 mr-3" @click="handleEdit(item)">编辑</button>
                <button class="text-gray-600 hover:text-gray-800" @click="handleDownload(item)">下载</button>
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
