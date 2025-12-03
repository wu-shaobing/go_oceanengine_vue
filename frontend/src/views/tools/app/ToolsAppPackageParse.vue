<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  packageUrl: '',
  appName: ''
})

const parseResult = ref<null | {
  packageName: string
  versionName: string
  versionCode: number
  minSdk: number
  targetSdk: number
  fileSize: string
  permissions: string[]
}>(null)

const loading = ref(false)

const handleParse = () => {
  loading.value = true
  setTimeout(() => {
    parseResult.value = {
      packageName: 'com.example.app',
      versionName: '2.5.0',
      versionCode: 250,
      minSdk: 21,
      targetSdk: 33,
      fileSize: '45.6 MB',
      permissions: ['网络访问', '存储读写', '相机', '位置信息', '通知']
    }
    loading.value = false
  }, 1500)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: '包解析' }]" />
      <h1 class="text-3xl font-bold text-gray-900">应用包解析</h1>
      <p class="mt-2 text-gray-600">提交解析应用包任务，获取应用信息</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">提交解析任务</h3>
        <form @submit.prevent="handleParse" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">广告主ID</label>
            <input v-model="form.advertiserId" type="text" placeholder="请输入广告主ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">应用包URL</label>
            <input v-model="form.packageUrl" type="url" placeholder="https://example.com/app.apk"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">应用名称</label>
            <input v-model="form.appName" type="text" placeholder="请输入应用名称"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <button type="submit" :disabled="loading"
                  class="w-full px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:bg-gray-400">
            {{ loading ? '解析中...' : '开始解析' }}
          </button>
        </form>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">解析结果</h3>
        <div v-if="parseResult" class="space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="p-3 bg-gray-50 rounded-lg">
              <p class="text-xs text-gray-500">包名</p>
              <p class="text-sm font-medium text-gray-900 font-mono">{{ parseResult.packageName }}</p>
            </div>
            <div class="p-3 bg-gray-50 rounded-lg">
              <p class="text-xs text-gray-500">版本号</p>
              <p class="text-sm font-medium text-gray-900">{{ parseResult.versionName }} ({{ parseResult.versionCode }})</p>
            </div>
            <div class="p-3 bg-gray-50 rounded-lg">
              <p class="text-xs text-gray-500">最低SDK</p>
              <p class="text-sm font-medium text-gray-900">API {{ parseResult.minSdk }}</p>
            </div>
            <div class="p-3 bg-gray-50 rounded-lg">
              <p class="text-xs text-gray-500">目标SDK</p>
              <p class="text-sm font-medium text-gray-900">API {{ parseResult.targetSdk }}</p>
            </div>
            <div class="p-3 bg-gray-50 rounded-lg col-span-2">
              <p class="text-xs text-gray-500">文件大小</p>
              <p class="text-sm font-medium text-gray-900">{{ parseResult.fileSize }}</p>
            </div>
          </div>
          <div>
            <p class="text-xs text-gray-500 mb-2">权限列表</p>
            <div class="flex flex-wrap gap-2">
              <span v-for="perm in parseResult.permissions" :key="perm"
                    class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">
                {{ perm }}
              </span>
            </div>
          </div>
        </div>
        <div v-else class="text-center py-12 text-gray-500">
          提交应用包后显示解析结果
        </div>
      </div>
    </div>
  </div>
</template>
