<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const targets = ref([
  { id: '1234567', name: '北京科技有限公司', type: 'advertiser', balance: 1250000, status: 'normal' },
  { id: '1234568', name: '上海网络公司', type: 'advertiser', balance: 850000, status: 'normal' },
  { id: '2345678', name: '深圳子代理', type: 'agent', balance: 2500000, status: 'normal' },
  { id: '1234569', name: '广州传媒集团', type: 'advertiser', balance: 180000, status: 'low_balance' },
  { id: '1234570', name: '杭州互联网公司', type: 'advertiser', balance: 50000, status: 'frozen' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工作台' }, { name: '可转账列表' }]" />
      <h1 class="text-3xl font-bold text-gray-900">可转账目标列表</h1>
      <p class="mt-2 text-gray-600">查看所有可进行转账的广告主和代理商账户</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <input type="text" placeholder="搜索账户名称或ID" class="flex-1 min-w-[200px] px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <select class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">全部类型</option>
          <option value="advertiser">广告主</option>
          <option value="agent">子代理</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">全部状态</option>
          <option value="normal">正常</option>
          <option value="low_balance">余额不足</option>
          <option value="frozen">已冻结</option>
        </select>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账户ID</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账户名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">当前余额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="target in targets" :key="target.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ target.id }}</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ target.name }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded text-xs font-medium',
                  target.type === 'advertiser' ? 'bg-blue-100 text-blue-700' : 'bg-purple-100 text-purple-700']">
                  {{ target.type === 'advertiser' ? '广告主' : '子代理' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm" :class="target.balance < 200000 ? 'text-yellow-600 font-medium' : 'text-gray-900'">
                ¥{{ target.balance.toLocaleString() }}
              </td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  target.status === 'normal' ? 'bg-green-100 text-green-800' :
                  target.status === 'low_balance' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                  {{ target.status === 'normal' ? '正常' : target.status === 'low_balance' ? '余额不足' : '已冻结' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <router-link :to="`/workspace/transfer/create?target=${target.id}`"
                             class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                  发起转账
                </router-link>
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
