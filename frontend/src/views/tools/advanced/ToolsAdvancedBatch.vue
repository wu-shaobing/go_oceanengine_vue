<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const batchTasks = ref([
  { id: 'BT001', name: '批量调整出价', type: 'bid', count: 25, status: 'completed', createTime: '2025-11-28 10:00', progress: 100 },
  { id: 'BT002', name: '批量修改预算', type: 'budget', count: 15, status: 'running', createTime: '2025-11-28 10:30', progress: 60 },
  { id: 'BT003', name: '批量暂停广告', type: 'pause', count: 10, status: 'pending', createTime: '2025-11-28 11:00', progress: 0 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateTask = () => {
  alert('创建新的批量任务')
}

const handleViewDetail = (task: typeof batchTasks.value[0]) => {
  alert(`查看任务详情: ${task.name}\n进度: ${task.progress}%`)
}

const handleCancelTask = (task: typeof batchTasks.value[0]) => {
  if (confirm(`确定取消任务「${task.name}」吗？`)) {
    const idx = batchTasks.value.findIndex(t => t.id === task.id)
    if (idx > -1) batchTasks.value.splice(idx, 1)
    alert('任务已取消')
  }
}

const handleBatchAction = (action: string) => {
  alert(`执行批量操作: ${action}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '批量操作' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">批量操作</h1>
          <p class="mt-2 text-gray-600">批量管理广告投放</p>
        </div>
        <button @click="handleCreateTask" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          新建批量任务
        </button>
      </div>
    </div>

<div class="grid grid-cols-4 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6 cursor-pointer hover:shadow-md transition-shadow" @click="handleBatchAction('批量调价')">
        <div class="text-3xl mb-3">💰</div>
        <h4 class="font-medium text-gray-900">批量调价</h4>
        <p class="text-sm text-gray-500 mt-1">批量调整广告出价</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6 cursor-pointer hover:shadow-md transition-shadow" @click="handleBatchAction('批量调预算')">
        <div class="text-3xl mb-3">📊</div>
        <h4 class="font-medium text-gray-900">批量调预算</h4>
        <p class="text-sm text-gray-500 mt-1">批量修改广告预算</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6 cursor-pointer hover:shadow-md transition-shadow" @click="handleBatchAction('批量暂停')">
        <div class="text-3xl mb-3">⏸️</div>
        <h4 class="font-medium text-gray-900">批量暂停</h4>
        <p class="text-sm text-gray-500 mt-1">批量暂停广告投放</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6 cursor-pointer hover:shadow-md transition-shadow" @click="handleBatchAction('批量复制')">
        <div class="text-3xl mb-3">📋</div>
        <h4 class="font-medium text-gray-900">批量复制</h4>
        <p class="text-sm text-gray-500 mt-1">批量复制广告计划</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">任务列表</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">任务名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">数量</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="task in batchTasks" :key="task.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ task.name }}</div>
              <div class="text-xs text-gray-500">{{ task.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ task.type }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ task.count }} 个</td>
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-20 h-2 bg-gray-200 rounded-full mr-2">
                  <div class="h-2 bg-blue-600 rounded-full" :style="{ width: task.progress + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ task.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ task.createTime }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     task.status === 'completed' ? 'bg-green-100 text-green-800' :
                     task.status === 'running' ? 'bg-blue-100 text-blue-800' : 'bg-yellow-100 text-yellow-800']">
                {{ task.status === 'completed' ? '已完成' : task.status === 'running' ? '执行中' : '等待中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button @click="handleViewDetail(task)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button v-if="task.status === 'pending'" @click="handleCancelTask(task)" class="text-red-600 hover:text-red-800">取消</button>
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
