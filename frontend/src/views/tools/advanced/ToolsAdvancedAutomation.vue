<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const rules = ref([
  { id: 'R001', name: '低效广告自动暂停', trigger: 'CTR < 1%', action: '暂停广告', scope: '全部广告', execCount: 25, status: 'active' },
  { id: 'R002', name: '预算耗尽提醒', trigger: '预算 < 10%', action: '发送通知', scope: '全部广告组', execCount: 12, status: 'active' },
  { id: 'R003', name: '高效广告加预算', trigger: 'ROI > 5', action: '预算+20%', scope: '选定广告', execCount: 8, status: 'inactive' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建自动化规则')
}

const handleEdit = (rule: typeof rules.value[0]) => {
  alert(`编辑规则: ${rule.name}`)
}

const handleLogs = (rule: typeof rules.value[0]) => {
  alert(`查看日志: ${rule.name}`)
}

const handleToggle = (rule: typeof rules.value[0]) => {
  rule.status = rule.status === 'active' ? 'inactive' : 'active'
  alert(`规则已${rule.status === 'active' ? '启用' : '禁用'}: ${rule.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '自动化规则' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">自动化规则</h1>
          <p class="mt-2 text-gray-600">配置自动执行的投放规则</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建规则
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">自动化规则</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已启用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日执行</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">45</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">节省工时</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">12h</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">规则名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">触发条件</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">执行动作</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">作用范围</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">执行次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="rule in rules" :key="rule.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ rule.name }}</div>
              <div class="text-xs text-gray-500">{{ rule.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-800">{{ rule.trigger }}</code>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ rule.action }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ rule.scope }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ rule.execCount }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     rule.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ rule.status === 'active' ? '已启用' : '未启用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleEdit(rule)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleLogs(rule)" class="text-gray-600 hover:text-gray-800 mr-3">日志</button>
              <button @click="handleToggle(rule)" :class="rule.status === 'active' ? 'text-yellow-600' : 'text-green-600'">
                {{ rule.status === 'active' ? '禁用' : '启用' }}
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
