<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const form = ref({
  nickname: authStore.userInfo?.nickname || '',
  email: authStore.userInfo?.email || '',
  phone: '',
  notification: true,
  language: 'zh-CN'
})

const saving = ref(false)

const handleSave = async () => {
  saving.value = true
  // TODO: 调用API保存设置
  setTimeout(() => {
    saving.value = false
    alert('设置已保存')
  }, 1000)
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-2xl mx-auto px-4">
      <h1 class="text-2xl font-bold text-gray-900 mb-6">个人设置</h1>
      
      <div class="bg-white rounded-lg shadow p-6 space-y-6">
        <!-- 基本信息 -->
        <div>
          <h2 class="text-lg font-medium text-gray-900 mb-4">基本信息</h2>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">昵称</label>
              <input
                v-model="form.nickname"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
              <input
                v-model="form.email"
                type="email"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">手机号</label>
              <input
                v-model="form.phone"
                type="tel"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
          </div>
        </div>

        <!-- 偏好设置 -->
        <div>
          <h2 class="text-lg font-medium text-gray-900 mb-4">偏好设置</h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">接收通知</span>
              <button
                @click="form.notification = !form.notification"
                :class="[
                  'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                  form.notification ? 'bg-blue-600' : 'bg-gray-200'
                ]"
              >
                <span
                  :class="[
                    'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                    form.notification ? 'translate-x-6' : 'translate-x-1'
                  ]"
                />
              </button>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">语言</label>
              <select
                v-model="form.language"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="zh-CN">简体中文</option>
                <option value="en-US">English</option>
              </select>
            </div>
          </div>
        </div>

        <!-- 保存按钮 -->
        <div class="pt-4">
          <button
            @click="handleSave"
            :disabled="saving"
            class="w-full py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
          >
            {{ saving ? '保存中...' : '保存设置' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
