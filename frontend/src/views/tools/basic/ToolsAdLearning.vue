<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  adIds: ''
})

const results = ref<Array<{
  adId: string
  adName: string
  learningPhase: string
  learningStatus: string
  conversions: number
  targetConversions: number
  remainingDays: number
}>>([])

const handleQuery = () => {
  results.value = [
    { adId: '174812345678', adName: '新品推广-iOS端', learningPhase: 'learning', learningStatus: '学习中', conversions: 35, targetConversions: 50, remainingDays: 3 },
    { adId: '174812345679', adName: '品牌曝光-Android', learningPhase: 'learned', learningStatus: '学习完成', conversions: 58, targetConversions: 50, remainingDays: 0 },
    { adId: '174812345680', adName: '618活动推广', learningPhase: 'failed', learningStatus: '学习失败', conversions: 12, targetConversions: 50, remainingDays: 0 }
  ]
}

const getStatusClass = (phase: string) => {
  switch (phase) {
    case 'learned': return 'bg-green-100 text-green-800'
    case 'learning': return 'bg-blue-100 text-blue-800'
    case 'failed': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '学习期状态' }]" />
      <h1 class="text-3xl font-bold text-gray-900">广告计划学习期状态</h1>
      <p class="mt-2 text-gray-600">查询广告计划的学习期状态和进度</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleQuery" class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <input v-model="form.advertiserId" type="text" placeholder="广告主ID"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div class="flex-1 min-w-[300px]">
          <input v-model="form.adIds" type="text" placeholder="广告计划ID（多个用逗号分隔）"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
          查询
        </button>
      </form>
    </div>

    <div v-if="results.length" class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">计划ID</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">计划名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">学习状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">学习进度</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">剩余天数</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in results" :key="item.adId" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ item.adId }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.adName }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusClass(item.learningPhase)]">
                  {{ item.learningStatus }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="flex-1 w-32 bg-gray-200 rounded-full h-2">
                    <div :class="['h-2 rounded-full', item.learningPhase === 'learned' ? 'bg-green-500' : item.learningPhase === 'failed' ? 'bg-red-500' : 'bg-blue-500']"
                         :style="{ width: Math.min((item.conversions / item.targetConversions) * 100, 100) + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-600">{{ item.conversions }}/{{ item.targetConversions }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-sm" :class="item.remainingDays > 0 ? 'text-blue-600' : 'text-gray-500'">
                {{ item.remainingDays > 0 ? item.remainingDays + ' 天' : '-' }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center text-gray-500">
      请输入广告主ID和广告计划ID进行查询
    </div>

    <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
      <h4 class="font-medium text-blue-800 mb-2">学习期说明</h4>
      <ul class="text-sm text-blue-700 space-y-1">
        <li>• 广告计划在7天内获得50个转化即可完成学习期</li>
        <li>• 学习期间请勿频繁调整出价和定向，以免影响模型学习</li>
        <li>• 学习失败后可重置学习期，建议优化创意后再次尝试</li>
      </ul>
    </div>
  </div>
</template>
