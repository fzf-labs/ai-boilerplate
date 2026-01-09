package auth

import (
	"context"
	"fmt"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/service"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var AppPrefixPathToWhiteList = map[string][]string{
	// protobuf 路由 (operation 格式: /app.v1.ServiceName/MethodName)
	"/app.": {
		pb.OperationUserLogin,
		pb.OperationUserSendVerifyCode,
	},
}

// AppAuthSelectorMiddleware 创建路由中间件
func AppAuthSelectorMiddleware(
	appV1UserService *service.AppV1UserService,
) middleware.Middleware {
	return selector.Server(
		AppAuth(appV1UserService),
	).Match(whiteListMatcher(AppPrefixPathToWhiteList)).Build()
}

// AppAuth 权限校验
func AppAuth(
	appV1UserService *service.AppV1UserService,
) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := http.RequestFromServerContext(ctx); ok {
				// 获取header头部中的Authorization的值
				authorization := tr.Header.Get("Authorization")
				// 不存在则报错
				if authorization == "" {
					return nil, pb.ErrorReasonTokenNotRequest()
				}
				// token截取
				var token string
				_, err = fmt.Sscanf(authorization, "Bearer %s", &token)
				if err != nil {
					return nil, pb.ErrorReasonTokenFormatErr()
				}
				// token解析
				checkToken, err := appV1UserService.CheckToken(ctx, &pb.CheckTokenReq{
					Token: token,
				})
				if err != nil {
					return nil, err
				}
				// 将JwtUID参数写进context中
				ctx = meta.SetMetadata(ctx, constant.XMdUserID, checkToken.UserId)
				ctx = meta.SetMetadata(ctx, constant.XMdWxGzhUserID, checkToken.WxGzhUserId)
				ctx = meta.SetMetadata(ctx, constant.XMdWxXcxUserID, checkToken.WxGzhXcxId)
			}
			return handler(ctx, req)
		}
	}
}
