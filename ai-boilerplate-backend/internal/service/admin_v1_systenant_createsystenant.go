package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/cryptutil"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/google/uuid"
)

// CreateSysTenant 系统-租户-创建一条数据
func (a *AdminV1SysTenantService) CreateSysTenant(ctx context.Context, req *pb.CreateSysTenantReq) (*pb.CreateSysTenantReply, error) {
	resp := &pb.CreateSysTenantReply{}
	err := a.commonRepo.Transaction(ctx, func(_ *ai_boilerplate_dao.Query) error {
		// 查询用户名是否存在
		sysAdminExist, err := a.sysAdminRepo.FindOneCacheByUsername(ctx, req.GetUsername())
		if err != nil {
			return pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		if sysAdminExist != nil && sysAdminExist.ID != "" {
			return pb.ErrorReasonAccountAlreadyExists()
		}
		password, err := cryptutil.Encrypt(req.GetPassword())
		if err != nil {
			return pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		sysAdminID := uuid.New().String()
		sysTenantID := uuid.New().String()
		// 创建用户
		sysAdmin := ai_boilerplate_model.SysAdmin{
			ID:       sysAdminID,
			Username: req.GetUsername(),
			Password: password,
			Nickname: req.GetUsername(),
			DeptID:   sysTenantID,
			Status:   int16(constant.SysStatusDisable),
		}
		err = a.sysAdminRepo.CreateOneCache(ctx, &sysAdmin)
		if err != nil {
			return pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		// 创建租户
		menuIDs, err := jsonutil.Marshal(req.GetMenuIds())
		if err != nil {
			return pb.ErrorReasonDataFormattingError(pb.WithError(err))
		}
		sysTenant := ai_boilerplate_model.SysTenant{
			ID:         sysTenantID,
			Name:       req.GetName(),
			Remark:     req.GetRemark(),
			AdminID:    sysAdminID,
			ExpireTime: timeutil.TimeToSQLNullTime(carbon.Parse(req.GetExpireTime()).StdTime()),
			MenuIds:    menuIDs,
			Status:     int16(req.GetStatus()),
		}
		err = a.sysTenantRepo.CreateOneCache(ctx, &sysTenant)
		if err != nil {
			return pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		resp.Id = sysTenantID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
