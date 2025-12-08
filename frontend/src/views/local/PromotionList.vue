<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '推广管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">推广管理</h1>
        <p class="text-gray-600 mt-1">管理本地推广计划</p>
      </div>
      <button @click="handleCreate" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        新建推广
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="推广名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="running">投放中</option>
          <option value="paused">已暂停</option>
          <option value="ended">已结束</option>
        </select>
        <select v-model="filters.project" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部项目</option>
          <option value="1">北京朝阳店</option>
          <option value="2">上海静安店</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 推广列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">推广信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">所属项目</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">今日消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">今日线索</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索成本</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="promo in promotions" :key="promo.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ promo.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ promo.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ promo.project }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(promo.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(promo.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ promo.todayCost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ promo.todayClues }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ promo.clueCost }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button @click="handleDetail(promo)" class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button v-if="promo.status === 'running'" @click="handlePause(promo)" class="text-orange-600 hover:text-orange-800 text-sm">暂停</button>
                <button v-else @click="handleStart(promo)" class="text-green-600 hover:text-green-800 text-sm">启动</button>
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
const filters = ref({ keyword: '', status: '', project: '' })

const promotions = ref([
  { id: 'PM001', name: '618限时优惠推广', project: '北京朝阳店', status: 'running', todayCost: 1580, todayClues: 28, clueCost: 56.4 },
  { id: 'PM002', name: '新品尝鲜活动', project: '北京朝阳店', status: 'running', todayCost: 650, todayClues: 12, clueCost: 54.2 },
  { id: 'PM003', name: '会员专享折扣', project: '北京朝阳店', status: 'paused', todayCost: 0, todayClues: 0, clueCost: 0 },
  { id: 'PM004', name: '周末特惠推广', project: '上海静安店', status: 'running', todayCost: 2100, todayClues: 35, clueCost: 60.0 },
  { id: 'PM005', name: '新店开业活动', project: '上海静安店', status: 'ended', todayCost: 0, todayClues: 0, clueCost: 0 }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    running: 'bg-green-100 text-green-800',
    paused: 'bg-orange-100 text-orange-800',
    ended: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    running: '投放中',
    paused: '已暂停',
    ended: '已结束'
  }
  return texts[status] || status
}

const handleCreate = () => {
  router.push('/local/promotion/create')
}

const handleSearch = () => {
  console.log('搜索推广:', filters.value)
  alert('查询完成')
}

const handleDetail = (promo: typeof promotions.value[0]) => {
  router.push(`/local/promotion/${promo.id}`)
}

const handlePause = (promo: typeof promotions.value[0]) => {
  if (confirm(`确定暂停推广「${promo.name}」吗？`)) {
    promo.status = 'paused'
    alert('推广已暂停')
  }
}

const handleStart = (promo: typeof promotions.value[0]) => {
  promo.status = 'running'
  alert(`推广「${promo.name}」已启动`)
}
</script>
