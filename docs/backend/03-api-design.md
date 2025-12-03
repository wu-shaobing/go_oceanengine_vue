# API 接口设计

## API 规范

### URL 规范
- 基础路径：`/api/v1`
- 资源命名：复数名词，小写，连字符分隔
- 层级关系：`/资源/id/子资源`

### HTTP 方法
| 方法 | 用途 | 示例 |
|------|------|------|
| GET | 查询资源 | GET /api/v1/advertisers |
| POST | 创建资源 | POST /api/v1/advertisers |
| PUT | 全量更新 | PUT /api/v1/advertisers/1 |
| PATCH | 部分更新 | PATCH /api/v1/advertisers/1 |
| DELETE | 删除资源 | DELETE /api/v1/advertisers/1 |

### 请求头
```
Content-Type: application/json
Authorization: Bearer <token>
X-Request-ID: <uuid>    // 请求追踪ID
```

### 统一响应格式

#### 成功响应
```json
{
    "code": 0,
    "message": "success",
    "data": { ... },
    "timestamp": 1703001234567
}
```

#### 分页响应
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "list": [ ... ],
        "total": 100,
        "page": 1,
        "pageSize": 10
    },
    "timestamp": 1703001234567
}
```

#### 错误响应
```json
{
    "code": 10001,
    "message": "参数错误",
    "details": [
        { "field": "name", "message": "不能为空" }
    ],
    "timestamp": 1703001234567
}
```

### 错误码定义
| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 10001 | 参数错误 |
| 10002 | 数据不存在 |
| 10003 | 数据已存在 |
| 20001 | 未登录 |
| 20002 | Token 过期 |
| 20003 | 权限不足 |
| 30001 | Ocean Engine API 错误 |
| 30002 | Token 刷新失败 |
| 50001 | 系统内部错误 |

---

## 接口详细设计

### 1. 认证模块

#### 1.1 用户登录
```
POST /api/v1/auth/login
```

**请求参数**
```json
{
    "username": "admin",
    "password": "123456",
    "captchaId": "xxx",
    "captchaCode": "abc123"
}
```

**响应**
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIs...",
        "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
        "expiresIn": 7200,
        "user": {
            "id": 1,
            "username": "admin",
            "nickname": "管理员",
            "avatar": "https://...",
            "roles": ["admin"]
        }
    }
}
```

#### 1.2 刷新 Token
```
POST /api/v1/auth/refresh
```

**请求参数**
```json
{
    "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
}
```

#### 1.3 退出登录
```
POST /api/v1/auth/logout
```

#### 1.4 获取当前用户信息
```
GET /api/v1/auth/userinfo
```

#### 1.5 获取验证码
```
GET /api/v1/auth/captcha
```

**响应**
```json
{
    "code": 0,
    "data": {
        "captchaId": "xxx",
        "captchaImage": "data:image/png;base64,..."
    }
}
```

---

### 2. 广告主模块

#### 2.1 获取广告主列表
```
GET /api/v1/advertisers
```

**查询参数**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认1 |
| pageSize | int | 否 | 每页数量，默认10，最大100 |
| keyword | string | 否 | 搜索关键词 |
| status | string | 否 | 状态筛选 |
| sync | bool | 否 | 是否同步远程数据 |

**响应**
```json
{
    "code": 0,
    "data": {
        "list": [
            {
                "id": 1,
                "advertiserId": 123456789,
                "name": "测试广告主",
                "company": "测试公司",
                "status": "STATUS_ENABLE",
                "balance": 10000.00,
                "createdAt": "2024-01-01 12:00:00"
            }
        ],
        "total": 100,
        "page": 1,
        "pageSize": 10
    }
}
```

#### 2.2 获取广告主详情
```
GET /api/v1/advertisers/:id
```

**响应**
```json
{
    "code": 0,
    "data": {
        "id": 1,
        "advertiserId": 123456789,
        "name": "测试广告主",
        "company": "测试公司",
        "status": "STATUS_ENABLE",
        "role": "ROLE_ADVERTISER",
        "balance": 10000.00,
        "validBalance": 9000.00,
        "cashBalance": 8000.00,
        "industry": "电商",
        "contactName": "张三",
        "contactPhone": "13800138000",
        "contactEmail": "test@example.com",
        "address": "北京市朝阳区",
        "lastSyncAt": "2024-01-01 12:00:00",
        "createdAt": "2024-01-01 12:00:00",
        "updatedAt": "2024-01-01 12:00:00"
    }
}
```

#### 2.3 添加广告主（OAuth 授权）
```
POST /api/v1/advertisers/oauth/authorize
```

**说明**：跳转到 Ocean Engine OAuth 授权页面

**响应**
```json
{
    "code": 0,
    "data": {
        "authorizeUrl": "https://ad.oceanengine.com/openapi/audit/oauth.html?..."
    }
}
```

#### 2.4 OAuth 回调
```
GET /api/v1/advertisers/oauth/callback
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| auth_code | string | 授权码 |
| state | string | 状态参数 |

#### 2.5 同步广告主数据
```
POST /api/v1/advertisers/:id/sync
```

**响应**
```json
{
    "code": 0,
    "message": "同步成功",
    "data": {
        "advertiserId": 123456789,
        "syncFields": ["balance", "status", "name"],
        "syncAt": "2024-01-01 12:00:00"
    }
}
```

#### 2.6 获取广告主余额
```
GET /api/v1/advertisers/:id/balance
```

**响应**
```json
{
    "code": 0,
    "data": {
        "balance": 10000.00,
        "validBalance": 9000.00,
        "cashBalance": 8000.00,
        "grantBalance": 1000.00,
        "syncAt": "2024-01-01 12:00:00"
    }
}
```

#### 2.7 获取资金流水
```
GET /api/v1/advertisers/:id/funds
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| page | int | 页码 |
| pageSize | int | 每页数量 |
| startDate | string | 开始日期 |
| endDate | string | 结束日期 |
| transactionType | string | 交易类型 |

---

### 3. 广告系列模块

#### 3.1 获取广告系列列表
```
GET /api/v1/campaigns
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID（必填） |
| page | int | 页码 |
| pageSize | int | 每页数量 |
| keyword | string | 搜索关键词 |
| status | string | 状态筛选 |
| landingType | string | 推广类型 |

**响应**
```json
{
    "code": 0,
    "data": {
        "list": [
            {
                "id": 1,
                "campaignId": 1234567890,
                "advertiserId": 123456789,
                "name": "双十一大促",
                "budgetMode": "BUDGET_MODE_DAY",
                "budget": 10000.00,
                "landingType": "LINK",
                "status": "CAMPAIGN_STATUS_ENABLE",
                "optStatus": "ENABLE",
                "createdAt": "2024-01-01 12:00:00"
            }
        ],
        "total": 50,
        "page": 1,
        "pageSize": 10
    }
}
```

#### 3.2 获取广告系列详情
```
GET /api/v1/campaigns/:id
```

#### 3.3 创建广告系列
```
POST /api/v1/campaigns
```

**请求参数**
```json
{
    "advertiserId": 123456789,
    "name": "双十一大促",
    "budgetMode": "BUDGET_MODE_DAY",
    "budget": 10000.00,
    "landingType": "LINK",
    "marketingGoal": "LIVE_PROM_GOODS"
}
```

**响应**
```json
{
    "code": 0,
    "data": {
        "id": 1,
        "campaignId": 1234567890
    }
}
```

#### 3.4 更新广告系列
```
PUT /api/v1/campaigns/:id
```

**请求参数**
```json
{
    "name": "双十一大促-更新",
    "budget": 20000.00
}
```

#### 3.5 更新广告系列状态
```
PATCH /api/v1/campaigns/:id/status
```

**请求参数**
```json
{
    "optStatus": "DISABLE"
}
```

#### 3.6 删除广告系列
```
DELETE /api/v1/campaigns/:id
```

#### 3.7 批量更新状态
```
PATCH /api/v1/campaigns/batch/status
```

**请求参数**
```json
{
    "ids": [1, 2, 3],
    "optStatus": "DISABLE"
}
```

---

### 4. 广告组模块

#### 4.1 获取广告组列表
```
GET /api/v1/ads
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID（必填） |
| campaignId | int | 广告系列ID |
| page | int | 页码 |
| pageSize | int | 每页数量 |
| keyword | string | 搜索关键词 |
| status | string | 状态筛选 |

#### 4.2 获取广告组详情
```
GET /api/v1/ads/:id
```

#### 4.3 创建广告组
```
POST /api/v1/ads
```

**请求参数**
```json
{
    "advertiserId": 123456789,
    "campaignId": 1234567890,
    "name": "广告组1",
    "deliveryRange": "DEFAULT",
    "budgetMode": "BUDGET_MODE_DAY",
    "budget": 1000.00,
    "bidType": "BIDTYPE_CPC",
    "bid": 1.50,
    "pricing": "PRICING_CPC",
    "startTime": "2024-01-01 00:00:00",
    "endTime": "2024-12-31 23:59:59",
    "audience": {
        "district": "CITY",
        "city": [110000, 310000],
        "gender": "GENDER_MALE",
        "age": ["AGE_18_23", "AGE_24_30"]
    }
}
```

#### 4.4 更新广告组
```
PUT /api/v1/ads/:id
```

#### 4.5 更新广告组状态
```
PATCH /api/v1/ads/:id/status
```

#### 4.6 更新广告组预算
```
PATCH /api/v1/ads/:id/budget
```

**请求参数**
```json
{
    "budget": 2000.00
}
```

#### 4.7 更新广告组出价
```
PATCH /api/v1/ads/:id/bid
```

**请求参数**
```json
{
    "bid": 2.00,
    "cpaBid": 50.00
}
```

---

### 5. 创意模块

#### 5.1 获取创意列表
```
GET /api/v1/creatives
```

#### 5.2 获取创意详情
```
GET /api/v1/creatives/:id
```

#### 5.3 创建创意
```
POST /api/v1/creatives
```

**请求参数**
```json
{
    "advertiserId": 123456789,
    "adId": 1234567890,
    "creativeMaterialMode": "STATIC_ASSEMBLE",
    "title": "限时特惠，点击了解更多",
    "source": "品牌官方",
    "imageMode": "CREATIVE_IMAGE_MODE_LARGE",
    "imageIds": ["image_id_1", "image_id_2"],
    "actionText": "立即购买",
    "actionUrl": "https://example.com/landing"
}
```

#### 5.4 更新创意状态
```
PATCH /api/v1/creatives/:id/status
```

---

### 6. 素材模块

#### 6.1 获取图片素材列表
```
GET /api/v1/materials/images
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| page | int | 页码 |
| pageSize | int | 每页数量 |
| keyword | string | 文件名搜索 |

#### 6.2 上传图片素材
```
POST /api/v1/materials/images/upload
```

**请求类型**：multipart/form-data

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| file | file | 图片文件 |
| filename | string | 文件名（可选） |

**响应**
```json
{
    "code": 0,
    "data": {
        "imageId": "xxx",
        "url": "https://...",
        "width": 1200,
        "height": 628,
        "size": 102400,
        "format": "jpg"
    }
}
```

#### 6.3 获取视频素材列表
```
GET /api/v1/materials/videos
```

#### 6.4 上传视频素材
```
POST /api/v1/materials/videos/upload
```

#### 6.5 获取素材详情
```
GET /api/v1/materials/:type/:id
```

---

### 7. 人群定向模块

#### 7.1 获取定向包列表
```
GET /api/v1/audience/packages
```

#### 7.2 创建定向包
```
POST /api/v1/audience/packages
```

**请求参数**
```json
{
    "advertiserId": 123456789,
    "name": "高价值用户定向",
    "description": "18-35岁，一线城市",
    "landingType": "LINK",
    "deliveryRange": "DEFAULT",
    "audience": {
        "district": "CITY",
        "city": [110000, 310000, 440100, 440300],
        "gender": "GENDER_FEMALE",
        "age": ["AGE_18_23", "AGE_24_30"],
        "interestCategories": [1001, 1002],
        "behaviorCategories": [2001, 2002]
    }
}
```

#### 7.3 获取自定义人群列表
```
GET /api/v1/audience/custom
```

#### 7.4 上传自定义人群
```
POST /api/v1/audience/custom/upload
```

**请求类型**：multipart/form-data

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| name | string | 人群包名称 |
| file | file | 数据文件（CSV格式） |
| dataType | string | 数据类型：IMEI/OAID/PHONE |

---

### 8. 数据报告模块

#### 8.1 获取广告主报告
```
GET /api/v1/reports/advertiser
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| startDate | string | 开始日期 YYYY-MM-DD |
| endDate | string | 结束日期 YYYY-MM-DD |
| timeGranularity | string | 时间粒度：STAT_TIME_GRANULARITY_DAILY/HOURLY |
| groupBy | array | 分组维度 |

**响应**
```json
{
    "code": 0,
    "data": {
        "list": [
            {
                "statDate": "2024-01-01",
                "cost": 10000.00,
                "showCnt": 500000,
                "clickCnt": 15000,
                "ctr": 3.00,
                "avgClickCost": 0.67,
                "convertCnt": 300,
                "convertCost": 33.33,
                "convertRate": 2.00
            }
        ],
        "summary": {
            "totalCost": 70000.00,
            "totalShow": 3500000,
            "totalClick": 105000,
            "avgCtr": 3.00,
            "totalConvert": 2100,
            "avgConvertCost": 33.33
        }
    }
}
```

#### 8.2 获取广告系列报告
```
GET /api/v1/reports/campaign
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| campaignIds | array | 广告系列ID列表 |
| startDate | string | 开始日期 |
| endDate | string | 结束日期 |
| timeGranularity | string | 时间粒度 |

#### 8.3 获取广告组报告
```
GET /api/v1/reports/ad
```

#### 8.4 获取创意报告
```
GET /api/v1/reports/creative
```

#### 8.5 导出报告
```
POST /api/v1/reports/export
```

**请求参数**
```json
{
    "advertiserId": 123456789,
    "reportType": "advertiser",
    "startDate": "2024-01-01",
    "endDate": "2024-01-31",
    "format": "xlsx",
    "fields": ["stat_date", "cost", "show_cnt", "click_cnt", "ctr"]
}
```

**响应**
```json
{
    "code": 0,
    "data": {
        "taskId": "export_task_123",
        "status": "processing"
    }
}
```

#### 8.6 获取导出任务状态
```
GET /api/v1/reports/export/:taskId
```

**响应**
```json
{
    "code": 0,
    "data": {
        "taskId": "export_task_123",
        "status": "completed",
        "downloadUrl": "https://...",
        "expireAt": "2024-01-02 12:00:00"
    }
}
```

#### 8.7 获取实时数据
```
GET /api/v1/reports/realtime
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| advertiserId | int | 广告主ID |
| level | string | 层级：advertiser/campaign/ad |
| ids | array | ID列表（可选） |

---

### 9. 系统管理模块

#### 9.1 用户管理

##### 获取用户列表
```
GET /api/v1/system/users
```

##### 创建用户
```
POST /api/v1/system/users
```

##### 更新用户
```
PUT /api/v1/system/users/:id
```

##### 删除用户
```
DELETE /api/v1/system/users/:id
```

##### 重置密码
```
POST /api/v1/system/users/:id/reset-password
```

#### 9.2 角色管理

##### 获取角色列表
```
GET /api/v1/system/roles
```

##### 创建角色
```
POST /api/v1/system/roles
```

##### 更新角色
```
PUT /api/v1/system/roles/:id
```

##### 配置角色权限
```
PUT /api/v1/system/roles/:id/menus
```

**请求参数**
```json
{
    "menuIds": [1, 2, 3, 4, 5]
}
```

#### 9.3 菜单管理

##### 获取菜单树
```
GET /api/v1/system/menus/tree
```

##### 创建菜单
```
POST /api/v1/system/menus
```

##### 更新菜单
```
PUT /api/v1/system/menus/:id
```

#### 9.4 字典管理

##### 获取字典类型列表
```
GET /api/v1/system/dict/types
```

##### 获取字典数据
```
GET /api/v1/system/dict/data/:type
```

#### 9.5 操作日志

##### 获取操作日志列表
```
GET /api/v1/system/logs/operation
```

**查询参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| page | int | 页码 |
| pageSize | int | 每页数量 |
| userId | int | 用户ID |
| module | string | 模块 |
| startTime | string | 开始时间 |
| endTime | string | 结束时间 |

---

## API 版本管理

### 版本策略
- URL 路径版本：`/api/v1/`, `/api/v2/`
- 不兼容的修改发布新版本
- 旧版本保持向后兼容至少 6 个月

### 版本升级通知
- API 文档标注废弃接口
- 响应头返回 `X-API-Deprecated: true`
- 提供升级指南文档
