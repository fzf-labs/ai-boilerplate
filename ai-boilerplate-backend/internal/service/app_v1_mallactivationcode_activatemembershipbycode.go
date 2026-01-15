package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// ActivateMembershipByCode 会员激活码激活
func (s *AppV1MallActivationCodeService) ActivateMembershipByCode(ctx context.Context, req *pb.ActivateMembershipByCodeReq) (*pb.ActivateMembershipByCodeReply, error) {
	resp := &pb.ActivateMembershipByCodeReply{}
	// 获取用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}
	code := strings.TrimSpace(req.GetCode())
	if code == "" {
		return nil, pb.ErrorReasonParamError()
	}
	// 获取激活码信息
	activationCode, err := s.mallActivationCodeRepo.FindOneCacheByCode(ctx, code)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if activationCode == nil || activationCode.ID == "" {
		return nil, pb.ErrorReasonActivationCodeNotFound()
	}
	if activationCode.ProductType != constant.MallProductTypeMembership.String() {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid(pb.WithError(errors.New("product type not supported")))
	}
	// 判断激活码是否可兑换
	if !s.mallActivationCodeRepo.IsActivationCodeRedeemable(activationCode) {
		return nil, pb.ErrorReasonActivationCodeNotRedeemable()
	}
	// 获取用户会员信息
	userMembership, err := s.userMembershipRepo.FindOneCacheByUserID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if userMembership == nil || userMembership.ID == "" {
		return nil, pb.ErrorReasonUserMembershipNotFound()
	}
	// 判断用户会员状态是否正常
	if userMembership.Status != int32(constant.StatusEnable) {
		return nil, pb.ErrorReasonUserMembershipStatusInvalid(pb.WithError(errors.New("user membership status invalid")))
	}
	// 获取商品信息
	product, err := s.mallProductRepo.FindOneCacheByID(ctx, activationCode.ProductID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if product == nil || product.ID == "" || product.ProductType != constant.MallProductTypeMembership.String() {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid(pb.WithError(errors.New("product not supported")))
	}
	productConfig, err := s.mallProductRepo.GetMembershipProductConfig(product)
	if err != nil {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid(pb.WithError(err))
	}
	if productConfig.Membership.MembershipType == "" || productConfig.Membership.DurationDays <= 0 {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid(pb.WithError(err))
	}
	// 更新数据信息
	err = s.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		// 激活码信息更新
		oldActivationCode := s.mallActivationCodeRepo.DeepCopy(activationCode)
		activationCode.UserID = userID
		activationCode.Status = int32(constant.ActivationCodeStatusActivated)
		activationCode.ActivatedAt = timeutil.TimeToSQLNullTime(carbon.Now().StdTime())
		if err := s.mallActivationCodeRepo.UpdateOneCacheWithZeroByTx(ctx, tx, activationCode, oldActivationCode); err != nil {
			return err
		}
		// 用户会员信息更新
		oldUserMembership := s.userMembershipRepo.DeepCopy(userMembership)
		userMembership.MembershipType = productConfig.Membership.MembershipType
		userMembership.ExpiredAt = timeutil.TimeToSQLNullTime(carbon.Now().AddDays(int(productConfig.Membership.DurationDays)).StdTime())
		if err := s.userMembershipRepo.UpdateOneCacheWithZeroByTx(ctx, tx, userMembership, oldUserMembership); err != nil {
			return err
		}
		// 用户会员变更记录更新
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp.MembershipType = productConfig.Membership.MembershipType
	resp.ExpiredAt = userMembership.ExpiredAt.Time.Format(time.RFC3339)
	return resp, nil
}
