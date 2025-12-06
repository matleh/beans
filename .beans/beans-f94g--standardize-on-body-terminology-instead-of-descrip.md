---
title: Standardize on 'body' terminology instead of 'description'
status: open
created_at: 2025-12-06T23:30:26Z
updated_at: 2025-12-06T23:30:26Z
---

The codebase and CLI inconsistently use 'body' and 'description' to refer to the same thing (the markdown content after frontmatter). We should standardize on 'body' everywhere.

## Areas to check
- [ ] CLI flag names (e.g., `--description` should become `--body`)
- [ ] Help text and command descriptions
- [ ] Code comments and variable names
- [ ] Documentation (CLAUDE.md, etc.)
- [ ] JSON output field names (verify it's already 'body')

## Notes
- 'body' is already used in the Bean struct and JSON output
- Main inconsistency is likely in CLI flags and help text