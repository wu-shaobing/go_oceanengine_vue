<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  catalogId: '',
  uploadType: 'file'
})

const uploadHistory = ref([
  { id: 'U001', fileName: 'products_20251125.csv', count: 5000, status: 'success', time: '2025-11-25 10:30' },
  { id: 'U002', fileName: 'products_20251126.csv', count: 3200, status: 'success', time: '2025-11-26 14:20' },
  { id: 'U003', fileName: 'products_20251127.csv', count: 1500, status: 'processing', time: '2025-11-27 09:15' }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: '商品上传' }]" />
      <h1 class="text-3xl font-bold text-gray-900">商品上传</h1>
      <p class="mt-2 text-gray-600">批量上传商品到DPA目录</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">选择商品目录</label>
            <select v-model="formData.catalogId" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="">请选择商品目录</option>
              <option value="CAT001">主商品目录</option>
              <option value="CAT002">促销商品目录</option>
              <option value="CAT003">测试目录</option>
            </select>
          </div>

          <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center hover:border-blue-400 transition-colors cursor-pointer">
            <div class="text-5xl mb-4">📦</div>
            <p class="text-lg font-medium text-gray-700">拖拽文件到此处上传</p>
            <p class="text-sm text-gray-500 mt-2">或点击选择文件</p>
            <p class="text-xs text-gray-400 mt-4">支持 CSV、XML 格式，最大 50MB</p>
            <button class="mt-6 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">选择文件</button>
          </div>

          <div class="flex items-center gap-4 p-4 bg-blue-50 rounded-lg">
            <span class="text-2xl">📄</span>
            <div>
              <p class="text-sm font-medium text-blue-800">下载模板文件</p>
              <p class="text-xs text-blue-600">使用标准模板可确保数据格式正确</p>
            </div>
            <button class="ml-auto px-4 py-1 text-sm text-blue-600 border border-blue-300 rounded hover:bg-blue-100">下载模板</button>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">上传记录</h4>
        <div class="space-y-3">
          <div v-for="record in uploadHistory" :key="record.id" class="p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">{{ record.fileName }}</span>
              <span :class="['px-2 py-0.5 rounded text-xs',
                     record.status === 'success' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700']">
                {{ record.status === 'success' ? '成功' : '处理中' }}
              </span>
            </div>
            <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
              <span>{{ record.count.toLocaleString() }} 条</span>
              <span>{{ record.time }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
