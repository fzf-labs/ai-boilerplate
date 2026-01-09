package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/samber/lo"
)

// SyncWxGzhTag 公众号标签表-同步标签
func (a *AdminV1WxGzhTagService) SyncWxGzhTag(ctx context.Context, req *pb.SyncWxGzhTagReq) (*pb.SyncWxGzhTagReply, error) {
	resp := &pb.SyncWxGzhTagReply{}
	// 查询账号信息
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 微信公众号同步标签
	wxTagList, err := officialAccount.UserTag.List(ctx)
	if err != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	wxTagDatas := make([]*ai_boilerplate_model.WxGzhTag, 0)
	// 获取微信标签 id 列表
	wxTagIDs := make([]int32, 0)
	for _, v := range wxTagList.Tags {
		wxTagIDs = append(wxTagIDs, int32(v.ID))
		wxTagDatas = append(wxTagDatas, &ai_boilerplate_model.WxGzhTag{
			AppID: wxGzhAccount.AppID,
			TagID: int32(v.ID),
			Name:  v.Name,
		})
	}
	dbTagDatas, err := a.wxGzhTagRepo.FindMultiCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	dbTagIDs := lo.Map(dbTagDatas, func(item *ai_boilerplate_model.WxGzhTag, _ int) int32 {
		return item.TagID
	})
	// 取交集
	updateTagIDs := lo.Intersect(wxTagIDs, dbTagIDs)
	// 取差集
	createTagIDs, deleteTagIDs := lo.Difference(wxTagIDs, dbTagIDs)
	// 创建数据
	if len(createTagIDs) > 0 {
		createDatas := lo.Filter(wxTagDatas, func(item *ai_boilerplate_model.WxGzhTag, _ int) bool {
			return lo.Contains(createTagIDs, item.TagID)
		})
		if len(createDatas) > 0 {
			err = a.wxGzhTagRepo.CreateBatchCache(ctx, createDatas, 100)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}
	}
	// 删除数据
	if len(deleteTagIDs) > 0 {
		deleteDatas := lo.Filter(dbTagDatas, func(item *ai_boilerplate_model.WxGzhTag, _ int) bool {
			return lo.Contains(deleteTagIDs, item.TagID)
		})
		if len(deleteDatas) > 0 {
			deleteIDs := lo.Map(deleteDatas, func(item *ai_boilerplate_model.WxGzhTag, _ int) string {
				return item.ID
			})
			err = a.wxGzhTagRepo.DeleteMultiCacheByIDS(ctx, deleteIDs)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}
	}
	// 更新数据
	if len(updateTagIDs) > 0 {
		updateDatas := lo.Filter(dbTagDatas, func(item *ai_boilerplate_model.WxGzhTag, _ int) bool {
			return lo.Contains(updateTagIDs, item.TagID)
		})
		tagIDToName := make(map[int32]string)
		for _, v := range wxTagDatas {
			tagIDToName[v.TagID] = v.Name
		}
		if len(updateDatas) > 0 {
			for _, v := range updateDatas {
				oldData := a.wxGzhTagRepo.DeepCopy(v)
				v.Name = tagIDToName[v.TagID]
				err = a.wxGzhTagRepo.UpdateOneCacheWithZero(ctx, v, oldData)
				if err != nil {
					return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
				}
			}
		}
	}
	return resp, nil
}
