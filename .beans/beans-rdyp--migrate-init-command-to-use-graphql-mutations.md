---
title: Migrate 'init' command to use GraphQL mutations
status: todo
type: task
priority: normal
created_at: 2025-12-09T12:04:36Z
updated_at: 2025-12-09T12:06:08Z
links:
    - blocks: beans-wp2o
    - parent: beans-7ao1
---

## Summary

Migrate the `beans init` command to use GraphQL mutations instead of directly calling `beancore.Init()`.

## Current Implementation

In `cmd/init.go`:
- Line 47: `beancore.Init(dir)`

## Considerations

The `init` command is a special case - it creates the `.beans/` directory structure and potentially a `.beans.yml` config file. This happens *before* a valid beans repository exists.

## Options

1. **Add GraphQL mutation:**
   ```graphql
   type Mutation {
     initRepository(path: String): Boolean!
   }
   ```

2. **Keep as direct beancore call:**
   - Init is a bootstrapping operation
   - GraphQL resolver assumes a valid repository exists
   - May be cleaner to keep init outside GraphQL

## Recommendation

Consider keeping `init` as a direct beancore call since:
- It's a one-time bootstrapping operation
- GraphQL resolver expects an existing repository
- The operation is self-contained

If consistency is paramount, add it to GraphQL but handle the case where no repository exists yet.

## Blocked By

- beans-wp2o (Add GraphQL mutations for bean CRUD operations) - if we decide to add GraphQL support

## Checklist

- [ ] Decide: Add to GraphQL or keep as direct core call
- [ ] If GraphQL: Add mutation to schema
- [ ] If GraphQL: Implement resolver with special handling
- [ ] Document the decision