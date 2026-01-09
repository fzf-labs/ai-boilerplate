package service

import (
	"context"
	"encoding/json"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSelfAppRelease 自应用版本发布表-更新一条数据
func (a *AdminV1SelfAppReleaseService) UpdateSelfAppRelease(ctx context.Context, req *pb.UpdateSelfAppReleaseReq) (*pb.UpdateSelfAppReleaseReply, error) {
	resp := &pb.UpdateSelfAppReleaseReply{}
	data, err := a.selfAppReleaseRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	graySns, err := json.Marshal(req.GetGraySns())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data.GraySns = graySns
	oldData := a.selfAppReleaseRepo.DeepCopy(data)
	data.PackageName = req.GetPackageName()
	data.Version = req.GetVersion()
	data.BuildNum = req.GetBuildNum()
	data.Channel = req.GetChannel()
	data.UpdateType = req.GetUpdateType()
	data.Title = req.GetTitle()
	data.Changelog = req.GetChangelog()
	data.PackageURL = req.GetPackageURL()
	data.PackageSize = req.GetPackageSize()
	data.PackageMd5 = req.GetPackageMd5()
	data.MinOsVersion = req.GetMinOsVersion()
	data.GrayStrategy = req.GetGrayStrategy()
	data.GraySns = graySns
	data.PublishTime = carbon.Parse(req.GetPublishTime()).StdTime()
	data.Status = req.GetStatus()
	err = a.selfAppReleaseRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
