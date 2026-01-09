package service

import (
	"context"
	"encoding/json"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSelfAppRelease 自应用版本发布表-创建一条数据
func (a *AdminV1SelfAppReleaseService) CreateSelfAppRelease(ctx context.Context, req *pb.CreateSelfAppReleaseReq) (*pb.CreateSelfAppReleaseReply, error) {
	resp := &pb.CreateSelfAppReleaseReply{}
	graySns, err := json.Marshal(req.GetGraySns())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := a.selfAppReleaseRepo.NewData()
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
	err = a.selfAppReleaseRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
