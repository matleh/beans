---
title: Add search command
status: open
---

Add a `beans search` command for full-text search across bean titles and bodies.

Implementation:
- Add `beans search <query>` command
- Search across title and body content (case-insensitive)
- Support `--json` output
- Consider supporting regex patterns
- Show matching context/snippets in results

Example usage:
```
beans search authentication
beans search "login bug" --status open
```