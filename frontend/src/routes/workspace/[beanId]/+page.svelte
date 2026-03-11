<script lang="ts">
  import { page } from '$app/state';
  import { beansStore } from '$lib/beans.svelte';
  import WorkspaceView from '$lib/components/WorkspaceView.svelte';

  const worktreeId = $derived(page.params.beanId);
  const bean = $derived(beansStore.get(worktreeId) ?? undefined);
</script>

{#key worktreeId}
  <WorkspaceView {worktreeId} {bean} />
{/key}
