<script setup lang="ts">
import { reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const agentInfo = reactive({
  id: '1234567890',
  name: '北京巨量引擎代理商',
  company: '北京巨量网络科技有限公司',
  type: '一级代理商',
  status: 'active',
  creditCode: '91110000MA01XXXX',
  legalPerson: '张三',
  contactName: '李四',
  contactPhone: '138****8888',
  contactEmail: 'lisi@example.com',
  address: '北京市朝阳区望京科技园A座18层',
  createdAt: '2023-01-15',
  balance: 5680000,
  creditLimit: 10000000,
  totalAdvertisers: 128,
  activeAdvertisers: 95,
  totalChildAgents: 25
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '代理商信息' }]" />
      <h1 class="text-3xl font-bold text-gray-900">代理商信息</h1>
      <p class="mt-2 text-gray-600">查看代理商详细信息</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Basic Info -->
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-lg font-semibold text-gray-900">基本信息</h2>
            <span :class="['px-3 py-1 rounded-full text-sm font-medium',
              agentInfo.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
              {{ agentInfo.status === 'active' ? '正常运营' : '已暂停' }}
            </span>
          </div>
          <div class="grid grid-cols-2 gap-6">
            <div>
              <p class="text-sm text-gray-500">代理商ID</p>
              <p class="mt-1 text-sm font-medium text-gray-900 font-mono">{{ agentInfo.id }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">代理商名称</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.name }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">公司全称</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.company }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">代理商类型</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.type }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">统一社会信用代码</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.creditCode }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">法人代表</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.legalPerson }}</p>
            </div>
            <div class="col-span-2">
              <p class="text-sm text-gray-500">公司地址</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.address }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h2 class="text-lg font-semibold text-gray-900 mb-6">联系人信息</h2>
          <div class="grid grid-cols-3 gap-6">
            <div>
              <p class="text-sm text-gray-500">联系人姓名</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.contactName }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">联系电话</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.contactPhone }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">联系邮箱</p>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ agentInfo.contactEmail }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats -->
      <div class="space-y-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
          <p class="text-sm text-blue-100">账户余额</p>
          <p class="text-3xl font-bold mt-2">¥{{ (agentInfo.balance / 10000).toFixed(0) }}万</p>
          <div class="mt-4 pt-4 border-t border-blue-400">
            <div class="flex justify-between text-sm">
              <span class="text-blue-100">授信额度</span>
              <span>¥{{ (agentInfo.creditLimit / 10000).toFixed(0) }}万</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-medium text-gray-900 mb-4">账户统计</h3>
          <div class="space-y-4">
            <div class="flex justify-between">
              <span class="text-gray-500">总广告主</span>
              <span class="font-medium text-gray-900">{{ agentInfo.totalAdvertisers }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">活跃广告主</span>
              <span class="font-medium text-green-600">{{ agentInfo.activeAdvertisers }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">子代理商</span>
              <span class="font-medium text-gray-900">{{ agentInfo.totalChildAgents }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">创建时间</span>
              <span class="font-medium text-gray-900">{{ agentInfo.createdAt }}</span>
            </div>
          </div>
        </div>

        <div class="flex flex-col gap-3">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            编辑信息
          </button>
          <button class="w-full px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
            查看资质
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
