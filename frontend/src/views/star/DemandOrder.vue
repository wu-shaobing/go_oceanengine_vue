<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '需求单管理' }]" />
    
    <div class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">需求单管理</h1>
        <p class="text-gray-600 mt-1">管理达人合作需求单</p>
      </div>
      <button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">创建需求单</button>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="filters.keyword" placeholder="搜索需求单" class="border border-gray-300 rounded px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="draft">草稿</option>
          <option value="pending">待审核</option>
          <option value="approved">已通过</option>
          <option value="rejected">已拒绝</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">需求信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">预算</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">创建时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ order.title }}</div>
              <div class="text-sm text-gray-500">{{ order.desc }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ order.type }}</td>
            <td class="px-4 py-3 text-right font-medium">¥{{ order.budget.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs rounded">{{ getStatusText(order.status) }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ order.createTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button @click="handleView(order)" class="text-blue-600 hover:text-blue-800 text-sm">查看</button>
                <button v-if="order.status === 'draft'" @click="handleEdit(order)" class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button v-if="order.status === 'draft'" @click="handleDelete(order)" class="text-red-600 hover:text-red-800 text-sm">删除</button>
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
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const router = useRouter()
const filters = ref({ keyword: '', status: '' })

const orders = ref([
  { id: 1, title: '618美妆种草视频需求', desc: '寻找美妆领域达人进行产品种草', type: '视频推广', budget: 100000, status: 'approved', createTime: '2024-06-01 10:00' },
  { id: 2, title: '新品护肤测评需求', desc: '护肤新品上市测评推广', type: '测评推广', budget: 50000, status: 'pending', createTime: '2024-06-10 14:30' },
  { id: 3, title: '品牌直播带货需求', desc: '618直播带货专场', type: '直播推广', budget: 200000, status: 'draft', createTime: '2024-06-15 09:00' },
  { id: 4, title: '日常种草内容需求', desc: '日常产品种草内容产出', type: '图文推广', budget: 30000, status: 'rejected', createTime: '2024-05-20 16:00' }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = { draft: 'bg-gray-100 text-gray-800', pending: 'bg-orange-100 text-orange-800', approved: 'bg-green-100 text-green-800', rejected: 'bg-red-100 text-red-800' }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = { draft: '草稿', pending: '待审核', approved: '已通过', rejected: '已拒绝' }
  return texts[status] || status
}

const handleCreate = () => {
  router.push('/star/demand/create')
}

const handleSearch = () => {
  console.log('搜索:', filters.value)
  alert('查询完成')
}

const handleView = (order: typeof orders.value[0]) => {
  alert(`查看需求单: ${order.title}`)
}

const handleEdit = (order: typeof orders.value[0]) => {
  alert(`编辑需求单: ${order.title}`)
}

const handleDelete = (order: typeof orders.value[0]) => {
  if (confirm(`确定删除需求单「${order.title}」吗？`)) {
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx > -1) orders.value.splice(idx, 1)
    alert('删除成功')
  }
}
</script>
