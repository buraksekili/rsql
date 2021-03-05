package client

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// displayDBStats shows the current db's status
func (c *DbClient) displayDBStats() {
	stats := c.db.Stats()
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "\n===== STATS =====")
	fmt.Fprintln(w, "Max Open Connections:\t", stats.MaxOpenConnections)
	fmt.Fprintln(w, "Open Connection:\t", stats.OpenConnections)
	fmt.Fprintln(w, "Idle:\t", stats.Idle)
	fmt.Fprintln(w, "In Use:\t", stats.InUse)
	w.Flush()
}
