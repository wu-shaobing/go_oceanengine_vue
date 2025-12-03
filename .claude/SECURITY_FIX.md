# 🔒 Claude Code 配置安全修复指南

## 🚨 发现的问题

### 1. 硬编码API密钥
**文件**: `.claude/mcp-servers.json`
**问题**: Stripe测试密钥被硬编码
```json
"STRIPE_SECRET_KEY": "pk_test_51SSf3L0SW9997rmE..."
```

### 2. 敏感配置文件未被Git忽略
**文件**: `.gitignore`
**问题**: `.claude/*` 和 `.env*` 没有被忽略
```bash
# 当前的 .gitignore
.env
.env.local
.env.*.local

# 缺少
.claude/mcp-servers.json
.claude/settings.json
.claude/settings.local.json
```

---

## ✅ 修复步骤

### 步骤1: 更新 .gitignore
```bash
# 追加到 .gitignore
cat >> .gitignore << 'EOF'

# Claude Code configuration files
.claude/mcp-servers.json
.claude/settings.json
.claude/settings.local.json
.claude/hooks/

# Environment files
.env
.env.local
.env.*.local
.env.production
.env.staging
EOF

# 重新初始化git跟踪
git rm --cached .claude/mcp-servers.json 2>/dev/null || true
git rm --cached .claude/settings.json 2>/dev/null || true
git rm --cached .claude/settings.local.json 2>/dev/null || true
git rm --cached .env 2>/dev/null || true
git rm --cached .env.local 2>/dev/null || true

git commit -m "security: remove sensitive files from git tracking"
```

### 步骤2: 替换MCP配置文件
```bash
# 备份当前配置
cp .claude/mcp-servers.json .claude/mcp-servers.json.backup

# 使用安全版本
cp .claude/mcp-servers.fixed.json .claude/mcp-servers.json

# 设置环境变量
export GITHUB_TOKEN="your_github_token"
export STRIPE_SECRET_KEY="your_stripe_key"
export FIGMA_TOKEN="your_figma_token"
export POSTGRES_URL="postgresql://..."
export SEMANTIC_SCHOLAR_API_KEY="your_api_key"

# 永久保存环境变量
cat >> ~/.bashrc << 'EOF'
export GITHUB_TOKEN="your_github_token"
export STRIPE_SECRET_KEY="your_stripe_key"
export FIGMA_TOKEN="your_figma_token"
export POSTGRES_URL="postgresql://..."
export SEMANTIC_SCHOLAR_API_KEY="your_api_key"
EOF

source ~/.bashrc
```

### 步骤3: 验证修复
```bash
# 检查配置文件
cat .claude/mcp-servers.json | grep -A1 STRIPE_SECRET_KEY
# 应该显示:
# "STRIPE_SECRET_KEY": "${STRIPE_SECRET_KEY}"

# 检查Git状态
git status --ignored | grep ".claude"
# 应该显示这些文件为ignored

# 测试MCP服务
cd /Users/wushaobing911/Desktop/douyin
npx -y @modelcontextprotocol/server-filesystem --help
```

---

## 🛡️ 安全最佳实践

### 1. 环境变量管理
```bash
# 创建 .env.example
cat > .env.example << 'EOF'
# Qianchuan API
APP_ID=your_app_id
APP_SECRET=your_app_secret

# Development
PORT=8080
GIN_MODE=debug

# Session
COOKIE_SECRET=random_32_byte_secret

# OAuth
VITE_OAUTH_APP_ID=your_app_id
VITE_OAUTH_REDIRECT_URI=http://localhost:3000/auth/callback

# MCP Services (optional)
GITHUB_TOKEN=your_github_token
STRIPE_SECRET_KEY=your_stripe_key
FIGMA_TOKEN=your_figma_token
POSTGRES_URL=postgresql://user:pass@host:port/db
SEMANTIC_SCHOLAR_API_KEY=your_api_key
EOF
```

### 2. MCP服务权限管理
```json
// .claude/settings.local.json
{
  "permissions": {
    "allow": [
      "Bash(find:*)",
      "Bash(git:*)",
      "Bash(npm:*)",
      "Bash(go:*)",
      "Bash(python3:*)"
    ],
    "deny": [
      "Bash(sudo:*)",
      "Bash(chmod:*)",
      "Bash(chown:*)"
    ],
    "ask": [
      "Bash(rm:*)",
      "Bash(del:*)"
    ]
  }
}
```

### 3. 定期安全检查
```bash
#!/bin/bash
# security-check.sh

echo "🔍 Running security checks..."

# 检查是否有硬编码密钥
if grep -r "pk_test_\|sk_\|AKIA" .claude/ 2>/dev/null; then
  echo "❌ WARNING: Hardcoded keys found!"
  exit 1
fi

# 检查配置文件是否被忽略
if git check-ignore .claude/mcp-servers.json 2>/dev/null; then
  echo "✅ mcp-servers.json is ignored"
else
  echo "❌ WARNING: mcp-servers.json is not ignored!"
  exit 1
fi

# 检查权限
if [ -f ".claude/settings.local.json" ]; then
  if [ "$(stat -c %a .claude/settings.local.json)" != "600" ]; then
    echo "⚠️  WARNING: settings.local.json should be 600"
    chmod 600 .claude/settings.local.json
  fi
fi

echo "✅ Security checks passed"
```

### 4. 自动化安全扫描
```yaml
# .github/workflows/security.yml
name: Security Check
on: [push, pull_request]
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Scan for hardcoded secrets
        run: |
          if grep -r "pk_\|sk_\|AKIA" .claude/ 2>/dev/null; then
            echo "❌ Hardcoded secrets found!"
            exit 1
          fi
          echo "✅ No hardcoded secrets"
```

---

## 📋 检查清单

### 安全配置
- [ ] 硬编码密钥已移除
- [ ] 环境变量已配置
- [ ] .gitignore 已更新
- [ ] 敏感文件已从Git跟踪中移除
- [ ] 配置文件权限为 600
- [ ] .env.example 已创建

### 功能验证
- [ ] MCP服务器可以正常启动
- [ ] 所有必需的环境变量已设置
- [ ] 代码质量分析技能可以运行
- [ ] 开发环境启动命令有效

### 文档更新
- [ ] README.md 已更新
- [ ] 安全修复指南已阅读
- [ ] 团队已了解安全最佳实践

---

## 🚨 如果密钥已泄露

### 立即行动
1. **撤销密钥**
   - GitHub: 在Settings > Developer settings > Personal access tokens中撤销
   - Stripe: 在Dashboard > Developers > API keys中撤销密钥
   - Figma: 在Account settings > Personal access tokens中撤销

2. **生成新密钥**
   - 创建新的API密钥
   - 更新环境变量
   - 更新密码管理器

3. **清理Git历史**
   ```bash
   # 永久删除敏感文件 (谨慎操作)
   git filter-branch --force --index-filter \
     'git rm --cached --ignore-unmatch .claude/mcp-servers.json' \
     --prune-empty --tag-name-filter cat -- --all

   # 强制推送 (仅在必要时)
   git push origin --force --all
   ```

---

## 📞 紧急联系

如果发现安全问题或需要帮助：
- 创建GitHub Issue: [Security] + 描述
- 联系安全负责人
- 立即撤销所有可能泄露的密钥

---

## 参考资源

- [Git忽略文件指南](https://git-scm.com/docs/gitignore)
- [MCP服务器文档](https://modelcontextprotocol.io/)
- [GitHub安全最佳实践](https://docs.github.com/en/authentication)
- [Stripe API密钥管理](https://stripe.com/docs/keys)

**更新日期**: 2025-11-13
**维护者**: 千川SDK管理平台团队
