package client

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func (c *DbClient) source(dbFile string) error {
	if !strings.HasSuffix(dbFile, ".sql") {
		return fmt.Errorf("you need .sql file, want=*.sql, got=%s", dbFile)
	}

	data, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return err
	}
	_, err = c.db.Query(string(data))
	if err != nil {
		return err
	}
	fmt.Printf("%s successfully executed", dbFile)

	return nil
}
