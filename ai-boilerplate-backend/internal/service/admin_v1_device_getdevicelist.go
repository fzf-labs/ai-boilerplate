package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/samber/lo"
)

// GetDeviceList 设备表-列表数据查询
func (a *AdminV1DeviceService) GetDeviceList(ctx context.Context, req *pb.GetDeviceListReq) (*pb.GetDeviceListReply, error) {
	resp := &pb.GetDeviceListReply{
		Total: 0,
		List:  []*pb.DeviceInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetSn() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "sn",
			Value: "%" + req.GetSn() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetRegistryTime()) > 0 {
		param.Query = append(param.Query,
			&condition.QueryParam{
				Field: "registry_time",
				Value: req.GetRegistryTime()[0],
				Exp:   condition.GTE,
				Logic: condition.AND,
			},
			&condition.QueryParam{
				Field: "registry_time",
				Value: req.GetRegistryTime()[1],
				Exp:   condition.LTE,
				Logic: condition.AND,
			},
		)
	}
	switch req.GetOnlineSearch() {
	case "online":
		sns, err := a.deviceHeartbeatRepo.GetAllOnlineDevices(ctx)
		if err != nil {
			return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
		}
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "sn",
			Value: sns,
			Exp:   condition.IN,
			Logic: condition.AND,
		})
	case "offline":
		sns, err := a.deviceHeartbeatRepo.GetAllOnlineDevices(ctx)
		if err != nil {
			return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
		}
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "sn",
			Value: sns,
			Exp:   condition.NOTIN,
			Logic: condition.AND,
		})
	default:
	}
	list, p, err := a.deviceRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		sns := lo.Map(list, func(item *ai_boilerplate_model.Device, _ int) string {
			return item.Sn
		})
		onlineMap, err := a.deviceHeartbeatRepo.IsDeviceOnlineBatch(ctx, sns)
		if err != nil {
			return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
		}
		for _, v := range list {
			devicePush := &pb.DevicePush{}
			if v.Push.String() != "" {
				if err := jsonutil.Unmarshal(v.Push, &devicePush); err != nil {
					return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.DeviceInfo{
				Id:             v.ID,
				Sn:             v.Sn,
				Name:           v.Name,
				Desc:           v.Desc,
				Brand:          v.Brand,
				Model:          v.Model,
				Network:        v.Network,
				Imei:           v.Imei,
				CPU:            v.CPU,
				Mac:            v.Mac,
				AppVersion:     v.AppVersion,
				AndroidVersion: v.AndroidVersion,
				RAMSize:        v.RAMSize,
				DdrSize:        v.DdrSize,
				Certificate:    v.Certificate,
				SecureKey:      v.SecureKey,
				RegistryTime:   v.RegistryTime.Time.Format(time.RFC3339),
				Push:           devicePush,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
				Online:         onlineMap[v.Sn],
			})
		}
	}
	return resp, nil
}
