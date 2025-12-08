<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()

const formInfo = ref({
  id: 'CF12345',
  name: '产品咨询表单',
  status: 'active',
  createdAt: '2025-11-15',
  totalClues: 3256,
  todayClues: 128,
  convertRate: 12.5
})

const fields = ref([
  { name: '姓名', type: 'text', required: true },
  { name: '手机号', type: 'phone', required: true },
  { name: '城市', type: 'select', required: false },
  { name: '需求描述', type: 'textarea', required: false }
])

const recentClues = ref([
  { id: 'CL001', name: '张三', phone: '138****1234', city: '北京', time: '2025-11-28 09:30', status: 'new' },
  { id: 'CL002', name: '李四', phone: '139****5678', city: '上海', time: '2025-11-28 09:25', status: 'followed' },
  { id: 'CL003', name: '王五', phone: '137****9012', city: '广州', time: '2025-11-28 09:20', status: 'new' }
])

const handleEdit = () => {
  router.push(`/clue/form/edit/${formInfo.value.id}`)
}

const handleExport = () => {
  alert(`正在导出表单「${formInfo.value.name}」的 ${formInfo.value.totalClues} 条线索...`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '表单详情' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ formInfo.name }}</h1>
          <p class="mt-2 text-gray-600">ID: {{ formInfo.id }}</p>
        </div>
        <div class="flex gap-3">
          <button @click="handleEdit" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">编辑</button>
          <button @click="handleExport" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">导出线索</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">累计线索</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ formInfo.totalClues.toLocaleString() }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日新增</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ formInfo.todayClues }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化率</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ formInfo.convertRate }}%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">状态</p>
        <span class="px-2 py-1 bg-green-100 text-green-800 rounded text-sm font-medium">运行中</span>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">表单字段</h4>
        <div class="space-y-3">
          <div v-for="field in fields" :key="field.name"
               class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center gap-2">
              <span class="text-sm font-medium text-gray-900">{{ field.name }}</span>
              <span v-if="field.required" class="text-red-500 text-xs">*必填</span>
            </div>
            <span class="text-xs text-gray-500 px-2 py-1 bg-gray-200 rounded">{{ field.type }}</span>
          </div>
        </div>
      </div>

      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">最近线索</h4>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">姓名</th>
              <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">手机号</th>
              <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">城市</th>
              <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">时间</th>
              <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">状态</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="clue in recentClues" :key="clue.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-sm text-gray-900">{{ clue.name }}</td>
              <td class="px-4 py-3 text-sm text-gray-600">{{ clue.phone }}</td>
              <td class="px-4 py-3 text-sm text-gray-600">{{ clue.city }}</td>
              <td class="px-4 py-3 text-sm text-gray-500">{{ clue.time }}</td>
              <td class="px-4 py-3">
                <span :class="['px-2 py-1 rounded text-xs',
                       clue.status === 'new' ? 'bg-blue-100 text-blue-700' : 'bg-green-100 text-green-700']">
                  {{ clue.status === 'new' ? '新线索' : '已跟进' }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
