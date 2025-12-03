<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const router = useRouter()

const form = reactive({
  name: '',
  type: 'custom',
  templateId: '',
  domain: '',
  description: ''
})

const templates = ref([
  { id: 'T001', name: '电商促销', preview: '🛒', desc: '适合商品推广' },
  { id: 'T002', name: '表单收集', preview: '📝', desc: '适合线索收集' },
  { id: 'T003', name: 'APP下载', preview: '📱', desc: '适合应用推广' },
  { id: 'T004', name: '品牌展示', preview: '🏢', desc: '适合品牌宣传' }
])

const handleSubmit = () => {
  console.log('Create site:', form)
  router.push('/site/list')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '创建落地页' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建落地页</h1>
      <p class="mt-2 text-gray-600">创建新的广告落地页</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">落地页名称</label>
            <input v-model="form.name" type="text" placeholder="请输入落地页名称"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">创建方式</label>
            <select v-model="form.type" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="custom">自定义创建</option>
              <option value="template">使用模板</option>
              <option value="import">导入已有</option>
            </select>
          </div>
        </div>

        <div v-if="form.type === 'template'">
          <label class="block text-sm font-medium text-gray-700 mb-4">选择模板</label>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div v-for="tpl in templates" :key="tpl.id" 
                 :class="['p-4 border-2 rounded-lg cursor-pointer transition-all',
                   form.templateId === tpl.id ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
                 @click="form.templateId = tpl.id">
              <div class="text-4xl text-center mb-2">{{ tpl.preview }}</div>
              <h4 class="font-medium text-gray-900 text-center">{{ tpl.name }}</h4>
              <p class="text-xs text-gray-500 text-center mt-1">{{ tpl.desc }}</p>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">绑定域名（可选）</label>
          <input v-model="form.domain" type="text" placeholder="例如: landing.example.com"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">描述</label>
          <textarea v-model="form.description" rows="3" placeholder="落地页用途描述"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
        </div>

        <div class="flex justify-end gap-4">
          <router-link to="/site/list" class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50">
            取消
          </router-link>
          <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            创建并进入编辑
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
