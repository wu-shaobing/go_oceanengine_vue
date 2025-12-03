<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const saving = ref(false)
const currentStep = ref(1)

const form = reactive({
  name: '',
  objective: 'conversion',
  budget: '',
  budgetType: 'daily',
  startDate: '',
  endDate: '',
  bidType: 'cpc',
  bidAmount: '',
  targetAudience: 'all',
  placements: ['feed', 'video']
})

const steps = [
  { id: 1, name: '基本信息', description: '设置广告系列名称和目标' },
  { id: 2, name: '预算与排期', description: '设置预算和投放时间' },
  { id: 3, name: '投放设置', description: '设置出价和定向' }
]

const objectives = [
  { value: 'conversion', label: '转化', description: '优化转化效果' },
  { value: 'traffic', label: '流量', description: '获取更多点击' },
  { value: 'brand', label: '品牌', description: '提升品牌认知' }
]

const placements = [
  { value: 'feed', label: '信息流' },
  { value: 'video', label: '短视频' },
  { value: 'search', label: '搜索广告' },
  { value: 'splash', label: '开屏广告' }
]

const nextStep = () => {
  if (currentStep.value < 3) {
    currentStep.value++
  }
}

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

const handleSubmit = async () => {
  saving.value = true
  await new Promise(resolve => setTimeout(resolve, 1500))
  saving.value = false
  router.push('/campaigns')
}

const togglePlacement = (value: string) => {
  const index = form.placements.indexOf(value)
  if (index > -1) {
    form.placements.splice(index, 1)
  } else {
    form.placements.push(value)
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告系列', path: '/campaigns' },
        { name: '创建广告系列' }
      ]" />
      <h1 class="text-3xl font-bold text-gray-900">创建广告系列</h1>
      <p class="mt-2 text-gray-600">按步骤创建新的广告投放系列</p>
    </div>

    <!-- Steps -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <nav aria-label="Progress">
        <ol class="flex items-center">
          <li v-for="(step, index) in steps" :key="step.id" class="relative" :class="{ 'flex-1': index < steps.length - 1 }">
            <div class="flex items-center">
              <span
                class="w-10 h-10 flex items-center justify-center rounded-full border-2 transition-colors"
                :class="{
                  'bg-blue-600 border-blue-600 text-white': currentStep >= step.id,
                  'border-gray-300 text-gray-500': currentStep < step.id
                }"
              >
                {{ step.id }}
              </span>
              <div class="ml-4" :class="{ 'hidden sm:block': index < steps.length - 1 }">
                <p class="text-sm font-medium" :class="{ 'text-blue-600': currentStep >= step.id, 'text-gray-500': currentStep < step.id }">
                  {{ step.name }}
                </p>
                <p class="text-xs text-gray-500">{{ step.description }}</p>
              </div>
            </div>
            <div v-if="index < steps.length - 1" class="hidden sm:block absolute top-5 left-14 w-full h-0.5 bg-gray-200">
              <div class="h-full bg-blue-600 transition-all" :style="{ width: currentStep > step.id ? '100%' : '0%' }"></div>
            </div>
          </li>
        </ol>
      </nav>
    </div>

    <!-- Form -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <!-- Step 1: Basic Info -->
      <div v-show="currentStep === 1" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">广告系列名称</label>
          <input
            v-model="form.name"
            type="text"
            placeholder="输入广告系列名称"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">推广目标</label>
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
            <div
              v-for="obj in objectives"
              :key="obj.value"
              @click="form.objective = obj.value"
              class="p-4 border rounded-lg cursor-pointer transition-colors"
              :class="{
                'border-blue-500 bg-blue-50': form.objective === obj.value,
                'border-gray-200 hover:border-gray-300': form.objective !== obj.value
              }"
            >
              <p class="font-medium" :class="{ 'text-blue-600': form.objective === obj.value }">
                {{ obj.label }}
              </p>
              <p class="text-sm text-gray-500 mt-1">{{ obj.description }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Step 2: Budget & Schedule -->
      <div v-show="currentStep === 2" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">预算类型</label>
          <div class="flex gap-4">
            <label class="flex items-center">
              <input v-model="form.budgetType" type="radio" value="daily" class="mr-2">
              <span>日预算</span>
            </label>
            <label class="flex items-center">
              <input v-model="form.budgetType" type="radio" value="total" class="mr-2">
              <span>总预算</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            {{ form.budgetType === 'daily' ? '日预算金额' : '总预算金额' }}
          </label>
          <div class="relative">
            <span class="absolute left-4 top-2 text-gray-500">¥</span>
            <input
              v-model="form.budget"
              type="number"
              placeholder="输入预算金额"
              class="w-full pl-8 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">开始日期</label>
            <input
              v-model="form.startDate"
              type="date"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">结束日期</label>
            <input
              v-model="form.endDate"
              type="date"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
          </div>
        </div>
      </div>

      <!-- Step 3: Targeting -->
      <div v-show="currentStep === 3" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">出价方式</label>
          <select
            v-model="form.bidType"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="cpc">按点击付费 (CPC)</option>
            <option value="cpm">按展示付费 (CPM)</option>
            <option value="ocpc">优化点击付费 (oCPC)</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">出价金额</label>
          <div class="relative">
            <span class="absolute left-4 top-2 text-gray-500">¥</span>
            <input
              v-model="form.bidAmount"
              type="number"
              step="0.01"
              placeholder="输入出价金额"
              class="w-full pl-8 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">投放版位</label>
          <div class="grid grid-cols-2 gap-3">
            <div
              v-for="placement in placements"
              :key="placement.value"
              @click="togglePlacement(placement.value)"
              class="p-3 border rounded-lg cursor-pointer transition-colors flex items-center"
              :class="{
                'border-blue-500 bg-blue-50': form.placements.includes(placement.value),
                'border-gray-200 hover:border-gray-300': !form.placements.includes(placement.value)
              }"
            >
              <input
                type="checkbox"
                :checked="form.placements.includes(placement.value)"
                class="mr-2"
                @click.stop
              >
              <span :class="{ 'text-blue-600': form.placements.includes(placement.value) }">
                {{ placement.label }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex items-center justify-between mt-8 pt-6 border-t border-gray-200">
        <button
          v-if="currentStep > 1"
          @click="prevStep"
          class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        >
          上一步
        </button>
        <div v-else></div>

        <div class="flex gap-3">
          <router-link
            to="/campaigns"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
          >
            取消
          </router-link>
          <button
            v-if="currentStep < 3"
            @click="nextStep"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            下一步
          </button>
          <button
            v-else
            @click="handleSubmit"
            :disabled="saving"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
          >
            {{ saving ? '创建中...' : '创建广告系列' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
