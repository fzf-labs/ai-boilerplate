---
name: app-api-gen
description: "App API 代码生成技能。从后端 Swagger 文件生成 TypeScript API 客户端代码。触发场景：(1) 后端 API 更新后需要同步 (2) 新增 API 接口后生成客户端代码 (3) 需要刷新/重新生成 API 类型定义"
allowed-tools: Bash, Read, Glob
---

# App API Gen

从后端 Swagger 文件生成 TypeScript API 客户端代码。

## 前置条件

1. 后端 API 已实现并生成 Swagger 文件
2. Swagger 文件位于 `ai-boilerplate-backend/doc/swagger/app/` 目录

## 命令

```bash
cd ai-boilerplate-uniapp && pnpm api:gen
```

## 工作原理

1. 扫描 `ai-boilerplate-backend/doc/swagger/app/` 下所有 `.swagger.json` 文件
2. 使用 `openapi-ts-request` 生成 TypeScript 代码
3. 输出到 `ai-boilerplate-uniapp/src/api/` 目录

## 生成文件

**输入：**
- `ai-boilerplate-backend/doc/swagger/app/v1/{module}.swagger.json`

**输出：**
- `ai-boilerplate-uniapp/src/api/v1/{module}/index.ts` - API 函数
- `ai-boilerplate-uniapp/src/api/v1/{module}/types.ts` - TypeScript 类型
- `ai-boilerplate-uniapp/src/api/v1/index.ts` - 统一导出

## 使用示例

```typescript
// 导入生成的 API
import { createUser, getUserList } from '@/api/v1/user'
import type { CreateUserReq, UserListReply } from '@/api/v1/user'

// 调用 API
const res = await createUser({ body: { name: 'test' } })
```

## 注意事项

- **禁止手动修改** `src/api/` 下的生成文件，下次生成会被覆盖
- 如果 API 缺失或类型错误，需要后端更新 Swagger 后重新生成
- 函数命名规则：`operationId` 的后半段转小驼峰（如 `User_CreateUser` → `createUser`）

## 输出

```
## App API 生成
- 命令: cd ai-boilerplate-uniapp && pnpm api:gen
- 状态: ✅ 已生成
- 生成模块: [列出生成的模块]
```
