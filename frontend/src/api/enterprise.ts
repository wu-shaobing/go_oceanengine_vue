import request, { PageResponse } from './request'

// ==================== 企业号信息 ====================

export interface EnterpriseInfo {
  account_id: string
  account_name: string
  avatar_url: string
  douyin_id: string
  fans_count: number
  following_count: number
  total_favorited: number
  aweme_count: number
  is_verified: boolean
  verification_type: string
  enterprise_verify_reason: string
  bind_time: string
}

export interface EnterpriseBind {
  bind_id: string
  account_id: string
  account_name: string
  avatar_url: string
  bind_type: string
  bind_status: string
  bind_time: string
}

// ==================== 视频相关 ====================

export interface EnterpriseItem {
  item_id: string
  title: string
  cover_url: string
  video_url: string
  create_time: string
  share_count: number
  comment_count: number
  digg_count: number
  play_count: number
  is_top: boolean
  item_status: string
}

// ==================== 评论相关 ====================

export interface EnterpriseComment {
  comment_id: string
  item_id: string
  content: string
  create_time: string
  digg_count: number
  reply_count: number
  user_id: string
  nickname: string
  avatar_url: string
  is_replied: boolean
  reply_content?: string
  reply_time?: string
}

export interface CommentReplyParams {
  comment_id: string
  content: string
}

// ==================== 数据报表 ====================

export interface EnterpriseOverviewData {
  date: string
  fans_count: number
  fans_increase: number
  play_count: number
  share_count: number
  comment_count: number
  digg_count: number
}

// ==================== 操作日志 ====================

export interface OperationLog {
  log_id: string
  operation_type: string
  operation_desc: string
  operator_name: string
  operation_time: string
  operation_result: string
}

// ==================== 快捷回复模板 ====================

export interface ReplyTemplate {
  template_id: string
  content: string
  use_count: number
  create_time: string
}

// ==================== API 方法 ====================

export const enterpriseApi = {
  // 企业号信息
  getInfo(account_id: string) {
    return request.get<EnterpriseInfo>('/enterprise/info', { account_id })
  },

  // 绑定列表
  getBindList(params: { page: number; page_size: number }) {
    return request.get<PageResponse<EnterpriseBind>>('/enterprise/binds', params)
  },

  bindAccount(data: { account_id: string; bind_type: string }) {
    return request.post<{ bind_id: string }>('/enterprise/binds', data)
  },

  unbindAccount(bind_id: string) {
    return request.delete<void>(`/enterprise/binds/${bind_id}`)
  },

  // 视频管理
  getItemList(params: { account_id: string; page: number; page_size: number; status?: string }) {
    return request.get<PageResponse<EnterpriseItem>>('/enterprise/items', params)
  },

  getItemDetail(item_id: string) {
    return request.get<EnterpriseItem>(`/enterprise/items/${item_id}`)
  },

  setTopItem(item_id: string, is_top: boolean) {
    return request.post<void>(`/enterprise/items/${item_id}/top`, { is_top })
  },

  deleteItem(item_id: string) {
    return request.delete<void>(`/enterprise/items/${item_id}`)
  },

  // 评论管理
  getCommentList(params: { account_id: string; item_id?: string; page: number; page_size: number; is_replied?: boolean }) {
    return request.get<PageResponse<EnterpriseComment>>('/enterprise/comments', params)
  },

  replyComment(data: CommentReplyParams) {
    return request.post<void>('/enterprise/comments/reply', data)
  },

  batchReplyComments(data: { comment_ids: string[]; content: string }) {
    return request.post<{ success_count: number; fail_count: number }>('/enterprise/comments/batch-reply', data)
  },

  updateReply(comment_id: string, content: string) {
    return request.put<void>(`/enterprise/comments/${comment_id}/reply`, { content })
  },

  hideComment(comment_id: string) {
    return request.post<void>(`/enterprise/comments/${comment_id}/hide`)
  },

  deleteComment(comment_id: string) {
    return request.delete<void>(`/enterprise/comments/${comment_id}`)
  },

  // 快捷回复模板
  getReplyTemplates() {
    return request.get<ReplyTemplate[]>('/enterprise/reply-templates')
  },

  createReplyTemplate(content: string) {
    return request.post<{ template_id: string }>('/enterprise/reply-templates', { content })
  },

  deleteReplyTemplate(template_id: string) {
    return request.delete<void>(`/enterprise/reply-templates/${template_id}`)
  },

  // 数据概览
  getOverviewData(params: { account_id: string; start_date: string; end_date: string }) {
    return request.get<EnterpriseOverviewData[]>('/enterprise/overview', params)
  },

  getDashboardStats(account_id: string) {
    return request.get<{
      fans_count: number
      fans_today: number
      play_total: number
      play_today: number
      digg_total: number
      comment_unreplied: number
    }>('/enterprise/dashboard', { account_id })
  },

  // 流量来源
  getTrafficSource(params: { account_id: string; start_date: string; end_date: string }) {
    return request.get<{
      source: string
      count: number
      ratio: number
    }[]>('/enterprise/traffic-source', params)
  },

  // 操作日志
  getOperationLogs(params: { account_id: string; page: number; page_size: number; operation_type?: string }) {
    return request.get<PageResponse<OperationLog>>('/enterprise/logs', params)
  }
}
