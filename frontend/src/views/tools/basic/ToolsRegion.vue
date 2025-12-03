<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchKeyword = ref('')
const selectedRegions = ref<string[]>([])

const regions = ref([
  { id: '110000', name: '北京市', code: 'BJ', parent: null, level: 1, children: [
    { id: '110100', name: '北京市', code: 'BJ', parent: '110000', level: 2 }
  ]},
  { id: '310000', name: '上海市', code: 'SH', parent: null, level: 1, children: [
    { id: '310100', name: '上海市', code: 'SH', parent: '310000', level: 2 }
  ]},
  { id: '440000', name: '广东省', code: 'GD', parent: null, level: 1, children: [
    { id: '440100', name: '广州市', code: 'GZ', parent: '440000', level: 2 },
    { id: '440300', name: '深圳市', code: 'SZ', parent: '440000', level: 2 },
    { id: '440600', name: '佛山市', code: 'FS', parent: '440000', level: 2 }
  ]},
  { id: '330000', name: '浙江省', code: 'ZJ', parent: null, level: 1, children: [
    { id: '330100', name: '杭州市', code: 'HZ', parent: '330000', level: 2 },
    { id: '330200', name: '宁波市', code: 'NB', parent: '330000', level: 2 }
  ]},
  { id: '320000', name: '江苏省', code: 'JS', parent: null, level: 1, children: [
    { id: '320100', name: '南京市', code: 'NJ', parent: '320000', level: 2 },
    { id: '320500', name: '苏州市', code: 'SZ', parent: '320000', level: 2 }
  ]}
])

const toggleRegion = (id: string) => {
  const index = selectedRegions.value.indexOf(id)
  if (index > -1) {
    selectedRegions.value.splice(index, 1)
  } else {
    selectedRegions.value.push(id)
  }
}

const clearSelection = () => {
  selectedRegions.value = []
}

const copyIds = () => {
  navigator.clipboard.writeText(selectedRegions.value.join(','))
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '地域定向' }]" />
      <h1 class="text-3xl font-bold text-gray-900">地域定向工具</h1>
      <p class="mt-2 text-gray-600">查询和管理广告投放地域定向数据</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Region Tree -->
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200">
        <div class="p-4 border-b border-gray-200">
          <input v-model="searchKeyword" type="text" placeholder="搜索地域..."
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div class="p-4 max-h-[600px] overflow-y-auto">
          <div v-for="province in regions" :key="province.id" class="mb-4">
            <div class="flex items-center gap-2 py-2 px-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100"
                 @click="toggleRegion(province.id)">
              <input type="checkbox" :checked="selectedRegions.includes(province.id)" class="w-4 h-4 rounded border-gray-300">
              <span class="font-medium text-gray-900">{{ province.name }}</span>
              <span class="text-xs text-gray-500 ml-auto">{{ province.code }}</span>
            </div>
            <div v-if="province.children" class="ml-6 mt-2 space-y-1">
              <div v-for="city in province.children" :key="city.id" 
                   class="flex items-center gap-2 py-1.5 px-3 rounded cursor-pointer hover:bg-gray-50"
                   @click="toggleRegion(city.id)">
                <input type="checkbox" :checked="selectedRegions.includes(city.id)" class="w-4 h-4 rounded border-gray-300">
                <span class="text-gray-700">{{ city.name }}</span>
                <span class="text-xs text-gray-400 ml-auto">ID: {{ city.id }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Selected Panel -->
      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <div class="flex items-center justify-between mb-4">
            <h3 class="font-semibold text-gray-900">已选地域</h3>
            <span class="text-sm text-gray-500">{{ selectedRegions.length }} 项</span>
          </div>
          <div v-if="selectedRegions.length" class="space-y-2 max-h-64 overflow-y-auto">
            <div v-for="id in selectedRegions" :key="id" 
                 class="flex items-center justify-between px-3 py-2 bg-blue-50 rounded-lg">
              <span class="text-sm text-blue-700">{{ id }}</span>
              <button @click="toggleRegion(id)" class="text-blue-500 hover:text-blue-700">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                </svg>
              </button>
            </div>
          </div>
          <div v-else class="text-center py-8 text-gray-500">
            暂未选择地域
          </div>
          <div v-if="selectedRegions.length" class="mt-4 flex gap-2">
            <button @click="copyIds" class="flex-1 px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700">
              复制ID
            </button>
            <button @click="clearSelection" class="px-4 py-2 border border-gray-300 text-gray-700 text-sm rounded-lg hover:bg-gray-50">
              清空
            </button>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-semibold text-gray-900 mb-3">使用说明</h3>
          <ul class="space-y-2 text-sm text-gray-600">
            <li>• 点击省份或城市进行选择</li>
            <li>• 支持多选，可跨省选择城市</li>
            <li>• 选择后可复制地域ID用于API调用</li>
            <li>• 地域ID遵循国家行政区划代码</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
