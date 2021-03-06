package client

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/buraksekili/rsql/data"
	_ "github.com/go-sql-driver/mysql"
)

func TestDbClient_ShowTables(t *testing.T) {
	c := NewDBClient()
	if err := MockDB(c); err != nil {
		t.Fatal("cannot mock db: ", err.Error())
	}

	tables := c.showTables()
	for _, row := range tables {
		if strings.Trim(row, " ") != "db_test" {
			t.Fatalf("got %s; want %s", row, "db_test")
		}
	}
}

// NewDBClient returns DbClient by using temp logger
func NewDBClient() *DbClient {
	return NewDbClient(nil)
}

// MockDB mocks DB with temporary database
func MockDB(c *DbClient) error {
	// you can create new MySQL container to test it with temporary credentials.
	conn := data.ConnInfo{
		User:     "test",
		Password: "yourtestpassword",
		HostAddr: "127.0.0.1",
		Port:     "8080",
		DbName:   "posts_test",
	}
	c.ConnInfo = &conn

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.ConnInfo.User, c.ConnInfo.Password, c.ConnInfo.HostAddr, c.ConnInfo.Port, c.ConnInfo.DbName))
	if err != nil {
		return fmt.Errorf("cannot establish mysql connection: %s", err.Error())
	}
	c.db = db
	return nil
}

