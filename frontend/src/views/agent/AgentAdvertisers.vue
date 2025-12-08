<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const loading = ref(false)
const searchForm = reactive({
  keyword: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 128
})

const advertisers = ref([
  { id: 1234567, name: '北京科技有限公司', company: '北京科技有限公司', balance: 1250000, todaySpend: 45680, status: 'active', createdAt: '2024-01-15' },
  { id: 1234568, name: '上海网络科技', company: '上海网络科技有限公司', balance: 850000, todaySpend: 38920, status: 'active', createdAt: '2024-02-20' },
  { id: 1234569, name: '深圳电商公司', company: '深圳电商有限公司', balance: 180000, todaySpend: 52340, status: 'warning', createdAt: '2024-03-10' },
  { id: 1234570, name: '广州传媒集团', company: '广州传媒集团有限公司', balance: 2350000, todaySpend: 89500, status: 'active', createdAt: '2024-01-08' },
  { id: 1234571, name: '杭州互联网公司', company: '杭州互联网有限公司', balance: 50000, todaySpend: 0, status: 'inactive', createdAt: '2024-04-05' }
])

const handleSearch = () => {
  loading.value = true
  setTimeout(() => loading.value = false, 500)
}

const handlePageChange = (page: number) => {
  pagination.page = page
  handleSearch()
}

const handleAdvDetail = (adv: typeof advertisers.value[0]) => {
  alert(`广告主详情: ${adv.name} (ID: ${adv.id})`)
}

const handleAdvRecharge = (adv: typeof advertisers.value[0]) => {
  alert(`为 ${adv.name} 充值`)
}

const handleAdvEdit = (adv: typeof advertisers.value[0]) => {
  alert(`编辑广告主: ${adv.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '广告主列表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告主管理</h1>
          <p class="mt-2 text-gray-600">管理您代理的广告主账户</p>
        </div>
        <router-link to="/agent/advertiser/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          新增广告主
        </router-link>
      </div>
    </div>

    <!-- Filter Bar -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <input v-model="searchForm.keyword" type="text" placeholder="搜索广告主名称或ID" 
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        </div>
        <select v-model="searchForm.status" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">全部状态</option>
          <option value="active">正常</option>
          <option value="warning">余额不足</option>
          <option value="inactive">已暂停</option>
        </select>
        <button @click="handleSearch" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
          搜索
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总广告主</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃广告主</p>
        <p class="text-2xl font-bold text-green-600 mt-1">95</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日总消耗</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥125.8万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总余额</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥568万</p>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告主</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账户余额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日消耗</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="adv in advertisers" :key="adv.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ adv.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ adv.id }}</div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm" :class="adv.balance < 200000 ? 'text-yellow-600 font-medium' : 'text-gray-900'">
                ¥{{ adv.balance.toLocaleString() }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">¥{{ adv.todaySpend.toLocaleString() }}</td>
              <td class="px-6 py-4">
                <span :class="['inline-flex px-2 py-0.5 rounded-full text-xs font-medium',
                  adv.status === 'active' ? 'bg-green-100 text-green-800' :
                  adv.status === 'warning' ? 'bg-yellow-100 text-yellow-800' : 'bg-gray-100 text-gray-800']">
                  {{ adv.status === 'active' ? '正常' : adv.status === 'warning' ? '余额不足' : '已暂停' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ adv.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
                <button @click="handleAdvDetail(adv)" class="text-blue-600 hover:text-blue-900 mr-3">详情</button>
                <button @click="handleAdvRecharge(adv)" class="text-green-600 hover:text-green-900 mr-3">充值</button>
                <button @click="handleAdvEdit(adv)" class="text-gray-600 hover:text-gray-900">编辑</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
