#!/bin/bash

# Git发布分支创建脚本
# 基于最新标签创建新的发布分支

set -e  # 遇到错误立即退出

# 颜色定义
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
}

# 检查是否在git仓库中
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    log_error "当前目录不是git仓库"
    exit 1
fi

# 检查是否有未提交的更改
if ! git diff-index --quiet HEAD --; then
    log_error "存在未提交的更改，请先提交或暂存"
    exit 1
fi

# 获取主分支名称
main_branch=${1:-master}

# 检查主分支是否存在
if ! git show-ref --verify --quiet refs/heads/"$main_branch" && ! git show-ref --verify --quiet refs/remotes/origin/"$main_branch"; then
    log_error "分支 ${main_branch} 不存在"
    exit 1
fi

log_info "基于分支: ${main_branch}"

# 获取远程最新信息
log_info "获取远程标签和分支信息..."
git fetch --tags origin

# 获取最新标签
latest_tag=$(git describe --tags --abbrev=0 origin/"$main_branch" 2>/dev/null || echo "")

if [ -z "$latest_tag" ]; then
    log_warning "未找到有效标签，将基于 v0.0.0 创建发布分支"
    latest_tag="v0.0.0"
fi

log_info "当前最新标签: ${latest_tag}"

# 解析版本号
version_without_v=${latest_tag#v}
IFS='.' read -r -a version_parts <<< "$version_without_v"

if [ ${#version_parts[@]} -ne 3 ]; then
    log_error "标签格式不正确: ${latest_tag} (期望格式: vX.Y.Z)"
    exit 1
fi

major=${version_parts[0]}
minor=${version_parts[1]}
patch=${version_parts[2]}

# 获取版本递增类型
echo
echo "选择版本递增类型:"
echo "1) patch - ${major}.${minor}.$((patch + 1))"
echo "2) minor - ${major}.$((minor + 1)).0"
echo "3) major - $((major + 1)).0.0"
echo

read -p "请选择 (1-3, 默认为1): " version_choice
version_choice=${version_choice:-1}

case $version_choice in
    1)
        new_patch=$((patch + 1))
        new_tag="v${major}.${minor}.${new_patch}"
        ;;
    2)
        new_minor=$((minor + 1))
        new_tag="v${major}.${new_minor}.0"
        ;;
    3)
        new_major=$((major + 1))
        new_tag="v${new_major}.0.0"
        ;;
    *)
        log_error "无效选择"
        exit 1
        ;;
esac

release_branch="release/${new_tag}"

log_info "将创建发布分支: ${release_branch}"

# 检查分支是否已存在
if git show-ref --verify --quiet refs/heads/"$release_branch"; then
    log_error "分支 ${release_branch} 已存在"
    exit 1
fi

if git show-ref --verify --quiet refs/remotes/origin/"$release_branch"; then
    log_error "远程分支 origin/${release_branch} 已存在"
    exit 1
fi

# 询问用户确认
echo
read -p "是否创建发布分支 ${release_branch}? (y/N): " confirm

if [[ ! $confirm =~ ^[Yy]$ ]]; then
    log_info "操作已取消"
    exit 0
fi

# 创建发布分支
log_info "创建发布分支..."
git checkout -b "$release_branch" origin/"$main_branch"

# 取消上游跟踪
git branch --unset-upstream

log_success "发布分支 ${release_branch} 创建成功"

# 询问是否推送到远程
echo
read -p "是否推送分支到远程仓库? (y/N): " push_confirm

if [[ $push_confirm =~ ^[Yy]$ ]]; then
    log_info "推送分支到远程仓库..."
    git push origin "$release_branch"
    
    # 设置上游跟踪
    git branch --set-upstream-to=origin/"$release_branch" "$release_branch"
    
    log_success "分支 ${release_branch} 已推送并设置上游跟踪"
else
    log_info "分支未推送，可以稍后使用命令推送: git push origin ${release_branch}"
fi

echo
log_info "当前分支: $(git branch --show-current)"
log_info "可以开始在此分支上进行发布准备工作"
log_success "脚本执行完成！"