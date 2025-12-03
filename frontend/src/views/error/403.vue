<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  router.back()
}

const goHome = () => {
  router.push('/')
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="text-center">
      <div class="mb-8">
        <svg class="mx-auto w-24 h-24 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" 
            d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
        </svg>
      </div>
      
      <h1 class="text-6xl font-bold text-gray-900 mb-4">403</h1>
      <h2 class="text-2xl font-semibold text-gray-700 mb-4">访问被拒绝</h2>
      <p class="text-gray-500 mb-8 max-w-md mx-auto">
        抱歉，您没有权限访问此页面。请联系管理员获取相应权限。
      </p>
      
      <div class="flex items-center justify-center gap-4">
        <button
          @click="goBack"
          class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        >
          返回上页
        </button>
        <button
          @click="goHome"
          class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
        >
          返回首页
        </button>
      </div>
    </div>
  </div>
</template>
