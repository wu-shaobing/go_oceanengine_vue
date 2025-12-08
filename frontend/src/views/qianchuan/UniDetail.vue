<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '全域推广' }, { name: '推广详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ detail.name }}</h1>
        <p class="text-gray-600 mt-1">推广ID: {{ detail.id }}</p>
      </div>
      <div class="flex space-x-3">
<button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleEdit">编辑</button>
        <button v-if="detail.status === 'running'" class="px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600" @click="handlePause">暂停</button>
        <button v-else class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600" @click="handleStart">启动</button>
      </div>
    </div>

    <!-- 核心数据 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend >= 0 ? 'text-green-500' : 'text-red-500'" class="text-sm">
          {{ stat.trend >= 0 ? '+' : '' }}{{ stat.trend }}%
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <!-- 数据趋势 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">数据趋势</h3>
          <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
            <span class="text-gray-400">消耗/GMV/ROI趋势图表</span>
          </div>
        </div>

        <!-- 素材效果 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">素材效果排行</h3>
          <table class="min-w-full">
            <thead>
              <tr class="border-b">
                <th class="text-left py-2 text-sm font-medium text-gray-500">素材</th>
                <th class="text-right py-2 text-sm font-medium text-gray-500">消耗</th>
                <th class="text-right py-2 text-sm font-medium text-gray-500">转化</th>
                <th class="text-right py-2 text-sm font-medium text-gray-500">ROI</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="material in topMaterials" :key="material.id" class="border-b">
                <td class="py-2">
                  <div class="flex items-center">
                    <img :src="material.cover" class="w-12 h-8 rounded mr-2 object-cover" alt="">
                    <span class="text-sm">{{ material.name }}</span>
                  </div>
                </td>
                <td class="py-2 text-sm text-right">¥{{ material.cost.toLocaleString() }}</td>
                <td class="py-2 text-sm text-right">{{ material.conversions }}</td>
                <td class="py-2 text-sm text-right font-medium">{{ material.roi }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="space-y-6">
        <!-- 推广设置 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">推广设置</h3>
          <div class="space-y-3">
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">状态</span>
              <span :class="detail.status === 'running' ? 'text-green-600' : 'text-gray-600'" class="font-medium">
                {{ detail.status === 'running' ? '投放中' : '已暂停' }}
              </span>
            </div>
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">类型</span>
              <span class="font-medium">{{ detail.type === 'live' ? '直播' : '短视频' }}</span>
            </div>
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">日预算</span>
              <span class="font-medium">¥{{ detail.budget.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">ROI目标</span>
              <span class="font-medium">{{ detail.roiTarget }}</span>
            </div>
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">实际ROI</span>
              <span :class="detail.actualRoi >= detail.roiTarget ? 'text-green-600' : 'text-red-600'" class="font-medium">
                {{ detail.actualRoi }}
              </span>
            </div>
            <div class="flex justify-between py-2 border-b">
              <span class="text-gray-500">创建时间</span>
              <span class="font-medium">{{ detail.createTime }}</span>
            </div>
          </div>
        </div>

        <!-- 关联店铺 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">关联店铺</h3>
          <div class="flex items-center">
            <img :src="detail.shop.logo" class="w-12 h-12 rounded-lg mr-3" alt="">
            <div>
              <div class="font-medium">{{ detail.shop.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ detail.shop.id }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const coreStats = ref([
  { label: '消耗', value: '¥28,600', trend: 15.2 },
  { label: 'GMV', value: '¥156,800', trend: 22.5 },
  { label: 'ROI', value: '5.48', trend: 8.3 },
  { label: '转化数', value: '568', trend: 18.6 },
  { label: '曝光', value: '1.2M', trend: 12.1 }
])

const detail = ref({
  id: 'U001',
  name: '618大促直播全域推广',
  status: 'running',
  type: 'live',
  budget: 30000,
  roiTarget: 4.5,
  actualRoi: 5.48,
  createTime: '2024-03-10 10:00:00',
  shop: {
    id: '7001',
    name: '品牌官方旗舰店',
    logo: 'https://via.placeholder.com/48'
  }
})

const topMaterials = ref([
  { id: 1, name: '直播片段1', cover: 'https://via.placeholder.com/96x64', cost: 8600, conversions: 186, roi: '6.2' },
  { id: 2, name: '直播片段2', cover: 'https://via.placeholder.com/96x64', cost: 6800, conversions: 145, roi: '5.8' },
  { id: 3, name: '商品展示', cover: 'https://via.placeholder.com/96x64', cost: 5400, conversions: 98, roi: '4.9' },
  { id: 4, name: '达人推荐', cover: 'https://via.placeholder.com/96x64', cost: 4200, conversions: 86, roi: '5.1' }
])

const handleEdit = () => {
  alert(`编辑推广: ${detail.value.name}`)
}

const handlePause = () => {
  alert(`暂停推广: ${detail.value.name}`)
}

const handleStart = () => {
  alert(`启动推广: ${detail.value.name}`)
}
</script>
