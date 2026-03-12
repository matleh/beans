<script lang="ts">
  import { fade } from 'svelte/transition';
  import { AgentActionsStore } from '$lib/agentActions.svelte';

  interface Props {
    beanId: string;
    agentBusy: boolean;
  }

  let { beanId, agentBusy }: Props = $props();

  const store = new AgentActionsStore();

  $effect(() => {
    store.fetch(beanId);
  });

  $effect(() => {
    store.notifyAgentStatus(beanId, agentBusy);
  });
</script>

{#if agentBusy}
  <div class="loader mr-2" transition:fade={{ duration: 200 }}></div>
{/if}
{#each store.actions as action (action.id)}
  <button
    class="btn-toggle btn-toggle-inactive ml-1"
    disabled={agentBusy || !!store.executingAction}
    title={action.description ?? undefined}
    onclick={() => store.execute(beanId, action.id)}
  >
    {action.label}
  </button>
{/each}
