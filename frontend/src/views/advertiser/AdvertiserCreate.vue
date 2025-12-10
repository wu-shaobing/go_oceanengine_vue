<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const submitting = ref(false)

const form = reactive({
  name: '',
  company: '',
  industry: '',
  email: '',
  phone: '',
  address: '',
  description: ''
})

const industries = [
  '电商零售',
  '游戏娱乐',
  '金融服务',
  '教育培训',
  '医疗健康',
  '房产家居',
  '汽车交通',
  '旅游出行',
  '本地生活',
  '其他'
]

const handleSubmit = async () => {
  if (!form.name) {
    alert('请输入广告主名称')
    return
  }
  if (!form.company) {
    alert('请输入公司名称')
    return
  }
  
  submitting.value = true
  await new Promise(resolve => setTimeout(resolve, 800))
  submitting.value = false
  
  alert('创建成功')
  router.push('/advertisers')
}

const handleCancel = () => {
  router.back()
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '广告主管理', path: '/advertisers' },
        { name: '创建广告主' }
      ]" />
      <h1 class="text-3xl font-bold text-gray-900">创建广告主</h1>
      <p class="mt-2 text-gray-600">添加新的广告主账户</p>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">广告主名称 *</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入广告主名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">公司名称 *</label>
            <input
              v-model="form.company"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入公司名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">所属行业</label>
            <select
              v-model="form.industry"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">请选择行业</option>
              <option v-for="ind in industries" :key="ind" :value="ind">{{ ind }}</option>
            </select>
          </div>
        </div>
      </div>

      <!-- 联系信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">联系信息</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
            <input
              v-model="form.email"
              type="email"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="contact@example.com"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">电话</label>
            <input
              v-model="form.phone"
              type="tel"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入联系电话"
            />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-1">地址</label>
            <input
              v-model="form.address"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入公司地址"
            />
          </div>
        </div>
      </div>

      <!-- 其他信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">其他信息</h2>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">备注</label>
          <textarea
            v-model="form.description"
            rows="4"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="请输入备注信息"
          ></textarea>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end gap-4">
        <button
          type="button"
          @click="handleCancel"
          class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors"
        >
          取消
        </button>
        <button
          type="submit"
          :disabled="submitting"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors"
        >
          {{ submitting ? '创建中...' : '创建' }}
        </button>
      </div>
    </form>
  </div>
</template>
