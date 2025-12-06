---
title: Add priority field to beans
status: open
---

Add an optional priority field to beans (high, medium, low or 1-5 scale). This helps agents and humans prioritize which beans to work on first.

Implementation:
- Add `priority` field to Bean struct and front matter
- Update `beans new` to accept `--priority` flag
- Update `beans list` to show priority and support `--sort priority`
- Consider adding `beans priority <id> <priority>` command similar to status