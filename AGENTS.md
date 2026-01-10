<!-- OPENSPEC:START -->
# OpenSpec Instructions

These instructions are for AI assistants working in this project.

Always open `@/openspec/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/openspec/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

Keep this managed block so 'openspec update' can refresh the instructions.

<!-- OPENSPEC:END -->



<skills_system priority="1">

## Available Skills

<!-- SKILLS_TABLE_START -->
<usage>
When users ask you to perform tasks, check if any of the available skills below can help complete the task more effectively. Skills provide specialized capabilities and domain knowledge.

How to use skills:
- Invoke: Bash("openskills read <skill-name>")
- The skill content will load with detailed instructions on how to complete the task
- Base directory provided in output for resolving bundled resources (references/, scripts/, assets/)

Usage notes:
- Only use skills listed in <available_skills> below
- Do not invoke a skill that is already loaded in your context
- Each skill invocation is stateless
</usage>

<available_skills>

<skill>
<name>admin-codeing</name>
<description>Admin 管理后台开发技能。当用户需要开发前端页面、对接后端接口、实现 CRUD 功能时使用此技能。触发场景包括：(1) 新增管理页面 (2) 表单和表格开发 (3) API 接口对接 (4) 权限控制实现 (5) 组件封装 (6) 完整的前端功能开发流程</description>
<location>project</location>
</skill>

<skill>
<name>api-schema-test</name>
<description>API Schema 契约测试技能。基于 schemathesis 进行 API 自动化测试。触发场景包括：(1) 测试单个 API 接口 (2) 批量测试所有 API (3) 测试特定 HTTP 方法 (4) 验证 API 契约合规性 (5) 集成测试和回归测试</description>
<location>project</location>
</skill>

<skill>
<name>app-codeing</name>
<description>"UniApp mobile application development workflow for app project. Use when users need to develop mobile app features, pages, or components. Triggers on: (1) Creating new pages (list/detail/form/tab pages) (2) Implementing business logic (3) API integration (4) Multi-platform adaptation (H5/WeChat/App) (5) Performance optimization (6) Form development (7) List with pagination (8) Complex UI interactions (9) State management with Pinia (10) Any mobile app development tasks using Vue 3 + TypeScript + UniApp stack"</description>
<location>project</location>
</skill>

<skill>
<name>backend-api-gen</name>
<description>后端 API 代码生成技能。用于从 Proto 文件生成 Go 代码（pb.go、grpc.pb.go、http.pb.go 等）。触发场景：(1) Proto 文件修改后 (2) 需要重新生成 API 代码</description>
<location>project</location>
</skill>

<skill>
<name>backend-audit</name>
<description>后端开发前置审计技能。用于验证开发前置条件、审计现有工件状态、确定开发起点。触发场景：(1) 开始后端开发任务前 (2) 检查表/API 工件是否存在 (3) 确定从哪个步骤开始开发</description>
<location>project</location>
</skill>

<skill>
<name>backend-codeing</name>
<description>Backend development skill for this repo. Use when implementing backend features, generating CRUD code, or writing service/data business logic (including after table schema changes).</description>
<location>project</location>
</skill>

<skill>
<name>backend-database</name>
<description>PostgreSQL 数据库表设计技能。触发场景：(1) 创建新表 (2) 修改现有表 (3) 设计表关系 (4) 查询表结构</description>
<location>project</location>
</skill>

<skill>
<name>backend-gorm</name>
<description>后端 GORM 代码生成技能。用于验证数据库表存在性并生成 GORM 模型、DAO、Repo 代码。触发场景：(1) 新建表后生成 GORM 代码 (2) 表结构变更后重新生成 (3) 检查 GORM 工件状态</description>
<location>project</location>
</skill>

<skill>
<name>backend-proto-edit</name>
<description>后端 Protobuf API 编辑技能。用于修改生成的 Proto 文件，添加过滤器、验证规则或业务 RPC。触发场景：(1) 需要添加列表过滤条件 (2) 调整验证规则 (3) 添加/删除 RPC 方法</description>
<location>project</location>
</skill>

<skill>
<name>backend-proto-gen</name>
<description>后端 Protobuf 定义生成技能。用于从 SQL 表结构自动生成 Proto 文件。触发场景：(1) 新表需要生成 API 定义 (2) 表结构变更后重新生成 Proto</description>
<location>project</location>
</skill>

<skill>
<name>backend-quality</name>
<description>后端代码质量检查技能。用于执行依赖注入、代码格式化、Lint 检查和验证。触发场景：(1) 业务逻辑实现后 (2) 代码提交前的质量检查</description>
<location>project</location>
</skill>

<skill>
<name>git-workflow</name>
<description>Git workflow and version control best practices for Kratos Admin project. Use when users need help with git operations, branching, commits, merges, pull requests, conflict resolution, or version control workflows. Triggers on keywords like git, branch, commit, merge, rebase, pull request, PR, push, checkout, conflict, or when users ask about version control operations and git best practices.</description>
<location>project</location>
</skill>

<skill>
<name>interview</name>
<description>This skill conducts discovery conversations to understand user intent and agree on approach before taking action. It should be used when the user explicitly calls /interview, asks for recommendations, needs exploration, wants to clarify, or when the request could be misunderstood. Prevents building the wrong thing by uncovering WHY behind WHAT.</description>
<location>project</location>
</skill>

<skill>
<name>skill-creator</name>
<description>Guide for creating effective skills. This skill should be used when users want to create a new skill (or update an existing skill) that extends Claude's capabilities with specialized knowledge, workflows, or tool integrations.</description>
<location>project</location>
</skill>

<skill>
<name>tech-decision</name>
<description>技术选型决策技能。当用户需要进行技术选型、技术对比、架构决策时使用此技能。触发场景包括：(1) 选择技术框架/库 (2) 对比技术方案 (3) 架构设计决策 (4) 第三方服务选型 (5) 技术栈升级评估</description>
<location>project</location>
</skill>

<skill>
<name>ui-ux-pro-max</name>
<description>"UI/UX design intelligence. 50 styles, 21 palettes, 50 font pairings, 20 charts, 8 stacks (React, Next.js, Vue, Svelte, SwiftUI, React Native, Flutter, Tailwind). Actions: plan, build, create, design, implement, review, fix, improve, optimize, enhance, refactor, check UI/UX code. Projects: website, landing page, dashboard, admin panel, e-commerce, SaaS, portfolio, blog, mobile app, .html, .tsx, .vue, .svelte. Elements: button, modal, navbar, sidebar, card, table, form, chart. Styles: glassmorphism, claymorphism, minimalism, brutalism, neumorphism, bento grid, dark mode, responsive, skeuomorphism, flat design. Topics: color palette, accessibility, animation, layout, typography, font pairing, spacing, hover, shadow, gradient."</description>
<location>project</location>
</skill>

</available_skills>
<!-- SKILLS_TABLE_END -->

</skills_system>
