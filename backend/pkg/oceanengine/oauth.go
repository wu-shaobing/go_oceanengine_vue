package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// OAuthService OAuth 服务
type OAuthService struct {
	client *Client
}

// NewOAuthService 创建 OAuth 服务
func NewOAuthService(client *Client) *OAuthService {
	return &OAuthService{client: client}
}

// AccessTokenRequest 获取 Access Token 请求
type AccessTokenRequest struct {
	AppID     string `json:"app_id"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
	AuthCode  string `json:"auth_code,omitempty"`
}

// AccessTokenResponse Access Token 响应
type AccessTokenResponse struct {
	AccessToken           string  `json:"access_token"`
	RefreshToken          string  `json:"refresh_token"`
	ExpiresIn             int64   `json:"expires_in"`
	RefreshTokenExpiresIn int64   `json:"refresh_token_expires_in"`
	AdvertiserIDs         []int64 `json:"advertiser_ids"`
}

// GetAccessToken 获取 Access Token
func (s *OAuthService) GetAccessToken(ctx context.Context, authCode string) (*AccessTokenResponse, error) {
	req := AccessTokenRequest{
		AppID:     s.client.appID,
		Secret:    s.client.secret,
		GrantType: "auth_code",
		AuthCode:  authCode,
	}

	resp, err := s.client.Post(ctx, "/oauth2/access_token/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AccessTokenResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// RefreshTokenRequest 刷新 Token 请求
type RefreshTokenRequest struct {
	AppID        string `json:"app_id"`
	Secret       string `json:"secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshAccessToken 刷新 Access Token
func (s *OAuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResponse, error) {
	req := RefreshTokenRequest{
		AppID:        s.client.appID,
		Secret:       s.client.secret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}

	resp, err := s.client.Post(ctx, "/oauth2/refresh_token/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AccessTokenResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// GetAuthURL 获取授权 URL
func (s *OAuthService) GetAuthURL(state, redirectURI string) string {
	return fmt.Sprintf(
		"https://ad.oceanengine.com/openapi/audit/oauth.html?app_id=%s&state=%s&redirect_uri=%s",
		s.client.appID,
		state,
		redirectURI,
	)
}

// GetAuthURLWithScope 获取带权限范围的授权 URL
func (s *OAuthService) GetAuthURLWithScope(state, redirectURI string, scope []string, materialAuth bool) string {
	scopeStr := ""
	if len(scope) > 0 {
		for i, sc := range scope {
			if i > 0 {
				scopeStr += ","
			}
			scopeStr += sc
		}
	}
	authURL := fmt.Sprintf(
		"https://ad.oceanengine.com/openapi/audit/oauth.html?app_id=%s&state=%s&redirect_uri=%s",
		s.client.appID,
		state,
		redirectURI,
	)
	if scopeStr != "" {
		authURL += "&scope=" + scopeStr
	}
	if materialAuth {
		authURL += "&material_auth=1"
	}
	return authURL
}

// ==================== 用户信息 ====================

// UserInfo 用户信息
type UserInfo struct {
	ID          uint64 `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

// GetUserInfo 获取授权用户信息
func (s *OAuthService) GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error) {
	path := "/oauth2/user_info/"
	params := map[string]interface{}{
		"access_token": accessToken,
	}

	resp, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		Data UserInfo `json:"data"`
	}
	if err := json.Unmarshal(resp.Data, &result.Data); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result.Data, nil
}

// ==================== 已授权账户 ====================

// Advertiser 广告主信息
type Advertiser struct {
	AdvertiserID   uint64 `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	IsValid        bool   `json:"is_valid"`
	AccountRole    string `json:"account_role"`
}

// GetAdvertiserList 获取已授权的广告主列表
func (s *OAuthService) GetAdvertiserList(ctx context.Context, accessToken string, appID string) ([]Advertiser, error) {
	path := "/oauth2/advertiser/get/"
	params := map[string]interface{}{
		"access_token": accessToken,
		"app_id":       appID,
	}

	resp, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		List []Advertiser `json:"list"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.List, nil
}

// ==================== App Access Token ====================

// AppAccessTokenResponse App Access Token 响应
type AppAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetAppAccessToken 获取 App Access Token (应用级别，非用户级别)
func (s *OAuthService) GetAppAccessToken(ctx context.Context) (*AppAccessTokenResponse, error) {
	req := map[string]string{
		"app_id":     s.client.appID,
		"secret":     s.client.secret,
		"grant_type": "client_credentials",
	}

	resp, err := s.client.Post(ctx, "/oauth2/app_access_token/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result AppAccessTokenResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== Token 状态校验 ====================

// TokenStatus Token 状态
type TokenStatus struct {
	IsValid       bool    `json:"is_valid"`
	ExpiresIn     int64   `json:"expires_in"`
	AdvertiserIDs []int64 `json:"advertiser_ids"`
}

// ValidateToken 校验 Token 是否有效
func (s *OAuthService) ValidateToken(ctx context.Context, accessToken string) (*TokenStatus, error) {
	// 通过获取用户信息来判断Token是否有效
	_, err := s.GetUserInfo(ctx, accessToken)
	if err != nil {
		return &TokenStatus{IsValid: false}, nil
	}
	return &TokenStatus{IsValid: true}, nil
}
