<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const exportTasks = ref([
  { id: 'EX001', name: '11月广告数据报表', type: 'ad', period: '2025-11-01 ~ 2025-11-28', status: 'completed', fileSize: '2.5MB', createdAt: '2025-11-28 10:30' },
  { id: 'EX002', name: '转化数据周报', type: 'convert', period: '2025-11-22 ~ 2025-11-28', status: 'completed', fileSize: '1.2MB', createdAt: '2025-11-28 09:00' },
  { id: 'EX003', name: '素材效果报表', type: 'creative', period: '2025-11-01 ~ 2025-11-28', status: 'processing', fileSize: '-', createdAt: '2025-11-28 11:00' },
  { id: 'EX004', name: '账户消耗明细', type: 'cost', period: '2025-10-01 ~ 2025-10-31', status: 'completed', fileSize: '5.8MB', createdAt: '2025-11-01 10:00' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表中心' }, { name: '报表导出' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">报表导出</h1>
          <p class="mt-2 text-gray-600">管理报表导出任务</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建导出任务
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总任务数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">42</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">处理中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日导出</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">8</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">报表名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">数据周期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">文件大小</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="task in exportTasks" :key="task.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ task.name }}</div>
              <div class="text-xs text-gray-500">{{ task.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                {{ task.type === 'ad' ? '广告数据' : task.type === 'convert' ? '转化数据' : task.type === 'creative' ? '素材数据' : '消耗数据' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ task.period }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     task.status === 'completed' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800']">
                {{ task.status === 'completed' ? '已完成' : '处理中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ task.fileSize }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ task.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button v-if="task.status === 'completed'" class="text-blue-600 hover:text-blue-800">下载</button>
              <span v-else class="text-gray-400">等待中...</span>
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
