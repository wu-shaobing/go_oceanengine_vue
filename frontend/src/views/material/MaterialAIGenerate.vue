<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  prompt: '',
  style: 'realistic',
  ratio: '9:16',
  count: 4
})

const styles = [
  { value: 'realistic', label: '写实风格', icon: '📷' },
  { value: 'cartoon', label: '卡通风格', icon: '🎨' },
  { value: 'minimalist', label: '简约风格', icon: '◻️' },
  { value: '3d', label: '3D风格', icon: '🎲' }
]

const ratios = [
  { value: '9:16', label: '竖版 9:16' },
  { value: '16:9', label: '横版 16:9' },
  { value: '1:1', label: '方形 1:1' }
]

const generatedImages = ref<string[]>([])

const handleGenerate = () => {
  if (!formData.value.prompt) {
    alert('请输入描述内容')
    return
  }
  alert('开始生成素材')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: 'AI生成' }]" />
      <h1 class="text-3xl font-bold text-gray-900">AI素材生成</h1>
      <p class="mt-2 text-gray-600">使用AI智能生成广告素材</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">描述你想要的素材</label>
            <textarea v-model="formData.prompt" rows="4" placeholder="例如：一款时尚的智能手表，放在木质桌面上，自然光照射，突出表盘细节..."
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">选择风格</label>
            <div class="grid grid-cols-4 gap-3">
              <div v-for="style in styles" :key="style.value"
                   :class="['p-3 border-2 rounded-lg cursor-pointer text-center transition-all',
                            formData.style === style.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                   @click="formData.style = style.value">
                <span class="text-2xl">{{ style.icon }}</span>
                <p class="text-sm font-medium text-gray-700 mt-1">{{ style.label }}</p>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">尺寸比例</label>
              <select v-model="formData.ratio" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option v-for="ratio in ratios" :key="ratio.value" :value="ratio.value">{{ ratio.label }}</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">生成数量</label>
              <select v-model="formData.count" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">1张</option>
                <option :value="2">2张</option>
                <option :value="4">4张</option>
              </select>
            </div>
          </div>

          <button class="w-full py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-lg hover:from-blue-700 hover:to-purple-700 font-medium" @click="handleGenerate">
            🚀 开始生成
          </button>
        </div>

        <div v-if="generatedImages.length > 0" class="mt-6">
          <h4 class="font-medium text-gray-900 mb-3">生成结果</h4>
          <div class="grid grid-cols-2 gap-4">
            <div v-for="(_img, index) in generatedImages" :key="index"
                 class="aspect-[9/16] bg-gray-100 rounded-lg"></div>
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">生成额度</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">今日已用</span>
              <span class="text-gray-900">15 / 50</span>
            </div>
            <div class="w-full h-2 bg-gray-200 rounded-full">
              <div class="h-full bg-blue-500 rounded-full" style="width: 30%"></div>
            </div>
          </div>
        </div>

        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">使用技巧</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 描述越详细，效果越好</li>
            <li>• 包含产品特点和场景</li>
            <li>• 说明光线和色调偏好</li>
            <li>• 可参考竞品素材风格</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
