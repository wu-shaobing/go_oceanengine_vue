<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 25 })

const audits = ref([
  { id: 'AD001', name: '新电商广告主', type: '企业', submitTime: '2025-11-28 09:30', docs: ['营业执照', '资质证明'], status: 'pending' },
  { id: 'AD002', name: '游戏公司B', type: '企业', submitTime: '2025-11-27 15:45', docs: ['营业执照', '游戏版号'], status: 'approved' },
  { id: 'AD003', name: '个人广告主C', type: '个人', submitTime: '2025-11-27 10:20', docs: ['身份证'], status: 'rejected' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告主管理' }, { name: '资质审核' }]" />
      <h1 class="text-3xl font-bold text-gray-900">资质审核</h1>
      <p class="mt-2 text-gray-600">审核广告主资质材料</p>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">审核总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">待审核</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已通过</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已拒绝</p>
        <p class="text-2xl font-bold text-red-600 mt-1">2</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告主</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提交时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提交材料</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="audit in audits" :key="audit.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ audit.name }}</div>
              <div class="text-xs text-gray-500">{{ audit.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     audit.type === '企业' ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700']">
                {{ audit.type }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ audit.submitTime }}</td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="doc in audit.docs" :key="doc"
                      class="px-2 py-0.5 bg-gray-100 text-gray-700 rounded text-xs">{{ doc }}</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     audit.status === 'pending' ? 'bg-yellow-100 text-yellow-800' :
                     audit.status === 'approved' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ audit.status === 'pending' ? '待审核' : audit.status === 'approved' ? '已通过' : '已拒绝' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">查看</button>
              <button v-if="audit.status === 'pending'" class="text-green-600 hover:text-green-800 mr-2">通过</button>
              <button v-if="audit.status === 'pending'" class="text-red-600 hover:text-red-800">拒绝</button>
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
