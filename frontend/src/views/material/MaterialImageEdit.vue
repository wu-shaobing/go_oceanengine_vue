<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const imageInfo = ref({
  name: '产品图片.jpg',
  resolution: '1200x800',
  size: '2.5MB',
  format: 'JPEG'
})

const editOptions = ref({
  brightness: 50,
  contrast: 50,
  saturation: 50,
  cropRatio: '16:9'
})

const cropRatios = ['自由', '1:1', '4:3', '16:9', '9:16']
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '图片编辑' }]" />
      <h1 class="text-3xl font-bold text-gray-900">图片编辑</h1>
      <p class="mt-2 text-gray-600">编辑和处理图片素材</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-6">
        <div class="bg-gray-100 rounded-lg aspect-video flex items-center justify-center">
          <div class="text-center text-gray-500">
            <span class="text-5xl">🖼️</span>
            <p class="mt-4">图片预览区域</p>
            <p class="text-sm mt-2">{{ imageInfo.name }}</p>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="font-medium text-gray-900 mb-4">图像调整</h4>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">亮度</span>
                <span class="text-gray-900">{{ editOptions.brightness }}</span>
              </div>
              <input v-model="editOptions.brightness" type="range" min="0" max="100" class="w-full">
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">对比度</span>
                <span class="text-gray-900">{{ editOptions.contrast }}</span>
              </div>
              <input v-model="editOptions.contrast" type="range" min="0" max="100" class="w-full">
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">饱和度</span>
                <span class="text-gray-900">{{ editOptions.saturation }}</span>
              </div>
              <input v-model="editOptions.saturation" type="range" min="0" max="100" class="w-full">
            </div>
          </div>
        </div>

        <div class="flex justify-end gap-4">
          <button class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">保存编辑</button>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">图片信息</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-500">文件名</span>
              <span class="text-gray-900">{{ imageInfo.name }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">分辨率</span>
              <span class="text-gray-900">{{ imageInfo.resolution }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">文件大小</span>
              <span class="text-gray-900">{{ imageInfo.size }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">格式</span>
              <span class="text-gray-900">{{ imageInfo.format }}</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">裁剪比例</h4>
          <div class="flex flex-wrap gap-2">
            <button v-for="ratio in cropRatios" :key="ratio"
                    :class="['px-3 py-1 text-sm rounded border',
                             editOptions.cropRatio === ratio ? 'border-blue-500 bg-blue-50 text-blue-700' : 'border-gray-300 text-gray-700 hover:bg-gray-50']"
                    @click="editOptions.cropRatio = ratio">
              {{ ratio }}
            </button>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">快捷操作</h4>
          <div class="grid grid-cols-2 gap-2">
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">旋转 90°</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">水平翻转</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">垂直翻转</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">重置</button>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">滤镜</h4>
          <div class="grid grid-cols-3 gap-2">
            <div v-for="i in 6" :key="i" class="aspect-square bg-gray-200 rounded cursor-pointer hover:ring-2 ring-blue-500">
              <div class="h-full flex items-center justify-center text-xs text-gray-500">滤镜{{ i }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
