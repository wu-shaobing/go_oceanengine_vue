<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  strategy: 'bid',
  endpoint: '',
  timeout: 100,
  enabled: true
})

const strategies = [
  { value: 'bid', label: '出价调整', desc: '实时调整广告出价' },
  { value: 'filter', label: '流量过滤', desc: '过滤低质量流量' },
  { value: 'creative', label: '创意优选', desc: '实时选择最优创意' }
]

const handleCancel = () => {
  alert('取消创建')
}

const handleCreate = () => {
  alert(`创建RTA策略: ${formData.value.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'RTA管理' }, { name: '创建RTA策略' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建RTA策略</h1>
      <p class="mt-2 text-gray-600">配置实时竞价接口</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">策略名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入策略名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">策略类型</label>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div v-for="item in strategies" :key="item.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                          formData.strategy === item.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.strategy = item.value">
              <h4 class="font-medium text-gray-900">{{ item.label }}</h4>
              <p class="text-sm text-gray-500 mt-1">{{ item.desc }}</p>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">接口地址</label>
          <input v-model="formData.endpoint" type="url" placeholder="https://your-rta-endpoint.com/api"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <p class="text-xs text-gray-500 mt-1">RTA服务接口URL，需支持HTTPS</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">超时时间 (ms)</label>
          <input v-model="formData.timeout" type="number" min="50" max="500"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <p class="text-xs text-gray-500 mt-1">建议设置100ms以内，超时将使用默认策略</p>
        </div>

        <div>
          <label class="flex items-center gap-2 cursor-pointer">
            <input v-model="formData.enabled" type="checkbox" class="rounded text-blue-600">
            <span class="text-sm text-gray-700">创建后立即启用</span>
          </label>
        </div>

        <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
          <p class="text-sm text-yellow-800">
            <span class="font-medium">提示：</span>
            创建RTA策略前，请确保您的服务接口已完成对接并通过测试。
          </p>
        </div>

<div class="flex justify-end gap-4 pt-4 border-t">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建策略</button>
        </div>
      </div>
    </div>
  </div>
</template>
