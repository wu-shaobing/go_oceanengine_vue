# Ocean Engine SDK 集成方案

## 概述

本文档描述如何在后端服务中集成 Ocean Engine Marketing API SDK，实现与巨量引擎广告平台的数据交互。

## SDK 封装设计

### 目录结构
```
pkg/oceanengine/
├── client.go           # 客户端主入口
├── config.go           # 配置定义
├── token.go            # Token 管理
├── request.go          # 请求封装
├── response.go         # 响应处理
├── errors.go           # 错误定义
├── types.go            # 公共类型
│
├── advertiser/         # 广告主相关
│   ├── advertiser.go
│   ├── fund.go
│   └── qualification.go
│
├── campaign/           # 广告系列相关
│   ├── campaign.go
│   └── types.go
│
├── ad/                 # 广告组相关
│   ├── ad.go
│   └── types.go
│
├── creative/           # 创意相关
│   ├── creative.go
│   └── types.go
│
├── material/           # 素材相关
│   ├── image.go
│   ├── video.go
│   └── types.go
│
├── audience/           # 人群定向
│   ├── package.go
│   ├── custom.go
│   └── types.go
│
├── report/             # 数据报告
│   ├── report.go
│   └── types.go
│
└── tools/              # 工具接口
    ├── region.go
    ├── interest.go
    └── behavior.go
```

### 核心组件

#### 1. 客户端 (Client)

```go
// pkg/oceanengine/client.go
package oceanengine

import (
    "context"
    "net/http"
    "sync"
    "time"
)

// Client Ocean Engine API 客户端
type Client struct {
    config      *Config
    httpClient  *http.Client
    tokenStore  TokenStore
    rateLimiter *RateLimiter
    
    // 各模块 API
    Advertiser  *AdvertiserAPI
    Campaign    *CampaignAPI
    Ad          *AdAPI
    Creative    *CreativeAPI
    Material    *MaterialAPI
    Audience    *AudienceAPI
    Report      *ReportAPI
    Tools       *ToolsAPI
}

// Config 客户端配置
type Config struct {
    AppID       string        // 应用ID
    Secret      string        // 应用密钥
    BaseURL     string        // API 基础地址
    Timeout     time.Duration // 请求超时
    RetryCount  int           // 重试次数
    RetryDelay  time.Duration // 重试间隔
    Debug       bool          // 调试模式
}

// NewClient 创建客户端
func NewClient(cfg *Config, tokenStore TokenStore) *Client {
    if cfg.BaseURL == "" {
        cfg.BaseURL = "https://ad.oceanengine.com/open_api"
    }
    if cfg.Timeout == 0 {
        cfg.Timeout = 30 * time.Second
    }
    if cfg.RetryCount == 0 {
        cfg.RetryCount = 3
    }
    
    c := &Client{
        config:     cfg,
        httpClient: &http.Client{Timeout: cfg.Timeout},
        tokenStore: tokenStore,
        rateLimiter: NewRateLimiter(100), // 100 QPS
    }
    
    // 初始化各模块 API
    c.Advertiser = &AdvertiserAPI{client: c}
    c.Campaign = &CampaignAPI{client: c}
    c.Ad = &AdAPI{client: c}
    c.Creative = &CreativeAPI{client: c}
    c.Material = &MaterialAPI{client: c}
    c.Audience = &AudienceAPI{client: c}
    c.Report = &ReportAPI{client: c}
    c.Tools = &ToolsAPI{client: c}
    
    return c
}

// WithAdvertiser 设置当前操作的广告主
func (c *Client) WithAdvertiser(advertiserID int64) *ClientContext {
    return &ClientContext{
        client:       c,
        advertiserID: advertiserID,
    }
}

// ClientContext 带广告主上下文的客户端
type ClientContext struct {
    client       *Client
    advertiserID int64
}
```

#### 2. Token 管理

```go
// pkg/oceanengine/token.go
package oceanengine

import (
    "context"
    "sync"
    "time"
)

// TokenStore Token 存储接口
type TokenStore interface {
    // GetToken 获取 Token
    GetToken(ctx context.Context, advertiserID int64) (*Token, error)
    
    // SaveToken 保存 Token
    SaveToken(ctx context.Context, advertiserID int64, token *Token) error
    
    // RefreshToken 刷新 Token
    RefreshToken(ctx context.Context, advertiserID int64) (*Token, error)
}

// Token 访问令牌
type Token struct {
    AccessToken           string    `json:"access_token"`
    RefreshToken          string    `json:"refresh_token"`
    AccessTokenExpiresIn  int64     `json:"access_token_expires_in"`
    RefreshTokenExpiresIn int64     `json:"refresh_token_expires_in"`
    ExpiresAt             time.Time `json:"expires_at"`
}

// IsExpired 检查 Token 是否过期
func (t *Token) IsExpired() bool {
    return time.Now().After(t.ExpiresAt.Add(-5 * time.Minute)) // 提前5分钟刷新
}

// DBTokenStore 数据库存储实现
type DBTokenStore struct {
    db     *gorm.DB
    client *Client
    mu     sync.RWMutex
    cache  map[int64]*Token // 内存缓存
}

func NewDBTokenStore(db *gorm.DB, client *Client) *DBTokenStore {
    return &DBTokenStore{
        db:     db,
        client: client,
        cache:  make(map[int64]*Token),
    }
}

func (s *DBTokenStore) GetToken(ctx context.Context, advertiserID int64) (*Token, error) {
    // 1. 先从缓存获取
    s.mu.RLock()
    if token, ok := s.cache[advertiserID]; ok && !token.IsExpired() {
        s.mu.RUnlock()
        return token, nil
    }
    s.mu.RUnlock()
    
    // 2. 从数据库获取
    var adv model.Advertiser
    err := s.db.WithContext(ctx).
        Select("access_token", "refresh_token", "token_expire_at").
        Where("advertiser_id = ?", advertiserID).
        First(&adv).Error
    if err != nil {
        return nil, err
    }
    
    token := &Token{
        AccessToken:  adv.AccessToken,
        RefreshToken: adv.RefreshToken,
        ExpiresAt:    adv.TokenExpireAt,
    }
    
    // 3. 检查是否需要刷新
    if token.IsExpired() {
        return s.RefreshToken(ctx, advertiserID)
    }
    
    // 4. 更新缓存
    s.mu.Lock()
    s.cache[advertiserID] = token
    s.mu.Unlock()
    
    return token, nil
}

func (s *DBTokenStore) RefreshToken(ctx context.Context, advertiserID int64) (*Token, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    // 再次检查是否已经刷新
    if token, ok := s.cache[advertiserID]; ok && !token.IsExpired() {
        return token, nil
    }
    
    // 从数据库获取 refresh_token
    var adv model.Advertiser
    err := s.db.WithContext(ctx).
        Select("refresh_token").
        Where("advertiser_id = ?", advertiserID).
        First(&adv).Error
    if err != nil {
        return nil, err
    }
    
    // 调用 API 刷新 Token
    newToken, err := s.client.RefreshAccessToken(ctx, adv.RefreshToken)
    if err != nil {
        return nil, fmt.Errorf("refresh token failed: %w", err)
    }
    
    // 保存到数据库
    err = s.SaveToken(ctx, advertiserID, newToken)
    if err != nil {
        return nil, err
    }
    
    // 更新缓存
    s.cache[advertiserID] = newToken
    
    return newToken, nil
}

func (s *DBTokenStore) SaveToken(ctx context.Context, advertiserID int64, token *Token) error {
    return s.db.WithContext(ctx).
        Model(&model.Advertiser{}).
        Where("advertiser_id = ?", advertiserID).
        Updates(map[string]interface{}{
            "access_token":    token.AccessToken,
            "refresh_token":   token.RefreshToken,
            "token_expire_at": token.ExpiresAt,
            "updated_at":      time.Now(),
        }).Error
}

// RefreshAccessToken 刷新访问令牌
func (c *Client) RefreshAccessToken(ctx context.Context, refreshToken string) (*Token, error) {
    resp, err := c.Post(ctx, "/oauth2/refresh_token/", map[string]interface{}{
        "app_id":        c.config.AppID,
        "secret":        c.config.Secret,
        "grant_type":    "refresh_token",
        "refresh_token": refreshToken,
    })
    if err != nil {
        return nil, err
    }
    
    var result struct {
        Data Token `json:"data"`
    }
    if err := resp.Decode(&result); err != nil {
        return nil, err
    }
    
    result.Data.ExpiresAt = time.Now().Add(
        time.Duration(result.Data.AccessTokenExpiresIn) * time.Second,
    )
    
    return &result.Data, nil
}
```

#### 3. 请求封装

```go
// pkg/oceanengine/request.go
package oceanengine

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

// Request 请求封装
type Request struct {
    Method      string
    Path        string
    Query       url.Values
    Body        interface{}
    Headers     map[string]string
    AccessToken string
}

// Response 响应封装
type Response struct {
    Code      int             `json:"code"`
    Message   string          `json:"message"`
    Data      json.RawMessage `json:"data"`
    RequestID string          `json:"request_id"`
}

// Do 执行请求
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
    // 1. 限流控制
    if err := c.rateLimiter.Wait(ctx); err != nil {
        return nil, fmt.Errorf("rate limit exceeded: %w", err)
    }
    
    // 2. 构建 URL
    u, err := url.Parse(c.config.BaseURL + req.Path)
    if err != nil {
        return nil, err
    }
    if req.Query != nil {
        u.RawQuery = req.Query.Encode()
    }
    
    // 3. 构建请求体
    var bodyReader io.Reader
    if req.Body != nil {
        bodyBytes, err := json.Marshal(req.Body)
        if err != nil {
            return nil, err
        }
        bodyReader = bytes.NewReader(bodyBytes)
    }
    
    // 4. 创建 HTTP 请求
    httpReq, err := http.NewRequestWithContext(ctx, req.Method, u.String(), bodyReader)
    if err != nil {
        return nil, err
    }
    
    // 5. 设置请求头
    httpReq.Header.Set("Content-Type", "application/json")
    if req.AccessToken != "" {
        httpReq.Header.Set("Access-Token", req.AccessToken)
    }
    for k, v := range req.Headers {
        httpReq.Header.Set(k, v)
    }
    
    // 6. 执行请求（带重试）
    var resp *http.Response
    for i := 0; i <= c.config.RetryCount; i++ {
        resp, err = c.httpClient.Do(httpReq)
        if err == nil {
            break
        }
        if i < c.config.RetryCount {
            time.Sleep(c.config.RetryDelay)
        }
    }
    if err != nil {
        return nil, fmt.Errorf("request failed after %d retries: %w", c.config.RetryCount, err)
    }
    defer resp.Body.Close()
    
    // 7. 读取响应
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    // 8. 解析响应
    var result Response
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }
    
    // 9. 检查业务错误
    if result.Code != 0 {
        return nil, &APIError{
            Code:      result.Code,
            Message:   result.Message,
            RequestID: result.RequestID,
        }
    }
    
    return &result, nil
}

// Get 发起 GET 请求
func (c *Client) Get(ctx context.Context, path string, query url.Values) (*Response, error) {
    return c.Do(ctx, &Request{
        Method: http.MethodGet,
        Path:   path,
        Query:  query,
    })
}

// GetWithToken 发起带 Token 的 GET 请求
func (c *Client) GetWithToken(ctx context.Context, advertiserID int64, path string, query url.Values) (*Response, error) {
    token, err := c.tokenStore.GetToken(ctx, advertiserID)
    if err != nil {
        return nil, err
    }
    
    return c.Do(ctx, &Request{
        Method:      http.MethodGet,
        Path:        path,
        Query:       query,
        AccessToken: token.AccessToken,
    })
}

// Post 发起 POST 请求
func (c *Client) Post(ctx context.Context, path string, body interface{}) (*Response, error) {
    return c.Do(ctx, &Request{
        Method: http.MethodPost,
        Path:   path,
        Body:   body,
    })
}

// PostWithToken 发起带 Token 的 POST 请求
func (c *Client) PostWithToken(ctx context.Context, advertiserID int64, path string, body interface{}) (*Response, error) {
    token, err := c.tokenStore.GetToken(ctx, advertiserID)
    if err != nil {
        return nil, err
    }
    
    return c.Do(ctx, &Request{
        Method:      http.MethodPost,
        Path:        path,
        Body:        body,
        AccessToken: token.AccessToken,
    })
}
```

#### 4. 广告主 API

```go
// pkg/oceanengine/advertiser/advertiser.go
package advertiser

import (
    "context"
    "net/url"
)

type AdvertiserAPI struct {
    client *oceanengine.Client
}

// Info 广告主信息
type Info struct {
    ID           int64   `json:"id"`
    Name         string  `json:"name"`
    Role         string  `json:"role"`
    Status       string  `json:"status"`
    Company      string  `json:"company"`
    Industry     string  `json:"industry"`
    Balance      float64 `json:"balance"`
    ValidBalance float64 `json:"valid_balance"`
    CashBalance  float64 `json:"cash_balance"`
    CreateTime   string  `json:"create_time"`
}

// GetInfo 获取广告主信息
func (a *AdvertiserAPI) GetInfo(ctx context.Context, advertiserID int64) (*Info, error) {
    query := url.Values{}
    query.Set("advertiser_ids", fmt.Sprintf("[%d]", advertiserID))
    
    resp, err := a.client.GetWithToken(ctx, advertiserID, "/2/advertiser/info/", query)
    if err != nil {
        return nil, err
    }
    
    var result struct {
        List []*Info `json:"list"`
    }
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return nil, err
    }
    
    if len(result.List) == 0 {
        return nil, errors.New("advertiser not found")
    }
    
    return result.List[0], nil
}

// Balance 账户余额
type Balance struct {
    Balance           float64 `json:"balance"`
    ValidBalance      float64 `json:"valid_balance"`
    CashBalance       float64 `json:"cash_balance"`
    GrantBalance      float64 `json:"grant_balance"`
    GrantValidBalance float64 `json:"grant_valid_balance"`
}

// GetBalance 获取账户余额
func (a *AdvertiserAPI) GetBalance(ctx context.Context, advertiserID int64) (*Balance, error) {
    query := url.Values{}
    query.Set("advertiser_id", fmt.Sprintf("%d", advertiserID))
    
    resp, err := a.client.GetWithToken(ctx, advertiserID, "/2/advertiser/fund/get/", query)
    if err != nil {
        return nil, err
    }
    
    var result Balance
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return nil, err
    }
    
    return &result, nil
}

// FundTransaction 资金流水
type FundTransaction struct {
    TransactionSeq  string  `json:"transaction_seq"`
    TransactionType string  `json:"transaction_type"`
    Amount          float64 `json:"amount"`
    Cash            float64 `json:"cash"`
    Grant           float64 `json:"grant"`
    TransactionTime string  `json:"transaction_time"`
    Remitter        string  `json:"remitter"`
}

// GetFundTransactions 获取资金流水
func (a *AdvertiserAPI) GetFundTransactions(ctx context.Context, req *FundTransactionReq) ([]*FundTransaction, error) {
    query := url.Values{}
    query.Set("advertiser_id", fmt.Sprintf("%d", req.AdvertiserID))
    query.Set("start_date", req.StartDate)
    query.Set("end_date", req.EndDate)
    if req.TransactionType != "" {
        query.Set("transaction_type", req.TransactionType)
    }
    query.Set("page", fmt.Sprintf("%d", req.Page))
    query.Set("page_size", fmt.Sprintf("%d", req.PageSize))
    
    resp, err := a.client.GetWithToken(ctx, req.AdvertiserID, "/2/advertiser/fund/transaction/get/", query)
    if err != nil {
        return nil, err
    }
    
    var result struct {
        List []*FundTransaction `json:"list"`
    }
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return nil, err
    }
    
    return result.List, nil
}
```

#### 5. 广告系列 API

```go
// pkg/oceanengine/campaign/campaign.go
package campaign

import (
    "context"
    "encoding/json"
    "net/url"
)

type CampaignAPI struct {
    client *oceanengine.Client
}

// Campaign 广告系列
type Campaign struct {
    ID                 int64   `json:"id"`
    Name               string  `json:"name"`
    BudgetMode         string  `json:"budget_mode"`
    Budget             float64 `json:"budget"`
    LandingType        string  `json:"landing_type"`
    MarketingGoal      string  `json:"marketing_goal"`
    Status             string  `json:"status"`
    OptStatus          string  `json:"opt_status"`
    DeliveryRelatedNum string  `json:"delivery_related_num"`
    CreateTime         string  `json:"create_time"`
    ModifyTime         string  `json:"modify_time"`
}

// ListReq 列表请求
type ListReq struct {
    AdvertiserID int64    `json:"advertiser_id"`
    IDs          []int64  `json:"ids,omitempty"`
    Filtering    *Filter  `json:"filtering,omitempty"`
    Page         int      `json:"page"`
    PageSize     int      `json:"page_size"`
}

// Filter 筛选条件
type Filter struct {
    Status           string   `json:"status,omitempty"`
    LandingType      string   `json:"landing_type,omitempty"`
    CampaignName     string   `json:"campaign_name,omitempty"`
    CampaignCreateTime string `json:"campaign_create_time,omitempty"`
}

// List 获取广告系列列表
func (a *CampaignAPI) List(ctx context.Context, req *ListReq) ([]*Campaign, *PageInfo, error) {
    resp, err := a.client.PostWithToken(ctx, req.AdvertiserID, "/2/campaign/get/", req)
    if err != nil {
        return nil, nil, err
    }
    
    var result struct {
        List     []*Campaign `json:"list"`
        PageInfo *PageInfo   `json:"page_info"`
    }
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return nil, nil, err
    }
    
    return result.List, result.PageInfo, nil
}

// CreateReq 创建请求
type CreateReq struct {
    AdvertiserID       int64   `json:"advertiser_id"`
    CampaignName       string  `json:"campaign_name"`
    BudgetMode         string  `json:"budget_mode"`
    Budget             float64 `json:"budget,omitempty"`
    LandingType        string  `json:"landing_type"`
    MarketingGoal      string  `json:"marketing_goal,omitempty"`
    DeliveryRelatedNum string  `json:"delivery_related_num,omitempty"`
}

// Create 创建广告系列
func (a *CampaignAPI) Create(ctx context.Context, req *CreateReq) (int64, error) {
    resp, err := a.client.PostWithToken(ctx, req.AdvertiserID, "/2/campaign/create/", req)
    if err != nil {
        return 0, err
    }
    
    var result struct {
        CampaignID int64 `json:"campaign_id"`
    }
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return 0, err
    }
    
    return result.CampaignID, nil
}

// UpdateReq 更新请求
type UpdateReq struct {
    AdvertiserID int64   `json:"advertiser_id"`
    CampaignID   int64   `json:"campaign_id"`
    CampaignName string  `json:"campaign_name,omitempty"`
    Budget       float64 `json:"budget,omitempty"`
    BudgetMode   string  `json:"budget_mode,omitempty"`
}

// Update 更新广告系列
func (a *CampaignAPI) Update(ctx context.Context, req *UpdateReq) error {
    _, err := a.client.PostWithToken(ctx, req.AdvertiserID, "/2/campaign/update/", req)
    return err
}

// UpdateStatusReq 更新状态请求
type UpdateStatusReq struct {
    AdvertiserID int64   `json:"advertiser_id"`
    CampaignIDs  []int64 `json:"campaign_ids"`
    OptStatus    string  `json:"opt_status"` // ENABLE, DISABLE, DELETE
}

// UpdateStatus 更新广告系列状态
func (a *CampaignAPI) UpdateStatus(ctx context.Context, req *UpdateStatusReq) error {
    _, err := a.client.PostWithToken(ctx, req.AdvertiserID, "/2/campaign/update/status/", req)
    return err
}
```

#### 6. 数据报告 API

```go
// pkg/oceanengine/report/report.go
package report

import (
    "context"
    "encoding/json"
)

type ReportAPI struct {
    client *oceanengine.Client
}

// IntegratedReq 综合报告请求
type IntegratedReq struct {
    AdvertiserID    int64    `json:"advertiser_id"`
    StartDate       string   `json:"start_date"`        // YYYY-MM-DD
    EndDate         string   `json:"end_date"`          // YYYY-MM-DD
    TimeGranularity string   `json:"time_granularity"`  // STAT_TIME_GRANULARITY_DAILY
    GroupBy         []string `json:"group_by"`          // ["STAT_GROUP_BY_FIELD_ID", "STAT_GROUP_BY_FIELD_STAT_TIME"]
    OrderField      string   `json:"order_field,omitempty"`
    OrderType       string   `json:"order_type,omitempty"`
    Page            int      `json:"page"`
    PageSize        int      `json:"page_size"`
    Filtering       *IntegratedFilter `json:"filtering,omitempty"`
}

// IntegratedFilter 筛选条件
type IntegratedFilter struct {
    CampaignIDs  []int64 `json:"campaign_ids,omitempty"`
    AdIDs        []int64 `json:"ad_ids,omitempty"`
    CreativeIDs  []int64 `json:"creative_ids,omitempty"`
}

// ReportData 报告数据
type ReportData struct {
    StatDatetime      string  `json:"stat_datetime"`
    AdvertiserID      int64   `json:"advertiser_id"`
    CampaignID        int64   `json:"campaign_id,omitempty"`
    CampaignName      string  `json:"campaign_name,omitempty"`
    AdID              int64   `json:"ad_id,omitempty"`
    AdName            string  `json:"ad_name,omitempty"`
    
    // 核心指标
    Cost              float64 `json:"cost"`
    ShowCnt           int64   `json:"show"`
    ClickCnt          int64   `json:"click"`
    Ctr               float64 `json:"ctr"`
    AvgClickCost      float64 `json:"avg_click_cost"`
    AvgShowCost       float64 `json:"avg_show_cost"`
    
    // 转化指标
    ConvertCnt        int64   `json:"convert"`
    ConvertCost       float64 `json:"convert_cost"`
    ConvertRate       float64 `json:"convert_rate"`
    DeepConvertCnt    int64   `json:"deep_convert"`
    DeepConvertCost   float64 `json:"deep_convert_cost"`
    DeepConvertRate   float64 `json:"deep_convert_rate"`
    
    // 互动指标
    LikeCnt           int64   `json:"like"`
    CommentCnt        int64   `json:"comment"`
    ShareCnt          int64   `json:"share"`
    FollowCnt         int64   `json:"follow"`
}

// GetIntegrated 获取综合报告
func (a *ReportAPI) GetIntegrated(ctx context.Context, req *IntegratedReq) ([]*ReportData, *PageInfo, error) {
    resp, err := a.client.PostWithToken(ctx, req.AdvertiserID, "/2/report/integrated/get/", req)
    if err != nil {
        return nil, nil, err
    }
    
    var result struct {
        List     []*ReportData `json:"list"`
        PageInfo *PageInfo     `json:"page_info"`
    }
    if err := json.Unmarshal(resp.Data, &result); err != nil {
        return nil, nil, err
    }
    
    return result.List, result.PageInfo, nil
}

// GetAdvertiserReport 获取广告主维度报告
func (a *ReportAPI) GetAdvertiserReport(ctx context.Context, req *IntegratedReq) ([]*ReportData, error) {
    req.GroupBy = []string{"STAT_GROUP_BY_FIELD_STAT_TIME"}
    list, _, err := a.GetIntegrated(ctx, req)
    return list, err
}

// GetCampaignReport 获取广告系列维度报告
func (a *ReportAPI) GetCampaignReport(ctx context.Context, req *IntegratedReq) ([]*ReportData, error) {
    req.GroupBy = []string{"STAT_GROUP_BY_FIELD_ID", "STAT_GROUP_BY_FIELD_STAT_TIME"}
    list, _, err := a.GetIntegrated(ctx, req)
    return list, err
}
```

---

## 使用示例

### 1. 初始化客户端

```go
// 在应用启动时初始化
func initOceanEngine(cfg *config.OceanConfig, db *gorm.DB) *oceanengine.Client {
    clientCfg := &oceanengine.Config{
        AppID:      cfg.AppID,
        Secret:     cfg.Secret,
        Timeout:    30 * time.Second,
        RetryCount: 3,
        RetryDelay: time.Second,
    }
    
    client := oceanengine.NewClient(clientCfg, nil)
    tokenStore := oceanengine.NewDBTokenStore(db, client)
    client.SetTokenStore(tokenStore)
    
    return client
}
```

### 2. 在 Service 中使用

```go
// internal/app/advertiser/service/advertiser.go
type AdvertiserService struct {
    repo     repository.AdvertiserRepository
    oceanSDK *oceanengine.Client
}

func (s *AdvertiserService) SyncAdvertiserInfo(ctx context.Context, advertiserID int64) error {
    // 调用 SDK 获取最新数据
    info, err := s.oceanSDK.Advertiser.GetInfo(ctx, advertiserID)
    if err != nil {
        return fmt.Errorf("failed to get advertiser info: %w", err)
    }
    
    balance, err := s.oceanSDK.Advertiser.GetBalance(ctx, advertiserID)
    if err != nil {
        return fmt.Errorf("failed to get balance: %w", err)
    }
    
    // 更新本地数据
    return s.repo.Update(ctx, &model.Advertiser{
        AdvertiserID:  advertiserID,
        Name:          info.Name,
        Company:       info.Company,
        Status:        info.Status,
        Balance:       balance.Balance,
        ValidBalance:  balance.ValidBalance,
        CashBalance:   balance.CashBalance,
        LastSyncAt:    time.Now(),
    })
}
```

### 3. 创建广告系列

```go
func (s *CampaignService) Create(ctx context.Context, req *dto.CampaignCreateReq) (int64, error) {
    // 调用 SDK 创建
    campaignID, err := s.oceanSDK.Campaign.Create(ctx, &campaign.CreateReq{
        AdvertiserID:  req.AdvertiserID,
        CampaignName:  req.Name,
        BudgetMode:    req.BudgetMode,
        Budget:        req.Budget,
        LandingType:   req.LandingType,
        MarketingGoal: req.MarketingGoal,
    })
    if err != nil {
        return 0, fmt.Errorf("failed to create campaign: %w", err)
    }
    
    // 保存到本地数据库
    localCampaign := &model.Campaign{
        CampaignID:   campaignID,
        AdvertiserID: req.AdvertiserID,
        Name:         req.Name,
        BudgetMode:   req.BudgetMode,
        Budget:       req.Budget,
        LandingType:  req.LandingType,
        Status:       "CAMPAIGN_STATUS_ENABLE",
        LastSyncAt:   time.Now(),
    }
    
    if err := s.repo.Create(ctx, localCampaign); err != nil {
        return 0, err
    }
    
    return localCampaign.ID, nil
}
```

---

## 错误处理

```go
// pkg/oceanengine/errors.go
package oceanengine

import "fmt"

// APIError Ocean Engine API 错误
type APIError struct {
    Code      int    `json:"code"`
    Message   string `json:"message"`
    RequestID string `json:"request_id"`
}

func (e *APIError) Error() string {
    return fmt.Sprintf("ocean engine api error: code=%d, message=%s, request_id=%s",
        e.Code, e.Message, e.RequestID)
}

// 常见错误码
const (
    ErrCodeInvalidToken     = 40001 // Token 无效
    ErrCodeTokenExpired     = 40002 // Token 过期
    ErrCodeRateLimit        = 40100 // 请求频率限制
    ErrCodeInvalidParam     = 40003 // 参数错误
    ErrCodeNoPermission     = 40004 // 无权限
    ErrCodeResourceNotFound = 40006 // 资源不存在
)

// IsTokenError 判断是否为 Token 错误
func IsTokenError(err error) bool {
    if apiErr, ok := err.(*APIError); ok {
        return apiErr.Code == ErrCodeInvalidToken || apiErr.Code == ErrCodeTokenExpired
    }
    return false
}

// IsRateLimitError 判断是否为限流错误
func IsRateLimitError(err error) bool {
    if apiErr, ok := err.(*APIError); ok {
        return apiErr.Code == ErrCodeRateLimit
    }
    return false
}
```

---

## 限流控制

```go
// pkg/oceanengine/ratelimit.go
package oceanengine

import (
    "context"
    "golang.org/x/time/rate"
)

// RateLimiter 限流器
type RateLimiter struct {
    limiter *rate.Limiter
}

// NewRateLimiter 创建限流器
func NewRateLimiter(qps int) *RateLimiter {
    return &RateLimiter{
        limiter: rate.NewLimiter(rate.Limit(qps), qps*2), // 突发量为 QPS 的 2 倍
    }
}

// Wait 等待获取令牌
func (r *RateLimiter) Wait(ctx context.Context) error {
    return r.limiter.Wait(ctx)
}
```
