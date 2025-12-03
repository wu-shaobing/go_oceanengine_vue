# MCP Servers Configuration

本项目包含15个MCP（Model Context Protocol）服务器的配置。

## 已配置的MCP服务器列表

1. **Chrome DevTools** - Chrome开发者工具MCP服务器
2. **Neon** - Neon数据库MCP服务器
3. **Supabase Memory** - Supabase内存MCP服务器
4. **Figma** - Figma设计工具MCP服务器
5. **Context7** - Context7内存MCP服务器
6. **Mem0 Memory** - Mem0内存MCP服务器 (@memOai/mem0-memory-mcp)
7. **Replicate** - Replicate AI模型MCP服务器
8. **Vercel Weather** - Vercel天气服务MCP服务器
9. **Perplexity** - Perplexity搜索MCP服务器
10. **Cloudflare** - Cloudflare服务MCP服务器
11. **GitHub** - GitHub集成MCP服务器
12. **Stripe** - Stripe支付MCP服务器
13. **Databutton** - Databutton应用构建MCP服务器
14. **Semgrep** - Semgrep代码分析MCP服务器
15. **Semantic Scholar** - 学术论文搜索MCP服务器 (@hamid-vakilzadeh/mcpsemanticscholar)

## 安装说明

### 方法1: 使用npx直接运行（推荐）

这些MCP服务器已经配置为使用`npx`直接运行，无需预先安装。当需要使用时，会自动下载并运行。

### 方法2: 全局安装

如果需要全局安装某个MCP服务器，可以使用：

```bash
npm install -g <package-name>
```

### 方法3: 本地安装

在项目目录中安装：

```bash
cd /Users/wushaobing911/Desktop/douyin/.mcp
npm install
```

## 配置说明

配置文件 `mcp.json` 包含了所有15个MCP服务器的配置。每个服务器都使用`npx`命令运行，确保始终使用最新版本。

## 使用说明

在支持MCP的客户端（如Claude Desktop）中，可以将此配置文件路径添加到MCP服务器配置中。

## 注意事项

- 某些MCP服务器可能需要API密钥或环境变量配置
- 请参考各个MCP服务器的官方文档获取详细的配置说明
- 部分包名可能需要根据实际情况调整




