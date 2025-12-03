import { ref, type UnwrapRef } from 'vue'

export function useRequest<T, P extends unknown[]>(
  fn: (...args: P) => Promise<T>
) {
  const data = ref<T | null>(null)
  const loading = ref(false)
  const error = ref<Error | null>(null)

  const execute = async (...args: P): Promise<T | null> => {
    loading.value = true
    error.value = null

    try {
      const result = await fn(...args)
      data.value = result as UnwrapRef<T>
      return result
    } catch (e) {
      error.value = e as Error
      return null
    } finally {
      loading.value = false
    }
  }

  const reset = () => {
    data.value = null
    error.value = null
    loading.value = false
  }

  return {
    data,
    loading,
    error,
    execute,
    reset
  }
}
