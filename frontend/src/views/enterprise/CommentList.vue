<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '评论管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">评论管理</h1>
        <p class="text-gray-600 mt-1">管理视频和直播间评论</p>
      </div>
      <div class="flex space-x-2">
        <button @click="batchReply" :disabled="selectedComments.length === 0" 
          :class="selectedComments.length > 0 ? 'bg-blue-600 hover:bg-blue-700' : 'bg-gray-300 cursor-not-allowed'"
          class="px-4 py-2 text-white rounded-lg">
          批量回复 ({{ selectedComments.length }})
        </button>
      </div>
    </div>

    <!-- 统计 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日评论</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stats.todayComments }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">待回复</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.pending }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">已回复</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.replied }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">回复率</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.replyRate }}%</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="filters.keyword" placeholder="搜索评论内容" 
          class="border border-gray-300 rounded-lg px-3 py-2 w-48" @keyup.enter="searchComments">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="pending">待回复</option>
          <option value="replied">已回复</option>
          <option value="hidden">已隐藏</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部来源</option>
          <option value="video">视频评论</option>
          <option value="live">直播评论</option>
        </select>
        <button @click="searchComments" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
        <button @click="resetFilters" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">重置</button>
      </div>
    </div>

    <!-- 快捷回复模板 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex items-center justify-between mb-3">
        <span class="font-medium text-gray-700">快捷回复模板</span>
        <button @click="showTemplateModal = true" class="text-blue-600 text-sm hover:underline">管理模板</button>
      </div>
      <div class="flex flex-wrap gap-2">
        <button v-for="tpl in replyTemplates" :key="tpl.id" @click="useTemplate(tpl.content)"
          class="px-3 py-1 bg-gray-100 text-gray-700 text-sm rounded-full hover:bg-gray-200">
          {{ tpl.name }}
        </button>
      </div>
    </div>

    <!-- 评论列表 -->
    <div class="space-y-4">
      <div v-for="comment in filteredComments" :key="comment.id" class="bg-white rounded-lg shadow p-4">
        <div class="flex">
          <div class="mr-3 flex items-start">
            <input type="checkbox" v-model="selectedComments" :value="comment.id" 
              :disabled="comment.status !== 'pending'"
              class="w-4 h-4 text-blue-600 rounded focus:ring-blue-500">
          </div>
          <img :src="comment.avatar" class="w-10 h-10 rounded-full mr-4" alt="">
          <div class="flex-1">
            <div class="flex items-center justify-between">
              <div>
                <span class="font-medium">{{ comment.user }}</span>
                <span class="text-sm text-gray-400 ml-2">{{ comment.time }}</span>
              </div>
              <span :class="getStatusClass(comment.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(comment.status) }}
              </span>
            </div>
            <div class="mt-2 text-gray-700">{{ comment.content }}</div>
            <div class="mt-2 text-sm text-gray-500">
              来自: {{ comment.source }} · {{ comment.videoTitle }}
            </div>
            
            <!-- 已有回复展示 -->
            <div v-if="comment.reply" class="mt-3 pl-4 border-l-2 border-blue-200 bg-blue-50 p-3 rounded">
              <div class="flex items-center justify-between">
                <span class="text-sm text-blue-600 font-medium">已回复:</span>
                <span class="text-xs text-gray-400">{{ comment.replyTime }}</span>
              </div>
              <div class="text-sm text-gray-700 mt-1">{{ comment.reply }}</div>
            </div>
            
            <!-- 回复输入框 -->
            <div v-if="replyingCommentId === comment.id" class="mt-3">
              <textarea v-model="replyContent" rows="3" 
                class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="请输入回复内容..."></textarea>
              <div class="flex items-center justify-between mt-2">
                <div class="flex flex-wrap gap-1">
                  <button v-for="tpl in replyTemplates.slice(0, 3)" :key="tpl.id" @click="replyContent = tpl.content"
                    class="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded hover:bg-gray-200">
                    {{ tpl.name }}
                  </button>
                </div>
                <div class="flex space-x-2">
                  <button @click="cancelReply" class="px-3 py-1 border border-gray-300 rounded text-sm hover:bg-gray-50">取消</button>
                  <button @click="submitReply(comment.id)" :disabled="!replyContent.trim() || submitting"
                    :class="replyContent.trim() && !submitting ? 'bg-blue-600 hover:bg-blue-700' : 'bg-gray-300 cursor-not-allowed'"
                    class="px-3 py-1 text-white rounded text-sm">
                    {{ submitting ? '提交中...' : '发送回复' }}
                  </button>
                </div>
              </div>
            </div>
            
            <!-- 操作按钮 -->
            <div class="mt-3 flex space-x-3">
              <button v-if="comment.status === 'pending' && replyingCommentId !== comment.id" 
                @click="startReply(comment.id)" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                回复
              </button>
              <button v-if="comment.status === 'replied'" @click="startReply(comment.id, comment.reply)" 
                class="text-blue-600 hover:text-blue-800 text-sm">修改回复</button>
              <button @click="toggleHide(comment)" class="text-gray-500 hover:text-gray-700 text-sm">
                {{ comment.status === 'hidden' ? '取消隐藏' : '隐藏' }}
              </button>
              <button @click="deleteComment(comment.id)" class="text-red-500 hover:text-red-700 text-sm">删除</button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 空状态 -->
      <div v-if="filteredComments.length === 0" class="bg-white rounded-lg shadow p-12 text-center">
        <div class="text-4xl mb-4">💬</div>
        <div class="text-gray-500">暂无评论数据</div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="pagination.current" :total="pagination.total" :page-size="pagination.pageSize" 
        @change="handlePageChange" />
    </div>

    <!-- 回复模板管理弹窗 -->
    <div v-if="showTemplateModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg w-[500px] max-h-[80vh] overflow-y-auto">
        <div class="flex justify-between items-center p-4 border-b">
          <h3 class="text-lg font-medium">管理回复模板</h3>
          <button @click="showTemplateModal = false" class="text-gray-400 hover:text-gray-600">✕</button>
        </div>
        <div class="p-4">
          <div class="mb-4">
            <div class="flex space-x-2">
              <input type="text" v-model="newTemplate.name" placeholder="模板名称" 
                class="flex-1 border border-gray-300 rounded px-3 py-2">
              <button @click="addTemplate" :disabled="!newTemplate.name.trim() || !newTemplate.content.trim()"
                class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-gray-300">添加</button>
            </div>
            <textarea v-model="newTemplate.content" rows="2" placeholder="模板内容"
              class="w-full border border-gray-300 rounded px-3 py-2 mt-2"></textarea>
          </div>
          <div class="space-y-2">
            <div v-for="tpl in replyTemplates" :key="tpl.id" 
              class="flex items-center justify-between p-3 bg-gray-50 rounded">
              <div>
                <div class="font-medium text-sm">{{ tpl.name }}</div>
                <div class="text-xs text-gray-500 mt-1">{{ tpl.content }}</div>
              </div>
              <button @click="removeTemplate(tpl.id)" class="text-red-500 hover:text-red-700 text-sm">删除</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 批量回复弹窗 -->
    <div v-if="showBatchReplyModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg w-[500px]">
        <div class="flex justify-between items-center p-4 border-b">
          <h3 class="text-lg font-medium">批量回复 ({{ selectedComments.length }}条评论)</h3>
          <button @click="showBatchReplyModal = false" class="text-gray-400 hover:text-gray-600">✕</button>
        </div>
        <div class="p-4">
          <div class="mb-3">
            <label class="block text-sm text-gray-600 mb-2">选择回复模板</label>
            <div class="flex flex-wrap gap-2">
              <button v-for="tpl in replyTemplates" :key="tpl.id" @click="batchReplyContent = tpl.content"
                :class="batchReplyContent === tpl.content ? 'bg-blue-100 border-blue-500' : 'bg-gray-100 border-gray-200'"
                class="px-3 py-1 text-sm rounded-full border">
                {{ tpl.name }}
              </button>
            </div>
          </div>
          <textarea v-model="batchReplyContent" rows="4" placeholder="请输入回复内容..." 
            class="w-full border border-gray-300 rounded-lg px-3 py-2"></textarea>
        </div>
        <div class="flex justify-end space-x-3 p-4 border-t bg-gray-50">
          <button @click="showBatchReplyModal = false" class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">取消</button>
          <button @click="submitBatchReply" :disabled="!batchReplyContent.trim() || submitting"
            class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-gray-300">
            {{ submitting ? '提交中...' : '确认回复' }}
          </button>
        </div>
      </div>
    </div>

    <!-- 删除确认弹窗 -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg w-[400px]">
        <div class="p-6 text-center">
          <div class="text-4xl mb-4">⚠️</div>
          <div class="text-lg font-medium mb-2">确认删除评论？</div>
          <div class="text-gray-500 text-sm">删除后将无法恢复</div>
        </div>
        <div class="flex justify-center space-x-3 p-4 border-t">
          <button @click="showDeleteModal = false" class="px-6 py-2 border border-gray-300 rounded hover:bg-gray-50">取消</button>
          <button @click="confirmDelete" class="px-6 py-2 bg-red-600 text-white rounded hover:bg-red-700">确认删除</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

interface Comment {
  id: number
  user: string
  avatar: string
  content: string
  time: string
  status: 'pending' | 'replied' | 'hidden'
  source: string
  videoTitle: string
  reply: string
  replyTime?: string
}

interface ReplyTemplate {
  id: number
  name: string
  content: string
}

const stats = ref({
  todayComments: 256,
  pending: 68,
  replied: 188,
  replyRate: 73.4
})

const filters = ref({
  keyword: '',
  status: '',
  type: ''
})

const pagination = ref({
  current: 1,
  total: 200,
  pageSize: 10
})

const comments = ref<Comment[]>([
  { id: 1, user: '小明', avatar: 'https://via.placeholder.com/40', content: '产品质量真的很好，已经回购三次了！强烈推荐给大家～', time: '10分钟前', status: 'pending', source: '视频评论', videoTitle: '新品发布会精彩回顾', reply: '' },
  { id: 2, user: '用户A', avatar: 'https://via.placeholder.com/40', content: '请问这款什么时候补货？等了好久了', time: '30分钟前', status: 'replied', source: '视频评论', videoTitle: '产品使用教程', reply: '亲，预计下周会补货，您可以先加购物车哦～', replyTime: '25分钟前' },
  { id: 3, user: '小红', avatar: 'https://via.placeholder.com/40', content: '客服态度很好，点赞！下次还会购买', time: '1小时前', status: 'pending', source: '直播评论', videoTitle: '直播带货专场', reply: '' },
  { id: 4, user: '用户B', avatar: 'https://via.placeholder.com/40', content: '收到货了，包装很精美，物流也快', time: '2小时前', status: 'replied', source: '视频评论', videoTitle: '开箱测评', reply: '感谢您的认可，期待您的下次光临！', replyTime: '1小时前' },
  { id: 5, user: '路人甲', avatar: 'https://via.placeholder.com/40', content: '这个价格还可以再便宜点吗？', time: '3小时前', status: 'pending', source: '直播评论', videoTitle: '周末特惠专场', reply: '' },
  { id: 6, user: '老顾客', avatar: 'https://via.placeholder.com/40', content: '一直支持，质量稳定', time: '5小时前', status: 'pending', source: '视频评论', videoTitle: '品牌介绍', reply: '' }
])

const replyTemplates = ref<ReplyTemplate[]>([
  { id: 1, name: '感谢支持', content: '感谢您的支持与认可，我们会继续努力提供更好的产品和服务！' },
  { id: 2, name: '补货通知', content: '亲，这款商品预计近期补货，您可以先关注店铺，补货后会第一时间通知您～' },
  { id: 3, name: '售后咨询', content: '如有任何问题，请联系我们的客服，我们会尽快为您处理。' },
  { id: 4, name: '欢迎再来', content: '感谢您的购买，期待您的下次光临！有任何问题随时联系我们～' }
])

// 回复相关状态
const replyingCommentId = ref<number | null>(null)
const replyContent = ref('')
const submitting = ref(false)
const selectedComments = ref<number[]>([])

// 弹窗状态
const showTemplateModal = ref(false)
const showBatchReplyModal = ref(false)
const showDeleteModal = ref(false)
const deleteTargetId = ref<number | null>(null)
const batchReplyContent = ref('')

// 新模板
const newTemplate = ref({ name: '', content: '' })

// 筛选后的评论
const filteredComments = computed(() => {
  return comments.value.filter(c => {
    if (filters.value.keyword && !c.content.includes(filters.value.keyword)) return false
    if (filters.value.status && c.status !== filters.value.status) return false
    if (filters.value.type) {
      const typeMatch = filters.value.type === 'video' ? '视频评论' : '直播评论'
      if (c.source !== typeMatch) return false
    }
    return true
  })
})

// 方法
const startReply = (commentId: number, existingReply = '') => {
  replyingCommentId.value = commentId
  replyContent.value = existingReply
}

const cancelReply = () => {
  replyingCommentId.value = null
  replyContent.value = ''
}

const submitReply = async (commentId: number) => {
  if (!replyContent.value.trim()) return
  submitting.value = true
  
  // 模拟API调用
  await new Promise(resolve => setTimeout(resolve, 500))
  
  const comment = comments.value.find(c => c.id === commentId)
  if (comment) {
    comment.reply = replyContent.value
    comment.status = 'replied'
    comment.replyTime = '刚刚'
    
    // 更新统计
    stats.value.pending--
    stats.value.replied++
    stats.value.replyRate = Math.round((stats.value.replied / stats.value.todayComments) * 1000) / 10
  }
  
  submitting.value = false
  cancelReply()
}

const useTemplate = (content: string) => {
  if (replyingCommentId.value) {
    replyContent.value = content
  }
}

const batchReply = () => {
  if (selectedComments.value.length === 0) return
  showBatchReplyModal.value = true
}

const submitBatchReply = async () => {
  if (!batchReplyContent.value.trim()) return
  submitting.value = true
  
  await new Promise(resolve => setTimeout(resolve, 800))
  
  let repliedCount = 0
  selectedComments.value.forEach(id => {
    const comment = comments.value.find(c => c.id === id)
    if (comment && comment.status === 'pending') {
      comment.reply = batchReplyContent.value
      comment.status = 'replied'
      comment.replyTime = '刚刚'
      repliedCount++
    }
  })
  
  stats.value.pending -= repliedCount
  stats.value.replied += repliedCount
  stats.value.replyRate = Math.round((stats.value.replied / stats.value.todayComments) * 1000) / 10
  
  submitting.value = false
  showBatchReplyModal.value = false
  selectedComments.value = []
  batchReplyContent.value = ''
}

const toggleHide = (comment: Comment) => {
  comment.status = comment.status === 'hidden' ? 'pending' : 'hidden'
}

const deleteComment = (id: number) => {
  deleteTargetId.value = id
  showDeleteModal.value = true
}

const confirmDelete = () => {
  if (deleteTargetId.value) {
    const idx = comments.value.findIndex(c => c.id === deleteTargetId.value)
    if (idx !== -1) {
      const comment = comments.value[idx]
      if (comment.status === 'pending') stats.value.pending--
      else if (comment.status === 'replied') stats.value.replied--
      stats.value.todayComments--
      comments.value.splice(idx, 1)
    }
  }
  showDeleteModal.value = false
  deleteTargetId.value = null
}

const addTemplate = () => {
  if (!newTemplate.value.name.trim() || !newTemplate.value.content.trim()) return
  replyTemplates.value.push({
    id: Date.now(),
    name: newTemplate.value.name,
    content: newTemplate.value.content
  })
  newTemplate.value = { name: '', content: '' }
}

const removeTemplate = (id: number) => {
  const idx = replyTemplates.value.findIndex(t => t.id === id)
  if (idx !== -1) replyTemplates.value.splice(idx, 1)
}

const searchComments = () => {
  pagination.value.current = 1
  // 实际项目中这里会调用API
}

const resetFilters = () => {
  filters.value = { keyword: '', status: '', type: '' }
  searchComments()
}

const handlePageChange = (page: number) => {
  pagination.value.current = page
  // 实际项目中这里会调用API
}

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    pending: 'bg-orange-100 text-orange-800',
    replied: 'bg-green-100 text-green-800',
    hidden: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    pending: '待回复',
    replied: '已回复',
    hidden: '已隐藏'
  }
  return texts[status] || status
}
</script>
