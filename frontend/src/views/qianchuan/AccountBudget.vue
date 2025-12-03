<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '账户管理' }, { name: '预算管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">预算管理</h1>
        <p class="text-gray-600 mt-1">设置和管理千川广告账户预算</p>
      </div>
      <button @click="showBudgetModal = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        设置日预算
      </button>
    </div>

    <!-- 预算概览 -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">账户日预算</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">¥{{ budgetInfo.dailyBudget.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-1">不限: {{ budgetInfo.dailyBudget === 0 ? '是' : '否' }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日已消耗</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">¥{{ budgetInfo.todayCost.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-1">消耗进度: {{ budgetProgress }}%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">预算剩余</div>
        <div class="text-2xl font-bold text-green-600 mt-1">¥{{ budgetRemaining.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-1">预计可投放至: {{ estimatedEndTime }}</div>
      </div>
    </div>

    <!-- 预算消耗趋势 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <h3 class="text-lg font-medium mb-4">预算消耗趋势</h3>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
        <span class="text-gray-400">消耗趋势图表</span>
      </div>
    </div>

    <!-- 历史预算记录 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b">
        <h3 class="text-lg font-medium">预算调整记录</h3>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">调整时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">调整前</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">调整后</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作人</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">备注</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="record in budgetHistory" :key="record.id">
            <td class="px-4 py-3 text-sm">{{ record.time }}</td>
            <td class="px-4 py-3 text-sm">¥{{ record.before.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">¥{{ record.after.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ record.operator }}</td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ record.remark }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 设置预算弹窗 -->
    <div v-if="showBudgetModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-96">
        <h3 class="text-lg font-medium mb-4">设置日预算</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">预算类型</label>
            <select v-model="newBudget.type" class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option value="limited">指定预算</option>
              <option value="unlimited">不限预算</option>
            </select>
          </div>
          <div v-if="newBudget.type === 'limited'">
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算金额</label>
            <input type="number" v-model="newBudget.amount" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="请输入预算金额">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">备注</label>
            <input type="text" v-model="newBudget.remark" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="可选">
          </div>
        </div>
        <div class="flex justify-end space-x-3 mt-6">
          <button @click="showBudgetModal = false" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="saveBudget" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">确定</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const showBudgetModal = ref(false)

const budgetInfo = ref({
  dailyBudget: 50000,
  todayCost: 28600,
  avgHourlyCost: 2800
})

const newBudget = ref({
  type: 'limited',
  amount: 50000,
  remark: ''
})

const budgetProgress = computed(() => {
  if (budgetInfo.value.dailyBudget === 0) return 0
  return Math.round((budgetInfo.value.todayCost / budgetInfo.value.dailyBudget) * 100)
})

const budgetRemaining = computed(() => {
  return Math.max(0, budgetInfo.value.dailyBudget - budgetInfo.value.todayCost)
})

const estimatedEndTime = computed(() => {
  const remaining = budgetRemaining.value
  const hourlyRate = budgetInfo.value.avgHourlyCost
  if (hourlyRate === 0) return '24:00'
  const hoursLeft = remaining / hourlyRate
  const now = new Date()
  now.setHours(now.getHours() + hoursLeft)
  return now.toTimeString().slice(0, 5)
})

const budgetHistory = ref([
  { id: 1, time: '2024-03-15 14:30:00', before: 30000, after: 50000, operator: '管理员', remark: '增加投放预算' },
  { id: 2, time: '2024-03-10 09:15:00', before: 20000, after: 30000, operator: '管理员', remark: '活动期间提量' },
  { id: 3, time: '2024-03-01 10:00:00', before: 0, after: 20000, operator: '系统', remark: '首次设置预算' }
])

const saveBudget = () => {
  showBudgetModal.value = false
}
</script>
