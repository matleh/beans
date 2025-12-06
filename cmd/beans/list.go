package beans

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all beans",
	Long:    `Lists all beans in the .beans directory, grouped by path.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		beans, err := store.FindAll()
		if err != nil {
			return fmt.Errorf("failed to list beans: %w", err)
		}

		if len(beans) == 0 {
			fmt.Println("No beans found")
			return nil
		}

		// Sort by path for grouping
		sort.Slice(beans, func(i, j int) bool {
			return beans[i].Path < beans[j].Path
		})

		table := tablewriter.NewTable(os.Stdout,
			tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
				Borders: tw.BorderNone,
				Settings: tw.Settings{
					Separators: tw.Separators{
						BetweenRows:    tw.Off,
						BetweenColumns: tw.Off,
					},
				},
			})),
			tablewriter.WithRowAlignment(tw.AlignLeft),
			tablewriter.WithHeaderAlignment(tw.AlignLeft),
			tablewriter.WithPadding(tw.Padding{Left: "", Right: "  "}),
		)

		table.Header("ID", "Status", "Path", "Title")

		for _, b := range beans {
			// Show directory path without filename for cleaner output
			dir := filepath.Dir(b.Path)
			if dir == "." {
				dir = ""
			}

			table.Append(b.ID, b.Status, dir, truncate(b.Title, 50))
		}

		table.Render()
		return nil
	},
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func init() {
	rootCmd.AddCommand(listCmd)
}
