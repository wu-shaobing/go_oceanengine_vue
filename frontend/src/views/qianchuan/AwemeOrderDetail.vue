<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '随心推' }, { name: '订单详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">订单详情</h1>
        <p class="text-gray-600 mt-1">订单ID: {{ order.id }}</p>
      </div>
      <div v-if="order.status === 'running'" class="flex space-x-3">
        <button class="px-4 py-2 border border-orange-500 text-orange-500 rounded-lg hover:bg-orange-50">停止投放</button>
      </div>
    </div>

    <!-- 核心数据 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- 左侧详情 -->
      <div class="lg:col-span-2 space-y-6">
        <!-- 数据趋势 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">数据趋势</h3>
          <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
            <span class="text-gray-400">播放/互动/消耗趋势图表</span>
          </div>
        </div>

        <!-- 订单信息 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">订单信息</h3>
          <div class="grid grid-cols-2 gap-4">
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">订单状态</div>
              <div class="font-medium">
                <span :class="order.status === 'running' ? 'text-green-600' : 'text-blue-600'">
                  {{ order.status === 'running' ? '投放中' : '已完成' }}
                </span>
              </div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">推广目标</div>
              <div class="font-medium">{{ order.goal }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">投放预算</div>
              <div class="font-medium">¥{{ order.budget }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">已消耗</div>
              <div class="font-medium text-blue-600">¥{{ order.cost }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">投放时长</div>
              <div class="font-medium">{{ order.duration }}小时</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">创建时间</div>
              <div class="font-medium">{{ order.createTime }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">预计结束</div>
              <div class="font-medium">{{ order.endTime }}</div>
            </div>
            <div class="py-2 border-b">
              <div class="text-sm text-gray-500">投放进度</div>
              <div class="font-medium">{{ Math.round(order.cost / order.budget * 100) }}%</div>
            </div>
          </div>
        </div>

        <!-- 定向设置 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">定向设置</h3>
          <div class="space-y-2">
            <div class="flex"><span class="text-gray-500 w-20">地域:</span><span>{{ order.targeting.region }}</span></div>
            <div class="flex"><span class="text-gray-500 w-20">性别:</span><span>{{ order.targeting.gender }}</span></div>
            <div class="flex"><span class="text-gray-500 w-20">年龄:</span><span>{{ order.targeting.age }}</span></div>
          </div>
        </div>
      </div>

      <!-- 右侧视频预览 -->
      <div class="space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">推广视频</h3>
          <div class="bg-gray-100 rounded-lg overflow-hidden">
            <div class="aspect-[9/16] bg-gray-200 flex items-center justify-center">
              <img :src="order.video.cover" alt="" class="w-full h-full object-cover">
            </div>
          </div>
          <div class="mt-4">
            <div class="font-medium line-clamp-2">{{ order.video.title }}</div>
            <div class="text-sm text-gray-500 mt-2">
              <div>发布时间: {{ order.video.publishTime }}</div>
              <div>原始播放: {{ order.video.originalPlays }}</div>
            </div>
          </div>
        </div>

        <!-- 效果对比 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">效果对比</h3>
          <div class="space-y-3">
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-500">播放量</span>
                <span class="text-green-600">+{{ order.effect.playsIncrease }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" style="width: 75%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-500">点赞数</span>
                <span class="text-green-600">+{{ order.effect.likesIncrease }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" style="width: 60%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-500">评论数</span>
                <span class="text-green-600">+{{ order.effect.commentsIncrease }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" style="width: 45%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-500">新增粉丝</span>
                <span class="text-green-600">+{{ order.effect.fansIncrease }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" style="width: 30%"></div>
              </div>
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
  { label: '播放量', value: '25,600' },
  { label: '点赞', value: '1,580' },
  { label: '评论', value: '256' },
  { label: '分享', value: '89' },
  { label: '新增粉丝', value: '126' },
  { label: '商品点击', value: '368' }
])

const order = ref({
  id: 'AO001',
  status: 'running',
  goal: '商品购买',
  budget: 500,
  cost: 380,
  duration: 24,
  createTime: '2024-03-15 10:00:00',
  endTime: '2024-03-16 10:00:00',
  targeting: {
    region: '全国',
    gender: '女',
    age: '18-40岁'
  },
  video: {
    title: '新品美妆测评分享 超好用必入!',
    cover: 'https://via.placeholder.com/270x480',
    publishTime: '2024-03-14',
    originalPlays: '8,560'
  },
  effect: {
    playsIncrease: '17,040',
    likesIncrease: '1,256',
    commentsIncrease: '198',
    fansIncrease: '126'
  }
})
</script>
