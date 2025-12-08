<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'

const loading = ref(true)

interface User {
  id: number
  username: string
  nickname: string
  email: string
  phone: string
  role: string
  status: string
  created_at: string
  last_login: string
}

const users = ref<User[]>([])
const searchKeyword = ref('')
const roleFilter = ref('all')

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

const columns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'username', title: '用户名' },
  { key: 'nickname', title: '昵称' },
  { key: 'email', title: '邮箱' },
  { key: 'role', title: '角色' },
  { key: 'status', title: '状态', width: 100 },
  { key: 'last_login', title: '最后登录' },
  { key: 'actions', title: '操作', width: 150, align: 'center' as const }
]

const roleOptions = [
  { value: 'all', label: '全部角色' },
  { value: 'admin', label: '管理员' },
  { value: 'operator', label: '运营' },
  { value: 'viewer', label: '观察者' }
]

const fetchUsers = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  users.value = Array.from({ length: 10 }, (_, i) => ({
    id: i + 1,
    username: `user${i + 1}`,
    nickname: `用户 ${i + 1}`,
    email: `user${i + 1}@example.com`,
    phone: `138****${String(1000 + i).slice(-4)}`,
    role: ['admin', 'operator', 'viewer'][i % 3],
    status: i % 5 === 0 ? 'disabled' : 'active',
    created_at: '2024-01-15 10:30:00',
    last_login: '2024-01-20 14:25:00'
  }))
  
  pagination.total = 50
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchUsers()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchUsers()
}

const getRoleLabel = (role: string) => {
  const map: Record<string, string> = {
    admin: '管理员',
    operator: '运营',
    viewer: '观察者'
  }
  return map[role] || role
}

const handleAddUser = () => {
  alert('添加用户')
}

const handleEditUser = (user: User) => {
  alert(`编辑用户: ${user.username}`)
}

const handleResetPassword = (user: User) => {
  if (confirm(`确定重置 ${user.username} 的密码?`)) {
    alert('密码重置成功')
  }
}

const handleDeleteUser = (user: User) => {
  if (confirm(`确定删除用户 ${user.username}?`)) {
    alert('删除成功')
  }
}

onMounted(fetchUsers)
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '用户管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">用户管理</h1>
          <p class="mt-2 text-gray-600">管理系统用户和权限</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2" @click="handleAddUser">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
          </svg>
          添加用户
        </button>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex items-center gap-4">
        <div class="relative flex-1 max-w-md">
          <input
            v-model="searchKeyword"
            type="text"
            placeholder="搜索用户名、邮箱..."
            @keyup.enter="handleSearch"
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
          </svg>
        </div>
        <select
          v-model="roleFilter"
          @change="handleSearch"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option v-for="opt in roleOptions" :key="opt.value" :value="opt.value">
            {{ opt.label }}
          </option>
        </select>
      </div>
    </div>

    <!-- 表格 -->
    <div class="bg-white rounded-lg border border-gray-200">
      <DataTable :columns="columns" :data="users" :loading="loading">
        <template #username="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #role="{ value }">
          <span class="px-2 py-1 text-xs rounded-full" :class="{
            'bg-purple-100 text-purple-700': value === 'admin',
            'bg-blue-100 text-blue-700': value === 'operator',
            'bg-gray-100 text-gray-700': value === 'viewer'
          }">
            {{ getRoleLabel(value) }}
          </span>
        </template>

        <template #status="{ value }">
          <StatusBadge
            :status="value === 'active' ? 'success' : 'danger'"
            :text="value === 'active' ? '正常' : '禁用'"
            show-icon
          />
        </template>

        <template #actions="{ row }">
          <div class="flex items-center justify-center gap-2">
            <button class="text-blue-600 hover:text-blue-800" @click="handleEditUser(row)">编辑</button>
            <button class="text-gray-600 hover:text-gray-800" @click="handleResetPassword(row)">重置密码</button>
            <button class="text-red-600 hover:text-red-800" @click="handleDeleteUser(row)">删除</button>
          </div>
        </template>
      </DataTable>

      <div class="p-4 border-t border-gray-200">
        <Pagination
          :current="pagination.current"
          :total="pagination.total"
          :page-size="pagination.pageSize"
          @change="handlePageChange"
        />
      </div>
    </div>
  </div>
</template>
