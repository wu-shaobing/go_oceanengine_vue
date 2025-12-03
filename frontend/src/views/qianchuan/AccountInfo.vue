<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '账户管理' }, { name: '账户信息' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">账户信息</h1>
      <p class="text-gray-600 mt-1">查看千川广告账户基本信息</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">基本信息</h3>
        <div class="space-y-4">
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">广告主ID</span>
            <span class="font-medium">{{ accountInfo.advertiserId }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">广告主名称</span>
            <span class="font-medium">{{ accountInfo.name }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">账户状态</span>
            <span :class="accountInfo.status === '正常' ? 'text-green-600' : 'text-red-600'" class="font-medium">
              {{ accountInfo.status }}
            </span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">账户类型</span>
            <span class="font-medium">{{ accountInfo.type }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">创建时间</span>
            <span class="font-medium">{{ accountInfo.createTime }}</span>
          </div>
        </div>
      </div>

      <!-- 资金信息 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">资金信息</h3>
        <div class="space-y-4">
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">账户余额</span>
            <span class="font-medium text-blue-600">¥{{ accountInfo.balance.toLocaleString() }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">可用余额</span>
            <span class="font-medium">¥{{ accountInfo.availableBalance.toLocaleString() }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">冻结金额</span>
            <span class="font-medium text-orange-600">¥{{ accountInfo.frozenBalance.toLocaleString() }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">赠款余额</span>
            <span class="font-medium">¥{{ accountInfo.grantBalance.toLocaleString() }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">累计消耗</span>
            <span class="font-medium">¥{{ accountInfo.totalCost.toLocaleString() }}</span>
          </div>
        </div>
      </div>

      <!-- 资质信息 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">资质信息</h3>
        <div class="space-y-4">
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">主体类型</span>
            <span class="font-medium">{{ accountInfo.entityType }}</span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">营业执照</span>
            <span :class="accountInfo.licenseStatus === '已认证' ? 'text-green-600' : 'text-orange-600'" class="font-medium">
              {{ accountInfo.licenseStatus }}
            </span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">行业资质</span>
            <span class="font-medium">{{ accountInfo.industryQualification }}</span>
          </div>
        </div>
      </div>

      <!-- 投放权限 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">投放权限</h3>
        <div class="space-y-4">
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">全域推广</span>
            <span :class="accountInfo.permissions.uniPromotion ? 'text-green-600' : 'text-gray-400'" class="font-medium">
              {{ accountInfo.permissions.uniPromotion ? '已开通' : '未开通' }}
            </span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">随心推</span>
            <span :class="accountInfo.permissions.awemeOrder ? 'text-green-600' : 'text-gray-400'" class="font-medium">
              {{ accountInfo.permissions.awemeOrder ? '已开通' : '未开通' }}
            </span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">直播间投放</span>
            <span :class="accountInfo.permissions.liveRoom ? 'text-green-600' : 'text-gray-400'" class="font-medium">
              {{ accountInfo.permissions.liveRoom ? '已开通' : '未开通' }}
            </span>
          </div>
          <div class="flex justify-between py-2 border-b">
            <span class="text-gray-500">商品投放</span>
            <span :class="accountInfo.permissions.product ? 'text-green-600' : 'text-gray-400'" class="font-medium">
              {{ accountInfo.permissions.product ? '已开通' : '未开通' }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import { qianchuanApi } from '@/api/qianchuan'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()
const loading = ref(false)
const error = ref('')

const accountInfo = ref({
  advertiserId: '',
  name: '',
  status: '',
  type: '',
  createTime: '',
  balance: 0,
  availableBalance: 0,
  frozenBalance: 0,
  grantBalance: 0,
  totalCost: 0,
  entityType: '',
  licenseStatus: '',
  industryQualification: '',
  permissions: {
    uniPromotion: false,
    awemeOrder: false,
    liveRoom: false,
    product: false
  }
})

const fetchAccountInfo = async () => {
  const advertiserId = advertiserStore.currentAdvertiserId
  if (!advertiserId) {
    error.value = '请先选择广告主账户'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const res = await qianchuanApi.getAccountInfo(advertiserId)
    if (res) {
      const data = res as any
      accountInfo.value = {
        advertiserId: String(data.advertiser_id || advertiserId),
        name: data.advertiser_name || '',
        status: data.account_type === 1 ? '正常' : '异常',
        type: '千川电商',
        createTime: data.create_time || '',
        balance: (data.balance || 0) / 100,
        availableBalance: (data.valid_balance || 0) / 100,
        frozenBalance: 0,
        grantBalance: 0,
        totalCost: 0,
        entityType: '企业',
        licenseStatus: '已认证',
        industryQualification: '',
        permissions: {
          uniPromotion: true,
          awemeOrder: true,
          liveRoom: true,
          product: true
        }
      }
    }
  } catch (e: any) {
    error.value = e.message || '获取账户信息失败'
    console.error('Failed to fetch account info:', e)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchAccountInfo()
})
</script>
