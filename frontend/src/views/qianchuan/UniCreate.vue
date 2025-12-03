<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '全域推广' }, { name: '创建推广' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创建全域推广</h1>
      <p class="text-gray-600 mt-1">智能托管投放，全域流量覆盖</p>
    </div>

    <div class="max-w-4xl">
      <!-- 推广类型 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">选择推广类型</h3>
        <div class="grid grid-cols-2 gap-4">
          <div @click="form.type = 'live'" 
            :class="form.type === 'live' ? 'border-blue-500 bg-blue-50' : 'border-gray-200'"
            class="border-2 rounded-lg p-6 cursor-pointer hover:border-blue-300 transition-colors">
            <div class="text-3xl mb-3">📺</div>
            <div class="font-medium text-lg">直播全域推广</div>
            <div class="text-sm text-gray-500 mt-2">为直播间引流，全域流量覆盖，智能优化ROI</div>
          </div>
          <div @click="form.type = 'video'" 
            :class="form.type === 'video' ? 'border-blue-500 bg-blue-50' : 'border-gray-200'"
            class="border-2 rounded-lg p-6 cursor-pointer hover:border-blue-300 transition-colors">
            <div class="text-3xl mb-3">🎬</div>
            <div class="font-medium text-lg">短视频全域推广</div>
            <div class="text-sm text-gray-500 mt-2">通过短视频推广商品，智能分发至全域流量</div>
          </div>
        </div>
      </div>

      <!-- 基本信息 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">基本设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">推广名称 <span class="text-red-500">*</span></label>
            <input type="text" v-model="form.name" class="w-full border border-gray-300 rounded-lg px-3 py-2" placeholder="请输入推广名称">
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">关联店铺 <span class="text-red-500">*</span></label>
              <select v-model="form.shopId" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">请选择店铺</option>
                <option value="7001">品牌官方旗舰店</option>
                <option value="7002">美妆专营店</option>
              </select>
            </div>
            <div v-if="form.type === 'live'">
              <label class="block text-sm font-medium text-gray-700 mb-1">直播间 <span class="text-red-500">*</span></label>
              <select v-model="form.liveRoomId" class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option value="">请选择直播间</option>
                <option value="L001">主播A直播间</option>
                <option value="L002">品牌直播间</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 预算与ROI -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">预算与目标</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">日预算 <span class="text-red-500">*</span></label>
            <div class="flex items-center">
              <input type="number" v-model="form.budget" class="border border-gray-300 rounded-lg px-3 py-2 w-48" placeholder="请输入日预算">
              <span class="ml-2 text-gray-500">元/天，最低300元</span>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">ROI目标 <span class="text-red-500">*</span></label>
            <div class="flex items-center">
              <input type="number" v-model="form.roiTarget" step="0.1" class="border border-gray-300 rounded-lg px-3 py-2 w-48" placeholder="请输入ROI目标">
              <span class="ml-2 text-gray-500">建议范围: 2.0 - 10.0</span>
            </div>
            <div class="text-sm text-gray-500 mt-1">系统将自动优化投放，尽量达到您设定的ROI目标</div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">投放时间</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.scheduleType" value="unlimited" class="mr-2"> 长期投放
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.scheduleType" value="scheduled" class="mr-2"> 指定时间
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- 素材设置 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-medium mb-4">素材设置</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">素材来源</label>
            <div class="flex items-center space-x-4">
              <label class="flex items-center">
                <input type="radio" v-model="form.materialSource" value="auto" class="mr-2"> 系统自动生成
              </label>
              <label class="flex items-center">
                <input type="radio" v-model="form.materialSource" value="manual" class="mr-2"> 手动上传
              </label>
            </div>
            <div class="text-sm text-gray-500 mt-1">系统自动生成将从直播间/商品中智能提取素材</div>
          </div>
          <div v-if="form.materialSource === 'manual'">
            <label class="block text-sm font-medium text-gray-700 mb-1">上传素材</label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
              <div class="text-gray-400 mb-2">点击或拖拽上传视频素材</div>
              <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">选择文件</button>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex justify-end space-x-4">
        <router-link to="/qianchuan/uni" class="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">取消</router-link>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">创建推广</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = ref({
  type: 'live',
  name: '',
  shopId: '',
  liveRoomId: '',
  budget: 1000,
  roiTarget: 4.0,
  scheduleType: 'unlimited',
  materialSource: 'auto'
})
</script>
