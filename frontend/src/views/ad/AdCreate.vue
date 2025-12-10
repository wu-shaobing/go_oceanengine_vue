<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const loading = ref(false)

const form = reactive({
  name: '',
  campaignId: '',
  budget: '',
  bidType: 'cpc',
  bidAmount: '',
  startTime: '',
  endTime: '',
  targetAge: [] as string[],
  targetGender: 'all',
  targetRegion: [] as string[],
  creativeName: '',
  creativeTitle: '',
  creativeDescription: '',
  landingUrl: ''
})

// 模拟广告系列数据
const campaigns = ref([
  { id: '1', name: '品牌推广系列' },
  { id: '2', name: '产品促销系列' },
  { id: '3', name: '新品上市系列' },
  { id: '4', name: '节日活动系列' },
  { id: '5', name: '日常投放系列' }
])

const ageOptions = [
  { value: '18-23', label: '18-23岁' },
  { value: '24-30', label: '24-30岁' },
  { value: '31-40', label: '31-40岁' },
  { value: '41-50', label: '41-50岁' },
  { value: '50+', label: '50岁以上' }
]

const regionOptions = [
  { value: 'beijing', label: '北京' },
  { value: 'shanghai', label: '上海' },
  { value: 'guangzhou', label: '广州' },
  { value: 'shenzhen', label: '深圳' },
  { value: 'hangzhou', label: '杭州' },
  { value: 'chengdu', label: '成都' },
  { value: 'wuhan', label: '武汉' },
  { value: 'xian', label: '西安' }
]

const handleSubmit = async () => {
  if (!form.name || !form.campaignId) {
    alert('请填写必填项')
    return
  }
  
  loading.value = true
  
  // 模拟提交
  await new Promise(resolve => setTimeout(resolve, 1500))
  
  alert('广告创建成功！')
  loading.value = false
  router.push('/ads')
}

const handleCancel = () => {
  router.back()
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '广告管理', path: '/ads' }, { name: '创建广告' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建广告</h1>
      <p class="mt-2 text-gray-600">创建新的广告投放计划</p>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- 基础信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">基础信息</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              广告名称 <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.name"
              type="text"
              placeholder="请输入广告名称"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              所属广告系列 <span class="text-red-500">*</span>
            </label>
            <select
              v-model="form.campaignId"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required
            >
              <option value="">请选择广告系列</option>
              <option v-for="campaign in campaigns" :key="campaign.id" :value="campaign.id">
                {{ campaign.name }}
              </option>
            </select>
          </div>
        </div>
      </div>

      <!-- 预算与出价 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">预算与出价</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">日预算（元）</label>
            <input
              v-model="form.budget"
              type="number"
              placeholder="不限"
              min="0"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            <p class="mt-1 text-sm text-gray-500">留空表示不限预算</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">出价方式</label>
            <select
              v-model="form.bidType"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="cpc">CPC（点击付费）</option>
              <option value="cpm">CPM（千次展示付费）</option>
              <option value="ocpm">oCPM（优化千次展示）</option>
              <option value="cpa">CPA（转化付费）</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">出价金额（元）</label>
            <input
              v-model="form.bidAmount"
              type="number"
              placeholder="请输入出价"
              min="0"
              step="0.01"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>
      </div>

      <!-- 投放时间 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">投放时间</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">开始时间</label>
            <input
              v-model="form.startTime"
              type="datetime-local"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">结束时间</label>
            <input
              v-model="form.endTime"
              type="datetime-local"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            <p class="mt-1 text-sm text-gray-500">留空表示长期投放</p>
          </div>
        </div>
      </div>

      <!-- 定向设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">定向设置</h2>
        
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">性别</label>
            <div class="flex gap-4">
              <label class="flex items-center">
                <input v-model="form.targetGender" type="radio" value="all" class="mr-2" />
                不限
              </label>
              <label class="flex items-center">
                <input v-model="form.targetGender" type="radio" value="male" class="mr-2" />
                男性
              </label>
              <label class="flex items-center">
                <input v-model="form.targetGender" type="radio" value="female" class="mr-2" />
                女性
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">年龄段</label>
            <div class="flex flex-wrap gap-4">
              <label v-for="option in ageOptions" :key="option.value" class="flex items-center">
                <input 
                  v-model="form.targetAge" 
                  type="checkbox" 
                  :value="option.value" 
                  class="mr-2"
                />
                {{ option.label }}
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">投放地域</label>
            <div class="flex flex-wrap gap-4">
              <label v-for="option in regionOptions" :key="option.value" class="flex items-center">
                <input 
                  v-model="form.targetRegion" 
                  type="checkbox" 
                  :value="option.value" 
                  class="mr-2"
                />
                {{ option.label }}
              </label>
            </div>
            <p class="mt-1 text-sm text-gray-500">不选择表示全国投放</p>
          </div>
        </div>
      </div>

      <!-- 创意设置 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">创意设置</h2>
        
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">创意标题</label>
            <input
              v-model="form.creativeTitle"
              type="text"
              placeholder="请输入创意标题（最多30字）"
              maxlength="30"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">创意描述</label>
            <textarea
              v-model="form.creativeDescription"
              placeholder="请输入创意描述（最多100字）"
              maxlength="100"
              rows="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">落地页链接</label>
            <input
              v-model="form.landingUrl"
              type="url"
              placeholder="https://"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">上传素材</label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center hover:border-blue-500 transition-colors cursor-pointer">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
              </svg>
              <p class="mt-2 text-sm text-gray-600">点击上传或拖拽文件到此处</p>
              <p class="mt-1 text-xs text-gray-500">支持 JPG、PNG、MP4 格式，最大 50MB</p>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end gap-4">
        <button
          type="button"
          @click="handleCancel"
          class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        >
          取消
        </button>
        <button
          type="submit"
          :disabled="loading"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
        >
          <svg v-if="loading" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ loading ? '提交中...' : '创建广告' }}
        </button>
      </div>
    </form>
  </div>
</template>
