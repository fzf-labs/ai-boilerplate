#!/bin/bash

# Git 标签创建脚本
# 自动创建新的语义化版本标签并推送到远程仓库

set -e  # 出错时立即退出

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
Git 标签创建脚本

用法:
    $0 [选项]

选项:
    --major         创建主版本号标签 (1.0.0 -> 2.0.0)
    --minor         创建次版本号标签 (1.0.0 -> 1.1.0)
    --patch         创建修订版本号标签 (1.0.0 -> 1.0.1) [默认]
    --help, -h      显示此帮助信息

示例:
    $0                  # 创建补丁版本号标签 (默认)
    $0 --minor          # 创建次版本号标签
    $0 --major          # 创建主版本号标签

说明:
    - 脚本会自动检查当前代码状态
    - 确保没有未提交和未推送的代码
    - 自动生成新的语义化版本号
    - 创建标签并推送到远程仓库
EOF
}

# 检查是否在Git仓库中
check_git_repo() {
    if ! git rev-parse --git-dir > /dev/null 2>&1; then
        log_error "必须在 git 仓库中运行此脚本"
    fi
}

# 检查当前分支是否为master
check_master_branch() {
    local current_branch
    current_branch=$(git rev-parse --abbrev-ref HEAD)
    if [ "$current_branch" != "master" ]; then
        log_error "只能在 master 分支创建生产版本标签，当前分支: $current_branch"
    fi
    log_info "当前分支: $current_branch ✓"
}

# 检查代码状态
check_code_status() {
    # 检测当前是否有未提交的代码
    if [ -n "$(git status --porcelain)" ]; then
        log_error "当前有未提交的代码，请先提交代码"
    fi
    
    # 检测当前是否有未推送的代码
    if [ -n "$(git cherry -v 2>/dev/null)" ]; then
        log_error "当前有未推送的代码，请先推送代码"
    fi
    
    log_success "代码状态检查通过"
}

# 获取最新标签
get_latest_tag() {
    local latest_tag
    latest_tag=$(git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 1)
    
    if [ -z "$latest_tag" ]; then
        # 如果没有找到标签，使用默认版本
        echo "v0.0.0"
    else
        echo "$latest_tag"
    fi
}

# 生成新版本号
generate_new_version() {
    local latest_tag="$1"
    local bump_type="$2"
    
    # 移除 'v' 前缀
    local version=${latest_tag#v}
    
    # 分解版本号
    local major minor patch
    IFS='.' read -r major minor patch <<< "$version"
    
    case "$bump_type" in
        "major")
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        "minor")
            minor=$((minor + 1))
            patch=0
            ;;
        "patch")
            patch=$((patch + 1))
            ;;
        *)
            log_error "无效的版本类型: $bump_type"
            ;;
    esac
    
    echo "v${major}.${minor}.${patch}"
}

# 检查是否有新提交
check_new_commits() {
    local latest_tag="$1"
    local commit_count
    
    if [ "$latest_tag" = "v0.0.0" ]; then
        # 如果是初始标签，检查是否有任何提交
        commit_count=$(git rev-list --count HEAD)
    else
        # 检查自最新标签以来的提交数量
        commit_count=$(git rev-list ${latest_tag}..HEAD --count)
    fi
    
    if [ "$commit_count" -eq "0" ]; then
        log_error "当前代码与最新标签 ${latest_tag} 相同，无需创建新版本"
    fi
    
    log_info "检测到 ${commit_count} 个新提交，可以创建新版本"
}

# 创建并推送标签
create_and_push_tag() {
    local new_version="$1"
    
    # 创建标签
    git tag -a "$new_version" -m "release $new_version"
    log_success "创建标签: ${new_version}"
    
    # 推送标签
    git push origin "$new_version"
    log_success "推送标签到远程仓库: ${new_version}"
}

# 主函数
main() {
    local bump_type="patch"  # 默认为补丁版本号
    
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
            --major)
                bump_type="major"
                shift
                ;;
            --minor)
                bump_type="minor"
                shift
                ;;
            --patch)
                bump_type="patch"
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
    
    # 前置检查
    check_git_repo
    check_master_branch
    check_code_status
    
    # 获取当前最新标签
    local latest_tag
    latest_tag=$(get_latest_tag)
    log_info "当前最新标签: ${latest_tag}"
    
    # 检查是否有新提交
    check_new_commits "$latest_tag"
    
    # 生成新版本号
    local new_version
    new_version=$(generate_new_version "$latest_tag" "$bump_type")
    log_info "新版本号: ${new_version} (${bump_type} 升级)"
    
    # 创建并推送标签
    create_and_push_tag "$new_version"
    
    log_success "🎉 标签创建完成！"
    echo
    log_info "可以使用以下命令查看标签："
    echo "  git tag -l"
    echo "  git describe --tags"
}

# 运行主函数
main "$@"
