---
title: Unify bean list rendering across CLI and TUI
status: done
type: task
created_at: 2025-12-07T17:30:02Z
updated_at: 2025-12-07T17:31:11Z
---

Create a shared bean row renderer and add Type column to TUI.

## Locations
- CLI list (cmd/list.go) - already shows ID, Status, Type, Title
- TUI list (internal/tui/list.go) - missing Type
- TUI relationships (internal/tui/detail.go) - shows link info, not bean type

## Plan
1. Create shared BeanRow renderer in internal/ui/
2. Update TUI list to include Type column
3. Update TUI relationships to show bean type