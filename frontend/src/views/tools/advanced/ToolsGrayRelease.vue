<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 12 })

const grayConfigs = ref([
  { id: 'GR001', name: '新素材测试', feature: 'new_creative', ratio: 10, status: 'running', startAt: '2025-11-25' },
  { id: 'GR002', name: '出价策略优化', feature: 'bid_strategy', ratio: 20, status: 'running', startAt: '2025-11-20' },
  { id: 'GR003', name: '落地页新版本', feature: 'landing_v2', ratio: 50, status: 'completed', startAt: '2025-11-15' },
  { id: 'GR004', name: '智能投放测试', feature: 'smart_delivery', ratio: 5, status: 'paused', startAt: '2025-11-22' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '灰度发布' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">灰度发布</h1>
          <p class="mt-2 text-gray-600">管理功能灰度发布配置</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建灰度
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总配置数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">运行中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">5</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">4</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已暂停</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">3</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">灰度名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">功能标识</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">灰度比例</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">开始时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="config in grayConfigs" :key="config.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ config.name }}</div>
              <div class="text-xs text-gray-500">{{ config.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-700">{{ config.feature }}</code>
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-20 h-2 bg-gray-200 rounded-full">
                  <div class="h-full bg-blue-500 rounded-full" :style="{ width: config.ratio + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ config.ratio }}%</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     config.status === 'running' ? 'bg-green-100 text-green-800' :
                     config.status === 'completed' ? 'bg-blue-100 text-blue-800' : 'bg-yellow-100 text-yellow-800']">
                {{ config.status === 'running' ? '运行中' : config.status === 'completed' ? '已完成' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ config.startAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button v-if="config.status === 'running'" class="text-yellow-600 hover:text-yellow-800">暂停</button>
              <button v-else-if="config.status === 'paused'" class="text-green-600 hover:text-green-800">恢复</button>
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
