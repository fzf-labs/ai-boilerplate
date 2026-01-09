# Admin 开发示例

本目录包含 Admin 管理后台开发的完整示例，从基础的 CRUD 到进阶功能，帮助你快速掌握前端开发技能。

---

## 📚 示例列表

### [01-basic-crud.md](./01-basic-crud.md) - 基础 CRUD 页面开发

**适用场景**：为后端已有的 CRUD 接口创建标准的前端页面

**包含内容**：
- ✅ API 接口定义（增删改查）
- ✅ 表格列配置（支持树形结构）
- ✅ 表单配置（创建/编辑复用）
- ✅ 列表页面实现
- ✅ 表单组件实现
- ✅ 路由配置
- ✅ 权限控制

**何时使用**：
- 第一次开发 Admin 页面
- 需要标准的增删改查功能
- 熟悉项目架构和开发流程

---

### [02-advanced-features.md](./02-advanced-features.md) - 进阶功能实现

**适用场景**：在基础 CRUD 之上添加更多功能

**包含内容**：
- 🔍 搜索表单（多条件筛选）
- ☑️ 批量操作（批量删除、批量导出）
- 📊 数据导入/导出（Excel）
- 🎯 高级筛选（自定义筛选面板）
- 📈 排序功能（单列/多列排序）
- 🔄 刷新策略（自动刷新、条件刷新）

**何时使用**：
- 基础功能已完成，需要添加搜索
- 需要批量操作功能
- 需要导入导出Excel
- 需要复杂的筛选条件

---

### [03-detail-page.md](./03-detail-page.md) - 详情页实现

**适用场景**：需要展示数据详情

**包含内容**：
- 🎨 风格 1：简单版（Descriptions 组件）
- ✨ 风格 2：华丽版（Card + 自定义样式）
- 🔗 列表页集成
- 📑 Tab 切换
- 👥 关联数据展示
- 📊 数据图表

**何时使用**：
- 需要查看数据详情
- 需要展示关联信息
- 需要丰富的视觉效果

---

## 🚀 快速开始

### 第一次开发建议的学习路径

```
1. 阅读 01-basic-crud.md
   ├─ 了解项目结构
   ├─ 学习标准 CRUD 实现
   └─ 完成第一个页面开发

2. 阅读 02-advanced-features.md
   ├─ 添加搜索表单
   ├─ 实现批量操作
   └─ 了解导入导出

3. 阅读 03-detail-page.md
   ├─ 选择合适的详情页风格
   ├─ 实现详情页
   └─ 添加关联数据展示
```

### 根据需求选择示例

| 需求 | 推荐示例 |
|------|---------|
| 我要开发一个部门管理页面 | [01-basic-crud.md](./01-basic-crud.md) |
| 我要添加搜索功能 | [02-advanced-features.md](./02-advanced-features.md#1-搜索表单) |
| 我要实现批量删除 | [02-advanced-features.md](./02-advanced-features.md#2-批量操作) |
| 我要导出 Excel | [02-advanced-features.md](./02-advanced-features.md#3-数据导入导出) |
| 我要实现详情页 | [03-detail-page.md](./03-detail-page.md) |

---

## 📖 参考文档

除了示例之外，还有以下参考文档可以帮助你：

### [../references/components-guide.md](../references/components-guide.md) - 组件使用指南

包含所有可用组件的详细文档：
- 表单组件（Input、Select、DatePicker 等）
- 表格组件（列配置、分页、树形结构）
- 弹窗组件（Modal 配置和 API）
- 权限控制（指令和函数方式）
- 国际化（翻译使用）
- 文件上传（单文件/多文件）

### [../references/best-practices.md](../references/best-practices.md) - 最佳实践

包含代码规范和优化建议：
- 项目结构规范
- 代码组织规范
- TypeScript 最佳实践
- API 接口规范
- 性能优化
- 状态管理
- 样式规范
- 错误处理

---

## 💡 开发建议

### 1. 循序渐进

不要一次性尝试所有功能，建议按以下顺序开发：

1. **先实现基础 CRUD**（01-basic-crud.md）
2. **测试基础功能是否正常**
3. **根据需求添加进阶功能**（02-advanced-features.md）
4. **最后实现详情页**（03-detail-page.md）

### 2. 复制粘贴 vs 理解原理

- ✅ **推荐**：先复制示例代码，确保功能正常，再逐步理解原理
- ❌ **不推荐**：完全不理解就复制，遇到问题无从下手

### 3. 遇到问题怎么办

1. **检查控制台错误**：大部分问题都会在控制台有错误提示
2. **查看参考文档**：components-guide.md 和 best-practices.md
3. **对比示例代码**：看看自己的代码和示例有什么不同
4. **查看相似页面**：项目中可能已有类似的页面实现

### 4. 代码风格

- 保持与项目现有代码风格一致
- 遵循 best-practices.md 中的规范
- 使用 ESLint 和 Prettier 格式化代码

---

## 🎯 常见场景速查

### 场景 1：我要创建一个用户管理页面

1. 阅读 [01-basic-crud.md](./01-basic-crud.md)
2. 复制示例代码，替换 `Dept` 为 `User`
3. 修改字段和表单项
4. 配置路由和权限
5. 测试功能

### 场景 2：我要在用户管理页面添加搜索

1. 已完成基础 CRUD
2. 阅读 [02-advanced-features.md](./02-advanced-features.md#1-搜索表单)
3. 添加搜索表单配置
4. 更新列表页代码
5. 测试搜索功能

### 场景 3：我要实现用户详情页

1. 已完成基础 CRUD
2. 阅读 [03-detail-page.md](./03-detail-page.md)
3. 选择详情页风格（简单版 or 华丽版）
4. 创建详情组件
5. 在列表页集成详情按钮
6. 测试详情功能

### 场景 4：我要实现树形部门选择

1. 参考 [01-basic-crud.md](./01-basic-crud.md) 中的表单配置
2. 使用 `TreeSelect` 组件
3. 提供部门树数据
4. 配置 `fieldNames`

### 场景 5：我要实现批量删除

1. 阅读 [02-advanced-features.md](./02-advanced-features.md#2-批量操作)
2. 在表格中启用多选
3. 实现批量删除逻辑
4. 添加批量删除按钮
5. 测试功能

---

## ⚠️ 注意事项

### 1. 权限代码

确保在后端配置了对应的权限代码，否则按钮会被隐藏：

```typescript
v-access:code="['system:user:create']"  // 需要在后端配置此权限
```

### 2. API 接口

确保后端接口已经实现，接口路径和参数要匹配：

```typescript
// 前端
export function getUserList(params) {
  return requestClient.get('/admin/v1/sys_user/list', { params });
}

// 后端需要实现对应的接口
GET /admin/v1/sys_user/list
```

### 3. 类型定义

使用 TypeScript 时，确保类型定义完整：

```typescript
export namespace SystemUserApi {
  export interface User {
    id: string;
    username: string;
    // ... 其他字段
  }
}
```

### 4. 路由配置

确保路由路径唯一，避免冲突：

```typescript
{
  path: '/system/user',  // 确保路径唯一
  name: 'SystemUser',    // 确保名称唯一
}
```

---

## 📝 版本历史

- **v1.0.0** (2024-01-01)
  - 初始版本
  - 拆分原 crud-example.md 为三个独立示例
  - 添加详细的使用说明

---

## 🤝 贡献

如果你发现示例中有错误或有改进建议，欢迎提交 Issue 或 Pull Request。

---

## 相关链接

- [Admin 开发技能文档](../SKILL.md) - 技能总览
- [组件使用指南](../references/components-guide.md) - 组件详细文档
- [最佳实践](../references/best-practices.md) - 代码规范和优化
