#!/bin/bash

# PostgreSQL 数据库备份脚本
# 包含表结构和数据备份功能
# 日期: $(date +"%Y-%m-%d")

set -e  # 遇到错误时退出脚本

# 配置变量
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
BACKUP_DIR="$PROJECT_ROOT/doc/sql/backup"
LOG_FILE="$BACKUP_DIR/backup.log"

# 数据库连接配置 (从环境变量获取)
DB_HOST="${DB_HOST}"
DB_PORT="${DB_PORT}"
DB_USER="${DB_USER}"
DB_PASSWORD="${DB_PASSWORD}"
DB_NAME="${DB_NAME}"

# 备份文件命名
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="ai_boilerplate_backup_${TIMESTAMP}.sql"
SCHEMA_ONLY_FILE="ai_boilerplate_schema_${TIMESTAMP}.sql"
DATA_ONLY_FILE="ai_boilerplate_data_${TIMESTAMP}.sql"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    local message="$1"
    echo -e "${GREEN}[INFO]${NC} $(date '+%Y-%m-%d %H:%M:%S') - $message"
    # 确保日志目录存在
    [ ! -d "$(dirname "$LOG_FILE")" ] && mkdir -p "$(dirname "$LOG_FILE")"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - INFO - $message" >> "$LOG_FILE"
}

log_warn() {
    local message="$1"
    echo -e "${YELLOW}[WARN]${NC} $(date '+%Y-%m-%d %H:%M:%S') - $message"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - WARN - $message" >> "$LOG_FILE"
}

log_error() {
    local message="$1"
    echo -e "${RED}[ERROR]${NC} $(date '+%Y-%m-%d %H:%M:%S') - $message"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - ERROR - $message" >> "$LOG_FILE"
}

# 检查依赖
check_dependencies() {
    log_info "检查依赖工具..."
    
    if ! command -v pg_dump &> /dev/null; then
        log_error "pg_dump 未找到，请安装 PostgreSQL 客户端工具"
        exit 1
    fi
    
    if ! command -v psql &> /dev/null; then
        log_error "psql 未找到，请安装 PostgreSQL 客户端工具"
        exit 1
    fi
    
    log_info "依赖工具检查完成"
}

# 创建备份目录
create_backup_dir() {
    if [ ! -d "$BACKUP_DIR" ]; then
        mkdir -p "$BACKUP_DIR"
        log_info "创建备份目录: $BACKUP_DIR"
    fi
}

# 测试数据库连接
test_connection() {
    log_info "测试数据库连接..."
    
    export PGPASSWORD="$DB_PASSWORD"
    
    if ! psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT version();" &> /dev/null; then
        log_error "无法连接到数据库 $DB_NAME"
        log_error "请检查数据库配置: Host=$DB_HOST, Port=$DB_PORT, User=$DB_USER, DB=$DB_NAME"
        exit 1
    fi
    
    log_info "数据库连接测试成功"
}

# 获取数据库信息
get_db_info() {
    log_info "获取数据库信息..."
    
    export PGPASSWORD="$DB_PASSWORD"
    
    local db_size=$(psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT pg_size_pretty(pg_database_size('$DB_NAME'));" | xargs)
    local table_count=$(psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public';" | xargs)
    
    log_info "数据库大小: $db_size"
    log_info "表数量: $table_count"
}

# 执行完整备份 (结构 + 数据)
backup_full() {
    log_info "开始完整备份 (结构 + 数据)..."
    
    export PGPASSWORD="$DB_PASSWORD"
    
    local backup_path="$BACKUP_DIR/$BACKUP_FILE"
    
    if pg_dump \
        --host="$DB_HOST" \
        --port="$DB_PORT" \
        --username="$DB_USER" \
        --dbname="$DB_NAME" \
        --clean \
        --if-exists \
        --create \
        --format=plain \
        --encoding=UTF8 \
        --no-owner \
        --no-privileges \
        --no-comments \
        --no-publications \
        --no-subscriptions \
        --no-security-labels \
        --no-tablespaces \
        --inserts \
        --file="$backup_path"; then
        
        local file_size=$(du -h "$backup_path" | cut -f1)
        log_info "完整备份成功: $backup_path (大小: $file_size)"
        
        # 清理注释
        clean_sql_comments "$backup_path"
        
        return 0
    else
        log_error "完整备份失败"
        return 1
    fi
}

# 执行仅结构备份
backup_schema_only() {
    log_info "开始结构备份..."
    
    export PGPASSWORD="$DB_PASSWORD"
    
    local backup_path="$BACKUP_DIR/$SCHEMA_ONLY_FILE"
    
    if pg_dump \
        --host="$DB_HOST" \
        --port="$DB_PORT" \
        --username="$DB_USER" \
        --dbname="$DB_NAME" \
        --clean \
        --if-exists \
        --create \
        --schema-only \
        --format=plain \
        --encoding=UTF8 \
        --no-owner \
        --no-privileges \
        --no-comments \
        --no-publications \
        --no-subscriptions \
        --no-security-labels \
        --no-tablespaces \
        --file="$backup_path"; then
        
        local file_size=$(du -h "$backup_path" | cut -f1)
        log_info "结构备份成功: $backup_path (大小: $file_size)"
        
        # 清理注释
        clean_sql_comments "$backup_path"
        
        return 0
    else
        log_error "结构备份失败"
        return 1
    fi
}

# 执行仅数据备份
backup_data_only() {
    log_info "开始数据备份..."
    
    export PGPASSWORD="$DB_PASSWORD"
    
    local backup_path="$BACKUP_DIR/$DATA_ONLY_FILE"
    
    if pg_dump \
        --host="$DB_HOST" \
        --port="$DB_PORT" \
        --username="$DB_USER" \
        --dbname="$DB_NAME" \
        --data-only \
        --format=plain \
        --encoding=UTF8 \
        --no-owner \
        --no-privileges \
        --disable-triggers \
        --no-comments \
        --no-publications \
        --no-subscriptions \
        --no-security-labels \
        --no-tablespaces \
        --inserts \
        --file="$backup_path"; then
        
        local file_size=$(du -h "$backup_path" | cut -f1)
        log_info "数据备份成功: $backup_path (大小: $file_size)"
        
        # 清理注释
        clean_sql_comments "$backup_path"
        
        return 0
    else
        log_error "数据备份失败"
        return 1
    fi
}


# 清理旧备份文件 (保留最近7天)
cleanup_old_backups() {
    log_info "清理旧备份文件 (保留最近7天)..."
    
    local deleted_count=0
    
    # 删除7天前的备份文件
    find "$BACKUP_DIR" -name "ai_boilerplate_*.sql" -mtime +7 -type f | while read -r file; do
        rm -f "$file"
        log_info "删除旧备份: $(basename "$file")"
        ((deleted_count++))
    done
    
    # 删除7天前的日志文件
    find "$BACKUP_DIR" -name "backup.log.*" -mtime +7 -type f | while read -r file; do
        rm -f "$file"
        log_info "删除旧日志: $(basename "$file")"
    done
    
    log_info "清理完成"
}

# 清理SQL文件中的注释和不必要的设置
clean_sql_comments() {
    local sql_file="$1"
    if [ -f "$sql_file" ]; then
        log_info "清理SQL文件中的注释和设置: $(basename "$sql_file")"
        
        # 创建临时文件
        local temp_file="${sql_file}.tmp"
        
        # 去掉pg_dump生成的标准注释和不必要的设置，但保留COPY语句和数据
        awk '
            # 跳过这些行
            /^--$/ { next }
            /^-- PostgreSQL database dump$/ { next }
            /^-- Dumped from database version/ { next }
            /^-- Dumped by pg_dump version/ { next }
            /^-- Name: .+; Type: .+; Schema: .+; Owner: -$/ { next }
            /^-- Name: .+; Type: .+; Schema: -; Owner: -$/ { next }
            /^-- Data for Name: .+; Type: .+; Schema: .+; Owner: -$/ { next }
            /^-- PostgreSQL database dump complete$/ { next }
            /^\\restrict .+$/ { next }
            /^\\unrestrict .+$/ { next }
            /^\\connect .+$/ { next }
            /^SET statement_timeout = 0;$/ { next }
            /^SET lock_timeout = 0;$/ { next }
            /^SET idle_in_transaction_session_timeout = 0;$/ { next }
            /^SET client_encoding = .UTF8.;$/ { next }
            /^SET standard_conforming_strings = on;$/ { next }
            /^SELECT pg_catalog\.set_config\(.search_path., .., false\);$/ { next }
            /^SET check_function_bodies = false;$/ { next }
            /^SET xmloption = content;$/ { next }
            /^SET client_min_messages = warning;$/ { next }
            /^SET row_security = off;$/ { next }
            /^SET default_tablespace = ..;$/ { next }
            /^SET default_table_access_method = heap;$/ { next }
            /^DROP DATABASE IF EXISTS .+;$/ { next }
            /^CREATE DATABASE .+ WITH TEMPLATE = template0 ENCODING = .UTF8. LOCALE = .+;$/ { next }
            # 保留所有其他行，包括COPY语句和数据
            { print }
        ' "$sql_file" > "$temp_file"
        
        # 去掉多余的空行，但保留单个空行
        awk '
            /^$/ { 
                if (empty_count < 1) {
                    print
                    empty_count++
                }
                next
            }
            { 
                empty_count = 0
                print 
            }
        ' "$temp_file" > "${temp_file}.clean"
        
        # 替换原文件
        mv "${temp_file}.clean" "$sql_file"
        rm -f "$temp_file"
        
        local new_size=$(du -h "$sql_file" | cut -f1)
        log_info "SQL文件清理完成，新大小: $new_size"
    fi
}

# 轮转日志文件
rotate_log() {
    if [ -f "$LOG_FILE" ] && [ $(stat -f%z "$LOG_FILE" 2>/dev/null || stat -c%s "$LOG_FILE" 2>/dev/null || echo 0) -gt 10485760 ]; then
        mv "$LOG_FILE" "${LOG_FILE}.$(date +%Y%m%d_%H%M%S)"
        log_info "日志文件已轮转"
    fi
}

# 显示帮助信息
show_help() {
    echo -e "${BLUE}PostgreSQL 数据库备份脚本${NC}"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  -h, --help              显示帮助信息"
    echo "  -f, --full              完整备份 (结构 + 数据) [默认]"
    echo "  -s, --schema-only       仅备份结构"
    echo "  -d, --data-only         仅备份数据"
    echo ""
    echo "环境变量:"
    echo "  DB_HOST                 数据库主机"
    echo "  DB_PORT                 数据库端口"
    echo "  DB_USER                 数据库用户"
    echo "  DB_PASSWORD             数据库密码"
    echo "  DB_NAME                 数据库名称"
    echo ""
    echo "示例:"
    echo "  $0                      # 完整备份"
    echo "  $0 -s                   # 仅备份结构"
    echo "  $0 -d                   # 仅备份数据"
}

# 主函数
main() {
    local backup_type="full"
    
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                show_help
                exit 0
                ;;
            -f|--full)
                backup_type="full"
                shift
                ;;
            -s|--schema-only)
                backup_type="schema"
                shift
                ;;
            -d|--data-only)
                backup_type="data"
                shift
                ;;
            *)
                log_error "未知选项: $1"
                show_help
                exit 1
                ;;
        esac
    done
    
    # 开始备份流程
    log_info "开始 PostgreSQL 数据库备份"
    log_info "备份类型: $backup_type"
    log_info "数据库: $DB_HOST:$DB_PORT/$DB_NAME"
    
    # 轮转日志
    rotate_log
    
    # 创建备份目录
    create_backup_dir
    
    # 检查依赖
    check_dependencies
    
    # 测试连接
    test_connection
    
    # 获取数据库信息
    get_db_info
    
    # 执行备份
    local backup_success=true
    
    case $backup_type in
        "full")
            backup_full || backup_success=false
            ;;
        "schema")
            backup_schema_only || backup_success=false
            ;;
        "data")
            backup_data_only || backup_success=false
            ;;
    esac
    
    # 清理旧备份
    cleanup_old_backups
    
    # 完成
    if [ "$backup_success" = true ]; then
        log_info "数据库备份完成"
        log_info "备份文件位置: $BACKUP_DIR"
        log_info "备份文件将保存在 doc/sql/backup/ 目录下"
        exit 0
    else
        log_error "数据库备份失败"
        exit 1
    fi
}

# 捕获中断信号
trap 'log_error "备份过程被中断"; exit 1' INT TERM

# 执行主函数
main "$@"
