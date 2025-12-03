<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const packages = ref([
  { id: 'PKG001', name: '高端消费人群', type: 'custom', source: 'DMP', coverage: 1520000, status: 'available', createdAt: '2025-10-01' },
  { id: 'PKG002', name: '电商活跃用户', type: 'lookalike', source: '种子扩展', coverage: 3200000, status: 'available', createdAt: '2025-10-05' },
  { id: 'PKG003', name: 'APP活跃30天', type: 'retargeting', source: '行为数据', coverage: 890000, status: 'calculating', createdAt: '2025-10-08' },
  { id: 'PKG004', name: '已转化用户', type: 'converted', source: '转化追踪', coverage: 156000, status: 'available', createdAt: '2025-10-10' }
])

const formatNumber = (num: number) => {
  if (num >= 10000000) return (num / 10000000).toFixed(1) + '千万'
  if (num >= 10000) return (num / 10000).toFixed(0) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '人群工具' }, { name: '人群包' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">人群包管理</h1>
          <p class="mt-2 text-gray-600">创建和管理广告定向人群包</p>
        </div>
        <router-link to="/tools/audience/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建人群包
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总人群包</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">28</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总覆盖</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">5.8千万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联计划</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">156</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">人群包</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">来源</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">覆盖人数</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in packages" :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                <div class="text-xs text-gray-500">{{ item.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ item.type }}</span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ item.source }}</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ formatNumber(item.coverage) }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  item.status === 'available' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                  {{ item.status === 'available' ? '可用' : '计算中' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
                <button class="text-gray-600 hover:text-gray-800">删除</button>
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
