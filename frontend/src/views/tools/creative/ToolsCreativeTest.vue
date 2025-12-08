<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const tests = ref([
  { id: 'T001', name: '双十一素材测试', variants: 3, status: 'running', startTime: '2025-11-25', winner: null, improvement: '+15.2%' },
  { id: 'T002', name: '标题文案测试', variants: 2, status: 'completed', startTime: '2025-11-20', winner: 'A', improvement: '+8.5%' },
  { id: 'T003', name: '落地页测试', variants: 4, status: 'draft', startTime: '-', winner: null, improvement: '-' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建A/B测试')
}

const handleDetail = (test: typeof tests.value[0]) => {
  alert(`查看测试详情: ${test.name}`)
}

const handleStart = (test: typeof tests.value[0]) => {
  alert(`启动测试: ${test.name}`)
}

const handleStop = (test: typeof tests.value[0]) => {
  alert(`停止测试: ${test.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '创意工具' }, { name: 'A/B测试' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">创意A/B测试</h1>
          <p class="mt-2 text-gray-600">对比测试不同创意效果</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreate">
          创建测试
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">测试总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">运行中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">5</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">11</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均提升</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">+12.3%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">测试名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">变体数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">开始时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">优胜方案</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">效果提升</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="test in tests" :key="test.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ test.name }}</div>
              <div class="text-xs text-gray-500">{{ test.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ test.variants }} 个</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ test.startTime }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     test.status === 'running' ? 'bg-blue-100 text-blue-800' :
                     test.status === 'completed' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ test.status === 'running' ? '运行中' : test.status === 'completed' ? '已完成' : '草稿' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <span v-if="test.winner" class="px-2 py-1 bg-yellow-100 text-yellow-800 rounded font-medium">方案 {{ test.winner }}</span>
              <span v-else class="text-gray-400">-</span>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ test.improvement }}</td>
            <td class="px-6 py-4 text-sm">
<button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleDetail(test)">详情</button>
              <button v-if="test.status === 'draft'" class="text-green-600 hover:text-green-800" @click="handleStart(test)">启动</button>
              <button v-if="test.status === 'running'" class="text-yellow-600 hover:text-yellow-800" @click="handleStop(test)">停止</button>
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
