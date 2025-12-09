---
title: Add GraphQL mutations for bean CRUD operations
status: todo
type: feature
priority: normal
created_at: 2025-12-09T12:03:46Z
updated_at: 2025-12-09T12:06:08Z
links:
    - parent: beans-7ao1
---

## Summary

The GraphQL schema currently only supports queries (read operations). To complete the migration of CLI commands to use GraphQL internally, we need to add mutations for create, update, and delete operations.

## Why This Is Needed

The following CLI commands need mutations to migrate to GraphQL:
- `beans create` - needs a `createBean` mutation
- `beans update` - needs an `updateBean` mutation  
- `beans delete` - needs a `deleteBean` mutation
- `beans archive` - needs `deleteBean` mutation (batch delete)

## Proposed Mutations

```graphql
type Mutation {
  """Create a new bean"""
  createBean(input: CreateBeanInput!): Bean!
  
  """Update an existing bean"""
  updateBean(id: ID!, input: UpdateBeanInput!): Bean!
  
  """Delete a bean by ID"""
  deleteBean(id: ID!): Boolean!
  
  """Remove all links pointing to a bean"""
  removeLinksTo(id: ID!): Int!
}

input CreateBeanInput {
  title: String!
  type: String!
  status: String
  priority: String
  tags: [String!]
  body: String
  links: [LinkInput!]
}

input UpdateBeanInput {
  title: String
  status: String
  type: String
  priority: String
  tags: [String!]
  body: String
  addLinks: [LinkInput!]
  removeLinks: [LinkInput!]
}

input LinkInput {
  type: String!
  target: String!
}
```

## Checklist

- [ ] Add mutation types to schema.graphqls
- [ ] Run `mise codegen` to regenerate code
- [ ] Implement createBean resolver
- [ ] Implement updateBean resolver
- [ ] Implement deleteBean resolver
- [ ] Implement removeLinksTo resolver
- [ ] Add tests for all mutations