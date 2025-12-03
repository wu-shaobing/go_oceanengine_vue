import { ref, reactive, computed } from 'vue'
import type { PageResponse } from '@/api/request'

export interface TableState<T> {
  data: T[]
  total: number
  loading: boolean
}

interface ReactiveTableState {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data: any[]
  total: number
  loading: boolean
}

export interface PaginationState {
  page: number
  pageSize: number
}

export function useTable<T, P extends object>(
  fetchFn: (params: P & { page: number; page_size: number }) => Promise<PageResponse<T>>,
  defaultParams?: Partial<P>
) {
  const state = reactive<ReactiveTableState>({
    data: [] as T[],
    total: 0,
    loading: false
  })

  const pagination = reactive<PaginationState>({
    page: 1,
    pageSize: 20
  })

  const params = ref<Partial<P>>(defaultParams || {})

  const totalPages = computed(() => Math.ceil(state.total / pagination.pageSize))

  const fetchData = async () => {
    state.loading = true
    try {
      const res = await fetchFn({
        ...params.value,
        page: pagination.page,
        page_size: pagination.pageSize
      } as P & { page: number; page_size: number })
      
      state.data = res.list
      state.total = res.total
    } catch (error) {
      console.error('Fetch table data error:', error)
      state.data = []
      state.total = 0
    } finally {
      state.loading = false
    }
  }

  const setPage = (page: number) => {
    pagination.page = page
    fetchData()
  }

  const setPageSize = (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    fetchData()
  }

  const setParams = (newParams: Partial<P>) => {
    params.value = { ...params.value, ...newParams }
    pagination.page = 1
    fetchData()
  }

  const refresh = () => {
    fetchData()
  }

  const reset = () => {
    params.value = defaultParams || {}
    pagination.page = 1
    fetchData()
  }

  return {
    state,
    pagination,
    params,
    totalPages,
    fetchData,
    setPage,
    setPageSize,
    setParams,
    refresh,
    reset
  }
}
