# AI Boilerplate UniApp

UniApp 移动端模板，基于 unibest。

## 本地开发

从本目录执行命令。使用 H5 开发时默认连接
`http://127.0.0.1:8000`；真机或小程序调试时，把
`VITE_SERVER_BASEURL` 改成设备可访问的局域网地址或平台允许的域名。

## 常用命令

```bash
pnpm install
pnpm dev:h5
pnpm dev:app
pnpm dev:mp-weixin
pnpm build:h5
pnpm build:app
pnpm check:type
```

## 目录说明

- `src/pages`：业务页面
- `src/pages-fg`：登录、404 等公共页面
- `src/router`：路由与登录拦截
- `src/api`：由后端 app Swagger 生成的 API 客户端
- `scripts/postupgrade.js`：升级后清理不需要的平台包
- `uno.config.ts`：UnoCSS 配置

## 说明

默认环境文件位于 `env/.env`。首次运行前按目标平台设置：

- `VITE_UNI_APPID`：DCloud/uni-app 应用 ID。
- `VITE_WX_APPID`：微信小程序 AppID，仅在构建微信小程序时需要。
- `VITE_SERVER_BASEURL`：后端 API 地址，默认指向本机后端。

平台地址检查：

- H5 本地开发可以使用 `http://127.0.0.1:8000`。
- App 真机调试需要使用手机可访问的局域网地址。
- 微信小程序需要使用平台允许的 HTTPS 域名，并配置对应 AppID。
- 不要把真实小程序密钥、生产域名或平台账号凭据提交到仓库。

如果不需要某些小程序平台，可以在 `scripts/postupgrade.js` 里直接删掉对应依赖。

当 `../ai-boilerplate-backend/doc/swagger/app` 发生变化时，重新生成移动端
API 客户端：

```bash
pnpm api:gen
pnpm check:type
```

仅页面文案、样式、路由标题或本地环境默认值变化时，不需要重新生成客户端。

## 更多文档

- 根目录操作流程：`../docs/operations.md`
- 技术选型与模板取舍：`../docs/technical-decisions.md`
- 生成 API 客户端流程：`../docs/generated-artifacts.md`
- 环境配置说明：`../docs/environment.md`
- 常见问题排查：`../docs/troubleshooting.md`
- 验证命令矩阵：`../docs/verification.md`
