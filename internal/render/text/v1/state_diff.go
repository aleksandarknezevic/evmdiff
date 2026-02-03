package text

import (
	"fmt"
	"strings"

	"github.com/aleksandarknezevic/evmdiff/internal/diff"
)

// RenderStateDiffText formats a state diff in human readable text.
func RenderStateDiffText(d diff.Diff) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("STATE DIFF: blocks %d -> %d\n\n", d.BlockFrom(), d.BlockTo()))

	// Balances
	if len(d.BalanceChanges()) > 0 {
		sb.WriteString("BALANCE CHANGES\n")
		for addr, delta := range d.BalanceChanges() {
			sb.WriteString(fmt.Sprintf("  %s: %s -> %s\n", addr, delta.Before, delta.After))
		}
		sb.WriteString("\n")
	}

	// Storage
	if len(d.StorageChanges()) > 0 {
		sb.WriteString("STORAGE CHANGES\n")
		for addr, slots := range d.StorageChanges() {
			sb.WriteString(fmt.Sprintf("  %s:\n", addr))
			for slot, delta := range slots {
				sb.WriteString(fmt.Sprintf("    %s: %s -> %s\n", slot, delta.Before, delta.After))
			}
			sb.WriteString("\n")
		}
	}

	// Code
	if len(d.CodeChanges()) > 0 {
		sb.WriteString("CODE CHANGES\n")
		for addr, delta := range d.CodeChanges() {
			sb.WriteString(fmt.Sprintf("  %s: %s -> %s\n", addr, delta.Before, delta.After))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
