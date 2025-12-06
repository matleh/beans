# Beans - Agentic Issue Tracker

This project uses **beans**, an agentic-first issue tracker. Issues are called "beans", and you can
use the "beans" CLI to manage them.

All commands support --json for machine-readable output. Use this flag to parse responses easily.

## Core Rules

- Track ALL work using beans (no TodoWrite tool, no markdown TODOs)
- Use `beans new` to create issues, not TodoWrite tool
- After compaction or clear, run `beans prompt` to re-sync
- When completing work, mark the bean as done using `beans status <bean-id> done`

## Finding work

- `beans list --json` to list all beans

## Creating new beans

- `beans new --help`
- When creating new beans, include a useful description. If you're not sure what to write, ask the user.
