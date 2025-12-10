<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const submitting = ref(false)
const adId = route.params.id as string

const form = reactive({
  name: '',
  campaignId: '',
  budget: 0,
  budgetMode: 'daily',
  bidType: 'cpc',
  bidAmount: 0,
  startTime: '',
  endTime: '',
  status: 'active'
})

const campaigns = ref([
  { id: '1', name: '品牌推广系列' },
  { id: '2', name: '效果投放系列' },
  { id: '3', name: '新品推广' }
])

const fetchAdDetail = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 300))
  
  // 模拟数据
  Object.assign(form, {
    name: `广告创意 ${adId}`,
    campaignId: '1',
    budget: 5000,
    budgetMode: 'daily',
    bidType: 'cpc',
    bidAmount: 2.5,
    startTime: '2025-01-01',
    endTime: '2025-12-31',
    status: 'active'
  })
  
  loading.value = false
}

const handleSubmit = async () => {
  if (!form.name) {
    alert('请输入广告名称')
    return
  }
  
  submitting.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  submitting.value = false
  
  alert('保存成功')
  router.push(`/ads/${adId}`)
}

const handleCancel = () => {
  router.back()
}

onMounted(() => {
  fetchAdDetail()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告管理', path: '/ads' },
        { name: '广告详情', path: `/ads/${adId}` },
        { name: '编辑广告' }
      ]" />
      <h1 class="text-3xl font-bold text-gray-900">编辑广告</h1>
    </div>

    <div v-if="loading" class="bg-white rounded-lg border border-gray-200 p-12 text-center">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
      <p class="mt-4 text-gray-500">加载中...</p>
    </div>

    <form v-else @submit.prevent="handleSubmit" class="space-y-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">广告名称 *</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入广告名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">所属系列</label>
            <select
              v-model="form.campaignId"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option v-for="c in campaigns" :key="c.id" :value="c.id">{{ c.name }}</option>
            </select>
          </div>
        </div>
      </div>

      <!-- 预算与出价 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">预算与出价</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算</label>
            <div class="flex">
              <span class="inline-flex items-center px-3 border border-r-0 border-gray-300 rounded-l-lg bg-gray-50 text-gray-500">¥</span>
              <input
                v-model.number="form.budget"
                type="number"
                min="0"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-r-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">出价方式</label>
            <select
              v-model="form.bidType"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="cpc">CPC (按点击)</option>
              <option value="cpm">CPM (按展示)</option>
              <option value="ocpc">OCPC (优化点击)</option>
              <option value="ocpm">OCPM (优化展示)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">出价金额</label>
            <div class="flex">
              <span class="inline-flex items-center px-3 border border-r-0 border-gray-300 rounded-l-lg bg-gray-50 text-gray-500">¥</span>
              <input
                v-model.number="form.bidAmount"
                type="number"
                min="0"
                step="0.01"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-r-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- 投放时间 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">投放时间</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">开始日期</label>
            <input
              v-model="form.startTime"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">结束日期</label>
            <input
              v-model="form.endTime"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="mt-1 text-xs text-gray-500">留空表示长期投放</p>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end gap-4">
        <button
          type="button"
          @click="handleCancel"
          class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors"
        >
          取消
        </button>
        <button
          type="submit"
          :disabled="submitting"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors"
        >
          {{ submitting ? '保存中...' : '保存' }}
        </button>
      </div>
    </form>
  </div>
</template>
