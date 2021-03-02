package client

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// OpenConnection opens connection, pings the DB.
// In case of any error while pinging, the app prints fatal
// error and exits.
//
// If there is no error after pinging the DB, this function takes
// input to execute on the DB specified by the user.
func (c *DbClient) OpenConnection() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.ConnInfo.User, c.ConnInfo.Password, c.ConnInfo.HostAddr, c.ConnInfo.Port, c.ConnInfo.DbName))

	if err != nil {
		return err
	}
	c.db = db
	defer db.Close()
	if err := db.Ping(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)

	c.displayDBStats()
	fmt.Print("> ")

	for scanner.Scan() {
		if scanner.Err() != nil {
			c.Log.Error("cannot scan input line: %v", scanner.Err())
		}

		line := strings.ToLower(strings.Trim(scanner.Text(), " "))
		cmds := strings.Split(line, " ")
		switch cmds[0] {
		case "info":
			fields := c.tableInfo(cmds[1])
			if len(fields) != 0 {
				fmt.Printf("\nFETCHING INFORMATION FOR TABLE: %s\n", cmds[1])
				tWriter := tablewriter.NewWriter(os.Stdout)
				tWriter.SetHeader([]string{"Field", "Type", "Null", "Key", "Default", "Extra"})

				for _, f := range fields {
					tWriter.Append([]string{f.Field, f.Type, f.Null, f.Key, f.Default, f.Extra})
				}
				tWriter.Render()
			}
		case "add":
			c.addData(cmds[1])
		case "display":
			c.displayTable(cmds[1])
		case "tables":
			c.showTables()
		case "stats":
			c.displayDBStats()
		case "q":
			return nil
		case "exit":
			return nil
		default:
			fmt.Println("INVALID SYNTAX")
		}
		fmt.Print("> ")
	}
	return nil
}
