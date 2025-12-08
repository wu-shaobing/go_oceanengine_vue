<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '广告报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">广告报表</h1>
      <p class="text-gray-600 mt-1">查看广告维度的投放数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="enable">投放中</option>
          <option value="disable">已暂停</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700" @click="handleQuery">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50" @click="handleExport">导出</button>
      </div>
    </div>

    <!-- 汇总数据 -->
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">消耗</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.cost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">展示</div>
        <div class="text-xl font-bold mt-1">{{ summary.show.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">点击</div>
        <div class="text-xl font-bold mt-1">{{ summary.click.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化</div>
        <div class="text-xl font-bold mt-1">{{ summary.convert.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化成本</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.convertCost }}</div>
      </div>
    </div>

    <!-- 数据表格 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">广告ID</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">广告名称</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">展示</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">CTR</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化成本</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in reportData" :key="item.adId">
            <td class="px-4 py-3 text-sm">{{ item.adId }}</td>
            <td class="px-4 py-3 text-sm">{{ item.adName }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.show.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.click.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.ctr }}%</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.convert }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.convertCost }}</td>
          </tr>
        </tbody>
      </table>
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

const filters = ref({ startDate: '', endDate: '', status: '' })

const handleQuery = () => {
  alert('查询广告报表')
}

const handleExport = () => {
  alert('导出广告报表')
}

const summary = ref({
  cost: 125680,
  show: 5680000,
  click: 168500,
  convert: 2856,
  convertCost: 44.0
})

const reportData = ref([
  { adId: '1001', adName: '618直播间引流广告', cost: 25680, show: 1250000, click: 38500, ctr: 3.08, convert: 856, convertCost: 30.0 },
  { adId: '1002', adName: '商品推广-爆款', cost: 18560, show: 980000, click: 28600, ctr: 2.92, convert: 625, convertCost: 29.7 },
  { adId: '1003', adName: '短视频带货', cost: 15230, show: 756000, click: 22500, ctr: 2.98, convert: 486, convertCost: 31.3 },
  { adId: '1004', adName: '品牌曝光广告', cost: 12860, show: 650000, click: 18900, ctr: 2.91, convert: 325, convertCost: 39.6 },
  { adId: '1005', adName: '新品首发推广', cost: 10580, show: 520000, click: 15600, ctr: 3.0, convert: 268, convertCost: 39.5 }
])
</script>
