<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 22 })

const strategies = ref([
  { id: 'S001', name: '新品冷启动', target: '日耗1万', duration: '7天', status: 'active', progress: 68, currentSpend: 6800 },
  { id: 'S002', name: '爆款加速', target: '日耗5万', duration: '3天', status: 'completed', progress: 100, currentSpend: 50000 },
  { id: 'S003', name: '测试期起量', target: '日耗5000', duration: '5天', status: 'paused', progress: 35, currentSpend: 1750 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateStrategy = () => {
  alert('创建策略')
}

const handleStrategyDetail = (strategy: typeof strategies.value[0]) => {
  alert(`查看详情: ${strategy.name}`)
}

const handlePause = (strategy: typeof strategies.value[0]) => {
  strategy.status = 'paused'
  alert(`已暂停: ${strategy.name}`)
}

const handleResume = (strategy: typeof strategies.value[0]) => {
  strategy.status = 'active'
  alert(`已恢复: ${strategy.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '起量工具' }, { name: '起量策略' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">起量策略</h1>
          <p class="mt-2 text-gray-600">管理广告起量策略</p>
        </div>
<button @click="handleCreateStrategy" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建策略
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">策略总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">执行中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">成功率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">85%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">策略名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">周期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">当前消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="strategy in strategies" :key="strategy.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ strategy.name }}</div>
              <div class="text-xs text-gray-500">{{ strategy.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ strategy.target }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ strategy.duration }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-24 h-2 bg-gray-200 rounded-full mr-2">
                  <div class="h-2 bg-blue-600 rounded-full" :style="{ width: strategy.progress + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ strategy.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ strategy.currentSpend.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     strategy.status === 'active' ? 'bg-blue-100 text-blue-800' :
                     strategy.status === 'completed' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ strategy.status === 'active' ? '执行中' : strategy.status === 'completed' ? '已完成' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleStrategyDetail(strategy)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button v-if="strategy.status === 'active'" @click="handlePause(strategy)" class="text-yellow-600 hover:text-yellow-800">暂停</button>
              <button v-if="strategy.status === 'paused'" @click="handleResume(strategy)" class="text-green-600 hover:text-green-800">恢复</button>
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
