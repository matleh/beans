---
title: Update prompt.tmpl to nudge agents toward GraphQL
status: completed
type: task
priority: normal
created_at: 2025-12-11T20:10:00Z
updated_at: 2025-12-11T20:11:10Z
---

Update the cmd/prompt.tmpl template to encourage agents to use `beans query` with GraphQL as the primary method for reading/querying beans, while keeping CLI commands for mutations.

## Changes

- Reframe "Core Rules" to recommend GraphQL for reads
- Update "Finding work" section to show GraphQL as primary approach
- Update "Working on a bean" to suggest GraphQL for reads
- Move GraphQL section higher in the document (before relationship filtering details)
- Keep CLI mutation commands (create, update) as-is since they're simpler for single operations