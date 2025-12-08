<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const stats = reactive({
  totalAdvertisers: 128,
  activeAdvertisers: 95,
  todaySpend: 1258600,
  monthSpend: 32580000,
  balance: 5680000,
  pendingTasks: 23
})

const recentAdvertisers = ref([
  { id: 1234567, name: '北京科技有限公司', initial: '北', color: 'blue', todaySpend: 45680, balance: 1250000, activeAds: 156, status: '正常' },
  { id: 1234568, name: '上海网络公司', initial: '上', color: 'green', todaySpend: 38920, balance: 850000, activeAds: 89, status: '正常' },
  { id: 1234569, name: '深圳电商公司', initial: '深', color: 'purple', todaySpend: 52340, balance: 180000, activeAds: 234, status: '余额不足' }
])

const pendingTasks = ref([
  { type: 'warning', icon: 'alert', title: '5个广告主账户余额不足', desc: '需要及时充值以保证广告正常投放', action: '立即处理' },
  { type: 'info', icon: 'document', title: '3个新广告主开户申请待审核', desc: '请及时审核新客户资料', action: '去审核' }
])

const handleViewAdvertiser = (adv: typeof recentAdvertisers.value[0]) => {
  alert(`查看广告主: ${adv.name}`)
}

const handleAdvAction = (adv: typeof recentAdvertisers.value[0]) => {
  if (adv.status === '余额不足') {
    alert(`为 ${adv.name} 充值`)
  } else {
    alert(`管理广告主: ${adv.name}`)
  }
}

const handleTaskAction = (task: typeof pendingTasks.value[0]) => {
  alert(`处理任务: ${task.title}`)
}

onMounted(() => {
  // Load data
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '代理商管理' }, { name: '代理商工作台' }]" />
      <h1 class="text-3xl font-bold text-gray-900">代理商工作台</h1>
      <p class="mt-2 text-gray-600">查看代理商数据概览和管理广告主</p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
        <p class="text-sm text-blue-100 mb-1">管理广告主</p>
        <p class="text-3xl font-bold">{{ stats.totalAdvertisers }}</p>
        <div class="mt-3 flex items-center text-sm">
          <span class="text-green-300">{{ stats.activeAdvertisers }} 个活跃</span>
        </div>
      </div>
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-lg p-6 text-white">
        <p class="text-sm text-green-100 mb-1">今日消耗</p>
        <p class="text-3xl font-bold">¥{{ (stats.todaySpend / 10000).toFixed(1) }}万</p>
        <div class="mt-3 flex items-center text-sm">
          <span>本月累计: ¥{{ (stats.monthSpend / 10000).toFixed(0) }}万</span>
        </div>
      </div>
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-lg p-6 text-white">
        <p class="text-sm text-purple-100 mb-1">账户余额</p>
        <p class="text-3xl font-bold">¥{{ (stats.balance / 10000).toFixed(0) }}万</p>
        <div class="mt-3 flex items-center text-sm">
          <span>可用余额充足</span>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">快捷操作</h2>
      <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
        <router-link to="/agent/advertisers" class="flex flex-col items-center gap-3 p-4 rounded-lg border-2 border-gray-200 hover:border-blue-500 hover:bg-blue-50 transition-all">
          <div class="p-3 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">新增广告主</span>
        </router-link>
        <router-link to="/agent/recharge" class="flex flex-col items-center gap-3 p-4 rounded-lg border-2 border-gray-200 hover:border-green-500 hover:bg-green-50 transition-all">
          <div class="p-3 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">账户充值</span>
        </router-link>
        <router-link to="/reports" class="flex flex-col items-center gap-3 p-4 rounded-lg border-2 border-gray-200 hover:border-purple-500 hover:bg-purple-50 transition-all">
          <div class="p-3 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">查看报表</span>
        </router-link>
        <router-link to="/system/users" class="flex flex-col items-center gap-3 p-4 rounded-lg border-2 border-gray-200 hover:border-orange-500 hover:bg-orange-50 transition-all">
          <div class="p-3 bg-orange-100 rounded-lg">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">系统设置</span>
        </router-link>
      </div>
    </div>

    <!-- Recent Advertisers Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900">最近活跃广告主</h2>
        <router-link to="/agent/advertisers" class="text-sm text-blue-600 hover:text-blue-700 font-medium">
          查看全部 →
        </router-link>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">广告主</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">今日消耗</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">账户余额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">活跃广告</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">操作</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="adv in recentAdvertisers" :key="adv.id" class="hover:bg-blue-50 transition-colors">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div :class="['flex-shrink-0 h-10 w-10 rounded-full flex items-center justify-center', `bg-${adv.color}-100`]">
                    <span :class="[`text-${adv.color}-600`, 'font-medium text-sm']">{{ adv.initial }}</span>
                  </div>
                  <div class="ml-4">
                    <div class="text-sm font-medium text-gray-900">{{ adv.name }}</div>
                    <div class="text-sm text-gray-500">ID: {{ adv.id }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">¥{{ adv.todaySpend.toLocaleString() }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm" :class="adv.balance < 200000 ? 'text-yellow-600 font-medium' : 'text-gray-900'">¥{{ adv.balance.toLocaleString() }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ adv.activeAds }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="['inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium', adv.status === '正常' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                  {{ adv.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm">
                <button @click="handleViewAdvertiser(adv)" class="text-blue-600 hover:text-blue-900 mr-3">查看</button>
                <button @click="handleAdvAction(adv)" :class="adv.status === '余额不足' ? 'text-orange-600 hover:text-orange-900' : 'text-gray-600 hover:text-gray-900'">
                  {{ adv.status === '余额不足' ? '充值' : '管理' }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Pending Tasks -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">待处理事项</h2>
      <div class="space-y-3">
        <div v-for="(task, index) in pendingTasks" :key="index" 
             :class="['flex items-center justify-between p-4 border rounded-lg', task.type === 'warning' ? 'bg-orange-50 border-orange-200' : 'bg-blue-50 border-blue-200']">
          <div class="flex items-center gap-3">
            <div :class="['p-2 rounded-lg', task.type === 'warning' ? 'bg-orange-100' : 'bg-blue-100']">
              <svg :class="['w-5 h-5', task.type === 'warning' ? 'text-orange-600' : 'text-blue-600']" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="task.icon === 'alert'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
              </svg>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-900">{{ task.title }}</p>
              <p class="text-xs text-gray-600 mt-1">{{ task.desc }}</p>
            </div>
          </div>
          <button @click="handleTaskAction(task)" :class="['px-4 py-2 text-white text-sm rounded-lg transition-colors', task.type === 'warning' ? 'bg-orange-600 hover:bg-orange-700' : 'bg-blue-600 hover:bg-blue-700']">
            {{ task.action }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
