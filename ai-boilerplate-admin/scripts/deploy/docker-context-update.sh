#!/bin/bash

# Docker 远程更新脚本 - 简化版
# 功能：通过 Docker Context 连接远程 Docker，拉取最新镜像并重启容器

set -e  # 出错时立即退出

# 加载配置函数
load_config() {
    local script_dir
    script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    
    # 配置文件优先级（从高到低）：
    # 1. 环境变量（优先级最高）
    # 2. .docker-config.local（本地配置）
    # 3. .docker-config（默认配置）
    
    local config_files=(
        "${script_dir}/.docker-config.local"
        "${script_dir}/.docker-config"
    )
    
    local config_loaded=false
    for config_file in "${config_files[@]}"; do
        if [ -f "$config_file" ]; then
            # shellcheck source=/dev/null
            source "$config_file"
            config_loaded=true
            break
        fi
    done
    
    if [ "$config_loaded" = false ]; then
        echo "警告: 未找到配置文件，请创建 ${script_dir}/.docker-config"
        echo "可以复制 ${script_dir}/.docker-config.example 进行配置"
        return 1
    fi
}

# 加载配置
load_config || exit 1

# 允许环境变量覆盖配置文件中的值
DOCKER_CONTEXT_PREFIX="${DOCKER_CONTEXT_PREFIX}"
DOCKER_REGISTRY="${DOCKER_REGISTRY}"
DOCKER_NAMESPACE="${DOCKER_NAMESPACE}"
DOCKER_IMAGE_NAME="${DOCKER_IMAGE_NAME}"
DOCKER_USERNAME="${DOCKER_USERNAME}"
DOCKER_PASSWORD="${DOCKER_PASSWORD}"

# 颜色输出
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m'

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO] $1${NC}"
}

log_success() {
    echo -e "${GREEN}[SUCCESS] $1${NC}"
}

log_error() {
    echo -e "${RED}[ERROR] $1${NC}"
    exit 1
}

# 显示帮助信息
show_help() {
    cat << EOF
Docker 远程更新脚本

用法: $0 [--version <版本号>]

选项:
    --version VERSION   指定版本号，不指定则根据当前分支自动获取

示例:
    # 自动根据当前分支更新（最简单）
    $0
    
    # 更新到指定版本
    $0 --version 324b8a3

分支与 Context 自动映射:
    development 分支 → ${DOCKER_CONTEXT_PREFIX}-dev  (开发环境)
    testing 分支     → ${DOCKER_CONTEXT_PREFIX}-test (测试环境)
    master 分支      → ${DOCKER_CONTEXT_PREFIX}-prod (生产环境)

版本号获取规则:
    - 如果指定 --version，使用指定的版本号
    - 如果未指定，则根据当前 git 分支自动获取:
      * development/testing 分支: 使用当前 commit_id (短格式)
      * master 分支: 使用最新的 git tag

EOF
}

# 检测当前分支
get_current_branch() {
    git rev-parse --abbrev-ref HEAD 2>/dev/null || echo ""
}

# 根据分支检测环境
detect_environment_from_branch() {
    local branch="$1"
    case "$branch" in
        "development")
            echo "development"
            ;;
        "testing")
            echo "testing"
            ;;
        "master")
            echo "production"
            ;;
        *)
            echo ""
            ;;
    esac
}

# 根据分支推断 Docker Context
infer_context_from_branch() {
    local branch="$1"
    case "$branch" in
        "development")
            echo "${DOCKER_CONTEXT_PREFIX}-dev"
            ;;
        "testing")
            echo "${DOCKER_CONTEXT_PREFIX}-test"
            ;;
        "master")
            echo "${DOCKER_CONTEXT_PREFIX}-prod"
            ;;
        *)
            echo ""
            ;;
    esac
}

# 根据当前分支获取版本号
get_version_from_branch() {
    local branch="$1"
    case "$branch" in
        "development"|"testing")
            # 开发和测试环境使用 commit_id
            git rev-parse --short HEAD 2>/dev/null || echo ""
            ;;
        "master")
            # 生产环境使用 git tag
            git describe --tags --always 2>/dev/null || echo ""
            ;;
        *)
            echo ""
            ;;
    esac
}

# 构建镜像标签
build_image_tag() {
    local env="$1"
    local version="$2"
    if [ -z "$version" ]; then
        echo "${DOCKER_REGISTRY}/${DOCKER_NAMESPACE}/${DOCKER_IMAGE_NAME}:${env}-latest"
    else
        echo "${DOCKER_REGISTRY}/${DOCKER_NAMESPACE}/${DOCKER_IMAGE_NAME}:${env}-${version}"
    fi
}

# 切换回本地 Docker context
switch_to_local_context() {
    local local_contexts=("default" "desktop-linux" "orbstack")
    
    for context in "${local_contexts[@]}"; do
        if docker context ls --format '{{.Name}}' 2>/dev/null | grep -q "^${context}$"; then
            if docker context use "$context" >/dev/null 2>&1; then
                log_info "已切换回本地 Docker context: ${context}"
                return 0
            fi
        fi
    done
    
    # 如果所有本地 context 都失败，至少尝试 default
    docker context use default >/dev/null 2>&1 || true
}

# 主函数
main() {
    local version=""
    
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
            --version)
                version="$2"
                shift 2
                ;;
            --help|-h)
                show_help
                exit 0
                ;;
            *)
                log_error "未知参数: $1。使用 --help 查看帮助信息"
                ;;
        esac
    done
    
    # 获取当前分支
    local current_branch
    current_branch=$(get_current_branch)
    if [ -z "$current_branch" ]; then
        log_error "无法获取当前 git 分支，请确保在 git 仓库中运行"
    fi
    log_info "当前分支: ${current_branch}"
    
    # 根据分支自动推断 Docker Context
    local context_name
    context_name=$(infer_context_from_branch "$current_branch")
    if [ -z "$context_name" ]; then
        log_error "无法根据分支 '${current_branch}' 推断 Docker Context，当前分支必须是 development/testing/master"
    fi
    log_info "根据分支自动选择 Docker Context: ${context_name}"
    
    # 根据分支推断环境
    local env
    env=$(detect_environment_from_branch "$current_branch")
    log_info "根据当前分支推断环境: ${env}"
    
    # 如果未指定版本号，尝试从分支自动获取
    if [ -z "$version" ] && [ -n "$current_branch" ]; then
        local auto_version
        auto_version=$(get_version_from_branch "$current_branch")
        if [ -n "$auto_version" ]; then
            version="$auto_version"
            log_info "根据当前分支自动获取版本号: ${version}"
        fi
    fi
    
    # 构建镜像标签
    local image_tag
    image_tag=$(build_image_tag "$env" "$version")
    
    log_info "=========================================="
    log_info "项目: ${DOCKER_IMAGE_NAME}"
    log_info "Context: ${context_name}"
    log_info "环境: ${env}"
    log_info "镜像: ${image_tag}"
    log_info "=========================================="
    
    # 保存当前 context
    local previous_context
    previous_context=$(docker context show)
    
    # 切换到目标 context
    log_info "切换到 Docker Context: ${context_name}"
    docker context use "$context_name" >/dev/null || log_error "切换 Context 失败"
    
    # 确保退出时切换回本地 Docker context
    # 优先使用 default，如果失败则尝试其他本地 context
    trap 'switch_to_local_context' EXIT
    
    # 检查远程连接
    log_info "检查远程 Docker 连接..."
    docker info >/dev/null || log_error "无法连接到远程 Docker"
    
    # 登录镜像仓库
    log_info "登录容器镜像服务..."
    echo "$DOCKER_PASSWORD" | docker login "$DOCKER_REGISTRY" \
        --username="$DOCKER_USERNAME" --password-stdin >/dev/null || log_error "Docker 登录失败"
    
    # 拉取镜像
    log_info "拉取镜像: ${image_tag}"
    docker pull "$image_tag" || log_error "拉取镜像失败"
    
    # 检查是否有旧容器在运行
    if docker ps -a --format '{{.Names}}' | grep -q "^${DOCKER_IMAGE_NAME}$"; then
        log_info "停止并删除旧容器..."
        docker stop "$DOCKER_IMAGE_NAME" 2>/dev/null || true
        docker rm "$DOCKER_IMAGE_NAME" 2>/dev/null || true
    fi
    
    # 使用 docker run 启动新容器
    log_info "启动新容器..."
    docker run -d \
        --name "$DOCKER_IMAGE_NAME" \
        --network green-orange-guard_guard \
        --restart always \
        -p 7000:8080 \
        "$image_tag" \
        || log_error "容器启动失败"
    
    log_success "容器启动成功"
    
    # 等待服务启动
    sleep 3
    
    # 检查服务状态
    log_info "检查服务状态..."
    if docker ps --format '{{.Names}}' | grep -q "^${DOCKER_IMAGE_NAME}$"; then
        log_success "服务 ${DOCKER_IMAGE_NAME} 更新成功！"
    else
        log_error "服务 ${DOCKER_IMAGE_NAME} 未正常运行"
    fi
    
    # 显示服务状态
    echo ""
    docker ps --filter "name=${DOCKER_IMAGE_NAME}" --format "table {{.ID}}\t{{.Names}}\t{{.Status}}\t{{.Image}}"
    
    # 切换回本地 Docker context（EXIT trap 会自动执行）
    
    log_success "=========================================="
    log_success "更新完成！"
    log_success "=========================================="
}

# 运行主函数
main "$@"
