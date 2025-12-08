<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '服务市场', path: '/servemarket' }, { name: '工作台' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">服务市场</h1>
      <p class="text-gray-600 mt-1">一站式营销服务平台</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已订阅服务</div>
        <div class="text-2xl font-bold mt-1">{{ stats.subscribed }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">进行中订单</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.orders }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">功能点余额</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.points.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">累计消耗</div>
        <div class="text-2xl font-bold mt-1">¥{{ stats.spent.toLocaleString() }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">已订阅服务</h3>
          <router-link to="/servemarket/subscribe" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <div class="space-y-3">
          <div v-for="service in subscribedServices" :key="service.id" class="flex items-center justify-between p-3 bg-gray-50 rounded">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-blue-100 rounded flex items-center justify-center mr-3">
                <span class="text-blue-600">{{ service.icon }}</span>
              </div>
              <div>
                <div class="font-medium">{{ service.name }}</div>
                <div class="text-sm text-gray-500">到期: {{ service.expireDate }}</div>
              </div>
            </div>
            <span class="px-2 py-1 text-xs bg-green-100 text-green-800 rounded">使用中</span>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">最近订单</h3>
          <router-link to="/servemarket/order" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <div class="space-y-3">
          <div v-for="order in recentOrders" :key="order.id" class="flex items-center justify-between p-3 bg-gray-50 rounded">
            <div>
              <div class="font-medium">{{ order.name }}</div>
              <div class="text-sm text-gray-500">{{ order.time }}</div>
            </div>
            <span :class="order.status === 'completed' ? 'text-green-600' : 'text-blue-600'" class="text-sm">
              {{ order.status === 'completed' ? '已完成' : '进行中' }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">热门服务推荐</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <div v-for="service in hotServices" :key="service.id" class="border rounded-lg p-4 hover:shadow-md transition-shadow">
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-3">
            <span class="text-2xl">{{ service.icon }}</span>
          </div>
          <h4 class="font-medium mb-1">{{ service.name }}</h4>
          <p class="text-sm text-gray-500 mb-3">{{ service.desc }}</p>
          <div class="flex justify-between items-center">
            <span class="text-blue-600 font-medium">¥{{ service.price }}/月</span>
            <button @click="handleSubscribe(service)" class="px-3 py-1 text-sm text-blue-600 border border-blue-600 rounded hover:bg-blue-50">订阅</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const stats = ref({ subscribed: 5, orders: 3, points: 12500, spent: 25680 })

const handleSubscribe = (service: typeof hotServices.value[0]) => {
  alert(`订阅服务: ${service.name}`)
}

const subscribedServices = ref([
  { id: 1, name: '智能投放助手', icon: '🤖', expireDate: '2024-12-31' },
  { id: 2, name: '创意素材生成', icon: '🎨', expireDate: '2024-09-15' },
  { id: 3, name: '数据分析Pro', icon: '📊', expireDate: '2024-08-20' }
])

const recentOrders = ref([
  { id: 1, name: '视频制作服务', time: '2024-06-18', status: 'processing' },
  { id: 2, name: '账户诊断服务', time: '2024-06-15', status: 'completed' },
  { id: 3, name: '代运营服务', time: '2024-06-10', status: 'processing' }
])

const hotServices = ref([
  { id: 1, name: '智能投放', icon: '🚀', desc: 'AI智能优化投放效果', price: 299 },
  { id: 2, name: '素材制作', icon: '🎬', desc: '专业视频素材制作', price: 199 },
  { id: 3, name: '数据分析', icon: '📈', desc: '深度数据洞察分析', price: 399 },
  { id: 4, name: '代运营', icon: '👨‍💼', desc: '专业团队代运营', price: 999 }
])
</script>
