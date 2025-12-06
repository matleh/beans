package beans

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Edit a bean in your editor",
	Long:  `Opens a bean in your $EDITOR for editing.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			return fmt.Errorf("$EDITOR not set")
		}

		b, err := store.FindByID(args[0])
		if err != nil {
			return fmt.Errorf("failed to find bean: %w", err)
		}

		path := store.FullPath(b)
		editorCmd := exec.Command(editor, path)
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		if err := editorCmd.Run(); err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
