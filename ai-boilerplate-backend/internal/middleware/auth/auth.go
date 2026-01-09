package auth

import (
	"context"
	"net"
	"strings"

	"github.com/fzf-labs/goutil/iputil"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/samber/lo"
)

// whiteListMatcher 路由白名单匹配器  true 校验 false 不校验
func whiteListMatcher(whiteList map[string][]string) selector.MatchFunc {
	return func(_ context.Context, operation string) bool {
		for k, v := range whiteList {
			if strings.HasPrefix(operation, k) {
				return !lo.Contains(v, operation)
			}
		}
		return false
	}
}

// GetClientIP 获取客户端真实IP，优先从 X-Forwarded-For 和 X-Real-IP 获取，都不存在则从 RemoteAddr 获取
func GetClientIP(r *http.Request) string {
	// 尝试从 X-Forwarded-For 获取
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if ip != "" && !iputil.IsInternalIP(ip) {
				return ip
			}
		}
	}
	// 尝试从 X-Real-IP 获取
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" && !iputil.IsInternalIP(xRealIP) {
		return xRealIP
	}

	// 从 RemoteAddr 获取
	if r.RemoteAddr != "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil && !iputil.IsInternalIP(ip) {
			return ip
		}
	}
	return ""
}
