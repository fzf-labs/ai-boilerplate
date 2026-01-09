package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSelfAppReleaseRepo(
	logger log.Logger,
	data *Data,
	selfAppReleaseRepo *ai_boilerplate_repo.SelfAppReleaseRepo,
) *SelfAppReleaseRepo {
	l := log.NewHelper(log.With(logger, "module", "data/selfAppRelease"))
	return &SelfAppReleaseRepo{
		log:                l,
		data:               data,
		SelfAppReleaseRepo: selfAppReleaseRepo,
	}
}

type SelfAppReleaseRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SelfAppReleaseRepo
}

// AppReleaseInfo 版本信息
type AppReleaseInfo struct {
	Version     string
	UpdateDesc  string
	DownloadURL string
	ForceUpdate int32
}

// FindLatestByPlatform 根据平台查询最新版本
func (r *SelfAppReleaseRepo) FindLatestByPlatform(ctx context.Context, platform string) (*AppReleaseInfo, error) {
	// 根据平台映射到渠道
	channel := "android"
	if platform == "ios" {
		channel = "ios"
	}
	// 查询最新发布的版本
	releases, _, err := r.FindMultiByCondition(ctx, &condition.Req{
		Query: []*condition.QueryParam{
			{Field: "channel", Value: channel, Exp: condition.EQ, Logic: condition.AND},
			{Field: "status", Value: 1, Exp: condition.EQ, Logic: condition.AND},
		},
		Order: []*condition.OrderParam{
			{Field: "build_num", Order: condition.DESC},
		},
		Page:     1,
		PageSize: 1,
	})
	if err != nil {
		return nil, err
	}
	if len(releases) == 0 {
		return nil, nil
	}
	release := releases[0]
	forceUpdate := int32(0)
	if release.UpdateType == 1 {
		forceUpdate = 1
	}
	return &AppReleaseInfo{
		Version:     release.Version,
		UpdateDesc:  release.Changelog,
		DownloadURL: release.PackageURL,
		ForceUpdate: forceUpdate,
	}, nil
}
