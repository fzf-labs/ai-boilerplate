package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSensitiveWord 敏感词-更新一条数据
func (a *AdminV1SensitiveWordService) UpdateSensitiveWord(ctx context.Context, req *pb.UpdateSensitiveWordReq) (*pb.UpdateSensitiveWordReply, error) {
	resp := &pb.UpdateSensitiveWordReply{}
	// 查询敏感词是否重复
	sensitiveWord, err := a.sensitiveWordRepo.FindOneCacheByWord(ctx, req.GetWord())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if sensitiveWord != nil && sensitiveWord.ID != "" && sensitiveWord.ID != req.GetId() {
		return nil, pb.ErrorReasonDataDuplicateRecord()
	}
	// 查询敏感词是否存在
	data, err := a.sensitiveWordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sensitiveWordRepo.DeepCopy(data)
	data.Word = req.GetWord()
	data.Lab = req.GetLab()
	err = a.sensitiveWordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
