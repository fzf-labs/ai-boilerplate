# Release Checklist

Use this checklist before publishing a change, opening a pull request, pushing a
subtree, or handing a branch to another developer.

## Scope

Classify the change before running checks:

| Scope | Required Checks |
| --- | --- |
| Single template | Run that template's command from `docs/verification.md`. |
| Backend API or database contract | Run backend tests, regenerate affected clients, and type-check each affected frontend client. |
| Shared root docs or policy only | Run a clean `git status --short` check and any command needed to validate touched examples. |
| Cross-template behavior | Run every affected template check, then run the full verification set when local SDKs are available. |

## Pre-Publish Steps

1. Confirm the task source and acceptance criteria are recorded in the issue,
   PR, task summary, or commit context.
2. Run the smallest relevant verification command while developing.
3. Regenerate generated artifacts from the owning source files.
4. Run every affected verification command after the final edit.
5. Run the full verification set for broad cross-template changes when local
   SDKs and dependencies are available.
6. Check `git status --short` and remove unrelated generated output or local
   build artifacts.
7. Commit with a concise message naming the affected template or workflow.

## Full Verification Set

Run these from each listed directory:

```bash
cd ai-boilerplate-backend && go test ./...
cd ../ai-boilerplate-admin && pnpm check:type --filter=@vben/web-antd
cd ../ai-boilerplate-uniapp && pnpm check:type
cd ../ai-boilerplate-pc && pnpm type-check
cd ../ai-boilerplate-electron && npm run typecheck
cd ../ai-boilerplate-tauri && npm run build
cd ../ai-boilerplate-chrome-extension && npm run type-check && npm test
cd ../ai-boilerplate-ios && swift build
cd ../ai-boilerplate-android && ./gradlew detekt test assembleDebug -x validateSigningDebug
```

If a local SDK or signing prerequisite is missing, record the exact command and
failure output, then run the closest static check available for the edited
files.

## Git Hygiene

- Never include unrelated user changes in a task commit.
- Do not rewrite subtree or repository history unless the repository owner asks
  for it explicitly.
- Keep generated source and generated output in the same commit.
- For documentation-only changes, avoid committing dependency lockfile or build
  output churn.
