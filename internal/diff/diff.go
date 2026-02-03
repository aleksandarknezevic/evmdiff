package diff

import "github.com/aleksandarknezevic/evmdiff/pkg/types"

type diffImpl struct {
	blockFrom uint64
	blockTo   uint64

	balances map[types.Address]types.Delta
	storage  map[types.Address]map[types.Slot]types.Delta
	code     map[types.Address]types.Delta
}

func NewDiff(
	from, to uint64,
	balances map[types.Address]types.Delta,
	storage map[types.Address]map[types.Slot]types.Delta,
	code map[types.Address]types.Delta,
) Diff {
	return &diffImpl{
		blockFrom: from,
		blockTo:   to,
		balances:  balances,
		storage:   storage,
		code:      code,
	}
}

func (d *diffImpl) BlockFrom() uint64 {
	return d.blockFrom
}

func (d *diffImpl) BlockTo() uint64 {
	return d.blockTo
}

func (d *diffImpl) BalanceChanges() map[types.Address]types.Delta {
	return d.balances
}

func (d *diffImpl) StorageChanges() map[types.Address]map[types.Slot]types.Delta {
	return d.storage
}

func (d *diffImpl) CodeChanges() map[types.Address]types.Delta {
	return d.code
}

func (d *diffImpl) HasChanges() bool {
	return !d.IsEmpty()
}

func (d *diffImpl) IsEmpty() bool {
	if len(d.balances) > 0 {
		return false
	}
	if len(d.code) > 0 {
		return false
	}
	if len(d.storage) > 0 {
		return false
	}
	return true
}
