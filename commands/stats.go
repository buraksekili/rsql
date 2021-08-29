package commands

import (
	"database/sql"
	"fmt"
	"io"
	"text/tabwriter"
)

// displayDBStats shows the current db's status
func (c *DbClient) displayDBStats(w io.Writer) {
	stats := getDBStatus(c.db)
	tw := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	tw.Init(w, 0, 8, 0, '\t', 0)
	fmt.Fprintln(tw, "\n===== STATS =====")
	fmt.Fprintln(tw, "Max Open Connections:\t", stats.MaxOpenConnections)
	fmt.Fprintln(tw, "Open Connection:\t", stats.OpenConnections)
	fmt.Fprintln(tw, "Idle:\t", stats.Idle)
	fmt.Fprintln(tw, "In Use:\t", stats.InUse)
	tw.Flush()
}

func getDBStatus(DB *sql.DB) sql.DBStats {
	return DB.Stats()
}
