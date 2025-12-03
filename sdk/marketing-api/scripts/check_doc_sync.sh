#!/bin/bash
# 文档-代码同步检查脚本
# 用途: 检查 api/ 下的函数是否都在对应的 .md 文档中有记录
# 用法: ./scripts/check_doc_sync.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

cd "$PROJECT_ROOT"

echo "=========================================="
echo "巨量引擎 SDK 文档-代码同步检查"
echo "=========================================="
echo ""

# 文档列表
DOCS="OCEANENGINE.md QIANCHUAN.md ENTERPRISE.md LOCAL.md STAR.md SERVE_MARKET.md"

# 获取文档对应的API目录
get_api_dirs() {
    local doc="$1"
    case "$doc" in
        "OCEANENGINE.md")
            echo "api/oauth api/advertiser api/agent api/campaign api/ad api/creative api/report api/dmp api/dpa api/file api/tools api/eventmanager api/audiencepackage api/keyword api/privativeword api/customercenter api/businessplatform api/majordomo api/sharedwallet api/clue api/assets api/v3 api/duoplus api/conversion api/track"
            ;;
        "QIANCHUAN.md")
            echo "api/qianchuan"
            ;;
        "ENTERPRISE.md")
            echo "api/enterprise"
            ;;
        "LOCAL.md")
            echo "api/local"
            ;;
        "STAR.md")
            echo "api/star"
            ;;
        "SERVE_MARKET.md")
            echo "api/servemarket"
            ;;
    esac
}

TOTAL_ISSUES=0

# 提取文档中的函数名
extract_doc_funcs() {
    local doc_file="$1"
    grep -oE '\[ *[A-Za-z0-9_./]+\(' "$doc_file" 2>/dev/null | \
        sed -E 's/\[ *([A-Za-z0-9_./]+)\(/\1/' | \
        awk -F'.' '{print $NF}' | \
        awk -F'/' '{print $NF}' | \
        sort -u
}

# 提取API目录中的公开函数名
extract_api_funcs() {
    local api_dirs="$1"
    for dir in $api_dirs; do
        if [ -d "$dir" ]; then
            grep -rE 'func\s+[A-Z][A-Za-z0-9_]*\s*\(' "$dir" 2>/dev/null | \
                sed -E 's/.*func\s+([A-Z][A-Za-z0-9_]*)\s*\(.*/\1/'
        fi
    done | sort -u
}

echo "检查各文档与对应API目录的同步状态..."
echo ""

for doc in $DOCS; do
    if [ ! -f "$doc" ]; then
        echo "⚠️  文档不存在: $doc"
        continue
    fi
    
    api_dirs=$(get_api_dirs "$doc")
    
    echo "📄 $doc"
    echo "   对应目录: $api_dirs"
    
    # 提取函数列表
    doc_funcs=$(extract_doc_funcs "$doc")
    api_funcs=$(extract_api_funcs "$api_dirs")
    
    doc_count=$(echo "$doc_funcs" | grep -c . || echo 0)
    api_count=$(echo "$api_funcs" | grep -c . || echo 0)
    
    echo "   文档中函数数: $doc_count"
    echo "   API中函数数: $api_count"
    
    # 检查API中有但文档中没有的函数
    missing_in_doc=$(comm -23 <(echo "$api_funcs") <(echo "$doc_funcs") 2>/dev/null | head -20)
    if [ -n "$missing_in_doc" ]; then
        missing_count=$(echo "$missing_in_doc" | wc -l | tr -d ' ')
        echo "   ⚠️  API中有但文档未记录的函数 (显示前20个):"
        echo "$missing_in_doc" | while read func; do
            echo "      - $func"
        done
        TOTAL_ISSUES=$((TOTAL_ISSUES + missing_count))
    fi
    
    # 检查文档中有但API中没有的函数
    missing_in_api=$(comm -13 <(echo "$api_funcs") <(echo "$doc_funcs") 2>/dev/null | head -20)
    if [ -n "$missing_in_api" ]; then
        # 过滤掉一些特殊情况 (如 Url, Active, WxaActive, Conversion, Attribution 等无 clt 参数的)
        filtered=$(echo "$missing_in_api" | grep -vE '^(Url|Active|WxaActive|Conversion|Attribution)$' || true)
        if [ -n "$filtered" ]; then
            filtered_count=$(echo "$filtered" | wc -l | tr -d ' ')
            echo "   ⚠️  文档中有但API目录中未找到的函数:"
            echo "$filtered" | while read func; do
                echo "      - $func"
            done
            TOTAL_ISSUES=$((TOTAL_ISSUES + filtered_count))
        fi
    fi
    
    if [ -z "$missing_in_doc" ] && [ -z "$filtered" ]; then
        echo "   ✅ 同步状态良好"
    fi
    
    echo ""
done

echo "=========================================="
if [ $TOTAL_ISSUES -eq 0 ]; then
    echo "✅ 所有文档与代码同步状态良好!"
else
    echo "⚠️  发现 $TOTAL_ISSUES 处潜在不同步"
    echo "   注意: 部分差异可能是正常的(如内部函数、别名等)"
fi
echo "=========================================="

exit 0
