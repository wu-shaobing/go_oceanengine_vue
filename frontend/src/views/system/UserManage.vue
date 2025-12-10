<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import { systemApi, type User, type Role } from '@/api'

const loading = ref(true)
const users = ref<User[]>([])
const roles = ref<Role[]>([])
const searchKeyword = ref('')
const roleFilter = ref<number | undefined>(undefined)
const statusFilter = ref<number | undefined>(undefined)

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
  { key: 'role_name', title: '角色' },
  { key: 'status', title: '状态', width: 100 },
  { key: 'last_login_at', title: '最后登录' },
  { key: 'actions', title: '操作', width: 200, align: 'center' as const }
]

// 获取用户列表
const fetchUsers = async () => {
  loading.value = true
  try {
    const params: Record<string, unknown> = {
      page: pagination.current,
      page_size: pagination.pageSize
    }
    if (searchKeyword.value) params.username = searchKeyword.value
    if (roleFilter.value !== undefined) params.role_id = roleFilter.value
    if (statusFilter.value !== undefined) params.status = statusFilter.value

    const res = await systemApi.getUserList(params)
    users.value = res.list || []
    pagination.total = res.total || 0
  } catch (error) {
    console.error('获取用户列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取角色列表
const fetchRoles = async () => {
  try {
    roles.value = await systemApi.getAllRoles()
  } catch (error) {
    console.error('获取角色列表失败:', error)
  }
}

const handleSearch = () => {
  pagination.current = 1
  fetchUsers()
}

const handleReset = () => {
  searchKeyword.value = ''
  roleFilter.value = undefined
  statusFilter.value = undefined
  pagination.current = 1
  fetchUsers()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchUsers()
}

// 用户表单弹窗
const showUserDialog = ref(false)
const userFormMode = ref<'add' | 'edit'>('add')
const userForm = reactive({
  id: 0,
  username: '',
  password: '',
  nickname: '',
  phone: '',
  email: '',
  role_id: 0,
  status: 1,
  remark: ''
})

const resetUserForm = () => {
  userForm.id = 0
  userForm.username = ''
  userForm.password = ''
  userForm.nickname = ''
  userForm.phone = ''
  userForm.email = ''
  userForm.role_id = 0
  userForm.status = 1
  userForm.remark = ''
}

const handleAddUser = () => {
  userFormMode.value = 'add'
  resetUserForm()
  showUserDialog.value = true
}

const handleEditUser = (user: User) => {
  userFormMode.value = 'edit'
  userForm.id = user.id
  userForm.username = user.username
  userForm.password = ''
  userForm.nickname = user.nickname
  userForm.phone = user.phone
  userForm.email = user.email
  userForm.role_id = user.role_id
  userForm.status = user.status
  showUserDialog.value = true
}

const handleSaveUser = async () => {
  try {
    if (userFormMode.value === 'add') {
      if (!userForm.username || !userForm.password || !userForm.role_id) {
        alert('请填写必填字段')
        return
      }
      await systemApi.createUser({
        username: userForm.username,
        password: userForm.password,
        nickname: userForm.nickname,
        phone: userForm.phone,
        email: userForm.email,
        role_id: userForm.role_id,
        status: userForm.status,
        remark: userForm.remark
      })
      alert('创建成功')
    } else {
      await systemApi.updateUser(userForm.id, {
        id: userForm.id,
        nickname: userForm.nickname,
        phone: userForm.phone,
        email: userForm.email,
        role_id: userForm.role_id,
        status: userForm.status,
        remark: userForm.remark
      })
      alert('更新成功')
    }
    showUserDialog.value = false
    fetchUsers()
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败')
  }
}

// 重置密码弹窗
const showResetDialog = ref(false)
const resetUserId = ref(0)
const resetUsername = ref('')
const newPassword = ref('')

const handleResetPassword = (user: User) => {
  resetUserId.value = user.id
  resetUsername.value = user.username
  newPassword.value = ''
  showResetDialog.value = true
}

const confirmResetPassword = async () => {
  if (!newPassword.value || newPassword.value.length < 6) {
    alert('密码至少6位')
    return
  }
  try {
    await systemApi.resetPassword(resetUserId.value, newPassword.value)
    alert('密码重置成功')
    showResetDialog.value = false
  } catch (error) {
    console.error('重置密码失败:', error)
    alert('重置密码失败')
  }
}

// 删除用户
const handleDeleteUser = async (user: User) => {
  if (!confirm(`确定删除用户 ${user.username}?`)) return
  try {
    await systemApi.deleteUser(user.id)
    alert('删除成功')
    fetchUsers()
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// 切换用户状态
const handleToggleStatus = async (user: User) => {
  const newStatus = user.status === 1 ? 0 : 1
  try {
    await systemApi.updateUser(user.id, { id: user.id, status: newStatus })
    user.status = newStatus
  } catch (error) {
    console.error('更新状态失败:', error)
    alert('更新状态失败')
  }
}

onMounted(() => {
  fetchUsers()
  fetchRoles()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '用户管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">用户管理</h1>
          <p class="mt-2 text-gray-600">管理系统用户和权限，共 {{ pagination.total }} 人</p>
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
      <div class="flex flex-wrap items-center gap-4">
        <div class="relative flex-1 min-w-48 max-w-md">
          <input
            v-model="searchKeyword"
            type="text"
            placeholder="搜索用户名..."
            @keyup.enter="handleSearch"
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
          </svg>
        </div>
        <select
          v-model="roleFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg"
        >
          <option :value="undefined">全部角色</option>
          <option v-for="role in roles" :key="role.id" :value="role.id">{{ role.name }}</option>
        </select>
        <select
          v-model="statusFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg"
        >
          <option :value="undefined">全部状态</option>
          <option :value="1">正常</option>
          <option :value="0">禁用</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
        <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="handleReset">重置</button>
      </div>
    </div>

    <!-- 表格 -->
    <div class="bg-white rounded-lg border border-gray-200">
      <DataTable :columns="columns" :data="users" :loading="loading">
        <template #username="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #role_name="{ value }">
          <span class="px-2 py-1 text-xs rounded-full bg-blue-100 text-blue-700">{{ value || '-' }}</span>
        </template>

        <template #status="{ row }">
          <button @click="handleToggleStatus(row)" class="focus:outline-none">
            <StatusBadge
              :status="row.status === 1 ? 'success' : 'danger'"
              :text="row.status === 1 ? '正常' : '禁用'"
              show-icon
            />
          </button>
        </template>

        <template #actions="{ row }">
          <div class="flex items-center justify-center gap-2">
            <button class="text-blue-600 hover:text-blue-800 text-sm" @click="handleEditUser(row)">编辑</button>
            <button class="text-gray-600 hover:text-gray-800 text-sm" @click="handleResetPassword(row)">重置密码</button>
            <button class="text-red-600 hover:text-red-800 text-sm" @click="handleDeleteUser(row)">删除</button>
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

    <!-- 用户表单弹窗 -->
    <div v-if="showUserDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[500px] max-h-[90vh] overflow-y-auto">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">{{ userFormMode === 'add' ? '添加用户' : '编辑用户' }}</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">用户名 <span class="text-red-500">*</span></label>
            <input v-model="userForm.username" type="text" :disabled="userFormMode === 'edit'"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg disabled:bg-gray-100" placeholder="请输入用户名">
          </div>
          <div v-if="userFormMode === 'add'">
            <label class="block text-sm font-medium text-gray-700 mb-1">密码 <span class="text-red-500">*</span></label>
            <input v-model="userForm.password" type="password"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入密码（6位以上）">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">昵称</label>
            <input v-model="userForm.nickname" type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入昵称">
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">手机号</label>
              <input v-model="userForm.phone" type="text"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入手机号">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
              <input v-model="userForm.email" type="email"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入邮箱">
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">角色 <span class="text-red-500">*</span></label>
              <select v-model="userForm.role_id" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="0" disabled>请选择角色</option>
                <option v-for="role in roles" :key="role.id" :value="role.id">{{ role.name }}</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
              <select v-model="userForm.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">正常</option>
                <option :value="0">禁用</option>
              </select>
            </div>
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showUserDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSaveUser">保存</button>
        </div>
      </div>
    </div>

    <!-- 重置密码弹窗 -->
    <div v-if="showResetDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-96">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">重置密码</h3>
        <p class="text-sm text-gray-600 mb-4">为用户 <span class="font-medium">{{ resetUsername }}</span> 设置新密码</p>
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">新密码</label>
          <input v-model="newPassword" type="password" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入新密码（6位以上）">
        </div>
        <div class="flex justify-end gap-2">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showResetDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="confirmResetPassword">确认重置</button>
        </div>
      </div>
    </div>
  </div>
</template>
