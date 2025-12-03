<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const saving = ref(false)

const campaignId = computed(() => route.params.id as string)

interface CampaignForm {
  name: string
  status: string
  budget: number
  budgetMode: string
  startDate: string
  endDate: string
  targetingType: string
  bidType: string
  bidAmount: number
  landingPage: string
  trackingUrl: string
}

const form = ref<CampaignForm>({
  name: '',
  status: 'active',
  budget: 1000,
  budgetMode: 'daily',
  startDate: '',
  endDate: '',
  targetingType: 'auto',
  bidType: 'ocpm',
  bidAmount: 10,
  landingPage: '',
  trackingUrl: ''
})

const statusOptions = [
  { value: 'active', label: '投放中' },
  { value: 'paused', label: '已暂停' },
  { value: 'draft', label: '草稿' }
]

const budgetModeOptions = [
  { value: 'daily', label: '日预算' },
  { value: 'total', label: '总预算' }
]

const targetingOptions = [
  { value: 'auto', label: '自动定向' },
  { value: 'custom', label: '自定义定向' }
]

const bidTypeOptions = [
  { value: 'ocpm', label: 'OCPM (优化千次展示)' },
  { value: 'ocpc', label: 'OCPC (优化点击)' },
  { value: 'cpm', label: 'CPM (千次展示)' },
  { value: 'cpc', label: 'CPC (点击)' }
]

const fetchCampaign = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // 模拟获取数据
  form.value = {
    name: '双十一促销活动',
    status: 'active',
    budget: 5000,
    budgetMode: 'daily',
    startDate: '2024-11-01',
    endDate: '2024-11-11',
    targetingType: 'custom',
    bidType: 'ocpm',
    bidAmount: 15,
    landingPage: 'https://example.com/landing',
    trackingUrl: 'https://track.example.com/click?id=123'
  }
  
  loading.value = false
}

const handleSubmit = async () => {
  saving.value = true
  await new Promise(resolve => setTimeout(resolve, 1000))
  saving.value = false
  
  router.push(`/campaigns/${campaignId.value}`)
}

const handleCancel = () => {
  router.push(`/campaigns/${campaignId.value}`)
}

onMounted(fetchCampaign)
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告计划', path: '/campaigns' },
        { name: '编辑广告计划' }
      ]" />
      <h1 class="text-3xl font-bold text-gray-900">编辑广告计划</h1>
      <p class="mt-2 text-gray-600">修改广告计划配置信息</p>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center h-64">
      <div class="text-gray-500">加载中...</div>
    </div>

    <template v-else>
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              计划名称 <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
              placeholder="请输入计划名称"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
            <select
              v-model="form.status"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="opt in statusOptions" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
        </div>
      </div>

      <!-- 预算设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">预算设置</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">预算类型</label>
            <select
              v-model="form.budgetMode"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="opt in budgetModeOptions" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              预算金额 <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500">¥</span>
              <input
                v-model.number="form.budget"
                type="number"
                min="100"
                class="w-full pl-8 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
                placeholder="请输入预算金额"
              />
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">开始日期</label>
            <input
              v-model="form.startDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">结束日期</label>
            <input
              v-model="form.endDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- 出价设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">出价设置</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">定向方式</label>
            <select
              v-model="form.targetingType"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="opt in targetingOptions" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">出价方式</label>
            <select
              v-model="form.bidType"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="opt in bidTypeOptions" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              出价金额 <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500">¥</span>
              <input
                v-model.number="form.bidAmount"
                type="number"
                min="0.01"
                step="0.01"
                class="w-full pl-8 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
                placeholder="请输入出价金额"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- 落地页设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">落地页设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              落地页 URL <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.landingPage"
              type="url"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
              placeholder="https://example.com/landing"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">监测链接（可选）</label>
            <input
              v-model="form.trackingUrl"
              type="url"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
              placeholder="https://track.example.com/click"
            />
            <p class="mt-1 text-sm text-gray-500">用于第三方监测点击和转化</p>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex items-center justify-end gap-3">
        <button
          @click="handleCancel"
          class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        >
          取消
        </button>
        <button
          @click="handleSubmit"
          :disabled="saving"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
        >
          <svg v-if="saving" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ saving ? '保存中...' : '保存修改' }}
        </button>
      </div>
    </template>
  </div>
</template>
