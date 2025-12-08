<script setup lang="ts">
import { ref, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchKeyword = ref('')

const countries = ref([
  { code: 'CN', name: '中国', region: '亚洲', currency: 'CNY', timezone: 'UTC+8' },
  { code: 'US', name: '美国', region: '北美', currency: 'USD', timezone: 'UTC-5~-8' },
  { code: 'JP', name: '日本', region: '亚洲', currency: 'JPY', timezone: 'UTC+9' },
  { code: 'KR', name: '韩国', region: '亚洲', currency: 'KRW', timezone: 'UTC+9' },
  { code: 'GB', name: '英国', region: '欧洲', currency: 'GBP', timezone: 'UTC+0' },
  { code: 'DE', name: '德国', region: '欧洲', currency: 'EUR', timezone: 'UTC+1' },
  { code: 'FR', name: '法国', region: '欧洲', currency: 'EUR', timezone: 'UTC+1' },
  { code: 'AU', name: '澳大利亚', region: '大洋洲', currency: 'AUD', timezone: 'UTC+10' },
  { code: 'SG', name: '新加坡', region: '亚洲', currency: 'SGD', timezone: 'UTC+8' },
  { code: 'TH', name: '泰国', region: '亚洲', currency: 'THB', timezone: 'UTC+7' },
  { code: 'VN', name: '越南', region: '亚洲', currency: 'VND', timezone: 'UTC+7' },
  { code: 'ID', name: '印度尼西亚', region: '亚洲', currency: 'IDR', timezone: 'UTC+7~+9' }
])

const filteredCountries = computed(() => {
  if (!searchKeyword.value) return countries.value
  const keyword = searchKeyword.value.toLowerCase()
  return countries.value.filter(c =>
    c.name.toLowerCase().includes(keyword) ||
    c.code.toLowerCase().includes(keyword) ||
    c.region.toLowerCase().includes(keyword)
  )
})

const handleCopyCode = (code: string) => {
  navigator.clipboard.writeText(code)
  alert(`已复制代码: ${code}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '国家信息' }]" />
      <h1 class="text-3xl font-bold text-gray-900">国家/区域信息</h1>
      <p class="mt-2 text-gray-600">查询支持投放的国家和区域信息</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <input v-model="searchKeyword" type="text" placeholder="搜索国家名称、代码或区域..."
             class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">国家代码</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">国家名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">所属区域</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">货币</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时区</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="country in filteredCountries" :key="country.code" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono font-medium text-gray-900">{{ country.code }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ country.name }}</td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ country.region }}</span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-700">{{ country.currency }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ country.timezone }}</td>
              <td class="px-6 py-4">
<button @click="handleCopyCode(country.code)" class="text-blue-600 hover:text-blue-800 text-sm">复制代码</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
