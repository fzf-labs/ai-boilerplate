# Spec Delta: Settings Center (设置中心)

本规范定义了设置中心的功能需求,包括通知设置、关于页面等系统设置项。

---

## ADDED Requirements

### Requirement: 设置中心首页
系统SHALL在设置中心为用户提供各类应用设置的管理入口,包括通知、隐私、账号安全等。

#### Scenario: 查看设置列表
**Given** 用户进入设置中心首页
**When** 页面加载完成
**Then** 系统应展示:
- 账号设置组:个人信息、账号安全
- 通知设置组:通知偏好
- 隐私设置组:隐私权限
- 通用设置组:语言、清除缓存、关于我们
- 每个设置项包含:图标、标题、右箭头(可跳转项)

#### Scenario: 跳转到设置子页面
**Given** 用户在设置中心首页
**When** 用户点击某个设置项(如"通知偏好")
**Then** 系统应跳转到对应的设置详情页

#### Scenario: 清除缓存
**Given** 用户在设置中心首页
**When** 用户点击"清除缓存"
**Then** 系统应:
- 弹出确认弹窗"确定要清除缓存吗?"
- 用户确认后清除应用缓存
- 展示"缓存已清除"提示
- 更新缓存大小显示为"0KB"

---

### Requirement: 通知设置
系统SHALL允许用户自定义通知偏好,控制接收哪些类型的通知。

#### Scenario: 查看通知设置
**Given** 用户进入通知设置页面
**When** 页面加载完成
**Then** 系统应展示:
- 系统通知开关(总开关)
- 活动通知开关(营销推送)
- 订单通知开关(订单状态更新)
- 消息通知开关(客服消息、评论回复)
- 勿扰模式设置(可选时间段)
- 每个开关显示当前状态(开/关)

#### Scenario: 切换通知开关
**Given** 用户在通知设置页面
**When** 用户切换某个通知开关
**Then** 系统应:
- 立即切换开关状态(UI反馈)
- 调用保存设置API
- 如保存成功,保持新状态
- 如保存失败,回退到原状态并提示"保存失败,请重试"

#### Scenario: 关闭系统通知总开关
**Given** 用户在通知设置页面
**When** 用户关闭"系统通知"总开关
**Then** 系统应:
- 同时禁用所有子通知开关(灰色不可操作)
- 调用系统API关闭推送权限(如支持)
- 展示提示"关闭后将不再接收任何通知"

#### Scenario: 设置勿扰时段
**Given** 用户在通知设置页面
**When** 用户点击"勿扰时段"
**Then** 系统应:
- 展示时间选择器(开始时间和结束时间)
- 用户选择时间后保存设置
- 在该时段内不推送通知
- 展示当前设置的勿扰时段(如"22:00 - 08:00")

---

### Requirement: 关于我们
系统SHALL在关于页面展示应用的版本信息、用户协议、隐私政策等内容。

#### Scenario: 查看关于页面
**Given** 用户进入关于页面
**When** 页面加载完成
**Then** 系统应展示:
- 应用Logo
- 应用名称
- 当前版本号(如"v1.2.3")
- "检查更新"按钮
- 用户协议链接
- 隐私政策链接
- 官方网站链接(可选)
- 客服电话/邮箱(可选)

#### Scenario: 检查更新
**Given** 用户在关于页面
**When** 用户点击"检查更新"
**Then** 系统应:
- 调用版本检查API
- 如有新版本,展示更新提示弹窗,包含:
  - 新版本号
  - 更新内容
  - "立即更新"和"稍后提醒"按钮
- 如无新版本,展示"已是最新版本"提示

#### Scenario: 立即更新(应用商店)
**Given** 检测到新版本
**When** 用户点击"立即更新"
**Then** 系统应:
- 跳转到应用商店下载页面(App)
- 或刷新页面(H5/小程序,如支持热更新)

#### Scenario: 查看用户协议/隐私政策
**Given** 用户在关于页面
**When** 用户点击"用户协议"或"隐私政策"
**Then** 系统应:
- 跳转到协议详情页(WebView或内部页面)
- 展示完整的协议内容
- 支持滚动阅读

---

### Requirement: 语言设置
系统SHALL允许用户切换应用语言(如支持国际化)。此功能为可选实现,如应用支持多语言时启用。

#### Scenario: 切换语言
**Given** 应用支持多语言
**When** 用户在设置中心点击"语言"
**Then** 系统应:
- 展示语言选择列表(简体中文、English等)
- 标识当前使用的语言
- 用户选择后重新加载应用
- 所有文本切换到新语言

---

## API Contracts (接口契约)

### GET /api/v1/settings/notification
获取通知设置

**Request Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "systemNotification": true,
  "activityNotification": true,
  "orderNotification": true,
  "messageNotification": false,
  "doNotDisturb": {
    "enabled": true,
    "startTime": "22:00",
    "endTime": "08:00"
  }
}
```

---

### POST /api/v1/settings/notification
保存通知设置

**Request Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "systemNotification": true,
  "activityNotification": false,
  "orderNotification": true,
  "messageNotification": true,
  "doNotDisturb": {
    "enabled": true,
    "startTime": "22:00",
    "endTime": "08:00"
  }
}
```

**Response 200:**
```json
{
  "message": "设置已保存"
}
```

---

### GET /api/v1/app/version
检查应用版本

**Response 200:**
```json
{
  "currentVersion": "1.2.3",
  "latestVersion": "1.3.0",
  "hasUpdate": true,
  "updateUrl": "https://example.com/download",
  "updateContent": "1. 新增帮助中心\n2. 优化设置中心\n3. 修复已知问题",
  "forceUpdate": false
}
```

**Response 200 (无更新):**
```json
{
  "currentVersion": "1.3.0",
  "latestVersion": "1.3.0",
  "hasUpdate": false
}
```

---

### GET /api/v1/app/about
获取关于信息

**Response 200:**
```json
{
  "appName": "示例商城",
  "logo": "https://cdn.example.com/logo.png",
  "version": "1.2.3",
  "userAgreementUrl": "https://example.com/agreement.html",
  "privacyPolicyUrl": "https://example.com/privacy.html",
  "officialWebsite": "https://example.com",
  "customerService": {
    "phone": "400-123-4567",
    "email": "support@example.com"
  }
}
```

---

### POST /api/v1/app/cache/clear
清除缓存(可选,可纯客户端处理)

**Request Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "缓存已清除"
}
```

---

## UI/UX Requirements

### 设置中心首页
- 设置项分组展示,组之间有分隔线
- 每个设置项左侧图标,右侧箭头(可跳转项)
- 开关类设置直接在列表项右侧展示Switch组件
- 缓存大小实时显示(如"缓存 45.6MB")

### 通知设置页面
- 使用Cell组件+Switch组件
- 总开关关闭时,子开关置灰不可点击
- 勿扰时段点击后弹出时间选择器
- 设置变更后立即保存,无需"保存"按钮

### 关于页面
- Logo居中显示,尺寸120rpx x 120rpx
- 版本号使用灰色小字体
- "检查更新"按钮使用主题色
- 协议链接使用蓝色下划线文字
- 客服信息使用图标+文字展示,支持一键拨号/复制

---

## Error Handling
- 通知设置保存失败:回退到原状态,提示"保存失败,请重试"
- 版本检查失败:提示"检查更新失败,请稍后重试"
- 清除缓存失败:提示"清除失败,请重试"
- 协议链接打开失败:提示"页面加载失败"

---

## Performance Requirements
- 通知设置变更立即保存,响应时间 < 1秒
- 版本检查超时时间5秒
- 清除缓存操作使用异步,避免阻塞UI
- 关于页面数据缓存1天

---

## Security & Privacy
- 清除缓存不清除登录态(token)
- 协议链接使用HTTPS
- 版本检查接口防止中间人攻击
- 通知设置仅同步到云端,不暴露给第三方

---

## MODIFIED Requirements

### Requirement: 个人信息编辑(已存在)
系统SHALL保持现有个人信息编辑功能不变,确保从设置中心可正常跳转。

#### Scenario: 从设置中心进入个人信息编辑
**Given** 用户在设置中心首页
**When** 用户点击"个人信息"
**Then** 系统应跳转到现有的个人信息编辑页面(/pages/profile/edit)

---

### Requirement: 账号安全(已存在)
系统SHALL保持现有账号安全功能不变,确保从设置中心可正常跳转。

#### Scenario: 从设置中心进入账号安全
**Given** 用户在设置中心首页
**When** 用户点击"账号安全"
**Then** 系统应跳转到现有的账号安全页面(/pages/security/index)

---

### Requirement: 隐私设置(已存在)
系统SHALL保持现有隐私设置功能不变,确保从设置中心可正常跳转。

#### Scenario: 从设置中心进入隐私设置
**Given** 用户在设置中心首页
**When** 用户点击"隐私设置"
**Then** 系统应跳转到现有的隐私设置页面(/pages/privacy/index)

---

*Last updated: 2026-01-08*
