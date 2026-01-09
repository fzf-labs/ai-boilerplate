package service

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSelfAppReleaseList 自应用版本发布表-列表数据查询
func (a *AdminV1SelfAppReleaseService) GetSelfAppReleaseList(ctx context.Context, req *pb.GetSelfAppReleaseListReq) (*pb.GetSelfAppReleaseListReply, error) {
	resp := &pb.GetSelfAppReleaseListReply{
		Total: 0,
		List:  []*pb.SelfAppReleaseInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "package_name",
				Value: req.GetPackageName(),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetChannel() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "channel",
			Value: req.GetChannel(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetBuildNum() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "build_num",
			Value: req.GetBuildNum(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.selfAppReleaseRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			graySns := []string{}
			if v.GraySns.String() != "" {
				err := json.Unmarshal(v.GraySns, &graySns)
				if err != nil {
					return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.SelfAppReleaseInfo{
				Id:           v.ID,
				PackageName:  v.PackageName,
				Version:      v.Version,
				BuildNum:     v.BuildNum,
				Channel:      v.Channel,
				UpdateType:   v.UpdateType,
				Title:        v.Title,
				Changelog:    v.Changelog,
				PackageURL:   v.PackageURL,
				PackageSize:  v.PackageSize,
				PackageMd5:   v.PackageMd5,
				MinOsVersion: v.MinOsVersion,
				GrayStrategy: v.GrayStrategy,
				GraySns:      graySns,
				Status:       v.Status,
				PublishTime:  timeutil.RFC3339(v.PublishTime),
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
