package auth

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/service"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
)

var AdminPrefixPathToWhiteList = map[string][]string{
	// protobuf 路由 (operation 格式: /admin.v1.ServiceName/MethodName)
	"/admin.": {
		pb.OperationSysAuthSysAuthLogin,
		pb.OperationSysAuthSysAuthLogout,
	},
}

// AdminAuthSelectorMiddleware 创建路由中间件
func AdminAuthSelectorMiddleware(
	adminV1SysAuthService *service.AdminV1SysAuthService,
	adminV1SysOperateLogService *service.AdminV1SysOperateLogService,
) middleware.Middleware {
	return selector.Server(
		AdminAuth(adminV1SysAuthService, adminV1SysOperateLogService),
	).Match(whiteListMatcher(AdminPrefixPathToWhiteList)).Build()
}

// AdminAuth 权限校验
func AdminAuth(
	adminV1SysAuthService *service.AdminV1SysAuthService,
	adminV1SysOperationLogService *service.AdminV1SysOperateLogService,
) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := http.RequestFromServerContext(ctx); ok {
				// Do something on entering
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
				checkToken, err := adminV1SysAuthService.SysAuthCheckToken(ctx, &pb.SysAuthCheckTokenReq{
					Token: token,
				})
				if err != nil {
					return nil, err
				}
				// 将JwtUID参数写进context中
				ctx = meta.SetMetadata(ctx, constant.XMdAdminID, checkToken.AdminId)
				ctx = meta.SetMetadata(ctx, constant.XMdTenantID, checkToken.TenantId)
				// 获取真实IP，使用自定义函数
				ip := GetClientIP(tr)
				defer func() {
					// Do something on exiting
					if tr.Method != "GET" {
						var traceID string
						if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
							traceID = span.TraceID().String()
						}
						var replyJSON []byte
						if err != nil {
							replyJSON, _ = json.Marshal(errors.FromError(err))
						} else {
							replyJSON, _ = json.Marshal(reply)
						}
						headerJSON, _ := json.Marshal(tr.Header)
						reqJSON, _ := json.Marshal(req)
						_, _ = adminV1SysOperationLogService.CreateSysOperateLog(ctx, &pb.CreateSysOperateLogReq{
							TenantId:  checkToken.TenantId,
							TraceId:   traceID,
							AdminId:   checkToken.AdminId,
							IP:        ip,
							URI:       tr.RequestURI,
							Useragent: tr.UserAgent(),
							Header:    string(headerJSON),
							Req:       string(reqJSON),
							Resp:      string(replyJSON),
						})
					}
				}()
			}
			return handler(ctx, req)
		}
	}
}
