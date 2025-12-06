package beans

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	forceDelete bool
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Aliases: []string{"rm"},
	Short:   "Delete a bean",
	Long:    `Deletes a bean after confirmation (use -f to skip confirmation).`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := store.FindByID(args[0])
		if err != nil {
			return fmt.Errorf("failed to find bean: %w", err)
		}

		if !forceDelete {
			fmt.Printf("Delete '%s' (%s)? [y/N] ", b.Title, b.Path)
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))

			if response != "y" && response != "yes" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		if err := store.Delete(args[0]); err != nil {
			return fmt.Errorf("failed to delete bean: %w", err)
		}

		fmt.Printf("Deleted %s\n", b.Path)
		return nil
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation")
	rootCmd.AddCommand(deleteCmd)
}
