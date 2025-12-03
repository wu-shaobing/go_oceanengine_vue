<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  landingType: 'app',
  deliveryRange: 'default',
  audience: ''
})

const result = ref<null | { suggestedBudget: number; expectedCost: number; confidenceLevel: string }>(null)

const handleQuery = () => {
  result.value = {
    suggestedBudget: 5000,
    expectedCost: 45.8,
    confidenceLevel: '高'
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '建议日预算' }]" />
      <h1 class="text-3xl font-bold text-gray-900">建议日预算及预期成本</h1>
      <p class="mt-2 text-gray-600">获取广告投放建议日预算和预期转化成本</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">查询参数</h3>
        <form @submit.prevent="handleQuery" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">广告主ID</label>
            <input v-model="form.advertiserId" type="text" placeholder="请输入广告主ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">推广类型</label>
            <select v-model="form.landingType" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="app">应用下载</option>
              <option value="link">落地页推广</option>
              <option value="aweme">抖音号推广</option>
              <option value="live">直播推广</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">投放范围</label>
            <select v-model="form.deliveryRange" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="default">默认</option>
              <option value="union">穿山甲</option>
              <option value="toutiao">头条</option>
            </select>
          </div>
          <button type="submit" class="w-full px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            查询建议
          </button>
        </form>
      </div>

      <div class="space-y-4">
        <div v-if="result" class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-semibold text-gray-900 mb-4">查询结果</h3>
          <div class="space-y-4">
            <div class="flex justify-between items-center p-4 bg-blue-50 rounded-lg">
              <span class="text-gray-700">建议日预算</span>
              <span class="text-2xl font-bold text-blue-600">¥{{ result.suggestedBudget.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between items-center p-4 bg-green-50 rounded-lg">
              <span class="text-gray-700">预期转化成本</span>
              <span class="text-2xl font-bold text-green-600">¥{{ result.expectedCost }}</span>
            </div>
            <div class="flex justify-between items-center p-4 bg-gray-50 rounded-lg">
              <span class="text-gray-700">置信度</span>
              <span :class="['px-3 py-1 rounded-full text-sm font-medium',
                result.confidenceLevel === '高' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ result.confidenceLevel }}
              </span>
            </div>
          </div>
        </div>
        <div v-else class="bg-white rounded-lg border border-gray-200 p-6 text-center text-gray-500">
          请填写查询参数后点击"查询建议"
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-semibold text-gray-900 mb-3">说明</h3>
          <ul class="space-y-2 text-sm text-gray-600">
            <li>• 建议日预算基于行业数据和历史投放效果计算</li>
            <li>• 预期成本仅供参考，实际成本可能有所浮动</li>
            <li>• 置信度越高，预估越准确</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
