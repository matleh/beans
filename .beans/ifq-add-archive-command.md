---
title: Add archive command
status: open
---

Add a `beans archive` command to move completed beans to an archive folder instead of deleting them.

Implementation:
- `beans archive <id>` moves a bean to `.beans/_archive/`
- `beans archive --all-done` archives all beans with status 'done'
- `beans list` excludes archived beans by default
- `beans list --archived` shows only archived beans
- `beans unarchive <id>` moves a bean back from archive
- Support `--json` output

This keeps the main bean list clean while preserving history.