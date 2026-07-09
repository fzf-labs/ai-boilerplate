package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
	"gorm.io/datatypes"
)

const sensitiveMask = security.SensitiveMask

func maskSensitiveValue(value string) string {
	return security.MaskSecret(value)
}

func aiProviderPlatformInfoFromModel(data *ai_boilerplate_model.AiProviderPlatform) *pb.AiProviderPlatformInfo {
	if data == nil {
		return nil
	}
	return &pb.AiProviderPlatformInfo{
		Id:        data.ID,
		Platform:  data.Platform,
		Name:      data.Name,
		APIURL:    data.APIURL,
		APIKey:    maskSensitiveValue(data.APIKey),
		DocURL:    data.DocURL,
		Sort:      data.Sort,
		Status:    data.Status,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
}

func smsChannelInfoFromModel(data *ai_boilerplate_model.SmsChannel) *pb.SmsChannelInfo {
	if data == nil {
		return nil
	}
	return &pb.SmsChannelInfo{
		Id:           data.ID,
		Name:         data.Name,
		Operator:     data.Operator,
		Remark:       data.Remark,
		APIKey:       maskSensitiveValue(data.APIKey),
		APISecret:    maskSensitiveValue(data.APISecret),
		CallbackURL:  data.CallbackURL,
		Status:       int32(data.Status),
		CreatedAt:    timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:    timeutil.RFC3339(data.UpdatedAt),
		OperatorName: constant.SmsChannelCodeToName[data.Operator],
	}
}

func mailAccountInfoFromModel(data *ai_boilerplate_model.MailAccount) *pb.MailAccountInfo {
	if data == nil {
		return nil
	}
	return &pb.MailAccountInfo{
		Id:        data.ID,
		Mail:      data.Mail,
		Username:  data.Username,
		Password:  maskSensitiveValue(data.Password),
		Host:      data.Host,
		Port:      data.Port,
		SslEnable: data.SslEnable,
		Status:    data.Status,
		Remark:    data.Remark,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
}

func wxGzhAccountInfoFromModel(data *ai_boilerplate_model.WxGzhAccount) *pb.WxGzhAccountInfo {
	if data == nil {
		return nil
	}
	return &pb.WxGzhAccountInfo{
		Id:             data.ID,
		Name:           data.Name,
		Account:        data.Account,
		AppId:          data.AppID,
		AppSecret:      maskSensitiveValue(data.AppSecret),
		URL:            data.URL,
		Token:          maskSensitiveValue(data.Token),
		EncodingAesKey: maskSensitiveValue(data.EncodingAesKey),
		QrCodeURL:      data.QrCodeURL,
		Remark:         data.Remark,
		CreatedAt:      timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:      timeutil.RFC3339(data.UpdatedAt),
	}
}

func storageConfigFromJSON(configJSON datatypes.JSON, redact bool) (*pb.StorageConfig, error) {
	config := &pb.StorageConfig{}
	if configJSON.String() == "" {
		return config, nil
	}
	raw := []byte(configJSON)
	var err error
	if redact {
		raw = security.RedactJSONBytes(raw)
	} else {
		raw, err = security.DecryptJSONSecrets(raw)
		if err != nil {
			return nil, err
		}
	}
	if err := jsonutil.Unmarshal(raw, config); err != nil {
		return nil, err
	}
	return config, nil
}
