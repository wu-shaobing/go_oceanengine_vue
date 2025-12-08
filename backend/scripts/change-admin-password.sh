#!/bin/bash
# OceanEngine 管理员密码修改脚本
# 
# 使用方式:
#   chmod +x change-admin-password.sh
#   ./change-admin-password.sh
#
# 或者直接提供新密码:
#   ./change-admin-password.sh "your_new_secure_password"
#
# 支持通过环境变量配置数据库连接:
#   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME

set -e

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BACKEND_DIR="$(dirname "$SCRIPT_DIR")"

# 加载 .env 文件 (如果存在)
if [ -f "$BACKEND_DIR/.env" ]; then
    log_info "加载环境变量从 $BACKEND_DIR/.env"
    export $(grep -v '^#' "$BACKEND_DIR/.env" | xargs)
fi

# 数据库配置 (环境变量或默认值)
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_USER="${DB_USER:-root}"
DB_PASSWORD="${DB_PASSWORD:-}"
DB_NAME="${DB_NAME:-oceanengine}"
ADMIN_USERNAME="${ADMIN_USERNAME:-admin}"

# 检查 mysql 客户端
check_mysql() {
    if ! command -v mysql &> /dev/null; then
        log_error "mysql 客户端未安装，请先安装 mysql-client"
        echo "Ubuntu/Debian: sudo apt install mysql-client"
        echo "CentOS/RHEL: sudo yum install mysql"
        echo "macOS: brew install mysql-client"
        exit 1
    fi
}

# 检查数据库连接
check_db_connection() {
    log_info "检查数据库连接..."
    if ! mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" ${DB_PASSWORD:+-p"$DB_PASSWORD"} -e "SELECT 1" &>/dev/null; then
        log_error "无法连接数据库"
        echo "请检查以下配置:"
        echo "  - DB_HOST: $DB_HOST"
        echo "  - DB_PORT: $DB_PORT"
        echo "  - DB_USER: $DB_USER"
        echo "  - DB_NAME: $DB_NAME"
        exit 1
    fi
    log_info "数据库连接成功"
}

# 密码强度检查
check_password_strength() {
    local password="$1"
    local errors=()
    
    if [ ${#password} -lt 8 ]; then
        errors+=("密码长度至少 8 位")
    fi
    
    if [ ${#password} -gt 32 ]; then
        errors+=("密码长度不超过 32 位")
    fi
    
    if ! echo "$password" | grep -q '[0-9]'; then
        errors+=("密码需包含数字")
    fi
    
    if ! echo "$password" | grep -q '[a-z]'; then
        errors+=("密码需包含小写字母")
    fi
    
    if ! echo "$password" | grep -q '[A-Z]'; then
        errors+=("密码需包含大写字母")
    fi
    
    if [ ${#errors[@]} -gt 0 ]; then
        log_error "密码强度不足:"
        for err in "${errors[@]}"; do
            echo "  - $err"
        done
        return 1
    fi
    
    return 0
}

# 生成 bcrypt 哈希
generate_bcrypt_hash() {
    local password="$1"
    
    # 尝试使用 Go 生成 bcrypt 哈希
    if command -v go &> /dev/null; then
        local hash=$(go run -ldflags="-s -w" - "$password" 2>/dev/null << 'GOCODE'
package main

import (
    "fmt"
    "os"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    if len(os.Args) < 2 {
        os.Exit(1)
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), 10)
    if err != nil {
        os.Exit(1)
    }
    fmt.Print(string(hash))
}
GOCODE
)
        if [ -n "$hash" ]; then
            echo "$hash"
            return 0
        fi
    fi
    
    # 尝试使用 Python 生成 bcrypt 哈希
    if command -v python3 &> /dev/null; then
        local hash=$(python3 -c "
import sys
try:
    import bcrypt
    password = sys.argv[1].encode('utf-8')
    hashed = bcrypt.hashpw(password, bcrypt.gensalt(rounds=10))
    print(hashed.decode('utf-8'), end='')
except Exception as e:
    sys.exit(1)
" "$password" 2>/dev/null)
        if [ -n "$hash" ]; then
            echo "$hash"
            return 0
        fi
    fi
    
    # 尝试使用 htpasswd
    if command -v htpasswd &> /dev/null; then
        local hash=$(htpasswd -bnBC 10 "" "$password" 2>/dev/null | tr -d ':\n' | sed 's/$2y/$2a/')
        if [ -n "$hash" ]; then
            echo "$hash"
            return 0
        fi
    fi
    
    log_error "无法生成 bcrypt 哈希"
    echo "请安装以下工具之一:"
    echo "  - Go: 推荐，安装后自动使用"
    echo "  - Python + bcrypt: pip3 install bcrypt"
    echo "  - htpasswd: sudo apt install apache2-utils"
    return 1
}

# 更新密码
update_password() {
    local new_password="$1"
    local bcrypt_hash="$2"
    
    log_info "更新管理员密码..."
    
    # SQL 安全转义
    local escaped_hash=$(printf '%s' "$bcrypt_hash" | sed "s/'/\\\\'/g")
    
    mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" ${DB_PASSWORD:+-p"$DB_PASSWORD"} "$DB_NAME" << EOF
UPDATE sys_users 
SET password = '$escaped_hash', 
    updated_at = NOW()
WHERE username = '$ADMIN_USERNAME';
EOF
    
    if [ $? -eq 0 ]; then
        log_info "密码更新成功"
        return 0
    else
        log_error "密码更新失败"
        return 1
    fi
}

# 验证用户存在
check_admin_exists() {
    local count=$(mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" ${DB_PASSWORD:+-p"$DB_PASSWORD"} "$DB_NAME" -N -e \
        "SELECT COUNT(*) FROM sys_users WHERE username = '$ADMIN_USERNAME';" 2>/dev/null)
    
    if [ "$count" != "1" ]; then
        log_error "未找到管理员账号: $ADMIN_USERNAME"
        return 1
    fi
    return 0
}

# 主流程
main() {
    echo ""
    echo "=============================================="
    echo "  OceanEngine 管理员密码修改工具"
    echo "=============================================="
    echo ""
    
    # 检查依赖
    check_mysql
    check_db_connection
    check_admin_exists
    
    # 获取新密码
    local new_password=""
    
    if [ -n "$1" ]; then
        # 从参数获取
        new_password="$1"
    else
        # 交互式输入
        echo "密码要求:"
        echo "  - 长度 8-32 位"
        echo "  - 包含大写字母"
        echo "  - 包含小写字母"
        echo "  - 包含数字"
        echo ""
        
        while true; do
            read -s -p "请输入新密码: " new_password
            echo ""
            read -s -p "请再次输入新密码: " confirm_password
            echo ""
            
            if [ "$new_password" != "$confirm_password" ]; then
                log_error "两次输入的密码不一致，请重试"
                continue
            fi
            
            if check_password_strength "$new_password"; then
                break
            fi
        done
    fi
    
    # 检查密码强度
    if ! check_password_strength "$new_password"; then
        exit 1
    fi
    
    # 生成哈希
    log_info "生成密码哈希..."
    local bcrypt_hash=$(generate_bcrypt_hash "$new_password")
    
    if [ -z "$bcrypt_hash" ]; then
        exit 1
    fi
    
    # 更新数据库
    if update_password "$new_password" "$bcrypt_hash"; then
        echo ""
        echo "=============================================="
        log_info "管理员密码已更新!"
        echo "=============================================="
        echo ""
        echo "用户名: $ADMIN_USERNAME"
        echo ""
        log_warn "请妥善保管新密码，此处不再显示"
        echo ""
    else
        exit 1
    fi
}

main "$@"
