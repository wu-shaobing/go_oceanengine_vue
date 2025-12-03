<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 23 })

const accounts = ref([
  { id: 'IES001', name: '主账号', type: 'main', platform: '抖音', bindStatus: 'bound', lastSync: '2025-11-28 10:00' },
  { id: 'IES002', name: '子账号1', type: 'sub', platform: '抖音', bindStatus: 'bound', lastSync: '2025-11-28 09:30' },
  { id: 'IES003', name: '头条账号', type: 'main', platform: '今日头条', bindStatus: 'bound', lastSync: '2025-11-27 18:00' },
  { id: 'IES004', name: '西瓜账号', type: 'main', platform: '西瓜视频', bindStatus: 'unbound', lastSync: '-' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '巨量账号' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">巨量账号管理</h1>
          <p class="mt-2 text-gray-600">管理绑定的巨量引擎账号</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          绑定新账号
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总账号数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已绑定</p>
        <p class="text-2xl font-bold text-green-600 mt-1">20</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">未绑定</p>
        <p class="text-2xl font-bold text-gray-400 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">异常账号</p>
        <p class="text-2xl font-bold text-red-600 mt-1">0</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账号名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平台</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">绑定状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">最后同步</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="account in accounts" :key="account.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ account.name }}</div>
              <div class="text-xs text-gray-500">{{ account.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     account.type === 'main' ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700']">
                {{ account.type === 'main' ? '主账号' : '子账号' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ account.platform }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     account.bindStatus === 'bound' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ account.bindStatus === 'bound' ? '已绑定' : '未绑定' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ account.lastSync }}</td>
            <td class="px-6 py-4 text-sm">
              <button v-if="account.bindStatus === 'bound'" class="text-blue-600 hover:text-blue-800 mr-3">同步</button>
              <button v-else class="text-green-600 hover:text-green-800 mr-3">绑定</button>
              <button class="text-gray-600 hover:text-gray-800">设置</button>
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
