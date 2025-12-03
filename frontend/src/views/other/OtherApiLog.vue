<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 20, total: 1256 })

const logs = ref([
  { id: 'LOG001', endpoint: '/api/v3/ad/get', method: 'GET', status: 200, duration: 125, time: '2025-11-28 10:30:15', ip: '192.168.1.100' },
  { id: 'LOG002', endpoint: '/api/v3/campaign/update', method: 'POST', status: 200, duration: 256, time: '2025-11-28 10:29:45', ip: '192.168.1.100' },
  { id: 'LOG003', endpoint: '/api/v3/creative/create', method: 'POST', status: 400, duration: 89, time: '2025-11-28 10:28:30', ip: '192.168.1.101' },
  { id: 'LOG004', endpoint: '/api/v3/report/get', method: 'GET', status: 200, duration: 568, time: '2025-11-28 10:27:15', ip: '192.168.1.100' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他工具' }, { name: 'API日志' }]" />
      <h1 class="text-3xl font-bold text-gray-900">API调用日志</h1>
      <p class="mt-2 text-gray-600">查看API请求记录</p>
    </div>

    <div class="grid grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日请求</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total.toLocaleString() }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">成功请求</p>
        <p class="text-2xl font-bold text-green-600 mt-1">1,198</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">失败请求</p>
        <p class="text-2xl font-bold text-red-600 mt-1">58</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均耗时</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">156ms</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">95.4%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input type="text" placeholder="搜索接口路径..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部状态</option>
          <option value="200">成功 (2xx)</option>
          <option value="400">客户端错误 (4xx)</option>
          <option value="500">服务端错误 (5xx)</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部方法</option>
          <option value="GET">GET</option>
          <option value="POST">POST</option>
          <option value="PUT">PUT</option>
          <option value="DELETE">DELETE</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">筛选</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">方法</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接口</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">耗时</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">IP</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="log in logs" :key="log.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm text-gray-500">{{ log.time }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs font-medium',
                     log.method === 'GET' ? 'bg-green-100 text-green-700' :
                     log.method === 'POST' ? 'bg-blue-100 text-blue-700' :
                     log.method === 'PUT' ? 'bg-yellow-100 text-yellow-700' : 'bg-red-100 text-red-700']">
                {{ log.method }}
              </span>
            </td>
            <td class="px-6 py-4">
              <code class="text-sm text-gray-800">{{ log.endpoint }}</code>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     log.status >= 200 && log.status < 300 ? 'bg-green-100 text-green-800' :
                     log.status >= 400 && log.status < 500 ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                {{ log.status }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm" :class="log.duration > 500 ? 'text-red-600' : log.duration > 200 ? 'text-yellow-600' : 'text-green-600'">
              {{ log.duration }}ms
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ log.ip }}</td>
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
