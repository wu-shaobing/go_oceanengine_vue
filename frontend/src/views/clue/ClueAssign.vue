<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const rules = ref([
  { id: 'R001', name: '地区分配规则', condition: '按省份', assignees: ['北京团队', '上海团队', '广州团队'], status: 'active', priority: 1 },
  { id: 'R002', name: '产品分配规则', condition: '按意向产品', assignees: ['产品A组', '产品B组'], status: 'active', priority: 2 },
  { id: 'R003', name: '轮流分配规则', condition: '轮询分配', assignees: ['销售1', '销售2', '销售3'], status: 'inactive', priority: 3 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '线索分配' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">线索分配</h1>
          <p class="mt-2 text-gray-600">配置线索自动分配规则</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加规则
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">分配规则</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已启用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">28</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日分配</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">256</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">分配成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">100%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">规则名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分配条件</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分配对象</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">优先级</th>
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
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ rule.condition }}</span>
            </td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="assignee in rule.assignees.slice(0, 2)" :key="assignee"
                      class="px-2 py-0.5 bg-blue-100 text-blue-700 rounded text-xs">{{ assignee }}</span>
                <span v-if="rule.assignees.length > 2" class="px-2 py-0.5 bg-gray-100 text-gray-600 rounded text-xs">
                  +{{ rule.assignees.length - 2 }}
                </span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-yellow-100 text-yellow-700 rounded text-xs font-medium">{{ rule.priority }}</span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     rule.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ rule.status === 'active' ? '已启用' : '未启用' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button :class="rule.status === 'active' ? 'text-yellow-600 hover:text-yellow-800' : 'text-green-600 hover:text-green-800'">
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
