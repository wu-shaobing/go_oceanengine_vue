<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  product: '',
  industry: '',
  style: 'promotion',
  keywords: ''
})

const generatedTexts = ref<string[]>([])
const isGenerating = ref(false)

const styles = [
  { value: 'promotion', label: '促销风格' },
  { value: 'brand', label: '品牌风格' },
  { value: 'emotion', label: '情感风格' },
  { value: 'humor', label: '幽默风格' },
  { value: 'story', label: '故事风格' }
]

const handleGenerate = () => {
  isGenerating.value = true
  setTimeout(() => {
    generatedTexts.value = [
      '🔥限时特惠，错过再等一年！智能手表低至5折，健康生活从手腕开始~',
      '⌚️ 年度爆款智能手表，10万+用户的共同选择！今日下单立减300元',
      '🎁 送礼首选！高颜值智能手表，让TA感受你的用心 点击领券享优惠',
      '💪 24小时守护健康，血氧/心率/睡眠全监测，这款手表太值了！',
      '✨ 科技与时尚的完美结合，戴上它你就是人群中最亮的星'
    ]
    isGenerating.value = false
  }, 1500)
}

const handleCopy = (text: string) => {
  navigator.clipboard.writeText(text)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '创意工具' }, { name: '文案生成' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创意文案生成</h1>
      <p class="mt-2 text-gray-600">AI智能生成广告创意文案</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">输入信息</h3>
        <form @submit.prevent="handleGenerate" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">产品/服务名称</label>
            <input v-model="form.product" type="text" placeholder="例如: 智能手表"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">所属行业</label>
            <select v-model="form.industry" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="">请选择行业</option>
              <option value="ecommerce">电商零售</option>
              <option value="education">教育培训</option>
              <option value="finance">金融服务</option>
              <option value="game">游戏娱乐</option>
              <option value="app">应用推广</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">文案风格</label>
            <div class="flex flex-wrap gap-2">
              <button v-for="style in styles" :key="style.value" type="button"
                      :class="['px-4 py-2 rounded-lg text-sm transition-colors',
                        form.style === style.value ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200']"
                      @click="form.style = style.value">
                {{ style.label }}
              </button>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">关键词（选填）</label>
            <input v-model="form.keywords" type="text" placeholder="例如: 健康、时尚、性价比"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <button type="submit" :disabled="isGenerating"
                  class="w-full py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50">
            {{ isGenerating ? '生成中...' : '生成文案' }}
          </button>
        </form>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">生成结果</h3>
        <div v-if="generatedTexts.length === 0" class="py-12 text-center text-gray-500">
          <div class="text-5xl mb-4">✍️</div>
          <p>填写信息后点击生成</p>
        </div>
        <div v-else class="space-y-3">
          <div v-for="(text, idx) in generatedTexts" :key="idx"
               class="p-4 bg-gray-50 rounded-lg border border-gray-200 group hover:border-blue-300">
            <p class="text-gray-700">{{ text }}</p>
            <div class="flex items-center justify-between mt-3">
              <span class="text-xs text-gray-400">{{ text.length }}字</span>
              <button @click="handleCopy(text)" 
                      class="text-blue-600 text-sm hover:text-blue-800 opacity-0 group-hover:opacity-100 transition-opacity">
                复制
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
