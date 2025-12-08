<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 25 })

const syncConfigs = ref([
  { id: 'SC001', name: '主站商品同步', source: 'API', frequency: 'hourly', lastSync: '2025-11-28 09:00', status: 'active', syncCount: 12560 },
  { id: 'SC002', name: '促销商品同步', source: 'Feed', frequency: 'daily', lastSync: '2025-11-28 00:00', status: 'active', syncCount: 3200 },
  { id: 'SC003', name: '测试商品同步', source: 'API', frequency: 'manual', lastSync: '2025-11-27 15:30', status: 'paused', syncCount: 500 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddSyncConfig = () => {
  alert('添加同步配置')
}

const handleSyncNow = (config: typeof syncConfigs.value[0]) => {
  alert(`立即同步: ${config.name}`)
}

const handleEditSyncConfig = (config: typeof syncConfigs.value[0]) => {
  alert(`编辑配置: ${config.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: '商品同步' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">商品同步配置</h1>
          <p class="mt-2 text-gray-600">管理商品数据自动同步</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleAddSyncConfig">
          添加同步配置
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">同步配置</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">运行中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">22</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日同步商品</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">16,260</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">同步成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">99.5%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">配置名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">数据源</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">同步频率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">上次同步</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">同步商品数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="config in syncConfigs" :key="config.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ config.name }}</div>
              <div class="text-xs text-gray-500">{{ config.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ config.source }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">
              {{ config.frequency === 'hourly' ? '每小时' : config.frequency === 'daily' ? '每天' : '手动' }}
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ config.lastSync }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ config.syncCount.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     config.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ config.status === 'active' ? '运行中' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleSyncNow(config)">立即同步</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleEditSyncConfig(config)">编辑</button>
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
