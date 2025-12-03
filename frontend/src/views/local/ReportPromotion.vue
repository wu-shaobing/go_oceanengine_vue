<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '数据报表' }, { name: '推广报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">推广报表</h1>
      <p class="text-gray-600 mt-1">查看推广计划的详细数据</p>
    </div>

    <!-- 筛选条件 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-end">
        <div>
          <label class="block text-sm text-gray-600 mb-1">日期范围</label>
          <div class="flex items-center space-x-2">
            <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
            <span>至</span>
            <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
          </div>
        </div>
        <div>
          <label class="block text-sm text-gray-600 mb-1">项目</label>
          <select v-model="filters.project" class="border border-gray-300 rounded px-3 py-2 w-40">
            <option value="">全部项目</option>
            <option value="1">北京朝阳店</option>
            <option value="2">上海静安店</option>
          </select>
        </div>
        <div>
          <label class="block text-sm text-gray-600 mb-1">状态</label>
          <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2 w-32">
            <option value="">全部状态</option>
            <option value="running">投放中</option>
            <option value="paused">已暂停</option>
          </select>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <!-- 汇总数据 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">消耗</div>
        <div class="text-xl font-bold text-gray-900 mt-1">¥{{ summary.cost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">曝光量</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ formatNumber(summary.impressions) }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">点击量</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ formatNumber(summary.clicks) }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">点击率</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ summary.ctr }}%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">线索数</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ summary.clues }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">线索成本</div>
        <div class="text-xl font-bold text-gray-900 mt-1">¥{{ summary.clueCost }}</div>
      </div>
    </div>

    <!-- 数据表格 -->
    <div class="bg-white rounded-lg shadow">
      <div class="overflow-x-auto">
        <table class="min-w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">推广名称</th>
              <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">所属项目</th>
              <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">曝光量</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击量</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击率</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索数</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索成本</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化率</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in reportData" :key="item.id">
              <td class="px-4 py-3">
                <div class="font-medium text-blue-600 hover:underline cursor-pointer">{{ item.name }}</div>
                <div class="text-xs text-gray-400">{{ item.id }}</div>
              </td>
              <td class="px-4 py-3 text-sm">{{ item.project }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(item.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ getStatusText(item.status) }}
                </span>
              </td>
              <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ formatNumber(item.impressions) }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ formatNumber(item.clicks) }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ item.ctr }}%</td>
              <td class="px-4 py-3 text-sm text-right">{{ item.clues }}</td>
              <td class="px-4 py-3 text-sm text-right">¥{{ item.clueCost }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ item.cvr }}%</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({
  startDate: '',
  endDate: '',
  project: '',
  status: ''
})

const summary = ref({
  cost: 45680,
  impressions: 1256800,
  clicks: 85600,
  ctr: 6.81,
  clues: 568,
  clueCost: 80.4
})

const reportData = ref([
  { id: 'PM001', name: '618限时优惠推广', project: '北京朝阳店', status: 'running', cost: 15800, impressions: 458600, clicks: 28560, ctr: 6.23, clues: 186, clueCost: 84.9, cvr: 0.65 },
  { id: 'PM002', name: '新品尝鲜活动', project: '北京朝阳店', status: 'running', cost: 6500, impressions: 186500, clicks: 12800, ctr: 6.86, clues: 98, clueCost: 66.3, cvr: 0.77 },
  { id: 'PM003', name: '会员专享折扣', project: '北京朝阳店', status: 'paused', cost: 8600, impressions: 256800, clicks: 18560, ctr: 7.23, clues: 125, clueCost: 68.8, cvr: 0.67 },
  { id: 'PM004', name: '周末特惠推广', project: '上海静安店', status: 'running', cost: 12100, impressions: 325600, clicks: 21580, ctr: 6.63, clues: 145, clueCost: 83.4, cvr: 0.67 },
  { id: 'PM005', name: '新店开业活动', project: '上海静安店', status: 'ended', cost: 2680, impressions: 29300, clicks: 4100, ctr: 14.0, clues: 14, clueCost: 191.4, cvr: 0.34 }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    running: 'bg-green-100 text-green-800',
    paused: 'bg-orange-100 text-orange-800',
    ended: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    running: '投放中',
    paused: '已暂停',
    ended: '已结束'
  }
  return texts[status] || status
}

const formatNumber = (num: number) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w'
  }
  return num.toLocaleString()
}
</script>
