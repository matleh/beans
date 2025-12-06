package beans

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/bean"
)

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new bean",
	Long:  `Creates a new bean (issue) with a generated ID and optional title.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		if title == "" {
			title = "Untitled"
		}

		b := &bean.Bean{
			ID:     bean.NewID(),
			Slug:   bean.Slugify(title),
			Title:  title,
			Status: "open",
		}

		if err := store.Save(b); err != nil {
			return fmt.Errorf("failed to create bean: %w", err)
		}

		fmt.Printf("Created %s\n", b.Path)

		// Open in editor if $EDITOR is set
		editor := os.Getenv("EDITOR")
		if editor != "" {
			path := store.FullPath(b)
			editorCmd := exec.Command(editor, path)
			editorCmd.Stdin = os.Stdin
			editorCmd.Stdout = os.Stdout
			editorCmd.Stderr = os.Stderr
			if err := editorCmd.Run(); err != nil {
				return fmt.Errorf("editor failed: %w", err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
