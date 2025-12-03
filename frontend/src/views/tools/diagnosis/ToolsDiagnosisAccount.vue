<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const accountHealth = ref({
  score: 85,
  status: 'good',
  lastCheck: '2025-11-28 10:00'
})

const issues = ref([
  { id: 'I001', type: 'warning', title: '预算即将耗尽', desc: '3个广告组预算剩余不足20%', impact: '中', action: '调整预算' },
  { id: 'I002', type: 'info', title: '出价建议', desc: '5个广告出价低于建议值', impact: '低', action: '优化出价' },
  { id: 'I003', type: 'error', title: '素材审核未通过', desc: '2个创意审核被拒', impact: '高', action: '修改素材' }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '诊断工具' }, { name: '账户诊断' }]" />
      <h1 class="text-3xl font-bold text-gray-900">账户诊断</h1>
      <p class="mt-2 text-gray-600">全面分析账户健康状况</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="text-center">
          <div class="relative inline-flex items-center justify-center">
            <svg class="w-32 h-32">
              <circle cx="64" cy="64" r="56" stroke="#e5e7eb" stroke-width="8" fill="none" />
              <circle cx="64" cy="64" r="56" stroke="#22c55e" stroke-width="8" fill="none"
                      stroke-dasharray="352" :stroke-dashoffset="352 - (352 * accountHealth.score / 100)"
                      transform="rotate(-90 64 64)" />
            </svg>
            <span class="absolute text-3xl font-bold text-gray-900">{{ accountHealth.score }}</span>
          </div>
          <p class="mt-4 text-lg font-medium text-green-600">健康状况良好</p>
          <p class="text-sm text-gray-500 mt-1">上次检查: {{ accountHealth.lastCheck }}</p>
        </div>
      </div>

      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">诊断概览</h3>
        <div class="grid grid-cols-4 gap-4">
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-2xl font-bold text-gray-900">15</p>
            <p class="text-sm text-gray-500">广告计划</p>
          </div>
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-2xl font-bold text-green-600">12</p>
            <p class="text-sm text-gray-500">正常运行</p>
          </div>
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-2xl font-bold text-yellow-600">2</p>
            <p class="text-sm text-gray-500">需要关注</p>
          </div>
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-2xl font-bold text-red-600">1</p>
            <p class="text-sm text-gray-500">存在问题</p>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-900">待处理问题</h3>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 text-sm">
          重新诊断
        </button>
      </div>
      <div class="divide-y divide-gray-200">
        <div v-for="issue in issues" :key="issue.id" class="px-6 py-4 flex items-center">
          <div :class="['w-10 h-10 rounded-full flex items-center justify-center mr-4',
                 issue.type === 'error' ? 'bg-red-100' : issue.type === 'warning' ? 'bg-yellow-100' : 'bg-blue-100']">
            <span :class="['text-lg',
                   issue.type === 'error' ? 'text-red-600' : issue.type === 'warning' ? 'text-yellow-600' : 'text-blue-600']">
              {{ issue.type === 'error' ? '❌' : issue.type === 'warning' ? '⚠️' : 'ℹ️' }}
            </span>
          </div>
          <div class="flex-1">
            <div class="flex items-center">
              <h4 class="text-sm font-medium text-gray-900">{{ issue.title }}</h4>
              <span :class="['ml-2 px-2 py-0.5 rounded text-xs',
                     issue.impact === '高' ? 'bg-red-100 text-red-700' :
                     issue.impact === '中' ? 'bg-yellow-100 text-yellow-700' : 'bg-gray-100 text-gray-700']">
                影响: {{ issue.impact }}
              </span>
            </div>
            <p class="text-sm text-gray-500 mt-1">{{ issue.desc }}</p>
          </div>
          <button class="px-4 py-2 text-sm text-blue-600 border border-blue-300 rounded-lg hover:bg-blue-50">
            {{ issue.action }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
