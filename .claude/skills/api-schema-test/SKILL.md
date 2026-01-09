---
name: api-schema-test
description: API Schema 契约测试技能。基于 schemathesis 进行 API 自动化测试。触发场景包括：(1) 测试单个 API 接口 (2) 批量测试所有 API (3) 测试特定 HTTP 方法 (4) 验证 API 契约合规性 (5) 集成测试和回归测试
---

# API Schema 契约测试技能

本技能提供 API Schema 契约测试工作流程，基于 schemathesis 自动化测试工具，验证 API 实现与 Swagger 文档的一致性。

## 项目结构

```
ai-boilerplate-backend/
├── doc/swagger/admin/v1/      # Swagger API 文档
│   ├── sys_admin.swagger.json
│   ├── sys_auth.swagger.json
│   └── ...
├── scripts/
│   └── api-schema-test.sh     # 测试脚本
└── Makefile                   # 包含 api-schema-test 目标
```

## 核心工作流程

### 前置条件检查

在开始测试之前，确认：
- [ ] API 服务已启动（默认：http://localhost:8000）
- [ ] schemathesis 已安装（可使用 `--install` 选项自动安装）
- [ ] Swagger 文档已生成（使用 `make api` 生成）
- [ ] 测试账号可用（默认：admin/123456）

### 流程 1：测试单个 API 文件

此流程适用于测试特定模块的 API 接口。

#### 第 1 步：确认测试目标

**询问用户：**
1. 要测试的 Swagger 文件路径（例如：doc/swagger/admin/v1/sys_admin.swagger.json）
2. 测试的 HTTP 方法（可选：GET/POST/PUT/DELETE/ALL，默认：GET）
3. 是否需要详细输出（-v 参数）

**自动执行：**
- 验证 Swagger 文件是否存在
- 检查 API 服务是否可访问

#### 第 2 步：执行测试

执行命令测试指定文件：

```bash
cd ai-boilerplate-backend
make api-schema-test FILE=doc/swagger/admin/v1/sys_admin.swagger.json
```

**可选参数：**
- 指定 HTTP 方法：`make api-schema-test FILE=xxx.json METHOD=POST`
- 详细输出：`./scripts/api-schema-test.sh -v xxx.json`
- 自定义 API URL：`./scripts/api-schema-test.sh --url http://api.example.com xxx.json`

#### 第 3 步：分析测试结果

**成功标志：**
- ✅ 所有测试用例通过
- ✅ API 响应符合 Schema 定义
- ✅ 状态码符合预期

**失败处理：**
- 检查 API 实现是否与 Swagger 文档一致
- 验证数据类型、必填字段、枚举值等
- 查看详细错误信息（使用 -v 参数）

### 流程 2：批量测试所有 API

此流程适用于全面的回归测试和集成测试。

#### 第 1 步：确认测试范围

**询问用户：**
1. 测试的 HTTP 方法（GET/POST/PUT/DELETE/ALL，默认：GET）
2. 是否需要详细输出
3. 是否需要认证（默认：需要）

#### 第 2 步：执行批量测试

执行命令测试所有 Swagger 文件：

```bash
cd ai-boilerplate-backend
make api-schema-test
```

**可选参数：**
- 测试所有方法：`make api-schema-test METHOD=ALL`
- 仅测试 POST：`make api-schema-test METHOD=POST`
- 详细输出：`./scripts/api-schema-test.sh -v`

#### 第 3 步：汇总测试结果

**输出信息：**
- 测试文件总数
- 通过/失败的接口数量
- 每个文件的测试结果

**生成报告：**
- 记录失败的接口和原因
- 提供修复建议
- 标记需要更新的文档

### 流程 3：自定义测试配置

此流程适用于特殊测试场景。

#### 第 1 步：配置测试环境

**环境变量：**
```bash
export TEST_API_URL=http://api.example.com
export TEST_ADMIN_USER=admin
export TEST_ADMIN_PASS=123456
```

#### 第 2 步：执行自定义测试

**脚本参数：**
```bash
cd ai-boilerplate-backend
./scripts/api-schema-test.sh [options] [swagger_file]

Options:
  -u, --url URL       API URL (默认: http://localhost:8000)
  -m, --method M      方法过滤: GET/POST/PUT/DELETE/ALL (默认: GET)
  -a, --all           测试所有 Swagger 文件
  -v, --verbose       详细输出
  --user USER         用户名 (默认: admin)
  --pass PASS         密码 (默认: 123456)
  --no-auth           不使用认证
  --install           安装 schemathesis
  -h, --help          显示帮助
```

**示例：**
```bash
# 测试所有文件的 GET 接口
./scripts/api-schema-test.sh

# 测试指定文件
./scripts/api-schema-test.sh doc/swagger/admin/v1/sys_admin.swagger.json

# 测试 POST 接口
./scripts/api-schema-test.sh -m POST

# 测试所有方法，详细输出
./scripts/api-schema-test.sh -m ALL -v

# 指定 API URL
./scripts/api-schema-test.sh --url http://api.example.com

# 不使用认证
./scripts/api-schema-test.sh --no-auth
```

## 常见问题处理

### 问题 1：schemathesis 未安装

**错误信息：**
```
❌ schemathesis 未安装
```

**解决方案：**
```bash
# 方式 1：使用脚本安装
cd ai-boilerplate-backend
./scripts/api-schema-test.sh --install

# 方式 2：手动安装
pip install schemathesis
# 或
pipx install schemathesis
```

### 问题 2：API 服务未启动

**错误信息：**
```
⚠️  获取 Token 失败，继续测试无认证接口
```

**解决方案：**
```bash
# 启动 API 服务
cd ai-boilerplate-backend
make run
```

### 问题 3：Schema 验证失败

**错误信息：**
```
Schema validation failed: ...
```

**解决方案：**
1. 检查 API 实现是否与 Swagger 文档一致
2. 验证数据类型、必填字段、枚举值等
3. 更新 Swagger 文档或修复 API 实现
4. 重新生成 API 代码：`make api`

### 问题 4：认证失败

**错误信息：**
```
401 Unauthorized
```

**解决方案：**
1. 检查用户名和密码是否正确
2. 验证登录接口是否正常
3. 使用 `--no-auth` 跳过认证测试公开接口

## 最佳实践

### 1. 测试策略

**开发阶段：**
- 每次修改 API 后测试相关接口
- 使用单文件测试快速验证

**集成测试：**
- 测试所有 GET 接口（快速验证）
- 逐步测试 POST/PUT/DELETE 接口

**回归测试：**
- 测试所有方法（METHOD=ALL）
- 使用详细输出记录问题

### 2. CI/CD 集成

**在 CI 流程中添加：**
```yaml
# .github/workflows/test.yml
- name: API Schema Test
  run: |
    cd ai-boilerplate-backend
    make api-schema-test METHOD=GET
```

### 3. 文档同步

**保持文档与代码一致：**
1. 修改 proto 文件后重新生成 Swagger：`make api`
2. 运行 schema 测试验证一致性
3. 同步文档到 Apifox：`make apidoc`

### 4. 测试覆盖

**优先级：**
1. 高频接口（登录、列表、详情）
2. 核心业务接口
3. 边界情况和错误处理

## 技能使用示例

### 示例 1：测试用户管理 API

**用户请求：**
"测试 sys_admin 的 API 接口"

**执行步骤：**
1. 确认 API 服务已启动
2. 执行：`make api-schema-test FILE=doc/swagger/admin/v1/sys_admin.swagger.json`
3. 分析测试结果
4. 报告通过/失败的接口

### 示例 2：全面回归测试

**用户请求：**
"运行所有 API 的契约测试"

**执行步骤：**
1. 确认测试环境
2. 执行：`make api-schema-test METHOD=ALL`
3. 汇总测试结果
4. 生成问题清单

### 示例 3：测试新增接口

**用户请求：**
"测试新增的 POST 接口"

**执行步骤：**
1. 确认 Swagger 文档已更新
2. 执行：`make api-schema-test METHOD=POST`
3. 验证新接口的 Schema 合规性
4. 提供修复建议（如有问题）

## 工具链说明

### schemathesis

**功能：**
- 基于 OpenAPI/Swagger 规范自动生成测试用例
- 验证 API 响应是否符合 Schema 定义
- 支持多种 HTTP 方法和认证方式

**优势：**
- 自动化测试，无需手写测试用例
- 发现 Schema 不一致问题
- 支持 CI/CD 集成

**文档：**
- 官网：https://schemathesis.readthedocs.io/
- GitHub：https://github.com/schemathesis/schemathesis

## 注意事项

1. **测试数据：** 测试会调用真实 API，可能产生测试数据，建议使用测试环境
2. **性能影响：** 批量测试会发送大量请求，注意 API 服务的性能
3. **认证 Token：** Token 有过期时间，长时间测试可能需要重新获取
4. **Schema 严格性：** schemathesis 会严格验证 Schema，确保文档准确性
5. **错误处理：** 测试失败不会中断，会继续测试其他接口

## 相关技能

- **backend-crud：** 生成 CRUD 代码后，使用本技能测试 API
- **backend-database：** 创建数据库表后，生成 API 并测试
- **tech-decision：** 选择 API 测试工具和策略

## 总结

本技能提供完整的 API Schema 契约测试工作流程，帮助开发者：
- 快速验证 API 实现与文档的一致性
- 自动化回归测试
- 提高 API 质量和可靠性
- 保持文档与代码同步

使用本技能可以显著提升 API 开发效率和质量。
