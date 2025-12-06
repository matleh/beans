package beans

import (
	"fmt"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/output"
	"hmans.dev/beans/internal/ui"
)

var statusJSON bool

var statusCmd = &cobra.Command{
	Use:   "status <id> <status>",
	Short: "Change a bean's status",
	Long: `Changes the status of a bean.

Valid statuses: open, in-progress, done`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		newStatus := args[1]

		// Validate status
		if !validStatuses[newStatus] {
			if statusJSON {
				return output.Error(output.ErrInvalidStatus, fmt.Sprintf("invalid status: %s (must be open, in-progress, or done)", newStatus))
			}
			return fmt.Errorf("invalid status: %s (must be open, in-progress, or done)", newStatus)
		}

		// Find the bean
		b, err := store.FindByID(id)
		if err != nil {
			if statusJSON {
				return output.Error(output.ErrNotFound, err.Error())
			}
			return fmt.Errorf("failed to find bean: %w", err)
		}

		// Update the status
		oldStatus := b.Status
		b.Status = newStatus

		if err := store.Save(b); err != nil {
			if statusJSON {
				return output.Error(output.ErrFileError, err.Error())
			}
			return fmt.Errorf("failed to save bean: %w", err)
		}

		if statusJSON {
			return output.Success(b, "Status updated")
		}

		fmt.Printf("%s %s → %s\n",
			ui.ID.Render(b.ID),
			ui.Muted.Render(oldStatus),
			ui.RenderStatusText(newStatus),
		)
		return nil
	},
}

func init() {
	statusCmd.Flags().BoolVar(&statusJSON, "json", false, "Output as JSON")
	rootCmd.AddCommand(statusCmd)
}
