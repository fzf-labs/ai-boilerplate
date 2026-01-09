package rpc

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"time"

	"github.com/fzf-labs/goutil/httputil"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/go-kratos/kratos/v2/log"
)

func NewBaiduPushHTTPRPC(
	cfg *conf.Bootstrap,
	logger log.Logger,
	client *httputil.Client,
) *BaiduPushHTTPRPC {
	l := log.NewHelper(log.With(logger, "module", "data/baiduPushHttpRpc"))
	return &BaiduPushHTTPRPC{
		cfg:    cfg,
		log:    l,
		client: client,
		appKey: cfg.GetBusiness()["baiduPush"].GetFields()["apiKey"].GetStringValue(),
		secret: cfg.GetBusiness()["baiduPush"].GetFields()["secretKey"].GetStringValue(),
	}
}

type BaiduPushHTTPRPC struct {
	cfg    *conf.Bootstrap
	log    *log.Helper
	client *httputil.Client
	appKey string
	secret string
}

type BaiduSingleDevicePushReq struct {
	ChannelID  string `json:"channel_id"` // 必须为端上初始化channel成功之后返回的channel_id
	MsgType    string `json:"msg_type"`   // 取值如下：0：消息；1：通知。默认为0
	Msg        string `json:"msg"`
	MsgExpires string `json:"msg_expires,omitempty"` // 消息过期时间，0~604800(86400*7)，默认为5小时(18000秒)
}

// 通知消息
type NotifyMessage struct {
	Title                  string `json:"title"`                              // 必选 通知标题
	Description            string `json:"description"`                        // 必选 通知文本内容
	NotificationBuilderID  int64  `json:"notification_builder_id,omitempty"`  // 可选
	NotificationBasicStyle int64  `json:"notification_basic_style,omitempty"` // 可选
	OpenType               int64  `json:"open_type,omitempty"`                // 可选
	URL                    string `json:"url,omitempty"`                      // 可选
	PkgContent             string `json:"pkg_content,omitempty"`              // 可选
	CustomContent          string `json:"custom_content,omitempty"`           // 可选
	TargetChannelID        string `json:"target_channel_id,omitempty"`        // 可选
}

type BaiduSingleDevicePushReply struct {
	RequestID      uint32 `json:"request_id"`
	ResponseParams struct {
		MsgID    int64 `json:"msg_id"`
		SendTime int64 `json:"send_time"`
	} `json:"response_params"`
}

// GetBaiduPushQueryParams 获取百度推送查询参数
// 参考百度推送API签名算法：
// 1. 拼接 method + url
// 2. 对参数按key排序并拼接为 key=value 格式
// 3. 最后拼接 secret_key
// 4. 对整个字符串进行URL编码
// 5. 计算MD5哈希
func (h *BaiduPushHTTPRPC) GetBaiduPushQueryParams(method, curl string, params map[string]string) map[string]string {
	if params == nil {
		params = make(map[string]string)
	}
	// 添加必需的参数
	ts := fmt.Sprintf("%d", time.Now().Unix())
	params["apikey"] = h.appKey
	params["timestamp"] = ts
	// 构建签名字符串
	gather := method + curl
	// 对参数按key排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接参数
	for _, key := range keys {
		gather += key + "=" + params[key]
	}
	// 拼接secret_key
	gather += h.secret
	// URL编码并计算MD5
	sign := fmt.Sprintf("%x", md5.Sum([]byte(url.QueryEscape(gather))))
	// 返回查询参数（包含签名）
	result := make(map[string]string)
	for k, v := range params {
		result[k] = v
	}
	result["sign"] = sign
	return result
}

// PushSingleDevice 单设备推送
func (h *BaiduPushHTTPRPC) PushSingleDevice(ctx context.Context, param *BaiduSingleDevicePushReq) (*BaiduSingleDevicePushReply, error) {
	reply := &BaiduSingleDevicePushReply{}
	curl := "https://api.tuisong.baidu.com/rest/3.0/push/single_device"
	// 构建POST请求的参数（包括请求体的参数）
	queryParams := make(map[string]string)
	queryParams["channel_id"] = param.ChannelID
	queryParams["msg"] = param.Msg
	// 获取查询参数
	queryParams = h.GetBaiduPushQueryParams("POST", curl, queryParams)
	response, err := h.client.R().SetContext(ctx).SetHeader("Content-Type", "application/json").SetQueryParams(queryParams).SetSuccessResult(reply).Post(curl)
	if err != nil {
		return nil, err
	}
	if response.IsSuccessState() {
		return nil, errors.New(response.Err.Error())
	}
	return reply, nil
}
