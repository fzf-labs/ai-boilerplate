package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateDictDatum 字典数据表-更新一条数据
func (a *AdminV1DictDatumService) UpdateDictDatum(ctx context.Context, req *pb.UpdateDictDatumReq) (*pb.UpdateDictDatumReply, error) {
	resp := &pb.UpdateDictDatumReply{}
	data, err := a.dictDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.dictDatumRepo.DeepCopy(data)
	data.Type = req.GetType()
	data.Label = req.GetLabel()
	data.Key = req.GetKey()
	data.Value = req.GetValue()
	data.Remark = req.GetRemark()
	data.CSSColor = req.GetCSSColor()
	data.CSSClass = req.GetCSSClass()
	data.Status = req.GetStatus()
	err = a.dictDatumRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
