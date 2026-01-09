# /app-dev - App 移动端开发完整工作流

## 描述
移动端开发完整工作流，基于 uni-app + Vue 3 + TypeScript，支持 H5/小程序/App 多平台。根据需求场景自动选择开发模式，完成从需求到交付的全流程开发。

## 使用方式
```
/app-dev <需求描述>
```

## 工作流决策

根据用户需求，自动选择对应流程：

### 1. 标准页面（列表/详情/表单）
- **列表页**：调用 `app-codeing` 技能 → API 接口 → z-paging 分页 → 搜索筛选 → 测试
- **详情页**：调用 `app-codeing` 技能 → API 接口 → 页面展示 → 操作按钮 → 测试
- **表单页**：调用 `app-codeing` 技能 → API 接口 → wot-design-uni 表单 → 验证 → 提交 → 测试

### 2. 复杂业务页面
- 需求不清晰 → 调用 `interview` 探索方案
- 需要技术选型 → 调用 `tech-decision` 对比方案
- 调用 `app-codeing` 技能 → 组件拆分 → 状态管理 → 多平台适配 → 测试

### 3. API 对接
- 调用 `app-codeing` 技能 → 分析接口文档 → TypeScript 类型定义 → alova 封装 → 页面调用 → 错误处理

### 4. 性能优化
- 识别瓶颈 → 调用 `tech-decision`（如需对比方案）→ 调用 `app-codeing` 技能 → 实施优化（懒加载/缓存/分包/虚拟列表）→ 验证

### 5. 多平台适配
- 调用 `app-codeing` 技能 → 平台差异分析 → 条件编译（#ifdef）→ 功能适配 → 各平台测试

## 开发规范

### 目录结构
```
src/pages/{module}/{name}/
├── index.vue           # 页面主文件
├── components/         # 页面组件
└── composables/        # 组合式函数
```

### API 接口
```typescript
// src/api/{module}/{name}/index.ts
import { alova } from '@/http'

export interface GetUserListReq {
  page: number
  page_size: number
}

export interface GetUserListResp {
  list: User[]
  total: number
}

export const getUserList = (params: GetUserListReq) => {
  return alova.Get<GetUserListResp>('/api/v1/user/list', { params })
}
```

### 页面开发
- 使用 `<script setup lang="ts">`
- Composition API
- 避免 any 类型
- 组件命名 PascalCase

### 多平台适配
```vue
<!-- #ifdef H5 -->
<view>H5 专属</view>
<!-- #endif -->

<!-- #ifdef MP-WEIXIN -->
<view>小程序专属</view>
<!-- #endif -->
```

## 技术栈

- **框架**：uni-app + Vue 3 + TypeScript
- **UI**：wot-design-uni + z-paging
- **状态**：Pinia + pinia-plugin-persistedstate
- **请求**：alova（缓存/重试）
- **样式**：UnoCSS + SCSS

## 常用命令

```bash
pnpm dev           # H5 开发
pnpm dev:mp        # 小程序开发
pnpm build         # H5 构建
pnpm build:mp      # 小程序构建
pnpm type-check    # 类型检查
pnpm lint          # 代码检查
```

## 质量检查

- [ ] TypeScript 类型完整，无 any
- [ ] API 错误处理
- [ ] 加载/空状态
- [ ] 表单验证
- [ ] 图片懒加载
- [ ] 长列表分页/虚拟滚动
- [ ] H5/小程序测试通过
- [ ] ESLint/TypeScript 检查通过

## 相关技能

- `app-codeing` - App 移动端开发技能
- `interview` - 方案探索
- `tech-decision` - 技术选型
- `ui-ux-pro-max` - UI/UX 设计
- `bug-detective` - 问题调试
