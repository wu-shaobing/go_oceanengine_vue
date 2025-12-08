<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()

const form = reactive({
  name: '',
  description: '',
  fields: [] as Array<{ name: string; type: string; required: boolean }>
})

const fieldTypes = [
  { value: 'text', label: '文本输入' },
  { value: 'phone', label: '手机号码' },
  { value: 'select', label: '下拉选择' },
  { value: 'radio', label: '单选' },
  { value: 'checkbox', label: '多选' },
  { value: 'date', label: '日期' }
]

const addField = () => {
  form.fields.push({ name: '', type: 'text', required: true })
}

const removeField = (index: number) => {
  form.fields.splice(index, 1)
}

const loading = ref(false)

const validateForm = () => {
  if (!form.name.trim()) {
    alert('请输入表单名称')
    return false
  }
  if (form.fields.length === 0) {
    alert('请至少添加一个字段')
    return false
  }
  return true
}

const handleSubmit = async () => {
  if (!validateForm()) return
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert('表单创建成功')
    router.push('/clue/form/list')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '创建表单' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建线索表单</h1>
      <p class="mt-2 text-gray-600">自定义表单字段收集用户线索</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-semibold text-gray-900 mb-4">基本信息</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">表单名称</label>
              <input v-model="form.name" type="text" placeholder="例如: 活动报名表"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">表单描述</label>
              <textarea v-model="form.description" rows="2" placeholder="表单用途说明"
                        class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="font-semibold text-gray-900">表单字段</h3>
            <button @click="addField" class="px-3 py-1.5 bg-blue-600 text-white text-sm rounded hover:bg-blue-700">
              + 添加字段
            </button>
          </div>
          
          <div v-if="form.fields.length === 0" class="py-8 text-center text-gray-500">
            暂无字段，点击"添加字段"开始配置
          </div>
          
          <div v-else class="space-y-4">
            <div v-for="(field, idx) in form.fields" :key="idx" 
                 class="p-4 bg-gray-50 rounded-lg border border-gray-200">
              <div class="flex items-center gap-4">
                <span class="text-sm text-gray-400 w-8">{{ idx + 1 }}.</span>
                <input v-model="field.name" type="text" placeholder="字段名称"
                       class="flex-1 px-3 py-1.5 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500">
                <select v-model="field.type" class="px-3 py-1.5 border border-gray-300 rounded text-sm">
                  <option v-for="t in fieldTypes" :key="t.value" :value="t.value">{{ t.label }}</option>
                </select>
                <label class="flex items-center gap-1 text-sm">
                  <input v-model="field.required" type="checkbox" class="rounded">
                  必填
                </label>
                <button @click="removeField(idx)" class="text-red-500 hover:text-red-700">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-semibold text-gray-900 mb-4">表单预览</h3>
          <div class="border border-gray-200 rounded-lg p-4 bg-gray-50">
            <h4 class="font-medium mb-4">{{ form.name || '表单名称' }}</h4>
            <div v-for="(field, idx) in form.fields" :key="idx" class="mb-3">
              <label class="block text-sm text-gray-600 mb-1">
                {{ field.name || '字段名称' }}
                <span v-if="field.required" class="text-red-500">*</span>
              </label>
              <input type="text" disabled class="w-full px-3 py-1.5 border border-gray-300 rounded bg-white text-sm">
            </div>
            <button class="w-full py-2 bg-blue-600 text-white rounded text-sm mt-4">提交</button>
          </div>
        </div>

        <div class="flex flex-col gap-3">
          <button @click="handleSubmit" class="w-full py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            保存表单
          </button>
          <router-link to="/clue/form/list" class="w-full py-2 border border-gray-300 rounded-lg text-gray-700 text-center hover:bg-gray-50">
            取消
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
