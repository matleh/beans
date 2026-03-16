---
# beans-let4
title: Add Claude model chooser to agent composer
status: completed
type: feature
priority: normal
created_at: 2026-03-16T16:48:23Z
updated_at: 2026-03-16T16:57:49Z
---

Add a model selector dropdown to the AgentComposer that lets users pick which Claude model to use (sonnet, opus, haiku). The selection flows through GraphQL to the backend, which passes --model to the spawned claude CLI process.

## Summary of Changes

### Backend
- Added `Model` field to `Session` struct (`internal/agent/types.go`)
- Pass `--model <name>` flag in `buildClaudeArgs()` when model is set (`internal/agent/claude.go`)
- Added `SetModel()` method to agent manager (`internal/agent/manager.go`)
- Added `model` field to `AgentSession` GraphQL type
- Added `model` optional param to `sendAgentMessage` mutation
- Added `setAgentModel` mutation for changing model without sending
- Updated `agentSessionToModel` helper to include model
- Implemented `SetAgentModel` resolver

### Frontend
- Added `model` to `AgentSessionFields` fragment
- Added `SetAgentModel` mutation and `model` param to `SendAgentMessage`
- Added `setModel()` method to `AgentChatStore`
- Added model chooser button group (Sonnet/Opus/Haiku) to `AgentComposer`
- Wired model through `AgentChat` to composer and send

### Tests
- Added `TestBuildClaudeArgs_Model`, `TestBuildClaudeArgs_NoModel`
- Added `TestSetModel_CreatesSession`, `TestSetModel_UpdatesExisting`, `TestSetModel_NoopWhenSame`
