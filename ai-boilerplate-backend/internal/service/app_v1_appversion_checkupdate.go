package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// CheckUpdate 检测版本更新
func (a *AppV1AppVersionService) CheckUpdate(ctx context.Context, req *pb.CheckUpdateReq) (*pb.CheckUpdateReply, error) {
	resp := &pb.CheckUpdateReply{HasUpdate: false}
	// 查询该包名和渠道下最新的已发布版本
	releases, _, err := a.selfAppReleaseRepo.FindMultiByCondition(ctx, &condition.Req{
		Query: []*condition.QueryParam{
			{Field: "package_name", Value: req.GetPackageName(), Exp: condition.EQ, Logic: condition.AND},
			{Field: "channel", Value: req.GetChannel(), Exp: condition.EQ, Logic: condition.AND},
			{Field: "status", Value: 1, Exp: condition.EQ, Logic: condition.AND},
		},
		Order: []*condition.OrderParam{
			{Field: "build_num", Order: condition.DESC},
		},
		Page:     1,
		PageSize: 1,
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 无可用版本
	if len(releases) == 0 {
		return resp, nil
	}
	latest := releases[0]
	// 比较 build 号，判断是否需要更新
	if latest.BuildNum <= req.GetBuildNum() {
		return resp, nil
	}
	// 有新版本
	resp.HasUpdate = true
	resp.Info = &pb.AppVersionInfo{
		Version:    latest.Version,
		BuildNum:   latest.BuildNum,
		Title:      latest.Title,
		Changelog:  latest.Changelog,
		PackageUrl: latest.PackageURL,
		UpdateType: latest.UpdateType,
	}
	return resp, nil
}
