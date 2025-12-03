<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '代理商广告主' }]" />
    
    <div class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">代理商广告主</h1>
        <p class="text-gray-600 mt-1">管理代理的广告主账户</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">添加广告主</button>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="filters.keyword" placeholder="搜索广告主" class="border border-gray-300 rounded px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="active">正常</option>
          <option value="inactive">停用</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">广告主总数</div>
        <div class="text-2xl font-bold mt-1">{{ stats.total }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">活跃广告主</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.active }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">本月消耗</div>
        <div class="text-2xl font-bold mt-1">¥{{ (stats.monthSpent / 10000).toFixed(1) }}万</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">进行中任务</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.activeTasks }}</div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">广告主</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">行业</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">本月消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="adv in advertisers" :key="adv.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ adv.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ adv.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ adv.industry }}</td>
            <td class="px-4 py-3 text-right">¥{{ adv.balance.toLocaleString() }}</td>
            <td class="px-4 py-3 text-right">¥{{ adv.monthSpent.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <span :class="adv.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" class="px-2 py-1 text-xs rounded">
                {{ adv.status === 'active' ? '正常' : '停用' }}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">充值</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="50" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ keyword: '', status: '' })

const stats = ref({ total: 25, active: 18, monthSpent: 1568000, activeTasks: 32 })

const advertisers = ref([
  { id: 'ADV001', name: '美妆品牌A', industry: '美妆护肤', balance: 125600, monthSpent: 356800, status: 'active' },
  { id: 'ADV002', name: '服装品牌B', industry: '服装配饰', balance: 89500, monthSpent: 268500, status: 'active' },
  { id: 'ADV003', name: '食品品牌C', industry: '食品饮料', balance: 56800, monthSpent: 185600, status: 'active' },
  { id: 'ADV004', name: '数码品牌D', industry: '数码3C', balance: 12500, monthSpent: 125800, status: 'inactive' }
])
</script>
