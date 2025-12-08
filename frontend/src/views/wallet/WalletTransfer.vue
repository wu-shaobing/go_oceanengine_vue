<script setup lang="ts">
import { reactive, ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  targetType: 'advertiser',
  targetId: '',
  amount: '',
  remark: ''
})

const walletInfo = reactive({
  balance: 5680000,
  available: 5500000
})

const loading = ref(false)

const validateForm = () => {
  if (!form.targetId) {
    alert('请选择目标账户')
    return false
  }
  const amount = parseFloat(form.amount)
  if (!amount || amount <= 0) {
    alert('请输入有效的转账金额')
    return false
  }
  if (amount > walletInfo.available) {
    alert('转账金额不能超过可转金额')
    return false
  }
  return true
}

const handleSubmit = async () => {
  if (!validateForm()) return
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert(`转账成功！金额: ¥${parseFloat(form.amount).toLocaleString()}`)
    form.targetId = ''
    form.amount = ''
    form.remark = ''
  } finally {
    loading.value = false
  }
}

const handleCancel = () => {
  form.targetId = ''
  form.amount = ''
  form.remark = ''
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '财务管理' }, { name: '钱包转账' }]" />
      <h1 class="text-3xl font-bold text-gray-900">钱包转账</h1>
      <p class="mt-2 text-gray-600">从共享钱包转账到广告主账户</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Transfer Form -->
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-6">转账信息</h2>
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">转账类型</label>
            <div class="flex gap-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.targetType" value="advertiser" class="mr-2">
                <span>转至广告主</span>
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.targetType" value="agent" class="mr-2">
                <span>转至子代理</span>
              </label>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">选择目标账户</label>
            <select v-model="form.targetId" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="">请选择账户</option>
              <option value="1234567">北京科技有限公司 (ID: 1234567)</option>
              <option value="1234568">上海网络公司 (ID: 1234568)</option>
              <option value="1234569">深圳电商公司 (ID: 1234569)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">转账金额 (元)</label>
            <input v-model="form.amount" type="number" placeholder="请输入转账金额" 
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            <p class="mt-1 text-sm text-gray-500">可转金额: ¥{{ walletInfo.available.toLocaleString() }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">备注</label>
            <textarea v-model="form.remark" rows="3" placeholder="请输入备注信息（选填）"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          <div class="flex gap-4">
            <button type="submit" :disabled="loading" class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50">
              {{ loading ? '处理中...' : '确认转账' }}
            </button>
            <button type="button" @click="handleCancel" class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
              取消
            </button>
          </div>
        </form>
      </div>

      <!-- Wallet Info -->
      <div class="space-y-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
          <p class="text-sm text-blue-100">钱包余额</p>
          <p class="text-3xl font-bold mt-2">¥{{ (walletInfo.balance / 10000).toFixed(0) }}万</p>
          <div class="mt-4 pt-4 border-t border-blue-400">
            <div class="flex justify-between text-sm">
              <span class="text-blue-100">可转金额</span>
              <span>¥{{ (walletInfo.available / 10000).toFixed(0) }}万</span>
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-medium text-gray-900 mb-4">转账须知</h3>
          <ul class="space-y-2 text-sm text-gray-600">
            <li class="flex items-start gap-2">
              <svg class="w-5 h-5 text-blue-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              转账金额需大于0且不超过可转金额
            </li>
            <li class="flex items-start gap-2">
              <svg class="w-5 h-5 text-blue-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              转账实时到账，请核实目标账户
            </li>
            <li class="flex items-start gap-2">
              <svg class="w-5 h-5 text-blue-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              转账记录可在流水明细中查看
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
