<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'

const loading = ref(true)

interface Menu {
  id: number
  name: string
  path: string
  icon: string
  sort: number
  status: string
  parentId: number | null
  children?: Menu[]
}

const menus = ref<Menu[]>([])

const fetchMenus = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  menus.value = [
    {
      id: 1,
      name: '仪表盘',
      path: '/dashboard',
      icon: 'home',
      sort: 1,
      status: 'active',
      parentId: null
    },
    {
      id: 2,
      name: '广告主管理',
      path: '/advertiser',
      icon: 'users',
      sort: 2,
      status: 'active',
      parentId: null,
      children: [
        { id: 21, name: '广告主列表', path: '/advertiser/list', icon: '', sort: 1, status: 'active', parentId: 2 },
        { id: 22, name: '广告主详情', path: '/advertiser/detail', icon: '', sort: 2, status: 'active', parentId: 2 }
      ]
    },
    {
      id: 3,
      name: '投放管理',
      path: '/campaign',
      icon: 'chart',
      sort: 3,
      status: 'active',
      parentId: null,
      children: [
        { id: 31, name: '广告计划', path: '/campaign/list', icon: '', sort: 1, status: 'active', parentId: 3 },
        { id: 32, name: '创意管理', path: '/campaign/creative', icon: '', sort: 2, status: 'active', parentId: 3 }
      ]
    },
    {
      id: 4,
      name: '数据报表',
      path: '/report',
      icon: 'document',
      sort: 4,
      status: 'active',
      parentId: null
    },
    {
      id: 5,
      name: '人群管理',
      path: '/audience',
      icon: 'group',
      sort: 5,
      status: 'active',
      parentId: null
    },
    {
      id: 6,
      name: '系统管理',
      path: '/system',
      icon: 'cog',
      sort: 6,
      status: 'active',
      parentId: null,
      children: [
        { id: 61, name: '用户管理', path: '/system/user', icon: '', sort: 1, status: 'active', parentId: 6 },
        { id: 62, name: '角色管理', path: '/system/role', icon: '', sort: 2, status: 'active', parentId: 6 },
        { id: 63, name: '菜单管理', path: '/system/menu', icon: '', sort: 3, status: 'active', parentId: 6 }
      ]
    }
  ]
  
  loading.value = false
}

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

onMounted(() => {
  fetchMenus()
  // 默认展开所有有子菜单的项
  expandedIds.value = [2, 3, 6]
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
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
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
          <button class="text-sm text-blue-600 hover:text-blue-800" @click="expandedIds = menus.filter(m => m.children).map(m => m.id)">
            展开全部
          </button>
          <span class="text-gray-300">|</span>
          <button class="text-sm text-blue-600 hover:text-blue-800" @click="expandedIds = []">
            折叠全部
          </button>
        </div>
      </div>
      
      <div v-if="loading" class="p-8 text-center text-gray-500">
        加载中...
      </div>
      
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
                  <p class="text-sm text-gray-500">{{ menu.path }}</p>
                </div>
              </div>
              
              <div class="flex items-center gap-4">
                <span class="text-sm text-gray-500">排序: {{ menu.sort }}</span>
                <StatusBadge
                  :status="menu.status === 'active' ? 'success' : 'danger'"
                  :text="menu.status === 'active' ? '启用' : '禁用'"
                />
                <div class="flex items-center gap-2">
                  <button class="text-sm text-blue-600 hover:text-blue-800">编辑</button>
                  <button class="text-sm text-green-600 hover:text-green-800" v-if="!menu.children?.length">添加子菜单</button>
                  <button class="text-sm text-red-600 hover:text-red-800">删除</button>
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
                  <div class="w-6 h-6 bg-gray-200 rounded flex items-center justify-center">
                    <svg class="w-3 h-3 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                  </div>
                  <div>
                    <p class="font-medium text-gray-700">{{ child.name }}</p>
                    <p class="text-sm text-gray-500">{{ child.path }}</p>
                  </div>
                </div>
                
                <div class="flex items-center gap-4">
                  <span class="text-sm text-gray-500">排序: {{ child.sort }}</span>
                  <StatusBadge
                    :status="child.status === 'active' ? 'success' : 'danger'"
                    :text="child.status === 'active' ? '启用' : '禁用'"
                  />
                  <div class="flex items-center gap-2">
                    <button class="text-sm text-blue-600 hover:text-blue-800">编辑</button>
                    <button class="text-sm text-red-600 hover:text-red-800">删除</button>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </template>
      </div>
    </div>
  </div>
</template>
