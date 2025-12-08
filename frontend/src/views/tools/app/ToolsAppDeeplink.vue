<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })
const testDeeplinkUrl = ref('')

const deeplinks = ref([
  { id: 'DL001', name: '首页跳转', app: '购物 App', scheme: 'myapp://home', status: 'active', createTime: '2025-10-15' },
  { id: 'DL002', name: '商品详情页', app: '购物 App', scheme: 'myapp://product/{id}', status: 'active', createTime: '2025-10-16' },
  { id: 'DL003', name: '活动页面', app: '购物 App', scheme: 'myapp://activity/{id}', status: 'inactive', createTime: '2025-10-20' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddDeeplink = () => {
  alert('添加新的Deeplink配置')
}

const handleEdit = (link: typeof deeplinks.value[0]) => {
  alert(`编辑Deeplink: ${link.name}`)
}

const handleTest = (link: typeof deeplinks.value[0]) => {
  alert(`测试Deeplink: ${link.scheme}`)
}

const handleDelete = (link: typeof deeplinks.value[0]) => {
  if (confirm(`确定删除Deeplink「${link.name}」吗？`)) {
    const idx = deeplinks.value.findIndex(l => l.id === link.id)
    if (idx > -1) deeplinks.value.splice(idx, 1)
    alert('删除成功')
  }
}

const handleValidate = () => {
  alert('正在验证链接...')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: 'Deeplink配置' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">Deeplink配置</h1>
          <p class="mt-2 text-gray-600">管理应用深度链接</p>
        </div>
        <button @click="handleAddDeeplink" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加Deeplink
        </button>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">链接名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">Scheme</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="link in deeplinks" :key="link.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ link.name }}</div>
                <div class="text-xs text-gray-500">{{ link.id }}</div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ link.app }}</td>
              <td class="px-6 py-4">
                <code class="px-2 py-1 bg-gray-100 rounded text-xs text-gray-800">{{ link.scheme }}</code>
              </td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                       link.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                  {{ link.status === 'active' ? '已启用' : '未启用' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm">
                <button @click="handleEdit(link)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
                <button @click="handleTest(link)" class="text-gray-600 hover:text-gray-800 mr-3">测试</button>
                <button @click="handleDelete(link)" class="text-red-600 hover:text-red-800">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="px-6 py-4 border-t border-gray-200">
          <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">Deeplink说明</h4>
          <div class="space-y-3 text-sm text-gray-600">
            <p>Deeplink用于实现从广告直接跳转到应用内指定页面。</p>
            <div class="p-3 bg-blue-50 rounded-lg">
              <p class="font-medium text-blue-800">支持的协议：</p>
              <ul class="mt-2 space-y-1 text-blue-700">
                <li>• URL Scheme</li>
                <li>• Universal Links (iOS)</li>
                <li>• App Links (Android)</li>
              </ul>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">测试工具</h4>
          <div class="space-y-3">
            <input v-model="testDeeplinkUrl" type="text" placeholder="输入Deeplink URL" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            <button @click="handleValidate" class="w-full px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 text-sm">
              验证链接
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
