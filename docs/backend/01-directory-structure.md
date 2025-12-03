# 项目目录结构设计

## 整体目录结构

```
backend/
├── cmd/                        # 应用入口
│   ├── server/                 # API 服务入口
│   │   └── main.go
│   ├── migrate/                # 数据库迁移工具
│   │   └── main.go
│   └── task/                   # 定时任务入口
│       └── main.go
│
├── config/                     # 配置文件
│   ├── config.go              # 配置结构定义
│   ├── settings.yml           # 默认配置
│   ├── settings.dev.yml       # 开发环境配置
│   ├── settings.test.yml      # 测试环境配置
│   └── settings.prod.yml      # 生产环境配置
│
├── internal/                   # 内部代码（不对外暴露）
│   ├── app/                   # 应用模块
│   │   ├── admin/             # 系统管理模块
│   │   │   ├── api/           # API 控制器
│   │   │   ├── service/       # 业务逻辑
│   │   │   ├── repository/    # 数据访问
│   │   │   ├── model/         # 数据模型
│   │   │   ├── dto/           # 数据传输对象
│   │   │   └── router/        # 路由注册
│   │   │
│   │   ├── advertiser/        # 广告主模块
│   │   │   ├── api/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   ├── dto/
│   │   │   └── router/
│   │   │
│   │   ├── campaign/          # 广告系列模块
│   │   │   ├── api/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   ├── dto/
│   │   │   └── router/
│   │   │
│   │   ├── creative/          # 创意模块
│   │   │   ├── api/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   ├── dto/
│   │   │   └── router/
│   │   │
│   │   ├── audience/          # 人群定向模块
│   │   │   ├── api/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   ├── dto/
│   │   │   └── router/
│   │   │
│   │   ├── report/            # 数据报告模块
│   │   │   ├── api/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── model/
│   │   │   ├── dto/
│   │   │   └── router/
│   │   │
│   │   └── media/             # 素材管理模块
│   │       ├── api/
│   │       ├── service/
│   │       ├── repository/
│   │       ├── model/
│   │       ├── dto/
│   │       └── router/
│   │
│   ├── middleware/            # 中间件
│   │   ├── auth.go            # JWT 认证
│   │   ├── cors.go            # 跨域处理
│   │   ├── logger.go          # 请求日志
│   │   ├── recovery.go        # 异常恢复
│   │   ├── ratelimit.go       # 限流
│   │   ├── permission.go      # 权限检查
│   │   └── trace.go           # 链路追踪
│   │
│   └── router/                # 路由汇总
│       └── router.go
│
├── pkg/                       # 可复用的公共包
│   ├── auth/                  # 认证相关
│   │   ├── jwt.go
│   │   └── casbin.go
│   │
│   ├── database/              # 数据库
│   │   ├── mysql.go
│   │   └── redis.go
│   │
│   ├── response/              # 统一响应
│   │   └── response.go
│   │
│   ├── errors/                # 错误定义
│   │   ├── errors.go
│   │   └── codes.go
│   │
│   ├── utils/                 # 工具函数
│   │   ├── crypto.go          # 加密
│   │   ├── validator.go       # 校验
│   │   ├── pagination.go      # 分页
│   │   └── time.go            # 时间处理
│   │
│   ├── logger/                # 日志
│   │   └── logger.go
│   │
│   └── oceanengine/           # SDK 封装
│       ├── client.go          # 客户端
│       ├── token.go           # Token 管理
│       ├── advertiser.go      # 广告主 API
│       ├── campaign.go        # 广告系列 API
│       ├── creative.go        # 创意 API
│       ├── report.go          # 报告 API
│       └── types.go           # 类型定义
│
├── api/                       # API 文档
│   └── swagger/
│       └── swagger.json
│
├── scripts/                   # 脚本
│   ├── build.sh               # 构建脚本
│   ├── deploy.sh              # 部署脚本
│   └── migrate.sh             # 迁移脚本
│
├── deployments/               # 部署配置
│   ├── docker/
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   └── kubernetes/
│       ├── deployment.yaml
│       └── service.yaml
│
├── test/                      # 测试
│   ├── integration/           # 集成测试
│   └── mock/                  # Mock 数据
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 各层职责说明

### 1. cmd/ - 应用入口层

每个子目录对应一个可执行程序：

```go
// cmd/server/main.go
package main

import (
    "oceanengine/config"
    "oceanengine/internal/router"
    "oceanengine/pkg/database"
    "oceanengine/pkg/logger"
)

func main() {
    // 1. 加载配置
    cfg := config.Load()
    
    // 2. 初始化日志
    logger.Init(cfg.Logger)
    
    // 3. 初始化数据库
    database.Init(cfg.Database)
    
    // 4. 初始化路由
    r := router.Init()
    
    // 5. 启动服务
    r.Run(cfg.Server.Addr)
}
```

### 2. config/ - 配置层

配置文件采用 YAML 格式，支持多环境：

```go
// config/config.go
package config

type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Redis    RedisConfig    `yaml:"redis"`
    JWT      JWTConfig      `yaml:"jwt"`
    Logger   LoggerConfig   `yaml:"logger"`
    Ocean    OceanConfig    `yaml:"ocean"` // Ocean Engine 配置
}

type ServerConfig struct {
    Mode         string `yaml:"mode"`         // dev, test, prod
    Host         string `yaml:"host"`
    Port         int    `yaml:"port"`
    ReadTimeout  int    `yaml:"readTimeout"`
    WriteTimeout int    `yaml:"writeTimeout"`
}

type OceanConfig struct {
    AppID       string `yaml:"appId"`
    Secret      string `yaml:"secret"`
    RedirectURI string `yaml:"redirectUri"`
}
```

### 3. internal/app/ - 业务模块层

每个模块遵循相同的目录结构：

#### api/ - 控制器

```go
// internal/app/advertiser/api/advertiser.go
package api

type AdvertiserAPI struct {
    service *service.AdvertiserService
}

// GetList 获取广告主列表
// @Summary 获取广告主列表
// @Tags 广告主
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.Response{data=dto.AdvertiserListResp}
// @Router /api/v1/advertisers [get]
func (a *AdvertiserAPI) GetList(c *gin.Context) {
    var req dto.AdvertiserListReq
    if err := c.ShouldBindQuery(&req); err != nil {
        response.Fail(c, errors.ParamError)
        return
    }
    
    data, total, err := a.service.GetList(c, &req)
    if err != nil {
        response.Fail(c, err)
        return
    }
    
    response.PageOK(c, data, total, req.Page, req.PageSize)
}
```

#### service/ - 业务逻辑

```go
// internal/app/advertiser/service/advertiser.go
package service

type AdvertiserService struct {
    repo       repository.AdvertiserRepository
    oceanSDK   *oceanengine.Client
}

func (s *AdvertiserService) GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*model.Advertiser, int64, error) {
    // 1. 从数据库获取本地数据
    list, total, err := s.repo.GetList(ctx, req)
    if err != nil {
        return nil, 0, err
    }
    
    // 2. 如需同步，从 Ocean Engine API 获取最新数据
    if req.Sync {
        for _, adv := range list {
            info, err := s.oceanSDK.GetAdvertiserInfo(ctx, adv.AdvertiserID)
            if err == nil {
                adv.Balance = info.Balance
                adv.Status = info.Status
            }
        }
    }
    
    return list, total, nil
}
```

#### repository/ - 数据访问

```go
// internal/app/advertiser/repository/advertiser.go
package repository

type AdvertiserRepository interface {
    GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*model.Advertiser, int64, error)
    GetByID(ctx context.Context, id int64) (*model.Advertiser, error)
    Create(ctx context.Context, advertiser *model.Advertiser) error
    Update(ctx context.Context, advertiser *model.Advertiser) error
    Delete(ctx context.Context, id int64) error
}

type advertiserRepository struct {
    db *gorm.DB
}

func (r *advertiserRepository) GetList(ctx context.Context, req *dto.AdvertiserListReq) ([]*model.Advertiser, int64, error) {
    var list []*model.Advertiser
    var total int64
    
    query := r.db.WithContext(ctx).Model(&model.Advertiser{})
    
    if req.Keyword != "" {
        query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
    }
    if req.Status != "" {
        query = query.Where("status = ?", req.Status)
    }
    
    err := query.Count(&total).Error
    if err != nil {
        return nil, 0, err
    }
    
    err = query.Offset(req.GetOffset()).Limit(req.GetLimit()).Find(&list).Error
    return list, total, err
}
```

#### model/ - 数据模型

```go
// internal/app/advertiser/model/advertiser.go
package model

type Advertiser struct {
    ID            int64     `gorm:"primaryKey" json:"id"`
    AdvertiserID  int64     `gorm:"uniqueIndex" json:"advertiserId"`   // Ocean Engine 广告主ID
    Name          string    `gorm:"size:255" json:"name"`
    Company       string    `gorm:"size:255" json:"company"`
    Status        string    `gorm:"size:50" json:"status"`
    Balance       float64   `json:"balance"`
    Industry      string    `gorm:"size:100" json:"industry"`
    AccessToken   string    `gorm:"size:500" json:"-"`                 // 不对外暴露
    RefreshToken  string    `gorm:"size:500" json:"-"`
    TokenExpireAt time.Time `json:"-"`
    CreatedAt     time.Time `json:"createdAt"`
    UpdatedAt     time.Time `json:"updatedAt"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Advertiser) TableName() string {
    return "sys_advertiser"
}
```

#### dto/ - 数据传输对象

```go
// internal/app/advertiser/dto/advertiser.go
package dto

// AdvertiserListReq 列表请求
type AdvertiserListReq struct {
    Page     int    `form:"page" binding:"min=1"`
    PageSize int    `form:"pageSize" binding:"min=1,max=100"`
    Keyword  string `form:"keyword"`
    Status   string `form:"status"`
    Sync     bool   `form:"sync"` // 是否同步远程数据
}

func (r *AdvertiserListReq) GetOffset() int {
    return (r.Page - 1) * r.PageSize
}

func (r *AdvertiserListReq) GetLimit() int {
    if r.PageSize == 0 {
        return 10
    }
    return r.PageSize
}

// AdvertiserListResp 列表响应
type AdvertiserListResp struct {
    List  []*AdvertiserItem `json:"list"`
    Total int64             `json:"total"`
}

// AdvertiserItem 列表项
type AdvertiserItem struct {
    ID           int64   `json:"id"`
    AdvertiserID int64   `json:"advertiserId"`
    Name         string  `json:"name"`
    Company      string  `json:"company"`
    Status       string  `json:"status"`
    Balance      float64 `json:"balance"`
    CreatedAt    string  `json:"createdAt"`
}

// AdvertiserCreateReq 创建请求
type AdvertiserCreateReq struct {
    AdvertiserID int64  `json:"advertiserId" binding:"required"`
    Name         string `json:"name" binding:"required,max=255"`
    Company      string `json:"company" binding:"max=255"`
    Industry     string `json:"industry" binding:"max=100"`
}

// AdvertiserUpdateReq 更新请求
type AdvertiserUpdateReq struct {
    ID       int64  `json:"id" binding:"required"`
    Name     string `json:"name" binding:"max=255"`
    Company  string `json:"company" binding:"max=255"`
    Industry string `json:"industry" binding:"max=100"`
}
```

#### router/ - 路由注册

```go
// internal/app/advertiser/router/router.go
package router

func Register(r *gin.RouterGroup, api *api.AdvertiserAPI) {
    group := r.Group("/advertisers")
    {
        group.GET("", api.GetList)           // 列表
        group.GET("/:id", api.GetDetail)     // 详情
        group.POST("", api.Create)           // 创建
        group.PUT("/:id", api.Update)        // 更新
        group.DELETE("/:id", api.Delete)     // 删除
        group.POST("/:id/sync", api.Sync)    // 同步远程数据
        group.GET("/:id/balance", api.GetBalance)  // 获取余额
    }
}
```

### 4. pkg/ - 公共包层

可被其他项目复用的公共代码：

```go
// pkg/response/response.go
package response

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type PageResponse struct {
    Response
    Total    int64 `json:"total"`
    Page     int   `json:"page"`
    PageSize int   `json:"pageSize"`
}

func OK(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

func Fail(c *gin.Context, err error) {
    var e *errors.AppError
    if errors.As(err, &e) {
        c.JSON(http.StatusOK, Response{
            Code:    e.Code,
            Message: e.Message,
        })
        return
    }
    c.JSON(http.StatusOK, Response{
        Code:    -1,
        Message: err.Error(),
    })
}

func PageOK(c *gin.Context, data interface{}, total int64, page, pageSize int) {
    c.JSON(http.StatusOK, PageResponse{
        Response: Response{
            Code:    0,
            Message: "success",
            Data:    data,
        },
        Total:    total,
        Page:     page,
        PageSize: pageSize,
    })
}
```

## 模块依赖关系

```
┌──────────────────────────────────────────────────────────┐
│                        cmd/server                         │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│                    internal/router                        │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│                    internal/middleware                    │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│                    internal/app/*/api                     │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│                  internal/app/*/service                   │
├────────────────────────┬─────────────────────────────────┤
│                        │                                  │
▼                        ▼                                  │
┌─────────────────┐  ┌─────────────────┐                   │
│ */repository    │  │ pkg/oceanengine │◄──────────────────┘
└────────┬────────┘  └────────┬────────┘
         │                    │
         ▼                    ▼
┌─────────────────┐  ┌─────────────────┐
│   pkg/database  │  │  Ocean Engine   │
│   (MySQL/Redis) │  │      API        │
└─────────────────┘  └─────────────────┘
```
