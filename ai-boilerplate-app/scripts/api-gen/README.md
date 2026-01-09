# API 代码生成工具

基于 `openapi-ts-request` 的 API 接口代码自动生成工具，从 Swagger 文件自动生成类型安全的 API 调用代码。

## 功能特性

- ✅ 自动扫描 backend 项目的 swagger 文件
- ✅ 生成完整的 TypeScript 类型定义
- ✅ 支持多版本 API（v1, v2 等）
- ✅ 自动生成 index.ts 导出文件
- ✅ 自定义函数命名转换
- ✅ 完整的类型推导支持

## 使用方法

### 快速开始

```bash
# 生成所有 API 接口代码
pnpm api:gen
```

### 生成流程

1. 扫描 `../ai-boilerplate-backend/doc/swagger/app` 目录下的所有 `.swagger.json` 文件
2. 按照目录结构在 `src/api` 下生成对应的接口文件
3. 自动生成 `index.ts` 导出文件

### 目录结构

```
src/api/
├── v1/
│   ├── user/          # 用户相关接口
│   │   ├── index.ts   # 导出文件
│   │   ├── types.ts   # 类型定义
│   │   └── wode.ts    # API 函数
│   ├── home/          # 首页相关接口
│   ├── profile/       # 个人资料相关接口
│   └── index.ts       # v1 版本统一导出
└── index.ts           # 所有版本统一导出（如果有多个版本）
```

## 配置说明

### 脚本配置

配置文件位于：`scripts/api-gen/index.ts`

关键配置项：

```typescript
{
  requestLibPath: 'import request from \'@/http\'\nimport type { CustomRequestOptions_ } from \'@/http/types\'',
  requestOptionsType: 'CustomRequestOptions_',
  isGenReactQuery: false,
  reactQueryMode: 'vue',
  isGenJavaScript: false,
  isCamelCase: true,
}
```

### 命名转换规则

1. **文件名转目录名**
   - `user.swagger.json` → `user/`
   - `sys_auth.swagger.json` → `sys-auth/`

2. **函数命名**
   - OperationId: `User_CreateUser` → 函数名: `createUser`
   - OperationId: `DeleteAccount` → 函数名: `deleteAccount`

3. **模块导出名**
   - 目录名: `error-reason` → 导出名: `ErrorReason`
   - 目录名: `sys-auth` → 导出名: `SysAuth`

## 使用示例

### 导入 API

```typescript
// 方式 1: 从版本目录导入
import { User, Home, Profile } from '@/api/v1'

// 方式 2: 直接导入模块
import * as UserApi from '@/api/v1/user'
import * as HomeApi from '@/api/v1/home'

// 方式 3: 导入具体函数
import { deleteAccount, sendVerifyCode } from '@/api/v1/user'
```

### 调用 API

```typescript
import { deleteAccount } from '@/api/v1/user'
import type { DeleteAccountReq } from '@/api/v1/user'

// 调用接口
const response = await deleteAccount({
  body: {
    // 请求参数
  },
  options: {
    // 自定义请求选项
  }
})
```

### 类型推导

生成的代码支持完整的 TypeScript 类型推导：

```typescript
import type {
  DeleteAccountReq,
  DeleteAccountReply
} from '@/api/v1/user'

// 请求类型自动推导
const req: DeleteAccountReq = {
  // IDE 会自动提示可用字段
}

// 响应类型自动推导
const res = await deleteAccount({ body: req })
// res 的类型为 DeleteAccountReply，IDE 会自动提示可用字段
```

## 开发工作流

### 添加新接口

1. 在 backend 项目中定义新的 Protobuf 接口
2. 运行 `make api` 生成 swagger 文件
3. 在 app 项目中运行 `pnpm api:gen`
4. 新的 API 函数和类型定义会自动生成

### 更新现有接口

1. 在 backend 项目中修改 Protobuf 定义
2. 运行 `make api` 更新 swagger 文件
3. 在 app 项目中运行 `pnpm api:gen`
4. API 代码会自动更新

## 注意事项

1. **不要手动修改生成的文件**：生成的 API 文件会在下次运行 `pnpm api:gen` 时被覆盖
2. **保持 backend 项目在同级目录**：脚本假设 backend 项目在 `../ai-boilerplate-backend`
3. **检查导入路径**：确保 `@/http` 和 `@/http/types` 正确配置
4. **版本管理**：建议在每次生成后提交代码，方便追踪变更

## 相关工具

- [openapi-ts-request](https://github.com/dream2023/openapi-ts-request) - OpenAPI/Swagger 代码生成工具
- TypeScript - 类型系统
- Alova - HTTP 请求库

## 故障排查

### 问题：找不到 swagger 文件

检查 backend 项目路径是否正确：
```bash
ls ../ai-boilerplate-backend/doc/swagger/app
```

### 问题：生成失败

1. 检查 swagger 文件格式是否正确
2. 查看控制台错误信息
3. 确认 `openapi-ts-request` 依赖已安装

### 问题：类型导入错误

确保 `@/http/types` 中定义了 `CustomRequestOptions_` 类型。
