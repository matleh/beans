---
title: Simplify beans show --json output
status: backlog
type: bug
priority: normal
created_at: 2025-12-12T23:53:33Z
updated_at: 2025-12-12T23:53:49Z
parent: beans-xnp8
---

## Problem

The `beans show --json` command currently wraps the bean data in an unnecessary envelope:

```json
{
  "success": true,
  "bean": {
    "id": "beans-iggk",
    ...
  }
}
```

## Expected Behavior

The command should return the bean object directly without any wrapper:

```json
{
  "id": "beans-iggk",
  ...
}
```

This is simpler and more consistent with typical JSON API design where the data type is already known from context.

## Checklist

- [ ] Update `cmd/show.go` to output bean data directly when `--json` flag is used
- [ ] Verify error handling still works appropriately (errors can still use a different format if needed)