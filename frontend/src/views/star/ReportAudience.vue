<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '受众分析' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">受众分析</h1>
      <p class="text-gray-600 mt-1">分析任务投放的受众画像</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <select v-model="filters.task" class="border border-gray-300 rounded px-3 py-2 w-48">
          <option value="">选择任务</option>
          <option value="1">618美妆种草任务</option>
          <option value="2">新品护肤推广</option>
        </select>
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总触达人数</div>
        <div class="text-2xl font-bold mt-1">{{ (stats.reach / 10000).toFixed(1) }}万</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">互动人数</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ (stats.interaction / 10000).toFixed(1) }}万</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化人数</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.conversion.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">人均互动成本</div>
        <div class="text-2xl font-bold mt-1">¥{{ stats.costPerInteraction }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">性别分布</h3>
        <div class="space-y-4">
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>女性</span><span>{{ audience.gender.female }}%</span>
            </div>
            <div class="bg-gray-200 rounded-full h-3">
              <div class="bg-pink-500 h-3 rounded-full" :style="{width: audience.gender.female+'%'}"></div>
            </div>
          </div>
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>男性</span><span>{{ audience.gender.male }}%</span>
            </div>
            <div class="bg-gray-200 rounded-full h-3">
              <div class="bg-blue-500 h-3 rounded-full" :style="{width: audience.gender.male+'%'}"></div>
            </div>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">年龄分布</h3>
        <div class="space-y-3">
          <div v-for="age in audience.age" :key="age.range">
            <div class="flex justify-between text-sm mb-1">
              <span>{{ age.range }}</span><span>{{ age.percent }}%</span>
            </div>
            <div class="bg-gray-200 rounded-full h-3">
              <div class="bg-blue-500 h-3 rounded-full" :style="{width: age.percent+'%'}"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">地域分布 TOP10</h3>
        <div class="space-y-2">
          <div v-for="(city, index) in audience.cities" :key="city.name" class="flex items-center">
            <span class="w-6 text-gray-500">{{ index + 1 }}</span>
            <span class="flex-1">{{ city.name }}</span>
            <span class="text-sm text-gray-500">{{ city.percent }}%</span>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">兴趣标签 TOP10</h3>
        <div class="flex flex-wrap gap-2">
          <span v-for="tag in audience.interests" :key="tag.name" class="px-3 py-1 bg-blue-50 text-blue-600 rounded-full text-sm">
            {{ tag.name }} ({{ tag.percent }}%)
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const filters = ref({ task: '', startDate: '', endDate: '' })

const stats = ref({ reach: 5680000, interaction: 286000, conversion: 3560, costPerInteraction: 0.35 })

const audience = ref({
  gender: { female: 72, male: 28 },
  age: [
    { range: '18-24岁', percent: 35 },
    { range: '25-34岁', percent: 42 },
    { range: '35-44岁', percent: 15 },
    { range: '45岁以上', percent: 8 }
  ],
  cities: [
    { name: '上海', percent: 12.5 },
    { name: '北京', percent: 10.8 },
    { name: '广州', percent: 8.6 },
    { name: '深圳', percent: 7.2 },
    { name: '杭州', percent: 5.8 },
    { name: '成都', percent: 5.2 },
    { name: '南京', percent: 4.5 },
    { name: '武汉', percent: 3.8 },
    { name: '重庆', percent: 3.5 },
    { name: '西安', percent: 3.2 }
  ],
  interests: [
    { name: '美妆', percent: 68 },
    { name: '护肤', percent: 62 },
    { name: '时尚', percent: 55 },
    { name: '穿搭', percent: 48 },
    { name: '购物', percent: 45 },
    { name: '生活', percent: 38 },
    { name: '旅行', percent: 32 },
    { name: '美食', percent: 28 },
    { name: '健身', percent: 22 },
    { name: '摄影', percent: 18 }
  ]
})
</script>
