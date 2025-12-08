<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '商品管理' }, { name: '商品列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">商品管理</h1>
        <p class="text-gray-600 mt-1">管理可投放的店铺商品</p>
      </div>
      <button @click="handleSync" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        同步商品
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="商品名称/ID" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.shop" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部店铺</option>
          <option value="7001">品牌官方旗舰店</option>
          <option value="7002">美妆专营店</option>
        </select>
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="online">在售</option>
          <option value="offline">下架</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 商品列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">商品信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">所属店铺</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">价格</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">库存</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">销量</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">投放数据</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="product in products" :key="product.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="product.image" class="w-16 h-16 rounded mr-3 object-cover" alt="">
                <div>
                  <div class="font-medium line-clamp-2 w-48">{{ product.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ product.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ product.shopName }}</td>
            <td class="px-4 py-3 text-sm text-red-600 font-medium">¥{{ product.price }}</td>
            <td class="px-4 py-3 text-sm">{{ product.stock }}</td>
            <td class="px-4 py-3">
              <span :class="product.status === 'online' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" 
                class="px-2 py-1 text-xs rounded-full">
                {{ product.status === 'online' ? '在售' : '下架' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">{{ product.sales }}</td>
            <td class="px-4 py-3">
              <div class="text-sm">
                <div>消耗: ¥{{ product.cost.toLocaleString() }}</div>
                <div class="text-gray-500">ROI: {{ product.roi }}</div>
              </div>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button @click="handlePromote(product)" class="text-blue-600 hover:text-blue-800 text-sm">投放</button>
                <button @click="handleAnalyze(product)" class="text-blue-600 hover:text-blue-800 text-sm">分析</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="500" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({
  keyword: '',
  shop: '',
  status: ''
})

const products = ref([
  { id: 'P001', name: '【爆款】美白保湿面霜 补水滋润护肤品', image: 'https://via.placeholder.com/64', shopName: '美妆专营店', price: 199, stock: 5680, status: 'online', sales: 12580, cost: 45680, roi: '4.2' },
  { id: 'P002', name: '夏季新款连衣裙 修身显瘦气质', image: 'https://via.placeholder.com/64', shopName: '服装精品店', price: 299, stock: 3200, status: 'online', sales: 8900, cost: 32450, roi: '3.8' },
  { id: 'P003', name: '网红零食大礼包 办公室休闲小吃', image: 'https://via.placeholder.com/64', shopName: '食品特卖店', price: 49.9, stock: 15000, status: 'online', sales: 25600, cost: 18760, roi: '5.6' },
  { id: 'P004', name: '无线蓝牙耳机 降噪运动', image: 'https://via.placeholder.com/64', shopName: '数码配件店', price: 159, stock: 2800, status: 'online', sales: 6780, cost: 28900, roi: '3.2' },
  { id: 'P005', name: '精华液套装 修复抗衰老', image: 'https://via.placeholder.com/64', shopName: '美妆专营店', price: 399, stock: 0, status: 'offline', sales: 4560, cost: 12340, roi: '4.5' }
])

const handleSync = () => {
  alert('正在同步商品...')
}

const handleSearch = () => {
  console.log('搜索:', filters.value)
  alert('查询完成')
}

const handlePromote = (product: typeof products.value[0]) => {
  if (product.status === 'offline') {
    alert('下架商品无法投放')
    return
  }
  alert(`创建商品推广: ${product.name}`)
}

const handleAnalyze = (product: typeof products.value[0]) => {
  alert(`查看商品分析: ${product.name}\nROI: ${product.roi}\n销量: ${product.sales}`)
}
</script>
