<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const apps = ref([
  { id: 'UA001', name: '主APP', platform: 'android', packageName: 'com.example.main', slotCount: 12, dau: 850000, status: 'approved' },
  { id: 'UA002', name: '主APP-iOS', platform: 'ios', packageName: 'com.example.mainios', slotCount: 10, dau: 620000, status: 'approved' },
  { id: 'UA003', name: '游戏APP', platform: 'android', packageName: 'com.example.game', slotCount: 8, dau: 320000, status: 'approved' },
  { id: 'UA004', name: '工具APP', platform: 'android', packageName: 'com.example.tool', slotCount: 5, dau: 180000, status: 'pending' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAdd = () => {
  alert('添加应用')
}

const handleManage = (app: typeof apps.value[0]) => {
  alert(`管理应用: ${app.name}`)
}

const handleDetail = (app: typeof apps.value[0]) => {
  alert(`查看详情: ${app.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '穿山甲应用' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">应用管理</h1>
          <p class="mt-2 text-gray-600">管理穿山甲变现应用</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleAdd">
          添加应用
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">应用总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已上线</p>
        <p class="text-2xl font-bold text-green-600 mt-1">25</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总DAU</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">197万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">广告位数</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">85</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平台</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">包名</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告位数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">DAU</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="app in apps" :key="app.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ app.name }}</div>
              <div class="text-xs text-gray-500">{{ app.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     app.platform === 'android' ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-700']">
                {{ app.platform === 'android' ? 'Android' : 'iOS' }}
              </span>
            </td>
            <td class="px-6 py-4">
              <code class="text-xs bg-gray-100 px-2 py-1 rounded">{{ app.packageName }}</code>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ app.slotCount }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (app.dau / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     app.status === 'approved' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ app.status === 'approved' ? '已上线' : '审核中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleManage(app)">管理</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleDetail(app)">详情</button>
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
