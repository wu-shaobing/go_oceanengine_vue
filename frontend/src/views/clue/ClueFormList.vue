<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 32 })

const forms = ref([
  { id: 'F001', name: '618活动报名表', type: '活动报名', fields: 5, submissions: 1256, conversionRate: 4.2, status: 'active', createdAt: '2025-11-01' },
  { id: 'F002', name: '产品咨询表单', type: '咨询留资', fields: 4, submissions: 892, conversionRate: 3.8, status: 'active', createdAt: '2025-10-28' },
  { id: 'F003', name: '预约试驾表单', type: '预约服务', fields: 6, submissions: 456, conversionRate: 2.5, status: 'active', createdAt: '2025-10-25' },
  { id: 'F004', name: '课程试听申请', type: '活动报名', fields: 5, submissions: 0, conversionRate: 0, status: 'draft', createdAt: '2025-11-10' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '线索管理' }, { name: '表单列表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">青鸟表单</h1>
          <p class="mt-2 text-gray-600">管理线索收集表单</p>
        </div>
        <router-link to="/clue/form/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建表单
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总表单</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日提交</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">156</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">本周提交</p>
        <p class="text-2xl font-bold text-green-600 mt-1">1,024</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">3.5%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">表单名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">字段数</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提交数</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="form in forms" :key="form.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ form.name }}</div>
                <div class="text-xs text-gray-500">ID: {{ form.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ form.type }}</span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-700">{{ form.fields }} 个</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ form.submissions.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm" :class="form.conversionRate >= 3 ? 'text-green-600' : 'text-gray-700'">
                {{ form.conversionRate }}%
              </td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  form.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                  {{ form.status === 'active' ? '启用' : '草稿' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ form.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
                <button class="text-green-600 hover:text-green-800 mr-3">数据</button>
                <button class="text-gray-600 hover:text-gray-800">复制</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
