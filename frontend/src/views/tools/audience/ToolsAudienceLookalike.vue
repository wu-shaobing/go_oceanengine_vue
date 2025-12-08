<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  name: '',
  sourceAudience: '',
  expansionScale: 'medium'
})

const sourceAudiences = ref([
  { id: 'AU001', name: '高价值用户包', coverage: 125 },
  { id: 'AU002', name: '历史购买用户', coverage: 85 },
  { id: 'AU003', name: '活跃用户包', coverage: 256 }
])

const scales = [
  { value: 'small', label: '精准', coverage: '约500万', similarity: '高相似度', icon: '🎯' },
  { value: 'medium', label: '平衡', coverage: '约1000万', similarity: '中等相似度', icon: '⚖️' },
  { value: 'large', label: '扩展', coverage: '约2000万', similarity: '较低相似度', icon: '🚀' }
]

const handleCancel = () => {
  alert('取消')
}

const handleCreate = () => {
  alert(`创建Lookalike扩展包: ${formData.value.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: 'Lookalike扩展' }]" />
      <h1 class="text-3xl font-bold text-gray-900">Lookalike人群扩展</h1>
      <p class="mt-2 text-gray-600">基于种子人群寻找相似用户</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">扩展配置</h3>
          <div class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">扩展包名称</label>
              <input v-model="formData.name" type="text" placeholder="请输入扩展包名称"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">选择种子人群</label>
              <select v-model="formData.sourceAudience"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option value="">请选择种子人群包</option>
                <option v-for="audience in sourceAudiences" :key="audience.id" :value="audience.id">
                  {{ audience.name }} ({{ audience.coverage }}万人)
                </option>
              </select>
              <p class="text-xs text-gray-500 mt-1">选择一个现有人群包作为Lookalike的种子人群</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">扩展规模</label>
              <div class="grid grid-cols-3 gap-4">
                <div v-for="scale in scales" :key="scale.value"
                     :class="['p-4 border-2 rounded-lg cursor-pointer text-center transition-all',
                              formData.expansionScale === scale.value ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                     @click="formData.expansionScale = scale.value">
                  <span class="text-3xl">{{ scale.icon }}</span>
                  <h4 class="font-medium text-gray-900 mt-2">{{ scale.label }}</h4>
                  <p class="text-sm text-blue-600 mt-1">{{ scale.coverage }}</p>
                  <p class="text-xs text-gray-500 mt-1">{{ scale.similarity }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

<div class="flex justify-end gap-4">
          <button @click="handleCancel" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button @click="handleCreate" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建Lookalike扩展包</button>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">Lookalike原理</h4>
          <p class="text-sm text-gray-600">
            Lookalike通过机器学习分析种子人群的共同特征，在全平台用户中寻找具有相似特征的潜在用户，帮助扩大目标受众规模。
          </p>
        </div>

        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">使用建议</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 种子人群不少于1000人</li>
            <li>• 种子人群特征越明显效果越好</li>
            <li>• 建议先小规模测试</li>
            <li>• 定期更新种子人群</li>
          </ul>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">预估效果</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">预估覆盖</span>
              <span class="text-gray-900 font-medium">1,000万</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">预估相似度</span>
              <span class="text-green-600 font-medium">75%</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">处理时间</span>
              <span class="text-gray-900 font-medium">约2小时</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
