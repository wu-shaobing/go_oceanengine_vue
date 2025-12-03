<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  type: 'custom',
  source: 'file',
  idType: 'phone',
  description: ''
})

const packageTypes = [
  { value: 'custom', label: '自定义人群', desc: '上传自有数据创建' },
  { value: 'lookalike', label: '相似人群', desc: '基于种子人群扩展' },
  { value: 'retarget', label: '重定向人群', desc: '已有互动用户' }
]

const sourceTypes = [
  { value: 'file', label: '文件上传', icon: '📁' },
  { value: 'api', label: 'API推送', icon: '🔌' },
  { value: 'dmp', label: 'DMP同步', icon: '🔄' }
]
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '人群管理' }, { name: '创建人群包' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建人群包</h1>
      <p class="mt-2 text-gray-600">上传或创建定向人群包</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">人群包名称</label>
          <input v-model="formData.name" type="text" placeholder="请输入人群包名称"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">人群类型</label>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div v-for="type in packageTypes" :key="type.value"
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                          formData.type === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="formData.type = type.value">
              <h4 class="font-medium text-gray-900">{{ type.label }}</h4>
              <p class="text-sm text-gray-500 mt-1">{{ type.desc }}</p>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">数据来源</label>
          <div class="flex gap-4">
            <div v-for="source in sourceTypes" :key="source.value"
                 :class="['flex items-center gap-2 px-4 py-3 border-2 rounded-lg cursor-pointer',
                          formData.source === source.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200']"
                 @click="formData.source = source.value">
              <span class="text-xl">{{ source.icon }}</span>
              <span class="text-sm font-medium text-gray-900">{{ source.label }}</span>
            </div>
          </div>
        </div>

        <div v-if="formData.source === 'file'">
          <label class="block text-sm font-medium text-gray-700 mb-2">ID类型</label>
          <select v-model="formData.idType" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            <option value="phone">手机号</option>
            <option value="imei">IMEI</option>
            <option value="idfa">IDFA</option>
            <option value="oaid">OAID</option>
          </select>
        </div>

        <div v-if="formData.source === 'file'" class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
          <div class="text-4xl mb-2">📤</div>
          <p class="text-gray-600">拖拽文件到此处或点击上传</p>
          <p class="text-sm text-gray-400 mt-1">支持 CSV、TXT 格式，单次最大10万条</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">描述（选填）</label>
          <textarea v-model="formData.description" rows="3" placeholder="请输入人群包描述"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
        </div>

        <div class="flex justify-end gap-4 pt-4 border-t">
          <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建人群包</button>
        </div>
      </div>
    </div>
  </div>
</template>
