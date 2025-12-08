<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const schedules = ref([
  { id: 'SC001', name: '每日消耗报表', frequency: 'daily', time: '09:00', recipients: 3, lastSent: '2025-11-28 09:00', status: 'active' },
  { id: 'SC002', name: '每周效果汇总', frequency: 'weekly', time: '周一 10:00', recipients: 5, lastSent: '2025-11-25 10:00', status: 'active' },
  { id: 'SC003', name: '月度财务报表', frequency: 'monthly', time: '每月1日 08:00', recipients: 2, lastSent: '2025-11-01 08:00', status: 'active' },
  { id: 'SC004', name: '转化数据日报', frequency: 'daily', time: '18:00', recipients: 4, lastSent: '2025-11-27 18:00', status: 'paused' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建定时任务')
}

const handleEdit = (schedule: typeof schedules.value[0]) => {
  alert(`编辑: ${schedule.name}`)
}

const handlePause = (schedule: typeof schedules.value[0]) => {
  alert(`暂停: ${schedule.name}`)
}

const handleEnable = (schedule: typeof schedules.value[0]) => {
  alert(`启用: ${schedule.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表中心' }, { name: '定时报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">定时报表</h1>
          <p class="mt-2 text-gray-600">配置自动发送的定时报表</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreate">
          创建定时任务
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">定时任务数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">运行中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已暂停</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日发送</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">12</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">报表名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发送频率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发送时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接收人</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">上次发送</th>
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
                {{ schedule.frequency === 'daily' ? '每日' : schedule.frequency === 'weekly' ? '每周' : '每月' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ schedule.time }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ schedule.recipients }} 人</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ schedule.lastSent }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     schedule.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ schedule.status === 'active' ? '运行中' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleEdit(schedule)">编辑</button>
              <button v-if="schedule.status === 'active'" class="text-yellow-600 hover:text-yellow-800" @click="handlePause(schedule)">暂停</button>
              <button v-else class="text-green-600 hover:text-green-800" @click="handleEnable(schedule)">启用</button>
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
