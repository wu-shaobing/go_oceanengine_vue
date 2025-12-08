<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const anchors = ref([
  { id: 'NA001', name: '商品购买锚点', type: 'product', adsCount: 45, clicks: 15600, conversions: 890, status: 'active' },
  { id: 'NA002', name: 'APP下载锚点', type: 'app', adsCount: 32, clicks: 12300, conversions: 1256, status: 'active' },
  { id: 'NA003', name: '表单收集锚点', type: 'form', adsCount: 28, clicks: 8900, conversions: 456, status: 'active' },
  { id: 'NA004', name: '电话咨询锚点', type: 'phone', adsCount: 15, clicks: 5600, conversions: 234, status: 'paused' }
])

const getTypeConfig = (type: string) => {
  switch (type) {
    case 'product': return { label: '商品', class: 'bg-blue-100 text-blue-700' }
    case 'app': return { label: '应用', class: 'bg-green-100 text-green-700' }
    case 'form': return { label: '表单', class: 'bg-purple-100 text-purple-700' }
    default: return { label: '电话', class: 'bg-orange-100 text-orange-700' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建原生锚点')
}

const handleEdit = (anchor: typeof anchors.value[0]) => {
  alert(`编辑锚点: ${anchor.name}`)
}

const handleData = (anchor: typeof anchors.value[0]) => {
  alert(`查看数据: ${anchor.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '原生锚点' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">原生锚点管理</h1>
          <p class="mt-2 text-gray-600">创建和管理广告原生锚点组件</p>
        </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建锚点
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总锚点</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联广告</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">120</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">42.4K</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总转化</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2,836</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">锚点名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联广告</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="anchor in anchors" :key="anchor.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ anchor.name }}</div>
              <div class="text-xs text-gray-500">{{ anchor.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs', getTypeConfig(anchor.type).class]">
                {{ getTypeConfig(anchor.type).label }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ anchor.adsCount }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ anchor.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ anchor.conversions }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                anchor.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ anchor.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleEdit(anchor)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleData(anchor)" class="text-gray-600 hover:text-gray-800">数据</button>
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
