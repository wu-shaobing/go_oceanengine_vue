<script setup lang="ts">
/**
 * 主布局组件
 * 包含侧边栏、顶部导航和内容区域
 */
import AppHeader from './AppHeader.vue'
import AppSidebar from './AppSidebar.vue'
import { ref } from 'vue'

const sidebarCollapsed = ref(false)

const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 侧边栏 -->
    <AppSidebar :collapsed="sidebarCollapsed" />
    
    <!-- 主内容区 -->
    <div
      class="transition-all duration-300"
      :class="sidebarCollapsed ? 'ml-16' : 'ml-64'"
    >
      <!-- 顶部导航 -->
      <AppHeader @toggle-sidebar="toggleSidebar" />
      
      <!-- 页面内容 -->
      <main class="p-6">
        <router-view />
      </main>
    </div>
  </div>
</template>
