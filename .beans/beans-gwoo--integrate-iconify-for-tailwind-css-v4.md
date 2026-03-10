---
# beans-gwoo
title: Integrate Iconify for Tailwind CSS v4
status: todo
type: task
priority: normal
created_at: 2026-03-10T15:14:04Z
updated_at: 2026-03-10T15:14:37Z
---

Integrate the @iconify/tailwind4 plugin into the frontend for icon support.

## Steps

- [ ] Install `@iconify/tailwind4` and `@iconify/json` as dev dependencies
- [ ] Add `@plugin "@iconify/tailwind4"` directive to the CSS file
- [ ] Verify icons render correctly (e.g. `<span class="icon-[mdi-light--home]"></span>`)
- [ ] Consider configuring prefix/scale options and clean selectors if needed

## Reference

- Docs: https://iconify.design/docs/usage/css/tailwind/tailwind4/
- Usage syntax: `icon-[prefix--name]` (e.g. `icon-[mdi-light--home]`)
- Clean selector alternative requires specifying icon set prefixes in config
