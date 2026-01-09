package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetDictDatumList 字典数据表-列表数据查询
func (a *AdminV1DictDatumService) GetDictDatumList(ctx context.Context, req *pb.GetDictDatumListReq) (*pb.GetDictDatumListReply, error) {
	resp := &pb.GetDictDatumListReply{
		Total: 0,
		List:  []*pb.DictDatumInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "type",
				Value: req.GetType(),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.dictDatumRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.DictDatumInfo{
				Id:        v.ID,
				Type:      v.Type,
				Label:     v.Label,
				Key:       v.Key,
				Value:     v.Value,
				Remark:    v.Remark,
				CSSColor:  v.CSSColor,
				CSSClass:  v.CSSClass,
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
