package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMallActivationCodeInfo 激活码管理表-单条数据查询
func (a *AdminV1MallActivationCodeService) GetMallActivationCodeInfo(ctx context.Context, req *pb.GetMallActivationCodeInfoReq) (*pb.GetMallActivationCodeInfoReply, error) {
	resp := &pb.GetMallActivationCodeInfoReply{}
	data, err := a.mallActivationCodeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	userChange := &pb.UserChange{}
	if data.UserChange.String() != "" {
		err = jsonutil.Unmarshal(data.UserChange, userChange)
		if err != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
		}
	}
	resp.Info = &pb.MallActivationCodeInfo{
		Id:                data.ID,
		BatchNo:           data.BatchNo,
		Code:              data.Code,
		ProductType:       data.ProductType,
		ProductId:         data.ProductID,
		ValidSt:           timeutil.RFC3339(data.ValidSt),
		ValidEd:           timeutil.RFC3339(data.ValidEd),
		ActivatedAt:       timeutil.RFC3339(data.ActivatedAt.Time),
		UserId:            data.UserID,
		UserChange:        userChange,
		Platform:          data.Platform,
		PlatformSoldAt:    timeutil.RFC3339(data.PlatformSoldAt.Time),
		PlatformOrderNo:   data.PlatformOrderNo,
		PlatformBuyerId:   data.PlatformBuyerID,
		PlatformBuyerName: data.PlatformBuyerName,
		Remark:            data.Remark,
		Status:            data.Status,
		CreatedAt:         data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
