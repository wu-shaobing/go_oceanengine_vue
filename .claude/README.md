# Claude Code Configuration

这个目录包含了 Claude Code AI 助手的自定义配置，用于增强千川SDK管理平台的开发体验。

## 📂 目录结构

```
.claude/
├── agents/                      # AI 智能体（6个）
│   ├── backend-developer.md     # Go 后端开发专家 ⭐ NEW
│   ├── frontend-developer.md    # React 前端开发专家
│   ├── api-tester.md           # API 测试专家 ⭐ NEW
│   ├── code-reviewer.md        # 代码审查专家
│   ├── researcher.md           # 研究助手
│   └── factory-guide.md        # 工厂编排器
├── commands/                    # 自定义命令（3个）
│   ├── analyze-quality.md      # 代码质量分析
│   ├── dev-start.md           # 启动开发环境 ⭐ NEW
│   └── dev-stop.md            # 停止开发环境 ⭐ NEW
├── skills/                      # 自定义技能（1个）
│   └── code-quality-analyzer/  # 代码质量分析器
├── hooks/                       # 工作流钩子（空）
├── mcp-servers.json            # MCP 服务器配置 ⚠️ 包含凭据
├── mcp-servers.example.json    # MCP 配置模板 ⭐ NEW
├── settings.json               # 全局设置
├── settings.local.json         # 本地权限设置
└── README.md                   # 本文件
```

## 🚀 快速开始

### 1. 安全配置

**⚠️ 重要：移除明文凭据**

```bash
# 1. 备份当前配置
cp .claude/mcp-servers.json .claude/mcp-servers.json.backup

# 2. 使用模板替换
cp .claude/mcp-servers.example.json .claude/mcp-servers.json

# 3. 设置环境变量
export GITHUB_TOKEN="your_github_token"
export STRIPE_KEY="your_stripe_key"

# 4. 添加到 .gitignore
echo ".claude/mcp-servers.json" >> .gitignore
echo ".claude/settings.json" >> .gitignore
echo ".claude/settings.local.json" >> .gitignore
```

### 2. 使用智能体

智能体会根据你的请求自动激活，无需手动调用。

#### **Backend Developer** (Go 专家)
```
触发词：
- "Build a backend API"
- "Create an endpoint"
- "Implement a handler"

示例：
> 帮我创建一个获取广告列表的 API 端点
```

#### **Frontend Developer** (React 专家)
```
触发词：
- "Build a React component"
- "Create a UI for..."
- "Implement frontend logic"

示例：
> 创建一个广告列表展示组件
```

#### **API Tester** (测试专家)
```
触发词：
- "Test the API"
- "Write tests for..."
- "Add unit tests"

示例：
> 为 AdHandler.List 方法写单元测试
```

### 3. 使用命令

#### 代码质量分析
```bash
# 分析整个项目
/analyze-quality --language go --path ./backend/internal

# 只检查安全性
/analyze-quality --language typescript --path ./frontend/src --check security

# 设置质量门槛
/analyze-quality --language go --path ./backend --min-score 80
```

#### 开发环境管理
```bash
# 启动开发环境（前后端同时启动）
/dev-start

# 停止所有服务
/dev-stop
```

## 🤖 智能体详情

### 1. Backend Developer (Go 专家)
- **模型**: Sonnet
- **专长**: Gin、REST API、qianchuanSDK 集成
- **功能**:
  - 创建 Handler/Service/Middleware
  - 实现 Session 管理
  - SDK 错误处理
  - 编写单元测试

### 2. Frontend Developer (React 专家)
- **模型**: Sonnet
- **专长**: React 18、TypeScript、Zustand、Tailwind
- **功能**:
  - 组件开发
  - 状态管理
  - API 集成
  - 性能优化

### 3. API Tester (测试专家)
- **模型**: Sonnet
- **专长**: Go 测试、Vitest、Playwright
- **功能**:
  - 单元测试
  - 集成测试
  - E2E 测试
  - Mock 数据生成

### 4. Code Reviewer
- **模型**: Haiku
- **专长**: 代码审查、最佳实践
- **功能**: 代码质量检查、重构建议

### 5. Researcher
- **模型**: Haiku
- **专长**: 技术研究、文档查找
- **功能**: API 文档查询、技术方案调研

### 6. Factory Guide
- **模型**: Haiku
- **专长**: 工作流编排
- **功能**: 引导创建 Skills、Agents、Prompts、Hooks

## 🛠️ MCP 服务集成

项目集成了 14 个 MCP 服务：

| 服务 | 用途 | 状态 |
|------|------|------|
| `filesystem` | 文件系统访问 | ✅ 活跃 |
| `grep` | 代码搜索 | ✅ 活跃 |
| `duckduckgo` | 网络搜索 | ✅ 活跃 |
| `sqlite` | 本地数据库 | ✅ 活跃 |
| `postgres` | PostgreSQL | 🟡 可选 |
| `browser` | 浏览器自动化 | ✅ 活跃 |
| `github` | GitHub API | ⚠️ 需要 Token |
| `stripe` | Stripe 支付 | ⚠️ 需要 Key |
| `figma` | Figma 设计 | 🟡 可选 |
| `ollama-local` | 本地 LLM | 🟡 可选 |
| `semgrep` | 安全扫描 | ✅ 活跃 |
| `semantic-scholar` | 学术研究 | 🟡 可选 |
| `vercel-weather` | 天气 API | 🟡 可选 |
| `databutton` | 数据分析 | 🟡 可选 |

## 📋 代码质量技能

### Code Quality Analyzer

全面的代码质量评估工具，支持 Python、JavaScript/TypeScript、Java、Go。

**分析维度**：
1. **复杂度** (25%): 圈复杂度、函数长度、代码重复
2. **安全性** (35%): SQL注入、XSS、硬编码密钥
3. **测试覆盖** (20%): 单元测试覆盖率
4. **最佳实践** (20%): 命名规范、错误处理

**使用示例**：
```bash
# 分析后端代码
/analyze-quality --language go --path ./backend/internal

# 完整分析前端
/analyze-quality --language typescript --path ./frontend/src \
  --check security complexity tests practices \
  --output frontend-quality.json

# CI/CD 质量门槛
/analyze-quality --language go --path ./backend --min-score 80
```

**评分系统**：
- **A (90-100)**: 优秀
- **B (80-89)**: 良好
- **C (70-79)**: 可接受
- **D (60-69)**: 较差
- **F (0-59)**: 严重问题

## 🎯 最佳实践

### 开发工作流

1. **启动开发环境**
   ```bash
   /dev-start
   ```

2. **开发新功能**
   ```
   > 帮我创建一个更新广告计划预算的 API
   # Backend Developer 会自动处理
   ```

3. **添加测试**
   ```
   > 为新的 API 端点写测试
   # API Tester 会创建完整测试套件
   ```

4. **代码审查**
   ```bash
   /analyze-quality --language go --path ./backend/internal/handler
   ```

5. **提交前检查**
   ```bash
   make test          # 运行所有测试
   make fmt           # 格式化代码
   /analyze-quality --min-score 75
   ```

### 安全注意事项

1. **永远不要提交凭据**
   - 使用 `.env` 文件
   - 使用环境变量
   - 添加到 `.gitignore`

2. **定期安全扫描**
   ```bash
   /analyze-quality --language go --path ./backend --check security
   ```

3. **审查 MCP 服务权限**
   - 检查 `settings.local.json` 的权限列表
   - 只允许必要的命令

## 🔧 自定义扩展

### 添加新智能体

1. 在 `.claude/agents/` 创建 Markdown 文件
2. 使用 YAML frontmatter 定义元数据：
   ```yaml
   ---
   name: my-agent
   description: Agent description
   tools: Read, Write, Edit, Bash
   model: sonnet
   color: blue
   field: implementation
   expertise: expert
   ---
   ```
3. 编写智能体的指令和示例

### 添加新命令

1. 在 `.claude/commands/` 创建 Markdown 文件
2. 定义命令元数据和用法
3. 命令会自动在 Claude Code 中可用

### 添加新技能

1. 使用 Factory Guide:
   ```
   /build-skill
   ```
2. 按提示创建技能定义、输入/输出格式
3. 技能会保存在 `.claude/skills/`

## 📚 相关文档

- **项目文档**: `../docs/`
- **项目规则**: `../WARP.md`
- **SDK 文档**: `../qianchuanSDK/README.md`
- **前端文档**: `../frontend/README.md`
- **后端文档**: `../backend/README.md`

## 🆘 故障排除

### Claude 没有调用正确的智能体
- 使用更明确的触发词
- 直接提及智能体名称："Use backend-developer agent to..."

### MCP 服务连接失败
```bash
# 检查服务是否可用
npx -y @modelcontextprotocol/server-filesystem --help

# 重启 Claude Code
```

### 权限被拒绝
- 检查 `settings.local.json` 的权限配置
- 确保命令在白名单中

## 🚧 待开发功能

- [ ] **Hooks**: 自动化工作流（提交前检查、自动测试）
- [ ] **更多智能体**: DevOps 专家、数据库专家
- [ ] **自定义规则**: 项目特定的代码规范检查
- [ ] **CI/CD 集成**: GitHub Actions 工作流

## 📝 更新日志

### 2024-11-13
- ✨ 新增 Backend Developer 智能体
- ✨ 新增 API Tester 智能体
- ✨ 新增 dev-start/dev-stop 命令
- 🔒 创建 mcp-servers.example.json 模板
- 📚 添加完整 README 文档

### 初始版本
- ✨ Frontend Developer 智能体
- ✨ Code Quality Analyzer 技能
- ✨ Factory Guide 编排器
- ✨ 14 个 MCP 服务集成

---

**维护者**: 千川SDK管理平台团队  
**更新日期**: 2024-11-13
