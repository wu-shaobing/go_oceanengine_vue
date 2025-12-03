package oceanengine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	// BaseURL Ocean Engine API 基础地址
	BaseURL = "https://ad.oceanengine.com/open_api"
	// DefaultTimeout 默认超时时间
	DefaultTimeout = 30 * time.Second
)

// Client Ocean Engine API 客户端
type Client struct {
	appID       string
	secret      string
	httpClient  *http.Client
	accessToken string
}

// NewClient 创建客户端
func NewClient(appID, secret string) *Client {
	return &Client{
		appID:  appID,
		secret: secret,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// OAuth 返回OAuth服务
func (c *Client) OAuth() *OAuthService {
	return &OAuthService{client: c}
}

// V3 返回V3体验版客户端
func (c *Client) V3() *V3Client {
	return NewV3Client(c)
}

// DMP 返回DMP人群包服务
func (c *Client) DMP() *DMPService {
	return NewDMPService(c)
}

// DPA 返回DPA商品广告客户端
func (c *Client) DPA() *DPAClient {
	return NewDPAClient(c)
}

// SetAccessToken 设置 Access Token
func (c *Client) SetAccessToken(token string) {
	c.accessToken = token
}

// SetTimeout 设置超时时间
func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// BaseResponse API 基础响应
type BaseResponse struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

// IsSuccess 是否成功
func (r *BaseResponse) IsSuccess() bool {
	return r.Code == 0
}

// Get 发送 GET 请求
func (c *Client) Get(ctx context.Context, path string, params map[string]interface{}) (*BaseResponse, error) {
	reqURL, err := url.Parse(BaseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse url failed: %w", err)
	}

	query := reqURL.Query()
	for k, v := range params {
		query.Set(k, fmt.Sprintf("%v", v))
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	c.setHeaders(req)

	return c.doRequest(req)
}

// Post 发送 POST 请求
func (c *Client) Post(ctx context.Context, path string, body interface{}) (*BaseResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal body failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseURL+path, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	c.setHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(req)
}

// setHeaders 设置请求头
func (c *Client) setHeaders(req *http.Request) {
	if c.accessToken != "" {
		req.Header.Set("Access-Token", c.accessToken)
	}
}

// doRequest 执行请求
func (c *Client) doRequest(req *http.Request) (*BaseResponse, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %d: %s", resp.StatusCode, string(body))
	}

	var result BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	return &result, nil
}

// GetWithToken 发送带 Token 的 GET 请求并直接解码响应
func (c *Client) GetWithToken(ctx context.Context, accessToken, path string, params map[string]interface{}, result interface{}) error {
	reqURL, err := url.Parse(BaseURL + path)
	if err != nil {
		return fmt.Errorf("parse url failed: %w", err)
	}

	query := reqURL.Query()
	for k, v := range params {
		query.Set(k, fmt.Sprintf("%v", v))
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	if accessToken != "" {
		req.Header.Set("Access-Token", accessToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %d: %s", resp.StatusCode, string(body))
	}

	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !baseResp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", baseResp.Code, baseResp.Message)
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("unmarshal result failed: %w", err)
		}
	}

	return nil
}

// PostWithToken 发送带 Token 的 POST 请求并直接解码响应
func (c *Client) PostWithToken(ctx context.Context, accessToken, path string, data interface{}, result interface{}) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal body failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseURL+path, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	if accessToken != "" {
		req.Header.Set("Access-Token", accessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %d: %s", resp.StatusCode, string(body))
	}

	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return fmt.Errorf("unmarshal response failed: %w", err)
	}

	if !baseResp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", baseResp.Code, baseResp.Message)
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("unmarshal result failed: %w", err)
		}
	}

	return nil
}

// UploadFile 上传文件 (multipart/form-data)
func (c *Client) UploadFile(ctx context.Context, accessToken, path, fileField, filePath string, extraFields map[string]string) (*BaseResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段
	part, err := writer.CreateFormFile(fileField, filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("copy file failed: %w", err)
	}

	// 添加其他字段
	for key, val := range extraFields {
		if err := writer.WriteField(key, val); err != nil {
			return nil, fmt.Errorf("write field %s failed: %w", key, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if accessToken != "" {
		req.Header.Set("Access-Token", accessToken)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.doRequest(req)
}

// UploadFileFromReader 从Reader上传文件
func (c *Client) UploadFileFromReader(ctx context.Context, accessToken, path, fileField, fileName string, fileReader io.Reader, extraFields map[string]string) (*BaseResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段
	part, err := writer.CreateFormFile(fileField, fileName)
	if err != nil {
		return nil, fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := io.Copy(part, fileReader); err != nil {
		return nil, fmt.Errorf("copy file failed: %w", err)
	}

	// 添加其他字段
	for key, val := range extraFields {
		if err := writer.WriteField(key, val); err != nil {
			return nil, fmt.Errorf("write field %s failed: %w", key, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BaseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if accessToken != "" {
		req.Header.Set("Access-Token", accessToken)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.doRequest(req)
}
