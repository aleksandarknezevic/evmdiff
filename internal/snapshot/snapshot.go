package snapshot

import "github.com/aleksandarknezevic/evmdiff/pkg/types"

type snapshotImpl struct {
	blockNumber uint64
	balances    map[types.Address]string
	storage     map[types.Address]map[types.Slot]types.Word
	codeHash    map[types.Address]types.Hash
}

func NewSnapshot(
	block uint64,
	balances map[types.Address]string,
	storage map[types.Address]map[types.Slot]types.Word,
	code map[types.Address]types.Hash,
) Snapshot {
	return &snapshotImpl{
		blockNumber: block,
		balances:    balances,
		storage:     storage,
		codeHash:    code,
	}
}

func (s *snapshotImpl) BlockNumber() uint64 {
	return s.blockNumber
}

func (s *snapshotImpl) Balances() map[types.Address]string {
	return s.balances
}

func (s *snapshotImpl) Storage() map[types.Address]map[types.Slot]types.Word {
	return s.storage
}

func (s *snapshotImpl) CodeHash() map[types.Address]types.Hash {
	return s.codeHash
}
