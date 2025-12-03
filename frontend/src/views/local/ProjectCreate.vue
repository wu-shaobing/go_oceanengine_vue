<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '项目管理', path: '/local/project' }, { name: '创建项目' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建项目</h1>
      <p class="text-gray-600 mt-1">新建本地推广项目</p>
    </div>

    <div class="max-w-3xl">
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">基本信息</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">项目名称</label>
            <input type="text" v-model="form.name" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入项目名称">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">行业分类</label>
            <select v-model="form.industry" class="w-full border border-gray-300 rounded px-3 py-2">
              <option value="">请选择行业</option>
              <option value="food">餐饮美食</option>
              <option value="beauty">丽人美业</option>
              <option value="education">教育培训</option>
              <option value="medical">医疗健康</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">投放城市</label>
            <select v-model="form.city" class="w-full border border-gray-300 rounded px-3 py-2">
              <option value="">请选择城市</option>
              <option value="beijing">北京市</option>
              <option value="shanghai">上海市</option>
              <option value="guangzhou">广州市</option>
              <option value="shenzhen">深圳市</option>
            </select>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">门店信息</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">门店名称</label>
            <input type="text" v-model="form.storeName" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入门店名称">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">门店地址</label>
            <input type="text" v-model="form.storeAddress" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入详细地址">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">联系电话</label>
            <input type="text" v-model="form.storePhone" class="w-full border border-gray-300 rounded px-3 py-2" placeholder="请输入联系电话">
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">预算设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算</label>
            <div class="flex items-center">
              <span class="mr-2">¥</span>
              <input type="number" v-model="form.dailyBudget" class="w-48 border border-gray-300 rounded px-3 py-2">
              <span class="ml-2 text-gray-500">元/天</span>
            </div>
          </div>
        </div>
      </div>

      <div class="flex justify-end space-x-4">
        <button class="px-6 py-2 border border-gray-300 rounded hover:bg-gray-50">取消</button>
        <button class="px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">创建项目</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = ref({
  name: '',
  industry: '',
  city: '',
  storeName: '',
  storeAddress: '',
  storePhone: '',
  dailyBudget: 1000
})
</script>
