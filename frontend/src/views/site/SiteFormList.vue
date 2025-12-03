<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 68 })

const forms = ref([
  { id: 'SF001', name: '产品咨询表单', siteName: '官网落地页', fields: 5, submissions: 1256, convRate: 12.5, status: 'active' },
  { id: 'SF002', name: '预约试用表单', siteName: '活动页A', fields: 4, submissions: 890, convRate: 15.2, status: 'active' },
  { id: 'SF003', name: '留言表单', siteName: '品牌页', fields: 3, submissions: 456, convRate: 8.5, status: 'active' },
  { id: 'SF004', name: '调研问卷', siteName: '调研页', fields: 10, submissions: 125, convRate: 5.2, status: 'paused' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '表单管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">表单管理</h1>
          <p class="mt-2 text-gray-600">管理落地页表单配置</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建表单
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">表单总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总提交数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">2,727</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-green-600 mt-1">10.4%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日提交</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">156</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">表单名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">所属落地页</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">字段数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提交数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="form in forms" :key="form.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ form.name }}</div>
              <div class="text-xs text-gray-500">{{ form.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ form.siteName }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ form.fields }} 个</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ form.submissions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm" :class="form.convRate >= 10 ? 'text-green-600' : 'text-yellow-600'">{{ form.convRate }}%</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     form.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ form.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-gray-600 hover:text-gray-800">数据</button>
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
