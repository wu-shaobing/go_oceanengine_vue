<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '全域推广' }, { name: '推广列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">全域推广</h1>
        <p class="text-gray-600 mt-1">智能托管，全域流量覆盖</p>
      </div>
      <router-link to="/qianchuan/uni/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建全域推广
      </router-link>
    </div>

    <!-- 数据概览 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日消耗</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">¥{{ stats.todayCost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日GMV</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">¥{{ stats.todayGmv.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">整体ROI</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.roi }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化数</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.conversions.toLocaleString() }}</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="推广名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="running">投放中</option>
          <option value="pause">已暂停</option>
          <option value="end">已结束</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="live">直播</option>
          <option value="video">短视频</option>
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
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">日预算</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">ROI目标</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">GMV</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">实际ROI</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in promotions" :key="item.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ item.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ item.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ item.type === 'live' ? '直播' : '短视频' }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(item.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(item.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">¥{{ item.budget.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ item.roiTarget }}</td>
            <td class="px-4 py-3 text-sm">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">¥{{ item.gmv.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <span :class="item.actualRoi >= item.roiTarget ? 'text-green-600' : 'text-red-600'" class="font-medium">
                {{ item.actualRoi }}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <router-link :to="`/qianchuan/uni/${item.id}`" class="text-blue-600 hover:text-blue-800 text-sm">详情</router-link>
                <button @click="handleEdit(item)" class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button v-if="item.status === 'running'" @click="handlePause(item)" class="text-orange-600 hover:text-orange-800 text-sm">暂停</button>
                <button v-else @click="handleStart(item)" class="text-green-600 hover:text-green-800 text-sm">启动</button>
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

const stats = ref({
  todayCost: 68500,
  todayGmv: 356800,
  roi: '5.21',
  conversions: 1256
})

const filters = ref({
  keyword: '',
  status: '',
  type: ''
})

const promotions = ref([
  { id: 'U001', name: '618大促直播全域推广', type: 'live', status: 'running', budget: 30000, roiTarget: 4.5, cost: 28600, gmv: 156800, actualRoi: 5.48 },
  { id: 'U002', name: '新品短视频全域推广', type: 'video', status: 'running', budget: 15000, roiTarget: 3.5, cost: 12800, gmv: 58600, actualRoi: 4.58 },
  { id: 'U003', name: '日常直播全域托管', type: 'live', status: 'running', budget: 20000, roiTarget: 4.0, cost: 18500, gmv: 86500, actualRoi: 4.68 },
  { id: 'U004', name: '爆款商品推广', type: 'video', status: 'pause', budget: 10000, roiTarget: 5.0, cost: 8600, gmv: 38600, actualRoi: 4.49 },
  { id: 'U005', name: '清仓特卖直播', type: 'live', status: 'end', budget: 5000, roiTarget: 3.0, cost: 5000, gmv: 16300, actualRoi: 3.26 }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    running: 'bg-green-100 text-green-800',
    pause: 'bg-orange-100 text-orange-800',
    end: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    running: '投放中',
    pause: '已暂停',
    end: '已结束'
  }
  return texts[status] || status
}

const handleSearch = () => {
  console.log('搜索:', filters.value)
  alert('查询完成')
}

const handleEdit = (item: typeof promotions.value[0]) => {
  alert(`编辑推广: ${item.name}`)
}

const handlePause = (item: typeof promotions.value[0]) => {
  item.status = 'pause'
  alert(`推广「${item.name}」已暂停`)
}

const handleStart = (item: typeof promotions.value[0]) => {
  if (item.status !== 'end') {
    item.status = 'running'
    alert(`推广「${item.name}」已启动`)
  } else {
    alert('已结束的推广无法重新启动')
  }
}
</script>
