<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const coupons = ref([
  { id: 'CPN001', name: '新用户专享券', type: 'discount', value: 50, minSpend: 200, total: 10000, used: 3256, status: 'active', expireAt: '2025-12-31' },
  { id: 'CPN002', name: '满减优惠券', type: 'reduction', value: 100, minSpend: 500, total: 5000, used: 1890, status: 'active', expireAt: '2025-12-15' },
  { id: 'CPN003', name: '限时折扣券', type: 'percent', value: 20, minSpend: 0, total: 2000, used: 2000, status: 'expired', expireAt: '2025-11-20' },
  { id: 'CPN004', name: 'VIP专属券', type: 'discount', value: 200, minSpend: 1000, total: 1000, used: 456, status: 'active', expireAt: '2025-12-31' }
])

const getTypeLabel = (type: string) => {
  switch (type) {
    case 'discount': return '立减券'
    case 'reduction': return '满减券'
    case 'percent': return '折扣券'
    default: return type
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '优惠券' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">优惠券管理</h1>
          <p class="mt-2 text-gray-600">创建和管理营销优惠券</p>
        </div>
        <router-link to="/clue/coupon/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建优惠券
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总优惠券</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">进行中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">28</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已发放</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">18,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已核销</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">7,602</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">优惠券</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">面值/门槛</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发放/使用</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">有效期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="coupon in coupons" :key="coupon.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ coupon.name }}</div>
              <div class="text-xs text-gray-500">{{ coupon.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getTypeLabel(coupon.type) }}</span>
            </td>
            <td class="px-6 py-4 text-sm">
              <div class="font-medium text-red-600">
                {{ coupon.type === 'percent' ? coupon.value + '%折' : '¥' + coupon.value }}
              </div>
              <div class="text-xs text-gray-500">满{{ coupon.minSpend }}可用</div>
            </td>
            <td class="px-6 py-4 text-sm">
              <div class="text-gray-900">{{ coupon.used }}/{{ coupon.total }}</div>
              <div class="w-20 h-1.5 bg-gray-200 rounded-full mt-1">
                <div class="h-full bg-blue-500 rounded-full" :style="{ width: (coupon.used / coupon.total * 100) + '%' }"></div>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ coupon.expireAt }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                coupon.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ coupon.status === 'active' ? '进行中' : '已结束' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-gray-600 hover:text-gray-800">数据</button>
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
