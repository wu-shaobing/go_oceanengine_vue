<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const projects = ref([
  { id: 'PJ001', name: '品牌词否定', wordsCount: 156, campaignsCount: 12, status: 'active', updatedAt: '2025-11-25' },
  { id: 'PJ002', name: '竞品词否定', wordsCount: 89, campaignsCount: 8, status: 'active', updatedAt: '2025-11-24' },
  { id: 'PJ003', name: '无效流量过滤', wordsCount: 234, campaignsCount: 15, status: 'active', updatedAt: '2025-11-23' },
  { id: 'PJ004', name: '测试项目', wordsCount: 12, campaignsCount: 2, status: 'paused', updatedAt: '2025-11-20' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建否定词项目')
}

const handleManage = (project: typeof projects.value[0]) => {
  alert(`管理项目: ${project.name}`)
}

const handleEdit = (project: typeof projects.value[0]) => {
  alert(`编辑项目: ${project.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '否定词项目' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">否定词项目</h1>
          <p class="mt-2 text-gray-600">按项目管理否定关键词</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建项目
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总项目</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">38</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">否定词总数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">491</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联计划</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">37</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">项目名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">否定词数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联计划</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">更新时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="project in projects" :key="project.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ project.name }}</div>
              <div class="text-xs text-gray-500">{{ project.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ project.wordsCount }} 个</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ project.campaignsCount }} 个</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                project.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ project.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ project.updatedAt }}</td>
            <td class="px-6 py-4 text-sm">
<button @click="handleManage(project)" class="text-blue-600 hover:text-blue-800 mr-3">管理</button>
              <button @click="handleEdit(project)" class="text-gray-600 hover:text-gray-800">编辑</button>
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
