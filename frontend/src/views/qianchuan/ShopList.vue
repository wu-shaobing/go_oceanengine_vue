<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '账户管理' }, { name: '店铺管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">店铺管理</h1>
        <p class="text-gray-600 mt-1">管理已授权的抖音小店</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        添加店铺授权
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="店铺名称/ID" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="active">授权有效</option>
          <option value="expired">授权过期</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="self">自营店铺</option>
          <option value="pop">POP店铺</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">重置</button>
      </div>
    </div>

    <!-- 店铺列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">店铺信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">店铺类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">授权状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">授权到期时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">关联达人数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">累计消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="shop in shops" :key="shop.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="shop.logo" class="w-10 h-10 rounded-lg mr-3" alt="">
                <div>
                  <div class="font-medium">{{ shop.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ shop.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ shop.type }}</td>
            <td class="px-4 py-3">
              <span :class="shop.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" 
                class="px-2 py-1 text-xs rounded-full">
                {{ shop.status === 'active' ? '授权有效' : '授权过期' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">{{ shop.expireTime }}</td>
            <td class="px-4 py-3 text-sm">{{ shop.awemeCount }}</td>
            <td class="px-4 py-3 text-sm">¥{{ shop.totalCost.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">查看详情</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">续期授权</button>
                <button class="text-red-600 hover:text-red-800 text-sm">取消授权</button>
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

const filters = ref({
  keyword: '',
  status: '',
  type: ''
})

const shops = ref([
  { id: '7001', name: '品牌官方旗舰店', logo: 'https://via.placeholder.com/40', type: '自营店铺', status: 'active', expireTime: '2024-12-31', awemeCount: 15, totalCost: 568000 },
  { id: '7002', name: '美妆专营店', logo: 'https://via.placeholder.com/40', type: 'POP店铺', status: 'active', expireTime: '2024-10-15', awemeCount: 8, totalCost: 325000 },
  { id: '7003', name: '食品特卖店', logo: 'https://via.placeholder.com/40', type: 'POP店铺', status: 'expired', expireTime: '2024-03-01', awemeCount: 5, totalCost: 186000 },
  { id: '7004', name: '数码配件店', logo: 'https://via.placeholder.com/40', type: '自营店铺', status: 'active', expireTime: '2024-08-20', awemeCount: 12, totalCost: 428000 },
  { id: '7005', name: '服装精品店', logo: 'https://via.placeholder.com/40', type: 'POP店铺', status: 'active', expireTime: '2024-11-30', awemeCount: 20, totalCost: 756000 }
])
</script>
