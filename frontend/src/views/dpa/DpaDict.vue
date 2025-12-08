<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 68 })
const searchKeyword = ref('')

const dictItems = ref([
  { id: 'D001', field: 'product_name', mapping: '商品名称', type: 'string', required: true, example: 'iPhone 15 Pro' },
  { id: 'D002', field: 'price', mapping: '商品价格', type: 'number', required: true, example: '7999' },
  { id: 'D003', field: 'original_price', mapping: '原价', type: 'number', required: false, example: '8999' },
  { id: 'D004', field: 'category', mapping: '商品分类', type: 'string', required: true, example: '数码电子' },
  { id: 'D005', field: 'image_url', mapping: '商品图片', type: 'url', required: true, example: 'https://...' },
  { id: 'D006', field: 'landing_url', mapping: '落地页', type: 'url', required: true, example: 'https://...' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddField = () => {
  alert('添加字段')
}

const handleSearch = () => {
  alert('搜索字段')
}

const handleEditField = (item: typeof dictItems.value[0]) => {
  alert(`编辑字段: ${item.field}`)
}

const handleDeleteField = (item: typeof dictItems.value[0]) => {
  if (confirm(`确定删除字段 ${item.field}?`)) {
    alert('删除成功')
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品广告' }, { name: '商品词典' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">商品词典</h1>
          <p class="mt-2 text-gray-600">定义商品数据字段映射</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleAddField">
          添加字段
        </button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="searchKeyword" type="text" placeholder="搜索字段名..."
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">字段总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">必填字段</p>
        <p class="text-2xl font-bold text-red-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可选字段</p>
        <p class="text-2xl font-bold text-gray-600 mt-1">56</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">字段名</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">中文名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">数据类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">是否必填</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">示例值</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in dictItems" :key="item.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <code class="px-2 py-1 bg-blue-50 text-blue-700 rounded text-sm">{{ item.field }}</code>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ item.mapping }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ item.type }}</span>
            </td>
            <td class="px-6 py-4">
              <span v-if="item.required" class="text-red-500 font-medium">是</span>
              <span v-else class="text-gray-400">否</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500 truncate max-w-xs">{{ item.example }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleEditField(item)">编辑</button>
              <button class="text-red-600 hover:text-red-800" @click="handleDeleteField(item)">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
