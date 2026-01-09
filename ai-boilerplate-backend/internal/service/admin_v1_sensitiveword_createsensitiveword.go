package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSensitiveWord 敏感词-创建一条数据
func (a *AdminV1SensitiveWordService) CreateSensitiveWord(ctx context.Context, req *pb.CreateSensitiveWordReq) (*pb.CreateSensitiveWordReply, error) {
	resp := &pb.CreateSensitiveWordReply{}
	// 查询敏感词是否存在
	sensitiveWord, err := a.sensitiveWordRepo.FindOneCacheByWord(ctx, req.GetWord())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if sensitiveWord != nil && sensitiveWord.ID != "" {
		return nil, pb.ErrorReasonDataDuplicateRecord()
	}
	data := a.sensitiveWordRepo.NewData()
	data.Word = req.GetWord()
	data.Lab = req.GetLab()
	err = a.sensitiveWordRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
