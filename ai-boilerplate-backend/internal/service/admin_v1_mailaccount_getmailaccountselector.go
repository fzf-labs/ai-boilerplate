package service

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// GetMailAccountSelector 邮箱账号表-选择器
func (a *AdminV1MailAccountService) GetMailAccountSelector(ctx context.Context, _ *pb.GetMailAccountSelectorReq) (*pb.GetMailAccountSelectorReply, error) {
	resp := &pb.GetMailAccountSelectorReply{}
	list, err := a.mailAccountRepo.FindMultiCacheByStatus(ctx, int32(constant.StatusEnable))
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 按照创建时间倒序
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	for _, v := range list {
		resp.List = append(resp.List, &pb.MailAccountSelector{
			Id:   v.ID,
			Mail: v.Mail,
		})
	}
	return resp, nil
}
