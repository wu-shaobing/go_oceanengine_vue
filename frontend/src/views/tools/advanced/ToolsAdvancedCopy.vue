<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })
const copyForm = reactive({
  sourceAd: '',
  targetAccount: '',
  copyCount: 1
})

const copyHistory = ref([
  { id: 'CP001', source: '促销广告A', target: '促销广告A-副本', copyItems: ['广告组', '创意', '定向'], status: 'success', time: '2025-11-28 10:00' },
  { id: 'CP002', source: '推广计划B', target: '推广计划B-副本', copyItems: ['广告组'], status: 'success', time: '2025-11-27 15:30' },
  { id: 'CP003', source: '品牌广告C', target: '品牌广告C-副本', copyItems: ['广告组', '创意'], status: 'failed', time: '2025-11-27 14:00' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleStartCopy = () => {
  alert('开始复制广告')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '广告复制' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">广告复制</h1>
          <p class="mt-2 text-gray-600">快速复制广告计划和创意</p>
        </div>
<button @click="handleStartCopy" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          开始复制
        </button>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h4 class="font-medium text-gray-900 mb-4">选择复制内容</h4>
        <div class="space-y-3">
          <label class="flex items-center p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100">
            <input type="checkbox" checked class="rounded text-blue-600 mr-3">
            <span class="text-sm text-gray-700">广告计划</span>
          </label>
          <label class="flex items-center p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100">
            <input type="checkbox" checked class="rounded text-blue-600 mr-3">
            <span class="text-sm text-gray-700">广告组</span>
          </label>
          <label class="flex items-center p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100">
            <input type="checkbox" checked class="rounded text-blue-600 mr-3">
            <span class="text-sm text-gray-700">创意素材</span>
          </label>
          <label class="flex items-center p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100">
            <input type="checkbox" class="rounded text-blue-600 mr-3">
            <span class="text-sm text-gray-700">定向设置</span>
          </label>
        </div>
      </div>

      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <h4 class="font-medium text-gray-900 mb-4">复制设置</h4>
        <div class="space-y-4">
          <div>
            <label class="block text-sm text-gray-600 mb-2">选择源广告</label>
            <select v-model="copyForm.sourceAd" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="">请选择要复制的广告</option>
              <option value="1">促销广告A</option>
              <option value="2">推广计划B</option>
              <option value="3">品牌广告C</option>
            </select>
          </div>
          <div>
            <label class="block text-sm text-gray-600 mb-2">目标账户</label>
            <select v-model="copyForm.targetAccount" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="">当前账户</option>
              <option value="1">子账户A</option>
              <option value="2">子账户B</option>
            </select>
          </div>
          <div>
            <label class="block text-sm text-gray-600 mb-2">复制份数</label>
            <input v-model="copyForm.copyCount" type="number" min="1" max="10" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
          </div>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">复制历史</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">源广告</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标广告</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">复制内容</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">复制时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="copy in copyHistory" :key="copy.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm text-gray-900">{{ copy.source }}</td>
            <td class="px-6 py-4 text-sm text-blue-600">{{ copy.target }}</td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="item in copy.copyItems" :key="item"
                      class="px-2 py-0.5 bg-gray-100 text-gray-700 rounded text-xs">{{ item }}</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ copy.time }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     copy.status === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ copy.status === 'success' ? '成功' : '失败' }}
              </span>
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
