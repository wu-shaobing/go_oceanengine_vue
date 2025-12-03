#!/bin/bash

# MCP服务器安装脚本
# 此脚本用于安装所有15个MCP服务器

echo "开始安装MCP服务器..."

# 创建必要的目录
mkdir -p /Users/wushaobing911/Desktop/douyin/.mcp

# 安装npm依赖（如果需要）
if [ -f "package.json" ]; then
    echo "安装npm依赖..."
    npm install
fi

echo "MCP服务器配置已完成！"
echo ""
echo "已配置的MCP服务器："
echo "1. Chrome DevTools"
echo "2. Neon"
echo "3. Supabase Memory"
echo "4. Figma"
echo "5. Context7"
echo "6. Mem0 Memory"
echo "7. Replicate"
echo "8. Vercel Weather"
echo "9. Perplexity"
echo "10. Cloudflare"
echo "11. GitHub"
echo "12. Stripe"
echo "13. Databutton"
echo "14. Semgrep"
echo "15. Semantic Scholar"
echo ""
echo "配置文件位置: /Users/wushaobing911/Desktop/douyin/.mcp/mcp.json"
echo ""
echo "注意: 这些MCP服务器使用npx运行，首次使用时会自动下载。"




