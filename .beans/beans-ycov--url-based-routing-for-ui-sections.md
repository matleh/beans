---
# beans-ycov
title: URL-based routing for UI sections
status: completed
type: feature
priority: normal
created_at: 2026-03-11T17:38:53Z
updated_at: 2026-03-11T18:02:28Z
order: "7"
---

Use URL routes to reflect which part of the UI the user is in (planning view, bean workspace, etc.). This enables deep linking, browser back/forward navigation, and shareable URLs for specific views.

## Plan

URL scheme:
- `/planning` → Planning backlog (default)
- `/planning/board` → Planning board
- `/workspace/:beanId` → Workspace view
- `/` → redirects to `/planning`
- `?bean=<id>` preserved as query param for selected bean

## Tasks

- [x] Create SvelteKit route files (`/planning`, `/planning/board`, `/workspace/[beanId]`)
- [x] Move Sidebar from `+page.svelte` into `+layout.svelte`
- [x] Update `UIState` navigation methods to use `goto()` instead of localStorage
- [x] Sync `UIState.activeView` and `planningView` from URL path
- [x] Update `+layout.ts` to remove activeView/planningView from localStorage init
- [x] Handle browser back/forward via SvelteKit's built-in routing
- [x] Verify `?bean=<id>` deep linking still works
- [x] Update e2e tests if needed (all 24 pass without changes)
- [x] Run build and verify no warnings
