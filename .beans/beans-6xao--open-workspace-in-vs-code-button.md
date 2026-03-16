---
# beans-6xao
title: Open workspace in VS Code button
status: completed
type: feature
priority: normal
created_at: 2026-03-16T14:27:49Z
updated_at: 2026-03-16T14:32:30Z
---

Add a button to the workspace toolbar that opens the current worktree (or project root for main workspace) in VS Code via `code <path>`.

## Summary of Changes

Added an "Open in VS Code" button to the workspace toolbar that launches `code <path>` for the current workspace directory.

### Files changed:
- `internal/graph/schema.graphqls` — added `openInEditor(workspaceId: ID!): Boolean!` mutation
- `internal/graph/schema.resolvers.go` — resolver that looks up workspace path and runs `code <dir>`
- `frontend/src/lib/graphql/operations.graphql` — frontend mutation definition
- `frontend/src/lib/components/WorkspaceView.svelte` — VS Code button in toolbar
