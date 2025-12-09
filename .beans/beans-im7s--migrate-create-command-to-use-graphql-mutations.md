---
title: Migrate 'create' command to use GraphQL mutations
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

Migrate the `beans create` command to use GraphQL mutations instead of directly calling `core.Create()`.

## Current Implementation

In `cmd/create.go`:
- Lines 76-83: Creates `bean.Bean` struct directly
- Line 102: `core.Create(b)`

## Target Implementation

Once GraphQL mutations are available (beans-wp2o), use:
```go
resolver := &graph.Resolver{Core: core}
input := model.CreateBeanInput{
  Title: title,
  Type: beanType,
  Status: &status,
  // ...
}
b, err := resolver.Mutation().CreateBean(context.Background(), input)
```

## Blocked By

- beans-wp2o (Add GraphQL mutations for bean CRUD operations)

## Checklist

- [ ] Wait for GraphQL mutations to be implemented
- [ ] Import graph and model packages
- [ ] Build CreateBeanInput from CLI flags
- [ ] Replace direct bean creation with mutation call
- [ ] Update output handling for GraphQL model
- [ ] Run tests