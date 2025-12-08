<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 25 })

const promotions = ref([
  { id: 'AP001', app: '购物 App', platform: 'iOS', type: '应用下载', budget: 50000, spend: 35600, installs: 8560, cpi: 4.16, status: 'active' },
  { id: 'AP002', app: '游戏 App', platform: 'Android', type: '应用激活', budget: 30000, spend: 22800, installs: 5230, cpi: 4.36, status: 'active' },
  { id: 'AP003', app: '工具 App', platform: '全平台', type: '应用下载', budget: 20000, spend: 15600, installs: 3890, cpi: 4.01, status: 'paused' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建推广')
}

const handleEdit = (promo: typeof promotions.value[0]) => {
  alert(`编辑推广: ${promo.app}`)
}

const handleData = (promo: typeof promotions.value[0]) => {
  alert(`查看数据: ${promo.app}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: '应用推广' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">应用推广</h1>
          <p class="mt-2 text-gray-600">管理应用推广计划</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建推广
        </button>
      </div>
    </div>

    <div class="grid grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">推广计划</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总预算</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥100,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥74,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总安装</p>
        <p class="text-2xl font-bold text-green-600 mt-1">17,680</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CPI</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥4.18</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平台</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">推广类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">预算</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">安装数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CPI</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="promo in promotions" :key="promo.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ promo.app }}</div>
              <div class="text-xs text-gray-500">{{ promo.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     promo.platform === 'iOS' ? 'bg-gray-800 text-white' :
                     promo.platform === 'Android' ? 'bg-green-100 text-green-700' : 'bg-blue-100 text-blue-700']">
                {{ promo.platform }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ promo.type }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ promo.budget.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ promo.spend.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ promo.installs.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">¥{{ promo.cpi.toFixed(2) }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     promo.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ promo.status === 'active' ? '投放中' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleEdit(promo)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleData(promo)" class="text-gray-600 hover:text-gray-800">数据</button>
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
