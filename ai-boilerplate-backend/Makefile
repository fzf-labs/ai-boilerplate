# Go Environment
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
# Database Configuration
DB_TYPE=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=ai_boilerplate
DB_DSN=host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable
DB_TABLES=
# Apifox Configuration
APIFOX_PROJECT_ID=xxxx
APIFOX_PROJECT_TOKEN=xxx
# Docker Configuration
DOCKER_REGISTRY=xxx
DOCKER_USERNAME=xxx
DOCKER_PASSWORD=xxx
DOCKER_NAMESPACE=ai-boilerplate
DOCKER_IMAGE_NAME=ai-boilerplate-backend
# API Testing Configuration
TEST_API_URL=http://localhost:8000
TEST_ADMIN_USER=admin
TEST_ADMIN_PASS=123456
TEST_LOGIN_PATH=/admin/v1/sys_auth/login


ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto -not -name apifox.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
	BUF_INSTALLED=$(shell command -v buf 2> /dev/null)
	GCI_INSTALLED=$(shell command -v gci 2> /dev/null)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: api
# generate api proto
api: buf gci
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   --validate_out=paths=source_relative,lang=go:./api \
		   --go-errors_out=paths=source_relative:./api \
		   --openapiv2_out ./doc/swagger \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   --openapiv2_opt fqn_for_openapi_name=true \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && GOPROXY="https://goproxy.cn,direct" go mod tidy && go build -ldflags "-w -s -X main.Version=$(VERSION)" -o ./bin/ ./cmd/...

.PHONY: docker-build
# docker build
docker-build:
	@./scripts/docker-build.sh

.PHONY: git-new-tag
# 创建新的git标签
git-new-tag:
	@./scripts/git-new-tag.sh

.PHONY: new-release-branch
# 创建发布分支
new-release-branch:
	@./scripts/new-release-branch.sh

.PHONY: wire
# wire
wire:
	wire ./...

.PHONY: run
# run
run:
	@export APP_ENV=development && kratos run

.PHONY: lint
# golang lint 检查
lint:
	@golangci-lint run --config .golangci.yml ./... -v

.PHONY: sqldump
# 导出sql文件
sqldump:
	@godb sqldump --db $(DB_TYPE) --dsn "$(DB_DSN)" --tables "$(DB_TABLES)" -f true

.PHONY: sqlbackup
# 数据库备份 (使用 MODE=schema 仅结构, MODE=data 仅数据, 默认完整备份)
sqlbackup:
	@./scripts/sql-backup.sh $(if $(filter schema,$(MODE)),--schema-only,$(if $(filter data,$(MODE)),--data-only,))


.PHONY: sqltopb
# sql转为pb，需要传入位置参数: make sqltopb admin table1,table2
sqltopb:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		echo "错误: 必须指定 POSITION 参数 (admin/app)"; \
		echo "使用方法: make sqltopb admin table1,table2"; \
		exit 1; \
	fi
	@if [ -z "$(word 3,$(MAKECMDGOALS))" ]; then \
		echo "错误: 必须指定 TABLES 参数"; \
		echo "使用方法: make sqltopb admin table1,table2"; \
		exit 1; \
	fi
	@POSITION=$(word 2,$(MAKECMDGOALS)); \
	TABLES=$(word 3,$(MAKECMDGOALS)); \
	godb sqltopb --db "$(DB_TYPE)" --dsn "$(DB_DSN)" --tables "$$TABLES" -p "$$POSITION.v1" -g "github.com/fzf-labs/ai-boilerplate-backend/api/$$POSITION/v1;v1" -o "./api/$$POSITION/v1"
%: # 防止位置参数被当作目标处理
	@:

.PHONY: gorm
# 生成 GORM 数据库代码
gorm:
	@godb ormgen --db "$(DB_TYPE)" --dsn "$(DB_DSN)" --tables "$(DB_TABLES)"

.PHONY: pbtocode
# pb转为代码
pbtocode:
	@kratos-gen data --db "$(DB_TYPE)" --dsn "$(DB_DSN)" --tables "$(DB_TABLES)" --partitionTable true
	@kratos-gen service

.PHONY: apidoc
# 同步接口文档
apidoc:
	@api-import swagger apifox -t "$(APIFOX_PROJECT_ID)" -p "$(APIFOX_PROJECT_TOKEN)" -i ./doc/swagger

.PHONY: pbdoc
# sql转为pb
pbdoc:
	@rm -rf ./doc/pb/*
	@godb sqltopb --db "$(DB_TYPE)" --dsn "$(DB_DSN)" -p 'admin.v1' -g 'github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1;v1' -o './doc/pb'

.PHONY: doc
# 生成文档
doc:
	make sqldump
	make sqlbackup
	make pbdoc
	make apidoc

.PHONY: buf
# buf 格式化 proto
buf:
	@if [ -n "$(BUF_INSTALLED)" ]; then \
        cd ./api  && \
        buf format -w && \
        echo "proto format finish"; \
    else \
        echo "please installation buf: https://buf.build/docs/installation"; \
    fi

.PHONY: gci
# buf 格式化 proto
gci:
	@if [ -n "$(GCI_INSTALLED)" ]; then \
        gci write ./internal --skip-generated && \
        echo "gci format finish"; \
    else \
        echo "please installation gci: https://github.com/daixiang0/gci"; \
    fi

# ============ API Schema 契约测试 ============

.PHONY: api-schema-test
# API Schema 契约测试 (用法: make api-schema-test [FILE=xxx.swagger.json] [METHOD=GET])
api-schema-test:
	@./scripts/api-schema-test.sh $(if $(METHOD),-m $(METHOD),) $(FILE)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
