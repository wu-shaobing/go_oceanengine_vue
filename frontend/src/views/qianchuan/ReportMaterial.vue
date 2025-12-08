<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '素材报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">素材报表</h1>
      <p class="text-gray-600 mt-1">查看素材维度的投放效果数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频素材</option>
          <option value="image">图片素材</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700" @click="handleQuery">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50" @click="handleExport">导出</button>
      </div>
    </div>

    <!-- 素材表格 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">素材</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">展示</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">CTR</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化成本</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">关联广告</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in materials" :key="item.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="item.cover" class="w-16 h-10 rounded object-cover mr-3" alt="">
                <div class="text-sm">{{ item.name }}</div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ item.type === 'video' ? '视频' : '图片' }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.show.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.click.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.ctr }}%</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.convert }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.convertCost }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.adCount }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="80" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', type: '' })

const handleQuery = () => {
  alert('查询素材报表')
}

const handleExport = () => {
  alert('导出素材报表')
}

const materials = ref([
  { id: 'M001', name: '618大促主视频', cover: 'https://via.placeholder.com/64x40', type: 'video', cost: 28560, show: 1580000, click: 48500, ctr: 3.07, convert: 986, convertCost: 29.0, adCount: 8 },
  { id: 'M002', name: '产品展示视频', cover: 'https://via.placeholder.com/64x40', type: 'video', cost: 22360, show: 1250000, click: 36800, ctr: 2.94, convert: 756, convertCost: 29.6, adCount: 6 },
  { id: 'M003', name: '直播间引流图', cover: 'https://via.placeholder.com/64x40', type: 'image', cost: 15680, show: 890000, click: 26500, ctr: 2.98, convert: 486, convertCost: 32.3, adCount: 5 },
  { id: 'M004', name: '爆款商品展示', cover: 'https://via.placeholder.com/64x40', type: 'video', cost: 18960, show: 1080000, click: 32500, ctr: 3.01, convert: 625, convertCost: 30.3, adCount: 4 },
  { id: 'M005', name: '品牌故事视频', cover: 'https://via.placeholder.com/64x40', type: 'video', cost: 12560, show: 680000, click: 19800, ctr: 2.91, convert: 386, convertCost: 32.5, adCount: 3 }
])
</script>
