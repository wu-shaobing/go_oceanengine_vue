<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  siteId: '',
  submitText: '立即提交',
  successMessage: '提交成功，我们会尽快与您联系'
})

const fields = ref([
  { id: 1, label: '姓名', type: 'text', required: true },
  { id: 2, label: '手机号', type: 'phone', required: true }
])

const fieldTypes = [
  { value: 'text', label: '单行文本' },
  { value: 'phone', label: '手机号' },
  { value: 'email', label: '邮箱' },
  { value: 'textarea', label: '多行文本' },
  { value: 'select', label: '下拉选择' },
  { value: 'radio', label: '单选' },
  { value: 'checkbox', label: '多选' }
]

const addField = () => {
  fields.value.push({
    id: Date.now(),
    label: '',
    type: 'text',
    required: false
  })
}

const removeField = (fieldId: number) => {
  fields.value = fields.value.filter(f => f.id !== fieldId)
}

const handleCancel = () => {
  alert('取消')
}

const handleCreateForm = () => {
  alert('创建表单')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '创建表单' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建表单</h1>
      <p class="mt-2 text-gray-600">配置落地页收集表单</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">表单名称</label>
              <input v-model="formData.name" type="text" placeholder="请输入表单名称"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">所属落地页</label>
              <select v-model="formData.siteId" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option value="">请选择落地页</option>
                <option value="S001">官网落地页</option>
                <option value="S002">活动页A</option>
                <option value="S003">品牌页</option>
              </select>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">表单字段</h3>
            <button class="px-3 py-1 text-sm text-blue-600 hover:bg-blue-50 rounded" @click="addField">
              + 添加字段
            </button>
          </div>
          <div class="space-y-4">
            <div v-for="(field, index) in fields" :key="field.id"
                 class="flex items-center gap-4 p-4 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-500">{{ index + 1 }}</span>
              <input v-model="field.label" type="text" placeholder="字段名称"
                     class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <select v-model="field.type" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                <option v-for="type in fieldTypes" :key="type.value" :value="type.value">{{ type.label }}</option>
              </select>
              <label class="flex items-center gap-1 text-sm">
                <input v-model="field.required" type="checkbox" class="rounded text-blue-600">
                必填
              </label>
              <button class="text-red-500 hover:text-red-700" @click="removeField(field.id)">×</button>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">提交设置</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">提交按钮文字</label>
              <input v-model="formData.submitText" type="text"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">成功提示语</label>
              <input v-model="formData.successMessage" type="text"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
          </div>
        </div>

        <div class="flex justify-end gap-4">
          <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleCancel">取消</button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateForm">创建表单</button>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h4 class="font-medium text-gray-900 mb-4">表单预览</h4>
        <div class="border border-gray-200 rounded-lg p-4">
          <div class="space-y-4">
            <div v-for="field in fields" :key="field.id">
              <label class="block text-sm font-medium text-gray-700 mb-1">
                {{ field.label || '未命名字段' }}
                <span v-if="field.required" class="text-red-500">*</span>
              </label>
              <input v-if="field.type === 'text' || field.type === 'phone' || field.type === 'email'"
                     type="text" disabled
                     class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded text-sm">
              <textarea v-else-if="field.type === 'textarea'" disabled
                        class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded text-sm" rows="2"></textarea>
            </div>
          </div>
          <button class="w-full mt-4 py-2 bg-blue-600 text-white rounded-lg text-sm">
            {{ formData.submitText }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
