<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '账户管理' }, { name: '达人授权' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">达人授权管理</h1>
        <p class="text-gray-600 mt-1">管理抖音达人账号授权</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        添加达人授权
      </button>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已授权达人</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stats.totalAweme }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">授权有效</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.activeAweme }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">即将过期</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.expiringAweme }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已过期</div>
        <div class="text-2xl font-bold text-red-600 mt-1">{{ stats.expiredAweme }}</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="达人昵称/抖音号" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="active">授权有效</option>
          <option value="expiring">即将过期</option>
          <option value="expired">已过期</option>
        </select>
        <select v-model="filters.shop" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部店铺</option>
          <option value="7001">品牌官方旗舰店</option>
          <option value="7002">美妆专营店</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 达人列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">达人信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">粉丝数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">关联店铺</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">授权状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">授权到期</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">投放权限</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="aweme in awemeList" :key="aweme.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="aweme.avatar" class="w-10 h-10 rounded-full mr-3" alt="">
                <div>
                  <div class="font-medium">{{ aweme.nickname }}</div>
                  <div class="text-sm text-gray-500">@{{ aweme.uniqueId }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ formatFans(aweme.fans) }}</td>
            <td class="px-4 py-3 text-sm">{{ aweme.shopName }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(aweme.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(aweme.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">{{ aweme.expireTime }}</td>
            <td class="px-4 py-3">
              <div class="flex flex-wrap gap-1">
                <span v-for="perm in aweme.permissions" :key="perm" class="px-2 py-0.5 bg-blue-100 text-blue-800 text-xs rounded">
                  {{ perm }}
                </span>
              </div>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">编辑权限</button>
                <button class="text-blue-600 hover:text-blue-800 text-sm">续期</button>
                <button class="text-red-600 hover:text-red-800 text-sm">取消授权</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const stats = ref({
  totalAweme: 48,
  activeAweme: 42,
  expiringAweme: 3,
  expiredAweme: 3
})

const filters = ref({
  keyword: '',
  status: '',
  shop: ''
})

const awemeList = ref([
  { id: 1, nickname: '美妆达人小美', uniqueId: 'beauty_xiaomei', avatar: 'https://via.placeholder.com/40', fans: 1560000, shopName: '美妆专营店', status: 'active', expireTime: '2024-12-31', permissions: ['直播', '短视频', '商品'] },
  { id: 2, nickname: '时尚穿搭师', uniqueId: 'fashion_master', avatar: 'https://via.placeholder.com/40', fans: 2380000, shopName: '服装精品店', status: 'active', expireTime: '2024-10-15', permissions: ['直播', '短视频'] },
  { id: 3, nickname: '美食探店家', uniqueId: 'food_explorer', avatar: 'https://via.placeholder.com/40', fans: 890000, shopName: '食品特卖店', status: 'expiring', expireTime: '2024-04-01', permissions: ['短视频', '商品'] },
  { id: 4, nickname: '数码测评师', uniqueId: 'tech_review', avatar: 'https://via.placeholder.com/40', fans: 3250000, shopName: '数码配件店', status: 'active', expireTime: '2024-08-20', permissions: ['直播', '短视频', '商品'] },
  { id: 5, nickname: '生活好物分享', uniqueId: 'life_good', avatar: 'https://via.placeholder.com/40', fans: 560000, shopName: '品牌官方旗舰店', status: 'expired', expireTime: '2024-02-15', permissions: ['短视频'] }
])

const formatFans = (fans: number) => {
  if (fans >= 10000) {
    return (fans / 10000).toFixed(1) + 'w'
  }
  return fans.toString()
}

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    active: 'bg-green-100 text-green-800',
    expiring: 'bg-orange-100 text-orange-800',
    expired: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    active: '授权有效',
    expiring: '即将过期',
    expired: '已过期'
  }
  return texts[status] || status
}
</script>
