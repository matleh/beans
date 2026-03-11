---
# beans-2j8h
title: Add agent.enabled config setting to hide agent UI
status: completed
type: feature
priority: normal
created_at: 2026-03-11T22:11:22Z
updated_at: 2026-03-11T22:15:17Z
---

Add agent.enabled config setting (default true). When disabled, hide agent chats, status panes, and worktree functionality in the web UI. beans init should set it to true by default.

## Summary of Changes

### Backend (Go)
- Added `Enabled *bool` field to `AgentConfig` struct in `pkg/config/config.go`
- Added `IsAgentEnabled()` getter that defaults to `true` when unset
- `beans init` sets `agent.enabled: true` by default (via `Default()`)
- Added `agentEnabled` GraphQL query to expose the setting to the frontend
- Added tests for `IsAgentEnabled`, YAML serialization, and config loading

### Frontend (Svelte)
- Created `config.svelte.ts` store that fetches `agentEnabled` on mount
- **Sidebar**: Hides Workspaces section and agent status indicators when disabled
- **PlanningView**: Hides Status and Agent toggle buttons when disabled
- **WorkspaceView**: Imports config store (workspace route is already unreachable when disabled)
- **BeanDetail**: Hides "Start Work" button when agents are disabled
- **Layout**: Redirects to planning view when agents are disabled and on a workspace route
