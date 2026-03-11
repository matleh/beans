---
# beans-vyjw
title: Add Iconify icons to agent message composer buttons
status: completed
type: feature
priority: normal
created_at: 2026-03-10T17:44:33Z
updated_at: 2026-03-10T19:31:46Z
order: k
---

Add icons from the UIL library (via Iconify) to the action buttons below the agent message composer in the web UI.

## Tasks

- [x] Add Iconify dependency to frontend (already installed)
- [x] Identify which UIL icons to use for each button
- [x] Add icons to the composer action buttons
- [x] Verify visual appearance and alignment (build passes)

## Summary of Changes

Added UIL icons (via Iconify) to all agent message composer buttons:
- Send: uil--message
- Stop: uil--stop-circle
- Plan: uil--eye
- Act: uil--play
- YOLO: uil--bolt
- Compact: uil--compress-arrows
- Clear: uil--trash-alt

Also added flex alignment (inline-flex, items-center, gap) to Send/Stop buttons and the .btn-tab-sm class.
