#!/bin/bash
# OceanEngine SSL 证书配置脚本
# 使用 Let's Encrypt 免费证书
# 
# 使用方式:
#   chmod +x setup-ssl.sh
#   sudo ./setup-ssl.sh your-domain.com your-email@example.com

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

# 检查参数
if [ -z "$1" ] || [ -z "$2" ]; then
    echo "使用方式: $0 <domain> <email>"
    echo "示例: $0 oceanengine.example.com admin@example.com"
    exit 1
fi

DOMAIN=$1
EMAIL=$2
NGINX_CONF_SRC="$(dirname "$0")/oceanengine.conf"
NGINX_CONF_DEST="/etc/nginx/sites-available/oceanengine.conf"
CERTBOT_WEBROOT="/var/www/certbot"
FRONTEND_ROOT="/var/www/oceanengine"

# 检查 root 权限
if [ "$EUID" -ne 0 ]; then
    log_error "请使用 sudo 运行此脚本"
    exit 1
fi

# 检查系统
if [ -f /etc/debian_version ]; then
    PKG_MANAGER="apt"
    PKG_UPDATE="apt update"
    PKG_INSTALL="apt install -y"
elif [ -f /etc/redhat-release ]; then
    PKG_MANAGER="yum"
    PKG_UPDATE="yum check-update || true"
    PKG_INSTALL="yum install -y"
else
    log_error "不支持的系统，请手动安装 certbot 和 nginx"
    exit 1
fi

log_info "开始配置 SSL 证书 for ${DOMAIN}"

# 1. 安装依赖
log_info "安装依赖..."
$PKG_UPDATE
if [ "$PKG_MANAGER" = "apt" ]; then
    $PKG_INSTALL nginx certbot python3-certbot-nginx
else
    $PKG_INSTALL nginx certbot python3-certbot-nginx epel-release
fi

# 2. 创建必要目录
log_info "创建必要目录..."
mkdir -p $CERTBOT_WEBROOT
mkdir -p $FRONTEND_ROOT
mkdir -p /etc/nginx/sites-available
mkdir -p /etc/nginx/sites-enabled

# 3. 复制并配置 Nginx
log_info "配置 Nginx..."
if [ -f "$NGINX_CONF_SRC" ]; then
    cp "$NGINX_CONF_SRC" "$NGINX_CONF_DEST"
else
    log_error "Nginx 配置模板不存在: $NGINX_CONF_SRC"
    exit 1
fi

# 替换域名
sed -i "s/YOUR_DOMAIN/${DOMAIN}/g" "$NGINX_CONF_DEST"

# 创建软链接
ln -sf "$NGINX_CONF_DEST" /etc/nginx/sites-enabled/oceanengine.conf

# 删除默认站点 (如果存在)
rm -f /etc/nginx/sites-enabled/default 2>/dev/null || true

# 4. 临时禁用 SSL 配置以便验证
log_info "准备 HTTP 验证..."
# 创建临时的 HTTP-only 配置
cat > /etc/nginx/sites-available/oceanengine-temp.conf << EOF
server {
    listen 80;
    listen [::]:80;
    server_name ${DOMAIN} www.${DOMAIN};
    
    location /.well-known/acme-challenge/ {
        root ${CERTBOT_WEBROOT};
    }
    
    location / {
        root ${FRONTEND_ROOT};
        index index.html;
    }
}
EOF

ln -sf /etc/nginx/sites-available/oceanengine-temp.conf /etc/nginx/sites-enabled/oceanengine-temp.conf
rm -f /etc/nginx/sites-enabled/oceanengine.conf

# 测试并重载 Nginx
nginx -t || {
    log_error "Nginx 配置测试失败"
    exit 1
}
systemctl reload nginx

# 5. 申请证书
log_info "申请 Let's Encrypt 证书..."
certbot certonly \
    --webroot \
    --webroot-path=$CERTBOT_WEBROOT \
    -d $DOMAIN \
    -d www.$DOMAIN \
    --email $EMAIL \
    --agree-tos \
    --non-interactive \
    --force-renewal || {
        log_warn "www 子域名可能不可用，尝试仅申请主域名..."
        certbot certonly \
            --webroot \
            --webroot-path=$CERTBOT_WEBROOT \
            -d $DOMAIN \
            --email $EMAIL \
            --agree-tos \
            --non-interactive
    }

# 6. 切换到完整的 HTTPS 配置
log_info "启用 HTTPS 配置..."
rm -f /etc/nginx/sites-enabled/oceanengine-temp.conf
ln -sf "$NGINX_CONF_DEST" /etc/nginx/sites-enabled/oceanengine.conf

# 测试并重载 Nginx
nginx -t || {
    log_error "HTTPS Nginx 配置测试失败"
    exit 1
}
systemctl reload nginx

# 7. 设置自动续期
log_info "配置证书自动续期..."
if [ "$PKG_MANAGER" = "apt" ]; then
    # Debian/Ubuntu 自动配置了 systemd timer
    systemctl enable certbot.timer
    systemctl start certbot.timer
else
    # CentOS/RHEL 使用 cron
    if ! crontab -l 2>/dev/null | grep -q "certbot renew"; then
        (crontab -l 2>/dev/null; echo "0 0,12 * * * /usr/bin/certbot renew --quiet --post-hook 'systemctl reload nginx'") | crontab -
    fi
fi

# 8. 验证 SSL
log_info "验证 SSL 配置..."
sleep 2
if curl -s -o /dev/null -w "%{http_code}" "https://${DOMAIN}" | grep -q "200\|301\|302"; then
    log_info "HTTPS 配置成功!"
else
    log_warn "HTTPS 可能需要等待 DNS 生效或检查防火墙配置"
fi

# 完成
echo ""
echo "=============================================="
log_info "SSL 证书配置完成!"
echo "=============================================="
echo ""
echo "证书路径:"
echo "  - 证书: /etc/letsencrypt/live/${DOMAIN}/fullchain.pem"
echo "  - 私钥: /etc/letsencrypt/live/${DOMAIN}/privkey.pem"
echo ""
echo "Nginx 配置: ${NGINX_CONF_DEST}"
echo "前端文件目录: ${FRONTEND_ROOT}"
echo ""
echo "下一步:"
echo "  1. 将前端构建文件复制到 ${FRONTEND_ROOT}"
echo "     cp -r frontend/dist/* ${FRONTEND_ROOT}/"
echo ""
echo "  2. 更新 OceanEngine OAuth 回调地址为:"
echo "     https://${DOMAIN}/auth/callback"
echo ""
echo "  3. 重启后端服务:"
echo "     cd backend && docker compose restart"
echo ""
echo "证书将每 12 小时自动检查续期"
echo "=============================================="
