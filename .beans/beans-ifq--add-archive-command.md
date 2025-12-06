---
title: Add archive command
status: done
---

Add a `beans archive` command.

Implementation:

- Deletes all beans that have their status set to "done".
- Show the number of beans that are going to be deleted, and ask the user for confirmation.
- If the `--force` flag is provided, skip the confirmation step.

This keeps the main bean list clean while preserving history.
