package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// OpenConnection opens connection, pings the DB.
// In case of any error while pinging, it returns error.
//
// If there is no error after pinging the DB, this function takes
// input(cmds) to execute on the DB specified by the user.
// To see available cmds;
// 		> help
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

	c.displayDBStats(os.Stdout)
	printHelp()
	fmt.Print("rsql> ")

	for scanner.Scan() {
		if scanner.Err() != nil {
			c.Log.Error("cannot scan input line: %v", scanner.Err())
		}

		line := strings.Trim(scanner.Text(), " ")
		cmds := strings.Split(line, " ")

		if len(cmds) > 2 {
			fmt.Printf("Invalid number of commands\n")
		} else if len(cmds) == 1 {
			switch cmds[0] {
			case "tables":
				c.showTables(os.Stdout)
			case "stats":
				c.displayDBStats(os.Stdout)
			case "help":
				printHelp()
			case "q":
				return nil
			case "exit":
				return nil
			default:
				fmt.Println("INVALID SYNTAX")
			}
		} else {
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
				c.displayTable(cmds[1], os.Stdout)
			case "source":
				if err := c.source(cmds[1]); err != nil {
					fmt.Println("error is ", err)
				}
			default:
				fmt.Println("INVALID SYNTAX")
			}
		}
		fmt.Print("rsql> ")
	}
	return nil
}

func printHelp() {
	fmt.Println("COMMANDS\n" +
		"\tadd <TABLE>\t: adds data to <TABLE>\n" +
		"\tinfo <TABLE>\t: displays the column informations of the <TABLE>\n" +
		"\tdisplay <TABLE>\t: displays the data of the <TABLE>\n" +
		"\ttables\t\t: displays available tables under the <DB> specified by user\n" +
		"\thelp\t\t: displays this message\n" +
		"\tq, exit \t: exits the program")
}
