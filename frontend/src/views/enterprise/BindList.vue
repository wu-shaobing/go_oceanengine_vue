<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '绑定管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">绑定管理</h1>
        <p class="text-gray-600 mt-1">管理企业号绑定的账号和应用</p>
      </div>
      <button @click="handleAddBind" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">添加绑定</button>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">账号信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">绑定类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">绑定时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in bindings" :key="item.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="item.avatar" class="w-10 h-10 rounded-full mr-3" alt="">
                <div>
                  <div class="font-medium">{{ item.name }}</div>
                  <div class="text-sm text-gray-500">{{ item.accountId }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ item.type }}</td>
            <td class="px-4 py-3">
              <span :class="item.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" class="px-2 py-1 text-xs rounded">
                {{ item.status === 'active' ? '已绑定' : '已解绑' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ item.bindTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button @click="handleBindDetail(item)" class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button @click="handleUnbind(item)" class="text-red-600 hover:text-red-800 text-sm">解绑</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="20" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const handleAddBind = () => {
  alert('添加绑定')
}

const handleBindDetail = (item: typeof bindings.value[0]) => {
  alert(`绑定详情: ${item.name}`)
}

const handleUnbind = (item: typeof bindings.value[0]) => {
  if (confirm(`确定解绑 ${item.name} 吗？`)) {
    alert(`已解绑: ${item.name}`)
  }
}

const bindings = ref([
  { id: 1, name: '品牌官方账号', accountId: 'brand_001', avatar: 'https://via.placeholder.com/40', type: '抖音号', status: 'active', bindTime: '2024-01-15' },
  { id: 2, name: '品牌小店', accountId: 'shop_001', avatar: 'https://via.placeholder.com/40', type: '抖店', status: 'active', bindTime: '2024-02-20' },
  { id: 3, name: '直播间账号', accountId: 'live_001', avatar: 'https://via.placeholder.com/40', type: '直播', status: 'active', bindTime: '2024-03-10' },
  { id: 4, name: '客服账号', accountId: 'service_001', avatar: 'https://via.placeholder.com/40', type: '客服', status: 'inactive', bindTime: '2024-01-20' }
])
</script>
