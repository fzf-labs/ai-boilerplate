package service

import (
	"context"
	"encoding/json"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// CheckUpdate 检测App版本更新
func (a *AppV1SelfAppReleaseService) CheckUpdate(ctx context.Context, req *pb.CheckUpdateReq) (*pb.CheckUpdateReply, error) {
	resp := &pb.CheckUpdateReply{HasUpdate: false}

	// 查询最新发布的版本（状态启用、build号大于当前版本）
	releases, _, err := a.selfAppReleaseRepo.FindMultiByCondition(ctx, &condition.Req{
		Query: []*condition.QueryParam{
			{Field: "package_name", Value: req.GetPackageName(), Exp: condition.EQ, Logic: condition.AND},
			{Field: "channel", Value: req.GetChannel(), Exp: condition.EQ, Logic: condition.AND},
			{Field: "build_num", Value: req.GetBuildNum(), Exp: condition.GT, Logic: condition.AND},
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

	if len(releases) == 0 {
		return resp, nil
	}

	release := releases[0]

	// 检查灰度策略
	if release.GrayStrategy == 2 && req.GetDeviceSn() != "" {
		// 灰度策略为自定义设备，检查设备是否在灰度列表中
		var graySns []string
		if release.GraySns != nil {
			if err := json.Unmarshal(release.GraySns, &graySns); err != nil {
				a.log.Warnf("parse gray_sns failed: %v", err)
			}
		}
		if len(graySns) > 0 {
			inGray := false
			for _, sn := range graySns {
				if sn == req.GetDeviceSn() {
					inGray = true
					break
				}
			}
			if !inGray {
				return resp, nil
			}
		}
	}

	// 检查最低系统版本要求
	if release.MinOsVersion != "" && req.GetOsVersion() != "" {
		if !isOsVersionSatisfied(req.GetOsVersion(), release.MinOsVersion) {
			return resp, nil
		}
	}

	// 有可用更新
	resp.HasUpdate = true
	resp.UpdateInfo = &pb.UpdateInfo{
		Version:      release.Version,
		BuildNum:     release.BuildNum,
		UpdateType:   release.UpdateType,
		Title:        release.Title,
		Changelog:    release.Changelog,
		PackageURL:   release.PackageURL,
		PackageSize:  release.PackageSize,
		PackageMd5:   release.PackageMd5,
		MinOsVersion: release.MinOsVersion,
	}

	return resp, nil
}

// isOsVersionSatisfied 检查系统版本是否满足最低要求
func isOsVersionSatisfied(current, minimum string) bool {
	// 简单的版本比较，假设版本格式为 x.y.z
	// 实际项目中可能需要更复杂的版本比较逻辑
	return current >= minimum
}
