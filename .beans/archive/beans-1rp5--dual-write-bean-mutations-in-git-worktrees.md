---
# beans-1rp5
title: Dual-write bean mutations in git worktrees
status: completed
type: feature
priority: normal
created_at: 2026-03-10T16:30:31Z
updated_at: 2026-03-10T17:42:16Z
---

When running in a secondary git worktree, bean mutations should write to BOTH the main repo's .beans/ (primary) AND the worktree's own .beans/ (mirror). This allows worktree branches to commit bean changes alongside code.

## Summary of Changes

- Added `mirrorRoot` field to `Core` with `SetMirrorRoot()`/`MirrorRoot()` accessors and a `mirrorOp()` helper
- All mutations (`saveToDisk`, `Delete`, `Archive`, `Unarchive`, `LoadAndUnarchive`, `Init`) now mirror to the secondary path (best-effort)
- Updated `resolveBeansPath()` in `root.go` to return both primary and mirror paths; mirror is set when running in a secondary git worktree
- Added 7 tests covering mirror create, update, delete, archive, unarchive, failure tolerance, and no-mirror no-op
- Updated existing `root_test.go` for the new 3-return-value signature, plus a new test for empty mirror in non-worktree context

## Reverted

Mirror functionality removed — worktree agents should only interact with the main repo's .beans/ (which the existing redirect already handles).
