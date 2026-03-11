---
# beans-tk4x
title: Add executeAgentAction mutation
status: completed
type: task
priority: normal
created_at: 2026-03-11T19:11:47Z
updated_at: 2026-03-11T19:34:02Z
parent: beans-x7bd
---

Add a mutation that executes an agent action by ID, mapping it to a prompt and injecting it into the conversation.

## Tasks
- [x] Add `executeAgentAction(beanId: ID!, actionId: ID!): Boolean!` mutation to schema
- [x] Implement resolver: look up action by ID, build prompt, call `AgentMgr.SendMessage()`
- [x] Run codegen (`mise codegen`)
- [x] Add tests for the mutation

## Summary of Changes

Added executeAgentAction mutation with resolver that looks up actions from the consolidated registry and injects prompts via AgentMgr.SendMessage. Merged into main repo alongside beans-6hvi.
