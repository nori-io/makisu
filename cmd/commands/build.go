package commands

import (
	"fmt"
	"github.com/nori-io/makisu/internal/build"
	"github.com/spf13/cobra"
	"os"
)

func buildCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "build",
		Run: func(cmd *cobra.Command, args []string) {
			// get plugin dir path to build
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				os.Exit(1)
			}

			err = build.Build(pwd)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		},
	}
}