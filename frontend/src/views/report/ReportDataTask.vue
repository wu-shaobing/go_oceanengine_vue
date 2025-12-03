<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 56 })

const tasks = ref([
  { id: 'TASK001', name: '11月投放数据导出', type: 'export', status: 'completed', progress: 100, fileSize: '15.2MB', createdAt: '2025-11-25 10:30', completedAt: '2025-11-25 10:32' },
  { id: 'TASK002', name: '素材效果分析报告', type: 'analysis', status: 'running', progress: 68, fileSize: '-', createdAt: '2025-11-26 09:15', completedAt: '-' },
  { id: 'TASK003', name: '人群包数据导入', type: 'import', status: 'completed', progress: 100, fileSize: '8.5MB', createdAt: '2025-11-26 08:00', completedAt: '2025-11-26 08:15' },
  { id: 'TASK004', name: 'DPA商品批量更新', type: 'update', status: 'failed', progress: 45, fileSize: '-', createdAt: '2025-11-26 11:00', completedAt: '-' }
])

const getTypeConfig = (type: string) => {
  switch (type) {
    case 'export': return { label: '数据导出', class: 'bg-blue-100 text-blue-700' }
    case 'import': return { label: '数据导入', class: 'bg-green-100 text-green-700' }
    case 'analysis': return { label: '数据分析', class: 'bg-purple-100 text-purple-700' }
    default: return { label: '批量更新', class: 'bg-orange-100 text-orange-700' }
  }
}

const getStatusConfig = (status: string) => {
  switch (status) {
    case 'completed': return { label: '已完成', class: 'bg-green-100 text-green-800' }
    case 'running': return { label: '进行中', class: 'bg-blue-100 text-blue-800' }
    case 'failed': return { label: '失败', class: 'bg-red-100 text-red-800' }
    default: return { label: '等待中', class: 'bg-yellow-100 text-yellow-800' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表中心' }, { name: '数据任务' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">数据任务</h1>
          <p class="mt-2 text-gray-600">查看数据导入导出及分析任务</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建任务
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总任务</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">进行中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">48</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">失败</p>
        <p class="text-2xl font-bold text-red-600 mt-1">5</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">任务</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">文件大小</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="task in tasks" :key="task.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ task.name }}</div>
              <div class="text-xs text-gray-500">{{ task.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs', getTypeConfig(task.type).class]">
                {{ getTypeConfig(task.type).label }}
              </span>
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-24 h-2 bg-gray-200 rounded-full overflow-hidden">
                  <div :class="['h-full rounded-full', task.status === 'failed' ? 'bg-red-500' : 'bg-blue-500']"
                       :style="{ width: task.progress + '%' }"></div>
                </div>
                <span class="text-xs text-gray-600">{{ task.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ task.fileSize }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ task.createdAt }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusConfig(task.status).class]">
                {{ getStatusConfig(task.status).label }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button v-if="task.status === 'completed'" class="text-blue-600 hover:text-blue-800 mr-3">下载</button>
              <button v-if="task.status === 'failed'" class="text-orange-600 hover:text-orange-800 mr-3">重试</button>
              <button class="text-gray-600 hover:text-gray-800">详情</button>
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
