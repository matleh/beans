---
title: Change bean identifier format
status: done
created_at: 2025-12-06T15:05:31Z
updated_at: 2025-12-06T15:10:42Z
---



Change how bean identifiers work:

## Requirements

1. **Flexible identifiers**: Bean identifiers can be any string (but must be unique in the project)

2. **New filename format**: Change from `id-slug.md` to `id.slug.md` (dot separator instead of dash)
   - Example: `bean-x2y.change-id-format.md`

3. **Rename existing beans**: Update all existing bean files to use the new format

4. **Configurable prefix**: Add a `prefix` setting to beans.toml for newly created beans
   - Example: `prefix = "bean-"` would generate IDs like `bean-x2y`

5. **Smart default prefix**: When `beans init` creates the config file, the default prefix should be the name of the current directory
   - Example: In `/home/user/myproject`, default prefix would be `myproject-`

## Implementation Notes

- Update store.go to handle new filename format (dot separator)
- Update ID generation to include the configured prefix
- Update beans.toml schema to include `[beans]` section with `prefix` setting
- Update `beans init` to detect directory name and set default prefix
- Migrate existing beans by renaming files