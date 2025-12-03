<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  type: 'single',
  layout: 'vertical',
  showPrice: true,
  showDiscount: true,
  buttonText: '立即购买'
})

const templateTypes = [
  { value: 'single', label: '单品模板', icon: '🎯', desc: '展示单个商品' },
  { value: 'multi', label: '多品模板', icon: '🎪', desc: '展示多个商品' },
  { value: 'carousel', label: '轮播模板', icon: '🎠', desc: '滑动展示商品' }
]

const layouts = [
  { value: 'vertical', label: '竖版', ratio: '9:16' },
  { value: 'horizontal', label: '横版', ratio: '16:9' },
  { value: 'square', label: '方形', ratio: '1:1' }
]
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品广告' }, { name: '创建模板' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建DPA模板</h1>
      <p class="mt-2 text-gray-600">创建商品动态创意模板</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">模板名称</label>
            <input v-model="formData.name" type="text" placeholder="请输入模板名称"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">模板类型</label>
            <div class="grid grid-cols-3 gap-4">
              <div v-for="type in templateTypes" :key="type.value"
                   :class="['p-4 border-2 rounded-lg cursor-pointer text-center transition-all',
                            formData.type === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                   @click="formData.type = type.value">
                <span class="text-3xl">{{ type.icon }}</span>
                <p class="mt-2 font-medium text-gray-900">{{ type.label }}</p>
                <p class="text-xs text-gray-500">{{ type.desc }}</p>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">布局方向</label>
            <div class="flex gap-4">
              <div v-for="layout in layouts" :key="layout.value"
                   :class="['px-4 py-3 border-2 rounded-lg cursor-pointer',
                            formData.layout === layout.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200']"
                   @click="formData.layout = layout.value">
                <span class="font-medium text-gray-900">{{ layout.label }}</span>
                <span class="text-xs text-gray-500 ml-2">{{ layout.ratio }}</span>
              </div>
            </div>
          </div>

          <div class="space-y-3">
            <label class="block text-sm font-medium text-gray-700">展示元素</label>
            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="formData.showPrice" type="checkbox" class="rounded text-blue-600">
              <span class="text-sm text-gray-700">显示价格</span>
            </label>
            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="formData.showDiscount" type="checkbox" class="rounded text-blue-600">
              <span class="text-sm text-gray-700">显示折扣标签</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">按钮文案</label>
            <input v-model="formData.buttonText" type="text"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>

          <div class="flex justify-end gap-4 pt-4 border-t">
            <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
            <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建模板</button>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">模板预览</h4>
        <div class="bg-gray-100 rounded-lg aspect-[9/16] flex items-center justify-center">
          <div class="text-center text-gray-400">
            <span class="text-4xl">📱</span>
            <p class="mt-2 text-sm">模板预览区域</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
