<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import DataTable from '@/components/common/DataTable.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import { systemApi, type Role, type MenuTree } from '@/api'

const loading = ref(true)
const roles = ref<Role[]>([])
const menus = ref<MenuTree[]>([])

const columns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'name', title: '角色名称' },
  { key: 'key', title: '角色标识' },
  { key: 'remark', title: '描述' },
  { key: 'sort', title: '排序', width: 80 },
  { key: 'status', title: '状态', width: 100 },
  { key: 'actions', title: '操作', width: 200, align: 'center' as const }
]

// 获取角色列表
const fetchRoles = async () => {
  loading.value = true
  try {
    const res = await systemApi.getRoleList({ page: 1, page_size: 100 })
    roles.value = res.list || []
  } catch (error) {
    console.error('获取角色列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取菜单树
const fetchMenuTree = async () => {
  try {
    menus.value = await systemApi.getMenuTree()
  } catch (error) {
    console.error('获取菜单树失败:', error)
  }
}

// 角色表单弹窗
const showRoleDialog = ref(false)
const roleFormMode = ref<'add' | 'edit'>('add')
const roleForm = reactive({
  id: 0,
  name: '',
  code: '',
  sort: 0,
  status: 1,
  remark: ''
})

const resetRoleForm = () => {
  roleForm.id = 0
  roleForm.name = ''
  roleForm.code = ''
  roleForm.sort = 0
  roleForm.status = 1
  roleForm.remark = ''
}

const handleCreateRole = () => {
  roleFormMode.value = 'add'
  resetRoleForm()
  showRoleDialog.value = true
}

const handleEditRole = (role: Role) => {
  roleFormMode.value = 'edit'
  roleForm.id = role.id
  roleForm.name = role.name
  roleForm.code = role.key
  roleForm.sort = role.sort
  roleForm.status = role.status
  roleForm.remark = role.remark
  showRoleDialog.value = true
}

const handleSaveRole = async () => {
  try {
    if (!roleForm.name || !roleForm.code) {
      alert('请填写必填字段')
      return
    }
    if (roleFormMode.value === 'add') {
      await systemApi.createRole({
        name: roleForm.name,
        code: roleForm.code,
        sort: roleForm.sort,
        status: roleForm.status,
        remark: roleForm.remark
      })
      alert('创建成功')
    } else {
      await systemApi.updateRole(roleForm.id, {
        id: roleForm.id,
        name: roleForm.name,
        code: roleForm.code,
        sort: roleForm.sort,
        status: roleForm.status,
        remark: roleForm.remark
      })
      alert('更新成功')
    }
    showRoleDialog.value = false
    fetchRoles()
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败')
  }
}

// 删除角色
const handleDeleteRole = async (role: Role) => {
  if (!confirm(`确定删除角色 ${role.name}?`)) return
  try {
    await systemApi.deleteRole(role.id)
    alert('删除成功')
    fetchRoles()
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// 权限配置弹窗
const showPermDialog = ref(false)
const permRoleId = ref(0)
const permRoleName = ref('')
const selectedMenuIds = ref<number[]>([])

const handlePermissionConfig = async (role: Role) => {
  permRoleId.value = role.id
  permRoleName.value = role.name
  try {
    selectedMenuIds.value = await systemApi.getRoleMenus(role.id)
  } catch (error) {
    console.error('获取角色菜单失败:', error)
    selectedMenuIds.value = []
  }
  showPermDialog.value = true
}

const toggleMenuSelect = (menuId: number) => {
  const index = selectedMenuIds.value.indexOf(menuId)
  if (index === -1) {
    selectedMenuIds.value.push(menuId)
  } else {
    selectedMenuIds.value.splice(index, 1)
  }
}

const handleSavePermission = async () => {
  try {
    await systemApi.updateRoleMenus(permRoleId.value, selectedMenuIds.value)
    alert('权限保存成功')
    showPermDialog.value = false
  } catch (error) {
    console.error('保存权限失败:', error)
    alert('保存权限失败')
  }
}

// 递归获取所有菜单 ID
const getAllMenuIds = (menuList: MenuTree[]): number[] => {
  const ids: number[] = []
  const traverse = (items: MenuTree[]) => {
    for (const item of items) {
      ids.push(item.id)
      if (item.children && item.children.length > 0) {
        traverse(item.children)
      }
    }
  }
  traverse(menuList)
  return ids
}

const handleSelectAll = () => {
  selectedMenuIds.value = getAllMenuIds(menus.value)
}

const handleClearAll = () => {
  selectedMenuIds.value = []
}

onMounted(() => {
  fetchRoles()
  fetchMenuTree()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '角色管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">角色管理</h1>
          <p class="mt-2 text-gray-600">管理系统角色和权限分配，共 {{ roles.length }} 个角色</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2" @click="handleCreateRole">
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
            <code class="text-sm text-gray-500 mt-1 bg-gray-100 px-1 rounded">{{ role.key }}</code>
          </div>
          <StatusBadge
            :status="role.status === 1 ? 'success' : 'danger'"
            :text="role.status === 1 ? '启用' : '禁用'"
          />
        </div>
        <p class="mt-3 text-sm text-gray-600">{{ role.remark || '暂无描述' }}</p>
        <div class="mt-4 flex items-center justify-between">
          <span class="text-sm text-gray-500">排序: {{ role.sort }}</span>
          <div class="flex items-center gap-2">
            <button class="text-sm text-blue-600 hover:text-blue-800" @click="handlePermissionConfig(role)">权限配置</button>
            <button class="text-sm text-gray-600 hover:text-gray-800" @click="handleEditRole(role)">编辑</button>
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

        <template #status="{ row }">
          <StatusBadge
            :status="row.status === 1 ? 'success' : 'danger'"
            :text="row.status === 1 ? '启用' : '禁用'"
          />
        </template>

        <template #actions="{ row }">
          <div class="flex items-center justify-center gap-2">
            <button class="text-blue-600 hover:text-blue-800 text-sm" @click="handlePermissionConfig(row)">权限</button>
            <button class="text-gray-600 hover:text-gray-800 text-sm" @click="handleEditRole(row)">编辑</button>
            <button class="text-red-600 hover:text-red-800 text-sm" @click="handleDeleteRole(row)">删除</button>
          </div>
        </template>
      </DataTable>
    </div>

    <!-- 角色表单弹窗 -->
    <div v-if="showRoleDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[450px]">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">{{ roleFormMode === 'add' ? '新建角色' : '编辑角色' }}</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">角色名称 <span class="text-red-500">*</span></label>
            <input v-model="roleForm.name" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入角色名称">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">角色标识 <span class="text-red-500">*</span></label>
            <input v-model="roleForm.code" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="如 admin, operator">
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">排序</label>
              <input v-model.number="roleForm.sort" type="number" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
              <select v-model="roleForm.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">启用</option>
                <option :value="0">禁用</option>
              </select>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">描述</label>
            <textarea v-model="roleForm.remark" rows="3" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="角色描述"></textarea>
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showRoleDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSaveRole">保存</button>
        </div>
      </div>
    </div>

    <!-- 权限配置弹窗 -->
    <div v-if="showPermDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[600px] max-h-[80vh] overflow-hidden flex flex-col">
        <h3 class="text-lg font-semibold text-gray-900 mb-2">权限配置 - {{ permRoleName }}</h3>
        <div class="flex gap-2 mb-4">
          <button class="px-3 py-1 text-sm border border-blue-600 text-blue-600 rounded hover:bg-blue-50" @click="handleSelectAll">全选</button>
          <button class="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50" @click="handleClearAll">清空</button>
        </div>
        <div class="flex-1 overflow-y-auto border border-gray-200 rounded-lg p-4">
          <div v-if="menus.length === 0" class="text-center text-gray-500 py-8">暂无菜单数据</div>
          <template v-else>
            <div v-for="menu in menus" :key="menu.id" class="mb-4">
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="checkbox" :checked="selectedMenuIds.includes(menu.id)" @change="toggleMenuSelect(menu.id)" class="rounded">
                <span class="font-medium">{{ menu.name }}</span>
              </label>
              <div v-if="menu.children && menu.children.length" class="ml-6 mt-2 space-y-2">
                <div v-for="child in menu.children" :key="child.id">
                  <label class="flex items-center gap-2 cursor-pointer">
                    <input type="checkbox" :checked="selectedMenuIds.includes(child.id)" @change="toggleMenuSelect(child.id)" class="rounded">
                    <span>{{ child.name }}</span>
                  </label>
                  <div v-if="child.children && child.children.length" class="ml-6 mt-1 space-y-1">
                    <label v-for="subChild in child.children" :key="subChild.id" class="flex items-center gap-2 cursor-pointer text-sm text-gray-600">
                      <input type="checkbox" :checked="selectedMenuIds.includes(subChild.id)" @change="toggleMenuSelect(subChild.id)" class="rounded">
                      <span>{{ subChild.name }}</span>
                    </label>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </div>
        <div class="flex justify-end gap-2 mt-4">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showPermDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSavePermission">保存权限</button>
        </div>
      </div>
    </div>
  </div>
</template>
