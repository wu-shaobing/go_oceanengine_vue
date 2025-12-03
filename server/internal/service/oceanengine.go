package service

import (
	"context"
	"sync"

	"github.com/bububa/oceanengine/marketing-api/api/advertiser"
	"github.com/bububa/oceanengine/marketing-api/api/oauth"
	"github.com/bububa/oceanengine/marketing-api/core"
	advertiserModel "github.com/bububa/oceanengine/marketing-api/model/advertiser"
	oauthModel "github.com/bububa/oceanengine/marketing-api/model/oauth"
)

// OceanEngineService 巨量引擎服务
type OceanEngineService struct {
	client *core.SDKClient
	mu     sync.RWMutex
}

// NewOceanEngineService 创建服务实例
func NewOceanEngineService(appID uint64, appSecret string) *OceanEngineService {
	return &OceanEngineService{
		client: core.NewSDKClient(appID, appSecret),
	}
}

// GetClient 获取SDK客户端
func (s *OceanEngineService) GetClient() *core.SDKClient {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.client
}

// --- OAuth 相关 ---

// GetAuthURL 获取授权链接
func (s *OceanEngineService) GetAuthURL(redirectURL, state string) string {
	return oauth.Url(s.client, redirectURL, state, false)
}

// GetAccessToken 使用授权码获取AccessToken
func (s *OceanEngineService) GetAccessToken(ctx context.Context, authCode string) (*oauthModel.AccessTokenResponseData, error) {
	return oauth.AccessToken(ctx, s.client, authCode)
}

// RefreshToken 刷新Token
func (s *OceanEngineService) RefreshToken(ctx context.Context, refreshToken string) (*oauthModel.AccessTokenResponseData, error) {
	return oauth.RefreshToken(ctx, s.client, refreshToken)
}

// GetAdvertisers 获取已授权的广告主列表
func (s *OceanEngineService) GetAdvertisers(ctx context.Context, accessToken string) ([]oauthModel.Advertiser, error) {
	return oauth.AdvertiserGet(ctx, s.client, accessToken)
}

// --- 广告主相关 ---

// GetAdvertiserInfo 获取广告主信息
func (s *OceanEngineService) GetAdvertiserInfo(ctx context.Context, accessToken string, advertiserIDs []uint64) ([]advertiserModel.Info, error) {
	req := &advertiserModel.InfoRequest{
		AdvertiserIDs: advertiserIDs,
	}
	return advertiser.Info(ctx, s.client, accessToken, req)
}
