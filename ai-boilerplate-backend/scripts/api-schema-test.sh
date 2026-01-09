#!/bin/bash
# API 契约测试脚本
# 基于 schemathesis 进行 API 自动化测试
# 用法: ./scripts/api-schema-test.sh [options] [swagger_file]

set -e

# ==================== 配置 ====================

API_URL="${TEST_API_URL}"
ADMIN_USER="${TEST_ADMIN_USER}"
ADMIN_PASS="${TEST_ADMIN_PASS}"
LOGIN_PATH="${TEST_LOGIN_PATH}"
SWAGGER_DIR="doc/swagger"
METHOD="GET"
VERBOSE=""
TOKEN=""

# 颜色
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# ==================== 帮助信息 ====================

show_help() {
    cat << EOF
Swagger Schema API 契约测试工具 (基于 schemathesis)

用法: $0 [options] [swagger_file]

Options:
  -u, --url URL       API URL (默认: $API_URL)
  -m, --method M      方法过滤: GET/POST/PUT/DELETE/ALL (默认: GET)
  -a, --all           测试所有 Swagger 文件
  -v, --verbose       详细输出
  --user USER         用户名 (默认: admin)
  --pass PASS         密码 (默认: 123456)
  --no-auth           不使用认证
  --install           安装 schemathesis
  -h, --help          显示帮助

示例:
  $0                                              # 测试所有文件的 GET 接口
  $0 doc/swagger/admin/v1/sys_admin.swagger.json  # 测试指定文件
  $0 -m POST                                      # 测试 POST 接口
  $0 -m ALL -v                                    # 测试所有方法，详细输出
  $0 --url http://api.example.com                 # 指定 API URL

环境变量:
  TEST_API_URL      API 基础 URL
  TEST_ADMIN_USER   管理员用户名
  TEST_ADMIN_PASS   管理员密码
EOF
}

# ==================== 工具函数 ====================

# 检查 schemathesis 是否安装
check_schemathesis() {
    if ! command -v schemathesis &> /dev/null; then
        echo -e "${RED}❌ schemathesis 未安装${NC}"
        echo ""
        echo "安装方式:"
        echo "  pip install schemathesis"
        echo "  # 或"
        echo "  pipx install schemathesis"
        echo "  # 或"
        echo "  $0 --install"
        exit 1
    fi
}

# 安装 schemathesis
install_schemathesis() {
    echo -e "${YELLOW}安装 schemathesis...${NC}"
    
    if command -v pipx &> /dev/null; then
        pipx install schemathesis
    elif command -v pip3 &> /dev/null; then
        pip3 install --user schemathesis
    elif command -v pip &> /dev/null; then
        pip install --user schemathesis
    else
        echo -e "${RED}错误: 需要 pip 或 pipx${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}✅ 安装完成${NC}"
}

# 获取认证 Token
get_token() {
    echo -e "${YELLOW}获取认证 Token...${NC}"
    
    local response
    response=$(curl -s -X POST "$API_URL$LOGIN_PATH" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$ADMIN_USER\",\"password\":\"$ADMIN_PASS\"}" 2>/dev/null)
    
    TOKEN=$(echo "$response" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    
    if [ -z "$TOKEN" ]; then
        echo -e "${YELLOW}⚠️  获取 Token 失败，继续测试无认证接口${NC}"
        return 1
    fi
    
    echo -e "${GREEN}✅ Token 获取成功${NC}"
    return 0
}

# 测试单个 Swagger 文件
test_file() {
    local file="$1"
    
    echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${BLUE}📄 $(basename "$file")${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    
    # 构建参数
    local args=("run" "$file" "--url" "$API_URL")
    
    # 方法过滤
    if [ "$METHOD" != "ALL" ]; then
        args+=("--include-method" "$METHOD")
    fi
    
    # 认证 Header
    if [ -n "$TOKEN" ]; then
        args+=("-H" "Authorization: Bearer $TOKEN")
    fi
    
    # 详细输出
    if [ -n "$VERBOSE" ]; then
        args+=("-v")
    fi
    
    # 执行测试
    schemathesis "${args[@]}" || true
}

# 收集 Swagger 文件
collect_files() {
    find "$SWAGGER_DIR" -name "*.swagger.json" -type f | grep -v "error_reason" | sort
}

# ==================== 主逻辑 ====================

# 解析参数
SWAGGER_FILE=""
TEST_ALL=""
NO_AUTH=""
DO_INSTALL=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -u|--url)
            API_URL="$2"
            shift 2
            ;;
        -m|--method)
            METHOD="$2"
            shift 2
            ;;
        -a|--all)
            TEST_ALL="true"
            shift
            ;;
        -v|--verbose)
            VERBOSE="true"
            shift
            ;;
        --user)
            ADMIN_USER="$2"
            shift 2
            ;;
        --pass)
            ADMIN_PASS="$2"
            shift 2
            ;;
        --no-auth)
            NO_AUTH="true"
            shift
            ;;
        --install)
            DO_INSTALL="true"
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        -*)
            echo -e "${RED}未知选项: $1${NC}"
            show_help
            exit 1
            ;;
        *)
            SWAGGER_FILE="$1"
            shift
            ;;
    esac
done

# 安装模式
if [ -n "$DO_INSTALL" ]; then
    install_schemathesis
    exit 0
fi

# 检查依赖
check_schemathesis

# 打印配置
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  Schema 契约测试 (schemathesis)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "  API URL:  $API_URL"
echo -e "  方法:     $METHOD"
echo -e "  用户:     $ADMIN_USER"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 获取 Token
if [ -z "$NO_AUTH" ]; then
    get_token || true
fi

# 执行测试
if [ -n "$SWAGGER_FILE" ]; then
    # 测试指定文件
    if [ ! -f "$SWAGGER_FILE" ]; then
        echo -e "${RED}❌ 文件不存在: $SWAGGER_FILE${NC}"
        exit 1
    fi
    test_file "$SWAGGER_FILE"
else
    # 测试所有文件
    files=$(collect_files)
    count=$(echo "$files" | wc -l | tr -d ' ')
    
    echo -e "\n📁 找到 $count 个 Swagger 文件"
    
    for file in $files; do
        test_file "$file"
    done
fi

echo -e "\n${GREEN}✅ 测试完成${NC}"
