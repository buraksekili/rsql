package client

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// showTables displays the available tables of the db
// specified by user.
func (c *DbClient) showTables() []string {
	rows, err := c.db.Query("show tables")
	if err != nil {
		c.Log.Error("cannot fetch tables: %v", err)
		return nil
	}

	cols, err := rows.Columns()
	if err != nil {
		c.Log.Error("cannot fetch columns: %v", err)
		return nil
	}
	defer rows.Close()

	tables := []string{}
	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("couldn't scan ", err.Error())
		}

		for _, v := range vals {
			s := v.(*sql.RawBytes)
			tables = append(tables, string(*s))
			fmt.Printf("value is %v\n", string(*s))
		}
	}

	tWriter := tablewriter.NewWriter(os.Stdout)
	for _, row := range tables {
		tWriter.Append([]string{row})
	}
	tWriter.Render()

	return tables
}
