package config

import (
	"strings"
	"testing"
	"time"

	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestValidateSecurityRejectsProductionUnsafeDefaults(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := validProductionConfig(t)
	cfg.Server.Http.EnablePprof = true
	cfg.Server.Http.Cors.Origins = []string{"*"}
	cfg.Server.Http.Middleware.EnableRateLimiter = false
	cfg.Server.Http.Middleware.EnableMetrics = false
	setBusinessStringValue(cfg, []string{"jwt", "admin", "accessSecret"}, "your_jwt_admin_secret_here_min_32_chars")
	cfg.Data.Redis.Password = "your_redis_password_here"

	err := ValidateSecurity(cfg)
	if err == nil {
		t.Fatal("expected production validation error")
	}

	message := err.Error()
	for _, want := range []string{
		"server.http.enablePprof must be false in production",
		"server.http.cors.origins must not contain wildcards in production",
		"server.http.middleware.enableRateLimiter must be enabled in production",
		"server.http.middleware.enableMetrics must be enabled in production",
		"business.jwt.admin.accessSecret must not use a placeholder value",
		"data.redis.password must not use a placeholder value in production",
	} {
		if !strings.Contains(message, want) {
			t.Fatalf("expected error to contain %q, got %q", want, message)
		}
	}
}

func TestValidateSecurityRejectsProductionLocalDataEndpoints(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := validProductionConfig(t)
	cfg.Data.Gorm.DataSourceName = "host=0.0.0.0 port=5432 user=postgres password=prod-password dbname=ai_boilerplate sslmode=disable"
	cfg.Data.Redis.Addr = "127.0.0.1:6379"

	err := ValidateSecurity(cfg)
	if err == nil {
		t.Fatal("expected local endpoint validation error")
	}

	message := err.Error()
	for _, want := range []string{
		"data.gorm.dataSourceName must not point to local or wildcard hosts in production",
		"data.gorm.dataSourceName must not disable TLS in production",
		"data.redis.addr must not point to local or wildcard hosts in production",
	} {
		if !strings.Contains(message, want) {
			t.Fatalf("expected error to contain %q, got %q", want, message)
		}
	}
}

func TestValidateSecurityRejectsShortJWTSecretInDevelopment(t *testing.T) {
	t.Setenv("APP_ENV", "development")

	cfg := validDevelopmentConfig(t)
	setBusinessStringValue(cfg, []string{"jwt", "kid", "accessSecret"}, "too-short")

	err := ValidateSecurity(cfg)
	if err == nil {
		t.Fatal("expected JWT secret length validation error")
	}
	if !strings.Contains(err.Error(), "business.jwt.kid.accessSecret must be at least 32 characters") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateSecurityRejectsMissingProductionJWTSecret(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := validProductionConfig(t)
	delete(cfg.Business["jwt"].GetFields()["parent"].GetStructValue().GetFields(), "accessSecret")

	err := ValidateSecurity(cfg)
	if err == nil {
		t.Fatal("expected missing JWT secret validation error")
	}
	if !strings.Contains(err.Error(), "business.jwt.parent.accessSecret is required") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateSecurityAllowsDevelopmentWideCorsAndPprof(t *testing.T) {
	t.Setenv("APP_ENV", "development")

	cfg := validDevelopmentConfig(t)
	cfg.Server.Http.EnablePprof = true
	cfg.Server.Http.Cors.Origins = []string{"*"}
	cfg.Server.Http.Middleware.EnableRateLimiter = false
	cfg.Server.Http.Middleware.EnableMetrics = false

	if err := ValidateSecurity(cfg); err != nil {
		t.Fatalf("expected development config to pass, got %v", err)
	}
}

func TestValidateSecurityAllowsHardenedProductionConfig(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := validProductionConfig(t)
	if err := ValidateSecurity(cfg); err != nil {
		t.Fatalf("expected production config to pass, got %v", err)
	}
}

func TestValidateSecurityRejectsEnvironmentMismatch(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	cfg := validProductionConfig(t)
	cfg.Env = "development"

	err := ValidateSecurity(cfg)
	if err == nil {
		t.Fatal("expected environment mismatch error")
	}
	if !strings.Contains(err.Error(), `env "development" does not match APP_ENV "production"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestApplyEnvironmentOverridesInjectsSecrets(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_DB_DSN", "host=db.internal port=5432 user=app password=from-env dbname=ai_boilerplate sslmode=require")
	t.Setenv("AI_BOILERPLATE_REDIS_ADDR", "redis.internal:6379")
	t.Setenv("AI_BOILERPLATE_REDIS_PASSWORD", "redis-from-env")
	t.Setenv("AI_BOILERPLATE_JWT_ADMIN_SECRET", strings.Repeat("x", 32))
	t.Setenv("AI_BOILERPLATE_CORS_ORIGINS", "https://admin.example.com, https://app.example.com")

	cfg := validProductionConfig(t)
	cfg.Data.Gorm.DataSourceName = "host=0.0.0.0 port=5432 user=postgres password=your_db_password_here dbname=ai_boilerplate sslmode=disable"
	cfg.Data.Redis.Addr = "0.0.0.0:6379"
	cfg.Data.Redis.Password = "your_redis_password_here"
	setBusinessStringValue(cfg, []string{"jwt", "admin", "accessSecret"}, "your_jwt_admin_secret_here_min_32_chars")

	ApplyEnvironmentOverrides(cfg)

	if got := cfg.Data.Gorm.DataSourceName; !strings.Contains(got, "from-env") {
		t.Fatalf("expected DB DSN override, got %q", got)
	}
	if got := cfg.Data.Redis.Addr; got != "redis.internal:6379" {
		t.Fatalf("expected Redis addr override, got %q", got)
	}
	adminSecret, ok := businessString(cfg, "jwt", "admin", "accessSecret")
	if !ok || adminSecret != strings.Repeat("x", 32) {
		t.Fatalf("expected admin JWT secret override, got %q", adminSecret)
	}
	if got := cfg.Server.Http.Cors.Origins; len(got) != 2 || got[1] != "https://app.example.com" {
		t.Fatalf("expected CORS origins override, got %#v", got)
	}
}

func validDevelopmentConfig(t *testing.T) *conf.Bootstrap {
	t.Helper()

	cfg := validProductionConfig(t)
	cfg.Env = "development"
	cfg.Server.Http.EnablePprof = true
	cfg.Server.Http.Cors.Origins = []string{"*"}
	return cfg
}

func validProductionConfig(t *testing.T) *conf.Bootstrap {
	t.Helper()

	business, err := structpb.NewStruct(map[string]any{
		"wx": map[string]any{
			"defaultGzhAppId": "wx-prod-gzh-app-id",
			"defaultXcxAppId": "wx-prod-xcx-app-id",
		},
		"jwt": map[string]any{
			"admin": map[string]any{
				"accessSecret": strings.Repeat("a", 32),
			},
			"kid": map[string]any{
				"accessSecret": strings.Repeat("b", 32),
			},
			"parent": map[string]any{
				"accessSecret": strings.Repeat("c", 32),
			},
		},
		"baiduPush": map[string]any{
			"apiKey":    "prod-baidu-push-api-key",
			"secretKey": "prod-baidu-push-secret-key",
		},
	})
	if err != nil {
		t.Fatalf("build business config: %v", err)
	}

	return &conf.Bootstrap{
		Name: "ai-boilerplate-backend",
		Env:  "production",
		Server: &conf.Server{
			Http: &conf.Server_HTTP{
				EnableCors:  true,
				EnablePprof: false,
				Middleware: &conf.Middleware{
					EnableLogging:     true,
					EnableRecovery:    true,
					EnableRateLimiter: true,
					EnableMetrics:     true,
					Limiter: &conf.RateLimiter{
						Window:       durationpb.New(time.Second),
						Bucket:       100,
						CpuThreshold: 80,
					},
				},
				Cors: &conf.Server_HTTP_CORS{
					Origins: []string{"https://admin.example.com"},
				},
			},
			Grpc: &conf.Server_GRPC{
				Middleware: &conf.Middleware{
					EnableLogging:     true,
					EnableRecovery:    true,
					EnableRateLimiter: true,
					EnableMetrics:     true,
				},
			},
		},
		Data: &conf.Data{
			Gorm: &conf.Data_Gorm{
				DataSourceName: "host=db.internal port=5432 user=app password=prod-db-password dbname=ai_boilerplate sslmode=require",
			},
			Redis: &conf.Data_Redis{
				Addr:     "redis.internal:6379",
				Password: "prod-redis-password",
			},
		},
		Business: map[string]*structpb.Struct{
			"wx":        business.GetFields()["wx"].GetStructValue(),
			"jwt":       business.GetFields()["jwt"].GetStructValue(),
			"baiduPush": business.GetFields()["baiduPush"].GetStructValue(),
		},
	}
}
