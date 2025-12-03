#!/bin/bash

# OceanEngine Backend 数据库迁移脚本

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
CONFIG_PATH="${CONFIG_PATH:-config/settings.yml}"
ACTION="${1:-migrate}"

usage() {
    echo "Usage: $0 [ACTION]"
    echo ""
    echo "Actions:"
    echo "  migrate   Run database migrations (default)"
    echo "  fresh     Drop all tables and re-run migrations"
    echo "  seed      Seed the database with initial data"
    echo "  status    Show migration status"
    echo ""
    echo "Environment variables:"
    echo "  CONFIG_PATH   Path to config file (default: config/settings.yml)"
    echo ""
    echo "Examples:"
    echo "  $0 migrate"
    echo "  $0 fresh"
    echo "  $0 seed"
    echo "  CONFIG_PATH=config/settings.prod.yml $0 migrate"
}

# 检查配置文件
check_config() {
    if [ ! -f "$CONFIG_PATH" ]; then
        echo -e "${RED}Config file not found: ${CONFIG_PATH}${NC}"
        exit 1
    fi
}

# 构建迁移工具
build_migrate() {
    echo -e "${YELLOW}Building migrate tool...${NC}"
    go build -o ./bin/migrate ./cmd/migrate/main.go
}

# 运行迁移
run_migrate() {
    echo -e "${GREEN}============================================${NC}"
    echo -e "${GREEN}  OceanEngine Database Migration${NC}"
    echo -e "${GREEN}============================================${NC}"
    echo -e "${YELLOW}Config: ${CONFIG_PATH}${NC}"
    echo -e "${YELLOW}Action: ${ACTION}${NC}"
    echo ""

    check_config
    build_migrate

    echo -e "${YELLOW}Running migration...${NC}"
    ./bin/migrate -config "$CONFIG_PATH" -action "$ACTION"

    echo ""
    echo -e "${GREEN}Migration completed successfully!${NC}"
}

# 主逻辑
case $ACTION in
    migrate|fresh|seed)
        run_migrate
        ;;
    status)
        echo -e "${YELLOW}Migration status check is not implemented yet.${NC}"
        ;;
    -h|--help|help)
        usage
        exit 0
        ;;
    *)
        echo -e "${RED}Unknown action: ${ACTION}${NC}"
        usage
        exit 1
        ;;
esac
