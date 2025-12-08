<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 15 })
const searchKeyword = ref('')
const filterType = ref('')

const authList = ref([
  { id: 1, awemeId: '28745612', nickname: '品牌官方号', avatar: '', authType: '授权', advertisers: 5, createdAt: '2025-10-15' },
  { id: 2, awemeId: '38956123', nickname: '达人合作号', avatar: '', authType: '达人', advertisers: 3, createdAt: '2025-10-20' },
  { id: 3, awemeId: '45678912', nickname: '企业蓝V号', avatar: '', authType: '企业', advertisers: 8, createdAt: '2025-09-10' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleSearch = () => {
  alert(`搜索: ${searchKeyword.value}`)
}

const handleDetail = (item: typeof authList.value[0]) => {
  alert(`查看详情: ${item.nickname}`)
}

const handleUnbind = (item: typeof authList.value[0]) => {
  alert(`解除授权: ${item.nickname}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '抖音授权' }]" />
      <h1 class="text-3xl font-bold text-gray-900">抖音授权关系</h1>
      <p class="mt-2 text-gray-600">查看和管理抖音号授权关系</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已授权抖音号</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联广告主</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">16</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">企业号</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">5</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-4 border-b border-gray-200">
        <div class="flex flex-wrap gap-4">
<input v-model="searchKeyword" type="text" placeholder="搜索抖音号" class="flex-1 min-w-[200px] px-4 py-2 border border-gray-300 rounded-lg">
          <select v-model="filterType" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="">全部类型</option>
            <option value="auth">授权号</option>
            <option value="daren">达人号</option>
            <option value="enterprise">企业号</option>
          </select>
<button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          搜索
        </button>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">抖音号</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">授权类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联广告主</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">授权时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in authList" :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 bg-pink-100 rounded-full flex items-center justify-center">
                    <span class="text-pink-600 text-sm">抖</span>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ item.nickname }}</div>
                    <div class="text-xs text-gray-500">ID: {{ item.awemeId }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded text-xs font-medium',
                  item.authType === '企业' ? 'bg-blue-100 text-blue-700' :
                  item.authType === '达人' ? 'bg-purple-100 text-purple-700' : 'bg-green-100 text-green-700']">
                  {{ item.authType }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.advertisers }} 个</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ item.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
<button @click="handleDetail(item)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button @click="handleUnbind(item)" class="text-red-600 hover:text-red-800">解除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
