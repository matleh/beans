---
# beans-vts8
title: Add sidebar navigation to web UI
status: completed
type: feature
priority: normal
created_at: 2026-03-10T18:56:08Z
updated_at: 2026-03-10T19:31:46Z
order: s
---

Add a sidebar to the beans-serve web UI with:
- Planning item at top (current Backlog/Board view + central agent)
- List of beans with active worktrees
- Each workspace shows agent chat + bean detail in split pane

## Tasks
- [x] Add activeView state to uiState.svelte.ts
- [x] Update +layout.ts and +layout.svelte for persistence/fallback
- [x] Extract PlanningView.svelte from +page.svelte
- [x] Create Sidebar.svelte component
- [x] Create WorkspaceView.svelte component
- [x] Rewrite +page.svelte as thin shell

## Summary of Changes

Added sidebar navigation to the beans-serve web UI:
- Fixed-width sidebar with "Planning" item (Backlog/Board + central agent) and per-bean workspace items
- Extracted PlanningView.svelte from +page.svelte (pure refactor)
- Created Sidebar.svelte with worktree bean list and running agent pulse indicators
- Created WorkspaceView.svelte with agent chat + bean detail in a resizable split pane
- Added activeView state to UIState with localStorage persistence
- Auto-fallback to planning view when a workspace's worktree is removed
