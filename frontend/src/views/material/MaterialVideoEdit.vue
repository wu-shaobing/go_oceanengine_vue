<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const videoInfo = ref({
  name: '产品宣传视频.mp4',
  duration: '00:30',
  resolution: '1920x1080',
  size: '45MB'
})

const editOptions = ref({
  trimStart: 0,
  trimEnd: 30,
  addText: false,
  textContent: '',
  addMusic: false,
  musicVolume: 50
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '视频编辑' }]" />
      <h1 class="text-3xl font-bold text-gray-900">视频编辑</h1>
      <p class="mt-2 text-gray-600">编辑和处理视频素材</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-6">
        <div class="bg-black rounded-lg aspect-video flex items-center justify-center">
          <div class="text-center text-white">
            <span class="text-5xl">▶️</span>
            <p class="mt-4">视频预览区域</p>
            <p class="text-sm text-gray-400 mt-2">{{ videoInfo.name }}</p>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-4">时间轴</h4>
          <div class="h-16 bg-gray-100 rounded-lg flex items-center px-4">
            <div class="flex-1 h-8 bg-blue-200 rounded relative">
              <div class="absolute left-0 top-0 bottom-0 w-1 bg-blue-600 rounded cursor-ew-resize"></div>
              <div class="absolute right-0 top-0 bottom-0 w-1 bg-blue-600 rounded cursor-ew-resize"></div>
            </div>
          </div>
          <div class="flex justify-between text-xs text-gray-500 mt-2">
            <span>00:00</span>
            <span>00:15</span>
            <span>00:30</span>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="font-medium text-gray-900 mb-4">裁剪设置</h4>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm text-gray-600 mb-1">开始时间</label>
              <input v-model="editOptions.trimStart" type="number" min="0" max="30"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm text-gray-600 mb-1">结束时间</label>
              <input v-model="editOptions.trimEnd" type="number" min="0" max="30"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg">
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
          <h4 class="font-medium text-gray-900 mb-3">视频信息</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-500">文件名</span>
              <span class="text-gray-900">{{ videoInfo.name }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">时长</span>
              <span class="text-gray-900">{{ videoInfo.duration }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">分辨率</span>
              <span class="text-gray-900">{{ videoInfo.resolution }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">文件大小</span>
              <span class="text-gray-900">{{ videoInfo.size }}</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">添加元素</h4>
          <div class="space-y-3">
            <label class="flex items-center justify-between p-3 bg-gray-50 rounded-lg cursor-pointer">
              <span class="text-sm text-gray-700">添加文字</span>
              <input v-model="editOptions.addText" type="checkbox" class="rounded text-blue-600">
            </label>
            <div v-if="editOptions.addText" class="pl-3">
              <input v-model="editOptions.textContent" type="text" placeholder="输入文字内容"
                     class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            </div>
            <label class="flex items-center justify-between p-3 bg-gray-50 rounded-lg cursor-pointer">
              <span class="text-sm text-gray-700">添加背景音乐</span>
              <input v-model="editOptions.addMusic" type="checkbox" class="rounded text-blue-600">
            </label>
            <div v-if="editOptions.addMusic" class="pl-3">
              <label class="block text-xs text-gray-500 mb-1">音量 {{ editOptions.musicVolume }}%</label>
              <input v-model="editOptions.musicVolume" type="range" min="0" max="100" class="w-full">
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">快捷操作</h4>
          <div class="grid grid-cols-2 gap-2">
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">旋转</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">裁剪</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">调速</button>
            <button class="p-2 text-sm text-gray-700 bg-gray-50 rounded hover:bg-gray-100">滤镜</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
