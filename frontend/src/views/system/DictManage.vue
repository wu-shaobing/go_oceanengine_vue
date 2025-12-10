<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import { systemApi } from '@/api/system'
import type { DictType, DictData } from '@/api/system'

// ==================== 字典类型 ====================

const typeLoading = ref(true)
const typeList = ref<DictType[]>([])
const currentType = ref<DictType | null>(null)
const typeSearch = ref('')

const typePagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

const typeColumns = [
  { key: 'name', title: '名称' },
  { key: 'type', title: '类型' },
  { key: 'status', title: '状态', width: 80 },
  { key: 'actions', title: '操作', width: 120, align: 'center' as const }
]

// 加载字典类型列表
async function loadDictTypes() {
  typeLoading.value = true
  try {
    const res = await systemApi.getDictTypeList({
      page: typePagination.current,
      page_size: typePagination.pageSize,
      name: typeSearch.value || undefined
    })
    typeList.value = res.list || []
    typePagination.total = res.total || 0
  } catch (error) {
    console.error('加载字典类型失败:', error)
  } finally {
    typeLoading.value = false
  }
}

// 选择字典类型
function handleTypeSelect(row: DictType) {
  currentType.value = row
  dataPagination.current = 1
  loadDictData()
}

// 字典类型表单
const showTypeDialog = ref(false)
const typeFormMode = ref<'add' | 'edit'>('add')
const typeForm = reactive({
  id: 0,
  name: '',
  type: '',
  status: 1,
  remark: ''
})

function resetTypeForm() {
  typeForm.id = 0
  typeForm.name = ''
  typeForm.type = ''
  typeForm.status = 1
  typeForm.remark = ''
}

function handleAddType() {
  typeFormMode.value = 'add'
  resetTypeForm()
  showTypeDialog.value = true
}

function handleEditType(row: DictType) {
  typeFormMode.value = 'edit'
  typeForm.id = row.id
  typeForm.name = row.name
  typeForm.type = row.type
  typeForm.status = row.status
  typeForm.remark = row.remark
  showTypeDialog.value = true
}

async function handleSaveType() {
  if (!typeForm.name || !typeForm.type) {
    alert('请填写必填字段')
    return
  }
  try {
    if (typeFormMode.value === 'edit') {
      await systemApi.updateDictType(typeForm.id, {
        id: typeForm.id,
        name: typeForm.name,
        type: typeForm.type,
        status: typeForm.status,
        remark: typeForm.remark
      })
      alert('更新成功')
    } else {
      await systemApi.createDictType({
        name: typeForm.name,
        type: typeForm.type,
        status: typeForm.status,
        remark: typeForm.remark
      })
      alert('创建成功')
    }
    showTypeDialog.value = false
    loadDictTypes()
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败')
  }
}

async function handleDeleteType(row: DictType) {
  if (!confirm(`确定删除字典类型 ${row.name}?`)) return
  try {
    await systemApi.deleteDictType(row.id)
    alert('删除成功')
    if (currentType.value?.id === row.id) {
      currentType.value = null
      dataList.value = []
    }
    loadDictTypes()
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// ==================== 字典数据 ====================

const dataLoading = ref(false)
const dataList = ref<DictData[]>([])
const dataSearch = ref('')

const dataPagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

const dataColumns = [
  { key: 'label', title: '标签' },
  { key: 'value', title: '键值' },
  { key: 'sort', title: '排序', width: 70 },
  { key: 'status', title: '状态', width: 70 },
  { key: 'is_default', title: '默认', width: 70 },
  { key: 'remark', title: '备注' },
  { key: 'actions', title: '操作', width: 120, align: 'center' as const }
]

// 加载字典数据列表
async function loadDictData() {
  if (!currentType.value) return
  dataLoading.value = true
  try {
    const res = await systemApi.getDictDataList({
      page: dataPagination.current,
      page_size: dataPagination.pageSize,
      dict_type: currentType.value.type,
      label: dataSearch.value || undefined
    })
    dataList.value = res.list || []
    dataPagination.total = res.total || 0
  } catch (error) {
    console.error('加载字典数据失败:', error)
  } finally {
    dataLoading.value = false
  }
}

// 字典数据表单
const showDataDialog = ref(false)
const dataFormMode = ref<'add' | 'edit'>('add')
const dataForm = reactive({
  id: 0,
  label: '',
  value: '',
  sort: 0,
  status: 1,
  is_default: false,
  css_class: '',
  list_class: '',
  remark: ''
})

function resetDataForm() {
  dataForm.id = 0
  dataForm.label = ''
  dataForm.value = ''
  dataForm.sort = 0
  dataForm.status = 1
  dataForm.is_default = false
  dataForm.css_class = ''
  dataForm.list_class = ''
  dataForm.remark = ''
}

function handleAddData() {
  if (!currentType.value) {
    alert('请先选择字典类型')
    return
  }
  dataFormMode.value = 'add'
  resetDataForm()
  showDataDialog.value = true
}

function handleEditData(row: DictData) {
  dataFormMode.value = 'edit'
  dataForm.id = row.id
  dataForm.label = row.label
  dataForm.value = row.value
  dataForm.sort = row.sort
  dataForm.status = row.status
  dataForm.is_default = row.is_default
  dataForm.css_class = row.css_class
  dataForm.list_class = row.list_class
  dataForm.remark = row.remark
  showDataDialog.value = true
}

async function handleSaveData() {
  if (!currentType.value || !dataForm.label || !dataForm.value) {
    alert('请填写必填字段')
    return
  }
  try {
    if (dataFormMode.value === 'edit') {
      await systemApi.updateDictData(dataForm.id, {
        id: dataForm.id,
        label: dataForm.label,
        value: dataForm.value,
        sort: dataForm.sort,
        status: dataForm.status,
        is_default: dataForm.is_default,
        css_class: dataForm.css_class,
        list_class: dataForm.list_class,
        remark: dataForm.remark
      })
      alert('更新成功')
    } else {
      await systemApi.createDictData({
        dict_type: currentType.value.type,
        label: dataForm.label,
        value: dataForm.value,
        sort: dataForm.sort,
        status: dataForm.status,
        is_default: dataForm.is_default,
        css_class: dataForm.css_class,
        list_class: dataForm.list_class,
        remark: dataForm.remark
      })
      alert('创建成功')
    }
    showDataDialog.value = false
    loadDictData()
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败')
  }
}

async function handleDeleteData(row: DictData) {
  if (!confirm(`确定删除字典数据 ${row.label}?`)) return
  try {
    await systemApi.deleteDictData(row.id)
    alert('删除成功')
    loadDictData()
  } catch (error) {
    console.error('删除失败:', error)
    alert('删除失败')
  }
}

// 分页变化
function handleTypePageChange(page: number) {
  typePagination.current = page
  loadDictTypes()
}

function handleDataPageChange(page: number) {
  dataPagination.current = page
  loadDictData()
}

// 搜索
function handleTypeSearch() {
  typePagination.current = 1
  loadDictTypes()
}

function handleDataSearch() {
  dataPagination.current = 1
  loadDictData()
}

// 初始化
onMounted(() => {
  loadDictTypes()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '系统管理' }, { name: '字典管理' }]" />
      <h1 class="text-3xl font-bold text-gray-900">字典管理</h1>
      <p class="mt-2 text-gray-600">管理系统字典类型和字典数据</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 左侧：字典类型列表 -->
      <div class="bg-white rounded-lg border border-gray-200">
        <div class="px-4 py-3 border-b border-gray-200 flex items-center justify-between">
          <h3 class="font-semibold text-gray-900">字典类型</h3>
          <button
            class="px-3 py-1.5 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 flex items-center gap-1"
            @click="handleAddType"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
            </svg>
            新增
          </button>
        </div>

        <!-- 搜索 -->
        <div class="px-4 py-3 border-b border-gray-100">
          <div class="flex gap-2">
            <input
              v-model="typeSearch"
              type="text"
              placeholder="搜索字典名称..."
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm"
              @keyup.enter="handleTypeSearch"
            />
            <button
              class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
              @click="handleTypeSearch"
            >
              搜索
            </button>
          </div>
        </div>

        <!-- 表格 -->
        <DataTable :columns="typeColumns" :data="typeList" :loading="typeLoading">
          <template #name="{ row }">
            <button
              class="text-blue-600 hover:text-blue-800 font-medium text-left"
              @click="handleTypeSelect(row)"
            >
              {{ row.name }}
            </button>
          </template>
          <template #status="{ row }">
            <StatusBadge
              :status="row.status === 1 ? 'success' : 'danger'"
              :text="row.status === 1 ? '启用' : '停用'"
            />
          </template>
          <template #actions="{ row }">
            <div class="flex items-center justify-center gap-2">
              <button class="text-blue-600 hover:text-blue-800 text-sm" @click="handleEditType(row)">编辑</button>
              <button class="text-red-600 hover:text-red-800 text-sm" @click="handleDeleteType(row)">删除</button>
            </div>
          </template>
        </DataTable>

        <div class="p-4 border-t border-gray-200">
          <Pagination
            :current="typePagination.current"
            :total="typePagination.total"
            :page-size="typePagination.pageSize"
            @change="handleTypePageChange"
          />
        </div>
      </div>

      <!-- 右侧：字典数据列表 -->
      <div class="bg-white rounded-lg border border-gray-200">
        <div class="px-4 py-3 border-b border-gray-200 flex items-center justify-between">
          <h3 class="font-semibold text-gray-900">
            字典数据
            <span v-if="currentType" class="ml-2 px-2 py-0.5 text-xs bg-blue-100 text-blue-700 rounded-full">
              {{ currentType.name }} ({{ currentType.type }})
            </span>
          </h3>
          <button
            class="px-3 py-1.5 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 flex items-center gap-1 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="!currentType"
            @click="handleAddData"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
            </svg>
            新增
          </button>
        </div>

        <template v-if="currentType">
          <!-- 搜索 -->
          <div class="px-4 py-3 border-b border-gray-100">
            <div class="flex gap-2">
              <input
                v-model="dataSearch"
                type="text"
                placeholder="搜索标签..."
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm"
                @keyup.enter="handleDataSearch"
              />
              <button
                class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
                @click="handleDataSearch"
              >
                搜索
              </button>
            </div>
          </div>

          <!-- 表格 -->
          <DataTable :columns="dataColumns" :data="dataList" :loading="dataLoading">
            <template #status="{ row }">
              <StatusBadge
                :status="row.status === 1 ? 'success' : 'danger'"
                :text="row.status === 1 ? '启用' : '停用'"
              />
            </template>
            <template #is_default="{ row }">
              <span v-if="row.is_default" class="px-2 py-0.5 text-xs bg-blue-100 text-blue-700 rounded-full">是</span>
              <span v-else class="text-gray-400">-</span>
            </template>
            <template #actions="{ row }">
              <div class="flex items-center justify-center gap-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm" @click="handleEditData(row)">编辑</button>
                <button class="text-red-600 hover:text-red-800 text-sm" @click="handleDeleteData(row)">删除</button>
              </div>
            </template>
          </DataTable>

          <div class="p-4 border-t border-gray-200">
            <Pagination
              :current="dataPagination.current"
              :total="dataPagination.total"
              :page-size="dataPagination.pageSize"
              @change="handleDataPageChange"
            />
          </div>
        </template>

        <div v-else class="px-4 py-16 text-center text-gray-500">
          请先从左侧选择一个字典类型
        </div>
      </div>
    </div>

    <!-- 字典类型表单弹窗 -->
    <div v-if="showTypeDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[450px]">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">
          {{ typeFormMode === 'add' ? '新增字典类型' : '编辑字典类型' }}
        </h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">名称 <span class="text-red-500">*</span></label>
            <input
              v-model="typeForm.name"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg"
              placeholder="请输入字典名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">类型 <span class="text-red-500">*</span></label>
            <input
              v-model="typeForm.type"
              type="text"
              :disabled="typeFormMode === 'edit'"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg disabled:bg-gray-100"
              placeholder="请输入字典类型（英文标识，如 sys_status）"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
            <select v-model="typeForm.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
              <option :value="1">启用</option>
              <option :value="0">停用</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">备注</label>
            <textarea
              v-model="typeForm.remark"
              rows="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg"
              placeholder="请输入备注"
            ></textarea>
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showTypeDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSaveType">保存</button>
        </div>
      </div>
    </div>

    <!-- 字典数据表单弹窗 -->
    <div v-if="showDataDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-[500px] max-h-[90vh] overflow-y-auto">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">
          {{ dataFormMode === 'add' ? '新增字典数据' : '编辑字典数据' }}
        </h3>
        <div class="space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">标签 <span class="text-red-500">*</span></label>
              <input
                v-model="dataForm.label"
                type="text"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg"
                placeholder="请输入标签名称"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">键值 <span class="text-red-500">*</span></label>
              <input
                v-model="dataForm.value"
                type="text"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg"
                placeholder="请输入键值"
              />
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">排序</label>
              <input
                v-model.number="dataForm.sort"
                type="number"
                min="0"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
              <select v-model="dataForm.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option :value="1">启用</option>
                <option :value="0">停用</option>
              </select>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">是否默认</label>
              <label class="flex items-center gap-2 mt-2">
                <input v-model="dataForm.is_default" type="checkbox" class="w-4 h-4 text-blue-600 rounded" />
                <span class="text-sm text-gray-700">设为默认值</span>
              </label>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">列表样式</label>
              <select v-model="dataForm.list_class" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
                <option value="">默认</option>
                <option value="primary">主要</option>
                <option value="success">成功</option>
                <option value="warning">警告</option>
                <option value="danger">危险</option>
                <option value="info">信息</option>
              </select>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">CSS类名</label>
            <input
              v-model="dataForm.css_class"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg"
              placeholder="自定义CSS类名（可选）"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">备注</label>
            <textarea
              v-model="dataForm.remark"
              rows="2"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg"
              placeholder="请输入备注"
            ></textarea>
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50" @click="showDataDialog = false">取消</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSaveData">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>
