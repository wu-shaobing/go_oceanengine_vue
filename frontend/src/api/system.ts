import request from './request'
import type { PageResponse } from './request'

// ==================== 用户管理 ====================

export interface User {
  id: number
  username: string
  nickname: string
  phone: string
  email: string
  avatar: string
  status: number
  role_id: number
  role_name?: string
  last_login_at: string
  created_at: string
}

export interface UserListParams {
  page?: number
  page_size?: number
  username?: string
  nickname?: string
  phone?: string
  status?: number
  role_id?: number
}

export interface UserCreateParams {
  username: string
  password: string
  nickname?: string
  phone?: string
  email?: string
  avatar?: string
  status?: number
  role_id: number
  remark?: string
}

export interface UserUpdateParams {
  id: number
  nickname?: string
  phone?: string
  email?: string
  avatar?: string
  status?: number
  role_id?: number
  remark?: string
}

// ==================== 角色管理 ====================

export interface Role {
  id: number
  name: string
  key: string
  sort: number
  status: number
  data_scope: number
  remark: string
  created_at: string
}

export interface RoleListParams {
  page?: number
  page_size?: number
  name?: string
  code?: string
  status?: number
}

export interface RoleCreateParams {
  name: string
  code: string
  sort?: number
  status?: number
  remark?: string
}

export interface RoleUpdateParams {
  id: number
  name?: string
  code?: string
  sort?: number
  status?: number
  remark?: string
}

// ==================== 菜单管理 ====================

export interface Menu {
  id: number
  parent_id: number
  name: string
  path: string
  component: string
  icon: string
  sort: number
  type: number
  permission: string
  status: number
  visible: number
  is_frame: number
  is_cache: number
  remark: string
  children?: Menu[]
}

export interface MenuTree {
  id: number
  parent_id: number
  name: string
  path: string
  component: string
  icon: string
  sort: number
  type: number
  permission: string
  status: number
  hidden: number
  children?: MenuTree[]
}

export interface MenuCreateParams {
  parent_id?: number
  name: string
  path?: string
  component?: string
  icon?: string
  sort?: number
  type: number
  permission?: string
  status?: number
  hidden?: number
  remark?: string
}

export interface MenuUpdateParams {
  id: number
  parent_id?: number
  name?: string
  path?: string
  component?: string
  icon?: string
  sort?: number
  type?: number
  permission?: string
  status?: number
  hidden?: number
  remark?: string
}

// ==================== 操作日志 ====================

export interface OperationLog {
  id: number
  user_id: number
  username: string
  module: string
  action: string
  method: string
  path: string
  ip: string
  user_agent: string
  request: string
  response: string
  status: number
  duration: number
  created_at: string
}

export interface OperationLogListParams {
  page?: number
  page_size?: number
  user_id?: number
  username?: string
  module?: string
  action?: string
  status?: number
  start_time?: string
  end_time?: string
}

// ==================== 用户设置 ====================

export interface UserSettingResp {
  language: string
  timezone: string
  theme: string
  notifications_enabled: boolean
  email_alerts_enabled: boolean
  sms_alerts_enabled: boolean
  auto_refresh_enabled: boolean
  refresh_interval: number
}

export interface UserSettingUpdateReq {
  language?: string
  timezone?: string
  theme?: string
  notifications_enabled?: boolean
  email_alerts_enabled?: boolean
  sms_alerts_enabled?: boolean
  auto_refresh_enabled?: boolean
  refresh_interval?: number
}

// ==================== 消息通知 ====================

export interface NotificationItem {
  id: number
  title: string
  content: string
  type: 'success' | 'warning' | 'error' | 'info'
  is_read: boolean
  link: string
  created_at: string
}

export interface NotificationListParams {
  page?: number
  page_size?: number
  type?: string
  is_read?: boolean
  keyword?: string
}

export interface NotificationStatsResp {
  total: number
  unread: number
  today_new: number
  important: number
}

// ==================== 字典管理 ====================

export interface DictType {
  id: number
  name: string
  type: string
  status: number
  remark: string
  created_at: string
}

export interface DictTypeListParams {
  page?: number
  page_size?: number
  name?: string
  type?: string
  status?: number
}

export interface DictTypeCreateParams {
  name: string
  type: string
  status?: number
  remark?: string
}

export interface DictTypeUpdateParams {
  id: number
  name?: string
  type?: string
  status?: number
  remark?: string
}

export interface DictData {
  id: number
  dict_type: string
  label: string
  value: string
  sort: number
  status: number
  is_default: boolean
  remark: string
  css_class: string
  list_class: string
  created_at: string
}

export interface DictDataListParams {
  page?: number
  page_size?: number
  dict_type: string
  label?: string
  status?: number
}

export interface DictDataCreateParams {
  dict_type: string
  label: string
  value: string
  sort?: number
  status?: number
  is_default?: boolean
  remark?: string
  css_class?: string
  list_class?: string
}

export interface DictDataUpdateParams {
  id: number
  label?: string
  value?: string
  sort?: number
  status?: number
  is_default?: boolean
  remark?: string
  css_class?: string
  list_class?: string
}

// ==================== API 定义 ====================

export const systemApi = {
  // ===== 用户管理 =====
  getUserList(params: UserListParams) {
    return request.get<PageResponse<User>>('/system/users', params)
  },

  getUserById(id: number) {
    return request.get<User>(`/system/users/${id}`)
  },

  createUser(data: UserCreateParams) {
    return request.post<User>('/system/users', data)
  },

  updateUser(id: number, data: UserUpdateParams) {
    return request.put<User>(`/system/users/${id}`, data)
  },

  deleteUser(id: number) {
    return request.delete<void>(`/system/users/${id}`)
  },

  resetPassword(id: number, password: string) {
    return request.post<void>(`/system/users/${id}/reset-password`, { id, password })
  },

  changePassword(data: { old_password: string; new_password: string }) {
    return request.post<void>('/system/users/change-password', data)
  },

  // ===== 角色管理 =====
  getRoleList(params: RoleListParams) {
    return request.get<PageResponse<Role>>('/system/roles', params)
  },

  getAllRoles() {
    return request.get<Role[]>('/system/roles/all')
  },

  getRoleById(id: number) {
    return request.get<Role>(`/system/roles/${id}`)
  },

  createRole(data: RoleCreateParams) {
    return request.post<Role>('/system/roles', data)
  },

  updateRole(id: number, data: RoleUpdateParams) {
    return request.put<Role>(`/system/roles/${id}`, data)
  },

  deleteRole(id: number) {
    return request.delete<void>(`/system/roles/${id}`)
  },

  getRoleMenus(roleId: number) {
    return request.get<number[]>(`/system/roles/${roleId}/menus`)
  },

  updateRoleMenus(roleId: number, menuIds: number[]) {
    return request.put<void>(`/system/roles/${roleId}/menus`, { menu_ids: menuIds })
  },

  // ===== 菜单管理 =====
  getMenuList() {
    return request.get<Menu[]>('/system/menus')
  },

  getMenuTree() {
    return request.get<MenuTree[]>('/system/menus/tree')
  },

  getUserMenuTree() {
    return request.get<MenuTree[]>('/system/menus/user')
  },

  getMenuById(id: number) {
    return request.get<Menu>(`/system/menus/${id}`)
  },

  createMenu(data: MenuCreateParams) {
    return request.post<Menu>('/system/menus', data)
  },

  updateMenu(id: number, data: MenuUpdateParams) {
    return request.put<Menu>(`/system/menus/${id}`, data)
  },

  deleteMenu(id: number) {
    return request.delete<void>(`/system/menus/${id}`)
  },

  // ===== 操作日志 =====
  getOperationLogList(params: OperationLogListParams) {
    return request.get<PageResponse<OperationLog>>('/system/logs/operation', params)
  },

  getOperationLogModules() {
    return request.get<string[]>('/system/logs/modules')
  },

  deleteOperationLogs(beforeTime: string) {
    return request.delete<void>('/system/logs/operation', { before_time: beforeTime })
  },

  // ===== 用户设置 =====
  getUserSettings() {
    return request.get<UserSettingResp>('/system/settings')
  },

  updateUserSettings(data: UserSettingUpdateReq) {
    return request.put<void>('/system/settings', data)
  },

  // ===== 消息通知 =====
  getNotificationList(params: NotificationListParams) {
    return request.get<PageResponse<NotificationItem>>('/system/notifications', params)
  },

  getNotificationStats() {
    return request.get<NotificationStatsResp>('/system/notifications/stats')
  },

  markNotificationsAsRead(ids: number[]) {
    return request.post<void>('/system/notifications/read', { ids })
  },

  markAllNotificationsAsRead() {
    return request.post<void>('/system/notifications/read-all')
  },

  deleteNotifications(ids: number[]) {
    return request.delete<void>('/system/notifications', { ids })
  },

  // ===== 字典类型 =====
  getDictTypeList(params: DictTypeListParams) {
    return request.get<PageResponse<DictType>>('/system/dict/types', params)
  },

  getDictTypeById(id: number) {
    return request.get<DictType>(`/system/dict/types/${id}`)
  },

  createDictType(data: DictTypeCreateParams) {
    return request.post<void>('/system/dict/types', data)
  },

  updateDictType(id: number, data: DictTypeUpdateParams) {
    return request.put<void>(`/system/dict/types/${id}`, data)
  },

  deleteDictType(id: number) {
    return request.delete<void>(`/system/dict/types/${id}`)
  },

  // ===== 字典数据 =====
  getDictDataList(params: DictDataListParams) {
    return request.get<PageResponse<DictData>>('/system/dict/data', params)
  },

  getDictDataByType(type: string) {
    return request.get<DictData[]>(`/system/dict/data/${type}`)
  },

  createDictData(data: DictDataCreateParams) {
    return request.post<void>('/system/dict/data', data)
  },

  updateDictData(id: number, data: DictDataUpdateParams) {
    return request.put<void>(`/system/dict/data/${id}`, data)
  },

  deleteDictData(id: number) {
    return request.delete<void>(`/system/dict/data/${id}`)
  }
}

export default systemApi
