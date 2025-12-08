<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const showPasswordModal = ref(false)
const passwordForm = ref({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const handleChangePassword = () => {
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    alert('两次输入的密码不一致')
    return
  }
  // TODO: 调用API修改密码
  alert('密码修改成功')
  showPasswordModal.value = false
  passwordForm.value = { oldPassword: '', newPassword: '', confirmPassword: '' }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-2xl mx-auto px-4">
      <h1 class="text-2xl font-bold text-gray-900 mb-6">账户信息</h1>
      
      <div class="bg-white rounded-lg shadow divide-y">
        <!-- 账户概览 -->
        <div class="p-6">
          <h2 class="text-lg font-medium text-gray-900 mb-4">账户概览</h2>
          <div class="flex items-center gap-4">
            <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold text-gray-900">{{ authStore.userInfo?.nickname || authStore.userInfo?.username || '用户' }}</h3>
              <p class="text-sm text-gray-500">{{ authStore.userInfo?.email || '未设置邮箱' }}</p>
            </div>
          </div>
        </div>

        <!-- 账户详情 -->
        <div class="p-6">
          <h2 class="text-lg font-medium text-gray-900 mb-4">账户详情</h2>
          <dl class="space-y-4">
            <div class="flex justify-between">
              <dt class="text-sm text-gray-500">用户名</dt>
              <dd class="text-sm text-gray-900">{{ authStore.userInfo?.username || '-' }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-sm text-gray-500">角色</dt>
              <dd class="text-sm text-gray-900">{{ authStore.userInfo?.role || '普通用户' }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-sm text-gray-500">账户ID</dt>
              <dd class="text-sm text-gray-900">{{ authStore.userInfo?.id || '-' }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-sm text-gray-500">注册时间</dt>
              <dd class="text-sm text-gray-900">{{ (authStore.userInfo as any)?.created_at || '-' }}</dd>
            </div>
          </dl>
        </div>

        <!-- 安全设置 -->
        <div class="p-6">
          <h2 class="text-lg font-medium text-gray-900 mb-4">安全设置</h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-sm font-medium text-gray-900">登录密码</h3>
                <p class="text-sm text-gray-500">定期更换密码可以提高账户安全性</p>
              </div>
              <button
                @click="showPasswordModal = true"
                class="px-4 py-2 text-sm text-blue-600 hover:bg-blue-50 rounded-md"
              >
                修改密码
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 修改密码弹窗 -->
    <div v-if="showPasswordModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-md mx-4 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">修改密码</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">当前密码</label>
            <input
              v-model="passwordForm.oldPassword"
              type="password"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">新密码</label>
            <input
              v-model="passwordForm.newPassword"
              type="password"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">确认新密码</label>
            <input
              v-model="passwordForm.confirmPassword"
              type="password"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>
        <div class="flex gap-3 mt-6">
          <button
            @click="showPasswordModal = false"
            class="flex-1 py-2 px-4 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50"
          >
            取消
          </button>
          <button
            @click="handleChangePassword"
            class="flex-1 py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            确认修改
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
