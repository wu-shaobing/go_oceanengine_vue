<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  type: 'form',
  source: 'landing_page',
  deepConvert: false,
  optimizationType: 'convert',
  convertPrice: ''
})

const convertTypes = [
  { value: 'form', label: '表单提交', icon: '📝' },
  { value: 'button', label: '按钮点击', icon: '🔘' },
  { value: 'download', label: '下载安装', icon: '📥' },
  { value: 'payment', label: '付款成功', icon: '💰' },
  { value: 'register', label: '注册成功', icon: '👤' }
]

const handleCancel = () => {
  alert('取消创建')
}

const handleCreate = () => {
  alert(`创建转化: ${formData.value.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '转化追踪' }, { name: '创建转化' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建转化</h1>
      <p class="mt-2 text-gray-600">设置广告转化事件追踪</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">转化名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入转化名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">转化类型</label>
          <div class="grid grid-cols-2 md:grid-cols-5 gap-3">
            <div v-for="type in convertTypes" :key="type.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer text-center transition-all',
                          formData.type === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.type = type.value">
              <span class="text-2xl">{{ type.icon }}</span>
              <p class="mt-1 text-sm font-medium text-gray-900">{{ type.label }}</p>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">转化来源</label>
          <select v-model="formData.source" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            <option value="landing_page">落地页</option>
            <option value="app">应用内</option>
            <option value="mini_program">小程序</option>
            <option value="api">API上报</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">优化目标</label>
          <div class="flex gap-4">
            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="formData.optimizationType" type="radio" value="convert" class="text-blue-600">
              <span class="text-sm text-gray-700">转化量</span>
            </label>
            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="formData.optimizationType" type="radio" value="price" class="text-blue-600">
              <span class="text-sm text-gray-700">转化成本</span>
            </label>
          </div>
        </div>

        <div>
          <label class="flex items-center gap-2 cursor-pointer">
            <input v-model="formData.deepConvert" type="checkbox" class="rounded text-blue-600">
            <span class="text-sm text-gray-700">启用深度转化优化</span>
          </label>
        </div>

<div class="flex justify-end gap-4 pt-4 border-t">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建转化</button>
        </div>
      </div>
    </div>
  </div>
</template>
