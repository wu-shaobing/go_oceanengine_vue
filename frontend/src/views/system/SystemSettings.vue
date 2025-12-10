<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import { systemApi, type UserSettingResp } from '@/api/system'

const activeTab = ref('general')
const loading = ref(false)
const saving = ref(false)

const settings = reactive<UserSettingResp>({
  language: 'zh-CN',
  timezone: 'Asia/Shanghai',
  theme: 'light',
  notifications_enabled: true,
  email_alerts_enabled: true,
  sms_alerts_enabled: false,
  auto_refresh_enabled: true,
  refresh_interval: 30
})

const tabs = [
  { key: 'general', label: '基本设置' },
  { key: 'notifications', label: '通知设置' },
  { key: 'security', label: '安全设置' },
  { key: 'privacy', label: '隐私设置' }
]

// 修改密码表单
const passwordForm = reactive({
  old_password: '',
  new_password: '',
  confirm_password: ''
})
const showPasswordDialog = ref(false)
const changingPassword = ref(false)

// 加载设置
const loadSettings = async () => {
  loading.value = true
  try {
    const res = await systemApi.getUserSettings() as UserSettingResp
    if (res) {
      Object.assign(settings, res)
    }
  } catch (error: any) {
    alert(error.message || '加载设置失败')
  } finally {
    loading.value = false
  }
}

// 保存设置
const handleSaveSettings = async () => {
  saving.value = true
  try {
    await systemApi.updateUserSettings({
      language: settings.language,
      timezone: settings.timezone,
      theme: settings.theme,
      notifications_enabled: settings.notifications_enabled,
      email_alerts_enabled: settings.email_alerts_enabled,
      sms_alerts_enabled: settings.sms_alerts_enabled,
      auto_refresh_enabled: settings.auto_refresh_enabled,
      refresh_interval: settings.refresh_interval
    })
    alert('设置已保存')
  } catch (error: any) {
    alert(error.message || '保存设置失败')
  } finally {
    saving.value = false
  }
}

// 打开修改密码对话框
const handleChangePassword = () => {
  passwordForm.old_password = ''
  passwordForm.new_password = ''
  passwordForm.confirm_password = ''
  showPasswordDialog.value = true
}

// 提交修改密码
const submitChangePassword = async () => {
  if (!passwordForm.old_password || !passwordForm.new_password || !passwordForm.confirm_password) {
    alert('请填写完整的密码信息')
    return
  }
  if (passwordForm.new_password !== passwordForm.confirm_password) {
    alert('两次输入的新密码不一致')
    return
  }
  if (passwordForm.new_password.length < 6 || passwordForm.new_password.length > 32) {
    alert('新密码长度应为6-32位')
    return
  }

  changingPassword.value = true
  try {
    await systemApi.changePassword({
      old_password: passwordForm.old_password,
      new_password: passwordForm.new_password
    })
    alert('密码修改成功')
    showPasswordDialog.value = false
  } catch (error: any) {
    alert(error.message || '修改密码失败')
  } finally {
    changingPassword.value = false
  }
}

// 启用双因素认证（示例）
const handleEnable2FA = () => {
  alert('双因素认证功能正在开发中，敬请期待')
}

onMounted(() => {
  loadSettings()
})
</script>

<template>
  <div class="space-y-6" v-loading="loading">
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
            <select v-model="settings.language" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="zh-CN">简体中文</option>
              <option value="zh-TW">繁体中文</option>
              <option value="en-US">English</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">时区</label>
            <select v-model="settings.timezone" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
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
          <div>
            <label class="flex items-center justify-between">
              <div>
                <p class="font-medium text-gray-900">自动刷新</p>
                <p class="text-sm text-gray-500">开启后数据将自动更新</p>
              </div>
              <input type="checkbox" v-model="settings.auto_refresh_enabled" class="w-5 h-5 text-blue-600 rounded">
            </label>
          </div>
          <div v-if="settings.auto_refresh_enabled">
            <label class="block text-sm font-medium text-gray-700 mb-2">刷新间隔（秒）</label>
            <input type="number" v-model.number="settings.refresh_interval" min="10" max="300" step="10"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
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
              <input type="checkbox" v-model="settings.notifications_enabled" class="w-5 h-5 text-blue-600 rounded">
            </label>
            <label class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium text-gray-900">邮件通知</p>
                <p class="text-sm text-gray-500">重要消息发送到邮箱</p>
              </div>
              <input type="checkbox" v-model="settings.email_alerts_enabled" class="w-5 h-5 text-blue-600 rounded">
            </label>
            <label class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium text-gray-900">短信通知</p>
                <p class="text-sm text-gray-500">紧急消息短信提醒</p>
              </div>
              <input type="checkbox" v-model="settings.sms_alerts_enabled" class="w-5 h-5 text-blue-600 rounded">
            </label>
          </div>
        </div>

        <div v-if="activeTab === 'security'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">安全设置</h3>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="font-medium text-gray-900">修改密码</p>
            <p class="text-sm text-gray-500 mt-1">定期更换密码以确保账号安全</p>
            <button class="mt-3 px-4 py-2 border border-gray-300 rounded-lg hover:bg-white transition-colors" @click="handleChangePassword">修改密码</button>
          </div>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="font-medium text-gray-900">双因素认证</p>
            <p class="text-sm text-gray-500 mt-1">启用后需要手机验证码登录</p>
            <button class="mt-3 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="handleEnable2FA">启用</button>
          </div>
        </div>

        <div v-if="activeTab === 'privacy'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">隐私设置</h3>
          <div class="p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
            <p class="text-sm text-yellow-800">您的数据受到严格保护，仅用于广告投放服务</p>
          </div>
        </div>

        <div class="flex justify-end pt-6 mt-6 border-t">
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
                  :disabled="saving"
                  @click="handleSaveSettings">
            {{ saving ? '保存中...' : '保存设置' }}
          </button>
        </div>
      </div>
    </div>

    <!-- 修改密码对话框 -->
    <div v-if="showPasswordDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg w-[420px] p-6">
        <h3 class="text-lg font-semibold mb-4">修改密码</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">当前密码</label>
            <input type="password" v-model="passwordForm.old_password" placeholder="请输入当前密码"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">新密码</label>
            <input type="password" v-model="passwordForm.new_password" placeholder="请输入新密码（6-32位）"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">确认新密码</label>
            <input type="password" v-model="passwordForm.confirm_password" placeholder="请再次输入新密码"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
        <div class="flex justify-end gap-3 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showPasswordDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
                  :disabled="changingPassword"
                  @click="submitChangePassword">
            {{ changingPassword ? '提交中...' : '确认修改' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
