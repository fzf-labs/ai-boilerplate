package service

import (
	"context"
	"errors"
	"net/http"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/fileutil"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/volcengine/volc-sdk-golang/service/sts"
)

// UploadFileOSSDefaultPolicy 客户端上传-默认上传到 OSS 的方式和凭证获取
func (a *AdminV1FileDatumService) UploadFileOSSDefaultPolicy(ctx context.Context, req *pb.UploadFileOSSDefaultPolicyReq) (*pb.UploadFileOSSDefaultPolicyReply, error) {
	resp := &pb.UploadFileOSSDefaultPolicyReply{
		FileId:     "",
		Storage:    "",
		Volcengine: &pb.VolcenginePolicy{},
		Tencent:    &pb.TencentPolicy{},
		Aliyun:     &pb.AliyunPolicy{},
		Qiniu:      &pb.QiniuPolicy{},
	}
	fileConfig, err := a.fileConfigRepo.FindMasterConfig(ctx)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	url := ""
	ext := ""
	switch fileConfig.Storage {
	case constant.FileStorageVolcengine.String():
		resp.Storage = constant.FileStorageVolcengine.String()
		resp.Volcengine, err = a.volcengineConfig(ctx, fileConfig)
		if err != nil {
			return nil, pb.ErrorReasonStorageGetConfigFailed(pb.WithError(err))
		}
		url = "https://" + resp.Volcengine.Bucket + "." + resp.Volcengine.Endpoint + "/" + req.GetPath()
		ext = fileutil.Ext(req.GetPath())
	case constant.FileStorageTencent.String():
		resp.Storage = constant.FileStorageTencent.String()
		resp.Tencent, err = a.tencentConfig(ctx, fileConfig)
		if err != nil {
			return nil, pb.ErrorReasonStorageGetConfigFailed(pb.WithError(err))
		}
	case constant.FileStorageAliyun.String():
		resp.Storage = constant.FileStorageAliyun.String()
		resp.Aliyun, err = a.aliyunConfig(ctx, fileConfig)
		if err != nil {
			return nil, pb.ErrorReasonStorageGetConfigFailed(pb.WithError(err))
		}
	case constant.FileStorageQiniu.String():
		resp.Storage = constant.FileStorageQiniu.String()
		resp.Qiniu, err = a.qiniuConfig(ctx, fileConfig)
		if err != nil {
			return nil, pb.ErrorReasonStorageGetConfigFailed(pb.WithError(err))
		}
	default:
		return nil, pb.ErrorReasonDataRecordNotFound(pb.WithError(errors.New("storage is not found")))
	}
	// 创建文件记录
	fileDatum := &ai_boilerplate_model.FileDatum{
		ConfigID: fileConfig.ID,
		Name:     req.GetName(),
		Path:     req.GetPath(),
		URL:      url,
		Ext:      ext,
		Size:     req.GetSize(),
		Status:   1,
	}
	err = a.fileDatumRepo.CreateOneCache(ctx, fileDatum)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.FileId = fileDatum.ID
	return resp, nil
}

// 火山云配置
func (a *AdminV1FileDatumService) volcengineConfig(_ context.Context, fileConfig *ai_boilerplate_model.FileConfig) (*pb.VolcenginePolicy, error) {
	resp := &pb.VolcenginePolicy{
		AccessKeyId:     "",
		SecretAccessKey: "",
		SessionToken:    "",
		Endpoint:        "",
		Region:          "",
		Bucket:          "",
	}
	config := pb.StorageConfig{}
	if fileConfig.Config.String() != "" {
		err := jsonutil.Unmarshal(fileConfig.Config, &config)
		if err != nil {
			return nil, err
		}
	}
	instance := sts.NewInstance()
	instance.Client.SetAccessKey(config.Volcengine.AccessKey)
	instance.Client.SetSecretKey(config.Volcengine.SecretKey)
	assumeRole, status, err := instance.AssumeRole(&sts.AssumeRoleRequest{
		DurationSeconds: 300, // 300秒
		RoleTrn:         "trn:iam::" + config.Volcengine.AccountID + ":role/" + config.Volcengine.RoleName,
		RoleSessionName: "tos_role_session",
	})
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New("assume role failed")
	}
	resp.AccessKeyId = assumeRole.Result.Credentials.AccessKeyId
	resp.SecretAccessKey = assumeRole.Result.Credentials.SecretAccessKey
	resp.SessionToken = assumeRole.Result.Credentials.SessionToken
	resp.Endpoint = config.Volcengine.Endpoint
	resp.Region = config.Volcengine.Region
	resp.Bucket = config.Volcengine.Bucket
	return resp, nil
}

// 腾讯云配置
func (a *AdminV1FileDatumService) tencentConfig(_ context.Context, _ *ai_boilerplate_model.FileConfig) (*pb.TencentPolicy, error) {
	resp := &pb.TencentPolicy{}
	return resp, nil
}

// 阿里云配置
func (a *AdminV1FileDatumService) aliyunConfig(_ context.Context, _ *ai_boilerplate_model.FileConfig) (*pb.AliyunPolicy, error) {
	resp := &pb.AliyunPolicy{}
	return resp, nil
}

// 七牛云配置
func (a *AdminV1FileDatumService) qiniuConfig(_ context.Context, _ *ai_boilerplate_model.FileConfig) (*pb.QiniuPolicy, error) {
	resp := &pb.QiniuPolicy{}
	return resp, nil
}
