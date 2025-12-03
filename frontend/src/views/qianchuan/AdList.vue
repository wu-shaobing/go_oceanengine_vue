<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '广告列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">广告列表</h1>
        <p class="text-gray-600 mt-1">管理所有千川广告</p>
      </div>
      <router-link to="/qianchuan/ad/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建广告
      </router-link>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="广告名称/ID" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="delivery">投放中</option>
          <option value="audit">审核中</option>
          <option value="reject">审核拒绝</option>
          <option value="disable">已暂停</option>
        </select>
        <select v-model="filters.campaignId" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部计划</option>
          <option value="10001">618大促直播计划</option>
          <option value="10002">新品短视频推广</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 广告列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">广告信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">所属计划</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">出价</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">曝光</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="ad in ads" :key="ad.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="ad.cover" class="w-16 h-10 rounded mr-3 object-cover" alt="">
                <div>
                  <div class="font-medium">{{ ad.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ ad.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ ad.campaignName }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(ad.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(ad.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">¥{{ ad.bid }}</td>
            <td class="px-4 py-3 text-sm">¥{{ ad.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ ad.impression.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ ad.click.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ ad.conversion }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <router-link :to="`/qianchuan/ad/${ad.id}`" class="text-blue-600 hover:text-blue-800 text-sm">详情</router-link>
                <button class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">复制</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="200" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { qianchuanApi } from '@/api/qianchuan'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()

const filters = ref({
  keyword: '',
  status: '',
  campaignId: ''
})

const loading = ref(false)
const error = ref('')
const pagination = ref({ page: 1, pageSize: 10, total: 0 })

interface AdItem {
  id: string
  name: string
  cover: string
  campaignName: string
  status: string
  bid: number
  cost: number
  impression: number
  click: number
  conversion: number
}

const ads = ref<AdItem[]>([])

const fetchAds = async () => {
  const advertiserId = advertiserStore.currentAdvertiserId
  if (!advertiserId) {
    error.value = '请先选择广告主'
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const params: { advertiser_id: number; page: number; page_size: number; status?: string } = {
      advertiser_id: advertiserId,
      page: pagination.value.page,
      page_size: pagination.value.pageSize
    }
    if (filters.value.status) {
      params.status = filters.value.status
    }
    
    const res = await qianchuanApi.getAdList(params)
    const data = res as any
    
    if (data?.list) {
      ads.value = data.list.map((item: any) => ({
        id: String(item.ad_id),
        name: item.ad_name,
        cover: 'https://via.placeholder.com/160x100',
        campaignName: item.campaign_name || '-',
        status: item.status || 'unknown',
        bid: item.bid || 0,
        cost: item.cost || 0,
        impression: item.show_cnt || 0,
        click: item.click_cnt || 0,
        conversion: item.convert_cnt || 0
      }))
      pagination.value.total = data.page_info?.total_number || 0
    }
  } catch (e: any) {
    error.value = e.message || '加载失败'
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.page = 1
  fetchAds()
}

const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchAds()
}

onMounted(() => {
  fetchAds()
})

defineExpose({ handleSearch, handlePageChange })

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    delivery: 'bg-green-100 text-green-800',
    audit: 'bg-yellow-100 text-yellow-800',
    reject: 'bg-red-100 text-red-800',
    disable: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    delivery: '投放中',
    audit: '审核中',
    reject: '审核拒绝',
    disable: '已暂停'
  }
  return texts[status] || status
}
</script>
