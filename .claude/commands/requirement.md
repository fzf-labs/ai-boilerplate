# /requirement - 产品需求全流程

## 描述
完整的产品需求工作流，包括需求澄清、需求分析、PRD 生成和原型设计。通过系统化的交互式流程，确保需求质量达标后自动生成专业的产品需求文档和高保真原型。

## 使用方式
```
/requirement
```

执行后会自动启动产品需求工作流，通过交互式对话完成从需求收集到原型交付的全过程。

## 适用场景

### 1. 新功能规划
当需要开发新功能，需要系统化地分析和文档化需求时使用。

### 2. 需求澄清
当产品想法模糊，需要通过结构化对话明确需求细节时使用。

### 3. PRD 文档生成
当需要输出专业的产品需求文档供团队评审和开发时使用。

### 4. 原型设计
当需要基于需求快速生成可交互的 UI 原型进行验证时使用。

### 5. 产品评审准备
当需要准备完整的产品评审材料（需求文档 + 原型）时使用。

## 执行流程

### 阶段 1：需求收集与分析（调用 product-requirements）

#### 步骤 1：项目上下文理解
1. 自动读取项目 README、package.json 等文件
2. 理解技术栈和现有架构
3. 向用户确认需求理解是否正确

#### 步骤 2：质量评分（100 分制）
基于五个维度进行需求质量评估：
- **业务价值与目标**（30 分）：问题陈述、成功指标、ROI 说明
- **功能需求**（25 分）：用户故事、功能描述、边界情况
- **用户体验**（20 分）：用户画像、用户旅程、交互流程
- **技术约束**（15 分）：性能要求、安全合规、集成需求
- **范围与优先级**（10 分）：MVP 定义、分阶段计划、优先级排序

#### 步骤 3：迭代式需求澄清
- 如果质量分 < 90 分，自动提出针对性问题
- 每轮提问 2-3 个问题，避免信息过载
- 展示分数提升进度，激励用户完善需求
- 持续迭代直到达到 90+ 分质量阈值

#### 步骤 4：生成 PRD 文档
- 质量达标后自动生成专业 PRD 文档
- 保存至 `docs/{功能名}-prd.md`
- 包含：执行摘要、问题陈述、成功指标、用户故事、功能需求、技术约束、风险评估等
- 展示最终质量分数和文档路径

**阶段 1 输出**：
- ✅ 质量评分报告（90+ 分）
- ✅ 专业 PRD 文档（`docs/{功能名}-prd.md`）

---

### 阶段 2：原型设计（调用 prototype-design）

#### 步骤 1：分析 PRD 文档
1. 自动读取生成的 PRD 文档
2. 提取关键信息：
   - 用户画像和使用场景
   - 核心功能和用户流程
   - UI/UX 偏好和约束
   - 平台要求（移动端/PC 端）
3. 确定原型范围和优先级

#### 步骤 2：选择设计系统
根据项目特点推荐设计系统：
- **企业微信风格**（WeChat Work）：中国企业用户、B2B 应用
- **iOS 原生风格**（iOS Native）：iOS 应用、西方消费者
- **Material Design**：Android 应用、跨平台 Web
- **Ant Design Mobile**：企业移动应用、表单密集型

向用户说明推荐理由并确认选择。

#### 步骤 3：设计页面结构
- 根据 PRD 中的用户故事映射 UI 组件
- 规划页面层级和导航流程
- 设计常见页面模式：
  - 首页/仪表盘
  - 列表/浏览页
  - 详情页
  - 表单/输入页
  - 设置/个人页

#### 步骤 4：生成原型代码
- 创建完整的 HTML/CSS/JavaScript 原型
- 使用选定的设计系统样式
- 实现核心交互功能：
  - 页面/标签导航
  - 表单验证
  - 模态对话框
  - Toast 通知
  - 加载状态
- 确保响应式设计

#### 步骤 5：测试与优化
- 视觉测试：颜色、字体、间距、对齐
- 交互测试：按钮、导航、表单、模态框
- 响应式测试：多种屏幕尺寸适配
- 内容测试：真实数据、各种状态展示
- 可访问性测试：对比度、键盘操作、语义化

#### 步骤 6：交付原型
- 保存至 `prototypes/{功能名}/` 目录
- 提供使用说明（如何查看、功能演示、交互说明）
- 说明设计决策和假设
- 提供迭代优化选项

**阶段 2 输出**：
- ✅ 可交互的 HTML 原型（`prototypes/{功能名}/index.html`）
- ✅ 使用说明文档（`prototypes/{功能名}/README.md`）
- ✅ 设计系统规范说明

---

## 完整示例

### 示例 1：新功能需求
```
用户: /requirement
助手: Hi! I'm Sarah, your Product Owner. I'll help define clear requirements for your feature.

Let me first understand your project context...
[自动读取项目文件]

用户: 我想做一个用户任务管理功能
助手:
📊 Requirements Quality Score: 45/100

Breakdown:
- Business Value & Goals: 10/30
- Functional Requirements: 15/25
- User Experience: 8/20
- Technical Constraints: 7/15
- Scope & Priorities: 5/10

Let me ask targeted questions to improve clarity...
[使用 AskUserQuestion 工具提问 2-3 个问题]

用户: [回答问题]
助手: Great! That improved Business Value from 10 to 22.
📊 Requirements Quality Score: 67/100
[继续提问...]

[经过 3-4 轮迭代]

助手:
📊 Final Quality Score: 92/100

Excellent! Ready to generate PRD.
[生成 PRD 文档]
✅ PRD saved to docs/user-task-management-prd.md

Now let's create a prototype based on this PRD...
[切换到 prototype-design]
[分析 PRD，选择设计系统，生成原型]
✅ Prototype saved to prototypes/user-task-management/index.html

Review the prototype and let me know if any adjustments are needed.
```

### 示例 2：已有想法但需要细化
```
用户: /requirement
助手: [需求收集流程...]

用户: 我们需要一个数据分析看板，用户可以查看销售数据
助手:
📊 Requirements Quality Score: 58/100

I understand you want a sales analytics dashboard. Let me clarify:
- What specific metrics should the dashboard show?
- Who are the primary users (sales team, managers, executives)?
- What time ranges should be supported (daily, weekly, monthly)?

[迭代澄清...]
[达到 90+ 分]
[生成 PRD]
[生成原型，推荐使用 Ant Design Mobile 风格]
```

## 输出物

### 阶段 1 输出
- **PRD 文档**：`docs/{功能名}-prd.md`
  - 执行摘要
  - 问题陈述
  - 成功指标
  - 用户画像
  - 用户故事
  - 功能需求
  - 技术约束
  - MVP 范围
  - 风险评估

### 阶段 2 输出
- **原型文件**：`prototypes/{功能名}/`
  - `index.html`（可交互原型）
  - `README.md`（使用说明）
  - 可选：`styles.css`、`script.js`（分离的样式和脚本）
- **设计规范**：颜色、字体、组件样式说明

## 质量保证

### PRD 质量阈值
- **最低要求**：90 分（满分 100）
- **优秀水平**：95+ 分
- **不达标处理**：继续迭代提问，直到达标

### 原型质量标准
- ✅ 符合选定的设计系统规范
- ✅ 实现 PRD 中的核心用户流程
- ✅ 多种状态展示（空态、加载、错误、成功）
- ✅ 响应式设计（移动端/PC 端适配）
- ✅ 可访问性基本达标（对比度、键盘操作）

## 相关技能

- **interview** - 创意探索和需求头脑风暴技能（可选前置）
- **product-requirements** - 需求收集、分析和 PRD 生成技能
- **prototype-design** - 基于 PRD 的原型设计技能

## 注意事项

- 每个阶段完成后会显示明确的完成标志（✅）
- 质量评分会透明展示，用户可随时了解进度
- 遇到不确定的地方会主动提问，不做假设
- 原型是验证工具，不是生产代码
- 可随时基于反馈进行迭代优化

## 最佳实践

1. **不要着急**：宁可多问几轮问题，也不要在需求不清晰时强行生成文档
2. **质量优先**：坚持 90+ 分阈值，不降低标准
3. **用户视角**：始终从用户需求出发，而非技术实现
4. **真实数据**：原型使用真实的功能名称和场景数据，不用占位符
5. **迭代友好**：PRD 和原型都支持快速迭代修改
