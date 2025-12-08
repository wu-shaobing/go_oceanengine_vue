<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '创建广告' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建广告</h1>
      <p class="text-gray-600 mt-1">设置广告定向和创意</p>
    </div>

    <div class="max-w-4xl">
      <!-- 广告计划选择 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">选择广告计划</h3>
        <select v-model="form.campaignId" class="w-full border border-gray-300 rounded-lg px-3 py-2">
          <option value="">请选择广告计划</option>
          <option value="10001">618大促直播计划</option>
          <option value="10002">新品短视频推广</option>
          <option value="10003">爆款商品推广</option>
        </select>
      </div>

      <!-- 定向设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">定向设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">地域</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.region" value="all" class="mr-2"> 不限
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.region" value="custom" class="mr-2"> 指定地域
              </label>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">年龄</label>
            <div class="flex flex-wrap gap-2">
              <label v-for="age in ageOptions" :key="age.value" class="flex items-center">
                <input type="checkbox" v-model="form.ages" :value="age.value" class="mr-1 rounded"> {{ age.label }}
              </label>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">性别</label>
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
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">兴趣行为</label>
            <button @click="addInterestTag" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">+ 添加兴趣标签</button>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">人群包</label>
            <select v-model="form.audiencePackage" class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option value="">不使用人群包</option>
              <option value="1">高价值用户</option>
              <option value="2">近期活跃买家</option>
              <option value="3">店铺粉丝</option>
            </select>
          </div>
        </div>
      </div>

      <!-- 出价设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">出价设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">转化出价</label>
            <div class="flex items-center">
              <input type="number" v-model="form.bid" class="border border-gray-300 rounded-lg px-3 py-2 w-32" placeholder="出价">
              <span class="ml-2 text-gray-500">元/转化</span>
            </div>
            <div class="text-sm text-gray-500 mt-1">建议出价范围: ¥15 - ¥50</div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">ROI目标</label>
            <div class="flex items-center">
              <input type="number" v-model="form.roiGoal" step="0.1" class="border border-gray-300 rounded-lg px-3 py-2 w-32" placeholder="ROI">
              <span class="ml-2 text-gray-500">预期ROI</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 创意设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">创意设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">广告名称 <span class="text-red-500">*</span></label>
            <input type="text" v-model="form.name" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="请输入广告名称">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">创意形式</label>
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
            <label class="block text-sm font-medium text-gray-700 mb-1">上传素材</label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
              <div class="text-gray-400 mb-2">点击或拖拽上传</div>
              <button @click="selectFile" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">选择文件</button>
              <div class="text-sm text-gray-500 mt-2">支持 MP4、MOV 格式，建议9:16竖版视频</div>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">广告文案</label>
            <textarea v-model="form.copywriting" rows="3" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="请输入广告文案"></textarea>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end space-x-4">
        <router-link to="/qianchuan/ad" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</router-link>
        <button @click="saveDraft" :disabled="loading" class="px-6 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50">{{ loading ? '保存中...' : '保存草稿' }}</button>
        <button @click="submitAd" :disabled="loading" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50">{{ loading ? '提交中...' : '提交审核' }}</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const loading = ref(false)

const ageOptions = [
  { value: '18-23', label: '18-23岁' },
  { value: '24-30', label: '24-30岁' },
  { value: '31-40', label: '31-40岁' },
  { value: '41-49', label: '41-49岁' },
  { value: '50+', label: '50岁以上' }
]

const form = ref({
  campaignId: '',
  name: '',
  region: 'all',
  ages: [] as string[],
  gender: 'all',
  audiencePackage: '',
  bid: 25,
  roiGoal: 3,
  creativeType: 'video',
  copywriting: ''
})

const validateForm = () => {
  if (!form.value.campaignId) {
    alert('请选择广告计划')
    return false
  }
  if (!form.value.name.trim()) {
    alert('请输入广告名称')
    return false
  }
  if (!form.value.bid || form.value.bid < 1) {
    alert('请设置有效的出价')
    return false
  }
  return true
}

const saveDraft = async () => {
  if (!form.value.name.trim()) {
    alert('请先输入广告名称')
    return
  }
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert('草稿保存成功')
  } finally {
    loading.value = false
  }
}

const submitAd = async () => {
  if (!validateForm()) return
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert('广告提交成功，等待审核')
    router.push('/qianchuan/ad')
  } finally {
    loading.value = false
  }
}

const addInterestTag = () => {
  alert('兴趣标签选择功能开发中...')
}

const selectFile = () => {
  alert('文件上传功能开发中...')
}
</script>
