---
# beans-jpst
title: Change default worktree base_ref from origin/main to main
status: completed
type: task
priority: normal
created_at: 2026-03-11T17:54:24Z
updated_at: 2026-03-11T18:02:28Z
order: zzzzz
---

The default base ref for worktree creation is currently `origin/main`. It should be `main` instead — the local branch is always available and doesn't require a remote. This affects the hardcoded default in `pkg/config/config.go` (DefaultWorktreeBaseRef), the comment/documentation strings, and the test expectations in `internal/gitutil/worktree_test.go`.

## Summary of Changes

Changed the default worktree base_ref from `origin/main` to `main` across the codebase: the `DefaultWorktreeBaseRef` constant, `.beans.yml` config, and all related doc comments.
