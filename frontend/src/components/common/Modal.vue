<script setup lang="ts">
import { watch } from 'vue'

interface Props {
  modelValue: boolean
  title?: string
  width?: string
  closable?: boolean
  maskClosable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  width: '500px',
  closable: true,
  maskClosable: true
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
  (e: 'confirm'): void
}>()

const close = () => {
  emit('update:modelValue', false)
  emit('close')
}

const handleMaskClick = () => {
  if (props.maskClosable) {
    close()
  }
}

const handleConfirm = () => {
  emit('confirm')
}

// 禁止背景滚动
watch(() => props.modelValue, (val) => {
  document.body.style.overflow = val ? 'hidden' : ''
})
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="modelValue" class="fixed inset-0 z-50 flex items-center justify-center">
        <!-- 遮罩层 -->
        <div
          class="absolute inset-0 bg-black/50"
          @click="handleMaskClick"
        />
        
        <!-- 对话框 -->
        <div
          class="relative bg-white rounded-xl shadow-xl max-h-[90vh] overflow-hidden"
          :style="{ width }"
        >
          <!-- 头部 -->
          <div v-if="title || closable" class="flex items-center justify-between px-6 py-4 border-b border-gray-200">
            <h3 class="text-lg font-semibold text-gray-900">{{ title }}</h3>
            <button
              v-if="closable"
              @click="close"
              class="text-gray-400 hover:text-gray-600 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>
          
          <!-- 内容 -->
          <div class="px-6 py-4 overflow-y-auto max-h-[60vh]">
            <slot />
          </div>
          
          <!-- 底部 -->
          <div v-if="$slots.footer" class="px-6 py-4 border-t border-gray-200 bg-gray-50">
            <slot name="footer" :close="close" :confirm="handleConfirm" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-from > div:last-child,
.modal-leave-to > div:last-child {
  transform: scale(0.95);
}
</style>
