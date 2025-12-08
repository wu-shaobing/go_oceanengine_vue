<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })
const activeTab = ref('all')

const notifications = ref([
  { id: 'N001', title: '广告计划审核通过', content: '您的广告计划"双11大促"已通过审核', type: 'success', time: '2025-11-28 09:30', read: false },
  { id: 'N002', title: '账户余额预警', content: '您的账户余额已不足1000元，请及时充值', type: 'warning', time: '2025-11-28 08:15', read: false },
  { id: 'N003', title: '系统升级通知', content: '系统将于今晚22:00-23:00进行升级维护', type: 'info', time: '2025-11-27 18:00', read: true },
  { id: 'N004', title: '素材审核未通过', content: '您上传的素材"产品图A"未通过审核，请修改后重新提交', type: 'error', time: '2025-11-27 15:30', read: true }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const markAllRead = () => {
  notifications.value.forEach(n => n.read = true)
}

const handleDismissNotification = (notification: typeof notifications.value[0]) => {
  notifications.value = notifications.value.filter(n => n.id !== notification.id)
}

const handleViewNotification = (notification: typeof notifications.value[0]) => {
  notification.read = true
  alert(`查看通知: ${notification.title}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统' }, { name: '消息通知' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">消息通知</h1>
          <p class="mt-2 text-gray-600">查看系统通知和提醒</p>
        </div>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="markAllRead">
          全部标记已读
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">全部消息</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">未读消息</p>
        <p class="text-2xl font-bold text-red-600 mt-1">{{ notifications.filter(n => !n.read).length }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日新增</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">重要通知</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">3</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex gap-4">
          <button v-for="tab in [{ key: 'all', label: '全部' }, { key: 'unread', label: '未读' }, { key: 'important', label: '重要' }]"
                  :key="tab.key"
                  :class="['px-4 py-2 rounded-lg transition-colors',
                           activeTab === tab.key ? 'bg-blue-50 text-blue-700' : 'text-gray-600 hover:bg-gray-100']"
                  @click="activeTab = tab.key">
            {{ tab.label }}
          </button>
        </div>
      </div>

      <div class="divide-y divide-gray-200">
        <div v-for="notification in notifications" :key="notification.id"
             :class="['px-6 py-4 hover:bg-gray-50 cursor-pointer', !notification.read ? 'bg-blue-50' : '']"
             @click="handleViewNotification(notification)">
          <div class="flex items-start gap-4">
            <div :class="['w-2 h-2 mt-2 rounded-full', !notification.read ? 'bg-blue-500' : 'bg-transparent']"></div>
            <div class="flex-1">
              <div class="flex items-center gap-2">
                <span :class="['px-2 py-0.5 rounded text-xs font-medium',
                       notification.type === 'success' ? 'bg-green-100 text-green-700' :
                       notification.type === 'warning' ? 'bg-yellow-100 text-yellow-700' :
                       notification.type === 'error' ? 'bg-red-100 text-red-700' :
                       'bg-blue-100 text-blue-700']">
                  {{ notification.type === 'success' ? '成功' : notification.type === 'warning' ? '警告' : notification.type === 'error' ? '错误' : '通知' }}
                </span>
                <h4 class="font-medium text-gray-900">{{ notification.title }}</h4>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ notification.content }}</p>
              <p class="text-xs text-gray-400 mt-2">{{ notification.time }}</p>
            </div>
            <button class="text-gray-400 hover:text-gray-600" @click.stop="handleDismissNotification(notification)">×</button>
          </div>
        </div>
      </div>

      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
