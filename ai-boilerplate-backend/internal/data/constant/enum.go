package constant

// go install github.com/abice/go-enum@latest

//go:generate go-enum --marshal --names --values --ptr --nocomments --flag --output-suffix .gen

// Status 状态
/*ENUM(
disable=-1 // 禁用
enable=1 // 启用
)*/
type Status int32

// SysStatus 系统状态
/*ENUM(
disable=-1 // 禁用
enable=1 // 启用
)*/
type SysStatus int32

// SysRoleDataPermissionType
/*ENUM(
all // 全部数据
deptAndBelow // 本部门及以下数据
dept // 本部门数据
self // 仅本人数据
)*/
type SysRoleDataPermissionType string

// SysMenuType 菜单类型
/*ENUM(
dir // 目录
menu // 菜单
button // 按钮
)*/
type SysMenuType string

// SysDeptType 部门类型 公司:company 区域:area 门店:store
/*ENUM(
root // 根节点
child // 子节点
leaf // 叶子节点
)*/
type SysDeptType string

// SmsChannelCode 短信渠道编码
/*ENUM(
ALIYUN // 阿里云
TENCENT // 腾讯云
HUAWEI // 华为云
QINIU // 七牛云
YUNPIAN // 云片
)*/
type SmsChannelCode string

// DeviceStatus 设备状态
/*ENUM(
disable=-1 // 禁用
enable=1 // 启用
)*/
type DeviceStatus int32

// UserBindDeviceIdentity 用户绑定设备身份
/*ENUM(
admin // 管理员
subAdmin // 子管理员
)*/
type UserBindDeviceIdentity string

// FileStorage 存储引擎
/*ENUM(
volcengine // 火山云
tencent // 腾讯云
aliyun // 阿里云
qiniu // 七牛云
)*/
type FileStorage string

// MembershipType 会员类型
/*
ENUM(
normal // 普通会员
vip // 会员
svip // 超级会员
)
*/
type MembershipType string

// MembershipBenefitKey 会员权益配置表-权益标识
/*
ENUM(
location // 定位
screen // 截图
knock // 敲一敲
)
*/
type MembershipBenefitKey string

/*
ENUM(
refunded=-2 // 已退款
disable=-1 // 禁用
stock=0 // 库存
sold=1 // 已售出
activated=2 // 已激活
expired=3 // 已过期
)
*/
type ActivationCodeStatus int32

/*
ENUM(
membership // 会员
service // 服务
goods // 商品
)
*/
type MallProductType string
