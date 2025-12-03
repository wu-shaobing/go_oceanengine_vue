<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 156 })
const selectedComments = ref<string[]>([])
const replyTemplate = ref('')

const comments = ref([
  { id: 'CM001', content: '这个产品在哪里可以买？', author: '用户A', time: '2025-11-28 09:15', adTitle: '新品上市活动', status: 'pending' },
  { id: 'CM002', content: '价格是多少？', author: '用户B', time: '2025-11-28 09:10', adTitle: '双11促销', status: 'pending' },
  { id: 'CM003', content: '质量怎么样？', author: '用户C', time: '2025-11-28 08:55', adTitle: '新品上市活动', status: 'pending' },
  { id: 'CM004', content: '有优惠吗？', author: '用户D', time: '2025-11-28 08:40', adTitle: '品牌推广', status: 'replied' }
])

const templates = ref([
  { id: 'T1', name: '购买引导', content: '您好，可以点击视频下方链接直接购买哦～' },
  { id: 'T2', name: '价格咨询', content: '感谢关注！活动价格请查看商品详情页～' },
  { id: 'T3', name: '通用回复', content: '感谢您的关注，如有问题可私信咨询～' }
])

const toggleSelect = (id: string) => {
  const idx = selectedComments.value.indexOf(id)
  if (idx > -1) selectedComments.value.splice(idx, 1)
  else selectedComments.value.push(id)
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '评论管理' }, { name: '评论回复' }]" />
      <h1 class="text-3xl font-bold text-gray-900">评论回复</h1>
      <p class="mt-2 text-gray-600">批量管理和回复广告评论</p>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div class="col-span-2 space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm text-gray-500">待回复 {{ comments.filter(c => c.status === 'pending').length }} 条</span>
            <button v-if="selectedComments.length > 0" class="px-4 py-1.5 bg-blue-600 text-white text-sm rounded hover:bg-blue-700">
              批量回复 ({{ selectedComments.length }})
            </button>
          </div>
          
          <div class="space-y-3">
            <div v-for="comment in comments" :key="comment.id"
                 :class="['p-4 border rounded-lg', selectedComments.includes(comment.id) ? 'border-blue-500 bg-blue-50' : 'border-gray-200']">
              <div class="flex items-start gap-3">
                <input type="checkbox" :checked="selectedComments.includes(comment.id)"
                       @change="toggleSelect(comment.id)" class="mt-1 rounded text-blue-600">
                <div class="flex-1">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-900">{{ comment.author }}</span>
                    <span class="text-xs text-gray-400">{{ comment.time }}</span>
                  </div>
                  <p class="text-gray-700 mt-1">{{ comment.content }}</p>
                  <div class="flex items-center justify-between mt-2">
                    <span class="text-xs text-gray-500">广告: {{ comment.adTitle }}</span>
                    <span :class="['px-2 py-0.5 rounded text-xs',
                           comment.status === 'pending' ? 'bg-yellow-100 text-yellow-700' : 'bg-green-100 text-green-700']">
                      {{ comment.status === 'pending' ? '待回复' : '已回复' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <div class="mt-4">
            <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">快捷回复</h4>
          <div class="space-y-2">
            <div v-for="tpl in templates" :key="tpl.id"
                 class="p-3 bg-gray-50 rounded cursor-pointer hover:bg-gray-100"
                 @click="replyTemplate = tpl.content">
              <p class="text-sm font-medium text-gray-700">{{ tpl.name }}</p>
              <p class="text-xs text-gray-500 mt-1">{{ tpl.content }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h4 class="font-medium text-gray-900 mb-3">回复内容</h4>
          <textarea v-model="replyTemplate" rows="4" placeholder="请输入回复内容..."
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"></textarea>
          <button class="w-full mt-3 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            发送回复
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
