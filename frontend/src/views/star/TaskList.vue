<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '任务管理' }, { name: '任务列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">任务管理</h1>
        <p class="text-gray-600 mt-1">管理达人营销任务</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建任务
      </button>
    </div>

    <!-- 统计 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">全部任务</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stats.total }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">进行中</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.running }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已完成</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.completed }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总预算</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">¥{{ stats.budget.toLocaleString() }}</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="任务名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="draft">草稿</option>
          <option value="running">进行中</option>
          <option value="completed">已完成</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频种草</option>
          <option value="live">直播带货</option>
          <option value="post">图文推广</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 任务列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">任务信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">预算</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">参与达人</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">内容产出</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">进度</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="task in tasks" :key="task.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ task.name }}</div>
              <div class="text-sm text-gray-500">{{ task.period }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ task.type }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(task.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(task.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ task.budget.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ task.influencers }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ task.contents }}</td>
            <td class="px-4 py-3 text-right">
              <div class="flex items-center justify-end">
                <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: task.progress + '%' }"></div>
                </div>
                <span class="text-sm">{{ task.progress }}%</span>
              </div>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="50" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { starApi } from '@/api/star'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()

const loading = ref(false)
const error = ref('')
const pagination = ref({ page: 1, pageSize: 10, total: 0 })

interface TaskItem {
  id: number | string
  name: string
  period: string
  type: string
  status: string
  budget: number
  influencers: number
  contents: number
  progress: number
}

const tasks = ref<TaskItem[]>([])

const stats = computed(() => {
  const total = pagination.value.total
  const running = tasks.value.filter(t => t.status === 'running').length
  const completed = tasks.value.filter(t => t.status === 'completed').length
  const budget = tasks.value.reduce((sum, t) => sum + t.budget, 0)
  return { total, running, completed, budget }
})

const filters = ref({
  keyword: '',
  status: '',
  type: ''
})

const fetchTasks = async () => {
  const starId = advertiserStore.currentAdvertiserId
  if (!starId) {
    error.value = '请先选择账户'
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const params: { star_id: number; page: number; page_size: number; status?: string } = {
      star_id: starId,
      page: pagination.value.page,
      page_size: pagination.value.pageSize
    }
    if (filters.value.status) {
      params.status = filters.value.status
    }
    
    const res = await starApi.getTaskList(params)
    const data = res as any
    
    if (data?.list) {
      tasks.value = data.list.map((item: any) => ({
        id: item.task_id,
        name: item.task_name,
        period: `${item.start_time?.substring(0, 10) || '-'} - ${item.end_time?.substring(0, 10) || '-'}`,
        type: item.content_type || '视频种草',
        status: item.status || 'draft',
        budget: item.budget || 0,
        influencers: item.author_count || 0,
        contents: item.video_count || 0,
        progress: item.completed_count && item.video_count ? Math.round(item.completed_count / item.video_count * 100) : 0
      }))
      pagination.value.total = data.page_info?.total_number || 0
    }
  } catch (e: any) {
    error.value = e.message || '加载失败'
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.page = 1
  fetchTasks()
}

const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchTasks()
}

onMounted(() => {
  fetchTasks()
})

defineExpose({ handleSearch, handlePageChange })

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    draft: 'bg-gray-100 text-gray-800',
    running: 'bg-blue-100 text-blue-800',
    completed: 'bg-green-100 text-green-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    draft: '草稿',
    running: '进行中',
    completed: '已完成'
  }
  return texts[status] || status
}
</script>
