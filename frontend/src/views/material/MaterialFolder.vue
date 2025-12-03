<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const folders = ref([
  { id: 'F001', name: '默认文件夹', itemCount: 256, size: '1.2GB', createdAt: '2025-01-01' },
  { id: 'F002', name: '促销素材', itemCount: 89, size: '450MB', createdAt: '2025-10-15' },
  { id: 'F003', name: '品牌素材', itemCount: 45, size: '280MB', createdAt: '2025-09-20' },
  { id: 'F004', name: '产品图片', itemCount: 156, size: '680MB', createdAt: '2025-08-10' },
  { id: 'F005', name: '视频素材', itemCount: 34, size: '2.1GB', createdAt: '2025-11-01' }
])

const showCreateModal = ref(false)
const newFolderName = ref('')
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '文件夹管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">文件夹管理</h1>
          <p class="mt-2 text-gray-600">管理素材库文件夹结构</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
                @click="showCreateModal = true">
          新建文件夹
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">文件夹总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ folders.length }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">素材总数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">580</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已用空间</p>
        <p class="text-2xl font-bold text-green-600 mt-1">4.7GB</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">剩余空间</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">15.3GB</p>
      </div>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <div v-for="folder in folders" :key="folder.id"
           class="bg-white rounded-lg border border-gray-200 p-4 hover:shadow-md transition-shadow cursor-pointer">
        <div class="flex items-start justify-between">
          <span class="text-4xl">📁</span>
          <button class="text-gray-400 hover:text-gray-600">⋮</button>
        </div>
        <h4 class="font-medium text-gray-900 mt-3">{{ folder.name }}</h4>
        <div class="flex items-center justify-between mt-2 text-sm text-gray-500">
          <span>{{ folder.itemCount }} 项</span>
          <span>{{ folder.size }}</span>
        </div>
      </div>

      <div class="bg-gray-50 rounded-lg border-2 border-dashed border-gray-300 p-4 flex flex-col items-center justify-center cursor-pointer hover:border-blue-400 transition-colors"
           @click="showCreateModal = true">
        <span class="text-3xl text-gray-400">➕</span>
        <span class="text-sm text-gray-500 mt-2">新建文件夹</span>
      </div>
    </div>

    <div v-if="showCreateModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-96">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">新建文件夹</h3>
        <input v-model="newFolderName" type="text" placeholder="请输入文件夹名称"
               class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <div class="flex justify-end gap-3 mt-4">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
                  @click="showCreateModal = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建</button>
        </div>
      </div>
    </div>
  </div>
</template>
