package config

import (
	"fmt"
	"os"
	"sort"
	"strings"

	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

const minJWTSecretLength = 32

type environment string

const (
	environmentLocal       environment = "local"
	environmentDevelopment environment = "development"
	environmentStaging     environment = "staging"
	environmentProduction  environment = "production"
)

// ValidationError reports all security configuration problems found at startup.
type ValidationError struct {
	Problems []string
}

func (e *ValidationError) Error() string {
	return "security configuration validation failed: " + strings.Join(e.Problems, "; ")
}

// ValidateSecurity validates unsafe defaults and required secrets before the app starts.
func ValidateSecurity(cfg *conf.Bootstrap) error {
	if cfg == nil {
		return &ValidationError{Problems: []string{"config is nil"}}
	}

	var problems []string
	env := validateEnvironment(cfg.GetEnv(), os.Getenv("APP_ENV"), &problems)

	validateJWTSecrets(cfg, &problems)
	if env == environmentProduction {
		validateProductionHTTP(cfg, &problems)
		validateProductionGRPC(cfg, &problems)
		validateProductionSensitiveConfig(cfg, &problems)
	}

	if len(problems) > 0 {
		return &ValidationError{Problems: uniqueProblems(problems)}
	}
	return nil
}

func validateEnvironment(configEnv, appEnv string, problems *[]string) environment {
	configEnv = strings.TrimSpace(configEnv)
	appEnv = strings.TrimSpace(appEnv)

	cfgEnv, cfgOK := normalizeEnvironment(configEnv)
	runtimeEnv, runtimeOK := normalizeEnvironment(appEnv)

	if configEnv == "" && appEnv == "" {
		*problems = append(*problems, "env must be set to local, development, staging, or production")
		return ""
	}
	if configEnv != "" && !cfgOK {
		*problems = append(*problems, fmt.Sprintf("env %q is unsupported; use local, development, staging, or production", configEnv))
	}
	if appEnv != "" && !runtimeOK {
		*problems = append(*problems, fmt.Sprintf("APP_ENV %q is unsupported; use local, development, staging, or production", appEnv))
	}
	if cfgOK && runtimeOK && cfgEnv != runtimeEnv {
		*problems = append(*problems, fmt.Sprintf("env %q does not match APP_ENV %q", configEnv, appEnv))
	}

	if cfgEnv == environmentProduction || runtimeEnv == environmentProduction {
		return environmentProduction
	}
	if runtimeOK {
		return runtimeEnv
	}
	if cfgOK {
		return cfgEnv
	}
	return ""
}

func normalizeEnvironment(value string) (environment, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "":
		return "", false
	case "local":
		return environmentLocal, true
	case "dev", "development":
		return environmentDevelopment, true
	case "stage", "staging", "test", "testing":
		return environmentStaging, true
	case "prod", "production":
		return environmentProduction, true
	default:
		return "", false
	}
}

func validateJWTSecrets(cfg *conf.Bootstrap, problems *[]string) {
	for _, audience := range []string{"admin", "kid", "parent"} {
		path := []string{"jwt", audience, "accessSecret"}
		name := "business." + strings.Join(path, ".")
		value, ok := businessString(cfg, path...)
		if !ok || strings.TrimSpace(value) == "" {
			*problems = append(*problems, name+" is required")
			continue
		}
		if len(value) < minJWTSecretLength {
			*problems = append(*problems, fmt.Sprintf("%s must be at least %d characters", name, minJWTSecretLength))
		}
		if isPlaceholderValue(value) {
			*problems = append(*problems, name+" must not use a placeholder value")
		}
	}
}

func validateProductionHTTP(cfg *conf.Bootstrap, problems *[]string) {
	httpCfg := cfg.GetServer().GetHttp()
	if httpCfg == nil {
		*problems = append(*problems, "server.http is required in production")
		return
	}

	if httpCfg.GetEnablePprof() {
		*problems = append(*problems, "server.http.enablePprof must be false in production")
	}

	middleware := httpCfg.GetMiddleware()
	if middleware == nil {
		*problems = append(*problems, "server.http.middleware is required in production")
	} else {
		requireEnabled("server.http.middleware.enableLogging", middleware.GetEnableLogging(), problems)
		requireEnabled("server.http.middleware.enableRecovery", middleware.GetEnableRecovery(), problems)
		requireEnabled("server.http.middleware.enableRateLimiter", middleware.GetEnableRateLimiter(), problems)
		requireEnabled("server.http.middleware.enableMetrics", middleware.GetEnableMetrics(), problems)

		if middleware.GetEnableRateLimiter() {
			limiter := middleware.GetLimiter()
			if limiter == nil {
				*problems = append(*problems, "server.http.middleware.limiter is required when rate limiter is enabled")
			} else {
				if limiter.GetWindow() == nil || limiter.GetWindow().AsDuration() <= 0 {
					*problems = append(*problems, "server.http.middleware.limiter.window must be greater than 0 in production")
				}
				if limiter.GetBucket() <= 0 {
					*problems = append(*problems, "server.http.middleware.limiter.bucket must be greater than 0 in production")
				}
				if limiter.GetCpuThreshold() <= 0 {
					*problems = append(*problems, "server.http.middleware.limiter.cpuThreshold must be greater than 0 in production")
				}
			}
		}
	}

	if httpCfg.GetEnableCors() {
		cors := httpCfg.GetCors()
		if cors == nil || len(cors.GetOrigins()) == 0 {
			*problems = append(*problems, "server.http.cors.origins must be explicitly configured in production")
			return
		}
		for _, origin := range cors.GetOrigins() {
			origin = strings.TrimSpace(origin)
			if origin == "" {
				*problems = append(*problems, "server.http.cors.origins must not contain empty origins in production")
				continue
			}
			if strings.Contains(origin, "*") {
				*problems = append(*problems, "server.http.cors.origins must not contain wildcards in production")
			}
		}
	}
}

func validateProductionGRPC(cfg *conf.Bootstrap, problems *[]string) {
	grpcCfg := cfg.GetServer().GetGrpc()
	if grpcCfg == nil {
		*problems = append(*problems, "server.grpc is required in production")
		return
	}

	middleware := grpcCfg.GetMiddleware()
	if middleware == nil {
		*problems = append(*problems, "server.grpc.middleware is required in production")
		return
	}
	requireEnabled("server.grpc.middleware.enableLogging", middleware.GetEnableLogging(), problems)
	requireEnabled("server.grpc.middleware.enableRecovery", middleware.GetEnableRecovery(), problems)
	requireEnabled("server.grpc.middleware.enableRateLimiter", middleware.GetEnableRateLimiter(), problems)
	requireEnabled("server.grpc.middleware.enableMetrics", middleware.GetEnableMetrics(), problems)
}

func validateProductionSensitiveConfig(cfg *conf.Bootstrap, problems *[]string) {
	requireProductionSecret("data.gorm.dataSourceName", cfg.GetData().GetGorm().GetDataSourceName(), problems)
	requireProductionSecret("data.redis.addr", cfg.GetData().GetRedis().GetAddr(), problems)
	requireProductionSecret("data.redis.password", cfg.GetData().GetRedis().GetPassword(), problems)
	if isInsecureProductionEndpoint(cfg.GetData().GetGorm().GetDataSourceName()) {
		*problems = append(*problems, "data.gorm.dataSourceName must not point to local or wildcard hosts in production")
	}
	if isInsecureProductionEndpoint(cfg.GetData().GetRedis().GetAddr()) {
		*problems = append(*problems, "data.redis.addr must not point to local or wildcard hosts in production")
	}
	if strings.Contains(strings.ToLower(cfg.GetData().GetGorm().GetDataSourceName()), "sslmode=disable") {
		*problems = append(*problems, "data.gorm.dataSourceName must not disable TLS in production")
	}

	for _, path := range [][]string{
		{"wx", "defaultGzhAppId"},
		{"wx", "defaultXcxAppId"},
		{"baiduPush", "apiKey"},
		{"baiduPush", "secretKey"},
	} {
		name := "business." + strings.Join(path, ".")
		value, ok := businessString(cfg, path...)
		if !ok {
			*problems = append(*problems, name+" is required in production")
			continue
		}
		requireProductionSecret(name, value, problems)
	}

	seenRequired := map[string]struct{}{
		"wx.defaultGzhAppId":      {},
		"wx.defaultXcxAppId":      {},
		"baiduPush.apiKey":        {},
		"baiduPush.secretKey":     {},
		"jwt.admin.accessSecret":  {},
		"jwt.kid.accessSecret":    {},
		"jwt.parent.accessSecret": {},
	}
	walkBusinessStrings(cfg.GetBusiness(), func(path, value string) {
		if _, ok := seenRequired[path]; ok {
			return
		}
		if isPlaceholderValue(value) {
			*problems = append(*problems, "business."+path+" must not use a placeholder value in production")
		}
	})
}

func requireEnabled(name string, enabled bool, problems *[]string) {
	if !enabled {
		*problems = append(*problems, name+" must be enabled in production")
	}
}

func requireProductionSecret(name, value string, problems *[]string) {
	value = strings.TrimSpace(value)
	if value == "" {
		*problems = append(*problems, name+" is required in production")
		return
	}
	if isPlaceholderValue(value) {
		*problems = append(*problems, name+" must not use a placeholder value in production")
	}
	if isNonProductionExampleValue(value) {
		*problems = append(*problems, name+" must not use local, dev, test, or sample values in production")
	}
}

func businessString(cfg *conf.Bootstrap, path ...string) (string, bool) {
	if cfg == nil || len(path) < 2 {
		return "", false
	}
	current := cfg.GetBusiness()[path[0]]
	if current == nil {
		return "", false
	}
	for i, key := range path[1:] {
		value := current.GetFields()[key]
		if value == nil {
			return "", false
		}
		if i == len(path)-2 {
			if stringValue, ok := value.GetKind().(*structpb.Value_StringValue); ok {
				return stringValue.StringValue, true
			}
			return "", false
		}
		current = value.GetStructValue()
		if current == nil {
			return "", false
		}
	}
	return "", false
}

func walkBusinessStrings(business map[string]*structpb.Struct, visit func(path, value string)) {
	keys := make([]string, 0, len(business))
	for key := range business {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		walkStructStrings(key, business[key], visit)
	}
}

func walkStructStrings(prefix string, value *structpb.Struct, visit func(path, value string)) {
	if value == nil {
		return
	}
	keys := make([]string, 0, len(value.GetFields()))
	for key := range value.GetFields() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		field := value.GetFields()[key]
		if field == nil {
			continue
		}
		path := prefix + "." + key
		switch kind := field.GetKind().(type) {
		case *structpb.Value_StringValue:
			visit(path, kind.StringValue)
		case *structpb.Value_StructValue:
			walkStructStrings(path, kind.StructValue, visit)
		case *structpb.Value_ListValue:
			walkListStrings(path, kind.ListValue, visit)
		}
	}
}

func walkListStrings(prefix string, value *structpb.ListValue, visit func(path, value string)) {
	if value == nil {
		return
	}
	for index, field := range value.GetValues() {
		if field == nil {
			continue
		}
		path := fmt.Sprintf("%s[%d]", prefix, index)
		switch kind := field.GetKind().(type) {
		case *structpb.Value_StringValue:
			visit(path, kind.StringValue)
		case *structpb.Value_StructValue:
			walkStructStrings(path, kind.StructValue, visit)
		case *structpb.Value_ListValue:
			walkListStrings(path, kind.ListValue, visit)
		}
	}
}

func isPlaceholderValue(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	if normalized == "" {
		return false
	}
	placeholderTokens := []string{
		"your_",
		"your-",
		"_here",
		"-here",
		"changeme",
		"change_me",
		"change-me",
		"placeholder",
		"replace_me",
		"replace-me",
		"<",
		">",
	}
	for _, token := range placeholderTokens {
		if strings.Contains(normalized, token) {
			return true
		}
	}
	return normalized == "todo" || normalized == "xxx" || normalized == "xxxx"
}

func isNonProductionExampleValue(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	unsafeTokens := []string{
		"dev-",
		"dev_",
		"local-",
		"local_",
		"test-",
		"test_",
		"dummy",
		"sample",
	}
	for _, token := range unsafeTokens {
		if strings.Contains(normalized, token) {
			return true
		}
	}
	return false
}

func isInsecureProductionEndpoint(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	for _, token := range []string{"host=0.0.0.0", "host=127.0.0.1", "host=localhost", "0.0.0.0:", "127.0.0.1:", "localhost:"} {
		if strings.Contains(normalized, token) {
			return true
		}
	}
	return false
}

func uniqueProblems(problems []string) []string {
	seen := make(map[string]struct{}, len(problems))
	unique := make([]string, 0, len(problems))
	for _, problem := range problems {
		if _, ok := seen[problem]; ok {
			continue
		}
		seen[problem] = struct{}{}
		unique = append(unique, problem)
	}
	return unique
}
