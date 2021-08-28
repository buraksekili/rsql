package commands

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

// displayTable writes the content of the given table into io.Writer.
// Returns content of the table or error.
func (c *DbClient) displayTable(table string, w io.Writer) ([][]string, error) {
	q := fmt.Sprintf("select * from %s;", table)
	rows, err := c.db.Query(q)
	if err != nil {
		c.Log.Error("cannot display table %s: %v\n", table, err)
		return nil, fmt.Errorf("cannot display table %s: %v\n", table, err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		c.Log.Error("cannot fetch columns: %v\n", err)
	}

	res := [][]string{}
	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("couldn't scan ", err.Error())
		}

		row := []string{}
		for _, v := range vals {
			s := v.(*sql.RawBytes)
			row = append(row, string(*s))
		}
		res = append(res, row)
	}

	tWriter := tablewriter.NewWriter(w)
	tWriter.SetHeader(cols)

	for _, row := range res {
		tWriter.Append(row)
	}
	tWriter.Render()

	return res, nil
}
