# ✅ 修复检查清单

**日期**: 2025-11-13  
**状态**: [x] 已完成

---

## 立即执行 (1-2分钟)

- [x] **修复MCP配置文件硬编码密钥**
  - [x] Stripe API密钥 → `${STRIPE_SECRET_KEY}`
  - [x] GitHub Token → `${GITHUB_TOKEN}`
  - [x] Figma Token → `${FIGMA_TOKEN}`
  - [x] Postgres URL → `${POSTGRES_URL}`
  - [x] Semantic Scholar API Key → `${SEMANTIC_SCHOLAR_API_KEY}`

- [x] **更新 .gitignore**
  - [x] 添加 `.claude/mcp-servers.json`
  - [x] 添加 `.claude/settings.json`
  - [x] 添加 `.claude/settings.local.json`
  - [x] 添加 `.claude/hooks/`

- [x] **清理Git跟踪**
  - [x] `git rm --cached .claude/mcp-servers.json`
  - [x] `git rm --cached .claude/settings.json`
  - [x] `git rm --cached .env .env.local`

- [x] **设置文件权限**
  - [x] `chmod 600 .claude/settings.local.json`
  - [x] `chmod 600 .claude/mcp-servers.json`
  - [x] `chmod +x .claude/skills/code-quality-analyzer/analyzer.py`

---

## 环境变量设置 (5分钟)

- [ ] **设置必需环境变量**
  ```bash
  export GITHUB_TOKEN="your_github_token"
  export STRIPE_SECRET_KEY="your_stripe_key"
  export FIGMA_TOKEN="your_figma_token"  # 可选
  export POSTGRES_URL="postgresql://..."  # 可选
  export SEMANTIC_SCHOLAR_API_KEY="..."   # 可选
  ```

- [ ] **保存到shell配置文件**
  ```bash
  echo 'export GITHUB_TOKEN="your_token"' >> ~/.bashrc
  echo 'export STRIPE_SECRET_KEY="your_key"' >> ~/.bashrc
  source ~/.bashrc
  ```

---

## 验证 (1分钟)

- [x] **检查硬编码密钥**
  ```bash
  grep -r "pk_\|sk_\|AKIA" .claude/*.json
  # 应该无结果
  ```

- [x] **验证JSON语法**
  ```bash
  python3 -m json.tool .claude/mcp-servers.json
  # 应该无错误
  ```

- [x] **检查Git忽略**
  ```bash
  git check-ignore .claude/mcp-servers.json
  # 应该显示 .gitignore:xx:.claude/mcp-servers.json
  ```

- [x] **运行安全检查**
  ```bash
  bash scripts/security-check.sh
  ```

---

## 可选增强 (10分钟)

- [ ] **创建 .env 文件**
  ```bash
  cp .env.example .env
  # 编辑 .env 填入真实值
  ```

- [ ] **测试MCP服务**
  ```bash
  npx -y @modelcontextprotocol/server-filesystem --help
  ```

- [ ] **测试开发环境**
  ```bash
  /dev-start
  ```

---

## 团队协作

- [ ] **分享环境变量模板**
  - 发送 `.env.example` 给团队成员
  - 说明需要设置的环境变量

- [ ] **培训团队**
  - 介绍如何使用智能体
  - 演示代码质量分析功能

- [ ] **设置CI/CD** (可选)
  - 添加安全扫描到GitHub Actions
  - 配置代码质量检查

---

## 故障排除

### 权限被拒绝
```bash
# 重新设置权限
chmod 600 .claude/*.json
```

### MCP服务连接失败
```bash
# 检查环境变量
echo $GITHUB_TOKEN
echo $STRIPE_SECRET_KEY

# 如果为空，重新设置
export GITHUB_TOKEN="your_token"
```

### JSON语法错误
```bash
# 验证JSON
python3 -m json.tool .claude/mcp-servers.json

# 如果有错误，检查第X行
```

### Git忽略不生效
```bash
# 清理Git缓存
git rm --cached .claude/mcp-servers.json
git add .gitignore
git commit -m "fix: update gitignore"
```

---

## 📚 相关文档

- `.claude/README.md` - 完整使用指南
- `.claude/SECURITY_FIX.md` - 详细安全修复指南
- `.claude/VERIFICATION_REPORT.md` - 完整验证报告
- `SECURITY_FIX_COMPLETE.md` - 修复完成报告

---

**全部完成** ✅

**下一步**: 
1. 设置环境变量
2. 测试配置
3. 开始使用智能体

**预计总时间**: 10-15分钟

**支持**: 查看 `.claude/README.md` 或运行 `bash scripts/security-check.sh`
