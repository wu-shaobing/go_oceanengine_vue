<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  type: 'discount',
  value: '',
  minAmount: '',
  quantity: '',
  startDate: '',
  endDate: '',
  description: ''
})

const couponTypes = [
  { value: 'discount', label: '满减券', icon: '💰' },
  { value: 'percent', label: '折扣券', icon: '📊' },
  { value: 'gift', label: '赠品券', icon: '🎁' }
]
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '创建优惠券' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建优惠券</h1>
      <p class="mt-2 text-gray-600">创建广告投放使用的优惠券</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">优惠券名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入优惠券名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">优惠券类型</label>
          <div class="grid grid-cols-3 gap-4">
            <div v-for="type in couponTypes" :key="type.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer text-center transition-all',
                          formData.type === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.type = type.value">
              <span class="text-3xl">{{ type.icon }}</span>
              <p class="mt-2 text-sm font-medium text-gray-900">{{ type.label }}</p>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              {{ formData.type === 'percent' ? '折扣比例' : '优惠金额' }}
            </label>
            <div class="relative">
              <input v-model="formData.value" type="number" placeholder="请输入"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <span class="absolute right-4 top-2 text-gray-500">
                {{ formData.type === 'percent' ? '%' : '元' }}
              </span>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">最低消费</label>
            <div class="relative">
              <input v-model="formData.minAmount" type="number" placeholder="请输入"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <span class="absolute right-4 top-2 text-gray-500">元</span>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">发放数量</label>
          <input v-model="formData.quantity" type="number" placeholder="请输入发放数量"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">生效时间</label>
            <input v-model="formData.startDate" type="datetime-local"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">失效时间</label>
            <input v-model="formData.endDate" type="datetime-local"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">使用说明</label>
          <textarea v-model="formData.description" rows="3" placeholder="请输入优惠券使用说明"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
        </div>

        <div class="flex justify-end gap-4 pt-4 border-t">
          <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建优惠券</button>
        </div>
      </div>
    </div>
  </div>
</template>
