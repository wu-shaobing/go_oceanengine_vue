<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '自定义报表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">自定义报表</h1>
        <p class="text-gray-600 mt-1">自定义数据维度和指标生成报表</p>
      </div>
<button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">新建报表</button>
    </div>

    <!-- 报表列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">报表名称</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">数据维度</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">创建时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">更新时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="report in reports" :key="report.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ report.name }}</div>
            </td>
            <td class="px-4 py-3 text-sm">
              <div class="flex flex-wrap gap-1">
                <span v-for="dim in report.dimensions" :key="dim" class="px-2 py-0.5 text-xs bg-gray-100 rounded">{{ dim }}</span>
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ report.createTime }}</td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ report.updateTime }}</td>
            <td class="px-4 py-3">
              <span :class="report.status === 'ready' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'" class="px-2 py-1 text-xs rounded">
                {{ report.status === 'ready' ? '已完成' : '生成中' }}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
<button @click="handleView(report)" class="text-blue-600 hover:text-blue-800 text-sm">查看</button>
                <button @click="handleDownload(report)" class="text-blue-600 hover:text-blue-800 text-sm">下载</button>
                <button @click="handleDelete(report)" class="text-red-600 hover:text-red-800 text-sm">删除</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="20" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const reports = ref([
  { id: 1, name: '618大促效果分析', dimensions: ['日期', '广告', '创意'], createTime: '2024-06-19 10:00', updateTime: '2024-06-19 10:05', status: 'ready' },
  { id: 2, name: '直播间ROI周报', dimensions: ['日期', '直播间'], createTime: '2024-06-18 09:00', updateTime: '2024-06-18 09:10', status: 'ready' },
  { id: 3, name: '素材效果对比', dimensions: ['素材', '广告'], createTime: '2024-06-17 14:00', updateTime: '2024-06-17 14:15', status: 'ready' },
  { id: 4, name: '关键词转化分析', dimensions: ['关键词', '日期'], createTime: '2024-06-16 11:00', updateTime: '2024-06-16 11:08', status: 'ready' },
  { id: 5, name: '商品销售报表', dimensions: ['商品', '日期'], createTime: '2024-06-15 16:00', updateTime: '2024-06-15 16:00', status: 'pending' }
])

const handleCreate = () => {
  alert('新建报表')
}

const handleView = (report: typeof reports.value[0]) => {
  alert(`查看报表: ${report.name}`)
}

const handleDownload = (report: typeof reports.value[0]) => {
  alert(`下载报表: ${report.name}`)
}

const handleDelete = (report: typeof reports.value[0]) => {
  alert(`删除报表: ${report.name}`)
}
</script>
