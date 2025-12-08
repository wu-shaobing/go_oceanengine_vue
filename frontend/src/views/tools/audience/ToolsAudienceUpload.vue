<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  idType: 'phone',
  description: ''
})

const idTypes = [
  { value: 'phone', label: '手机号', icon: '📱', desc: 'MD5加密的手机号' },
  { value: 'imei', label: 'IMEI', icon: '📟', desc: '设备IMEI号' },
  { value: 'idfa', label: 'IDFA', icon: '🍎', desc: 'iOS广告标识符' },
  { value: 'oaid', label: 'OAID', icon: '🤖', desc: '安卓匿名设备标识' }
]

const handleSelectFile = () => {
  alert('选择文件')
}

const handleCancel = () => {
  alert('取消')
}

const handleCreate = () => {
  alert(`创建人群包: ${formData.value.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '上传人群包' }]" />
      <h1 class="text-3xl font-bold text-gray-900">上传人群包</h1>
      <p class="mt-2 text-gray-600">通过文件上传创建自定义人群包</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">人群包名称</label>
              <input v-model="formData.name" type="text" placeholder="请输入人群包名称"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ID类型</label>
              <div class="grid grid-cols-2 gap-3">
                <div v-for="type in idTypes" :key="type.value"
                     :class="['p-3 border-2 rounded-lg cursor-pointer transition-all',
                              formData.idType === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                     @click="formData.idType = type.value">
                  <div class="flex items-center gap-2">
                    <span class="text-xl">{{ type.icon }}</span>
                    <div>
                      <p class="font-medium text-gray-900">{{ type.label }}</p>
                      <p class="text-xs text-gray-500">{{ type.desc }}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">描述（选填）</label>
              <textarea v-model="formData.description" rows="2" placeholder="请输入人群包描述"
                        class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">上传文件</h3>
<div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center hover:border-blue-400 transition-colors cursor-pointer" @click="handleSelectFile">
            <div class="text-5xl mb-4">📤</div>
            <p class="text-lg font-medium text-gray-700">拖拽文件到此处上传</p>
            <p class="text-sm text-gray-500 mt-2">或点击选择文件</p>
            <p class="text-xs text-gray-400 mt-4">支持 CSV、TXT 格式，每行一个ID，最大100MB</p>
<button @click="handleSelectFile" class="mt-6 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">选择文件</button>
          </div>
        </div>

<div class="flex justify-end gap-4">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建人群包</button>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">上传说明</h4>
          <ul class="text-sm text-gray-600 space-y-2">
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              <span>文件格式：CSV或TXT</span>
            </li>
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              <span>每行一个ID，不含标题行</span>
            </li>
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              <span>手机号需MD5加密</span>
            </li>
            <li class="flex items-start gap-2">
              <span class="text-blue-500">•</span>
              <span>文件大小不超过100MB</span>
            </li>
          </ul>
        </div>

        <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
          <h4 class="font-medium text-yellow-800 mb-2">数据安全</h4>
          <p class="text-sm text-yellow-700">上传的数据将进行加密存储，仅用于广告投放定向，不会泄露给第三方。</p>
        </div>
      </div>
    </div>
  </div>
</template>
