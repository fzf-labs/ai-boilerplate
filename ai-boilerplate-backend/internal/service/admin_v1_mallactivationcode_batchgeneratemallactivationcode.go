package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

// BatchGenerateMallActivationCode 激活码管理表-批量生成激活码
func (a *AdminV1MallActivationCodeService) BatchGenerateMallActivationCode(ctx context.Context, req *pb.BatchGenerateMallActivationCodeReq) (*pb.BatchGenerateMallActivationCodeReply, error) {
	resp := &pb.BatchGenerateMallActivationCodeReply{
		BatchNo: "",
	}
	batchNo, err := a.mallActivationCodeRepo.GetBatchNo(ctx)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	codes, err := a.mallActivationCodeRepo.GenerateCode(ctx, req.GetNum())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := make([]*ai_boilerplate_model.MallActivationCode, 0)
	for _, code := range codes {
		data = append(data, &ai_boilerplate_model.MallActivationCode{
			ProductType: req.GetProductType(),
			ProductID:   req.GetProductId(),
			BatchNo:     batchNo,
			Code:        code,
			ValidSt:     carbon.Parse(req.GetValidSt()).StdTime(),
			ValidEd:     carbon.Parse(req.GetValidEd()).StdTime(),
			Platform:    req.GetPlatform(),
			Remark:      req.GetRemark(),
			Status:      int32(constant.ActivationCodeStatusStock),
		})
	}
	err = a.mallActivationCodeRepo.CreateBatchCache(ctx, data, 100)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.BatchNo = batchNo
	return resp, nil
}
