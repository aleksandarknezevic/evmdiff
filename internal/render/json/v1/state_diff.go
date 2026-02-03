package v1

import (
	"github.com/aleksandarknezevic/evmdiff/internal/diff"
)

// StateDiffOutput is the v1 JSON format for state diffs.
type StateDiffOutput struct {
	Version string `json:"version"` // "v1"
	Type    string `json:"type"`    // "state_diff"

	BlockFrom uint64 `json:"blockFrom"`
	BlockTo   uint64 `json:"blockTo"`

	Changes StateChanges `json:"changes"`

	// Optional: additional metadata can be added later without breaking v1.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type StateChanges struct {
	Balances map[string]Delta            `json:"balances,omitempty"`
	Storage  map[string]map[string]Delta `json:"storage,omitempty"`
	Code     map[string]Delta            `json:"code,omitempty"`
}

// RenderStateDiff converts a diff.Diff into the v1 JSON format.
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
