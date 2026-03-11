---
# beans-6hvi
title: Add AgentAction GraphQL type and query
status: completed
type: task
priority: normal
created_at: 2026-03-11T19:11:46Z
updated_at: 2026-03-11T19:34:02Z
parent: beans-x7bd
---

Define the backend types and GraphQL schema for agent actions.

## Tasks
- [x] Add `AgentAction` type to GraphQL schema: `{ id: ID!, label: String!, description: String }`
- [x] Add query field `agentActions(beanId: ID!): [AgentAction!]!`
- [x] Implement resolver that returns available actions (hardcode Commit and Review for now)
- [x] Run codegen (`mise codegen`)
- [x] Add tests for the resolver

## Summary of Changes

Added AgentAction GraphQL type, agentActions query, and consolidated action registry in agent_helpers.go. Merged into main repo alongside beans-tk4x.
