package service

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSelfAppReleaseInfo 自应用版本发布表-单条数据查询
func (a *AdminV1SelfAppReleaseService) GetSelfAppReleaseInfo(ctx context.Context, req *pb.GetSelfAppReleaseInfoReq) (*pb.GetSelfAppReleaseInfoReply, error) {
	resp := &pb.GetSelfAppReleaseInfoReply{}
	data, err := a.selfAppReleaseRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	graySns := []string{}
	if data.GraySns.String() != "" {
		err := json.Unmarshal(data.GraySns, &graySns)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	resp.Info = &pb.SelfAppReleaseInfo{
		Id:           data.ID,
		PackageName:  data.PackageName,
		Version:      data.Version,
		BuildNum:     data.BuildNum,
		Channel:      data.Channel,
		UpdateType:   data.UpdateType,
		Title:        data.Title,
		Changelog:    data.Changelog,
		PackageURL:   data.PackageURL,
		PackageSize:  data.PackageSize,
		PackageMd5:   data.PackageMd5,
		MinOsVersion: data.MinOsVersion,
		GrayStrategy: data.GrayStrategy,
		GraySns:      graySns,
		PublishTime:  timeutil.RFC3339(data.PublishTime),
		Status:       data.Status,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
