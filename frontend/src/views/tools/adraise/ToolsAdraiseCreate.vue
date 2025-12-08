<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  adIds: '',
  targetType: 'cost',
  targetValue: '',
  duration: '3',
  autoStop: true
})

const targetTypes = [
  { value: 'cost', label: '目标消耗', desc: '设置期望达到的消耗金额' },
  { value: 'convert', label: '目标转化', desc: '设置期望达到的转化数' },
  { value: 'impression', label: '目标曝光', desc: '设置期望达到的曝光量' }
]

const durations = [
  { value: '1', label: '1天' },
  { value: '3', label: '3天' },
  { value: '7', label: '7天' },
  { value: 'custom', label: '自定义' }
]

const handleCancel = () => {
  alert('取消创建')
}

const handleCreate = () => {
  alert(`创建任务: ${formData.value.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '起量工具' }, { name: '创建起量任务' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建起量任务</h1>
      <p class="mt-2 text-gray-600">为广告创建自动起量优化任务</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">任务名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入任务名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">广告ID</label>
          <textarea v-model="formData.adIds" rows="3" placeholder="请输入广告ID，多个ID用逗号或换行分隔"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
          <p class="text-xs text-gray-500 mt-1">最多支持50个广告同时起量</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">起量目标</label>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div v-for="type in targetTypes" :key="type.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                          formData.targetType === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.targetType = type.value">
              <h4 class="font-medium text-gray-900">{{ type.label }}</h4>
              <p class="text-sm text-gray-500 mt-1">{{ type.desc }}</p>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">目标值</label>
          <div class="relative">
            <input v-model="formData.targetValue" type="number" placeholder="请输入目标值"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            <span class="absolute right-4 top-2 text-gray-500">
              {{ formData.targetType === 'cost' ? '元' : formData.targetType === 'convert' ? '次' : '次' }}
            </span>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">起量周期</label>
          <div class="flex gap-3">
            <button v-for="d in durations" :key="d.value"
                    :class="['px-4 py-2 rounded-lg border-2',
                             formData.duration === d.value ? 'border-blue-500 bg-blue-50 text-blue-700' : 'border-gray-200']"
                    @click="formData.duration = d.value">
              {{ d.label }}
            </button>
          </div>
        </div>

        <div>
          <label class="flex items-center gap-2 cursor-pointer">
            <input v-model="formData.autoStop" type="checkbox" class="rounded text-blue-600">
            <span class="text-sm text-gray-700">达到目标后自动停止</span>
          </label>
        </div>

        <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
          <p class="text-sm text-yellow-800">
            <span class="font-medium">提示：</span>
            起量过程中系统会自动调整出价和预算，可能产生额外消耗，请确认账户余额充足。
          </p>
        </div>

<div class="flex justify-end gap-4 pt-4 border-t">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建任务</button>
        </div>
      </div>
    </div>
  </div>
</template>
