<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const accounts = ref([
  { id: 'AWM001', nickname: '品牌官方号', avatar: '🎯', fans: 1256000, videos: 358, status: 'active', bindTime: '2025-06-15' },
  { id: 'AWM002', nickname: '产品测评号', avatar: '📱', fans: 568000, videos: 126, status: 'active', bindTime: '2025-08-20' },
  { id: 'AWM003', nickname: '活动推广号', avatar: '🎉', fans: 235000, videos: 89, status: 'expired', bindTime: '2025-09-10' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleBind = () => {
  alert('绑定新账号')
}

const handleDetail = (account: typeof accounts.value[0]) => {
  alert(`查看详情: ${account.nickname}`)
}

const handleReauth = (account: typeof accounts.value[0]) => {
  alert(`重新授权: ${account.nickname}`)
}

const handleUnbind = (account: typeof accounts.value[0]) => {
  if (confirm(`确定解绑「${account.nickname}」吗？`)) {
    alert('解绑成功')
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '抖音工具' }, { name: '账号管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">抖音账号管理</h1>
          <p class="mt-2 text-gray-600">管理已绑定的抖音号</p>
        </div>
<button @click="handleBind" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          绑定新账号
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">绑定账号</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">正常状态</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总粉丝数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">2,059,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总视频数</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">573</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账号信息</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">粉丝数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">绑定时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="account in accounts" :key="account.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="flex items-center">
                <span class="text-3xl mr-3">{{ account.avatar }}</span>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ account.nickname }}</div>
                  <div class="text-xs text-gray-500">{{ account.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ account.fans.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ account.videos }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ account.bindTime }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     account.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ account.status === 'active' ? '正常' : '已过期' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleDetail(account)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button v-if="account.status === 'expired'" @click="handleReauth(account)" class="text-green-600 hover:text-green-800 mr-3">重新授权</button>
              <button @click="handleUnbind(account)" class="text-red-600 hover:text-red-800">解绑</button>
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
