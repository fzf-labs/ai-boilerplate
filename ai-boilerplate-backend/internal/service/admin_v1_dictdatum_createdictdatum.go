package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateDictDatum 字典数据表-创建一条数据
func (a *AdminV1DictDatumService) CreateDictDatum(ctx context.Context, req *pb.CreateDictDatumReq) (*pb.CreateDictDatumReply, error) {
	resp := &pb.CreateDictDatumReply{}
	data := a.dictDatumRepo.NewData()
	data.Type = req.GetType()
	data.Label = req.GetLabel()
	data.Key = req.GetKey()
	data.Value = req.GetValue()
	data.Remark = req.GetRemark()
	data.CSSColor = req.GetCSSColor()
	data.CSSClass = req.GetCSSClass()
	data.Status = req.GetStatus()
	err := a.dictDatumRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
