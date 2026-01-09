package service

import (
	"context"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/material/request"
	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/ctxutil"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// SyncWxGzhMaterial 公众号素材表-素材同步
func (a *AdminV1WxGzhMaterialService) SyncWxGzhMaterial(ctx context.Context, req *pb.SyncWxGzhMaterialReq) (*pb.SyncWxGzhMaterialReply, error) {
	resp := &pb.SyncWxGzhMaterialReply{}
	ctxn := ctxutil.CopyContextWithOutTimeOut(ctx)
	go func() {
		syncErr := a.AsyncSyncWxGzhMaterial(ctxn, req.GetAppId())
		if syncErr != nil {
			a.log.Errorf("SyncWxGzhMaterial error: %v", syncErr)
		}
	}()
	return resp, nil
}

// 异步处理素材同步
func (a *AdminV1WxGzhMaterialService) AsyncSyncWxGzhMaterial(ctx context.Context, appID string) error {
	// 查询账号信息
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, appID)
	if err != nil {
		return pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		return pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 数据库中的统计信息
	dbStats, err := a.wxGzhMaterialRepo.GetWxGzhMaterialStats(ctx, appID)
	if err != nil {
		return pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 微信的素材统计信息
	wxStats, err := officialAccount.Material.Stats(ctx)
	if err != nil {
		return pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	datas := make([]*ai_boilerplate_model.WxGzhMaterial, 0)
	// 图片不为空，则获取图片素材列表
	if wxStats.ImageCount > 0 {
		for i := 0; i < wxStats.ImageCount-int(dbStats["image"]); i += 20 {
			imgs, listErr := officialAccount.Material.List(ctx, &request.RequestMaterialBatchGetMaterial{
				Type:   "image",
				Offset: int64(i),
				Count:  20, // 每次获取20条
			})
			if listErr != nil {
				return pb.ErrorReasonAPIThirdErr(pb.WithError(listErr))
			}
			for _, img := range imgs.Item {
				imgMap := img.ToHashMap()
				mediaID := cast.ToString(imgMap.Get("media_id"))
				name := cast.ToString(imgMap.Get("name"))
				tags := cast.ToStringSlice(imgMap.Get("tags"))
				tagsJSON, marshalErr := jsonutil.Marshal(tags)
				if marshalErr != nil {
					return pb.ErrorReasonDataFormattingError(pb.WithError(marshalErr))
				}
				updateTime := carbon.CreateFromTimestamp(cast.ToInt64(imgMap.Get("update_time"))).StdTime()
				url := cast.ToString(imgMap.Get("url"))
				datas = append(datas, &ai_boilerplate_model.WxGzhMaterial{
					AppID:      appID,
					Type:       "image",
					MediaID:    mediaID,
					Tags:       tagsJSON,
					UpdateTime: updateTime,
					Name:       name,
					URL:        url,
				})
			}
		}
	}
	// 音频不为空，则获取音频素材列表
	if wxStats.VoiceCount > 0 {
		for i := 0; i < wxStats.VoiceCount-int(dbStats["voice"]); i += 20 {
			voices, listErr := officialAccount.Material.List(ctx, &request.RequestMaterialBatchGetMaterial{
				Type:   "voice",
				Offset: int64(i),
				Count:  20, // 每次获取20条
			})
			if listErr != nil {
				return pb.ErrorReasonAPIThirdErr(pb.WithError(listErr))
			}
			for _, voice := range voices.Item {
				voiceMap := voice.ToHashMap()
				mediaID := cast.ToString(voiceMap.Get("media_id"))
				name := cast.ToString(voiceMap.Get("name"))
				tags := cast.ToStringSlice(voiceMap.Get("tags"))
				tagsJSON, marshalErr := jsonutil.Marshal(tags)
				if marshalErr != nil {
					return pb.ErrorReasonDataFormattingError(pb.WithError(marshalErr))
				}
				updateTime := carbon.CreateFromTimestamp(cast.ToInt64(voiceMap.Get("update_time"))).StdTime()
				datas = append(datas, &ai_boilerplate_model.WxGzhMaterial{
					AppID:      appID,
					Type:       "voice",
					MediaID:    mediaID,
					Tags:       tagsJSON,
					UpdateTime: updateTime,
					Name:       name,
				})
			}
		}
	}
	// 视频不为空，则获取视频素材列表
	if wxStats.VideoCount > 0 {
		for i := 0; i < wxStats.VideoCount-int(dbStats["video"]); i += 20 {
			videos, listErr := officialAccount.Material.List(ctx, &request.RequestMaterialBatchGetMaterial{
				Type:   "video",
				Offset: int64(i),
				Count:  20, // 每次获取20条
			})
			if listErr != nil {
				return pb.ErrorReasonAPIThirdErr(pb.WithError(listErr))
			}
			for _, video := range videos.Item {
				videoMap := video.ToHashMap()
				mediaID := cast.ToString(videoMap.Get("media_id"))
				name := cast.ToString(videoMap.Get("name"))
				tags := cast.ToStringSlice(videoMap.Get("tags"))
				tagsJSON, marshalErr := jsonutil.Marshal(tags)
				if marshalErr != nil {
					return pb.ErrorReasonDataFormattingError(pb.WithError(marshalErr))
				}
				updateTime := carbon.CreateFromTimestamp(cast.ToInt64(videoMap.Get("update_time"))).StdTime()
				coverURL := cast.ToString(videoMap.Get("cover_url"))
				description := cast.ToString(videoMap.Get("description"))
				newcat := cast.ToString(videoMap.Get("newcat"))
				newsubcat := cast.ToString(videoMap.Get("newsubcat"))
				vid := cast.ToString(videoMap.Get("vid"))
				// 获取视频下载 URL
				videoResp, getErr := officialAccount.Material.GetVideo(ctx, mediaID)
				if getErr != nil {
					return pb.ErrorReasonAPIThirdErr(pb.WithError(getErr))
				}
				datas = append(datas, &ai_boilerplate_model.WxGzhMaterial{
					AppID:       appID,
					Type:        "video",
					MediaID:     mediaID,
					Tags:        tagsJSON,
					UpdateTime:  updateTime,
					Name:        name,
					CoverURL:    coverURL,
					Description: description,
					Newcat:      newcat,
					Newsubcat:   newsubcat,
					Vid:         vid,
					URL:         videoResp.DownUrl,
				})
			}
		}
	}
	// 过滤已经存在的素材
	mediaIDs := lo.Map(datas, func(item *ai_boilerplate_model.WxGzhMaterial, _ int) string {
		return item.MediaID
	})
	existingMaterials, err := a.wxGzhMaterialRepo.FindMultiByMediaIDS(ctx, mediaIDs)
	if err != nil {
		return pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	datas = lo.Filter(datas, func(item *ai_boilerplate_model.WxGzhMaterial, _ int) bool {
		return !lo.ContainsBy(existingMaterials, func(existingMaterial *ai_boilerplate_model.WxGzhMaterial) bool {
			return existingMaterial.MediaID == item.MediaID
		})
	})
	// 插入素材
	if len(datas) > 0 {
		err := a.wxGzhMaterialRepo.CreateBatchCache(ctx, datas, 100)
		if err != nil {
			return pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	return nil
}

// type Image struct {
// 	MediaID    string        `json:"media_id"`
// 	Name       string        `json:"name"`
// 	Tags       []interface{} `json:"tags"`
// 	UpdateTime int64         `json:"update_time"`
// 	URL        string        `json:"url"`
// }

// type Voice struct {
// 	MediaID    string        `json:"media_id"`
// 	Name       string        `json:"name"`
// 	Tags       []interface{} `json:"tags"`
// 	UpdateTime int64         `json:"update_time"`
// }

// type Video struct {
// 	MediaID     string        `json:"media_id"`    // 媒体文件 ID
// 	Name        string        `json:"name"`        // 媒体文件名称
// 	Tags        []interface{} `json:"tags"`        // 媒体文件标签
// 	UpdateTime  int64         `json:"update_time"` // 媒体文件更新时间戳
// 	CoverURL    string        `json:"cover_url"`   // 视频封面 URL
// 	Description string        `json:"description"` // 视频描述
// 	Newcat      string        `json:"newcat"`      // 视频分类
// 	Newsubcat   string        `json:"newsubcat"`   // 视频子分类
// 	Vid         string        `json:"vid"`         // 视频 ID
// }
