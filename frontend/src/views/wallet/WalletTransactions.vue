<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 20, total: 256 })
const searchKeyword = ref('')
const filterType = ref('')
const startDate = ref('')
const endDate = ref('')

const transactions = ref([
  { id: 'W20251111001', type: '广告消耗', advertiser: '北京科技有限公司', amount: -5680, balance: 5674320, time: '2025-11-11 15:30:25' },
  { id: 'W20251111002', type: '充值', advertiser: '上海网络公司', amount: 100000, balance: 5680000, time: '2025-11-11 14:30:00' },
  { id: 'W20251111003', type: '广告消耗', advertiser: '深圳电商公司', amount: -3240, balance: 5580000, time: '2025-11-11 13:45:18' },
  { id: 'W20251111004', type: '转账', advertiser: '广州传媒集团', amount: -50000, balance: 5583240, time: '2025-11-11 11:20:00' },
  { id: 'W20251111005', type: '服务费', advertiser: '系统', amount: -1200, balance: 5633240, time: '2025-11-11 00:00:00' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleSearch = () => {
  alert('搜索流水记录')
}

const handleExport = () => {
  alert('导出流水记录')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '财务管理' }, { name: '流水明细' }]" />
      <h1 class="text-3xl font-bold text-gray-900">流水明细</h1>
      <p class="mt-2 text-gray-600">查看共享钱包所有交易流水</p>
    </div>

    <!-- Filter -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <input v-model="searchKeyword" type="text" placeholder="搜索流水号" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 w-48">
        <select v-model="filterType" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">全部类型</option>
          <option value="spend">广告消耗</option>
          <option value="recharge">充值</option>
          <option value="transfer">转账</option>
          <option value="fee">服务费</option>
        </select>
        <input v-model="startDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <input v-model="endDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
<button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
        <button class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50" @click="handleExport">导出</button>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">流水号</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联方</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">金额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">余额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="tx in transactions" :key="tx.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ tx.id }}</td>
              <td class="px-6 py-4">
                <span :class="['inline-flex px-2 py-0.5 rounded-full text-xs font-medium',
                  tx.type === '充值' ? 'bg-green-100 text-green-800' :
                  tx.type === '广告消耗' ? 'bg-blue-100 text-blue-800' :
                  tx.type === '转账' ? 'bg-orange-100 text-orange-800' : 'bg-gray-100 text-gray-800']">
                  {{ tx.type }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ tx.advertiser }}</td>
              <td class="px-6 py-4 text-sm font-medium" :class="tx.amount > 0 ? 'text-green-600' : 'text-red-600'">
                {{ tx.amount > 0 ? '+' : '' }}¥{{ tx.amount.toLocaleString() }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">¥{{ tx.balance.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ tx.time }}</td>
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
