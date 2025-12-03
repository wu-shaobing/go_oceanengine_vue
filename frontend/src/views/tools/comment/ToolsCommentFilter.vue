<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })
const newKeyword = ref('')

const filterRules = ref([
  { id: 'FR001', keyword: '骗子', type: 'block', matchType: 'contain', blockedCount: 156, status: 'active' },
  { id: 'FR002', keyword: '垃圾', type: 'block', matchType: 'contain', blockedCount: 89, status: 'active' },
  { id: 'FR003', keyword: '微信', type: 'block', matchType: 'contain', blockedCount: 234, status: 'active' },
  { id: 'FR004', keyword: '竞品名', type: 'block', matchType: 'exact', blockedCount: 45, status: 'paused' }
])

const stats = ref({ totalRules: 45, activeRules: 38, blockedToday: 67, blockedTotal: 1256 })

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '评论管理' }, { name: '评论过滤' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">评论过滤规则</h1>
          <p class="mt-2 text-gray-600">管理评论过滤关键词和规则</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加规则
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总规则数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ stats.totalRules }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ stats.activeRules }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日拦截</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ stats.blockedToday }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">累计拦截</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">{{ stats.blockedTotal.toLocaleString() }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="newKeyword" type="text" placeholder="输入关键词快速添加..."
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">添加</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">匹配方式</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">拦截次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="rule in filterRules" :key="rule.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 font-medium text-gray-900">{{ rule.keyword }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                {{ rule.matchType === 'contain' ? '包含' : '精确' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ rule.blockedCount }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     rule.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ rule.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-red-600 hover:text-red-800">删除</button>
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
