<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 56 })
const checkUrl = ref('')

const checkResults = ref([
  { id: 'AC001', url: 'https://example.com/page1', anchorCount: 3, validCount: 3, status: 'pass', checkedAt: '2025-11-28 10:30' },
  { id: 'AC002', url: 'https://example.com/page2', anchorCount: 5, validCount: 4, status: 'warning', checkedAt: '2025-11-28 10:25' },
  { id: 'AC003', url: 'https://example.com/page3', anchorCount: 2, validCount: 0, status: 'fail', checkedAt: '2025-11-28 10:20' },
  { id: 'AC004', url: 'https://example.com/page4', anchorCount: 4, validCount: 4, status: 'pass', checkedAt: '2025-11-28 10:15' }
])

const stats = ref({ totalChecks: 56, passRate: 85, avgAnchors: 3.5 })

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleStartCheck = () => {
  alert('开始检测')
}

const handleDetail = (item: typeof checkResults.value[0]) => {
  alert(`查看详情: ${item.url}`)
}

const handleRecheck = (item: typeof checkResults.value[0]) => {
  alert(`重新检测: ${item.url}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '锚点检测' }]" />
      <h1 class="text-3xl font-bold text-gray-900">锚点检测</h1>
      <p class="mt-2 text-gray-600">检测落地页锚点配置是否正确</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="checkUrl" type="url" placeholder="输入落地页URL进行检测..."
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <button @click="handleStartCheck" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">开始检测</button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总检测次数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ stats.totalChecks }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">通过率</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ stats.passRate }}%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均锚点数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ stats.avgAnchors }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">URL</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">锚点数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">有效数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">检测时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in checkResults" :key="item.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <a :href="item.url" target="_blank" class="text-sm text-blue-600 hover:underline">
                {{ item.url }}
              </a>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ item.anchorCount }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ item.validCount }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     item.status === 'pass' ? 'bg-green-100 text-green-800' :
                     item.status === 'warning' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                {{ item.status === 'pass' ? '通过' : item.status === 'warning' ? '警告' : '失败' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ item.checkedAt }}</td>
            <td class="px-6 py-4 text-sm">
<button @click="handleDetail(item)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button @click="handleRecheck(item)" class="text-gray-600 hover:text-gray-800">重新检测</button>
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
