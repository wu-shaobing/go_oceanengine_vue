<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 356 })
const dateRange = ref('7d')

const filterType = ref('')
const sortBy = ref('')
const searchKeyword = ref('')

const creativeData = ref([
  { id: 'CRE001', name: '产品展示视频A', type: 'video', impressions: 1560000, clicks: 62400, ctr: 4.0, conversions: 856, cost: 12560 },
  { id: 'CRE002', name: '品牌宣传图片', type: 'image', impressions: 980000, clicks: 29400, ctr: 3.0, conversions: 325, cost: 8960 },
  { id: 'CRE003', name: '促销Banner', type: 'image', impressions: 756000, clicks: 26460, ctr: 3.5, conversions: 425, cost: 6580 },
  { id: 'CRE004', name: '直播预告视频', type: 'video', impressions: 580000, clicks: 20300, ctr: 3.5, conversions: 286, cost: 4560 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleExport = () => {
  alert('导出创意报表')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表' }, { name: '创意报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">创意报表</h1>
          <p class="mt-2 text-gray-600">分析各创意素材的投放效果</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
<button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleExport">导出</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">创意总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">视频创意</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">156</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">图片创意</p>
        <p class="text-2xl font-bold text-green-600 mt-1">200</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">最高CTR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">5.2%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">3.5%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <select v-model="filterType" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部类型</option>
          <option value="video">视频</option>
          <option value="image">图片</option>
        </select>
        <select v-model="sortBy" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">按CTR排序</option>
          <option value="ctr_desc">CTR从高到低</option>
          <option value="impressions_desc">展示从高到低</option>
          <option value="conversions_desc">转化从高到低</option>
        </select>
        <input v-model="searchKeyword" type="text" placeholder="搜索创意..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创意名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="creative in creativeData" :key="creative.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ creative.name }}</div>
              <div class="text-xs text-gray-500">{{ creative.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     creative.type === 'video' ? 'bg-purple-100 text-purple-700' : 'bg-blue-100 text-blue-700']">
                {{ creative.type === 'video' ? '视频' : '图片' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (creative.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (creative.clicks / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm font-medium" :class="creative.ctr >= 3.5 ? 'text-green-600' : 'text-yellow-600'">
              {{ creative.ctr }}%
            </td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ creative.conversions }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ creative.cost.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
