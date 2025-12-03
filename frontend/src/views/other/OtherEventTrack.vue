<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const events = ref([
  { id: 'E001', name: '页面浏览', code: 'page_view', category: '基础事件', count: 125600, status: 'active' },
  { id: 'E002', name: '按钮点击', code: 'button_click', category: '交互事件', count: 89500, status: 'active' },
  { id: 'E003', name: '表单提交', code: 'form_submit', category: '转化事件', count: 5680, status: 'active' },
  { id: 'E004', name: '商品加购', code: 'add_to_cart', category: '电商事件', count: 12560, status: 'inactive' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他工具' }, { name: '事件追踪' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">事件追踪</h1>
          <p class="mt-2 text-gray-600">管理自定义追踪事件</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加事件
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">追踪事件</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已启用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">38</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日触发</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">233,340</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化事件</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">12</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件代码</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分类</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日触发</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="event in events" :key="event.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ event.name }}</div>
              <div class="text-xs text-gray-500">{{ event.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-800">{{ event.code }}</code>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ event.category }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ event.count.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     event.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ event.status === 'active' ? '已启用' : '未启用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
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
