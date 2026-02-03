package v1

import "github.com/aleksandarknezevic/evmdiff/internal/diff"

type TxDiffOutput struct {
	Version string `json:"version"` // "v1"
	Type    string `json:"type"`    // "tx_diff"

	TxHash string `json:"txHash"`
	Block  uint64 `json:"block"`

	Changes TxChanges `json:"changes"`

	// Optional fields
	Events      []Event  `json:"events,omitempty"`
	CallSummary []string `json:"callSummary,omitempty"`
}

type TxChanges struct {
	Balances map[string]Delta            `json:"balances,omitempty"`
	Storage  map[string]map[string]Delta `json:"storage,omitempty"`
	Code     map[string]Delta            `json:"code,omitempty"`
}

type Event struct {
	Address string                 `json:"address"`
	Name    string                 `json:"name"`
	Args    map[string]interface{} `json:"args"`
}

// RenderTxDiff converts diff.Diff into v1 JSON format.
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
