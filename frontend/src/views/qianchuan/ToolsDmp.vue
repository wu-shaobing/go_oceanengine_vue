<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '工具' }, { name: 'DMP人群' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">DMP人群管理</h1>
        <p class="text-gray-600 mt-1">管理和使用DMP人群包</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">创建人群</button>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">人群名称</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">覆盖人数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">更新时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in audiences" :key="item.id">
            <td class="px-4 py-3 font-medium">{{ item.name }}</td>
            <td class="px-4 py-3 text-sm">{{ item.type }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.coverage.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <span :class="item.status === 'ready' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'" class="px-2 py-1 text-xs rounded">
                {{ item.status === 'ready' ? '已就绪' : '计算中' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ item.updateTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">推送</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="50" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const audiences = ref([
  { id: 1, name: '高价值用户', type: '自定义人群', coverage: 1250000, status: 'ready', updateTime: '2024-06-18' },
  { id: 2, name: '近30天成交用户', type: '行为人群', coverage: 580000, status: 'ready', updateTime: '2024-06-17' },
  { id: 3, name: '直播间互动用户', type: '行为人群', coverage: 860000, status: 'ready', updateTime: '2024-06-16' },
  { id: 4, name: '竞品用户', type: '自定义人群', coverage: 450000, status: 'pending', updateTime: '2024-06-15' }
])
</script>
