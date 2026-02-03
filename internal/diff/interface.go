package diff

import (
	"github.com/aleksandarknezevic/evmdiff/pkg/types"
)

type Diff interface {
	BlockFrom() uint64
	BlockTo() uint64

	BalanceChanges() map[types.Address]types.Delta
	StorageChanges() map[types.Address]map[types.Slot]types.Delta
	CodeChanges() map[types.Address]types.Delta

	// Optional helpers
	HasChanges() bool
	IsEmpty() bool
}
