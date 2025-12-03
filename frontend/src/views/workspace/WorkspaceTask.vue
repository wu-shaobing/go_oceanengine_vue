<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const tasks = ref([
  { id: 'TK001', name: '审核双十一创意', type: 'review', priority: 'high', assignee: '张三', dueDate: '2025-11-28', status: 'pending' },
  { id: 'TK002', name: '优化广告出价', type: 'optimize', priority: 'medium', assignee: '李四', dueDate: '2025-11-29', status: 'in_progress' },
  { id: 'TK003', name: '生成周报', type: 'report', priority: 'low', assignee: '王五', dueDate: '2025-11-30', status: 'completed' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工作台' }, { name: '任务管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">任务管理</h1>
          <p class="mt-2 text-gray-600">管理团队任务和待办事项</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建任务
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总任务</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">待处理</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">进行中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">任务名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">优先级</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">负责人</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">截止日期</th>
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
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ task.type }}</span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs font-medium',
                     task.priority === 'high' ? 'bg-red-100 text-red-700' :
                     task.priority === 'medium' ? 'bg-yellow-100 text-yellow-700' : 'bg-gray-100 text-gray-700']">
                {{ task.priority === 'high' ? '高' : task.priority === 'medium' ? '中' : '低' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ task.assignee }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ task.dueDate }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     task.status === 'completed' ? 'bg-green-100 text-green-800' :
                     task.status === 'in_progress' ? 'bg-blue-100 text-blue-800' : 'bg-yellow-100 text-yellow-800']">
                {{ task.status === 'completed' ? '已完成' : task.status === 'in_progress' ? '进行中' : '待处理' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
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
