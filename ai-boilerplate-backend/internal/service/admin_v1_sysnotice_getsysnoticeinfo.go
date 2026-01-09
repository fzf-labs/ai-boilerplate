package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysNoticeInfo 系统-公告-单条数据查询
func (a *AdminV1SysNoticeService) GetSysNoticeInfo(ctx context.Context, req *pb.GetSysNoticeInfoReq) (*pb.GetSysNoticeInfoReply, error) {
	resp := &pb.GetSysNoticeInfoReply{}
	data, err := a.sysNoticeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysNoticeInfo{
		Id:        data.ID,
		Type:      data.Type,
		Title:     data.Title,
		Content:   data.Content,
		Status:    int32(data.Status),
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
