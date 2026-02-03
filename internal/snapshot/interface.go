package snapshot

import "github.com/aleksandarknezevic/evmdiff/pkg/types"

type Snapshot interface {
	BlockNumber() uint64
	Balances() map[types.Address]string
	Storage() map[types.Address]map[types.Slot]types.Word
	CodeHash() map[types.Address]types.Hash
}
