<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

interface MenuItem {
  name: string
  path: string
  icon: string
}

interface MenuSection {
  title: string
  items: MenuItem[]
}

interface ProductLine {
  id: string
  name: string
  color: string
  icon: string
  sections: MenuSection[]
}

// 产品线定义
const productLines: ProductLine[] = [
  {
    id: 'overview',
    name: '概览',
    color: 'blue',
    icon: 'home',
    sections: [
      {
        title: '',
        items: [
          { name: '工作台', path: '/', icon: 'home' }
        ]
      }
    ]
  },
  {
    id: 'qianchuan',
    name: '巨量千川',
    color: 'orange',
    icon: 'shop',
    sections: [
      {
        title: '推广管理',
        items: [
          { name: '千川工作台', path: '/qianchuan', icon: 'shop' },
          { name: '广告计划', path: '/qianchuan/campaign', icon: 'campaign' },
          { name: '广告管理', path: '/qianchuan/ad', icon: 'ad' }
        ]
      },
      {
        title: '电商资产',
        items: [
          { name: '商品管理', path: '/qianchuan/product', icon: 'list' },
          { name: '全域推广', path: '/qianchuan/uni', icon: 'globe' },
          { name: '随心推', path: '/qianchuan/aweme-order', icon: 'heart' }
        ]
      }
    ]
  },
  {
    id: 'local',
    name: '本地推',
    color: 'green',
    icon: 'location',
    sections: [
      {
        title: '推广管理',
        items: [
          { name: '本地推工作台', path: '/local', icon: 'location' },
          { name: '项目管理', path: '/local/project', icon: 'list' },
          { name: '推广管理', path: '/local/promotion', icon: 'campaign' }
        ]
      },
      {
        title: '线索转化',
        items: [
          { name: '线索管理', path: '/local/clue', icon: 'audience' }
        ]
      }
    ]
  },
  {
    id: 'enterprise',
    name: '企业号',
    color: 'purple',
    icon: 'building',
    sections: [
      {
        title: '内容运营',
        items: [
          { name: '企业号工作台', path: '/enterprise', icon: 'building' },
          { name: '评论管理', path: '/enterprise/comment', icon: 'chat' },
          { name: '视频列表', path: '/enterprise/item', icon: 'media' }
        ]
      },
      {
        title: '数据分析',
        items: [
          { name: '数据概览', path: '/enterprise/overview', icon: 'chart' }
        ]
      }
    ]
  },
  {
    id: 'star',
    name: '巨量星图',
    color: 'yellow',
    icon: 'star',
    sections: [
      {
        title: '达人营销',
        items: [
          { name: '星图工作台', path: '/star', icon: 'star' },
          { name: '任务管理', path: '/star/task', icon: 'list' },
          { name: '需求管理', path: '/star/demand', icon: 'campaign' }
        ]
      },
      {
        title: '财务数据',
        items: [
          { name: '资金管理', path: '/star/fund', icon: 'wallet' },
          { name: '数据报表', path: '/star/report', icon: 'chart' }
        ]
      }
    ]
  },
  {
    id: 'tools',
    name: '工具与服务',
    color: 'gray',
    icon: 'tools',
    sections: [
      {
        title: '代理商管理',
        items: [
          { name: '代理商总览', path: '/agent', icon: 'agent' },
          { name: '广告主管理', path: '/agent/advertisers', icon: 'building' },
          { name: '充值管理', path: '/agent/recharge', icon: 'wallet' },
          { name: '转账记录', path: '/agent/transfer', icon: 'transfer' }
        ]
      },
      {
        title: '财务管理',
        items: [
          { name: '日流水统计', path: '/wallet', icon: 'chart' },
          { name: '流水明细', path: '/wallet/transactions', icon: 'list' },
          { name: '钱包转账', path: '/wallet/transfer', icon: 'transfer' }
        ]
      },
      {
        title: '广告资产',
        items: [
          { name: '广告主', path: '/advertisers', icon: 'building' },
          { name: '广告计划', path: '/campaigns', icon: 'campaign' },
          { name: '广告', path: '/ads', icon: 'ad' }
        ]
      },
      {
        title: '素材资产',
        items: [
          { name: '创意', path: '/creatives', icon: 'creative' },
          { name: '媒体库', path: '/media', icon: 'media' }
        ]
      },
      {
        title: '基础工具',
        items: [
          { name: '地域定向', path: '/tools/region', icon: 'location' },
          { name: '行业分类', path: '/tools/industry', icon: 'category' },
          { name: '接口配额', path: '/tools/quota', icon: 'quota' },
          { name: '人群包', path: '/audiences', icon: 'audience' }
        ]
      },
      {
        title: '数据分析',
        items: [
          { name: '数据报表', path: '/reports', icon: 'chart' }
        ]
      }
    ]
  }
]

// 展开/折叠状态
const expandedProducts = ref<Set<string>>(new Set(['overview', 'qianchuan', 'local', 'enterprise', 'star', 'tools']))

const toggleProduct = (id: string) => {
  if (expandedProducts.value.has(id)) {
    expandedProducts.value.delete(id)
  } else {
    expandedProducts.value.add(id)
  }
}

const isExpanded = (id: string) => expandedProducts.value.has(id)

const isActive = (path: string) => {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}

</script>

<template>
  <aside class="w-64 bg-white border-r border-gray-200 min-h-[calc(100vh-64px)] sticky top-16 overflow-y-auto">
    <nav class="py-3">
      <!-- 产品线循环 -->
      <div v-for="product in productLines" :key="product.id" class="mb-1">
        <!-- 产品线标题 (可折叠) -->
        <button
          v-if="product.id !== 'overview'"
          @click="toggleProduct(product.id)"
          :class="[
            'w-full flex items-center justify-between px-4 py-2.5 text-left transition-colors',
            'hover:bg-gray-50',
            isExpanded(product.id) ? 'border-l-2 border-blue-500' : 'border-l-2 border-transparent'
          ]"
        >
          <div class="flex items-center gap-2">
            <span :class="['w-2 h-2 rounded-full', `bg-${product.color}-500`]"></span>
            <span class="text-sm font-semibold text-gray-800">{{ product.name }}</span>
          </div>
          <svg
            :class="['w-4 h-4 text-gray-400 transition-transform', isExpanded(product.id) ? 'rotate-180' : '']"
            fill="none" stroke="currentColor" viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
          </svg>
        </button>

        <!-- 产品线内容 -->
        <div v-show="product.id === 'overview' || isExpanded(product.id)" class="mt-1">
          <div v-for="section in product.sections" :key="section.title" class="mb-3">
            <!-- 二级分组标题 -->
            <h4 v-if="section.title" class="px-4 py-1.5 text-xs font-medium text-gray-400 uppercase tracking-wide">
              {{ section.title }}
            </h4>
            <!-- 菜单项 -->
            <ul class="space-y-0.5">
              <li v-for="item in section.items" :key="item.path">
                <router-link
                  :to="item.path"
                  :class="[
                    'group flex items-center gap-3 mx-2 px-3 py-2 rounded-md transition-all text-sm',
                    isActive(item.path)
                      ? 'bg-blue-600 text-white font-medium shadow-sm'
                      : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
                  ]"
                >
                  <!-- Home Icon -->
                  <svg v-if="item.icon === 'home'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
                  </svg>
                  <!-- Building Icon -->
                  <svg v-else-if="item.icon === 'building'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"/>
                  </svg>
                  <!-- Campaign Icon -->
                  <svg v-else-if="item.icon === 'campaign'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"/>
                  </svg>
                  <!-- Ad Icon -->
                  <svg v-else-if="item.icon === 'ad'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"/>
                  </svg>
                  <!-- Creative Icon -->
                  <svg v-else-if="item.icon === 'creative'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"/>
                  </svg>
                  <!-- Media Icon -->
                  <svg v-else-if="item.icon === 'media'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                  </svg>
                  <!-- Audience Icon -->
                  <svg v-else-if="item.icon === 'audience'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
                  </svg>
                  <!-- Chart Icon -->
                  <svg v-else-if="item.icon === 'chart'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                  </svg>
                  <!-- Agent Icon -->
                  <svg v-else-if="item.icon === 'agent'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                  </svg>
                  <!-- Wallet Icon -->
                  <svg v-else-if="item.icon === 'wallet'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/>
                  </svg>
                  <!-- Transfer Icon -->
                  <svg v-else-if="item.icon === 'transfer'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"/>
                  </svg>
                  <!-- List Icon -->
                  <svg v-else-if="item.icon === 'list'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
                  </svg>
                  <!-- Location Icon -->
                  <svg v-else-if="item.icon === 'location'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
                  </svg>
                  <!-- Category Icon -->
                  <svg v-else-if="item.icon === 'category'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"/>
                  </svg>
                  <!-- Quota Icon -->
                  <svg v-else-if="item.icon === 'quota'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"/>
                  </svg>
                  <!-- Shop Icon -->
                  <svg v-else-if="item.icon === 'shop'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
                  </svg>
                  <!-- Globe Icon -->
                  <svg v-else-if="item.icon === 'globe'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                  </svg>
                  <!-- Heart Icon -->
                  <svg v-else-if="item.icon === 'heart'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
                  </svg>
                  <!-- Star Icon -->
                  <svg v-else-if="item.icon === 'star'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/>
                  </svg>
                  <!-- Chat Icon -->
                  <svg v-else-if="item.icon === 'chat'" class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"/>
                  </svg>
                  <!-- Default Icon -->
                  <svg v-else class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                  </svg>
                  <span>{{ item.name }}</span>
                </router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </nav>
  </aside>
</template>
