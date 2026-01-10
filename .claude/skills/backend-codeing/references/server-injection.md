# Server 注入参考

基于 `internal/server/http.go` 风格。

## 检查是否已注入

搜索 `NewHTTPServer` 函数参数和注册代码：
- 参数: `{position}V1{Table}Service *service.{Position}V1{Table}Service`
- 注册: `{position}v1.Register{Table}HTTPServer(srv, {position}V1{Table}Service)`

## 1. 添加函数参数

在 `NewHTTPServer` 函数参数中添加：

```go
// Admin 示例
adminV1UserService *service.AdminV1UserService,

// App 示例
appV1UserService *service.AppV1UserService,
```

**位置**: 按模块分组，Admin 在前，App 在后。

## 2. 添加服务注册

在函数体中添加注册代码：

```go
// Admin v1 服务注册
adminv1.RegisterUserHTTPServer(srv, adminV1UserService)

// App v1 服务注册
appv1.RegisterUserHTTPServer(srv, appV1UserService)
```

**位置**: 在对应的注释块下方添加。

## 3. 自定义路由（可选）

如需 SSE 流式返回或特殊处理：

```go
// 自定义路由
adminRoute := srv.Route("/admin")
adminRoute.POST("/v1/xxx/completions", adminV1XxxService.XxxHandler)
```
