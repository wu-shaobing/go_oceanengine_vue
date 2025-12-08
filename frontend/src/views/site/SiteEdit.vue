<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const siteData = ref({
  id: 'LP12345',
  name: '双11促销活动页',
  url: 'https://example.com/promo/1111',
  status: 'published',
  template: 'promo',
  components: ['header', 'banner', 'products', 'form', 'footer']
})

const componentList = ref([
  { id: 'header', name: '页面头部', icon: '📋' },
  { id: 'banner', name: '轮播图', icon: '🖼️' },
  { id: 'products', name: '商品列表', icon: '🛒' },
  { id: 'form', name: '表单收集', icon: '📝' },
  { id: 'video', name: '视频模块', icon: '🎬' },
  { id: 'footer', name: '页面底部', icon: '📌' }
])

const handlePreview = () => {
  alert('预览页面')
}

const handleSave = () => {
  alert('保存成功')
}

const handlePublish = () => {
  alert('发布成功')
}

const handleEditComponent = (compId: string) => {
  alert(`编辑组件: ${compId}`)
}

const handleDeleteComponent = (compId: string) => {
  if (confirm('确定删除该组件?')) {
    siteData.value.components = siteData.value.components.filter(c => c !== compId)
  }
}

const handleAddComponent = (comp: typeof componentList.value[0]) => {
  if (!siteData.value.components.includes(comp.id)) {
    siteData.value.components.push(comp.id)
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页管理' }, { name: '编辑落地页' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ siteData.name }}</h1>
          <p class="mt-2 text-gray-600">ID: {{ siteData.id }}</p>
        </div>
        <div class="flex gap-3">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handlePreview">预览</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSave">保存</button>
          <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700" @click="handlePublish">发布</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">组件库</h4>
        <div class="space-y-2">
          <div v-for="comp in componentList" :key="comp.id"
               class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100"
               @click="handleAddComponent(comp)">
            <span class="text-xl">{{ comp.icon }}</span>
            <span class="text-sm text-gray-700">{{ comp.name }}</span>
          </div>
        </div>
      </div>

      <div class="col-span-2 bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">页面编辑区</h4>
        <div class="border-2 border-dashed border-gray-300 rounded-lg min-h-[500px] p-4">
          <div v-for="comp in siteData.components" :key="comp"
               class="mb-3 p-4 bg-gray-50 border border-gray-200 rounded-lg flex items-center justify-between">
            <div class="flex items-center gap-3">
              <span class="text-xl">{{ componentList.find(c => c.id === comp)?.icon }}</span>
              <span class="text-sm font-medium text-gray-700">
                {{ componentList.find(c => c.id === comp)?.name }}
              </span>
            </div>
            <div class="flex gap-2">
              <button class="text-blue-600 text-sm hover:text-blue-800" @click="handleEditComponent(comp)">编辑</button>
              <button class="text-red-600 text-sm hover:text-red-800" @click="handleDeleteComponent(comp)">删除</button>
            </div>
          </div>
          <div class="text-center text-gray-400 py-4">
            拖拽组件到此处
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <h4 class="font-medium text-gray-900 mb-4">属性设置</h4>
        <div class="space-y-4">
          <div>
            <label class="block text-sm text-gray-600 mb-1">页面名称</label>
            <input v-model="siteData.name" type="text"
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
          </div>
          <div>
            <label class="block text-sm text-gray-600 mb-1">页面URL</label>
            <input v-model="siteData.url" type="text" readonly
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm bg-gray-50">
          </div>
          <div>
            <label class="block text-sm text-gray-600 mb-1">模板类型</label>
            <select v-model="siteData.template" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="promo">促销活动</option>
              <option value="product">产品展示</option>
              <option value="form">表单收集</option>
              <option value="brand">品牌展示</option>
            </select>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
