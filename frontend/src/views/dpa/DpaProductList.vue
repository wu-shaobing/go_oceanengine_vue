<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 156 })

const products = ref([
  { id: 'P001', name: 'Apple iPhone 15 Pro Max 256GB', category: '手机数码', price: 9999, stock: 1250, status: 'active', image: '' },
  { id: 'P002', name: '华为 Mate 60 Pro 512GB', category: '手机数码', price: 6999, stock: 890, status: 'active', image: '' },
  { id: 'P003', name: '小米14 Ultra 1TB', category: '手机数码', price: 5999, stock: 0, status: 'out_of_stock', image: '' },
  { id: 'P004', name: 'MacBook Pro 14寸 M3 Pro', category: '电脑办公', price: 16999, stock: 320, status: 'active', image: '' },
  { id: 'P005', name: 'iPad Pro 12.9寸 M2', category: '平板电脑', price: 8999, stock: 45, status: 'low_stock', image: '' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: '商品列表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">DPA商品库</h1>
          <p class="mt-2 text-gray-600">管理动态商品广告的商品信息</p>
        </div>
        <div class="flex gap-3">
          <button class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
            批量导入
          </button>
          <router-link to="/dpa/product/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            添加商品
          </router-link>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总商品数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">在售商品</p>
        <p class="text-2xl font-bold text-green-600 mt-1">142</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">缺货商品</p>
        <p class="text-2xl font-bold text-red-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">低库存</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">6</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <input type="text" placeholder="搜索商品名称或ID" class="flex-1 min-w-[200px] px-4 py-2 border border-gray-300 rounded-lg">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部分类</option>
          <option value="phone">手机数码</option>
          <option value="computer">电脑办公</option>
          <option value="tablet">平板电脑</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部状态</option>
          <option value="active">在售</option>
          <option value="out_of_stock">缺货</option>
          <option value="low_stock">低库存</option>
        </select>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left">
                <input type="checkbox" class="w-4 h-4 rounded border-gray-300">
              </th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">商品</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分类</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">价格</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">库存</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="product in products" :key="product.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <input type="checkbox" class="w-4 h-4 rounded border-gray-300">
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-12 h-12 bg-gray-100 rounded flex items-center justify-center text-gray-400 text-xs">
                    图片
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ product.name }}</div>
                    <div class="text-xs text-gray-500">ID: {{ product.id }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-700">{{ product.category }}</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ product.price.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm" :class="product.stock === 0 ? 'text-red-600' : product.stock < 50 ? 'text-yellow-600' : 'text-gray-900'">
                {{ product.stock }}
              </td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  product.status === 'active' ? 'bg-green-100 text-green-800' :
                  product.status === 'out_of_stock' ? 'bg-red-100 text-red-800' : 'bg-yellow-100 text-yellow-800']">
                  {{ product.status === 'active' ? '在售' : product.status === 'out_of_stock' ? '缺货' : '低库存' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
                <button class="text-gray-600 hover:text-gray-800">详情</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
