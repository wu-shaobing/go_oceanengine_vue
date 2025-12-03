# MCP服务器错误修复指南

## 问题诊断

根据测试结果，以下8个MCP服务器无法正常工作，主要原因是：

1. **缺少API密钥或环境变量配置**
2. **包名错误**（mem0-memory使用了错误的包名）
3. **缺少必要的命令行参数**（如GitHub需要stdio参数）

## 已修复的问题

### ✅ 1. 包名修正
- **mem0-memory**: 从 `openmemory` 更正为 `@mem0/mcp`

### ✅ 2. 添加了必要的环境变量配置
所有需要API密钥的服务器都已添加环境变量占位符。

### ✅ 3. 添加了必要的命令行参数
- **GitHub**: 添加了 `stdio` 参数
- **Stripe**: 添加了 `--tools=all` 参数

## 需要配置API密钥的服务器

### 🔑 1. Neon
**环境变量**: `NEON_API_KEY`
**获取方式**: 
- 访问 https://console.neon.tech
- 登录后进入 Settings > API Keys
- 创建新的API密钥

### 🔑 2. Supabase Memory
**环境变量**: 
- `SUPABASE_URL` - 您的Supabase项目URL
- `SUPABASE_KEY` - 您的Supabase API密钥
**获取方式**:
- 访问 https://supabase.com
- 登录后进入项目设置
- 在API设置中找到项目URL和anon/public key

### 🔑 3. Mem0 Memory
**环境变量**: `MEM0_API_KEY`
**获取方式**:
- 访问 https://mem0.ai
- 注册账号并获取API密钥

### 🔑 4. Replicate
**环境变量**: `REPLICATE_API_TOKEN`
**获取方式**:
- 访问 https://replicate.com/account/api-tokens
- 登录后创建API令牌

### 🔑 5. Perplexity
**环境变量**: `PERPLEXITY_API_KEY`
**获取方式**:
- 访问 https://www.perplexity.ai/settings/api
- 登录后创建API密钥

### 🔑 6. Cloudflare
**环境变量**: 
- `CLOUDFLARE_API_TOKEN` - Cloudflare API令牌
- `CLOUDFLARE_ACCOUNT_ID` - Cloudflare账户ID
**获取方式**:
- 访问 https://developers.cloudflare.com/fundamentals/api/get-started/create-token/
- 创建API令牌并获取账户ID

### 🔑 7. GitHub
**环境变量**: `GITHUB_PERSONAL_ACCESS_TOKEN`
**获取方式**:
- 访问 https://github.com/settings/tokens
- 点击 "Generate new token (classic)"
- 选择以下权限：
  - ✅ `repo` - 完整访问私有仓库
  - ✅ `read:org` - 读取组织和团队成员信息
  - ✅ `user` - 更新用户数据

### 🔑 8. Stripe
**环境变量**: `STRIPE_SECRET_KEY`
**获取方式**:
- 访问 https://dashboard.stripe.com/apikeys
- 登录后复制Secret key（以sk_开头）

## 配置步骤

### 方法1: 在Cursor配置文件中直接填写（推荐）

编辑 `/Users/wushaobing911/.cursor/mcp.json` 文件，将空字符串 `""` 替换为您的实际API密钥：

```json
{
  "mcpServers": {
    "neon": {
      "command": "npx",
      "args": ["-y", "@neondatabase/mcp-server-neon"],
      "env": {
        "NEON_API_KEY": "your_actual_api_key_here"
      }
    }
  }
}
```

### 方法2: 使用环境变量（更安全）

在您的shell配置文件中（如 `~/.zshrc` 或 `~/.bash_profile`）添加：

```bash
export NEON_API_KEY="your_api_key"
export SUPABASE_URL="your_supabase_url"
export SUPABASE_KEY="your_supabase_key"
# ... 其他环境变量
```

然后在配置文件中引用：

```json
{
  "mcpServers": {
    "neon": {
      "command": "npx",
      "args": ["-y", "@neondatabase/mcp-server-neon"],
      "env": {
        "NEON_API_KEY": "${NEON_API_KEY}"
      }
    }
  }
}
```

## 验证配置

配置完成后：

1. **重启Cursor应用** - 使配置生效
2. **检查MCP服务器状态** - 在Cursor设置中查看服务器是否正常启动
3. **查看错误日志** - 如果仍有错误，点击 "Show Output" 查看详细错误信息

## 已知问题

### Cloudflare MCP服务器
`@cloudflare/mcp-server-cloudflare` 可能存在依赖问题（zod版本冲突）。如果无法解决，可以：
- 等待包更新
- 或暂时禁用该服务器

## 安全提示

⚠️ **重要**: 
- 永远不要将API密钥提交到版本控制系统（如Git）
- 使用环境变量或安全的密钥管理工具
- 定期轮换API密钥
- 只授予必要的权限

## 配置完成后

配置完所有API密钥后，重启Cursor，这些MCP服务器应该能够正常工作。如果仍有问题，请查看具体的错误日志。


