<script setup lang="ts">
interface BreadcrumbItem {
  name: string
  path?: string
}

defineProps<{
  items: BreadcrumbItem[]
}>()
</script>

<template>
  <div class="flex items-center text-sm text-gray-500 mb-2">
    <router-link to="/" class="hover:text-blue-600">首页</router-link>
    <template v-for="(item, index) in items" :key="index">
      <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
      </svg>
      <router-link
        v-if="item.path && index < items.length - 1"
        :to="item.path"
        class="hover:text-blue-600"
      >
        {{ item.name }}
      </router-link>
      <span v-else class="text-gray-900">{{ item.name }}</span>
    </template>
  </div>
</template>
