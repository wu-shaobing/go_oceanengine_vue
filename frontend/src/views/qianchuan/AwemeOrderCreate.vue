<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '随心推' }, { name: '创建订单' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建随心推订单</h1>
      <p class="text-gray-600 mt-1">快速为视频创建投放任务</p>
    </div>

    <div class="max-w-4xl">
      <!-- 选择视频 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">选择推广视频</h3>
        <div class="border-2 border-dashed border-gray-300 rounded-lg p-8">
          <div v-if="!form.videoId" class="text-center">
            <div class="text-gray-400 mb-4">从抖音账号选择要推广的视频</div>
            <button @click="showVideoModal = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
              选择视频
            </button>
          </div>
          <div v-else class="flex items-center">
            <img src="https://via.placeholder.com/120x160" class="w-24 h-32 rounded object-cover mr-4" alt="">
            <div class="flex-1">
              <div class="font-medium">{{ form.videoTitle }}</div>
              <div class="text-sm text-gray-500 mt-1">发布时间: 2024-03-15</div>
              <div class="text-sm text-gray-500">播放: 12.5w 点赞: 8560</div>
            </div>
            <button @click="form.videoId = ''" class="text-red-500 hover:text-red-700">更换</button>
          </div>
        </div>
      </div>

      <!-- 推广目标 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">推广目标</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div v-for="goal in promotionGoals" :key="goal.value"
            @click="form.goal = goal.value"
            :class="form.goal === goal.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200'"
            class="border-2 rounded-lg p-4 cursor-pointer hover:border-blue-300 transition-colors text-center">
            <div class="text-2xl mb-2">{{ goal.icon }}</div>
            <div class="font-medium">{{ goal.label }}</div>
          </div>
        </div>
      </div>

      <!-- 投放设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">投放设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">投放预算</label>
            <div class="flex gap-3">
              <button v-for="budget in budgetOptions" :key="budget"
                @click="form.budget = budget"
                :class="form.budget === budget ? 'border-blue-500 bg-blue-50 text-blue-600' : 'border-gray-300'"
                class="px-6 py-2 border-2 rounded-lg hover:border-blue-300 transition-colors">
                ¥{{ budget }}
              </button>
              <input type="number" v-model="form.customBudget" 
                class="border border-gray-300 rounded-lg px-3 py-2 w-32" 
                placeholder="自定义"
                @focus="form.budget = 0">
            </div>
            <div class="text-sm text-gray-500 mt-2">最低投放金额100元</div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">投放时长</label>
            <div class="flex gap-3">
              <button v-for="duration in durationOptions" :key="duration.value"
                @click="form.duration = duration.value"
                :class="form.duration === duration.value ? 'border-blue-500 bg-blue-50 text-blue-600' : 'border-gray-300'"
                class="px-6 py-2 border-2 rounded-lg hover:border-blue-300 transition-colors">
                {{ duration.label }}
              </button>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">期望效果</label>
            <div class="flex items-center gap-4">
              <select v-model="form.expectType" class="border border-gray-300 rounded-lg px-3 py-2">
                <option value="speed">优先跑量</option>
                <option value="cost">优先成本</option>
              </select>
              <span class="text-sm text-gray-500">{{ form.expectType === 'speed' ? '快速曝光，优先保量' : '控制成本，优化转化' }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 定向设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">定向设置（可选）</h3>
        <div class="space-y-4">
          <div class="flex items-center space-x-4">
            <label class="flex items-center">
              <input type="radio" v-model="form.targetType" value="auto" class="mr-2"> 系统智能推荐
            </label>
            <label class="flex items-center">
              <input type="radio" v-model="form.targetType" value="custom" class="mr-2"> 自定义定向
            </label>
          </div>
          <div v-if="form.targetType === 'custom'" class="pl-4 space-y-3">
            <div>
              <label class="block text-sm text-gray-700 mb-1">性别</label>
              <div class="flex gap-2">
                <label class="flex items-center"><input type="radio" v-model="form.gender" value="all" class="mr-1"> 不限</label>
                <label class="flex items-center"><input type="radio" v-model="form.gender" value="male" class="mr-1"> 男</label>
                <label class="flex items-center"><input type="radio" v-model="form.gender" value="female" class="mr-1"> 女</label>
              </div>
            </div>
            <div>
              <label class="block text-sm text-gray-700 mb-1">年龄</label>
              <select v-model="form.age" class="border border-gray-300 rounded-lg px-3 py-2">
                <option value="">不限</option>
                <option value="18-23">18-23岁</option>
                <option value="24-30">24-30岁</option>
                <option value="31-40">31-40岁</option>
                <option value="40+">40岁以上</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end space-x-4">
        <router-link to="/qianchuan/aweme-order" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</router-link>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">提交订单 (¥{{ totalBudget }})</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const showVideoModal = ref(false)

const promotionGoals = [
  { value: 'video', label: '视频播放', icon: '▶️' },
  { value: 'interaction', label: '互动提升', icon: '💬' },
  { value: 'fans', label: '涨粉', icon: '👥' },
  { value: 'product', label: '商品购买', icon: '🛒' }
]

const budgetOptions = [100, 200, 300, 500, 1000]

const durationOptions = [
  { value: 6, label: '6小时' },
  { value: 12, label: '12小时' },
  { value: 24, label: '24小时' }
]

const form = ref({
  videoId: 'V123456',
  videoTitle: '新品美妆测评分享 超好用必入!',
  goal: 'video',
  budget: 300,
  customBudget: null as number | null,
  duration: 24,
  expectType: 'speed',
  targetType: 'auto',
  gender: 'all',
  age: ''
})

const totalBudget = computed(() => {
  return form.value.customBudget || form.value.budget
})
</script>
