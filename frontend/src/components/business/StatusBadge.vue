<script setup lang="ts">
defineProps<{
  status: 'success' | 'warning' | 'error' | 'danger' | 'info' | 'default'
  text: string
  showIcon?: boolean
}>()
</script>

<template>
  <span
    :class="[
      'px-2.5 py-0.5 text-xs font-medium rounded-full ring-1 ring-inset flex items-center gap-1 w-fit',
      {
        'text-green-700 bg-green-50 ring-green-600/20': status === 'success',
        'text-yellow-700 bg-yellow-50 ring-yellow-600/20': status === 'warning',
        'text-red-700 bg-red-50 ring-red-600/20': status === 'error' || status === 'danger',
        'text-blue-700 bg-blue-50 ring-blue-600/20': status === 'info',
        'text-gray-700 bg-gray-50 ring-gray-600/20': status === 'default',
      }
    ]"
  >
    <!-- Success Icon -->
    <svg v-if="showIcon && status === 'success'" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
    </svg>
    <!-- Error Icon -->
    <svg v-if="showIcon && status === 'error'" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
    </svg>
    <!-- Warning Icon -->
    <svg v-if="showIcon && status === 'warning'" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
    </svg>
    <!-- Info Icon -->
    <svg v-if="showIcon && status === 'info'" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
    </svg>
    {{ text }}
  </span>
</template>
