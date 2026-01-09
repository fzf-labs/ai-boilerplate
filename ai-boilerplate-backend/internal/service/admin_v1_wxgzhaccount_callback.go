package service

import (
	"context"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	kernelModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/samber/lo"
)

// OfficialAccountCallback 公众号回调
func (a *AdminV1WxGzhAccountService) OfficialAccountCallback(hw http.ResponseWriter, r *http.Request) {
	// 获取 appId
	appID := r.URL.Query().Get("appId")
	if appID == "" {
		a.log.WithContext(r.Context()).Errorf("officialAccountCallback appId is empty")
		return
	}
	// 获取账号信息
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(r.Context(), appID)
	if err != nil {
		a.log.WithContext(r.Context()).Errorf("officialAccountCallback err: %v", err)
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		a.log.WithContext(r.Context()).Errorf("officialAccountCallback err: %v", err)
		return
	}
	// 如果是 Get 请求，说明是微信服务器的验证请求
	if r.Method == http.MethodGet {
		// 验证微信服务器
		verifyResp, verifyErr := officialAccount.Server.VerifyURL(r)
		if verifyErr != nil {
			a.log.WithContext(r.Context()).Errorf("officialAccountCallback err: %v", verifyErr)
			return
		}
		defer func() {
			_ = verifyResp.Body.Close()
		}()
		if sendErr := helper.HttpResponseSend(verifyResp, hw); sendErr != nil {
			a.log.WithContext(r.Context()).Errorf("officialAccountCallback helper.HttpResponseSend err: %v", sendErr)
			return
		}
		return
	}
	// 获取请求参数
	notifyResp, notifyErr := officialAccount.Server.Notify(r, func(event contract.EventInterface) interface{} {
		if event.GetMsgType() == kernelModels.CALLBACK_MSG_TYPE_EVENT {
			// 获取 openid
			openID := event.GetFromUserName()
			// 处理用户信息
			processErr := a.OfficialAccountUserProcess(r.Context(), officialAccount, appID, openID)
			if processErr != nil {
				a.log.WithContext(r.Context()).Errorf("officialAccountCallback err: %v", processErr)
			}
			if event.GetEvent() == "subscribe" {
				return messages.NewText(a.GetSubscribeMessage(appID))
			}
		}
		// 回复消息
		if lo.Contains([]string{
			kernelModels.CALLBACK_MSG_TYPE_TEXT,
			kernelModels.CALLBACK_MSG_TYPE_IMAGE,
			kernelModels.CALLBACK_MSG_TYPE_VOICE,
			kernelModels.CALLBACK_MSG_TYPE_VIDEO,
			kernelModels.CALLBACK_MSG_TYPE_LOCATION,
			kernelModels.CALLBACK_MSG_TYPE_LINK,
		}, event.GetMsgType()) {
			return messages.NewText(a.GetReplyMessage(appID, event.GetMsgType(), string(event.GetContent())))
		}
		return kernel.SUCCESS_EMPTY_RESPONSE
	})
	if notifyErr != nil {
		a.log.WithContext(r.Context()).Errorf("officialAccountCallback err: %v", notifyErr)
		return
	}
	defer func() {
		_ = notifyResp.Body.Close()
	}()
	if sendErr := helper.HttpResponseSend(notifyResp, hw); sendErr != nil {
		a.log.WithContext(r.Context()).Errorf("officialAccountCallback helper.HttpResponseSend err: %v", sendErr)
		return
	}
}

// OfficialAccountUserProcess 公众号用户信息处理
func (a *AdminV1WxGzhAccountService) OfficialAccountUserProcess(ctx context.Context, officialAccount *officialAccount.OfficialAccount, appID, openID string) error {
	// 使用 openid 获取用户信息
	userInfo, err := officialAccount.User.Get(ctx, openID, "zh_CN")
	if err != nil {
		return err
	}
	// 保存用户信息到数据库
	userWx, err := a.wxGzhUserRepo.FindOneCacheByAppIDOpenid(ctx, appID, openID)
	if err != nil {
		return err
	}
	if userWx == nil || userWx.ID == "" {
		// 保存用户信息到数据库
		err = a.wxGzhUserRepo.CreateOneCache(ctx, &ai_boilerplate_model.WxGzhUser{
			AppID:           appID,
			Openid:          openID,
			Unionid:         userInfo.UnionID,
			SubscribeStatus: int32(userInfo.Subscribe),
		})
		if err != nil {
			return err
		}
	} else {
		// 如果关注状态或 openid 发生变化，则更新数据库
		if userWx.SubscribeStatus != int32(userInfo.Subscribe) || userWx.Openid != openID {
			oldUserWx := a.wxGzhUserRepo.DeepCopy(userWx)
			userWx.Openid = openID
			userWx.SubscribeStatus = int32(userInfo.Subscribe)
			err = a.wxGzhUserRepo.UpdateOneCacheWithZero(ctx, userWx, oldUserWx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSubscribeMessage 获取关注消息
func (a *AdminV1WxGzhAccountService) GetSubscribeMessage(appID string) string {
	_ = appID
	return ""
}

// GetReplyMessage 获取回复消息
func (a *AdminV1WxGzhAccountService) GetReplyMessage(appID, msgType, str string) string {
	_ = appID
	_ = msgType
	_ = str
	return ""
}
