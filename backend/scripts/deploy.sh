#!/bin/bash

# OceanEngine Backend 部署脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

# 默认配置
DEPLOY_ENV="${DEPLOY_ENV:-dev}"
DEPLOY_TYPE="${DEPLOY_TYPE:-docker}"
IMAGE_NAME="${IMAGE_NAME:-oceanengine-api}"
IMAGE_TAG="${IMAGE_TAG:-latest}"

usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -e, --env       Deploy environment (dev|test|prod), default: dev"
    echo "  -t, --type      Deploy type (docker|k8s), default: docker"
    echo "  -v, --version   Image version tag, default: latest"
    echo "  -h, --help      Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 -e dev -t docker"
    echo "  $0 -e prod -t k8s -v v1.0.0"
}

# 解析参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -e|--env)
            DEPLOY_ENV="$2"
            shift 2
            ;;
        -t|--type)
            DEPLOY_TYPE="$2"
            shift 2
            ;;
        -v|--version)
            IMAGE_TAG="$2"
            shift 2
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            usage
            exit 1
            ;;
    esac
done

echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN}  OceanEngine Backend Deploy Script${NC}"
echo -e "${GREEN}============================================${NC}"
echo -e "${YELLOW}Environment: ${DEPLOY_ENV}${NC}"
echo -e "${YELLOW}Deploy Type: ${DEPLOY_TYPE}${NC}"
echo -e "${YELLOW}Image Tag: ${IMAGE_TAG}${NC}"
echo ""

# Docker 部署
deploy_docker() {
    echo -e "${YELLOW}Deploying with Docker Compose...${NC}"
    
    cd deployments/docker
    
    # 构建镜像
    echo -e "${YELLOW}Building Docker image...${NC}"
    docker build -t "${IMAGE_NAME}:${IMAGE_TAG}" -f Dockerfile ../..
    
    # 停止旧容器
    echo -e "${YELLOW}Stopping old containers...${NC}"
    docker-compose down || true
    
    # 启动新容器
    echo -e "${YELLOW}Starting new containers...${NC}"
    docker-compose up -d
    
    # 等待服务启动
    echo -e "${YELLOW}Waiting for services to start...${NC}"
    sleep 10
    
    # 检查健康状态
    echo -e "${YELLOW}Checking health status...${NC}"
    if curl -sf http://localhost:8080/health > /dev/null; then
        echo -e "${GREEN}Service is healthy!${NC}"
    else
        echo -e "${RED}Service health check failed!${NC}"
        docker-compose logs
        exit 1
    fi
    
    echo -e "${GREEN}Docker deployment completed!${NC}"
}

# Kubernetes 部署
deploy_k8s() {
    echo -e "${YELLOW}Deploying to Kubernetes...${NC}"
    
    cd deployments/kubernetes
    
    # 创建命名空间
    kubectl create namespace oceanengine --dry-run=client -o yaml | kubectl apply -f -
    
    # 更新镜像标签
    sed -i.bak "s|image: .*|image: ${IMAGE_NAME}:${IMAGE_TAG}|g" deployment.yaml
    
    # 应用配置
    echo -e "${YELLOW}Applying Kubernetes manifests...${NC}"
    kubectl apply -f service.yaml
    kubectl apply -f deployment.yaml
    
    # 等待部署完成
    echo -e "${YELLOW}Waiting for deployment to complete...${NC}"
    kubectl -n oceanengine rollout status deployment/oceanengine-api --timeout=300s
    
    # 检查 Pod 状态
    echo -e "${YELLOW}Checking pod status...${NC}"
    kubectl -n oceanengine get pods -l app=oceanengine-api
    
    echo -e "${GREEN}Kubernetes deployment completed!${NC}"
}

# 执行部署
case $DEPLOY_TYPE in
    docker)
        deploy_docker
        ;;
    k8s|kubernetes)
        deploy_k8s
        ;;
    *)
        echo -e "${RED}Unknown deploy type: ${DEPLOY_TYPE}${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN}  Deployment completed successfully!${NC}"
echo -e "${GREEN}============================================${NC}"
