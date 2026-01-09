package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetFileDatumInfo 文件表-单条数据查询
func (a *AdminV1FileDatumService) GetFileDatumInfo(ctx context.Context, req *pb.GetFileDatumInfoReq) (*pb.GetFileDatumInfoReply, error) {
	resp := &pb.GetFileDatumInfoReply{}
	data, err := a.fileDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.FileDatumInfo{
		Id:        data.ID,
		ConfigId:  data.ConfigID,
		Name:      data.Name,
		Path:      data.Path,
		URL:       data.URL,
		Ext:       data.Ext,
		Size:      data.Size,
		Status:    data.Status,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
