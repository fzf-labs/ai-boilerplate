package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMallActivationCodeList 激活码管理表-列表数据查询
func (a *AdminV1MallActivationCodeService) GetMallActivationCodeList(ctx context.Context, req *pb.GetMallActivationCodeListReq) (*pb.GetMallActivationCodeListReply, error) {
	resp := &pb.GetMallActivationCodeListReply{
		Total: 0,
		List:  []*pb.MallActivationCodeInfo{},
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
	if req.GetProductType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_type",
			Value: req.GetProductType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetProductId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_id",
			Value: req.GetProductId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetBatchNo() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "batch_no",
			Value: req.GetBatchNo(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetCode() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "code",
			Value: req.GetCode(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetUserId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "user_id",
			Value: req.GetUserId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetPlatform() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform",
			Value: req.GetPlatform(),
		})
	}
	if req.GetPlatformOrderNo() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform_order_no",
			Value: req.GetPlatformOrderNo(),
			Exp:   condition.EQ,
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
	if len(req.GetActivatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "activated_at",
			Value: req.GetActivatedAt()[0],
			Exp:   condition.GTE,
			Logic: condition.AND,
		}, &condition.QueryParam{
			Field: "activated_at",
			Value: req.GetActivatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[0],
			Exp:   condition.GTE,
			Logic: condition.AND,
		}, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.mallActivationCodeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		productIDs := make([]string, 0)
		userIDs := make([]string, 0)
		for _, v := range list {
			productIDs = append(productIDs, v.ProductID)
			userIDs = append(userIDs, v.UserID)
		}
		productIDToProductName, err := a.mallProductRepo.ProductIDToProductName(ctx, productIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		userIDToNickname, err := a.userRepo.UserIDToNickname(ctx, userIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range list {
			userChange := &pb.UserChange{}
			if v.UserChange.String() != "" {
				err = jsonutil.Unmarshal(v.UserChange, userChange)
				if err != nil {
					return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.MallActivationCodeInfo{
				Id:                v.ID,
				ProductType:       v.ProductType,
				ProductId:         v.ProductID,
				BatchNo:           v.BatchNo,
				Code:              v.Code,
				ValidSt:           timeutil.RFC3339(v.ValidSt),
				ValidEd:           timeutil.RFC3339(v.ValidEd),
				ActivatedAt:       timeutil.RFC3339(v.ActivatedAt.Time),
				UserId:            v.UserID,
				UserChange:        userChange,
				Platform:          v.Platform,
				PlatformSoldAt:    timeutil.RFC3339(v.PlatformSoldAt.Time),
				PlatformOrderNo:   v.PlatformOrderNo,
				PlatformBuyerId:   v.PlatformBuyerID,
				PlatformBuyerName: v.PlatformBuyerName,
				Remark:            v.Remark,
				Status:            v.Status,
				CreatedAt:         v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:         v.UpdatedAt.Format(time.RFC3339),
				ProductName:       productIDToProductName[v.ProductID],
				UserNickname:      userIDToNickname[v.UserID],
			})
		}
	}
	return resp, nil
}
