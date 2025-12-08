<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '线索管理' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">线索管理</h1>
      <p class="text-gray-600 mt-1">管理达人合作产生的线索</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="pending">待跟进</option>
          <option value="following">跟进中</option>
          <option value="converted">已转化</option>
        </select>
<button @click="handleQuery" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总线索</div>
        <div class="text-xl font-bold mt-1">{{ stats.total }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">待跟进</div>
        <div class="text-xl font-bold text-orange-600 mt-1">{{ stats.pending }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已转化</div>
        <div class="text-xl font-bold text-green-600 mt-1">{{ stats.converted }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化率</div>
        <div class="text-xl font-bold mt-1">{{ stats.conversionRate }}%</div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">线索信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">来源</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">创建时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="clue in clues" :key="clue.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ clue.name }}</div>
              <div class="text-sm text-gray-500">{{ clue.phone }}</div>
            </td>
            <td class="px-4 py-3 text-sm">
              <div>达人: {{ clue.influencer }}</div>
              <div class="text-gray-500">任务: {{ clue.task }}</div>
            </td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(clue.status)" class="px-2 py-1 text-xs rounded">{{ getStatusText(clue.status) }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ clue.createTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
<button @click="handleDetail(clue)" class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button @click="handleFollow(clue)" class="text-blue-600 hover:text-blue-800 text-sm">跟进</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', status: '' })

const stats = ref({ total: 856, pending: 125, converted: 586, conversionRate: 68.5 })

const clues = ref([
  { id: 1, name: '张先生', phone: '138****1234', influencer: '美妆达人小美', task: '618种草任务', status: 'pending', createTime: '2024-06-18 15:30' },
  { id: 2, name: '李女士', phone: '139****5678', influencer: '时尚博主Amy', task: '新品推广', status: 'following', createTime: '2024-06-18 14:20' },
  { id: 3, name: '王先生', phone: '137****9012', influencer: '生活家小王', task: '品牌故事', status: 'converted', createTime: '2024-06-17 18:45' }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = { pending: 'bg-orange-100 text-orange-800', following: 'bg-blue-100 text-blue-800', converted: 'bg-green-100 text-green-800' }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = { pending: '待跟进', following: '跟进中', converted: '已转化' }
  return texts[status] || status
}

const handleQuery = () => {
  alert('查询线索')
}

const handleExport = () => {
  alert('导出线索数据')
}

const handleDetail = (clue: typeof clues.value[0]) => {
  alert(`查看线索详情: ${clue.name}`)
}

const handleFollow = (clue: typeof clues.value[0]) => {
  alert(`跟进线索: ${clue.name}`)
}
</script>
