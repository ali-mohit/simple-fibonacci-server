package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "print the version info",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
}

func printVersion() {
	fmt.Println("v1.0.0")
}
