package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/kratos-contrib/meta"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type membershipProductConfig struct {
	Membership struct {
		MembershipType string `json:"membershipType"`
		DurationDays   int32  `json:"duration_days"`
	} `json:"membership"`
}

// ActivateMembershipByCode 会员激活码激活
func (s *AppV1MallActivationCodeService) ActivateMembershipByCode(ctx context.Context, req *pb.ActivateMembershipByCodeReq) (*pb.ActivateMembershipByCodeReply, error) {
	resp := &pb.ActivateMembershipByCodeReply{}

	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	code := strings.TrimSpace(req.GetCode())
	if code == "" {
		return nil, pb.ErrorReasonParamError()
	}

	activationCode, err := s.mallActivationCodeRepo.FindOneCacheByCode(ctx, code)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if activationCode == nil || activationCode.ID == "" {
		return nil, pb.ErrorReasonActivationCodeNotFound()
	}
	if activationCode.ProductType != constant.MallProductTypeMembership.String() {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid()
	}

	product, err := s.mallProductRepo.FindOneCacheByID(ctx, activationCode.ProductID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if product == nil || product.ID == "" || product.ProductType != constant.MallProductTypeMembership.String() {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid()
	}

	membershipType, durationDays, err := parseMembershipProductConfig(product.ProductConfig)
	if err != nil {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid()
	}
	if _, err := constant.ParseMembershipType(membershipType); err != nil {
		return nil, pb.ErrorReasonActivationCodeProductConfigInvalid()
	}

	var expiredAt time.Time
	err = s.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		now := time.Now()
		codeDAO := tx.MallActivationCode
		codeData, err := codeDAO.WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where(codeDAO.Code.Eq(code)).
			First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return pb.ErrorReasonActivationCodeNotFound()
			}
			return err
		}
		if !isActivationCodeRedeemable(codeData, now) {
			return pb.ErrorReasonActivationCodeNotRedeemable()
		}

		oldCode := s.mallActivationCodeRepo.DeepCopy(codeData)
		codeData.Status = int32(constant.ActivationCodeStatusActivated)
		codeData.UserID = userID
		codeData.ActivatedAt = sql.NullTime{Time: now, Valid: true}
		if err := s.mallActivationCodeRepo.UpdateOneCacheWithZeroByTx(ctx, tx, codeData, oldCode); err != nil {
			return err
		}

		membershipDAO := tx.UserMembership
		membershipData, err := membershipDAO.WithContext(ctx).Where(membershipDAO.UserID.Eq(userID)).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return pb.ErrorReasonDataRecordNotFound()
			}
			return err
		}

		oldMembership := s.userMembershipRepo.DeepCopy(membershipData)
		baseTime := now
		if membershipData.ExpiredAt.Valid && membershipData.ExpiredAt.Time.After(now) {
			baseTime = membershipData.ExpiredAt.Time
		}
		expiredAt = baseTime.AddDate(0, 0, int(durationDays))
		membershipData.MembershipType = membershipType
		membershipData.ExpiredAt = sql.NullTime{Time: expiredAt, Valid: true}
		membershipData.Status = 1
		if err := s.userMembershipRepo.UpdateOneCacheWithZeroByTx(ctx, tx, membershipData, oldMembership); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if kerrors.FromError(err) != nil {
			return nil, err
		}
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.MembershipType = membershipType
	resp.ExpiredAt = expiredAt.Format(time.RFC3339)
	return resp, nil
}

func parseMembershipProductConfig(raw []byte) (string, int32, error) {
	if len(raw) == 0 {
		return "", 0, errors.New("empty product config")
	}

	var cfg membershipProductConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return "", 0, err
	}
	if cfg.Membership.MembershipType == "" || cfg.Membership.DurationDays <= 0 {
		return "", 0, errors.New("invalid membership config")
	}
	return cfg.Membership.MembershipType, cfg.Membership.DurationDays, nil
}

func isActivationCodeRedeemable(data *ai_boilerplate_model.MallActivationCode, now time.Time) bool {
	switch constant.ActivationCodeStatus(data.Status) {
	case constant.ActivationCodeStatusDisable,
		constant.ActivationCodeStatusRefunded,
		constant.ActivationCodeStatusActivated,
		constant.ActivationCodeStatusExpired:
		return false
	default:
	}
	if now.Before(data.ValidSt) || now.After(data.ValidEd) {
		return false
	}
	return true
}
