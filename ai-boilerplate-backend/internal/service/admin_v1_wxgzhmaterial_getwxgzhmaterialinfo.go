package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetWxGzhMaterialInfo 公众号素材表-单条数据查询
func (a *AdminV1WxGzhMaterialService) GetWxGzhMaterialInfo(ctx context.Context, req *pb.GetWxGzhMaterialInfoReq) (*pb.GetWxGzhMaterialInfoReply, error) {
	resp := &pb.GetWxGzhMaterialInfoReply{}
	data, err := a.wxGzhMaterialRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	tags := make([]string, 0)
	if data.Tags.String() != "" {
		if err := jsonutil.Unmarshal(data.Tags, &tags); err != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
		}
	}
	resp.Info = &pb.WxGzhMaterialInfo{
		Id:          data.ID,
		AppId:       data.AppID,
		Type:        data.Type,
		MediaId:     data.MediaID,
		Tags:        tags,
		UpdateTime:  data.UpdateTime.Format(time.RFC3339),
		Name:        data.Name,
		URL:         data.URL,
		CoverURL:    data.CoverURL,
		Description: data.Description,
		Newcat:      data.Newcat,
		Newsubcat:   data.Newsubcat,
		Vid:         data.Vid,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	// // 查询账号信息
	// wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, data.AppID)
	// if err != nil {
	// 	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	// }
	// officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	// if err != nil {
	// 	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	// }
	// // 获取素材信息
	// materialResp, err := officialAccount.Material.Get(ctx, data.MediaID)
	// if err != nil {
	// 	return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	// }
	// defer materialResp.Body.Close()
	// // 读取素材响应的 body 内容
	// if materialResp != nil && materialResp.Body != nil {
	// 	bodyBytes, err := io.ReadAll(materialResp.Body)
	// 	if err != nil {
	// 		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	// 	}
	// 	fmt.Println(string(bodyBytes))
	// }
	return resp, nil
}
