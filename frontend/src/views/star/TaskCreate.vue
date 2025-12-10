<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()
const submitting = ref(false)

const form = reactive({
  name: '',
  type: 'video',
  description: '',
  budget: 0,
  influencerCount: 1,
  startDate: '',
  endDate: '',
  requirements: '',
  targetAudience: ''
})

const taskTypes = [
  { value: 'video', label: '视频推广', desc: '达人发布短视频进行推广' },
  { value: 'live', label: '直播推广', desc: '达人在直播间进行带货推广' },
  { value: 'article', label: '图文推广', desc: '达人发布图文内容推广' }
]

const handleSubmit = async () => {
  if (!form.name) {
    alert('请输入任务名称')
    return
  }
  if (!form.budget || form.budget <= 0) {
    alert('请输入有效预算')
    return
  }
  if (!form.startDate || !form.endDate) {
    alert('请选择任务时间')
    return
  }
  
  submitting.value = true
  await new Promise(resolve => setTimeout(resolve, 800))
  submitting.value = false
  
  alert('任务创建成功')
  router.push('/star/tasks')
}

const handleCancel = () => {
  router.back()
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[
        { name: '星图', path: '/star' },
        { name: '任务列表', path: '/star/tasks' },
        { name: '创建任务' }
      ]" />
      <h1 class="text-3xl font-bold text-gray-900">创建任务</h1>
      <p class="mt-2 text-gray-600">创建达人合作推广任务</p>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">任务名称 *</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="例如：618美妆种草推广"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">任务类型</label>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div
                v-for="type in taskTypes"
                :key="type.value"
                @click="form.type = type.value"
                :class="[
                  'p-4 border-2 rounded-lg cursor-pointer transition-all',
                  form.type === type.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
                ]"
              >
                <h4 class="font-medium text-gray-900">{{ type.label }}</h4>
                <p class="text-sm text-gray-500 mt-1">{{ type.desc }}</p>
              </div>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">任务描述</label>
            <textarea
              v-model="form.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请描述任务内容和推广要求"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- 预算和时间 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">预算与时间</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">任务预算 *</label>
            <div class="flex">
              <span class="inline-flex items-center px-3 border border-r-0 border-gray-300 rounded-l-lg bg-gray-50 text-gray-500">¥</span>
              <input
                v-model.number="form.budget"
                type="number"
                min="0"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-r-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="请输入预算金额"
              />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">达人数量</label>
            <input
              v-model.number="form.influencerCount"
              type="number"
              min="1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">开始日期 *</label>
            <input
              v-model="form.startDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">结束日期 *</label>
            <input
              v-model="form.endDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- 投放要求 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">投放要求</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">内容要求</label>
            <textarea
              v-model="form.requirements"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请描述对创作内容的要求，如时长、风格、关键信息点等"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">目标受众</label>
            <textarea
              v-model="form.targetAudience"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请描述目标受众特征，如年龄、性别、兴趣等"
            ></textarea>
          </div>
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
          {{ submitting ? '创建中...' : '创建任务' }}
        </button>
      </div>
    </form>
  </div>
</template>
