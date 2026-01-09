#!/bin/bash

# Git新标签创建脚本
# 自动检测最新标签，生成新版本号并创建标签

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

# 获取当前分支
current_branch=$(git branch --show-current)
log_info "当前分支: ${current_branch}"

# 获取当前最新的 tag
LATEST_TAG=$(git tag -l | grep -E '^v?[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 1)

if [ -z "$LATEST_TAG" ]; then
    log_warning "未找到有效的版本标签，将创建初始版本 v1.0.0"
    LATEST_TAG="v0.0.0"
fi

log_info "当前最新的 tag: ${LATEST_TAG}"

# 检查当前代码是否比最新 tag 更新
if [ "$LATEST_TAG" = "v0.0.0" ]; then
    # 如果没有找到标签，检查是否有提交
    commit_count=$(git rev-list HEAD --count)
    if [ "$commit_count" -eq "0" ]; then
        log_error "仓库中没有任何提交，无法创建标签"
        exit 1
    fi
    log_success "检测到 ${commit_count} 个提交，可以创建初始版本标签"
else
    # 如果有现有标签，检查是否有新提交
    commit_count=$(git rev-list ${LATEST_TAG}..HEAD --count 2>/dev/null || echo "0")
    
    if [ "$commit_count" -eq "0" ]; then
        log_error "当前代码与最新 tag ${LATEST_TAG} 相同，无需创建新版本"
        exit 1
    fi
    
    log_success "检测到 ${commit_count} 个新提交，可以创建新版本"
fi

# 生成新的版本号
if [ "$LATEST_TAG" = "v0.0.0" ]; then
    # 如果没有现有标签，创建初始版本
    increment_type=${1:-patch}
    case $increment_type in
        major)
            NEW_VERSION="v1.0.0"
            ;;
        minor)
            NEW_VERSION="v0.1.0"
            ;;
        patch)
            NEW_VERSION="v0.0.1"
            ;;
        *)
            log_error "无效的版本递增类型: $increment_type (支持: major, minor, patch)"
            exit 1
            ;;
    esac
else
    # 移除 v 前缀（如果有）
    version_without_v=${LATEST_TAG#v}
    
    # 解析版本号
    IFS='.' read -r major minor patch <<< "$version_without_v"
    
    # 默认递增补丁版本
    increment_type=${1:-patch}
    
    case $increment_type in
        major)
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        minor)
            minor=$((minor + 1))
            patch=0
            ;;
        patch)
            patch=$((patch + 1))
            ;;
        *)
            log_error "无效的版本递增类型: $increment_type (支持: major, minor, patch)"
            exit 1
            ;;
    esac
    
    NEW_VERSION="v${major}.${minor}.${patch}"
fi
log_info "新的版本号: ${NEW_VERSION}"

# 询问用户确认
echo
read -p "是否创建新标签 ${NEW_VERSION}? (y/N): " confirm

if [[ ! $confirm =~ ^[Yy]$ ]]; then
    log_info "操作已取消"
    exit 0
fi

# 获取标签消息
echo
read -p "请输入标签消息 (默认: release ${NEW_VERSION}): " tag_message
tag_message=${tag_message:-"release ${NEW_VERSION}"}

# 创建新的 tag
log_info "创建标签: ${NEW_VERSION}"
git tag -a "${NEW_VERSION}" -m "${tag_message}"

log_success "标签 ${NEW_VERSION} 创建成功"

# 询问是否推送
echo
read -p "是否推送标签到远程仓库? (y/N): " push_confirm

if [[ $push_confirm =~ ^[Yy]$ ]]; then
    log_info "推送标签到远程仓库..."
    git push origin "${NEW_VERSION}"
    log_success "标签 ${NEW_VERSION} 已推送到远程仓库"
else
    log_info "标签未推送，可以稍后使用命令推送: git push origin ${NEW_VERSION}"
fi

echo
log_success "脚本执行完成！"