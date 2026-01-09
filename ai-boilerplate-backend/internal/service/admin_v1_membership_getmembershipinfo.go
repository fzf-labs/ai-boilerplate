package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMembershipInfo 会员类型配置表-单条数据查询
func (a *AdminV1MembershipService) GetMembershipInfo(ctx context.Context, req *pb.GetMembershipInfoReq) (*pb.GetMembershipInfoReply, error) {
	resp := &pb.GetMembershipInfoReply{}
	data, err := a.membershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MembershipInfo{
		Id:          data.ID,
		Name:        data.Name,
		Type:        data.Type,
		Description: data.Description,
		Sort:        data.Sort,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
