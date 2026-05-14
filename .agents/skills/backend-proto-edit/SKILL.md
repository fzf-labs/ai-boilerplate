---
name: backend-proto-edit
description: 后端 Protobuf API 编辑技能。用于修改生成的 Proto 文件，添加过滤器、验证规则或业务 RPC。触发场景：(1) 删除不需要的 RPC 方法 (2) 需要添加列表过滤条件 (3) 调整验证规则 (4) 添加业务 RPC
---

# Backend Proto Edit (Step 4)

编辑 Proto 定义。

## 可编辑文件

`api/{position}/v1/{table}.proto`

- `{position}`: `admin` 或 `app`
- `{table}`: 表名（如 `user`、`mall_order`）

## 常见编辑

| 场景 | 操作 | 参考 |
|------|------|------|
| 标准 CRUD 足够 | 标记 N/A | - |
| 删除不需要的方法 | 删除 rpc 及 Message | [delete-rpc.md](references/delete-rpc.md) |
| 添加过滤条件 | 修改 `Get{Table}ListReq` | [proto-examples.md#5](references/proto-examples.md) |
| 调整验证 | 修改 `buf.validate` | [proto-examples.md#3-4](references/proto-examples.md) |
| 添加业务 RPC | 添加 rpc 定义 | [proto-examples.md#1](references/proto-examples.md) |

## 输出

```
## Step 4: Proto 编辑
- 状态: N/A / 已修改
- 修改内容: [删除了 Update/Delete 方法] / [添加了过滤条件] / [其他]
- Proto 已锁定: ✅
```
