<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const callbacks = ref([
  { id: 'FC001', name: '主回调配置', crmName: '飞鱼CRM', syncType: 'realtime', successRate: 99.2, todaySync: 1256, status: 'active' },
  { id: 'FC002', name: '备用回调', crmName: '飞鱼CRM-备用', syncType: 'batch', successRate: 98.5, todaySync: 890, status: 'active' },
  { id: 'FC003', name: '测试回调', crmName: '测试环境', syncType: 'realtime', successRate: 95.0, todaySync: 45, status: 'paused' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddCallback = () => {
  alert('创建新的飞鱼CRM回调配置')
}

const handleEditCallback = (cb: typeof callbacks.value[0]) => {
  alert(`编辑回调配置: ${cb.name}`)
}

const handleViewLogs = (cb: typeof callbacks.value[0]) => {
  alert(`查看「${cb.name}」的同步日志`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '飞鱼回调' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">飞鱼CRM回调</h1>
          <p class="mt-2 text-gray-600">配置飞鱼CRM线索同步回调</p>
        </div>
        <button @click="handleAddCallback" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加回调
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">回调配置</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">32</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日同步</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">2,191</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">98.6%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">配置名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CRM名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">同步方式</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">成功率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日同步</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="cb in callbacks" :key="cb.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ cb.name }}</div>
              <div class="text-xs text-gray-500">{{ cb.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ cb.crmName }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     cb.syncType === 'realtime' ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700']">
                {{ cb.syncType === 'realtime' ? '实时' : '批量' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm font-medium" :class="cb.successRate >= 99 ? 'text-green-600' : 'text-yellow-600'">
              {{ cb.successRate }}%
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ cb.todaySync.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     cb.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ cb.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button @click="handleEditCallback(cb)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleViewLogs(cb)" class="text-gray-600 hover:text-gray-800">日志</button>
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
