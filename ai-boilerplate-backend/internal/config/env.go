package config

import (
	"os"
	"strings"

	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

const envPrefix = "AI_BOILERPLATE_"

// ApplyEnvironmentOverrides injects deployment secrets without committing real config files.
func ApplyEnvironmentOverrides(cfg *conf.Bootstrap) {
	if cfg == nil {
		return
	}

	setStringFromEnv(envPrefix+"DB_DSN", func(value string) {
		ensureData(cfg).Gorm.DataSourceName = value
	})
	setStringFromEnv(envPrefix+"REDIS_ADDR", func(value string) {
		ensureData(cfg).Redis.Addr = value
	})
	setStringFromEnv(envPrefix+"REDIS_PASSWORD", func(value string) {
		ensureData(cfg).Redis.Password = value
	})
	setStringFromEnv(envPrefix+"JWT_ADMIN_SECRET", func(value string) {
		setBusinessStringValue(cfg, []string{"jwt", "admin", "accessSecret"}, value)
	})
	setStringFromEnv(envPrefix+"JWT_KID_SECRET", func(value string) {
		setBusinessStringValue(cfg, []string{"jwt", "kid", "accessSecret"}, value)
	})
	setStringFromEnv(envPrefix+"JWT_PARENT_SECRET", func(value string) {
		setBusinessStringValue(cfg, []string{"jwt", "parent", "accessSecret"}, value)
	})
	setStringFromEnv(envPrefix+"WX_DEFAULT_GZH_APP_ID", func(value string) {
		setBusinessStringValue(cfg, []string{"wx", "defaultGzhAppId"}, value)
	})
	setStringFromEnv(envPrefix+"WX_DEFAULT_XCX_APP_ID", func(value string) {
		setBusinessStringValue(cfg, []string{"wx", "defaultXcxAppId"}, value)
	})
	setStringFromEnv(envPrefix+"BAIDU_PUSH_API_KEY", func(value string) {
		setBusinessStringValue(cfg, []string{"baiduPush", "apiKey"}, value)
	})
	setStringFromEnv(envPrefix+"BAIDU_PUSH_SECRET_KEY", func(value string) {
		setBusinessStringValue(cfg, []string{"baiduPush", "secretKey"}, value)
	})
	setStringFromEnv(envPrefix+"CORS_ORIGINS", func(value string) {
		origins := splitCommaSeparated(value)
		if len(origins) == 0 {
			return
		}
		ensureHTTP(cfg).Cors.Origins = origins
	})
}

func setStringFromEnv(name string, set func(string)) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return
	}
	set(value)
}

func splitCommaSeparated(value string) []string {
	parts := strings.Split(value, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			values = append(values, part)
		}
	}
	return values
}

func ensureData(cfg *conf.Bootstrap) *conf.Data {
	if cfg.Data == nil {
		cfg.Data = &conf.Data{}
	}
	if cfg.Data.Gorm == nil {
		cfg.Data.Gorm = &conf.Data_Gorm{}
	}
	if cfg.Data.Redis == nil {
		cfg.Data.Redis = &conf.Data_Redis{}
	}
	return cfg.Data
}

func ensureHTTP(cfg *conf.Bootstrap) *conf.Server_HTTP {
	if cfg.Server == nil {
		cfg.Server = &conf.Server{}
	}
	if cfg.Server.Http == nil {
		cfg.Server.Http = &conf.Server_HTTP{}
	}
	if cfg.Server.Http.Cors == nil {
		cfg.Server.Http.Cors = &conf.Server_HTTP_CORS{}
	}
	return cfg.Server.Http
}

func setBusinessStringValue(cfg *conf.Bootstrap, path []string, value string) {
	if len(path) < 2 {
		return
	}
	if cfg.Business == nil {
		cfg.Business = map[string]*structpb.Struct{}
	}

	current := cfg.Business[path[0]]
	if current == nil {
		current = &structpb.Struct{Fields: map[string]*structpb.Value{}}
		cfg.Business[path[0]] = current
	}
	for _, key := range path[1 : len(path)-1] {
		fields := ensureStructFields(current)
		nextValue := fields[key]
		next := nextValue.GetStructValue()
		if next == nil {
			next = &structpb.Struct{Fields: map[string]*structpb.Value{}}
			fields[key] = structpb.NewStructValue(next)
		}
		current = next
	}
	ensureStructFields(current)[path[len(path)-1]] = structpb.NewStringValue(value)
}

func ensureStructFields(value *structpb.Struct) map[string]*structpb.Value {
	if value.Fields == nil {
		value.Fields = map[string]*structpb.Value{}
	}
	return value.Fields
}
