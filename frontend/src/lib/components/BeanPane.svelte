<script lang="ts">
  import type { Bean } from '$lib/beans.svelte';
  import BeanDetail from './BeanDetail.svelte';
  import AgentChat from './AgentChat.svelte';

  interface Props {
    bean: Bean;
    onSelect?: (bean: Bean) => void;
    onEdit?: (bean: Bean) => void;
    onClose?: () => void;
  }

  let { bean, onSelect, onEdit, onClose }: Props = $props();

  const hasWorktree = $derived(!!bean.worktreeId);

  // Track explicit tab selection per bean; activeTab derives from it
  let tabSelection = $state<{ beanId: string; tab: 'bean' | 'chat' } | null>(null);
  const activeTab = $derived.by(() => {
    if (tabSelection?.beanId === bean.id) {
      // Fall back to 'bean' if chat tab selected but worktree was removed
      if (tabSelection.tab === 'chat' && !hasWorktree) return 'bean';
      return tabSelection.tab;
    }
    return 'bean';
  });

  function setTab(tab: 'bean' | 'chat') {
    tabSelection = { beanId: bean.id, tab };
  }
</script>

<div class="flex h-full flex-col bg-surface">
  <div class="toolbar">
    <div class="flex">
      <button
        onclick={() => setTab('bean')}
        class={[
          'btn-tab',
          hasWorktree ? 'rounded-l-md' : 'rounded-md',
          activeTab === 'bean' ? 'btn-tab-active' : 'btn-tab-inactive'
        ]}
      >
        Bean
      </button>
      {#if hasWorktree}
        <button
          onclick={() => setTab('chat')}
          class={[
            'btn-tab rounded-r-md border-l-0',
            activeTab === 'chat' ? 'btn-tab-active' : 'btn-tab-inactive'
          ]}
        >
          Chat
        </button>
      {/if}
    </div>
    {#if onClose}
      <div class="flex-1"></div>
      <button onclick={onClose} class="btn-icon" title="Close"> &#x2715; </button>
    {/if}
  </div>

  <div class="min-h-0 flex-1">
    {#if activeTab === 'bean'}
      <BeanDetail {bean} {onSelect} {onEdit} />
    {:else if activeTab === 'chat' && bean.worktreeId}
      <AgentChat beanId={bean.worktreeId} />
    {/if}
  </div>
</div>
