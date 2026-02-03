package v1

import (
	"github.com/aleksandarknezevic/evmdiff/internal/diff"
	"github.com/aleksandarknezevic/evmdiff/pkg/types"
)

type TxDiffOutput struct {
	Version string `json:"version"`
	Type    string `json:"type"`

	TxHash string `json:"txHash"`
	Block  uint64 `json:"block"`

	Changes TxChanges `json:"changes"`

	Events      []Event  `json:"events,omitempty"`
	CallSummary []string `json:"callSummary,omitempty"`
}

type TxChanges struct {
	Balances map[string]types.Delta            `json:"balances,omitempty"`
	Storage  map[string]map[string]types.Delta `json:"storage,omitempty"`
	Code     map[string]types.Delta            `json:"code,omitempty"`
}

type Event struct {
	Address string                 `json:"address"`
	Name    string                 `json:"name"`
	Args    map[string]interface{} `json:"args"`
}

func RenderTxDiff(d diff.Diff, txHash string) TxDiffOutput {
	return TxDiffOutput{
		Version: "v1",
		Type:    "tx_diff",
		TxHash:  txHash,
		Block:   d.BlockTo(),
		Changes: TxChanges{
			Balances: d.BalanceChanges(),
			Storage:  d.StorageChanges(),
			Code:     d.CodeChanges(),
		},
	}
}
