---
# beans-a115
title: Reduce verbose agent logging
status: completed
type: task
priority: normal
created_at: 2026-03-11T22:07:58Z
updated_at: 2026-03-11T22:10:36Z
---

Server logs excessive agent message data. Key issues: unhandled events logged as raw JSON, all stderr from Claude Code logged, spawn args include session IDs. Fix by truncating/suppressing verbose output and using proper log levels.

## Summary of Changes

Reduced verbose agent logging in internal/agent/claude.go:
- Silenced stderr drain (Claude Code writes verbose progress info that overwhelmed logs)
- Unhandled events now log only the event type, not the full JSON payload
- Removed CLI args from process spawn log (session IDs, permission flags)
