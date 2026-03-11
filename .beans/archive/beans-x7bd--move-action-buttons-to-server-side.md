---
# beans-x7bd
title: Move action buttons to server-side
status: completed
type: epic
priority: normal
created_at: 2026-03-11T19:11:38Z
updated_at: 2026-03-11T19:39:54Z
---

Move the web UI action buttons (Commit, Review) from being hardcoded in the frontend to being served by the backend via GraphQL. The backend becomes the single source of truth for available actions and their prompt templates.

## Design

- Backend provides available actions via a new GraphQL query
- Frontend's Status pane fetches actions when the agent finishes working (re-fetches on session status change)
- Clicking a button sends a mutation (executeAgentAction) back to the backend with the action ID
- Backend maps actionId to prompt, injects it into the agent conversation as a user message
- Frontend sees it arrive via the existing conversation subscription

## Goals
- Decouple frontend from knowing what actions exist or what prompts they generate
- Allow backend to conditionally show/hide actions based on bean state, agent status, etc.
- Keep it simple: no new subscription, just re-fetch on agent status change

## Summary of Changes

Action buttons are now fully server-driven:
- Backend: agentActions query + executeAgentAction mutation with consolidated action registry
- Frontend: ChangesPane fetches actions dynamically, fires mutation on click
- Removed: ActionButton.svelte, actionContext.ts, hardcoded prompts in frontend
