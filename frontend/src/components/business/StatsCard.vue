<script setup lang="ts">
defineProps<{
  title: string
  value: string | number
  icon?: string
  color?: 'blue' | 'green' | 'orange' | 'red' | 'purple'
  trend?: 'up' | 'down' | 'neutral'
  trendValue?: string
}>()
</script>

<template>
  <div class="bg-white rounded-lg border border-gray-200 p-6">
    <div class="flex items-start justify-between">
      <div>
        <p class="text-sm font-medium text-gray-500">{{ title }}</p>
        <p class="mt-2 text-3xl font-bold text-gray-900">{{ value }}</p>
        <div v-if="trend && trendValue" class="mt-2 flex items-center gap-1">
          <!-- Up trend -->
          <svg v-if="trend === 'up'" class="w-4 h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"/>
          </svg>
          <!-- Down trend -->
          <svg v-else-if="trend === 'down'" class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"/>
          </svg>
          <!-- Neutral -->
          <svg v-else class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>
          </svg>
          <span
            :class="[
              'text-sm font-medium',
              trend === 'up' ? 'text-green-600' : trend === 'down' ? 'text-red-600' : 'text-gray-500'
            ]"
          >
            {{ trendValue }}
          </span>
          <span class="text-sm text-gray-400">较上期</span>
        </div>
      </div>
      <div v-if="icon" class="w-10 h-10 rounded-lg flex items-center justify-center"
        :class="[
          {
            'bg-blue-100': color === 'blue' || !color,
            'bg-green-100': color === 'green',
            'bg-orange-100': color === 'orange',
            'bg-red-100': color === 'red',
            'bg-purple-100': color === 'purple',
          }
        ]"
      >
        <svg class="w-5 h-5"
          :class="[
            {
              'text-blue-600': color === 'blue' || !color,
              'text-green-600': color === 'green',
              'text-orange-600': color === 'orange',
              'text-red-600': color === 'red',
              'text-purple-600': color === 'purple',
            }
          ]"
          fill="none" stroke="currentColor" viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
        </svg>
      </div>
    </div>
  </div>
</template>
