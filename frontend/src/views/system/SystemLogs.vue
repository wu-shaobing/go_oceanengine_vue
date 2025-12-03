<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 20, total: 1256 })

const logs = ref([
  { id: 'L001', action: '登录系统', module: '用户管理', operator: '管理员', ip: '192.168.1.100', time: '2025-11-28 09:30:25', status: 'success' },
  { id: 'L002', action: '创建广告计划', module: '广告管理', operator: '张三', ip: '192.168.1.101', time: '2025-11-28 09:28:15', status: 'success' },
  { id: 'L003', action: '修改预算', module: '财务管理', operator: '李四', ip: '192.168.1.102', time: '2025-11-28 09:25:08', status: 'success' },
  { id: 'L004', action: '删除素材', module: '素材库', operator: '王五', ip: '192.168.1.103', time: '2025-11-28 09:20:42', status: 'warning' },
  { id: 'L005', action: '导出报表', module: '报表中心', operator: '管理员', ip: '192.168.1.100', time: '2025-11-28 09:15:33', status: 'success' }
])

const filterModule = ref('')
const filterDate = ref('')

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统' }, { name: '操作日志' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">操作日志</h1>
          <p class="mt-2 text-gray-600">查看系统操作记录</p>
        </div>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
          导出日志
        </button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <select v-model="filterModule" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部模块</option>
          <option value="user">用户管理</option>
          <option value="ad">广告管理</option>
          <option value="finance">财务管理</option>
          <option value="material">素材库</option>
          <option value="report">报表中心</option>
        </select>
        <input v-model="filterDate" type="date" class="px-4 py-2 border border-gray-300 rounded-lg">
        <input type="text" placeholder="搜索操作内容..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作内容</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">模块</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作人</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">IP地址</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="log in logs" :key="log.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ log.action }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ log.module }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ log.operator }}</td>
            <td class="px-6 py-4 text-sm text-gray-500 font-mono">{{ log.ip }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ log.time }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     log.status === 'success' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ log.status === 'success' ? '成功' : '警告' }}
              </span>
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
