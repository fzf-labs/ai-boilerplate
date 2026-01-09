package constant

// SmsChannelCodeList 短信渠道编码列表
var SmsChannelCodeToName = map[string]string{
	SmsChannelCodeALIYUN.String():  "阿里云",
	SmsChannelCodeTENCENT.String(): "腾讯云",
	SmsChannelCodeHUAWEI.String():  "华为云",
	SmsChannelCodeQINIU.String():   "七牛云",
	SmsChannelCodeYUNPIAN.String(): "云片",
}

// FileStorageToName 存储引擎名称
var FileStorageToName = map[string]string{
	FileStorageVolcengine.String(): "火山引擎",
	FileStorageTencent.String():    "腾讯云",
	FileStorageAliyun.String():     "阿里云",
	FileStorageQiniu.String():      "七牛云",
}

// MembershipTypeToName 会员类型名称
var MembershipTypeToName = map[string]string{
	MembershipTypeNormal.String(): "普通会员",
	MembershipTypeVip.String():    "会员",
	MembershipTypeSvip.String():   "超级会员",
}

// MembershipBenefitKeyToName 会员权益配置表-权益标识名称
var MembershipBenefitKeyToName = map[string]string{
	MembershipBenefitKeyLocation.String(): "定位",
	MembershipBenefitKeyScreen.String():   "截图",
	MembershipBenefitKeyKnock.String():    "敲一敲",
}
