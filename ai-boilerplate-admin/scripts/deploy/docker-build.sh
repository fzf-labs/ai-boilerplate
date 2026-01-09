#!/bin/bash

# Docker 镜像构建脚本
# 支持按环境区分构建，自动检测分支和版本策略

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
DOCKER_REGISTRY="${DOCKER_REGISTRY}"
DOCKER_NAMESPACE="${DOCKER_NAMESPACE}"
DOCKER_IMAGE_NAME="${DOCKER_IMAGE_NAME}"
DOCKER_USERNAME="${DOCKER_USERNAME}"
DOCKER_PASSWORD="${DOCKER_PASSWORD}"
DOCKERFILE_PATH="${DOCKERFILE_PATH}"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

log_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

log_error() {
    echo -e "${RED}❌ $1${NC}"
    exit 1
}

# 显示帮助信息
show_help() {
    cat << EOF
Docker 镜像构建脚本

用法:
    $0 [环境]

环境:
    development     构建开发环境镜像 (需在 development 分支)
    testing         构建测试环境镜像 (需在 testing 分支)
    production      构建生产环境镜像 (需在 master 分支)
    
如果不指定环境，脚本会自动根据当前分支确定对应的环境

选项:
    --help, -h      显示此帮助信息

示例:
    $0                          # 自动检测当前分支并构建对应环境
    $0 development              # 强制构建开发环境镜像
    $0 testing                  # 强制构建测试环境镜像
    $0 production               # 强制构建生产环境镜像

分支与环境自动映射:
    development 分支 → development 环境 (使用 commit_id)
    testing 分支     → testing 环境     (使用 commit_id)
    master 分支      → production 环境  (使用 git tag)

注意:
    - 脚本会自动验证当前分支与目标环境是否匹配
    - 生产环境构建前会检查代码提交和推送状态
    - 生产环境会自动创建新的 git tag (可选)
EOF
}

# 检测当前分支
get_current_branch() {
    git rev-parse --abbrev-ref HEAD
}

# 检测环境
detect_environment() {
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
            log_error "当前分支 '$branch' 不是有效的环境分支 (development/testing/master)"
            ;;
    esac
}

# 获取版本号
get_version() {
    local env="$1"
    if [ "$env" = "production" ]; then
        # 生产环境使用 git tag
        git describe --tags --always
    else
        # 开发和测试环境使用 commit_id
        git rev-parse --short HEAD
    fi
}

# 验证分支和环境匹配
validate_branch_env() {
    local current_branch="$1"
    local target_env="$2"
    
    case "$target_env" in
        "development")
            if [ "$current_branch" != "development" ]; then
                log_error "当前分支不是 development，无法构建 development 环境镜像"
            fi
            ;;
        "testing")
            if [ "$current_branch" != "testing" ]; then
                log_error "当前分支不是 testing，无法构建 testing 环境镜像"
            fi
            ;;
        "production")
            if [ "$current_branch" != "master" ]; then
                log_error "当前分支不是 master，无法构建 production 环境镜像"
            fi
            ;;
    esac
}

# 检查生产环境的代码状态
check_production_ready() {
    # 检测当前是否有未提交的代码
    if [ -n "$(git status --porcelain)" ]; then
        log_error "当前有未提交的代码，请先提交代码"
    fi
    
    # 检测当前是否有未推送的代码
    if [ -n "$(git cherry -v 2>/dev/null)" ]; then
        log_error "当前有未推送的代码，请先推送代码"
    fi
}

# 创建新的 git tag (仅生产环境)
create_git_tag() {
    # 获取当前最新的 tag
    local latest_tag=$(git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 1)
    
    # 如果没有找到符合规范的tag，提示用户
    if [ -z "$latest_tag" ]; then
        log_error "未找到符合版本规范(v1.2.3)的tag，请先手动创建一个初始版本tag"
    fi
    
    log_info "当前最新的 tag: ${latest_tag}"
    
    # 检查当前代码是否比最新 tag 更新
    if [ "$(git rev-list ${latest_tag}..HEAD --count)" -eq "0" ]; then
        log_warning "当前代码与最新 tag ${latest_tag} 相同，将使用现有tag版本"
        return
    fi
    
    log_info "检测到 $(git rev-list ${latest_tag}..HEAD --count) 个新提交，可以创建新版本"
    
    # 提取版本号并验证格式
    if [[ ! "$latest_tag" =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
        log_error "最新tag ${latest_tag} 格式不正确，应为 v1.2.3 格式"
    fi
    
    local major="${BASH_REMATCH[1]}"
    local minor="${BASH_REMATCH[2]}"
    local patch="${BASH_REMATCH[3]}"
    
    # 生成新的版本号 (自动递增补丁版本)
    local new_patch=$((patch + 1))
    local new_version="v${major}.${minor}.${new_patch}"
    log_info "新的版本号: ${new_version}"
    
    # 自动创建新tag
    log_info "自动创建新tag ${new_version}"
    
    # 创建新的 tag
    if ! git tag -a "${new_version}" -m "release ${new_version}"; then
        log_error "创建tag失败"
    fi
    
    if ! git push origin "${new_version}"; then
        log_error "推送tag失败"
    fi
    
    log_success "创建并推送新 tag: ${new_version}"
}

# 确保使用本地 Docker
ensure_local_docker() {
    log_info "检查 Docker context..."
    
    # 获取当前 Docker context
    local current_context
    current_context=$(docker context show 2>/dev/null || echo "default")
    
    # 检查是否是本地 context（default、desktop-linux 或 orbstack）
    case "$current_context" in
        default|desktop-linux|orbstack)
            log_info "当前使用本地 Docker context: ${current_context}"
            ;;
        *)
            log_warning "当前 Docker context 为 '${current_context}'，切换到本地 Docker..."
            # 尝试切换到 default context
            if ! docker context use default >/dev/null 2>&1; then
                log_error "无法切换到本地 Docker context，请手动执行: docker context use default"
            fi
            log_success "已切换到本地 Docker context"
            ;;
    esac
}

# Docker 登录
docker_login() {
    log_info "登录容器镜像服务..."
    
    # 检查Docker守护进程是否运行
    if ! docker info >/dev/null 2>&1; then
        log_error "Docker 守护进程未运行，请启动 Docker"
    fi
    
    # 检查登录凭证是否设置
    if [ -z "$DOCKER_USERNAME" ] || [ -z "$DOCKER_PASSWORD" ]; then
        log_error "Docker 登录凭证未设置，请检查 DOCKER_USERNAME 和 DOCKER_PASSWORD"
    fi
    
    # 尝试登录
    if ! echo "$DOCKER_PASSWORD" | docker login "$DOCKER_REGISTRY" --username="$DOCKER_USERNAME" --password-stdin; then
        log_error "Docker 登录失败，请检查凭证和网络连接"
    fi
    
    log_success "Docker 登录成功"
}

# 构建 Docker 镜像
build_docker_image() {
    local version="$1"
    local env="$2"
    local image_tag="${DOCKER_IMAGE_NAME}:${env}-${version}"
    local remote_tag="${DOCKER_REGISTRY}/${DOCKER_NAMESPACE}/${image_tag}"
    
    # 使用配置的 DOCKERFILE_PATH 拼接完整路径
    local dockerfile_path="${DOCKERFILE_PATH}/Dockerfile"
    
    # 检查Dockerfile是否存在
    if [ ! -f "$dockerfile_path" ]; then
        log_error "Dockerfile 不存在: ${dockerfile_path}"
    fi
    
    log_info "构建 Docker 镜像 ${image_tag}..."
    
    # 设置清理变量
    export IMAGE_TAG_TO_CLEAN="${image_tag}"
    
    # 获取构建命令
    local build_command=""
    case "$env" in
        "development")
            build_command="build:antd"
            ;;
        "testing")
            build_command="build-testing:antd"
            ;;
        "production")
            build_command="build:antd"
            ;;
    esac
    
    # 获取项目根目录（假设脚本在 scripts/deploy 目录下）
    local project_root="$(cd "${script_dir}/../.." && pwd)"
    
    # 构建镜像，添加更多构建参数和错误处理
    if ! docker buildx build \
        --platform linux/amd64 \
        --build-arg APP_VERSION="${version}" \
        --build-arg APP_ENV="${env}" \
        --build-arg BUILD_COMMAND="${build_command}" \
        -f "${dockerfile_path}" \
        -t "${image_tag}" \
        --load \
        "${project_root}"; then
        log_error "Docker 镜像构建失败"
    fi
    
    # 构建成功后清除清理变量
    unset IMAGE_TAG_TO_CLEAN
    
    log_info "标记镜像为 ${remote_tag}..."
    if ! docker tag "${image_tag}" "${remote_tag}"; then
        log_error "镜像标记失败"
    fi
    
    log_info "推送镜像到远程仓库..."
    if ! docker push "${remote_tag}"; then
        log_error "镜像推送失败"
    fi
    
    log_success "${env} 环境镜像构建并推送成功: ${remote_tag}"
    
    # 显示镜像信息
    log_info "镜像详情:"
    echo "  本地标签: ${image_tag}" >&2
    echo "  远程标签: ${remote_tag}" >&2
    echo "  镜像大小: $(docker images --format "table {{.Size}}" "${image_tag}" | tail -n +2)" >&2
}


# 验证输入参数
validate_inputs() {
    local env="$1"
    
    # 验证环境参数
    case "$env" in
        development|testing|production)
            # 有效环境
            ;;
        *)
            log_error "无效的环境参数: $env"
            ;;
    esac
}

# 清理临时文件和镜像
cleanup() {
    local exit_code=$?
    
    # 如果构建失败，清理可能创建的镜像
    if [ $exit_code -ne 0 ] && [ -n "$IMAGE_TAG_TO_CLEAN" ]; then
        log_warning "构建失败，清理临时镜像..."
        docker rmi "$IMAGE_TAG_TO_CLEAN" 2>/dev/null || true
    fi
    
    exit $exit_code
}

# 主函数
main() {
    local target_env=""
    
    # 设置清理函数
    trap cleanup EXIT
    
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
            development|testing|production)
                target_env="$1"
                shift
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
    log_info "当前分支: ${current_branch}"
    
    # 检测或确认环境
    if [ -z "$target_env" ]; then
        # 没有指定环境参数，自动检测环境
        target_env=$(detect_environment "$current_branch")
        log_info "根据分支 '${current_branch}' 自动检测环境: ${target_env}"
    fi
    
    log_success "目标环境: ${target_env}"
    
    # 验证输入参数
    validate_inputs "$target_env"
    
    # 验证分支和环境匹配
    validate_branch_env "$current_branch" "$target_env"
    
    # 生产环境特殊处理
    if [ "$target_env" = "production" ]; then
        log_info "生产环境构建，检查代码状态..."
        check_production_ready
        create_git_tag
    fi
    
    # 获取版本号
    local version
    version=$(get_version "$target_env")
    if [ "$target_env" = "production" ]; then
        log_info "生产环境使用 tag 版本: ${version}"
    else
        log_info "${target_env} 环境使用 commit_id 版本: ${version}"
    fi
    
    # 确保使用本地 Docker
    ensure_local_docker
    
    # 登录 Docker
    docker_login
    
    # 构建镜像
    build_docker_image "$version" "$target_env"
    
    log_success "🎉 镜像构建完成！"
}

# 确保在 git 仓库中运行
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    log_error "必须在 git 仓库中运行此脚本"
fi

# 运行主函数
main "$@"
