<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 48 })

const sites = ref([
  { id: 'S001', name: '618大促活动页', type: '活动页', template: '电商模板', pv: 125600, conversionRate: 3.2, status: 'published', updatedAt: '2025-11-10' },
  { id: 'S002', name: '新品发布落地页', type: '产品页', template: '品牌模板', pv: 89200, conversionRate: 2.8, status: 'published', updatedAt: '2025-11-09' },
  { id: 'S003', name: '双11预热页面', type: '活动页', template: '促销模板', pv: 0, conversionRate: 0, status: 'draft', updatedAt: '2025-11-08' },
  { id: 'S004', name: '品牌故事页', type: '品牌页', template: '品牌模板', pv: 45800, conversionRate: 1.5, status: 'published', updatedAt: '2025-11-05' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页管理' }, { name: '站点列表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">橙子建站</h1>
          <p class="mt-2 text-gray-600">管理您的落地页站点</p>
        </div>
        <router-link to="/site/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建站点
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总站点</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已发布</p>
        <p class="text-2xl font-bold text-green-600 mt-1">42</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日PV</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">26.1万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2.4%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <input type="text" placeholder="搜索站点名称" class="flex-1 min-w-[200px] px-4 py-2 border border-gray-300 rounded-lg">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部类型</option>
          <option value="activity">活动页</option>
          <option value="product">产品页</option>
          <option value="brand">品牌页</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部状态</option>
          <option value="published">已发布</option>
          <option value="draft">草稿</option>
        </select>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="site in sites" :key="site.id" class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow">
        <div class="h-40 bg-gradient-to-br from-blue-100 to-purple-100 flex items-center justify-center">
          <span class="text-4xl text-gray-300">📄</span>
        </div>
        <div class="p-4">
          <div class="flex items-start justify-between">
            <div>
              <h3 class="font-semibold text-gray-900">{{ site.name }}</h3>
              <p class="text-sm text-gray-500 mt-1">{{ site.type }} · {{ site.template }}</p>
            </div>
            <span :class="['px-2 py-1 rounded-full text-xs font-medium',
              site.status === 'published' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
              {{ site.status === 'published' ? '已发布' : '草稿' }}
            </span>
          </div>
          <div class="mt-4 flex items-center justify-between text-sm">
            <span class="text-gray-500">PV: {{ site.pv.toLocaleString() }}</span>
            <span class="text-gray-500">转化率: {{ site.conversionRate }}%</span>
          </div>
          <div class="mt-4 flex gap-2">
            <button class="flex-1 px-3 py-1.5 text-sm border border-gray-300 text-gray-700 rounded hover:bg-gray-50">
              编辑
            </button>
            <button class="flex-1 px-3 py-1.5 text-sm border border-gray-300 text-gray-700 rounded hover:bg-gray-50">
              预览
            </button>
            <button class="px-3 py-1.5 text-sm text-gray-500 hover:text-gray-700">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="flex justify-center">
      <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
    </div>
  </div>
</template>
