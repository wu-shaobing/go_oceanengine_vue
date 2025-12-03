<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const activeTab = ref('general')

const settings = ref({
  language: 'zh-CN',
  timezone: 'Asia/Shanghai',
  theme: 'light',
  notifications: true,
  emailAlerts: true,
  smsAlerts: false,
  autoRefresh: true,
  refreshInterval: 30
})

const tabs = [
  { key: 'general', label: '基本设置' },
  { key: 'notifications', label: '通知设置' },
  { key: 'security', label: '安全设置' },
  { key: 'privacy', label: '隐私设置' }
]
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统' }, { name: '系统设置' }]" />
      <h1 class="text-3xl font-bold text-gray-900">系统设置</h1>
      <p class="mt-2 text-gray-600">配置系统偏好和个性化选项</p>
    </div>

    <div class="flex gap-6">
      <div class="w-48 space-y-1">
        <button v-for="tab in tabs" :key="tab.key"
                :class="['w-full px-4 py-2 text-left rounded-lg transition-colors',
                         activeTab === tab.key ? 'bg-blue-50 text-blue-700 font-medium' : 'text-gray-600 hover:bg-gray-100']"
                @click="activeTab = tab.key">
          {{ tab.label }}
        </button>
      </div>

      <div class="flex-1 bg-white rounded-lg border border-gray-200 p-6">
        <div v-if="activeTab === 'general'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">基本设置</h3>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">语言</label>
            <select v-model="settings.language" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="zh-CN">简体中文</option>
              <option value="zh-TW">繁体中文</option>
              <option value="en-US">English</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">时区</label>
            <select v-model="settings.timezone" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="Asia/Shanghai">中国标准时间 (UTC+8)</option>
              <option value="Asia/Hong_Kong">香港时间</option>
              <option value="UTC">UTC</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">主题</label>
            <div class="flex gap-4">
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="radio" v-model="settings.theme" value="light" class="text-blue-600">
                <span class="text-sm">浅色</span>
              </label>
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="radio" v-model="settings.theme" value="dark" class="text-blue-600">
                <span class="text-sm">深色</span>
              </label>
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="radio" v-model="settings.theme" value="auto" class="text-blue-600">
                <span class="text-sm">跟随系统</span>
              </label>
            </div>
          </div>
        </div>

        <div v-if="activeTab === 'notifications'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">通知设置</h3>
          <div class="space-y-4">
            <label class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium text-gray-900">系统通知</p>
                <p class="text-sm text-gray-500">接收系统消息和更新通知</p>
              </div>
              <input type="checkbox" v-model="settings.notifications" class="toggle">
            </label>
            <label class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium text-gray-900">邮件通知</p>
                <p class="text-sm text-gray-500">重要消息发送到邮箱</p>
              </div>
              <input type="checkbox" v-model="settings.emailAlerts" class="toggle">
            </label>
            <label class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium text-gray-900">短信通知</p>
                <p class="text-sm text-gray-500">紧急消息短信提醒</p>
              </div>
              <input type="checkbox" v-model="settings.smsAlerts" class="toggle">
            </label>
          </div>
        </div>

        <div v-if="activeTab === 'security'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">安全设置</h3>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="font-medium text-gray-900">修改密码</p>
            <p class="text-sm text-gray-500 mt-1">定期更换密码以确保账号安全</p>
            <button class="mt-3 px-4 py-2 border border-gray-300 rounded-lg hover:bg-white">修改密码</button>
          </div>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="font-medium text-gray-900">双因素认证</p>
            <p class="text-sm text-gray-500 mt-1">启用后需要手机验证码登录</p>
            <button class="mt-3 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">启用</button>
          </div>
        </div>

        <div v-if="activeTab === 'privacy'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">隐私设置</h3>
          <div class="p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
            <p class="text-sm text-yellow-800">您的数据受到严格保护，仅用于广告投放服务</p>
          </div>
        </div>

        <div class="flex justify-end pt-6 mt-6 border-t">
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">保存设置</button>
        </div>
      </div>
    </div>
  </div>
</template>
