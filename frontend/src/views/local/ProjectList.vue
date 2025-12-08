<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '项目管理' }, { name: '项目列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">项目列表</h1>
        <p class="text-gray-600 mt-1">管理本地推广项目</p>
      </div>
      <router-link to="/local/project/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建项目
      </router-link>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="项目名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="running">投放中</option>
          <option value="pause">已暂停</option>
        </select>
        <select v-model="filters.industry" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部行业</option>
          <option value="food">餐饮</option>
          <option value="beauty">丽人</option>
          <option value="education">教育</option>
          <option value="auto">汽车</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 项目列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">项目信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">行业</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">日预算</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索数</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">线索成本</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="project in projects" :key="project.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ project.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ project.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ project.industry }}</td>
            <td class="px-4 py-3">
              <span :class="project.status === 'running' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" 
                class="px-2 py-1 text-xs rounded-full">
                {{ project.status === 'running' ? '投放中' : '已暂停' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ project.budget.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ project.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ project.leads }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ project.leadCost }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <router-link :to="`/local/project/${project.id}`" class="text-blue-600 hover:text-blue-800 text-sm">详情</router-link>
                <button @click="handleEdit(project)" class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button v-if="project.status === 'running'" @click="handlePause(project)" class="text-orange-600 hover:text-orange-800 text-sm">暂停</button>
                <button v-else @click="handleStart(project)" class="text-green-600 hover:text-green-800 text-sm">启动</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="50" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { localApi, type LocalProject } from '@/api/local'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()
const loading = ref(false)
const error = ref('')

const filters = ref({
  keyword: '',
  status: '',
  industry: ''
})

const pagination = ref({
  page: 1,
  pageSize: 20,
  total: 0
})

interface ProjectItem {
  id: string
  name: string
  industry: string
  status: string
  budget: number
  cost: number
  leads: number
  leadCost: number
}

const projects = ref<ProjectItem[]>([])

const fetchProjects = async () => {
  const advertiserId = advertiserStore.currentAdvertiserId
  if (!advertiserId) {
    error.value = '请先选择广告主账户'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const res = await localApi.getProjectList({
      advertiser_id: advertiserId,
      page: pagination.value.page,
      page_size: pagination.value.pageSize,
      status: filters.value.status || undefined
    })

    if (res) {
      const data = res as any
      projects.value = (data.list || []).map((item: LocalProject) => ({
        id: String(item.project_id),
        name: item.project_name,
        industry: '',
        status: item.opt_status === 'enable' ? 'running' : 'pause',
        budget: item.budget || 0,
        cost: 0,
        leads: 0,
        leadCost: 0
      }))
      pagination.value.total = data.total || 0
    }
  } catch (e: any) {
    error.value = e.message || '获取项目列表失败'
    console.error('Failed to fetch projects:', e)
  } finally {
    loading.value = false
  }
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const handleSearch = () => {
  pagination.value.page = 1
  fetchProjects()
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars  
const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchProjects()
}

// 导出以便在模板中使用
defineExpose({ handleSearch, handlePageChange, loading, error })

onMounted(() => {
  fetchProjects()
})

watch(() => advertiserStore.currentAdvertiserId, () => {
  fetchProjects()
})

const handleEdit = (project: ProjectItem) => {
  alert(`编辑项目: ${project.name}`)
}

const handlePause = async (project: ProjectItem) => {
  if (confirm(`确定暂停项目「${project.name}」吗？`)) {
    project.status = 'pause'
    alert('项目已暂停')
  }
}

const handleStart = async (project: ProjectItem) => {
  project.status = 'running'
  alert(`项目「${project.name}」已启动`)
}
</script>
