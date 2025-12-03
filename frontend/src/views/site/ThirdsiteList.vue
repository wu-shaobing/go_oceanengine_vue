<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const sites = ref([
  { id: 'TS001', name: '京东商城落地页', url: 'https://m.jd.com/promo/xxx', domain: 'jd.com', status: 'verified', createdAt: '2025-11-20' },
  { id: 'TS002', name: '天猫活动页', url: 'https://m.tmall.com/act/xxx', domain: 'tmall.com', status: 'verified', createdAt: '2025-11-18' },
  { id: 'TS003', name: '拼多多专场', url: 'https://mobile.yangkeduo.com/xxx', domain: 'yangkeduo.com', status: 'pending', createdAt: '2025-11-25' },
  { id: 'TS004', name: '有赞商城页', url: 'https://shop.youzan.com/xxx', domain: 'youzan.com', status: 'failed', createdAt: '2025-11-22' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页管理' }, { name: '第三方落地页' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">第三方落地页</h1>
          <p class="mt-2 text-gray-600">管理第三方平台落地页</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加第三方落地页
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已验证</p>
        <p class="text-2xl font-bold text-green-600 mt-1">38</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">待验证</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">5</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">验证失败</p>
        <p class="text-2xl font-bold text-red-600 mt-1">2</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">落地页名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">域名</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">URL</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">添加时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="site in sites" :key="site.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ site.name }}</div>
              <div class="text-xs text-gray-500">{{ site.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ site.domain }}</td>
            <td class="px-6 py-4">
              <a :href="site.url" target="_blank" class="text-sm text-blue-600 hover:underline truncate block max-w-xs">
                {{ site.url }}
              </a>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     site.status === 'verified' ? 'bg-green-100 text-green-800' :
                     site.status === 'pending' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                {{ site.status === 'verified' ? '已验证' : site.status === 'pending' ? '待验证' : '验证失败' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ site.createdAt }}</td>
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
