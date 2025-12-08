<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '线索管理' }, { name: '线索列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">线索管理</h1>
        <p class="text-gray-600 mt-1">管理获取的潜在客户线索</p>
      </div>
      <button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
        导出线索
      </button>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日线索</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stats.todayLeads }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">待跟进</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.pending }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已成交</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.converted }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化率</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.conversionRate }}%</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="姓名/电话" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="new">新线索</option>
          <option value="following">跟进中</option>
          <option value="converted">已成交</option>
          <option value="invalid">无效</option>
        </select>
        <select v-model="filters.source" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部来源</option>
          <option value="form">表单提交</option>
          <option value="phone">电话咨询</option>
          <option value="chat">在线咨询</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 线索列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">客户信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">来源</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">所属项目</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">获取时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">跟进人</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="clue in clues" :key="clue.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="font-medium">{{ clue.name }}</div>
              <div class="text-sm text-gray-500">{{ clue.phone }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ clue.source }}</td>
            <td class="px-4 py-3 text-sm">{{ clue.project }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(clue.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(clue.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">{{ clue.createTime }}</td>
            <td class="px-4 py-3 text-sm">{{ clue.follower || '-' }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button @click="handleFollow(clue)" class="text-blue-600 hover:text-blue-800 text-sm">跟进</button>
                <button @click="handleDetail(clue)" class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button @click="handleInvalid(clue)" class="text-red-600 hover:text-red-800 text-sm">无效</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="500" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const stats = ref({
  todayLeads: 68,
  pending: 125,
  converted: 86,
  conversionRate: 18.5
})

const filters = ref({
  keyword: '',
  status: '',
  source: ''
})

const clues = ref([
  { id: 1, name: '张女士', phone: '138****5678', source: '表单提交', project: '餐饮推广项目', status: 'new', createTime: '2024-03-15 18:30', follower: '' },
  { id: 2, name: '李先生', phone: '139****1234', source: '电话咨询', project: '美容美发推广', status: 'following', createTime: '2024-03-15 16:20', follower: '销售A' },
  { id: 3, name: '王女士', phone: '136****9012', source: '在线咨询', project: '教育培训招生', status: 'converted', createTime: '2024-03-15 14:10', follower: '销售B' },
  { id: 4, name: '赵先生', phone: '135****3456', source: '表单提交', project: '汽车服务推广', status: 'following', createTime: '2024-03-15 11:45', follower: '销售A' },
  { id: 5, name: '陈女士', phone: '137****7890', source: '表单提交', project: '健身房获客', status: 'invalid', createTime: '2024-03-15 09:30', follower: '销售C' }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    new: 'bg-blue-100 text-blue-800',
    following: 'bg-orange-100 text-orange-800',
    converted: 'bg-green-100 text-green-800',
    invalid: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    new: '新线索',
    following: '跟进中',
    converted: '已成交',
    invalid: '无效'
  }
  return texts[status] || status
}

const handleSearch = () => {
  console.log('搜索线索:', filters.value)
  alert('查询完成')
}

const handleExport = () => {
  alert('线索导出中...')
}

const handleFollow = (clue: typeof clues.value[0]) => {
  alert(`开始跟进线索: ${clue.name}`)
}

const handleDetail = (clue: typeof clues.value[0]) => {
  alert(`查看线索详情: ${clue.name}\n电话: ${clue.phone}\n来源: ${clue.source}`)
}

const handleInvalid = (clue: typeof clues.value[0]) => {
  if (confirm(`确定将线索「${clue.name}」标记为无效吗？`)) {
    clue.status = 'invalid'
    alert('已标记为无效')
  }
}
</script>
