# OceanEngine 广告管理平台 - 生产环境部署文档

**部署服务器**: 1.12.234.253  
**部署时间**: 2025-12-03  
**环境**: Docker Compose  

---

## 📋 已配置的应用信息

### 1. 巨量广告（代理商）
- **状态**: ✅ 已上线
- **APP_ID**: `1850175799663708`
- **Secret**: `45d3a92c5384089f0204b14c1375c28da790a02d`
- **回调地址**: `http://1.12.234.253/auth/callback`
- **应用类型**: 巨量广告-自研投放系统-代理商
- **授权URL**: https://open.oceanengine.com/audit/oauth.html?app_id=1850175799663708&state=your_custom_params&material_auth=1&rid=m6f4y7tt4pi

**应用能力**:
- ✅ 可服务账户生效主体数：10个不同主体（代理主体）
- ✅ 可授权用户数：50个
- ✅ 可申请接口：除工作台管理类接口

---

### 2. 巨量千川（代理商）
- **状态**: ✅ 已上线
- **APP_ID**: `1850228280031387`
- **Secret**: `a30dd267362779428e97330f49d73216208233a5`
- **回调地址**: `http://1.12.234.253/auth/callback`
- **应用类型**: 巨量千川-自研投放系统-代理商
- **接入能力范围**: 千川PC版
- **授权URL**: https://qianchuan.jinritemai.com/openapi/qc/audit/oauth.html?app_id=1850228280031387&state=your_custom_params&material_auth=1&rid=sfkr52gq2jb

**应用能力**:
- ✅ 可服务账户生效主体数：10个不同主体（代理主体）
- ✅ 可授权用户数：50个
- ✅ 可申请接口：除工作台管理类接口

---

## 🔐 安全配置

### 生成的密钥（已配置在 .env）

```bash
# JWT Secret (32字节)
JWT_SECRET_KEY=fNP+uIAuqmVYJmL2qZ7su/jILKWejqf1E2bSjqx9PEw=

# MySQL Root 密码
MYSQL_ROOT_PASSWORD=UQZLy6JwOmzPizL0brcZby+RsGVkleL+

# MySQL 应用密码
MYSQL_PASSWORD=UQZLy6JwOmzPizL0brcZby+RsGVkleL+

# Redis 密码
REDIS_PASSWORD=5LAFAtqmkJaqXOqsjHHTSC1WLo4IA55E
```

⚠️ **重要**: 这些密钥已经配置在 `backend/.env` 文件中，请妥善保管！

---

## 🚀 部署步骤

### 前置条件

确保服务器已安装：
- Docker 20.10+
- Docker Compose 2.0+

### 1. 上传项目到服务器

```bash
# 在本地打包项目（排除不必要的文件）
cd /Users/wushaobing911/Desktop/oceanengine
tar --exclude='node_modules' \
    --exclude='.git' \
    --exclude='frontend/dist' \
    --exclude='backend/logs' \
    -czf oceanengine-prod.tar.gz .

# 上传到服务器
scp oceanengine-prod.tar.gz root@1.12.234.253:/opt/

# 在服务器上解压
ssh root@1.12.234.253
cd /opt
tar -xzf oceanengine-prod.tar.gz
mv oceanengine-prod oceanengine
cd oceanengine
```

### 2. 启动后端服务

```bash
cd backend

# 环境变量文件已配置好（.env）
# 验证配置
cat .env

# 启动所有服务（MySQL + Redis + 后端）
docker compose up -d

# 查看日志
docker compose logs -f app

# 等待服务启动完成（约30秒）
docker compose ps
```

### 3. 运行数据库迁移

```bash
# 进入应用容器
docker compose exec app sh

# 运行迁移（创建表结构）
./main -action=migrate

# 填充初始数据（创建管理员账号）
./main -action=seed

# 退出容器
exit
```

### 4. 构建并部署前端

```bash
cd ../frontend

# 安装依赖
npm install

# 构建生产版本
npm run build

# 部署到 Nginx（假设使用 Nginx）
sudo mkdir -p /var/www/oceanengine
sudo cp -r dist/* /var/www/oceanengine/
```

### 5. 配置 Nginx（可选）

创建 Nginx 配置文件 `/etc/nginx/sites-available/oceanengine`:

```nginx
server {
    listen 80;
    server_name 1.12.234.253;

    # 前端静态文件
    location / {
        root /var/www/oceanengine;
        index index.html;
        try_files $uri $uri/ /index.html;
    }

    # 后端 API 代理
    location /api/ {
        proxy_pass http://localhost:8080/api/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # OAuth 回调
    location /auth/callback {
        proxy_pass http://localhost:8080/api/v1/advertisers/oauth/callback;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # 健康检查
    location /health {
        proxy_pass http://localhost:8080/health;
    }
}
```

启用配置：

```bash
sudo ln -s /etc/nginx/sites-available/oceanengine /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

---

## 🔍 验证部署

### 1. 检查服务状态

```bash
# 检查容器状态
docker compose ps

# 应该看到：
# NAME                    STATUS
# oceanengine-backend     Up (healthy)
# oceanengine-mysql       Up (healthy)
# oceanengine-redis       Up (healthy)
```

### 2. 测试健康检查

```bash
curl http://localhost:8080/health
# 应该返回: {"status":"ok"}
```

### 3. 访问前端

浏览器打开: `http://1.12.234.253`

### 4. 测试登录

**默认管理员账号**:
- 用户名: `admin`
- 密码: `admin123`

### 5. 测试 OAuth 授权

#### 巨量广告授权
访问: https://open.oceanengine.com/audit/oauth.html?app_id=1850175799663708&state=test&material_auth=1&rid=m6f4y7tt4pi

#### 巨量千川授权
访问: https://qianchuan.jinritemai.com/openapi/qc/audit/oauth.html?app_id=1850228280031387&state=test&material_auth=1&rid=sfkr52gq2jb

---

## 📊 服务地址

| 服务 | 地址 | 说明 |
|------|------|------|
| 前端界面 | http://1.12.234.253 | 主应用界面 |
| 后端 API | http://1.12.234.253/api | RESTful API |
| 健康检查 | http://1.12.234.253/health | 服务健康状态 |
| phpMyAdmin | http://1.12.234.253:8081 | 数据库管理 |
| Redis Commander | http://1.12.234.253:8082 | Redis 管理 |

---

## 📝 日常运维

### 查看日志

```bash
# 所有服务日志
docker compose logs -f

# 仅查看应用日志
docker compose logs -f app

# 仅查看 MySQL 日志
docker compose logs -f mysql

# 最近 100 行
docker compose logs --tail=100 app
```

### 重启服务

```bash
# 重启所有服务
docker compose restart

# 仅重启应用
docker compose restart app
```

### 停止服务

```bash
# 停止所有服务
docker compose down

# 停止并删除数据卷（⚠️ 危险操作）
docker compose down -v
```

### 备份数据库

```bash
# 创建备份
docker compose exec mysql mysqldump -u root -p oceanengine > backup_$(date +%Y%m%d_%H%M%S).sql

# 恢复备份
docker compose exec -T mysql mysql -u root -p oceanengine < backup_20251203_100000.sql
```

---

## 🔒 安全建议

### 1. 启用 HTTPS

⚠️ **强烈建议生产环境使用 HTTPS**

使用 Let's Encrypt 免费证书：

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com
```

### 2. 更新回调地址

如果启用了 HTTPS，需要在巨量引擎平台更新回调地址：
- 从 `http://1.12.234.253/auth/callback`
- 改为 `https://yourdomain.com/auth/callback`

### 3. 配置防火墙

```bash
# 仅开放必要端口
sudo ufw allow 80/tcp    # HTTP
sudo ufw allow 443/tcp   # HTTPS
sudo ufw allow 22/tcp    # SSH
sudo ufw enable
```

### 4. 定期备份

设置定时任务每天备份数据库：

```bash
# 编辑 crontab
crontab -e

# 添加：每天凌晨 2 点备份
0 2 * * * cd /opt/oceanengine/backend && docker compose exec -T mysql mysqldump -u root -pUQZLy6JwOmzPizL0brcZby+RsGVkleL+ oceanengine > /backup/oceanengine_$(date +\%Y\%m\%d).sql
```

---

## 🐛 故障排查

### 问题 1: 容器无法启动

```bash
# 查看详细错误
docker compose logs app

# 检查端口占用
netstat -tlnp | grep :8080
```

### 问题 2: 数据库连接失败

```bash
# 检查 MySQL 健康状态
docker compose exec mysql mysqladmin ping -u root -p

# 验证环境变量
docker compose exec app env | grep DB_
```

### 问题 3: OAuth 回调失败

1. 确认回调地址配置正确：`http://1.12.234.253/auth/callback`
2. 检查 Nginx 代理配置
3. 查看应用日志：`docker compose logs -f app`

---

## 📞 支持联系

- **技术支持**: 11489573@qq.com
- **项目地址**: https://github.com/wu-shaobing/go_oceanengine_vue
- **问题反馈**: https://github.com/wu-shaobing/go_oceanengine_vue/issues

---

**部署完成时间**: 2025-12-03  
**配置状态**: ✅ 已完成  
**服务状态**: 待启动  
