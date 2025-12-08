<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 12 })

const configs = ref([
  { id: 'CC001', name: '主站线索回传', url: 'https://api.example.com/callback', method: 'POST', status: 'active', lastTest: '2025-11-28 09:00', successRate: 99.5 },
  { id: 'CC002', name: '落地页线索回传', url: 'https://crm.example.com/lead', method: 'POST', status: 'active', lastTest: '2025-11-28 08:30', successRate: 98.2 },
  { id: 'CC003', name: '表单线索回传', url: 'https://form.example.com/submit', method: 'POST', status: 'inactive', lastTest: '2025-11-27 15:00', successRate: 95.8 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddConfig = () => {
  alert('添加配置功能开发中...')
}

const handleTest = (config: any) => {
  alert(`正在测试配置: ${config.name}`)
}

const handleEdit = (config: any) => {
  alert(`编辑配置: ${config.name}`)
}

const handleDelete = (config: any) => {
  if (confirm(`确定要删除配置 "${config.name}" 吗？`)) {
    const index = configs.value.findIndex(c => c.id === config.id)
    if (index > -1) {
      configs.value.splice(index, 1)
      alert('删除成功')
    }
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '回传配置' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">线索回传配置</h1>
          <p class="mt-2 text-gray-600">配置线索数据回传接口</p>
        </div>
        <button @click="handleAddConfig" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加配置
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">回传配置</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已启用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">10</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日回传</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">1,256</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">97.8%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">配置名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">回传地址</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">请求方式</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">上次测试</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">成功率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="config in configs" :key="config.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ config.name }}</div>
              <div class="text-xs text-gray-500">{{ config.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="text-xs text-gray-600 bg-gray-100 px-2 py-1 rounded max-w-xs truncate block">{{ config.url }}</code>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ config.method }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ config.lastTest }}</td>
            <td class="px-6 py-4 text-sm">
              <span :class="config.successRate >= 98 ? 'text-green-600' : config.successRate >= 95 ? 'text-yellow-600' : 'text-red-600'">
                {{ config.successRate }}%
              </span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     config.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ config.status === 'active' ? '已启用' : '未启用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button @click="handleTest(config)" class="text-blue-600 hover:text-blue-800 mr-3">测试</button>
              <button @click="handleEdit(config)" class="text-gray-600 hover:text-gray-800 mr-3">编辑</button>
              <button @click="handleDelete(config)" class="text-red-600 hover:text-red-800">删除</button>
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
