<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { systemApi, type OperationLog } from '@/api'

const loading = ref(false)
const logs = ref<OperationLog[]>([])
const modules = ref<string[]>([])
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// 筛选条件
const filterModule = ref('')
const filterStatus = ref<number | undefined>(undefined)
const filterStartDate = ref('')
const filterEndDate = ref('')
const searchKeyword = ref('')

// 获取日志列表
const fetchLogs = async () => {
  loading.value = true
  try {
    const params: Record<string, unknown> = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (filterModule.value) params.module = filterModule.value
    if (filterStatus.value !== undefined) params.status = filterStatus.value
    if (filterStartDate.value) params.start_time = filterStartDate.value + ' 00:00:00'
    if (filterEndDate.value) params.end_time = filterEndDate.value + ' 23:59:59'
    if (searchKeyword.value) params.username = searchKeyword.value

    const res = await systemApi.getOperationLogList(params)
    logs.value = res.list || []
    pagination.total = res.total || 0
  } catch (error) {
    console.error('获取日志失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取模块列表
const fetchModules = async () => {
  try {
    modules.value = await systemApi.getOperationLogModules()
  } catch (error) {
    console.error('获取模块列表失败:', error)
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
  fetchLogs()
}

const handleSearch = () => {
  pagination.page = 1
  fetchLogs()
}

const handleReset = () => {
  filterModule.value = ''
  filterStatus.value = undefined
  filterStartDate.value = ''
  filterEndDate.value = ''
  searchKeyword.value = ''
  pagination.page = 1
  fetchLogs()
}

// 删除旧日志
const showDeleteDialog = ref(false)
const deleteBeforeDate = ref('')

const handleDeleteLogs = async () => {
  if (!deleteBeforeDate.value) {
    alert('请选择截止日期')
    return
  }
  if (!confirm(`确定删除 ${deleteBeforeDate.value} 之前的所有日志吗？此操作不可撤销！`)) {
    return
  }
  try {
    await systemApi.deleteOperationLogs(deleteBeforeDate.value + ' 23:59:59')
    showDeleteDialog.value = false
    deleteBeforeDate.value = ''
    fetchLogs()
    alert('删除成功')
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// 导出日志
const handleExportLogs = () => {
  const csvContent = [
    ['ID', '用户名', '模块', '操作', '方法', '路径', 'IP', '状态', '耗时(ms)', '时间'].join(','),
    ...logs.value.map(log => [
      log.id,
      log.username,
      log.module,
      log.action,
      log.method,
      log.path,
      log.ip,
      log.status,
      log.duration,
      log.created_at
    ].join(','))
  ].join('\n')

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = `operation_logs_${new Date().toISOString().split('T')[0]}.csv`
  link.click()
}

// 状态映射
const getStatusInfo = (status: number) => {
  if (status >= 500) return { text: '服务器错误', class: 'bg-red-100 text-red-800' }
  if (status >= 400) return { text: '客户端错误', class: 'bg-yellow-100 text-yellow-800' }
  return { text: '成功', class: 'bg-green-100 text-green-800' }
}

onMounted(() => {
  fetchLogs()
  fetchModules()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统' }, { name: '操作日志' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">操作日志</h1>
          <p class="mt-2 text-gray-600">查看系统操作记录，共 {{ pagination.total }} 条</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleExportLogs">
            导出日志
          </button>
          <button class="px-4 py-2 border border-red-300 text-red-600 rounded-lg hover:bg-red-50" @click="showDeleteDialog = true">
            清理旧日志
          </button>
        </div>
      </div>
    </div>

    <!-- 筛选条件 -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <select v-model="filterModule" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部模块</option>
          <option v-for="m in modules" :key="m" :value="m">{{ m }}</option>
        </select>
        <select v-model="filterStatus" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option :value="undefined">全部状态</option>
          <option :value="200">成功 (2xx)</option>
          <option :value="400">客户端错误 (4xx)</option>
          <option :value="500">服务器错误 (5xx)</option>
        </select>
        <input v-model="filterStartDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg" placeholder="开始日期">
        <span class="self-center text-gray-500">至</span>
        <input v-model="filterEndDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg" placeholder="结束日期">
        <input v-model="searchKeyword" type="text" placeholder="搜索用户名..." class="flex-1 min-w-48 px-4 py-2 border border-gray-300 rounded-lg">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleReset">重置</button>
      </div>
    </div>

    <!-- 日志列表 -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div v-if="loading" class="flex items-center justify-center py-20">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <span class="ml-2 text-gray-500">加载中...</span>
      </div>
      <table v-else class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">用户</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">模块</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">方法</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">路径</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">IP</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">耗时</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-if="logs.length === 0">
            <td colspan="9" class="px-4 py-8 text-center text-gray-500">暂无日志记录</td>
          </tr>
          <tr v-for="log in logs" :key="log.id" class="hover:bg-gray-50">
            <td class="px-4 py-3 text-sm text-gray-900">{{ log.username || '-' }}</td>
            <td class="px-4 py-3">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ log.module }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-600">{{ log.action }}</td>
            <td class="px-4 py-3">
              <span :class="['px-2 py-1 rounded text-xs font-medium',
                     log.method === 'POST' ? 'bg-green-100 text-green-700' :
                     log.method === 'PUT' ? 'bg-yellow-100 text-yellow-700' :
                     log.method === 'DELETE' ? 'bg-red-100 text-red-700' : 'bg-gray-100 text-gray-700']">
                {{ log.method }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500 font-mono max-w-48 truncate" :title="log.path">{{ log.path }}</td>
            <td class="px-4 py-3 text-sm text-gray-500 font-mono">{{ log.ip }}</td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ log.duration }}ms</td>
            <td class="px-4 py-3">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusInfo(log.status).class]">
                {{ log.status }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500 whitespace-nowrap">{{ log.created_at }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>

    <!-- 删除日志弹窗 -->
    <div v-if="showDeleteDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-96">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">清理旧日志</h3>
        <p class="text-sm text-gray-600 mb-4">删除指定日期之前的所有操作日志，此操作不可撤销。</p>
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">截止日期</label>
          <input v-model="deleteBeforeDate" type="date" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
        </div>
        <div class="flex justify-end gap-2">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showDeleteDialog = false">取消</button>
          <button class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700" @click="handleDeleteLogs">确认删除</button>
        </div>
      </div>
    </div>
  </div>
</template>
