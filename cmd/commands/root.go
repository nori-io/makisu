package commands

import (
	"github.com/spf13/cobra"
	"os"
)

// root command
var rootCmd = &cobra.Command{
	Use:   "makisu [command]",
	Short: "v0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(buildCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}