<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  amount: '',
  remark: ''
})

const recentRecords = ref([
  { id: 1, advertiser: '北京科技有限公司', amount: 100000, time: '2025-11-11 14:30', status: '成功' },
  { id: 2, advertiser: '上海网络公司', amount: 50000, time: '2025-11-11 10:15', status: '成功' },
  { id: 3, advertiser: '深圳电商公司', amount: 200000, time: '2025-11-10 16:45', status: '成功' }
])

const handleSubmit = () => {
  console.log('Submit:', form)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '账户充值' }]" />
      <h1 class="text-3xl font-bold text-gray-900">账户充值</h1>
      <p class="mt-2 text-gray-600">为广告主账户进行充值操作</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Recharge Form -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-6">充值信息</h2>
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">选择广告主</label>
            <select v-model="form.advertiserId" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">请选择广告主</option>
              <option value="1234567">北京科技有限公司 (ID: 1234567)</option>
              <option value="1234568">上海网络公司 (ID: 1234568)</option>
              <option value="1234569">深圳电商公司 (ID: 1234569)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">充值金额 (元)</label>
            <input v-model="form.amount" type="number" placeholder="请输入充值金额" 
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <p class="mt-1 text-sm text-gray-500">最小充值金额：1000元</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">备注</label>
            <textarea v-model="form.remark" rows="3" placeholder="请输入备注信息（选填）"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
          </div>
          <div class="flex gap-4">
            <button type="submit" class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
              确认充值
            </button>
            <button type="button" class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
              取消
            </button>
          </div>
        </form>
      </div>

      <!-- Recent Records -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-6">最近充值记录</h2>
        <div class="space-y-4">
          <div v-for="record in recentRecords" :key="record.id" class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
            <div>
              <p class="text-sm font-medium text-gray-900">{{ record.advertiser }}</p>
              <p class="text-xs text-gray-500 mt-1">{{ record.time }}</p>
            </div>
            <div class="text-right">
              <p class="text-lg font-semibold text-green-600">+¥{{ record.amount.toLocaleString() }}</p>
              <span class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">{{ record.status }}</span>
            </div>
          </div>
        </div>
        <router-link to="/agent/transfer" class="block mt-4 text-center text-sm text-blue-600 hover:text-blue-700">
          查看全部充值记录 →
        </router-link>
      </div>
    </div>
  </div>
</template>
