package commands

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

// showTables displays the available tables of the db
// specified by user.
func (c *DbClient) showTables(w io.Writer) []string {
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
		}
	}

	tWriter := tablewriter.NewWriter(w)
	for _, row := range tables {
		tWriter.Append([]string{row})
	}
	tWriter.Render()

	return tables
}
