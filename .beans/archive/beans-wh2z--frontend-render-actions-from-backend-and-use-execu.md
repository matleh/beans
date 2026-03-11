---
# beans-wh2z
title: 'Frontend: render actions from backend and use executeAgentAction'
status: completed
type: task
priority: normal
created_at: 2026-03-11T19:11:52Z
updated_at: 2026-03-11T19:39:40Z
parent: beans-x7bd
blocked_by:
    - beans-6hvi
    - beans-tk4x
---

Replace the hardcoded ActionButton instances in ChangesPane with a dynamic list fetched from the backend, and wire clicks to the new mutation.

## Tasks
- [x] Add `agentActions` query to frontend GraphQL operations
- [x] Add `executeAgentAction` mutation to frontend GraphQL operations
- [x] Update ChangesPane (or StatusPane) to fetch actions via the query
- [x] Re-fetch actions when agent session status changes to IDLE
- [x] Replace hardcoded ActionButton instances with dynamic rendering from query results
- [x] Wire button clicks to call `executeAgentAction` mutation instead of `sendMessage`
- [x] Remove old ActionButton prompt props and unused ActionContext code
- [x] Run `pnpm build` and fix any warnings
- [x] Update e2e tests if needed (no e2e tests referenced ActionButton)

## Summary of Changes

- Replaced hardcoded ActionButton instances with dynamic buttons fetched from backend via agentActions query
- Added executeAgentAction mutation call on button click
- ChangesPane now takes beanId prop instead of onAction callback
- Deleted ActionButton.svelte and actionContext.ts
- Updated WorkspaceView and PlanningView to pass beanId instead of onAction
- Actions re-fetch when agent transitions to IDLE
