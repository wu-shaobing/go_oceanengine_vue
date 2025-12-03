<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()

const form = reactive({
  targetId: route.query.target as string || '',
  targetType: 'advertiser',
  amount: '',
  remark: ''
})

const walletBalance = 5680000

const handleSubmit = () => {
  console.log('Submit transfer:', form)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工作台' }, { name: '发起转账' }]" />
      <h1 class="text-3xl font-bold text-gray-900">发起转账</h1>
      <p class="mt-2 text-gray-600">创建新的转账交易</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">转账类型</label>
            <div class="flex gap-4">
              <label class="flex items-center p-4 border-2 rounded-lg cursor-pointer transition-all"
                     :class="form.targetType === 'advertiser' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'">
                <input type="radio" v-model="form.targetType" value="advertiser" class="sr-only">
                <div>
                  <p class="font-medium text-gray-900">转至广告主</p>
                  <p class="text-sm text-gray-500">充值到广告主账户</p>
                </div>
              </label>
              <label class="flex items-center p-4 border-2 rounded-lg cursor-pointer transition-all"
                     :class="form.targetType === 'agent' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'">
                <input type="radio" v-model="form.targetType" value="agent" class="sr-only">
                <div>
                  <p class="font-medium text-gray-900">转至子代理</p>
                  <p class="text-sm text-gray-500">划转到子代理账户</p>
                </div>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">目标账户ID</label>
            <input v-model="form.targetId" type="text" placeholder="请输入目标账户ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">转账金额 (元)</label>
            <input v-model="form.amount" type="number" placeholder="请输入转账金额"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            <p class="mt-1 text-sm text-gray-500">可转余额: ¥{{ walletBalance.toLocaleString() }}</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">备注</label>
            <textarea v-model="form.remark" rows="3" placeholder="请输入备注信息（选填）"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          
          <div class="flex gap-4 pt-4">
            <button type="submit" class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
              确认转账
            </button>
            <router-link to="/workspace/transfer-targets" 
                         class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors text-center">
              取消
            </router-link>
          </div>
        </form>
      </div>

      <div class="space-y-4">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
          <p class="text-sm text-blue-100">钱包余额</p>
          <p class="text-3xl font-bold mt-2">¥{{ (walletBalance / 10000).toFixed(0) }}万</p>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-medium text-gray-900 mb-3">转账须知</h3>
          <ul class="space-y-2 text-sm text-gray-600">
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              转账实时到账，请核实目标账户
            </li>
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              单笔转账最低100元
            </li>
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              转账记录可在流水明细中查看
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
