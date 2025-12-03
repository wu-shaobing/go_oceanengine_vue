<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 25 })

const childAgents = ref([
  { id: '2345671', name: '深圳子代理A', company: '深圳网络科技有限公司', advertisers: 45, balance: 2500000, monthSpend: 8500000, status: 'active', createdAt: '2024-01-15' },
  { id: '2345672', name: '广州子代理B', company: '广州传媒有限公司', advertisers: 32, balance: 1800000, monthSpend: 5200000, status: 'active', createdAt: '2024-02-20' },
  { id: '2345673', name: '杭州子代理C', company: '杭州互联网有限公司', advertisers: 18, balance: 500000, monthSpend: 2100000, status: 'warning', createdAt: '2024-03-10' },
  { id: '2345674', name: '成都子代理D', company: '成都科技有限公司', advertisers: 0, balance: 100000, monthSpend: 0, status: 'inactive', createdAt: '2024-04-05' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '二级代理商' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">二级代理商列表</h1>
          <p class="mt-2 text-gray-600">管理您的二级代理商账户</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          新增子代理
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总子代理</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃子代理</p>
        <p class="text-2xl font-bold text-green-600 mt-1">22</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">子代理总余额</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥490万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">本月总消耗</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥1,580万</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">子代理</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告主数</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账户余额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">本月消耗</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="agent in childAgents" :key="agent.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ agent.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ agent.id }}</div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ agent.advertisers }}</td>
              <td class="px-6 py-4 text-sm" :class="agent.balance < 500000 ? 'text-yellow-600 font-medium' : 'text-gray-900'">
                ¥{{ agent.balance.toLocaleString() }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">¥{{ agent.monthSpend.toLocaleString() }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  agent.status === 'active' ? 'bg-green-100 text-green-800' :
                  agent.status === 'warning' ? 'bg-yellow-100 text-yellow-800' : 'bg-gray-100 text-gray-800']">
                  {{ agent.status === 'active' ? '活跃' : agent.status === 'warning' ? '余额不足' : '未激活' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ agent.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
                <button class="text-green-600 hover:text-green-800 mr-3">充值</button>
                <button class="text-gray-600 hover:text-gray-800">编辑</button>
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
