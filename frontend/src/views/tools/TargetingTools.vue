<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const activeTab = ref('geo')

const tabs = [
  { id: 'geo', name: '地域定向', icon: 'M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z' },
  { id: 'interest', name: '兴趣定向', icon: 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z' },
  { id: 'behavior', name: '行为定向', icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z' },
  { id: 'device', name: '设备定向', icon: 'M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z' },
  { id: 'custom', name: '自定义人群', icon: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z' }
]

const geoData = reactive({
  selectedRegions: ['北京', '上海', '广州', '深圳'],
  regions: ['北京', '上海', '广州', '深圳', '杭州', '南京', '成都', '武汉', '西安', '重庆']
})

const interestCategories = [
  { id: 'shopping', name: '购物', count: 1580 },
  { id: 'games', name: '游戏', count: 2340 },
  { id: 'travel', name: '旅游', count: 980 },
  { id: 'food', name: '美食', count: 1250 },
  { id: 'sports', name: '运动', count: 890 },
  { id: 'tech', name: '科技', count: 1670 },
  { id: 'fashion', name: '时尚', count: 1120 },
  { id: 'finance', name: '金融', count: 760 }
]

const selectedInterests = ref<string[]>(['shopping', 'tech'])

const deviceOptions = reactive({
  os: ['iOS', 'Android'],
  network: ['WiFi', '4G', '5G'],
  brand: ['Apple', 'Huawei', 'Xiaomi', 'OPPO', 'VIVO', 'Samsung']
})

const selectedDevice = reactive({
  os: ['iOS', 'Android'],
  network: ['WiFi', '4G', '5G'],
  brand: [] as string[]
})

const toggleRegion = (region: string) => {
  const index = geoData.selectedRegions.indexOf(region)
  if (index > -1) {
    geoData.selectedRegions.splice(index, 1)
  } else {
    geoData.selectedRegions.push(region)
  }
}

const toggleInterest = (id: string) => {
  const index = selectedInterests.value.indexOf(id)
  if (index > -1) {
    selectedInterests.value.splice(index, 1)
  } else {
    selectedInterests.value.push(id)
  }
}

const toggleDevice = (type: 'os' | 'network' | 'brand', value: string) => {
  const arr = selectedDevice[type]
  const index = arr.indexOf(value)
  if (index > -1) {
    arr.splice(index, 1)
  } else {
    arr.push(value)
  }
}

const handleUpload = () => {
  alert('上传人群包')
}

const handleCreateRule = () => {
  alert('创建规则')
}

const handleReset = () => {
  geoData.selectedRegions = []
  selectedInterests.value = []
  selectedDevice.os = ['iOS', 'Android']
  selectedDevice.network = ['WiFi', '4G', '5G']
  selectedDevice.brand = []
  alert('已重置所有定向')
}

const handleSave = () => {
  alert('保存定向配置')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '定向工具' }]" />
      <h1 class="text-3xl font-bold text-gray-900">定向工具</h1>
      <p class="mt-2 text-gray-600">配置广告投放的目标受众</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <!-- Tabs -->
      <div class="border-b border-gray-200">
        <nav class="flex -mb-px">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            class="flex items-center gap-2 px-6 py-4 border-b-2 text-sm font-medium transition-colors"
            :class="{
              'border-blue-500 text-blue-600': activeTab === tab.id,
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== tab.id
            }"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="tab.icon"/>
            </svg>
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <div class="p-6">
        <!-- 地域定向 -->
        <div v-show="activeTab === 'geo'" class="space-y-6">
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">选择投放地域</h3>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="region in geoData.regions"
                :key="region"
                @click="toggleRegion(region)"
                class="px-4 py-2 rounded-lg border transition-colors"
                :class="{
                  'bg-blue-50 border-blue-500 text-blue-700': geoData.selectedRegions.includes(region),
                  'border-gray-300 text-gray-700 hover:bg-gray-50': !geoData.selectedRegions.includes(region)
                }"
              >
                {{ region }}
              </button>
            </div>
          </div>
          
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-600">
              已选择 <span class="font-medium text-gray-900">{{ geoData.selectedRegions.length }}</span> 个地区，
              预估覆盖人群 <span class="font-medium text-blue-600">2,580万</span>
            </p>
          </div>
        </div>

        <!-- 兴趣定向 -->
        <div v-show="activeTab === 'interest'" class="space-y-6">
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">选择兴趣标签</h3>
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
              <div
                v-for="interest in interestCategories"
                :key="interest.id"
                @click="toggleInterest(interest.id)"
                class="p-4 rounded-lg border cursor-pointer transition-colors"
                :class="{
                  'bg-blue-50 border-blue-500': selectedInterests.includes(interest.id),
                  'border-gray-200 hover:border-gray-300': !selectedInterests.includes(interest.id)
                }"
              >
                <p class="font-medium" :class="{ 'text-blue-700': selectedInterests.includes(interest.id) }">
                  {{ interest.name }}
                </p>
                <p class="text-sm text-gray-500 mt-1">{{ interest.count.toLocaleString() }}万人</p>
              </div>
            </div>
          </div>
        </div>

        <!-- 行为定向 -->
        <div v-show="activeTab === 'behavior'" class="space-y-6">
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">用户行为筛选</h3>
            <div class="space-y-4">
              <div class="p-4 border border-gray-200 rounded-lg">
                <label class="flex items-center justify-between">
                  <span class="text-gray-700">近期有购物行为</span>
                  <input type="checkbox" class="w-5 h-5 text-blue-600 rounded">
                </label>
                <p class="text-sm text-gray-500 mt-1">最近7天内有电商购物行为的用户</p>
              </div>
              <div class="p-4 border border-gray-200 rounded-lg">
                <label class="flex items-center justify-between">
                  <span class="text-gray-700">活跃App用户</span>
                  <input type="checkbox" class="w-5 h-5 text-blue-600 rounded" checked>
                </label>
                <p class="text-sm text-gray-500 mt-1">日均使用App超过2小时的用户</p>
              </div>
              <div class="p-4 border border-gray-200 rounded-lg">
                <label class="flex items-center justify-between">
                  <span class="text-gray-700">高消费能力</span>
                  <input type="checkbox" class="w-5 h-5 text-blue-600 rounded">
                </label>
                <p class="text-sm text-gray-500 mt-1">月均消费超过5000元的用户</p>
              </div>
            </div>
          </div>
        </div>

        <!-- 设备定向 -->
        <div v-show="activeTab === 'device'" class="space-y-6">
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">操作系统</h3>
            <div class="flex gap-3">
              <button
                v-for="os in deviceOptions.os"
                :key="os"
                @click="toggleDevice('os', os)"
                class="px-4 py-2 rounded-lg border transition-colors"
                :class="{
                  'bg-blue-50 border-blue-500 text-blue-700': selectedDevice.os.includes(os),
                  'border-gray-300 text-gray-700': !selectedDevice.os.includes(os)
                }"
              >
                {{ os }}
              </button>
            </div>
          </div>

          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">网络类型</h3>
            <div class="flex gap-3">
              <button
                v-for="network in deviceOptions.network"
                :key="network"
                @click="toggleDevice('network', network)"
                class="px-4 py-2 rounded-lg border transition-colors"
                :class="{
                  'bg-blue-50 border-blue-500 text-blue-700': selectedDevice.network.includes(network),
                  'border-gray-300 text-gray-700': !selectedDevice.network.includes(network)
                }"
              >
                {{ network }}
              </button>
            </div>
          </div>

          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">设备品牌（可选）</h3>
            <div class="flex flex-wrap gap-3">
              <button
                v-for="brand in deviceOptions.brand"
                :key="brand"
                @click="toggleDevice('brand', brand)"
                class="px-4 py-2 rounded-lg border transition-colors"
                :class="{
                  'bg-blue-50 border-blue-500 text-blue-700': selectedDevice.brand.includes(brand),
                  'border-gray-300 text-gray-700': !selectedDevice.brand.includes(brand)
                }"
              >
                {{ brand }}
              </button>
            </div>
          </div>
        </div>

        <!-- 自定义人群 -->
        <div v-show="activeTab === 'custom'" class="space-y-6">
          <div class="text-center py-8">
            <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/>
            </svg>
            <h3 class="mt-2 text-lg font-medium text-gray-900">自定义人群包</h3>
            <p class="mt-1 text-gray-500">上传自定义人群数据或创建人群规则</p>
<div class="mt-6 flex justify-center gap-3">
              <button @click="handleUpload" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
                上传人群包
              </button>
              <button @click="handleCreateRule" class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
                创建规则
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="px-6 py-4 bg-gray-50 border-t border-gray-200 flex items-center justify-between">
        <p class="text-sm text-gray-600">
          预估覆盖人群: <span class="font-medium text-blue-600">1,850万</span>
        </p>
<div class="flex gap-3">
          <button @click="handleReset" class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
            重置
          </button>
          <button @click="handleSave" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            保存定向
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
