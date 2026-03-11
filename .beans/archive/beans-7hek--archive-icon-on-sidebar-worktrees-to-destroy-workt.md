---
# beans-7hek
title: Archive icon on sidebar worktrees to destroy worktree
status: completed
type: feature
priority: normal
created_at: 2026-03-11T22:13:21Z
updated_at: 2026-03-11T22:17:09Z
---

Add an archive/destroy icon to worktree items in the sidebar. When clicked, it should confirm and then destroy the worktree using the existing removeWorktree/stopWork mutation.

## Summary of Changes

Added an archive icon button to each worktree item in the sidebar. The icon appears on hover, and clicking it opens a confirmation modal before destroying the worktree via `removeWorktree`. If the destroyed worktree was the active view, navigation redirects to Planning.
