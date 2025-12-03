<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const callbacks = ref([
  { id: 'CB001', name: '转化回调', url: 'https://api.example.com/callback', events: ['purchase', 'register'], successRate: 99.5, lastCall: '2025-11-28 10:30', status: 'active' },
  { id: 'CB002', name: '线索回调', url: 'https://api.example.com/clue', events: ['form_submit'], successRate: 98.2, lastCall: '2025-11-28 10:25', status: 'active' },
  { id: 'CB003', name: '订单回调', url: 'https://api.example.com/order', events: ['order_create', 'payment'], successRate: 97.8, lastCall: '2025-11-28 10:20', status: 'error' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他工具' }, { name: '追踪回调' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">追踪回调配置</h1>
          <p class="mt-2 text-gray-600">管理转化追踪回调接口</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加回调
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">回调配置数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">正常运行</p>
        <p class="text-2xl font-bold text-green-600 mt-1">25</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">异常</p>
        <p class="text-2xl font-bold text-red-600 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日调用</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">12,560</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">回调名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">URL</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">成功率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">最后调用</th>
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
            <td class="px-6 py-4 text-sm text-blue-600 truncate max-w-xs">{{ cb.url }}</td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="event in cb.events" :key="event" class="px-2 py-0.5 bg-gray-100 text-gray-700 rounded text-xs">
                  {{ event }}
                </span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm font-medium" :class="cb.successRate >= 99 ? 'text-green-600' : 'text-yellow-600'">
              {{ cb.successRate }}%
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ cb.lastCall }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     cb.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ cb.status === 'active' ? '正常' : '异常' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-gray-600 hover:text-gray-800">测试</button>
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
