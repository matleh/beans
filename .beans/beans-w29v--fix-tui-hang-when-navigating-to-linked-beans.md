---
title: Fix TUI hang when navigating to linked beans
status: done
type: bug
created_at: 2025-12-07T17:24:03Z
updated_at: 2025-12-07T17:24:35Z
---

When pressing Enter to jump to a linked bean, the TUI hangs for several seconds.

## Root Cause
resolveAllLinks() calls store.FindByID() for each outgoing link, and FindByID() internally calls FindAll() - reading ALL bean files from disk each time. With N outgoing links, we do N+1 full directory scans.

## Fix
Load all beans once with FindAll(), build a lookup map, and use that for all link resolution.