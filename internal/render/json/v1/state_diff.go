package v1

import (
	"github.com/aleksandarknezevic/evmdiff/internal/diff"
	"github.com/aleksandarknezevic/evmdiff/pkg/types"
)

type StateDiffOutput struct {
	Version string `json:"version"`
	Type    string `json:"type"`

	BlockFrom uint64 `json:"blockFrom"`
	BlockTo   uint64 `json:"blockTo"`

	Changes StateChanges `json:"changes"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type StateChanges struct {
	Balances map[string]types.Delta            `json:"balances,omitempty"`
	Storage  map[string]map[string]types.Delta `json:"storage,omitempty"`
	Code     map[string]types.Delta            `json:"code,omitempty"`
}

func RenderStateDiff(d diff.Diff) StateDiffOutput {
	return StateDiffOutput{
		Version:   "v1",
		Type:      "state_diff",
		BlockFrom: d.BlockFrom(),
		BlockTo:   d.BlockTo(),
		Changes: StateChanges{
			Balances: d.BalanceChanges(),
			Storage:  d.StorageChanges(),
			Code:     d.CodeChanges(),
		},
	}
}
