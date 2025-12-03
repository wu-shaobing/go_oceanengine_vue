<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const formData = ref({
  title: '',
  tags: '',
  folder: 'default',
  autoGenerate: false
})

const uploadList = ref<{ name: string; progress: number; status: string }[]>([])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '上传视频' }]" />
      <h1 class="text-3xl font-bold text-gray-900">上传视频素材</h1>
      <p class="mt-2 text-gray-600">上传视频素材到素材库</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <div class="border-2 border-dashed border-gray-300 rounded-lg p-12 text-center hover:border-blue-400 transition-colors cursor-pointer">
          <div class="text-5xl mb-4">🎬</div>
          <p class="text-lg font-medium text-gray-700">拖拽视频到此处上传</p>
          <p class="text-sm text-gray-500 mt-2">或点击选择文件</p>
          <p class="text-xs text-gray-400 mt-4">支持 MP4、MOV 格式，单文件最大 500MB</p>
          <button class="mt-6 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            选择视频
          </button>
        </div>

        <div v-if="uploadList.length > 0" class="mt-6 space-y-3">
          <h4 class="font-medium text-gray-900">上传队列</h4>
          <div v-for="item in uploadList" :key="item.name"
               class="flex items-center gap-4 p-3 bg-gray-50 rounded-lg">
            <span class="text-2xl">🎥</span>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-900">{{ item.name }}</p>
              <div class="w-full h-1.5 bg-gray-200 rounded-full mt-2">
                <div class="h-full bg-blue-500 rounded-full transition-all" :style="{ width: item.progress + '%' }"></div>
              </div>
            </div>
            <span class="text-sm text-gray-500">{{ item.progress }}%</span>
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
              </select>
            </div>
            <div>
              <label class="flex items-center gap-2 cursor-pointer">
                <input v-model="formData.autoGenerate" type="checkbox" class="rounded text-blue-600">
                <span class="text-sm text-gray-700">自动生成封面</span>
              </label>
            </div>
          </div>
        </div>

        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">上传规范</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 视频时长 5-60秒</li>
            <li>• 分辨率 720P 以上</li>
            <li>• 码率 2Mbps 以上</li>
            <li>• 竖版 9:16 或横版 16:9</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
