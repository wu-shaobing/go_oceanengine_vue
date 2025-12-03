package oceanengine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// FileService 文件服务
type FileService struct {
	client *Client
}

// NewFileService 创建文件服务
func NewFileService(client *Client) *FileService {
	return &FileService{client: client}
}

// ImageInfo 图片信息
type ImageInfo struct {
	ImageID   string `json:"id"`
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Size      int64  `json:"size"`
	Format    string `json:"format"`
	Signature string `json:"signature"`
}

// ImageUploadRequest 图片上传请求
type ImageUploadRequest struct {
	AdvertiserID   int64  `json:"advertiser_id"`
	UploadType     string `json:"upload_type"` // UPLOAD_BY_FILE, UPLOAD_BY_URL
	ImageFile      string `json:"image_file,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	Filename       string `json:"filename,omitempty"`
	ImageSignature string `json:"image_signature,omitempty"` // md5
}

// UploadImageByFile 上传图片文件
func (s *FileService) UploadImageByFile(ctx context.Context, advertiserID int64, filePath string) (*ImageInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("advertiser_id", fmt.Sprintf("%d", advertiserID)); err != nil {
		return nil, fmt.Errorf("write advertiser_id failed: %w", err)
	}

	part, err := writer.CreateFormFile("image_file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("copy file content failed: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/2/file/image/ad/", body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Access-Token", s.client.accessToken)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	var apiResp BaseResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !apiResp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", apiResp.Code, apiResp.Message)
	}

	var result ImageInfo
	if err := json.Unmarshal(apiResp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// UploadImageByURL 通过URL上传图片
func (s *FileService) UploadImageByURL(ctx context.Context, advertiserID int64, imageURL string) (*ImageInfo, error) {
	reqBody := map[string]interface{}{
		"advertiser_id": advertiserID,
		"image_url":     imageURL,
	}

	resp, err := s.client.Post(ctx, "/2/file/image/ad/", reqBody)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ImageInfo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// VideoInfo 视频信息
type VideoInfo struct {
	VideoID   string  `json:"video_id"`
	URL       string  `json:"video_url"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Size      int64   `json:"size"`
	Duration  float64 `json:"duration"`
	Signature string  `json:"signature"`
	PosterURL string  `json:"poster_url"`
}

// UploadVideoByFile 上传视频文件
func (s *FileService) UploadVideoByFile(ctx context.Context, advertiserID int64, filePath string) (*VideoInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("advertiser_id", fmt.Sprintf("%d", advertiserID)); err != nil {
		return nil, fmt.Errorf("write advertiser_id failed: %w", err)
	}

	part, err := writer.CreateFormFile("video_file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("copy file content failed: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/2/file/video/ad/", body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Access-Token", s.client.accessToken)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	var apiResp BaseResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !apiResp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", apiResp.Code, apiResp.Message)
	}

	var result VideoInfo
	if err := json.Unmarshal(apiResp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ImageGetRequest 获取图片请求
type ImageGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		ImageIDs   []string `json:"image_ids,omitempty"`
		Signatures []string `json:"signatures,omitempty"`
		Width      int      `json:"width,omitempty"`
		Height     int      `json:"height,omitempty"`
		StartTime  string   `json:"start_time,omitempty"`
		EndTime    string   `json:"end_time,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// ImageGetResponse 获取图片响应
type ImageGetResponse struct {
	List     []ImageInfo `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetImages 获取图片列表
func (s *FileService) GetImages(ctx context.Context, req *ImageGetRequest) (*ImageGetResponse, error) {
	resp, err := s.client.Post(ctx, "/2/file/image/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result ImageGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// VideoGetRequest 获取视频请求
type VideoGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Filtering    struct {
		VideoIDs   []string `json:"video_ids,omitempty"`
		Signatures []string `json:"signatures,omitempty"`
		Width      int      `json:"width,omitempty"`
		Height     int      `json:"height,omitempty"`
		StartTime  string   `json:"start_time,omitempty"`
		EndTime    string   `json:"end_time,omitempty"`
	} `json:"filtering,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// VideoGetResponse 获取视频响应
type VideoGetResponse struct {
	List     []VideoInfo `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetVideos 获取视频列表
func (s *FileService) GetVideos(ctx context.Context, req *VideoGetRequest) (*VideoGetResponse, error) {
	resp, err := s.client.Post(ctx, "/2/file/video/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result VideoGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ImageAdGetRequest 获取广告图片请求
type ImageAdGetRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	ImageIDs     []string `json:"image_ids"`
}

// GetAdImages 获取广告图片
func (s *FileService) GetAdImages(ctx context.Context, req *ImageAdGetRequest) ([]ImageInfo, error) {
	resp, err := s.client.Post(ctx, "/2/file/image/ad/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []ImageInfo `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// VideoAdGetRequest 获取广告视频请求
type VideoAdGetRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	VideoIDs     []string `json:"video_ids"`
}

// GetAdVideos 获取广告视频
func (s *FileService) GetAdVideos(ctx context.Context, req *VideoAdGetRequest) ([]VideoInfo, error) {
	resp, err := s.client.Post(ctx, "/2/file/video/ad/get/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []VideoInfo `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// UploadImageByBytes 通过字节数组上传图片
func (s *FileService) UploadImageByBytes(ctx context.Context, advertiserID int64, filename string, data []byte) (*ImageInfo, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("advertiser_id", fmt.Sprintf("%d", advertiserID)); err != nil {
		return nil, fmt.Errorf("write advertiser_id failed: %w", err)
	}

	part, err := writer.CreateFormFile("image_file", filename)
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := part.Write(data); err != nil {
		return nil, fmt.Errorf("write file content failed: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/2/file/image/ad/", body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Access-Token", s.client.accessToken)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	var apiResp BaseResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !apiResp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", apiResp.Code, apiResp.Message)
	}

	var result ImageInfo
	if err := json.Unmarshal(apiResp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// UploadVideoByBytes 通过字节数组上传视频
func (s *FileService) UploadVideoByBytes(ctx context.Context, advertiserID int64, filename string, data []byte) (*VideoInfo, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("advertiser_id", fmt.Sprintf("%d", advertiserID)); err != nil {
		return nil, fmt.Errorf("write advertiser_id failed: %w", err)
	}

	part, err := writer.CreateFormFile("video_file", filename)
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := part.Write(data); err != nil {
		return nil, fmt.Errorf("write file content failed: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/2/file/video/ad/", body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Access-Token", s.client.accessToken)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	var apiResp BaseResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !apiResp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", apiResp.Code, apiResp.Message)
	}

	var result VideoInfo
	if err := json.Unmarshal(apiResp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}
