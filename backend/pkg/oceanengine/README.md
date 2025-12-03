# OceanEngine SDK 封装

本目录包含巨量引擎 API 的 Go 语言封装。

## 与 sdk/ 目录的关系

项目中存在两套 SDK 实现：

| 位置 | 用途 | 特点 |
|------|------|------|
| `backend/pkg/oceanengine/` | 后端业务封装 | 简化接口，业务优先 |
| `sdk/marketing-api/` | 完整官方SDK | 全量API，类型完整 |

### 设计说明

1. **本目录 (`backend/pkg/oceanengine/`)** 
   - 对核心业务功能进行简化封装
   - 便于后端服务直接调用
   - 包含错误处理、重试逻辑等业务层面优化

2. **sdk 目录 (`sdk/marketing-api/`)**
   - 完整的官方 SDK 实现
   - 包含全量 API（870+ 个接口）
   - 提供完整的请求/响应类型定义

### 使用建议

- **简单业务场景**：直接使用本目录的封装
- **复杂业务场景**：引用 `sdk/` 目录获取完整功能

```go
// 方式一：使用简化封装
import "oceanengine-backend/pkg/oceanengine"

client := oceanengine.NewClient(appID, secret)
info, err := client.GetAdvertiserInfo(ctx, advertiserID)

// 方式二：引用完整SDK（需在 go.mod 中配置）
import "oceanengine-sdk/api/advertiser"

req := &advertiser.InfoRequest{AdvertiserIds: []uint64{123}}
resp, err := advertiser.Info(ctx, clt, accessToken, req)
```

### 整合方案（可选）

如需统一使用 `sdk/` 目录，在 `backend/go.mod` 中添加：

```go
require oceanengine-sdk v0.0.0

replace oceanengine-sdk => ../../sdk/marketing-api
```

## 文件说明

| 文件 | 功能 |
|------|------|
| `client.go` | HTTP 客户端封装 |
| `oauth.go` | OAuth2.0 授权流程 |
| `advertiser.go` | 广告主管理 |
| `agent.go` | 代理商管理 |
| `campaign.go` | 广告系列管理 |
| `ad.go` | 广告计划管理 |
| `creative.go` | 创意管理 |
| `report.go` | 数据报表 |
| `async_report.go` | 异步报表 |
| `file.go` | 素材上传 |
| `dmp.go` | DMP 人群管理 |
| `qianchuan.go` | 千川投放 |
| `enterprise.go` | 企业号管理 |
| `local.go` | 本地推广告 |
| `star.go` | 星图营销 |
| `servemarket.go` | 服务市场 |
| `v3.go` | 巨量广告升级版 |

## 完整SDK引用

如需使用完整官方SDK的全量API，可从 `sdk/` 目录引入：

- 千川 (`sdk/marketing-api/api/qianchuan/`)
- 企业号 (`sdk/marketing-api/api/enterprise/`)
- 本地推 (`sdk/marketing-api/api/local/`)
- 星图 (`sdk/marketing-api/api/star/`)
- 服务市场 (`sdk/marketing-api/api/servemarket/`)

## 配置

在 `config/settings.yml` 中配置 Ocean Engine 凭证：

```yaml
ocean:
  app_id: "your_app_id"
  secret: "your_secret"
  redirect_uri: "http://localhost:8080/api/v1/advertisers/oauth/callback"
  base_url: "https://ad.oceanengine.com/open_api"
  timeout: 30s
  retry_count: 3
```
