<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')

const topMaterials = ref([
  { id: 'M001', name: '双11促销视频A', type: 'video', plays: 1250000, ctr: 4.5, cvr: 3.2, cost: 0.85 },
  { id: 'M002', name: '产品展示图-红色', type: 'image', plays: 890000, ctr: 3.8, cvr: 2.8, cost: 0.92 },
  { id: 'M003', name: '品牌故事视频', type: 'video', plays: 756000, ctr: 4.2, cvr: 2.5, cost: 1.05 },
  { id: 'M004', name: '限时优惠Banner', type: 'image', plays: 680000, ctr: 3.5, cvr: 3.0, cost: 0.88 },
  { id: 'M005', name: '用户评价合集', type: 'video', plays: 590000, ctr: 4.8, cvr: 3.5, cost: 0.78 }
])

const stats = ref({
  totalMaterials: 256,
  activeMaterials: 89,
  avgCtr: 3.8,
  avgCvr: 2.9
})

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handleViewDetail = (item: typeof topMaterials.value[0]) => {
  alert(`查看详情: ${item.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '素材分析' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">素材效果分析</h1>
          <p class="mt-2 text-gray-600">分析素材投放效果，优化创意策略</p>
        </div>
        <div class="flex gap-2">
          <button v-for="d in ['7d', '14d', '30d']" :key="d"
                  :class="['px-4 py-2 rounded-lg', dateRange === d ? 'bg-blue-600 text-white' : 'bg-gray-100']"
                  @click="dateRange = d">
            {{ d === '7d' ? '近7天' : d === '14d' ? '近14天' : '近30天' }}
          </button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">素材总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ stats.totalMaterials }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">投放中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ stats.activeMaterials }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ stats.avgCtr }}%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CVR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">{{ stats.avgCvr }}%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="font-semibold text-gray-900 mb-4">效果TOP素材</h3>
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">素材</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CVR</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CPC</th>
            <th class="px-4 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in topMaterials" :key="item.id" class="hover:bg-gray-50">
            <td class="px-4 py-3">
              <div class="flex items-center gap-3">
                <span class="text-2xl">{{ item.type === 'video' ? '🎬' : '🖼️' }}</span>
                <span class="text-sm font-medium text-gray-900">{{ item.name }}</span>
              </div>
            </td>
            <td class="px-4 py-3">
              <span :class="['px-2 py-1 rounded text-xs',
                     item.type === 'video' ? 'bg-purple-100 text-purple-700' : 'bg-blue-100 text-blue-700']">
                {{ item.type === 'video' ? '视频' : '图片' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-900">{{ formatNumber(item.plays) }}</td>
            <td class="px-4 py-3 text-sm font-medium text-green-600">{{ item.ctr }}%</td>
            <td class="px-4 py-3 text-sm font-medium text-blue-600">{{ item.cvr }}%</td>
            <td class="px-4 py-3 text-sm text-gray-900">¥{{ item.cost }}</td>
            <td class="px-4 py-3">
              <button class="text-sm text-blue-600 hover:text-blue-800" @click="handleViewDetail(item)">详情</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
