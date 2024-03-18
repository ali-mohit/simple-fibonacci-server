package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(getServeCommand())
	rootCmd.AddCommand(getConsoleFibonacciCommand())
	rootCmd.AddCommand(getConsoleFibonacciV2Command())
	rootCmd.AddCommand(getVersionCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
