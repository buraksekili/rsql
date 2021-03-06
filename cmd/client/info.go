package client

import (
	"fmt"

	"github.com/buraksekili/rsql/data"
)

// tableInfo returns fields of given table.
func (c *DbClient) tableInfo(table string) []data.TableField {
	if c.db == nil {
		c.Log.Error("cannot fetch table: c.db == nil\n")
		return nil
	}

	rows, err := c.db.Query(fmt.Sprintf("show columns from %s", table))
	if err != nil {
		c.Log.Error("cannot fetch table info: %v\n", err)
		return nil
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		c.Log.Error("cannot fetch columns: %v\n", err)
	}

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
	return fields
}
