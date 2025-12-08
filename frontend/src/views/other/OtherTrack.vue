<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 32 })

const tracks = ref([
  { id: 'TRK001', name: 'APP激活追踪', type: 'app', platform: 'android', events: ['激活', '注册', '付费'], status: 'active', todayConversions: 1256, createdAt: '2025-09-01' },
  { id: 'TRK002', name: '网页表单追踪', type: 'web', platform: '-', events: ['表单提交', '电话咨询'], status: 'active', todayConversions: 342, createdAt: '2025-09-15' },
  { id: 'TRK003', name: '小程序购买追踪', type: 'miniapp', platform: 'wechat', events: ['下单', '付款'], status: 'active', todayConversions: 89, createdAt: '2025-10-01' },
  { id: 'TRK004', name: '测试追踪点', type: 'web', platform: '-', events: ['点击'], status: 'inactive', todayConversions: 0, createdAt: '2025-10-20' }
])

const getTypeLabel = (type: string) => {
  switch (type) {
    case 'app': return 'APP'
    case 'web': return '网页'
    case 'miniapp': return '小程序'
    default: return type
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateTrack = () => {
  alert('创建转化追踪')
}

const handleConfigTrack = (track: any) => {
  alert(`配置追踪: ${track.name}`)
}

const handleViewTrackData = (track: any) => {
  alert(`查看数据: ${track.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他功能' }, { name: '转化追踪' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">转化追踪管理</h1>
          <p class="mt-2 text-gray-600">配置和管理广告转化追踪点</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateTrack">
          创建追踪
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">追踪点</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">28</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日转化</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">1,687</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">本月转化</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">45,320</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">追踪点</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平台</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化事件</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="track in tracks" :key="track.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ track.name }}</div>
              <div class="text-xs text-gray-500">{{ track.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getTypeLabel(track.type) }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ track.platform }}</td>
            <td class="px-6 py-4">
              <div class="flex flex-wrap gap-1">
                <span v-for="event in track.events" :key="event" 
                      class="px-2 py-0.5 bg-gray-100 text-gray-700 rounded text-xs">{{ event }}</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ track.todayConversions.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                track.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ track.status === 'active' ? '启用' : '停用' }}
              </span>
            </td>
<td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleConfigTrack(track)">配置</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleViewTrackData(track)">数据</button>
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
