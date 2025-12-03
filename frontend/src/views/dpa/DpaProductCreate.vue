<script setup lang="ts">
import { reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  name: '',
  category: '',
  price: '',
  originalPrice: '',
  stock: '',
  description: '',
  imageUrl: '',
  landingUrl: ''
})

const handleSubmit = () => {
  console.log('Submit product:', form)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: '创建商品' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建DPA商品</h1>
      <p class="mt-2 text-gray-600">添加新的动态商品广告商品</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品名称 *</label>
            <input v-model="form.name" type="text" placeholder="请输入商品名称"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品分类 *</label>
            <select v-model="form.category" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="">请选择分类</option>
              <option value="phone">手机数码</option>
              <option value="computer">电脑办公</option>
              <option value="home">家居生活</option>
              <option value="fashion">服饰鞋包</option>
            </select>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">现价 (元) *</label>
              <input v-model="form.price" type="number" step="0.01" placeholder="0.00"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">原价 (元)</label>
              <input v-model="form.originalPrice" type="number" step="0.01" placeholder="0.00"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">库存数量 *</label>
            <input v-model="form.stock" type="number" placeholder="请输入库存数量"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品描述</label>
            <textarea v-model="form.description" rows="4" placeholder="请输入商品描述"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品图片URL</label>
            <input v-model="form.imageUrl" type="url" placeholder="https://"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">落地页URL *</label>
            <input v-model="form.landingUrl" type="url" placeholder="https://"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div class="flex gap-4 pt-4">
            <button type="submit" class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
              创建商品
            </button>
            <router-link to="/dpa/products" class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 text-center">
              取消
            </router-link>
          </div>
        </form>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-medium text-gray-900 mb-3">商品预览</h3>
          <div class="border border-gray-200 rounded-lg overflow-hidden">
            <div class="h-40 bg-gray-100 flex items-center justify-center">
              <span class="text-gray-400">商品图片</span>
            </div>
            <div class="p-3">
              <p class="font-medium text-gray-900 text-sm">{{ form.name || '商品名称' }}</p>
              <p class="text-red-600 font-bold mt-1">¥{{ form.price || '0.00' }}</p>
            </div>
          </div>
        </div>
        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">创建须知</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 商品名称需与落地页一致</li>
            <li>• 图片建议尺寸 800x800</li>
            <li>• 价格需与实际销售价格一致</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
