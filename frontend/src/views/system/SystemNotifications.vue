<script setup lang="ts">
import { ref, reactive, onMounted, watch, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { systemApi, type NotificationItem, type NotificationStatsResp } from '@/api/system'
import type { PageResponse } from '@/api/request'

const loading = ref(false)
const pagination = reactive({ page: 1, pageSize: 10, total: 0 })
const activeTab = ref<'all' | 'unread' | 'important'>('all')
const notifications = ref<NotificationItem[]>([])
const stats = ref<NotificationStatsResp>({
  total: 0,
  unread: 0,
  today_new: 0,
  important: 0
})

// 根据 tab 计算筛选参数
const filterParams = computed(() => {
  const params: { is_read?: boolean; type?: string } = {}
  if (activeTab.value === 'unread') {
    params.is_read = false
  } else if (activeTab.value === 'important') {
    params.type = 'warning' // warning 和 error 都是重要通知
  }
  return params
})

// 加载通知列表
const loadNotifications = async () => {
  loading.value = true
  try {
    const res = await systemApi.getNotificationList({
      page: pagination.page,
      page_size: pagination.pageSize,
      ...filterParams.value
    }) as PageResponse<NotificationItem>
    if (res) {
      notifications.value = res.list || []
      pagination.total = res.total || 0
    }
  } catch (error: any) {
    alert(error.message || '加载通知失败')
  } finally {
    loading.value = false
  }
}

// 加载统计数据
const loadStats = async () => {
  try {
    const res = await systemApi.getNotificationStats() as NotificationStatsResp
    if (res) {
      stats.value = res
    }
  } catch (error) {
    // 统计加载失败不显示错误
  }
}

// 分页变化
const handlePageChange = (page: number) => {
  pagination.page = page
  loadNotifications()
}

// Tab 切换
watch(activeTab, () => {
  pagination.page = 1
  loadNotifications()
})

// 全部标记已读
const markAllRead = async () => {
  try {
    await systemApi.markAllNotificationsAsRead()
    alert('已全部标记为已读')
    loadNotifications()
    loadStats()
  } catch (error: any) {
    alert(error.message || '操作失败')
  }
}

// 查看通知（标记已读）
const handleViewNotification = async (notification: NotificationItem) => {
  if (!notification.is_read) {
    try {
      await systemApi.markNotificationsAsRead([notification.id])
      notification.is_read = true
      stats.value.unread = Math.max(0, stats.value.unread - 1)
    } catch (error) {
      // 标记失败不影响查看
    }
  }
  // 如果有链接，跳转
  if (notification.link) {
    window.open(notification.link, '_blank')
  }
}

// 删除通知
const handleDismissNotification = async (notification: NotificationItem) => {
  if (!confirm('确定要删除这条通知吗？')) {
    return
  }
  try {
    await systemApi.deleteNotifications([notification.id])
    alert('删除成功')
    loadNotifications()
    loadStats()
  } catch (error: any) {
    alert(error.message || '删除失败')
  }
}

// 获取通知类型标签
const getTypeLabel = (type: string) => {
  const labels: Record<string, string> = {
    success: '成功',
    warning: '警告',
    error: '错误',
    info: '通知'
  }
  return labels[type] || '通知'
}

onMounted(() => {
  loadNotifications()
  loadStats()
})
</script>

<template>
  <div class="space-y-6" v-loading="loading">
    <div>
      <Breadcrumb :items="[{ name: '系统' }, { name: '消息通知' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">消息通知</h1>
          <p class="mt-2 text-gray-600">查看系统通知和提醒</p>
        </div>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
                :disabled="stats.unread === 0"
                @click="markAllRead">
          全部标记已读
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">全部消息</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ stats.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">未读消息</p>
        <p class="text-2xl font-bold text-red-600 mt-1">{{ stats.unread }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日新增</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ stats.today_new }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">重要通知</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">{{ stats.important }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex gap-4">
          <button v-for="tab in [{ key: 'all', label: '全部' }, { key: 'unread', label: '未读' }, { key: 'important', label: '重要' }]"
                  :key="tab.key"
                  :class="['px-4 py-2 rounded-lg transition-colors',
                           activeTab === tab.key ? 'bg-blue-50 text-blue-700' : 'text-gray-600 hover:bg-gray-100']"
                  @click="activeTab = tab.key as 'all' | 'unread' | 'important'">
            {{ tab.label }}
          </button>
        </div>
      </div>

      <div class="divide-y divide-gray-200">
        <template v-if="notifications.length > 0">
          <div v-for="notification in notifications" :key="notification.id"
               :class="['px-6 py-4 hover:bg-gray-50 cursor-pointer', !notification.is_read ? 'bg-blue-50' : '']"
               @click="handleViewNotification(notification)">
            <div class="flex items-start gap-4">
              <div :class="['w-2 h-2 mt-2 rounded-full flex-shrink-0', !notification.is_read ? 'bg-blue-500' : 'bg-transparent']"></div>
              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-2">
                  <span :class="['px-2 py-0.5 rounded text-xs font-medium flex-shrink-0',
                         notification.type === 'success' ? 'bg-green-100 text-green-700' :
                         notification.type === 'warning' ? 'bg-yellow-100 text-yellow-700' :
                         notification.type === 'error' ? 'bg-red-100 text-red-700' :
                         'bg-blue-100 text-blue-700']">
                    {{ getTypeLabel(notification.type) }}
                  </span>
                  <h4 class="font-medium text-gray-900 truncate">{{ notification.title }}</h4>
                </div>
                <p class="text-sm text-gray-600 mt-1 line-clamp-2">{{ notification.content }}</p>
                <p class="text-xs text-gray-400 mt-2">{{ notification.created_at }}</p>
              </div>
              <button class="text-gray-400 hover:text-red-500 flex-shrink-0 text-xl leading-none"
                      @click.stop="handleDismissNotification(notification)">×</button>
            </div>
          </div>
        </template>
        <div v-else class="px-6 py-12 text-center text-gray-500">
          <p>暂无通知消息</p>
        </div>
      </div>

      <div v-if="pagination.total > 0" class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
