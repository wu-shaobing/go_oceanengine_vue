<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '创建计划' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建广告计划</h1>
      <p class="text-gray-600 mt-1">配置千川广告投放计划</p>
    </div>

    <div class="max-w-4xl">
      <!-- 营销目标 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">营销目标</h3>
        <div class="grid grid-cols-3 gap-4">
          <div v-for="goal in marketingGoals" :key="goal.value" 
            @click="form.marketingGoal = goal.value"
            :class="form.marketingGoal === goal.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200'"
            class="border-2 rounded-lg p-4 cursor-pointer hover:border-blue-300 transition-colors">
            <div class="text-2xl mb-2">{{ goal.icon }}</div>
            <div class="font-medium">{{ goal.label }}</div>
            <div class="text-sm text-gray-500">{{ goal.desc }}</div>
          </div>
        </div>
      </div>

      <!-- 基本信息 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">基本信息</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">计划名称 <span class="text-red-500">*</span></label>
            <input type="text" v-model="form.name" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="请输入计划名称">
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">关联店铺 <span class="text-red-500">*</span></label>
              <select v-model="form.shopId" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">请选择店铺</option>
                <option value="7001">品牌官方旗舰店</option>
                <option value="7002">美妆专营店</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">推广类型</label>
              <select v-model="form.promotionType" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="live">直播间</option>
                <option value="video">短视频</option>
                <option value="product">商品</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 投放设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">投放设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">投放时间</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.scheduleType" value="unlimited" class="mr-2"> 不限
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.scheduleType" value="scheduled" class="mr-2"> 指定时间
              </label>
            </div>
            <div v-if="form.scheduleType === 'scheduled'" class="mt-2 flex items-center space-x-2">
              <input type="date" v-model="form.startDate" class="border border-gray-300 rounded-lg px-3 py-2">
              <span>至</span>
              <input type="date" v-model="form.endDate" class="border border-gray-300 rounded-lg px-3 py-2">
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.budgetType" value="unlimited" class="mr-2"> 不限
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.budgetType" value="limited" class="mr-2"> 指定预算
              </label>
            </div>
            <div v-if="form.budgetType === 'limited'" class="mt-2">
              <input type="number" v-model="form.budget" class="border border-gray-300 rounded-lg px-3 py-2 w-48" placeholder="请输入日预算">
              <span class="text-sm text-gray-500 ml-2">元/天，最低100元</span>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">出价方式</label>
            <select v-model="form.bidType" class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option value="ocpm">OCPM（目标转化出价）</option>
              <option value="cpm">CPM（千次曝光出价）</option>
              <option value="cpc">CPC（点击出价）</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">转化目标</label>
            <select v-model="form.convertType" class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option value="pay">成交</option>
              <option value="click">商品点击</option>
              <option value="form">表单提交</option>
              <option value="live_enter">直播间进入</option>
            </select>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end space-x-4">
        <router-link to="/qianchuan/campaign" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</router-link>
        <button 
          @click="saveDraft" 
          :disabled="loading"
          class="px-6 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50"
        >
          {{ loading ? '保存中...' : '保存草稿' }}
        </button>
        <button 
          @click="handleNext" 
          :disabled="loading"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
        >
          {{ loading ? '提交中...' : '提交创建' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()

const marketingGoals = [
  { value: 'live', label: '直播带货', desc: '为直播间引流，提升成交', icon: '📺' },
  { value: 'video', label: '短视频带货', desc: '通过短视频推广商品', icon: '🎬' },
  { value: 'product', label: '商品推广', desc: '直接推广商品详情页', icon: '🛒' }
]

const form = ref({
  marketingGoal: 'live',
  name: '',
  shopId: '',
  promotionType: 'live',
  scheduleType: 'unlimited',
  startDate: '',
  endDate: '',
  budgetType: 'limited',
  budget: 1000,
  bidType: 'ocpm',
  convertType: 'pay'
})

const loading = ref(false)
const errors = ref<Record<string, string>>({})

// 表单验证
const validateForm = () => {
  errors.value = {}
  
  if (!form.value.name.trim()) {
    errors.value.name = '请输入计划名称'
    return false
  }
  
  if (!form.value.shopId) {
    errors.value.shopId = '请选择关联店铺'
    return false
  }
  
  if (form.value.budgetType === 'limited' && (!form.value.budget || form.value.budget < 100)) {
    errors.value.budget = '日预算最低100元'
    return false
  }
  
  return true
}

// 保存草稿
const saveDraft = async () => {
  if (!form.value.name.trim()) {
    alert('请先输入计划名称')
    return
  }
  
  loading.value = true
  try {
    // TODO: 调用API保存草稿
    await new Promise(resolve => setTimeout(resolve, 500))
    alert('草稿保存成功')
  } catch (error) {
    alert('保存失败，请重试')
  } finally {
    loading.value = false
  }
}

// 下一步
const handleNext = async () => {
  if (!validateForm()) {
    // 显示验证错误
    const firstError = Object.values(errors.value)[0]
    if (firstError) {
      alert(firstError)
    }
    return
  }
  
  loading.value = true
  try {
    // TODO: 调用API创建计划
    await new Promise(resolve => setTimeout(resolve, 500))
    alert('广告计划创建成功！')
    router.push('/qianchuan/campaign')
  } catch (error) {
    alert('创建失败，请重试')
  } finally {
    loading.value = false
  }
}
</script>
