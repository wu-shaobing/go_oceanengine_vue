<script setup lang="ts">
/**
 * 认证页面布局组件
 * 用于登录、注册等认证相关页面
 */
</script>

<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
    <div class="w-full max-w-md">
      <!-- Logo -->
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-blue-600 rounded-2xl mb-4">
          <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-gray-900">巨量引擎</h1>
        <p class="text-gray-500 mt-1">营销管理平台</p>
      </div>
      
      <!-- 内容区域 -->
      <div class="bg-white rounded-2xl shadow-xl p-8">
        <slot />
      </div>
      
      <!-- 底部信息 -->
      <div class="text-center mt-8 text-sm text-gray-500">
        <p>&copy; {{ new Date().getFullYear() }} 巨量引擎. All rights reserved.</p>
      </div>
    </div>
    
    <!-- 装饰元素 -->
    <div class="fixed top-0 left-0 w-full h-full overflow-hidden pointer-events-none -z-10">
      <div class="absolute -top-40 -right-40 w-80 h-80 bg-blue-200 rounded-full opacity-20 blur-3xl"></div>
      <div class="absolute -bottom-40 -left-40 w-80 h-80 bg-indigo-200 rounded-full opacity-20 blur-3xl"></div>
    </div>
  </div>
</template>
