package beancore

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWorktreeWatcher(t *testing.T) {
	t.Run("watches worktree beans dir and merges changes", func(t *testing.T) {
		core, _ := setupTestCore(t)

		// Create a bean in the main repo first
		createTestBean(t, core, "wt-test-1", "Original", "todo")

		// Start watching
		if err := core.StartWatching(); err != nil {
			t.Fatalf("StartWatching() error = %v", err)
		}
		defer core.Unwatch()

		// Create a fake worktree directory with a .beans/ subdir
		wtDir := t.TempDir()
		wtBeansDir := filepath.Join(wtDir, BeansDir)
		if err := os.MkdirAll(wtBeansDir, 0755); err != nil {
			t.Fatalf("failed to create worktree .beans dir: %v", err)
		}

		// Subscribe to events
		events, unsub := core.Subscribe()
		defer unsub()

		// Start watching the worktree
		if err := core.WatchWorktreeBeans(wtDir); err != nil {
			t.Fatalf("WatchWorktreeBeans() error = %v", err)
		}
		defer core.UnwatchWorktreeBeans(wtDir)

		// Write a modified version of the bean in the worktree
		content := `---
title: Updated in Worktree
status: in-progress
type: task
---

Working on this in a worktree.
`
		beanPath := filepath.Join(wtBeansDir, "wt-test-1--original.md")
		if err := os.WriteFile(beanPath, []byte(content), 0644); err != nil {
			t.Fatalf("failed to write worktree bean: %v", err)
		}

		// Wait for the event to propagate
		select {
		case batch := <-events:
			found := false
			for _, ev := range batch {
				if ev.BeanID == "wt-test-1" && ev.Type == EventUpdated {
					found = true
					if ev.Bean.Title != "Updated in Worktree" {
						t.Errorf("Title = %q, want %q", ev.Bean.Title, "Updated in Worktree")
					}
					if ev.Bean.Status != "in-progress" {
						t.Errorf("Status = %q, want %q", ev.Bean.Status, "in-progress")
					}
				}
			}
			if !found {
				t.Error("expected EventUpdated for wt-test-1")
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timed out waiting for worktree bean change event")
		}

		// Bean should be dirty (came from worktree, not persisted to main)
		if !core.IsDirty("wt-test-1") {
			t.Error("bean should be dirty after worktree update")
		}

		// In-memory state should reflect the worktree's version
		got, err := core.Get("wt-test-1")
		if err != nil {
			t.Fatalf("Get() error = %v", err)
		}
		if got.Title != "Updated in Worktree" {
			t.Errorf("Title = %q, want %q", got.Title, "Updated in Worktree")
		}
	})

	t.Run("does not crash when worktree has no .beans dir", func(t *testing.T) {
		core, _ := setupTestCore(t)

		wtDir := t.TempDir()
		// No .beans/ dir inside

		err := core.WatchWorktreeBeans(wtDir)
		if err != nil {
			t.Errorf("WatchWorktreeBeans() should return nil for missing .beans/ dir, got %v", err)
		}
	})

	t.Run("delete in worktree reverts to main-repo version", func(t *testing.T) {
		core, _ := setupTestCore(t)

		// Create a bean in the main repo
		createTestBean(t, core, "wt-del-1", "Original Title", "todo")

		// Create a worktree with a modified version of the bean
		wtDir := t.TempDir()
		wtBeansDir := filepath.Join(wtDir, BeansDir)
		os.MkdirAll(wtBeansDir, 0755)

		content := "---\ntitle: Modified in Worktree\nstatus: in-progress\ntype: task\n---\n"
		beanPath := filepath.Join(wtBeansDir, "wt-del-1--original-title.md")
		os.WriteFile(beanPath, []byte(content), 0644)

		// Start watching the worktree
		if err := core.WatchWorktreeBeans(wtDir); err != nil {
			t.Fatalf("WatchWorktreeBeans() error = %v", err)
		}
		defer core.UnwatchWorktreeBeans(wtDir)

		// Wait for initial load to merge the worktree version
		time.Sleep(100 * time.Millisecond)

		// Verify the worktree version was merged
		got, err := core.Get("wt-del-1")
		if err != nil {
			t.Fatalf("Get() error = %v", err)
		}
		if got.Title != "Modified in Worktree" {
			t.Errorf("Title = %q, want %q (worktree version should be active)", got.Title, "Modified in Worktree")
		}

		// Subscribe to events
		events, unsub := core.Subscribe()
		defer unsub()

		// Delete the bean from the worktree
		os.Remove(beanPath)

		// Should emit an Updated event reverting to the main-repo version
		select {
		case batch := <-events:
			found := false
			for _, ev := range batch {
				if ev.BeanID == "wt-del-1" && ev.Type == EventUpdated {
					found = true
					if ev.Bean.Title != "Original Title" {
						t.Errorf("reverted Title = %q, want %q", ev.Bean.Title, "Original Title")
					}
				}
			}
			if !found {
				t.Error("expected EventUpdated reverting to main-repo version")
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timed out waiting for revert event")
		}

		// Bean should still exist and be reverted to main-repo version
		got, err = core.Get("wt-del-1")
		if err != nil {
			t.Fatalf("bean should still exist after worktree delete, got error: %v", err)
		}
		if got.Title != "Original Title" {
			t.Errorf("Title = %q, want %q", got.Title, "Original Title")
		}

		// Should no longer be dirty
		if core.IsDirty("wt-del-1") {
			t.Error("bean should not be dirty after reverting to main-repo version")
		}
	})

	t.Run("delete worktree-only bean removes from runtime", func(t *testing.T) {
		core, _ := setupTestCore(t)

		// Create a worktree with a bean that doesn't exist in main
		wtDir := t.TempDir()
		wtBeansDir := filepath.Join(wtDir, BeansDir)
		os.MkdirAll(wtBeansDir, 0755)

		content := "---\ntitle: Worktree Only\nstatus: todo\ntype: task\n---\n"
		beanPath := filepath.Join(wtBeansDir, "wt-only-1--worktree-only.md")
		os.WriteFile(beanPath, []byte(content), 0644)

		// Start watching — initial scan loads the bean
		if err := core.WatchWorktreeBeans(wtDir); err != nil {
			t.Fatalf("WatchWorktreeBeans() error = %v", err)
		}
		defer core.UnwatchWorktreeBeans(wtDir)

		// Verify the bean was loaded
		got, err := core.Get("wt-only-1")
		if err != nil {
			t.Fatalf("bean should exist after initial load, got error: %v", err)
		}
		if got.Title != "Worktree Only" {
			t.Errorf("Title = %q, want %q", got.Title, "Worktree Only")
		}

		// Subscribe to events
		events, unsub := core.Subscribe()
		defer unsub()

		// Delete the bean from the worktree
		os.Remove(beanPath)

		// Should emit a Deleted event
		select {
		case batch := <-events:
			found := false
			for _, ev := range batch {
				if ev.BeanID == "wt-only-1" && ev.Type == EventDeleted {
					found = true
				}
			}
			if !found {
				t.Error("expected EventDeleted for worktree-only bean")
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timed out waiting for delete event")
		}

		// Bean should no longer exist
		if _, err := core.Get("wt-only-1"); err != ErrNotFound {
			t.Errorf("expected ErrNotFound after delete, got %v", err)
		}
	})

	t.Run("UnwatchAllWorktrees stops all watchers", func(t *testing.T) {
		core, _ := setupTestCore(t)

		// Create two worktrees
		for i := 0; i < 2; i++ {
			wtDir := t.TempDir()
			wtBeansDir := filepath.Join(wtDir, BeansDir)
			os.MkdirAll(wtBeansDir, 0755)
			if err := core.WatchWorktreeBeans(wtDir); err != nil {
				t.Fatalf("WatchWorktreeBeans() error = %v", err)
			}
		}

		core.UnwatchAllWorktrees()

		// Verify no worktree watchers remain
		core.mu.RLock()
		count := len(core.worktreeWatchers)
		core.mu.RUnlock()

		if count != 0 {
			t.Errorf("expected 0 worktree watchers after UnwatchAll, got %d", count)
		}
	})
}
