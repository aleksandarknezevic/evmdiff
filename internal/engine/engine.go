package engine

import (
	"github.com/aleksandarknezevic/evmdiff/internal/diff"
	"github.com/aleksandarknezevic/evmdiff/internal/snapshot"
	"github.com/aleksandarknezevic/evmdiff/pkg/types"
)

type engineImpl struct{}

func NewEngine() Engine {
	return &engineImpl{}
}

// Snapshot returns a fake snapshot (placeholder)
func (e *engineImpl) Snapshot(block uint64, addresses []string) (snapshot.Snapshot, error) {
	// placeholder data
	balances := make(map[types.Address]string)
	storage := make(map[types.Address]map[types.Slot]types.Word)
	code := make(map[types.Address]types.Hash)

	for _, a := range addresses {
		balances[types.Address(a)] = "0"
		storage[types.Address(a)] = map[types.Slot]types.Word{}
		code[types.Address(a)] = ""
	}

	return snapshot.NewSnapshot(block, balances, storage, code), nil
}

// Diff returns a fake diff (placeholder)
func (e *engineImpl) Diff(fromBlock uint64, toBlock uint64, addresses []string) (diff.Diff, error) {
	// placeholder diff data
	balances := make(map[types.Address]types.Delta)
	storage := make(map[types.Address]map[types.Slot]types.Delta)
	code := make(map[types.Address]types.Delta)

	return diff.NewDiff(fromBlock, toBlock, balances, storage, code), nil
}
