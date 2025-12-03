<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const showUserMenu = ref(false)

// 产品线快捷导航
const productNavItems = [
  { name: '千川', path: '/qianchuan', color: 'orange' },
  { name: '本地推', path: '/local', color: 'green' },
  { name: '企业号', path: '/enterprise', color: 'purple' },
  { name: '星图', path: '/star', color: 'yellow' },
  { name: '代理商', path: '/agent', color: 'blue' }
]

const isProductActive = (path: string) => {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}
</script>

<template>
  <header class="sticky top-0 z-40 bg-white/95 backdrop-blur border-b border-gray-200">
    <div class="flex h-16 items-center justify-between px-6">
      <!-- Logo + Product Nav -->
      <div class="flex items-center gap-6">
        <router-link to="/" class="flex items-center gap-3">
          <div class="flex h-9 w-9 items-center justify-center rounded-lg bg-blue-600">
            <svg class="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
          </div>
          <div class="hidden lg:block">
            <h1 class="text-base font-bold text-gray-900">巨量引擎</h1>
          </div>
        </router-link>
        
        <!-- Product Line Quick Nav -->
        <nav class="hidden md:flex items-center gap-1">
          <router-link
            v-for="item in productNavItems"
            :key="item.path"
            :to="item.path"
            :class="[
              'px-3 py-1.5 text-sm font-medium rounded-md transition-colors',
              isProductActive(item.path)
                ? 'bg-gray-100 text-gray-900'
                : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
            ]"
          >
            {{ item.name }}
          </router-link>
        </nav>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-3">
        <!-- Search -->
        <button class="hidden lg:flex items-center justify-center w-10 h-10 rounded-lg hover:bg-gray-100 transition-colors">
          <svg class="h-5 w-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
          </svg>
        </button>

        <!-- Notifications -->
        <button class="relative flex items-center justify-center w-10 h-10 rounded-lg hover:bg-gray-100 transition-colors">
          <svg class="h-5 w-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>
          </svg>
          <span class="absolute right-1.5 top-1.5 h-2 w-2 rounded-full bg-red-500"></span>
        </button>

        <!-- User Menu -->
        <div class="relative">
          <button
            @click="showUserMenu = !showUserMenu"
            class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100 transition-colors"
          >
            <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-white">
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
            </div>
            <span class="hidden text-sm font-medium md:inline-block">管理员</span>
            <svg class="h-4 w-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>

          <!-- Dropdown Menu -->
          <div
            v-if="showUserMenu"
            class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 py-1 z-50"
          >
            <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">个人设置</a>
            <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">账户信息</a>
            <hr class="my-1">
            <a href="#" class="block px-4 py-2 text-sm text-red-600 hover:bg-gray-100">退出登录</a>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>
