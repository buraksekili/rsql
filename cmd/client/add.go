package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (c *DbClient) addData(table string) {
	fields := c.tableInfo(table)
	if len(fields) == 0 {
		c.Log.Error("cannot fetch table information for table %s\n", table)
		return
	}

	req := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for _, f := range fields {
		// if type is date, add default date
		if f.Extra != "auto_increment" {
			fmt.Printf("%s (%s): ", f.Field, f.Type)
			if strings.ToLower(f.Type) == "date" {
				t := time.Now().Format(time.RFC3339)
				fmt.Printf("default: %s", t)
			}
			if scanner.Scan() {
				if scanner.Err() != nil {
					c.Log.Error("cannot scan input line for adding information %s: %v", f.Field, scanner.Err())
					break
				}
				if strings.ToLower(f.Type) == "date" && strings.Trim(scanner.Text(), " ") == "" {
					fmt.Println("using default time")
					req[f.Field] = time.Now().Format(time.RFC3339)
				} else {
					req[f.Field] = scanner.Text()
				}
			}
		}
	}

	q := fmt.Sprintf("insert into %s (", table)

	var keys []string
	for k, _ := range req {
		keys = append(keys, k)
		q += k + ","
	}

	// remove ending comma
	q = q[:len(q)-1] + ") VALUES ("
	for _, k := range keys {
		q += fmt.Sprintf("'%s',", req[k])
	}
	q = q[:len(q)-1] + ")"

	fmt.Println("query is q ", q)

	_, err := c.db.Exec(q)
	if err != nil {
		c.Log.Error("cannot insert into table `%s`: %v", table, err)
	}
	fmt.Println("successfully inserted")
}
