---
title: Make beans list responsive to terminal width
status: completed
type: feature
priority: normal
created_at: 2025-12-13T11:56:39Z
updated_at: 2025-12-13T12:12:18Z
---

Add terminal width detection to the CLI list command so it can adjust its output (column widths, tags visibility) similar to how the TUI does it.

## Approach
1. Detect terminal width using golang.org/x/term
2. Use existing CalculateResponsiveColumns() function from internal/ui/styles.go
3. Update RenderTree to accept and use responsive columns
4. Pass calculated columns through to RenderBeanRow

## Checklist
- [ ] Add terminal width detection to cmd/list.go
- [ ] Update RenderTree signature to accept terminal width
- [ ] Calculate responsive columns in RenderTree
- [ ] Pass responsive column config to RenderBeanRow calls
- [ ] Test with various terminal widths