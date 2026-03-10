package gitutil

import (
	"os/exec"
	"path/filepath"
	"testing"
)

// initTestRepo creates a temporary git repo with an initial commit.
func initTestRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	commands := [][]string{
		{"git", "init"},
		{"git", "config", "user.email", "test@test.com"},
		{"git", "config", "user.name", "Test"},
		{"git", "commit", "--allow-empty", "-m", "initial"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v failed: %s: %v", args, out, err)
		}
	}

	return dir
}

func TestMainWorktreeRoot_MainWorktree(t *testing.T) {
	repoDir := initTestRepo(t)

	root, isSecondary := MainWorktreeRoot(repoDir)
	if isSecondary {
		t.Errorf("expected main worktree, got secondary with root %q", root)
	}
	if root != "" {
		t.Errorf("expected empty root, got %q", root)
	}
}

func TestMainWorktreeRoot_SecondaryWorktree(t *testing.T) {
	repoDir := initTestRepo(t)

	// Create a secondary worktree
	wtPath := filepath.Join(t.TempDir(), "secondary")
	cmd := exec.Command("git", "worktree", "add", wtPath, "-b", "test-branch")
	cmd.Dir = repoDir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git worktree add failed: %s: %v", out, err)
	}

	root, isSecondary := MainWorktreeRoot(wtPath)
	if !isSecondary {
		t.Fatal("expected secondary worktree, got main")
	}

	// Resolve repoDir to handle symlinks (e.g., /tmp -> /private/tmp on macOS)
	expectedRoot, err := filepath.EvalSymlinks(repoDir)
	if err != nil {
		t.Fatalf("EvalSymlinks: %v", err)
	}
	actualRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatalf("EvalSymlinks: %v", err)
	}

	if actualRoot != expectedRoot {
		t.Errorf("got root %q, want %q", actualRoot, expectedRoot)
	}
}

func TestMainWorktreeRoot_NotGitRepo(t *testing.T) {
	dir := t.TempDir()

	root, isSecondary := MainWorktreeRoot(dir)
	if isSecondary {
		t.Errorf("expected not secondary, got secondary with root %q", root)
	}
	if root != "" {
		t.Errorf("expected empty root, got %q", root)
	}
}
