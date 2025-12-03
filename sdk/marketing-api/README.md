# OceanEngine Marketing API Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](../LICENSE)

巨量引擎营销 API Go SDK，支持巨量引擎、巨量千川、企业号、本地推、星图、服务市场等平台。

## 安装

```bash
go get github.com/bububa/oceanengine/marketing-api
```

## 快速开始

### 1. 初始化客户端

```go
import "github.com/bububa/oceanengine/marketing-api/core"

// appID 为 uint64 类型
var appID uint64 = 12345678
client := core.NewSDKClient(appID, "your_app_secret")
```

### 2. OAuth2 授权

```go
import (
    "context"
    "github.com/bububa/oceanengine/marketing-api/api/oauth"
)

// 生成授权链接
authURL := oauth.Url(client, "https://your-callback.com/callback", "state", false)

// 使用授权码获取 AccessToken
ctx := context.Background()
tokenResp, err := oauth.AccessToken(ctx, client, "auth_code")

// 刷新 Token
newTokenResp, err := oauth.RefreshToken(ctx, client, tokenResp.RefreshToken)
```

### 3. 调用 API

```go
import (
    "github.com/bububa/oceanengine/marketing-api/api/advertiser"
    "github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// 获取广告主信息
req := &advertiser.InfoRequest{
    AdvertiserIDs: []uint64{12345},
}
infos, err := advertiser.Info(ctx, client, accessToken, req)
```

## 支持的平台

| 平台 | 文档 | 示例 |
|------|------|------|
| 巨量引擎 | [OCEANENGINE.md](./OCEANENGINE.md) | [examples/oceanengine](./examples/oceanengine) |
| 巨量千川 | [QIANCHUAN.md](./QIANCHUAN.md) | [examples/qianchuan](./examples/qianchuan) |
| 企业号 | [ENTERPRISE.md](./ENTERPRISE.md) | [examples/enterprise](./examples/enterprise) |
| 本地推 | [LOCAL.md](./LOCAL.md) | [examples/local](./examples/local) |
| 星图 | [STAR.md](./STAR.md) | [examples/star](./examples/star) |
| 服务市场 | [SERVE_MARKET.md](./SERVE_MARKET.md) | - |

## 项目结构

```
marketing-api/
├── api/            # API 实现
│   ├── oauth/      # OAuth 授权
│   ├── advertiser/ # 广告主管理
│   ├── ad/         # 广告计划
│   ├── campaign/   # 广告组
│   ├── report/     # 数据报表
│   ├── qianchuan/  # 千川 API
│   ├── enterprise/ # 企业号 API
│   ├── local/      # 本地推 API
│   ├── star/       # 星图 API
│   └── ...
├── core/           # SDK 核心
├── enum/           # 枚举定义
├── model/          # 请求/响应模型
├── examples/       # 使用示例
├── scripts/        # 工具脚本
├── CHANGELOG.md    # 变更日志
└── VERSION         # 版本号
```

## 主要功能

### 巨量引擎
- OAuth2 授权
- 广告主管理
- 广告组/计划/创意管理
- 定向包管理
- 素材库管理
- 数据报表
- 转化追踪

### 巨量千川
- 店铺管理
- 商品管理
- 电商广告投放
- 直播间推广
- 短视频推广
- 数据报表

### 企业号
- 账号管理
- 线索管理
- 粉丝管理
- 数据统计

### 本地推
- 门店管理
- 团购/优惠券
- 本地广告投放

### 星图
- 达人搜索
- 任务管理
- 线索管理
- 数据分析

## 环境变量配置

建议使用环境变量管理敏感信息：

```bash
export OCEANENGINE_APP_ID="your_app_id"
export OCEANENGINE_APP_SECRET="your_app_secret"
export OCEANENGINE_ACCESS_TOKEN="your_access_token"
```

## 开发工具

### 文档同步检查

检查文档与代码实现是否一致：

```bash
./scripts/check_doc_sync.sh
```

### 代码检查

```bash
go vet ./...
go build ./...
```

## 注意事项

1. **Token 管理**: AccessToken 有效期为 24 小时，请在过期前使用 RefreshToken 刷新
2. **并发限制**: 请遵守 API 调用频率限制
3. **错误处理**: 所有 API 调用都应检查返回的 error
4. **Context**: 大多数 API 调用需要传入 `context.Context` 参数

## 变更日志

查看 [CHANGELOG.md](./CHANGELOG.md) 了解版本变更记录。

## License

本项目采用 Apache License 2.0 许可证，详见 [LICENSE](../LICENSE) 文件。
