# MCP服务器包名修正说明

## 问题诊断

检查发现之前配置的MCP服务器包名不正确，导致服务器无法启动。已将所有包名更新为正确的npm包名。

## 已修正的包名

### ✅ 已确认存在的包：

1. **chrome-devtools**: `chrome-devtools-mcp` ✅
2. **figma**: `figma-mcp` ✅
3. **supabase-memory**: `@supabase/mcp-server-supabase` ✅
4. **context7**: `@upstash/context7-mcp` ✅
5. **mem0-memory**: `openmemory` ✅
6. **replicate**: `replicate-mcp` ✅
7. **perplexity**: `perplexity-mcp` ✅
8. **cloudflare**: `@cloudflare/mcp-server-cloudflare` ✅
9. **github**: `github-mcp-custom` ✅
10. **stripe**: `@stripe/mcp` ✅
11. **databutton**: `@iflow-mcp/databutton-mcp` ✅
12. **semgrep**: `mcp-server-semgrep` ✅
13. **semantic-scholar**: `researchmcp` ✅ (支持Semantic Scholar、arXiv、PubMed)

### ⚠️ 需要验证的包：

1. **neon**: `neon-init` - 这个包可能不是MCP服务器，可能需要使用其他方式配置
2. **vercel-weather**: `@iflow-mcp/weather-mcp` - 已更新为天气MCP服务器

## 系统要求

- ✅ Node.js v22.19.0 (已安装)
- ✅ npm 10.9.3 (已安装)
- ✅ npx 10.9.3 (已安装)

## 下一步

1. 重启Cursor应用以使配置生效
2. 检查MCP服务器状态
3. 某些服务器可能需要API密钥，请参考ENV_CONFIG.md

## 注意事项

- 所有包都使用`npx -y`运行，首次使用时会自动下载
- 如果某些服务器仍然无法启动，可能需要：
  - 检查是否需要API密钥
  - 验证包是否正确支持MCP协议
  - 查看服务器日志以获取详细错误信息



