<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  title: '',
  tags: '',
  folder: 'default'
})

const previewImages = ref<string[]>([])

const handleSelectImages = () => {
  alert('选择图片')
}

const handleStartUpload = () => {
  alert('开始上传')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '上传图片' }]" />
      <h1 class="text-3xl font-bold text-gray-900">上传图片素材</h1>
      <p class="mt-2 text-gray-600">批量上传图片素材到素材库</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
<div class="border-2 border-dashed border-gray-300 rounded-lg p-12 text-center hover:border-blue-400 transition-colors cursor-pointer" @click="handleSelectImages">
          <div class="text-5xl mb-4">🖼️</div>
          <p class="text-lg font-medium text-gray-700">拖拽图片到此处上传</p>
          <p class="text-sm text-gray-500 mt-2">或点击选择文件，支持批量上传</p>
          <p class="text-xs text-gray-400 mt-4">支持 JPG、PNG、GIF 格式，单张最大 10MB</p>
          <button class="mt-6 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSelectImages">
            选择图片
          </button>
        </div>

        <div v-if="previewImages.length > 0" class="mt-6">
          <h4 class="font-medium text-gray-900 mb-3">预览</h4>
          <div class="grid grid-cols-4 gap-4">
            <div v-for="(img, index) in previewImages" :key="index"
                 class="aspect-square bg-gray-100 rounded-lg relative group">
              <img :src="img" class="w-full h-full object-cover rounded-lg">
              <button @click.stop="previewImages.splice(index, 1)" class="absolute top-2 right-2 w-6 h-6 bg-red-500 text-white rounded-full opacity-0 group-hover:opacity-100 transition-opacity">
                ×
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-4">素材信息</h4>
          <div class="space-y-4">
            <div>
              <label class="block text-sm text-gray-600 mb-1">素材标题</label>
              <input v-model="formData.title" type="text" placeholder="可选"
                     class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            </div>
            <div>
              <label class="block text-sm text-gray-600 mb-1">标签</label>
              <input v-model="formData.tags" type="text" placeholder="多个标签用逗号分隔"
                     class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            </div>
            <div>
              <label class="block text-sm text-gray-600 mb-1">存储文件夹</label>
              <select v-model="formData.folder" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                <option value="default">默认文件夹</option>
                <option value="promo">促销素材</option>
                <option value="brand">品牌素材</option>
                <option value="product">产品素材</option>
              </select>
            </div>
          </div>
        </div>

        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">图片规范</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 横版: 1280×720 / 1200×628</li>
            <li>• 竖版: 720×1280 / 628×1200</li>
            <li>• 方形: 800×800</li>
            <li>• 文件大小不超过 10MB</li>
          </ul>
        </div>

        <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleStartUpload">
          开始上传
        </button>
      </div>
    </div>
  </div>
</template>
