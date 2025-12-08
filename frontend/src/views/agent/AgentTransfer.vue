<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 56 })

const searchKeyword = ref('')
const filterType = ref('')
const filterDate = ref('')

const records = ref([
  { id: 'T20251111001', from: '代理商账户', to: '北京科技有限公司', amount: 100000, time: '2025-11-11 14:30', status: '成功', type: '充值' },
  { id: 'T20251111002', from: '代理商账户', to: '上海网络公司', amount: 50000, time: '2025-11-11 10:15', status: '成功', type: '充值' },
  { id: 'T20251110001', from: '代理商账户', to: '深圳电商公司', amount: 200000, time: '2025-11-10 16:45', status: '成功', type: '充值' },
  { id: 'T20251110002', from: '深圳电商公司', to: '代理商账户', amount: 30000, time: '2025-11-10 09:20', status: '成功', type: '退款' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleSearch = () => {
  alert('搜索转账记录')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '转账记录' }]" />
      <h1 class="text-3xl font-bold text-gray-900">转账记录</h1>
      <p class="mt-2 text-gray-600">查看所有资金转账记录</p>
    </div>

    <!-- Filter -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <input v-model="searchKeyword" type="text" placeholder="搜索交易单号" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 w-48">
        <select v-model="filterType" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">全部类型</option>
          <option value="recharge">充值</option>
          <option value="refund">退款</option>
          <option value="transfer">转账</option>
        </select>
        <input v-model="filterDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <button @click="handleSearch" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">交易单号</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转出方</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转入方</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">金额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="record in records" :key="record.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ record.id }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ record.from }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ record.to }}</td>
              <td class="px-6 py-4 text-sm font-medium" :class="record.type === '退款' ? 'text-red-600' : 'text-green-600'">
                {{ record.type === '退款' ? '-' : '+' }}¥{{ record.amount.toLocaleString() }}
              </td>
              <td class="px-6 py-4">
                <span :class="['inline-flex px-2 py-0.5 rounded-full text-xs font-medium', record.type === '充值' ? 'bg-green-100 text-green-800' : 'bg-orange-100 text-orange-800']">
                  {{ record.type }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ record.time }}</td>
              <td class="px-6 py-4">
                <span class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">{{ record.status }}</span>
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
