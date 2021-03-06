package client

import (
	"strings"
	"testing"
)

func TestDbClient_TableInfo(t *testing.T) {
	c := NewDBClient()
	if err := MockDB(c); err != nil {
		t.Fatal("cannot mock db: ", err.Error())
	}

	fields := c.tableInfo("db_test")
	for _, f := range fields {
		if strings.Trim(f.Field, " ") == "id" && strings.Trim(f.Type, " ") != "int" {
			t.Fatalf("id must be int; want int, got %s", f.Type)
		}

		if strings.Trim(f.Field, " ") == "body" && strings.Trim(f.Type, " ") != "mediumtext" {
			t.Fatalf("body must be mediumtext; want mediumtext, got %s", f.Type)
		}
	}
}
