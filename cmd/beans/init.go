package beans

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/bean"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a .beans directory",
	Long:  `Creates a .beans directory in the current directory for storing issues.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		if err := bean.Init(dir); err != nil {
			return fmt.Errorf("failed to initialize: %w", err)
		}

		fmt.Println("Initialized .beans directory")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
