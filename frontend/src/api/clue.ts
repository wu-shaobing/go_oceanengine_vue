import request, { PageResponse } from './request'

// ==================== 飞鱼线索类型 ====================

export interface Clue {
  clue_id: number
  advertiser_id: number
  name: string
  telephone: string
  gender: number
  age: number
  city: string
  address: string
  remark: string
  clue_type: number
  clue_source: string
  clue_state: number
  follow_state: number
  create_time: string
  update_time: string
  ad_id: number
  ad_name: string
  creative_id: number
  creative_title: string
}

export interface ClueListParams {
  advertiser_id: number
  start_time?: string
  end_time?: string
  clue_state?: number
  follow_state?: number
  search_type?: number
  search_value?: string
  page: number
  page_size: number
}

export interface ClueCallbackParams {
  advertiser_id: number
  clue_id: number
  clue_state: number
  follow_state?: number
  remark?: string
}

export interface KeyAction {
  action_id: number
  action_name: string
  action_type: number
  create_time: string
}

export interface SmartPhone {
  phone_id: number
  phone_number: string
  area: string
  status: number
  expire_time: string
}

export interface Form {
  form_id: number
  form_name: string
  form_type: number
  status: number
  create_time: string
}

export interface FormDetail extends Form {
  advertiser_id: number
  fields: FormField[]
}

export interface FormField {
  field_id: number
  field_name: string
  field_type: string
  required: boolean
  options?: string[]
}

export interface ClueStore {
  store_id: number
  store_name: string
  address: string
  province: string
  city: string
  district: string
  longitude: number
  latitude: number
  telephone: string
  status: number
}

// ==================== 青鸟线索通类型 ====================

export interface QingniaoForm {
  form_id: number
  form_name: string
  form_type: string
  status: string
  page_type: string
  create_time: string
  update_time: string
}

export interface QingniaoFormDetail extends QingniaoForm {
  advertiser_id: number
  components: FormComponent[]
  submit_config: SubmitConfig
}

export interface FormComponent {
  component_id: number
  component_type: string
  label: string
  placeholder?: string
  required: boolean
  options?: string[]
  validation?: object
}

export interface SubmitConfig {
  button_text: string
  success_text: string
  redirect_url?: string
}

export interface Coupon {
  coupon_id: number
  coupon_name: string
  coupon_type: string
  discount_value: number
  min_amount: number
  stock: number
  used_count: number
  start_time: string
  end_time: string
  status: string
  create_time: string
}

export interface CouponCreateParams {
  advertiser_id: number
  coupon_name: string
  coupon_type: string
  discount_value: number
  min_amount?: number
  stock: number
  start_time: string
  end_time: string
  description?: string
}

export interface Employee {
  employee_id: number
  employee_name: string
  phone: string
  role: string
  department?: string
  status: string
  bind_time: string
}

export interface EmployeeCreateParams {
  advertiser_id: number
  employee_name: string
  phone: string
  role: string
  department?: string
}

export interface QingniaoSmartPhone {
  phone_id: number
  phone_number: string
  area: string
  call_count: number
  duration: number
  status: string
  expire_time: string
}

export interface SmartPhoneRecord {
  record_id: string
  phone_number: string
  caller: string
  callee: string
  duration: number
  call_time: string
  record_url?: string
  status: string
}

export interface WechatPool {
  pool_id: number
  pool_name: string
  wechat_count: number
  online_count: number
  status: string
  create_time: string
}

export interface WechatInstance {
  instance_id: number
  wechat_id: string
  nickname: string
  avatar: string
  status: string
  friend_count: number
  today_add: number
  bindTime: string
}

// ==================== API 方法 ====================

export const clueApi = {
  // ========== 飞鱼线索 ==========
  
  // 获取线索列表
  getClueList(params: ClueListParams) {
    return request.get<PageResponse<Clue>>('/clue/list', params)
  },

  // 线索状态回调
  clueCallback(data: ClueCallbackParams) {
    return request.post<void>('/clue/callback', data)
  },

  // 批量线索状态回调
  batchClueCallback(data: { advertiser_id: number; callbacks: ClueCallbackParams[] }) {
    return request.post<{ success_count: number; fail_list: number[] }>('/clue/callback/batch', data)
  },

  // 获取关键行为
  getKeyAction(advertiser_id: number) {
    return request.get<KeyAction[]>('/clue/key-action', { advertiser_id })
  },

  // 获取智能电话信息
  getSmartPhone(advertiser_id: number) {
    return request.get<SmartPhone[]>('/clue/smart-phone', { advertiser_id })
  },

  // 获取表单列表
  getFormList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<Form>>('/clue/form/list', params)
  },

  // 获取表单详情
  getFormDetail(form_id: number, advertiser_id: number) {
    return request.get<FormDetail>('/clue/form/detail', { form_id, advertiser_id })
  },

  // 获取门店列表
  getClueStoreList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<ClueStore>>('/clue/store/list', params)
  },

  // ========== 青鸟线索通 - 表单管理 ==========

  // 获取青鸟表单列表
  getQingniaoFormList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<QingniaoForm>>('/clue/qingniao/forms', params)
  },

  // 获取青鸟表单详情
  getQingniaoFormDetail(form_id: number, advertiser_id: number) {
    return request.get<QingniaoFormDetail>(`/clue/qingniao/forms/${form_id}`, { advertiser_id })
  },

  // 创建青鸟表单
  createQingniaoForm(data: { advertiser_id: number; form_name: string; form_type: string; components: FormComponent[] }) {
    return request.post<{ form_id: number }>('/clue/qingniao/forms', data)
  },

  // 更新青鸟表单
  updateQingniaoForm(form_id: number, data: Partial<QingniaoFormDetail>) {
    return request.put<void>(`/clue/qingniao/forms/${form_id}`, data)
  },

  // 删除青鸟表单
  deleteQingniaoForm(form_id: number, advertiser_id: number) {
    return request.delete<void>(`/clue/qingniao/forms/${form_id}`, { advertiser_id })
  },

  // ========== 青鸟线索通 - 优惠券管理 ==========

  // 获取优惠券列表
  getCouponList(params: { advertiser_id: number; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<Coupon>>('/clue/qingniao/coupons', params)
  },

  // 获取优惠券详情
  getCouponDetail(coupon_id: number, advertiser_id: number) {
    return request.get<Coupon>(`/clue/qingniao/coupons/${coupon_id}`, { advertiser_id })
  },

  // 创建优惠券
  createCoupon(data: CouponCreateParams) {
    return request.post<{ coupon_id: number }>('/clue/qingniao/coupons', data)
  },

  // 更新优惠券
  updateCoupon(coupon_id: number, data: Partial<Coupon>) {
    return request.put<void>(`/clue/qingniao/coupons/${coupon_id}`, data)
  },

  // 删除优惠券
  deleteCoupon(coupon_id: number, advertiser_id: number) {
    return request.delete<void>(`/clue/qingniao/coupons/${coupon_id}`, { advertiser_id })
  },

  // ========== 青鸟线索通 - 员工管理 ==========

  // 获取员工列表
  getEmployeeList(params: { advertiser_id: number; page: number; page_size: number }) {
    return request.get<PageResponse<Employee>>('/clue/qingniao/employees', params)
  },

  // 创建员工
  createEmployee(data: EmployeeCreateParams) {
    return request.post<{ employee_id: number }>('/clue/qingniao/employees', data)
  },

  // 更新员工
  updateEmployee(employee_id: number, data: Partial<Employee>) {
    return request.put<void>(`/clue/qingniao/employees/${employee_id}`, data)
  },

  // 删除员工
  deleteEmployee(employee_id: number, advertiser_id: number) {
    return request.delete<void>(`/clue/qingniao/employees/${employee_id}`, { advertiser_id })
  },

  // ========== 青鸟线索通 - 智能电话 ==========

  // 获取青鸟智能电话信息
  getQingniaoSmartPhone(advertiser_id: number) {
    return request.get<QingniaoSmartPhone[]>('/clue/qingniao/smart-phone', { advertiser_id })
  },

  // 更新青鸟智能电话配置
  updateQingniaoSmartPhone(data: { advertiser_id: number; phone_id: number; config: object }) {
    return request.put<void>('/clue/qingniao/smart-phone', data)
  },

  // 获取通话记录
  getSmartPhoneRecords(params: { advertiser_id: number; phone_id?: number; start_time?: string; end_time?: string; page: number; page_size: number }) {
    return request.get<PageResponse<SmartPhoneRecord>>('/clue/qingniao/smart-phone/records', params)
  },

  // ========== 青鸟线索通 - 微信相关 ==========

  // 获取微信号池
  getWechatPool(advertiser_id: number) {
    return request.get<WechatPool[]>('/clue/qingniao/wechat/pool', { advertiser_id })
  },

  // 更新微信号池配置
  updateWechatPool(data: { advertiser_id: number; pool_id: number; config: object }) {
    return request.put<void>('/clue/qingniao/wechat/pool', data)
  },

  // 获取微信号实例列表
  getWechatInstances(params: { advertiser_id: number; pool_id?: number; page: number; page_size: number }) {
    return request.get<PageResponse<WechatInstance>>('/clue/qingniao/wechat/instances', params)
  }
}
