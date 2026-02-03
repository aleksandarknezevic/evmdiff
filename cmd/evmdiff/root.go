package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "evmdiff",
	Short: "evmdiff - EVM state diff CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(txCmd)
	rootCmd.AddCommand(stateCmd)
}
