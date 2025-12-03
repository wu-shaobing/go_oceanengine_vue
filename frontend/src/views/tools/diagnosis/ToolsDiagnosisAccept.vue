<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const acceptedItems = ref([
  { id: 'DA001', suggestion: '提高出价至建议值', adId: 'AD12345', acceptedAt: '2025-11-28 10:30', effect: '+25%', status: 'effective' },
  { id: 'DA002', suggestion: '更换低效素材', adId: 'AD12346', acceptedAt: '2025-11-27 15:20', effect: '+18%', status: 'effective' },
  { id: 'DA003', suggestion: '调整定向人群', adId: 'AD12347', acceptedAt: '2025-11-26 09:15', effect: '+12%', status: 'effective' },
  { id: 'DA004', suggestion: '增加预算上限', adId: 'AD12348', acceptedAt: '2025-11-25 14:00', effect: '-5%', status: 'ineffective' },
  { id: 'DA005', suggestion: '优化投放时段', adId: 'AD12349', acceptedAt: '2025-11-24 11:30', effect: '评估中', status: 'pending' }
])

const stats = ref({ total: 35, effective: 28, ineffective: 4, pending: 3 })

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '诊断工具' }, { name: '采纳记录' }]" />
      <h1 class="text-3xl font-bold text-gray-900">诊断采纳记录</h1>
      <p class="mt-2 text-gray-600">查看已采纳的优化建议及效果</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总采纳</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ stats.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">有效</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ stats.effective }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">无效</p>
        <p class="text-2xl font-bold text-red-600 mt-1">{{ stats.ineffective }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">评估中</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">{{ stats.pending }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">建议内容</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告ID</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">采纳时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">效果</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in acceptedItems" :key="item.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm text-gray-900">{{ item.suggestion }}</td>
            <td class="px-6 py-4 text-sm text-blue-600">{{ item.adId }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ item.acceptedAt }}</td>
            <td class="px-6 py-4">
              <span :class="['font-medium', 
                item.effect.startsWith('+') ? 'text-green-600' : 
                item.effect.startsWith('-') ? 'text-red-600' : 'text-gray-600']">
                {{ item.effect }}
              </span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                item.status === 'effective' ? 'bg-green-100 text-green-800' :
                item.status === 'ineffective' ? 'bg-red-100 text-red-800' : 'bg-yellow-100 text-yellow-800']">
                {{ item.status === 'effective' ? '有效' : item.status === 'ineffective' ? '无效' : '评估中' }}
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
