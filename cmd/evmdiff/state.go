package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aleksandarknezevic/evmdiff/internal/engine"
)

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "State related commands",
}

var stateInspectCmd = &cobra.Command{
	Use:   "inspect [address] [block]",
	Short: "Inspect state at a block",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := args[0]
		// parse block from args[1]...

		e := engine.NewEngine()
		_, err := e.Snapshot(0, []string{addr})
		if err != nil {
			return err
		}

		fmt.Println("state inspect placeholder")
		return nil
	},
}

var stateDiffCmd = &cobra.Command{
	Use:   "diff [address] [from] [to]",
	Short: "Diff state between blocks",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		// parse args...

		e := engine.NewEngine()
		_, err := e.Diff(0, 0, []string{args[0]})
		if err != nil {
			return err
		}

		fmt.Println("state diff placeholder")
		return nil
	},
}

func init() {
	stateCmd.AddCommand(stateInspectCmd)
	stateCmd.AddCommand(stateDiffCmd)
}
