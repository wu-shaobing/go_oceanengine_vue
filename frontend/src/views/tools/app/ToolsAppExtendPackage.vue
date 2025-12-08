<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const packages = ref([
  { id: 'EP001', name: '高活跃用户包', baseApp: 'APP001', coverage: 2560000, matchRate: 78, status: 'available', createdAt: '2025-11-20' },
  { id: 'EP002', name: '高消费用户包', baseApp: 'APP001', coverage: 1890000, matchRate: 65, status: 'available', createdAt: '2025-11-18' },
  { id: 'EP003', name: '流失用户召回包', baseApp: 'APP002', coverage: 890000, matchRate: 45, status: 'available', createdAt: '2025-11-15' },
  { id: 'EP004', name: '新增潜力包', baseApp: 'APP001', coverage: 3200000, matchRate: 82, status: 'calculating', createdAt: '2025-11-25' }
])

const formatNumber = (num: number) => {
  if (num >= 10000000) return (num / 10000000).toFixed(1) + '千万'
  if (num >= 10000) return (num / 10000).toFixed(0) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建拓展包')
}

const handleUse = (pkg: typeof packages.value[0]) => {
  alert(`使用拓展包: ${pkg.name}`)
}

const handleDetail = (pkg: typeof packages.value[0]) => {
  alert(`查看详情: ${pkg.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: '拓展包' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">应用拓展包</h1>
          <p class="mt-2 text-gray-600">基于已有用户数据拓展相似人群</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建拓展包
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">拓展包数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总覆盖人数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">854万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均匹配率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">68%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">拓展包</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">基准应用</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">覆盖人数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">匹配率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="pkg in packages" :key="pkg.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ pkg.name }}</div>
              <div class="text-xs text-gray-500">{{ pkg.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ pkg.baseApp }}</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ formatNumber(pkg.coverage) }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-16 h-1.5 bg-gray-200 rounded-full">
                  <div class="h-full bg-blue-500 rounded-full" :style="{ width: pkg.matchRate + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ pkg.matchRate }}%</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                pkg.status === 'available' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ pkg.status === 'available' ? '可用' : '计算中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ pkg.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
<button @click="handleUse(pkg)" class="text-blue-600 hover:text-blue-800 mr-3">使用</button>
              <button @click="handleDetail(pkg)" class="text-gray-600 hover:text-gray-800">详情</button>
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
