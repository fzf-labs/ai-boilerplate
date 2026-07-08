# App 认证 Token 模式说明

## 概述

App 端支持两种 token 认证模式：**单 Token 模式**和**双 Token 模式**。两种模式的行为和安全性有所不同。

## 模式对比

| 特性 | 单 Token 模式 | 双 Token 模式 |
|------|-------------|-------------|
| Token 数量 | 1 个 | 2 个（accessToken + refreshToken） |
| AccessToken 有效期 | 长期（如 7 天） | 短期（如 30 分钟） |
| RefreshToken 有效期 | N/A | 长期（如 30 天） |
| 刷新能力 | ❌ 不支持 | ✅ 支持自动刷新 |
| Token 过期处理 | 需要重新登录 | 自动刷新，无感继续请求 |
| 安全性 | 中等 | 高（accessToken 泄露损失有限） |
| 实现复杂度 | 低 | 中等 |

## 配置方式

### 环境变量

在 `ai-boilerplate-uniapp/env/.env` 中配置：

```env
# 单 Token 模式（默认）
VITE_AUTH_MODE = 'single'

# 双 Token 模式
VITE_AUTH_MODE = 'double'
```

### 前端识别

```typescript
import { isDoubleTokenMode } from '@/utils'

if (isDoubleTokenMode) {
  // 双 Token 模式逻辑
} else {
  // 单 Token 模式逻辑
}
```

## 单 Token 模式

### 流程

```
登录 → 获取 token → 使用 token 访问 API
            ↓
      Token 过期 → 需要重新登录
```

### 实现

- **登录**：调用 `POST /app/v1/user/login`，获取单个 token
- **API 请求**：在 `Authorization: Bearer {token}` 头中发送 token
- **Token 过期**：HTTP 返回 401，前端跳转到登录页

### 代码示例

```typescript
const tokenStore = useTokenStore()

// 登录
const res = await tokenStore.login({ username, password })
// tokenStore.tokenInfo = { token: '...', expiresIn: 604800 }

// 使用 token
const token = tokenStore.validToken  // 自动返回有效 token，或空串

// Token 过期
// HTTP 返回 401 → HTTP 拦截器检测 → 跳转登录页
```

## 双 Token 模式

### 流程

```
登录 → 获取 accessToken + refreshToken 
  ↓
使用 accessToken 访问 API
  ↓
AccessToken 过期 → 自动用 refreshToken 刷新 → 获取新 accessToken
  ↓
继续使用新 accessToken 访问 API
  ↓
RefreshToken 也过期 → 需要重新登录
```

### 实现

- **登录**：调用 `POST /app/v1/user/login`，获取 accessToken + refreshToken
- **Token 刷新**：调用 `POST /app/v1/user/refresh_token`，获取新 tokens
- **API 请求**：在 `Authorization: Bearer {accessToken}` 头中发送
- **自动重试**：HTTP 拦截器检测 401 → 自动刷新 → 重试请求

### 代码示例

```typescript
const tokenStore = useTokenStore()

// 登录
const res = await tokenStore.login({ username, password })
// tokenStore.tokenInfo = {
//   accessToken: '...',
//   accessExpiresIn: 1800,
//   refreshToken: '...',
//   refreshExpiresIn: 2592000
// }

// 使用 accessToken
const token = tokenStore.validToken  // 返回有效的 accessToken

// AccessToken 过期
// HTTP 返回 401（尝试用过期 accessToken）
//   ↓
// HTTP 拦截器自动调用 refreshToken()
//   ↓
// 刷新成功：获取新 accessToken → 自动重试之前的请求
// 刷新失败：跳转登录页
```

## Token 失效场景

以下场景会使用户的所有 token 失效：

### 1. 用户修改密码
- 后端：修改密码后清理该用户的所有 token
- 前端：密码修改成功 → token 自动失效 → 下次请求自动重新登录

### 2. 用户注销账号
- 后端：账号注销后清理 token
- 前端：注销后本地 token 清空，跳转登录页

### 3. 管理员禁用用户
- 后端：用户被禁用后清理 token
- 前端：下次请求时返回 401 → 自动跳转登录页

### 4. 用户主动登出
- 前端：调用 `logout()` 方法
- 后端：`POST /app/v1/user/logout` 清理会话
- 本地：删除 token 和用户信息，跳转登录页

## 最佳实践

### 选择合适的模式

- **单 Token 模式**：适合 token 有效期较短（如 1-2 天）的场景，或对无感刷新无要求的应用
- **双 Token 模式**：推荐用于实际生产环境，提供更好的安全性和用户体验

### 实现建议

1. **默认使用双 Token 模式**（更安全、用户体验更好）
2. **合理设置过期时间**：
   - accessToken：15-30 分钟
   - refreshToken：7-30 天
3. **监听登出事件**：某些操作需要即时生效的登出（如改密、注销）
4. **处理网络错误**：刷新 token 时可能网络中断，需要妥善处理

## API 接口

### 登录
```
POST /app/v1/user/login
请求: { username, password }
响应: { token, expiredAt, refreshAt }
```

### 刷新 Token（双 Token 模式）
```
POST /app/v1/user/refresh_token
请求: { refreshToken }
响应: { token, expiredAt, refreshAt }
```

### 登出
```
POST /app/v1/user/logout
请求: {}
响应: {}
```

## 总结

| 需求 | 选择模式 |
|------|--------|
| 简单应用，频繁登录可接受 | 单 Token |
| 需要无感刷新，提升体验 | 双 Token |
| 安全性要求高 | 双 Token |
| 移动 App（流量节省） | 双 Token（减少登录流程） |
