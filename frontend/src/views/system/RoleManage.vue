<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import DataTable from '@/components/common/DataTable.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'

const loading = ref(true)

interface Role {
  id: number
  name: string
  key: string
  description: string
  userCount: number
  status: string
  created_at: string
}

const roles = ref<Role[]>([])

const columns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'name', title: '角色名称' },
  { key: 'key', title: '角色标识' },
  { key: 'description', title: '描述' },
  { key: 'userCount', title: '用户数', width: 100, align: 'center' as const },
  { key: 'status', title: '状态', width: 100 },
  { key: 'actions', title: '操作', width: 150, align: 'center' as const }
]

const fetchRoles = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  roles.value = [
    { id: 1, name: '超级管理员', key: 'super_admin', description: '拥有所有权限', userCount: 2, status: 'active', created_at: '2024-01-01' },
    { id: 2, name: '管理员', key: 'admin', description: '管理广告主和投放', userCount: 5, status: 'active', created_at: '2024-01-01' },
    { id: 3, name: '运营', key: 'operator', description: '日常运营操作', userCount: 12, status: 'active', created_at: '2024-01-05' },
    { id: 4, name: '观察者', key: 'viewer', description: '只读权限', userCount: 8, status: 'active', created_at: '2024-01-10' },
    { id: 5, name: '测试角色', key: 'test', description: '测试用角色', userCount: 0, status: 'disabled', created_at: '2024-01-15' }
  ]
  
  loading.value = false
}

onMounted(fetchRoles)
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '角色管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">角色管理</h1>
          <p class="mt-2 text-gray-600">管理系统角色和权限分配</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
          </svg>
          新建角色
        </button>
      </div>
    </div>

    <!-- 角色卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div
        v-for="role in roles"
        :key="role.id"
        class="bg-white rounded-lg border border-gray-200 p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between">
          <div>
            <h3 class="text-lg font-semibold text-gray-900">{{ role.name }}</h3>
            <p class="text-sm text-gray-500 mt-1">{{ role.key }}</p>
          </div>
          <StatusBadge
            :status="role.status === 'active' ? 'success' : 'danger'"
            :text="role.status === 'active' ? '启用' : '禁用'"
          />
        </div>
        <p class="mt-3 text-sm text-gray-600">{{ role.description }}</p>
        <div class="mt-4 flex items-center justify-between">
          <span class="text-sm text-gray-500">
            <span class="font-medium text-gray-900">{{ role.userCount }}</span> 位用户
          </span>
          <div class="flex items-center gap-2">
            <button class="text-sm text-blue-600 hover:text-blue-800">权限配置</button>
            <button class="text-sm text-gray-600 hover:text-gray-800">编辑</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 角色列表 -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">角色列表</h3>
      </div>
      <DataTable :columns="columns" :data="roles" :loading="loading">
        <template #name="{ value }">
          <span class="font-medium text-gray-900">{{ value }}</span>
        </template>

        <template #key="{ value }">
          <code class="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded">{{ value }}</code>
        </template>

        <template #userCount="{ value }">
          <span class="text-gray-600">{{ value }}</span>
        </template>

        <template #status="{ value }">
          <StatusBadge
            :status="value === 'active' ? 'success' : 'danger'"
            :text="value === 'active' ? '启用' : '禁用'"
          />
        </template>

        <template #actions>
          <div class="flex items-center justify-center gap-2">
            <button class="text-blue-600 hover:text-blue-800">权限</button>
            <button class="text-gray-600 hover:text-gray-800">编辑</button>
            <button class="text-red-600 hover:text-red-800">删除</button>
          </div>
        </template>
      </DataTable>
    </div>
  </div>
</template>
