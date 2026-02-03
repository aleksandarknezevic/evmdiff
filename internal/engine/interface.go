package engine

import (
	"github.com/aleksandarknezevic/evmdiff/internal/diff"
	"github.com/aleksandarknezevic/evmdiff/internal/snapshot"
)

type Engine interface {
	Snapshot(block uint64, addresses []string) (snapshot.Snapshot, error)
	Diff(fromBlock uint64, toBlock uint64, addresses []string) (diff.Diff, error)
}
