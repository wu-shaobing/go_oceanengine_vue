<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const behaviors = ref([
  { id: 'BH001', name: '商品浏览', eventType: 'view_product', count: 125600, users: 45600, status: 'active' },
  { id: 'BH002', name: '加入购物车', eventType: 'add_cart', count: 35800, users: 18900, status: 'active' },
  { id: 'BH003', name: '发起结算', eventType: 'checkout', count: 12500, users: 8900, status: 'active' },
  { id: 'BH004', name: '完成购买', eventType: 'purchase', count: 8900, users: 6700, status: 'active' },
  { id: 'BH005', name: '商品收藏', eventType: 'favorite', count: 28900, users: 15600, status: 'paused' }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddBehavior = () => {
  alert('添加行为')
}

const handleEditBehavior = (behavior: typeof behaviors.value[0]) => {
  alert(`编辑行为: ${behavior.name}`)
}

const handleViewBehaviorData = (behavior: typeof behaviors.value[0]) => {
  alert(`查看数据: ${behavior.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品广告' }, { name: '行为设置' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">DPA行为追踪</h1>
          <p class="mt-2 text-gray-600">配置商品行为追踪事件</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleAddBehavior">
          添加行为
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">行为类型</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日事件</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">21.2万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃用户</p>
        <p class="text-2xl font-bold text-green-600 mt-1">9.6万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">7.1%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">行为名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">事件数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">用户数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="behavior in behaviors" :key="behavior.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ behavior.name }}</div>
              <div class="text-xs text-gray-500">{{ behavior.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-700">{{ behavior.eventType }}</code>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-blue-600">{{ formatNumber(behavior.count) }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ formatNumber(behavior.users) }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     behavior.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ behavior.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleEditBehavior(behavior)">编辑</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleViewBehaviorData(behavior)">数据</button>
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
