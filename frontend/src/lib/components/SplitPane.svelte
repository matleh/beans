<script lang="ts">
  import type { Snippet } from 'svelte';

  interface Props {
    direction?: 'horizontal' | 'vertical';
    side?: 'start' | 'end';
    initialSize?: number;
    minSize?: number;
    maxSize?: number;
    persistKey?: string;
    collapsed?: boolean;
    children: Snippet;
    aside: Snippet;
  }

  let {
    direction = 'horizontal',
    side = 'end',
    initialSize = 350,
    minSize = 40,
    maxSize,
    persistKey,
    collapsed = false,
    children,
    aside
  }: Props = $props();

  let size = $state(getInitialSize());

  function getInitialSize(): number {
    return initialSize;
  }
  let isDragging = $state(false);
  let containerEl: HTMLDivElement | undefined = $state();

  // Load persisted size on mount
  $effect(() => {
    if (persistKey) {
      const saved = localStorage.getItem(`beans-split-${persistKey}`);
      if (saved) {
        const parsed = parseInt(saved, 10);
        if (!Number.isNaN(parsed)) {
          size = Math.max(minSize, maxSize ? Math.min(parsed, maxSize) : parsed);
        }
      }
    }
  });

  function startDrag(e: MouseEvent) {
    if (collapsed) return;
    isDragging = true;
    e.preventDefault();
  }

  function onDrag(e: MouseEvent) {
    if (!isDragging || !containerEl) return;
    const rect = containerEl.getBoundingClientRect();
    const isHorizontal = direction === 'horizontal';
    const mousePos = isHorizontal ? e.clientX : e.clientY;
    const containerStart = isHorizontal ? rect.left : rect.top;
    const containerEnd = isHorizontal ? rect.right : rect.bottom;
    const containerSize = containerEnd - containerStart;

    let newSize: number;
    if (side === 'start') {
      newSize = mousePos - containerStart;
    } else {
      newSize = containerEnd - mousePos;
    }

    // Clamp: aside pane gets at least minSize, and leave minSize for the main pane too
    const upperBound = maxSize
      ? Math.min(maxSize, containerSize - minSize)
      : containerSize - minSize;
    size = Math.max(minSize, Math.min(upperBound, newSize));
  }

  function stopDrag() {
    if (isDragging) {
      isDragging = false;
      if (persistKey) {
        localStorage.setItem(`beans-split-${persistKey}`, size.toString());
      }
    }
  }

  const isHorizontal = $derived(direction === 'horizontal');
  const isReversed = $derived(side === 'end');
  const displaySize = $derived(collapsed ? 0 : size);
</script>

<svelte:window onmousemove={onDrag} onmouseup={stopDrag} />

<div
  bind:this={containerEl}
  class={[
    'flex min-h-0 min-w-0 flex-1',
    isHorizontal ? (isReversed ? 'flex-row-reverse' : 'flex-row') : isReversed ? 'flex-col-reverse' : 'flex-col'
  ]}
>
  <!-- Fixed-size aside pane -->
  <div
    class="flex shrink-0 flex-col overflow-hidden"
    style="{isHorizontal ? 'width' : 'height'}: {displaySize}px"
  >
    {@render aside()}
  </div>

  <!-- Resize handle -->
  {#if !collapsed}
    <div
      class={[
        'shrink-0 transition-colors',
        isHorizontal ? 'w-px cursor-col-resize' : 'h-1 cursor-row-resize',
        isDragging ? 'bg-surface-dim' : 'bg-border hover:bg-surface-dim'
      ]}
      role="slider"
      aria-orientation={isHorizontal ? 'horizontal' : 'vertical'}
      aria-valuenow={size}
      aria-valuemin={minSize}
      aria-valuemax={maxSize ?? 999}
      tabindex="0"
      onmousedown={startDrag}
    ></div>
  {/if}

  <!-- Main content pane -->
  <div class="flex min-h-0 min-w-0 flex-1 flex-col">
    {@render children()}
  </div>
</div>
