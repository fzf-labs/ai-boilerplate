# Spec Delta: Help Center (帮助中心)

本规范定义了帮助与客服功能需求,包括帮助文档、常见问题和用户反馈。

---

## ADDED Requirements

### Requirement: 帮助中心首页
系统SHALL提供帮助中心首页,用户可通过该页面快速找到常见问题的答案和获取支持。

#### Scenario: 浏览帮助分类
**Given** 用户进入帮助中心首页
**When** 页面加载完成
**Then** 系统应展示:
- 搜索框(占位符:"搜索帮助内容")
- 热门问题列表(最多5个)
- 帮助分类网格(如:账号问题、订单问题、支付问题、配送问题)
- 每个分类包含图标和名称
- 底部联系客服入口

#### Scenario: 点击帮助分类
**Given** 用户在帮助中心首页
**When** 用户点击某个分类(如"订单问题")
**Then** 系统应:
- 跳转到该分类的FAQ列表页
- 展示该分类下的所有问题
- 顶部显示分类名称

#### Scenario: 点击热门问题
**Given** 用户在帮助中心首页
**When** 用户点击热门问题
**Then** 系统应:
- 跳转到问题详情页
- 展示问题标题和详细答案
- 提供"有帮助"和"无帮助"反馈按钮

---

### Requirement: 常见问题(FAQ)
系统SHALL允许用户浏览和搜索常见问题,快速找到解决方案。

#### Scenario: 浏览FAQ列表
**Given** 用户进入某个分类的FAQ页面
**When** 页面加载完成
**Then** 系统应展示:
- 该分类下的所有问题标题列表
- 每个问题支持展开/折叠
- 默认状态为折叠,仅显示标题

#### Scenario: 展开问题查看答案
**Given** 用户在FAQ列表页
**When** 用户点击某个问题标题
**Then** 系统应:
- 展开该问题,显示详细答案
- 答案支持富文本(文字、图片、链接)
- 底部显示"有帮助"和"无帮助"按钮
- 其他已展开的问题保持展开状态(手风琴模式可选)

#### Scenario: 搜索FAQ
**Given** 用户在帮助中心首页或FAQ页面
**When** 用户在搜索框输入关键词并搜索
**Then** 系统应:
- 展示包含关键词的问题列表
- 高亮显示匹配的关键词
- 如无结果展示"未找到相关问题,请联系客服"
- 提供"联系客服"快捷按钮

#### Scenario: 反馈问题是否有帮助
**Given** 用户查看了问题答案
**When** 用户点击"有帮助"或"无帮助"
**Then** 系统应:
- 记录用户反馈
- 展示"感谢您的反馈"提示
- 如果点击"无帮助",引导用户"联系客服"或"提交反馈"

---

### Requirement: 问题反馈
系统SHALL允许用户提交问题反馈,上报遇到的问题和建议。

#### Scenario: 填写反馈表单
**Given** 用户进入问题反馈页面
**When** 页面加载完成
**Then** 系统应展示:
- 问题分类选择器(必填,如:功能异常、建议、其他)
- 问题描述输入框(必填,多行文本,最多500字)
- 图片上传区域(可选,最多3张)
- 联系方式输入框(可选,默认使用账号手机号)
- "提交"按钮

#### Scenario: 上传问题截图
**Given** 用户在反馈表单页面
**When** 用户点击"上传图片"
**Then** 系统应:
- 调用相册/相机选择图片
- 支持最多3张图片
- 每张图片不超过5MB
- 展示缩略图和删除按钮
- 自动压缩图片到合适大小

#### Scenario: 提交反馈
**Given** 用户填写了必填项
**When** 用户点击"提交"
**Then** 系统应:
- 验证表单(分类和描述不能为空)
- 上传图片到服务器
- 调用提交反馈API
- 展示提交成功提示
- 自动返回上一页或帮助中心首页

#### Scenario: 提交失败
**Given** 用户提交反馈时网络异常
**When** API调用失败
**Then** 系统应:
- 展示"提交失败,请重试"提示
- 保留用户已填写的内容
- 提供"重试"按钮

#### Scenario: 查看反馈历史
**Given** 用户曾提交过反馈
**When** 用户进入"我的反馈"页面
**Then** 系统应展示:
- 反馈记录列表(时间倒序)
- 每条记录包含:问题分类、描述摘要、提交时间、处理状态
- 点击可查看详情和回复

---

### Requirement: 在线客服
系统SHALL提供在线客服功能,用户可通过客服获取实时帮助。此功能为可选实现,可在后期版本中添加。

#### Scenario: 打开客服聊天
**Given** 用户在帮助中心或"我的"页面
**When** 用户点击"联系客服"按钮
**Then** 系统应:
- 打开客服聊天窗口(新页面或弹窗)
- 展示客服欢迎语
- 提供快捷问题选项(可选)
- 展示输入框和发送按钮

#### Scenario: 发送消息
**Given** 用户在客服聊天窗口
**When** 用户输入文字并点击发送
**Then** 系统应:
- 展示用户消息气泡(右侧)
- 调用客服API发送消息
- 展示"发送中"状态
- 收到客服回复后展示客服消息气泡(左侧)

#### Scenario: 快捷回复
**Given** 客服聊天窗口显示快捷问题
**When** 用户点击快捷问题
**Then** 系统应:
- 自动发送该问题文本
- 客服返回预设答案或转人工

#### Scenario: 离线客服
**Given** 客服不在线或超出服务时间
**When** 用户打开客服聊天
**Then** 系统应:
- 展示"客服暂时离线,请留言或提交反馈"
- 允许用户留言
- 引导用户查看FAQ或提交反馈

---

## API Contracts (接口契约)

### GET /api/v1/help/categories
获取帮助分类列表

**Response 200:**
```json
{
  "categories": [
    {
      "id": 1,
      "name": "账号问题",
      "icon": "user",
      "order": 1
    },
    {
      "id": 2,
      "name": "订单问题",
      "icon": "order",
      "order": 2
    },
    {
      "id": 3,
      "name": "支付问题",
      "icon": "payment",
      "order": 3
    },
    {
      "id": 4,
      "name": "配送问题",
      "icon": "delivery",
      "order": 4
    }
  ],
  "hotQuestions": [
    {
      "id": 101,
      "categoryId": 2,
      "question": "如何取消订单?",
      "order": 1
    },
    {
      "id": 102,
      "categoryId": 3,
      "question": "支持哪些支付方式?",
      "order": 2
    }
  ]
}
```

---

### GET /api/v1/help/faqs
获取FAQ列表

**Request Parameters:**
```
categoryId?: number (分类ID,可选)
keyword?: string (搜索关键词,可选)
page?: number (页码,默认1)
pageSize?: number (每页数量,默认20)
```

**Response 200:**
```json
{
  "total": 25,
  "list": [
    {
      "id": 101,
      "categoryId": 2,
      "question": "如何取消订单?",
      "answer": "在订单详情页点击"取消订单"按钮即可。<br>注意:已发货的订单无法取消。",
      "helpful": 156,
      "notHelpful": 12,
      "createdAt": "2025-12-01"
    },
    {
      "id": 102,
      "categoryId": 2,
      "question": "订单多久发货?",
      "answer": "一般情况下,订单会在24小时内发货。<br>如遇节假日可能延迟。",
      "helpful": 98,
      "notHelpful": 5,
      "createdAt": "2025-12-01"
    }
  ]
}
```

---

### POST /api/v1/help/feedback/helpful
提交FAQ反馈(是否有帮助)

**Request Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "faqId": 101,
  "helpful": true
}
```

**Response 200:**
```json
{
  "message": "感谢您的反馈"
}
```

---

### POST /api/v1/help/feedback/submit
提交问题反馈

**Request Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "category": "功能异常",
  "description": "支付时页面卡住了",
  "images": [
    "https://cdn.example.com/feedback/img1.jpg",
    "https://cdn.example.com/feedback/img2.jpg"
  ],
  "contact": "13800138000"
}
```

**Response 200:**
```json
{
  "feedbackId": 5001,
  "message": "提交成功,我们会尽快处理"
}
```

---

### GET /api/v1/help/feedback/my
获取我的反馈记录

**Request Headers:**
```
Authorization: Bearer {token}
```

**Request Parameters:**
```
page?: number (页码,默认1)
pageSize?: number (每页数量,默认20)
```

**Response 200:**
```json
{
  "total": 8,
  "list": [
    {
      "id": 5001,
      "category": "功能异常",
      "description": "支付时页面卡住了",
      "images": ["https://cdn.example.com/feedback/img1.jpg"],
      "status": "pending",
      "statusText": "处理中",
      "reply": null,
      "createdAt": "2026-01-08 14:30:00",
      "updatedAt": "2026-01-08 14:30:00"
    },
    {
      "id": 5000,
      "category": "建议",
      "description": "希望增加暗黑模式",
      "images": [],
      "status": "replied",
      "statusText": "已回复",
      "reply": "感谢您的建议,我们会考虑在未来版本中添加",
      "createdAt": "2026-01-05 10:00:00",
      "updatedAt": "2026-01-06 09:00:00"
    }
  ]
}
```

---

### POST /api/v1/help/image/upload
上传反馈图片

**Request Headers:**
```
Authorization: Bearer {token}
Content-Type: multipart/form-data
```

**Request Body:**
```
file: (binary)
```

**Response 200:**
```json
{
  "url": "https://cdn.example.com/feedback/img_20260108143000.jpg"
}
```

---

## UI/UX Requirements

### 帮助中心首页
- 搜索框固定在顶部,占位符文字灰色
- 热门问题使用卡片列表,带右箭头
- 分类使用2x2或3x3网格,每个分类带图标
- "联系客服"按钮浮动在右下角

### FAQ页面
- 问题列表使用手风琴展开模式
- 问题标题使用粗体,答案使用常规字体
- 搜索结果中关键词高亮显示(黄色背景)
- "有帮助"/"无帮助"按钮使用图标+文字

### 反馈表单
- 分类使用下拉选择器或Picker
- 描述输入框支持多行,显示字数统计(X/500)
- 图片上传区域显示虚线边框,带"+"图标
- 提交按钮在填写完必填项后才启用

### 客服聊天(可选)
- 用户消息右对齐,蓝色气泡
- 客服消息左对齐,白色气泡
- 底部输入框固定,带发送按钮
- 快捷问题使用横向滚动按钮组

---

## Error Handling
- FAQ加载失败:展示"加载失败,请重试"和重试按钮
- 搜索无结果:展示"未找到相关问题"和建议联系客服
- 图片上传失败:展示"上传失败,请重新选择"提示
- 反馈提交失败:保留用户输入,提供重试按钮

---

## Performance Requirements
- 帮助分类和热门问题缓存1小时
- FAQ列表支持分页,首屏加载不超过2秒
- 图片上传前自动压缩到1MB以内
- 搜索使用防抖,300ms后才触发请求

---

*Last updated: 2026-01-08*
