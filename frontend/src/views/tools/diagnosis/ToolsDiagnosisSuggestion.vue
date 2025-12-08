<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  adId: ''
})

const diagnosisResult = ref<null | {
  score: number
  level: string
  issues: Array<{ type: string; severity: string; description: string; suggestion: string }>
}>(null)

const handleDiagnose = () => {
  diagnosisResult.value = {
    score: 72,
    level: '一般',
    issues: [
      { type: '出价', severity: 'high', description: '出价低于行业均值30%', suggestion: '建议将出价提升至50元以上' },
      { type: '素材', severity: 'medium', description: '视频素材点击率偏低', suggestion: '建议更换开头3秒内容，增加吸引力' },
      { type: '定向', severity: 'low', description: '定向范围较窄', suggestion: '可适当放宽年龄定向' }
    ]
  }
}

const getSeverityClass = (severity: string) => {
  switch (severity) {
    case 'high': return 'bg-red-100 text-red-700'
    case 'medium': return 'bg-yellow-100 text-yellow-700'
    default: return 'bg-blue-100 text-blue-700'
  }
}

const handleAccept = (issue: { type: string; suggestion: string }) => {
  alert(`采纳建议: ${issue.type} - ${issue.suggestion}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '诊断工具' }, { name: '计划诊断' }]" />
      <h1 class="text-3xl font-bold text-gray-900">计划诊断建议</h1>
      <p class="mt-2 text-gray-600">获取广告计划的诊断分析和优化建议</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleDiagnose" class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <input v-model="form.advertiserId" type="text" placeholder="广告主ID"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div class="flex-1 min-w-[200px]">
          <input v-model="form.adId" type="text" placeholder="广告计划ID"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          开始诊断
        </button>
      </form>
    </div>

    <div v-if="diagnosisResult" class="space-y-6">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="text-center">
            <p class="text-sm text-gray-500">健康分数</p>
            <p class="text-5xl font-bold mt-2" :class="diagnosisResult.score >= 80 ? 'text-green-600' : diagnosisResult.score >= 60 ? 'text-yellow-600' : 'text-red-600'">
              {{ diagnosisResult.score }}
            </p>
          </div>
        </div>
        <div class="lg:col-span-3 bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="font-semibold text-gray-900 mb-4">诊断结果: {{ diagnosisResult.level }}</h3>
          <div class="flex items-center gap-2 mb-4">
            <span class="px-2 py-1 bg-red-100 text-red-700 rounded text-xs">高优先级 {{ diagnosisResult.issues.filter(i => i.severity === 'high').length }}</span>
            <span class="px-2 py-1 bg-yellow-100 text-yellow-700 rounded text-xs">中优先级 {{ diagnosisResult.issues.filter(i => i.severity === 'medium').length }}</span>
            <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">低优先级 {{ diagnosisResult.issues.filter(i => i.severity === 'low').length }}</span>
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div v-for="(issue, idx) in diagnosisResult.issues" :key="idx" 
             class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-start gap-4">
            <span :class="['px-2 py-1 rounded text-xs font-medium', getSeverityClass(issue.severity)]">
              {{ issue.type }}
            </span>
            <div class="flex-1">
              <p class="font-medium text-gray-900">{{ issue.description }}</p>
              <p class="text-sm text-gray-600 mt-2">
                <span class="font-medium text-blue-600">建议：</span>{{ issue.suggestion }}
              </p>
            </div>
<button class="px-4 py-1.5 bg-blue-600 text-white text-sm rounded hover:bg-blue-700" @click="handleAccept(issue)">
              一键采纳
            </button>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center text-gray-500">
      输入广告主ID和计划ID后开始诊断
    </div>
  </div>
</template>
