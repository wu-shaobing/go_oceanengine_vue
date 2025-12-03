# MCP服务器环境变量配置指南

## 概述

某些MCP服务器可能需要API密钥或环境变量才能正常工作。以下是各个服务器可能需要的配置：

## 需要配置的MCP服务器

### 1. Chrome DevTools
- **通常不需要API密钥**
- 直接使用即可

### 2. Neon
- **NEON_API_KEY**: Neon数据库API密钥
- **NEON_DATABASE_URL**: Neon数据库连接URL

### 3. Supabase Memory
- **SUPABASE_URL**: Supabase项目URL
- **SUPABASE_KEY**: Supabase API密钥

### 4. Figma
- **FIGMA_ACCESS_TOKEN**: Figma访问令牌
- 获取方式: https://www.figma.com/developers/api#access-tokens

### 5. Context7 Memory
- **CONTEXT7_API_KEY**: Context7 API密钥

### 6. Mem0 Memory (@memOai/mem0-memory-mcp)
- **MEM0_API_KEY**: Mem0 API密钥

### 7. Replicate
- **REPLICATE_API_TOKEN**: Replicate API令牌
- 获取方式: https://replicate.com/account/api-tokens

### 8. Vercel Weather
- **通常不需要API密钥**
- 直接使用即可

### 9. Perplexity
- **PERPLEXITY_API_KEY**: Perplexity API密钥
- 获取方式: https://www.perplexity.ai/settings/api

### 10. Cloudflare
- **CLOUDFLARE_API_TOKEN**: Cloudflare API令牌
- **CLOUDFLARE_ACCOUNT_ID**: Cloudflare账户ID
- 获取方式: https://developers.cloudflare.com/fundamentals/api/get-started/create-token/

### 11. GitHub
- **GITHUB_PERSONAL_ACCESS_TOKEN**: GitHub个人访问令牌
- 获取方式: https://github.com/settings/tokens

### 12. Stripe
- **STRIPE_API_KEY**: Stripe API密钥
- 获取方式: https://dashboard.stripe.com/apikeys

### 13. Databutton
- **DATABUTTON_API_KEY**: Databutton API密钥

### 14. Semgrep
- **SEMGREP_API_TOKEN**: Semgrep API令牌
- 获取方式: https://semgrep.dev/orgs/-/settings/api-tokens

### 15. Semantic Scholar (@hamid-vakilzadeh/mcpsemanticscholar)
- **SEMANTIC_SCHOLAR_API_KEY**: Semantic Scholar API密钥（可选）
- 获取方式: https://www.semanticscholar.org/product/api

## 配置方法

### 方法1: 环境变量
在运行MCP服务器的环境中设置环境变量：

```bash
export NEON_API_KEY="your_key_here"
export SUPABASE_URL="your_url_here"
# ... 其他环境变量
```

### 方法2: 在mcp.json中配置
某些MCP服务器支持在配置文件中直接设置环境变量：

```json
{
  "mcpServers": {
    "neon": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-neon"],
      "env": {
        "NEON_API_KEY": "your_key_here"
      }
    }
  }
}
```

### 方法3: 使用.env文件
如果MCP服务器支持，可以创建`.env`文件（注意：不要提交到版本控制）：

```bash
NEON_API_KEY=your_key_here
SUPABASE_URL=your_url_here
```

## 注意事项

1. **安全性**: 永远不要将API密钥提交到版本控制系统
2. **权限**: 确保API密钥具有适当的权限范围
3. **限制**: 某些API可能有使用限制或费用
4. **文档**: 请参考各个MCP服务器的官方文档获取最新配置要求

## 验证配置

配置完成后，可以通过以下方式验证：

1. 检查MCP服务器是否正常启动
2. 查看日志输出是否有错误信息
3. 测试MCP服务器的基本功能




