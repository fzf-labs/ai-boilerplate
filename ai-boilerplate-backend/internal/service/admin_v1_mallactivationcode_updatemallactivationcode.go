package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// UpdateMallActivationCode 激活码管理表-更新一条数据
func (a *AdminV1MallActivationCodeService) UpdateMallActivationCode(ctx context.Context, req *pb.UpdateMallActivationCodeReq) (*pb.UpdateMallActivationCodeReply, error) {
	resp := &pb.UpdateMallActivationCodeReply{}
	data, err := a.mallActivationCodeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mallActivationCodeRepo.DeepCopy(data)
	data.Platform = req.GetPlatform()
	data.PlatformOrderNo = req.GetPlatformOrderNo()
	data.PlatformBuyerID = req.GetPlatformBuyerId()
	data.PlatformBuyerName = req.GetPlatformBuyerName()
	data.Remark = req.GetRemark()
	if req.GetPlatformSoldAt() != "" {
		data.PlatformSoldAt = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetPlatformSoldAt()).StdTime())
	}
	err = a.mallActivationCodeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
