package service

import (
	"context"
	"encoding/json"
	"slices"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetMembershipProductList 获取会员商品列表
func (a *AppV1MallProductService) GetMembershipProductList(ctx context.Context, req *pb.GetMembershipProductListReq) (*pb.GetMembershipProductListReply, error) {
	resp := &pb.GetMembershipProductListReply{}
	userId := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	// 查询用户会员信息
	userMembership, err := a.userMembershipRepo.GetUserActualMembershipInfo(ctx, userId)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 会员信息
	membership, err := a.membershipRepo.FindOneCacheByType(ctx, userMembership.MembershipType)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	memberships, _, err := a.membershipRepo.FindMultiCacheByCondition(ctx, &condition.Req{
		Query: []*condition.QueryParam{
			{
				Field: "level",
				Value: membership.Level,
				Exp:   condition.GTE,
				Logic: condition.AND,
			},
			{
				Field: "status",
				Value: constant.StatusEnable,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	membershipTypes := make([]string, 0)
	for _, membership := range memberships {
		membershipTypes = append(membershipTypes, membership.Type)
	}
	// 查询会员商品列表
	products, err := a.mallProductRepo.FindMultiCacheByProductType(ctx, string(constant.MallProductTypeMembership))
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 只获取会员商品等级大于等于用户会员等级的商品
	for _, product := range products {
		productImages := make([]string, 0)
		if product.ProductImages.String() != "" {
			err = json.Unmarshal(product.ProductImages, &productImages)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}
		productDetail := make([]string, 0)
		if product.ProductDetail.String() != "" {
			err = json.Unmarshal(product.ProductDetail, &productDetail)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}
		productConfig := &pb.ProductConfig{}
		if product.ProductConfig.String() != "" {
			err = json.Unmarshal(product.ProductConfig, productConfig)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}
		if slices.Contains(membershipTypes, productConfig.Membership.MembershipType) {
			resp.List = append(resp.List, &pb.MallProductInfo{
				Id:            product.ID,
				ProductType:   product.ProductType,
				ProductName:   product.ProductName,
				ProductDesc:   product.ProductDesc,
				ProductImages: productImages,
				ProductDetail: productDetail,
				ProductConfig: productConfig,
				OriginalPrice: product.OriginalPrice,
				CurrentPrice:  product.CurrentPrice,
				StockQuantity: product.StockQuantity,
				SoldQuantity:  product.SoldQuantity,
				Sort:          product.Sort,
				Status:        product.Status,
				CreatedAt:     product.CreatedAt.Format(time.RFC3339),
				UpdatedAt:     product.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	// 按排序升序
	slices.SortFunc(resp.List, func(a, b *pb.MallProductInfo) int { return int(a.Sort - b.Sort) })
	return resp, nil
}
