<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '账户信息' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">账户信息</h1>
      <p class="text-gray-600 mt-1">查看和管理星图账户信息</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">基本信息</h3>
          <div class="flex items-start">
            <img src="https://via.placeholder.com/80" class="w-20 h-20 rounded mr-6" alt="">
            <div class="flex-1 space-y-3">
              <div class="flex items-center">
                <span class="text-gray-500 w-24">账户名称</span>
                <span class="font-medium">{{ account.name }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">账户ID</span>
                <span>{{ account.id }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">账户类型</span>
                <span>{{ account.type }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">认证状态</span>
                <span class="px-2 py-0.5 text-xs bg-green-100 text-green-800 rounded">已认证</span>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">企业信息</h3>
          <div class="space-y-3">
            <div class="flex items-center">
              <span class="text-gray-500 w-24">企业名称</span>
              <span>{{ account.companyName }}</span>
            </div>
            <div class="flex items-center">
              <span class="text-gray-500 w-24">营业执照</span>
              <span>{{ account.license }}</span>
            </div>
            <div class="flex items-center">
              <span class="text-gray-500 w-24">联系人</span>
              <span>{{ account.contact }}</span>
            </div>
            <div class="flex items-center">
              <span class="text-gray-500 w-24">联系电话</span>
              <span>{{ account.phone }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">资金概览</h3>
          <div class="space-y-4">
            <div class="flex justify-between">
              <span class="text-gray-500">账户余额</span>
              <span class="font-bold text-xl">¥{{ account.balance.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">冻结金额</span>
              <span class="text-orange-600">¥{{ account.frozen.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">累计消耗</span>
              <span>¥{{ account.totalSpent.toLocaleString() }}</span>
            </div>
            <button class="w-full mt-2 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">充值</button>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">快捷操作</h3>
          <div class="space-y-2">
            <button class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">修改信息</button>
            <button class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">资金流水</button>
            <button class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">发票管理</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const account = ref({
  id: 'STAR001',
  name: '品牌星图账户',
  type: '品牌方',
  companyName: '某某科技有限公司',
  license: '91110000XXXXXXXX',
  contact: '张经理',
  phone: '138****8888',
  balance: 256800,
  frozen: 45000,
  totalSpent: 1568000
})
</script>
