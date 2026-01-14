# 错误码编写参考

> 何时阅读: 当需要定义业务错误码时阅读此文件。

## 错误码文件位置

- Admin API: `api/admin/v1/error_reason.proto`
- App API: `api/app/v1/error_reason.proto`

## 错误码定义格式

```protobuf
// 错误码注释（必须）
ErrorCodeName = 序号 [
  (errors.code) = HTTP状态码,
  (errors.message) = "ErrorCodeName",
  (errors.lang) = "zh_CN",
  (errors.i18n) = {
    zh_CN: "中文错误信息"
    en_US: "English error message"
  }
];
```

## 完整示例

```protobuf
// 订单不存在
OrderNotFound = 29 [
  (errors.code) = 409,
  (errors.message) = "OrderNotFound",
  (errors.lang) = "zh_CN",
  (errors.i18n) = {
    zh_CN: "订单不存在"
    en_US: "Order not found"
  }
];

// 订单已支付
OrderAlreadyPaid = 30 [
  (errors.code) = 409,
  (errors.message) = "OrderAlreadyPaid",
  (errors.lang) = "zh_CN",
  (errors.i18n) = {
    zh_CN: "订单已支付"
    en_US: "Order already paid"
  }
];

// 支付金额不足（带参数）
PaymentAmountInsufficient = 31 [
  (errors.code) = 409,
  (errors.message) = "PaymentAmountInsufficient",
  (errors.lang) = "zh_CN",
  (errors.i18n) = {
    zh_CN: "支付金额不足，需要支付 %s 元"
    en_US: "Payment amount insufficient, need to pay %s yuan"
  }
];
```

## HTTP 状态码选择

| 场景 | 状态码 | 说明 |
|------|--------|------|
| 业务逻辑错误 | 409 | 数据冲突、状态不对、记录不存在 |
| 认证错误 | 401 | Token 相关错误 |
| 权限错误 | 403 | 无权限访问 |
| 限流错误 | 429 | 请求频率超限 |
| 服务器错误 | 500 | 数据库、Redis、MQ 等内部错误 |

## 生成代码

编辑完成后执行：

```bash
cd ai-boilerplate-backend && make api
```

## Service 中使用

```go
// 返回错误
return nil, pb.ErrorReasonOrderNotFound()

// 带额外信息
return nil, pb.ErrorReasonOrderNotFound(pb.WithError(err))

// 带参数（需要错误信息中有 %s）
return nil, pb.ErrorReasonPaymentAmountInsufficient("100.00")
```

## 序号规则

1. 查看当前文件最大序号
2. 新错误码使用 `最大序号 + 1`
3. 保持序号连续，不要跳号
