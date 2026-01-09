---
name: prd-to-testcase
description: Generate standardized test cases from PRD documents with comprehensive coverage of functional, boundary value, and UI/UX test scenarios. Use when users need to create test cases from product requirements, convert PRD to test documentation, or generate comprehensive test coverage for new features. Triggers when users ask to generate test cases, create test plans from PRD, or need testing documentation from requirements.
---

# PRD to Test Case Generator

## Overview

Generate comprehensive, standardized test case documentation from Product Requirement Documents (PRD). This skill transforms PRD content into structured test cases covering functional testing (positive and negative flows), boundary value testing, and UI/UX testing scenarios, outputting professional markdown documentation with complete traceability.

## Quick Start

**Basic usage:**

When user provides a PRD document:
1. Read the PRD content carefully
2. Identify key functional modules and test scenarios
3. Generate test cases using the standardized template from `assets/testcase-template.md`
4. Ensure comprehensive coverage across all test types

**Example invocation:**
- "Generate test cases from this PRD"
- "Create test documentation for the login feature described in requirements.md"
- "I need comprehensive test cases covering all scenarios in this PRD"

## Workflow

Follow this sequential workflow to generate high-quality test cases:

### Step 1: Analyze PRD Content

Read and understand the PRD document structure:

1. **Identify key sections** in the PRD:
   - Functional requirements → Generate functional test cases
   - Business rules → Generate functional + boundary tests
   - User interface specifications → Generate UI/UX tests
   - Data field definitions → Generate boundary value tests
   - Error handling → Generate exception flow tests
   - Performance/Security requirements → Generate corresponding test types

2. **Extract core information**:
   - Feature modules and functionalities
   - User stories and acceptance criteria
   - Business rules and validation logic
   - Data constraints and boundaries
   - UI interaction flows
   - Error scenarios

3. **Map requirements to test scenarios**:
   - Each user story → At least one positive flow test case
   - Each business rule → Validation test case
   - Each data field → Boundary value tests
   - Each UI element → Interaction test case
   - Each error condition → Exception handling test case

**Tip**: Refer to `references/testcase-guide.md` section "从PRD提取测试场景" for detailed extraction strategies.

### Step 2: Plan Test Case Coverage

Determine comprehensive test coverage based on extracted information:

#### Test Case Types to Generate

1. **Functional Test Cases (正向+异常流程)**
   - **Positive flows**: Normal business scenarios with valid inputs
   - **Negative flows**: Error scenarios with invalid inputs or exceptional conditions
   - Priority: P0-P1 for core features, P2-P3 for edge cases

2. **Boundary Value Test Cases**
   - Test all input constraints: length limits, numeric ranges, format restrictions
   - Apply boundary testing pattern: min-1, min, max, max+1
   - Include special cases: null, empty, special characters

3. **UI/UX Test Cases**
   - Verify interface display and layout
   - Test interaction feedback and states
   - Validate responsive behavior across devices
   - Check accessibility features (when applicable)

#### Coverage Checklist

For each functional requirement, ensure:
- [ ] At least 1 positive flow test case
- [ ] At least 1 negative flow test case
- [ ] Boundary value tests for all inputs
- [ ] UI interaction tests for all interface elements
- [ ] Error message validation tests

**Tip**: Aim for 60%+ test cases to be P0/P1 priority, focusing on critical business flows.

### Step 3: Generate Test Cases

Use the template from `assets/testcase-template.md` to create structured test case documentation:

#### Required Fields for Each Test Case

Complete all fields as specified in the template:

| Field | Description | Example |
|-------|-------------|---------|
| **用例ID** | TC-XXX format, sequential numbering | TC-001, TC-002 |
| **用例名称** | Concise scenario description (≤50 chars) | "使用有效凭据登录成功" |
| **关联需求ID** | Traceable to PRD requirement | REQ-LOGIN-001 |
| **测试类型** | Functional/Boundary/UI/Performance/Security | `功能测试` |
| **优先级** | P0-阻塞 / P1-高 / P2-中 / P3-低 | `P0-阻塞` |
| **负责人** | Test engineer name or team | QA Team |
| **自动化标识** | 可自动化 / 手工测试 / 待评估 | `可自动化` |

#### Writing Test Case Content

For each test case, provide:

1. **测试场景描述**: Detailed description of what this test validates

2. **前置条件**:
   - List all prerequisites
   - Include system state, data setup, environment config
   - Must be specific and verifiable

3. **测试步骤**:
   - Use table format: 步骤 | 操作描述 | 测试数据
   - One action per step
   - Include specific test data

4. **预期结果**:
   - Describe observable, verifiable outcomes
   - Include specific data or states
   - Match each test step or final outcome

5. **测试数据**:
   - Provide complete test data in readable format (JSON, table, code block)
   - Distinguish valid vs invalid data
   - Include boundary cases

6. **备注**: Additional notes, dependencies, known issues

**Example test case structure**:

```markdown
### TC-001: 使用有效凭据登录成功

#### 基本信息

| 字段 | 内容 |
|------|------|
| **用例ID** | TC-001 |
| **用例名称** | 使用有效凭据登录成功 |
| **关联需求ID** | REQ-LOGIN-001 |
| **测试类型** | `功能测试` |
| **优先级** | `P0-阻塞` |
| **负责人** | QA Team |
| **自动化标识** | `可自动化` |

#### 测试场景描述

验证用户使用正确的用户名和密码能够成功登录系统并跳转到首页

#### 前置条件

1. 系统已部署到测试环境 (http://test.example.com)
2. 数据库中存在测试用户: username=testuser, password=Test@123
3. 用户处于未登录状态
4. 浏览器已清除缓存和Cookie

#### 测试步骤

| 步骤 | 操作描述 | 测试数据 |
|------|----------|----------|
| 1 | 打开登录页面 | URL: http://test.example.com/login |
| 2 | 在用户名输入框输入用户名 | testuser |
| 3 | 在密码输入框输入密码 | Test@123 |
| 4 | 点击"登录"按钮 | - |
| 5 | 观察页面跳转和用户状态 | - |

#### 预期结果

1. 页面成功跳转到首页 (URL: http://test.example.com/home)
2. 页面右上角显示用户名 "testuser"
3. 登录按钮变更为"退出"按钮
4. 本地存储中保存了有效的 token

#### 测试数据

```json
{
  "username": "testuser",
  "password": "Test@123"
}
```

#### 备注

此为核心登录流程,必须通过才能进行其他功能测试
```

**Writing Guidelines**:
- Refer to `references/testcase-guide.md` for detailed field descriptions and quality standards
- Keep test case descriptions clear and unambiguous
- Use consistent terminology throughout the document
- Ensure test cases are independent and can be executed in any order

### Step 4: Quality Assurance

Before finalizing test case documentation, perform quality checks:

#### Completeness Check

- [ ] All required fields populated for each test case
- [ ] Test case IDs are sequential and unique
- [ ] All requirements from PRD have corresponding test cases
- [ ] Test data is complete and realistic
- [ ] Traceability established (test case ↔ requirement)

#### Coverage Check

Use the checklist from `references/testcase-guide.md` "编写质量标准" section:

- [ ] Positive flows fully covered
- [ ] Major exception scenarios included
- [ ] Boundary value tests for all constrained inputs
- [ ] UI interactions validated
- [ ] Priority distribution reasonable (60%+ P0/P1)

#### Quality Check

Verify each test case meets quality standards:

- [ ] Test case names are concise and descriptive
- [ ] Steps are specific and actionable
- [ ] Expected results are observable and verifiable
- [ ] Test data is realistic and complete
- [ ] No ambiguous language or technical jargon
- [ ] Test cases are executable without specialized knowledge

#### Update Statistics

Calculate and update the statistics section in the template:

```markdown
## 测试用例统计

| 测试类型 | 数量 | 占比 |
|---------|------|------|
| 功能测试 | 15 | 60% |
| 边界值测试 | 6 | 24% |
| UI/UX测试 | 4 | 16% |
| **总计** | **25** | **100%** |

| 优先级 | 数量 | 占比 |
|--------|------|------|
| P0-阻塞 | 8 | 32% |
| P1-高 | 10 | 40% |
| P2-中 | 5 | 20% |
| P3-低 | 2 | 8% |
| **总计** | **25** | **100%** |
```

### Step 5: Generate Final Document

1. **Use the template structure** from `assets/testcase-template.md`
2. **Fill in project information** at the top
3. **Add all generated test cases** in sequential order
4. **Update statistics section** with accurate counts
5. **Add appendix sections** if needed (environment requirements, risks, dependencies)

**Output file naming convention**:
- Format: `testcase_{feature_name}_{date}.md`
- Example: `testcase_user_login_20260101.md`
- Use lowercase with underscores for clarity

## Key Principles

### Comprehensive Coverage

- Generate test cases for ALL scenarios extracted from PRD
- Don't skip edge cases or exception handling
- Include both positive and negative test flows
- Cover all three test types: Functional, Boundary, UI/UX

### Standardization

- Strictly follow the template structure from `assets/testcase-template.md`
- Use consistent terminology and formatting
- Maintain uniform priority and categorization
- Apply the same level of detail across all test cases

### Traceability

- Each test case must link to specific PRD requirement(s)
- Use clear requirement IDs (REQ-XXX format)
- Ensure bidirectional traceability (requirement ↔ test case)
- Enable coverage analysis and impact assessment

### Quality Focus

- Write clear, unambiguous test steps
- Provide realistic, complete test data
- Ensure test cases are independently executable
- Make expected results observable and verifiable

## Resources

### assets/testcase-template.md

Professional test case documentation template with:
- Complete field structure (all 10 required fields)
- Markdown table formatting for better readability
- Project information header
- Statistics and appendix sections
- Example test case structure

Use this template as the foundation for all generated test documentation.

### references/testcase-guide.md

Comprehensive guide covering:
- **PRD extraction strategies**: How to identify test scenarios from different PRD sections
- **Field descriptions**: Detailed explanation of each test case field with examples
- **Test type distinctions**: Differences between functional, boundary, and UI tests
- **Quality standards**: Completeness, clarity, coverage, and executability checklists
- **Best practices**: Common mistakes to avoid and proven approaches

Refer to this guide when:
- Unsure how to extract test scenarios from PRD
- Need clarification on field values or formats
- Want to improve test case quality
- Looking for examples of good vs bad test cases

## Common Scenarios

### Scenario 1: Complete Feature Testing

User provides PRD for a complete feature (e.g., user authentication system):

1. Analyze entire feature scope from PRD
2. Generate comprehensive test suite covering:
   - All user flows (registration, login, password reset, etc.)
   - Boundary tests for all input fields
   - UI tests for all interface elements
   - Security tests for authentication mechanisms
3. Organize test cases by sub-feature or user flow
4. Ensure high coverage with balanced priority distribution

### Scenario 2: Targeted Scenario Testing

User requests test cases for specific PRD section (e.g., "generate test cases for the password validation rules"):

1. Focus on the specific requirement area
2. Generate targeted test cases:
   - Functional tests for the validation rules
   - Comprehensive boundary tests for password constraints
   - Error message validation tests
3. Keep focused but thorough within the scope
4. Link test cases to specific requirement IDs

### Scenario 3: Test Case Enhancement

User has existing test cases but wants to expand coverage:

1. Review existing test cases to understand current coverage
2. Identify gaps based on PRD requirements
3. Generate additional test cases to fill gaps
4. Maintain consistency with existing test case style and numbering
5. Update statistics to reflect new coverage

## Tips for Effective Test Case Generation

1. **Read PRD thoroughly first**: Don't start generating until you fully understand the requirements
2. **Think like a user**: Consider how real users will interact with the feature
3. **Be paranoid about edge cases**: Think about what could go wrong
4. **Use realistic test data**: Avoid trivial data like "test123"
5. **Keep test cases atomic**: One test case = one scenario
6. **Make steps reproducible**: Anyone should be able to execute your test cases
7. **Write clear expected results**: Avoid vague statements like "system works correctly"
8. **Balance coverage and maintainability**: Comprehensive but not redundant

---

**Remember**: Quality test cases are the foundation of quality software. Take time to ensure comprehensiveness, clarity, and traceability in every test case you generate.
