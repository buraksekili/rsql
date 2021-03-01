package client

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/olekukonko/tablewriter"
)

// OpenConnection opens connection, pings the DB.
// In case of any error while pinging, the app prints fatal
// error and exits.
//
// If there is no error after pinging the DB, this function takes
// input to execute on the DB specified by the user.
func (c *DbClient) OpenConnection() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.ConnInfo.User, c.ConnInfo.Password, c.ConnInfo.HostAddr, c.ConnInfo.Port, c.ConnInfo.DbName))

	if err != nil {
		c.Log.Fatal("cannot open connection to db: %v", err)
	}
	c.db = db
	defer db.Close()
	if err := db.Ping(); err != nil {
		c.Log.Fatal("cannot establish connection: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	stats := c.db.Stats()
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "===== STATS =====")
	fmt.Fprintln(w, "Max Open Connections:\t", stats.MaxOpenConnections)
	fmt.Fprintln(w, "Open Connection:\t", stats.OpenConnections)
	fmt.Fprintln(w, "Idle:\t", stats.Idle)
	fmt.Fprintln(w, "In Use:\t", stats.InUse)
	w.Flush()

	fmt.Print("> ")
	for scanner.Scan() {
		if scanner.Err() != nil {
			c.Log.Error("cannot scan input line: %v", scanner.Err())
		}
		line := strings.ToLower(strings.Trim(scanner.Text(), " "))
		if line == "q" {
			break
		}

		cmds := strings.Split(line, " ")
		switch cmds[0] {
		case "info":
			fmt.Printf("\nFETCHING INFORMATION FOR TABLE: %s\n", cmds[1])
			fields := c.tableInfo(cmds[1])
			tWriter := tablewriter.NewWriter(os.Stdout)
			tWriter.SetHeader([]string{"Field", "Type", "Null", "Key", "Default", "Extra"})

			for _, f := range fields {
				tWriter.Append([]string{f.Field, f.Type, f.Null, f.Key, f.Default, f.Extra})
			}
			tWriter.Render()
		case "add":
			c.addData(cmds[1])
		case "display":
			c.displayTable(cmds[1])
		case "q":
			return
		case "exit":
			return
		default:
			fmt.Println("INVALID SYNTAX")
		}
		fmt.Print("> ")
	}
}
