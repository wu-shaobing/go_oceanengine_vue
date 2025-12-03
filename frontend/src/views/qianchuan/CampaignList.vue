<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '广告计划' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">广告计划</h1>
        <p class="text-gray-600 mt-1">管理千川广告投放计划</p>
      </div>
      <router-link to="/qianchuan/campaign/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建计划
      </router-link>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="计划名称/ID" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="enable">投放中</option>
          <option value="disable">已暂停</option>
          <option value="delete">已删除</option>
        </select>
        <select v-model="filters.marketingGoal" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部目标</option>
          <option value="live">直播带货</option>
          <option value="video">短视频带货</option>
          <option value="product">商品推广</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">重置</button>
      </div>
    </div>

    <!-- 批量操作 -->
    <div class="flex items-center gap-4 mb-4">
      <button class="px-3 py-1 border border-gray-300 rounded hover:bg-gray-50 text-sm">批量启用</button>
      <button class="px-3 py-1 border border-gray-300 rounded hover:bg-gray-50 text-sm">批量暂停</button>
      <button class="px-3 py-1 border border-gray-300 rounded hover:bg-gray-50 text-sm">批量删除</button>
    </div>

    <!-- 计划列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">计划信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">营销目标</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">预算</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">转化数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="campaign in campaigns" :key="campaign.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="font-medium">{{ campaign.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ campaign.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ campaign.marketingGoal }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(campaign.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(campaign.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">¥{{ campaign.budget.toLocaleString() }}/天</td>
            <td class="px-4 py-3 text-sm">¥{{ campaign.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ campaign.conversions }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">数据</button>
                <button v-if="campaign.status === 'enable'" class="text-orange-600 hover:text-orange-800 text-sm">暂停</button>
                <button v-else class="text-green-600 hover:text-green-800 text-sm">启用</button>
              </div>
            </td>
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

const filters = ref({
  keyword: '',
  status: '',
  marketingGoal: ''
})

const campaigns = ref([
  { id: '10001', name: '618大促直播计划', marketingGoal: '直播带货', status: 'enable', budget: 10000, cost: 8560, conversions: 256 },
  { id: '10002', name: '新品短视频推广', marketingGoal: '短视频带货', status: 'enable', budget: 5000, cost: 3280, conversions: 128 },
  { id: '10003', name: '爆款商品推广', marketingGoal: '商品推广', status: 'disable', budget: 8000, cost: 6890, conversions: 345 },
  { id: '10004', name: '日常直播引流', marketingGoal: '直播带货', status: 'enable', budget: 3000, cost: 2100, conversions: 89 },
  { id: '10005', name: '达人带货计划', marketingGoal: '短视频带货', status: 'enable', budget: 15000, cost: 12450, conversions: 567 }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    enable: 'bg-green-100 text-green-800',
    disable: 'bg-gray-100 text-gray-800',
    delete: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    enable: '投放中',
    disable: '已暂停',
    delete: '已删除'
  }
  return texts[status] || status
}
</script>
