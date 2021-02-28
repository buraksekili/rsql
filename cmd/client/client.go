package client

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/buraksekili/rsql/data"

	"github.com/buraksekili/selog"
	"github.com/olekukonko/tablewriter"
)

type DbClient struct {
	ConnInfo *data.ConnInfo
	Log      *selog.Selog
	db       *sql.DB
}

func NewDbClient(l *selog.Selog) *DbClient {
	return &DbClient{&data.ConnInfo{}, l, nil}
}
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
			c.tableInfo(cmds[1])
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

func (c *DbClient) tableInfo(table string) {
	fmt.Printf("\nFETCHING INFORMATION FOR TABLE: %s\n", table)
	if c.db == nil {
		c.Log.Error("cannot fetch table: c.db == nil\n")
		return
	}

	rows, err := c.db.Query(fmt.Sprintf("show columns from %s", table))
	if err != nil {
		c.Log.Error("cannot fetch table info: %v\n", err)
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	colVals := make([]interface{}, len(cols))
	var fields []data.TableField
	var field data.TableField

	for rows.Next() {
		// https://www.farrellit.net/2018/08/12/golang-sql-unknown-rows.html
		colsAsSoc := make(map[string]interface{}, len(cols))

		for i, _ := range colVals {
			colVals[i] = new(interface{})
		}
		if err := rows.Scan(colVals...); err != nil {
			c.Log.Error("cannot get columns: %v", err)
			break
		}

		for i, col := range cols {
			colsAsSoc[col] = *colVals[i].(*interface{})
			switch col {
			case "Field":
				field.Field = fmt.Sprintf("%s", colsAsSoc[col])
			case "Type":
				field.Type = fmt.Sprintf("%s", colsAsSoc[col])
			case "Null":
				field.Null = fmt.Sprintf("%s", colsAsSoc[col])
			case "Default":
				if colsAsSoc[col] == nil {
					field.Default = ""
				} else {
					field.Default = fmt.Sprintf("%s", colsAsSoc[col])
				}
			case "Extra":
				field.Extra = fmt.Sprintf("%s", colsAsSoc[col])
				fields = append(fields, field)
			}
		}
	}

	tWriter := tablewriter.NewWriter(os.Stdout)
	tWriter.SetHeader([]string{"Field", "Type", "Null", "Key", "Default", "Extra"})

	for _, f := range fields {
		tWriter.Append([]string{f.Field, f.Type, f.Null, f.Key, f.Default, f.Extra})
	}
	tWriter.Render()
}
