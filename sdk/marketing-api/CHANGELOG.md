# Changelog

本文档记录 OceanEngine Marketing API Go SDK 的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [Unreleased]

## [1.0.0] - 2024-12-02

### Added
- 完整的巨量引擎广告投放 API 支持
  - OAuth2 授权认证
  - 广告主管理
  - 广告组/计划/创意管理
  - 定向包管理
  - 素材库管理
  - 数据报表查询
  - 转化追踪
- 巨量千川电商广告 API 支持
  - 店铺管理
  - 商品管理
  - 电商广告投放
  - 直播间/短视频推广
- 企业号 API 支持
  - 账号信息管理
  - 运营数据查询
  - 视频/评论管理
- 本地推广告 API 支持
  - 线索管理
  - POI 门店管理
  - 项目/推广管理
- 星图达人营销 API 支持
  - 账户信息查询
  - 需求/订单管理
  - 线索获取
  - 数据报表
- 服务市场 API 支持
- SDK 核心功能
  - 统一的客户端封装
  - 请求签名与验证
  - 错误处理
  - Debug 模式
  - Sandbox 环境支持
  - 限流支持
  - OpenTelemetry 链路追踪支持

### Documentation
- 添加 README.md 快速入门指南
- 添加各业务线 API 文档 (OCEANENGINE.md, QIANCHUAN.md, ENTERPRISE.md, LOCAL.md, STAR.md, SERVE_MARKET.md)
- 添加使用示例 (examples/)
- 添加文档同步检查脚本 (scripts/check_doc_sync.sh)

### Tests
- 添加 core/client_test.go 单元测试
- 添加 util/query/encode_test.go 编码测试
