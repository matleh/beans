package beans

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show a bean's contents",
	Long:  `Displays the full contents of a bean, including front matter and body.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := store.FindByID(args[0])
		if err != nil {
			return fmt.Errorf("failed to find bean: %w", err)
		}

		content, err := b.Render()
		if err != nil {
			return fmt.Errorf("failed to render bean: %w", err)
		}

		fmt.Print(string(content))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
