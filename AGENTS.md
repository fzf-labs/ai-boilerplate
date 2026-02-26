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
<name>admin-api-gen</name>
<description>"Admin API 代码生成技能。从后端 Swagger 文件生成 TypeScript API 客户端代码。触发场景：(1) 后端 API 更新后需要同步 (2) 新增 API 接口后生成客户端代码 (3) 需要刷新/重新生成 API 类型定义"</description>
<location>project</location>
</skill>

<skill>
<name>admin-codeing</name>
<description>Admin 管理后台开发技能。当用户需要开发前端页面、对接后端接口、实现 CRUD 功能时使用此技能。触发场景包括：(1) 新增管理页面 (2) 表单和表格开发 (3) API 接口对接 (4) 权限控制实现 (5) 组件封装 (6) 完整的前端功能开发流程</description>
<location>project</location>
</skill>

<skill>
<name>admin-dev</name>
<description>Admin 管理后台前端开发完整流程编排。自动编排执行：(1) interview 需求澄清 (2) admin-api-gen 生成 API 客户端 (3) admin-codeing 实现页面功能。触发场景：开发 Admin 管理后台页面、表单表格开发、完整的前端 CRUD 功能开发</description>
<location>project</location>
</skill>

<skill>
<name>app-api-gen</name>
<description>"App API 代码生成技能。从后端 Swagger 文件生成 TypeScript API 客户端代码。触发场景：(1) 后端 API 更新后需要同步 (2) 新增 API 接口后生成客户端代码 (3) 需要刷新/重新生成 API 类型定义"</description>
<location>project</location>
</skill>

<skill>
<name>app-codeing</name>
<description>App development skill for uni-app mobile application. Use when developing mobile pages, integrating backend APIs, implementing features with wot-design-uni components. Triggers include：(1) Creating new pages (2) Form and list development (3) API integration (4) State management (5) Component usage (6) Complete mobile app development workflow</description>
<location>project</location>
</skill>

<skill>
<name>app-dev</name>
<description>App 移动端前端开发完整流程编排。自动编排执行：(1) interview 需求澄清 (2) app-api-gen 生成 API 客户端 (3) app-codeing 实现页面功能。触发场景：开发 uni-app 移动端页面、表单列表开发、完整的移动端功能开发</description>
<location>project</location>
</skill>

<skill>
<name>backend-api-gen</name>
<description>后端 API 代码生成技能。用于从 Proto 文件生成 Go 代码（pb.go、grpc.pb.go、http.pb.go 等）。触发场景：(1) Proto 文件修改后 (2) 需要重新生成 API 代码</description>
<location>project</location>
</skill>

<skill>
<name>backend-api-schema-test</name>
<description>API Schema 契约测试技能。基于 schemathesis 进行 API 自动化测试。触发场景包括：(1) 测试单个 API 接口 (2) 批量测试所有 API (3) 测试特定 HTTP 方法 (4) 验证 API 契约合规性 (5) 集成测试和回归测试</description>
<location>project</location>
</skill>

<skill>
<name>backend-audit</name>
<description>后端开发前置审计技能。用于验证开发前置条件、审计现有工件状态、确定开发起点。触发场景：(1) 开始后端开发任务前 (2) 检查表/API 工件是否存在 (3) 确定从哪个步骤开始开发</description>
<location>project</location>
</skill>

<skill>
<name>backend-codeing</name>
<description>后端业务逻辑开发技能。触发场景：(1) 实现 Service 业务逻辑 (2) 定义业务错误码 (3) 注册 HTTP Server (4) 编写 Data 层代码 (5) 表结构变更后实现逻辑 (6) 后端 CRUD 功能开发</description>
<location>project</location>
</skill>

<skill>
<name>backend-database</name>
<description>PostgreSQL 数据库表设计技能。触发场景：(1) 创建新表 (2) 修改现有表 (3) 设计表关系 (4) 查询表结构</description>
<location>project</location>
</skill>

<skill>
<name>backend-dev</name>
<description>后端开发完整流程编排技能。自动编排执行 Step 0-7 的完整后端开发流程。触发场景：(1) 需要完整的后端 CRUD 开发 (2) 从零开始开发新功能 (3) 用户说"帮我开发xxx后端功能" (4) 需要自动化执行多个后端开发步骤</description>
<location>project</location>
</skill>

<skill>
<name>backend-gorm</name>
<description>后端 GORM 代码生成技能。用于验证数据库表存在性并生成 GORM 模型、DAO、Repo 代码。触发场景：(1) 新建表后生成 GORM 代码 (2) 表结构变更后重新生成 (3) 检查 GORM 工件状态</description>
<location>project</location>
</skill>

<skill>
<name>backend-proto-edit</name>
<description>后端 Protobuf API 编辑技能。用于修改生成的 Proto 文件，添加过滤器、验证规则或业务 RPC。触发场景：(1) 删除不需要的 RPC 方法 (2) 需要添加列表过滤条件 (3) 调整验证规则 (4) 添加业务 RPC</description>
<location>project</location>
</skill>

<skill>
<name>backend-proto-gen</name>
<description>后端 Protobuf API 定义生成技能（必选步骤）。基于 sqltopb 从数据库表自动生成 Proto 文件，禁止手动创建。触发场景：(1) 新建表后生成 API (2) 开发后端 CRUD 功能 (3) 需要 Proto/protobuf 文件 (4) 创建 gRPC/HTTP 接口定义 (5) 后端开发流程中的 Step 3</description>
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
