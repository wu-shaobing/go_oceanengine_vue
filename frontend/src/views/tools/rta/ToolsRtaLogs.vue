<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 20, total: 5680 })

const logs = ref([
  { id: 'RL001', strategyName: '高价值用户策略', requestId: 'req_abc123', responseTime: 32, result: 'accept', bidPrice: 15.5, time: '2025-11-28 09:30:25.123' },
  { id: 'RL002', strategyName: '新用户转化策略', requestId: 'req_def456', responseTime: 45, result: 'reject', bidPrice: 0, time: '2025-11-28 09:30:24.987' },
  { id: 'RL003', strategyName: '高价值用户策略', requestId: 'req_ghi789', responseTime: 28, result: 'accept', bidPrice: 18.2, time: '2025-11-28 09:30:24.856' },
  { id: 'RL004', strategyName: '促销活动策略', requestId: 'req_jkl012', responseTime: 120, result: 'timeout', bidPrice: 0, time: '2025-11-28 09:30:24.654' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: 'RTA日志' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">RTA调用日志</h1>
          <p class="mt-2 text-gray-600">查看RTA接口调用记录</p>
        </div>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
          导出日志
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总调用次数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total.toLocaleString() }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">接受出价</p>
        <p class="text-2xl font-bold text-green-600 mt-1">4,256</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">拒绝出价</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">1,380</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">超时</p>
        <p class="text-2xl font-bold text-red-600 mt-1">44</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均响应</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">38ms</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部策略</option>
          <option value="RTA001">高价值用户策略</option>
          <option value="RTA002">新用户转化策略</option>
          <option value="RTA003">促销活动策略</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部结果</option>
          <option value="accept">接受</option>
          <option value="reject">拒绝</option>
          <option value="timeout">超时</option>
        </select>
        <input type="text" placeholder="搜索请求ID..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">策略名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">请求ID</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">响应时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">结果</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="log in logs" :key="log.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm text-gray-900">{{ log.strategyName }}</td>
            <td class="px-6 py-4">
              <code class="text-xs bg-gray-100 px-2 py-1 rounded">{{ log.requestId }}</code>
            </td>
            <td class="px-6 py-4 text-sm" :class="log.responseTime <= 50 ? 'text-green-600' : log.responseTime <= 100 ? 'text-yellow-600' : 'text-red-600'">
              {{ log.responseTime }}ms
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     log.result === 'accept' ? 'bg-green-100 text-green-800' :
                     log.result === 'reject' ? 'bg-yellow-100 text-yellow-800' :
                     'bg-red-100 text-red-800']">
                {{ log.result === 'accept' ? '接受' : log.result === 'reject' ? '拒绝' : '超时' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ log.bidPrice > 0 ? `¥${log.bidPrice}` : '-' }}</td>
            <td class="px-6 py-4 text-xs text-gray-500 font-mono">{{ log.time }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800">详情</button>
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
