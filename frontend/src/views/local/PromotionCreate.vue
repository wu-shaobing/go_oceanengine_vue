<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '推广管理', path: '/local/promotion' }, { name: '创建推广' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建推广</h1>
      <p class="text-gray-600 mt-1">新建本地推广计划</p>
    </div>

    <div class="max-w-3xl">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">基本信息</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">推广名称 <span class="text-red-500">*</span></label>
            <input type="text" v-model="form.name" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入推广名称">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">所属项目 <span class="text-red-500">*</span></label>
            <select v-model="form.projectId" class="w-full border border-gray-300 rounded px-3 py-2">
              <option value="">请选择项目</option>
              <option v-for="project in projects" :key="project.id" :value="project.id">{{ project.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">推广目标 <span class="text-red-500">*</span></label>
            <div class="grid grid-cols-3 gap-4">
              <label v-for="goal in promotionGoals" :key="goal.value" class="border rounded-lg p-4 cursor-pointer hover:border-blue-500 transition-colors" :class="{ 'border-blue-500 bg-blue-50': form.goal === goal.value }">
                <input type="radio" v-model="form.goal" :value="goal.value" class="hidden">
                <div class="text-center">
                  <span class="text-2xl">{{ goal.icon }}</span>
                  <div class="font-medium mt-2">{{ goal.label }}</div>
                </div>
              </label>
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
                <input type="radio" v-model="form.scheduleType" value="continuous" class="mr-2"> 长期投放
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.scheduleType" value="fixed" class="mr-2"> 指定日期
              </label>
            </div>
            <div v-if="form.scheduleType === 'fixed'" class="mt-3 flex items-center space-x-2">
              <input type="date" v-model="form.startDate" class="border border-gray-300 rounded px-3 py-2">
              <span>至</span>
              <input type="date" v-model="form.endDate" class="border border-gray-300 rounded px-3 py-2">
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算</label>
            <div class="flex items-center">
              <span class="mr-2">¥</span>
              <input type="number" v-model="form.dailyBudget" class="w-48 border border-gray-300 rounded px-3 py-2" min="100">
              <span class="ml-2 text-gray-500">元/天（最低100元）</span>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">出价方式</label>
            <select v-model="form.bidType" class="w-full border border-gray-300 rounded px-3 py-2">
              <option value="ocpm">OCPM智能出价</option>
              <option value="cpc">CPC点击出价</option>
              <option value="cpm">CPM展示出价</option>
            </select>
          </div>
          <div v-if="form.bidType !== 'ocpm'">
            <label class="block text-sm font-medium text-gray-700 mb-1">出价金额</label>
            <div class="flex items-center">
              <span class="mr-2">¥</span>
              <input type="number" v-model="form.bid" class="w-48 border border-gray-300 rounded px-3 py-2" step="0.01">
              <span class="ml-2 text-gray-500">元</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 定向设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">定向设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">投放范围</label>
            <select v-model="form.locationRadius" class="w-full border border-gray-300 rounded px-3 py-2">
              <option value="3">门店周边3公里</option>
              <option value="5">门店周边5公里</option>
              <option value="10">门店周边10公里</option>
              <option value="20">门店周边20公里</option>
              <option value="city">全城投放</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">年龄定向</label>
            <div class="flex items-center space-x-2">
              <select v-model="form.ageMin" class="border border-gray-300 rounded px-3 py-2">
                <option v-for="age in ageOptions" :key="age" :value="age">{{ age }}岁</option>
              </select>
              <span>-</span>
              <select v-model="form.ageMax" class="border border-gray-300 rounded px-3 py-2">
                <option v-for="age in ageOptions" :key="age" :value="age">{{ age }}岁</option>
              </select>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">性别定向</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.gender" value="all" class="mr-2"> 不限
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.gender" value="male" class="mr-2"> 男
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.gender" value="female" class="mr-2"> 女
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- 创意设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">创意设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">创意类型</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.creativeType" value="video" class="mr-2"> 视频
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.creativeType" value="image" class="mr-2"> 图片
              </label>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">创意素材</label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
              <span class="text-4xl">📤</span>
              <div class="mt-2 text-gray-600">点击或拖拽上传素材</div>
              <div class="text-sm text-gray-400 mt-1">支持mp4、jpg、png格式</div>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">推广标题</label>
            <input type="text" v-model="form.title" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入推广标题" maxlength="30">
            <div class="text-sm text-gray-400 mt-1">{{ form.title.length }}/30</div>
          </div>
        </div>
      </div>

      <div class="flex justify-end space-x-4">
        <router-link to="/local/promotion" class="px-6 py-2 border border-gray-300 rounded hover:bg-gray-50">取消</router-link>
        <button @click="saveDraft" class="px-6 py-2 border border-blue-600 text-blue-600 rounded hover:bg-blue-50">保存草稿</button>
        <button @click="submit" class="px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">提交审核</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = ref({
  name: '',
  projectId: '',
  goal: 'clue',
  scheduleType: 'continuous',
  startDate: '',
  endDate: '',
  dailyBudget: 500,
  bidType: 'ocpm',
  bid: 0,
  locationRadius: '5',
  ageMin: 18,
  ageMax: 55,
  gender: 'all',
  creativeType: 'video',
  title: ''
})

const projects = ref([
  { id: '1', name: '北京朝阳店' },
  { id: '2', name: '上海静安店' },
  { id: '3', name: '广州天河店' }
])

const promotionGoals = [
  { value: 'clue', label: '获取线索', icon: '📋' },
  { value: 'visit', label: '到店访问', icon: '🏪' },
  { value: 'awareness', label: '品牌曝光', icon: '📢' }
]

const ageOptions = Array.from({ length: 50 }, (_, i) => i + 16)

const saveDraft = () => {
  console.log('保存草稿', form.value)
}

const submit = () => {
  console.log('提交审核', form.value)
}
</script>
