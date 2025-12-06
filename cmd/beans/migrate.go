package beans

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/bean"
	"hmans.dev/beans/internal/output"
	"hmans.dev/beans/internal/ui"
)

var (
	migrateJSON   bool
	migrateDryRun bool
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate beans to new filename format",
	Long:  `Renames existing bean files from the legacy format (id-slug.md) to the new format (id.slug.md).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		beans, err := store.FindAll()
		if err != nil {
			if migrateJSON {
				return output.Error(output.ErrFileError, err.Error())
			}
			return fmt.Errorf("failed to list beans: %w", err)
		}

		var renamed []map[string]string
		var skipped int

		for _, b := range beans {
			oldPath := b.Path
			oldFilename := filepath.Base(oldPath)
			nameWithoutExt := strings.TrimSuffix(oldFilename, ".md")

			// Check if already in new format (has double-dash)
			if strings.Contains(nameWithoutExt, "--") {
				skipped++
				continue
			}

			// Skip if no slug (ID-only files don't need migration)
			if b.Slug == "" {
				skipped++
				continue
			}

			// Build new filename
			newFilename := bean.BuildFilename(b.ID, b.Slug)
			if oldFilename == newFilename {
				skipped++
				continue
			}

			// Build new path (preserve directory structure)
			dir := filepath.Dir(oldPath)
			var newPath string
			if dir == "." {
				newPath = newFilename
			} else {
				newPath = filepath.Join(dir, newFilename)
			}

			oldFullPath := filepath.Join(store.Root, oldPath)
			newFullPath := filepath.Join(store.Root, newPath)

			if migrateDryRun {
				if !migrateJSON {
					fmt.Printf("Would rename: %s -> %s\n", oldPath, newPath)
				}
				renamed = append(renamed, map[string]string{
					"old": oldPath,
					"new": newPath,
				})
				continue
			}

			// Perform the rename
			if err := os.Rename(oldFullPath, newFullPath); err != nil {
				if migrateJSON {
					return output.Error(output.ErrFileError, fmt.Sprintf("failed to rename %s: %v", oldPath, err))
				}
				return fmt.Errorf("failed to rename %s: %w", oldPath, err)
			}

			renamed = append(renamed, map[string]string{
				"old": oldPath,
				"new": newPath,
			})

			if !migrateJSON {
				fmt.Println(ui.Muted.Render(oldPath) + " → " + ui.Success.Render(newPath))
			}
		}

		if migrateJSON {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{
				"success": true,
				"renamed": renamed,
				"skipped": skipped,
				"dry_run": migrateDryRun,
			})
		}

		if len(renamed) == 0 {
			fmt.Println("No beans to migrate")
		} else if migrateDryRun {
			fmt.Printf("\nDry run: %d beans would be renamed\n", len(renamed))
		} else {
			fmt.Printf("\nMigrated %d beans\n", len(renamed))
		}

		return nil
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&migrateJSON, "json", false, "Output as JSON")
	migrateCmd.Flags().BoolVar(&migrateDryRun, "dry-run", false, "Show what would be renamed without making changes")
	rootCmd.AddCommand(migrateCmd)
}
