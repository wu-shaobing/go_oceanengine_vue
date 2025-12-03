<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })
const activeTab = ref('pending')

const reviews = ref([
  { id: 'CR001', name: '双十一促销视频', type: 'video', submitTime: '2025-11-28 09:30', status: 'pending', reason: '' },
  { id: 'CR002', name: '新品发布海报', type: 'image', submitTime: '2025-11-28 08:45', status: 'approved', reason: '' },
  { id: 'CR003', name: '限时优惠Banner', type: 'image', submitTime: '2025-11-27 16:20', status: 'rejected', reason: '图片文字过多' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '创意工具' }, { name: '创意审核' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创意审核</h1>
      <p class="mt-2 text-gray-600">查看广告创意审核状态</p>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">待审核</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">审核中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已通过</p>
        <p class="text-2xl font-bold text-green-600 mt-1">156</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已拒绝</p>
        <p class="text-2xl font-bold text-red-600 mt-1">23</p>
      </div>
    </div>

    <div class="flex gap-2">
      <button :class="['px-4 py-2 rounded-lg', activeTab === 'pending' ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
              @click="activeTab = 'pending'">待审核</button>
      <button :class="['px-4 py-2 rounded-lg', activeTab === 'approved' ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
              @click="activeTab = 'approved'">已通过</button>
      <button :class="['px-4 py-2 rounded-lg', activeTab === 'rejected' ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
              @click="activeTab = 'rejected'">已拒绝</button>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创意名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提交时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">拒绝原因</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="review in reviews" :key="review.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-12 h-12 bg-gray-200 rounded flex items-center justify-center mr-3">
                  <span class="text-xl">{{ review.type === 'video' ? '🎬' : '🖼️' }}</span>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ review.name }}</div>
                  <div class="text-xs text-gray-500">{{ review.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     review.type === 'video' ? 'bg-purple-100 text-purple-700' : 'bg-blue-100 text-blue-700']">
                {{ review.type === 'video' ? '视频' : '图片' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ review.submitTime }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     review.status === 'pending' ? 'bg-yellow-100 text-yellow-800' :
                     review.status === 'approved' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ review.status === 'pending' ? '待审核' : review.status === 'approved' ? '已通过' : '已拒绝' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-red-600">{{ review.reason || '-' }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">查看</button>
              <button v-if="review.status === 'rejected'" class="text-green-600 hover:text-green-800">重新提交</button>
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
