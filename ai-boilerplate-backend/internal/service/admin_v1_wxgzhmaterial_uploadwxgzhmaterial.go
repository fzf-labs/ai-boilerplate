package service

import (
	"context"
	"io"
	"os"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/material/response"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func (a *AdminV1WxGzhMaterialService) UploadWxGzhMaterialHandler(ctx http.Context) error {
	// 设置操作路径，用于中间件鉴权
	http.SetOperation(ctx, "/admin.v1.WxGzhMaterial/UploadWxGzhMaterial")

	// 使用中间件包装处理逻辑
	h := ctx.Middleware(func(ctx context.Context, _ interface{}) (interface{}, error) {
		// 从 HTTP Context 获取原始的 http.Request 和 http.ResponseWriter
		tr, _ := http.RequestFromServerContext(ctx)
		r := tr
		appID := r.FormValue("appId")
		materialType := r.FormValue("type")
		materialTitle := r.FormValue("title")
		materialIntroduction := r.FormValue("introduction")
		// 获取文件并写入到服务器的临时目录
		materialFile, fileHeader, err := r.FormFile("file")
		if err != nil {
			return nil, err
		}
		defer materialFile.Close()

		// 创建临时文件
		tempFile, err := os.CreateTemp("", "upload_*_"+fileHeader.Filename)
		if err != nil {
			return nil, err
		}
		defer tempFile.Close()
		defer os.Remove(tempFile.Name()) // 清理临时文件

		// 将上传的文件内容复制到临时文件
		_, err = io.Copy(tempFile, materialFile)
		if err != nil {
			return nil, err
		}
		filePath := tempFile.Name()

		// 查询账号信息
		wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, appID)
		if err != nil {
			return nil, err
		}

		officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
		if err != nil {
			return nil, err
		}

		materialResp := &response.ResponseMaterialAddMaterial{}
		switch materialType {
		case "video":
			materialResp, err = officialAccount.Material.UploadVideo(ctx, filePath, materialTitle, materialIntroduction)
			if err != nil {
				return nil, err
			}
		case "image":
			materialResp, err = officialAccount.Material.UploadImage(ctx, filePath)
			if err != nil {
				return nil, err
			}
		case "voice":
			materialResp, err = officialAccount.Material.UploadVoice(ctx, filePath)
			if err != nil {
				return nil, err
			}
		}

		// 如果素材ID为空，则上传失败
		if materialResp.MediaID == "" {
			return nil, pb.ErrorReasonMaterialUploadFailed()
		}

		// 异步同步素材
		go func() {
			syncErr := a.AsyncSyncWxGzhMaterial(context.Background(), appID)
			if syncErr != nil {
				a.log.Errorf("AsyncSyncWxGzhMaterial error: %v", syncErr)
				return
			}
		}()

		return materialResp, nil
	})

	// 执行中间件链
	reply, err := h(ctx, nil)
	if err != nil {
		return err
	}

	// 返回成功响应
	return ctx.Result(200, reply)
}
