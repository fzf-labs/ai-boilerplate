# Subtree 配置列表（格式：名称|前缀|仓库|分支）
SUBTREES := \
	backend|ai-boilerplate-backend|git@github.com:fzf-labs/ai-boilerplate-backend.git|master \
	admin|ai-boilerplate-admin|git@github.com:fzf-labs/ai-boilerplate-admin.git|master \
	app|ai-boilerplate-app|git@github.com:fzf-labs/ai-boilerplate-app.git|master

# 颜色定义
COLOR_RESET := \033[0m
COLOR_GREEN := \033[32m
COLOR_YELLOW := \033[33m
COLOR_BLUE := \033[34m
COLOR_CYAN := \033[36m
COLOR_RED := \033[31m

# 辅助函数：从配置中提取字段
get-prefix = $(word 2,$(subst |, ,$(1)))
get-repo = $(word 3,$(subst |, ,$(1)))
get-branch = $(word 4,$(subst |, ,$(1)))
get-name = $(word 1,$(subst |, ,$(1)))

# 生成所有 subtree 名称列表
SUBTREE_NAMES := $(foreach st,$(SUBTREES),$(call get-name,$(st)))

# 检查是否在 git 主工作树中（防止在 worktree 中误操作 subtree）
# 主工作树的 .git 是目录，链接工作树的 .git 是文件
define check-main-worktree
	@GIT_ROOT=$$(cd "$$(git rev-parse --show-toplevel 2>/dev/null)" && pwd -P); \
	CURRENT_DIR=$$(pwd -P); \
	if [ -z "$$GIT_ROOT" ]; then \
		echo "$(COLOR_RED)❌ 错误：当前目录不在 git 仓库中$(COLOR_RESET)"; \
		exit 1; \
	fi; \
	if [ "$$CURRENT_DIR" != "$$GIT_ROOT" ]; then \
		echo "$(COLOR_RED)❌ 错误：请在仓库根目录执行 subtree 操作$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   当前目录：$$CURRENT_DIR$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   仓库根目录：$$GIT_ROOT$(COLOR_RESET)"; \
		exit 1; \
	fi; \
	if [ -f "$$GIT_ROOT/.git" ]; then \
		echo "$(COLOR_RED)❌ 错误：subtree 操作只能在主工作树 (main worktree) 中执行$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   当前位于链接工作树 (linked worktree)$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   请切换到主工作树后再执行此命令$(COLOR_RESET)"; \
		echo ""; \
		echo "$(COLOR_CYAN)   提示：使用 'git worktree list' 查看所有工作树$(COLOR_RESET)"; \
		exit 1; \
	fi
endef

# 动态生成 pull 目标
define make-pull-target
subtree-pull-$(call get-name,$(1)):
	$(check-main-worktree)
	@echo "$(COLOR_BLUE)正在更新 $(call get-prefix,$(1))...$(COLOR_RESET)"
	@OUTPUT=$$$$(git subtree pull --prefix=$(call get-prefix,$(1)) $(call get-repo,$(1)) $(call get-branch,$(1)) --squash 2>&1); \
	EXIT_CODE=$$$$?; \
	if [ $$$$EXIT_CODE -ne 0 ] && echo "$$$$OUTPUT" | grep -q "does not exist; use 'git subtree add'"; then \
		echo "$(COLOR_YELLOW)⚠️  $(call get-prefix,$(1)) 未添加，先执行添加操作...$(COLOR_RESET)"; \
		$(MAKE) subtree-add-$(call get-name,$(1)); \
		echo "$(COLOR_GREEN)✓ $(call get-prefix,$(1)) 添加完成，现在可以正常使用 pull 命令$(COLOR_RESET)"; \
	elif [ $$$$EXIT_CODE -ne 0 ]; then \
		echo "$$$$OUTPUT"; \
		exit 1; \
	else \
		echo "$$$$OUTPUT"; \
		echo "$(COLOR_GREEN)✓ $(call get-prefix,$(1)) 更新完成$(COLOR_RESET)"; \
	fi
endef

# 动态生成 push 目标
define make-push-target
subtree-push-$(call get-name,$(1)):
	$(check-main-worktree)
	@echo "$(COLOR_BLUE)正在推送 $(call get-prefix,$(1))...$(COLOR_RESET)"
	@OUTPUT=$$$$(git subtree push --prefix=$(call get-prefix,$(1)) $(call get-repo,$(1)) $(call get-branch,$(1)) 2>&1); \
	EXIT_CODE=$$$$?; \
	echo "$$$$OUTPUT"; \
	if echo "$$$$OUTPUT" | grep -q "Everything up-to-date"; then \
		echo "$(COLOR_YELLOW)⚠️  $(call get-prefix,$(1)) 没有新内容需要推送$(COLOR_RESET)"; \
	elif [ $$$$EXIT_CODE -ne 0 ] && echo "$$$$OUTPUT" | grep -q "non-fast-forward\|rejected"; then \
		echo "$(COLOR_YELLOW)⚠️  $(call get-prefix,$(1)) 推送被拒绝：远程有新的提交$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   请先执行: make subtree-pull-$(call get-name,$(1))$(COLOR_RESET)"; \
		exit 1; \
	elif [ $$$$EXIT_CODE -eq 0 ]; then \
		echo "$(COLOR_GREEN)✓ $(call get-prefix,$(1)) 推送完成$(COLOR_RESET)"; \
	fi
endef

# 动态生成 add 目标
define make-add-target
subtree-add-$(call get-name,$(1)):
	$(check-main-worktree)
	@echo "$(COLOR_BLUE)正在添加 $(call get-prefix,$(1)) 为 subtree...$(COLOR_RESET)"
	@if [ -n "$$$$(git status --porcelain)" ]; then \
		echo "$(COLOR_YELLOW)⚠️  工作区有未提交的更改，需要先提交或暂存$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)   请先执行: git add . && git commit -m 'chore: prepare for subtree add'$(COLOR_RESET)"; \
		exit 1; \
	fi
	@if [ -d "$(call get-prefix,$(1))" ]; then \
		echo "$(COLOR_YELLOW)⚠️  目录 $(call get-prefix,$(1)) 已存在，先移除后重新添加$(COLOR_RESET)"; \
		git rm -r --cached $(call get-prefix,$(1)) 2>/dev/null || true; \
		rm -rf $(call get-prefix,$(1)); \
		git commit -m "chore: remove $(call get-prefix,$(1)) for re-initialization" 2>/dev/null || true; \
	fi
	@git subtree add --prefix=$(call get-prefix,$(1)) $(call get-repo,$(1)) $(call get-branch,$(1)) --squash
	@echo "$(COLOR_GREEN)✓ $(call get-prefix,$(1)) 添加完成$(COLOR_RESET)"
endef

# 动态生成 diff 目标
define make-diff-target
subtree-diff-$(call get-name,$(1)):
	$(check-main-worktree)
	@echo "$(COLOR_CYAN)$(call get-prefix,$(1)) 的差异：$(COLOR_RESET)"
	@git diff HEAD -- $(call get-prefix,$(1))/
endef

# .PHONY 声明
.PHONY: help $(foreach name,$(SUBTREE_NAMES),subtree-pull-$(name) subtree-push-$(name) subtree-add-$(name) subtree-diff-$(name)) \
	subtree-pull-all subtree-push-all subtree-add-all subtree-status subtree-check-dirty subtree-list git-clean

# 默认目标
.DEFAULT_GOAL := help

help:
	@echo "$(COLOR_CYAN)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo "$(COLOR_CYAN)  Git Subtree 管理工具$(COLOR_RESET)"
	@echo "$(COLOR_CYAN)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_YELLOW)📥 拉取命令（从远程更新）：$(COLOR_RESET)"
	@$(foreach st,$(SUBTREES), \
		echo "  $(COLOR_GREEN)make subtree-pull-$(call get-name,$(st))$(COLOR_RESET)  - 更新 $(call get-prefix,$(st))";)
	@echo "  $(COLOR_GREEN)make subtree-pull-all$(COLOR_RESET)       - 更新所有 subtree"
	@echo ""
	@echo "$(COLOR_YELLOW)📤 推送命令（推送到远程）：$(COLOR_RESET)"
	@$(foreach st,$(SUBTREES), \
		echo "  $(COLOR_GREEN)make subtree-push-$(call get-name,$(st))$(COLOR_RESET)  - 推送 $(call get-prefix,$(st))";)
	@echo "  $(COLOR_GREEN)make subtree-push-all$(COLOR_RESET)       - 推送所有 subtree"
	@echo ""
	@echo "$(COLOR_YELLOW)🔍 查看命令：$(COLOR_RESET)"
	@echo "  $(COLOR_GREEN)make subtree-status$(COLOR_RESET)         - 查看所有 subtree 状态"
	@echo "  $(COLOR_GREEN)make subtree-list$(COLOR_RESET)           - 列出所有 subtree 配置"
	@echo "  $(COLOR_GREEN)make subtree-check-dirty$(COLOR_RESET)    - 检查是否有未提交的更改"
	@$(foreach st,$(SUBTREES), \
		echo "  $(COLOR_GREEN)make subtree-diff-$(call get-name,$(st))$(COLOR_RESET)   - 查看 $(call get-prefix,$(st)) 的差异";)
	@echo ""
	@echo "$(COLOR_YELLOW)➕ 添加命令（首次使用）：$(COLOR_RESET)"
	@$(foreach st,$(SUBTREES), \
		echo "  $(COLOR_GREEN)make subtree-add-$(call get-name,$(st))$(COLOR_RESET)    - 添加 $(call get-prefix,$(st))";)
	@echo ""
	@echo "$(COLOR_YELLOW)⚠️  危险命令：$(COLOR_RESET)"
	@echo "  $(COLOR_RED)make git-clean$(COLOR_RESET)            - 清空 git 历史并强制推送（不可恢复！）"
	@echo ""

# 列出所有 subtree 配置
subtree-list:
	@echo "$(COLOR_CYAN)配置的 Subtree 列表：$(COLOR_RESET)"
	@echo ""
	@$(foreach st,$(SUBTREES), \
		echo "$(COLOR_YELLOW)● $(call get-name,$(st))$(COLOR_RESET)"; \
		echo "  前缀：  $(call get-prefix,$(st))"; \
		echo "  仓库：  $(call get-repo,$(st))"; \
		echo "  分支：  $(call get-branch,$(st))"; \
		echo "";)

# 检查是否有未提交的更改
subtree-check-dirty:
	$(check-main-worktree)
	@echo "$(COLOR_BLUE)检查工作区状态...$(COLOR_RESET)"
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "$(COLOR_YELLOW)⚠️  警告：工作区有未提交的更改$(COLOR_RESET)"; \
		git status --short; \
		exit 1; \
	else \
		echo "$(COLOR_GREEN)✓ 工作区干净$(COLOR_RESET)"; \
	fi

# 查看所有 subtree 状态
subtree-status:
	$(check-main-worktree)
	@echo "$(COLOR_CYAN)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo "$(COLOR_CYAN)  Subtree 状态$(COLOR_RESET)"
	@echo "$(COLOR_CYAN)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo ""
	@$(foreach st,$(SUBTREES), \
		echo "$(COLOR_YELLOW)● $(call get-prefix,$(st))$(COLOR_RESET)"; \
		echo "  最近提交："; \
		git log --oneline -1 --color=always -- $(call get-prefix,$(st))/ 2>/dev/null | sed 's/^/    /' || echo "    $(COLOR_YELLOW)未找到提交记录$(COLOR_RESET)"; \
		echo "  本地更改："; \
		if [ -n "$$(git status --short $(call get-prefix,$(st))/ 2>/dev/null)" ]; then \
			git status --short $(call get-prefix,$(st))/ | sed 's/^/    /'; \
		else \
			echo "    $(COLOR_GREEN)无更改$(COLOR_RESET)"; \
		fi; \
		echo "";)

# 为每个 subtree 生成目标
$(foreach st,$(SUBTREES),$(eval $(call make-pull-target,$(st))))
$(foreach st,$(SUBTREES),$(eval $(call make-push-target,$(st))))
$(foreach st,$(SUBTREES),$(eval $(call make-add-target,$(st))))
$(foreach st,$(SUBTREES),$(eval $(call make-diff-target,$(st))))

# 添加所有 subtree
subtree-add-all: $(foreach name,$(SUBTREE_NAMES),subtree-add-$(name))
	@echo ""
	@echo "$(COLOR_GREEN)✓ 所有 subtree 添加完成$(COLOR_RESET)"

# 批量操作
subtree-pull-all: $(foreach name,$(SUBTREE_NAMES),subtree-pull-$(name))
	@echo ""
	@echo "$(COLOR_GREEN)✓ 所有 subtree 更新完成$(COLOR_RESET)"

# 推送所有 subtree 的更改到远程
subtree-push-all: $(foreach name,$(SUBTREE_NAMES),subtree-push-$(name))
	@echo ""
	@echo "$(COLOR_GREEN)✓ 所有 subtree 推送完成$(COLOR_RESET)"

# git 记录清除（危险操作：会清空所有历史记录并强制推送）
git-clean:
	$(check-main-worktree)
	@echo "$(COLOR_RED)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo "$(COLOR_RED)  ⚠️  警告：此操作将清空所有 git 历史记录！$(COLOR_RESET)"
	@echo "$(COLOR_RED)═══════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_YELLOW)此操作将执行以下步骤：$(COLOR_RESET)"
	@echo "  1. 创建孤立分支，丢弃所有历史"
	@echo "  2. 将当前文件作为初始提交"
	@echo "  3. 删除原 master 分支"
	@echo "  4. 强制推送到远程（不可恢复！）"
	@echo ""
	@echo "$(COLOR_RED)⚠️  此操作不可逆，远程仓库的历史将被永久删除！$(COLOR_RESET)"
	@echo ""
	@read -p "确认执行？输入 'yes' 继续: " confirm; \
	if [ "$$confirm" != "yes" ]; then \
		echo "$(COLOR_GREEN)已取消操作$(COLOR_RESET)"; \
		exit 0; \
	fi
	@echo ""
	@echo "$(COLOR_BLUE)正在清除 git 历史...$(COLOR_RESET)"
	@git checkout --orphan clean-branch
	@git add -A
	@git commit -am "chore: initial commit (history cleaned)"
	@git branch -D master
	@git branch -m master
	@echo "$(COLOR_BLUE)正在强制推送到远程...$(COLOR_RESET)"
	@git push -f origin master
	@echo ""
	@echo "$(COLOR_GREEN)✓ git 历史清除完成$(COLOR_RESET)"