<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const saving = ref(false)

const isEdit = computed(() => !!route.params.id)

interface AudienceForm {
  name: string
  description: string
  type: string
  tags: string[]
  rules: AudienceRule[]
}

interface AudienceRule {
  id: number
  field: string
  operator: string
  value: string
}

const form = ref<AudienceForm>({
  name: '',
  description: '',
  type: 'custom',
  tags: [],
  rules: []
})

const availableFields = [
  { value: 'age', label: '年龄' },
  { value: 'gender', label: '性别' },
  { value: 'city', label: '城市' },
  { value: 'interest', label: '兴趣' },
  { value: 'device', label: '设备' },
  { value: 'behavior', label: '行为' }
]

const operators = [
  { value: 'eq', label: '等于' },
  { value: 'ne', label: '不等于' },
  { value: 'in', label: '包含' },
  { value: 'nin', label: '不包含' },
  { value: 'gt', label: '大于' },
  { value: 'lt', label: '小于' }
]

const audienceTypes = [
  { value: 'custom', label: '自定义人群' },
  { value: 'lookalike', label: '相似人群' },
  { value: 'retargeting', label: '再营销人群' },
  { value: 'dmp', label: 'DMP人群' }
]

const tagInput = ref('')

const fetchAudience = async () => {
  if (!isEdit.value) {
    loading.value = false
    return
  }
  
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // 模拟数据
  form.value = {
    name: '高价值用户群',
    description: '近30天有购买行为的高消费用户',
    type: 'custom',
    tags: ['高价值', '电商', '活跃用户'],
    rules: [
      { id: 1, field: 'age', operator: 'in', value: '25-45' },
      { id: 2, field: 'behavior', operator: 'eq', value: '购买' },
      { id: 3, field: 'city', operator: 'in', value: '一线城市' }
    ]
  }
  
  loading.value = false
}

const addRule = () => {
  form.value.rules.push({
    id: Date.now(),
    field: 'age',
    operator: 'eq',
    value: ''
  })
}

const removeRule = (id: number) => {
  form.value.rules = form.value.rules.filter(r => r.id !== id)
}

const addTag = () => {
  if (tagInput.value && !form.value.tags.includes(tagInput.value)) {
    form.value.tags.push(tagInput.value)
    tagInput.value = ''
  }
}

const removeTag = (tag: string) => {
  form.value.tags = form.value.tags.filter(t => t !== tag)
}

const handleSave = async () => {
  saving.value = true
  await new Promise(resolve => setTimeout(resolve, 1000))
  saving.value = false
  router.push('/audience')
}

const handleCancel = () => {
  router.push('/audience')
}

onMounted(fetchAudience)
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '人群管理', path: '/audience' }, { name: isEdit ? '编辑人群' : '新建人群' }]" />
      <h1 class="text-3xl font-bold text-gray-900">{{ isEdit ? '编辑人群' : '新建人群' }}</h1>
      <p class="mt-2 text-gray-600">{{ isEdit ? '修改人群包配置信息' : '创建新的定向人群包' }}</p>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center h-64">
      <div class="text-gray-500">加载中...</div>
    </div>

    <template v-else>
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              人群名称 <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
              placeholder="请输入人群名称"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">人群描述</label>
            <textarea
              v-model="form.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500 resize-none"
              placeholder="请输入人群描述"
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">人群类型</label>
            <select
              v-model="form.type"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="type in audienceTypes" :key="type.value" :value="type.value">
                {{ type.label }}
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">标签</label>
            <div class="flex flex-wrap gap-2 mb-2">
              <span
                v-for="tag in form.tags"
                :key="tag"
                class="inline-flex items-center gap-1 px-2 py-1 bg-blue-100 text-blue-800 text-sm rounded"
              >
                {{ tag }}
                <button @click="removeTag(tag)" class="hover:text-blue-600">
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                  </svg>
                </button>
              </span>
            </div>
            <div class="flex gap-2">
              <input
                v-model="tagInput"
                type="text"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
                placeholder="输入标签名称"
                @keyup.enter="addTag"
              />
              <button
                @click="addTag"
                class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors"
              >
                添加
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 定向规则 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">定向规则</h3>
          <button
            @click="addRule"
            class="px-3 py-1.5 text-sm text-blue-600 border border-blue-300 rounded-lg hover:bg-blue-50 transition-colors"
          >
            + 添加规则
          </button>
        </div>
        
        <div v-if="form.rules.length === 0" class="text-center py-8 text-gray-500">
          暂无定向规则，点击上方按钮添加
        </div>
        
        <div v-else class="space-y-3">
          <div
            v-for="(rule, index) in form.rules"
            :key="rule.id"
            class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg"
          >
            <span class="text-sm text-gray-500 w-16">
              {{ index === 0 ? '满足' : '且满足' }}
            </span>
            
            <select
              v-model="rule.field"
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="field in availableFields" :key="field.value" :value="field.value">
                {{ field.label }}
              </option>
            </select>
            
            <select
              v-model="rule.operator"
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            >
              <option v-for="op in operators" :key="op.value" :value="op.value">
                {{ op.label }}
              </option>
            </select>
            
            <input
              v-model="rule.value"
              type="text"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
              placeholder="请输入值"
            />
            
            <button
              @click="removeRule(rule.id)"
              class="p-2 text-gray-400 hover:text-red-500 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex items-center justify-end gap-3">
        <button
          @click="handleCancel"
          class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        >
          取消
        </button>
        <button
          @click="handleSave"
          :disabled="saving"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
        >
          <svg v-if="saving" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ saving ? '保存中...' : '保存' }}
        </button>
      </div>
    </template>
  </div>
</template>
