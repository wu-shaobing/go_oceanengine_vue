# Cursor Rules 使用指南

这是一套完整的、科学的 AI 协作规则体系，帮助你使用 Cursor AI 生成高质量、规范化的代码。

## 📁 目录结构

```
.cursor/rules/
├── README.md                        # 本文件
├── ai-协作规则.mdc                   # 主协议（Always Apply）
├── basic/                           # 基础规范层（Always Apply）
│   ├── 代码质量规范.mdc              # 复杂度、命名、安全等
│   └── TypeScript规范.mdc           # TS 类型规范
├── modules/                         # 模块层（Apply to Specific Files）
│   └── React组件规范.mdc             # React/JSX 规范
├── workflow/                        # 工作流层（Apply Intelligently）
│   └── CRUD开发流程.mdc              # CRUD 完整流程
└── tools/                           # 工具层（Apply Manually）
    └── （待添加）
```

## 🎯 规则分类

### 1. Always Apply（始终应用）

**标识**: `alwaysApply: true`

**位置**: `ai-协作规则.mdc` 和 `basic/` 目录

**作用**: 每次对话自动加载，确保 AI 始终遵循项目的基础规范

**包含**:
- AI 协作总协议
- 代码质量规范
- TypeScript 规范

### 2. Apply to Specific Files（文件匹配）

**标识**: `globs: ["pattern"]`

**位置**: `modules/` 目录

**作用**: 当处理特定类型文件时自动应用

**示例**:
- `globs: ["**/*.tsx", "**/*.jsx"]` - 处理 React 组件时自动应用

### 3. Apply Intelligently（智能匹配）

**标识**: `description: "场景描述"`

**位置**: `workflow/` 目录

**作用**: AI 根据描述智能判断是否需要应用

**示例**:
- 当你说"创建用户管理功能"时，AI 会自动匹配 CRUD 开发流程

### 4. Apply Manually（手动应用）

**标识**: 无特殊字段

**位置**: `tools/` 目录

**作用**: 通过 `@规则名` 手动引用

**示例**:
- `@Git提交规范` - 手动引用 Git 提交相关规范

## 🚀 快速开始

### 1. 验证规则已加载

在 Cursor 中问：

```
列出当前项目的代码规范要求
```

AI 应该能够回答基础规范、TypeScript 规范等内容。

### 2. 测试自动匹配

创建一个 React 组件：

```
创建一个用户卡片组件 UserCard
```

AI 应该自动应用 React 组件规范，生成符合规范的代码。

### 3. 测试智能匹配

开发 CRUD 功能：

```
创建一个完整的用户管理功能，包括增删改查
```

AI 应该自动识别这是 CRUD 场景，并按照 `CRUD开发流程.mdc` 中的步骤执行。

## 📝 使用场景示例

### 场景 1: 创建新的 API 服务

```
帮我创建一个商品（Product）的 API 服务，包括 getList, getDetail, create, update, delete 方法
```

**AI 会**:
1. 应用 TypeScript 规范 ✅
2. 应用代码质量规范 ✅
3. 生成带完整类型定义的 API 服务
4. 包含错误处理和 JSDoc 注释

### 场景 2: 创建 React 组件

```
创建一个数据表格组件 DataTable，支持排序和分页
```

**AI 会**:
1. 匹配 React 组件规范 ✅（globs 匹配 .tsx）
2. 应用 TypeScript 规范 ✅
3. 生成符合 Props 定义、Hooks 使用规范的组件
4. 使用 CSS Modules

### 场景 3: 完整功能开发

```
开发一个订单管理功能，需要列表、搜索、筛选、创建、编辑、删除
```

**AI 会**:
1. 识别为 CRUD 场景 ✅（智能匹配）
2. 按照 types → services → hooks → components 的顺序开发
3. 生成完整的增删改查代码
4. 包含加载状态、错误处理、分页等

## 🔧 自定义规则

### 添加新的基础规范

1. 在 `basic/` 目录创建新的 `.mdc` 文件
2. 添加元数据：
   ```markdown
   ---
   alwaysApply: true
   ---
   
   # 你的规范标题
   ...
   ```

### 添加新的工作流

1. 在 `workflow/` 目录创建新的 `.mdc` 文件
2. 添加元数据：
   ```markdown
   ---
   description: "清晰描述这个工作流的适用场景"
   ---
   
   # 工作流标题
   ...
   ```

### 添加文件类型匹配规则

1. 在 `modules/` 目录创建新的 `.mdc` 文件
2. 添加元数据：
   ```markdown
   ---
   globs: ["**/*.your-ext"]
   ---
   
   # 规范标题
   ...
   ```

## ✅ 规则编写最佳实践

### 1. 使用代码示例

```markdown
## 强制行为

✅ 正确：
\`\`\`typescript
const user: User = { ... };
\`\`\`

❌ 错误：
\`\`\`typescript
const user: any = { ... };
\`\`\`
```

### 2. 明确分类

```markdown
## 强制行为
- 必须执行的操作

## 禁止行为
- 严格禁止的操作

## 最佳实践
- 推荐的做法
```

### 3. 提供完整示例

不要只描述规范，要给出完整的可运行代码示例。

### 4. 保持简洁

每个规则文件专注于一个主题，避免内容过于庞杂。

## 🎓 学习路径

如果你是新手，建议按以下顺序学习：

1. **阅读** `ai-协作规则.mdc` - 理解整体架构
2. **阅读** `basic/代码质量规范.mdc` - 了解基础规范
3. **阅读** `basic/TypeScript规范.mdc` - 掌握类型规范
4. **实践** `workflow/CRUD开发流程.mdc` - 完整功能开发
5. **扩展** 根据项目需要添加自定义规则

## 🐛 故障排除

### AI 没有应用规则

**可能原因**:
1. 规则文件格式错误（检查 `---` 元数据）
2. 描述不够清晰（智能匹配时）
3. globs 模式不匹配（文件匹配时）

**解决方法**:
- 检查 mdc 文件的语法
- 使描述更具体
- 手动 @ 引用规则

### AI 应用了错误的规则

**可能原因**:
- 多个规则的 description 相似
- 规则内容有冲突

**解决方法**:
- 明确规则的适用场景
- 在提示词中明确说明需求

### 规则太多导致上下文溢出

**解决方法**:
- 减少 `alwaysApply` 规则的数量
- 将通用规范合并到一个文件
- 使用智能匹配而非始终应用

## 📚 进阶资源

- [Cursor 官方文档](https://cursor.sh/docs)
- [TypeScript 最佳实践](https://www.typescriptlang.org/docs/handbook/declaration-files/do-s-and-don-ts.html)
- [React 最佳实践](https://react.dev/learn/thinking-in-react)

## 🤝 贡献

如果你发现规则有改进空间，或者想添加新的规则：

1. 测试你的规则
2. 确保遵循现有的格式规范
3. 提交 PR 或反馈

## 📄 许可

本规则集基于实际项目经验总结，可自由使用和修改。

---

**记住**: 规则是为了提高效率，不是限制创造力。在遵循规范的基础上，保持灵活性！ 🚀
