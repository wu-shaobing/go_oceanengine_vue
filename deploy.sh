#!/bin/bash

# OceanEngine 广告管理平台 - 一键部署脚本
# 使用方法: ./deploy.sh

set -e

echo "========================================"
echo "  OceanEngine 广告管理平台部署脚本"
echo "========================================"
echo ""

# 检查是否在项目根目录
if [ ! -d "backend" ] || [ ! -d "frontend" ]; then
    echo "❌ 错误: 请在项目根目录运行此脚本"
    exit 1
fi

# 检查 Docker 是否安装
if ! command -v docker &> /dev/null; then
    echo "❌ 错误: 未安装 Docker，请先安装 Docker"
    exit 1
fi

# 检查 Docker Compose 是否安装
if ! command -v docker compose &> /dev/null; then
    echo "❌ 错误: 未安装 Docker Compose，请先安装 Docker Compose"
    exit 1
fi

echo "✅ 环境检查通过"
echo ""

# 进入后端目录
cd backend

# 检查 .env 文件
if [ ! -f ".env" ]; then
    echo "❌ 错误: .env 文件不存在"
    echo "请先创建 .env 文件: cp .env.example .env"
    echo "并填写必要的配置信息"
    exit 1
fi

echo "✅ 配置文件检查通过"
echo ""

echo "📦 正在启动 Docker 服务..."
echo ""

# 启动所有服务
docker compose up -d

echo ""
echo "⏳ 等待服务启动..."
sleep 10

# 检查服务状态
echo ""
echo "📊 服务状态:"
docker compose ps

echo ""
echo "⏳ 等待数据库就绪..."
sleep 20

echo ""
echo "🔄 运行数据库迁移..."

# 运行数据库迁移
docker compose exec -T app sh -c "cd /app && ./main -action=migrate" || true

echo ""
echo "📝 填充初始数据..."

# 填充初始数据
docker compose exec -T app sh -c "cd /app && ./main -action=seed" || true

echo ""
echo "========================================"
echo "  ✅ 部署完成！"
echo "========================================"
echo ""
echo "📊 服务地址:"
echo "  - 后端 API:    http://localhost:8080"
echo "  - 健康检查:    http://localhost:8080/health"
echo "  - phpMyAdmin:  http://localhost:8081"
echo "  - Redis管理:   http://localhost:8082"
echo ""
echo "👤 默认管理员账号:"
echo "  - 用户名: admin"
echo "  - 密码:   admin123"
echo ""
echo "🔍 查看日志:"
echo "  docker compose logs -f app"
echo ""
echo "🛑 停止服务:"
echo "  docker compose down"
echo ""
