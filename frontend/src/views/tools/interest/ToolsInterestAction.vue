<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 128 })

const actions = ref([
  { id: 'IA001', name: '电商购物', category: '消费行为', coverage: 85000000, usageCount: 1256, status: 'active' },
  { id: 'IA002', name: '短视频观看', category: '内容消费', coverage: 120000000, usageCount: 2340, status: 'active' },
  { id: 'IA003', name: '游戏付费', category: '游戏行为', coverage: 25000000, usageCount: 890, status: 'active' },
  { id: 'IA004', name: '直播互动', category: '社交行为', coverage: 45000000, usageCount: 567, status: 'active' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '行为兴趣定向' }]" />
      <h1 class="text-3xl font-bold text-gray-900">行为兴趣定向</h1>
      <p class="mt-2 text-gray-600">基于用户行为的精准定向</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">行为标签</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">行为分类</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总覆盖人群</p>
        <p class="text-2xl font-bold text-green-600 mt-1">2.7亿</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">使用次数</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">5,053</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部分类</option>
          <option value="consume">消费行为</option>
          <option value="content">内容消费</option>
          <option value="game">游戏行为</option>
          <option value="social">社交行为</option>
        </select>
        <input type="text" placeholder="搜索行为标签..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">行为名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分类</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">覆盖人群</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">使用次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="action in actions" :key="action.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ action.name }}</div>
              <div class="text-xs text-gray-500">{{ action.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ action.category }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ (action.coverage / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ action.usageCount.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-green-100 text-green-800 rounded-full text-xs font-medium">可用</span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">应用</button>
              <button class="text-gray-600 hover:text-gray-800">详情</button>
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
