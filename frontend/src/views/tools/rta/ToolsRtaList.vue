<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 12 })

const rtaList = ref([
  { id: 'RTA001', name: '高意向用户竞价', endpoint: 'https://api.example.com/rta/bid', status: 'online', avgLatency: 45, qps: 1200, createdAt: '2025-09-15' },
  { id: 'RTA002', name: '实时人群过滤', endpoint: 'https://api.example.com/rta/filter', status: 'online', avgLatency: 32, qps: 2500, createdAt: '2025-09-20' },
  { id: 'RTA003', name: '动态出价策略', endpoint: 'https://api.example.com/rta/price', status: 'testing', avgLatency: 68, qps: 500, createdAt: '2025-10-01' },
  { id: 'RTA004', name: '频控策略', endpoint: 'https://api.example.com/rta/freq', status: 'offline', avgLatency: 0, qps: 0, createdAt: '2025-10-05' }
])

const getStatusConfig = (status: string) => {
  switch (status) {
    case 'online': return { label: '在线', class: 'bg-green-100 text-green-800' }
    case 'testing': return { label: '测试中', class: 'bg-yellow-100 text-yellow-800' }
    default: return { label: '离线', class: 'bg-gray-100 text-gray-800' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleEdit = (item: typeof rtaList.value[0]) => {
  alert(`编辑策略: ${item.name}`)
}

const handleMonitor = (item: typeof rtaList.value[0]) => {
  alert(`查看监控: ${item.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'RTA工具' }, { name: 'RTA策略' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">RTA策略管理</h1>
          <p class="mt-2 text-gray-600">管理实时竞价API策略</p>
        </div>
        <router-link to="/tools/rta/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建策略
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总策略</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">在线</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均延迟</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">38ms</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总QPS</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">4,200</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">策略名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接口地址</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平均延迟</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">QPS</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in rtaList" :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                <div class="text-xs text-gray-500">{{ item.id }}</div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600 font-mono text-xs truncate max-w-[200px]">{{ item.endpoint }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusConfig(item.status).class]">
                  {{ getStatusConfig(item.status).label }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.avgLatency }}ms</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.qps.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm">
<button @click="handleEdit(item)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
                <button @click="handleMonitor(item)" class="text-gray-600 hover:text-gray-800">监控</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
