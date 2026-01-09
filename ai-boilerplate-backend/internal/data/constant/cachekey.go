package constant

import (
	"time"

	"github.com/fzf-labs/godb/cache/keymanage"
)

var cacheKey = keymanage.New("v1")

var (
	LOCK                    = cacheKey.AddKey("LOCK", time.Minute*5, "锁")
	DeviceHeartbeat         = cacheKey.AddKey("deviceheartbeat", time.Minute*2, "设备心跳")
	DeviceHeartbeatSorted   = cacheKey.AddKey("deviceheartbeat_sorted", time.Minute*2, "设备心跳有序集合")
	DeviceControlLocation   = cacheKey.AddKey("devicecontrollocation", time.Minute*2, "设备定位管控")
	DeviceControlScreenshot = cacheKey.AddKey("devicecontrolscreenshot", time.Minute*2, "设备截图管控")

	// 短信验证码相关缓存键
	UserSmsCode           = cacheKey.AddKey("user_sms_code", time.Minute*5, "用户短信验证码")
	UserSmsCodeFrequency  = cacheKey.AddKey("user_sms_code_frequency", time.Hour*24, "用户短信验证码发送频率")
	ActivationCodeBatchNo = cacheKey.AddKey("activation_code_batch_no", time.Hour*24, "激活码批次号")
)
