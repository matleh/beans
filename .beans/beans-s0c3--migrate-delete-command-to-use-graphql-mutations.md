---
title: Migrate 'delete' command to use GraphQL mutations
status: completed
type: task
priority: normal
created_at: 2025-12-09T12:04:25Z
updated_at: 2025-12-09T13:02:28Z
links:
    - blocks: beans-wp2o
    - parent: beans-7ao1
---

## Summary

Migrate the `beans delete` command to use GraphQL mutations instead of directly calling core methods.

## Current Implementation

In `cmd/delete.go`:
- Line 28: `core.Get(args[0])` - fetch bean to verify it exists
- Line 37: `core.FindIncomingLinks(b.ID)` - find beans linking to this one
- Line 65: `core.RemoveLinksTo(b.ID)` - remove incoming links
- Line 77: `core.Delete(args[0])` - delete the bean

## Target Implementation

Once GraphQL mutations are available (beans-wp2o):
```go
resolver := &graph.Resolver{Core: core}

// Verify bean exists via query
b, err := resolver.Query().Bean(ctx, id)

// Check incoming links (need GraphQL support or keep direct call)
// This might need a new query: incomingLinks(id: ID!): [Link!]!

// Remove incoming links via mutation
resolver.Mutation().RemoveLinksTo(ctx, id)

// Delete via mutation
resolver.Mutation().DeleteBean(ctx, id)
```

## Note

Finding incoming links may need additional GraphQL query support.

## Blocked By

- beans-wp2o (Add GraphQL mutations for bean CRUD operations)

## Checklist

- [ ] Wait for GraphQL mutations to be implemented
- [ ] Determine how to query incoming links via GraphQL
- [ ] Replace `core.Get()` with GraphQL query
- [ ] Replace `core.RemoveLinksTo()` with mutation
- [ ] Replace `core.Delete()` with mutation
- [ ] Run tests