# 删除不需要的 RPC 方法

sqltopb 默认生成 5 个标准 CRUD 方法，根据业务需求删除不需要的。

## 常见删除场景

| 方法 | 常见删除场景 |
|------|-------------|
| `Create{Table}` | 有自定义创建逻辑（如订单用 PlaceOrder） |
| `Update{Table}` | 数据不可修改（如订单、日志） |
| `Delete{Table}` | 数据不可删除（如订单、审计记录） |
| `Get{Table}Info` | 不需要单条查询 |
| `Get{Table}List` | 不需要列表查询 |

## 删除步骤

1. 删除 service 中的 rpc 定义
2. 删除对应的 Request Message（如 `Delete{Table}Req`）
3. 删除对应的 Reply Message（如 `Delete{Table}Reply`）

## 示例 - 删除 Update 和 Delete

**删除前：**

```protobuf
service MallOrder {
  rpc CreateMallOrder(CreateMallOrderReq) returns (CreateMallOrderReply) {...}
  rpc UpdateMallOrder(UpdateMallOrderReq) returns (UpdateMallOrderReply) {...}  // 删除
  rpc DeleteMallOrder(DeleteMallOrderReq) returns (DeleteMallOrderReply) {...}  // 删除
  rpc GetMallOrderInfo(GetMallOrderInfoReq) returns (GetMallOrderInfoReply) {...}
  rpc GetMallOrderList(GetMallOrderListReq) returns (GetMallOrderListReply) {...}
}

// 以下 Message 也需要删除
message UpdateMallOrderReq {...}
message UpdateMallOrderReply {...}
message DeleteMallOrderReq {...}
message DeleteMallOrderReply {...}
```

**删除后：**

```protobuf
service MallOrder {
  rpc CreateMallOrder(CreateMallOrderReq) returns (CreateMallOrderReply) {...}
  rpc GetMallOrderInfo(GetMallOrderInfoReq) returns (GetMallOrderInfoReply) {...}
  rpc GetMallOrderList(GetMallOrderListReq) returns (GetMallOrderListReply) {...}
}
```

## 典型业务场景

### 订单表（只保留查询）

```protobuf
service MallOrder {
  // 保留
  rpc GetMallOrderInfo(...) returns (...) {...}
  rpc GetMallOrderList(...) returns (...) {...}
  
  // 删除 Create/Update/Delete，使用自定义业务方法替代
  // rpc PlaceOrder(...) returns (...) {...}  // 自定义下单
}
```

### 日志/审计表（只保留列表查询）

```protobuf
service OperationLog {
  // 只保留列表查询
  rpc GetOperationLogList(...) returns (...) {...}
  
  // 删除 Create/Update/Delete/GetInfo
}
```

### 配置表（只保留单条查询和更新）

```protobuf
service SystemConfig {
  // 保留查询和更新
  rpc GetSystemConfigInfo(...) returns (...) {...}
  rpc UpdateSystemConfig(...) returns (...) {...}
  
  // 删除 Create/Delete/GetList
}
```
