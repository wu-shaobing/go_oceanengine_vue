<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  url: '',
  verifyMethod: 'meta',
  verifyCode: ''
})

const verifyMethods = [
  { value: 'meta', label: 'Meta标签验证', desc: '在页面head中添加验证标签' },
  { value: 'file', label: '文件验证', desc: '上传验证文件到网站根目录' },
  { value: 'dns', label: 'DNS验证', desc: '添加DNS TXT记录' }
]

const generatedCode = ref('ocean_verify_abc123def456')

const handleDownloadFile = () => {
  alert('下载验证文件')
}

const handleCancel = () => {
  alert('取消')
}

const handleStartVerify = () => {
  alert('开始验证')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页管理' }, { name: '添加第三方落地页' }]" />
      <h1 class="text-3xl font-bold text-gray-900">添加第三方落地页</h1>
      <p class="mt-2 text-gray-600">添加并验证第三方平台落地页</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">落地页名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入落地页名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">落地页URL</label>
          <input v-model="formData.url" type="url" placeholder="https://example.com/landing"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <p class="text-xs text-gray-500 mt-1">请输入完整的落地页URL地址</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">验证方式</label>
          <div class="space-y-3">
            <div v-for="method in verifyMethods" :key="method.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                          formData.verifyMethod === method.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.verifyMethod = method.value">
              <div class="flex items-center gap-3">
                <div :class="['w-4 h-4 rounded-full border-2',
                             formData.verifyMethod === method.value ? 'border-blue-500 bg-blue-500' : 'border-gray-300']">
                  <div v-if="formData.verifyMethod === method.value" class="w-full h-full flex items-center justify-center">
                    <div class="w-1.5 h-1.5 bg-white rounded-full"></div>
                  </div>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ method.label }}</h4>
                  <p class="text-sm text-gray-500">{{ method.desc }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-4">
          <h4 class="font-medium text-gray-900 mb-3">验证步骤</h4>
          <div v-if="formData.verifyMethod === 'meta'" class="space-y-3">
            <p class="text-sm text-gray-600">1. 复制以下验证代码：</p>
            <div class="bg-gray-800 text-green-400 p-3 rounded font-mono text-sm">
              &lt;meta name="ocean-verify" content="{{ generatedCode }}"&gt;
            </div>
            <p class="text-sm text-gray-600">2. 将代码添加到落地页的 &lt;head&gt; 标签内</p>
            <p class="text-sm text-gray-600">3. 点击验证按钮完成验证</p>
          </div>
          <div v-else-if="formData.verifyMethod === 'file'" class="space-y-3">
            <p class="text-sm text-gray-600">1. 下载验证文件</p>
            <button class="px-4 py-2 bg-gray-200 rounded-lg text-sm" @click="handleDownloadFile">下载 ocean_verify.txt</button>
            <p class="text-sm text-gray-600">2. 将文件上传到网站根目录</p>
            <p class="text-sm text-gray-600">3. 点击验证按钮完成验证</p>
          </div>
          <div v-else class="space-y-3">
            <p class="text-sm text-gray-600">1. 添加以下DNS TXT记录：</p>
            <div class="bg-gray-800 text-green-400 p-3 rounded font-mono text-sm">
              TXT ocean-verify={{ generatedCode }}
            </div>
            <p class="text-sm text-gray-600">2. 等待DNS生效（可能需要数分钟到数小时）</p>
            <p class="text-sm text-gray-600">3. 点击验证按钮完成验证</p>
          </div>
        </div>

        <div class="flex justify-end gap-4 pt-4 border-t">
          <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleCancel">取消</button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleStartVerify">开始验证</button>
        </div>
      </div>
    </div>
  </div>
</template>
