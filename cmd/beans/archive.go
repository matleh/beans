package beans

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/output"
)

var (
	archiveForce bool
	archiveJSON  bool
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Delete all beans with status 'done'",
	Long:  `Deletes all beans that have their status set to "done". Asks for confirmation unless --force is provided.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		beans, err := store.FindAll()
		if err != nil {
			if archiveJSON {
				return output.Error(output.ErrFileError, err.Error())
			}
			return fmt.Errorf("failed to list beans: %w", err)
		}

		// Find beans with status "done"
		var doneBeans []string
		for _, b := range beans {
			if b.Status == "done" {
				doneBeans = append(doneBeans, b.ID)
			}
		}

		if len(doneBeans) == 0 {
			if archiveJSON {
				return output.SuccessMessage("No beans to archive")
			}
			fmt.Println("No beans with status 'done' to archive.")
			return nil
		}

		// JSON implies force (no prompts for machines)
		if !archiveForce && !archiveJSON {
			var confirm bool
			err := huh.NewConfirm().
				Title(fmt.Sprintf("Archive %d bean(s) with status 'done'?", len(doneBeans))).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm).
				Run()

			if err != nil {
				return err
			}

			if !confirm {
				fmt.Println("Cancelled")
				return nil
			}
		}

		// Delete all done beans
		var deleted []string
		for _, id := range doneBeans {
			if err := store.Delete(id); err != nil {
				if archiveJSON {
					return output.Error(output.ErrFileError, fmt.Sprintf("failed to delete bean %s: %s", id, err.Error()))
				}
				return fmt.Errorf("failed to delete bean %s: %w", id, err)
			}
			deleted = append(deleted, id)
		}

		if archiveJSON {
			return output.SuccessMessage(fmt.Sprintf("Archived %d bean(s)", len(deleted)))
		}

		fmt.Printf("Archived %d bean(s)\n", len(deleted))
		return nil
	},
}

func init() {
	archiveCmd.Flags().BoolVarP(&archiveForce, "force", "f", false, "Skip confirmation")
	archiveCmd.Flags().BoolVar(&archiveJSON, "json", false, "Output as JSON (implies --force)")
	rootCmd.AddCommand(archiveCmd)
}
