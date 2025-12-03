<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const router = useRouter()

const form = reactive({
  id: route.params.id || 'DPA001',
  name: '智能手表Pro Max',
  category: '3C数码',
  price: 1299,
  originalPrice: 1599,
  stock: 1520,
  description: '全新一代智能手表',
  status: 'active'
})

const handleSubmit = () => {
  console.log('Update product:', form)
  router.push('/dpa/products')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: '更新商品' }]" />
      <h1 class="text-3xl font-bold text-gray-900">更新商品信息</h1>
      <p class="mt-2 text-gray-600">编辑DPA商品库中的商品</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleSubmit" class="space-y-6 max-w-2xl">
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品名称</label>
            <input v-model="form.name" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">商品分类</label>
            <select v-model="form.category" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option value="3C数码">3C数码</option>
              <option value="服装鞋包">服装鞋包</option>
              <option value="美妆个护">美妆个护</option>
              <option value="家居家装">家居家装</option>
            </select>
          </div>
        </div>
        <div class="grid grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">售价 (元)</label>
            <input v-model.number="form.price" type="number" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">原价 (元)</label>
            <input v-model.number="form.originalPrice" type="number" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">库存</label>
            <input v-model.number="form.stock" type="number" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">商品描述</label>
          <textarea v-model="form.description" rows="4" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">状态</label>
          <select v-model="form.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            <option value="active">在售</option>
            <option value="inactive">下架</option>
          </select>
        </div>
        <div class="flex justify-end gap-4 pt-4">
          <button type="button" @click="router.back()" class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50">取消</button>
          <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>
