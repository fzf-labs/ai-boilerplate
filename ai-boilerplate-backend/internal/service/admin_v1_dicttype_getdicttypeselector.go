package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetDictTypeSelector 字典类型表-选择器
func (a *AdminV1DictTypeService) GetDictTypeSelector(ctx context.Context, req *pb.GetDictTypeSelectorReq) (*pb.GetDictTypeSelectorReply, error) {
	resp := &pb.GetDictTypeSelectorReply{}
	param := &condition.Req{
		Query: []*condition.QueryParam{},
		Order: []*condition.OrderParam{},
	}
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	list, _, err := a.dictTypeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.DictTypeSelector{
			Type: v.Type,
			Name: v.Name,
		})
	}
	return resp, nil
}
