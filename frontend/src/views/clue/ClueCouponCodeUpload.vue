<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  couponId: '',
  uploadType: 'file'
})

const uploadHistory = ref([
  { id: 'U001', fileName: 'codes_batch1.csv', count: 5000, uploadedAt: '2025-11-25 10:30', status: 'success' },
  { id: 'U002', fileName: 'codes_batch2.csv', count: 3000, uploadedAt: '2025-11-26 14:20', status: 'success' },
  { id: 'U003', fileName: 'codes_batch3.csv', count: 2000, uploadedAt: '2025-11-27 09:15', status: 'processing' }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '上传优惠码' }]" />
      <h1 class="text-3xl font-bold text-gray-900">上传优惠码</h1>
      <p class="mt-2 text-gray-600">批量上传优惠券兑换码</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">选择优惠券</label>
            <select v-model="formData.couponId" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="">请选择优惠券</option>
              <option value="CP001">双11满减券 - 满100减20</option>
              <option value="CP002">新客专享券 - 满50减10</option>
              <option value="CP003">会员折扣券 - 9折</option>
            </select>
          </div>

          <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center hover:border-blue-400 transition-colors cursor-pointer">
            <div class="text-5xl mb-4">📤</div>
            <p class="text-lg font-medium text-gray-700">拖拽文件到此处上传</p>
            <p class="text-sm text-gray-500 mt-2">或点击选择文件</p>
            <p class="text-xs text-gray-400 mt-4">支持 CSV、TXT 格式，每行一个优惠码</p>
            <button class="mt-6 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
              选择文件
            </button>
          </div>

          <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
            <h4 class="font-medium text-blue-800 mb-2">文件格式要求</h4>
            <ul class="text-sm text-blue-700 space-y-1">
              <li>• 每行一个优惠码，不含标题行</li>
              <li>• 优惠码长度 6-20 位</li>
              <li>• 仅支持字母、数字</li>
              <li>• 单次最多上传 10 万条</li>
            </ul>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">上传记录</h4>
        <div class="space-y-3">
          <div v-for="record in uploadHistory" :key="record.id"
               class="p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">{{ record.fileName }}</span>
              <span :class="['px-2 py-0.5 rounded text-xs',
                     record.status === 'success' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700']">
                {{ record.status === 'success' ? '成功' : '处理中' }}
              </span>
            </div>
            <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
              <span>{{ record.count.toLocaleString() }} 条</span>
              <span>{{ record.uploadedAt }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
