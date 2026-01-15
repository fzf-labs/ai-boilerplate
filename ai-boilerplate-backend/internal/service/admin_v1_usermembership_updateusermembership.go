package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// UpdateUserMembership 用户会员关系表-更新一条数据
func (a *AdminV1UserMembershipService) UpdateUserMembership(ctx context.Context, req *pb.UpdateUserMembershipReq) (*pb.UpdateUserMembershipReply, error) {
	resp := &pb.UpdateUserMembershipReply{}
	data, err := a.userMembershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	err = a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		// 更新用户会员关系表
		oldData := a.userMembershipRepo.DeepCopy(data)
		data.UserID = req.GetUserId()
		data.MembershipType = req.GetMembershipType()
		data.ExpiredAt = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetExpiredAt()).StdTime())
		data.AutoRenew = req.GetAutoRenew()
		data.AutoRenewDays = req.GetAutoRenewDays()
		data.Status = req.GetStatus()
		if err := a.userMembershipRepo.UpdateOneCacheWithZeroByTx(ctx, tx, data, oldData); err != nil {
			return err
		}
		// 创建用户会员变更记录表
		before, err := jsonutil.Marshal(oldData)
		if err != nil {
			return err
		}
		after, err := jsonutil.Marshal(data)
		if err != nil {
			return err
		}
		userMembershipChange := &ai_boilerplate_model.UserMembershipChange{
			UserID:     data.UserID,
			SourceType: constant.MembershipChangeSourceAdmin.String(),
			SourceID:   data.ID,
			Before:     before,
			After:      after,
		}
		if err := a.userMembershipChangeRepo.CreateOneCacheByTx(ctx, tx, userMembershipChange); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
