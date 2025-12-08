<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '工具' }, { name: '行业数据' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">行业数据</h1>
      <p class="text-gray-600 mt-1">查看行业投放数据和趋势</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <select v-model="filters.industry" class="border border-gray-300 rounded px-3 py-2">
          <option value="">选择行业</option>
          <option value="fashion">服饰鞋包</option>
          <option value="beauty">美妆个护</option>
          <option value="food">食品饮料</option>
          <option value="home">家居家装</option>
          <option value="digital">数码家电</option>
        </select>
        <select v-model="filters.period" class="border border-gray-300 rounded px-3 py-2">
          <option value="7d">近7天</option>
          <option value="30d">近30天</option>
          <option value="90d">近90天</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700" @click="handleQuery">查询</button>
      </div>
    </div>

    <!-- 行业指标 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">行业平均CPM</div>
        <div class="text-xl font-bold mt-1">¥{{ industryData.avgCpm }}</div>
        <div class="text-sm text-green-500 mt-1">较上周 -5.2%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">行业平均CTR</div>
        <div class="text-xl font-bold mt-1">{{ industryData.avgCtr }}%</div>
        <div class="text-sm text-green-500 mt-1">较上周 +2.1%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">行业平均CVR</div>
        <div class="text-xl font-bold mt-1">{{ industryData.avgCvr }}%</div>
        <div class="text-sm text-red-500 mt-1">较上周 -1.5%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">行业平均ROI</div>
        <div class="text-xl font-bold mt-1">{{ industryData.avgRoi }}</div>
        <div class="text-sm text-green-500 mt-1">较上周 +3.8%</div>
      </div>
    </div>

    <!-- 趋势图 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">CPM趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">CPM趋势图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">ROI趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">ROI趋势图表</span>
        </div>
      </div>
    </div>

    <!-- 热门品类 -->
    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">热门品类</h3>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">品类</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗占比</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">平均CPM</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">平均CTR</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">平均ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="cat in categories" :key="cat.name">
            <td class="px-4 py-3 text-sm font-medium">{{ cat.name }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ cat.costRatio }}%</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ cat.cpm }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ cat.ctr }}%</td>
            <td class="px-4 py-3 text-sm text-right">{{ cat.roi }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const filters = ref({ industry: '', period: '7d' })

const handleQuery = () => {
  alert('查询行业数据')
}

const industryData = ref({
  avgCpm: 25.6,
  avgCtr: 3.2,
  avgCvr: 2.1,
  avgRoi: 4.5
})

const categories = ref([
  { name: '女装', costRatio: 28.5, cpm: 22.5, ctr: 3.5, roi: 5.2 },
  { name: '美妆', costRatio: 22.3, cpm: 28.6, ctr: 3.2, roi: 4.8 },
  { name: '食品', costRatio: 18.6, cpm: 18.5, ctr: 2.8, roi: 3.9 },
  { name: '家居', costRatio: 15.2, cpm: 25.2, ctr: 2.5, roi: 4.2 },
  { name: '数码', costRatio: 15.4, cpm: 32.5, ctr: 2.2, roi: 3.5 }
])
</script>
