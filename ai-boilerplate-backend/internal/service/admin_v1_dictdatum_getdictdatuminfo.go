package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetDictDatumInfo 字典数据表-单条数据查询
func (a *AdminV1DictDatumService) GetDictDatumInfo(ctx context.Context, req *pb.GetDictDatumInfoReq) (*pb.GetDictDatumInfoReply, error) {
	resp := &pb.GetDictDatumInfoReply{}
	data, err := a.dictDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.DictDatumInfo{
		Id:        data.ID,
		Type:      data.Type,
		Label:     data.Label,
		Key:       data.Key,
		Value:     data.Value,
		Remark:    data.Remark,
		CSSColor:  data.CSSColor,
		CSSClass:  data.CSSClass,
		Status:    data.Status,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
