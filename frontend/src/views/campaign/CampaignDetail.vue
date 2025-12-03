<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import DataTable from '@/components/common/DataTable.vue'

const route = useRoute()
const router = useRouter()

const campaignId = computed(() => route.params.id as string)
const loading = ref(true)

const campaign = ref({
  id: '',
  name: '',
  status: 'active',
  objective: 'conversion',
  budget: 50000,
  budgetType: 'daily',
  spend: 32580,
  startDate: '2024-01-15',
  endDate: '2024-03-15'
})

const stats = ref({
  impressions: 1580000,
  clicks: 52600,
  ctr: 3.33,
  conversions: 2150,
  cvr: 4.09,
  cpc: 0.62
})

const ads = ref<any[]>([])

const adColumns = [
  { key: 'id', title: '广告ID' },
  { key: 'name', title: '广告名称' },
  { key: 'impressions', title: '展示', align: 'right' as const },
  { key: 'clicks', title: '点击', align: 'right' as const },
  { key: 'spend', title: '消耗', align: 'right' as const },
  { key: 'status', title: '状态' }
]

const fetchCampaignDetail = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  campaign.value = {
    id: campaignId.value,
    name: '双十一大促广告系列',
    status: 'active',
    objective: 'conversion',
    budget: 50000,
    budgetType: 'daily',
    spend: 32580,
    startDate: '2024-01-15',
    endDate: '2024-03-15'
  }
  
  ads.value = Array.from({ length: 5 }, (_, i) => ({
    id: `AD${200001 + i}`,
    name: `广告创意 ${i + 1}`,
    impressions: Math.floor(Math.random() * 300000) + 50000,
    clicks: Math.floor(Math.random() * 10000) + 2000,
    spend: Math.floor(Math.random() * 8000) + 2000,
    status: Math.random() > 0.2 ? 'active' : 'paused'
  }))
  
  loading.value = false
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

const toggleStatus = () => {
  campaign.value.status = campaign.value.status === 'active' ? 'paused' : 'active'
}

onMounted(() => {
  fetchCampaignDetail()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告系列', path: '/campaigns' },
        { name: '广告系列详情' }
      ]" />
      <div class="flex items-center justify-between">
        <div>
          <div class="flex items-center gap-3">
            <h1 class="text-3xl font-bold text-gray-900">{{ campaign.name || '加载中...' }}</h1>
            <StatusBadge
              v-if="!loading"
              :status="campaign.status === 'active' ? 'success' : 'warning'"
              :text="campaign.status === 'active' ? '投放中' : '已暂停'"
              show-icon
            />
          </div>
          <p class="mt-2 text-gray-600">广告系列 ID: {{ campaignId }}</p>
        </div>
        <div class="flex items-center gap-3">
          <button
            @click="toggleStatus"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
          >
            {{ campaign.status === 'active' ? '暂停投放' : '开始投放' }}
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

    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>

    <template v-else>
      <!-- Performance Stats -->
      <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
        <StatsCard title="展示量" :value="formatNumber(stats.impressions)" color="blue" />
        <StatsCard title="点击量" :value="formatNumber(stats.clicks)" color="green" />
        <StatsCard title="点击率" :value="`${stats.ctr}%`" color="purple" />
        <StatsCard title="转化数" :value="formatNumber(stats.conversions)" color="orange" />
        <StatsCard title="转化率" :value="`${stats.cvr}%`" color="blue" />
        <StatsCard title="平均CPC" :value="`¥${stats.cpc}`" color="green" />
      </div>

      <!-- Campaign Info -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">广告系列设置</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <p class="text-sm text-gray-500">推广目标</p>
            <p class="text-gray-900 font-medium">
              {{ campaign.objective === 'conversion' ? '转化' : campaign.objective === 'traffic' ? '流量' : '品牌' }}
            </p>
          </div>
          <div>
            <p class="text-sm text-gray-500">预算类型</p>
            <p class="text-gray-900 font-medium">
              {{ campaign.budgetType === 'daily' ? '日预算' : '总预算' }}
            </p>
          </div>
          <div>
            <p class="text-sm text-gray-500">预算金额</p>
            <p class="text-gray-900 font-medium">{{ formatMoney(campaign.budget) }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500">已消耗</p>
            <p class="text-gray-900 font-medium">{{ formatMoney(campaign.spend) }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500">开始日期</p>
            <p class="text-gray-900 font-medium">{{ campaign.startDate }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500">结束日期</p>
            <p class="text-gray-900 font-medium">{{ campaign.endDate }}</p>
          </div>
        </div>
      </div>

      <!-- Ads Table -->
      <div class="bg-white rounded-lg border border-gray-200">
        <div class="p-6 border-b border-gray-200 flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">广告列表</h3>
          <router-link
            to="/ads/create"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm"
          >
            添加广告
          </router-link>
        </div>

        <DataTable :columns="adColumns" :data="ads" :loading="false">
          <template #id="{ value }">
            <span class="font-medium text-gray-900">{{ value }}</span>
          </template>

          <template #name="{ value, row }">
            <router-link :to="`/ads/${row.id}`" class="text-blue-600 hover:text-blue-800">
              {{ value }}
            </router-link>
          </template>

          <template #impressions="{ value }">
            <span class="text-gray-600">{{ formatNumber(value) }}</span>
          </template>

          <template #clicks="{ value }">
            <span class="text-gray-600">{{ formatNumber(value) }}</span>
          </template>

          <template #spend="{ value }">
            <span class="text-gray-900">{{ formatMoney(value) }}</span>
          </template>

          <template #status="{ value }">
            <StatusBadge
              :status="value === 'active' ? 'success' : 'warning'"
              :text="value === 'active' ? '投放中' : '已暂停'"
              show-icon
            />
          </template>
        </DataTable>
      </div>
    </template>
  </div>
</template>
