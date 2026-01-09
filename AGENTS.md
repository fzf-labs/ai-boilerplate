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
<name>backend-codeing</name>
<description>后端代码开发技能。当用户需要开发后端功能、生成 CRUD 代码、编写业务逻辑时使用此技能。触发场景包括：(1) 为新表生成完整后端代码 (2) 开发新的 API 接口 (3) 编写业务逻辑代码 (4) 表字段改动后更新代码 (5) 完整的后端功能开发流程</description>
<location>project</location>
</skill>

<skill>
<name>backend-database</name>
<description>PostgreSQL 数据库表设计技能。当用户需要设计新的数据库表、添加字段、创建索引、或询问表结构相关问题时使用此技能。触发场景包括：(1) 创建新表结构 (2) 修改现有表 (3) 设计表关联关系 (4) 创建索引策略 (5) 字段类型选择 (6) 命名规范咨询</description>
<location>project</location>
</skill>

<skill>
<name>bug-detective</name>
<description>Systematic bug detection and troubleshooting for Kratos Admin project. Use when users encounter bugs, errors, or issues requiring debugging in Go/Kratos backend, Vue frontend, PostgreSQL database, or Redis cache. Triggers on keywords like bug, debug, error, exception, issue, troubleshooting, not working, failure, or when users report unexpected behavior. Provides structured debugging workflows for HTTP/gRPC services, database queries, caching, permissions, and common Kratos framework errors.</description>
<location>project</location>
</skill>

<skill>
<name>git-workflow</name>
<description>Git workflow and version control best practices for Kratos Admin project. Use when users need help with git operations, branching, commits, merges, pull requests, conflict resolution, or version control workflows. Triggers on keywords like git, branch, commit, merge, rebase, pull request, PR, push, checkout, conflict, or when users ask about version control operations and git best practices.</description>
<location>project</location>
</skill>

<skill>
<name>interview</name>
<description>|</description>
<location>project</location>
</skill>

<skill>
<name>prd-to-testcase</name>
<description>Generate standardized test cases from PRD documents with comprehensive coverage of functional, boundary value, and UI/UX test scenarios. Use when users need to create test cases from product requirements, convert PRD to test documentation, or generate comprehensive test coverage for new features. Triggers when users ask to generate test cases, create test plans from PRD, or need testing documentation from requirements.</description>
<location>project</location>
</skill>

<skill>
<name>product-requirements</name>
<description>Interactive Product Owner skill for requirements gathering, analysis, and PRD generation. Triggers when users request product requirements, feature specification, PRD creation, or need help understanding and documenting project requirements. Uses quality scoring and iterative dialogue to ensure comprehensive requirements before generating professional PRD documents.</description>
<location>project</location>
</skill>

<skill>
<name>prompt-optimizer</name>
<description>Prompt engineering expert that helps users craft optimized prompts using 57 proven frameworks. Use when users want to optimize prompts, improve AI instructions, create better prompts for specific tasks, or need help selecting the best prompt framework for their use case.</description>
<location>project</location>
</skill>

<skill>
<name>prototype-design</name>
<description>Prototype design tool for generating high-fidelity UI/UX prototypes from PRD documents. Trigger when users need to create interactive prototypes based on product requirements, validate UI/UX designs, or quickly generate prototype pages for requirement confirmation. Supports both mobile and PC applications with multiple design systems (WeChat Work, iOS Native, Material Design, Ant Design Mobile).</description>
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

<skill>
<name>webapp-testing</name>
<description>Toolkit for interacting with and testing local web applications using Playwright. Supports verifying frontend functionality, debugging UI behavior, capturing browser screenshots, and viewing browser logs.</description>
<location>project</location>
</skill>

</available_skills>
<!-- SKILLS_TABLE_END -->

</skills_system>
