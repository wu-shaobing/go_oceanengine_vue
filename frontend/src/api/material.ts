import request, { PageResponse } from './request'

// ==================== 图片素材 ====================

export interface ImageMaterial {
  id: number
  image_id: string
  advertiser_id: number
  filename: string
  url: string
  width: number
  height: number
  size: number
  format: string
  signature: string
  created_at: string
}

// ==================== 视频素材 ====================

export interface VideoMaterial {
  id: number
  video_id: string
  advertiser_id: number
  filename: string
  url: string
  poster_url: string
  width: number
  height: number
  duration: number
  size: number
  format: string
  bit_rate: number
  signature: string
  created_at: string
}

// ==================== 素材分组 ====================

export interface MaterialGroup {
  group_id: string
  group_name: string
  material_count: number
  created_at: string
}

// ==================== 上传参数 ====================

export interface UploadParams {
  advertiser_id: number
  upload_type?: string
  filename?: string
}

// ==================== API 方法 ====================

export const materialApi = {
  // 图片素材
  getImageList(params: { advertiser_id: number; page: number; page_size: number; group_id?: string }) {
    return request.get<PageResponse<ImageMaterial>>('/media/images', params)
  },

  getImageDetail(id: number) {
    return request.get<ImageMaterial>(`/media/images/${id}`)
  },

  uploadImage(advertiser_id: number, file: File, filename?: string) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    if (filename) {
      formData.append('filename', filename)
    }
    return request.post<{ id: number; image_id: string; url: string }>('/media/images/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  uploadImageByUrl(advertiser_id: number, image_url: string, filename?: string) {
    return request.post<{ id: number; image_id: string; url: string }>('/media/images/upload-url', {
      advertiser_id,
      image_url,
      filename
    })
  },

  deleteImage(id: number) {
    return request.delete<void>(`/media/images/${id}`)
  },

  batchDeleteImages(ids: number[]) {
    return request.delete<void>('/media/images/batch', { ids })
  },

  // 视频素材
  getVideoList(params: { advertiser_id: number; page: number; page_size: number; group_id?: string }) {
    return request.get<PageResponse<VideoMaterial>>('/media/videos', params)
  },

  getVideoDetail(id: number) {
    return request.get<VideoMaterial>(`/media/videos/${id}`)
  },

  uploadVideo(advertiser_id: number, file: File, filename?: string, onProgress?: (percent: number) => void) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('advertiser_id', String(advertiser_id))
    if (filename) {
      formData.append('filename', filename)
    }
    return request.post<{ id: number; video_id: string; url: string }>('/media/videos/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (progressEvent) => {
        if (onProgress && progressEvent.total) {
          const percent = Math.round((progressEvent.loaded * 100) / progressEvent.total)
          onProgress(percent)
        }
      }
    })
  },

  uploadVideoByUrl(advertiser_id: number, video_url: string, filename?: string) {
    return request.post<{ id: number; video_id: string; url: string }>('/media/videos/upload-url', {
      advertiser_id,
      video_url,
      filename
    })
  },

  deleteVideo(id: number) {
    return request.delete<void>(`/media/videos/${id}`)
  },

  batchDeleteVideos(ids: number[]) {
    return request.delete<void>('/media/videos/batch', { ids })
  },

  // 素材分组
  getGroupList(params: { advertiser_id: number; material_type: 'image' | 'video' }) {
    return request.get<MaterialGroup[]>('/media/groups', params)
  },

  createGroup(data: { advertiser_id: number; material_type: 'image' | 'video'; group_name: string }) {
    return request.post<{ group_id: string }>('/media/groups', data)
  },

  updateGroup(group_id: string, group_name: string) {
    return request.put<void>(`/media/groups/${group_id}`, { group_name })
  },

  deleteGroup(group_id: string) {
    return request.delete<void>(`/media/groups/${group_id}`)
  },

  moveToGroup(data: { material_ids: number[]; material_type: 'image' | 'video'; group_id: string }) {
    return request.post<void>('/media/groups/move', data)
  },

  // 智能裁剪
  smartCrop(params: { video_id: string; target_ratio: string }) {
    return request.post<{ video_id: string; url: string }>('/media/videos/smart-crop', params)
  },

  // 视频封面
  extractCovers(video_id: string, count?: number) {
    return request.post<{ covers: string[] }>('/media/videos/extract-covers', { video_id, count: count || 5 })
  },

  // 素材推荐
  getRecommend(params: { advertiser_id: number; material_type: 'image' | 'video'; industry_id?: number }) {
    return request.get<(ImageMaterial | VideoMaterial)[]>('/media/recommend', params)
  }
}
