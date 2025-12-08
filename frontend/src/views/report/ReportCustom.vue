<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 15 })

const customReports = ref([
  { id: 'CR001', name: '每日消耗汇总', metrics: ['消耗', '展示', '点击'], schedule: '每天 09:00', lastRun: '2025-11-28 09:00', status: 'active' },
  { id: 'CR002', name: '转化效果分析', metrics: ['转化', 'CVR', 'CPA'], schedule: '每周一 10:00', lastRun: '2025-11-25 10:00', status: 'active' },
  { id: 'CR003', name: '素材表现报表', metrics: ['展示', 'CTR', '消耗'], schedule: '手动', lastRun: '2025-11-27 15:30', status: 'active' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreate = () => {
  alert('创建新的自定义报表')
}

const handleRun = (report: typeof customReports.value[0]) => {
  alert(`运行报表: ${report.name}`)
}

const handleEdit = (report: typeof customReports.value[0]) => {
  alert(`编辑报表: ${report.name}`)
}

const handleDelete = (report: typeof customReports.value[0]) => {
  if (confirm(`确定删除报表「${report.name}」吗？`)) {
    const idx = customReports.value.findIndex(r => r.id === report.id)
    if (idx > -1) customReports.value.splice(idx, 1)
    alert('删除成功')
  }
}

const handleUseTemplate = (name: string) => {
  alert(`使用${name}模板创建报表`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '数据报表' }, { name: '自定义报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">自定义报表</h1>
          <p class="mt-2 text-gray-600">创建和管理自定义数据报表</p>
        </div>
        <button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建报表
        </button>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">报表名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">指标</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">执行计划</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">上次运行</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="report in customReports" :key="report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ report.name }}</div>
                <div class="text-xs text-gray-500">{{ report.id }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="flex flex-wrap gap-1">
                  <span v-for="metric in report.metrics" :key="metric"
                        class="px-2 py-0.5 bg-blue-100 text-blue-700 rounded text-xs">{{ metric }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ report.schedule }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ report.lastRun }}</td>
              <td class="px-6 py-4 text-sm">
                <button @click="handleRun(report)" class="text-blue-600 hover:text-blue-800 mr-3">运行</button>
                <button @click="handleEdit(report)" class="text-gray-600 hover:text-gray-800 mr-3">编辑</button>
                <button @click="handleDelete(report)" class="text-red-600 hover:text-red-800">删除</button>
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
          <h4 class="font-medium text-gray-900 mb-3">报表模板</h4>
          <div class="space-y-2">
            <button @click="handleUseTemplate('消耗分析')" class="w-full text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100">
              <p class="text-sm font-medium text-gray-900">消耗分析模板</p>
              <p class="text-xs text-gray-500">按账户/广告组分析消耗</p>
            </button>
            <button @click="handleUseTemplate('转化追踪')" class="w-full text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100">
              <p class="text-sm font-medium text-gray-900">转化追踪模板</p>
              <p class="text-xs text-gray-500">追踪转化漏斗数据</p>
            </button>
            <button @click="handleUseTemplate('ROI分析')" class="w-full text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100">
              <p class="text-sm font-medium text-gray-900">ROI分析模板</p>
              <p class="text-xs text-gray-500">分析投资回报率</p>
            </button>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">常用指标</h4>
          <div class="flex flex-wrap gap-2">
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">消耗</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">展示</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">点击</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">CTR</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">CVR</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">CPA</span>
            <span class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs">ROI</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
