<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'

const route = useRoute()
const router = useRouter()

const advertiserId = computed(() => route.params.id as string)
const isEditMode = computed(() => route.query.mode === 'edit')
const loading = ref(true)
const saving = ref(false)

const advertiser = ref({
  id: '',
  name: '',
  company: '',
  status: 'enabled',
  balance: 0,
  email: '',
  phone: '',
  address: '',
  industry: '',
  createTime: '',
  description: ''
})

const stats = ref({
  totalSpend: 0,
  avgCpc: 0,
  totalImpressions: 0,
  totalClicks: 0
})

const fetchAdvertiserDetail = async () => {
  loading.value = true
  // 模拟API调用
  await new Promise(resolve => setTimeout(resolve, 500))
  
  advertiser.value = {
    id: advertiserId.value,
    name: '测试广告主',
    company: '测试科技有限公司',
    status: 'enabled',
    balance: 158000,
    email: 'contact@test.com',
    phone: '13800138000',
    address: '北京市朝阳区xxx路xx号',
    industry: '电商零售',
    createTime: '2024-01-15',
    description: '主营电商业务的广告主'
  }
  
  stats.value = {
    totalSpend: 125600,
    avgCpc: 1.58,
    totalImpressions: 2580000,
    totalClicks: 85200
  }
  
  loading.value = false
}

const handleSave = async () => {
  saving.value = true
  await new Promise(resolve => setTimeout(resolve, 1000))
  saving.value = false
  router.push(`/advertisers/${advertiserId.value}`)
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

onMounted(() => {
  fetchAdvertiserDetail()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告主管理', path: '/advertisers' },
        { name: isEditMode ? '编辑广告主' : '广告主详情' }
      ]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">
            {{ isEditMode ? '编辑广告主' : '广告主详情' }}
          </h1>
          <p class="mt-2 text-gray-600">
            {{ isEditMode ? '修改广告主信息' : `广告主 ID: ${advertiserId}` }}
          </p>
        </div>
        <div class="flex items-center gap-3">
          <router-link
            v-if="!isEditMode"
            :to="`/advertisers/${advertiserId}?mode=edit`"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            编辑
          </router-link>
          <button
            v-if="isEditMode"
            @click="handleSave"
            :disabled="saving"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
          >
            {{ saving ? '保存中...' : '保存' }}
          </button>
          <button
            @click="router.back()"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
          >
            返回
          </button>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>

    <template v-else>
      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
        <StatsCard title="账户余额" :value="formatMoney(advertiser.balance)" color="blue">
          <template #icon>
            <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
          </template>
        </StatsCard>

        <StatsCard title="总消耗" :value="formatMoney(stats.totalSpend)" color="orange">
          <template #icon>
            <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/>
            </svg>
          </template>
        </StatsCard>

        <StatsCard title="总展示" :value="formatNumber(stats.totalImpressions)" color="purple">
          <template #icon>
            <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
            </svg>
          </template>
        </StatsCard>

        <StatsCard title="总点击" :value="formatNumber(stats.totalClicks)" color="green">
          <template #icon>
            <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122"/>
            </svg>
          </template>
        </StatsCard>
      </div>

      <!-- Detail Card -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">基本信息</h3>
          <StatusBadge
            :status="advertiser.status === 'enabled' ? 'success' : 'error'"
            :text="advertiser.status === 'enabled' ? '启用' : '禁用'"
            show-icon
          />
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">广告主名称</label>
            <input
              v-if="isEditMode"
              v-model="advertiser.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
            <p v-else class="text-gray-900">{{ advertiser.name }}</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">公司名称</label>
            <input
              v-if="isEditMode"
              v-model="advertiser.company"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
            <p v-else class="text-gray-900">{{ advertiser.company }}</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">所属行业</label>
            <select
              v-if="isEditMode"
              v-model="advertiser.industry"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="电商零售">电商零售</option>
              <option value="游戏">游戏</option>
              <option value="教育">教育</option>
              <option value="金融">金融</option>
              <option value="房产">房产</option>
              <option value="其他">其他</option>
            </select>
            <p v-else class="text-gray-900">{{ advertiser.industry }}</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">联系邮箱</label>
            <input
              v-if="isEditMode"
              v-model="advertiser.email"
              type="email"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
            <p v-else class="text-gray-900">{{ advertiser.email }}</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">联系电话</label>
            <input
              v-if="isEditMode"
              v-model="advertiser.phone"
              type="tel"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
            <p v-else class="text-gray-900">{{ advertiser.phone }}</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">创建时间</label>
            <p class="text-gray-900">{{ advertiser.createTime }}</p>
          </div>

          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-1">联系地址</label>
            <input
              v-if="isEditMode"
              v-model="advertiser.address"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
            <p v-else class="text-gray-900">{{ advertiser.address }}</p>
          </div>

          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-1">描述</label>
            <textarea
              v-if="isEditMode"
              v-model="advertiser.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            ></textarea>
            <p v-else class="text-gray-900">{{ advertiser.description }}</p>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>
