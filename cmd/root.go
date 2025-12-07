package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/beancore"
	"hmans.dev/beans/internal/config"
)

var core *beancore.Core
var cfg *config.Config
var beansPath string

var rootCmd = &cobra.Command{
	Use:   "beans",
	Short: "A file-based issue tracker for AI-first workflows",
	Long: `Beans is a lightweight issue tracker that stores issues as markdown files.
Track your work alongside your code and supercharge your coding agent with
a full view of your project.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip core initialization for init command
		if cmd.Name() == "init" {
			return nil
		}

		var root string
		var err error

		if beansPath != "" {
			// Use explicit path
			root = beansPath
			// Verify it exists
			if info, statErr := os.Stat(root); statErr != nil || !info.IsDir() {
				return fmt.Errorf("beans path does not exist or is not a directory: %s", root)
			}
		} else {
			// Search upward for .beans directory
			root, err = beancore.FindRoot()
			if err != nil {
				return fmt.Errorf("no .beans directory found (run 'beans init' to create one)")
			}
		}

		cfg, err = config.Load(root)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		core = beancore.New(root, cfg)
		if err := core.Load(); err != nil {
			return fmt.Errorf("loading beans: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&beansPath, "beans-path", "", "Path to data directory (default: searches upward for .beans/)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
