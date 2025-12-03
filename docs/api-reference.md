# API 参考

## 概述

本文档提供 OceanEngine 广告管理平台 API 的快速参考。完整的 API 规范请参阅 [后端 API 设计文档](./backend/03-api-design.md)。

**基础信息：**
- 基础路径：`/api/v1`
- 认证方式：Bearer Token（JWT）
- 内容类型：`application/json`

---

## 认证接口

### 登录
```
POST /auth/login
```

**请求：**
```json
{
  "username": "admin",
  "password": "password123"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_in": 7200
  }
}
```

### 刷新 Token
```
POST /auth/refresh
```

### 获取用户信息
```
GET /auth/info
```

### 登出
```
POST /auth/logout
```

---

## 广告主接口

### 获取广告主列表
```
GET /advertiser
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 20 |
| keyword | string | 否 | 搜索关键词 |
| status | string | 否 | 状态筛选 |

### 获取广告主详情
```
GET /advertiser/:id
```

### 同步广告主数据
```
POST /advertiser/:id/sync
```

### 获取广告主余额
```
GET /advertiser/:id/balance
```

**响应：**
```json
{
  "code": 0,
  "data": {
    "balance": 10000.00,
    "valid_balance": 8000.00,
    "cash_balance": 5000.00
  }
}
```

### 获取资金流水
```
GET /advertiser/:id/funds
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start_date | string | 是 | 开始日期 |
| end_date | string | 是 | 结束日期 |
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |

---

## 广告系列接口

### 获取系列列表
```
GET /campaign
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| advertiser_id | int | 是 | 广告主 ID |
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| status | string | 否 | 状态筛选 |
| name | string | 否 | 名称搜索 |

### 创建系列
```
POST /campaign
```

**请求：**
```json
{
  "advertiser_id": 123456,
  "name": "测试系列",
  "budget_mode": "BUDGET_MODE_DAY",
  "budget": 1000.00,
  "landing_type": "LANDING_PAGE"
}
```

### 获取系列详情
```
GET /campaign/:id
```

### 更新系列
```
PUT /campaign/:id
```

### 更新系列状态
```
POST /campaign/:id/status
```

**请求：**
```json
{
  "status": "enable"  // enable | disable | delete
}
```

### 批量更新状态
```
POST /campaign/batch/status
```

**请求：**
```json
{
  "ids": [1, 2, 3],
  "status": "disable"
}
```

---

## 广告组接口

### 获取广告组列表
```
GET /ad
```

### 创建广告组
```
POST /ad
```

### 获取广告组详情
```
GET /ad/:id
```

### 更新广告组
```
PUT /ad/:id
```

### 更新广告组状态
```
POST /ad/:id/status
```

---

## 创意接口

### 获取创意列表
```
GET /creative
```

### 创建创意
```
POST /creative
```

### 获取创意详情
```
GET /creative/:id
```

### 更新创意
```
PUT /creative/:id
```

### 更新创意状态
```
POST /creative/:id/status
```

---

## 报表接口

### 广告主报表
```
GET /report/advertiser
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| advertiser_id | int | 是 | 广告主 ID |
| start_date | string | 是 | 开始日期（YYYY-MM-DD） |
| end_date | string | 是 | 结束日期 |
| group_by | string[] | 否 | 分组维度 |

**响应：**
```json
{
  "code": 0,
  "data": [
    {
      "stat_datetime": "2024-01-15",
      "cost": 1000.00,
      "show": 50000,
      "click": 1500,
      "ctr": 3.00,
      "convert": 50,
      "convert_cost": 20.00,
      "convert_rate": 3.33
    }
  ]
}
```

### 广告系列报表
```
GET /report/campaign
```

### 实时数据
```
GET /report/realtime
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| advertiser_id | int | 是 | 广告主 ID |

### 导出报表
```
POST /report/export
```

**请求：**
```json
{
  "advertiser_id": 123456,
  "start_date": "2024-01-01",
  "end_date": "2024-01-31",
  "type": "campaign"
}
```

**响应：**
```json
{
  "code": 0,
  "data": {
    "task_id": "export_123456"
  }
}
```

### 获取导出结果
```
GET /report/export/:task_id
```

---

## 人群管理接口

### 获取人群列表
```
GET /audience
```

### 创建人群
```
POST /audience
```

### 获取人群详情
```
GET /audience/:id
```

### 更新人群
```
PUT /audience/:id
```

### 删除人群
```
DELETE /audience/:id
```

### 推送人群
```
POST /audience/:id/push
```

### 获取人群预估
```
GET /audience/:id/estimate
```

---

## 素材接口

### 上传图片
```
POST /material/image/upload
Content-Type: multipart/form-data
```

**参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | 图片文件 |
| advertiser_id | int | 是 | 广告主 ID |

### 上传视频
```
POST /material/video/upload
Content-Type: multipart/form-data
```

### 获取素材列表
```
GET /material
```

### 删除素材
```
DELETE /material/:id
```

---

## 系统管理接口

### 用户管理

```
GET    /system/user          # 用户列表
POST   /system/user          # 创建用户
GET    /system/user/:id      # 用户详情
PUT    /system/user/:id      # 更新用户
DELETE /system/user/:id      # 删除用户
```

### 角色管理

```
GET    /system/role          # 角色列表
POST   /system/role          # 创建角色
GET    /system/role/:id      # 角色详情
PUT    /system/role/:id      # 更新角色
DELETE /system/role/:id      # 删除角色
```

### 菜单管理

```
GET    /system/menu          # 菜单列表
POST   /system/menu          # 创建菜单
PUT    /system/menu/:id      # 更新菜单
DELETE /system/menu/:id      # 删除菜单
```

---

## 响应格式

### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": { },
  "request_id": "req_123456"
}
```

### 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

### 错误响应

```json
{
  "code": 10001,
  "message": "参数错误",
  "request_id": "req_123456"
}
```

---

## 错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 10001 | 参数错误 |
| 10002 | 数据不存在 |
| 10003 | 数据已存在 |
| 20001 | 未登录 |
| 20002 | Token 过期 |
| 20003 | 权限不足 |
| 30001 | 系统错误 |
| 30002 | 第三方服务错误 |
| 40001 | 巨量引擎 API 错误 |
| 40002 | 巨量引擎限流 |

---

## 限流规则

| 接口类型 | 限流规则 |
|----------|----------|
| 登录接口 | 5 次/分钟/IP |
| 普通接口 | 100 次/分钟/用户 |
| 报表接口 | 20 次/分钟/用户 |
| 导出接口 | 5 次/分钟/用户 |
| 上传接口 | 10 次/分钟/用户 |

---

## 更多资源

- [完整 API 设计文档](./backend/03-api-design.md)
- [认证授权设计](./backend/04-auth-design.md)
- [错误处理方案](./backend/07-error-handling.md)
- [Swagger 文档](http://localhost:8080/swagger/index.html)（开发环境）
