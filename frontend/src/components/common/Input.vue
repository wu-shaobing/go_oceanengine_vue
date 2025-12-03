<script setup lang="ts">
interface Props {
  modelValue?: string | number
  type?: string
  placeholder?: string
  disabled?: boolean
  readonly?: boolean
  error?: string
  prefix?: string
  suffix?: string
}

withDefaults(defineProps<Props>(), {
  type: 'text',
  disabled: false,
  readonly: false
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string | number): void
  (e: 'focus', event: FocusEvent): void
  (e: 'blur', event: FocusEvent): void
}>()

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<template>
  <div class="relative">
    <div v-if="prefix" class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
      <span class="text-gray-500 text-sm">{{ prefix }}</span>
    </div>
    
    <input
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :readonly="readonly"
      @input="handleInput"
      @focus="emit('focus', $event)"
      @blur="emit('blur', $event)"
      :class="[
        'w-full px-3 py-2 text-sm border rounded-lg transition-colors',
        'focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent',
        'disabled:bg-gray-100 disabled:cursor-not-allowed',
        error ? 'border-red-500' : 'border-gray-300',
        prefix ? 'pl-8' : '',
        suffix ? 'pr-8' : ''
      ]"
    />
    
    <div v-if="suffix" class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
      <span class="text-gray-500 text-sm">{{ suffix }}</span>
    </div>
    
    <p v-if="error" class="mt-1 text-xs text-red-500">{{ error }}</p>
  </div>
</template>
