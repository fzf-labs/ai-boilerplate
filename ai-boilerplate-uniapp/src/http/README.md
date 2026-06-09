# 请求库

目前 unibest 支持 3 种请求库：

- `简单版本http`，路径 `src/http/http.ts`，适合最小依赖场景。
- `alova`，路径 `src/http/alova.ts`，适合已经熟悉 alova 的项目。
- `vue-query`，路径 `src/http/vue-query.ts`，主要用于自动生成接口，示例在 `src/service/app` 文件夹。

## 如何选择

如果项目足够简单，优先使用 `简单版本http`。
如果团队已经在使用 alova 或 vue-query，可以直接沿用熟悉的方案。

## 说明
如果后续脚手架提供了统一的请求库选择方案，再按项目情况切换即可；当前这三种方式都能直接使用。
