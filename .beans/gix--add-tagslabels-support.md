---
title: Add tags/labels support
status: open
---

Add optional tags/labels to beans for flexible categorization beyond status and path.

Implementation:
- Add `tags` field to Bean struct (array of strings in front matter)
- Update `beans new` to accept `--tag` flag (repeatable)
- Update `beans list` to support `--tag` filter
- Consider `beans tag <id> <tag>` and `beans untag <id> <tag>` commands
- Show tags in list output

Example front matter:
```yaml
title: Fix login bug
status: open
tags: [bug, auth, urgent]
```