<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '广告详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ adDetail.name }}</h1>
        <p class="text-gray-600 mt-1">广告ID: {{ adDetail.id }}</p>
      </div>
      <div class="flex space-x-3">
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">编辑</button>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">复制广告</button>
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
      <!-- 广告信息 -->
      <div class="lg:col-span-2">
        <div class="bg-white rounded-lg shadow p-6 mb-6">
          <h3 class="text-lg font-medium mb-4">广告设置</h3>
          <div class="grid grid-cols-2 gap-4">
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">广告状态</div>
              <div class="font-medium">
                <span :class="adDetail.status === 'delivery' ? 'text-green-600' : 'text-gray-600'">
                  {{ adDetail.status === 'delivery' ? '投放中' : '已暂停' }}
                </span>
              </div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">所属计划</div>
              <div class="font-medium">{{ adDetail.campaignName }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">出价方式</div>
              <div class="font-medium">{{ adDetail.bidType }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">转化出价</div>
              <div class="font-medium">¥{{ adDetail.bid }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">转化目标</div>
              <div class="font-medium">{{ adDetail.convertType }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">ROI目标</div>
              <div class="font-medium">{{ adDetail.roiGoal }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">投放时间</div>
              <div class="font-medium">{{ adDetail.schedule }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">创建时间</div>
              <div class="font-medium">{{ adDetail.createTime }}</div>
            </div>
          </div>
        </div>

        <!-- 定向信息 -->
        <div class="bg-white rounded-lg shadow p-6 mb-6">
          <h3 class="text-lg font-medium mb-4">定向设置</h3>
          <div class="space-y-3">
            <div class="flex">
              <span class="text-gray-500 w-24">地域:</span>
              <span>{{ adDetail.targeting.region }}</span>
            </div>
            <div class="flex">
              <span class="text-gray-500 w-24">年龄:</span>
              <span>{{ adDetail.targeting.age }}</span>
            </div>
            <div class="flex">
              <span class="text-gray-500 w-24">性别:</span>
              <span>{{ adDetail.targeting.gender }}</span>
            </div>
            <div class="flex">
              <span class="text-gray-500 w-24">兴趣标签:</span>
              <div class="flex flex-wrap gap-1">
                <span v-for="tag in adDetail.targeting.interests" :key="tag" class="px-2 py-0.5 bg-blue-100 text-blue-800 text-sm rounded">
                  {{ tag }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- 数据趋势 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">数据趋势</h3>
          <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
            <span class="text-gray-400">消耗/转化趋势图表</span>
          </div>
        </div>
      </div>

      <!-- 创意预览 -->
      <div class="lg:col-span-1">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">创意预览</h3>
          <div class="bg-gray-100 rounded-lg overflow-hidden">
            <div class="aspect-[9/16] bg-gray-200 flex items-center justify-center">
              <img :src="adDetail.creative.cover" alt="" class="w-full h-full object-cover">
            </div>
          </div>
          <div class="mt-4">
            <div class="text-sm text-gray-500 mb-1">广告文案</div>
            <p class="text-gray-700">{{ adDetail.creative.copywriting }}</p>
          </div>
          <div class="mt-4">
            <div class="text-sm text-gray-500 mb-1">素材信息</div>
            <div class="text-sm">
              <div>时长: {{ adDetail.creative.duration }}秒</div>
              <div>尺寸: {{ adDetail.creative.resolution }}</div>
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
  { label: '消耗', value: '¥3,560', trend: 12.5 },
  { label: '曝光', value: '156,000', trend: 8.3 },
  { label: '点击', value: '4,580', trend: 15.2 },
  { label: '转化', value: '128', trend: 22.1 },
  { label: 'ROI', value: '3.85', trend: 5.6 }
])

const adDetail = ref({
  id: '20001',
  name: '直播间引流-美妆专场',
  status: 'delivery',
  campaignName: '618大促直播计划',
  bidType: 'OCPM',
  bid: 25,
  convertType: '成交',
  roiGoal: 3.5,
  schedule: '长期投放',
  createTime: '2024-03-10 14:30:00',
  targeting: {
    region: '全国',
    age: '18-40岁',
    gender: '女',
    interests: ['美妆护肤', '时尚穿搭', '生活方式']
  },
  creative: {
    cover: 'https://via.placeholder.com/270x480',
    copywriting: '限时特惠！爆款美妆套装低至5折，点击进入直播间抢购~',
    duration: 15,
    resolution: '720x1280'
  }
})
</script>
