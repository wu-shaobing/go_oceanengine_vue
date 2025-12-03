<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const phones = ref([
  { id: 'SP001', name: '主营销热线', number: '400-xxx-xxxx', callsToday: 156, answeredRate: 92, avgDuration: 185, status: 'active' },
  { id: 'SP002', name: '售后服务', number: '400-xxx-xxxx', callsToday: 89, answeredRate: 88, avgDuration: 245, status: 'active' },
  { id: 'SP003', name: '活动咨询', number: '400-xxx-xxxx', callsToday: 234, answeredRate: 78, avgDuration: 120, status: 'active' },
  { id: 'SP004', name: '测试号码', number: '400-xxx-xxxx', callsToday: 0, answeredRate: 0, avgDuration: 0, status: 'inactive' }
])

const formatDuration = (seconds: number) => {
  if (seconds === 0) return '-'
  const min = Math.floor(seconds / 60)
  const sec = seconds % 60
  return `${min}分${sec}秒`
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '智能电话' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">智能电话管理</h1>
          <p class="mt-2 text-gray-600">管理广告投放电话号码</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          申请号码
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">号码总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日来电</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">479</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">接通率</p>
        <p class="text-2xl font-bold text-green-600 mt-1">86%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均时长</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">3分12秒</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">号码</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日来电</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接通率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平均时长</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="phone in phones" :key="phone.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ phone.name }}</div>
              <div class="text-xs text-gray-500">{{ phone.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ phone.number }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ phone.callsToday }}</td>
            <td class="px-6 py-4">
              <span :class="['text-sm font-medium', phone.answeredRate >= 85 ? 'text-green-600' : phone.answeredRate >= 70 ? 'text-yellow-600' : 'text-red-600']">
                {{ phone.answeredRate }}%
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ formatDuration(phone.avgDuration) }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                phone.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ phone.status === 'active' ? '启用' : '停用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">配置</button>
              <button class="text-gray-600 hover:text-gray-800">通话记录</button>
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
