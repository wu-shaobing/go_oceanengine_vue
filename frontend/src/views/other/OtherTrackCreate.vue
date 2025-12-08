<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  targetUrl: '',
  trackType: 'click',
  parameters: ''
})

const trackTypes = [
  { value: 'click', label: '点击追踪', desc: '追踪广告点击行为' },
  { value: 'impression', label: '曝光追踪', desc: '追踪广告展示行为' },
  { value: 'conversion', label: '转化追踪', desc: '追踪用户转化行为' }
]

const handleCancel = () => {
  alert('取消')
}

const handleCreateLink = () => {
  alert('创建链接')
}

const handleCopyPreviewLink = () => {
  alert('复制链接')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他' }, { name: '创建追踪链接' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建追踪链接</h1>
      <p class="mt-2 text-gray-600">创建新的广告追踪链接</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">链接名称</label>
            <input v-model="formData.name" type="text" placeholder="请输入追踪链接名称"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">目标URL</label>
            <input v-model="formData.targetUrl" type="url" placeholder="https://example.com/landing"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            <p class="text-xs text-gray-500 mt-1">用户点击后将跳转到此地址</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">追踪类型</label>
            <div class="space-y-3">
              <div v-for="type in trackTypes" :key="type.value"
                   :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                            formData.trackType === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                   @click="formData.trackType = type.value">
                <div class="flex items-center gap-3">
                  <div :class="['w-4 h-4 rounded-full border-2',
                               formData.trackType === type.value ? 'border-blue-500 bg-blue-500' : 'border-gray-300']"></div>
                  <div>
                    <h4 class="font-medium text-gray-900">{{ type.label }}</h4>
                    <p class="text-sm text-gray-500">{{ type.desc }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">附加参数（选填）</label>
            <textarea v-model="formData.parameters" rows="3" placeholder="utm_source=oceanengine&utm_medium=cpc"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
            <p class="text-xs text-gray-500 mt-1">自定义URL参数，多个参数用&分隔</p>
          </div>

          <div class="flex justify-end gap-4 pt-4 border-t">
            <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleCancel">取消</button>
            <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateLink">创建链接</button>
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">追踪链接预览</h4>
          <div class="p-3 bg-gray-50 rounded-lg">
            <code class="text-xs text-gray-700 break-all">
              https://track.oceanengine.com/c/{{ formData.name ? formData.name.replace(/\s/g, '-') : 'example' }}
            </code>
          </div>
          <button class="mt-3 w-full px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50 text-sm" @click="handleCopyPreviewLink">
            复制链接
          </button>
        </div>

        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">使用说明</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 追踪链接用于统计广告点击</li>
            <li>• 可自定义UTM参数便于分析</li>
            <li>• 支持实时数据回传</li>
            <li>• 链接创建后可随时编辑</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
