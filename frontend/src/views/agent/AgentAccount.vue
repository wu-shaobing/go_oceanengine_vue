<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const accounts = ref([
  { id: 'ACC001', name: '主账号', email: 'admin@example.com', role: 'admin', lastLogin: '2025-11-28 10:30', status: 'active' },
  { id: 'ACC002', name: '运营账号', email: 'ops@example.com', role: 'operator', lastLogin: '2025-11-28 09:15', status: 'active' },
  { id: 'ACC003', name: '查看账号', email: 'viewer@example.com', role: 'viewer', lastLogin: '2025-11-27 16:45', status: 'inactive' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleAddAccount = () => {
  alert('添加账号')
}

const handleEditAccount = (account: typeof accounts.value[0]) => {
  alert(`编辑账号: ${account.name}`)
}

const handleAccountPermission = (account: typeof accounts.value[0]) => {
  alert(`权限管理: ${account.name}`)
}

const handleDisableAccount = (account: typeof accounts.value[0]) => {
  if (confirm(`确定禁用账号 ${account.name} 吗？`)) {
    alert(`已禁用: ${account.name}`)
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '账号管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">账号管理</h1>
          <p class="mt-2 text-gray-600">管理代理商下的账号权限</p>
        </div>
        <button @click="handleAddAccount" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加账号
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">账号总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活跃账号</p>
        <p class="text-2xl font-bold text-green-600 mt-1">25</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">管理员</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">3</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">待激活</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">2</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账号名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">邮箱</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">角色</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">最后登录</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="account in accounts" :key="account.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ account.name }}</div>
              <div class="text-xs text-gray-500">{{ account.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ account.email }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs font-medium',
                     account.role === 'admin' ? 'bg-red-100 text-red-700' :
                     account.role === 'operator' ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700']">
                {{ account.role === 'admin' ? '管理员' : account.role === 'operator' ? '运营' : '查看者' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ account.lastLogin }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     account.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ account.status === 'active' ? '活跃' : '未激活' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button @click="handleEditAccount(account)" class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button @click="handleAccountPermission(account)" class="text-gray-600 hover:text-gray-800 mr-3">权限</button>
              <button @click="handleDisableAccount(account)" class="text-red-600 hover:text-red-800">禁用</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
