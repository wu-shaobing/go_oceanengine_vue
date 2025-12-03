<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '数据报表' }, { name: '项目报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">项目报表</h1>
      <p class="text-gray-600 mt-1">查看项目维度的投放数据</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总消耗</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.cost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总线索</div>
        <div class="text-xl font-bold mt-1">{{ summary.clues.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">线索成本</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.clueCost }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">有效线索率</div>
        <div class="text-xl font-bold mt-1">{{ summary.validRate }}%</div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">项目</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">展示</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索成本</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">有效线索</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in projects" :key="item.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ item.name }}</div>
              <div class="text-sm text-gray-500">{{ item.city }}</div>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.show.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.click.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.clues }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.clueCost }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.validClues }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="30" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '' })

const summary = ref({ cost: 125680, clues: 2856, clueCost: 44.0, validRate: 68.5 })

const projects = ref([
  { id: 1, name: '北京朝阳店-618大促', city: '北京市', cost: 45680, show: 580000, click: 18500, clues: 986, clueCost: 46.3, validClues: 685 },
  { id: 2, name: '上海静安店', city: '上海市', cost: 38560, show: 456000, click: 15200, clues: 856, clueCost: 45.0, validClues: 598 },
  { id: 3, name: '广州天河店', city: '广州市', cost: 25680, show: 320000, click: 10500, clues: 625, clueCost: 41.1, validClues: 425 },
  { id: 4, name: '深圳福田店', city: '深圳市', cost: 15760, show: 198000, click: 6800, clues: 389, clueCost: 40.5, validClues: 268 }
])
</script>
