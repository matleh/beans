---
title: Implement 'g t' key chord for tag filtering in TUI
status: done
type: feature
created_at: 2025-12-07T19:40:17Z
updated_at: 2025-12-07T19:43:42Z
---

Add a key chord 'g t' (go tags) that:
1. Opens a tag selection page listing all tags used in the system
2. Lets the user select a tag
3. Returns to the bean list filtered by that tag

Build this in an extensible manner for future filter operations.

## Checklist
- [x] Add key chord state tracking to the App model
- [x] Create a new view state for tag selection
- [x] Implement tags list view that shows all unique tags
- [x] Wire up navigation from list -> tags -> filtered list
- [x] Add ability to clear filter and return to full list
- [x] Update key bindings and help text