<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '项目管理', path: '/local/project' }, { name: '项目详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ project.name }}</h1>
        <p class="text-gray-600 mt-1">项目ID: {{ project.id }}</p>
      </div>
      <div class="flex space-x-2">
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">编辑</button>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">新建推广</button>
      </div>
    </div>

    <!-- 项目数据 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日消耗</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">¥{{ project.todayCost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日线索</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ project.todayClues }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">线索成本</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">¥{{ project.clueCost }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">累计消耗</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">¥{{ project.totalCost.toLocaleString() }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <!-- 趋势图表 -->
        <div class="bg-white rounded-lg shadow p-4">
          <h3 class="text-lg font-medium mb-4">数据趋势</h3>
          <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
            <span class="text-gray-400">消耗与线索趋势图表</span>
          </div>
        </div>

        <!-- 推广列表 -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-4 border-b">
            <h3 class="text-lg font-medium">关联推广</h3>
          </div>
          <table class="min-w-full">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">推广名称</th>
                <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
                <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
                <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="promo in promotions" :key="promo.id">
                <td class="px-4 py-3 text-sm">{{ promo.name }}</td>
                <td class="px-4 py-3">
                  <span :class="promo.status === 'running' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" class="px-2 py-1 text-xs rounded-full">
                    {{ promo.status === 'running' ? '投放中' : '已暂停' }}
                  </span>
                </td>
                <td class="px-4 py-3 text-sm text-right">¥{{ promo.cost.toLocaleString() }}</td>
                <td class="px-4 py-3 text-sm text-right">{{ promo.clues }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- 项目信息 -->
      <div class="space-y-6">
        <div class="bg-white rounded-lg shadow p-4">
          <h3 class="text-lg font-medium mb-4">项目信息</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-gray-500">项目状态</span>
              <span class="px-2 py-0.5 text-xs bg-green-100 text-green-800 rounded">投放中</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">投放城市</span>
              <span>{{ project.cities }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">行业分类</span>
              <span>{{ project.industry }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">日预算</span>
              <span>¥{{ project.dailyBudget.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">创建时间</span>
              <span>{{ project.createTime }}</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-4">
          <h3 class="text-lg font-medium mb-4">门店信息</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-gray-500">门店名称</span>
              <span>{{ project.storeName }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">门店地址</span>
              <span class="text-sm">{{ project.storeAddress }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">联系电话</span>
              <span>{{ project.storePhone }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const project = ref({
  id: 'LP001',
  name: '北京朝阳店-618大促活动',
  todayCost: 2580,
  todayClues: 45,
  clueCost: 57.3,
  totalCost: 125600,
  cities: '北京市',
  industry: '餐饮/美食',
  dailyBudget: 5000,
  createTime: '2024-05-15',
  storeName: '某某餐厅(朝阳大悦城店)',
  storeAddress: '北京市朝阳区朝阳北路101号大悦城3层',
  storePhone: '010-12345678'
})

const promotions = ref([
  { id: 1, name: '618限时优惠推广', status: 'running', cost: 1580, clues: 28 },
  { id: 2, name: '新品尝鲜活动', status: 'running', cost: 650, clues: 12 },
  { id: 3, name: '会员专享折扣', status: 'paused', cost: 350, clues: 5 }
])
</script>
