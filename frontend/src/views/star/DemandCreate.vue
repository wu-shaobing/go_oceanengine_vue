<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '需求管理', path: '/star/demand' }, { name: '发布需求' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">发布需求</h1>
      <p class="text-gray-600 mt-1">创建达人合作需求，吸引优质达人申请</p>
    </div>

    <div class="bg-white rounded-lg shadow p-6">
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- 基本信息 -->
        <div>
          <h3 class="text-lg font-medium text-gray-900 mb-4">基本信息</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-1">需求标题 *</label>
              <input type="text" v-model="form.title" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="输入需求标题，如：618大促品牌种草视频招募" required>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-1">需求描述 *</label>
              <textarea v-model="form.description" rows="4" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="详细描述合作需求、达人要求等" required></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">需求类型 *</label>
              <select v-model="form.type" class="w-full border border-gray-300 rounded-lg px-3 py-2" required>
                <option value="">请选择</option>
                <option value="video">视频推广</option>
                <option value="live">直播带货</option>
                <option value="post">图文种草</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">行业领域</label>
              <select v-model="form.industry" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">请选择</option>
                <option value="beauty">美妆护肤</option>
                <option value="fashion">服饰穿搭</option>
                <option value="food">美食餐饮</option>
                <option value="tech">数码科技</option>
                <option value="life">生活方式</option>
                <option value="mother">母婴育儿</option>
              </select>
            </div>
          </div>
        </div>

        <!-- 招募设置 -->
        <div>
          <h3 class="text-lg font-medium text-gray-900 mb-4">招募设置</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">总预算 (元) *</label>
              <input type="number" v-model="form.budget" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="100000" min="1000" required>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">招募人数 *</label>
              <input type="number" v-model="form.targetCount" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="10" min="1" required>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">单人预算 (元)</label>
              <input type="number" v-model="form.perBudget" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="10000" min="100">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">截止日期 *</label>
              <input type="date" v-model="form.deadline" class="w-full border border-gray-300 rounded-lg px-3 py-2" required>
            </div>
          </div>
        </div>

        <!-- 达人要求 -->
        <div>
          <h3 class="text-lg font-medium text-gray-900 mb-4">达人要求</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">最低粉丝数</label>
              <select v-model="form.minFans" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">不限</option>
                <option value="10000">1万+</option>
                <option value="50000">5万+</option>
                <option value="100000">10万+</option>
                <option value="500000">50万+</option>
                <option value="1000000">100万+</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">达人等级</label>
              <select v-model="form.level" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">不限</option>
                <option value="1">L1</option>
                <option value="2">L2</option>
                <option value="3">L3</option>
                <option value="4">L4</option>
                <option value="5">L5</option>
              </select>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-1">其他要求</label>
              <textarea v-model="form.requirements" rows="3" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="其他具体要求，如：需要有品牌合作经验、需配合直播等"></textarea>
            </div>
          </div>
        </div>

        <!-- 提交按钮 -->
        <div class="flex justify-end space-x-4 pt-4 border-t">
          <button type="button" @click="handleCancel" class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
            取消
          </button>
          <button type="button" @click="handleSaveDraft" class="px-6 py-2 border border-blue-600 text-blue-600 rounded-lg hover:bg-blue-50">
            保存草稿
          </button>
          <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            发布需求
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()

const form = ref({
  title: '',
  description: '',
  type: '',
  industry: '',
  budget: null as number | null,
  targetCount: null as number | null,
  perBudget: null as number | null,
  deadline: '',
  minFans: '',
  level: '',
  requirements: ''
})

const handleSubmit = () => {
  // 表单验证
  if (!form.value.title || !form.value.description || !form.value.type || !form.value.budget || !form.value.targetCount || !form.value.deadline) {
    alert('请填写必填项')
    return
  }
  
  // 提交逻辑
  alert('需求发布成功')
  router.push('/star/demand')
}

const handleSaveDraft = () => {
  alert('草稿已保存')
}

const handleCancel = () => {
  router.push('/star/demand')
}
</script>
