package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// GetFileConfigStorageSelect 文件配置表-获取存储器选择器
func (a *AdminV1FileConfigService) GetFileConfigStorageSelect(_ context.Context, _ *pb.GetFileConfigStorageSelectReq) (*pb.GetFileConfigStorageSelectReply, error) {
	resp := &pb.GetFileConfigStorageSelectReply{}
	storageNames := constant.FileStorageNames()
	for _, storageName := range storageNames {
		resp.List = append(resp.List, &pb.FileConfigStorage{
			Label: constant.FileStorageToName[storageName],
			Value: storageName,
		})
	}
	return resp, nil
}
