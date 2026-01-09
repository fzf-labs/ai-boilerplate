#!/bin/bash

# 发布分支创建脚本
# 从 master 分支创建新的 release 分支，版本号基于最新 tag +1

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
发布分支创建脚本

用法:
    $0 [选项]

选项:
    --help, -h      显示此帮助信息

功能:
    - 从 master 分支创建新的 release 分支
    - 自动获取最新 tag 并将版本号 +1
    - 分支名格式: release/v{major}.{minor}.{patch}

示例:
    $0              # 创建新的 release 分支

说明:
    - 脚本会自动从远程获取最新标签
    - 基于最新标签的 patch 版本 +1
    - 从 origin/master 创建新分支
    - 取消上游分支关联
EOF
}

# 检查是否在Git仓库中
check_git_repo() {
    if ! git rev-parse --git-dir > /dev/null 2>&1; then
        log_error "必须在 git 仓库中运行此脚本"
    fi
}

# 主函数
main() {
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
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
    
    log_info "开始创建新的 release 分支..."
    
    # 获取最新标签
    log_info "获取远程标签..."
    git fetch --tags
    
    # 获取最新的标签
    local latest_tag
    latest_tag=$(git describe --tags --abbrev=0 origin/master)
    log_info "当前最新标签: ${latest_tag}"
    
    # 解析版本号
    local version_without_v=${latest_tag#v}
    IFS='.' read -r -a version_parts <<< "$version_without_v"
    local major=${version_parts[0]}
    local minor=${version_parts[1]}
    local patch=${version_parts[2]}
    
    # 计算新版本号
    local new_patch=$((patch + 1))
    local new_tag="v${major}.${minor}.${new_patch}"
    log_info "新版本号: ${new_tag}"
    
    # 检查分支是否已存在
    local branch_name="release/${new_tag}"
    log_info "检查分支: ${branch_name}"
    
    # 检查本地分支是否存在
    if git show-ref --verify --quiet refs/heads/"$branch_name"; then
        echo
        echo -e "${RED}❌ 本地分支 '${branch_name}' 已经存在${NC}"
        echo
        echo "解决方案："
        echo "  1. 删除本地分支: git branch -D ${branch_name}"
        echo "  2. 或者手动创建其他版本的分支"
        exit 1
    fi
    
    # 检查远程分支是否存在
    if git show-ref --verify --quiet refs/remotes/origin/"$branch_name"; then
        echo
        echo -e "${RED}❌ 远程分支 'origin/${branch_name}' 已经存在${NC}"
        echo
        echo "解决方案："
        echo "  1. 删除远程分支: git push origin --delete ${branch_name}"
        echo "  2. 或者手动创建其他版本的分支"
        exit 1
    fi
    
    # 创建新分支
    log_info "创建分支: ${branch_name}"
    git checkout -b "$branch_name" origin/master
    
    # 取消上游分支关联
    git branch --unset-upstream
    
    log_success "🎉 新 release 分支创建完成！"
    echo
    log_info "当前状态："
    echo "  分支名称: ${branch_name}"
    echo "  基于版本: ${latest_tag}"
    echo "  目标版本: ${new_tag}"
    echo
    log_info "后续操作建议："
    echo "  1. 在此分支进行发布前的最后调整"
    echo "  2. 完成后创建标签: make git-new-tag"
    echo "  3. 构建镜像: make docker-build"
}

# 运行主函数
main "$@"