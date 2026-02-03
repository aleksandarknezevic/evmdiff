package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aleksandarknezevic/evmdiff/internal/engine"
	"github.com/aleksandarknezevic/evmdiff/internal/render/json/v1"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Transaction related commands",
}

var txDiffCmd = &cobra.Command{
	Use:   "diff [txHash]",
	Short: "Diff state changes caused by a tx",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		txHash := args[0]

		// TODO: resolve tx -> blocks & addresses
		// For now, use placeholders
		e := engine.NewEngine()
		d, err := e.Diff(0, 0, []string{})
		if err != nil {
			return err
		}

		out := v1.RenderTxDiff(d, txHash)
		fmt.Println(out)

		return nil
	},
}

func init() {
	txCmd.AddCommand(txDiffCmd)
}
