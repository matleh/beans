// Package gitutil provides git-related utility functions.
package gitutil

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// MainWorktreeRoot returns the root directory of the main git worktree
// if the given directory is inside a secondary worktree.
// Returns ("", false) if in the main worktree, not a git repo, or git unavailable.
func MainWorktreeRoot(dir string) (string, bool) {
	commonDir, err := gitRevParse(dir, "--git-common-dir")
	if err != nil {
		return "", false
	}

	gitDir, err := gitRevParse(dir, "--git-dir")
	if err != nil {
		return "", false
	}

	// Resolve to absolute paths for reliable comparison.
	// git rev-parse may return relative or absolute paths.
	commonDir = resolveGitPath(dir, commonDir)
	gitDir = resolveGitPath(dir, gitDir)

	// If they're the same, we're in the main worktree
	if commonDir == gitDir {
		return "", false
	}

	// Main repo root is the parent of the common .git directory
	return filepath.Dir(commonDir), true
}

// resolveGitPath makes a git path absolute. If the path is already absolute,
// it's cleaned and returned as-is. Otherwise it's joined with the base dir.
func resolveGitPath(base, p string) string {
	if filepath.IsAbs(p) {
		return filepath.Clean(p)
	}
	return filepath.Join(base, p)
}

func gitRevParse(dir, flag string) (string, error) {
	cmd := exec.Command("git", "-C", dir, "rev-parse", flag)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
