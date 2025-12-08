<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 22 })

const schedules = ref([
  { id: 'S001', name: '工作日投放', type: 'weekday', hours: '09:00-22:00', status: 'active', ads: 15 },
  { id: 'S002', name: '全天投放', type: 'allday', hours: '00:00-24:00', status: 'active', ads: 8 },
  { id: 'S003', name: '周末高峰', type: 'weekend', hours: '10:00-20:00', status: 'inactive', ads: 5 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建排期')
}

const handleEdit = (schedule: typeof schedules.value[0]) => {
  alert(`编辑排期: ${schedule.name}`)
}

const handleToggle = (schedule: typeof schedules.value[0]) => {
  schedule.status = schedule.status === 'active' ? 'inactive' : 'active'
  alert(`已${schedule.status === 'active' ? '启用' : '禁用'}: ${schedule.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '排期工具' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">排期工具</h1>
          <p class="mt-2 text-gray-600">管理广告投放时段</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建排期
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">排期规则</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已启用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">18</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联广告</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">28</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">当前投放</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">23</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">排期名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">投放时段</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联广告</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="schedule in schedules" :key="schedule.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ schedule.name }}</div>
              <div class="text-xs text-gray-500">{{ schedule.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">
                {{ schedule.type === 'weekday' ? '工作日' : schedule.type === 'weekend' ? '周末' : '全天' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ schedule.hours }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ schedule.ads }} 个</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     schedule.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ schedule.status === 'active' ? '已启用' : '未启用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleEdit(schedule)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleToggle(schedule)" :class="schedule.status === 'active' ? 'text-yellow-600 hover:text-yellow-800' : 'text-green-600 hover:text-green-800'">
                {{ schedule.status === 'active' ? '禁用' : '启用' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
