---
title: Migrate 'update' command to use GraphQL mutations
status: todo
type: task
priority: normal
created_at: 2025-12-09T12:04:25Z
updated_at: 2025-12-09T12:06:08Z
links:
    - blocks: beans-wp2o
    - parent: beans-7ao1
---

## Summary

Migrate the `beans update` command to use GraphQL mutations instead of directly calling core methods.

## Current Implementation

In `cmd/update.go`:
- Line 52: `core.Get(id)` to fetch the bean
- Line 171: `core.Update(b)` to save changes

## Target Implementation

Once GraphQL mutations are available (beans-wp2o):
```go
resolver := &graph.Resolver{Core: core}

// Fetch bean via GraphQL query
b, err := resolver.Query().Bean(ctx, id)

// Update via GraphQL mutation
input := model.UpdateBeanInput{
  Title: &newTitle,
  Status: &newStatus,
  AddLinks: addLinks,
  RemoveLinks: removeLinks,
}
updated, err := resolver.Mutation().UpdateBean(ctx, id, input)
```

## Blocked By

- beans-wp2o (Add GraphQL mutations for bean CRUD operations)

## Checklist

- [ ] Wait for GraphQL mutations to be implemented
- [ ] Replace `core.Get()` with GraphQL query
- [ ] Build UpdateBeanInput from CLI flags
- [ ] Replace `core.Update()` with mutation call
- [ ] Handle link add/remove through mutation
- [ ] Run tests