# OceanEngine API Server

基于 OceanEngine SDK 构建的 HTTP API 服务，支持多种腾讯云部署方式。

## 快速开始

### 本地运行

```bash
# 1. 复制环境变量配置
cp .env.example .env

# 2. 编辑 .env 文件，填入你的 AppID 和 AppSecret
vim .env

# 3. 安装依赖
go mod tidy

# 4. 运行服务
go run ./cmd/api
```

服务默认运行在 `http://localhost:8080`

### API 端点

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |
| POST | `/api/v1/oauth/auth_url` | 获取授权URL |
| POST | `/api/v1/oauth/access_token` | 获取访问令牌 |
| POST | `/api/v1/oauth/refresh_token` | 刷新访问令牌 |
| GET | `/api/v1/advertiser/list` | 获取广告主列表 |
| POST | `/api/v1/advertiser/info` | 获取广告主详情 |

## 部署方式

### 1. Docker 部署

```bash
# 在 oceanengine 根目录执行
cd /path/to/oceanengine

# 构建镜像
docker build -f server/deployments/docker/Dockerfile -t oceanengine-api:latest .

# 运行容器
docker run -d \
  -p 8080:8080 \
  -e OCEANENGINE_APP_ID=your_app_id \
  -e OCEANENGINE_APP_SECRET=your_app_secret \
  oceanengine-api:latest
```

或使用 docker-compose:

```bash
cd server/deployments/docker

# 设置环境变量
export OCEANENGINE_APP_ID=your_app_id
export OCEANENGINE_APP_SECRET=your_app_secret

# 启动
docker-compose up -d
```

### 2. 腾讯云容器服务 TKE

```bash
# 1. 推送镜像到腾讯云容器镜像仓库
docker tag oceanengine-api:latest ccr.ccs.tencentyun.com/your-namespace/oceanengine-api:latest
docker push ccr.ccs.tencentyun.com/your-namespace/oceanengine-api:latest

# 2. 创建 Secret (修改 service.yaml 中的凭据)
kubectl apply -f deployments/k8s/service.yaml

# 3. 部署应用
kubectl apply -f deployments/k8s/deployment.yaml
```

### 3. 腾讯云 Serverless (SCF)

```bash
# 1. 安装 Serverless Framework
npm install -g serverless

# 2. 配置腾讯云凭据
serverless login

# 3. 构建 SCF 二进制
cd deployments/scf
GOOS=linux GOARCH=amd64 go build -o main main.go

# 4. 部署
export OCEANENGINE_APP_ID=your_app_id
export OCEANENGINE_APP_SECRET=your_app_secret
serverless deploy
```

### 4. 腾讯云 CVM / 轻量应用服务器

```bash
# 1. 编译二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-server ./cmd/api

# 2. 上传到服务器
scp api-server user@your-server:/opt/oceanengine/

# 3. 创建 systemd 服务 (在服务器上)
cat > /etc/systemd/system/oceanengine-api.service << EOF
[Unit]
Description=OceanEngine API Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/oceanengine
ExecStart=/opt/oceanengine/api-server
Restart=always
Environment=SERVER_MODE=release
Environment=OCEANENGINE_APP_ID=your_app_id
Environment=OCEANENGINE_APP_SECRET=your_app_secret

[Install]
WantedBy=multi-user.target
EOF

# 4. 启动服务
systemctl daemon-reload
systemctl enable oceanengine-api
systemctl start oceanengine-api
```

## API 使用示例

### 获取授权URL

```bash
curl -X POST http://localhost:8080/api/v1/oauth/auth_url \
  -H "Content-Type: application/json" \
  -d '{"redirect_uri": "https://your-domain.com/callback", "state": "random_state"}'
```

### 获取访问令牌

```bash
curl -X POST http://localhost:8080/api/v1/oauth/access_token \
  -H "Content-Type: application/json" \
  -d '{"auth_code": "AUTH_CODE_FROM_CALLBACK"}'
```

### 获取广告主列表

```bash
curl -X GET http://localhost:8080/api/v1/advertiser/list \
  -H "Access-Token: YOUR_ACCESS_TOKEN"
```

### 获取广告主详情

```bash
curl -X POST http://localhost:8080/api/v1/advertiser/info \
  -H "Content-Type: application/json" \
  -H "Access-Token: YOUR_ACCESS_TOKEN" \
  -d '{"advertiser_ids": [123456789]}'
```

## 项目结构

```
server/
├── cmd/api/                    # 主程序入口
│   └── main.go
├── internal/
│   ├── config/                 # 配置管理
│   │   └── config.go
│   ├── handler/                # HTTP 处理器
│   │   └── handler.go
│   ├── middleware/             # 中间件
│   │   └── middleware.go
│   └── service/                # 业务服务
│       └── oceanengine.go
├── deployments/
│   ├── docker/                 # Docker 配置
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   ├── k8s/                    # Kubernetes 配置
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   └── scf/                    # 腾讯云函数配置
│       ├── main.go
│       └── serverless.yml
├── .env.example                # 环境变量示例
├── go.mod
└── README.md
```

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `SERVER_PORT` | 服务端口 | 8080 |
| `SERVER_MODE` | Gin运行模式 (debug/release/test) | debug |
| `OCEANENGINE_APP_ID` | 巨量引擎 App ID | - |
| `OCEANENGINE_APP_SECRET` | 巨量引擎 App Secret | - |

## License

MIT
