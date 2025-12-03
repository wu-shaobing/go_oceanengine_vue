<script setup lang="ts">
/**
 * 面包屑导航组件
 */
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

interface BreadcrumbItem {
  name: string
  path?: string
}

const props = defineProps<{
  items?: BreadcrumbItem[]
}>()

const route = useRoute()
const router = useRouter()

// 如果没有传入 items，则根据路由自动生成
const breadcrumbs = computed(() => {
  if (props.items && props.items.length > 0) {
    return props.items
  }
  
  // 从路由 meta 生成
  const matched = route.matched
  return matched
    .filter(r => r.meta?.title)
    .map(r => ({
      name: r.meta.title as string,
      path: r.path
    }))
})

const goTo = (path?: string) => {
  if (path) {
    router.push(path)
  }
}
</script>

<template>
  <nav class="flex items-center space-x-2 text-sm mb-4">
    <!-- 首页链接 -->
    <router-link
      to="/"
      class="text-gray-500 hover:text-gray-700 transition-colors"
    >
      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
      </svg>
    </router-link>
    
    <!-- 面包屑项 -->
    <template v-for="(item, index) in breadcrumbs" :key="index">
      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
      </svg>
      
      <span
        v-if="index === breadcrumbs.length - 1"
        class="text-gray-900 font-medium"
      >
        {{ item.name }}
      </span>
      <button
        v-else
        @click="goTo(item.path)"
        class="text-gray-500 hover:text-gray-700 transition-colors"
      >
        {{ item.name }}
      </button>
    </template>
  </nav>
</template>
