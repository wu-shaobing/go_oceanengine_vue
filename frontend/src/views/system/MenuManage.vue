<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import { systemApi, type MenuTree } from '@/api'

const loading = ref(true)
const menus = ref<MenuTree[]>([])

// 获取菜单树
const fetchMenus = async () => {
  loading.value = true
  try {
    menus.value = await systemApi.getMenuTree()
  } catch (error) {
    console.error('获取菜单失败:', error)
  } finally {
    loading.value = false
  }
}

// 展开/折叠
const expandedIds = ref<number[]>([])

const toggleExpand = (id: number) => {
  const index = expandedIds.value.indexOf(id)
  if (index >= 0) {
    expandedIds.value.splice(index, 1)
  } else {
    expandedIds.value.push(id)
  }
}

const isExpanded = (id: number) => expandedIds.value.includes(id)

const expandAll = () => {
  const ids: number[] = []
  const traverse = (items: MenuTree[]) => {
    for (const item of items) {
      if (item.children && item.children.length > 0) {
        ids.push(item.id)
        traverse(item.children)
      }
    }
  }
  traverse(menus.value)
  expandedIds.value = ids
}

const collapseAll = () => {
  expandedIds.value = []
}

// 菜单表单
const showMenuDialog = ref(false)
const menuFormMode = ref<'add' | 'edit'>('add')
const menuForm = reactive({
  id: 0,
  parent_id: 0,
  name: '',
  path: '',
  component: '',
  icon: '',
  sort: 0,
  type: 1,
  permission: '',
  status: 1,
  hidden: 0,
  remark: ''
})

const resetMenuForm = () => {
  menuForm.id = 0
  menuForm.parent_id = 0
  menuForm.name = ''
  menuForm.path = ''
  menuForm.component = ''
  menuForm.icon = ''
  menuForm.sort = 0
  menuForm.type = 1
  menuForm.permission = ''
  menuForm.status = 1
  menuForm.hidden = 0
  menuForm.remark = ''
}

const handleCreateMenu = () => {
  menuFormMode.value = 'add'
  resetMenuForm()
  showMenuDialog.value = true
}

const handleAddChild = (parent: MenuTree) => {
  menuFormMode.value = 'add'
  resetMenuForm()
  menuForm.parent_id = parent.id
  showMenuDialog.value = true
}

const handleEditMenu = (menu: MenuTree) => {
  menuFormMode.value = 'edit'
  menuForm.id = menu.id
  menuForm.parent_id = menu.parent_id
  menuForm.name = menu.name
  menuForm.path = menu.path
  menuForm.component = menu.component
  menuForm.icon = menu.icon
  menuForm.sort = menu.sort
  menuForm.type = menu.type
  menuForm.permission = menu.permission
  menuForm.status = menu.status
  menuForm.hidden = menu.hidden
  showMenuDialog.value = true
}

const handleSaveMenu = async () => {
  try {
    if (!menuForm.name) {
      alert('请填写菜单名称')
      return
    }
    if (menuFormMode.value === 'add') {
      await systemApi.createMenu({
        parent_id: menuForm.parent_id || undefined,
        name: menuForm.name,
        path: menuForm.path,
        component: menuForm.component,
        icon: menuForm.icon,
        sort: menuForm.sort,
        type: menuForm.type,
        permission: menuForm.permission,
        status: menuForm.status,
        hidden: menuForm.hidden,
        remark: menuForm.remark
      })
      alert('创建成功')
    } else {
      await systemApi.updateMenu(menuForm.id, {
        id: menuForm.id,
        parent_id: menuForm.parent_id || undefined,
        name: menuForm.name,
        path: menuForm.path,
        component: menuForm.component,
        icon: menuForm.icon,
        sort: menuForm.sort,
        type: menuForm.type,
        permission: menuForm.permission,
        status: menuForm.status,
        hidden: menuForm.hidden,
        remark: menuForm.remark
      })
      alert('更新成功')
    }
    showMenuDialog.value = false
    fetchMenus()
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败')
  }
}

// 删除菜单
const handleDeleteMenu = async (menu: MenuTree) => {
  if (menu.children && menu.children.length > 0) {
    alert('请先删除子菜单')
    return
  }
  if (!confirm(`确定删除菜单「${menu.name}」吗？`)) return
  try {
    await systemApi.deleteMenu(menu.id)
    alert('删除成功')
    fetchMenus()
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// 菜单类型映射
const menuTypeMap: Record<number, string> = {
  1: '目录',
  2: '菜单',
  3: '按钮'
}

onMounted(() => {
  fetchMenus()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '菜单管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">菜单管理</h1>
          <p class="mt-2 text-gray-600">管理系统导航菜单配置</p>
        </div>
        <button @click="handleCreateMenu" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
          </svg>
          新建菜单
        </button>
      </div>
    </div>

    <!-- 菜单树 -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-4 border-b border-gray-200 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-900">菜单配置</h3>
        <div class="flex items-center gap-2">
          <button class="text-sm text-blue-600 hover:text-blue-800" @click="expandAll">展开全部</button>
          <span class="text-gray-300">|</span>
          <button class="text-sm text-blue-600 hover:text-blue-800" @click="collapseAll">折叠全部</button>
        </div>
      </div>
      
      <div v-if="loading" class="p-8 flex items-center justify-center">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <span class="ml-2 text-gray-500">加载中...</span>
      </div>

      <div v-else-if="menus.length === 0" class="p-8 text-center text-gray-500">暂无菜单数据</div>
      
      <div v-else class="divide-y divide-gray-100">
        <template v-for="menu in menus" :key="menu.id">
          <!-- 父级菜单 -->
          <div class="p-4 hover:bg-gray-50 transition-colors">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <button
                  v-if="menu.children?.length"
                  class="w-6 h-6 flex items-center justify-center text-gray-400 hover:text-gray-600"
                  @click="toggleExpand(menu.id)"
                >
                  <svg
                    class="w-4 h-4 transition-transform"
                    :class="{ 'rotate-90': isExpanded(menu.id) }"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                  </svg>
                </button>
                <span v-else class="w-6"></span>
                
                <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
                  </svg>
                </div>
                
                <div>
                  <p class="font-medium text-gray-900">{{ menu.name }}</p>
                  <p class="text-sm text-gray-500">{{ menu.path || '-' }}</p>
                </div>
              </div>
              
              <div class="flex items-center gap-4">
                <span class="text-xs px-2 py-1 rounded bg-gray-100 text-gray-600">{{ menuTypeMap[menu.type] || '未知' }}</span>
                <span class="text-sm text-gray-500">排序: {{ menu.sort }}</span>
                <StatusBadge
                  :status="menu.status === 1 ? 'success' : 'danger'"
                  :text="menu.status === 1 ? '启用' : '禁用'"
                />
                <div class="flex items-center gap-2">
                  <button @click="handleEditMenu(menu)" class="text-sm text-blue-600 hover:text-blue-800">编辑</button>
                  <button @click="handleAddChild(menu)" class="text-sm text-green-600 hover:text-green-800">添加子菜单</button>
                  <button @click="handleDeleteMenu(menu)" class="text-sm text-red-600 hover:text-red-800">删除</button>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 子菜单 -->
          <template v-if="menu.children?.length && isExpanded(menu.id)">
            <div
              v-for="child in menu.children"
              :key="child.id"
              class="p-4 pl-16 bg-gray-50 hover:bg-gray-100 transition-colors border-l-2 border-blue-200"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <button
                    v-if="child.children?.length"
                    class="w-6 h-6 flex items-center justify-center text-gray-400 hover:text-gray-600"
                    @click="toggleExpand(child.id)"
                  >
                    <svg class="w-4 h-4 transition-transform" :class="{ 'rotate-90': isExpanded(child.id) }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                  </button>
                  <div v-else class="w-6 h-6 bg-gray-200 rounded flex items-center justify-center">
                    <svg class="w-3 h-3 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                  </div>
                  <div>
                    <p class="font-medium text-gray-700">{{ child.name }}</p>
                    <p class="text-sm text-gray-500">{{ child.path || '-' }}</p>
                  </div>
                </div>
                
                <div class="flex items-center gap-4">
                  <span class="text-xs px-2 py-1 rounded bg-gray-100 text-gray-600">{{ menuTypeMap[child.type] || '未知' }}</span>
                  <span class="text-sm text-gray-500">排序: {{ child.sort }}</span>
                  <StatusBadge
                    :status="child.status === 1 ? 'success' : 'danger'"
                    :text="child.status === 1 ? '启用' : '禁用'"
                  />
                  <div class="flex items-center gap-2">
                    <button @click="handleEditMenu(child)" class="text-sm text-blue-600 hover:text-blue-800">编辑</button>
                    <button @click="handleAddChild(child)" class="text-sm text-green-600 hover:text-green-800">添加</button>
                    <button @click="handleDeleteMenu(child)" class="text-sm text-red-600 hover:text-red-800">删除</button>
                  </div>
                </div>
              </div>
              <!-- 三级菜单 -->
              <template v-if="child.children?.length && isExpanded(child.id)">
                <div v-for="subChild in child.children" :key="subChild.id" class="mt-2 ml-10 p-3 bg-white rounded border border-gray-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="font-medium text-gray-600">{{ subChild.name }}</p>
                      <p class="text-sm text-gray-400">{{ subChild.permission || subChild.path || '-' }}</p>
                    </div>
                    <div class="flex items-center gap-3">
                      <span class="text-xs px-2 py-1 rounded bg-purple-100 text-purple-600">{{ menuTypeMap[subChild.type] || '未知' }}</span>
                      <button @click="handleEditMenu(subChild)" class="text-sm text-blue-600 hover:text-blue-800">编辑</button>
                      <button @click="handleDeleteMenu(subChild)" class="text-sm text-red-600 hover:text-red-800">删除</button>
                    </div>
                  </div>
                </div>
              </template>
            </div>
          </template>
        </template>
      </div>
    </div>

    <!-- 菜单表单弹窗 -->
    <div v-if="showMenuDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[550px] max-h-[90vh] overflow-y-auto">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">{{ menuFormMode === 'add' ? '新建菜单' : '编辑菜单' }}</h3>
        <div class="space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-1">菜单名称 <span class="text-red-500">*</span></label>
              <input v-model="menuForm.name" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="请输入菜单名称">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">菜单类型</label>
              <select v-model="menuForm.type" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">目录</option>
                <option :value="2">菜单</option>
                <option :value="3">按钮</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
              <select v-model="menuForm.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">启用</option>
                <option :value="0">禁用</option>
              </select>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">路由地址</label>
              <input v-model="menuForm.path" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="如 /system/user">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">组件路径</label>
              <input v-model="menuForm.component" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="如 system/UserManage">
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">图标</label>
              <input v-model="menuForm.icon" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="图标名称">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">排序</label>
              <input v-model.number="menuForm.sort" type="number" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">权限标识</label>
            <input v-model="menuForm.permission" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg" placeholder="如 system:user:list">
          </div>
          <div class="flex items-center gap-4">
            <label class="flex items-center gap-2">
              <input v-model="menuForm.hidden" type="checkbox" :true-value="1" :false-value="0" class="rounded">
              <span class="text-sm text-gray-700">隐藏菜单</span>
            </label>
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showMenuDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSaveMenu">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>
