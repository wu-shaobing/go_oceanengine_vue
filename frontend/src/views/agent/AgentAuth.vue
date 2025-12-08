<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const auths = ref([
  { id: 'AU001', advertiser: '电商广告主A', authTime: '2025-10-15', expireTime: '2026-10-15', scope: ['投放', '报表', '素材'], status: 'active' },
  { id: 'AU002', advertiser: '游戏广告主B', authTime: '2025-08-20', expireTime: '2026-08-20', scope: ['投放', '报表'], status: 'active' },
  { id: 'AU003', advertiser: '品牌广告主C', authTime: '2025-06-10', expireTime: '2025-12-10', scope: ['报表'], status: 'expiring' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleApplyAuth = () => {
  alert('申请授权')
}

const handleAuthDetail = (auth: typeof auths.value[0]) => {
  alert(`授权详情: ${auth.advertiser}`)
}

const handleRenewAuth = (auth: typeof auths.value[0]) => {
  alert(`续期授权: ${auth.advertiser}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '授权管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">授权管理</h1>
          <p class="mt-2 text-gray-600">管理广告主授权</p>
        </div>
        <button @click="handleApplyAuth" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          申请授权
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">授权总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">有效授权</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">即将过期</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">2</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已过期</p>
        <p class="text-2xl font-bold text-red-600 mt-1">1</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告主</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">授权时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">过期时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">授权范围</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="auth in auths" :key="auth.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ auth.advertiser }}</div>
              <div class="text-xs text-gray-500">{{ auth.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ auth.authTime }}</td>
            <td class="px-6 py-4 text-sm" :class="auth.status === 'expiring' ? 'text-yellow-600' : 'text-gray-500'">
              {{ auth.expireTime }}
            </td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="scope in auth.scope" :key="scope"
                      class="px-2 py-0.5 bg-blue-100 text-blue-700 rounded text-xs">{{ scope }}</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     auth.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ auth.status === 'active' ? '有效' : '即将过期' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button @click="handleAuthDetail(auth)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button v-if="auth.status === 'expiring'" @click="handleRenewAuth(auth)" class="text-green-600 hover:text-green-800">续期</button>
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
