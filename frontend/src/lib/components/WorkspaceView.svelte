<script lang="ts">
  import { AgentChatStore } from '$lib/agentChat.svelte';
  import { ui } from '$lib/uiState.svelte';
  import { worktreeStore } from '$lib/worktrees.svelte';
  import { onDestroy } from 'svelte';
  import SplitPane from './SplitPane.svelte';
  import AgentChat from './AgentChat.svelte';
  import ChangesPane from './ChangesPane.svelte';
  import PaneHeader from './PaneHeader.svelte';
  import TerminalPane from './TerminalPane.svelte';
  import ViewToolbar from './ViewToolbar.svelte';

  interface Props {
    worktreeId: string;
  }

  let { worktreeId }: Props = $props();

  const agentStore = new AgentChatStore();

  $effect(() => {
    agentStore.subscribe(worktreeId);
  });

  onDestroy(() => {
    agentStore.unsubscribe();
  });

  const agentBusy = $derived(agentStore.session?.status === 'RUNNING');

  const worktreePath = $derived(
    worktreeStore.worktrees.find((wt) => wt.id === worktreeId)?.path
  );
</script>

{#snippet changesPanel()}
  <ChangesPane path={worktreePath} beanId={worktreeId} {agentBusy} />
{/snippet}

{#snippet agentChatPanel()}
  <div class="flex h-full flex-col bg-surface">
    <PaneHeader title="Agent" />
    <div class="min-h-0 flex-1">
      <AgentChat beanId={worktreeId} store={agentStore} />
    </div>
  </div>
{/snippet}

{#snippet terminalPanel()}
  {#if ui.terminalInitialized}
    <TerminalPane sessionId={worktreeId} onClose={() => ui.toggleTerminal()} />
  {/if}
{/snippet}

{#snippet mainContent()}
  <SplitPane
    direction="horizontal"
    panels={[
      { content: changesPanel, size: 420, collapsed: !ui.showChanges, persistKey: 'workspace-changes' },
      { content: agentChatPanel }
    ]}
  />
{/snippet}

<div class="flex h-full flex-col">
  <ViewToolbar />

  <div class="flex min-h-0 flex-1 flex-col">
    <SplitPane
      direction="vertical"
      panels={[
        { content: mainContent },
        { content: terminalPanel, size: 300, collapsed: !ui.showTerminal, persistKey: 'workspace-terminal' }
      ]}
    />
  </div>
</div>
