---
# beans-9df8
title: Fix failing workspace e2e tests
status: completed
type: bug
priority: normal
created_at: 2026-03-17T15:20:37Z
updated_at: 2026-03-17T15:24:39Z
---

All workspace-related e2e tests fail because the sidebar renders workspace items as div[role=button] instead of semantic <button> elements. The e2e tests use CSS selectors like button.font-medium which don't match divs.

## Summary of Changes

Fixed all 13 failing workspace e2e tests by:
1. Changed workspace sidebar items from `<div role="button">` to semantic `<button>` elements in `Sidebar.svelte`, matching what the e2e test selectors expect (`button.font-medium`)
2. Restructured the DOM to avoid invalid nested buttons — the destroy worktree button is now a sibling of the workspace button inside a flex wrapper, not a child
3. Updated e2e test selectors in `workspace-create.spec.ts` and `workspace-destroy-warning.spec.ts` to find the destroy button via the parent card (`div.rounded-md`) rather than via the workspace button
