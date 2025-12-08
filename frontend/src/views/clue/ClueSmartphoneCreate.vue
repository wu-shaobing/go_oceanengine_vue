<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const loading = ref(false)

const formData = ref({
  name: '',
  phoneNumber: '',
  welcomeText: '',
  workTime: 'all',
  customStartTime: '09:00',
  customEndTime: '18:00',
  recordCall: true
})

const workTimes = [
  { value: 'all', label: '全天候', desc: '24小时接听' },
  { value: 'business', label: '工作时间', desc: '9:00-18:00' },
  { value: 'custom', label: '自定义', desc: '设置特定时段' }
]

const validateForm = () => {
  if (!formData.value.name.trim()) {
    alert('请输入配置名称')
    return false
  }
  if (!formData.value.phoneNumber.trim()) {
    alert('请输入接听号码')
    return false
  }
  return true
}

const handleCancel = () => {
  router.push('/clue/smartphone/list')
}

const handleCreate = async () => {
  if (!validateForm()) return
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert('智能电话配置创建成功')
    router.push('/clue/smartphone/list')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '创建智能电话' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建智能电话</h1>
      <p class="mt-2 text-gray-600">配置广告智能电话拨打功能</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">配置名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入配置名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">接听号码</label>
          <input v-model="formData.phoneNumber" type="tel" placeholder="请输入接听电话号码"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <p class="text-xs text-gray-500 mt-1">用户点击拨打后将接入此号码</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">欢迎语</label>
          <textarea v-model="formData.welcomeText" rows="3" placeholder="请输入通话欢迎语（选填）"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">接听时段</label>
          <div class="space-y-3">
            <div v-for="time in workTimes" :key="time.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                          formData.workTime === time.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.workTime = time.value">
              <div class="flex items-center gap-3">
                <div :class="['w-4 h-4 rounded-full border-2',
                             formData.workTime === time.value ? 'border-blue-500 bg-blue-500' : 'border-gray-300']"></div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ time.label }}</h4>
                  <p class="text-sm text-gray-500">{{ time.desc }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="formData.workTime === 'custom'" class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-gray-600 mb-1">开始时间</label>
            <input v-model="formData.customStartTime" type="time"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg">
          </div>
          <div>
            <label class="block text-sm text-gray-600 mb-1">结束时间</label>
            <input v-model="formData.customEndTime" type="time"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg">
          </div>
        </div>

        <div>
          <label class="flex items-center gap-2 cursor-pointer">
            <input v-model="formData.recordCall" type="checkbox" class="rounded text-blue-600">
            <span class="text-sm text-gray-700">启用通话录音</span>
          </label>
        </div>

        <div class="flex justify-end gap-4 pt-4 border-t">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" :disabled="loading" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50">{{ loading ? '创建中...' : '创建配置' }}</button>
        </div>
      </div>
    </div>
  </div>
</template>
