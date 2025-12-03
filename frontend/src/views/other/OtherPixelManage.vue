<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const pixels = ref([
  { id: 'PX001', name: '官网像素', pixelId: 'CGBJWZ..', domain: 'www.example.com', events: 12, lastActive: '2025-11-28 10:15', status: 'active' },
  { id: 'PX002', name: '落地页像素', pixelId: 'CGHKTA..', domain: 'landing.example.com', events: 8, lastActive: '2025-11-28 09:45', status: 'active' },
  { id: 'PX003', name: '测试像素', pixelId: 'CGTESZ..', domain: 'test.example.com', events: 3, lastActive: '2025-11-27 16:30', status: 'inactive' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他工具' }, { name: '像素管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">像素管理</h1>
          <p class="mt-2 text-gray-600">管理追踪像素代码</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建像素
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">像素总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃像素</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日触发</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">58,960</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联事件</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">23</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">像素名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">像素ID</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联域名</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">最后活跃</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="pixel in pixels" :key="pixel.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ pixel.name }}</div>
              <div class="text-xs text-gray-500">{{ pixel.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-800">{{ pixel.pixelId }}</code>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ pixel.domain }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ pixel.events }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ pixel.lastActive }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     pixel.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ pixel.status === 'active' ? '活跃' : '未活跃' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">代码</button>
              <button class="text-gray-600 hover:text-gray-800 mr-3">编辑</button>
              <button class="text-red-600 hover:text-red-800">删除</button>
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
