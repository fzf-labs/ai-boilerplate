---
name: git-workflow
description: Git workflow and version control best practices for Kratos Admin project. Use when users need help with git operations, branching, commits, merges, pull requests, conflict resolution, or version control workflows. Triggers on keywords like git, branch, commit, merge, rebase, pull request, PR, push, checkout, conflict, or when users ask about version control operations and git best practices.
---

# Git Workflow

## Branch Naming Convention

| Branch Type | Format | Example |
|------------|--------|---------|
| Main | main/master | main |
| Development | develop | develop |
| Feature | feature/description | feature/user-login |
| Bug Fix | fix/description | fix/login-error |
| Hotfix | hotfix/description | hotfix/security-patch |
| Release | release/version | release/v1.0.0 |

## Commit Message Convention

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Commit Types

| Type | Description |
|------|-------------|
| feat | New feature |
| fix | Bug fix |
| docs | Documentation update |
| style | Code formatting (no functionality change) |
| refactor | Code refactoring (not feature or fix) |
| perf | Performance optimization |
| test | Test-related changes |
| chore | Build/tool changes |

### Examples

```bash
# New feature
git commit -m "feat(apitest): add batch import for test cases"

# Bug fix
git commit -m "fix(login): resolve timeout issue"

# Documentation
git commit -m "docs: update API documentation"

# Refactoring
git commit -m "refactor(service): restructure user service layer"
```

## Common Git Commands

### Branch Operations

```bash
# View local branches
git branch

# View all branches (local and remote)
git branch -a

# Create and switch to new branch
git checkout -b feature/xxx

# Switch branch
git checkout main

# Delete local branch
git branch -d feature/xxx

# Delete remote branch
git push origin --delete xxx
```

### Commit Operations

```bash
# Stage all changes
git add .

# Commit with message
git commit -m "message"

# Amend last commit
git commit --amend

# Push to remote
git push origin branch-name
```

### Sync Operations

```bash
# Fetch remote updates
git fetch origin

# Pull and merge
git pull origin main

# Pull and rebase
git pull --rebase origin main
```

### Merge Operations

```bash
# Merge branch
git merge feature/xxx

# Rebase onto main
git rebase main
```

### Undo Operations

```bash
# Undo commit, keep changes
git reset --soft HEAD^

# Undo commit, discard changes
git reset --hard HEAD^

# Discard file changes
git checkout -- file

# Stash current changes
git stash

# Restore stashed changes
git stash pop
```

### View History

```bash
# Concise history
git log --oneline

# Graph visualization
git log --graph

# View differences
git diff

# View file modification history
git blame file
```

## Standard Workflow

### Create Feature Branch

```bash
# 1. Start from main
git checkout main
git pull origin main

# 2. Create feature branch
git checkout -b feature/xxx
```

### Develop and Commit

```bash
# 3. Make changes and commit
git add .
git commit -m "feat: add new feature"

# 4. Push to remote
git push origin feature/xxx
```

### Create Pull Request

```bash
# 5. Create PR via GitHub/GitLab UI

# 6. After code review approval, merge

# 7. Clean up feature branch
git checkout main
git pull origin main
git branch -d feature/xxx
```

## Conflict Resolution

### Rebase Workflow

```bash
# 1. Fetch latest code
git fetch origin
git rebase origin/main

# 2. Resolve conflicts
# Edit conflicting files, keep desired code

# 3. Mark conflicts as resolved
git add .

# 4. Continue rebase
git rebase --continue

# 5. Push (may require force)
git push origin feature/xxx --force-with-lease
```

### Merge Workflow

```bash
# 1. Fetch and merge
git fetch origin
git merge origin/main

# 2. Resolve conflicts
# Edit conflicting files

# 3. Complete merge
git add .
git commit -m "merge: resolve conflicts with main"

# 4. Push
git push origin feature/xxx
```

## .gitignore Configuration

```gitignore
# Python
__pycache__/
*.py[cod]
*.egg-info/
.eggs/
venv/
.env

# Node
node_modules/
dist/
.npm

# Go
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# Logs
*.log
logs/

# System files
.DS_Store
Thumbs.db

# Local config
*.local
.env.local
.env.*.local

# Build artifacts
/build
/target
/out
```

## Git Aliases (Optional)

Add to `~/.gitconfig` for convenience:

```ini
[alias]
    co = checkout
    br = branch
    ci = commit
    st = status
    unstage = reset HEAD --
    last = log -1 HEAD
    visual = log --graph --oneline --all
    amend = commit --amend --no-edit
```

## Best Practices

**DO:**
- ✅ Create feature branch from latest main
- ✅ Use conventional commit messages
- ✅ Keep commits atomic and focused
- ✅ Pull latest changes before starting work
- ✅ Resolve conflicts promptly
- ✅ Request code review before merging
- ✅ Delete branches after merging

**DON'T:**
- ❌ Commit directly to main branch
- ❌ Commit sensitive information (passwords, keys)
- ❌ Commit large files (>10MB)
- ❌ Force push to shared branches
- ❌ Use unclear commit messages
- ❌ Leave merge conflicts unresolved
- ❌ Commit generated files (unless necessary)

## Pre-Commit Checklist

- [ ] Branch created from latest main
- [ ] Commit messages follow convention
- [ ] No sensitive information included
- [ ] All conflicts resolved
- [ ] Code passes tests/linting
- [ ] Changes reviewed by peer

## Troubleshooting

### Accidentally Committed to Wrong Branch

```bash
# 1. Undo commit but keep changes
git reset --soft HEAD^

# 2. Stash changes
git stash

# 3. Switch to correct branch
git checkout correct-branch

# 4. Apply changes
git stash pop

# 5. Commit
git add .
git commit -m "message"
```

### Need to Update Feature Branch with Main

```bash
# Option 1: Rebase (cleaner history)
git checkout feature/xxx
git fetch origin
git rebase origin/main

# Option 2: Merge (preserves history)
git checkout feature/xxx
git fetch origin
git merge origin/main
```

### Committed Sensitive Information

```bash
# Remove file from history
git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch path/to/sensitive-file" \
  --prune-empty --tag-name-filter cat -- --all

# Force push (use with extreme caution)
git push origin --force --all

# Rotate compromised credentials immediately
```

## GitHub/GitLab Integration

### Create PR from Command Line (GitHub)

```bash
# Using GitHub CLI (gh)
gh pr create --title "feat: add new feature" --body "Description of changes"

# With assignee and labels
gh pr create --assignee @me --label "enhancement,backend"
```

### View PR Status

```bash
# List PRs
gh pr list

# View PR details
gh pr view 123

# Check PR status
gh pr status
```
