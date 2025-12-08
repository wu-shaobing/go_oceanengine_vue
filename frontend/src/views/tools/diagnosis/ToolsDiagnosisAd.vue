<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const ads = ref([
  { id: 'AD001', name: '双十一促销广告', score: 92, status: 'healthy', issues: 0, suggestions: 1 },
  { id: 'AD002', name: '新品推广广告', score: 75, status: 'warning', issues: 2, suggestions: 3 },
  { id: 'AD003', name: '品牌曝光广告', score: 45, status: 'critical', issues: 5, suggestions: 6 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleBatchDiagnose = () => {
  alert('批量诊断广告')
}

const handleDetail = (ad: typeof ads.value[0]) => {
  alert(`诊断详情: ${ad.name}`)
}

const handleOptimize = (ad: typeof ads.value[0]) => {
  alert(`一键优化: ${ad.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '诊断工具' }, { name: '广告诊断' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告诊断</h1>
          <p class="mt-2 text-gray-600">分析广告投放效果问题</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleBatchDiagnose">
          批量诊断
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">广告总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">健康</p>
        <p class="text-2xl font-bold text-green-600 mt-1">32</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">需关注</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">10</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">异常</p>
        <p class="text-2xl font-bold text-red-600 mt-1">3</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">健康评分</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">问题数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">优化建议</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="ad in ads" :key="ad.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ ad.name }}</div>
              <div class="text-xs text-gray-500">{{ ad.id }}</div>
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-16 h-2 bg-gray-200 rounded-full mr-2">
                  <div :class="['h-2 rounded-full',
                         ad.score >= 80 ? 'bg-green-500' : ad.score >= 60 ? 'bg-yellow-500' : 'bg-red-500']"
                       :style="{ width: ad.score + '%' }"></div>
                </div>
                <span class="text-sm font-medium" :class="ad.score >= 80 ? 'text-green-600' : ad.score >= 60 ? 'text-yellow-600' : 'text-red-600'">
                  {{ ad.score }}
                </span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     ad.status === 'healthy' ? 'bg-green-100 text-green-800' :
                     ad.status === 'warning' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                {{ ad.status === 'healthy' ? '健康' : ad.status === 'warning' ? '需关注' : '异常' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <span :class="ad.issues > 0 ? 'text-red-600 font-medium' : 'text-gray-500'">{{ ad.issues }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-blue-600">{{ ad.suggestions }}</td>
            <td class="px-6 py-4 text-sm">
<button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleDetail(ad)">诊断详情</button>
              <button class="text-green-600 hover:text-green-800" @click="handleOptimize(ad)">一键优化</button>
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
