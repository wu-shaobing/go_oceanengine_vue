<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '操作日志' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">操作日志</h1>
      <p class="text-gray-600 mt-1">查看企业号账号操作记录</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频操作</option>
          <option value="comment">评论操作</option>
          <option value="setting">设置操作</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作内容</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作人</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">结果</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="log in logs" :key="log.id">
            <td class="px-4 py-3 text-sm text-gray-500">{{ log.time }}</td>
            <td class="px-4 py-3">
              <span :class="getTypeClass(log.type)" class="px-2 py-1 text-xs rounded">{{ log.type }}</span>
            </td>
            <td class="px-4 py-3 text-sm">{{ log.content }}</td>
            <td class="px-4 py-3 text-sm">{{ log.operator }}</td>
            <td class="px-4 py-3">
              <span :class="log.result === 'success' ? 'text-green-600' : 'text-red-600'" class="text-sm">
                {{ log.result === 'success' ? '成功' : '失败' }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', type: '' })

const logs = ref([
  { id: 1, time: '2024-06-18 15:30:25', type: '视频操作', content: '发布视频《618大促预告》', operator: '管理员', result: 'success' },
  { id: 2, time: '2024-06-18 14:20:10', type: '评论操作', content: '回复用户评论', operator: '客服', result: 'success' },
  { id: 3, time: '2024-06-18 11:15:30', type: '设置操作', content: '修改账号简介', operator: '管理员', result: 'success' },
  { id: 4, time: '2024-06-17 18:45:20', type: '视频操作', content: '删除视频', operator: '管理员', result: 'success' },
  { id: 5, time: '2024-06-17 16:30:00', type: '评论操作', content: '隐藏评论', operator: '客服', result: 'success' }
])

const getTypeClass = (type: string) => {
  const classes: Record<string, string> = {
    '视频操作': 'bg-blue-100 text-blue-800',
    '评论操作': 'bg-green-100 text-green-800',
    '设置操作': 'bg-orange-100 text-orange-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}
</script>
