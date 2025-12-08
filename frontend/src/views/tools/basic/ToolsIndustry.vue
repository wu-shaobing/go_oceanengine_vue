<script setup lang="ts">
import { ref, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchKeyword = ref('')
const expandedCategories = ref<string[]>(['100', '200', '300'])

const industries = ref([
  { id: '100', name: '电商零售', icon: '🛒', children: [
    { id: '101', name: '综合电商' },
    { id: '102', name: '垂直电商' },
    { id: '103', name: '跨境电商' },
    { id: '104', name: '社交电商' }
  ]},
  { id: '200', name: '游戏', icon: '🎮', children: [
    { id: '201', name: '手机游戏' },
    { id: '202', name: '网页游戏' },
    { id: '203', name: '主机游戏' },
    { id: '204', name: '电竞' }
  ]},
  { id: '300', name: '教育培训', icon: '📚', children: [
    { id: '301', name: 'K12教育' },
    { id: '302', name: '职业教育' },
    { id: '303', name: '语言培训' },
    { id: '304', name: '兴趣培训' }
  ]},
  { id: '400', name: '金融', icon: '💰', children: [
    { id: '401', name: '银行' },
    { id: '402', name: '保险' },
    { id: '403', name: '证券' },
    { id: '404', name: '互联网金融' }
  ]},
  { id: '500', name: '汽车', icon: '🚗', children: [
    { id: '501', name: '整车' },
    { id: '502', name: '汽车配件' },
    { id: '503', name: '汽车服务' },
    { id: '504', name: '新能源汽车' }
  ]},
  { id: '600', name: '房产家居', icon: '🏠', children: [
    { id: '601', name: '房地产开发' },
    { id: '602', name: '房产中介' },
    { id: '603', name: '家居建材' },
    { id: '604', name: '家装服务' }
  ]}
])

const filteredIndustries = computed(() => {
  if (!searchKeyword.value) return industries.value
  const keyword = searchKeyword.value.toLowerCase()
  return industries.value.filter(cat => 
    cat.name.toLowerCase().includes(keyword) ||
    cat.children.some(c => c.name.toLowerCase().includes(keyword))
  )
})

const toggleExpand = (id: string) => {
  const index = expandedCategories.value.indexOf(id)
  if (index > -1) {
    expandedCategories.value.splice(index, 1)
  } else {
    expandedCategories.value.push(id)
  }
}

const handleCopyId = (id: string) => {
  navigator.clipboard.writeText(id)
  alert(`已复制ID: ${id}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '行业分类' }]" />
      <h1 class="text-3xl font-bold text-gray-900">行业分类工具</h1>
      <p class="mt-2 text-gray-600">查询广告投放行业分类及对应代码</p>
    </div>

    <!-- Search -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <input v-model="searchKeyword" type="text" placeholder="搜索行业..."
             class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
    </div>

    <!-- Industry Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="category in filteredIndustries" :key="category.id" 
           class="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <div class="p-4 border-b border-gray-100 bg-gray-50 cursor-pointer hover:bg-gray-100"
             @click="toggleExpand(category.id)">
          <div class="flex items-center gap-3">
            <span class="text-2xl">{{ category.icon }}</span>
            <div class="flex-1">
              <h3 class="font-semibold text-gray-900">{{ category.name }}</h3>
              <p class="text-xs text-gray-500">ID: {{ category.id }}</p>
            </div>
            <svg :class="['w-5 h-5 text-gray-400 transition-transform', 
                         expandedCategories.includes(category.id) ? 'rotate-180' : '']" 
                 fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </div>
        </div>
        <div v-if="expandedCategories.includes(category.id)" class="p-2">
          <div v-for="child in category.children" :key="child.id"
               class="flex items-center justify-between px-3 py-2 rounded hover:bg-gray-50 cursor-pointer group">
            <span class="text-sm text-gray-700">{{ child.name }}</span>
            <div class="flex items-center gap-2">
              <span class="text-xs text-gray-400">{{ child.id }}</span>
<button @click.stop="handleCopyId(child.id)" class="opacity-0 group-hover:opacity-100 text-blue-500 text-xs">复制</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- API Info -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="font-semibold text-gray-900 mb-4">API 使用说明</h3>
      <div class="bg-gray-900 rounded-lg p-4 text-sm font-mono text-green-400 overflow-x-auto">
        <pre>GET /2/tools/industry/get/
Response:
{
  "code": 0,
  "message": "OK",
  "data": {
    "list": [
      { "industry_id": "100", "industry_name": "电商零售", "level": 1 },
      { "industry_id": "101", "industry_name": "综合电商", "level": 2, "parent_id": "100" }
    ]
  }
}</pre>
      </div>
    </div>
  </div>
</template>
