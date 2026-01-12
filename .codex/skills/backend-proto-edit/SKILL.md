---
name: backend-proto-edit
description: 后端 Protobuf API 编辑技能。用于修改生成的 Proto 文件，添加过滤器、验证规则或业务 RPC。触发场景：(1) 需要添加列表过滤条件 (2) 调整验证规则 (3) 添加/删除 RPC 方法
allowed-tools: Read, Edit, Glob
---

# Backend Proto Edit (Step 4)

编辑 Proto 定义。

## 可编辑文件

`api/{position}/v1/{table}.proto`

## 常见编辑

| 场景 | 操作 |
|------|------|
| 标准 CRUD 足够 | 标记 N/A |
| 添加过滤条件 | 修改 `Get{Table}ListReq` |
| 调整验证 | 修改 `buf.validate` |
| 添加业务 RPC | 添加 rpc 定义 |

## 示例

详细示例见 [references/proto-examples.md](references/proto-examples.md)

## 输出

```
## Step 4: Proto 编辑
- 状态: N/A / 已修改
- Proto 已锁定: ✅
```
