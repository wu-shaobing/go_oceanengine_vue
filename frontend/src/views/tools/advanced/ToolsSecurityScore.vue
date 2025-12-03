<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: ''
})

const scoreData = ref<null | {
  totalScore: number
  maxScore: number
  level: string
  records: Array<{ date: string; type: string; score: number; reason: string }>
}>(null)

const handleQuery = () => {
  scoreData.value = {
    totalScore: 15,
    maxScore: 100,
    level: '良好',
    records: [
      { date: '2025-11-05', type: '素材违规', score: 5, reason: '广告素材含夸大宣传内容' },
      { date: '2025-10-20', type: '落地页违规', score: 5, reason: '落地页与广告内容不符' },
      { date: '2025-09-15', type: '资质问题', score: 5, reason: '资质过期未及时更新' }
    ]
  }
}

const getLevelColor = (level: string) => {
  switch (level) {
    case '优秀': return 'text-green-600'
    case '良好': return 'text-blue-600'
    case '一般': return 'text-yellow-600'
    default: return 'text-red-600'
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '违规积分' }]" />
      <h1 class="text-3xl font-bold text-gray-900">违规积分查询</h1>
      <p class="mt-2 text-gray-600">查询账户累计违规积分情况</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleQuery" class="flex gap-4">
        <div class="flex-1">
          <input v-model="form.advertiserId" type="text" placeholder="请输入广告主ID"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          查询
        </button>
      </form>
    </div>

    <div v-if="scoreData" class="space-y-6">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="text-center">
            <p class="text-sm text-gray-500">当前积分</p>
            <p class="text-5xl font-bold text-gray-900 mt-2">{{ scoreData.totalScore }}</p>
            <p class="text-sm text-gray-500 mt-2">满分 {{ scoreData.maxScore }} 分</p>
            <div class="w-full bg-gray-200 rounded-full h-3 mt-4">
              <div class="bg-blue-500 h-3 rounded-full" :style="{ width: (scoreData.totalScore / scoreData.maxScore * 100) + '%' }"></div>
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="text-center">
            <p class="text-sm text-gray-500">账户等级</p>
            <p :class="['text-4xl font-bold mt-2', getLevelColor(scoreData.level)]">{{ scoreData.level }}</p>
            <p class="text-sm text-gray-500 mt-4">积分越低，等级越好</p>
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="text-center">
            <p class="text-sm text-gray-500">违规次数</p>
            <p class="text-4xl font-bold text-orange-600 mt-2">{{ scoreData.records.length }}</p>
            <p class="text-sm text-gray-500 mt-4">近90天内</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200">
        <div class="p-4 border-b border-gray-200">
          <h3 class="font-semibold text-gray-900">违规记录</h3>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead class="bg-gray-50 border-b border-gray-200">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">日期</th>
                <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">违规类型</th>
                <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">扣分</th>
                <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">原因</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="(record, idx) in scoreData.records" :key="idx" class="hover:bg-gray-50">
                <td class="px-6 py-4 text-sm text-gray-500">{{ record.date }}</td>
                <td class="px-6 py-4">
                  <span class="px-2 py-1 bg-red-100 text-red-700 rounded text-xs">{{ record.type }}</span>
                </td>
                <td class="px-6 py-4 text-sm font-medium text-red-600">+{{ record.score }}</td>
                <td class="px-6 py-4 text-sm text-gray-700">{{ record.reason }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center text-gray-500">
      请输入广告主ID进行查询
    </div>
  </div>
</template>
