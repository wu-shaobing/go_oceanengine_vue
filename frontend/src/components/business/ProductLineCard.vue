<script setup lang="ts">
defineProps<{
  title: string
  description: string
  icon: string
  color: 'orange' | 'green' | 'purple' | 'yellow' | 'blue' | 'gray'
  to: string
}>()

const colorClasses = {
  orange: {
    bg: 'bg-orange-50',
    iconBg: 'bg-orange-100',
    iconText: 'text-orange-600',
    border: 'border-orange-200',
    hoverBorder: 'hover:border-orange-400'
  },
  green: {
    bg: 'bg-green-50',
    iconBg: 'bg-green-100',
    iconText: 'text-green-600',
    border: 'border-green-200',
    hoverBorder: 'hover:border-green-400'
  },
  purple: {
    bg: 'bg-purple-50',
    iconBg: 'bg-purple-100',
    iconText: 'text-purple-600',
    border: 'border-purple-200',
    hoverBorder: 'hover:border-purple-400'
  },
  yellow: {
    bg: 'bg-yellow-50',
    iconBg: 'bg-yellow-100',
    iconText: 'text-yellow-600',
    border: 'border-yellow-200',
    hoverBorder: 'hover:border-yellow-400'
  },
  blue: {
    bg: 'bg-blue-50',
    iconBg: 'bg-blue-100',
    iconText: 'text-blue-600',
    border: 'border-blue-200',
    hoverBorder: 'hover:border-blue-400'
  },
  gray: {
    bg: 'bg-gray-50',
    iconBg: 'bg-gray-100',
    iconText: 'text-gray-600',
    border: 'border-gray-200',
    hoverBorder: 'hover:border-gray-400'
  }
}
</script>

<template>
  <router-link
    :to="to"
    :class="[
      'group block p-5 rounded-xl border-2 transition-all duration-200',
      'hover:shadow-md hover:-translate-y-0.5',
      colorClasses[color].border,
      colorClasses[color].hoverBorder,
      colorClasses[color].bg
    ]"
  >
    <div class="flex items-start gap-4">
      <!-- Icon -->
      <div :class="['w-12 h-12 rounded-lg flex items-center justify-center flex-shrink-0', colorClasses[color].iconBg]">
        <!-- Shop Icon (千川) -->
        <svg v-if="icon === 'shop'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
        </svg>
        <!-- Location Icon (本地推) -->
        <svg v-else-if="icon === 'location'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
        </svg>
        <!-- Building Icon (企业号) -->
        <svg v-else-if="icon === 'building'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"/>
        </svg>
        <!-- Star Icon (星图) -->
        <svg v-else-if="icon === 'star'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/>
        </svg>
        <!-- Tools Icon (工具) -->
        <svg v-else-if="icon === 'tools'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
        </svg>
        <!-- Agent Icon (代理商) -->
        <svg v-else-if="icon === 'agent'" :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
        </svg>
        <!-- Default Icon -->
        <svg v-else :class="['w-6 h-6', colorClasses[color].iconText]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
        </svg>
      </div>
      
      <!-- Content -->
      <div class="flex-1 min-w-0">
        <h3 class="text-base font-semibold text-gray-900 group-hover:text-gray-700 flex items-center gap-2">
          {{ title }}
          <svg class="w-4 h-4 text-gray-400 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </h3>
        <p class="mt-1 text-sm text-gray-500 line-clamp-2">{{ description }}</p>
      </div>
    </div>
  </router-link>
</template>
